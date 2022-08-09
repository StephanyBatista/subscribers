package database

import "gorm.io/gorm"

type IRepository[Entity any] interface {
	Create(obj *Entity) bool
	List(where Entity) *[]Entity
	GetById(id string) *Entity
	GetBy(where Entity) *Entity
}

type Repository[Entity any] struct {
	DB *gorm.DB
}

func (r *Repository[Entity]) Create(obj *Entity) bool {
	result := r.DB.Create(&obj)
	return result.Error == nil
}

func (r *Repository[Entity]) List(where Entity) *[]Entity {
	var entities *[]Entity
	result := r.DB.Where(where).Find(&entities)
	if result.RowsAffected == 0 {
		return nil
	}
	return entities
}

func (r *Repository[Entity]) GetById(id string) *Entity {
	var entity *Entity
	r.DB.First(&entity, id)
	return entity
}

func (r *Repository[Entity]) GetBy(where Entity) *Entity {
	var entity *Entity
	result := r.DB.Where(where).Find(&entity)
	if result.RowsAffected == 0 {
		return nil
	}
	return entity
}
