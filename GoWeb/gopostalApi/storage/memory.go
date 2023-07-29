package storage

import "com.loollab/postalapi/types"

type MemoryStorage struct {
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		Id:   1,
		Name: "Foo",
	}
}
func (s *MemoryStorage) GetFoo(id int) *types.User {
	return &types.User{
		Id:   1,
		Name: "Foo",
	}
}
func (s *MemoryStorage) GetBar(id int) *types.User {
	return &types.User{
		Id:   1,
		Name: "Bar",
	}
}
func (s *MemoryStorage) GetName(id int) *types.User {
	return &types.User{
		Id:   1,
		Name: "Name",
	}
}
