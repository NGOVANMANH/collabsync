package email

import (
	"bytes"
	"html/template"
)

var emailTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.Subject}}</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			color: #333;
			margin: 0;
			padding: 20px;
		}
		.container {
			max-width: 600px;
			margin: 0 auto;
			background-color: #fff;
			padding: 20px;
			border-radius: 5px;
			box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		}
		h1 {
			color: #333;
			font-size: 24px;
			margin-bottom: 20px;
		}
		p {
			font-size: 16px;
			line-height: 1.5;
			margin-bottom: 20px;
		}
		a.button {
			display: inline-block;
			padding: 12px 20px;
			background-color: #007BFF;
			color: #ffffff;
			text-decoration: none;
			border-radius: 4px;
			font-weight: bold;
		}
		.footer {
			font-size: 12px;
			color: #999;
			margin-top: 30px;
			text-align: center;
		}
	</style>
</head>
<body>
	<div class="container">
		<h1>Welcome, {{.UserName}}!</h1>
		<p>Thank you for registering with us. To complete your registration and activate your account, please verify your email address by clicking the button below:</p>
		<p style="text-align: center;">
			<a href="{{.VerificationURL}}" class="button">Verify My Account</a>
		</p>
		<p>If the button above doesn't work, you can copy and paste the following URL into your browser:</p>
		<p>{{.VerificationURL}}</p>
		<p>We’re excited to have you with us!</p>
		<p>Best regards,<br>Your Company Team</p>
		<div class="footer">
			If you did not create this account, please ignore this email.
		</div>
	</div>
</body>
</html>`

// EmailTemplateData holds the data for rendering the email template.
type EmailTemplateData struct {
	UserName        string
	Subject         string
	VerificationURL string
}

// NewEmailTemplateData creates a new EmailTemplateData instance.
func NewEmailTemplateData(userName, subject, verificationURL string) *EmailTemplateData {
	return &EmailTemplateData{
		UserName:        userName,
		Subject:         subject,
		VerificationURL: verificationURL,
	}
}

// GetEmailTemplate renders the email template with the given data.
func GetEmailTemplate(data *EmailTemplateData) (string, error) {
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
