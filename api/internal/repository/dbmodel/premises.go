// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodel

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
	"github.com/volatiletech/strmangle"
)

// Premise is an object representing the database table.
type Premise struct {
	ID           int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name         string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Location     string    `boil:"location" json:"location" toml:"location" yaml:"location"`
	PremisesCode string    `boil:"premises_code" json:"premises_code" toml:"premises_code" yaml:"premises_code"`
	Description  string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	CCTVCount    int       `boil:"cctv_count" json:"cctv_count" toml:"cctv_count" yaml:"cctv_count"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *premiseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L premiseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PremiseColumns = struct {
	ID           string
	Name         string
	Location     string
	PremisesCode string
	Description  string
	CCTVCount    string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	Name:         "name",
	Location:     "location",
	PremisesCode: "premises_code",
	Description:  "description",
	CCTVCount:    "cctv_count",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var PremiseTableColumns = struct {
	ID           string
	Name         string
	Location     string
	PremisesCode string
	Description  string
	CCTVCount    string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "premises.id",
	Name:         "premises.name",
	Location:     "premises.location",
	PremisesCode: "premises.premises_code",
	Description:  "premises.description",
	CCTVCount:    "premises.cctv_count",
	CreatedAt:    "premises.created_at",
	UpdatedAt:    "premises.updated_at",
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

var PremiseWhere = struct {
	ID           whereHelperint64
	Name         whereHelperstring
	Location     whereHelperstring
	PremisesCode whereHelperstring
	Description  whereHelperstring
	CCTVCount    whereHelperint
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
}{
	ID:           whereHelperint64{field: "\"premises\".\"id\""},
	Name:         whereHelperstring{field: "\"premises\".\"name\""},
	Location:     whereHelperstring{field: "\"premises\".\"location\""},
	PremisesCode: whereHelperstring{field: "\"premises\".\"premises_code\""},
	Description:  whereHelperstring{field: "\"premises\".\"description\""},
	CCTVCount:    whereHelperint{field: "\"premises\".\"cctv_count\""},
	CreatedAt:    whereHelpertime_Time{field: "\"premises\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"premises\".\"updated_at\""},
}

// PremiseRels is where relationship names are stored.
var PremiseRels = struct {
	CCTVDevices string
}{
	CCTVDevices: "CCTVDevices",
}

// premiseR is where relationships are stored.
type premiseR struct {
	CCTVDevices CCTVDeviceSlice `boil:"CCTVDevices" json:"CCTVDevices" toml:"CCTVDevices" yaml:"CCTVDevices"`
}

// NewStruct creates a new relationship struct
func (*premiseR) NewStruct() *premiseR {
	return &premiseR{}
}

func (r *premiseR) GetCCTVDevices() CCTVDeviceSlice {
	if r == nil {
		return nil
	}
	return r.CCTVDevices
}

// premiseL is where Load methods for each relationship are stored.
type premiseL struct{}

var (
	premiseAllColumns            = []string{"id", "name", "location", "premises_code", "description", "cctv_count", "created_at", "updated_at"}
	premiseColumnsWithoutDefault = []string{"id", "name", "location", "premises_code", "description", "cctv_count"}
	premiseColumnsWithDefault    = []string{"created_at", "updated_at"}
	premisePrimaryKeyColumns     = []string{"id"}
	premiseGeneratedColumns      = []string{}
)

type (
	// PremiseSlice is an alias for a slice of pointers to Premise.
	// This should almost always be used instead of []Premise.
	PremiseSlice []*Premise

	premiseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	premiseType                 = reflect.TypeOf(&Premise{})
	premiseMapping              = queries.MakeStructMapping(premiseType)
	premisePrimaryKeyMapping, _ = queries.BindMapping(premiseType, premiseMapping, premisePrimaryKeyColumns)
	premiseInsertCacheMut       sync.RWMutex
	premiseInsertCache          = make(map[string]insertCache)
	premiseUpdateCacheMut       sync.RWMutex
	premiseUpdateCache          = make(map[string]updateCache)
	premiseUpsertCacheMut       sync.RWMutex
	premiseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single premise record from the query.
func (q premiseQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Premise, error) {
	o := &Premise{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: failed to execute a one query for premises")
	}

	return o, nil
}

// All returns all Premise records from the query.
func (q premiseQuery) All(ctx context.Context, exec boil.ContextExecutor) (PremiseSlice, error) {
	var o []*Premise

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodel: failed to assign all query results to Premise slice")
	}

	return o, nil
}

// Count returns the count of all Premise records in the query.
func (q premiseQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to count premises rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q premiseQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: failed to check if premises exists")
	}

	return count > 0, nil
}

// CCTVDevices retrieves all the cctv_device's CCTVDevices with an executor.
func (o *Premise) CCTVDevices(mods ...qm.QueryMod) cctvDeviceQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"cctv_devices\".\"premise_id\"=?", o.ID),
	)

	return CCTVDevices(queryMods...)
}

// LoadCCTVDevices allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (premiseL) LoadCCTVDevices(ctx context.Context, e boil.ContextExecutor, singular bool, maybePremise interface{}, mods queries.Applicator) error {
	var slice []*Premise
	var object *Premise

	if singular {
		var ok bool
		object, ok = maybePremise.(*Premise)
		if !ok {
			object = new(Premise)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePremise)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePremise))
			}
		}
	} else {
		s, ok := maybePremise.(*[]*Premise)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePremise)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePremise))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &premiseR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &premiseR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`cctv_devices`),
		qm.WhereIn(`cctv_devices.premise_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load cctv_devices")
	}

	var resultSlice []*CCTVDevice
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice cctv_devices")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on cctv_devices")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for cctv_devices")
	}

	if singular {
		object.R.CCTVDevices = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &cctvDeviceR{}
			}
			foreign.R.Premise = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PremiseID {
				local.R.CCTVDevices = append(local.R.CCTVDevices, foreign)
				if foreign.R == nil {
					foreign.R = &cctvDeviceR{}
				}
				foreign.R.Premise = local
				break
			}
		}
	}

	return nil
}

