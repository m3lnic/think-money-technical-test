package repository

// > Define repositories to be simple generic interfaces
type IRepository[K comparable, D any] interface {
	Create(K, D) (D, error)
	Read(K) (D, error)
	Update(K, D) (D, error)
	Delete(K) error
}
