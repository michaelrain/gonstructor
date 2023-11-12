package response

type Message struct {
	Text        string
	Buttons     [][]Button
	Attachments []Attachments
}

type Response struct {
	Messages []Message
}

type Button struct {
	Code    string
	Display string
	Target  ButtonTarget
}

type ButtonTarget struct {
	Type     string
	Resource string
}

type Attachments struct {
	URL string
}
