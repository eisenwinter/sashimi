package resolver

//Predicate represents a predicate
type Predicate func(expresion string) bool

//EntityResolver enables to resolve entities from any given datasource
type EntityResolver interface {
	//ResolveByID gets a enitity by its unique id
	ResolveByID(entityType string, id int) (map[string]interface{}, error)
	//ResolveAll gets all avaiable entites of the given type
	ResolveAll(entityType string) ([]map[string]interface{}, error)
	//Resolve gets a all entities wich satisfy the predicate supplied
	Resolve(entityType string, predicate Predicate) ([]map[string]interface{}, error)
}
