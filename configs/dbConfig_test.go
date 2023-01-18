package configs_test

import (
	"os"
	"testing"

	"github.com/nurcholisnanda/wallet-record/configs"
	"github.com/stretchr/testify/assert"
)

func TestBuildDBConfig(t *testing.T) {
	// set test environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_NAME", "testdb")

	// invoke the function
	dbConfig := configs.BuildDBConfig()

	// assert that the values match the expected results
	assert.Equal(t, "localhost", dbConfig.Host, "Expected Host to be 'localhost'")
	assert.Equal(t, "3306", dbConfig.Port, "Expected Port to be '3306'")
	assert.Equal(t, "testuser", dbConfig.User, "Expected User to be 'testuser'")
	assert.Equal(t, "testpassword", dbConfig.Password, "Expected Password to be 'testpassword'")
	assert.Equal(t, "testdb", dbConfig.DBName, "Expected DBName to be 'testdb'")
}

func TestDatabaseURL(t *testing.T) {
	// define test input
	dbConfig := &configs.DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "testuser",
		Password: "testpassword",
		DBName:   "testdb",
	}

	// invoke the function
	url := configs.DatabaseURL(dbConfig)

	// assert that the returned URL matches the expected format
	expected := "testuser:testpassword@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
	assert.Equal(t, expected, url, "Expected URL to be '%s'", expected)
}
