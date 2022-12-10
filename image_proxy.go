package proxy

// The ImageProxy has the same interface as the Image.
// The ImageProxy delegates all operations to the Image.
type ImageProxyInterface interface {
	Load(filename string)
	HandleMouse(event *MouseEvent)
	Draw()
	SetX(x int)
	SetY(y int)
	SetWidth(width int)
	SetHeight(height int)
}
