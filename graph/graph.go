package graph

//CardinalityType defins whetver something exists in a one or many relationship
type CardinalityType uint8

const (
	//One represents the cardinality to one (1)
	One CardinalityType = 1
	//Many represents the cardinality to many (n)
	Many = 2
)

//Field is a field in a node
type Field struct {
	Name      string
	Alias     string
	Type      string
	Atoms     []string
	KeyValues map[string]string
}

//Edge represents a edge between nodes, it defines wich field of the given node is creating the edge  to wich entity with wich cardinality
type Edge struct {
	From        string
	To          string
	Cardinality CardinalityType
}

//Node is a node in the graph
type Node struct {
	Name         string
	ConstraintTo CardinalityType
	Fields       []*Field
	Edges        []*Edge
}

//SchemaGraph is a collection of a all nodes
type SchemaGraph struct {
	Nodes []*Node
}

//Diff allows to get the differences between two SchemaGraphs
func (s *SchemaGraph) Diff(graph *SchemaGraph) {}
