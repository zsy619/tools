package xcas

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ticket认证
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

	body, err := ioutil.ReadAll(res.Body)
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

// 退出登录
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("CASLogoutAction-->", err)
		return err
	}
	fmt.Println("CASLogoutAction-->", string(body))
	return nil
}
