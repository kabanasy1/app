package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kabanasy1/app/internal/application/models"
	log "github.com/sirupsen/logrus"
)

type UserStorage struct {
	dbpool *pgxpool.Pool
}

func NewUserStorage(pool *pgxpool.Pool) *UserStorage {
	us := new(UserStorage)
	us.dbpool = pool
	return us
}

func (st *UserStorage) GetListUsers(nameFilter string) []models.User {
	query := "select id, firstname, lastname, officenumber, phonenumber, position from users"
	agrs := make([]interface{}, 0)
	if nameFilter != "" {
		query += " where firstname like $1"
		agrs = append(agrs, fmt.Sprintf("%%%s%%", nameFilter))
	}

	var result []models.User

	err := pgxscan.Select(context.Background(), st.dbpool, &result, query, agrs...)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (st *UserStorage) GetUserById(id int64) models.User {
	query := "select id, firstname, lastname, officenumber, phonenumber, position from users where id = $1"

	var result models.User

	err := pgxscan.Get(context.Background(), st.dbpool, &result, query, id)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (st *UserStorage) CreateUser(user models.User) error {
	query := "insert into users(firstname, lastname, officenumber, phonenumber, position) values($1, $2, $3, $4, $5)"

	_, err := st.dbpool.Exec(context.Background(), query, user.Firstname, user.Lastname, user.Officenumber, user.Phonenumber, user.Position)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
