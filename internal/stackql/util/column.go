package util

import (
	"github.com/stackql/any-sdk/anysdk"
)

type Column interface {
	GetName() string
	GetSchema() anysdk.Schema
	GetWidth() int
}

type simpleColumn struct {
	name   string
	schema anysdk.Schema
}

func newSimpleColumn(name string, schema anysdk.Schema) Column {
	return &simpleColumn{
		name:   name,
		schema: schema,
	}
}

func newSimpleStringColumn(name string, m anysdk.OperationStore) Column {
	return newSimpleColumn(name, anysdk.NewStringSchema(
		m.GetService(),
		"",
		"",
	),
	)
}

func (sc simpleColumn) GetName() string {
	return sc.name
}

func (sc simpleColumn) GetWidth() int {
	return -1
}

func (sc simpleColumn) GetSchema() anysdk.Schema {
	return sc.schema
}
