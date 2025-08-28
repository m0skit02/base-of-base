package teststore

import (
	"log"
	"wb-task-L0/internal/app/model"
	"wb-task-L0/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		log.Printf("User with ID %d not found", id)
		return nil, store.ErrRecordNotFound
	}

	log.Printf("Found user with ID: %d", id)
	return u, nil
}
