package xcas

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// CasVersion2ServiceValidateAction 使用 CAS 2.0 serviceValidate 接口对票据进行校验。
// 参数 url 为已拼接好的 serviceValidate 完整 URL（含 service 与 ticket）。
// 返回解析后的 XmlServiceResponse；当响应体不包含 serviceResponse 字段时返回 invalid service response 错误。
// 网络或解析失败时将错误打印到标准输出，并返回相应错误。
func CasVersion2ServiceValidateAction(url string) (serviceResponse *XmlServiceResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	bodyx := string(body)
	fmt.Println("----->", bodyx)
	if strings.Contains(bodyx, "serviceResponse") {
		err = xml.Unmarshal(body, &serviceResponse)
		return
	}
	return nil, fmt.Errorf("invalid service response")
}

// CASLogoutAction 通过 POST 请求调用 CAS 登出接口，使服务端会话失效。
// 参数 logoutPath 为 CAS 登出接口的完整 URL。
// 任意步骤失败都会打印日志并返回错误；成功时打印响应体并返回 nil。
func CASLogoutAction(logoutPath string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", logoutPath, nil)
	if err != nil {
		fmt.Println("CASLogoutAction-->", err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("CASLogoutAction-->", err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("CASLogoutAction-->", err)
		return err
	}
	fmt.Println("CASLogoutAction-->", string(body))
	return nil
}
