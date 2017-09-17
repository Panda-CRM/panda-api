package helpers

import (
	"encoding/base64"
	"fmt"
	"github.com/matcornic/hermes"
	"log"
	"net/smtp"
	"os"
	"time"
)

const (
	ENV_EMAIL_SMTP     = "EMAIL_SMTP"
	ENV_EMAIL_PORT     = "EMAIL_PORT"
	ENV_EMAIL_FROM     = "EMAIL_FROM"
	ENV_EMAIL_USER     = "EMAIL_USER"
	ENV_EMAIL_PASSWORD = "EMAIL_PASSWORD"
)

var SMTP_SERVER SMTPServer

type SMTPServer struct {
	SMTP     string
	Port     string
	From     string
	User     string
	Password string
}

func (s *SMTPServer) address() string {
	return s.SMTP + ":" + s.Port
}

func (s *SMTPServer) auth() smtp.Auth {
	return smtp.PlainAuth("", s.User, s.Password, s.SMTP)
}

func (s *SMTPServer) SendEmail(email EmailInterface) error {
	if err := smtp.SendMail(s.address(), s.auth(), s.From, email.GetMailingList(), email.Message()); err != nil {
		return err
	}
	return nil
}

func init() {
	getSMTPServerConfig()
}

func getSMTPServerConfig() {
	log.Print("[CONFIG] Lendo configurações de e-mail")
	emailSMTP := os.Getenv(ENV_EMAIL_SMTP)
	emailPort := os.Getenv(ENV_EMAIL_PORT)
	emailFrom := os.Getenv(ENV_EMAIL_FROM)
	emailUser := os.Getenv(ENV_EMAIL_USER)
	emailPassword := os.Getenv(ENV_EMAIL_PASSWORD)
	if len(emailSMTP) > 0 {
		SMTP_SERVER.SMTP = emailSMTP
	}
	if len(emailPort) > 0 {
		SMTP_SERVER.Port = emailPort
	}
	if len(emailFrom) > 0 {
		SMTP_SERVER.From = emailFrom
	}
	if len(emailPassword) > 0 {
		SMTP_SERVER.User = emailUser
	}
	if len(emailPassword) > 0 {
		SMTP_SERVER.Password = emailPassword
	}
}

/*func Example() {
	var email resetEmail
	email.UserName = "Wilson Tamarozzi"
	email.ToEmails = []string{"wilson.tamarozzi@gmail.com"}
	email.ResetLink = "https://hermes-example.com/resetEmail-password?token=d9729feb74992cc3482b350163a1a010"
	SMTP_SERVER.SendEmail(&email)
}*/

type EmailInterface interface {
	Message() []byte
	GetMailingList() []string
}

type baseEmail struct {
	ToEmails        []string
	UserName        string
	CompanyNameUser string
	TitleMessage    string
	EmailMessage    hermes.Email
}

func (b *baseEmail) makeEmail() []byte {
	header := make(map[string]string)
	header["Subject"] = b.TitleMessage
	header["MIME-version"] = "1.0"
	header["Content-Type"] = `text/html; charset="UTF-8"`
	header["Content-Transfer-Encoding"] = "base64"
	header["X-Entity-Ref-ID"] = time.Now().String()

	var message string
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(b.makeTemplateHTML()))
	return []byte(message)
}

func (b *baseEmail) makeTemplateHTML() string {
	h := hermes.Hermes{
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			Name:        "Equipe Panda",
			Link:        "https://api.panda.com/",
			Logo:        "https://github.com/Panda-CRM/panda-web-angular/blob/master/img/logo-128x128.png?raw=true",
			Copyright:   "Copyright © 2017 Panda. Todos direitos reservados.",
			TroubleText: "Se você estiver tendo problemas com o botão '{ACTION}', copie e cole o URL abaixo em seu navegador.",
		},
	}

	body, err := h.GenerateHTML(b.EmailMessage)
	if err != nil {
		panic(err)
	}
	return body
}

type ResetEmail struct {
	baseEmail
	ResetLink string
}

func (r *ResetEmail) build() {
	r.TitleMessage = "Recuperar Senha"
	r.EmailMessage = hermes.Email{
		Body: hermes.Body{
			Greeting: "Oi",
			Name:     r.UserName,
			Intros: []string{
				"Você recebeu este e-mail porque houve um pedido de recuperação de senha para a sua conta.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Clique no botão abaixo para redefinir sua senha:",
					Button: hermes.Button{
						Color: "#DC4D2F",
						Text:  "Alterar minha senha",
						Link:  r.ResetLink,
					},
				},
			},
			Outros: []string{
				"Caso não tenha solicitado a troca de senha, ignore este e-mail.",
			},
			Signature: "Obrigado",
		},
	}
}

func (r ResetEmail) Message() []byte {
	r.build()
	return r.makeEmail()
}

func (r ResetEmail) GetMailingList() []string {
	return r.ToEmails
}
