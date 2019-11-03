package secondAssignment              //this file holds most of the functions that will be called, as well as some structs. Think of it as a class file.

type projectStorage interface {
	Init()
	Add(c Project) error
	Count() int
	Get(key string) (Project, bool)
	GetAll() []Project
}
type Project struct {
	Repository						string `json:"path_with_namespace"`
	Commits								int
}
type Commit struct {
	cID 									string `json:"id"`
}
type ProjectsDB struct {
	projects map[string]Project
}
func (db *ProjectsDB) Init() {
	db.projects = make(map[string]Project)
}
func (db *ProjectsDB) Add(p Project) error {
	db.projects[p.Repository] = p
	return nil
}
func (db *ProjectsDB) Count() int {
	return len(db.projects)
}
func (db *ProjectsDB) Get(Repository  string) (Project, bool) {
	s, ok := db.projects[Repository]
	return s, ok
}
func (db *ProjectsDB) GetAll() []Project {
	all := make([]Project, 0, db.Count())
	for _, s := range db.projects {
		all = append(all, s)
	}
	return all
}
func stringExists(a string, list []string) bool{
	for _, b := range list {
		if b == a{
			return true
		}
	}
	return false
}
