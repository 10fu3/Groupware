package Repo

import (
	"../Entity"
	"fmt"
)

func (r *Repository) GetUserRepository() *UserRepository {
	return &UserRepository{repo: r}
}

type UserRepository struct {
	repo *Repository
}

func (u *UserRepository) Add(usr *Entity.User) error {
	return u.repo.doDB(func() error {
		return u.repo.db.Create(usr).Error
	})
}

func (u *UserRepository) Update(usr *Entity.User) error {
	return u.repo.doDB(func() error {
		return u.repo.db.Save(usr).Error
	})
}

func (u *UserRepository) FindByID(uuid string) (*Entity.User, error) {
	var find Entity.User

	if nil == u.repo.doDB(func() error {
		return u.repo.db.Where("uuid = ?", uuid).First(&find).Error
	}) {
		return &find, nil
	}
	return nil, fmt.Errorf("getUser occurred error - %s", uuid)
}

func (u *UserRepository) Delete(uuid string) error {
	return u.repo.doDB(func() error {
		return u.repo.db.Delete(&Entity.User{}, "uuid = ?", uuid).Error
	})
}
