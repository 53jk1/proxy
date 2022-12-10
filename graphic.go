package proxy

// The Graphic class defines an interface for a graphical object:
// a Graphic is a "drawable" object.  The Graphic interface is
// implemented by all graphical objects.  The Graphic interface
// provides methods for drawing the object, for changing its
// appearance, and for handling events.
type Graphic interface {
	// Draw draws the graphical object.
	Draw()
	// HandleMouse handles mouse events.
	HandleMouse(event *MouseEvent)
	// Load loads the graphical object.
	Load(filename string)
	// SetHeight sets the height of the graphical object.
	SetHeight(height int)
	// SetWidth sets the width of the graphical object.
	SetWidth(width int)
	// SetX sets the x-coordinate of the graphical object.
	SetX(x int)
	// SetY sets the y-coordinate of the graphical object.
	SetY(y int)

	// ... other methods omitted ...
}

// The GraphicState class defines the
// state of a Graphic object.
type GraphicState struct {
	// ... other fields omitted ...
}

func (g GraphicState) Draw() {
	//TODO implement me
	panic("implement me")
}

func (g GraphicState) HandleMouse(event *MouseEvent) {
	//TODO implement me
	panic("implement me")
}

func (g GraphicState) Load(filename string) {
	//TODO implement me
	panic("implement me")
}

func (g GraphicState) SetHeight(height int) {
	//TODO implement me
	panic("implement me")
}

func (g GraphicState) SetWidth(width int) {
	//TODO implement me
	panic("implement me")
}

func (g GraphicState) SetX(x int) {
	//TODO implement me
	panic("implement me")
}

func (g GraphicState) SetY(y int) {
	//TODO implement me
	panic("implement me")
}

func NewGraphic() *GraphicState {
	return &GraphicState{}
}
