package client

import (
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/spf13/viper"
)

type PostgresClient struct {
	DB *pg.DB
}

func NewPostgresClient() (PostgresClient, error) {
	db := pg.Connect(&pg.Options{
		Addr:     viper.GetString("postgres.Addr"),
		User:     viper.GetString("postgres.User"),
		Password: viper.GetString("postgres.Password"),
		Database: viper.GetString("postgres.Database")})
	return PostgresClient{
		DB: db,
	}, createSchema(db)
}

// postgres Check
func (c PostgresClient) Check(ctx context.Context, userID int64) (bool, error) {
	if userID == 0 {
		return true, nil
	}

	count := Counter{
		UserID: userID,
	}

	err := c.DB.Model(&count).Where(`user_id = ?`, userID).Select()
	if err != nil {
		if !errors.Is(err, pg.ErrNoRows) {
			return false, err
		}
	}

	if count.Counter >= Cfg.Limit {
		return false, nil
	}
	count.Counter++

	_, err = c.DB.Model(&count).OnConflict("(user_id) DO UPDATE").Set("counter = EXCLUDED.counter").Insert()
	if err != nil {
		return false, err
	}

	return true, nil
}

type Counter struct {
	UserID  int64 `pg:"user_id,pk"`
	Counter int64 `pg:"counter"`
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*Counter)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
