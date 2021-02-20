# proto自动生成go代码

1. 下载protobuf的编译器protoc
    [下载地址](https://github.com/google/protobuf/releases)
    将bin目录下的`protoc`文件加入$(GOPATH)/bin文件夹下面

2. 将GOPATH加入环境变量
    ```
    export GOPATH=$HOME/go   #默认安装包的路径
    export PATH=$PATH:$GOPATH/bin
    ```
   
3. 安装protoc-gen-go插件： `go get -u github.com/golang/protobuf/protoc-gen-go`, 
安装完会在$(GOPATH)/bin下面生成`protoc-gen-go`二进制文件

4. 安装protoc-gen-micro插件： `go get github.com/micro/protoc-gen-micro/v2`

5. 创建proto文件，运行命令：
    ```
    protoc -I ./ --go_out=./ --micro_out=./ ./*.proto
    ```
    -I： 指定import路径，可以指定多个-I参数，编译时按顺序查找，不指定时默认查找当前目录
    --go_out: protoc-gen-go文件输出的目录
    --micro_out: protoc-gen-micro文件输出的目录

## 报错：
`WARNING: Missing ‘go_package‘ option in “message.proto“` 需要在proto文件的syntax下加入：`option go_package = "./;message";`
go_package="go文件存放的路径;go所属包名"