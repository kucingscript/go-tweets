package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/kucingscript/go-tweets/internal/config"
	"gopkg.in/gomail.v2"
)

var templateFS embed.FS

type Mailer struct {
	dialer *gomail.Dialer
	sender string
}

func NewMailer(cfg *config.Config) *Mailer {
	dialer := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPass)

	return &Mailer{
		dialer: dialer,
		sender: cfg.SMTPSender,
	}
}

func (m *Mailer) Send(recipient string, templateFile string, data interface{}) error {
	tmpl, err := template.New("email").ParseFS(templateFS, templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse email template '%s': %w", templateFile, err)
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return fmt.Errorf("failed to execute subject template: %w", err)
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return fmt.Errorf("failed to execute plainBody template: %w", err)
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return fmt.Errorf("failed to execute htmlBody template: %w", err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
