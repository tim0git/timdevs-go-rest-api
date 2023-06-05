package utils

import "os"

func GetTableName() string {
	return os.Getenv("TABLE_NAME")
}
