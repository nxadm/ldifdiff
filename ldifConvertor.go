package ldifdiff

//
//import (
//	"bytes"
//	"errors"
//	"fmt"
//	"strings"
//	"sync"
//)
//
//func createModifyStr(actionEntry actionEntry) (string, error) {
//	var buffer bytes.Buffer
//	subActions := make(map[modifyType]string)
//	subActions[modifyAdd] = "Add"
//	subActions[modifyDelete] = "Delete"
//	subActions[modifyReplace] = "replace"
//	for idx, subActionList := range actionEntry.SubActionAttrs {
//		for subAction, attrList := range subActionList {
//			if subAction == modifyNone {
//				return "", errors.New(("Invalid Subaction modifyNone for typeCheck Modify"))
//			}
//			for idxInner, attr := range attrList {
//				if idxInner != 0 || idx != 0 {
//					buffer.WriteString("-\n")
//				}
//				parts := strings.Split(attr, ":")
//				buffer.WriteString(subActions[subAction] +
//					": " + parts[0] + "\n")
//				buffer.WriteString(parts[0] + ":" +
//					strings.Join(parts[1:], ":") + "\n")
//
//			}
//		}
//	}
//	return buffer.String(), nil
//}
//
//func createLDIF(queue <-chan actionEntry, writer *bytes.Buffer, wg *sync.WaitGroup, err *error) {
//	defer wg.Done()
//	for actionEntry := range queue {
//		if *err != nil {
//			continue
//		}
//		switch actionEntry.Action {
//		case Add:
//			writer.WriteString(actionEntry.Dn + "\n") //dn
//			writer.WriteString("changetype: Add\n")
//			attrList := actionEntry.SubActionAttrs[0][modifyNone]
//			for _, attr := range attrList {
//				writer.WriteString(attr + "\n")
//			}
//		case Delete:
//			writer.WriteString(actionEntry.Dn + "\n") //dn
//			writer.WriteString("changetype: Delete\n")
//		case Modify:
//			writer.WriteString(actionEntry.Dn + "\n") //dn
//			writer.WriteString("changetype: Modify\n")
//			modifyStr, modifyErr := createModifyStr(actionEntry)
//			if modifyErr != nil {
//				*err = modifyErr
//				continue
//			}
//			writer.WriteString(modifyStr)
//		default:
//			*err = errors.New("Unexpected LDIF typeCheck value: " + fmt.Sprintf("%d", actionEntry.Action))
//			continue
//		}
//		writer.WriteString("\n") // empty line as record separator
//	}
//}
