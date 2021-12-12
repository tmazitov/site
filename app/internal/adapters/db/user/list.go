package user

import (
	"fmt"
	"site/app/internal/domain/models"
	"site/app/settings"
)

// GetAll get all from users table from the "limit" to the "offset"
func (bs *userStorage) GetAll(offset, limit int) ([]*models.User, error) {

	query := fmt.Sprintf("select username, register, random from users offset %v limit %v", offset, limit)

	// Select from db
	rows, err := settings.DB.Query(query)

	if err != nil {
		e := fmt.Errorf("error get rows from db: %s", err)
		return nil, e
	}

	// Slice of users
	users := []*models.User{}

	// For each row ( user )
	for rows.Next() {
		u := models.User{}
		// Reading from row user data and writing to u
		err := rows.Scan(&u.Username, &u.Register, &u.RandomId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, &u)
	}

	return users, err
}
