package xmath

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

func Decimal0(value float64) float64 {
	return DecimalDigit(value, 0)
}

func Decimal2(value float64) float64 {
	return DecimalDigit(value, 2)
}

func Decimal1(value float64) float64 {
	return DecimalDigit(value, 1)
}

// DecimalDigit
/**
 * @description: 保留小数位数
 * @param {float64} value
 * @param {int} digit
 * @return {*}
 */
func DecimalDigit[T ~float32 | ~float64](value T, digit int) T {
	if digit < 0 {
		return value
	}
	format := "%." + strconv.Itoa(digit) + "f"
	outValue, _ := strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return T(outValue)
}

// DivisionPrecision 是当运算结果无法精确整除时，结果保留的小数位数。
//
// 示例：
//
//	d1 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//	d1.String() // output: "0.6666666666666667"
//	d2 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(30000))
//	d2.String() // output: "0.0000666666666667"
//	d3 := decimal.NewFromFloat(20000).Div(decimal.NewFromFloat(3))
//	d3.String() // output: "6666.6666666666666667"
//	decimal.DivisionPrecision = 3
//	d4 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//	d4.String() // output: "0.667"
var DivisionPrecision = 16

// MarshalJSONWithoutQuotes 若希望 Decimal 以数字而非字符串的形式进行 JSON 序列化，
// 可将该变量设置为 true。
// 警告：对位数很多的 Decimal 而言这样做有风险——许多 JSON
// 反序列化器（例如 JavaScript 的）会把 JSON 数字解析为 IEEE 754
// 双精度浮点数，这意味着你可能在
// 不知不觉中丢失精度。
var MarshalJSONWithoutQuotes = false

// ExpMaxIterations 指定使用 ExpHullAbrham 方法计算精确的自然指数时
// 所需的最大迭代次数。
var ExpMaxIterations = 1000

// Zero 常量，用于加快运算速度。
// 注意：Zero 不应直接与 == 或 != 比较，请改用 decimal.Equal 或 decimal.Cmp。
var Zero = New(0, 1)

var (
	zeroInt   = big.NewInt(0)
	oneInt    = big.NewInt(1)
	twoInt    = big.NewInt(2)
	fourInt   = big.NewInt(4)
	fiveInt   = big.NewInt(5)
	tenInt    = big.NewInt(10)
	twentyInt = big.NewInt(20)
)

var factorials = []Decimal{New(1, 0)}

// Decimal 表示一个定点小数，且是不可变的。
// 其数值为 value * 10 ^ exp。
type Decimal struct {
	value *big.Int

	// NOTE(vadim)：该字段必须为 int32，因为在计算过程中会将其转换为 float64。
	// 若 exp 是 64 位，可能会丢失精度。
	// 若我们希望可以表示任意可能的小数，可将 exp 设为 *big.Int，
	// 但那样会影响性能，而且如此大的数字在实际中并不现实。
	exp int32
}

// New 返回一个新的定点小数，其值为 value * 10 ^ exp。
func New(value int64, exp int32) Decimal {
	return Decimal{
		value: big.NewInt(value),
		exp:   exp,
	}
}

// NewFromInt 将 int64 转换为 Decimal。
//
// 示例：
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(value int64) Decimal {
	return Decimal{
		value: big.NewInt(value),
		exp:   0,
	}
}

// NewFromInt32 将 int32 转换为 Decimal。
//
// 示例：
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(value int32) Decimal {
	return Decimal{
		value: big.NewInt(int64(value)),
		exp:   0,
	}
}

// NewFromBigInt 根据 big.Int 返回新的 Decimal，其值为 value * 10 ^ exp。
func NewFromBigInt(value *big.Int, exp int32) Decimal {
	return Decimal{
		value: new(big.Int).Set(value),
		exp:   exp,
	}
}

// NewFromString 根据字符串表示形式返回新的 Decimal。
// 末尾的零不会被去除。
//
// 示例：
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(value string) (Decimal, error) {
	originalInput := value
	var intString string
	var exp int64

	// 检查该数字是否使用了科学计数法
	eIndex := strings.IndexAny(value, "Ee")
	if eIndex != -1 {
		expInt, err := strconv.ParseInt(value[eIndex+1:], 10, 32)
		if err != nil {
			if e, ok := err.(*strconv.NumError); ok && e.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", value)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", value)
		}
		value = value[:eIndex]
		exp = expInt
	}

	pIndex := -1
	vLen := len(value)
	for i := 0; i < vLen; i++ {
		if value[i] == '.' {
			if pIndex > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", value)
			}
			pIndex = i
		}
	}

	if pIndex == -1 {
		// 没有小数点，可以直接把原字符串解析为
		// 整数
		intString = value
	} else {
		if pIndex+1 < vLen {
			intString = value[:pIndex] + value[pIndex+1:]
		} else {
			intString = value[:pIndex]
		}
		expInt := -len(value[pIndex+1:])
		exp += int64(expInt)
	}

	var dValue *big.Int
	// strconv.ParseInt 比 new(big.Int).SetString 更快，因此对已知不会溢出的字符串
	// 直接走这个捷径
		parsed64, err := strconv.ParseInt(intString, 10, 64)
		if err != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", value)
		}
		dValue = big.NewInt(parsed64)
	} else {
		dValue = new(big.Int)
		_, ok := dValue.SetString(intString, 10)
		if !ok {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", value)
		}
	}

	if exp < math.MinInt32 || exp > math.MaxInt32 {
	// NOTE(vadim)：一个字符串的长度几乎不可能真的达到这种程度
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", originalInput)
	}

	return Decimal{
		value: dValue,
		exp:   int32(exp),
	}, nil
}

// NewFromFormattedString 根据带格式的字符串表示形式返回新的 Decimal。
// 第二个参数 replRegexp 是一个正则表达式，用于找出应从给定十进制字符串表示形式中移除的字符，
// 所有匹配到的字符都会被替换为空字符串。
//
// 示例：
//
//	r := regexp.MustCompile("[$,]")
//	d1, err := NewFromFormattedString("$5,125.99", r)
//
//	r2 := regexp.MustCompile("[_]")
//	d2, err := NewFromFormattedString("1_000_000", r2)
//
//	r3 := regexp.MustCompile("[USD\\s]")
//	d3, err := NewFromFormattedString("5000 USD", r3)
func NewFromFormattedString(value string, replRegexp *regexp.Regexp) (Decimal, error) {
	parsedValue := replRegexp.ReplaceAllString(value, "")
	d, err := NewFromString(parsedValue)
	if err != nil {
		return Decimal{}, err
	}
	return d, nil
}

// RequireFromString 根据字符串表示形式返回新的 Decimal；
// 若 NewFromString 会返回错误，则该函数直接 panic。
//
// 示例：
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(value string) Decimal {
	dec, err := NewFromString(value)
	if err != nil {
		panic(err)
	}
	return dec
}

