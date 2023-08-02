package controller

import (
	"net/http"
	"time"

	"github.com/rherlt/reval/internal/api/evaluationapi"

	"github.com/gin-gonic/gin"
)

type EvaluationApiServerInterface struct {
	evaluationapi.ServerInterface
}

func GetSwagger(c *gin.Context) {

	swagger, error := evaluationapi.GetSwagger()

	if error != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.PureJSON(http.StatusOK, swagger)
}

var response1 = evaluationapi.GetEvaluationResponse{
	Id: 1,
	Response: evaluationapi.Message{
		From:    "fastchat/vicuna-7b",
		Subject: "",
		Body:    "Sehr geehrter [Name des Kunden],\n\nvielen Dank für Ihre freundliche E-Mail und das große Lob für unsere Zahlungen und deren Umfang. Es freut uns sehr, dass Sie mit unserem Service zufrieden sind und wir Ihre Erwartungen erfüllen können.\n\nAls langjähriger Kunde ist es uns besonders wichtig, Ihnen eine zuverlässige und professionelle Energieversorgung zu bieten. Daher sind wir sehr erfreut darüber, dass Sie unsere pünktlichen und vollständigen Zahlungen schätzen.\n\nIhr Lob werden wir selbstverständlich gerne an unsere Kollegen weitergeben, die an der Abwicklung Ihrer Zahlungen beteiligt sind. Es motiviert uns sehr, wenn unsere Arbeit von unseren Kunden anerkannt wird.\n\nWir möchten Ihnen ebenfalls für Ihre langjährige Treue danken. Es ist uns eine Freude, Sie als Kunden zu haben und Ihnen weiterhin einen ausgezeichneten Service bieten zu können.\n\nWenn Sie weitere Fragen oder Anliegen haben, stehen wir Ihnen natürlich jederzeit gerne zur Verfügung.\n\nMit freundlichen Grüßen\n\n[Dein Name]\nKundenbetreuer Energieversorgungsunternehmen",
		Date:    "2023-08-02T08:18:34.000Z",
	},
	Evaluations: evaluationapi.Evaluations{
		NumNegative: 0,
		NumNeutral:  1,
		NumPositive: 1,
	},
	Request: evaluationapi.Message{
		From:    "Max Mustermann <max@mustermann.de>",
		Subject: "Lob für Ihre Zahlung und deren Umfang",
		Body:    "Sehr geehrte Damen und Herren,\n\nich möchte mich heute bei Ihnen für Ihre vorbildliche Zahlung und deren Umfang bedanken. Es ist erfreulich zu sehen, dass Sie Ihre Rechnungen stets pünktlich und vollständig begleichen.\n\nAls langjähriger Kunde Ihres Energieversorgungsunternehmens schätze ich Ihre Zuverlässigkeit und Professionalität sehr. Es ist beruhigend zu wissen, dass ich mich auf Sie verlassen kann, wenn es um meine Energieversorgung geht.\n\nIch möchte Ihnen daher ein großes Lob aussprechen und mich für Ihre hervorragende Arbeit bedanken. Bitte geben Sie dieses Lob auch gerne an Ihre Kollegen weiter, die an der Abwicklung meiner Zahlungen beteiligt sind.\n\nMit freundlichen Grüßen\n\n[Name des Kunden]",
		Date:    "2023-08-02T08:12:34.000Z",
	},
}

