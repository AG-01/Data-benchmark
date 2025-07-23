package repository

import (
	"gorm.io/gorm"
	"benchmark-api/internal/models"
)

type QueryRepository struct {
	db *gorm.DB
}

func NewQueryRepository(db *gorm.DB) *QueryRepository {
	return &QueryRepository{db: db}
}

func (r *QueryRepository) Create(query *models.Query) error {
	return r.db.Create(query).Error
}

func (r *QueryRepository) GetByID(id uint) (*models.Query, error) {
	var query models.Query
	err := r.db.Preload("Executions").First(&query, id).Error
	return &query, err
}

func (r *QueryRepository) GetByBenchmarkID(benchmarkID uint) ([]models.Query, error) {
	var queries []models.Query
	err := r.db.Where("benchmark_id = ?", benchmarkID).Find(&queries).Error
	return queries, err
}

func (r *QueryRepository) List(filters map[string]interface{}, limit, offset int) ([]models.Query, error) {
	var queries []models.Query
	query := r.db.Model(&models.Query{})
	
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	
	err := query.Limit(limit).Offset(offset).Find(&queries).Error
	return queries, err
}

func (r *QueryRepository) Update(query *models.Query) error {
	return r.db.Save(query).Error
}

func (r *QueryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Query{}, id).Error
}
