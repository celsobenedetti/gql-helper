package main

type AnalyticDownloadParams struct {
	Timezone  string    `json:"timezone"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Test      bool      `json:"test"`
	Nids      []float32 `json:"nids"`
}
