// Package ldifdiff is a fast library that outputs the difference
// between two LDIF files as a valid and importable LDIF (e.g.
// by your LDAP server).
package ldifdiff

import "sync"

//Info base 64
//It is legal to base64-encode any value when representing it in LDIF,
//but any value that meets any of the following conditions is required
//to be base64 encoded for use in LDIF:
//
//Any value that begins with a space.
//Any value that begins with a colon.
//Any value that begins with a less-than character.
//Any value that includes a carriage-return or line-break character.
//Any value that includes the ASCII null character.
//Any value that contains non-ASCII data.
//
//Although it is not absolutely required, it is strongly recommended that
//values that end with a space should also be base64 encoded. It is also
//strongly recommended that values containing ASCII control characters
//(e.g., tabs, backspace, etc.) be base64 encoded.

/* Public functions */

//// Diff is a higher level function that compares a LDIF string with all the
//// supplied targets. A list is returned where each element is a ModifyLDIF
//// for the specific target. If attributes are supplied, they will be ignored
//// in the comparison. In case of failure, an error is provided.
//func Diff(source string, targets []string, attr []string) ([]ModifyLDIF, error) {
//	return nil, nil
//}
//
//// DiffFromFiles compares a LDIF file with all the supplied targets. A list is
//// returned where each element is a ModifyLDIF for the specific target. If
//// attributes are supplied, they will be ignored in the comparison. In case
//// of failure, an error is provided.
//func DiffFromFiles(source string, targets []string, attr []string) ([]ModifyLDIF, error) {
//	return nil, nil
//}
//
//// ListDiffDn compares a LDIF string with all the supplied targets. It outputs the
//// differences as a list of affected DNs (Dintinguished Names). If attributes
//// are supplied, they will be ignored in the comparison. In case of failure, an
//// error is provided.
//func ListDiffDn(source string, targets []string, attr []string) ([]DN, error) {
//	return nil, nil
//}
//
//// ListDiffDnFromFiles compares a LDIF file with all the supplied targets. It
//// outputs the differences as a list of affected DNs (Dintinguished Names). If
//// attributes are supplied, they will be ignored in the comparison. In case of
//// failure, an error is provided.
//func ListDiffDnFromFiles(source string, targets []string, attr []string) ([]DN, error) {
//	return nil, nil
//}

