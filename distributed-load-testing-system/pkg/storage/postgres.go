package storage

import (
	"database/sql"
	"distributed-load-testing-system/pkg/controller/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func SaveConfig(config models.TestConfig) error {
	_, err := db.Exec("INSERT INTO test_configs (id, name, description) VALUES ($1, $2, $3)", config.ID, config.Name, config.Description)
	return err
}

func GetConfig(id string) (models.TestConfig, error) {
	var config models.TestConfig
	err := db.QueryRow("SELECT id, name, description FROM test_configs WHERE id = $1", id).Scan(&config.ID, &config.Name, &config.Description)
	return config, err
}

func UpdateConfig(id string, config models.TestConfig) error {
	_, err := db.Exec("UPDATE test_configs SET name = $1, description = $2 WHERE id = $3", config.Name, config.Description, id)
	return err
}

func DeleteConfig(id string) error {
	_, err := db.Exec("DELETE FROM test_configs WHERE id = $1", id)
	return err
}
