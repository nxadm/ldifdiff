// Package ldiffdiff finds and outputs the difference between two LDIF files.
package ldifdiff

import (
	"bytes"
	"sort"
	"strings"
	"sync"
)

const Version = "v0.1.0"
const Author = "Claudio Ramirez <pub.claudio@gmail.com>"
const Repo = "https://github.com/nxadm/ldifdiff"

type fn func(string, []string) (Entries, error)

var skipDnForDelete map[string]bool

/* Public functions */

// Diff compares two LDIF strings (sourceStr and targetStr) and outputs the
// differences as a LDIF string. An array of attributes can be supplied. These
// attributes will be ignored when comparing the LDIF strings.
// The output is a string, a valid LDIF, and can be added to the "target"
// database (the one that created targetStr) in order to make it
// equal to the *source* database (the one that created sourceStr). In case of
// failure, an error is provided.
func Diff(sourceStr, targetStr string, ignoreAttr []string) (string, error) {
	return genericDiff(sourceStr, targetStr, ignoreAttr, convertLdifStr, nil)
}

// DiffFromFiles compares two LDIF files (sourceFile and targetFile) and
// outputs the differences as a LDIF string. An array of attributes can be
// supplied. These attributes will be ignored when comparing the LDIF strings.
// The output is a string, a valid LDIF, and can be added to the "target"
// database (the one that created targetFile) in order to make it equal to the
// *source* database (the one that created sourceFile). In case of failure, an
// error is provided.
func DiffFromFiles(sourceFile, targetFile string, ignoreAttr []string) (string, error) {
	return genericDiff(sourceFile, targetFile, ignoreAttr, importLdifFile, nil)
}

// ListDiffDn compares two LDIF strings (sourceStr and targetStr) and outputs
// the differences as a list of affected DNs (Dintinguished Names). An array of
// attributes can be supplied. These attributes will be ignored when comparing
// the LDIF strings.
// The output is a string slice. In case of failure, an error is provided.
func ListDiffDn(sourceStr, targetStr string, ignoreAttr []string) ([]string, error) {
	dnList := []string{}
	_, err := genericDiff(sourceStr, targetStr, ignoreAttr, convertLdifStr, &dnList)
	return dnList, err
}

// ListDiffDnFromFiles compares two LDIF files (sourceFile and targetFileStr)
// and outputs the differences as a list of affected DNs (Dintinguished Names).
// An array of attributes can be supplied. These attributes will be ignored
// when comparing the LDIF strings.
// The output is a string slice. In case of failure, an error is provided.
func ListDiffDnFromFiles(sourceFile, targetFile string, ignoreAttr []string) ([]string, error) {
	dnList := []string{}
	_, err := genericDiff(sourceFile, targetFile, ignoreAttr, importLdifFile, &dnList)
	return dnList, err
}

/* Package private functions */

