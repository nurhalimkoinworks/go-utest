```text
go test ./...
```

```text
go test -cover ./...
```

```text
go test -coverprofile fmtcoverage.html fmt
```

```text
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

```text
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

```text
go test -covermode=count -coverprofile=coverage.out
```