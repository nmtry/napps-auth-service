package auth

type MockDB struct {
	users map[string]User
}

func NewMockDB() *MockDB {
	return &MockDB{users: make(map[string]User)}
}

func (db *MockDB) Insert(u User) {
	db.users[u.Id] = u
}

func (db *MockDB) Find(id string) (User, bool) {
	u, ok := db.users[id]
	return u, ok
}
