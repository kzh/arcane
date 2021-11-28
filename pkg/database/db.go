package database

import (
	"embed"
	"github.com/gobuffalo/pop/v6"
	"go.uber.org/zap"
)

func Connect() (*pop.Connection, error) {
	config := &pop.ConnectionDetails{
		Dialect:  "postgres",
		Database: "arcane",
		User:     "root",
		Host:     "cockroachdb-public.crdb",
		Port:     "26257",
		Options: map[string]string{
			"sslmode":     "verify-full",
			"sslrootcert": "/cockroach/ca.crt",
			"sslcert":     "/cockroach/client.root.crt",
			"sslkey":      "/cockroach/client.root.key",
		},
	}

	zap.L().Info("Connecting to DB...")

	db, err := pop.NewConnection(config)
	if err != nil {
		return nil, err
	}

	err = db.Open()
	if err != nil {
		return nil, err
	}

	return db, nil
}

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(conn *pop.Connection) error {
	box, err := pop.NewMigrationBox(migrations, conn)
	if err != nil {
		return err
	}

	return box.Up()
}
