package secondAssignment

type userStorage interface {
	Init()
	Add(u User) error
	Count() int
	Get(key string) (User, bool)
	GetAll() []User
}

type User struct {
	Username						string `json:"username"`
	Count								int
}

type UsersDB struct {
	users map[string]User
}
func (db *UsersDB) Init() {
	db.users = make(map[string]User)
}
func (db *UsersDB) Add(u User) error {
	db.users[u.Username] = u
	return nil
}
func (db *UsersDB) Count() int {
	return len(db.users)
}
func (db *UsersDB) Get(Username  string) (User, bool) {
	s, ok := db.users[Username]
	return s, ok
}
func (db *UsersDB) GetAll() []User {
	all := make([]User, 0, db.Count())
	for _, s := range db.users {
		all = append(all, s)
	}
	return all
}
