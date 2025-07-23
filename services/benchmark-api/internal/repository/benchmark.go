package repository

import (
	"gorm.io/gorm"
	"benchmark-api/internal/models"
)

type BenchmarkRepository struct {
	db *gorm.DB
}

func NewBenchmarkRepository(db *gorm.DB) *BenchmarkRepository {
	return &BenchmarkRepository{db: db}
}

func (r *BenchmarkRepository) Create(benchmark *models.Benchmark) error {
	return r.db.Create(benchmark).Error
}

func (r *BenchmarkRepository) GetByID(id uint) (*models.Benchmark, error) {
	var benchmark models.Benchmark
	err := r.db.Preload("Queries").Preload("Results").First(&benchmark, id).Error
	return &benchmark, err
}

func (r *BenchmarkRepository) List(filters map[string]interface{}, limit, offset int) ([]models.Benchmark, error) {
	var benchmarks []models.Benchmark
	query := r.db.Model(&models.Benchmark{})
	
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	
	err := query.Limit(limit).Offset(offset).Find(&benchmarks).Error
	return benchmarks, err
}

func (r *BenchmarkRepository) Update(benchmark *models.Benchmark) error {
	return r.db.Save(benchmark).Error
}

func (r *BenchmarkRepository) Delete(id uint) error {
	return r.db.Delete(&models.Benchmark{}, id).Error
}
