package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/mocks"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
	service "github.com/nachoconques0/diagnosis_svc/internal/service/diagnosis"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDiagnosisRepository(ctrl)
	svc := service.New(mockRepo)
	ctx := context.Background()

	patientID := uuid.New().String()
	diag := "nachoisabittired"
	prescription := "gotosleeplol"

	req := model.CreateDiagnosisRequest{
		PatientID:    patientID,
		Diagnosis:    diag,
		Prescription: &prescription,
	}
	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&diagnosis.Entity{
			PatientID:    uuid.MustParse(patientID),
			Diagnosis:    diag,
			Prescription: &prescription,
			CreatedAt:    time.Now(),
		}, nil)

		res, err := svc.Create(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, diag, res.Diagnosis)
	})

	t.Run("error from repo", func(t *testing.T) {
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("errFroMRepoo"))

		_, err := svc.Create(ctx, req)
		assert.Error(t, err)
	})
}

func TestService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockDiagnosisRepository(ctrl)
	svc := service.New(repo)
	ctx := context.Background()

	filters := query.DiagnosisFilters{PatientName: "nachhh"}
	pagination := query.Pagination{Page: 1, PageSize: 5}

	t.Run("success", func(t *testing.T) {
		repo.EXPECT().Find(gomock.Any(), filters, pagination).Return([]diagnosis.Entity{{Diagnosis: "sleepyy"}}, nil)
		results, err := svc.Find(ctx, filters, pagination)
		assert.NoError(t, err)
		assert.Len(t, results, 1)
	})

	t.Run("error from repo", func(t *testing.T) {
		repo.EXPECT().Find(gomock.Any(), filters, pagination).Return(nil, errors.New("errFroMRepoo"))
		_, err := svc.Find(ctx, filters, pagination)
		assert.Error(t, err)
	})
}
