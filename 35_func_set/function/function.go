package function

import "fmt"

//Rpc服务端的处理流程设计的函数

type OP func(msg interface{}) (interface{}, error)

//请求解码
func decode(msg interface{}) OP {
	//编码
	return func(msg interface{}) (interface{}, error) {
		fmt.Println("decoding ...", msg)
		result := fmt.Sprintf("decode_%v", msg)
		fmt.Println("decoded success->", result)
		return result, nil
	}
}

//解码后的处理
func handler(msg interface{}) OP {
	//处理
	return func(msg interface{}) (interface{}, error) {
		fmt.Println("do handle ...", msg)
		result := fmt.Sprintf("handled_%v", msg)
		fmt.Println("handled success->", result)
		return result, nil
	}
}

//处理结果的编码
func encode(msg interface{}) OP {
	//编码
	return func(msg interface{}) (interface{}, error) {
		fmt.Println("encode ...", msg)
		result := fmt.Sprintf("encode_%v", msg)
		fmt.Println("encoded success->", result)
		return result, nil
	}
}
