package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type RegistrationConn struct {
	Conn *pgx.Conn
}

func (rc *RegistrationConn) InitDatabasesIfNotExist() (err error) {
	if _, err = rc.Conn.Exec(context.Background(), "create extension if not exists ltree;"); err!=nil{
		return err
	}

	if _, err = rc.Conn.Exec(context.Background(), UserTable); err!=nil{
		return err
	}

	if _, err = rc.Conn.Exec(context.Background(), PostsTable); err!=nil{
		return err
	}

	if _, err = rc.Conn.Exec(context.Background(), CommentsTable); err!=nil{
		return err
	}

	return nil
}

func (rc *RegistrationConn) CreateConnection(path string) (err error) {
	Conn, err := pgx.Connect(context.Background(), path)
	if err != nil {
		return err
	}
	rc.Conn = Conn
	return nil
}