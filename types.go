package ldifdiff

/* Types */
// Each Dn has an action (actionAdd, actionDelete, actionModify) and a set of subctions applied
// on the associated attributes;
// - action actionAdd and actionDelete have no subactions (typed as 1 subaction subActionNone) and
// are expected to have only 1 set of attributes. actionAdd will add the Dn with the
// supplied attributes, while actionDelete will delete the Dn completely (supplied
// attributes are ignored).
// - action actionModify is more complex and has 3 types of subactions (subActionModifyAdd,
// subActionModifyDelete and subActionModifyUpdate). A Dn with a actionModify action can have 
// multiple combinations of subActions and associated attributes. In the case of the
// subAction ModifyUpdate, the attribute must be unique in order to to respect possible 
// schema restrictions (rfc2849).

type action int
type subAction int
type subActionAttrs map[subAction][]string
type actionEntry struct {
	Dn             string
	Action         action
	SubActionAttrs []subActionAttrs
}

// Return map with dn as key and attribute array as value
type entries map[string][]string

const (
	actionAdd    action = iota
	actionDelete
	actionModify
)
const (
	subActionModifyAdd     subAction = iota
	subActionModifyDelete
	subActionModifyReplace
	subActionNone
)
