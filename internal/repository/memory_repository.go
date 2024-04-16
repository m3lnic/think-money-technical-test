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
func (m *MemoryRepository[K, D]) Update(key K, newData D) (D, error) {
	_, found := m.repository[key]
	if !found {
		return *new(D), ErrKeyNotFound
	}

	m.repository[key] = newData

	return newData, nil
}

// Delete implements IRepository.
func (m *MemoryRepository[K, D]) Delete(key K) error {
	if _, found := m.repository[key]; !found {
		return ErrKeyNotFound
	}

	delete(m.repository, key)

	return nil
}
