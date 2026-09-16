package migrations

import (
	"errors"
	"log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)


func RunMigrations(DBURL string){
	m, err := migrate.New("file://internal/migrations", DBURL)
	if err != nil{
		log.Fatalf("mig failed: %v", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange){
		log.Fatalf("Migration up failed: %v", err)

	}
	log.Println("Mig success")
}