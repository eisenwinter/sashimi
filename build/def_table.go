package build

type defTableEntry struct {
	Identfier  string
	IsDefined  bool
	Properties map[string]*propertyDef
}
