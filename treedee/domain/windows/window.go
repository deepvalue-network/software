package windows

type window struct {
	title        string
	width        uint
	height       uint
	isResizable  bool
	isFullscreen bool
}

func createWindow(
	title string,
	width uint,
	height uint,
	isResizable bool,
	isFullscreen bool,
) Window {
	out := window{
		title:        title,
		width:        width,
		height:       height,
		isResizable:  isResizable,
		isFullscreen: isFullscreen,
	}

	return &out
}

// Title returns the title
func (obj *window) Title() string {
	return obj.title
}

// Width returns the width
func (obj *window) Width() uint {
	return obj.width
}

// Height returns the height
func (obj *window) Height() uint {
	return obj.height
}

// IsResizable returns true if the window is resizable, false otherwise
func (obj *window) IsResizable() bool {
	return obj.isResizable
}

// IsFullscreen returns true if the window is fullscreen, false otherwise
func (obj *window) IsFullscreen() bool {
	return obj.isFullscreen
}
