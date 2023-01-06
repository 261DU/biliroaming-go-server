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

// PlayURLCach is an object representing the database table.
type PlayURLCach struct {
	ID         int        `boil:"id" json:"id" toml:"id" yaml:"id"`
	EpisodeID  int64      `boil:"episode_id" json:"episode_id" toml:"episode_id" yaml:"episode_id"`
	IsVip      bool       `boil:"is_vip" json:"is_vip" toml:"is_vip" yaml:"is_vip"`
	Area       int16      `boil:"area" json:"area" toml:"area" yaml:"area"`
	DeviceType int16      `boil:"device_type" json:"device_type" toml:"device_type" yaml:"device_type"`
	FormatType int16      `boil:"format_type" json:"format_type" toml:"format_type" yaml:"format_type"`
	Quality    int16      `boil:"quality" json:"quality" toml:"quality" yaml:"quality"`
	Data       types.JSON `boil:"data" json:"data" toml:"data" yaml:"data"`
	CreatedAt  time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *playURLCachR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L playURLCachL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PlayURLCachColumns = struct {
	ID         string
	EpisodeID  string
	IsVip      string
	Area       string
	DeviceType string
	FormatType string
	Quality    string
	Data       string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	EpisodeID:  "episode_id",
	IsVip:      "is_vip",
	Area:       "area",
	DeviceType: "device_type",
	FormatType: "format_type",
	Quality:    "quality",
	Data:       "data",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var PlayURLCachTableColumns = struct {
	ID         string
	EpisodeID  string
	IsVip      string
	Area       string
	DeviceType string
	FormatType string
	Quality    string
	Data       string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "play_url_caches.id",
	EpisodeID:  "play_url_caches.episode_id",
	IsVip:      "play_url_caches.is_vip",
	Area:       "play_url_caches.area",
	DeviceType: "play_url_caches.device_type",
	FormatType: "play_url_caches.format_type",
	Quality:    "play_url_caches.quality",
	Data:       "play_url_caches.data",
	CreatedAt:  "play_url_caches.created_at",
	UpdatedAt:  "play_url_caches.updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint16) IN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint16) NIN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PlayURLCachWhere = struct {
	ID         whereHelperint
	EpisodeID  whereHelperint64
	IsVip      whereHelperbool
	Area       whereHelperint16
	DeviceType whereHelperint16
	FormatType whereHelperint16
	Quality    whereHelperint16
	Data       whereHelpertypes_JSON
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperint{field: "\"play_url_caches\".\"id\""},
	EpisodeID:  whereHelperint64{field: "\"play_url_caches\".\"episode_id\""},
	IsVip:      whereHelperbool{field: "\"play_url_caches\".\"is_vip\""},
	Area:       whereHelperint16{field: "\"play_url_caches\".\"area\""},
	DeviceType: whereHelperint16{field: "\"play_url_caches\".\"device_type\""},
	FormatType: whereHelperint16{field: "\"play_url_caches\".\"format_type\""},
	Quality:    whereHelperint16{field: "\"play_url_caches\".\"quality\""},
	Data:       whereHelpertypes_JSON{field: "\"play_url_caches\".\"data\""},
	CreatedAt:  whereHelpertime_Time{field: "\"play_url_caches\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"play_url_caches\".\"updated_at\""},
}

// PlayURLCachRels is where relationship names are stored.
var PlayURLCachRels = struct {
}{}

// playURLCachR is where relationships are stored.
type playURLCachR struct {
}

// NewStruct creates a new relationship struct
func (*playURLCachR) NewStruct() *playURLCachR {
	return &playURLCachR{}
}

// playURLCachL is where Load methods for each relationship are stored.
type playURLCachL struct{}

var (
	playURLCachAllColumns            = []string{"id", "episode_id", "is_vip", "area", "device_type", "format_type", "quality", "data", "created_at", "updated_at"}
	playURLCachColumnsWithoutDefault = []string{"episode_id", "is_vip", "area", "device_type", "format_type", "quality", "data", "created_at", "updated_at"}
	playURLCachColumnsWithDefault    = []string{"id"}
	playURLCachPrimaryKeyColumns     = []string{"id"}
	playURLCachGeneratedColumns      = []string{}
)

type (
	// PlayURLCachSlice is an alias for a slice of pointers to PlayURLCach.
	// This should almost always be used instead of []PlayURLCach.
	PlayURLCachSlice []*PlayURLCach
	// PlayURLCachHook is the signature for custom PlayURLCach hook methods
	PlayURLCachHook func(context.Context, boil.ContextExecutor, *PlayURLCach) error

	playURLCachQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	playURLCachType                 = reflect.TypeOf(&PlayURLCach{})
	playURLCachMapping              = queries.MakeStructMapping(playURLCachType)
	playURLCachPrimaryKeyMapping, _ = queries.BindMapping(playURLCachType, playURLCachMapping, playURLCachPrimaryKeyColumns)
	playURLCachInsertCacheMut       sync.RWMutex
	playURLCachInsertCache          = make(map[string]insertCache)
	playURLCachUpdateCacheMut       sync.RWMutex
	playURLCachUpdateCache          = make(map[string]updateCache)
	playURLCachUpsertCacheMut       sync.RWMutex
	playURLCachUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var playURLCachAfterSelectHooks []PlayURLCachHook

var playURLCachBeforeInsertHooks []PlayURLCachHook
var playURLCachAfterInsertHooks []PlayURLCachHook

var playURLCachBeforeUpdateHooks []PlayURLCachHook
var playURLCachAfterUpdateHooks []PlayURLCachHook

var playURLCachBeforeDeleteHooks []PlayURLCachHook
var playURLCachAfterDeleteHooks []PlayURLCachHook

var playURLCachBeforeUpsertHooks []PlayURLCachHook
var playURLCachAfterUpsertHooks []PlayURLCachHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PlayURLCach) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PlayURLCach) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PlayURLCach) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PlayURLCach) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PlayURLCach) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PlayURLCach) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PlayURLCach) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PlayURLCach) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PlayURLCach) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range playURLCachAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPlayURLCachHook registers your hook function for all future operations.
func AddPlayURLCachHook(hookPoint boil.HookPoint, playURLCachHook PlayURLCachHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		playURLCachAfterSelectHooks = append(playURLCachAfterSelectHooks, playURLCachHook)
	case boil.BeforeInsertHook:
		playURLCachBeforeInsertHooks = append(playURLCachBeforeInsertHooks, playURLCachHook)
	case boil.AfterInsertHook:
		playURLCachAfterInsertHooks = append(playURLCachAfterInsertHooks, playURLCachHook)
	case boil.BeforeUpdateHook:
		playURLCachBeforeUpdateHooks = append(playURLCachBeforeUpdateHooks, playURLCachHook)
	case boil.AfterUpdateHook:
		playURLCachAfterUpdateHooks = append(playURLCachAfterUpdateHooks, playURLCachHook)
	case boil.BeforeDeleteHook:
		playURLCachBeforeDeleteHooks = append(playURLCachBeforeDeleteHooks, playURLCachHook)
	case boil.AfterDeleteHook:
		playURLCachAfterDeleteHooks = append(playURLCachAfterDeleteHooks, playURLCachHook)
	case boil.BeforeUpsertHook:
		playURLCachBeforeUpsertHooks = append(playURLCachBeforeUpsertHooks, playURLCachHook)
	case boil.AfterUpsertHook:
		playURLCachAfterUpsertHooks = append(playURLCachAfterUpsertHooks, playURLCachHook)
	}
}

