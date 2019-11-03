package secondAssignment

import (
  "encoding/json"
	"net/http"
  //"strings"
  //"bytes"
  "strconv"
)
var DBp = ProjectsDB{}											// variable used to store objects of a class


func replyWithAlls(w http.ResponseWriter, DB projectStorage, limit string, auth string){  

  limitINT, err := strconv.Atoi(limit)
	if err == nil {                               
	}
  for i := 1; i <= limitINT; i++ {
    if i == 0 {
      continue
    }

    url := "https://git.gvk.idi.ntnu.no/api/v4/projects/" + strconv.Itoa(i) + "?private_token=" + auth
  	resp, err := http.Get(url)								// retrieves the url, says it on the code aswell
  	if err != nil {														// error handler, returns not found error if the url isnt found
  		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
      ST.TestApi("Gitlab")
  	}
    defer resp.Body.Close()
  	var tempProject Project
    json.NewDecoder(resp.Body).Decode(&tempProject)

    url = "https://git.gvk.idi.ntnu.no/api/v4/projects/" + strconv.Itoa(i) + "/repository/commits?per_page=900&private_token=" + auth
    resp, err = http.Get(url)							
    if err != nil {													//pretty much the same as above
      http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
      ST.TestApi("Gitlab")
    }
    defer resp.Body.Close()
    var tempCommit[] Commit
    json.NewDecoder(resp.Body).Decode(&tempCommit)

    tempProject.Commits = len(tempCommit)
    if(tempProject.Repository != "") { DB.Add(tempProject) }
  }

	a := make([]Project, 0, DB.Count())		// makes a varable map to be printed 
	for _, s := range DB.GetAll() {				// for each DB
		a = append(a, s)										// append them to 'a'
	}
	json.NewEncoder(w).Encode(a)					// Display on browser
}








func HandlerCommits(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
//	parts := strings.Split(r.URL.Path, "/")
  var limit string = r.URL.Query().Get("limit")
  var commitAuth string = r.URL.Query().Get("auth")

  if limit == ""{
    limit = "5"
  }


	replyWithAlls(w, &DBp, limit, commitAuth)
}
