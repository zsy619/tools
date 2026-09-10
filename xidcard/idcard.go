package xidcard

// IdCard 表示一个身份证号码。
type IdCard string

// NewIdCard 将字符串 idCard 包装为 IdCard 类型。
// 当 idCard 为空字符串时仍会返回合法的 IdCard，校验请调用 Validate。
func NewIdCard(idCard string) IdCard {
	return IdCard(idCard)
}

// ToString 返回身份证号字符串形式。
func (id IdCard) ToString() string {
	return string(id)
}

// Validate 校验当前身份证号是否合法。
// 校验流程：格式 -> 地区码 -> 出生日期 -> 校验和。
// 返回 (true, nil) 表示通过；任一步失败返回 (false, 错误信息)。
func (id IdCard) Validate() (bool, error) {
	if flag, err := Validate(id.ToString()); err != nil {
		return flag, err
	}
	return true, nil
}

// Area 获取身份证号对应的地区信息。
// 返回 (地区码, 地区名, error)；地区码不存在时返回 ErrAddressInvalid。
func (id IdCard) Area() (string, string, error) {
	return Area(id.ToString())
}

// Birth 获取身份证号中的出生日期，格式为 YYYYMMDD。
// 解析失败或超出合法范围时返回相应错误。
func (id IdCard) Birth() (string, error) {
	return Birth(id.ToString())
}

// BirthYm 获取身份证号中的出生年月，格式为 YYYYMM。
// 内部基于 Birth 实现，错误情况与 Birth 一致。
func (id IdCard) BirthYm() (string, error) {
	ymd, err := Birth(id.ToString())
	if err != nil {
		return "", err
	}
	return ymd[:6], nil
}

// Sex 获取身份证号对应的性别信息。
// 返回 (性别码, "男"/"女", error)；倒数第二位无法解析时返回 ErrSexInvalid。
func (id IdCard) Sex() (string, string, error) {
	return Sex(id.ToString())
}