// NewFromFloat 将 float64 转换为 Decimal。
//
// 转换结果包含的位数，是可以在浮点数中可靠往返（roundtrip）表示的有效数字位数，
// 通常为 15 位，某些情况下可能更多。
// 更多信息参见 https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/ 。
//
// 若需要略快的转换，可使用 NewFromFloatWithExponent，以绝对方式指定精度。
//
// 注意：对 NaN、+/-inf 调用本函数会 panic。
func NewFromFloat(value float64) Decimal {
	if value == 0 {
		return New(0, 0)
	}
	return newFromFloat(value, math.Float64bits(value), &float64info)
}

// NewFromFloat32 将 float32 转换为 Decimal。
//
// 转换结果包含的位数，是可以在浮点数中可靠往返（roundtrip）表示的有效数字位数，
// 通常依输入不同为 6~8 位。
// 更多信息参见 https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/ 。
//
// 若需要略快的转换，可使用 NewFromFloatWithExponent，以绝对方式指定精度。
//
// 注意：对 NaN、+/-inf 调用本函数会 panic。
func NewFromFloat32(value float32) Decimal {
	if value == 0 {
		return New(0, 0)
	}
	// 使用 XOR 是对 https://github.com/golang/go/issues/26285 的规避方案
	a := math.Float32bits(value) ^ 0x80808080
	return newFromFloat(float64(value), uint64(a)^0x80808080, &float32info)
}

func newFromFloat(val float64, bits uint64, flt *floatInfo) Decimal {
	if math.IsNaN(val) || math.IsInf(val, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", val))
	}
	exp := int(bits>>flt.mantbits) & (1<<flt.expbits - 1)
	mant := bits & (uint64(1)<<flt.mantbits - 1)

	switch exp {
	case 0:
		// 非规格化（denormalized）数
		exp++

	default:
		// 加入隐含的最高位
		mant |= uint64(1) << flt.mantbits
	}
	exp += flt.bias

	var d decimal
	d.Assign(mant)
	d.Shift(exp - int(flt.mantbits))
	d.neg = bits>>(flt.expbits+flt.mantbits) != 0

	roundShortest(&d, mant, exp, flt)
	// 若少于 19 位数字，可在 int64 中完成计算。
	if d.nd < 19 {
		tmp := int64(0)
		m := int64(1)
		for i := d.nd - 1; i >= 0; i-- {
			tmp += m * int64(d.d[i]-'0')
			m *= 10
		}
		if d.neg {
			tmp *= -1
		}
		return Decimal{value: big.NewInt(tmp), exp: int32(d.dp) - int32(d.nd)}
	}
	dValue := new(big.Int)
	dValue, ok := dValue.SetString(string(d.d[:d.nd]), 10)
	if ok {
		return Decimal{value: dValue, exp: int32(d.dp) - int32(d.nd)}
	}

	return NewFromFloatWithExponent(val, int32(d.dp)-int32(d.nd))
}

// NewFromFloatWithExponent 将 float64 转换为 Decimal，并支持任意
// 小数位数。
//
// 示例：
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(value float64, exp int32) Decimal {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", value))
	}

	bits := math.Float64bits(value)
	mant := bits & (1<<52 - 1)
	exp2 := int32((bits >> 52) & (1<<11 - 1))
	sign := bits >> 63

	if exp2 == 0 {
		// 特殊情形（指数位为 0）
		if mant == 0 {
			return Decimal{}
		}
		// 次正规数（subnormal）
		exp2++
	} else {
		// 常规数
		mant |= 1 << 52
	}

	exp2 -= 1023 + 52

	// 对以 2 为底的尾数进行归一化
	for mant&1 == 0 {
		mant = mant >> 1
		exp2++
	}

	// 精确表示 2^N 所需的最大十进制小数位数：当 N<0 时不超过 -N
	if exp < 0 && exp < exp2 {
		if exp2 < 0 {
			exp = exp2
		} else {
			exp = 0
		}
	}

	// 将 10^M * 2^N 表示成 5^M * 2^(M+N)
	exp2 -= exp

	temp := big.NewInt(1)
	dMant := big.NewInt(int64(mant))

	// 计算 5^M
	if exp > 0 {
		temp = temp.SetInt64(int64(exp))
		temp = temp.Exp(fiveInt, temp, nil)
	} else if exp < 0 {
		temp = temp.SetInt64(-int64(exp))
		temp = temp.Exp(fiveInt, temp, nil)
		dMant = dMant.Mul(dMant, temp)
		temp = temp.SetUint64(1)
	}

	// 计算 2^(M+N)
	if exp2 > 0 {
		dMant = dMant.Lsh(dMant, uint(exp2))
	} else if exp2 < 0 {
		temp = temp.Lsh(temp, uint(-exp2))
	}

	// 舍入并缩减比例
	if exp > 0 || exp2 < 0 {
		halfDown := new(big.Int).Rsh(temp, 1)
		dMant = dMant.Add(dMant, halfDown)
		dMant = dMant.Quo(dMant, temp)
	}

	if sign == 1 {
		dMant = dMant.Neg(dMant)
	}

	return Decimal{
		value: dMant,
		exp:   exp,
	}
}

// Copy 返回该 decimal 的副本：值与指数相同，但指向新的 value 指针。
func (d Decimal) Copy() Decimal {
	d.ensureInitialized()
	return Decimal{
		value: &(*d.value),
		exp:   d.exp,
	}
}

// rescale 返回重新缩放（改变指数）后的 decimal。
// 若给定的指数大于原 Decimal 的初始指数，返回结果的精度可能降低。
// 注意：此处是截断，而不是四舍五入
//
// 示例：
//
//	d := New(12345, -4)
//	d2 := d.rescale(-1)
//	d3 := d2.rescale(-4)
//	println(d1)
//	println(d2)
//	println(d3)
//
// 输出：
//
//	1.2345
//	1.2
//	1.2000
func (d Decimal) rescale(exp int32) Decimal {
	d.ensureInitialized()

	if d.exp == exp {
		return Decimal{
			new(big.Int).Set(d.value),
			d.exp,
		}
	}

	// NOTE(vadim)：做减法前必须先把指数转为 float64，以防止溢出
	diff := math.Abs(float64(exp) - float64(d.exp))
	value := new(big.Int).Set(d.value)

	expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(diff)), nil)
	if exp > d.exp {
		value = value.Quo(value, expScale)
	} else if exp < d.exp {
		value = value.Mul(value, expScale)
	}

	return Decimal{
		value: value,
		exp:   exp,
	}
}

// Abs 返回该 decimal 的绝对值。
func (d Decimal) Abs() Decimal {
	if !d.IsNegative() {
		return d
	}
	d.ensureInitialized()
	d2Value := new(big.Int).Abs(d.value)
	return Decimal{
		value: d2Value,
		exp:   d.exp,
	}
}

