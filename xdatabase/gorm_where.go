package xdatabase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zsy619/tools/xinterface"
)

const (
	GormWhere_Exp            = "exp"
	GormWhere_ExpPrefix      = GormWhere_Exp + ":::"
	GormWhere_String         = "_string"
	GormWhere_StringPrefix   = GormWhere_String + ":::"
	GormWhere_StringOr       = "_string_or"
	GormWhere_StringOrPrefix = GormWhere_StringOr + ":::"
	GormWhere_Like           = "like"
	GormWhere_LikePrefix     = GormWhere_Like + ":::"
	GormWhere_Or             = "or"
	GormWhere_OrPrefix       = GormWhere_Or + ":::"
	GormWhere_In             = "in"
	GormWhere_InPrefix       = GormWhere_In + ":::"
)

type GormWhere struct {
	where map[string]interface{}
}

func NewGormWhere() *GormWhere {
	return &GormWhere{
		where: make(map[string]interface{}),
	}
}

// 销毁
func (e *GormWhere) Unset(key string) error {
	_, ok := e.where[key]
	if !ok {
		return fmt.Errorf("not exist key %s", key)
	}
	delete(e.where, key)
	return nil
}

// Length 获取长度
func (e *GormWhere) Length() int {
	if e == nil || e.where == nil {
		return 0
	}
	return len(e.where)
}

// Add 添加
func (e *GormWhere) Add(key string, value interface{}) *GormWhere {
	e.where[key] = value
	return e
}

// AddInt 添加整数
func (e *GormWhere) AddInt(key string, value int) *GormWhere {
	e.where[key] = strconv.Itoa(value)
	return e
}

// AddInt32 添加整数
func (e *GormWhere) AddInt32(key string, value int32) *GormWhere {
	e.where[key] = strconv.Itoa(int(value))
	return e
}

// AddInt64 添加整数
func (e *GormWhere) AddInt64(key string, value int64) *GormWhere {
	e.where[key] = strconv.FormatInt(value, 10)
	return e
}

// AddInt64s 添加整数
func (e *GormWhere) AddInt64s(key string, value []int64) *GormWhere {
	var str []string
	for _, v := range value {
		str = append(str, strconv.FormatInt(v, 10))
	}
	e.where[key] = strings.Join(str, ",")
	return e
}

// AddInts 添加整数
func (e *GormWhere) AddInts(key string, value []int) *GormWhere {
	var str []string
	for _, v := range value {
		str = append(str, strconv.Itoa(v))
	}
	e.where[key] = strings.Join(str, ",")
	return e
}

// AddString 添加字符串
func (e *GormWhere) AddString(key string, value string) *GormWhere {
	e.where[key] = e.ForamtString(value)
	return e
}

// AddExp 添加表达式
func (e *GormWhere) AddExp(expKey string, value string) *GormWhere {
	e.where[GormWhere_ExpPrefix+expKey] = e.ForamtString(value)
	return e
}

// ForamtString 格式化字符串
func (e *GormWhere) ForamtString(value string) string {
	return value
	// return SafeString(value)
}

// AddDatetime 添加日期时间
func (e *GormWhere) AddDatetime(key string, value string) *GormWhere {
	e.where[key] = e.ForamtString(value)
	return e
}

// ForamtDatetime 格式化日期时间
func (e *GormWhere) ForamtDatetime(value string) string {
	return ""
}

