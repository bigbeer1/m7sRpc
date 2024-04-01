# plugin-m7sRpc
让M7s支持rpc和grpc调用



## 插件地址
https://github.com/bigbeer1/m7sRpc

## 插件引入

```go
import (
   _ "github.com/bigbeer1/m7sRpc"
)
```

## 配置

```yaml

m7srpc:
  name: m7s.rpc
  listenon: 0.0.0.0:8810
  timeout: 30000
```

```go 
type M7sRpcConfig struct {
    Name     string
    ListenOn string
    Timeout  int64 `json:",default=2000"`
}
```