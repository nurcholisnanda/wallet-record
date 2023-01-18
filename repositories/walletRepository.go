package repositories

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/nurcholisnanda/wallet-record/configs"
	"github.com/nurcholisnanda/wallet-record/models"
)

// Repository defines the interface for the wallet record repository
type Repository interface {
	Insert(record *models.Record) error
	Update(record *models.Record) error
	GetByDate(start, end time.Time) ([]models.Record, error)
	GetLatest() (models.Record, error)
}

// repository implements the Repository interface
type repository struct {
	connection *gorm.DB
}

// NewRepository creates a new instance of the repository
func NewRepository() Repository {
	// Connect to the database
	db, err := gorm.Open("mysql", configs.DatabaseURL(configs.BuildDBConfig()))
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	// Automigrate the models
	db.AutoMigrate(&models.Record{})
	// Add index on datetime column
	db.Model(&models.Record{}).AddIndex("idx_datetime", "datetime")
	return &repository{connection: db}
}

// Insert insert new record to the database
func (r *repository) Insert(record *models.Record) error {
	return r.connection.Create(record).Error
}

// Update update record on the database
func (r *repository) Update(record *models.Record) error {
	return r.connection.Save(record).Error
}

// GetLatest get the latest record on the database
func (r *repository) GetLatest() (models.Record, error) {
	var record models.Record
	err := r.connection.Order("datetime desc").First(&record).Error
	return record, err
}

// GetByDate get records based on the start and end datetime
func (r *repository) GetByDate(start, end time.Time) ([]models.Record, error) {
	var records []models.Record
	// format the datetime to match the format in the database
	sTime := fmt.Sprintf("%d-%d-%d %d:00:00", start.Year(), start.Month(), start.Day(), start.Hour())
	eTime := fmt.Sprintf("%d-%d-%d %d:00:00", end.Year(), end.Month(), end.Day(), end.Hour())
	err := r.connection.Where("datetime >= ? and datetime<= ?", sTime, eTime).Find(&records).Error
	return records, err
}
