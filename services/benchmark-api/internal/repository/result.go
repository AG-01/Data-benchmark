package repository

import (
	"gorm.io/gorm"
	"benchmark-api/internal/models"
)

type ResultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) *ResultRepository {
	return &ResultRepository{db: db}
}

func (r *ResultRepository) Create(result *models.Result) error {
	return r.db.Create(result).Error
}

func (r *ResultRepository) GetByID(id uint) (*models.Result, error) {
	var result models.Result
	err := r.db.First(&result, id).Error
	return &result, err
}

func (r *ResultRepository) GetByBenchmarkID(benchmarkID uint) ([]models.Result, error) {
	var results []models.Result
	err := r.db.Where("benchmark_id = ?", benchmarkID).Find(&results).Error
	return results, err
}

func (r *ResultRepository) List(filters map[string]interface{}, limit, offset int) ([]models.Result, error) {
	var results []models.Result
	query := r.db
	
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	
	err := query.Limit(limit).Offset(offset).Find(&results).Error
	return results, err
}

func (r *ResultRepository) Update(result *models.Result) error {
	return r.db.Save(result).Error
}

func (r *ResultRepository) Delete(id uint) error {
	return r.db.Delete(&models.Result{}, id).Error
}
