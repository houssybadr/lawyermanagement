package repository

import (
	"test/internal/models"

	"gorm.io/gorm"
)

type Repository[T models.Admin | models.Avocat | models.Client | models.Dossier | models.Document] struct {
	db *gorm.DB
}

type transaction func(tx *gorm.DB) error

func (r *Repository[T]) SetDB(db *gorm.DB) {
	r.db = db
}

func (r Repository[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r Repository[T]) GetAll(items *[]T) error {
	return r.db.Find(items).Error
}

func (r Repository[T]) GetById(item *T,id uint) error {
	return r.db.First(item, id).Error
}

func (r Repository[T]) GetByField(items *[]T,field string,value interface{}) error {
	return r.db.Where(field+" = ?", value).Find(items).Error
}

func (r Repository[T]) Delete(id uint) error {
	var item *T
	return r.db.Delete(&item, id).Error
}

func (r Repository[T]) Updates(item T,id uint) error{
	err:=r.db.Model(&item).Where("id=?",id).Updates(item).Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository[T]) Count(count *int64) error {
	var item T
	return r.db.Model(&item).Count(count).Error
}

func (r Repository[T]) CountByField(count *int64,field string,value interface{}) error{
	var item T
	return r.db.Model(&item).Where(field+"=?",value).Count(count).Error
}

func (r Repository[T]) Transaction(fn transaction) error {
	err := r.db.Transaction(fn)
	if err != nil {
		return err
	}
	return nil
}
