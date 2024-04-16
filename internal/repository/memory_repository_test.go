package repository_test

import (
	"errors"
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/repository"
)

type TestStructPleaseIgnore struct {
	Data string
}

func TestMemoryRepositoryCRUDComplexType(t *testing.T) {
	t.Parallel()

	testKey := "test"
	testData := "value"

	memoryRepository := repository.NewMemory[string, TestStructPleaseIgnore]()

	// > Create
	createdData, err := memoryRepository.Create(testKey, TestStructPleaseIgnore{Data: testData})
	if err != nil {
		t.Errorf("expected nil, got error(%+v)", err)
	}
	if createdData.Data != testData {
		t.Errorf("expected string(%s), got string(%s)", testData, createdData)
	}

	_, err = memoryRepository.Create(testKey, TestStructPleaseIgnore{Data: testData})
	if err == nil {
		t.Errorf("expected error(%+v), got nil", repository.ErrKeyAlreadyExists)
	}
	if !errors.Is(err, repository.ErrKeyAlreadyExists) {
		t.Errorf("expected error(%+v), got error(%+v)", repository.ErrKeyAlreadyExists, err)
	}

	// > Read
	fetchedVal, err := memoryRepository.Read(testKey)
	if err != nil {
		t.Errorf("expected nil, got error(%+v)", err)
	}
	if fetchedVal.Data != testData {
		t.Errorf("expected string(%s), got string(%s)", testData, fetchedVal.Data)
	}

	_, err = memoryRepository.Read("invalid_key")
	if err == nil {
		t.Errorf("expected error(%+v), got nil", err)
	}
	if !errors.Is(err, repository.ErrKeyNotFound) {
		t.Errorf("expected error(%+v), got error(%+v)", repository.ErrKeyNotFound, err)
	}
}
