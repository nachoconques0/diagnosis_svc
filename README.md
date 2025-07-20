# Faceit Challenge
### Made with <3 by Juan Calcagno AKA Nacho. 
#### I love CS:GO BTW hardcode fan since 1.5 lol


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
5. Run `make mod` so you download needed pkgs
6. Run `make migration-run dir=up` this will run all needed migrations
7. Run `make run` and if all good. Project should be running ready to get some HTTP calls and also gRPC.


### You don't want to run it? But want to trust the tests? 😈
1. Have docker in your machine
2. `git clone` this repo
3. Once you are inside the repo
4 Run `make test` and this will trigger a docker compose file that will spin up a test DB + mgirations and then run all the needed tests. By the time of writing this test are passing lol. 🤞🏼


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
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hY2hvQGdtYWlsLmNvbSIsImV4cCI6MTc1MzA5MDQ4NCwibmlja25hbWUiOiJuYWNobyJ9.TPLj-qYRAV6fobVNP2dAHTSvjDfi7As-v2M6EE5GoLU"
}

```
#### Get Patients `GET /v1/patients`
- Query params can be used, ie: /v1/patients?name="nacho"
##### Response 200
```
[
  {
      "ID": "c9d9ec80-3cc9-40f0-a0b6-a69ed606edbb",
      "Name": "nachin patient",
      "DNI": "12345678",
      "Email": "nachinahcin@gmail.com",
      "Phone": "123456789",
      "Address": "barcelona"
  },
  {
      "ID": "02b6c0bf-3599-4d23-8178-c078dd019017",
      "Name": "patientttttttt",
      "DNI": "123123",
      "Email": "patieneee@gmail.com",
      "Phone": null,
      "Address": null
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
    "ID": "02b6c0bf-3599-4d23-8178-c078dd019017",
    "Name": "patientttttttt",
    "DNI": "123123",
    "Email": "patieneee@gmail.com",
    "Phone": null,
    "Address": null
}
```

#### Get Diagnosis `GET /v1/diagnosis`
- Query params can be used, ie: /v1/diagnosis?page_size=10&page=1
##### Response 200
```
[
    {
        "ID": "a659cd01-235d-46ef-abe5-791ca7d46f49",
        "PatientID": "c9d9ec80-3cc9-40f0-a0b6-a69ed606edbb",
        "Diagnosis": "gripe 1",
        "Prescription": null,
        "CreatedAt": "2025-07-20T11:04:24.480295+02:00"
    },
    {
        "ID": "b55dfe79-6e88-4f66-a845-395a58fbd86c",
        "PatientID": "c9d9ec80-3cc9-40f0-a0b6-a69ed606edbb",
        "Diagnosis": "gripe 2",
        "Prescription": "antigripepuntero2",
        "CreatedAt": "2025-07-19T11:04:24.480296+02:00"
    },
    {
        "ID": "e944000d-c44f-4fa6-83ec-cb7062650a57",
        "PatientID": "c9d9ec80-3cc9-40f0-a0b6-a69ed606edbb",
        "Diagnosis": "gripe 3",
        "Prescription": "antigripepuntero3",
        "CreatedAt": "2025-07-18T11:04:24.480328+02:00"
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
    "ID": "a850b612-8900-4f93-8340-61719e4a1f48",
    "PatientID": "02b6c0bf-3599-4d23-8178-c078dd019017",
    "Diagnosis": "probando desde postman",
    "Prescription": "testeandosecondpatient",
    "CreatedAt": "2025-07-20T12:13:17.957952395+02:00"
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