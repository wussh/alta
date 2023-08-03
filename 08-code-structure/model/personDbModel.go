package model

import (
	"gorm.io/gorm"
)

type PersonDbModel struct {
	db *gorm.DB
}

func NewPersonDbModel(db *gorm.DB) *PersonDbModel {
	db.AutoMigrate(&Person{})
	return &PersonDbModel{
		db: db,
	}
}

func (pm *PersonDbModel) GetAll() ([]Person, error) {
	var allPerson []Person
	err := pm.db.Find(&allPerson).Error
	return allPerson, err
}

func (pm *PersonDbModel) Add(p Person) (Person, error) {
	err := pm.db.Save(&p).Error
	return p, err
}

func (pm *PersonDbModel) Edit(id int, newP Person) (Person, error) {
	p := Person{}
	err := pm.db.First(&p, id).Error
	if err != nil {
		return p, err
	}
	p.Address = newP.Address
	p.Name = newP.Name
	p.Email = newP.Email
	p.Password = newP.Password
	err = pm.db.Save(&p).Error
	return p, err
}
