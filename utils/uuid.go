package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func GetUUID1() (string) {
	uuid := uuid.NewV1()
	return fmt.Sprintf("%s", uuid)
}

func GetUUID4() (string) {
	uuid := uuid.NewV4()
	return fmt.Sprintf("%s", uuid)
}
