package model

const (
	START_TAUNT     = ""
	PRIMARY_COLOR   = "#75CEDD"
	NAME            = "Snakey"
	SECONDARY_COLOR = "#F7D3A2"
)

type StartResponse struct {
	Color          string   `json:"color,omitempty"`
	Name           string   `json:"name,omitempty"`
	HeadURL        string   `json:"head_url,omitempty"`
	Taunt          string   `json:"taunt,omitempty"`
	HeadType       HeadType `json:"head_type,omitempty"`
	TailType       TailType `json:"tail_type,omitempty"`
	SecondaryColor string   `json:"secondary_color,omitempty"`
}
