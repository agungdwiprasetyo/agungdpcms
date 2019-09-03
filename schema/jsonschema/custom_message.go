package jsonschema

type customMessage struct{}

func (l customMessage) Required() string {
	return `{{.property}} is required`
}

func (l customMessage) InvalidType() string {
	return `Invalid type. Expected: {{.expected}}, given: {{.given}}`
}

func (l customMessage) NumberAnyOf() string {
	return `Must validate at least one schema (anyOf)`
}

func (l customMessage) NumberOneOf() string {
	return `Must validate one and only one schema (oneOf)`
}

func (l customMessage) NumberAllOf() string {
	return `Must validate all the schemas (allOf)`
}

func (l customMessage) NumberNot() string {
	return `Must not validate the schema (not)`
}

func (l customMessage) MissingDependency() string {
	return `Has a dependency on {{.dependency}}`
}

func (l customMessage) Internal() string {
	return `Internal Error {{.error}}`
}

func (l customMessage) Const() string {
	return `{{.field}} does not match: {{.allowed}}`
}

func (l customMessage) Enum() string {
	return `{{.field}} must be one of the following: {{.allowed}}`
}

func (l customMessage) ArrayNoAdditionalItems() string {
	return `No additional items allowed on array`
}

func (l customMessage) ArrayNotEnoughItems() string {
	return `Not enough items on array to match positional list of schema`
}

func (l customMessage) ArrayMinItems() string {
	return `Array must have at least {{.min}} items`
}

func (l customMessage) ArrayMaxItems() string {
	return `Array must have at most {{.max}} items`
}

func (l customMessage) Unique() string {
	return `{{.type}} items[{{.i}},{{.j}}] must be unique`
}

func (l customMessage) ArrayContains() string {
	return `At least one of the items must match`
}

func (l customMessage) ArrayMinProperties() string {
	return `Must have at least {{.min}} properties`
}

func (l customMessage) ArrayMaxProperties() string {
	return `Must have at most {{.max}} properties`
}

func (l customMessage) AdditionalPropertyNotAllowed() string {
	return `Additional property {{.property}} is not allowed`
}

func (l customMessage) InvalidPropertyPattern() string {
	return `Property "{{.property}}" does not match pattern {{.pattern}}`
}

func (l customMessage) InvalidPropertyName() string {
	return `Property name of "{{.property}}" does not match`
}

func (l customMessage) StringGTE() string {
	return `String length must be greater than or equal to {{.min}}`
}

func (l customMessage) StringLTE() string {
	return `Panjang karakter tidak boleh melebihi {{.max}}`
}

func (l customMessage) DoesNotMatchPattern() string {
	return `Tidak dalam format yang benar`
}

func (l customMessage) DoesNotMatchFormat() string {
	return `Tidak dalam format '{{.format}}'`
}

func (l customMessage) MultipleOf() string {
	return `Must be a multiple of {{.multiple}}`
}

func (l customMessage) NumberGTE() string {
	return `Must be greater than or equal to {{.min}}`
}

func (l customMessage) NumberGT() string {
	return `Must be greater than {{.min}}`
}

func (l customMessage) NumberLTE() string {
	return `Must be less than or equal to {{.max}}`
}

func (l customMessage) NumberLT() string {
	return `Must be less than {{.max}}`
}

// Schema validators
func (l customMessage) RegexPattern() string {
	return `Invalid regex pattern '{{.pattern}}'`
}

func (l customMessage) GreaterThanZero() string {
	return `{{.number}} must be strictly greater than 0`
}

func (l customMessage) MustBeOfA() string {
	return `{{.x}} must be of a {{.y}}`
}

func (l customMessage) MustBeOfAn() string {
	return `{{.x}} must be of an {{.y}}`
}

func (l customMessage) CannotBeUsedWithout() string {
	return `{{.x}} cannot be used without {{.y}}`
}

func (l customMessage) CannotBeGT() string {
	return `{{.x}} cannot be greater than {{.y}}`
}

func (l customMessage) MustBeOfType() string {
	return `{{.key}} must be of type {{.type}}`
}

func (l customMessage) MustBeValidRegex() string {
	return `{{.key}} must be a valid regex`
}

func (l customMessage) MustBeValidFormat() string {
	return `{{.key}} must be a valid format {{.given}}`
}

func (l customMessage) MustBeGTEZero() string {
	return `{{.key}} must be greater than or equal to 0`
}

func (l customMessage) KeyCannotBeGreaterThan() string {
	return `{{.key}} cannot be greater than {{.y}}`
}

func (l customMessage) KeyItemsMustBeOfType() string {
	return `{{.key}} items must be {{.type}}`
}

func (l customMessage) KeyItemsMustBeUnique() string {
	return `{{.key}} items must be unique`
}

func (l customMessage) ReferenceMustBeCanonical() string {
	return `Reference {{.reference}} must be canonical`
}

func (l customMessage) NotAValidType() string {
	return `has a primitive type that is NOT VALID -- given: {{.given}} Expected valid values are:{{.expected}}`
}

func (l customMessage) Duplicated() string {
	return `{{.type}} type is duplicated`
}

func (l customMessage) HttpBadStatus() string {
	return `Could not read schema from HTTP, response status is {{.status}}`
}

// Replacement options: field, description, context, value
func (l customMessage) ErrorFormat() string {
	return `{{.field}}: {{.description}}`
}

//Parse error
func (l customMessage) ParseError() string {
	return `Expected: {{.expected}}, given: Invalid JSON`
}

//If/Else
func (l customMessage) ConditionThen() string {
	return `Must validate "then" as "if" was valid`
}

func (l customMessage) ConditionElse() string {
	return `Must validate "else" as "if" was not valid`
}
