package storage

import (
	"sync"

	"github.com/gophers-latam/small-template/domain"
)

type Repository interface {
	Get(int) *domain.Product
	Save(*domain.Product)
	Delete(int)
}

type MemoryRepository struct {
	db     map[int]*domain.Product
	nextId int
	mutex  sync.Mutex
}

func NewMemoryRepository() *MemoryRepository {
	db := make(map[int]*domain.Product)
	return &MemoryRepository{db: db, mutex: sync.Mutex{}}
}

func (repo *MemoryRepository) Save(product *domain.Product) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.nextId++
	product.Id = repo.nextId
	repo.db[product.Id] = product
}

func (repo *MemoryRepository) Get(id int) *domain.Product {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.db[id]
}

func (repo *MemoryRepository) Delete(id int) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	delete(repo.db, id)
}
