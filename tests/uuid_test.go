package test
import (
	"testing"
	"cmdb/utils"
)

func TestGetUUID1(t *testing.T) {
	uuid := utils.GetUUID1()
	t.Logf("UUID1: %s\n", uuid)
}

func TestGetUUID4(t *testing.T) {
	uuid := utils.GetUUID4()
	t.Logf("UUID1: %s\n", uuid)
}
