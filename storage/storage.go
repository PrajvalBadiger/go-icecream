package storage

import "github.com/PrajvalBadiger/go-icecream/types"

// Contains interface for database
// This interface can be helpful when we want to change to diffrent database
type Storage interface {
	Create_flavour(*types.Flavour) error
	Get_flavours() ([]*types.Flavour, error)
	Get_flavour_by_id(int) (*types.Flavour, error)
	Delete_flavour(int) error
	Update_flavour(int, *types.Flavour) error
}
