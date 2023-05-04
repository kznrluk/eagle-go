package eagle

// ItemData holds the information about an item.
type ItemData struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Size             int       `json:"size"`
	Ext              string    `json:"ext"`
	Tags             []string  `json:"tags"`
	Folders          []string  `json:"folders"`
	IsDeleted        bool      `json:"isDeleted"`
	URL              string    `json:"url"`
	Annotation       string    `json:"annotation"`
	ModificationTime int64     `json:"modificationTime"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	NoThumbnail      bool      `json:"noThumbnail"`
	LastModified     int64     `json:"lastModified"`
	Palettes         []Palette `json:"palettes"`
	Star             int       `json:"star,omitempty"`
}

type Palette struct {
	Color []int   `json:"color"`
	Ratio float32 `json:"ratio"`
}
