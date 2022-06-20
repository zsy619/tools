package xerror

import "fmt"

// WrapError 包装错误
func WrapError(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s, err = %s", fmt.Sprintf(msg, args...), err)
}

// ReturnFirstError 返回第一个错误
func ReturnFirstError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// OneByOneUntilError 按序处理直到发生错误
func OneByOneUntilError(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
