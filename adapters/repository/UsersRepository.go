package repository

import (
	"aBet/model"
	"aBet/usecase/repository"

	_ "github.com/lib/pq"
)

type usersRepository struct {
	db *Orm
}

func NewUsersRepository(db *Orm) repository.UsersRepository {
	return &usersRepository{
		db: db,
	}
}
func (u *usersRepository) GetUsersByName(userName string) ([]model.Users, error) {
	var usersModel []model.Users
	u.db.pgdb.Model(&usersModel).Where("user_name = ?", userName).Select()
	return usersModel, nil
}
func (u *usersRepository) LoginUserAccount(userName string, password string) (model.Users, error) {
	var usersModel model.Users
	u.db.pgdb.Model(&usersModel).Where("password = ?", password).Where("user_name = ?", userName).Select()
	return usersModel, nil
}
func (u *usersRepository) AddUsers(Users model.Users) (model.Users, error) {
	_, err := u.db.pgdb.Model(&Users).Insert()
	return Users, err
}
func (u *usersRepository) EditUsers(Users model.Users) (model.Users, error) {
	_, e := u.db.pgdb.Model(&Users).Where("id = ?", Users.Id).Update()
	return Users, e
	// bundle
}
func (u *usersRepository) DeleteUsersById(Users model.Users) (model.Users, error) {
	_, err := u.db.pgdb.Model(&Users).Where("id = ?", Users.Id).Delete()
	return Users, err
}
func (u *usersRepository) GetByIDUsers(Id string) ([]model.Users, error) {
	var usersModel []model.Users
	if Id == "all" {
		u.db.pgdb.Model(&usersModel).Select()
	} else {
		u.db.pgdb.Model(&usersModel).Where("id = ?", Id).Select()
	}
	return usersModel, nil
	// bundle
}