// Add 返回 d + d2。
func (d Decimal) Add(d2 Decimal) Decimal {
	rd, rd2 := RescalePair(d, d2)

	d3Value := new(big.Int).Add(rd.value, rd2.value)
	return Decimal{
		value: d3Value,
		exp:   rd.exp,
	}
}

// Sub 返回 d - d2。
func (d Decimal) Sub(d2 Decimal) Decimal {
	rd, rd2 := RescalePair(d, d2)

	d3Value := new(big.Int).Sub(rd.value, rd2.value)
	return Decimal{
		value: d3Value,
		exp:   rd.exp,
	}
}

// Neg 返回 -d。
func (d Decimal) Neg() Decimal {
	d.ensureInitialized()
	val := new(big.Int).Neg(d.value)
	return Decimal{
		value: val,
		exp:   d.exp,
	}
}

// Mul 返回 d * d2。
func (d Decimal) Mul(d2 Decimal) Decimal {
	d.ensureInitialized()
	d2.ensureInitialized()

	expInt64 := int64(d.exp) + int64(d2.exp)
	if expInt64 > math.MaxInt32 || expInt64 < math.MinInt32 {
		// NOTE(vadim)：与其返回错误结果不如直接 panic，因为
		// Decimal 通常用于金额计算
		panic(fmt.Sprintf("exponent %v overflows an int32!", expInt64))
	}

	d3Value := new(big.Int).Mul(d.value, d2.value)
	return Decimal{
		value: d3Value,
		exp:   int32(expInt64),
	}
}

// Shift 以 10 为基数对 decimal 进行移位。
// shift 为正时左移，为负时右移。
// 简单来说，就是把给定的 shift 数值加到该 decimal 的指数上。
func (d Decimal) Shift(shift int32) Decimal {
	d.ensureInitialized()
	return Decimal{
		value: new(big.Int).Set(d.value),
		exp:   d.exp + shift,
	}
}

// Div 返回 d / d2。若无法整除，结果将保留
// DivisionPrecision 位小数。
func (d Decimal) Div(d2 Decimal) Decimal {
	return d.DivRound(d2, int32(DivisionPrecision))
}

// QuoRem 执行带余数的除法
// d.QuoRem(d2, precision) 返回商 q 与余数 r，满足：
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// 注意：允许 precision<0 作为输入。
func (d Decimal) QuoRem(d2 Decimal, precision int32) (Decimal, Decimal) {
	d.ensureInitialized()
	d2.ensureInitialized()
	if d2.value.Sign() == 0 {
		panic("decimal division by 0")
	}
	scale := -precision
	e := int64(d.exp - d2.exp - scale)
	if e > math.MaxInt32 || e < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var aa, bb, expo big.Int
	var scalerest int32
	// d = a * 10^ea
	// d2 = b * 10^eb
	if e < 0 {
		aa = *d.value
		expo.SetInt64(-e)
		bb.Exp(tenInt, &expo, nil)
		bb.Mul(d2.value, &bb)
		scalerest = d.exp
		// 此时 aa = a
		//      bb = b * 10^(scale + eb - ea)
	} else {
		expo.SetInt64(e)
		aa.Exp(tenInt, &expo, nil)
		aa.Mul(d.value, &aa)
		bb = *d2.value
		scalerest = scale + d2.exp
		// 此时 aa = a * 10^(ea - eb - scale)
		//      bb = b
	}
	var q, r big.Int
	q.QuoRem(&aa, &bb, &r)
	dq := Decimal{value: &q, exp: scale}
	dr := Decimal{value: &r, exp: scalerest}
	return dq, dr
}

// DivRound 执行除法并按给定精度四舍五入，
// 即结果将是 10^(-precision) 的整数倍
//
//	商为正时，数字 5 向上进位（远离 0）
//	商为负时，数字 5 向下进位（远离 0）
//
// 注意：允许 precision<0 作为输入。
func (d Decimal) DivRound(d2 Decimal, precision int32) Decimal {
	// QuoRem 内部已做初始化检查
	q, r := d.QuoRem(d2, precision)
	// 实际的舍入决策是比较 r*10^precision 与 d2/2 的大小；
	// 因此这里改为比较 2*r*10^precision 与 d2
	var rv2 big.Int
	rv2.Abs(r.value)
	rv2.Lsh(&rv2, 1)
	// 此时 rv2 = abs(r.value) * 2
	r2 := Decimal{value: &rv2, exp: r.exp + precision}
	// r2 现在是 2 * r * 10^precision
	c := r2.Cmp(d2.Abs())

	if c < 0 {
		return q
	}

	if d.value.Sign()*d2.value.Sign() < 0 {
		return q.Sub(New(1, -precision))
	}

	return q.Add(New(1, -precision))
}

// Mod 返回 d % d2。
func (d Decimal) Mod(d2 Decimal) Decimal {
	quo := d.DivRound(d2, -d.exp+1).Truncate(0)
	return d.Sub(d2.Mul(quo))
}

// Pow 返回 d 的 d2 次幂
func (d Decimal) Pow(d2 Decimal) Decimal {
	var temp Decimal
	if d2.IntPart() == 0 {
		return NewFromFloat(1)
	}
	temp = d.Pow(d2.Div(NewFromFloat(2)))
	if d2.IntPart()%2 == 0 {
		return temp.Mul(temp)
	}
	if d2.IntPart() > 0 {
		return temp.Mul(temp).Mul(d)
	}
	return temp.Mul(temp).Div(d)
}

