package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	container, db, err := SetupTestDatabase(ctx)
	if err != nil {
		log.Println("Failed container and db setup")
		os.Exit(1)
	}
	defer func() {
		container.Terminate(ctx)
	}()

	testDB = db

	os.Exit(m.Run())
}

func SetupTestDatabase(ctx context.Context) (testcontainers.Container, *gorm.DB, error) {
	containerReq := testcontainers.ContainerRequest{
		Name:         "test_db",
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_USER":     "test",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return nil, nil, err
	}

	port, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, err
	}
	host, err := dbContainer.Host(ctx)
	if err != nil {
		return nil, nil, err
	}

	dbURI := fmt.Sprintf("postgres://test:test@%v:%v/test", host, port.Port())

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		log.Println(err)
		return dbContainer, db, err
	}

	qr, err := os.ReadFile("./migrations_test/000001_init_tables_and_data.up.sql")
	if err != nil {
		log.Println(err)
		return dbContainer, db, err
	}

	if err := db.Exec(string(qr)).Error; err != nil {
		log.Println(err)
		return dbContainer, db, err
	}

	return dbContainer, db, nil
}
