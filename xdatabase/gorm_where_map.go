package xdatabase

import (
	"fmt"
	"strings"
)

// NullType 表示 SQL NULL 条件的特殊标记。
type NullType byte

const (
	_ NullType = iota
	// IsNull 等价于 SQL 中的 "IS NULL"。
	IsNull
	// IsNotNull 等价于 SQL 中的 "IS NOT NULL"。
	IsNotNull
)

// sql build where
// WhereBuild 根据 map 构造 WHERE 子句 SQL 与对应参数。
//
// 参数：
//   - where: 键可以是 "field"、"field op" 或 "field op xxx" 形式；值可以为 NullType 表示 IS NULL/IS NOT NULL。
//
// 返回值：
//   - whereSQL: 拼接后的 SQL 片段，多个条件以 " AND " 连接。
//   - vals: 与 whereSQL 中 "?" 占位符一一对应的参数列表。
//   - err: 键格式不合法（拆出超过 2 段）时返回错误。
func WhereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			outerr := fmt.Errorf("error in query condition: %s. ", k)
			return "", nil, outerr
		}

		if whereSQL != "" {
			whereSQL += " AND "
		}
		strings.Join(ks, ",")
		switch len(ks) {
		case 1:
			// fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case NullType:
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				vals = append(vals, v)
			}
		}
	}
	return
}

// WhereBuildExtension 是 WhereBuild 的扩展版本，根据值的类型拼接 WHERE SQL 片段（用于手工拼装原生 SQL）。
//
// 注意：返回值是直接拼接进 SQL 的字符串，请确保 map 中的 key/字符串值已做安全过滤，
// 避免 SQL 注入风险；当前实现仅对单引号做了 "”" 转义。
//
// 参数：
//   - where: 键值对 map；key 作为字段名，value 根据类型决定拼接方式（string/int/[]int/[]float 等）。
//
// 返回值：去掉最前 " and " 前缀后的 WHERE 片段。
func WhereBuildExtension(where map[string]interface{}) (result string) {
	for k, v := range where {
		switch v := v.(type) {
		case string:
			value := v
			if value != "" {
				if k == "or" || k == "and" {
					if result == "" {
						result = " and " + value + " " + k
					} else {
						result += " " + k + "  " + value
					}
				} else if strings.HasPrefix(value, "<") || strings.HasPrefix(value, "<=") ||
					strings.HasPrefix(value, ">") || strings.HasPrefix(value, ">=") ||
					strings.HasPrefix(value, "<>") || strings.HasPrefix(value, "between") ||
					strings.HasPrefix(value, "in") || strings.HasPrefix(value, "not in") ||
					strings.HasPrefix(value, "like") || strings.HasPrefix(value, "not like") {
					result += " and " + k + " " + value
				} else {
					result += " and " + k + " = '" + strings.Replace(value, "'", "''", -1) + "'"
				}
			}
		case []string:
			vv := fmt.Sprintf("%s", v)
			vv = strings.Replace(vv, "'", "''", -1)
			result += " and " + k + " in " + strings.Replace(strings.Replace(strings.Replace(vv, " ", "','", -1), "[", "('", -1), "]", "')", -1)
		case int, int16, int32, int64, uint, uint16, uint32, uint64, uint8:
			result += " and " + k + " = " + fmt.Sprintf("%d", v)
		case []int, []int16, []int32, []int64, []uint, []uint16, []uint32, []uint64, []uint8:
			result += " and " + k + " in " + strings.Replace(strings.Replace(strings.Replace(fmt.Sprintf("%d", v), " ", ",", -1), "[", "(", -1), "]", ")", -1)
		case []float32, []float64:
			result += " and " + k + " in " + strings.Replace(strings.Replace(strings.Replace(fmt.Sprintf("%f", v), " ", ",", -1), "[", "(", -1), "]", ")", -1)
		case float32, float64:
			result += " and " + k + " = " + fmt.Sprintf("%f", v)
		}
	}
	if result != "" {
		result = result[5:]
	}
	return
}
