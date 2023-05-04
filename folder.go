package eagle

type FolderData struct {
	ID                   string       `json:"id"`
	Name                 string       `json:"name"`
	Description          string       `json:"description"`
	Children             []FolderData `json:"children"`
	ModificationTime     int64        `json:"modificationTime"`
	Tags                 []string     `json:"tags"`
	ImageCount           int          `json:"imageCount"`
	DescendantImageCount int          `json:"descendantImageCount"`
	Pinyin               string       `json:"pinyin"`
	ExtendTags           []string     `json:"extendTags"`
	Size                 int          `json:"size"`
	VSType               string       `json:"vstype"`
	Styles               struct {
		Depth int  `json:"depth"`
		First bool `json:"first"`
		Last  bool `json:"last"`
	} `json:"styles"`
	IsVisible bool   `json:"isVisible"`
	HashKey   string `json:"$$hashKey"`
	Editable  bool   `json:"editable"`
}
