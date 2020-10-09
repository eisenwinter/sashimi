package graph

type CardinalityType uint8

const (
	One  CardinalityType = 1
	Many                 = 2
)

type Field struct {
	Name      string
	Alias     string
	Type      string
	Atoms     []string
	KeyValues map[string]string
}

type Edge struct {
	To          string
	Cardinality CardinalityType
}

type Node struct {
	Name   string
	Fields []*Field
	Edges  []*Edge
}