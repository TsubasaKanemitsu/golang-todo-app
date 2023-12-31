// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersions)
	t.Run("MTodoTasks", testMTodoTasks)
}

func TestDelete(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsDelete)
	t.Run("MTodoTasks", testMTodoTasksDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsQueryDeleteAll)
	t.Run("MTodoTasks", testMTodoTasksQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsSliceDeleteAll)
	t.Run("MTodoTasks", testMTodoTasksSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsExists)
	t.Run("MTodoTasks", testMTodoTasksExists)
}

func TestFind(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsFind)
	t.Run("MTodoTasks", testMTodoTasksFind)
}

func TestBind(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsBind)
	t.Run("MTodoTasks", testMTodoTasksBind)
}

func TestOne(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsOne)
	t.Run("MTodoTasks", testMTodoTasksOne)
}

func TestAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsAll)
	t.Run("MTodoTasks", testMTodoTasksAll)
}

func TestCount(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsCount)
	t.Run("MTodoTasks", testMTodoTasksCount)
}

func TestInsert(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsInsert)
	t.Run("GooseDBVersions", testGooseDBVersionsInsertWhitelist)
	t.Run("MTodoTasks", testMTodoTasksInsert)
	t.Run("MTodoTasks", testMTodoTasksInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsReload)
	t.Run("MTodoTasks", testMTodoTasksReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsReloadAll)
	t.Run("MTodoTasks", testMTodoTasksReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsSelect)
	t.Run("MTodoTasks", testMTodoTasksSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsUpdate)
	t.Run("MTodoTasks", testMTodoTasksUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsSliceUpdateAll)
	t.Run("MTodoTasks", testMTodoTasksSliceUpdateAll)
}
