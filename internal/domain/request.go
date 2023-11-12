package domain

type Request struct {
	FromID  string
	Source  string
	Command string
}

func NewRequest(fromID string, source string, command string) Request {
	return Request{
		FromID:  fromID,
		Source:  source,
		Command: command,
	}
}
