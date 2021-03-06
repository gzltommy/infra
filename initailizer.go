package infra

// 初始化接口
type Initializer interface {
	// 用于对象实例化后的初始化操作
	Init()
}

// 初始化注册器
type InitializeRegister struct {
	Initializers []Initializer
}

// 注册一个初始化对象
func (ir *InitializeRegister) Register(i Initializer) {
	ir.Initializers = append(ir.Initializers, i)
}
