package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/db/orm"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/db/queries"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
)

type DataBase struct {
	db *orm.DataBase
	m  *sync.RWMutex
}

var _ gateways.Repository = &DataBase{}

func convertError(err *error) {
	if *err == nil {
		return
	}

	if _, ok := (*err).(*usecases.Error); ok {
		return
	}

	*err = usecases.WrapError(usecases.CodeDataBase, *err)
}

func New(dbFile string) (r gateways.Repository, err error) {
	defer convertError(&err)

	db, err := orm.New(dbFile)
	if err != nil {
		return r, err
	}

	d := &DataBase{
		db: db,
		m:  &sync.RWMutex{},
	}

	if err := d.CreateTable(); err != nil {
		err = fmt.Errorf("faild to setup database: %w", err)
		return d, err
	}
	return d, nil
}

func (d *DataBase) Close() (err error) {
	defer convertError(&err)
	err = d.db.Close()
	return
}

func (d *DataBase) CreateRemote(ctx context.Context, r *domain.Remote) (_ *domain.Remote, err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			r, err = queries.InsertIntoRemotes(ctx, tx, r)
			return err
		},

		func(tx *sql.Tx) error {
			for _, button := range r.Buttons {
				_, err = queries.InsertIntoButton(ctx, tx, domain.RemoteID(r.ID), button)
			}
			return err
		},
	}, false)
	return r, err
}

func (d *DataBase) ReadRemote(ctx context.Context, remoteID domain.RemoteID) (r *domain.Remote, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			r, err = queries.SelectFromRemotesWhere(ctx, tx, remoteID)
			return err
		},
		func(tx *sql.Tx) error {
			buttons, err := queries.SelectFromButtons(ctx, tx, remoteID)
			r.Buttons = buttons
			return err
		},
	}, true)
	return
}

func (d *DataBase) ReadRemotes(ctx context.Context) (remotes []*domain.Remote, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			remotes, err = queries.SelectFromRemotes(ctx, tx)
			return err
		},
		func(tx *sql.Tx) error {
			for _, r := range remotes {
				buttons, err := queries.SelectFromButtons(ctx, tx, r.ID)
				if err != nil {
					return err
				}
				r.Buttons = buttons
			}
			return nil
		},
	}, true)
	return
}

func (d *DataBase) ReadButton(ctx context.Context, buttonID domain.ButtonID) (c *domain.Button, err error) {
	d.m.RLock()
	defer convertError(&err)
	d.m.RUnlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			c, err = queries.SelectFromButtonsWhere(ctx, tx, buttonID)
			return err
		},
	}, true)
	return
}

func (d *DataBase) ReadButtons(ctx context.Context, remoteID domain.RemoteID) (coms []*domain.Button, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			coms, err = queries.SelectFromButtons(ctx, tx, remoteID)
			return err
		},
	}, true)
	return
}

func (d *DataBase) ReadIRDataAndDeviceID(ctx context.Context, buttonID domain.ButtonID) (
	irData domain.IRData, deviceId domain.DeviceID, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			irData, deviceId, err = queries.SelectIRDataAndDeviceIDFromButtonsWhere(ctx, tx, buttonID)
			return err
		},
	}, true)
	return
}

func (d *DataBase) UpdateRemote(ctx context.Context, a *domain.Remote) (err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdateRemote(ctx, tx, a)
		},
	}, false)
	return
}

func (d *DataBase) LearnIR(ctx context.Context, buttonID domain.ButtonID, irData domain.IRData) (err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			return queries.LearnIRData(ctx, tx, buttonID, irData)
		},
	}, false)
	return
}

func (d *DataBase) DeleteRemote(ctx context.Context, remoteID domain.RemoteID) (err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			return queries.DeleteFromRemoteWhere(ctx, tx, remoteID)
		},
	}, false)
	return err
}
