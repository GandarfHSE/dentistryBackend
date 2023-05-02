package image

type Image struct {
	Ext  string `json:"ext"`
	Data []byte `json:"data"`
}
