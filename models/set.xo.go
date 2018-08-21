package models

// Set represents a set.
type Set struct {
	SetName string // set_name
}

// MySets runs a custom query, returning results as Set.
func MySets(db XODB, schema string) ([]*Set, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`DISTINCT column_name AS set_name ` +
		`FROM information_schema.columns ` +
		`WHERE data_type = 'set' AND table_schema = ?`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Set{}
	for q.Next() {
		e := Set{}

		// scan
		err = q.Scan(&e.SetName)
		if err != nil {
			return nil, err
		}

		res = append(res, &e)
	}

	return res, nil
}
