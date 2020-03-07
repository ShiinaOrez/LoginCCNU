package spoc

import (
	"net/http"
)

func AddHeaders(request *http.Request) *http.Request {
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Host", "spoc.ccnu.edu.cn")
	request.Header.Set("Origin", "http://spoc.ccnu.edu.cn")
	request.Header.Set("Referer", "http://spoc.ccnu.edu.cn/studentHomepage")
	return request
}
