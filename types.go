package ldifdiff

/* API types */
type (
	// ErrReadLDIF corresponds with an error encountered while reading a LDIF file.
	ErrReadLDIF error

	// ErrParseLDIF corresponds with an error encountered while parsing LDIF contents.
	ErrParseLDIF error

	// ModifyLDIF corresponds with the results of a LDIF comparison in a string
	// format usable by ldapmodify.
	ModifyLDIF string

	// DN corresponds with the Distinguished name of an LDIF entry.
	DN string

	// Action corresponds with the action that needs to be applied to a DN
	// in the target LDIF to match the source.
	Action int

	// ModifyType corresponds with the Modify type that needs to be applied
	// to an attribute of a DN in the target LDIF to match the source.
	ModifyType int
)

// DNInfo corresponds with the structured contents of an LDIF file or string.
type DNInfo map[DN][]Attribute

// DiffResult corresponds with the structured result of an LDIF comparison.
type DiffResult []DNAction

// DNAction hold the action to be done to a DN.
type DNAction struct {
	DN string
	Action
	Attributes []Attribute
}

// Attribute corresponds with an attribute of a DN
type Attribute struct {
	Name   string
	Value  string
	Base64 bool
	ModifyType
}

// AttributeAction corresponds holds the information needed to modify attributes.
type AttributeAction struct {
	Attribute
	ModifyType
}

// Action constants
const (
	None Action = iota
	Add
	Delete
	Modify
)

// Modify constants
const (
	ModifyNone ModifyType = iota
	ModifyAdd
	ModifyDelete
	ModifyReplace
)

/* Internal types */
type inputType int

const (
	// Clarify the input type
	inputStr inputType = iota
	inputFile
)
