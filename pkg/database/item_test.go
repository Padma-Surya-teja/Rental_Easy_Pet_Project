package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"rental_easy.in/m/pkg/models"
)

var mockItem1 = models.Item{
	Name:          "Iqoo Neo 6",
	Description:   "Phone is Superfast and Display is just amazing",
	Category:      "Mobile Phones",
	AvailableFrom: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	AvailableTo:   time.Date(2023, 01, 03, 0, 0, 0, 0, time.UTC),
	AmountPerDay:  500,
	UserId:        1,
}

func newMockDBClient() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		fmt.Println(err)
		panic("Got an unexpected error while opening mock sql connection.")
	}
	return gormDB, mock
}

func TestAddItem(t *testing.T) {
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	DBClient := DBClient{
		Db: db,
	}
	mock.ExpectBegin()
	t.Run("Check forinsertion", func(t *testing.T) {
		mock.ExpectQuery(`INSERT INTO "items" (.+)`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		x := models.Item{}
		id, err := DBClient.AddItem(x)
		assert.Nil(t, err)

		if id != 1 {
			t.Errorf("Unable to Add Item %v", err)
		}
	})

}
