
mockgen --source=internal/service/diagnosis/service.go --destination=internal/mocks/mock_diagnosis_service.go --package=mocks --mock_names=repository=MockDiagnosisRepository
mockgen --source=internal/service/user/service.go --destination=internal/mocks/mock_user_service.go --package=mocks --mock_names=repository=MockUserRepository
mockgen --source=internal/service/patient/service.go --destination=internal/mocks/mock_patient_service.go --package=mocks --mock_names=repository=MockPatientRepository

mockgen --source=internal/controller/diagnosis/controller.go --destination=internal/mocks/mock_diagnosis_controller.go --package=mocks --mock_names=service=MockDiagnosisService
mockgen --source=internal/controller/patient/controller.go --destination=internal/mocks/mock_patient_controller.go --package=mocks --mock_names=service=MockPatientService
mockgen --source=internal/controller/user/controller.go --destination=internal/mocks/mock_user_controller.go --package=mocks --mock_names=service=MockUserService
