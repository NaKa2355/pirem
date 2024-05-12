package db

import (
	"context"
	"database/sql"
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

type ContextKey int

const (
	transaction ContextKey = iota
)

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
		return d, err
	}
	return d, nil
}

func (d *DataBase) Close() (err error) {
	defer convertError(&err)
	err = d.db.Close()
	return
}

func (d *DataBase) Transaction(ctx context.Context, f func(context.Context, gateways.Repository) error) error {
	return d.db.BeginTransaction(ctx, func(tx *sql.Tx) error {
		return f(context.WithValue(ctx, transaction, tx), d)
	}, false)
}

func (d *DataBase) CreateRemote(ctx context.Context, r *domain.Remote) (_ *domain.Remote, err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	r, err = queries.InsertIntoRemotes(ctx, tx, r)
	if err != nil {
		return r, err
	}
	for _, button := range r.Buttons {
		_, err = queries.InsertIntoButton(ctx, tx, domain.RemoteID(r.ID), button)
		if err != nil {
			return r, err
		}
	}
	return r, err
}

func (d *DataBase) ReadRemote(ctx context.Context, remoteID domain.RemoteID) (r *domain.Remote, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	r, err = queries.SelectFromRemotesWhere(ctx, tx, remoteID)
	if err != nil {
		return r, err
	}
	buttons, err := queries.SelectFromButtons(ctx, tx, remoteID)
	r.Buttons = buttons
	return r, err
}

func (d *DataBase) ReadRemotes(ctx context.Context) (remotes []*domain.Remote, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	remotes, err = queries.SelectFromRemotes(ctx, tx)
	if err != nil {
		return remotes, err
	}
	for _, r := range remotes {
		buttons, err := queries.SelectFromButtons(ctx, tx, r.ID)
		if err != nil {
			return remotes, err
		}
		r.Buttons = buttons
	}
	return remotes, nil
}

func (d *DataBase) ReadButton(ctx context.Context, buttonID domain.ButtonID) (b *domain.Button, err error) {
	d.m.RLock()
	defer convertError(&err)
	d.m.RUnlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	return queries.SelectFromButtonsWhere(ctx, tx, buttonID)
}

func (d *DataBase) ReadButtons(ctx context.Context, remoteID domain.RemoteID) (buttons []*domain.Button, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	return queries.SelectFromButtons(ctx, tx, remoteID)
}

func (d *DataBase) ReadIRDataAndDeviceID(ctx context.Context, buttonID domain.ButtonID) (
	irData domain.IRData, deviceId domain.DeviceID, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	return queries.SelectIRDataAndDeviceIDFromButtonsWhere(ctx, tx, buttonID)
}

func (d *DataBase) UpdateRemote(ctx context.Context, a *domain.Remote) (err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	return queries.UpdateRemote(ctx, tx, a)
}

func (d *DataBase) LearnIR(ctx context.Context, buttonID domain.ButtonID, irData domain.IRData) (err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	return queries.LearnIRData(ctx, tx, buttonID, irData)
}

func (d *DataBase) DeleteRemote(ctx context.Context, remoteID domain.RemoteID) (err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	tx := ctx.Value(transaction).(*sql.Tx)
	return queries.DeleteFromRemoteWhere(ctx, tx, remoteID)
}
