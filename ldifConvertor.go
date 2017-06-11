package ldifdiff

import (
	"bytes"
	"errors"
	"strings"
	"sync"
)

func createModifyStr(actionEntry ActionEntry) (string, error) {
	var buffer bytes.Buffer
	subActions := make(map[SubAction]string)
	subActions[ModifyAdd] = "add"
	subActions[ModifyDelete] = "delete"
	subActions[ModifyReplace] = "replace"
	for idx, subActionList := range actionEntry.SubActionAttrs {
		for subAction, attrList := range subActionList {
			if subAction == None {
				return "", errors.New(("Invalid Subaction None for Action Modify"))
			}
			for idxInner, attr := range attrList {
				if idxInner != 0 || idx != 0 {
					buffer.WriteString("-\n")
				}
				parts := strings.Split(attr, ":")
				buffer.WriteString(subActions[subAction] +
					": " + parts[0] + "\n")
				buffer.WriteString(parts[0] + ":" +
					strings.Join(parts[1:], ":") + "\n")

			}
		}
	}
	return buffer.String(), nil
}

func writeLdif(queue <-chan ActionEntry, writer *bytes.Buffer, wg *sync.WaitGroup, err *error) {
	defer wg.Done()
	for actionEntry := range queue {
		if *err != nil {
			continue
		}
		switch actionEntry.Action {
		case Add:
			writer.WriteString(actionEntry.Dn + "\n") //dn
			writer.WriteString("changetype: add\n")
			attrList := actionEntry.SubActionAttrs[0][None]
			for _, attr := range attrList {
				writer.WriteString(attr + "\n")
			}
		case Delete:
			writer.WriteString(actionEntry.Dn + "\n") //dn
			writer.WriteString("changetype: delete\n")
		case Modify:
			writer.WriteString(actionEntry.Dn + "\n") //dn
			writer.WriteString("changetype: modify\n")
			modifyStr, modifyErr := createModifyStr(actionEntry)
			if modifyErr != nil {
				*err = modifyErr
				continue
			}
			writer.WriteString(modifyStr)
		default:
			*err = errors.New("Unexpected LDIF action value: " + string(actionEntry.Action))
			continue
		}
		writer.WriteString("\n") // empty line as record separator
	}
}
