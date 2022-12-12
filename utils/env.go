package utils

import (
	"fmt"
	"os"
)

var environment string

func GetEnvironment() string {
	if len(environment) == 0 {
		environment = os.Getenv("ENV")
		switch {
		case environment == "prod":
			return environment
		case environment == "dev":
			return environment
		default:
			panic(fmt.Errorf("invalid value for STAGE_STATUS on env"))
		}
	}
	return environment
}
