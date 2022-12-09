# go-beginner

# Build

建制環境時, 可以代入版本號

```shell
go build -ldflags="-X go-beginner/env.Version=1.0.0 -X go-beginner/env.BuildTime=$(TZ=UTC8; date +'%Y-%m-%dT%T%z')" main.go
```
