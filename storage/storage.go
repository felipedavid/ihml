package storage

import "github.com/teze-investimentos/saldopositivo-web/pkg/types"

type Storage interface {
	GetUser(userID int) *types.User
}
