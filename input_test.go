package ldifdiff

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadIntoChan(t *testing.T) {
	expectedBytes, _ := ioutil.ReadFile("t/modify.ldif")

	// String source
	input, err := readIntoChan(inputStr, string(expectedBytes))
	assert.NoError(t, err)

	var buffer bytes.Buffer

	for line := range input {
		buffer.WriteString(line)
	}

	assert.Equal(t, string(expectedBytes), buffer.String())

	// File source
	input, err = readIntoChan(inputFile, "t/modify.ldif")
	assert.NoError(t, err)
	buffer.Reset()

	for line := range input {
		buffer.WriteString(line)
	}

	assert.Equal(t, string(expectedBytes), buffer.String())
}

//func TestImportLdifFile(t *testing.T) {
//
//	entries, err := importLdifFile(testSourceLdifFile, testIgnoreAttr)
//	okLdifTests(t, entries, testIgnoreAttr, err)
//
//	_, err = importLdifFile(testInvalidLineContLdifFile, testIgnoreAttr)
//	invalidLineCont(t, err)
//
//	_, err = importLdifFile(testInvalidNoDnLdifFile, testIgnoreAttr)
//	invalidNoDn(t, err)
//
//}
//
//func TestConvertLdifStr(t *testing.T) {
//	entries, err := convertLdifStr(testSourceStr, testIgnoreAttr)
//	okLdifTests(t, entries, testIgnoreAttr, err)
//
//	_, err = convertLdifStr(testInvalidLineContStr, testIgnoreAttr)
//	invalidLineCont(t, err)
//
//	entries, err = convertLdifStr(testInvalidNoDnStr, testIgnoreAttr)
//	invalidNoDn(t, err)
//}
//
//func TestImportLdifFileBig(t *testing.T) {
//	if os.Getenv(testBigFilesEnv) != testBigFilesEnvValue {
//		t.Skip("Skipping big files test")
//	}
//	entries, err := importLdifFile(testSourceLdifFileBig, testIgnoreAttr)
//	okLdifTests(t, entries, testIgnoreAttr, err)
//}
//
//func TestConvertLdifStrBig(t *testing.T) {
//	if os.Getenv(testBigFilesEnv) != testBigFilesEnvValue {
//		t.Skip("Skipping big files test")
//	}
//	entries, err := convertLdifStr(testSourceStrBig, testIgnoreAttr)
//	okLdifTests(t, entries, testIgnoreAttr, err)
//}
//
///* Helper test evaluation */
//func okLdifTests(t *testing.T, entries entries, ignoreAttr []string, err error) {
//	if err != nil {
//		t.Error("Expected values, got error: ", err)
//	}
//	if len(entries) != testSourceNrEntries {
//		t.Error("Expected", testSourceNrEntries, "entries, got", strconv.Itoa(len(entries)))
//	}
//	for dn, attributes := range entries {
//		if !strings.HasPrefix(dn, "dn:") {
//			t.Error("Invalid dn:", dn)
//		}
//		for _, attr := range attributes {
//			if strings.HasPrefix(attr, "#") {
//				t.Error("Invalid comment:", attr)
//			}
//			if strings.HasPrefix(attr, " ") {
//				t.Error("Line continuation not correctly appended:", attr)
//			}
//			if strings.IndexAny(attr, " :") < 1 {
//				t.Error("Invalid attribute line:", attr)
//			}
//		}
//	}
//	if len(ignoreAttr) > 0 {
//		for _, attributes := range entries {
//			for _, attr := range attributes {
//				for _, ignoredAttr := range ignoreAttr {
//					if strings.HasPrefix(attr, ignoredAttr+":") {
//						t.Error("Attributed not ignored as requested:",
//							ignoredAttr, attr)
//					}
//				}
//			}
//		}
//	}
//}
//
//func invalidLineCont(t *testing.T, err error) {
//	if err == nil {
//		t.Error("Error expected (line continuation), but none received")
//	}
//}
//
//func invalidNoDn(t *testing.T, err error) {
//	if err == nil {
//		t.Error("Error expected (no dn), but none received")
//	}
//}
