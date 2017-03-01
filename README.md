# golang-study
go语言学习

##实例
- [fibonacci](https://github.com/GxlZ/golang-study/tree/master/fibonacci) `命令行 计算斐波那契数列第N位数字`

##常用技巧

- 变量值互换
```
i, j := 1, 2
i, j = j, i
```

- 类型转换 
> `对于每个类型T,都有一个对应的类型转换操作T(x),如果T是指针类型,需要用括号包装T (*T)(x)`
```  
type F float64
fmt.Println(reflect.TypeOf(F(123)))
```