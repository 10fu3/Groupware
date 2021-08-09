package Repo

import (
	"../Entity"
	"fmt"
)

func (r *Repository) GetGroupRepo() *GroupRepository {
	return &GroupRepository{repo: r}
}

type GroupRepository struct {
	repo *Repository
}

func (r *GroupRepository) Add(group Entity.Group) error {
	return r.repo.doDB(func() error {
		return r.repo.db.Create(group).Error
	})
}

func (r *GroupRepository) Update(group Entity.Group) error {
	return r.repo.doDB(func() error {
		return r.repo.db.Save(group).Error
	})
}

func (r *GroupRepository) Delete(uuid string) error {
	return r.repo.doDB(func() error {
		return r.repo.db.Delete(Entity.Group{}, "uuid = ?", uuid).Error
	})
}

func (r *GroupRepository) Get(uuid string) (*Entity.Group, error) {

	var g Entity.Group

	if nil == r.repo.doDB(func() error {
		return r.repo.db.Find(&g, "uuid = ?", uuid).Error
	}) {
		return nil, fmt.Errorf("GroupRepo.FindByID() occurred error - %s", uuid)
	}
	return &g, nil
}

func (r *GroupRepository) GetAll() ([]Entity.Group, error) {

	var groups []Entity.Group

	if nil == r.repo.doDB(func() error {
		return r.repo.db.Find(&groups).Error
	}) {
		return nil, fmt.Errorf("GroupRepo.GetAll() occurred error")
	}
	return groups, nil
}
