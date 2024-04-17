package repository

// > Define repositories to be simple generic interfaces
type IRepository[K comparable, D any] interface {
	All() map[K]D
	Create(K, D) (D, error)
	Read(K) (D, error)
	Update(K, D) (D, error)
	Delete(K) error
}
