package proxy

// TextDocument Finally, we have the TextDocument class, which contains Graphic objects:
type TextDocument struct {
	Graphic []Graphic
}

// NewTextDocument The TextDocument class has a constructor that initializes the Graphic array:
func NewTextDocument() *TextDocument {
	return &TextDocument{Graphic: []Graphic{}}
}

// AddGraphic The TextDocument class has a method that adds a Graphic object to the Graphic array:
func (textDocument *TextDocument) AddGraphic(graphic Graphic) {
	textDocument.Graphic = append(textDocument.Graphic, graphic)
}

// Draw The TextDocument class has a method that draws all the Graphic objects in the Graphic array:
func (textDocument *TextDocument) Draw() {
	for _, graphic := range textDocument.Graphic {
		graphic.Draw()
	}
}

// HandleMouse The TextDocument class has a method that handles mouse events for all the Graphic objects in the Graphic array:
func (textDocument *TextDocument) HandleMouse(event *MouseEvent) {
	for _, graphic := range textDocument.Graphic {
		graphic.HandleMouse(event)
	}
}

// Load The TextDocument class has a method that loads all the Graphic objects in the Graphic array:
func (textDocument *TextDocument) Load(filename string) {
	for _, graphic := range textDocument.Graphic {
		graphic.Load(filename)
	}
}

// SetHeight The TextDocument class has a method that sets the height of all the Graphic objects in the Graphic array:
func (textDocument *TextDocument) SetHeight(height int) {
	for _, graphic := range textDocument.Graphic {
		graphic.SetHeight(height)
	}
}

// SetWidth The TextDocument class has a method that sets the width of all the Graphic objects in the Graphic array:
func (textDocument *TextDocument) SetWidth(width int) {
	for _, graphic := range textDocument.Graphic {
		graphic.SetWidth(width)
	}
}

// SetX The TextDocument class has a method that sets the x-coordinate of all the Graphic objects in the Graphic array:
func (textDocument *TextDocument) SetX(x int) {
	for _, graphic := range textDocument.Graphic {
		graphic.SetX(x)
	}
}

// SetY The TextDocument class has a method that sets the y-coordinate of all the Graphic objects in the Graphic array:
func (textDocument *TextDocument) SetY(y int) {
	for _, graphic := range textDocument.Graphic {
		graphic.SetY(y)
	}
}

// AddImageProxy - add ImageProxy to TextDocument
func (textDocument *TextDocument) AddImageProxy(imageProxy ImageProxyInterface) {
	textDocument.Graphic = append(textDocument.Graphic, imageProxy)
}