// ExpHullAbrham 使用 Hull-Abraham 算法计算 decimal 的自然指数（即 e 的 d 次方）。
// OverallPrecision 参数指定结果的总体精度（整数部分 + 小数部分）。
//
// 精度较小时 ExpHullAbrham 比 ExpTaylor 更快，但精度很大时它会慢得多。
//
// 示例：
//
//	NewFromFloat(26.1).ExpHullAbrham(2).String()    // output: "220000000000"
//	NewFromFloat(26.1).ExpHullAbrham(20).String()   // output: "216314672147.05767284"
func (d Decimal) ExpHullAbrham(overallPrecision uint32) (Decimal, error) {
	// 算法基于《ACM Transactions on Mathematical Software》中
	// T. E. Hull 与 A. Abrham 发表的“变精度指数函数”一文。
	if d.IsZero() {
		return Decimal{oneInt, 0}, nil
	}

	currentPrecision := overallPrecision

	// 当 currentPrecision * 23 < |x| 时本算法不适用。
	// 此时会自动提高精度，从而可精确计算出结果。
	// 若新计算出的精度超过 ExpMaxIterations，则 currentPrecision 保持不变。
	f := d.Abs().InexactFloat64()
	if ncp := f / 23; ncp > float64(currentPrecision) && ncp < float64(ExpMaxIterations) {
		currentPrecision = uint32(math.Ceil(ncp))
	}

	// 若 abs(d) 超出上溢/下溢阈值则直接失败
	overflowThreshold := New(23*int64(currentPrecision), 0)
	if d.Abs().Cmp(overflowThreshold) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// 若 abs(d) 足够小则返回 1；这也可避免后续的上溢/下溢
	overflowThreshold2 := New(9, -int32(currentPrecision)-1)
	if d.Abs().Cmp(overflowThreshold2) <= 0 {
		return Decimal{oneInt, d.exp}, nil
	}

	// t 是使 abs(d/k) < 1 成立的最小非负整数
	t := d.exp + int32(d.NumDigits()) // 加上 d.NumDigits 是因为论文假设 d.value 在 [0.1, 1) 区间

	if t < 0 {
		t = 0
	}

	k := New(1, t)                                     // 缩减因子
	r := Decimal{new(big.Int).Set(d.value), d.exp - t} // 缩减后的自变量
	p := int32(currentPrecision) + t + 2               // 计算求和时的精度

	// 确定 n，即求和的项数；
	// 使用牛顿迭代第一步 (1.435p - 1.182) / log10(p/abs(r))
	// 来求解相应方程，并结合定向舍入以及
	// log10(p/abs(r)) 的简单有理数上界
	rf := r.Abs().InexactFloat64()
	pf := float64(p)
	nf := math.Ceil((1.453*pf - 1.182) / math.Log10(pf/rf))
	if nf > float64(ExpMaxIterations) || math.IsNaN(nf) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	n := int64(nf)

	tmp := New(0, 0)
	sum := New(1, 0)
	one := New(1, 0)
	for i := n - 1; i > 0; i-- {
		tmp.value.SetInt64(i)
		sum = sum.Mul(r.DivRound(tmp, p))
		sum = sum.Add(one)
	}

	ki := k.IntPart()
	res := New(1, 0)
	for i := ki; i > 0; i-- {
		res = res.Mul(sum)
	}

	resNumDigits := int32(res.NumDigits())

	var roundDigits int32
	if resNumDigits > abs(res.exp) {
		roundDigits = int32(currentPrecision) - resNumDigits - res.exp
	} else {
		roundDigits = int32(currentPrecision)
	}

	res = res.Round(roundDigits)

	return res, nil
}

// ExpTaylor 使用泰勒级数展开计算 decimal 的自然指数（即 e 的 d 次方）。
// Precision 参数指定结果的精度要求（小数点后的位数）。
// 允许传入负精度。
//
// 在精度较大时 ExpTaylor 比 ExpHullAbrham 快得多。
//
// 示例：
//
//	d, err := NewFromFloat(26.1).ExpTaylor(2).String()
//	d.String()  // output: "216314672147.06"
//
//	NewFromFloat(26.1).ExpTaylor(20).String()
//	d.String()  // output: "216314672147.05767284062928674083"
//
//	NewFromFloat(26.1).ExpTaylor(-10).String()
//	d.String()  // output: "220000000000"
func (d Decimal) ExpTaylor(precision int32) (Decimal, error) {
	// 注(mwoss)：该实现可进一步优化，仅使用 big.Int 的 API 即可
	if d.IsZero() {
		return Decimal{oneInt, 0}.Round(precision), nil
	}

	var epsilon Decimal
	var divPrecision int32
	if precision < 0 {
		epsilon = New(1, -1)
		divPrecision = 8
	} else {
		epsilon = New(1, -precision-1)
		divPrecision = precision + 1
	}

	decAbs := d.Abs()
	pow := d.Abs()
	factorial := New(1, 0)

	result := New(1, 0)

	for i := int64(1); ; {
		step := pow.DivRound(factorial, divPrecision)
		result = result.Add(step)

		// 当当前项小于 epsilon 时停止泰勒级数展开
		if step.Cmp(epsilon) < 0 {
			break
		}

		pow = pow.Mul(decAbs)

		i++

		// 计算下一个阶乘数，或从缓存中取出已缓存的值
		if len(factorials) >= int(i) && !factorials[i-1].IsZero() {
			factorial = factorials[i-1]
		} else {
			// 为避免并发竞态，先向切片追加零值占位，
			// 为新算出的阶乘留出位置；随后再通过下标
			// 用计算好的阶乘替换掉该零值。
			factorial = factorials[i-2].Mul(New(i, 0))
			factorials = append(factorials, Zero)
			factorials[i-1] = factorial
		}
	}

	if d.Sign() < 0 {
		result = New(1, 0).DivRound(result, precision+1)
	}

	result = result.Round(precision)
	return result, nil
}

// NumDigits 返回 decimal 系数（d.value）的位数。
// 注意：对于很大的 decimal，或小数部分很长的 decimal，当前实现会非常慢
func (d Decimal) NumDigits() int {
	// 注(mwoss)：此处可优化，把 big.Int 转成字符串并非必要
	if d.IsNegative() {
		return len(d.value.String()) - 1
	}
	return len(d.value.String())
}

// IsInteger 当 decimal 可以表示为整数值时返回 true，否则返回 false。
func (d Decimal) IsInteger() bool {
	// 最常见的情形：指数大于等于 0 的 decimal 都可以表示为整数
	if d.exp >= 0 {
		return true
	}
	// 当指数为负时，需要检查小数点后的每一位；
	// 若这些位全部为 0，即可确定该 decimal 可以表示为整数
	var r big.Int
	q := new(big.Int).Set(d.value)
	for z := abs(d.exp); z > 0; z-- {
		q.QuoRem(q, tenInt, &r)
		if r.Cmp(zeroInt) != 0 {
			return false
		}
	}
	return true
}

// abs 计算任意 int32 的绝对值，用于计算 decimal 指数的绝对值。
func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

// Cmp 比较 d 与 d2 所代表的数值并返回：
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (d Decimal) Cmp(d2 Decimal) int {
	d.ensureInitialized()
	d2.ensureInitialized()

	if d.exp == d2.exp {
		return d.value.Cmp(d2.value)
	}

	rd, rd2 := RescalePair(d, d2)

	return rd.value.Cmp(rd2.value)
}

// Equal 返回 d 与 d2 所代表的数值是否相等。
func (d Decimal) Equal(d2 Decimal) bool {
	return d.Cmp(d2) == 0
}

// Equals 已废弃，请改用 Equal 方法。
func (d Decimal) Equals(d2 Decimal) bool {
	return d.Equal(d2)
}

// GreaterThan（GT）当 d 大于 d2 时返回 true。
func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.Cmp(d2) == 1
}

// GreaterThanOrEqual（GTE）当 d 大于或等于 d2 时返回 true。
func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == 1 || cmp == 0
}

// LessThan（LT）当 d 小于 d2 时返回 true。
func (d Decimal) LessThan(d2 Decimal) bool {
	return d.Cmp(d2) == -1
}

// LessThanOrEqual（LTE）当 d 小于或等于 d2 时返回 true。
func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == -1 || cmp == 0
}

