package xdatabase

import "database/sql"

// DoQuery 将 *sql.Rows 中的全部行扫描为 map[string]interface{} 列表。
//
// 参数：
//   - rows: 已经执行查询得到的 *sql.Rows。
//
// 返回值：
//   - []map[string]interface{}: 每行数据用列名作为键，未命中任何记录时返回空切片。
//   - error: 当前实现始终返回 nil；Scan/Columns 错误被忽略。
//
// 副作用：函数内部会关闭 rows。
func DoQuery(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) // 临时存储每行数据
	for index := range cache {                 // 为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} // 返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) // 取实际类型
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list, nil
}

// DoQuery2 将 *sql.Rows 转换为 map[string]string 列表，同时返回列名。
//
// 参数：
//   - rows: 已经执行查询得到的 *sql.Rows。
//
// 返回值：
//   - result: 每行数据用列名作为键（NULL 值会被忽略）。
//   - columns: 列名列表；获取失败时返回 nil。
//
// panic 条件：rows.Scan 返回错误时会 panic。
//
// 副作用：函数内部会关闭 rows。
func DoQuery2(rows *sql.Rows) (result []map[string]string, columns []string) {
	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// 这个map用来存储一行数据，列名为map的key，map的value为列的值
		rowMap := make(map[string]string)
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col != nil {
				value = string(col)
				rowMap[columns[i]] = value
			}
		}
		result = append(result, rowMap)
	}
	_ = rows.Close()
	return
}

// Golang读取Rows到map[string]interface{}中
// DoQuery3 在 DoQuery 的基础上直接执行 SQL 查询并扫描结果。
//
// 参数：
//   - db: 已初始化的 *sql.DB。
//   - sqlInfo: 待执行的 SQL 语句（带占位符 ?）。
//   - args: SQL 占位符对应的参数列表。
//
// 返回值：
//   - []map[string]interface{}: 命中行转换为 map 列名 -> 值的形式。
//   - error: db.Query 失败时返回错误；后续 Scan/Columns 错误被忽略。
//
// 副作用：函数内部会关闭 rows。
func DoQuery3(db *sql.DB, sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sqlInfo, args...)
	if err != nil {
		return nil, err
	}
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index := range cache {
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{})
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list, nil
}
