package queries

//SQL database query wrapper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func InsertIntoRemotes(ctx context.Context, tx *sql.Tx, r *domain.Remote) (*domain.Remote, error) {
	_, err := tx.ExecContext(ctx, `INSERT INTO remotes(remote_id, name, device_id, tag) VALUES(?, ?, ?, ?)`, r.ID, r.Name, r.DeviceID, r.Tag)

	if sqlErr, ok := err.(*sqlite.Error); ok {
		if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = usecases.WrapError(
				usecases.CodeAlreadyExists,
				fmt.Errorf("same name domain already exists: %w", err),
			)
			return r, err
		}
	}

	return r, err
}

func SelectFromRemotesWhere(ctx context.Context, tx *sql.Tx, id domain.RemoteID) (r *domain.Remote, err error) {
	r = &domain.Remote{}

	rows, err := tx.QueryContext(ctx, `SELECT * FROM remotes a WHERE a.remote_id = ?`, id)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		err = usecases.WrapError(
			usecases.CodeNotFound,
			errors.New("domain not found"),
		)
		return
	}

	err = rows.Scan(&r.ID, &r.Name, &r.DeviceID, &r.Tag)
	return r, err
}

func selectCountFromRemotes(ctx context.Context, tx *sql.Tx) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM remotes`)
	if err != nil {
		return
	}
	err = row.Scan(&count)
	return
}

func SelectFromRemotes(ctx context.Context, tx *sql.Tx) (remotes []*domain.Remote, err error) {
	count, err := selectCountFromRemotes(ctx, tx)
	if err != nil {
		return
	}

	rows, err := tx.QueryContext(ctx, `SELECT * FROM remotes`)
	if err != nil {
		return
	}
	defer rows.Close()

	remotes = make([]*domain.Remote, 0, count)

	for rows.Next() {
		r := domain.Remote{}
		err = rows.Scan(&r.ID, &r.Name, &r.DeviceID, &r.Tag)
		if err != nil {
			return
		}
		remotes = append(remotes, &r)
	}

	return remotes, err
}

func UpdateRemote(ctx context.Context, tx *sql.Tx, r *domain.Remote) (err error) {
	_, err = tx.ExecContext(ctx, `UPDATE remotes SET name=?, device_id=? WHERE remote_id=?`, r.Name, r.DeviceID, r.ID)
	if sqlErr, ok := err.(*sqlite.Error); ok {
		if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = usecases.WrapError(
				usecases.CodeAlreadyExists,
				fmt.Errorf("same name domain already exists: %w", err),
			)
			return
		}
	}
	return
}

func DeleteFromRemoteWhere(ctx context.Context, tx *sql.Tx, id domain.RemoteID) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM remotes WHERE remote_id=?`, id)
	return
}
