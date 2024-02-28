# Patch

通过 Patch 更新结构体。

修改字段：

```console
$ go run github.com/BedivereZero/learning-notes/golang/json/patch -p patch-age.json
Original: {"name":"Flandre","age":495}
Modified: {"name":"Flandre","age":17}
```

新增字段：

```console
$ go run github.com/BedivereZero/learning-notes/golang/json/patch -p patch-sleep.json
Original: {"name":"Flandre","age":495}
Modified: {"name":"Flandre","age":495,"sleep":true}
```

修改多个字段：

```console
$ go run github.com/BedivereZero/learning-notes/golang/json/patch -p patch-sleep-and-age-0.json
Original: {"name":"Flandre","age":495}
Modified: {"name":"Flandre","sleep":true}
```