// Sign 返回：
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (d Decimal) Sign() int {
	if d.value == nil {
		return 0
	}
	return d.value.Sign()
}

// IsPositive 返回：
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (d Decimal) IsPositive() bool {
	return d.Sign() == 1
}

// IsNegative 返回：
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (d Decimal) IsNegative() bool {
	return d.Sign() == -1
}

// IsZero 返回：
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (d Decimal) IsZero() bool {
	return d.Sign() == 0
}

// Exponent 返回该 decimal 的指数（scale 分量）。
func (d Decimal) Exponent() int32 {
	return d.exp
}

// Coefficient 返回该 decimal 的系数，其缩放比例为 10^Exponent()。
func (d Decimal) Coefficient() *big.Int {
	d.ensureInitialized()
	// 对系数做一次拷贝，这样修改返回值不会影响原 Decimal。
	return new(big.Int).Set(d.value)
}

// CoefficientInt64 以 int64 返回该 decimal 的系数，其缩放比例为 10^Exponent()。
// 若系数无法在 int64 中表示，结果将是未定义的。
func (d Decimal) CoefficientInt64() int64 {
	d.ensureInitialized()
	return d.value.Int64()
}

// IntPart 返回该 decimal 的整数部分。
func (d Decimal) IntPart() int64 {
	scaledD := d.rescale(0)
	return scaledD.value.Int64()
}

// BigInt 以 big.Int 形式返回该 decimal 的整数部分。
func (d Decimal) BigInt() *big.Int {
	scaledD := d.rescale(0)
	i := &big.Int{}
	i.SetString(scaledD.String(), 10)
	return i
}

// BigFloat 以 big.Float 形式返回该 decimal。
// 注意：将 decimal 转换为 big.Float 可能会丢失精度。
func (d Decimal) BigFloat() *big.Float {
	f := &big.Float{}
	f.SetString(d.String())
	return f
}

// Rat 返回该 decimal 的有理数表示形式。
func (d Decimal) Rat() *big.Rat {
	d.ensureInitialized()
	if d.exp <= 0 {
		// NOTE(vadim)：必须先完成类型转换再取负，以防止 int32 溢出
		denom := new(big.Int).Exp(tenInt, big.NewInt(-int64(d.exp)), nil)
		return new(big.Rat).SetFrac(d.value, denom)
	}

	mul := new(big.Int).Exp(tenInt, big.NewInt(int64(d.exp)), nil)
	num := new(big.Int).Mul(d.value, mul)
	return new(big.Rat).SetFrac(num, oneInt)
}

// Float64 返回最接近 d 的 float64 值 f，并用布尔值指示
// f 是否精确等于 d。
// 更多细节参见 big.Rat.Float64 的文档。
func (d Decimal) Float64() (f float64, exact bool) {
	return d.Rat().Float64()
}

// InexactFloat64 返回最接近 d 的 float64 值，
// 但不指示该返回值是否精确等于 d。
func (d Decimal) InexactFloat64() float64 {
	f, _ := d.Float64()
	return f
}

// String 返回该 decimal 的定点（含小数点）字符串表示形式。
//
// 示例：
//
//	d := New(-12345, -3)
//	println(d.String())
//
// 输出：
//
//	-12.345
func (d Decimal) String() string {
	return d.string(true)
}

// StringFixed 返回经过四舍五入的定点字符串，小数点后保留 places 位数字。
//
// 示例：
//
//	NewFromFloat(0).StringFixed(2) // output: "0.00"
//	NewFromFloat(0).StringFixed(0) // output: "0"
//	NewFromFloat(5.45).StringFixed(0) // output: "5"
//	NewFromFloat(5.45).StringFixed(1) // output: "5.5"
//	NewFromFloat(5.45).StringFixed(2) // output: "5.45"
//	NewFromFloat(5.45).StringFixed(3) // output: "5.450"
//	NewFromFloat(545).StringFixed(-1) // output: "550"
func (d Decimal) StringFixed(places int32) string {
	rounded := d.Round(places)
	return rounded.string(false)
}

// StringFixedBank 返回按银行家舍入法（四舍六入五取偶）处理的定点字符串，
// 小数点后保留 places 位数字。
//
// 示例：
//
//	NewFromFloat(0).StringFixedBank(2) // output: "0.00"
//	NewFromFloat(0).StringFixedBank(0) // output: "0"
//	NewFromFloat(5.45).StringFixedBank(0) // output: "5"
//	NewFromFloat(5.45).StringFixedBank(1) // output: "5.4"
//	NewFromFloat(5.45).StringFixedBank(2) // output: "5.45"
//	NewFromFloat(5.45).StringFixedBank(3) // output: "5.450"
//	NewFromFloat(545).StringFixedBank(-1) // output: "540"
func (d Decimal) StringFixedBank(places int32) string {
	rounded := d.RoundBank(places)
	return rounded.string(false)
}

// StringFixedCash 返回按瑞典/现金舍入法处理的定点字符串。
// 更多细节参见 RoundCash 函数的说明。
func (d Decimal) StringFixedCash(interval uint8) string {
	rounded := d.RoundCash(interval)
	return rounded.string(false)
}

// Round 将 decimal 四舍五入到 places 位小数。
// 若 places < 0，则将整数部分舍入到最近的 10^(-places)。
//
// 示例：
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (d Decimal) Round(places int32) Decimal {
	if d.exp == -places {
		return d
	}
	// 先截断到 places + 1 位
	ret := d.rescale(-places - 1)

	// 加上 sign(d) * 0.5
	if ret.value.Sign() < 0 {
		ret.value.Sub(ret.value, fiveInt)
	} else {
		ret.value.Add(ret.value, fiveInt)
	}

	// 正数向下取整，负数向上取整
	_, m := ret.value.DivMod(ret.value, tenInt, new(big.Int))
	ret.exp++
	if ret.value.Sign() < 0 && m.Cmp(zeroInt) != 0 {
		ret.value.Add(ret.value, oneInt)
	}

	return ret
}

// RoundCeil 将 decimal 向正无穷方向舍入。
//
// 示例：
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.5"
func (d Decimal) RoundCeil(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}

	if d.value.Sign() > 0 {
		rescaled.value.Add(rescaled.value, oneInt)
	}

	return rescaled
}

// RoundFloor 将 decimal 向负无穷方向舍入。
//
// 示例：
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.4"
func (d Decimal) RoundFloor(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}

	if d.value.Sign() < 0 {
		rescaled.value.Sub(rescaled.value, oneInt)
	}

	return rescaled
}

// RoundUp 将 decimal 向远离零的方向舍入。
//
// 示例：
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.4"
func (d Decimal) RoundUp(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}

	if d.value.Sign() > 0 {
		rescaled.value.Add(rescaled.value, oneInt)
	} else if d.value.Sign() < 0 {
		rescaled.value.Sub(rescaled.value, oneInt)
	}

	return rescaled
}

