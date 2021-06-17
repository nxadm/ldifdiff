package ldifdiff

//
//import (
//	"errors"
//	"strings"
//	"sync"
//)
//
//func importRecords(inputType inputType, source string, ignoreAttr []string) (Entries, error) {
//	var (
//		//entries  Entries
//		readErr  error
//		//parseErr error
//		wg       sync.WaitGroup
//	)
//
//	queue := make(chan []string, 10)
//
//	// Read and Parse the file concurrently
//	//wg.Add(2) // 1 reader + 1 parser
//	wg.Add(1)
//
//	switch inputType {
//	case inputStr: // it's a ldifStr
//		go readStr(source, ignoreAttr, queue, &wg, &readErr)
//	default: // it's a file
//		//go readFile(source, ignoreAttr, queue, &wg, &readErr)
//	}
//
//	//go parse(entries, queue, &wg, &parseErr)
//
//	wg.Wait()
//
//	//
//	//// Return values
//	//switch {
//	//case readErr != nil:
//	//	return entries, ErrReadLDIF(readErr)
//	//case parseErr != nil:
//	//	return entries, ErrParseLDIF(parseErr)
//	//default:
//	//	return entries, nil
//	//}
//	return Entries{}, nil
//}
//////
///////* Package only functions */
//////
//////func convertLdifStr(ldifStr string, ignoreAttr []string) (entries, error) {
//////	return importRecords(ldifStr, "", ignoreAttr)
//////}
//////
//////func importLdifFile(file string, ignoreAttr []string) (entries, error) {
//////	entries, err := importRecords("", file, ignoreAttr)
//////	if err != nil {
//////		err = errors.New(err.Error() + " [" + file + "]")
//////	}
//////	return entries, err
//////}
//////
///////* Internal functions */
//////
//func addLineToRecord(line *string, record *[]string, ignoreAttr []string, prevAttrSkipped *bool) error {
//	var err error
//
//	// Append continuation lines to previous line
//	if strings.HasPrefix(*line, " ") {
//		if *prevAttrSkipped { // but not lines from a skipped attribute
//			return nil
//		}
//		switch len(*record) > 0 {
//		case true:
//			prevIdx := len(*record) - 1
//			prevLine := strings.TrimSuffix((*record)[prevIdx], "\n") +
//				strings.TrimPrefix(*line, " ")
//			(*record)[prevIdx] = prevLine
//			return nil
//		case false:
//			err = errors.New("Invalid modifyStr line continuation: \"" + *line + "\"")
//			return err
//		}
//	}
//
//	// Regular line
//	if len(*line) != 0 {
//		for _, attrName := range ignoreAttr {
//			if strings.HasPrefix(*line, attrName+":") {
//				*prevAttrSkipped = true
//				return nil
//			}
//		}
//		*record = append(*record, *line)
//		*prevAttrSkipped = false
//	}
//
//	return nil
//}
////
////func readFile(file string, ignoreAttr []string, queue chan<- []string, wg *sync.WaitGroup, err *error) {
////	defer wg.Done()
////	defer close(queue)
////	fh, osErr := os.Open(file)
////	if osErr != nil {
////		*err = osErr
////		return
////	}
////	defer fh.Close()
////
////	record := []string{}
////	scanner := bufio.NewScanner(fh)
////	var prevAttrSkipped bool // use to skip continuation lines of skipped attr
////	firstLine := true
////	for scanner.Scan() {
////
////		line := scanner.Text()
////
////		// Skip comments
////		if strings.HasPrefix(line, "#") {
////			continue
////		}
////
////		// Check if first line is a "version: *" line and skip it
////		if firstLine {
////			if strings.HasPrefix(line, "version: ") {
////				firstLine = false
////				continue
////			} else if line == "" {
////				continue
////			}
////			firstLine = false
////		}
////
////		// Import lines as records
////		*err = addLineToRecord(&line, &record, ignoreAttr, &prevAttrSkipped)
////		if *err != nil {
////			return
////		}
////
////		//  Dispatch the record to buffer & reset record
////		if len(line) == 0 && len(record) != 0 {
////			queue <- record
////			record = []string{}
////		}
////	}
////
////	// Last record may be a leftover (no empty line)
////	if len(record) != 0 {
////		queue <- record
////	}
////
////	if *err == nil {
////		*err = scanner.Err()
////	}
////}
////
//func readStr(source string, ignoreAttr []string, queue chan<- []string, wg *sync.WaitGroup, err *error) {
//	defer wg.Done()
//	defer close(queue)
//
//	// Records are separated by an empty line
//	for idx, recordStr := range strings.Split(source, "\n\n") {
//		var (
//			prevAttrSkipped bool
//			record []string
//		)
//
//		for _, line := range strings.Split(recordStr, "\n") {
//			if strings.HasPrefix(line, "#") { // Skip comments
//				continue
//			}
//
//			if idx == 0 { // First record only, skip version
//				if strings.HasPrefix(line, "version: ") {
//					continue
//				}
//			}
//
//			*err = addLineToRecord(&line, &record, ignoreAttr, &prevAttrSkipped)
//			if *err != nil {
//				return
//			}
//
//			//  Dispatch the record to buffer & reset record
//			if len(line) == 0 && len(record) != 0 {
//				queue <- record
//				record = []string{}
//			}
//		}
//		// Last record may be a leftover (no empty line)
//		if len(record) != 0 {
//			queue <- record
//		}
//	}
//}
////
////func parse(entries Entries, queue <-chan []string, wg *sync.WaitGroup, err *error) {
////	defer wg.Done()
////	for record := range queue {
////		dn := record[0] // Find dn, should be the first line
////		if !strings.HasPrefix(dn, "dn:") {
////			*err = errors.New("No dn could be retrieved")
////			continue
////		}
////
////		// Sort the entries
////		attr := record[1:]
////		sort.Strings(attr)
////		entries[dn] = attr
////	}
////}
