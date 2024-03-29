// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

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

// VerificationToken is an object representing the database table.
type VerificationToken struct {
	ID         string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Identifier string    `boil:"identifier" json:"identifier" toml:"identifier" yaml:"identifier"`
	Token      string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	Expires    time.Time `boil:"expires" json:"expires" toml:"expires" yaml:"expires"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *verificationTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L verificationTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VerificationTokenColumns = struct {
	ID         string
	Identifier string
	Token      string
	Expires    string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	Identifier: "identifier",
	Token:      "token",
	Expires:    "expires",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var VerificationTokenTableColumns = struct {
	ID         string
	Identifier string
	Token      string
	Expires    string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "verification_tokens.id",
	Identifier: "verification_tokens.identifier",
	Token:      "verification_tokens.token",
	Expires:    "verification_tokens.expires",
	CreatedAt:  "verification_tokens.created_at",
	UpdatedAt:  "verification_tokens.updated_at",
}

// Generated where

var VerificationTokenWhere = struct {
	ID         whereHelperstring
	Identifier whereHelperstring
	Token      whereHelperstring
	Expires    whereHelpertime_Time
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperstring{field: "`verification_tokens`.`id`"},
	Identifier: whereHelperstring{field: "`verification_tokens`.`identifier`"},
	Token:      whereHelperstring{field: "`verification_tokens`.`token`"},
	Expires:    whereHelpertime_Time{field: "`verification_tokens`.`expires`"},
	CreatedAt:  whereHelpertime_Time{field: "`verification_tokens`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`verification_tokens`.`updated_at`"},
}

// VerificationTokenRels is where relationship names are stored.
var VerificationTokenRels = struct {
}{}

// verificationTokenR is where relationships are stored.
type verificationTokenR struct {
}

// NewStruct creates a new relationship struct
func (*verificationTokenR) NewStruct() *verificationTokenR {
	return &verificationTokenR{}
}

// verificationTokenL is where Load methods for each relationship are stored.
type verificationTokenL struct{}

var (
	verificationTokenAllColumns            = []string{"id", "identifier", "token", "expires", "created_at", "updated_at"}
	verificationTokenColumnsWithoutDefault = []string{"id", "identifier", "token", "expires"}
	verificationTokenColumnsWithDefault    = []string{"created_at", "updated_at"}
	verificationTokenPrimaryKeyColumns     = []string{"id"}
	verificationTokenGeneratedColumns      = []string{}
)

type (
	// VerificationTokenSlice is an alias for a slice of pointers to VerificationToken.
	// This should almost always be used instead of []VerificationToken.
	VerificationTokenSlice []*VerificationToken
	// VerificationTokenHook is the signature for custom VerificationToken hook methods
	VerificationTokenHook func(context.Context, boil.ContextExecutor, *VerificationToken) error

	verificationTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	verificationTokenType                 = reflect.TypeOf(&VerificationToken{})
	verificationTokenMapping              = queries.MakeStructMapping(verificationTokenType)
	verificationTokenPrimaryKeyMapping, _ = queries.BindMapping(verificationTokenType, verificationTokenMapping, verificationTokenPrimaryKeyColumns)
	verificationTokenInsertCacheMut       sync.RWMutex
	verificationTokenInsertCache          = make(map[string]insertCache)
	verificationTokenUpdateCacheMut       sync.RWMutex
	verificationTokenUpdateCache          = make(map[string]updateCache)
	verificationTokenUpsertCacheMut       sync.RWMutex
	verificationTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var verificationTokenAfterSelectMu sync.Mutex
var verificationTokenAfterSelectHooks []VerificationTokenHook

var verificationTokenBeforeInsertMu sync.Mutex
var verificationTokenBeforeInsertHooks []VerificationTokenHook
var verificationTokenAfterInsertMu sync.Mutex
var verificationTokenAfterInsertHooks []VerificationTokenHook

var verificationTokenBeforeUpdateMu sync.Mutex
var verificationTokenBeforeUpdateHooks []VerificationTokenHook
var verificationTokenAfterUpdateMu sync.Mutex
var verificationTokenAfterUpdateHooks []VerificationTokenHook

var verificationTokenBeforeDeleteMu sync.Mutex
var verificationTokenBeforeDeleteHooks []VerificationTokenHook
var verificationTokenAfterDeleteMu sync.Mutex
var verificationTokenAfterDeleteHooks []VerificationTokenHook

var verificationTokenBeforeUpsertMu sync.Mutex
var verificationTokenBeforeUpsertHooks []VerificationTokenHook
var verificationTokenAfterUpsertMu sync.Mutex
var verificationTokenAfterUpsertHooks []VerificationTokenHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VerificationToken) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VerificationToken) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VerificationToken) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VerificationToken) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VerificationToken) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VerificationToken) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VerificationToken) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VerificationToken) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VerificationToken) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range verificationTokenAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVerificationTokenHook registers your hook function for all future operations.
func AddVerificationTokenHook(hookPoint boil.HookPoint, verificationTokenHook VerificationTokenHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		verificationTokenAfterSelectMu.Lock()
		verificationTokenAfterSelectHooks = append(verificationTokenAfterSelectHooks, verificationTokenHook)
		verificationTokenAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		verificationTokenBeforeInsertMu.Lock()
		verificationTokenBeforeInsertHooks = append(verificationTokenBeforeInsertHooks, verificationTokenHook)
		verificationTokenBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		verificationTokenAfterInsertMu.Lock()
		verificationTokenAfterInsertHooks = append(verificationTokenAfterInsertHooks, verificationTokenHook)
		verificationTokenAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		verificationTokenBeforeUpdateMu.Lock()
		verificationTokenBeforeUpdateHooks = append(verificationTokenBeforeUpdateHooks, verificationTokenHook)
		verificationTokenBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		verificationTokenAfterUpdateMu.Lock()
		verificationTokenAfterUpdateHooks = append(verificationTokenAfterUpdateHooks, verificationTokenHook)
		verificationTokenAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		verificationTokenBeforeDeleteMu.Lock()
		verificationTokenBeforeDeleteHooks = append(verificationTokenBeforeDeleteHooks, verificationTokenHook)
		verificationTokenBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		verificationTokenAfterDeleteMu.Lock()
		verificationTokenAfterDeleteHooks = append(verificationTokenAfterDeleteHooks, verificationTokenHook)
		verificationTokenAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		verificationTokenBeforeUpsertMu.Lock()
		verificationTokenBeforeUpsertHooks = append(verificationTokenBeforeUpsertHooks, verificationTokenHook)
		verificationTokenBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		verificationTokenAfterUpsertMu.Lock()
		verificationTokenAfterUpsertHooks = append(verificationTokenAfterUpsertHooks, verificationTokenHook)
		verificationTokenAfterUpsertMu.Unlock()
	}
}

// One returns a single verificationToken record from the query.
func (q verificationTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VerificationToken, error) {
	o := &VerificationToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for verification_tokens")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all VerificationToken records from the query.
func (q verificationTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (VerificationTokenSlice, error) {
	var o []*VerificationToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to VerificationToken slice")
	}

	if len(verificationTokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all VerificationToken records in the query.
func (q verificationTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count verification_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q verificationTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if verification_tokens exists")
	}

	return count > 0, nil
}

// VerificationTokens retrieves all the records using an executor.
func VerificationTokens(mods ...qm.QueryMod) verificationTokenQuery {
	mods = append(mods, qm.From("`verification_tokens`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`verification_tokens`.*"})
	}

	return verificationTokenQuery{q}
}

// FindVerificationToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVerificationToken(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*VerificationToken, error) {
	verificationTokenObj := &VerificationToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `verification_tokens` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, verificationTokenObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from verification_tokens")
	}

	if err = verificationTokenObj.doAfterSelectHooks(ctx, exec); err != nil {
		return verificationTokenObj, err
	}

	return verificationTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VerificationToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no verification_tokens provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(verificationTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	verificationTokenInsertCacheMut.RLock()
	cache, cached := verificationTokenInsertCache[key]
	verificationTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			verificationTokenAllColumns,
			verificationTokenColumnsWithDefault,
			verificationTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(verificationTokenType, verificationTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(verificationTokenType, verificationTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `verification_tokens` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `verification_tokens` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `verification_tokens` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, verificationTokenPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into verification_tokens")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for verification_tokens")
	}

CacheNoHooks:
	if !cached {
		verificationTokenInsertCacheMut.Lock()
		verificationTokenInsertCache[key] = cache
		verificationTokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the VerificationToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VerificationToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	verificationTokenUpdateCacheMut.RLock()
	cache, cached := verificationTokenUpdateCache[key]
	verificationTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			verificationTokenAllColumns,
			verificationTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update verification_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `verification_tokens` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, verificationTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(verificationTokenType, verificationTokenMapping, append(wl, verificationTokenPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update verification_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for verification_tokens")
	}

	if !cached {
		verificationTokenUpdateCacheMut.Lock()
		verificationTokenUpdateCache[key] = cache
		verificationTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q verificationTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for verification_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for verification_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VerificationTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verificationTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `verification_tokens` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, verificationTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in verificationToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all verificationToken")
	}
	return rowsAff, nil
}

var mySQLVerificationTokenUniqueColumns = []string{
	"id",
	"token",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VerificationToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no verification_tokens provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(verificationTokenColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLVerificationTokenUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	verificationTokenUpsertCacheMut.RLock()
	cache, cached := verificationTokenUpsertCache[key]
	verificationTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			verificationTokenAllColumns,
			verificationTokenColumnsWithDefault,
			verificationTokenColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			verificationTokenAllColumns,
			verificationTokenPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("entity: unable to upsert verification_tokens, could not build update column list")
		}

		ret := strmangle.SetComplement(verificationTokenAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`verification_tokens`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `verification_tokens` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(verificationTokenType, verificationTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(verificationTokenType, verificationTokenMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert for verification_tokens")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(verificationTokenType, verificationTokenMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "entity: unable to retrieve unique values for verification_tokens")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for verification_tokens")
	}

CacheNoHooks:
	if !cached {
		verificationTokenUpsertCacheMut.Lock()
		verificationTokenUpsertCache[key] = cache
		verificationTokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single VerificationToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VerificationToken) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no VerificationToken provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), verificationTokenPrimaryKeyMapping)
	sql := "DELETE FROM `verification_tokens` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from verification_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for verification_tokens")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q verificationTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no verificationTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from verification_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for verification_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VerificationTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(verificationTokenBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verificationTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `verification_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, verificationTokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from verificationToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for verification_tokens")
	}

	if len(verificationTokenAfterDeleteHooks) != 0 {
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
func (o *VerificationToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVerificationToken(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VerificationTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VerificationTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verificationTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `verification_tokens`.* FROM `verification_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, verificationTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in VerificationTokenSlice")
	}

	*o = slice

	return nil
}

// VerificationTokenExists checks if the VerificationToken row exists.
func VerificationTokenExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `verification_tokens` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if verification_tokens exists")
	}

	return exists, nil
}

// Exists checks if the VerificationToken row exists.
func (o *VerificationToken) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return VerificationTokenExists(ctx, exec, o.ID)
}
