package db

import (
	"embed"
	"fmt"
)

//go:embed queries/createTable.sql
var createTableQueries embed.FS

func (d *DataBase) CreateTable() error {
	query, err := createTableQueries.ReadFile("queries/createTable.sql")
	if err != nil {
		return err
	}

	if _, err := d.db.ExecStmt(string(query)); err != nil {
		err = fmt.Errorf("faild to create table: %w", err)
		return err
	}
	return nil
}
