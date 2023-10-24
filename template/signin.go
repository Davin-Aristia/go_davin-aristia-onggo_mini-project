package template

import (
	"html/template"
	"bytes"
)

func RenderSigninTemplate(date, name string) (string, error) {
    const signinTemplate = `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Signin Activity Notification</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                background-color: #f7f7f7;
                margin: 0;
                padding: 0;
            }

            .container {
                background-color: #ffffff;
                max-width: 600px;
                margin: 0 auto;
                padding: 20px;
                border-radius: 5px;
                box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            }

            h1 {
                color: #333;
            }

            p {
                font-size: 16px;
                line-height: 1.6;
                color: #555;
            }

            .notification {
                background-color: #f1f1f1;
                padding: 10px;
                margin-top: 20px;
            }

            .footer {
                background-color: #f1f1f1;
                padding: 10px;
                text-align: center;
            }

            .footer p {
                font-size: 14px;
                color: #777;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Signin Activity Notification</h1>
            <p>Hello, {{.Name}}</p>
            <p>Your Book Store API account was accessed on {{.Date}}.</p>
            <p>If this was you, you can safely ignore this message.</p>
            <p>If this wasn't you, please take immediate action to secure your account.</p>
            <div class="notification">
                <p>If you have any questions, please contact our support team.</p>
            </div>
        </div>
        <div class="footer">
            <p>Best regards, Book Store API Team</p>
        </div>
    </body>
    </html>

    `
    tmpl, err := template.New("signinTemplate").Parse(signinTemplate)
	if err != nil {
		return "", err
	}

	var emailBodyContent bytes.Buffer
	data := struct {
		Date string
		Name string
	}{
		Date: date,
		Name: name,
	}

	err = tmpl.Execute(&emailBodyContent, data)
	if err != nil {
		return "", err
	}

	return emailBodyContent.String(), nil

}