// RoundDown 将 decimal 向靠近零的方向舍入。
//
// 示例：
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.5"
func (d Decimal) RoundDown(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}
	return rescaled
}

// RoundBank 将 decimal 四舍五入到 places 位小数。
// 若待舍入的最后一位恰好处在两个相邻整数中间（即距离相等），
// 则取其中的偶数作为舍入结果
//
// 若 places < 0，则将整数部分舍入到最近的 10^(-places)。
//
// Examples:
//
//	NewFromFloat(5.45).RoundBank(1).String() // output: "5.4"
//	NewFromFloat(545).RoundBank(-1).String() // output: "540"
//	NewFromFloat(5.46).RoundBank(1).String() // output: "5.5"
//	NewFromFloat(546).RoundBank(-1).String() // output: "550"
//	NewFromFloat(5.55).RoundBank(1).String() // output: "5.6"
//	NewFromFloat(555).RoundBank(-1).String() // output: "560"
func (d Decimal) RoundBank(places int32) Decimal {
	round := d.Round(places)
	remainder := d.Sub(round).Abs()

	half := New(5, -places-1)
	if remainder.Cmp(half) == 0 && round.value.Bit(0) != 0 {
		if round.value.Sign() < 0 {
			round.value.Add(round.value, oneInt)
		} else {
			round.value.Sub(round.value, oneInt)
		}
	}

	return round
}

// RoundCash 又称 Cash/Penny/öre 舍入，将 decimal 舍入到指定的
// 间隔。现金交易中应支付的金额会舍入到可用最小货币单位的最近
// 倍数。支持的间隔为：5、10、25、50、100；传入其他数值会触发 panic。
//
//	  5：  按 5 分舍入，如 3.43 => 3.45
//	 10：  按 10 分舍入，如 3.45 => 3.50（5 向上进位）
//	 25：  按 25 分舍入，如 3.41 => 3.50
//	 50：  按 50 分舍入，如 3.75 => 4.00
//	100： 按 100 分舍入，如 3.50 => 4.00
//
// 更多细节参见：https://en.wikipedia.org/wiki/Cash_rounding
func (d Decimal) RoundCash(interval uint8) Decimal {
	var iVal *big.Int
	switch interval {
	case 5:
		iVal = twentyInt
	case 10:
		iVal = tenInt
	case 25:
		iVal = fourInt
	case 50:
		iVal = twoInt
	case 100:
		iVal = oneInt
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", interval))
	}
	dVal := Decimal{
		value: iVal,
	}

	// TODO：优化这些计算以减少大量内存分配（约 29 次分配）。
	return d.Mul(dVal).Round(0).Div(dVal).Truncate(2)
}

// Floor 返回小于或等于 d 的最近整数值。
func (d Decimal) Floor() Decimal {
	d.ensureInitialized()

	if d.exp >= 0 {
		return d
	}

	exp := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	exp.Exp(exp, big.NewInt(-int64(d.exp)), nil)

	z := new(big.Int).Div(d.value, exp)
	return Decimal{value: z, exp: 0}
}

// Ceil 返回大于或等于 d 的最近整数值。
func (d Decimal) Ceil() Decimal {
	d.ensureInitialized()

	if d.exp >= 0 {
		return d
	}

	exp := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	exp.Exp(exp, big.NewInt(-int64(d.exp)), nil)

	z, m := new(big.Int).DivMod(d.value, exp, new(big.Int))
	if m.Cmp(zeroInt) != 0 {
		z.Add(z, oneInt)
	}
	return Decimal{value: z, exp: 0}
}

// Truncate 直接截掉数字的若干位，不做四舍五入。
//
// 注意：precision 是不被截掉的最后一位（必须 >= 0）。
//
// 示例：
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (d Decimal) Truncate(precision int32) Decimal {
	d.ensureInitialized()
	if precision >= 0 && -precision > d.exp {
		return d.rescale(-precision)
	}
	return d
}

// UnmarshalJSON 实现 json.Unmarshaler 接口。
func (d *Decimal) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" {
		return nil
	}

	str, err := unquoteIfQuoted(decimalBytes)
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", decimalBytes, err)
	}

	decimal, err := NewFromString(str)
	*d = decimal
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", str, err)
	}
	return nil
}

// MarshalJSON 实现 json.Marshaler 接口。
func (d Decimal) MarshalJSON() ([]byte, error) {
	var str string
	if MarshalJSONWithoutQuotes {
		str = d.String()
	} else {
		str = "\"" + d.String() + "\""
	}
	return []byte(str), nil
}

// UnmarshalBinary 实现 encoding.BinaryUnmarshaler 接口。由于编码为文本时
// 已使用字符串表示，该方法以 []byte 形式存储该字符串。
func (d *Decimal) UnmarshalBinary(data []byte) error {
	// 校验指数至少占用 4 个字节。GOB 编码后的值可能为空。
	if len(data) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", data, len(data))
	}

	// 取出指数
	d.exp = int32(binary.BigEndian.Uint32(data[:4]))

	// 取出数值
	d.value = new(big.Int)
	if err := d.value.GobDecode(data[4:]); err != nil {
		return fmt.Errorf("error decoding binary %v: %s", data, err)
	}

	return nil
}

// MarshalBinary 实现 encoding.BinaryMarshaler 接口。
func (d Decimal) MarshalBinary() (data []byte, err error) {
	// 先写指数（其长度为定值）
	v1 := make([]byte, 4)
	binary.BigEndian.PutUint32(v1, uint32(d.exp))

	// 追加数值
	var v2 []byte
	if v2, err = d.value.GobEncode(); err != nil {
		return
	}

	// 返回字节数组
	data = append(v1, v2...)
	return
}

// Scan 实现 sql.Scanner 接口，用于数据库反序列化。
func (d *Decimal) Scan(value interface{}) error {
	// 首先尝试判断数据库中的数据是否为 Numeric 类型
	switch v := value.(type) {

	case float32:
		*d = NewFromFloat(float64(v))
		return nil

	case float64:
		// sqlite3 中的 numeric 类型会以 float64 传递
		*d = NewFromFloat(v)
		return nil

	case int64:
		// 至少在 sqlite3 中，当数据库中的值为 0 时，数据会以
		// int64 而不是 float64 传递给我们……
		*d = New(v, 0)
		return nil

	default:
		// 默认情况下尝试把存储的值解释为字符串
		str, err := unquoteIfQuoted(v)
		if err != nil {
			return err
		}
		*d, err = NewFromString(str)
		return err
	}
}

// Value 实现 driver.Valuer 接口，用于数据库序列化。
func (d Decimal) Value() (driver.Value, error) {
	return d.String(), nil
}

// UnmarshalText 实现 encoding.TextUnmarshaler 接口，用于 XML 反序列化。
func (d *Decimal) UnmarshalText(text []byte) error {
	str := string(text)

	dec, err := NewFromString(str)
	*d = dec
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", str, err)
	}

	return nil
}

