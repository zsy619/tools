package xdatabase

import (
	"fmt"
	"strings"
)

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

// sql build where
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
