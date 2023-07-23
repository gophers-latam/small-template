package storage

import (
	"sync"

	"github.com/gophers-latam/small-template/models"
)

type Repository interface {
	Get(int) *models.Product
	Save(*models.Product)
	Delete(int)
}

type MemoryRepository struct {
	db     map[int]*models.Product
	nextId int
	mutex  sync.Mutex
}

func NewMemoryRepository() *MemoryRepository {
	db := make(map[int]*models.Product)
	return &MemoryRepository{db: db, mutex: sync.Mutex{}}
}

func (repo *MemoryRepository) Save(product *models.Product) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.nextId++
	product.Id = repo.nextId
	repo.db[product.Id] = product
}

func (repo *MemoryRepository) Get(id int) *models.Product {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.db[id]
}

func (repo *MemoryRepository) Delete(id int) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	delete(repo.db, id)
}
