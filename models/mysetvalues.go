package models

// SetValue represents a set value. (string and binary)
type SetValue struct {
	SetValue   string // enum_value
	ConstValue int    // const_value
}

// MySetValue represents a row from '[custom my_set_value]'.
type MySetValue struct {
	SetValues string // set_values
}

// MySetValues runs a custom query, returning results as MySetValue.
func MySetValues(db XODB, schema string, set string) (*MySetValue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`SUBSTRING(column_type, 5, CHAR_LENGTH(column_type) - 5) AS set_values ` +
		`FROM information_schema.columns ` +
		`WHERE data_type = 'set' AND table_schema = ? AND column_name = ?`

	// run query
	XOLog(sqlstr, schema, set)
	var mev MySetValue
	err = db.QueryRow(sqlstr, schema, set).Scan(&mev.SetValues)
	if err != nil {
		return nil, err
	}

	return &mev, nil
}
