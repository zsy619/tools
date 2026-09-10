package xdatabase

import (
	"database/sql"
	"time"
)

// Queryer 抽象了数据库连接的最小查询能力，*sql.DB 与 *sql.Tx 都满足该接口。
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

/**
 * @description: 查询多条数据
 * @param {Queryer} db 数据库连接
 * @param {string} query 查询语句
 * @param {...interface{}} args 查询参数
 * @return {*}
 */
func Select(db Queryer, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return SelectScan(rows)
}

/**
 * @description: 查询多条数据
 * @param {*sql.Rows} rows 查询结果
 * @return {*}
 */
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

// Get 通过 Queryer 执行查询并返回第一条记录。
//
// 参数：
//   - db: 实现 Queryer 接口的数据库连接或事务。
//   - query: 待执行的 SQL 语句（带占位符 ?）。
//   - args: 占位符对应的参数列表。
//
// 返回值：
//   - map[string]interface{}: 命中行的列名 -> 值映射；无记录时为 nil。
//   - error: db.Query 失败时返回错误；无记录时 GetScan 返回 sql.ErrNoRows。
func Get(db Queryer, query string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return GetScan(rows)
}

// GetScan 从 *sql.Rows 中读取第一条记录并扫描为 map。
//
// 参数：
//   - rows: 已经执行查询得到的 *sql.Rows。
//
// 返回值：
//   - map[string]interface{}: 命中行的列名 -> 值映射。
//   - error: 无记录时返回 sql.ErrNoRows；其他错误来自 Columns/Scan/Rows.Err。
//
// 副作用：函数内部会关闭 rows。
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

// Rows2StringInterfaceMapSlice 将 *sql.Rows 转换为 map 列表，并对常见类型做规范化（[]byte 转 string、time.Time 格式化）。
//
// 参数：
//   - rows: 已经执行查询得到的 *sql.Rows。
//   - cols: 可选列名列表；为空时使用 rows.Columns() 自动获取。
//
// 返回值：
//   - []map[string]interface{}: 每行数据按 cols 顺序构造为 map。
//   - error: rows.Columns()/rows.Scan 失败时返回错误。
//
// 注意：函数不会关闭 rows，调用方需自行处理。
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
			case string:
				row[cols[i]] = string(v)
			case int64:
				row[cols[i]] = int64(v)
			case int32:
				row[cols[i]] = int32(v)
			case int16:
				row[cols[i]] = int16(v)
			case int8:
				row[cols[i]] = int8(v)
			case int:
				row[cols[i]] = int(v)
			case float32:
				row[cols[i]] = float32(v)
			case float64:
				row[cols[i]] = float64(v)
			case bool:
				row[cols[i]] = bool(v)
			case time.Time:
				row[cols[i]] = v.Format("2006-01-02 15:04:05")
			case nil:
				row[cols[i]] = ""
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
