// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// THSeason2Cach is an object representing the database table.
type THSeason2Cach struct {
	ID        int        `boil:"id" json:"id" toml:"id" yaml:"id"`
	SeasonID  int64      `boil:"season_id" json:"season_id" toml:"season_id" yaml:"season_id"`
	IsVip     bool       `boil:"is_vip" json:"is_vip" toml:"is_vip" yaml:"is_vip"`
	Data      types.JSON `boil:"data" json:"data" toml:"data" yaml:"data"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *thSeason2CachR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thSeason2CachL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var THSeason2CachColumns = struct {
	ID        string
	SeasonID  string
	IsVip     string
	Data      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	SeasonID:  "season_id",
	IsVip:     "is_vip",
	Data:      "data",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var THSeason2CachTableColumns = struct {
	ID        string
	SeasonID  string
	IsVip     string
	Data      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "th_season2_caches.id",
	SeasonID:  "th_season2_caches.season_id",
	IsVip:     "th_season2_caches.is_vip",
	Data:      "th_season2_caches.data",
	CreatedAt: "th_season2_caches.created_at",
	UpdatedAt: "th_season2_caches.updated_at",
}

// Generated where

var THSeason2CachWhere = struct {
	ID        whereHelperint
	SeasonID  whereHelperint64
	IsVip     whereHelperbool
	Data      whereHelpertypes_JSON
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"th_season2_caches\".\"id\""},
	SeasonID:  whereHelperint64{field: "\"th_season2_caches\".\"season_id\""},
	IsVip:     whereHelperbool{field: "\"th_season2_caches\".\"is_vip\""},
	Data:      whereHelpertypes_JSON{field: "\"th_season2_caches\".\"data\""},
	CreatedAt: whereHelpertime_Time{field: "\"th_season2_caches\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"th_season2_caches\".\"updated_at\""},
}

// THSeason2CachRels is where relationship names are stored.
var THSeason2CachRels = struct {
}{}

// thSeason2CachR is where relationships are stored.
type thSeason2CachR struct {
}

// NewStruct creates a new relationship struct
func (*thSeason2CachR) NewStruct() *thSeason2CachR {
	return &thSeason2CachR{}
}

// thSeason2CachL is where Load methods for each relationship are stored.
type thSeason2CachL struct{}

var (
	thSeason2CachAllColumns            = []string{"id", "season_id", "is_vip", "data", "created_at", "updated_at"}
	thSeason2CachColumnsWithoutDefault = []string{"season_id", "is_vip", "data", "created_at", "updated_at"}
	thSeason2CachColumnsWithDefault    = []string{"id"}
	thSeason2CachPrimaryKeyColumns     = []string{"id"}
	thSeason2CachGeneratedColumns      = []string{}
)

type (
	// THSeason2CachSlice is an alias for a slice of pointers to THSeason2Cach.
	// This should almost always be used instead of []THSeason2Cach.
	THSeason2CachSlice []*THSeason2Cach
	// THSeason2CachHook is the signature for custom THSeason2Cach hook methods
	THSeason2CachHook func(context.Context, boil.ContextExecutor, *THSeason2Cach) error

	thSeason2CachQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thSeason2CachType                 = reflect.TypeOf(&THSeason2Cach{})
	thSeason2CachMapping              = queries.MakeStructMapping(thSeason2CachType)
	thSeason2CachPrimaryKeyMapping, _ = queries.BindMapping(thSeason2CachType, thSeason2CachMapping, thSeason2CachPrimaryKeyColumns)
	thSeason2CachInsertCacheMut       sync.RWMutex
	thSeason2CachInsertCache          = make(map[string]insertCache)
	thSeason2CachUpdateCacheMut       sync.RWMutex
	thSeason2CachUpdateCache          = make(map[string]updateCache)
	thSeason2CachUpsertCacheMut       sync.RWMutex
	thSeason2CachUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thSeason2CachAfterSelectHooks []THSeason2CachHook

var thSeason2CachBeforeInsertHooks []THSeason2CachHook
var thSeason2CachAfterInsertHooks []THSeason2CachHook

var thSeason2CachBeforeUpdateHooks []THSeason2CachHook
var thSeason2CachAfterUpdateHooks []THSeason2CachHook

var thSeason2CachBeforeDeleteHooks []THSeason2CachHook
var thSeason2CachAfterDeleteHooks []THSeason2CachHook

var thSeason2CachBeforeUpsertHooks []THSeason2CachHook
var thSeason2CachAfterUpsertHooks []THSeason2CachHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *THSeason2Cach) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *THSeason2Cach) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *THSeason2Cach) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *THSeason2Cach) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *THSeason2Cach) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *THSeason2Cach) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *THSeason2Cach) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *THSeason2Cach) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *THSeason2Cach) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2CachAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTHSeason2CachHook registers your hook function for all future operations.
func AddTHSeason2CachHook(hookPoint boil.HookPoint, thSeason2CachHook THSeason2CachHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		thSeason2CachAfterSelectHooks = append(thSeason2CachAfterSelectHooks, thSeason2CachHook)
	case boil.BeforeInsertHook:
		thSeason2CachBeforeInsertHooks = append(thSeason2CachBeforeInsertHooks, thSeason2CachHook)
	case boil.AfterInsertHook:
		thSeason2CachAfterInsertHooks = append(thSeason2CachAfterInsertHooks, thSeason2CachHook)
	case boil.BeforeUpdateHook:
		thSeason2CachBeforeUpdateHooks = append(thSeason2CachBeforeUpdateHooks, thSeason2CachHook)
	case boil.AfterUpdateHook:
		thSeason2CachAfterUpdateHooks = append(thSeason2CachAfterUpdateHooks, thSeason2CachHook)
	case boil.BeforeDeleteHook:
		thSeason2CachBeforeDeleteHooks = append(thSeason2CachBeforeDeleteHooks, thSeason2CachHook)
	case boil.AfterDeleteHook:
		thSeason2CachAfterDeleteHooks = append(thSeason2CachAfterDeleteHooks, thSeason2CachHook)
	case boil.BeforeUpsertHook:
		thSeason2CachBeforeUpsertHooks = append(thSeason2CachBeforeUpsertHooks, thSeason2CachHook)
	case boil.AfterUpsertHook:
		thSeason2CachAfterUpsertHooks = append(thSeason2CachAfterUpsertHooks, thSeason2CachHook)
	}
}

// One returns a single thSeason2Cach record from the query.
func (q thSeason2CachQuery) One(ctx context.Context, exec boil.ContextExecutor) (*THSeason2Cach, error) {
	o := &THSeason2Cach{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for th_season2_caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all THSeason2Cach records from the query.
func (q thSeason2CachQuery) All(ctx context.Context, exec boil.ContextExecutor) (THSeason2CachSlice, error) {
	var o []*THSeason2Cach

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to THSeason2Cach slice")
	}

	if len(thSeason2CachAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all THSeason2Cach records in the query.
func (q thSeason2CachQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count th_season2_caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thSeason2CachQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if th_season2_caches exists")
	}

	return count > 0, nil
}

// THSeason2Caches retrieves all the records using an executor.
func THSeason2Caches(mods ...qm.QueryMod) thSeason2CachQuery {
	mods = append(mods, qm.From("\"th_season2_caches\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"th_season2_caches\".*"})
	}

	return thSeason2CachQuery{q}
}

// FindTHSeason2Cach retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTHSeason2Cach(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*THSeason2Cach, error) {
	thSeason2CachObj := &THSeason2Cach{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"th_season2_caches\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, thSeason2CachObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from th_season2_caches")
	}

	if err = thSeason2CachObj.doAfterSelectHooks(ctx, exec); err != nil {
		return thSeason2CachObj, err
	}

	return thSeason2CachObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *THSeason2Cach) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season2_caches provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thSeason2CachColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thSeason2CachInsertCacheMut.RLock()
	cache, cached := thSeason2CachInsertCache[key]
	thSeason2CachInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thSeason2CachAllColumns,
			thSeason2CachColumnsWithDefault,
			thSeason2CachColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thSeason2CachType, thSeason2CachMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thSeason2CachType, thSeason2CachMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"th_season2_caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"th_season2_caches\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into th_season2_caches")
	}

	if !cached {
		thSeason2CachInsertCacheMut.Lock()
		thSeason2CachInsertCache[key] = cache
		thSeason2CachInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the THSeason2Cach.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *THSeason2Cach) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thSeason2CachUpdateCacheMut.RLock()
	cache, cached := thSeason2CachUpdateCache[key]
	thSeason2CachUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thSeason2CachAllColumns,
			thSeason2CachPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update th_season2_caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"th_season2_caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, thSeason2CachPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thSeason2CachType, thSeason2CachMapping, append(wl, thSeason2CachPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update th_season2_caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for th_season2_caches")
	}

	if !cached {
		thSeason2CachUpdateCacheMut.Lock()
		thSeason2CachUpdateCache[key] = cache
		thSeason2CachUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thSeason2CachQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for th_season2_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for th_season2_caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o THSeason2CachSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeason2CachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"th_season2_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, thSeason2CachPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in thSeason2Cach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all thSeason2Cach")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *THSeason2Cach) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season2_caches provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thSeason2CachColumnsWithDefault, o)

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

	thSeason2CachUpsertCacheMut.RLock()
	cache, cached := thSeason2CachUpsertCache[key]
	thSeason2CachUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thSeason2CachAllColumns,
			thSeason2CachColumnsWithDefault,
			thSeason2CachColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			thSeason2CachAllColumns,
			thSeason2CachPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert th_season2_caches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(thSeason2CachPrimaryKeyColumns))
			copy(conflict, thSeason2CachPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"th_season2_caches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(thSeason2CachType, thSeason2CachMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thSeason2CachType, thSeason2CachMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert th_season2_caches")
	}

	if !cached {
		thSeason2CachUpsertCacheMut.Lock()
		thSeason2CachUpsertCache[key] = cache
		thSeason2CachUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single THSeason2Cach record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *THSeason2Cach) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no THSeason2Cach provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thSeason2CachPrimaryKeyMapping)
	sql := "DELETE FROM \"th_season2_caches\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from th_season2_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for th_season2_caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thSeason2CachQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no thSeason2CachQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from th_season2_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season2_caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o THSeason2CachSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(thSeason2CachBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeason2CachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"th_season2_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeason2CachPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thSeason2Cach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season2_caches")
	}

	if len(thSeason2CachAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *THSeason2Cach) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTHSeason2Cach(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *THSeason2CachSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := THSeason2CachSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeason2CachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"th_season2_caches\".* FROM \"th_season2_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeason2CachPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in THSeason2CachSlice")
	}

	*o = slice

	return nil
}

// THSeason2CachExists checks if the THSeason2Cach row exists.
func THSeason2CachExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"th_season2_caches\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if th_season2_caches exists")
	}

	return exists, nil
}

// Exists checks if the THSeason2Cach row exists.
func (o *THSeason2Cach) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return THSeason2CachExists(ctx, exec, o.ID)
}
