package config

import "os"

var (
	DB_HOST = os.Getenv("MONGO_HOST")
	DB_NAME = os.Getenv("MONGO_DB_NAME")
)
