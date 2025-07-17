package query

import "time"

type DiagnosisFilters struct {
	PatientName string
	Date        *time.Time
}
