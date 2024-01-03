package queries

import (
	"context"
	"database/sql"
	"fmt"

	adapter "github.com/NaKa2355/pirem/internal/app/pirem/adapter/proto"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
)

func SelectIRDataAndDeviceIDFromButtonsWhere(ctx context.Context, tx *sql.Tx, buttonID domain.ButtonID) (irData domain.IRData, deviceID domain.DeviceID, err error) {
	irData = &domain.RawData{}
	row := tx.QueryRowContext(
		ctx,
		`SELECT  buttons.irdata, device_id FROM buttons INNER JOIN remotes ON buttons.remote_id = remotes.remote_id WHERE button_id=?`,
		buttonID,
	)

	binaryIRData := []byte{}
	err = row.Scan(&binaryIRData, &deviceID)
	if err != nil {
		return
	}
	if len(binaryIRData) == 0 {
		err = usecases.WrapError(usecases.CodeNotFound, fmt.Errorf("irdata not learned"))
		return
	}
	irData, err = adapter.UnmarshalBinaryIRData(binaryIRData)
	return
}
