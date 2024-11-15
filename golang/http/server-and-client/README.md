# HTTP

通过 Unix Socket 文件使用 HTTP。

在当前目录执行命令 `go build` 编译得到二进制 `http`。

## 基于 Unix Socket

参数 `--network=unix` 指定使用 Unix Socket。

参数 `--address=test.socket` 指定网络地址即 Unix Socket 文件的路径。

启动服务端

```bash
./http server --network=unix --address=test.socket
```

客户端发送 HTTP Request、并把接收到的 HTTP Response 的 Body 打印到 STDOUT。

```bash
./http client --network=unix --address=test.socket --method=GET --path=/
```
