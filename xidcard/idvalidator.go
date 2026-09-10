package xidcard

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	// ErrFormatInvalid 身份证号整体格式不合法。
	ErrFormatInvalid = errors.New("格式错误")
	// ErrAddressInvalid 身份证号地区码不在已知表中。
	ErrAddressInvalid = errors.New("地址码错误")
	// ErrBirthFormatInvalid 出生日期字段无法解析为 YYYYMMDD。
	ErrBirthFormatInvalid = errors.New("出生日期格式错误")
	// ErrBirthRangeInvalid 出生日期超出 [min_date, max_date] 范围。
	ErrBirthRangeInvalid = errors.New("出生日期范围错误")
	// ErrSumInvalid 身份证号最后一位校验位与计算结果不一致。
	ErrSumInvalid = errors.New("校验和错误")
	// ErrSexInvalid 身份证号倒数第二位无法解析为性别。
	ErrSexInvalid = errors.New("性别错误")

	// reg 身份证号格式正则：6 位地区码 + 年份(18/19/20 可选) + 月日 + 3 位顺序码 + 校验位。
	reg = regexp.MustCompile(`^(\d{6})(18|19|20)?(\d{2})(0\d|10|11|12)([012]\d|3[01])(\d{3})(\d|X)?$`)
	// area 身份证前两位地区码 -> 地区名的映射。
	area = map[string]string{"11": "北京", "12": "天津", "13": "河北", "14": "山西", "15": "内蒙", "21": "辽宁", "22": "吉林", "23": "黑龙", " 31": "上海", "32": "江苏", "33": "浙江", "34": "安徽", "35": "福建", "36": "江西", "37": "山东", "41": "河南", "42": "湖北", "43": "湖南", "44": "广东", "45": "广西", "46": "海南", "50": "重庆", "51": "四川", "52": "贵州", "53": "云南", "54": "西藏", "61": "陕西", "62": "甘肃", "63": "青海", "64": "宁夏", "65": "新疆", "71": "台湾", "81": "香港", "82": "澳门", "91": "国外"}
	// min_date 允许的最小出生日期。
	min_date = time.Date(1890, 0, 0, 0, 0, 0, 0, time.Local)
	// max_date 允许的最大出生日期（当前时间）。
	max_date = time.Now()
	// weight 十七位本体码对应的权重。
	weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2} // 十七位数字本体码权重
	// code 加权和取模 11 后对应的校验位字符表。
	code = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

// ValidateReg 校验身份证号整体格式是否合法。
// 通过时返回 nil，否则返回 ErrFormatInvalid。
func ValidateReg(id string) error {
	if reg.MatchString(id) {
		return nil
	}
	return ErrFormatInvalid
}

// ValidateArea 校验身份证号前两位地区码是否在已知地区表中。
// 校验通过返回 nil，否则返回 ErrAddressInvalid。
func ValidateArea(id string) error {
	if _, ok := area[id[0:2]]; ok {
		return nil
	}
	return ErrAddressInvalid
}

// Area 返回身份证号对应的 (地区码, 地区名, error)。
// 地区码未命中时返回 ("", "", ErrAddressInvalid)。
func Area(id string) (string, string, error) {
	if val, ok := area[id[0:2]]; ok {
		return id[0:2], val, nil
	}
	return "", "", ErrAddressInvalid
}

// ValidateBirth 校验身份证号中的出生日期，包括格式和取值范围。
// id 长度不足 14 会触发 panic；格式错误返回 ErrBirthFormatInvalid，
// 范围错误返回 ErrBirthRangeInvalid。
func ValidateBirth(id string) error {
	birth := id[6:14]
	if date, err := time.Parse("20060102", birth); err != nil {
		return ErrBirthFormatInvalid
	} else if date.After(max_date) && date.Before(min_date) {
		return ErrBirthRangeInvalid
	}
	return nil
}

// Birth 返回身份证号中的出生日期字符串（YYYYMMDD）。
// 解析失败或超出范围时返回相应错误。
func Birth(id string) (string, error) {
	birth := id[6:14]
	if date, err := time.Parse("20060102", birth); err != nil {
		return "", ErrBirthFormatInvalid
	} else if date.After(max_date) && date.Before(min_date) {
		return "", ErrBirthRangeInvalid
	}
	return birth, nil
}

// Sex 根据身份证号倒数第二位判断性别。
// 返回 (性别码, "男"/"女", error)；倒数第二位无法解析时返回 ErrSexInvalid。
func Sex(id string) (string, string, error) {
	idLen := len(id)
	fmt.Println("----------?", idLen)
	idSex := id[idLen-2 : idLen-1]
	sex, err := strconv.Atoi(idSex)
	if err != nil {
		return "", "", ErrSexInvalid
	}
	if sex%2 == 0 {
		return "2", "女", nil
	}
	return "1", "男", nil
}

// ValidateSum 校验身份证号最后一位校验位是否正确。
// 通过返回 nil，否则返回 ErrSumInvalid。
func ValidateSum(id string) error {
	sum := 0
	for i, char := range id[:len(id)-1] {
		char_f, _ := strconv.ParseFloat(string(char), 64)
		sum += int(char_f) * weight[i]
	}
	if code[sum%11] == id[len(id)-1] {
		return nil
	}
	return ErrSumInvalid
}

// Validate 综合校验身份证号：格式 -> 地区码 -> 出生日期 -> 校验和。
// 全部通过返回 (true, nil)；任一步骤失败返回 (false, 对应错误)。
func Validate(id string) (flag bool, err error) {
	// fmt.Println(id)
	if err = ValidateReg(id); err != nil {
		return false, err
	}
	if err = ValidateArea(id); err != nil {
		return false, err
	}
	if err = ValidateBirth(id); err != nil {
		return false, err
	}
	if err = ValidateSum(id); err != nil {
		return false, err
	}
	return true, nil
}
