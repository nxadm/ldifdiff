package ldifdiff

/* Types */
// Each Dn has an action (actionAdd, actionDelete, actionModify) and a set of subctions applied
// on the associated attributes;
// - action actionAdd and actionDelete have no Subactions (typed as 1 Subaction subActionNone) and
// are expected to have only 1 set of attributes. actionAdd will actionAdd the Dn with the
// supplied attributes, while actionDelete will remove the Dn completely (supplied
// attributes are ignored).
// - action actionModify is more complex and has 3 types of Subactions (subActionModifyAdd,
// subActionModifyDelete and ModifyUpdate). A Dn with a actionModify action can have multiple
// combinations of SubActions and associated attributes. In the case of the
// subAction ModifyUpdate only 1 attribute is expected (rfc2849). This is done
// in order to respect possible schema restrictions.

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
