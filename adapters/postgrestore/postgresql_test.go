package postgrestore_test

import (
	"hexagon/adapters"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Info struct {
	CurrentUser string `db:"current_user"`
}

func TestConnection(t *testing.T) {
	dbName, dbUser, dbPass := "test1", "test1", "123456"
	db := adapters.CreateConnection(t, dbName, dbUser, dbPass)
	adapters.MigrateTestDatabase(t, db, "../../migrations")

	var info Info
	err := db.Get(&info, "SELECT current_user")
	assert.NoError(t, err)
	assert.Equal(t, dbUser, info.CurrentUser)
}
