package build

import "fmt"

//SashimiKind represents the kind of type
type SashimiKind string

const (
	//SashimiScalar represents a scalar value
	SashimiScalar SashimiKind = "Scalar"
	//SashimiUnion represents a union type
	SashimiUnion = "Union"
	//SashimiList represents an iterable abstract list type
	SashimiList = "List"
	//SashimiReference represents a followable reference
	SashimiReference = "Ref"
)

//SashimiType represents the underlying primitive type
type SashimiType string

const (
	//SashimiTypeText is a text type commonly known as string
	SashimiTypeText SashimiType = "text"
	//SashimiTypePicture is an image
	SashimiTypePicture = "picture"
	//SashimiTypeNumber is a abtiary number
	SashimiTypeNumber = "number"
	//SashimiTypeBool is a boolean type
	SashimiTypeBool = "bool"
	//SashimiTypeEntity is a reference typ for a compound type
	SashimiTypeEntity = "entity"
	//SashimiTypeTextUnion is union of a text type
	SashimiTypeTextUnion = "text-union"
)

type typeInf interface {
	String() string
	Kind() SashimiKind
	Type() SashimiType
	ResolveTypeName() string
	HKT() typeInf
}

type propertyDef struct {
	Alias string
	Type  typeInf
}

type unionType struct {
	isExcl bool
	tags   []string
}

func (t *unionType) Kind() SashimiKind {
	return "Union"
}

func (t *unionType) String() string {
	if t.isExcl {
		return "oneOf"
	}
	return "manyOf"
}

func (t *unionType) ResolveTypeName() string {
	return "Union<text>"
}

func (t *unionType) Type() SashimiType {
	return SashimiTypeTextUnion
}

func (t *unionType) HKT() typeInf {
	return nil
}

type listType struct {
	baseType typeInf
}

func (t *listType) String() string {
	return "list of " + t.baseType.String()
}

func (t *listType) Kind() SashimiKind {
	return "List"
}

func (t *listType) ResolveTypeName() string {
	return fmt.Sprintf("List<%s>", t.baseType.ResolveTypeName())
}

func (t *listType) Type() SashimiType {
	return t.baseType.Type()
}

func (t *listType) HKT() typeInf {
	return t.baseType
}

type scalarType struct {
	name       string
	constraint map[string]string
}

func (t *scalarType) String() string {
	if t.constraint != nil {
		constraint := ""
		for k, v := range t.constraint {
			constraint += " " + k + "=" + v
		}
		return t.name + "[" + constraint + "]"
	}
	return t.name
}

func (t *scalarType) Kind() SashimiKind {
	return "Scalar"
}

func (t *scalarType) Type() SashimiType {
	return SashimiType(t.name)
}

func (t *scalarType) HKT() typeInf {
	return nil
}

func (t *scalarType) ResolveTypeName() string {
	return t.name
}

type refType struct {
	refOf string
}

func (t *refType) String() string {
	return "->" + t.refOf
}

func (t *refType) Kind() SashimiKind {
	return "Ref"
}

func (t *refType) Type() SashimiType {
	return SashimiType(t.refOf)
}

func (t *refType) HKT() typeInf {
	return nil
}

func (t *refType) ResolveTypeName() string {
	return t.refOf
}

func newList(baseType typeInf) typeInf {
	return &listType{baseType}
}
func newUnion(excl bool, tags []string) typeInf {
	return &unionType{excl, tags}
}
func newRef(refTo string) typeInf {
	return &refType{refTo}
}

type typeBuilder interface {
	IsAwaitingConstraint() bool
	Union(excl bool, tags []string)
	List()
	Scalar(name string)
	Ref(refTo string)
	Constraint(key string, value string)
	Build() typeInf
	Reset()
	Accepts() bool
}

func newTypeBuilder() typeBuilder {
	return &defaultTypeBuilder{}
}

type defaultTypeBuilder struct {
	scalarRead bool
	accepts    bool

	final  typeInf
	l      bool
	scalar scalarType
}

func (d *defaultTypeBuilder) IsAwaitingConstraint() bool {
	return d.scalarRead
}
func (d *defaultTypeBuilder) Accepts() bool {
	return d.accepts
}

func (d *defaultTypeBuilder) Union(excl bool, tags []string) {
	d.final = newUnion(excl, tags)
	d.accepts = false
}

func (d *defaultTypeBuilder) List() {
	d.l = true
}

func (d *defaultTypeBuilder) Scalar(name string) {
	d.scalar = scalarType{
		name:       name,
		constraint: make(map[string]string, 0),
	}
	d.scalarRead = true
}

func (d *defaultTypeBuilder) Ref(refTo string) {
	d.final = newRef(refTo)
}

func (d *defaultTypeBuilder) Constraint(key string, value string) {
	d.scalar.constraint[key] = value
}

func (d *defaultTypeBuilder) Build() typeInf {
	if d.scalarRead {
		scalar := d.scalar
		d.final = &scalar
		d.accepts = false
	}
	if d.l {
		return newList(d.final)
	}
	return d.final

}

func (d *defaultTypeBuilder) Reset() {
	d.final = nil
	d.accepts = true
	d.l = false
	d.scalarRead = false
	d.scalar = scalarType{}
}