// ImportLDIF imports the contents of an LDIF file in a structured Entries that can be
// compared.
func ImportLDIF(in inputType, source string, ignoreAttr []string) (DNInfo, error) {
	//input, err := readLDIF(in, source)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}

// CompareEntries compares two Entries and returns the results as a DiffResult
// and an error if applicable. If attributes are supplied, they will be ignored
// in the comparison. In case of failure, an error is provided.
func CompareEntries(source, target DNInfo, attr []string) (DiffResult, error) {
	return nil, nil
}

// DiffToLDIF is a low level functions that converts a channel of DNAction
// into a channel of strings that create an LDIF, usable by ldapmodify.
// This allows to use (e.g. printing) the results as they come in.
func DiffToLDIF(input <-chan DNAction) chan string {
	wg := &sync.WaitGroup{}
	output := make(chan string, 10)

	wg.Add(1)

	go createLDIF(input, output, wg)

	return output
}

//
//func diff(inputType inputType, source, targets, ignoreAttr []string) ([]string, error) {
//	//var (
//	//	sourceEntries, targetEntries entries
//	//	sourceErr, targetErr         error
//	//	wg                           sync.WaitGroup
//	//)
//	//
//	//wg.Add(2)
//	//
//	//
//	//go func(entries *entries, wg *sync.WaitGroup, err *error) {
//	//	result, e := importRecords(inputType, source, ignoreAttr)
//	//	*entries = result
//	//	*err = e
//	//	wg.Done()
//	//}(&sourceEntries, &wg, &sourceErr)
//	//
//	//go func(entries *entries, wg *sync.WaitGroup, err *error) {
//	//	result, e := importRecords(inputType, target, ignoreAttr)
//	//	*entries = result
//	//	*err = e
//	//	wg.Done()
//	//}(&targetEntries, &wg, &targetErr)
//	//
//	//wg.Wait()
//	//
//	//if sourceErr != nil {
//	//	return "", sourceErr
//	//}
//	//if targetErr != nil {
//	//	return "", targetErr
//	//}
//	//
//	//// Compare the files
//	//return compare(&sourceEntries, &targetEntries, nil)
//	return nil, nil
//}
//
//func diffDn(inputType inputType, source, target, ignoreAttr []string) ([]string, error) {
//	return nil, nil
//}
//
//func genericDiff(sourceParam, targetParam string, ignoreAttr []string, fn fn, dnList *[]string) (string, error) {
//	// Read the files in memory as a Map with sorted attributes
//	var source, target entries
//	var sourceErr, targetErr error
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go func(entries *entries, wg *sync.WaitGroup, err *error) {
//		result, e := fn(sourceParam, ignoreAttr)
//		*entries = result
//		*err = e
//		wg.Done()
//	}(&source, &wg, &sourceErr)
//
//	go func(entries *entries, wg *sync.WaitGroup, err *error) {
//		result, e := fn(targetParam, ignoreAttr)
//		*entries = result
//		*err = e
//		wg.Done()
//	}(&target, &wg, &targetErr)
//
//	wg.Wait()
//
//	if sourceErr != nil {
//		return "", sourceErr
//	}
//	if targetErr != nil {
//		return "", targetErr
//	}
//
//	// Compare the files
//	return compare(&source, &target, dnList)
//}
//
///* Package private functions */
//
//func arraysEqual(a, b []string) bool {
//	if a == nil && b == nil {
//		return true
//	}
//	if a == nil || b == nil {
//		return false
//	}
//	if len(a) != len(b) {
//		return false
//	}
//	for i := range a {
//		if a[i] != b[i] {
//			return false
//		}
//	}
//	return true
//}
//
//// Ordering Logic:
//// Add: entries from source sorted S -> L. Otherwise is invalid.
//// Remove: entries from target sorted L -> S. Otherwise is invalid.
//// Modify:
//// - Keep S ->  L ordering
//// - If only 1 instance of attribute with different value on source and target:
//// update. This way we don't break the applicable LDAP schema.
//// - extra attribute on source: Add
//// - extra attribute on target: Delete
//
//func compare(source, target *entries, dnList *[]string) (string, error) {
//	var buffer bytes.Buffer
//	var err error
//	queue := make(chan actionEntry, 10)
//	var wg sync.WaitGroup
//
//	// Find the order in which operation must happen
//	orderedSourceShortToLong := sortDnByDepth(source, false)
//	orderedTargetLongToShort := sortDnByDepth(target, true)
//
//	// Write the file concurrently
//	wg.Add(1) // 1 writer
//	go createLDIF(queue, &buffer, &wg, &err)
//
//	// Dn only on source + removal of identical entries
//	skipDnForDelete = make(map[string]bool) // Keep track of dn to skip at Deletion
//	sendForAddition(&orderedSourceShortToLong, source, target, queue, dnList)
//
//	// Dn only on target
//	sendForDeletion(&orderedTargetLongToShort, source, target, queue, dnList)
//
//	// Dn on source and target
//	sendForModification(&orderedSourceShortToLong, source, target, queue, dnList)
//
//	// Done sending work
//	close(queue)
//
//	// Free some memory
//	*source = entries{}
//	*target = entries{}
//
//	// Wait for the creation of the LDIF
//	wg.Wait()
//
//	// Return the results
//	return buffer.String(), err
//}
//
////func elementInArray(a string, array []string) bool {
////	for _, b := range array {
////		if b == a {
////			return true
////		}
////	}
////	return false
////}
//
//func sendForAddition(
//	orderedSourceShortToLong *[]string,
//	source, target *entries,
//	queue chan<- actionEntry,
//	dnList *[]string) {
//	for _, dn := range *orderedSourceShortToLong {
//		// Ignore equal entries
//		if arraysEqual((*source)[dn], (*target)[dn]) {
//			Delete(*source, dn)
//			Delete(*target, dn)
//			continue
//		}
//		// Mark entries for addition if only on source
//		if _, ok := (*target)[dn]; !ok {
//			if dnList == nil {
//				subActionAttr := make(map[modifyType][]string)
//				subActionAttr[modifyNone] = (*source)[dn]
//				actionEntry :=
//					actionEntry{
//						Dn:             dn,
//						Action:         Add,
//						SubActionAttrs: []subActionAttrs{subActionAttr},
//					}
//				queue <- actionEntry
//			} else {
//				// Always Add (attributes not relevant)
//				*dnList = append(*dnList, dn)
//			}
//			Delete(*source, dn)
//		}
//		// Implicit else:
//		// It exists on target and it's not equal, so it's a modifyStr
//		skipDnForDelete[dn] = true
//	}
//}
//
//func sendForDeletion(
//	orderedTargetLongToShort *[]string,
//	source, target *entries,
//	queue chan<- actionEntry,
//	dnList *[]string) {
//	for _, dn := range *orderedTargetLongToShort {
//		if skipDnForDelete[dn] { // We know it's not a Delete operation
//			continue
//		}
//		if _, ok := (*target)[dn]; ok { // It has not been deleted above
//			if _, ok := (*source)[dn]; !ok { // does not exists on source
//				if dnList == nil {
//					subActionAttr := make(map[modifyType][]string)
//					subActionAttr[modifyNone] = nil
//					actionEntry :=
//						actionEntry{
//							Dn:             dn,
//							Action:         Delete,
//							SubActionAttrs: []subActionAttrs{subActionAttr},
//						}
//					queue <- actionEntry
//				} else {
//					// Always remove (attributes are not relevant)
//					*dnList = append(*dnList, dn)
//				}
//				Delete(*target, dn)
//			}
//			// Implicit else:
//			// It exists on source and it's not equal (tested on sendForAddition),
//			// so it's a modifyStr
//		}
//	}
//	// Free some memory
//	skipDnForDelete = nil
//}
//
//func sendForModification(
//	orderedSourceShortToLong *[]string, source,
//	target *entries,
//	queue chan<- actionEntry,
//	dnList *[]string) {
//	for _, dn := range *orderedSourceShortToLong {
//		// DN is present on source and target:
//		// sendForAdd/Remove clean up source and target
//		_, okSource := (*source)[dn]
//		_, okTarget := (*target)[dn]
//		if okSource && okTarget { // it hasn't been deleted
//			if dnList == nil {
//
//				// Store the attributes to be added, deleted or replaced
//				attrToModifyAdd := []string{}
//				attrToModifyDelete := []string{}
//				attrToModifyReplace := []string{}
//				// Put the attributes in a map for easy lookup
//				sourceAttr := make(map[string]bool)
//				targetAttr := make(map[string]bool)
//				for _, attr := range (*source)[dn] {
//					sourceAttr[attr] = true
//				}
//				for _, attr := range (*target)[dn] {
//					targetAttr[attr] = true
//				}
//
//				// Compare attribute values starting from the source
//				for _, attr := range (*source)[dn] { // Keep the order of the attributes
//					// Attribute is not equal on both sides
//					if _, ok := targetAttr[attr]; !ok {
//						// Is the attribute name (not value) unique?
//						switch uniqueAttrName(attr, sourceAttr, targetAttr) {
//						case true: // This is a Modify-Replace operation
//							attrToModifyReplace = append(attrToModifyReplace, attr)
//						case false: // This is just a Modify Add (only on source).
//							attrToModifyAdd = append(attrToModifyAdd, attr)
//						}
//					}
//				}
//
//				// Compare attribute values starting from the target.
//				for _, attr := range (*target)[dn] { // Keep the order of the attributes
//					// Looking for unique attributes
//					if !uniqueAttrName(attr, sourceAttr, targetAttr) {
//						if _, ok := sourceAttr[attr]; !ok {
//							attrToModifyDelete = append(attrToModifyDelete, attr)
//						}
//					}
//				}
//
//				// Send it
//				actionEntry := actionEntry{Dn: dn, Action: Modify}
//				subActionAttrArray := []subActionAttrs{}
//				switch {
//				case len(attrToModifyAdd) > 0:
//					subActionAttrArray = append(subActionAttrArray, subActionAttrs{modifyAdd: attrToModifyAdd})
//					fallthrough
//				case len(attrToModifyDelete) > 0:
//					subActionAttrArray = append(subActionAttrArray, subActionAttrs{modifyDelete: attrToModifyDelete})
//					fallthrough
//				case len(attrToModifyReplace) > 0:
//					subActionAttrArray = append(subActionAttrArray, subActionAttrs{modifyReplace: attrToModifyReplace})
//				}
//
//				actionEntry.SubActionAttrs = subActionAttrArray
//				queue <- actionEntry
//			} else {
//				// There must be something left to Modify
//				//if len((*source)[dn]) > 0 || len((*target)[dn]) > 0 {
//				*dnList = append(*dnList, dn)
//				//}
//			}
//			// Clean it up
//			Delete(*source, dn)
//			Delete(*target, dn)
//		}
//	}
//}
//
//func sortDnByDepth(entries *entries, longToShort bool) []string {
//	var sorted []string
//
//	dns := []string{}
//	for dn := range *entries {
//		dns = append(dns, dn)
//	}
//
//	splitByDc := make(map[string][]string)
//	longestDn := 0
//	// Split the components of the dn and remember the longest size
//	for _, dn := range dns {
//		parts := strings.Split(dn, ",")
//		if len(parts) > longestDn {
//			longestDn = len(parts)
//		}
//		splitByDc[dn] = parts
//	}
//
//	// Get the direction of the loop
//	componentSizes := []int{}
//	if longToShort {
//		for i := longestDn; i > 0; i-- {
//			componentSizes = append(componentSizes, i)
//		}
//	} else {
//		for i := 1; i <= longestDn; i++ {
//			componentSizes = append(componentSizes, i)
//		}
//	}
//
//	// Sort by size and alpahbetically within size
//	for _, size := range componentSizes {
//		sameSize := []string{}
//		for dn, components := range splitByDc {
//			if len(components) == size {
//				sameSize = append(sameSize, dn)
//			}
//		}
//		sort.Strings(sameSize)
//		sorted = append(sorted, sameSize...)
//	}
//
//	return sorted
//}
//
//func uniqueAttrName(attr string, sourceAttr, targetAttr map[string]bool) bool {
//
//	// Get the attribute name
//	parts := strings.Split(attr, ":")
//	attrName := parts[0]
//	//var base64 bool
//	//if parts[1] == "" {
//	//	base64 = true
//	//}
//	sourceCounter := 0
//	for attr := range sourceAttr {
//		if strings.HasPrefix(attr, attrName+":") {
//			sourceCounter++
//		}
//	}
//	targetCounter := 0
//	for attr := range targetAttr {
//		if strings.HasPrefix(attr, attrName+":") {
//			targetCounter++
//		}
//	}
//
//	return sourceCounter == 1 && targetCounter == 1
//}