// One returns a single playURLCach record from the query.
func (q playURLCachQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PlayURLCach, error) {
	o := &PlayURLCach{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for play_url_caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PlayURLCach records from the query.
func (q playURLCachQuery) All(ctx context.Context, exec boil.ContextExecutor) (PlayURLCachSlice, error) {
	var o []*PlayURLCach

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PlayURLCach slice")
	}

	if len(playURLCachAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PlayURLCach records in the query.
func (q playURLCachQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count play_url_caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q playURLCachQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if play_url_caches exists")
	}

	return count > 0, nil
}

// PlayURLCaches retrieves all the records using an executor.
func PlayURLCaches(mods ...qm.QueryMod) playURLCachQuery {
	mods = append(mods, qm.From("\"play_url_caches\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"play_url_caches\".*"})
	}

	return playURLCachQuery{q}
}

// FindPlayURLCach retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPlayURLCach(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*PlayURLCach, error) {
	playURLCachObj := &PlayURLCach{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"play_url_caches\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, playURLCachObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from play_url_caches")
	}

	if err = playURLCachObj.doAfterSelectHooks(ctx, exec); err != nil {
		return playURLCachObj, err
	}

	return playURLCachObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PlayURLCach) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no play_url_caches provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(playURLCachColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	playURLCachInsertCacheMut.RLock()
	cache, cached := playURLCachInsertCache[key]
	playURLCachInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			playURLCachAllColumns,
			playURLCachColumnsWithDefault,
			playURLCachColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(playURLCachType, playURLCachMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(playURLCachType, playURLCachMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"play_url_caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"play_url_caches\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into play_url_caches")
	}

	if !cached {
		playURLCachInsertCacheMut.Lock()
		playURLCachInsertCache[key] = cache
		playURLCachInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PlayURLCach.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PlayURLCach) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	playURLCachUpdateCacheMut.RLock()
	cache, cached := playURLCachUpdateCache[key]
	playURLCachUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			playURLCachAllColumns,
			playURLCachPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update play_url_caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"play_url_caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, playURLCachPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(playURLCachType, playURLCachMapping, append(wl, playURLCachPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update play_url_caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for play_url_caches")
	}

	if !cached {
		playURLCachUpdateCacheMut.Lock()
		playURLCachUpdateCache[key] = cache
		playURLCachUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q playURLCachQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for play_url_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for play_url_caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PlayURLCachSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), playURLCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"play_url_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, playURLCachPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in playURLCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all playURLCach")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PlayURLCach) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no play_url_caches provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(playURLCachColumnsWithDefault, o)

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

	playURLCachUpsertCacheMut.RLock()
	cache, cached := playURLCachUpsertCache[key]
	playURLCachUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			playURLCachAllColumns,
			playURLCachColumnsWithDefault,
			playURLCachColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			playURLCachAllColumns,
			playURLCachPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert play_url_caches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(playURLCachPrimaryKeyColumns))
			copy(conflict, playURLCachPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"play_url_caches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(playURLCachType, playURLCachMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(playURLCachType, playURLCachMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert play_url_caches")
	}

	if !cached {
		playURLCachUpsertCacheMut.Lock()
		playURLCachUpsertCache[key] = cache
		playURLCachUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PlayURLCach record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PlayURLCach) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PlayURLCach provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), playURLCachPrimaryKeyMapping)
	sql := "DELETE FROM \"play_url_caches\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from play_url_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for play_url_caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q playURLCachQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no playURLCachQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from play_url_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for play_url_caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PlayURLCachSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(playURLCachBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), playURLCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"play_url_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, playURLCachPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from playURLCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for play_url_caches")
	}

	if len(playURLCachAfterDeleteHooks) != 0 {
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
func (o *PlayURLCach) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPlayURLCach(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PlayURLCachSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PlayURLCachSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), playURLCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"play_url_caches\".* FROM \"play_url_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, playURLCachPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PlayURLCachSlice")
	}

	*o = slice

	return nil
}

// PlayURLCachExists checks if the PlayURLCach row exists.
func PlayURLCachExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"play_url_caches\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if play_url_caches exists")
	}

	return exists, nil
}

// Exists checks if the PlayURLCach row exists.
func (o *PlayURLCach) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PlayURLCachExists(ctx, exec, o.ID)
}
