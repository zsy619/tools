package xdatabase

import "database/sql"

type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func Select(db Queryer, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return SelectScan(rows)
}

func SelectScan(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	numColumns := len(columns)

	values := make([]interface{}, numColumns)
	for i := range values {
		values[i] = new(interface{})
	}

	var results []map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		dest := make(map[string]interface{}, numColumns)
		for i, column := range columns {
			dest[column] = *(values[i].(*interface{}))
		}
		results = append(results, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func Get(db Queryer, query string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return GetScan(rows)
}

func GetScan(rows *sql.Rows) (map[string]interface{}, error) {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	numColumns := len(columns)

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	values := make([]interface{}, numColumns)
	for i := range values {
		values[i] = new(interface{})
	}

	if err := rows.Scan(values...); err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, numColumns)
	for i, column := range columns {
		result[column] = *(values[i].(*interface{}))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func Rows2StringInterfaceMapSlice(rows *sql.Rows, cols ...string) ([]map[string]interface{}, error) {
	var err error
	if len(cols) == 0 {
		cols, err = rows.Columns()
		if err != nil {
			return nil, err
		}
	}
	rowsRes := make([]map[string]interface{}, 0)
	values, args := createEmptyResultSet(len(cols))
	for rows.Next() {
		row := make(map[string]interface{}, len(cols))
		for i := range values {
			values[i] = nil
		}
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		for i, v := range values {
			switch v := v.(type) {
			case []byte:
				row[cols[i]] = string(v)
			default:
				row[cols[i]] = v
			}
		}
		rowsRes = append(rowsRes, row)
	}
	return rowsRes, nil
}

func createEmptyResultSet(numOfCols int) (values []interface{}, args []interface{}) {
	values = make([]interface{}, numOfCols)
	args = make([]interface{}, numOfCols)
	for col := 0; col < numOfCols; col++ {
		args[col] = &values[col]
	}
	return
}