// MarshalText 实现 encoding.TextMarshaler 接口，用于 XML 序列化。
func (d Decimal) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

// GobEncode 实现 gob.GobEncoder 接口，用于 gob 序列化。
func (d Decimal) GobEncode() ([]byte, error) {
	return d.MarshalBinary()
}

// GobDecode 实现 gob.GobDecoder 接口，用于 gob 反序列化。
func (d *Decimal) GobDecode(data []byte) error {
	return d.UnmarshalBinary(data)
}

// StringScaled 先将 decimal 缩放，再对其调用 .String()。
// 注意：该函数有缺陷、不符合直觉且已被废弃！请改用 StringFixed。
func (d Decimal) StringScaled(exp int32) string {
	return d.rescale(exp).String()
}

func (d Decimal) string(trimTrailingZeros bool) string {
	if d.exp >= 0 {
		return d.rescale(0).value.String()
	}

	abs := new(big.Int).Abs(d.value)
	str := abs.String()

	var intPart, fractionalPart string

	// NOTE(vadim)：若 d.exp == INT_MIN 且运行在 32 位机器上，
	// 这里到 int 的转换会引发问题。此极端情况不再修复。
	dExpInt := int(d.exp)
	if len(str) > -dExpInt {
		intPart = str[:len(str)+dExpInt]
		fractionalPart = str[len(str)+dExpInt:]
	} else {
		intPart = "0"

		num0s := -dExpInt - len(str)
		fractionalPart = strings.Repeat("0", num0s) + str
	}

	if trimTrailingZeros {
		i := len(fractionalPart) - 1
		for ; i >= 0; i-- {
			if fractionalPart[i] != '0' {
				break
			}
		}
		fractionalPart = fractionalPart[:i+1]
	}

	number := intPart
	if len(fractionalPart) > 0 {
		number += "." + fractionalPart
	}

	if d.value.Sign() < 0 {
		return "-" + number
	}

	return number
}

func (d *Decimal) ensureInitialized() {
	if d.value == nil {
		d.value = new(big.Int)
	}
}

// MinDecimal 返回传入参数中最小的 Decimal。
//
// 若要以数组调用该函数，必须这样写：
//
//	MinDecimal(arr[0], arr[1:]...)
//
// 这样可避免不小心用 0 个参数调用 Min。
func MinDecimal(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) < 0 {
			ans = item
		}
	}
	return ans
}

// MaxDecimal 返回传入参数中最大的 Decimal。
//
// 若要以数组调用该函数，必须这样写：
//
//	MaxDecimal(arr[0], arr[1:]...)
//
// 这样可避免不小心用 0 个参数调用 Max。
func MaxDecimal(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) > 0 {
			ans = item
		}
	}
	return ans
}

// SumDecimal 返回 first 与其余 Decimals 的总和。
func SumDecimal(first Decimal, rest ...Decimal) Decimal {
	total := first
	for _, item := range rest {
		total = total.Add(item)
	}

	return total
}

// AvgDecimal 返回 first 与其余 Decimals 的平均值。
func AvgDecimal(first Decimal, rest ...Decimal) Decimal {
	count := New(int64(len(rest)+1), 0)
	sum := SumDecimal(first, rest...)
	return sum.Div(count)
}

// RescalePair 将两个 decimal 缩放到共同的指数（取两者中较小的指数）。
func RescalePair(d1 Decimal, d2 Decimal) (Decimal, Decimal) {
	d1.ensureInitialized()
	d2.ensureInitialized()

	if d1.exp == d2.exp {
		return d1, d2
	}

	baseScale := min(d1.exp, d2.exp)
	if baseScale != d1.exp {
		return d1.rescale(baseScale), d2
	}
	return d1, d2.rescale(baseScale)
}

func min(x, y int32) int32 {
	if x >= y {
		return y
	}
	return x
}

func unquoteIfQuoted(value interface{}) (string, error) {
	var bytes []byte

	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'",
			value, value)
	}

	// If the amount is quoted, strip the quotes
	if len(bytes) > 2 && bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}
	return string(bytes), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the database.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(d Decimal) NullDecimal {
	return NullDecimal{
		Decimal: d,
		Valid:   true,
	}
}

// Scan 实现 sql.Scanner 接口，用于数据库反序列化。
func (d *NullDecimal) Scan(value interface{}) error {
	if value == nil {
		d.Valid = false
		return nil
	}
	d.Valid = true
	return d.Decimal.Scan(value)
}

// Value 实现 driver.Valuer 接口，用于数据库序列化。
func (d NullDecimal) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}
	return d.Decimal.Value()
}

// UnmarshalJSON 实现 json.Unmarshaler 接口。
func (d *NullDecimal) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" {
		d.Valid = false
		return nil
	}
	d.Valid = true
	return d.Decimal.UnmarshalJSON(decimalBytes)
}

// MarshalJSON 实现 json.Marshaler 接口。
func (d NullDecimal) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}
	return d.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (d *NullDecimal) UnmarshalText(text []byte) error {
	str := string(text)

	// check for empty XML or XML without body e.g., <tag></tag>
	if str == "" {
		d.Valid = false
		return nil
	}
	if err := d.Decimal.UnmarshalText(text); err != nil {
		d.Valid = false
		return err
	}
	d.Valid = true
	return nil
}

// MarshalText 实现 encoding.TextMarshaler 接口，用于 XML 序列化。
func (d NullDecimal) MarshalText() (text []byte, err error) {
	if !d.Valid {
		return []byte{}, nil
	}
	return d.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (d Decimal) Atan() Decimal {
	if d.Equal(NewFromFloat(0.0)) {
		return d
	}
	if d.GreaterThan(NewFromFloat(0.0)) {
		return d.satan()
	}
	return d.Neg().satan().Neg()
}

func (d Decimal) xatan() Decimal {
	P0 := NewFromFloat(-8.750608600031904122785e-01)
	P1 := NewFromFloat(-1.615753718733365076637e+01)
	P2 := NewFromFloat(-7.500855792314704667340e+01)
	P3 := NewFromFloat(-1.228866684490136173410e+02)
	P4 := NewFromFloat(-6.485021904942025371773e+01)
	Q0 := NewFromFloat(2.485846490142306297962e+01)
	Q1 := NewFromFloat(1.650270098316988542046e+02)
	Q2 := NewFromFloat(4.328810604912902668951e+02)
	Q3 := NewFromFloat(4.853903996359136964868e+02)
	Q4 := NewFromFloat(1.945506571482613964425e+02)
	z := d.Mul(d)
	b1 := P0.Mul(z).Add(P1).Mul(z).Add(P2).Mul(z).Add(P3).Mul(z).Add(P4).Mul(z)
	b2 := z.Add(Q0).Mul(z).Add(Q1).Mul(z).Add(Q2).Mul(z).Add(Q3).Mul(z).Add(Q4)
	z = b1.Div(b2)
	z = d.Mul(z).Add(d)
	return z
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (d Decimal) satan() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	pi := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if d.LessThanOrEqual(NewFromFloat(0.66)) {
		return d.xatan()
	}
	if d.GreaterThan(Tan3pio8) {
		return pi.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(d).xatan()).Add(Morebits)
	}
	return pi.Div(NewFromFloat(4.0)).Add((d.Sub(NewFromFloat(1.0)).Div(d.Add(NewFromFloat(1.0)))).xatan()).Add(NewFromFloat(0.5).Mul(Morebits))
}

