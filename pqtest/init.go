package pqtest

import (
	"github.com/maprost/assertion"
	"log"
	"os"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqtime"
)

type DataInfo struct{}

func (d DataInfo) Driver() string {
	return "postgres"
}

func (d DataInfo) Database() string {
	return "test_pqx"
}

func (d DataInfo) Host() string {
	return "localhost"
}

func (d DataInfo) Port() string {
	return "5432"
}

func (d DataInfo) Username() string {
	return "postgres"
}

func InitDatabase() error {
	return pqx.OpenDatabaseConnection(DataInfo{})
}

func InitDatabaseTest(t assertion.TestEnvironment) assertion.Assert {
	pqtime.Reset()
	assert := assertion.New(t)

	err := InitDatabase()
	assert.Nil(err)

	return assert
}

func InitTransactionTest(t assertion.TestEnvironment) (pqx.Transaction, assertion.Assert) {
	assert := InitDatabaseTest(t)

	tx, err := pqx.New()
	assert.Nil(err)
	tx.AddLogger(log.New(os.Stdout, "", 0))

	return tx, assert
}
