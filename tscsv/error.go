package tscsv

import "fmt"

//ImportError 导入错误
type ImportError struct {
	Row      int
	Column   int
	ErrorMsg string
}

func (e ImportError) Error() string {

	return fmt.Sprintf("Row:%d,Column:%d,error:%s", e.Row, e.Column, e.ErrorMsg)
}

//ParseErrors 将错误数组解析为字符串
func ParseErrors(errs []ImportError) (errString string) {
	if len(errs) > 0 {
		for n, err := range errs {
			errString += err.Error()
			if n < len(errs)-1 {
				errString += "\n"
			}
		}
	}
	return
}
