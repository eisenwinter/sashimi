package resolver

//Predicate represents a predicate
type Predicate func(expresion string) bool

//NodeResolver enables to resolve nodes from any given graph datasource
type NodeResolver interface {
	//ResolveByID gets a enitity by its unique id
	ResolveByID(nodeType string, id int) (map[string]interface{}, error)
	//ResolveAll gets all avaiable entites of the given type
	ResolveAll(nodeType string) ([]map[string]interface{}, error)
	//Resolve gets a all entities wich satisfy the predicate supplied
	Resolve(nodeType string, predicate Predicate) ([]map[string]interface{}, error)
}
