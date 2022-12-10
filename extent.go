package proxy

type Extent struct {
	Width, Height int
}

// The GetExtent implementation returns a cached extent if possible; otherwise, the image is loaded from a file. `Draw` loads the image, `HandleMouse` redirects events to the real image.
func (image *ImageProxy) GetExtent() interface{} {
	if image._extent == nil {
		image.Load(image.filename)
	}
	return image._extent
}
