// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// Column represents column info.
type Column struct {
	FieldOrdinal int            // field_ordinal
	ColumnName   string         // column_name
	DataType     string         // data_type
	NotNull      bool           // not_null
	DefaultValue sql.NullString // default_value
	IsPrimaryKey bool           // is_primary_key
}

// PgTableColumns runs a custom query, returning results as Column.
func PgTableColumns(db XODB, schema string, table string, sys bool) ([]*Column, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`a.attnum, ` + // ::integer AS field_ordinal
		`a.attname, ` + // ::varchar AS column_name
		`format_type(a.atttypid, a.atttypmod), ` + // ::varchar AS data_type
		`a.attnotnull, ` + // ::boolean AS not_null
		`COALESCE(pg_get_expr(ad.adbin, ad.adrelid), ''), ` + // ::varchar AS default_value
		`COALESCE(ct.contype = 'p', false) ` + // ::boolean AS is_primary_key
		`FROM pg_attribute a ` +
		`JOIN ONLY pg_class c ON c.oid = a.attrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid AND a.attnum = ANY(ct.conkey) AND ct.contype = 'p' ` +
		`LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid AND ad.adnum = a.attnum ` +
		`WHERE a.attisdropped = false AND n.nspname = $1 AND c.relname = $2 AND ($3 OR a.attnum > 0) ` +
		`ORDER BY a.attnum`

	// run query
	XOLog(sqlstr, schema, table, sys)
	q, err := db.Query(sqlstr, schema, table, sys)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Column{}
	for q.Next() {
		c := Column{}

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}

// MyTableColumns runs a custom query, returning results as Column.
func MyTableColumns(db XODB, schema string, table string) ([]*Column, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ordinal_position AS field_ordinal, ` +
		`column_name, ` +
		`IF(data_type = 'enum' OR data_type = 'set', column_name, column_type) AS data_type, ` +
		`IF(is_nullable = 'YES', false, true) AS not_null, ` +
		`column_default AS default_value, ` +
		`IF(column_key = 'PRI', true, false) AS is_primary_key ` +
		`FROM information_schema.columns ` +
		`WHERE table_schema = ? AND table_name = ? ` +
		`ORDER BY ordinal_position`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Column{}
	for q.Next() {
		c := Column{}

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}

// MsTableColumns runs a custom query, returning results as Column.
func MsTableColumns(db XODB, schema string, table string) ([]*Column, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`c.colid AS field_ordinal, ` +
		`c.name AS column_name, ` +
		`TYPE_NAME(c.xtype)+IIF(c.prec > 0, '('+CAST(c.prec AS varchar)+IIF(c.scale > 0,','+CAST(c.scale AS varchar),'')+')', '') as data_type, ` +
		`IIF(c.isnullable=1, 0, 1) AS not_null, ` +
		`x.text AS default_value, ` +
		`IIF(COALESCE(( ` +
		`SELECT count(z.colid) ` +
		`FROM sysindexes i ` +
		`INNER JOIN sysindexkeys z ON i.id = z.id AND i.indid = z.indid AND z.colid = c.colid ` +
		`WHERE i.id = o.id AND i.name = k.name ` +
		`), 0) > 0, 1, 0) AS is_primary_key ` +
		`FROM syscolumns c ` +
		`JOIN sysobjects o ON o.id = c.id ` +
		`LEFT JOIN sysobjects k ON k.xtype='PK' AND k.parent_obj = o.id ` +
		`LEFT JOIN syscomments x ON x.id = c.cdefault ` +
		`WHERE o.type IN('U', 'V') AND SCHEMA_NAME(o.uid) = $1 AND o.name = $2 ` +
		`ORDER BY c.colid`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Column{}
	for q.Next() {
		c := Column{}

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}

// OrTableColumns runs a custom query, returning results as Column.
func OrTableColumns(db XODB, schema string, table string) ([]*Column, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`c.column_id AS field_ordinal, ` +
		`LOWER(c.column_name) AS column_name, ` +
		`LOWER(CASE c.data_type ` +
		`WHEN 'CHAR' THEN 'CHAR('||c.data_length||')' ` +
		`WHEN 'VARCHAR2' THEN 'VARCHAR2('||data_length||')' ` +
		`WHEN 'NUMBER' THEN ` +
		`(CASE WHEN c.data_precision IS NULL AND c.data_scale IS NULL THEN 'NUMBER' ` +
		`ELSE 'NUMBER('||NVL(c.data_precision, 38)||','||NVL(c.data_scale, 0)||')' END) ` +
		`ELSE c.data_type END) AS data_type, ` +
		`CASE WHEN c.nullable = 'N' THEN '1' ELSE '0' END AS not_null, ` +
		`COALESCE((SELECT CASE WHEN r.constraint_type = 'P' THEN '1' ELSE '0' END ` +
		`FROM all_cons_columns l, all_constraints r ` +
		`WHERE r.constraint_type = 'P' AND r.owner = c.owner AND r.table_name = c.table_name AND r.constraint_name = l.constraint_name ` +
		`AND l.owner = c.owner AND l.table_name = c.table_name AND l.column_name = c.column_name), '0') AS is_primary_key ` +
		`FROM all_tab_columns c ` +
		`WHERE c.owner = UPPER(:1) AND c.table_name = UPPER(:2) ` +
		`ORDER BY c.column_id`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Column{}
	for q.Next() {
		c := Column{}

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.IsPrimaryKey)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}
