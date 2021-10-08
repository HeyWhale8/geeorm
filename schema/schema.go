package schema

// Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema represents a table of database
type Schema  struct {
	Model     interface{}
	Name      string
	Fields    []*Field
	FieldName []string
	fieldMap  map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}
