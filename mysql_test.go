//go:build singlenode
// +build singlenode

package mysql

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ctx = context.TODO()

func dsn() string {
	if os.Getenv("MYSQL_DATABASE") != "" {
		return os.Getenv("MYSQL_DATABASE") + "?charset=utf8&parseTime=True&loc=Local"
	}

	return "root@tcp(localhost:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"
}

func TestAdapter_specs(t *testing.T) {
	var (
		adapter = MustOpen(dsn())
	)

	defer adapter.Close()
	AdapterSpecs(t, adapter)
}

func TestAdapter_Open(t *testing.T) {
	// with parameter
	assert.NotPanics(t, func() {
		adapter, _ := Open("root@tcp(localhost:3306)/rel_test?charset=utf8")
		defer adapter.Close()
	})

	// without paremeter
	assert.NotPanics(t, func() {
		adapter, _ := Open("root@tcp(localhost:3306)/rel_test")
		defer adapter.Close()
	})
}

func TestAdapter_Transaction_commitError(t *testing.T) {
	adapter, err := Open(dsn())
	assert.Nil(t, err)
	defer adapter.Close()

	assert.NotNil(t, adapter.Commit(ctx))
}

func TestAdapter_Transaction_rollbackError(t *testing.T) {
	adapter, err := Open(dsn())
	assert.Nil(t, err)
	defer adapter.Close()

	assert.NotNil(t, adapter.Rollback(ctx))
}

func TestAdapter_Exec_error(t *testing.T) {
	adapter, err := Open(dsn())
	assert.Nil(t, err)
	defer adapter.Close()

	_, _, err = adapter.Exec(ctx, "error", nil)
	assert.NotNil(t, err)
}

func TestCheck(t *testing.T) {
	assert.Panics(t, func() {
		check(errors.New("error"))
	})
}
