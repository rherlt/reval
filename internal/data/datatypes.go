package data

type Umbrella struct {
	Header  *Header  `json:"header,omitempty"`
	Entries *[]Entry `json:"entries,omitempty"`
}

type Header struct {
	Id           *string      `json:"id,omitempty"`
	Model        *string      `json:"model,omitempty"`
	Hyperparams  *Hyperparams `json:"hyperparams,omitempty"`
	SystemPrompt *string      `json:"systemPrompt,omitempty"`
	Date         *string      `json:"date,omitempty"`
}

type Hyperparams struct {
	Temperature *int `json:"temperature,omitempty"`
	MaxTokens   *int `json:"max_tokens,omitempty"`
}

type Entry struct {
	Rating   *Rating  `json:"rating,omitempty"`
	Id       *string  `json:"id,omitempty"`
	Request  *Message `json:"request,omitempty"`
	Response *Message `json:"response,omitempty"`
}

type Rating struct {
	From  *string `json:"from,omitempty"`
	Value *string `json:"value,omitempty"`
	Date  *string `json:"date,omitempty"`
}

type Message struct {
	Body    *string `json:"body,omitempty"`
	Subject *string `json:"subject,omitempty"`
}
