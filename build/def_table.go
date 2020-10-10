package build

type defTableEntry struct {
	Identfier  string
	IsDefined  bool
	IsUnique   bool
	Properties map[string]*propertyDef
}
