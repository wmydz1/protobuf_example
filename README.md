##
在 Golang 中使用 Protobuf
发布于 2014 年 10 月 11 日2014 年 12 月 09 日 作者 name5566
https://github.com/golang/protobuf 项目为 Golang 提供了 Protobuf 的支持。

----

安装 goprotobuf
从 https://github.com/google/protobuf/releases 获取 Protobuf 编译器 protoc（可下载到 Windows 下的二进制版本）
获取 goprotobuf 提供的 Protobuf 编译器插件 protoc-gen-go（被放置于 $GOPATH/bin 下，$GOPATH/bin 应该被加入 PATH 环境变量，以便 protoc 能够找到 protoc-gen-go）
go get github.com/golang/protobuf/protoc-gen-go
此插件被 protoc 使用，用于编译 .proto 文件为 Golang 源文件，通过此源文件可以使用定义在 .proto 文件中的消息。

获取 goprotobuf 提供的支持库，包含诸如编码（marshaling）、解码（unmarshaling）等功能
go get github.com/golang/protobuf/proto
使用 goprotobuf
这里通过一个例子来说明用法。先创建一个 .proto 文件 test.proto：

```
package example;
	
enum FOO { X = 17; };
	
message Test {
    required string label = 1;
    optional int32 type = 2 [default=77];
    repeated int64 reps = 3;
    optional group OptionalGroup = 4 {
        required string RequiredField = 5;
    }
}

```

编译此 .proto 文件：

protoc --go_out=. *.proto
这里通过 go_out 来使用 goprotobuf 提供的 Protobuf 编译器插件 protoc-gen-go。这时候我们会生成一个名为 test.pb.go 的源文件。

在使用之前，我们先了解一下每个 Protobuf 消息在 Golang 中有哪一些可用的接口：

每一个 Protobuf 消息对应一个 Golang 结构体
消息中域名字为 camel_case 在对应的 Golang 结构体中为 CamelCase
消息对应的 Golang 结构体中不存在 setter 方法，只需要直接对结构体赋值即可，赋值时可能使用到一些辅助函数，例如：
msg.Foo = proto.String("hello")
消息对应的 Golang 结构体中存在 getter 方法，用于返回域的值，如果域未设置值，则返回一个默认值
消息中非 repeated 的域都被实现为一个指针，指针为 nil 时表示域未设置
消息中 repeated 的域被实现为 slice
访问枚举值时，使用“枚举类型名_枚举名”的格式（更多内容可以直接阅读生成的源码）
使用 proto.Marshal 函数进行编码，使用 proto.Unmarshal 函数进行解码
现在我们编写一个小程序：

```
package main
 
import (
    "log"
    // 辅助库
    "github.com/golang/protobuf/proto"
    // test.pb.go 的路径
    "example"
)
 
func main() {
    // 创建一个消息 Test
    test := &example.Test{
        // 使用辅助函数设置域的值
        Label: proto.String("hello"),
        Type:  proto.Int32(17),
        Optionalgroup: &example.Test_OptionalGroup{
            RequiredField: proto.String("good bye"),
        },
    }
 
    // 进行编码
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
 
    // 进行解码
    newTest := &example.Test{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
 
    // 测试结果
    if test.GetLabel() != newTest.GetLabel() {
        log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
    }
}

```