package storage

import "com.loollab/postalapi/types"

type Storage interface {
	Get(int) *types.User
}
