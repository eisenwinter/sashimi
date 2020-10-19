package resolver

//TestCreateMockPageEntities creates mocked basic entites for testing
func TestCreateMockPageEntities() []*BaseEntity {
	payload := map[string]interface{}{"title": "Test Page One", "description": "Description"}
	entityOne := &BaseEntity{
		Entity:     payload,
		id:         1,
		slug:       "test-one",
		locale:     "en",
		entityType: "page",
		atoms:      map[string][]string{"description": {"markdown"}},
		kvps:       make(map[string][]string),
	}
	payloadTwo := map[string]interface{}{"title": "Test Page Two", "description": "# Description 2"}
	entityTwo := &BaseEntity{
		Entity:     payloadTwo,
		id:         2,
		slug:       "test-two",
		locale:     "en",
		entityType: "site",
		atoms:      map[string][]string{"description": {"markdown"}},
		kvps:       make(map[string][]string),
	}
	return []*BaseEntity{entityOne, entityTwo}

}
