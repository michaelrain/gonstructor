package components

type FileComponent struct {
	BaseComponent
	Media    string `json:"media"`
	FileText string `json:"fileText"`
}
