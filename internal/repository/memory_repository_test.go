package repository_test

import (
	"errors"
	"sync"
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/repository"
)

type TestStructPleaseIgnore struct {
	Data string
}

// > Step by step test
func TestMemoryRepositoryCRUD(t *testing.T) {
	t.Parallel()

	testKey := "test"
	testData := "value"

	memoryRepository := repository.NewMemory[string, *TestStructPleaseIgnore]()
	initialItem := &TestStructPleaseIgnore{Data: testData}

	t.Run("creates data", func(t *testing.T) {
		createdData, err := memoryRepository.Create(testKey, initialItem)
		if err != nil {
			t.Errorf("expected nil, got error(%+v)", err)
		}
		if createdData.Data != testData {
			t.Errorf("expected string(%s), got string(%s)", testData, createdData)
		}
	})

	// > TODO
	t.Run("all", func(t *testing.T) {
		expectedMap := map[string]*TestStructPleaseIgnore{
			testKey: initialItem,
		}

		outMap := memoryRepository.All()

		if len(expectedMap) != len(outMap) {
			t.Errorf("maps are mismatched")
		}

		matches := true
		for key, expectedCurrent := range expectedMap {
			retrievedCurrent, found := outMap[key]
			if !found || retrievedCurrent != expectedCurrent {
				matches = false
				break
			}
		}

		if !matches {
			t.Errorf("maps keys don't line up")
		}
	})

	t.Run("errors on creating data when key already exists", func(t *testing.T) {
		_, err := memoryRepository.Create(testKey, &TestStructPleaseIgnore{Data: testData})
		if err == nil {
			t.Errorf("expected error(%+v), got nil", repository.ErrKeyAlreadyExists)
		}
		if !errors.Is(err, repository.ErrKeyAlreadyExists) {
			t.Errorf("expected error(%+v), got error(%+v)", repository.ErrKeyAlreadyExists, err)
		}
	})

	t.Run("reads data", func(t *testing.T) {
		fetchedVal, err := memoryRepository.Read(testKey)
		if err != nil {
			t.Errorf("expected nil, got error(%+v)", err)
		}
		if fetchedVal.Data != testData {
			t.Errorf("expected string(%s), got string(%s)", testData, fetchedVal.Data)
		}
	})

	t.Run("errors on invalid key", func(t *testing.T) {
		_, err := memoryRepository.Read("invalid_key")
		if err == nil {
			t.Errorf("expected error(%+v), got nil", err)
		}
		if !errors.Is(err, repository.ErrKeyNotFound) {
			t.Errorf("expected error(%+v), got error(%+v)", repository.ErrKeyNotFound, err)
		}
	})

	// > Update
	t.Run("update test", func(t *testing.T) {
		updatedTestValue := "for honour"
		_, err := memoryRepository.Update(testKey, &TestStructPleaseIgnore{Data: updatedTestValue})
		if err != nil {
			t.Errorf("expected nil, got error(%+v)", err)
		}

		updatedTestData, _ := memoryRepository.Read(testKey)
		if updatedTestData.Data != updatedTestValue {
			t.Errorf("expected string(%s), got string(%s)", updatedTestValue, updatedTestData.Data)
		}
	})

	t.Run("errors when updating invalid key", func(t *testing.T) {
		_, err := memoryRepository.Update("invalid_key", &TestStructPleaseIgnore{Data: ""})
		if err == nil {
			t.Errorf("expected err(%+v), got nil", repository.ErrKeyNotFound)
		}
		if !errors.Is(err, repository.ErrKeyNotFound) {
			t.Errorf("expected error(%+v), got error(%+v)", repository.ErrKeyNotFound, err)
		}
	})

	t.Run("deletes key", func(t *testing.T) {
		err := memoryRepository.Delete(testKey)
		if err != nil {
			t.Errorf("expected nil, got error(%+v)", err)
		}

		_, err = memoryRepository.Read(testKey)
		if err == nil {
			t.Errorf("expected nil, got error(%+v)", err)
		}
		if !errors.Is(err, repository.ErrKeyNotFound) {
			t.Errorf("expected err(%+v), got error(%+v)", repository.ErrKeyNotFound, err)
		}
	})

	t.Run("errors when key not found on delete", func(t *testing.T) {
		err := memoryRepository.Delete(testKey)
		if err == nil {
			t.Errorf("expected error(%+v), got nil", repository.ErrKeyNotFound)
		}
		if !errors.Is(err, repository.ErrKeyNotFound) {
			t.Errorf("expected err(%+v), got err(%+v)", repository.ErrKeyNotFound, err)
		}
	})
}

func benchmarkRepository(b *testing.B, iterations int) {
	testData := "value"

	memoryRepository := repository.NewMemory[int, *TestStructPleaseIgnore]()

	wg := &sync.WaitGroup{}
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(myIteration int, repo repository.IRepository[int, *TestStructPleaseIgnore]) {
			repo.Create(myIteration, &TestStructPleaseIgnore{Data: testData})
			repo.Delete(myIteration)
			wg.Done()
		}(i, memoryRepository)
	}

	wg.Wait()
}

func BenchmarkRepository50(b *testing.B)      { benchmarkRepository(b, 50) }
func BenchmarkRepository100(b *testing.B)     { benchmarkRepository(b, 100) }
func BenchmarkRepository250(b *testing.B)     { benchmarkRepository(b, 250) }
func BenchmarkRepository500(b *testing.B)     { benchmarkRepository(b, 500) }
func BenchmarkRepository5000(b *testing.B)    { benchmarkRepository(b, 5000) }
func BenchmarkRepository50000(b *testing.B)   { benchmarkRepository(b, 50000) }
func BenchmarkRepository250000(b *testing.B)  { benchmarkRepository(b, 250000) }
func BenchmarkRepository1000000(b *testing.B) { benchmarkRepository(b, 1000000) }
