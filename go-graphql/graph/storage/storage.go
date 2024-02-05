package storage

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-graphql/graph/model"
	"sync"
	"time"
)

var (
	notFoundError = errors.New("not found")
	//	existsError   = errors.New("exists")
)

type UserStorage struct {
	lock sync.Mutex
	data map[string]model.User
	db   *sqlx.DB
}

func NewUserStroage(db *sqlx.DB) *UserStorage {
	return &UserStorage{
		data: make(map[string]model.User),
		db:   db,
	}
}

func (u *UserStorage) Get(userId string) (model.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	fmt.Println("user get")
	time.Sleep(time.Second)

	user, ok := u.data[userId]
	if !ok {
		return model.User{}, notFoundError
	}

	return user, nil
}

func (u *UserStorage) Put(user model.NewUser) (*model.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	fmt.Println("user put")

	fmt.Printf("Debug: User -  Name: %d\n", user.Name)

	res, err := u.db.NamedExec("INSERT INTO `users` (`name`) VALUES (:name)", user)

	fmt.Println("user put")

	if err != nil {
		return nil, err
	}

	fmt.Printf("Debug: User -  Name: %d\n", user.Name)

	id, err := res.LastInsertId()
	fmt.Println("user put")
	if err != nil {
		return nil, err
	}
	fmt.Println("user put")
	newUser := &model.User{
		ID:   id,
		Name: user.Name,
	}

	fmt.Printf("Debug: User - ID: %d, Name: %d\n", newUser.ID, newUser.Name)

	return newUser, nil
}
