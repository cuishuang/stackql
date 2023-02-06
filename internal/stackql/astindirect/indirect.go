package astindirect

import (
	"fmt"

	"github.com/stackql/go-openapistackql/openapistackql"
	"github.com/stackql/stackql/internal/stackql/drm"
	"github.com/stackql/stackql/internal/stackql/symtab"

	"github.com/stackql/stackql/internal/stackql/internal_data_transfer/internaldto"
	"vitess.io/vitess/go/vt/sqlparser"
)

var (
	_ Indirect = &view{}
)

type IndirectType int

const (
	ViewType IndirectType = iota
	SubqueryType
	CTEType
)

func NewViewIndirect(viewDTO internaldto.ViewDTO) (Indirect, error) {
	rv := &view{
		viewDTO:               viewDTO,
		underlyingSymbolTable: symtab.NewHashMapTreeSymTab(),
	}
	return rv, nil
}

func NewSubqueryIndirect(subQuery *sqlparser.Subquery) (Indirect, error) {
	if subQuery == nil {
		return nil, fmt.Errorf("cannot accomodate nil subquery")
	}
	rv := &subquery{
		subQuery:              subQuery,
		selectStmt:            subQuery.Select,
		underlyingSymbolTable: symtab.NewHashMapTreeSymTab(),
	}
	return rv, nil
}

type Indirect interface {
	Parse() error
	GetAssignedParameters() (internaldto.TableParameterCollection, bool)
	GetColumnByName(name string) (internaldto.ColumnMetadata, bool)
	GetColumns() []internaldto.ColumnMetadata
	GetName() string
	GetOptionalParameters() map[string]openapistackql.Addressable
	GetRequiredParameters() map[string]openapistackql.Addressable
	GetSelectAST() sqlparser.SelectStatement
	GetSelectContext() drm.PreparedStatementCtx
	GetType() IndirectType
	GetUnderlyingSymTab() symtab.SymTab
	SetAssignedParameters(internaldto.TableParameterCollection)
	SetSelectContext(drm.PreparedStatementCtx)
	SetUnderlyingSymTab(symtab.SymTab)
}
