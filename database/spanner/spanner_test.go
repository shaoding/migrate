package spanner

import (
	"fmt"
	"github.com/shaoding/migrate"
	"os"
	"testing"

	dt "github.com/shaoding/migrate/database/testing"
	_ "github.com/shaoding/migrate/source/file"
)

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db, ok := os.LookupEnv("SPANNER_DATABASE")
	if !ok {
		t.Skip("SPANNER_DATABASE not set, skipping test.")
	}

	s := &Spanner{}
	addr := fmt.Sprintf("spanner://%s", db)
	d, err := s.Open(addr)
	if err != nil {
		t.Fatalf("%v", err)
	}
	dt.Test(t, d, []byte("SELECT 1"))
}

func TestMigrate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db, ok := os.LookupEnv("SPANNER_DATABASE")
	if !ok {
		t.Skip("SPANNER_DATABASE not set, skipping test.")
	}

	s := &Spanner{}
	addr := fmt.Sprintf("spanner://%s", db)
	d, err := s.Open(addr)
	if err != nil {
		t.Fatalf("%v", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://./examples/migrations", db, d)
	if err != nil {
		t.Fatalf("%v", err)
	}
	dt.TestMigrate(t, m, []byte("SELECT 1"))
}
