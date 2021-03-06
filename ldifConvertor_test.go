package ldifdiff

import (
	"bytes"
	"sync"
	"testing"
)

func TestCreateModifyStr(t *testing.T) {

	// Regular use
	modifyLdif, _ := createModifyStr(testActionEntryTestData.ModifyOnlyAdd)
	if modifyLdif != testModifyAddStr {
		t.Error("Expected:\n[" + testModifyAddStr + "]\nGot:\n[" + modifyLdif + "]\n")
	}

	modifyLdif, _ = createModifyStr(testActionEntryTestData.ModifyOnlyDelete)
	if modifyLdif != testModifyDeleteStr {
		t.Error("Expected:\n[" + testModifyAddStr + "]\nGot:\n[" + modifyLdif + "]\n")
	}
	modifyLdif, _ = createModifyStr(testActionEntryTestData.ModifyOnlyReplace)
	if modifyLdif != testModifyReplaceStr {
		t.Error("Expected:\n[" + testModifyAddStr + "]\nGot:\n[" + modifyLdif + "]\n")
	}

	// Errors
	_, err := createModifyStr(testActionEntryTestData.ModifyNone)
	if err == nil {
		t.Error("Invalid Subaction subActionNone for action actionModify")
	}
}

func TestWriteLdif(t *testing.T) {
	var buffer bytes.Buffer
	var wg sync.WaitGroup
	var err error
	queue := make(chan actionEntry)
	wg.Add(2)
	go func(queue chan actionEntry) {
		queue <- testActionEntryTestData.Add
		queue <- testActionEntryTestData.Delete
		queue <- testActionEntryTestData.Modify
		close(queue)
		wg.Done()
	}(queue)

	go writeLdif(queue, &buffer, &wg, &err)
	wg.Wait()

	if err != nil {
		t.Error("Error not expected, got: ", err)
	}
	ldif := buffer.String()
	if ldif != testModifyStr {
		t.Error("Expected:\n[" + testModifyStr + "]\nGot:\n[" + ldif + "]\n")
	}
}

func TestWriteLdifError(t *testing.T) {
	var buffer bytes.Buffer
	var wg sync.WaitGroup
	var err error
	queue := make(chan actionEntry)
	wg.Add(2)
	go func(queue chan actionEntry) {
		actionEntry := actionEntry{Dn: testDn, Action: 100,
			SubActionAttrs:            []subActionAttrs{{subActionNone: testAttrList}}}
		queue <- actionEntry
		close(queue)
		wg.Done()
	}(queue)

	go writeLdif(queue, &buffer, &wg, &err)
	wg.Wait()

	if err == nil {
		t.Error("Error expected for invalid action")
	}
}
