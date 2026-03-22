package repositories

import (
	"database/sql"
	"errors"
	"postly/configs"
	"postly/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur UserRepository) CheckUser(username string) (bool, error) {
	var exists bool

	err := configs.DB.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`,
		username,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (ur UserRepository) CreateUser(register models.Register) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	_, err := configs.DB.Exec("Insert into users(username,password,email,firstname,lastname) values($1,$2,$3,$4,$5)",
		register.Username, hashedPassword, register.Email, register.Firstname, register.Lastname)
	if err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) GetUser(username string) (models.User, error) {
	var user models.User
	err := configs.DB.QueryRow("Select id,username,email,password,firstname, lastname from users where username = $1", username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Firstname,
		&user.Lastname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	return user, nil
}
