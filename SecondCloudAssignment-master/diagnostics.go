package secondAssignment

import (
  "time"
  "net/http"
  "encoding/json"
  "strings"
)

type statusStorage interface {
	Init()
  Get() (Status, error)
  CheckIfWorks(api string)
  GetAll() []Status
}

type Status struct {
  Gitlab             		int
  Database              int
  Uptime             		time.Duration
  Version             	string
}

var startTime time.Time                  // GLobal time variable for storing when service started
var ST = statusDB{}                        // Map vaiable for fiagnostics

type statusDB struct {                     // Diagnostics map stored in memory
	status map[int]Status
}

func (db *statusDB) Init() {               // Initialised for use
	db.status = make(map[int]Status)
  startTime = time.Now()                 // Stores application start Time
  var tempDiag Status                      // Temp to hold to be modified diagnostics value
  tempDiag = db.status[0]                  // Copies object
  tempDiag.Gitlab = http.StatusOK        // Assigns default start up values
  tempDiag.Database = http.StatusOK
  tempDiag.Version = "v1"
  db.status[0] = tempDiag                  // Places udated information into diag
}

func (db *statusDB) Get() (Status, bool){    // Get specific diagnostics
  s, ok := db.status[0]
	return s, ok
}

func (db *statusDB) TestApi(api string){   // Assigns 503 error code if api is not working
  var tempDiag Status
  tempDiag = db.status[0]
  if api == "Gitlab"{                   // For gitlab
    tempDiag.Gitlab = http.StatusServiceUnavailable
  }else if api == "Database"{            //for the database
    tempDiag.Database = http.StatusServiceUnavailable
  }

  db.status[0] = tempDiag
}

func (db *statusDB) GetAll() []Status {     // Fetchdes the diagnostics
  var tempDiag Status
  tempDiag = db.status[0]
  tempDiag.Uptime = time.Since(startTime) / time.Second
  db.status[0] = tempDiag
	all := make([]Status, 0, 1)
	for _, s := range db.status {
		all = append(all, s)
	}
	return all
}



                                        // Returns webservice on request
func printDiagnostics(w http.ResponseWriter) {
  a := make([]Status, 0, 1)
  for _, s := range ST.GetAll() {
    a = append(a, s)
  }
  json.NewEncoder(w).Encode(a)
}

func HandlerDiag(w http.ResponseWriter, r *http.Request) {
		http.Header.Add(w.Header(), "content-type", "application/json")

		parts := strings.Split(r.URL.Path, "/")

		if len(parts) == 6 || parts[1] == "conservation" {
			http.Error(w, "Bad request:", http.StatusBadRequest)
			return
		}

    printDiagnostics(w)
}
