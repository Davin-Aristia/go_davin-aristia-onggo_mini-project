package template

import (
    "go-mini-project/model"

	"html/template"
	"bytes"
    "fmt"
    "strings"
)

func numberFormat(value float64) string {
    // Format the number with a comma as the thousand separator and two decimal places
    return fmt.Sprintf("Rp %s", formatFloat(value))
}

func formatFloat(value float64) string {
    parts := strings.Split(fmt.Sprintf("%.2f", value), ".")
    integerPart := parts[0]
    decimalPart := parts[1]

    var result string
    for i := len(integerPart) - 1; i >= 0; i-- {
        result = string(integerPart[i]) + result
        if (len(integerPart)-i)%3 == 0 && i != 0 {
            result = "," + result
        }
    }

    return result + "." + decimalPart
}

func RenderCheckoutTemplate(invoice, date string, salesDetails []model.SalesDetail, total float64) (string, error) {
    const checkoutTemplate = `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Order Confirmation</title>
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
            <h1>Order Confirmation</h1>
            <p>Hello,</p>
            <p>Your order (Invoice: {{.Invoice}}) has been successfully placed on {{.Date}}.</p>
            <p>Order details:</p>
            <ul>
                {{range .SalesDetails}}
                    <li>{{.Quantity}}x {{.BookId}} - {{numberFormat .Price}}</li>
                {{end}}
            </ul>
            <p>Total: {{numberFormat .Total}}</p>
            <p>If you have any questions or need assistance, please don't hesitate to contact us.</p>
            <div class="notification">
                <p>Thank you for shopping with us!</p>
            </div>
        </div>
        <div class="footer">
            <p>Best regards, Book Store API Team</p>
        </div>
    </body>
    </html>
    `

    var emailBodyContent bytes.Buffer
    data := struct {
        Invoice      string
        Date         string
        SalesDetails []model.SalesDetail
        Total        float64
    }{
        Invoice:      invoice,
        Date:         date,
        SalesDetails: salesDetails,
        Total: total,
    }

    funcMap := template.FuncMap{
        "numberFormat": numberFormat,
    }
    
    tmpl, err := template.New("checkoutTemplate").Funcs(funcMap).Parse(checkoutTemplate)
    if err != nil {
        return "", err
    }

    err = tmpl.Execute(&emailBodyContent, data)
    if err != nil {
        return "", err
    }

    return emailBodyContent.String(), nil

}