package Mail

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"rental_easy.in/m/pkg/utils"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(filepath.Join(`/home/suryatejap/Documents/Go Lang Practice/RENTAL EASY/pkg/Mail/`, "mail.env"))
	utils.CheckErr(err)

	return os.Getenv(key)
}

func BookingConfirmationMail(email, item_name, start_date, end_date string, e chan error) {

	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"

	senderEmail := GoDotEnvVariable("SENDEREMAIL")

	mail := fmt.Sprintf(`{
	    "personalizations": [
	        {
	            "to": [
	                {
	                    "email": `+`"%s"`+`
	                }
	            ],
	            "subject": "Booking Details"
	        }
	    ],
	    "from": {
	        "email": `+`"%s"`+`
	    },
	    "content": [
	        {
	            "type": "text/plain",
	            "value":`+` "Your Booking is Successful.\n The Item booked is `+fmt.Sprint(item_name)+` is successful.\n The Starting Date of Your Booking is : `+fmt.Sprint(start_date)+`.\n The Ending Date of your Booking is : `+fmt.Sprint(end_date)+`.\n"`+`
			}
	    ]
	}`, email, senderEmail)

	payload := strings.NewReader(mail)
	req, err := http.NewRequest("POST", url, payload)
	utils.CheckErr(err)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", GoDotEnvVariable("RAPIDAPIKEY"))
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	utils.CheckErr(err)

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	utils.CheckErr(err)

	e <- err

}
