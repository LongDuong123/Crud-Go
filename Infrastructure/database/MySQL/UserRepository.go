package mysql

import (
	domain "crud/Domain"
	"crud/helpers"
	"database/sql"
)

type UserRepository struct {
	db *databaseMySQL
}

func NewUserRepository(_db *databaseMySQL) domain.UserRepository {
	return &UserRepository{db: _db}
}

func (ur *UserRepository) GetByID(id int) (*domain.User, error) {
	user, err := ur.db.Mysql.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	var getUser domain.User
	if user.Next() {
		var age sql.NullInt64
		var gender sql.NullString
		err := user.Scan(&getUser.ID, &getUser.Name, &age, &gender, &getUser.Email, &getUser.Role, &getUser.Password)
		if err != nil {
			return nil, err
		}
		getUser.Age = helpers.NullInt64ToInt(age)
		getUser.Gender = helpers.NullStringToString(gender)
	}
	return &getUser, nil
}

func (ur *UserRepository) GetByPage(start int, end int) ([]domain.User, error) {
	user, err := ur.db.Mysql.Query("SELECT * FROM users WHERE id >= ? and id <= ?", start, end)
	if err != nil {
		return nil, err
	}
	var getListUser []domain.User
	for user.Next() {
		var getUser domain.User
		err := user.Scan(&getUser.ID, &getUser.Name, &getUser.Age, &getUser.Gender, &getUser.Email, &getUser.Role, &getUser.Password)
		if err != nil {
			return nil, err
		}
		getListUser = append(getListUser, getUser)
	}
	return getListUser, nil
}

func (ur *UserRepository) GetByEmail(email string) (*domain.User, error) {
	user, err := ur.db.Mysql.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}
	var getUser domain.User
	if user.Next() {
		var age sql.NullInt64
		var gender sql.NullString
		err := user.Scan(&getUser.ID, &getUser.Name, &age, &gender, &getUser.Email, &getUser.Role, &getUser.Password)
		if err != nil {
			return nil, err
		}
		getUser.Age = helpers.NullInt64ToInt(age)
		getUser.Gender = helpers.NullStringToString(gender)
	}
	return &getUser, nil
}

func (ur *UserRepository) CheckEmail(email string) (bool, error) {
	user, err := ur.db.Mysql.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return false, err
	}
	if user.Next() {
		return false, nil
	}
	return true, nil
}

func (ur *UserRepository) Create(user *domain.User) error {
	age := helpers.IntToInt64Null(int64(user.Age))
	gender := helpers.StringToNullString(user.Gender)
	_, err := ur.db.Mysql.Exec("INSERT INTO users (name,age,gender,email,password) VALUES (?,?,?,?,?)", user.Name, age, gender, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdateByID(id int, user *domain.User) error {
	_, err := ur.db.Mysql.Exec("UPDATE users SET name = ?, age = ?, gender = ?, email = ? WHERE id = ?", user.Name, user.Age, user.Gender, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteByID(id int) error {
	_, err := ur.db.Mysql.Exec("DELETE FROM users WHERE id =?", id)
	if err != nil {
		return err
	}
	return nil
}
