package storage

import "com.loollab/postalapi/types"

type Storage interface {
	Get(int) *types.User
	GetFoo(int) *types.User
	GetBar(int) *types.User
	GetName(int) *types.User
}
