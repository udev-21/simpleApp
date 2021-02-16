package main

import (
	"fmt"
	"time"
)

func migrate() {
	checkErr(db.AutoMigrate(Author{}, Book{}))
	if dbSeedFlag {
		(Author{}).faker()
		(Book{}).faker()
	}
}

func (m Author) faker() {
	db.Model(&Author{}).Create(Author{ID: 1, FirstName: "John", LastName: "Doe", Email: "fakemail@mail.org", CreatedAt: time.Now(), UpdatedAt: time.Now(), BirthDate: time.Now()})
	db.Model(&Author{}).Create(Author{ID: 2, FirstName: "Doe", LastName: "John", Email: "fakemail1@mail.org", CreatedAt: time.Now(), UpdatedAt: time.Now(), BirthDate: time.Now()})
	db.Model(&Author{}).Create(Author{ID: 3, FirstName: "Doe", LastName: "John", Email: "fakemail2@mail.org", CreatedAt: time.Now(), UpdatedAt: time.Now(), BirthDate: time.Now()})
	fmt.Printf(InfoColor, "Table AUTHORS filled with fake data")
}

func (b Book) faker() {
	db.Model(&Book{}).Create(Book{ID: 1, Title: "Book Title 1", Description: "Book description 1", AuthorID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	db.Model(&Book{}).Create(Book{ID: 2, Title: "Book Title 2", Description: "Book description 2", AuthorID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	db.Model(&Book{}).Create(Book{ID: 3, Title: "Book Title 3", Description: "Book description 3", AuthorID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	db.Model(&Book{}).Create(Book{ID: 4, Title: "Book Title 4", Description: "Book description 4", AuthorID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	db.Model(&Book{}).Create(Book{ID: 5, Title: "Book Title 5", Description: "Book description 5", AuthorID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()})

	fmt.Printf(InfoColor, "Table BOOK filled with fake data")
}