// String 转换为字符串sql
func (e *GormWhere) String() string {
	where := ""
	for exp, v := range e.where {
		switch v := v.(type) {
		case string:
			if len(where) > 0 {
				switch exp {
				default:
					if strings.HasPrefix(exp, GormWhere_ExpPrefix) {
						exp := v
						if len(exp) > 0 {
							where += " AND " + exp[1:len(exp)-1]
						}
					} else {
						where += " AND " + exp + " = " + v
					}
				case GormWhere_String:
					where += " AND " + v
				case GormWhere_Exp:
					exp := v
					if len(exp) > 0 {
						where += " AND " + exp[1:len(exp)-1]
					}
				}
			} else {
				switch exp {
				default:
					if strings.HasPrefix(exp, GormWhere_ExpPrefix) {
						exp := v
						if len(exp) > 0 {
							where += exp[1 : len(exp)-1]
						}
					} else {
						where = exp + " = " + v
					}
				case GormWhere_String:
					where += v
				case GormWhere_Exp:
					exp := v
					if len(exp) > 0 {
						where += exp[1 : len(exp)-1]
					}
				}
			}
		case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
			val := xinterface.ToString(v)
			if len(where) > 0 {
				switch exp {
				default:
					where += " AND " + exp + " = " + val
				case GormWhere_String:
					where += val
				}
			} else {
				switch exp {
				default:
					where = exp + " = " + val
				case GormWhere_String:
					where += val
				}
			}
		case []string:
			wheres := v
			str := e.toString(exp, wheres)
			if len(str) > 0 {
				where += " AND " + str
			} else {
				where = str
			}
		case [][]string:
			wheres := v
			str := e.toStringMore(exp, wheres)
			if len(str) > 0 {
				where += " AND " + str
			} else {
				where = str
			}
		case *GormWhere:
			wheres := v
			str := wheres.String()
			if len(str) > 0 {
				if len(where) > 0 {
					where += " AND (" + str + ")"
				} else {
					where = str
				}
			} else {
				where = str
			}
		case GormWhere:
			wheres := v
			str := wheres.String()
			if len(str) > 0 {
				if len(where) > 0 {
					where += " AND (" + str + ")"
				} else {
					where = str
				}
			} else {
				where = str
			}
		}
	}
	if len(where) > 0 && strings.HasPrefix(where, " AND ") {
		return where[len(" AND "):]
	}
	return where
}

// toString 转换为字符串s
func (e *GormWhere) toString(field string, wheres []string) string {
	if field == GormWhere_String {
		return wheres[0]
	}
	if field == GormWhere_StringOr {
		return "(" + strings.Join(wheres[:], " or ") + ")"
	}
	length := len(wheres)
	switch {
	default:
		return ""
	case length == 1:
		return field + " = " + wheres[0]
	case length == 2:
		switch strings.ToLower(wheres[0]) {
		default:
			return field + " " + wheres[0] + " " + wheres[1]
		case GormWhere_In:
			return field + " in ( " + strings.Join(wheres[1:], ",") + " )"
		case GormWhere_Like:
			return field + " like '%" + wheres[1] + "%'"
		case GormWhere_Exp:
			return field + " " + wheres[1]
		}
	}
}

// toStringMore 转换为字符串
func (e *GormWhere) toStringMore(field string, wheres [][]string) string {
	length := len(wheres)
	rt := ""
	for i := 0; i < length; i++ {
		wh := wheres[i]
		length := len(wh)
		switch {
		case length == 1: // 如果是第一个元素，则等于 ，否则按逻辑关系处理
			if i == 0 {
				switch field {
				default:
					rt = field + " = " + wh[0]
				case GormWhere_String:
					rt = wh[0]
				case GormWhere_StringOr:
					rt = wh[0]
				}
			} else {
				rt += " " + wh[0]
			}
		case length >= 2: // 运算符表达式 值
			if len(rt) > 0 {
				rt += " and "
			}
			fmt.Println("--------------------------------", field)
			switch field {
			default:
				switch strings.ToLower(wh[0]) {
				default:
					rt += " " + field + " " + wh[0] + " " + wh[1]
				case GormWhere_In:
					rt += " " + field + " " + wh[0] + " (" + wh[1] + ")"
				case GormWhere_Like:
					rt += " " + field + " " + wh[0] + " '%" + wh[1] + "%')"
				}
			case GormWhere_String:
				rt += " " + strings.Join(wh[:], " and ")
			case GormWhere_StringOr:
				fmt.Println("_string_or")
				rt += " " + strings.Join(wh[:], " or ")
			}
		}
	}
	if len(rt) > 0 {
		return "(" + rt + ")"
	}
	return ""
}
