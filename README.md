# ldifdiff
ldiffdiff is a fast Go (golang) library and executable that outputs the 
difference between two LDIF files (LDAP). Binaries of the command implementation
are provided under [releases](https://github.com/nxadm/ldifdiff/releases).


[![Build Status](https://travis-ci.org/nxadm/ldifdiff.svg?branch=master)](https://travis-ci.org/nxadm/ldifdiff)

## Usage of the ldifdiff command
```
$ ./ldifdiff -h
ldifdiff v0.1.0 (Claudio Ramirez <pub.claudio@gmail.com>).
Compare two LDIF files and output the differences as a valid LDIF.
Bugs to https://github.com/nxadm/ldifdiff.

       _       _       _       _       _       _       _       _
    _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
  *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
  // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\

Usage:
  ldifdiff <source> <target> [-i <attributes> ...] [-d]
  ldifdiff -h
  ldifdiff -v
Options:
  -d, --dn
    Only print DNs instead of a full LDIF.
  -i <attributes>, --ignore <attributes>
	Comma separated attribute list.
	Multiple instances of this switch are allowed.
  -h, --help
  	Show this screen.
  -v, --version
  	Show version
```

## Installation

Download the library using go:
```go get github.com/nxadm/ldifdiff```

Import the library into your project:
```import "github.com/nxadm/ldifdiff"```

Compile the ldiff cmd implementation:
```cd cmd; go build -o ldifdiff ldifdiff.go```


## API ##

The API is simple and provides four functions.

```func Diff(sourceStr, targetStr string, ignoreAttr []string) (string, error)```

Diff compares two LDIF strings (sourceStr and targetStr) and outputs the
differences as a LDIF string. An array of attributes can be supplied.
These attributes will be ignored when comparing the LDIF strings. The
output is a string, a valid LDIF, and can be added to the "target"
database (the one that created targetStr) in order to make it equal to
the *source* database (the one that created sourceStr). In case of
failure, an error is provided.


```func DiffFromFiles(sourceFile, targetFile string, ignoreAttr []string) (string, error)```

DiffFromFiles compares two LDIF files (sourceFile and targetFile) and
outputs the differences as a LDIF string. An array of attributes can be
supplied. These attributes will be ignored when comparing the LDIF
strings. The output is a string, a valid LDIF, and can be added to the
"target" database (the one that created targetFile) in order to make it
equal to the *source* database (the one that created sourceFile). In
case of failure, an error is provided.

```func ListDiffDn(sourceStr, targetStr string, ignoreAttr []string) ([]string, error)```

ListDiffDn compares two LDIF strings (sourceStr and targetStr) and
outputs the differences as a list of affected DNs (Dintinguished Names).
An array of attributes can be supplied. These attributes will be ignored
when comparing the LDIF strings. The output is a string slice. In case
of failure, an error is provided.

```func ListDiffDnFromFiles(sourceFile, targetFile string, ignoreAttr []string) ([]string, error)```

ListDiffDnFromFiles compares two LDIF files (sourceFile and
targetFileStr) and outputs the differences as a list of affected DNs
(Dintinguished Names). An array of attributes can be supplied. These
attributes will be ignored when comparing the LDIF strings. The output
is a string slice. In case of failure, an error is provided.
