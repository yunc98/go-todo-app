package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (p *Todo) FirstById(id uint) (tx *gorm.DB) {
	return db.Where("id = ?", id).First(&p)
}

func (p *Todo) Create() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *Todo) Update() (tx *gorm.DB) {
	return db.Save(&p)
}