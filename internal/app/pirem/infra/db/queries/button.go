package queries

import (
	"context"
	"database/sql"
	"fmt"

	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/proto"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func InsertIntoButton(ctx context.Context, tx *sql.Tx, remoteID domain.RemoteID, b *domain.Button) (*domain.Button, error) {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO buttons(button_id, remote_id, name, tag, irdata) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return b, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.ID, remoteID, b.Name, b.Tag, []byte{})
	if err == nil {
		return b, err
	}

	if sqliteErr, ok := err.(*sqlite.Error); !ok {
		if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = usecases.WrapError(
				usecases.CodeAlreadyExists,
				fmt.Errorf("same name button already exists: %w", err),
			)
		}
		return b, err
	}

	return b, err
}

func LearnIRData(ctx context.Context, tx *sql.Tx, buttonID domain.ButtonID, domainIRData domain.IRData) (err error) {
	irData, err := adapter.MarshalIRDataToBinary(domainIRData)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE buttons SET irdata=? WHERE button_id=?`, irData, buttonID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = usecases.WrapError(
				usecases.CodeNotFound,
				fmt.Errorf("button not found: %w", err),
			)
		}
	}
	return
}

func SelectCountFromButtonsWhere(ctx context.Context, tx *sql.Tx, remoteID domain.RemoteID) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM buttons WHERE remote_id=?`, remoteID)
	err = row.Scan(&count)
	return
}

func SelectFromButtons(ctx context.Context, tx *sql.Tx, remoteID domain.RemoteID) (buttons []*domain.Button, err error) {
	count, err := SelectCountFromButtonsWhere(ctx, tx, remoteID)
	if err != nil {
		return
	}

	buttons = make([]*domain.Button, 0, count)
	rows, err := tx.QueryContext(
		ctx,
		`SELECT name, length(irdata) != 0, button_id, tag FROM buttons WHERE buttons.remote_id=?`,
		remoteID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b = domain.Button{}
		err = rows.Scan(&b.Name, &b.HasIRData, &b.ID, &b.Tag)
		if err != nil {
			return
		}
		buttons = append(buttons, &b)
	}
	return
}

func SelectFromButtonsWhere(ctx context.Context, tx *sql.Tx, buttonID domain.ButtonID) (b *domain.Button, err error) {
	b = &domain.Button{}

	rows := tx.QueryRowContext(
		ctx,
		`SELECT name, length(irdata) != 0, tag FROM buttons WHERE button_id=?`,
		buttonID,
	)

	err = rows.Scan(&b.Name, &b.HasIRData, &b.Tag)
	if err != nil {
		if err == sql.ErrNoRows {
			err = usecases.WrapError(
				usecases.CodeNotFound,
				fmt.Errorf("button not found: %w", err),
			)

		}
	}
	b.ID = buttonID
	return b, err
}
