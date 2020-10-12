package build

type defTableEntry struct {
	Identifier string
	IsDefined  bool
	IsUnique   bool
	Properties map[string]*propertyDef
}

type aliasEntry struct {
	Scope          string
	Source         string
	UnderlyingType string
}
