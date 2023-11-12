package domain

type Client struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Avatar        string         `json:"avatar"`
	Note          string         `json:"note"`
	Emails        []Email        `json:"emails"`
	Phones        []Phone        `json:"phones"`
	AdditionalIds []AdditionalID `json:"additional_ids"`
}

type Email struct {
	Email    string `json:"email"`
	ClientID int    `json:"client_id"`
}

type Phone struct {
	Phone    string `json:"phone"`
	Type     string `json:"type"`
	ClientID int    `json:"client_id"`
}

type AdditionalID struct {
	Value    string `json:"value"`
	ClientID int    `json:"client_id"`
}
