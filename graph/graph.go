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
	From        string
	To          string
	Cardinality CardinalityType
}

type Node struct {
	Name         string
	ConstraintTo CardinalityType
	Fields       []*Field
	Edges        []*Edge
}

type SchemaGraph struct {
	Nodes []*Node
}

func (s *SchemaGraph) Diff(graph *SchemaGraph) {}
