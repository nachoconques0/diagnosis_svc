# TopDoctor
### Made with <3 by Juan Calcagno AKA Nacho. 

### What I've done
I've built a microservice that is allowing to handle patients and their diagnosis. Still too much work to be done. But it is doing what it was asked :)

### Features
- [x] Create a Patient
- [x] Find Patients
- [x] Create a Diagnosis
- [x] Find Diagnosis
- [x] Login with an specific user
- [x] It has validations, at entity level and also at the controller layer (simple ones)
- [x] HTTP Endpoints, including a health check


### Postman Collection available :white_check_mark:
It is available in the repo.

### How to run it 🙀
1. Have docker in your machine
2. `git clone` this repo
3. Once you are inside the repo
4. Run `docker compose up -d` this will initiate a container with a running postgres DB
6. Run `make migration-run dir=up` this will run all needed migrations
8. Run `make seed` in order to have the USER and some extra data so you can test from postman
9. Run `make run` and if all good. Project should be running ready to get some HTTP calls


### You don't want to run it? But want to trust the tests? 😈
1. Have docker in your machine
2. `git clone` this repo
3. Once you are inside the repo
4. Run `make test` and this will trigger a docker compose file that will spin up a test DB + mgirations and then run all the needed tests. By the time of writing this test are passing lol. 🤞🏼


## HTTP Endpoints
#### Login `POST /v1/login`
- All fields must be in payload
##### Body
```
{
    "email":"nacho@gmail.com",
    "password":"testing123123"
}
```

##### Response 200
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hY2hvQGdtYWlsLmNvbSIsImV4cCI6MTc1MzAxMjQ0MSwibmlja25hbWUiOiJuYWNobyJ9.pHNzJaJbHdPyyEnNtyf9_peXFkUDZV5Lc8cZV5DPDLk"
}

```
#### Get Patients `GET /v1/patients`
- Query params can be used, ie: /v1/patients?name="nacho"
##### Response 200
```
[
    {
        "id": "9f6248b4-7e1e-4d56-84c0-2a814a303315",
        "name": "nachin patient",
        "email": "nachinahcin@gmail.com",
        "dni": "12345678",
        "phone": "123456789",
        "address": "barcelona"
    }
]
```

#### Create Patient `POST /v1/patients`
- Name, email and dni are required. Phone & Address are optional
##### Body
```
{
    "name": "patientttttttt",
    "email": "patieneee@gmail.com",
    "dni": "123123"
}
```

##### Response 200
```
{
    "name": "patientttttttt",
    "email": "patieneee@gmail.com",
    "dni": "123123"
}
```

#### Get Diagnosis `GET /v1/diagnosis`
- Query params can be used, ie: /v1/diagnosis?page_size=10&page=1
##### Response 200
```
[
    {
        "id": "b2d57eef-7fc0-4fd1-bda4-24226e96e694",
        "patient_id": "9f6248b4-7e1e-4d56-84c0-2a814a303315",
        "diagnosis": "gripe 1",
        "created_at": "2025-07-20T13:38:47.625962+02:00"
    },
    {
        "id": "281677af-6a3f-45bf-a3bc-62fbb64ac670",
        "patient_id": "9f6248b4-7e1e-4d56-84c0-2a814a303315",
        "diagnosis": "gripe 2",
        "prescription": "antigripepuntero2",
        "created_at": "2025-07-19T13:38:47.625962+02:00"
    },
    {
        "id": "30ee9079-d7bb-4b10-97f8-32d342e7f7b5",
        "patient_id": "9f6248b4-7e1e-4d56-84c0-2a814a303315",
        "diagnosis": "gripe 3",
        "prescription": "antigripepuntero3",
        "created_at": "2025-07-18T13:38:47.625996+02:00"
    }
]
```

#### Create Diagnosis `POST /v1/diagnosis`
- Name, email and dni are required. Phone & Address are optional
##### Body
```
  {
      "patient_id": "02b6c0bf-3599-4d23-8178-c078dd019017",
      "diagnosis": "probando desde postman",
      "prescription": "testeandosecondpatient"
  }
