package env

import "os"

func InProductionEnv() bool {
	appEnv := os.Getenv("APP_ENV")
	return appEnv == "production"
}
