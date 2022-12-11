package proxy

type ImageProxy struct {
	filename string
	_extent  interface{}
}

func (image *ImageProxy) Load(filename string) {
	image.filename = filename
}
