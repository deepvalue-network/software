package parsers

type variable struct {
	declaration Declaration
	assignment  Assignment
	concat      Concatenation
	delete      string
}

func createVariableWithDeclaration(declaration Declaration) Variable {
	return createVariableInternally(declaration, nil, nil, "")
}

func createVariableWithAssignment(assignment Assignment) Variable {
	return createVariableInternally(nil, assignment, nil, "")
}

func createVariableWithConcatenation(concat Concatenation) Variable {
	return createVariableInternally(nil, nil, concat, "")
}

func createVariableWithDelete(delete string) Variable {
	return createVariableInternally(nil, nil, nil, delete)
}

func createVariableInternally(
	declaration Declaration,
	assignment Assignment,
	concat Concatenation,
	delete string,
) Variable {
	out := variable{
		declaration: declaration,
		assignment:  assignment,
		concat:      concat,
		delete:      delete,
	}

	return &out
}

// IsDeclaration returns true if there is a declaration, false otherwise
func (obj *variable) IsDeclaration() bool {
	return obj.declaration != nil
}

// Declaration returns the declaration, if any
func (obj *variable) Declaration() Declaration {
	return obj.declaration
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *variable) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *variable) Assignment() Assignment {
	return obj.assignment
}

// IsConcatenation returns true if there is a concatenation, false otherwise
func (obj *variable) IsConcatenation() bool {
	return obj.concat != nil
}

// Concatenation returns the concatenation, if any
func (obj *variable) Concatenation() Concatenation {
	return obj.concat
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *variable) IsDelete() bool {
	return obj.delete != ""
}

// Delete returns the delete, if any
func (obj *variable) Delete() string {
	return obj.delete
}
