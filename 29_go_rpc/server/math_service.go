package server

// MathService 定义一个远程服务对象
type MathService struct {
}

// Args 定义一个参数结构体
type Args struct {
	A, B int
}

// Add 定义一个相加
func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
