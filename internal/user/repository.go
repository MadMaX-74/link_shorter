package user

import (
	"go_dev/pkg/db"
)

type UserRepository struct {
	Db *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}
func (repository *UserRepository) Create(user *User) (*User, error) {
	err := repository.Db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repository *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := repository.Db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repository *UserRepository) Update(user *User) (*User, error) {
	err := repository.Db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repository *UserRepository) Delete(id uint) error {
	return repository.Db.Delete(&User{}, id).Error
}
