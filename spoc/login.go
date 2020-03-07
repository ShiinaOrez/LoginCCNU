package spoc

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/json"
	"strings"
)

func Login(sno, password string) (response *Response, err error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}
	response, err = LoginSPOC(sno, password, client)
	if err != nil {
		return nil, err
	}
	return
}

func LoginSPOC(sno, password string, client *http.Client) (*Response, error) {
	v := url.Values{}
	v.Set("loginName", sno)
	v.Set("password", password)

	request, err := http.NewRequest("GET", "http://spoc.ccnu.edu.cn/userLoginController/getVerifCode", nil)
	if err != nil {
		return nil, err
	}
	_, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	request, err = http.NewRequest("POST", "http://spoc.ccnu.edu.cn/userLoginController/getUserProfile", strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	request = AddHeaders(request)
	_, err = client.Do(request)
	if err != nil {
		return nil, err
	}

	request, err = http.NewRequest("POST", "http://spoc.ccnu.edu.cn/userInfo/getUserInfo", strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	request = AddHeaders(request)
	resp, respErr := client.Do(request)
	if respErr != nil {
		return nil, respErr
	}
	if _, ok := resp.Header["Content-Language"]; !ok {
		request, err = http.NewRequest("GET", "http://spoc.ccnu.edu.cn/userInfo/getUserInfo", nil)
		resp, err = client.Do(request)
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		response := Response{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	}
	return nil, nil
}
