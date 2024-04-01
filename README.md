# plugin-reportor
数据同步上报插件——使得m7s实例可作把内存中信息同步到其他数据库中去：目前支持reids,redis-clusters



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

origin代表源服务器拉流地址前缀，可以由如下几种格式：
```go 
type M7sRpcConfig struct {
    Name     string
    ListenOn string
    Timeout  int64 `json:",default=2000"`
}
```