package img

type Image struct {
	Ext  string `json:"ext"`
	Data []byte `json:"data"`
}

type ImageUploadResponse struct {
	ImageSource string `json:"imageSource"`
}
