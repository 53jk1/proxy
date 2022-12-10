package proxy

type ImageProxy struct {
	filename string
	_extent  interface{}
}

func (image *ImageProxy) Load(filename string) {
	image.filename = filename
}

// NewImageProxy - the constructor writes a local copy of the filename that stores the image and initializes `_extent` and `_image`:

type Width int

type _extent interface {
	Width() Width
}

