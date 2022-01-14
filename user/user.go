package user

type UserStore interface {
	Insert(string, string) error
	Get(string) (string, error)
}

type UserService struct {
	store UserStore
}

func NewUserService(s UserStore) *UserService {
	return &UserService{
		store: s,
	}
}

func (u *UserService) CreateUser(key, value string) error {
	err := u.store.Insert(key, value)
	return err
}

func (u *UserService) RetrieveUser(key string) (string, error) {
	return u.store.Get(key)
}
