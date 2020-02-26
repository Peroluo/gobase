# Grpc

#### 1.安装grpc

```shell
go get -u google.golang.org/grpc
```

#### 2. 安装Protocol Buffers v3

安装用于生成gRPC服务代码的协议编译器，最简单的方法是从下面的链接：https://github.com/google/protobuf/releases下载适合你平台的预编译好的二进制文件（`protoc--.zip`）。

下载完之后，执行下面的步骤：

1. 解压下载好的文件
2. 把`protoc`二进制文件的路径加到环境变量中

#### 3. 安装执行proto文件编译插件

```she
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### 4. 执行编译

```she
protoc --go_out=plugins=grpc:. ./user.proto
```