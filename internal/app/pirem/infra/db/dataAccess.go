package dataAccess

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/db/orm"
	"github.com/NaKa2355/pirem/internal/app/pirem/infra/db/queries"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	gateway "github.com/NaKa2355/pirem/internal/app/pirem/usecases/gateways"
)

type DataAccess struct {
	db *orm.DataBase
	m  *sync.RWMutex
}

var _ gateway.Repository = &DataAccess{}

func convertError(err *error) {
	if *err == nil {
		return
	}

	if _, ok := (*err).(*usecases.Error); ok {
		return
	}

	*err = usecases.WrapError(usecases.CodeDataBase, *err)
}

func New(dbFile string) (d *DataAccess, err error) {
	defer convertError(&err)

	db, err := orm.New(dbFile)
	if err != nil {
		return d, err
	}

	d = &DataAccess{
		db: db,
		m:  &sync.RWMutex{},
	}

	if err := d.CreateTable(); err != nil {
		err = fmt.Errorf("faild to setup database: %w", err)
		return d, err
	}
	return d, nil
}

func (d *DataAccess) Close() (err error) {
	defer convertError(&err)
	err = d.db.Close()
	return
}

func (d *DataAccess) CreateRemote(ctx context.Context, r *remote.Remote) (_ *remote.Remote, err error) {
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
				_, err = queries.InsertIntoButton(ctx, tx, remote.ID(r.ID), button)
			}
			return err
		},
	}, false)
	return r, err
}

func (d *DataAccess) ReadRemote(ctx context.Context, remoteID remote.ID) (r *remote.Remote, err error) {
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

func (d *DataAccess) ReadRemotes(ctx context.Context) (remotes []*remote.Remote, err error) {
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

func (d *DataAccess) ReadButton(ctx context.Context, buttonID button.ID) (c *button.Button, err error) {
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

func (d *DataAccess) ReadButtons(ctx context.Context, appID remote.ID) (coms []*button.Button, err error) {
	d.m.RLock()
	defer convertError(&err)
	defer d.m.RUnlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			coms, err = queries.SelectFromButtons(ctx, tx, appID)
			return err
		},
	}, true)
	return
}

func (d *DataAccess) UpdateRemote(ctx context.Context, a *remote.Remote) (err error) {
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

func (d *DataAccess) UpdateButton(ctx context.Context, b *button.Button) (err error) {
	d.m.Lock()
	defer convertError(&err)
	defer d.m.Unlock()
	err = d.db.BeginTransaction(ctx, orm.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdateButton(ctx, tx, b)
		},
	}, false)
	return
}

func (d *DataAccess) DeleteRemote(ctx context.Context, remoteID remote.ID) (err error) {
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
