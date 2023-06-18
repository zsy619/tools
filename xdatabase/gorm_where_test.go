package xdatabase

import (
	"fmt"
	"testing"
)

func Test_GormWhere1(t *testing.T) {
	where := NewGormWhere()
	where.AddInt("a.zhaopinhuiId", 1)
	where.AddInt("a.checkState", 1)
	where.AddInt("a.is_participant_code", 0)

	{
		where1 := NewGormWhere()
		where1.AddInt("a.zhaopinhuiId", 1)
		where1.AddInt("a.checkState", 1)
		where1.AddInt("a.is_participant_code", 0)
		where.Add("---", where1)
	}
	searchText := "a"
	if searchText != "" {
		where.Add("_string_or", []string{"b.companyName like '%" + searchText + "%'", "a.place_num like '%" + searchText + "%'"})
	}
	where.AddString("exp", "a.participant_time >= ' 00:00:00' and a.participant_time <= ' 23:59:59'")
	wh := where.String()
	fmt.Println(wh)
}

func Test_GormWhere2(t *testing.T) {
	where := NewGormWhere()
	title := "sss"
	count := -1
	name := "xxxx"
	if title != "" {
		where.AddString("exp", "a.record_title like '%"+title+"%'")
	}
	where.AddString(GormWhere_ExpPrefix, "a.record_title like '%"+title+"%'")
	if count > 0 {
		where.AddInt("a.record_count", count)
	}
	if name != "" {
		where.AddString("exp", "b.companyName like '%"+name+"%'")
	}
	wh := where.String()
	fmt.Println(wh)
}
