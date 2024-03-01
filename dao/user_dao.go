package dao

import "twix/ent"

type UserDao interface {
	WithTx(tx *ent.Tx) (UserDao, *ent.Tx, error)
	Save() (*ent.User, error)
	Update() (*ent.User, error)
	GetByID() (*ent.User, error)
}
