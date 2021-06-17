package ldifdiff

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"sync"
)

func createLDIF(input <-chan DNAction, output chan<- string, wg *sync.WaitGroup) {
	defer close(output)
	defer wg.Done()

	var buffer bytes.Buffer

	for dnAction := range input {
		buffer.WriteString(fmt.Sprintf("dn: %s\n", dnAction.DN))
		buffer.WriteString("changetype: ")

		switch dnAction.Action { //nolint:exhaustive
		case Add:
			buffer.WriteString("add\n")

			for _, attr := range dnAction.Attributes {
				buffer.WriteString(fmt.Sprintf("%s: %s\n", attr.Name, attr.Value))
			}
		case Delete:
			buffer.WriteString("delete\n")
		case Modify:
			buffer.WriteString("modify\n")
			buffer.WriteString(createModifyStr(dnAction))
		}

		buffer.WriteString("\n") // empty line as record separator
	}

	output <- buffer.String()
}

func createModifyStr(dnAction DNAction) string {
	var buffer bytes.Buffer

	for idx, attr := range dnAction.Attributes {
		if idx != 0 {
			buffer.WriteString("-\n")
		}

		switch attr.ModifyType { //nolint:exhaustive
		case ModifyAdd:
			buffer.WriteString("add: ")
		case ModifyDelete:
			buffer.WriteString("delete: ")
		case ModifyReplace:
			buffer.WriteString("replace: ")
		}

		buffer.WriteString(attr.Name + "\n")

		if attr.ModifyType == ModifyDelete {
			continue
		}

		if attr.Base64 {
			value := b64.StdEncoding.EncodeToString([]byte(attr.Value))
			buffer.WriteString(fmt.Sprintf("%s:: %s\n", attr.Name, value))
		} else {
			buffer.WriteString(fmt.Sprintf("%s: %s\n", attr.Name, attr.Value))
		}
	}

	return buffer.String()
}
