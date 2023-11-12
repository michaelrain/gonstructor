package state

type State struct {
	UserID        string                 `json:"user_id"`
	Command       string                 `json:"command"`
	Message       string                 `json:"message"`
	Source        string                 `json:"source"`
	ModuleContext map[string]interface{} `json:"module_context"`
	CapturedBy    string                 `json:"caputred_by"`
	DataBag       map[string]interface{} `json:"data_bag"`
}
