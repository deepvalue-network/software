package windows

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the window builder
type Builder interface {
	Create() Builder
	WithTitle(title string) Builder
	WithWidth(width uint) Builder
	WithHeight(height uint) Builder
	IsResizable() Builder
	IsFullscreen() Builder
	Now() (Window, error)
}

// Window represents a windows
type Window interface {
	Title() string
	Width() uint
	Height() uint
	IsResizable() bool
	IsFullscreen() bool
}
