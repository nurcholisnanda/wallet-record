package services

import (
	"errors"
	"time"

	"github.com/nurcholisnanda/wallet-record/models"
	"github.com/nurcholisnanda/wallet-record/repositories"
	"go.uber.org/zap"
)

type Service interface {
	CreateRecord(record *models.Record) error
	GetHistory(start, end time.Time) ([]models.Record, error)
	GetLatest() (models.Record, error)
}

type service struct {
	repository repositories.Repository
}

// NewService is a constructor function that returns an instance of the service struct with the provided repository.
func NewService(r repositories.Repository) Service {
	return &service{
		repository: r,
	}
}

// CreateRecord function is used for inserting a new record into the database.
// It checks if the record's datetime is in the future or not, and if it is, it returns an error.
// It also checks if the record's datetime is before the latest record's datetime, and if it is, it returns an error.
// If the record's datetime is valid, it retrieves the latest record and adds the new record's amount to it.
// It then either inserts the new record into the database or updates the latest record with the new amount.
func (s *service) CreateRecord(record *models.Record) error {
	prevAmount := 0.0
	record.Datetime = record.Datetime.UTC()
	dateTime := time.Date(record.Datetime.Year(), record.Datetime.Month(), record.Datetime.Day(), record.Datetime.Hour()+1, 0, 0, 0, time.UTC)
	if dateTime.After(time.Now().UTC()) {
		return errors.New("Please make sure not to add future date")
	}
	res, err := s.repository.GetLatest()
	if err == nil {
		prevAmount = res.Amount
		if dateTime.Before(res.Datetime) {
			return errors.New("could not insert past hour transaction")
		}
	}

	data, err := s.repository.GetByDate(dateTime, dateTime)
	if err != nil || len(data) == 0 {
		record.Amount += prevAmount
		record.Datetime = dateTime
		err := s.repository.Insert(record)
		if err != nil {
			zap.Error(err)
			return err
		}
		return nil
	}
	data[0].Amount += record.Amount
	err = s.repository.Update(&data[0])
	if err != nil {
		zap.Error(err)
	}
	return err
}

// GetHistory function is used for retrieving a list of records within a given date range.
// It converts the start and end datetime to UTC time, then retrieves the records from the repository.
// If there is an error, it logs the error and returns an empty list of records and the error.
// Otherwise, it returns the list of records and no error.
func (s *service) GetHistory(start, end time.Time) ([]models.Record, error) {
	start = start.UTC()
	end = end.UTC()
	startTime := time.Date(start.Year(), start.Month(), start.Day(), start.Hour()+1, 0, 0, 0, time.UTC)
	endTime := time.Date(end.Year(), end.Month(), end.Day(), end.Hour(), 0, 0, 0, time.UTC)
	data, err := s.repository.GetByDate(startTime, endTime)
	if err != nil {
		zap.Error(err)
		return []models.Record{}, err
	}
	return data, nil
}

// GetLatest function is used for retrieving the latest record from the database.
// It retrieves the latest record from the repository and if there is an error, it logs the error and returns an empty record and the error.
// Otherwise, it returns the latest record and no error.
func (s *service) GetLatest() (models.Record, error) {
	rec, err := s.repository.GetLatest()
	if err != nil {
		zap.Error(err)
		return models.Record{}, err
	}
	return rec, nil
}
