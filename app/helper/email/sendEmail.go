package email

import (
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

// EmailConfig holds the configuration for sending emails using SMTP.
type EmailConfig struct {
	SMTPHost    string // SMTP server host (e.g., smtp.example.com)
	SMTPPort    string // SMTP server port (e.g., 587)
	SenderEmail string // The email address that will be used to send emails
	SenderPass  string // The password for the sender's email address
}

// EmailHelper provides methods for sending emails using the SMTP protocol.
type EmailHelper struct {
	config EmailConfig
}

// NewEmailHelper initializes a new EmailHelper by reading the configuration from a JSON file using Viper.
// configPath: Path to the configuration file (e.g., "./config.json").
// Returns an instance of EmailHelper.
func NewEmailHelper(configPath string) (*EmailHelper, error) {
	// Set up Viper to read from the config file
	viper.SetConfigFile(configPath)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Load the email configuration
	config := EmailConfig{
		SMTPHost:    viper.GetString("email.smtp_host"),
		SMTPPort:    viper.GetString("email.smtp_port"),
		SenderEmail: viper.GetString("email.sender_email"),
		SenderPass:  viper.GetString("email.sender_pass"),
	}

	return &EmailHelper{config: config}, nil
}

// SendEmail sends an email to the specified recipient with the given subject and body.
// to: Recipient's email address.
// subject: Subject of the email.
// body: Body content of the email.
// Returns an error if the email could not be sent.
func (e *EmailHelper) SendEmail(to string, subject string, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", e.config.SenderEmail, e.config.SenderPass, e.config.SMTPHost)

	// Create the email message.
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" + body + "\r\n")

	// Attempt to send the email.
	err := smtp.SendMail(e.config.SMTPHost+":"+e.config.SMTPPort, auth, e.config.SenderEmail, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
