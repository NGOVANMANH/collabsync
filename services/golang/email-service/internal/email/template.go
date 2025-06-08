package email

import (
	"bytes"
	"html/template"
)

// EmailTemplateData holds data to be injected into the email template.
type EmailTemplateData struct {
	RecipientName string
	Message       string
}

// ParseTemplate parses the given template string with the provided data.
func ParseTemplate(tmpl string, data EmailTemplateData) (string, error) {
	t, err := template.New("email").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Example usage:
// const tmpl = `<h1>Hello, {{.RecipientName}}</h1><p>{{.Message}}</p>`
// body, err := ParseTemplate(tmpl, EmailTemplateData{RecipientName: "John", Message: "Welcome!"})