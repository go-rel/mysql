package mysql

import (
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/specs"
	_ "github.com/go-sql-driver/mysql"
)

func AdapterSpecs(t *testing.T, adapter rel.Adapter) {
	repo := rel.New(adapter)

	// Prepare tables
	teardown := specs.Setup(t, repo)
	defer teardown()

	// Migration Specs
	// - Rename column is only supported by MySQL 8.0
	specs.Migrate(t, repo, specs.SkipRenameColumn)

	// Query Specs
	specs.Query(t, repo)
	specs.QueryJoin(t, repo)
	specs.QueryNotFound(t, repo)
	specs.QueryWhereSubQuery(t, repo)

	// Preload specs
	specs.PreloadHasMany(t, repo)
	specs.PreloadHasManyWithQuery(t, repo)
	specs.PreloadHasManySlice(t, repo)
	specs.PreloadHasOne(t, repo)
	specs.PreloadHasOneWithQuery(t, repo)
	specs.PreloadHasOneSlice(t, repo)
	specs.PreloadBelongsTo(t, repo)
	specs.PreloadBelongsToWithQuery(t, repo)
	specs.PreloadBelongsToSlice(t, repo)

	// Aggregate Specs
	specs.Aggregate(t, repo)

	// Insert Specs
	specs.Insert(t, repo)
	specs.InsertHasMany(t, repo)
	specs.InsertHasOne(t, repo)
	specs.InsertBelongsTo(t, repo)
	specs.Inserts(t, repo)
	specs.InsertAll(t, repo)
	specs.InsertAllPartialCustomPrimary(t, repo)

	// Update Specs
	specs.Update(t, repo)
	specs.UpdateNotFound(t, repo)
	specs.UpdateHasManyInsert(t, repo)
	specs.UpdateHasManyUpdate(t, repo)
	specs.UpdateHasManyReplace(t, repo)
	specs.UpdateHasOneInsert(t, repo)
	specs.UpdateHasOneUpdate(t, repo)
	specs.UpdateBelongsToInsert(t, repo)
	specs.UpdateBelongsToUpdate(t, repo)
	specs.UpdateAtomic(t, repo)
	specs.Updates(t, repo)
	specs.UpdateAny(t, repo)

	// Delete specs
	specs.Delete(t, repo)
	specs.DeleteBelongsTo(t, repo)
	specs.DeleteHasOne(t, repo)
	specs.DeleteHasMany(t, repo)
	specs.DeleteAll(t, repo)
	specs.DeleteAny(t, repo)

	// Constraint specs
	// - Check constraint is not supported by mysql
	specs.UniqueConstraintOnInsert(t, repo)
	specs.UniqueConstraintOnUpdate(t, repo)
	specs.ForeignKeyConstraintOnInsert(t, repo)
	specs.ForeignKeyConstraintOnUpdate(t, repo)
}
