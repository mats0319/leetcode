 - go tool compile –S main.go >> main.S 生成main.go的go汇编，并保存在main.S文件中  
 - defer只会压一个函数进栈；如果有一串连续的调用，最后一个函数之前的内容都会立刻执行  
 - 显式调用os.Exit(0)，defer不会被触发  
 - var _ Interface = (*interfaceImpl)(nil)，检查结构体是否实现了接口  
 - “值类型结构体不能使用定义在指针类型接收者上的方法”——如果要实现接口，尽量使用指针类型的结构体  
 - map在未初始化之前无法写入（panic：nil）；读操作可以正常执行，但使用v, ok := map[key]时，ok值为false，v为对应类型的0值  
 