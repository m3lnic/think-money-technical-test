package repository

import "errors"

func NewMemory[K comparable, D any]() IRepository[K, D] {
	return &MemoryRepository[K, D]{
		repository: make(map[K]D),
	}
}

type MemoryRepository[K comparable, D any] struct {
	repository map[K]D
}

var ErrKeyAlreadyExists error = errors.New("key already exists")

// Create implements IRepository.
func (m *MemoryRepository[K, D]) Create(key K, data D) (D, error) {
	_, found := m.repository[key]
	if found {
		return *new(D), ErrKeyAlreadyExists
	}

	m.repository[key] = data

	return data, nil
}

var ErrKeyNotFound error = errors.New("key not found")

// Read implements IRepository.
func (m *MemoryRepository[K, D]) Read(key K) (D, error) {
	val, found := m.repository[key]
	if !found {
		return *new(D), ErrKeyNotFound
	}

	return val, nil
}

// Update implements IRepository.
func (*MemoryRepository[K, D]) Update(K, D) (D, error) {
	panic("unimplemented")
}

// Delete implements IRepository.
func (*MemoryRepository[K, D]) Delete(K) error {
	panic("unimplemented")
}
