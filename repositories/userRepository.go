package repositories

import (
	"postly/configs"
	"postly/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur UserRepository) GetUserByUsernameOrEmail(username string, email string) (models.User, error) {
	var user models.User
	err := configs.DB.QueryRow("Select id,username,email,password,firstname, lastname from users where username = $1 or email = $2", username, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Firstname,
		&user.Lastname)
	return user, err
}

func (ur UserRepository) CreateUser(register models.Register) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	_, err := configs.DB.Exec("Insert into users(username,password,email,firstname,lastname) values($1,$2,$3,$4,$5)",
		register.Username, hashedPassword, register.Email, register.Firstname, register.Lastname)
	return err
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
	return user, err
}

func (ur UserRepository) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := configs.DB.QueryRow("Select id,username,email,password,firstname, lastname from users where id = $1", id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Firstname,
		&user.Lastname)
	return user, err
}

func (ur UserRepository) SearchUsers(search string) ([]models.UserResponse, error) {
	rows, err := configs.DB.Query(`
			select id,username,firstname,lastname
			from users
		  	where username ilike '%' || $1 || '%'
           	or firstname ilike '%' || $1 || '%'
           	or lastname ilike '%' || $1 || '%'
			order by id desc
	`, search)
	if err != nil {
		return nil, err
	}
	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Firstname,
			&user.Lastname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
