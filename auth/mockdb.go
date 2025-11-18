package auth

type MockDB struct {
	users        map[string]User
	usersByEmail map[string]User
}

func NewMockDB() *MockDB {
	return &MockDB{
		users:        make(map[string]User),
		usersByEmail: make(map[string]User),
	}
}

func (db *MockDB) Insert(u User) {
	db.users[u.Id] = u
	db.usersByEmail[u.Email] = u
}

func (db *MockDB) Find(id string) (User, bool) {
	u, ok := db.users[id]
	return u, ok
}

func (db *MockDB) FindUserByEmail(email string) (User, bool) {
	u, ok := db.usersByEmail[email]
	return u, ok
}
