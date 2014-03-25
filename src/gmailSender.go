//WITH SUPPORT FROM http://www.goinggo.net/2013/06/send-email-in-go-with-smtpsendmail.html

package main

import (
	"log"
	"net/smtp"
	"text/template"
	"bytes"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/binding"
)

type theEmail struct {
	TheMessage   string `form:"TheMessage"`
	TheSender    string `form:"TheSender"`
	TheSubject   string `form:"TheSubject"`
}

func main() {
	m := martini.Classic()

	// render html templates from templates directory
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
			r.HTML(200, "senderForm", "")
		})

	m.Post("/", binding.Form(theEmail{}), func(post theEmail, r render.Render) {

			sendEmail(post.TheMessage, post.TheSender, post.TheSubject)
			r.HTML(200, "senderForm", "")

		})

	m.Run()

}

func sendEmail(theMessage string, theSender string, theSubject string){
	log.Println(theSubject)

	parameters := &struct {
			From string
			To string
			Subject string
			Message string
		}{
		"EMAIL GOES HERE",
		theSender,
		theSubject,
		theMessage,
	}

	buffer := new(bytes.Buffer)

	template := template.Must(template.New("emailTemplate").Parse(_EmailScript()))
	template.Execute(buffer, parameters)

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"EMAIL GOES HERE",
		"PASSWORD GOES HERE",
		"SMTP SERVER GOES HERE",
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"SMTP SERVER GOES HERE : PORT",
		auth,
		"EMAIL GOES HERE",
		[]string{theSender},
		buffer.Bytes(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(theSender)
}

// _EmailScript returns a template for the email message to be sent
func _EmailScript() (script string) {
	return `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"

{{.Message}}`
}

