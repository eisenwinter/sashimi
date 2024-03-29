package build

//SashimiErrorCode represents possible errors as numbers
type SashimiErrorCode int

const (
	//SashminiUnknownPropertyPath Unknown property path: `%s`
	SashminiUnknownPropertyPath SashimiErrorCode = 100
	//SashimiUnknownEntity Unknown entity: `%s`
	SashimiUnknownEntity SashimiErrorCode = 104
	//SashminiUndefinedLayoutSection Undefined layout section `%s`
	SashminiUndefinedLayoutSection = 101
	//SashimiPropertyAlreadyDeclared Property with name `%s` has already been declared
	SashimiPropertyAlreadyDeclared = 103
	//SashimiInvalidAlias Invalid alias declaration at this point
	SashimiInvalidAlias = 104
	//SashimiInvalidScopingEnd occurs when a end tag has no matching start tag
	SashimiInvalidScopingEnd = 105
	//SashimiInvalidScopingStart occurs when a start tag has no matching end tag
	SashimiInvalidScopingStart = 106
	//SashminiUnusedLayoutSection Unused layout section `%s`
	SashminiUnusedLayoutSection = 200
	//SashimiEntityAlreadyDeclared Entity with name `%s` has already been declared
	SashimiEntityAlreadyDeclared = 201
)
