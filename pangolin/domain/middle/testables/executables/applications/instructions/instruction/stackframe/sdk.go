package stackframe

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSkipBuilder creates a new skip instance
func NewSkipBuilder() SkipBuilder {
	return createSkipBuilder()
}

// NewSaveBuilder creates a new save builder instance
func NewSaveBuilder() SaveBuilder {
	return createSaveBuilder()
}

// Builder represents the stackframe builder
type Builder interface {
	Create() Builder
	IsPush() Builder
	IsPop() Builder
	WithSkip(skip Skip) Builder
	WithIndex(indexVariable string) Builder
	WithSave(save Save) Builder
	WithSwitch(swtch string) Builder
	Now() (Stackframe, error)
}

// Stackframe represents a stackframe instruction
type Stackframe interface {
	IsPush() bool
	IsPop() bool
	IsIndex() bool
	Index() string
	IsSkip() bool
	Skip() Skip
	IsSave() bool
	Save() Save
	IsSwitch() bool
	Switch() string
}

// SkipBuilder represents a skip builder
type SkipBuilder interface {
	Create() SkipBuilder
	WithInt(intVal int64) SkipBuilder
	WithVariable(variable string) SkipBuilder
	Now() (Skip, error)
}

// Skip represents a skip
type Skip interface {
	IsInt() bool
	Int() int64
	IsVariable() bool
	Variable() string
}

// SaveBuilder represents a save builder
type SaveBuilder interface {
	Create() SaveBuilder
	From(from string) SaveBuilder
	To(to string) SaveBuilder
	Now() (Save, error)
}

// Save represents a save
type Save interface {
	To() string
	HasFrom() bool
	From() string
}
