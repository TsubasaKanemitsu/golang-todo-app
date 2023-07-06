// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MTodoTask is an object representing the database table.
type MTodoTask struct { // ID
	ID int `boil:"id" json:"id" toml:"id" yaml:"id"`
	// Todoタイトル
	Title string `boil:"title" json:"title" toml:"title" yaml:"title"`
	// 内容
	Contents null.String `boil:"contents" json:"contents,omitempty" toml:"contents" yaml:"contents,omitempty"`
	// 担当者
	Assignee null.String `boil:"assignee" json:"assignee,omitempty" toml:"assignee" yaml:"assignee,omitempty"`
	// ラベル
	Label null.String `boil:"label" json:"label,omitempty" toml:"label" yaml:"label,omitempty"`
	// 進捗状況
	Status string `boil:"status" json:"status" toml:"status" yaml:"status"`
	// 開始予定日
	StartDate time.Time `boil:"start_date" json:"startDate" toml:"startDate" yaml:"startDate"`
	// 終了予定日
	EndDate time.Time `boil:"end_date" json:"endDate" toml:"endDate" yaml:"endDate"`
	// 作成時刻
	CreatedAt null.Time `boil:"created_at" json:"createdAt,omitempty" toml:"createdAt" yaml:"createdAt,omitempty"`
	// 更新時刻
	UpdatedAt null.Time `boil:"updated_at" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`

	R *mTodoTaskR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mTodoTaskL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MTodoTaskColumns = struct {
	ID        string
	Title     string
	Contents  string
	Assignee  string
	Label     string
	Status    string
	StartDate string
	EndDate   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Title:     "title",
	Contents:  "contents",
	Assignee:  "assignee",
	Label:     "label",
	Status:    "status",
	StartDate: "start_date",
	EndDate:   "end_date",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var MTodoTaskTableColumns = struct {
	ID        string
	Title     string
	Contents  string
	Assignee  string
	Label     string
	Status    string
	StartDate string
	EndDate   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "m_todo_task.id",
	Title:     "m_todo_task.title",
	Contents:  "m_todo_task.contents",
	Assignee:  "m_todo_task.assignee",
	Label:     "m_todo_task.label",
	Status:    "m_todo_task.status",
	StartDate: "m_todo_task.start_date",
	EndDate:   "m_todo_task.end_date",
	CreatedAt: "m_todo_task.created_at",
	UpdatedAt: "m_todo_task.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MTodoTaskWhere = struct {
	ID        whereHelperint
	Title     whereHelperstring
	Contents  whereHelpernull_String
	Assignee  whereHelpernull_String
	Label     whereHelpernull_String
	Status    whereHelperstring
	StartDate whereHelpertime_Time
	EndDate   whereHelpertime_Time
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"m_todo_task\".\"id\""},
	Title:     whereHelperstring{field: "\"m_todo_task\".\"title\""},
	Contents:  whereHelpernull_String{field: "\"m_todo_task\".\"contents\""},
	Assignee:  whereHelpernull_String{field: "\"m_todo_task\".\"assignee\""},
	Label:     whereHelpernull_String{field: "\"m_todo_task\".\"label\""},
	Status:    whereHelperstring{field: "\"m_todo_task\".\"status\""},
	StartDate: whereHelpertime_Time{field: "\"m_todo_task\".\"start_date\""},
	EndDate:   whereHelpertime_Time{field: "\"m_todo_task\".\"end_date\""},
	CreatedAt: whereHelpernull_Time{field: "\"m_todo_task\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"m_todo_task\".\"updated_at\""},
}

// MTodoTaskRels is where relationship names are stored.
var MTodoTaskRels = struct {
}{}

// mTodoTaskR is where relationships are stored.
type mTodoTaskR struct {
}

// NewStruct creates a new relationship struct
func (*mTodoTaskR) NewStruct() *mTodoTaskR {
	return &mTodoTaskR{}
}

// mTodoTaskL is where Load methods for each relationship are stored.
type mTodoTaskL struct{}

var (
	mTodoTaskAllColumns            = []string{"id", "title", "contents", "assignee", "label", "status", "start_date", "end_date", "created_at", "updated_at"}
	mTodoTaskColumnsWithoutDefault = []string{"id", "title", "status", "end_date"}
	mTodoTaskColumnsWithDefault    = []string{"contents", "assignee", "label", "start_date", "created_at", "updated_at"}
	mTodoTaskPrimaryKeyColumns     = []string{"id"}
	mTodoTaskGeneratedColumns      = []string{}
)

type (
	// MTodoTaskSlice is an alias for a slice of pointers to MTodoTask.
	// This should almost always be used instead of []MTodoTask.
	MTodoTaskSlice []*MTodoTask

	mTodoTaskQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mTodoTaskType                 = reflect.TypeOf(&MTodoTask{})
	mTodoTaskMapping              = queries.MakeStructMapping(mTodoTaskType)
	mTodoTaskPrimaryKeyMapping, _ = queries.BindMapping(mTodoTaskType, mTodoTaskMapping, mTodoTaskPrimaryKeyColumns)
	mTodoTaskInsertCacheMut       sync.RWMutex
	mTodoTaskInsertCache          = make(map[string]insertCache)
	mTodoTaskUpdateCacheMut       sync.RWMutex
	mTodoTaskUpdateCache          = make(map[string]updateCache)
	mTodoTaskUpsertCacheMut       sync.RWMutex
	mTodoTaskUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single mTodoTask record from the query.
func (q mTodoTaskQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MTodoTask, error) {
	o := &MTodoTask{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for m_todo_task")
	}

	return o, nil
}

// All returns all MTodoTask records from the query.
func (q mTodoTaskQuery) All(ctx context.Context, exec boil.ContextExecutor) (MTodoTaskSlice, error) {
	var o []*MTodoTask

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to MTodoTask slice")
	}

	return o, nil
}

// Count returns the count of all MTodoTask records in the query.
func (q mTodoTaskQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count m_todo_task rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mTodoTaskQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if m_todo_task exists")
	}

	return count > 0, nil
}