func arraysEqual(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Ordering Logic:
// Add: entries from source sorted S -> L. Otherwise is invalid.
// Remove: entries from target sorted L -> S. Otherwise is invalid.
// Modify:
// - Keep S ->  L ordering
// - If only 1 instance of attribute with different value on source and target:
// update. This way we don't break the applicable LDAP schema.
// - extra attribute on source: add
// - extra attribute on target: delete

func compare(source, target *Entries, dnList *[]string) (string, error) {
	var buffer bytes.Buffer
	var err error
	queue := make(chan ActionEntry, 10)
	var wg sync.WaitGroup

	// Find the order in which operation must happen
	orderedSourceShortToLong := sortDnByDepth(source, false)
	orderedTargetLongToShort := sortDnByDepth(target, true)

	// Write the file concurrently
	wg.Add(1) // 1 writer
	go writeLdif(queue, &buffer, &wg, &err)

	// Dn only on source + removal of identical entries
	skipDnForDelete = make(map[string]bool) // Keep track of dn to skip at Deletion
	sendForAddition(&orderedSourceShortToLong, source, target, queue, dnList)

	// Dn only on target
	sendForDeletion(&orderedTargetLongToShort, source, target, queue, dnList)

	// Dn on source and target
	sendForModification(&orderedSourceShortToLong, source, target, queue, dnList)

	// Done sending work
	close(queue)

	// Free some memory
	*source = Entries{}
	*target = Entries{}

	// Wait for the creation of the LDIF
	wg.Wait()

	// Return the results
	return buffer.String(), err
}

//func elementInArray(a string, array []string) bool {
//	for _, b := range array {
//		if b == a {
//			return true
//		}
//	}
//	return false
//}

func genericDiff(sourceParam, targetParam string, ignoreAttr []string, fn fn, dnList *[]string) (string, error) {
	// Read the files in memory as a Map with sorted attributes
	var source, target Entries
	var sourceErr, targetErr error
	var wg sync.WaitGroup
	wg.Add(2)
	go func(entries *Entries, wg *sync.WaitGroup, err *error) {
		result, e := fn(sourceParam, ignoreAttr)
		*entries = result
		*err = e
		wg.Done()
	}(&source, &wg, &sourceErr)
	go func(entries *Entries, wg *sync.WaitGroup, err *error) {
		result, e := fn(targetParam, ignoreAttr)
		*entries = result
		*err = e
		wg.Done()
	}(&target, &wg, &targetErr)
	wg.Wait()

	if sourceErr != nil {
		return "", sourceErr
	}
	if targetErr != nil {
		return "", targetErr
	}

	// Compare the files
	return compare(&source, &target, dnList)
}

func sendForAddition(
	orderedSourceShortToLong *[]string,
	source, target *Entries,
	queue chan<- ActionEntry,
	dnList *[]string) {
	for _, dn := range *orderedSourceShortToLong {
		// Ignore equal entries
		if arraysEqual((*source)[dn], (*target)[dn]) {
			delete(*source, dn)
			delete(*target, dn)
			continue
		}
		// Mark entries for addition if only on source
		if _, ok := (*target)[dn]; !ok {
			if dnList == nil {
				subActionAttr := make(map[SubAction][]string)
				subActionAttr[None] = (*source)[dn]
				actionEntry :=
					ActionEntry{
						Dn:             dn,
						Action:         Add,
						SubActionAttrs: []SubActionAttr{subActionAttr},
					}
				queue <- actionEntry
			} else {
				// Always add (attributes not relevant)
				*dnList = append(*dnList, dn)
			}
			delete(*source, dn)
		}
		// Implict else:
		// It exists on target and it's not equal, so it's a modifyStr
		skipDnForDelete[dn] = true
	}
}

func sendForDeletion(
	orderedTargetLongToShort *[]string,
	source, target *Entries,
	queue chan<- ActionEntry,
	dnList *[]string) {
	for _, dn := range *orderedTargetLongToShort {
		if skipDnForDelete[dn] { // We know it's not a delete operation
			continue
		}
		if _, ok := (*target)[dn]; ok { // It has not been deleted above
			if _, ok := (*source)[dn]; !ok { // does not exists on source
				if dnList == nil {
					subActionAttr := make(map[SubAction][]string)
					subActionAttr[None] = nil
					actionEntry :=
						ActionEntry{
							Dn:             dn,
							Action:         Delete,
							SubActionAttrs: []SubActionAttr{subActionAttr},
						}
					queue <- actionEntry
				} else {
					// Always remove (attributes are not relevant)
					*dnList = append(*dnList, dn)
				}
				delete(*target, dn)
			}
			// Implict else:
			// It exists on source and it's not equal (tested on sendForAddition),
			// so it's a modifyStr
		}
	}
	// Free some memory
	skipDnForDelete = nil
}

func sendForModification(
	orderedSourceShortToLong *[]string, source,
	target *Entries,
	queue chan<- ActionEntry,
	dnList *[]string) {
	for _, dn := range *orderedSourceShortToLong {
		// DN is present on source and target:
		// sendForAdd/Remove clean up source and target
		_, okSource := (*source)[dn]
		_, okTarget := (*target)[dn]
		if okSource && okTarget { // it hasn't been deleted
			if dnList == nil {

				// Store the attributes to be added, deleted or replaced
				attrToModifyAdd := []string{}
				attrToModifyDelete := []string{}
				attrToModifyReplace := []string{}
				// Put the attributes in a map for easy lookup
				sourceAttr := make(map[string]bool)
				targetAttr := make(map[string]bool)
				for _, attr := range (*source)[dn] {
					sourceAttr[attr] = true
				}
				for _, attr := range (*target)[dn] {
					targetAttr[attr] = true
				}

				// Compare attribute values starting from the source
				for _, attr := range (*source)[dn] { // Keep the order of the attributes
					// Attribute is not equal on both sides
					if _, ok := targetAttr[attr]; !ok {
						// Is the attribute name (not value) unique?
						switch uniqueAttrName(attr, sourceAttr, targetAttr) {
						case true: // This is a Modify-Replace operation
							attrToModifyReplace = append(attrToModifyReplace, attr)
						case false: // This is just a Modify Add (only on source).
							attrToModifyAdd = append(attrToModifyAdd, attr)
						}
					}
				}

				// Compare attribute values starting from the target.
				for _, attr := range (*target)[dn] { // Keep the order of the attributes
					// Looking for unique attributes
					if !uniqueAttrName(attr, sourceAttr, targetAttr) {
						if _, ok := sourceAttr[attr]; !ok {
							attrToModifyDelete = append(attrToModifyDelete, attr)
						}
					}
				}

				// Send it
				actionEntry := ActionEntry{Dn: dn, Action: Modify}
				subActionAttrArray := []SubActionAttr{}
				switch {
				case len(attrToModifyAdd) > 0:
					subActionAttrArray = append(subActionAttrArray, SubActionAttr{ModifyAdd: attrToModifyAdd})
					fallthrough
				case len(attrToModifyDelete) > 0:
					subActionAttrArray = append(subActionAttrArray, SubActionAttr{ModifyDelete: attrToModifyDelete})
					fallthrough
				case len(attrToModifyReplace) > 0:
					subActionAttrArray = append(subActionAttrArray, SubActionAttr{ModifyReplace: attrToModifyReplace})
				}

				actionEntry.SubActionAttrs = subActionAttrArray
				queue <- actionEntry
			} else {
				// There must be something left to modify
				//if len((*source)[dn]) > 0 || len((*target)[dn]) > 0 {
				*dnList = append(*dnList, dn)
				//}
			}
			// Clean it up
			delete(*source, dn)
			delete(*target, dn)
		}
	}
}

func sortDnByDepth(entries *Entries, longToShort bool) []string {
	var sorted []string

	dns := []string{}
	for dn := range *entries {
		dns = append(dns, dn)
	}

	splitByDc := make(map[string][]string)
	longestDn := 0
	// Split the components of the dn and remember the longest size
	for _, dn := range dns {
		parts := strings.Split(dn, ",")
		if len(parts) > longestDn {
			longestDn = len(parts)
		}
		splitByDc[dn] = parts
	}

	// Get the direction of the loop
	componentSizes := []int{}
	if longToShort {
		for i := longestDn; i > 0; i-- {
			componentSizes = append(componentSizes, i)
		}
	} else {
		for i := 1; i <= longestDn; i++ {
			componentSizes = append(componentSizes, i)
		}
	}

	// Sort by size and alpahbetically within size
	for _, size := range componentSizes {
		sameSize := []string{}
		for dn, components := range splitByDc {
			if len(components) == size {
				sameSize = append(sameSize, dn)
			}
		}
		sort.Strings(sameSize)
		sorted = append(sorted, sameSize...)
	}

	return sorted
}

func uniqueAttrName(attr string, sourceAttr, targetAttr map[string]bool) bool {

	// Get the attribute name
	parts := strings.Split(attr, ":")
	attrName := parts[0]
	//var base64 bool
	//if parts[1] == "" {
	//	base64 = true
	//}
	sourceCounter := 0
	for attr := range sourceAttr {
		if strings.HasPrefix(attr, attrName+":") {
			sourceCounter++
		}
	}
	targetCounter := 0
	for attr := range targetAttr {
		if strings.HasPrefix(attr, attrName+":") {
			targetCounter++
		}
	}

	return sourceCounter == 1 && targetCounter == 1
}
