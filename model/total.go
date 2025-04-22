package model

import (
	"context"
	"fmt"

	"github.com/hysios/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Scope = *gorm.DB

var selectTotalKey = struct{}{}

type mysqlSelectTotals struct {
	tx    Scope
	Debug bool
}

func (*mysqlSelectTotals) ModifyStatement(stmt *gorm.Statement) {
	selectClause := stmt.Clauses["SELECT"]
	if selectClause.AfterExpression != nil {
		stmt.Selects = append([]string{"SQL_CALC_FOUND_ROWS *,"}, stmt.Selects...)
	} else {
		stmt.Selects = append([]string{"SQL_CALC_FOUND_ROWS *"}, stmt.Selects...)
	}
}

func (*mysqlSelectTotals) Build(builder clause.Builder) {
	log.Infof("builder %v", builder)
}

func (totals *mysqlSelectTotals) Total() (int, error) {
	var total int
	if err := totals.tx.Raw("SELECT FOUND_ROWS()").Scan(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func DebugSQL(debug bool, scope Scope, fn func(scope Scope) Scope) error {
	if debug {
		sql := scope.ToSQL(func(tx *gorm.DB) *gorm.DB {
			scope = fn(tx)

			return scope
		})
		fmt.Printf("SQL: %s", sql)
		return nil
	} else {
		scope = fn(scope)
	}

	return scope.Error
}

type GetTotal interface {
	Total() (int, error)
}

func GetScopeTotal(scope Scope) (GetTotal, bool) {
	if totals, ok := scope.Statement.Context.Value(selectTotalKey).(GetTotal); ok {
		return totals, ok
	}

	return nil, false
}

func WithTotal(scope Scope, debug bool) (GetTotal, Scope) {
	var totals = &mysqlSelectTotals{tx: scope, Debug: debug}
	scope = scope.WithContext(context.WithValue(scope.Statement.Context, selectTotalKey, totals))
	return totals, scope.Clauses(totals)
}
