# eval

eval 是一个表达式求值引擎，可以解析和计算表达式。表达式是返回单个值（主要但不限于布尔值）的一行代码。

该组件的目的是允许用户使用配置表达式的方式来实现更复杂的逻辑。为了方便“业务规则引擎”的开发而编写。

```go
// 算数运算
1 + 2 * 3
5 / 9 * (x - 32)
// 逻辑运算
a < 10 == b > 6
lang in ["cpp", "golang"]
// 内置函数
contains("golang is a beautiful language", x)
sin(pi / 2)
```



#### Install

```shell
go get github.com/Cauchy-NY/eval
```



#### Examples

**Ex.01-hello, world!** 

打印出 hello,world!

```go
input := "greet + name"

expr, err := Parse(input)
if err != nil {
	panic(err)
}

var env = parser.Env{
	"greet": "hello,",
	"name":  "world!",
}

got := expr.Eval(test.env)
fmt.Printf("%s", got)

// output:
// hello,world!
```

**Ex.02-video review** 

当视频的色情模型分大于 0.86 且用户不是特殊用户时，下架该视频

```go
input := "pron_predict > 0.86 && user_type not_in [\"big_v\", \"org\"]"

expr, err := Parse(input)
if err != nil {
	panic(err)
}

var videos = []parser.Env{
	{"pron_predict": 0.91, "user_type": "normal"},
	{"pron_predict": 0.93, "user_type": "big_v"},
	{"pron_predict": 0.66, "user_type": "normal"},
}

fmt.Printf("%s\n", input)

for _, video := range videos {
  got := expr.Eval(video)
  fmt.Printf("\t%v => %v\n", video, got)
}

// output:
// pron_predict > 0.86 && user_type not_in ["big_v", "org"]
//	map[pron_predict:0.91 user_type:normal] => true
//	map[pron_predict:0.93 user_type:big_v] => false
//	map[pron_predict:0.66 user_type:normal] => false
```



#### 支持的运算符

**算数**：`+`  `-`  `*`  `/`  `%`

**比较**：`< `  `lt`  `<=`  `le`  `>`  `gt`  `>=`  `ge`  `==`  `eq`  `!=`  `ne`

**逻辑**：`&&`  `and`  `AND`  `||`  `or`  `OR`  `in`  `not_in`

**单目**：`!`  `not`  `+`  `-`

**嵌套**：`(`  `)`



#### 支持的类型

**布尔类型**：`t`  `T`  `true`  `True`  `TRUE`  `f`  `F`  `false`  `False`  `FALSE`

**数值类型**：整型、浮点型，eg. `1.6`  `10`  

**字符类型**：字符、字符串，eg. `'a'`  `"awesome"`

**数组**: eg. `["Tom", "Jim", "Sam"]`



#### 支持的内置函数

```go
// 数学运算
pow(float, float)
sin(float)
sqrt(float)
// 返回一个字符串长度
len(string)
// 返回小写字符串
lower(string)
// 子串b在字符串a中第一次出现的位置，如果没有返回-1
index(string, string)
// 字符串a是否含有子串b
contains(a string, b string)
// 字符串a是否以子串b开始/结束
has_prefix(a string, b string)
has_suffix(a string, b string)
```

