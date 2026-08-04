package confirm_email_service

import (
	"errors"
	"fmt"
	"github.com/resend/resend-go/v3"
	"log"
	"os"
	"strings"
)

func SendEmail(to, confirmURL string) error {
	htmlBytes, err := os.ReadFile("templates/confirm-email-template.html")

	if err != nil {
		return err
	}

	html := strings.ReplaceAll(string(htmlBytes), "{{confirm_url}}", confirmURL)

	resendApiKey := os.Getenv("RESEND_API_KEY")
	if resendApiKey == "" {
		return errors.New("Please set RESEND_API_KEY environment variable")
	}

	from := os.Getenv("EMAIL_FROM")
	if from == "" {
		return errors.New("Please set EMAIL_FROM environment variable")
	}

	client := resend.NewClient(resendApiKey)

	toSlice := make([]string, 1)
	toSlice[0] = to

	params := &resend.SendEmailRequest{
		From:    from,
		To:      toSlice,
		Subject: "Confirm your email | CoFounders Match",
		Html:    html,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	log.Printf("Email sent: %v", sent)

	return nil
}
