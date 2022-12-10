package proxy

// The Image class implements an interface to show image files.
// Image overrides HandleMouse to allow the user to resize the image interactively.
type Image struct {
	Graphic // The Graphic object that this Image object is a proxy for
}

// ImageInterface Interface of Image
type ImageInterface interface {
	Load(filename string)
	HandleMouse(event *MouseEvent)
}

// Load loads the Image object.
func (image *Image) Load(filename string) {
	// Load the image file
	// ...
	// Create the Graphic object
	image.Graphic = NewGraphic()
	// Set the Graphic object's state
	// ...
}

// HandleMouse handles mouse events.
func (image *Image) HandleMouse(event *MouseEvent) {
	// Handle mouse events
	// ...
}
