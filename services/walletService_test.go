package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/nurcholisnanda/wallet-record/mocks"
	"github.com/nurcholisnanda/wallet-record/models"
	"github.com/nurcholisnanda/wallet-record/repositories"
	"github.com/stretchr/testify/mock"
)

var mockRepo = new(mocks.Repository)

func TestNewService(t *testing.T) {
	type args struct {
		r repositories.Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "Success test case",
			args: args{mockRepo},
			want: &service{mockRepo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetHistory(t *testing.T) {
	now := time.Now()
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		setup   func()
		want    []models.Record
		wantErr bool
	}{
		{
			name: "Success test case",
			s: &service{
				repository: mockRepo,
			},
			args: args{
				start: now.Add(-24 * time.Hour),
				end:   now,
			},
			setup: func() {
				mockRepo.On("GetByDate", mock.Anything, mock.Anything).Return([]models.Record{{Amount: 100, Datetime: now}}, nil)
			},
			want:    []models.Record{{Amount: 100, Datetime: now}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := tt.s.GetHistory(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetLatest(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		s       *service
		setup   func()
		want    models.Record
		wantErr bool
	}{
		{
			name: "Success test case",
			s: &service{
				repository: mockRepo},
			setup: func() {
				mockRepo.On("GetLatest").Return(models.Record{Amount: 100, Datetime: now}, nil)
			},
			want:    models.Record{Amount: 100, Datetime: now},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := tt.s.GetLatest()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetLatest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetLatest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateRecord(t *testing.T) {
	now := time.Now().Add(-2 * time.Hour)
	type args struct {
		record *models.Record
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		setup   func()
		wantErr bool
	}{
		{
			name: "Success test case",
			s: &service{
				repository: mockRepo,
			},
			args: args{
				record: &models.Record{Amount: 100, Datetime: now},
			},
			setup: func() {
				mockRepo.On("GetLatest", mock.Anything).Return(models.Record{Amount: 100, Datetime: now}, nil)
				mockRepo.On("GetByDate", mock.Anything, mock.Anything).Return([]models.Record{{Amount: 100, Datetime: now}}, nil)
				mockRepo.On("Update", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			if err := tt.s.CreateRecord(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
