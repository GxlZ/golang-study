# golang-study
go语言学习

##实例
- [fibonacci](https://github.com/GxlZ/golang-study/tree/master/fibonacci) `命令行 计算斐波那契数列第N位数字`
- [clock](https://github.com/GxlZ/golang-study/tree/master/clock) `网络时钟服务`


##常用技巧
- 时间日期格式化

| 01/Jan | 02 | 03/15  | 04 | 05 | 06 |-07\[00]\[:00] | PM | Mon |
| :--:   |:--:|:--:    |:--:|:--:|:--:|:--:           |:--:|:--: |
| 月      | 日 | 时     | 分 | 秒  | 年 |      时差     |上下午|星期几|

1234567 分别对应：月日时分秒年 时差
> 注意:  
  月：01或Jan都可以  
  小时：03表示12小时制，15表示24小时制。  
  时差：是 -07 ，不是 07,后边可以增加“00”或“:00”，表示更进一步的分秒时差。  
  上下午：使用PM，不是AM。  

- 变量值互换
```
i, j := 1, 2
i, j = j, i
```

- 类型转换
`对于每个类型T,都有一个对应的类型转换操作T(x),如果T是指针类型,需要用括号包装T (*T)(x)`
```  
type F float64
fmt.Println(reflect.TypeOf(F(123)))
```