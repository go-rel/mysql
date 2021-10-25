package mysql

import (
	"testing"

	"github.com/go-rel/primaryreplica"
)

func TestPrimaryReplica_specs(t *testing.T) {
	var (
		adapter = primaryreplica.New(
			MustOpen("root:my_root_password@tcp(localhost:23306)/rel_test?charset=utf8&parseTime=True&loc=Local"),
			MustOpen("root:my_root_password@tcp(localhost:23307)/rel_test?charset=utf8&parseTime=True&loc=Local"),
		)
	)

	defer adapter.Close()
	AdapterSpecs(t, adapter)
}
