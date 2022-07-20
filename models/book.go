package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Price       *int   `json:"price"`
	AuthorID    *int   `json:"author_id"`
	PublisherID *int   `json:"publisher_id"`
}

func (b *Book) AfterCreate(tx *gorm.DB) error {
	var err error
	publisher := new(Publisher)
	author := new(Author)

	if err = tx.First(publisher, b.PublisherID).Error; err != nil {
		return err
	}
	tx.Model(publisher).Update("how_many", publisher.HowMany+1)

	if err = tx.First(author, b.AuthorID).Error; err != nil {
		return err
	}
	tx.Model(author).Update("how_many", author.HowMany+1)

	return nil
}
