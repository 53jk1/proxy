package proxy

// The Save operation saves the cached image extent and image file along with the name to the pipeline.
type Save struct {
	// The name of the image file.
	Filename string

	// The image extent.
	Extent interface{}

	// The image file.
	Image interface{}

	// The name of the image.
	Name string
}

// The Load reads this information and initializes it to the corresponding members.
func (save *Save) Load(filename string, extent interface{}, image interface{}, name string) {
	save.Filename = filename
	save.Extent = extent
	save.Image = image
	save.Name = name
}