var response2 = evaluationapi.GetEvaluationResponse{
	Id: 2,
	Response: evaluationapi.Message{
		From:    "fastchat/vicuna-7b",
		Subject: "",
		Body:    "Sehr geehrter zufriedener Kunde,\n\nvielen Dank für Ihre freundliche E-Mail und das positive Feedback zu unserem Service.\n\nEs freut uns sehr zu hören, dass der Tarifwechsel reibungslos und unkompliziert für Sie verlaufen ist. Unsere Mitarbeiter sind stets bemüht, Ihnen den bestmöglichen Service zu bieten und Sie während des gesamten Prozesses zu unterstützen. Es ist schön zu hören, dass Sie sich bei uns gut aufgehoben fühlen und mit unserer Strom- und Gasversorgung zufrieden sind.\n\nWir legen großen Wert darauf, unseren Kunden Flexibilität und Zuverlässigkeit zu bieten. Daher freut es uns besonders, dass unser neues Tarifmodell Ihnen diese Flexibilität ermöglicht und Sie weiterhin mit qualitativ hochwertiger Energie versorgt sind.\n\nWir bedanken uns für Ihr Vertrauen in unser Unternehmen und freuen uns, Sie auch in Zukunft als zufriedenen Kunden bei uns begrüßen zu dürfen. Bei weiteren Fragen oder Anliegen stehen wir Ihnen selbstverständlich jederzeit zur Verfügung.\n\nMit freundlichen Grüßen\n\nIhr Kundenservice",
		Date:    "2023-08-01T12:44:00.000Z",
	},
	Evaluations: evaluationapi.Evaluations{
		NumNegative: 2,
		NumNeutral:  1,
		NumPositive: 1,
	},
	Request: evaluationapi.Message{
		From:    "Max Mustermann <max@mustermann.de>",
		Subject: "Lob für den reibungslosen Tarifwechsel",
		Body:    "Sehr geehrte Damen und Herren,\n\nich möchte mich auf diesem Wege bei Ihnen für den reibungslosen Tarifwechsel bedanken.\n\nEs war für mich eine wichtige Entscheidung, die mit einem gewissen Maß an Unsicherheit verbunden war. Jedoch wurde ich in jeder Hinsicht bestmöglich von Ihnen beraten und unterstützt. Obwohl ich mich für einen völlig neuen Tarif entschieden habe, verlief der Wechsel so unkompliziert und schnell, dass ich mich wieder für Sie als meinen Energieanbieter entscheiden würde.\n\nDas neue Tarifmodell bietet mir nun viel mehr Flexibilität bei gleichbleibender Zuverlässigkeit sowie qualitative Strom- und Gasversorgung. Ich fühle mich von Ihnen mehr als nur gut betreut und würde mich jederzeit wieder für Ihre Dienstleistungen entscheiden.\n\nVielen Dank für Ihren hervorragenden Service!\n\nMit freundlichen Grüßen,\n\nIhr zufriedener Kunde",
		Date:    "2023-08-01T12:00:00.000Z",
	},
}

func (si EvaluationApiServerInterface) GetServerOptions() evaluationapi.GinServerOptions {
	return evaluationapi.GinServerOptions{
		BaseURL: "/api/",
	}
}

func (si EvaluationApiServerInterface) GetEvaluation(c *gin.Context, params evaluationapi.GetEvaluationParams) {

	//find out wich response to return (by random)
	var response = response1
	if time.Now().Second()%2 == 1 {
		response = response2
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (si EvaluationApiServerInterface) PostEvaluation(c *gin.Context, params evaluationapi.PostEvaluationParams) {

	requestBody := new(evaluationapi.PostEvaluationRequest)

	if err := c.BindJSON(&requestBody); err != nil {
		var ei = evaluationapi.ErrorInformation{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, ei)
		return
	}

	var response *evaluationapi.GetEvaluationResponse

	if requestBody.Id == response1.Id {
		response = &response1
	} else if requestBody.Id == response2.Id {
		response = &response2
	} else {
		c.Status(http.StatusNotFound)
		return
	}

	switch requestBody.EvaluationResult {
	case evaluationapi.Negative:
		response.Evaluations.NumNegative++
	case evaluationapi.Positive:
		response.Evaluations.NumPositive++
	case evaluationapi.Neutral:
		response.Evaluations.NumNeutral++
	}

	c.Status(http.StatusOK)
}
