package ldifdiff

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func createLDIF(queue <-chan DNAction, buffer *bytes.Buffer, wg *sync.WaitGroup) {
func TestCreateLDIF(t *testing.T) {
	input := make(chan DNAction, 100)
	output := make(chan string, 100)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go createLDIF(input, output, wg)

	for _, dnAction := range testReturnDNActions() {
		input <- dnAction
	}

	close(input)

	var buffer bytes.Buffer

	for record := range output {
		fmt.Print(record)
		buffer.WriteString(record)
	}

	expectedBytes, _ := ioutil.ReadFile("t/modify.ldif")
	assert.Equal(t, string(expectedBytes), buffer.String())
}

func TestCreateModifyStr(t *testing.T) {
	expectedBytes, _ := ioutil.ReadFile("t/modifyStr.ldif")
	assert.Equal(t, string(expectedBytes), createModifyStr(testReturnDNActions()[0]))
}

func testReturnDNActions() []DNAction {
	var dnActions []DNAction

	attrs := []Attribute{}
	attrs = append(attrs,
		Attribute{
			Name:       "mail",
			Value:      "me@example.com",
			ModifyType: ModifyAdd,
		},
		Attribute{
			Name:       "address",
			Value:      "the sun",
			ModifyType: ModifyReplace,
		},
		Attribute{
			Name:       "phone",
			ModifyType: ModifyDelete,
		},
	)
	dnAction := DNAction{
		DN:         "uid=foo,ou=bar,dc=example,dc=com",
		Action:     Modify,
		Attributes: attrs,
	}
	dnActions = append(dnActions, dnAction)

	attrs = []Attribute{}
	attrs = append(attrs,
		Attribute{
			Name:  "mail",
			Value: "me2@example.com",
		},
		Attribute{
			Name:  "address",
			Value: "the moon",
		},
	)
	dnAction = DNAction{
		DN:         "uid=foo2,ou=bar,dc=example,dc=com",
		Action:     Add,
		Attributes: attrs,
	}
	dnActions = append(dnActions, dnAction)

	dnAction = DNAction{
		DN:     "uid=foo3,ou=bar,dc=example,dc=com",
		Action: Delete,
	}
	dnActions = append(dnActions, dnAction)

	return dnActions
}
