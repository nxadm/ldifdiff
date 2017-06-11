package ldifdiff

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strings"
	"sync"
)

/* Package only functions */

func convertLdifStr(ldifStr string, ignoreAttr []string) (Entries, error) {
	return importRecords(ldifStr, "", ignoreAttr)
}

func importLdifFile(file string, ignoreAttr []string) (Entries, error) {
	entries, err := importRecords("", file, ignoreAttr)
	if err != nil {
		err = errors.New(err.Error() + " [" + file + "]")
	}
	return entries, err
}

/* Internal functions */

func addLineToRecord(line *string, record *[]string, ignoreAttr []string, prevAttrSkipped *bool) error {
	var err error

	// Skip comments
	if strings.HasPrefix(*line, "#") {
		return nil
	}

	// Append continuation lines to previous line
	if strings.HasPrefix(*line, " ") {
		if *prevAttrSkipped { // but not lines from a skipped attribute
			return nil
		}
		switch len(*record) > 0 {
		case true:
			prevIdx := len(*record) - 1
			prevLine := strings.TrimSuffix((*record)[prevIdx], "\n") +
				strings.TrimPrefix(*line, " ")
			(*record)[prevIdx] = prevLine
			return nil
		case false:
			err = errors.New("Invalid modifyStr line continuation: \"" + *line + "\"")
			return err
		}
	}

	// Regular line
	if len(*line) != 0 {
		for _, attrName := range ignoreAttr {
			if strings.HasPrefix(*line, attrName+":") {
				*prevAttrSkipped = true
				return nil
			}
		}
		*record = append(*record, *line)
		*prevAttrSkipped = false
	}

	return nil
}

func importRecords(ldifStr, file string, ignoreAttr []string) (Entries, error) {
	var readErr, parseErr error
	queue := make(chan []string, 10)
	entries := make(map[string][]string)

	// Read and Parse the file concurrently
	var wg sync.WaitGroup
	wg.Add(2) // 1 reader + 1 parser
	switch file {
	case "": // it's a ldifStr
		go readStr(ldifStr, ignoreAttr, queue, &wg, &readErr)
	default: // it's a file
		go readFile(file, ignoreAttr, queue, &wg, &readErr)
	}
	go parse(entries, queue, &wg, &parseErr)
	wg.Wait()

	// Return values
	switch {
	case readErr != nil:
		return entries, readErr
	case parseErr != nil:
		return entries, parseErr
	default:
		return entries, nil
	}
}

func readFile(file string, ignoreAttr []string, queue chan<- []string, wg *sync.WaitGroup, err *error) {
	defer wg.Done()
	defer close(queue)
	fh, osErr := os.Open(file)
	if osErr != nil {
		*err = osErr
		return
	}
	defer fh.Close()

	record := []string{}
	scanner := bufio.NewScanner(fh)
	var prevAttrSkipped bool // use to skip continuation lines of skipped attr
	for scanner.Scan() {
		line := scanner.Text()
		*err = addLineToRecord(&line, &record, ignoreAttr, &prevAttrSkipped)
		if *err != nil {
			return
		}

		//  Dispatch the record to buffer & reset record
		if len(line) == 0 && len(record) != 0 {
			queue <- record
			record = []string{}
		}
	}

	// Last record may be a leftover (no empty line)
	if len(record) != 0 {
		queue <- record
	}

	if *err == nil {
		*err = scanner.Err()
	}
}

func readStr(ldifStr string, ignoreAttr []string, queue chan<- []string, wg *sync.WaitGroup, err *error) {
	defer wg.Done()
	defer close(queue)
	for _, recordStr := range strings.Split(ldifStr, "\n\n") {
		var prevAttrSkipped bool
		record := []string{}
		for _, line := range strings.Split(recordStr, "\n") {
			*err = addLineToRecord(&line, &record, ignoreAttr, &prevAttrSkipped)
			if *err != nil {
				return
			}

			//  Dispatch the record to buffer & reset record
			if len(line) == 0 && len(record) != 0 {
				queue <- record
				record = []string{}
			}
		}
		// Last record may be a leftover (no empty line)
		if len(record) != 0 {
			queue <- record
		}
	}
}

func parse(entries Entries, queue <-chan []string, wg *sync.WaitGroup, err *error) {
	defer wg.Done()
	for record := range queue {
		dn := record[0] // Find dn, should be the first line
		if !strings.HasPrefix(dn, "dn:") {
			*err = errors.New("No dn could be retrieved")
			continue
		}

		// Sort the entries
		attr := record[1:]
		sort.Strings(attr)
		entries[dn] = attr
	}
}
