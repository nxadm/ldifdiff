package ldifdiff

//import (
//	"io/ioutil"
//	"os"
//	"strings"
//)
//
////* Test data */
//const testBigFilesEnv = "LDIFDIFF_BIGFILES"
//const testBigFilesEnvValue = "1"
//const testDn = "dn: some_dn,ou=aAccounts,dc=domain,dc=ext"
//const testSourceLdifFile = "t/source.ldif"
//const testTargetLdifFile = "t/target.ldif"
//const testResultLdifFile = "t/result.ldif"
//const testResultDnFile = "t/result_dn"
//const testResultDnIgnoreAttrFile = "t/result_dn_ignore_attr"
//const testInvalidLineContLdifFile = "t/invalid_line_continuation.ldif"
//const testInvalidNoDnLdifFile = "t/invalid_no_dn.ldif"
//const testSourceLdifFileBig = "t/_source_big.ldif"
//const testTargetLdifFileBig = "t/_target_big.ldif"
//const testResultLdifFileBig = "t/_result_big.ldif"
//const testModifyAddLdifFile = "t/modifyAdd.ldif"
//const testModifyDeleteLdifFile = "t/modifyDelete.ldif"
//const testModifyReplaceLdifFile = "t/modifyReplace.ldif"
//const testModifyLdifFile = "t/Modify.ldif"
//
//var testSourceStr = testGetLdifeStr(testSourceLdifFile, false)
//var testSourceNrEntries = testGetNrOfEntries(testSourceStr)
//var testTargetStr = testGetLdifeStr(testTargetLdifFile, false)
//var testResultStr = testGetLdifeStr(testResultLdifFile, false)
//var testResultDnStr = testGetLdifeStr(testResultDnFile, false)
//var testResultDnIgnoreAttrStr = testGetLdifeStr(testResultDnIgnoreAttrFile, false)
//var testInvalidLineContStr = testGetLdifeStr(testInvalidLineContLdifFile, false)
//var testInvalidNoDnStr = testGetLdifeStr(testInvalidNoDnLdifFile, false)
//var testSourceStrBig = testGetLdifeStr(testSourceLdifFileBig, true)
//var testTargetStrBig = testGetLdifeStr(testTargetLdifFileBig, true)
//var testResultStrBig = testGetLdifeStr(testResultLdifFileBig, true)
//var testModifyAddStr = testGetLdifeStr(testModifyAddLdifFile, false)
//var testModifyDeleteStr = testGetLdifeStr(testModifyDeleteLdifFile, false)
//var testModifyReplaceStr = testGetLdifeStr(testModifyReplaceLdifFile, false)
//var testModifyStr = testGetLdifeStr(testModifyLdifFile, false)
//var testIgnoreAttr = []string{"sambaSID", "eduPersonEntitlement"}
//var testIgnoreAttrDn = []string{"sambaSID", "eduPersonEntitlement", "mail"}
//var testAttrList = []string{"mail: auth2@domain.ext", "phone: +32364564645"}
//var testAttrListModifyReplace = []string{testAttrList[0]}
//var testActionEntryTestData = testGetActionEntryMap()
//
///* Helper functions and types */
//type TestActionEntryData struct {
//	Add, Delete, Modify, ModifyOnlyAdd,
//	ModifyOnlyDelete, ModifyOnlyReplace,
//	ModifyNone, ModifyReplaceAttributes actionEntry
//}
//
//func testGetActionEntryMap() TestActionEntryData {
//	return TestActionEntryData{
//		Add: actionEntry{Dn: testDn, Action: Add,
//			SubActionAttrs: []subActionAttrs{{modifyNone: testAttrList}}},
//		Delete: actionEntry{Dn: testDn, Action: Delete,
//			SubActionAttrs: []subActionAttrs{{modifyNone: testAttrList}}},
//		Modify: actionEntry{Dn: testDn, Action: Modify,
//			SubActionAttrs: []subActionAttrs{
//				{modifyAdd: testAttrList},
//				{modifyDelete: testAttrList},
//				{modifyReplace: testAttrListModifyReplace}}},
//		ModifyOnlyAdd: actionEntry{Dn: testDn, Action: Modify,
//			SubActionAttrs: []subActionAttrs{{modifyAdd: testAttrList}}},
//		ModifyOnlyDelete: actionEntry{Dn: testDn, Action: Modify,
//			SubActionAttrs: []subActionAttrs{{modifyDelete: testAttrList}}},
//		ModifyOnlyReplace: actionEntry{Dn: testDn, Action: Modify,
//			SubActionAttrs: []subActionAttrs{{modifyReplace: testAttrListModifyReplace}}},
//		ModifyNone: actionEntry{Dn: testDn, Action: Modify,
//			SubActionAttrs: []subActionAttrs{{modifyNone: testAttrList}}},
//		ModifyReplaceAttributes: actionEntry{Dn: testDn, Action: Modify,
//			SubActionAttrs: []subActionAttrs{{modifyReplace: testAttrList}}},
//	}
//}
//
//func testGetLdifeStr(file string, big bool) string {
//	if big && os.Getenv(testBigFilesEnv) != testBigFilesEnvValue {
//		return ""
//	}
//	bytes, _ := ioutil.ReadFile(file)
//	return string(bytes) + "\n"
//}
//
//func testGetNrOfEntries(ldifStr string) int {
//	var counter int
//	for _, line := range strings.Split(ldifStr, "\n") {
//		if strings.HasPrefix(line, "dn:") {
//			counter++
//		}
//	}
//	return counter
//}