// AddCCTVDevices adds the given related objects to the existing relationships
// of the premise, optionally inserting them as new records.
// Appends related to o.R.CCTVDevices.
// Sets related.R.Premise appropriately.
func (o *Premise) AddCCTVDevices(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CCTVDevice) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PremiseID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"cctv_devices\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"premise_id"}),
				strmangle.WhereClause("\"", "\"", 2, cctvDevicePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PremiseID = o.ID
		}
	}

	if o.R == nil {
		o.R = &premiseR{
			CCTVDevices: related,
		}
	} else {
		o.R.CCTVDevices = append(o.R.CCTVDevices, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &cctvDeviceR{
				Premise: o,
			}
		} else {
			rel.R.Premise = o
		}
	}
	return nil
}

// Premises retrieves all the records using an executor.
func Premises(mods ...qm.QueryMod) premiseQuery {
	mods = append(mods, qm.From("\"premises\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"premises\".*"})
	}

	return premiseQuery{q}
}

// FindPremise retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPremise(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Premise, error) {
	premiseObj := &Premise{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"premises\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, premiseObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: unable to select from premises")
	}

	return premiseObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Premise) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no premises provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(premiseColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	premiseInsertCacheMut.RLock()
	cache, cached := premiseInsertCache[key]
	premiseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			premiseAllColumns,
			premiseColumnsWithDefault,
			premiseColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(premiseType, premiseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(premiseType, premiseMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"premises\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"premises\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dbmodel: unable to insert into premises")
	}

	if !cached {
		premiseInsertCacheMut.Lock()
		premiseInsertCache[key] = cache
		premiseInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Premise.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Premise) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	premiseUpdateCacheMut.RLock()
	cache, cached := premiseUpdateCache[key]
	premiseUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			premiseAllColumns,
			premisePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodel: unable to update premises, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"premises\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, premisePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(premiseType, premiseMapping, append(wl, premisePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodel: unable to update premises row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by update for premises")
	}

	if !cached {
		premiseUpdateCacheMut.Lock()
		premiseUpdateCache[key] = cache
		premiseUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q premiseQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all for premises")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected for premises")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PremiseSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodel: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), premisePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"premises\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, premisePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all in premise slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected all in update all premise")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Premise) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no premises provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(premiseColumnsWithDefault, o)

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

	premiseUpsertCacheMut.RLock()
	cache, cached := premiseUpsertCache[key]
	premiseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			premiseAllColumns,
			premiseColumnsWithDefault,
			premiseColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			premiseAllColumns,
			premisePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodel: unable to upsert premises, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(premisePrimaryKeyColumns))
			copy(conflict, premisePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"premises\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(premiseType, premiseMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(premiseType, premiseMapping, ret)
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
		return errors.Wrap(err, "dbmodel: unable to upsert premises")
	}

	if !cached {
		premiseUpsertCacheMut.Lock()
		premiseUpsertCache[key] = cache
		premiseUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Premise record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Premise) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodel: no Premise provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), premisePrimaryKeyMapping)
	sql := "DELETE FROM \"premises\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete from premises")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by delete for premises")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q premiseQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodel: no premiseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from premises")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for premises")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PremiseSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), premisePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"premises\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, premisePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from premise slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for premises")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Premise) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPremise(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PremiseSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PremiseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), premisePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"premises\".* FROM \"premises\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, premisePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to reload all in PremiseSlice")
	}

	*o = slice

	return nil
}

// PremiseExists checks if the Premise row exists.
func PremiseExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"premises\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: unable to check if premises exists")
	}

	return exists, nil
}

// Exists checks if the Premise row exists.
func (o *Premise) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PremiseExists(ctx, exec, o.ID)
}
