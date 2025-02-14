# tdd-practice-app

## Error

```sh
go test -timeout 30s -run ^TestGETPlayers$ github.com/mzzz-zzm/tdd-practice-app/httpserver
```

```
runtime
/usr/local/go/src/runtime/map.go:877:6: mapiterinit redeclared in this block
:
:
app/httpserver [build failed] FAIL
```

This error is related to the Go runtime and indicates that there are conflicting declarations within the Go standard library.

## Solution (but need to resolve without this workaround!!!)

```sh
sudo rm -rf /usr/local/go
wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source ~/.profile
go clean -cache -modcache -i -r
cd /workspaces/tdd-practice-app
go test ./httpserver -timeout 30s -run ^TestGETPlayers$
```
