package xdatabase

import (
	"fmt"
	"strconv"
	"strings"
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

func (e *GormWhere) Length() int {
	if e == nil || e.where == nil {
		return 0
	}
	return len(e.where)
}

func (e *GormWhere) Add(key string, value interface{}) {
	e.where[key] = value
}

func (e *GormWhere) AddInt(key string, value int) {
	e.where[key] = strconv.Itoa(value)
}

func (e *GormWhere) AddString(key string, value string) {
	e.where[key] = e.ForamtString(value)
}

func (e *GormWhere) AddExp(expKey string, value string) {
	e.where[GormWhere_ExpPrefix+expKey] = e.ForamtString(value)
}

func (e *GormWhere) ForamtString(value string) string {
	return "'" + value + "'"
}

func (e *GormWhere) AddDatetime(key string, value string) {
	e.where[key] = e.ForamtString(value)
}

func (e *GormWhere) ForamtDatetime(value string) string {
	return ""
}

func (e *GormWhere) String() string {
	where := ""
	for f, v := range e.where {
		switch v.(type) {
		case string:
			if len(where) > 0 {
				switch f {
				default:
					if strings.HasPrefix(f, GormWhere_ExpPrefix) {
						exp := v.(string)
						if len(exp) > 0 {
							where += " AND " + exp[1:len(exp)-1]
						}
					} else {
						where += " AND " + f + " = " + v.(string)
					}
				case GormWhere_String:
					where += " AND " + v.(string)
				case GormWhere_Exp:
					exp := v.(string)
					if len(exp) > 0 {
						where += " AND " + exp[1:len(exp)-1]
					}
				}
			} else {
				switch f {
				default:
					if strings.HasPrefix(f, GormWhere_ExpPrefix) {
						exp := v.(string)
						if len(exp) > 0 {
							where += exp[1 : len(exp)-1]
						}
					} else {
						where = f + " = " + v.(string)
					}
				case GormWhere_String:
					where += v.(string)
				case GormWhere_Exp:
					exp := v.(string)
					if len(exp) > 0 {
						where += exp[1 : len(exp)-1]
					}
				}
			}
		case int:
			if len(where) > 0 {
				switch f {
				default:
					where += " AND " + f + " = " + strconv.Itoa(v.(int))
				case GormWhere_String:
					where += strconv.Itoa(v.(int))
				}
			} else {
				switch f {
				default:
					where = f + " = " + strconv.Itoa(v.(int))
				case GormWhere_String:
					where += strconv.Itoa(v.(int))
				}
			}
		case []string:
			wheres := v.([]string)
			str := e.toString(f, wheres)
			if len(str) > 0 {
				where += " AND " + str
			} else {
				where = str
			}
		case [][]string:
			wheres := v.([][]string)
			str := e.toStringMore(f, wheres)
			if len(str) > 0 {
				where += " AND " + str
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
