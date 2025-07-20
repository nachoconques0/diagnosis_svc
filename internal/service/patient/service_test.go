package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/mocks"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
	service "github.com/nachoconques0/diagnosis_svc/internal/service/patient"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPatientRepository(ctrl)
	svc := service.New(mockRepo)
	ctx := context.Background()

	patientTestName := "nachin"
	patientTestDNI := "123123"
	patientTestEmail := "nacho@gmail.com"

	req := model.CreatePatientRequest{
		Name:  patientTestName,
		DNI:   patientTestDNI,
		Email: patientTestEmail,
	}
	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&patient.Entity{
			Name:  patientTestName,
			DNI:   patientTestDNI,
			Email: patientTestEmail,
		}, nil)

		res, err := svc.Create(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, patientTestName, res.Name)
	})

	t.Run("error from repo", func(t *testing.T) {
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("errFroMRepoo"))

		_, err := svc.Create(ctx, req)
		assert.Error(t, err)
	})
}

func TestService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockPatientRepository(ctrl)
	svc := service.New(repo)
	ctx := context.Background()

	filters := query.DiagnosisFilters{PatientName: "nachhh"}
	pagination := query.Pagination{Page: 1, PageSize: 5}

	t.Run("success", func(t *testing.T) {
		repo.EXPECT().Find(gomock.Any(), filters, pagination).Return([]patient.Entity{{Name: "nachin"}}, nil)

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
