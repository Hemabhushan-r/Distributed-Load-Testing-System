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

func SaveConfig(config models.TestConfig) (string, error) {
	var last_inserted_id string
	err := db.QueryRow("INSERT INTO test_configurations (id, name, description) VALUES (gen_random_uuid(), $1, $2) RETURNING id", config.Name, config.Description).Scan(&last_inserted_id)
	//fmt.Println(res.LastInsertId())
	return last_inserted_id, err
}

func GetConfig(id string) (models.TestConfig, error) {
	var config models.TestConfig
	err := db.QueryRow("SELECT id, name, description FROM test_configurations WHERE id = $1", id).Scan(&config.ID, &config.Name, &config.Description)
	return config, err
}

func UpdateConfig(id string, config models.TestConfig) error {
	_, err := db.Exec("UPDATE test_configurations SET name = $1, description = $2 WHERE id = $3", config.Name, config.Description, id)
	return err
}

func DeleteConfig(id string) error {
	_, err := db.Exec("DELETE FROM test_configurations WHERE id = $1", id)
	return err
}
