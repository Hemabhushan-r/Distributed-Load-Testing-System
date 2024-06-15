package models

type TestConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Other fields remaining
}

type TaskRequest struct {
	ConfigID string `json:"config_id"`
	//Other fields remaining
}
