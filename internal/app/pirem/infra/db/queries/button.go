package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/proto"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func InsertIntoButton(ctx context.Context, tx *sql.Tx, remoteID remote.ID, b *button.Button) (*button.Button, error) {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO buttons(button_id, remote_id, name, tag, irdata) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return b, err
	}
	defer stmt.Close()

	var sqliteErr *sqlite.Error

	b.ID = button.ID(genID())

	_, err = stmt.Exec(b.ID, remoteID, b.Name, b.Tag, []byte{})
	if err == nil {
		return b, err
	}

	if _, ok := err.(*sqlite.Error); !ok {
		return b, err
	}

	sqliteErr = err.(*sqlite.Error)
	if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
		err = usecases.WrapError(
			usecases.CodeAlreadyExists,
			fmt.Errorf("same name button already exists: %w", err),
		)
		return b, err
	}

	return b, err
}

func UpdateButton(ctx context.Context, tx *sql.Tx, b *button.Button) (err error) {
	irdata, err := adapter.MarshalIRDataToBinary(b.IRData)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE buttons SET name=?, irdata=? WHERE button_id=?`,
		b.Name, irdata, b.ID)

	if err, ok := err.(*sqlite.Error); ok {
		if err.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			return usecases.WrapError(
				usecases.CodeAlreadyExists,
				fmt.Errorf("same name button already exists: %w", err),
			)
		}
	}
	return
}

func SelectCountFromButtonsWhere(ctx context.Context, tx *sql.Tx, remoteID remote.ID) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM buttons WHERE remote_id=?`, remoteID)
	err = row.Scan(&count)
	return
}

func SelectFromButtons(ctx context.Context, tx *sql.Tx, remoteID remote.ID) (buttons []*button.Button, err error) {
	count, err := SelectCountFromButtonsWhere(ctx, tx, remoteID)
	if err != nil {
		return
	}

	buttons = make([]*button.Button, 0, count)

	rows, err := tx.QueryContext(
		ctx,
		`SELECT name, irdata, button_id, tag, device_id FROM buttons INNER LEFT JOIN remotes ON buttons.remote_id = remotes.remote_id WHERE remote_id=?`,
		remoteID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b = button.Button{}
		binaryIRData := []byte{}
		err = rows.Scan(&b.Name, binaryIRData, &b.ID, &b.Tag, &b.DeviceID)
		if err != nil {
			return
		}
		b.IRData, err = adapter.UnmarshalBinaryIRData(binaryIRData)
		if err != nil {
			b.IRData = nil
		}
		buttons = append(buttons, &b)
	}
	return
}

func SelectFromButtonsWhere(ctx context.Context, tx *sql.Tx, buttonID button.ID) (b *button.Button, err error) {
	b = &button.Button{}

	rows, err := tx.QueryContext(
		ctx,
		`SELECT name, irdata, tag FROM buttons INNER LEFT JOIN remotes ON buttons.remote_id = remotes.remote_id WHERE button_id=?`,
		buttonID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		return b, usecases.WrapError(usecases.CodeNotFound, errors.New("button not found"))
	}

	err = rows.Scan(&b.Name, &b.IRData, &b.Tag, &b.DeviceID)
	b.ID = buttonID
	return b, err
}
