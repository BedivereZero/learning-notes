# Golang Gin

## 服务端处理请求超时后，停止正在进行中的第三方调用

1. 在第三方调用时，把 `gin.Context.Request.Context()` 作为`context.Context` 参数哦。
2. 将 `gin.Engine.ContextWithFallback` 设为 `true`，`gin.Context` 可以作为 `context.Context` 参数。
