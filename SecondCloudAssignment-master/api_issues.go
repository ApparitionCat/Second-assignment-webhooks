package secondAssignment

import (
  "encoding/json"
	"net/http"
  //"strings"
)
type Issues struct {
//	labels		labelInfo `json:"labels"`
Repository						string `json:"title"`
 AuthName      authName    `json:"author"`
}

type labelInfo struct {
	labels		    string    `json:"labels"`
}

type authName struct {
  id            int       `json:"id"`
	username	  	string    `json:"username"`
}

var issueStructure[] Issues
var DBu = UsersDB{}
var DBl = LabelsDB{}

func replyWithAlll(w http.ResponseWriter, DB labelsStorage, auth string){
  url := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/labels?private_token=" + auth
	resp, err := http.Get(url)								
	if err != nil {												//these are everywhere, you know what they do		
  	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
	}
  defer resp.Body.Close()
	var tempLabel[] Label
  json.NewDecoder(resp.Body).Decode(&tempLabel)
  lurl := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/issues?private_token=" + auth
  resp, err = http.Get(lurl)								 
  if err != nil {										
    http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
  }
  defer resp.Body.Close()
  json.NewDecoder(resp.Body).Decode(&issueStructure)
  for idx, x := range tempLabel{     // For each occurrence
    if idx == 0 {														// Skip header
      println("")
    }
    DBl.Add(x)
  }
	a := make([]Label, 0, DBl.Count())		 //identical to the code on line 64
	for _, s := range DBl.GetAll() {	
		a = append(a, s)								
	}
	json.NewEncoder(w).Encode(a)					
}

func replyWithAllu(w http.ResponseWriter, DB userStorage, auth string){                           
  url := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/members/all?private_token=" + auth
	resp, err := http.Get(url)								
	if err != nil {														// gets url, gives a not found error if its not found
  	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
	}
  defer resp.Body.Close()
	var tempUser[] User
  json.NewDecoder(resp.Body).Decode(&tempUser)

  iurl := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/issues?private_token=" + auth
  resp, err = http.Get(iurl)								  // gets url, gives a not found error if its not found
  if err != nil {														
    http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
  }
  defer resp.Body.Close()
  json.NewDecoder(resp.Body).Decode(&issueStructure)
  for idx, x := range tempUser{     // For each occurrence
    if idx == 0 {														// this will Skip the header
      println("")
    }
    println(x.Username)                                             //prints username, should print all usernames 
    DBu.Add(x)
  }
  for idx, y := range issueStructure{
    if idx == 0 {														// again, skips header
      println("")
    }
    println(y.AuthName.username)                                        //prints authenticator name
  }
	a := make([]User, 0, DBu.Count())		  // variable map used to display
	for _, s := range DB.GetAll() {				// loops for each object
		a = append(a, s)										// and appends them to a list
	}
	json.NewEncoder(w).Encode(a)					// thats displayed
}
type findProject struct{
  Event     	string	`json:"event"`                     //struct 
}

func HandlerIssues(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
  var issueType string = r.URL.Query().Get("type")
  var issueAuth string = r.URL.Query().Get("auth")

  switch r.Method {                                                  
	case http.MethodPost:
		var myProject findProject
		err := json.NewDecoder(r.Body).Decode(&myProject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

  }

  if(issueType == "user"){
    replyWithAllu(w, &DBu, issueAuth)
  }
  if(issueType == "labels"){
	replyWithAlll(w, &DBl, issueAuth)
  }
}