```

##### Response 200
```
{
    "id": "c8a11bd3-ebef-4be7-b7e8-d7e60c05ca06",
    "patient_id": "9f6248b4-7e1e-4d56-84c0-2a814a303315",
    "diagnosis": "probando desde postman",
    "prescription": "testeandosecondpatient",
    "created_at": "2025-07-20T13:42:33.288480171+02:00"
}
```

### Project folder structure 🌴
```
📦diagnosis_svc
 ┣ 📂cmd
 ┃ ┣ 📂seed
 ┃ ┃ ┗ 📜main.go
 ┃ ┗ 📂server
 ┃ ┃ ┣ 📜dev.go
 ┃ ┃ ┗ 📜main.go
 ┣ 📂internal
 ┃ ┣ 📂app
 ┃ ┃ ┣ 📜instance.go
 ┃ ┃ ┗ 📜options.go
 ┃ ┣ 📂controller
 ┃ ┃ ┣ 📂diagnosis
 ┃ ┃ ┃ ┣ 📜controller.go
 ┃ ┃ ┃ ┗ 📜controller_test.go
 ┃ ┃ ┣ 📂patient
 ┃ ┃ ┃ ┣ 📜controler_test.go
 ┃ ┃ ┃ ┗ 📜controller.go
 ┃ ┃ ┗ 📂user
 ┃ ┃ ┃ ┣ 📜controller.go
 ┃ ┃ ┃ ┗ 📜controller_test.go
 ┃ ┣ 📂db
 ┃ ┃ ┣ 📜db.go
 ┃ ┃ ┗ 📜options.go
 ┃ ┣ 📂entity
 ┃ ┃ ┣ 📂diagnosis
 ┃ ┃ ┃ ┣ 📜diagnosis.go
 ┃ ┃ ┃ ┗ 📜diagnosis_test.go
 ┃ ┃ ┣ 📂patient
 ┃ ┃ ┃ ┣ 📜patient.go
 ┃ ┃ ┃ ┗ 📜patient_test.go
 ┃ ┃ ┗ 📂user
 ┃ ┃ ┃ ┗ 📜user.go
 ┃ ┣ 📂env
 ┃ ┃ ┗ 📜env.go
 ┃ ┣ 📂errors
 ┃ ┃ ┗ 📜errors.go
 ┃ ┣ 📂helpers
 ┃ ┃ ┣ 📂db
 ┃ ┃ ┃ ┗ 📜db.go
 ┃ ┃ ┗ 📂query
 ┃ ┃ ┃ ┣ 📜filters.go
 ┃ ┃ ┃ ┗ 📜pagination.go
 ┃ ┣ 📂http
 ┃ ┃ ┣ 📂middleware
 ┃ ┃ ┃ ┗ 📜authentication.go
 ┃ ┃ ┣ 📜http.go
 ┃ ┃ ┣ 📜options.go
 ┃ ┃ ┗ 📜routes.go
 ┃ ┣ 📂mocks
 ┃ ┃ ┣ 📜mock_diagnosis_controller.go
 ┃ ┃ ┣ 📜mock_diagnosis_service.go
 ┃ ┃ ┣ 📜mock_patient_controller.go
 ┃ ┃ ┣ 📜mock_patient_service.go
 ┃ ┃ ┣ 📜mock_user_controller.go
 ┃ ┃ ┗ 📜mock_user_service.go
 ┃ ┣ 📂repo
 ┃ ┃ ┣ 📂diagnosis
 ┃ ┃ ┃ ┣ 📜diagnosis.go
 ┃ ┃ ┃ ┗ 📜diagnosis_test.go
 ┃ ┃ ┣ 📂patient
 ┃ ┃ ┃ ┣ 📜patient.go
 ┃ ┃ ┃ ┗ 📜patient_test.go
 ┃ ┃ ┗ 📂user
 ┃ ┃ ┃ ┣ 📜user.go
 ┃ ┃ ┃ ┗ 📜user_test.go
 ┃ ┗ 📂service
 ┃ ┃ ┣ 📂diagnosis
 ┃ ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┃ ┗ 📜service_test.go
 ┃ ┃ ┣ 📂patient
 ┃ ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┃ ┗ 📜service_test.go
 ┃ ┃ ┗ 📂user
 ┃ ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┃ ┗ 📜service_test.go
 ┣ 📂migrations
 ┃ ┣ 📜20250716192739_init.down.sql
 ┃ ┣ 📜20250716192739_init.up.sql
 ┃ ┣ 📜20250716192919_patient-diagnose-tables.down.sql
 ┃ ┗ 📜20250716192919_patient-diagnose-tables.up.sql
 ┣ 📜.gitignore
 ┣ 📜.golangci.yml
 ┣ 📜Makefile
 ┣ 📜README.md
 ┣ 📜docker-compose.yml
 ┣ 📜docker-compose_test.yml
 ┣ 📜generate-mocks.sh
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┗ 📜postman_collection.json
 ```