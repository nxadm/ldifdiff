package ldifdiff

/* Types */
// Each Dn has an Action (Add, Delete, Modify) and a set of subctions applied
// on the associated attributes;
// - Action Add and Delete have no Subactions (typed as 1 Subaction None) and
// are expected to have only 1 set of attributes. Add will add the Dn with the
// supplied attributes, while Delete will remove the Dn completely (supplied
// attributes are ignored).
// - Action Modify is more complex and has 3 types of Subactions (ModifyAdd,
// ModifyDelete and ModifyUpdate). A Dn with a Modify Action can have multiple
// combinations of SubActions and associated attributes. In the case of the
// SubAction ModifyUpdate only 1 attribute is expected (rfc2849). This is done
// in order to respect possible schema restrictions.

type Action int
type SubAction int
type SubActionAttr map[SubAction][]string
type ActionEntry struct {
	Dn             string
	Action         Action
	SubActionAttrs []SubActionAttr
}

// Return map with dn as key and attribute array as value
type Entries map[string][]string

const (
	Add Action = iota
	Delete
	Modify
)
const (
	ModifyAdd SubAction = iota
	ModifyDelete
	ModifyReplace
	None
)
