package _6_http

import "net/url"

// ReqWithKV 请求url?k=v
func ReqWithKV(urlAddress *url.URL, key, value string) {
	// 获取Query对象
	query := urlAddress.Query()
	// query 本质是一个map ——> type Values map[string][]string
	query.Add(key, value)
	// 将当前query对象的values转为 k1=v1&k2=v2, 称为encode
	urlAddress.RawQuery = query.Encode()
}

// ReqWithPath 请求url/path
func ReqWithPath(urlAddress *url.URL, path string) {
	urlAddress.Path = path
}