// MTodoTasks retrieves all the records using an executor.
func MTodoTasks(mods ...qm.QueryMod) mTodoTaskQuery {
	mods = append(mods, qm.From("\"m_todo_task\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"m_todo_task\".*"})
	}

	return mTodoTaskQuery{q}
}

// FindMTodoTask retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMTodoTask(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*MTodoTask, error) {
	mTodoTaskObj := &MTodoTask{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"m_todo_task\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mTodoTaskObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from m_todo_task")
	}

	return mTodoTaskObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MTodoTask) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no m_todo_task provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(mTodoTaskColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mTodoTaskInsertCacheMut.RLock()
	cache, cached := mTodoTaskInsertCache[key]
	mTodoTaskInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mTodoTaskAllColumns,
			mTodoTaskColumnsWithDefault,
			mTodoTaskColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mTodoTaskType, mTodoTaskMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mTodoTaskType, mTodoTaskMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"m_todo_task\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"m_todo_task\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into m_todo_task")
	}

	if !cached {
		mTodoTaskInsertCacheMut.Lock()
		mTodoTaskInsertCache[key] = cache
		mTodoTaskInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the MTodoTask.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MTodoTask) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	mTodoTaskUpdateCacheMut.RLock()
	cache, cached := mTodoTaskUpdateCache[key]
	mTodoTaskUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mTodoTaskAllColumns,
			mTodoTaskPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update m_todo_task, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"m_todo_task\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mTodoTaskPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mTodoTaskType, mTodoTaskMapping, append(wl, mTodoTaskPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update m_todo_task row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for m_todo_task")
	}

	if !cached {
		mTodoTaskUpdateCacheMut.Lock()
		mTodoTaskUpdateCache[key] = cache
		mTodoTaskUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q mTodoTaskQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for m_todo_task")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for m_todo_task")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MTodoTaskSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mTodoTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"m_todo_task\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mTodoTaskPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in mTodoTask slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all mTodoTask")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MTodoTask) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no m_todo_task provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(mTodoTaskColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	mTodoTaskUpsertCacheMut.RLock()
	cache, cached := mTodoTaskUpsertCache[key]
	mTodoTaskUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mTodoTaskAllColumns,
			mTodoTaskColumnsWithDefault,
			mTodoTaskColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			mTodoTaskAllColumns,
			mTodoTaskPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert m_todo_task, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mTodoTaskPrimaryKeyColumns))
			copy(conflict, mTodoTaskPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"m_todo_task\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mTodoTaskType, mTodoTaskMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mTodoTaskType, mTodoTaskMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert m_todo_task")
	}

	if !cached {
		mTodoTaskUpsertCacheMut.Lock()
		mTodoTaskUpsertCache[key] = cache
		mTodoTaskUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single MTodoTask record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MTodoTask) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no MTodoTask provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mTodoTaskPrimaryKeyMapping)
	sql := "DELETE FROM \"m_todo_task\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from m_todo_task")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for m_todo_task")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mTodoTaskQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no mTodoTaskQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from m_todo_task")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for m_todo_task")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MTodoTaskSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mTodoTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"m_todo_task\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mTodoTaskPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from mTodoTask slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for m_todo_task")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MTodoTask) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMTodoTask(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MTodoTaskSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MTodoTaskSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mTodoTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"m_todo_task\".* FROM \"m_todo_task\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mTodoTaskPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in MTodoTaskSlice")
	}

	*o = slice

	return nil
}

// MTodoTaskExists checks if the MTodoTask row exists.
func MTodoTaskExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"m_todo_task\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if m_todo_task exists")
	}

	return exists, nil
}

// Exists checks if the MTodoTask row exists.
func (o *MTodoTask) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MTodoTaskExists(ctx, exec, o.ID)
}