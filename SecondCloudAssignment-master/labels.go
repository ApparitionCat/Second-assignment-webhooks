package secondAssignment

type labelsStorage interface {
	Init()
	Add(l Label) error
	Count() int
	Get(key string) (Label, bool)
	GetAll() []Label
}


type Label struct {
	Label						string `json:"name"`
	Count						int
}




type LabelsDB struct {
	labels map[string]Label
}

func (db *LabelsDB) Init() {
	db.labels = make(map[string]Label)
}

func (db *LabelsDB) Add(l Label) error {
	db.labels[l.Label] = l
	return nil
}

func (db *LabelsDB) Count() int {
	return len(db.labels)
}

func (db *LabelsDB) Get(Label  string) (Label, bool) {
	s, ok := db.labels[Label]
	return s, ok
}

func (db *LabelsDB) GetAll() []Label {
	all := make([]Label, 0, db.Count())
	for _, s := range db.labels {
		all = append(all, s)
	}
	return all
}
