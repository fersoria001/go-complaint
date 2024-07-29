package commands_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// pre-process
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5170")
	os.Setenv("ALLOWED_ORIGINS", "http://localhost:3000,localhost:3000,localhost")
	os.Setenv("CSRF_KEY", "ultrasecret")
	os.Setenv("DATABASE_URL", "postgres://postgres:sfdkwtf@localhost:5432/postgres?pool_max_conns=100&search_path=public&connect_timeout=5")
	os.Setenv("PORT", "5170")
	os.Setenv("DNS", "http://localhost:3000")
	os.Setenv("SEND_GRID_API_KEY", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	os.Exit(m.Run())
	// post-process
}
