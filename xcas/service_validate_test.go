package xcas

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func Test_CasVersion2ServiceValidateActionSuccess(t *testing.T) {
	rt :=
		`<serviceResponse xmlns="http://www.yale.edu/tp/cas">
  <authenticationSuccess>
    <user>13633861512</user>
    <attributes>
      <authenticationDate>2022-02-18T08:56:04.984551Z</authenticationDate>
      <longTermAuthenticationRequestTokenUsed>false</longTermAuthenticationRequestTokenUsed>
      <isFromNewLogin>true</isFromNewLogin>
      <userAttributes>
        <attribute name="username">13633861512</attribute>
        <attribute name="xm">朱书彦</attribute>
        <attribute name="jg"></attribute>
        <attribute name="xjd">河南郑州</attribute>
        <attribute name="zzmm">中共党员</attribute>
        <attribute name="csrq">0001-01-01</attribute>
        <attribute name="xz"></attribute>
        <attribute name="dz"></attribute>
        <attribute name="yx">zsy619@163.com</attribute>
        <attribute name="tx">http://localhost:8888/Uploads/icon/7fb467db9b42346e0a1dc2d549ccd0d9_resize.png</attribute>
        <attribute name="byyx">北京大学</attribute>
        <attribute name="dh">13633861512</attribute>
        <attribute name="nj"></attribute>
        <attribute name="sfzh">412924197602102558</attribute>
        <attribute name="usertype">student</attribute>
        <attribute name="xb">女</attribute>
        <attribute name="bysj">2021-04</attribute>
        <attribute name="zy">计算机</attribute>
        <attribute name="bj"></attribute>
        <attribute name="xl">硕士</attribute>
      </userAttributes>
    </attributes>
  </authenticationSuccess>
</serviceResponse>`
	var serviceResponse *XmlServiceResponse
	err := xml.Unmarshal([]byte(rt), &serviceResponse)
	if err != nil {
		fmt.Println("1 --->", err.Error())
		return
	}
	if serviceResponse.Success != nil {
		fmt.Println("2 --->", serviceResponse.Success.User)
		for _, attribute := range serviceResponse.Success.Attributes.UserAttributes.Attributes {
			fmt.Println("3--->", attribute.Name, attribute.Value)
		}
	}
}

func Test_CasVersion2ServiceValidateActionFail(t *testing.T) {
	rt :=
		`<serviceResponse xmlns="http://www.yale.edu/tp/cas">
  <authenticationFailure code="INVALID_TICKET">Ticket ST-aOHxOmsUmidqWhTghxiQxMejUMUzKSnXlNyblMbdGCEORoTfiOlEkUPCXdyjdehnwqiDOljSgXzjCgEkzsFCvzjQKqXhTOgwcqlq not recognized</authenticationFailure>
</serviceResponse>`
	var serviceResponse *XmlServiceResponse
	err := xml.Unmarshal([]byte(rt), &serviceResponse)
	if err != nil {
		fmt.Println("1 --->", err.Error())
		return
	}
	fmt.Println("2 --->", serviceResponse.Failure)
	if serviceResponse.Failure != nil {
		fmt.Println("3 --->", serviceResponse.Failure.Code)
		fmt.Println("3 --->", serviceResponse.Failure.Message)
	}
}
