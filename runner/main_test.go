package runner

import (
	"fault/util"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	util.LoadConfig()
	os.Exit(m.Run())
}
