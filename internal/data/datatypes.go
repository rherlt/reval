package data

type Evaluations struct {
	Rating   Rating  `json:"rating"`
	Id       string  `json:"id"`
	Request  Message `json:"request"`
	Response Message `json:"response"`
}

type Rating struct {
	From  string `json:"from"`
	Value string `json:"value"`
	Date  string `json:"date"`
}

type Message struct {
	Body    string `json:"body"`
	Date    string `json:"date"`
	From    string `json:"from"`
	Subject string `json:"subject"`
}
