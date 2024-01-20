### grpcurl 工具

参考：https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-08-grpcurl.html


1. 安装grpcul工具
```shell
$ go get github.com/fullstorydev/grpcurl
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```
2. 查看服务列表
```shell
$ grpcurl -plaintext localhost:50051 list
HelloService.HelloService
grpc.reflection.v1alpha.ServerReflection

```
2. 