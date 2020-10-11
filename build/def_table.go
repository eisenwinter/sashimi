package build

type defTableEntry struct {
	Identfier  string
	IsDefined  bool
	IsUnique   bool
	Properties map[string]*propertyDef
}

type callEntry struct {
	Scope  string
	Source string
}

type aliasEntry struct {
	Scope          string
	Source         string
	UnderlyingType string
}
