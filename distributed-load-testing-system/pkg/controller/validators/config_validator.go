package validators

import (
	"distributed-load-testing-system/pkg/controller/models"
	"fmt"
)

func ValidateConfig(config models.TestConfig) error {
	if config.Name == "" {
		return fmt.Errorf("Name should not be empty")
	}
	return nil
}
