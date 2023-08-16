package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rherlt/reval/internal/api/evaluationapi"
)

type MailItem struct {
	RequestMail  string `json:"requestMail"`
	MailResponse string `json:"mailResponse"`
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("../../tmp/output.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	bytes, _ := io.ReadAll(jsonFile)

	// we initialize our Users array
	var mails []MailItem

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(bytes, &mails)

	var transformedMails []evaluationapi.GetEvaluationResponse

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(mails); i++ {
		var currentMail = mails[i]
		transformedMails = append(transformedMails, evaluationapi.GetEvaluationResponse{
			Id: string(i),
			Request: evaluationapi.Message{
				Body:    GetStringAfterInBetween(currentMail.RequestMail, "Betreff: ", "\n\n"),
				From:    "chatgpt/gpt-3.5-turbo",
				Subject: GetStringInBetween(currentMail.RequestMail, "Betreff: ", "\n"),
			},
			Response: evaluationapi.Message{
				Body:    currentMail.MailResponse,
				From:    "chatgpt/gpt-3.5-turbo",
				Subject: "",
			},
		})
	}

	outFile, _ := json.MarshalIndent(transformedMails, "", "  ")

	_ = os.WriteFile("../../tmp/reval-transformed.json", outFile, 0644)
}

// GetStringInBetween Returns empty string if no start string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	fin := s + e
	return str[s:fin]
}

// GetStringInBetween Returns empty string if no start string found
func GetStringAfterInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return str
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return str
	}
	fin := s + e + len(end)
	return str[fin:]
}
