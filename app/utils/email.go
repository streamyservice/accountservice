package utils

import (
	"accountservice/app/domain/dao"
	"bytes"
	"crypto/tls"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *dao.User, data *EmailData) {

	// Sender data.
	from := os.Getenv("EMAIL_FROM")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")
	to := "mitchelketcha@gmail.com"
	smtpHost := os.Getenv("SMTP_HOST")
	//smtpPort := os.Getenv("SMTP_PORT")

	var body bytes.Buffer

	path, err := filepath.Abs("app/templates")
	if err != nil {
		log.Fatal(err)
	}
	template, err := template.ParseGlob(filepath.Join(path, "*.html"))
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	//template, err := ParseTemplateDir("templates")
	//if err != nil {
	//	log.Fatal("Could not parse template", err)
	//}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, 587, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}

func VerifyOTP(userCode string, savedCode string) bool {

	return strings.Compare(userCode, savedCode) == 0
}
