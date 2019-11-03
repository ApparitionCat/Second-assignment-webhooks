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

var startTime time.Time                  // this is a time variable, its used as a counter for how long the service has been operating
var ST = statusDB{}                        // variable for diagnostics
type statusDB struct {                     
	status map[int]Status
}

func (db *statusDB) Init() {             
	db.status = make(map[int]Status)
  startTime = time.Now()                 // starts counting time
  var tempDiag Status                      //holds diagnostic valyes
  tempDiag = db.status[0]                  
  tempDiag.Gitlab = http.StatusOK        // Assigns default start up values
  tempDiag.Database = http.StatusOK
  tempDiag.Version = "v1"
  db.status[0] = tempDiag                  // Places udated information into diag
}

func (db *statusDB) Get() (Status, bool){    // retrieves diagnostics
  s, ok := db.status[0]
	return s, ok
}

func (db *statusDB) TestApi(api string){   // this si a error handle if serivce is unavaliable, the 503 error basically
  var tempDiag Status
  tempDiag = db.status[0]
  if api == "Gitlab"{                   // For gitlab
    tempDiag.Gitlab = http.StatusServiceUnavailable
  }else if api == "Database"{            //for the database
    tempDiag.Database = http.StatusServiceUnavailable
  }

  db.status[0] = tempDiag
}

func (db *statusDB) GetAll() []Status {                                      // retrieves diags
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



                                                                    // should return webservice
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
