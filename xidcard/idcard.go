package xidcard

// IdCard 身份证号码
type IdCard string

// NewIdCard
/**
 * @description: 生成身份证号码
 * @param {string} idCard
 * @return {IdCard}
 */
func NewIdCard(idCard string) IdCard {
	return IdCard(idCard)
}

// ToString
/**
 * @description: 转换为字符串
 * @return {string}
 */
func (id IdCard) ToString() string {
	return string(id)
}

// Validate
/**
 * @description: 身份证号码校验
 * @return {bool, error}
 */
func (id IdCard) Validate() (bool, error) {
	if flag, err := Validate(id.ToString()); err != nil {
		return flag, err
	}
	return true, nil
}

// Area
/**
 * @description: 获取地区
 * @return {string, string, error}
 */
func (id IdCard) Area() (string, string, error) {
	return Area(id.ToString())
}

// Birth
/**
 * @description: 获取生日
 * @return {string, error}
 */
func (id IdCard) Birth() (string, error) {
	return Birth(id.ToString())
}

// Sex
/**
 * @description: 获取性别
 * @return {string, string, error}
 */
func (id IdCard) Sex() (string, string, error) {
	return Sex(id.ToString())
}
