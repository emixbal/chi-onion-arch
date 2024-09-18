package email

import (
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

// GmailConfig holds the configuration for sending emails using Gmail SMTP.
type GmailConfig struct {
	SMTPHost       string // SMTP server host (e.g., smtp.gmail.com)
	SMTPPort       string // SMTP server port (e.g., 587)
	SenderUsername string // The email address that will be used to send emails
	SenderPass     string // The password for the sender's email address
}

// GmailHelper provides methods for sending emails using the SMTP protocol.
type GmailHelper struct {
	config GmailConfig
}

// NewGmailHelper initializes a new GmailHelper by reading the configuration from Viper.
// Returns an instance of GmailHelper.
func NewGmailHelper() *GmailHelper {
	config := GmailConfig{
		SMTPHost:       viper.GetString("emailGmail.smtp_host"),
		SMTPPort:       viper.GetString("emailGmail.smtp_port"),
		SenderUsername: viper.GetString("emailGmail.sender_email"),
		SenderPass:     viper.GetString("emailGmail.sender_pass"),
	}

	return &GmailHelper{config: config}
}

// SendEmail sends an email to the specified recipient with the given subject and body.
// to: Recipient's email address.
// subject: Subject of the email.
// body: Body content of the email.
// from: Address to display as the sender.
// Returns an error if the email could not be sent.
func (e *GmailHelper) SendEmail(to string, subject string, body string, from string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", e.config.SenderUsername, e.config.SenderPass, e.config.SMTPHost)

	// Create the email message.
	msg := []byte(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nContent-Type: text/plain; charset=\"UTF-8\"\n\n%s", from, to, subject, body))

	// Attempt to send the email.
	err := smtp.SendMail(e.config.SMTPHost+":"+e.config.SMTPPort, auth, e.config.SenderUsername, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
