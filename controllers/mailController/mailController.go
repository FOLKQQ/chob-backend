package mailController

import (
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	mail "gopkg.in/gomail.v2"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	from := os.Getenv("MAILER_USERNAME")
	pass := os.Getenv("MAILER_PASSWORD")

	//get json data from request body
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	m := mail.NewMessage()

	m.SetHeader("From", from)

	m.SetHeader("To", data["email"].(string))

	//m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")

	m.SetHeader("Subject", data["subject"].(string))

	m.SetBody("text/html", data["message"].(string))

	//m.Attach("lolcat.jpg")

	d := mail.NewDialer("smtp.gmail.com", 587, from, pass)

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}

}
