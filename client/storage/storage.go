package storage

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/mano/models"
	"github.com/rs/zerolog/log"
)

type DbClient struct {
	Connection *pgx.Conn
}

func NewDBConnection(ctx context.Context) (*DbClient, error) {
	dsn := "postgresql://mano:GkESv8-Y_8tjLH9mCcfWqQ@cogent-raccoon-6964.6wr.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Error().Err(err).Msg("establish db connection")
		return nil, err
	}

	return &DbClient{
		Connection: conn,
	}, nil
}

func (db *DbClient) CreateNewUser(ctx context.Context, user *models.User) (pgconn.CommandTag, error) {
	sqlStatement := `INSERT INTO Users (user_id, first_name, last_name, user_token) VALUES ($1, $2, $3, $4)`
	res, err := db.Connection.Exec(ctx, sqlStatement, user.UserId, user.FirstName, user.LastName, user.UserToken)
	return res, err
}
