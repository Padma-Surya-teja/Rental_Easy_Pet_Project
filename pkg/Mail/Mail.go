package Mail

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Mail(email, item_name, start_date, end_date string) {

	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"

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
	        "email": "sender email"
	    },
	    "content": [
	        {
	            "type": "text/plain",
	            "value":`+` "Your Booking is Successful.\n The Item booked is `+fmt.Sprint(item_name)+` is successful.\n The Starting Date of Your Booking is : `+fmt.Sprint(start_date)+`.\n The Ending Date of your Booking is : `+fmt.Sprint(end_date)+`.\n"`+`
			}
	    ]
	}`, email)

	payload := strings.NewReader(mail)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "Enter Rapid-API-Key here")
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
