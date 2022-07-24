package function

import "testing"

func TestOP(t *testing.T) {
	//定义rpc请求的入参
	msg := "[rpc_consumer_data]"
	//获取OP函数名
	decodeOPName := decode(msg)
	//执行OP
	decodeResult, _ := decodeOPName(msg)

	handleOPName := handler(decodeResult)
	handleResult, _ := handleOPName(decodeResult)

	encodeOPName := encode(handleResult)
	_, err := encodeOPName(handleResult)
	if err != nil {
		return
	}

}
