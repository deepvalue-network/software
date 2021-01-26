package schemas

type propertyContent struct {
	isPrimaryKey bool
	foreignKey   Schema
	typ          Type
}

func createPropertyContentWithPrimaryKey() PropertyContent {
	return createPropertyContentInternally(true, nil, nil)
}

func createPropertyContentWithForeignKey(foreignKey Schema) PropertyContent {
	return createPropertyContentInternally(false, foreignKey, nil)
}

func createPropertyContentWithType(typ Type) PropertyContent {
	return createPropertyContentInternally(false, nil, typ)
}

func createPropertyContentInternally(
	isPrimaryKey bool,
	foreignKey Schema,
	typ Type,
) PropertyContent {
	out := propertyContent{
		isPrimaryKey: isPrimaryKey,
		foreignKey:   foreignKey,
		typ:          typ,
	}

	return &out
}

// IsPrimaryKey returns true if there is a primaryKey, false otherwise
func (obj *propertyContent) IsPrimaryKey() bool {
	return obj.isPrimaryKey
}

// IsForeignKey returns true if there is a foreignKey, false otherwise
func (obj *propertyContent) IsForeignKey() bool {
	return obj.foreignKey != nil
}

// ForeignKey returns the foreignKey, if any
func (obj *propertyContent) ForeignKey() Schema {
	return obj.foreignKey
}

// IsType returns true if there is a type, false otherwise
func (obj *propertyContent) IsType() bool {
	return obj.typ != nil
}

// Type returns the type, if any
func (obj *propertyContent) Type() Type {
	return obj.typ
}
