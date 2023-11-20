## 生成代码的命令
* protoc --go-grpc_out=. hello.proto
* protoc --go_out=. hello.proto
## Grpc 自己的理解 与 openFeign区别
* Grpc 自己理解：
    * 基于HTTP2.0，使用protobuf进行数据交互，使用HTTP2.0的特性进行数据传输，使用HTTP2.0的特性进行
    * 它更像是 java中的netty的使用方式 不能默认提供 restFul的api接口
    * 假如要提供restFul的api接口，
      * 1、需要使用 grpc的Gateway进行转发
      * 2、可以自己gin框架提供一个接口 这个接口包住rpc服务的接口即可
    