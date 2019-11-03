package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"secondAssignment"
)
func main() {

	secondAssignment.DBp.Init()
	secondAssignment.DBu.Init()
	secondAssignment.DBl.Init()
	secondAssignment.ST.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

																						// Runs functions based on url typed
	http.HandleFunc("/repocheck/v1/commits/", secondAssignment.HandlerCommits)
	http.HandleFunc("/repocheck/v1/issues/", secondAssignment.HandlerIssues)
	http.HandleFunc("/repocheck/v1/status/", secondAssignment.HandlerDiag)
	http.HandleFunc("/repocheck/v1/webhooks/", secondAssignment.WebhookHandler)
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
