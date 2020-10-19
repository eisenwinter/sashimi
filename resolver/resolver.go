package resolver

import (
	"fmt"
	"time"
)

//Predicate represents a predicate
type Predicate func(expresion string) bool

//Entity represents the loaded data
type Entity map[string]interface{}

//BaseEntity is the basic building block for all data
type BaseEntity struct {
	Entity
	id         int
	slug       string
	locale     string
	entityType string
	atoms      map[string][]string
	kvps       map[string][]string
}

//GetID returns the id of a entity
func (b *BaseEntity) GetID() int {
	return b.id
}

//GetSlug returns the slug of a entity
func (b *BaseEntity) GetSlug() string {
	return b.slug
}

//GetLocale gets the localization of the entity
func (b *BaseEntity) GetLocale() string {
	return b.locale
}

//GetType returns the schema / type of the entity
func (b *BaseEntity) GetType() string {
	return b.entityType
}

//GetUnderlyingData returns the underlying data of a entity
func (b *BaseEntity) GetUnderlyingData() Entity {
	return b.Entity
}

//HasAtom returns true if the property has the supplied atom
func (b *BaseEntity) HasAtom(property string, atom string) bool {
	if _, ok := b.atoms[atom]; ok {
		return true
	}
	return false
}

//GetText gets the underlying text field
func (b Entity) GetText(prop string) (string, error) {
	if val, ok := b[prop]; ok {
		return fmt.Sprintf("%v", val), nil
	}
	return "", fmt.Errorf("Field `%s` not found", prop)
}

//GetImage gets the underlying image field
func (b Entity) GetImage(prop string) (href string, alt string, err error) {
	if val, ok := b[prop]; ok {
		switch i := val.(type) {
		case map[string]string:
			return i["href"], i["alt"], nil
		default:
			return "", "", fmt.Errorf("Field `%s` is of wrong type", prop)
		}
	}
	return "", "", fmt.Errorf("Field `%s` not found", prop)
}

//GetNumber gets the underlying number field
func (b Entity) GetNumber(prop string) (float64, error) {
	if val, ok := b[prop]; ok {
		switch i := val.(type) {
		case float64:
			return i, nil
		default:
			return 0, fmt.Errorf("Field `%s` is of wrong type", prop)
		}
	}
	return 0, fmt.Errorf("Field `%s` not found", prop)
}

//GetBool gets the underlying boolean field
func (b Entity) GetBool(prop string) (bool, error) {
	if val, ok := b[prop]; ok {
		switch i := val.(type) {
		case bool:
			return i, nil
		default:
			return false, fmt.Errorf("Field `%s` is of wrong type", prop)
		}
	}
	return false, fmt.Errorf("Field `%s` not found", prop)
}

//GetDate gets the underlying date field
func (b Entity) GetDate(prop string) (time.Time, error) {
	if val, ok := b[prop]; ok {
		switch i := val.(type) {
		case time.Time:
			return i, nil
		default:
			return time.Now(), fmt.Errorf("Field `%s` is of wrong type", prop)
		}
	}
	return time.Now(), fmt.Errorf("Field `%s` not found", prop)
}

//NodeResolver enables to resolve nodes from any given graph datasource
type NodeResolver interface {
	//ResolveByID gets a enitity by its unique id
	ResolveByID(nodeType string, id int) (BaseEntity, error)
	//ResolveAll gets all avaiable entites of the given type
	ResolveAll(nodeType string) ([]BaseEntity, error)
	//Resolve gets a all entities wich satisfy the predicate supplied
	Resolve(nodeType string, predicate Predicate) ([]BaseEntity, error)
}
