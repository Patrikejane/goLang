package storage

import "com.loollab/postalapi/types"

type MongoStorage struct {
}

func (s *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