// sin coefficients
var _sin = [...]Decimal{
	NewFromFloat(1.58962301576546568060e-10), // 0x3de5d8fd1fd19ccd
	NewFromFloat(-2.50507477628578072866e-8), // 0xbe5ae5e5a9291f5d
	NewFromFloat(2.75573136213857245213e-6),  // 0x3ec71de3567d48a1
	NewFromFloat(-1.98412698295895385996e-4), // 0xbf2a01a019bfdf03
	NewFromFloat(8.33333333332211858878e-3),  // 0x3f8111111110f7d0
	NewFromFloat(-1.66666666666666307295e-1), // 0xbfc5555555555548
}

// Sin returns the sine of the radian argument x.
func (d Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if d.Equal(NewFromFloat(0.0)) {
		return d
	}
	// make argument positive but save the sign
	sign := false
	if d.LessThan(NewFromFloat(0.0)) {
		d = d.Neg()
		sign = true
	}

	j := d.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := NewFromFloat(float64(j)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if j&1 == 1 {
		j++
		y = y.Add(NewFromFloat(1.0))
	}
	j &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if j > 3 {
		sign = !sign
		j -= 4
	}
	z := d.Sub(y.Mul(PI4A)).Sub(y.Mul(PI4B)).Sub(y.Mul(PI4C)) // Extended precision modular arithmetic
	zz := z.Mul(z)

	if j == 1 || j == 2 {
		w := zz.Mul(zz).Mul(_cos[0].Mul(zz).Add(_cos[1]).Mul(zz).Add(_cos[2]).Mul(zz).Add(_cos[3]).Mul(zz).Add(_cos[4]).Mul(zz).Add(_cos[5]))
		y = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(zz)).Add(w)
	} else {
		y = z.Add(z.Mul(zz).Mul(_sin[0].Mul(zz).Add(_sin[1]).Mul(zz).Add(_sin[2]).Mul(zz).Add(_sin[3]).Mul(zz).Add(_sin[4]).Mul(zz).Add(_sin[5])))
	}
	if sign {
		y = y.Neg()
	}
	return y
}

// cos coefficients
var _cos = [...]Decimal{
	NewFromFloat(-1.13585365213876817300e-11), // 0xbda8fa49a0861a9b
	NewFromFloat(2.08757008419747316778e-9),   // 0x3e21ee9d7b4e3f05
	NewFromFloat(-2.75573141792967388112e-7),  // 0xbe927e4f7eac4bc6
	NewFromFloat(2.48015872888517045348e-5),   // 0x3efa01a019c844f5
	NewFromFloat(-1.38888888888730564116e-3),  // 0xbf56c16c16c14f91
	NewFromFloat(4.16666666666665929218e-2),   // 0x3fa555555555554b
}

// Cos returns the cosine of the radian argument x.
func (d Decimal) Cos() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	sign := false
	if d.LessThan(NewFromFloat(0.0)) {
		d = d.Neg()
	}

	j := d.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := NewFromFloat(float64(j)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if j&1 == 1 {
		j++
		y = y.Add(NewFromFloat(1.0))
	}
	j &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if j > 3 {
		sign = !sign
		j -= 4
	}
	if j > 1 {
		sign = !sign
	}

	z := d.Sub(y.Mul(PI4A)).Sub(y.Mul(PI4B)).Sub(y.Mul(PI4C)) // Extended precision modular arithmetic
	zz := z.Mul(z)

	if j == 1 || j == 2 {
		y = z.Add(z.Mul(zz).Mul(_sin[0].Mul(zz).Add(_sin[1]).Mul(zz).Add(_sin[2]).Mul(zz).Add(_sin[3]).Mul(zz).Add(_sin[4]).Mul(zz).Add(_sin[5])))
	} else {
		w := zz.Mul(zz).Mul(_cos[0].Mul(zz).Add(_cos[1]).Mul(zz).Add(_cos[2]).Mul(zz).Add(_cos[3]).Mul(zz).Add(_cos[4]).Mul(zz).Add(_cos[5]))
		y = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(zz)).Add(w)
	}
	if sign {
		y = y.Neg()
	}
	return y
}

var _tanP = [...]Decimal{
	NewFromFloat(-1.30936939181383777646e+4), // 0xc0c992d8d24f3f38
	NewFromFloat(1.15351664838587416140e+6),  // 0x413199eca5fc9ddd
	NewFromFloat(-1.79565251976484877988e+7), // 0xc1711fead3299176
}

var _tanQ = [...]Decimal{
	NewFromFloat(1.00000000000000000000e+0),
	NewFromFloat(1.36812963470692954678e+4),  // 0x40cab8a5eeb36572
	NewFromFloat(-1.32089234440210967447e+6), // 0xc13427bc582abc96
	NewFromFloat(2.50083801823357915839e+7),  // 0x4177d98fc2ead8ef
	NewFromFloat(-5.38695755929454629881e+7), // 0xc189afe03cbe5a31
}

// Tan returns the tangent of the radian argument x.
func (d Decimal) Tan() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if d.Equal(NewFromFloat(0.0)) {
		return d
	}

	// make argument positive but save the sign
	sign := false
	if d.LessThan(NewFromFloat(0.0)) {
		d = d.Neg()
		sign = true
	}

	j := d.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := NewFromFloat(float64(j)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if j&1 == 1 {
		j++
		y = y.Add(NewFromFloat(1.0))
	}

	z := d.Sub(y.Mul(PI4A)).Sub(y.Mul(PI4B)).Sub(y.Mul(PI4C)) // Extended precision modular arithmetic
	zz := z.Mul(z)

	if zz.GreaterThan(NewFromFloat(1e-14)) {
		w := zz.Mul(_tanP[0].Mul(zz).Add(_tanP[1]).Mul(zz).Add(_tanP[2]))
		x := zz.Add(_tanQ[1]).Mul(zz).Add(_tanQ[2]).Mul(zz).Add(_tanQ[3]).Mul(zz).Add(_tanQ[4])
		y = z.Add(z.Mul(w.Div(x)))
	} else {
		y = z
	}
	if j&2 == 2 {
		y = NewFromFloat(-1.0).Div(y)
	}
	if sign {
		y = y.Neg()
	}
	return y
}
