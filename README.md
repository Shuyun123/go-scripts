# go-scripts
> 总结开发中使用golang常用的命令



#### 使用

推荐将bin目录添加环境变量，每次添加新脚本直接编译可使用

```shell
# GoLang自定义脚本
export GOSCRIPT=/GO/src/scripts
export PATH=$GOSCRIPT/bin:$PATH:.
```



**注意：bin 下的二进制文件是Mac环境编译，如果运行Linux环境，需要执行类似命令**`env GOOS=linux GOARCH=amd64 go build -o ./bin/download ./download/download.go`



### 命令：

#### [jsondump](./json/README.MD)

说明：

```shell
jsondump -h
```



格式化json：

```shell
jsondump -d '{"name":  {"first":"Tom","last":"Anderson"},  "age":37,                                    [14:16:29]
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44}
  ]}'
```



将输出如下：

```json
{
	"name": {
		"first": "Tom",
		"last": "Anderson"
	},
	"age": 37,
	"children": ["Sara", "Alex", "Jack"],
	"fav.movie": "Deer Hunter",
	"friends": [
		{
			"first": "Janet",
			"last": "Murphy",
			"age": 44
		}
	]
}
```



压缩json：

```shell
jsondump -c=true -d '{"name":  {"first":"Tom","last":"Anderson"},  "age":37,                            [14:16:36]
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44}'
```



将输出如下：

```shell
{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Janet","last":"Murphy","age":44}]}
```



输出到文件：

```shell
jsondump -c=true -d '{"name":  {"first":"Tom","last":"Anderson"},  "age":37,                            [14:16:36]
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44} -o test.json
```



对文件json进行格式化或压缩：

格式化文件json：

```shell
jsondump -f test.json
```



压缩文件json：

```shell
jsondump -c=true -f test.json
```



格式化-i参数可以指定间隔，如：

```
jsondump -i "  " -f test.json
```



#### [download](./download/README.MD)

说明：

```
download -h
```



下载文件：

```
download -l https://studygolang.com/dl/golang/go1.15.4.src.tar.gz
```



通过-n可以指定并发数：

````
download -l https://studygolang.com/dl/golang/go1.15.4.src.tar.gz -n 8
````



通过-o指定下载目录地址，默认当前环境目录：

```
download -l https://studygolang.com/dl/golang/go1.15.4.darwin-amd64.tar.gz -n 5 -o test
```



多文件下载，通过-f指定URL集合文件，比如：urls.txt内容如下，

```
https://studygolang.com/dl/golang/go1.15.4.src.tar.gz
https://studygolang.com/dl/golang/go1.15.4.darwin-amd64.tar.gz
```



通过文件url地址集合下载：

```
download -f urls.txt -o test
```

## License
The Apache Software License, Version 2.0

Copyright [2020] [Anumbrella]

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

