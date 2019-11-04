package secondAssignment

import (
  //"bytes"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"strconv"
  "time"
)

type WebhookArchive struct {
  Id          int
	Event     	string	`json:"event"`
  Url		      string	`json:"url"`
  Timestamp   time.Time
}
var webhooks []WebhookArchive

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Println("Received POST request...")
	//	for _, v := range webhooks {
		//	go CallUrl(v.Url, "Response on registered event in webhook demo: " + v.Event)
	//	}
	default:
		http.Error(w, "Invalid method "+r.Method, http.StatusBadRequest)
	}
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
  http.Header.Add(w.Header(), "content-type", "application/json")
	switch r.Method {
	case http.MethodPost:
		webhook := WebhookArchive{}
		err := json.NewDecoder(r.Body).Decode(&webhook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
    webhook.Id = len(webhooks)
    t := time.Now().UTC()
    webhook.Timestamp = t
		webhooks = append(webhooks, webhook)
		fmt.Fprintln(w, len(webhooks)-1)
		fmt.Println("Webhook " + webhook.Url + " has been registered.")

	case http.MethodGet:
		err := json.NewEncoder(w).Encode(webhooks)
		if err != nil {
			http.Error(w, "Something went wrong: " + err.Error(), http.StatusInternalServerError)
		}
	default: http.Error(w, "Invalid method " + r.Method, http.StatusBadRequest)
	}
}


/*
 	Calls given URL with given content and awaits response (status and body).
 */
