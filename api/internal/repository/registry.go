package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nhan1603/CloneScc/api/internal/repository/alert"
	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
	"github.com/nhan1603/CloneScc/api/internal/repository/request"
	"github.com/nhan1603/CloneScc/api/internal/repository/user"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Registry interface {
	User() user.Repository
	Asset() asset.Repository
	Alert() alert.Repository
	Request() request.Repository
	DoInTx(ctx context.Context, txFunc TxFunc) error
}

// New returns an implementation instance which satisfying Registry
func New(pgConn *sql.DB) Registry {
	return impl{
		user:    user.New(pgConn),
		asset:   asset.New(pgConn),
		alert:   alert.New(pgConn),
		request: request.New(pgConn),
		pgConn:  pgConn,
	}
}

type impl struct {
	alert   alert.Repository
	asset   asset.Repository
	user    user.Repository
	request request.Repository
	txExec  boil.Transactor
	pgConn  *sql.DB
}

// TxFunc is a function that can be executed in a transaction
type TxFunc func(txRegistry Registry) error

// User returns user repo
func (i impl) User() user.Repository {
	return i.user
}

// Asset returns the asset repo
func (i impl) Asset() asset.Repository {
	return i.asset
}

// Alert returns the alert repo
func (i impl) Alert() alert.Repository {
	return i.alert
}

// Request returns the request repo
func (i impl) Request() request.Repository {
	return i.request
}

// DoInTx handles db operations in a transaction
func (i impl) DoInTx(ctx context.Context, txFunc TxFunc) error {
	if i.txExec != nil {
		return errors.New("db tx nested in db tx")
	}

	tx, err := i.pgConn.BeginTx(ctx, nil)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	var committed bool
	defer func() {
		if committed {
			return
		}

		_ = tx.Rollback()
	}()

	newI := impl{
		user:    user.New(tx),
		alert:   alert.New(tx),
		asset:   asset.New(tx),
		request: request.New(tx),
		txExec:  tx,
	}

	if err = txFunc(newI); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return pkgerrors.WithStack(err)
	}

	committed = true

	return nil
}
