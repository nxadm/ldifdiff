//Compare two LDIF files and output the differences as a valid LDIF.
//Bugs to https://github.com/nxadm/ldifdiff.
//
//    _       _       _       _       _       _       _       _
// _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
//*(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
// // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\
//
//Usage:
//ldifdiff <source> <target> [-i <attributes> ...] [-d]
//ldifdiff -h
//ldifdiff -v
//Options:
//-d, --dn
//  Only print DNs instead of a full LDIF.
//-i <attributes>, --ignore <attributes>
//  Comma separated attribute list to be ignored.
//  Multiple instances of this switch are allowed.
//-h, --help
//  Show this screen.
//-v, --version
//  Show version.
package main

//
//import (
//	"fmt"
//	"os"
//	"strings"
//
//	"github.com/docopt/docopt-go"
//	"github.com/nxadm/ldifdiff"
//)
//
//// Used by the implementation program in the cmd directory.
//const Version = "v0.2.0"
//
//// Used by the implementation program in the cmd directory.
//const Author = "Claudio Ramirez <pub.claudio@gmail.com>"
//
//// Used by the implementation program in the cmd directory.
//const Repo = "https://github.com/nxadm/ldifdiff"
//
//
//type Params struct {
//	Source, Target string
//	IgnoreAttr     []string
//	DnOnly         bool
//}
//
//var versionMsg = "ldifdiff " + Version + " (" + Author + ")."
//var usage = versionMsg + "\n" +
//	"Compare two LDIF files and output the differences as a valid LDIF.\n" +
//	"Bugs to " + Repo + ".\n" + `
//       _       _       _       _       _       _       _       _
//    _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
//  *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
//  // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\
//
//Usage:
//  ldifdiff <source> <target> [-i <attributes> ...] [-d]
//  ldifdiff -h
//  ldifdiff -v
//Options:
//  -d, --dn
//    Only print DNs instead of a full LDIF.
//  -i <attributes>, --ignore <attributes>
//	Comma separated attribute list to be ignored.
//	Multiple instances of this switch are allowed.
//  -h, --help
//  	Show this screen.
//  -v, --version
//  	Show version
//`
//
//func main() {
//	params := Params{}
//	params.parse()
//
//	/* DiffFromFiles the files */
//	var output string
//	var err error
//	switch params.DnOnly {
//	case true:
//		var outputList []string
//		outputList, err = ldifdiff.ListDiffDnFromFiles(params.Source, params.Target, params.IgnoreAttr)
//		output = strings.Join(outputList, "\n") + "\n"
//	default:
//		output, err = ldifdiff.DiffFromFiles(params.Source, params.Target, params.IgnoreAttr)
//	}
//
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "%s\n", err)
//		os.Exit(1)
//	}
//
//	fmt.Printf("%s", output)
//}
//
//func (params *Params) parse() {
//	args, err := docopt.Parse(usage, nil, true, versionMsg, false)
//
//	switch {
//	case err != nil:
//		fmt.Fprintf(os.Stderr, "Error parsing the command line arguments:\n%v\n", err)
//		os.Exit(1)
//	case args["--dn"].(bool):
//		params.DnOnly = true
//		fallthrough
//	case args["--ignore"].([]string) != nil:
//		params.IgnoreAttr = args["--ignore"].([]string)
//		fallthrough
//	case args["<source>"].(string) != "":
//		params.Source = args["<source>"].(string)
//		fallthrough
//	case args["<target>"].(string) != "":
//		params.Target = args["<target>"].(string)
//	}
//
//	errMsgs := []string{}
//	switch {
//	case params.Source == params.Target:
//		fmt.Fprintln(os.Stderr, "Source and target ldif files are the same.")
//		os.Exit(1)
//	default:
//		if _, err := os.Stat(params.Source); err != nil {
//			errMsgs = append(errMsgs, err.Error())
//		}
//		if _, err := os.Stat(params.Target); err != nil {
//			errMsgs = append(errMsgs, err.Error())
//		}
//	}
//	if len(errMsgs) > 0 {
//		for _, msg := range errMsgs {
//			fmt.Fprintf(os.Stderr, "%s\n", msg)
//		}
//		os.Exit(1)
//	}
//
//}
