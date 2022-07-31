# 生成api业务代码
goctl api go -api *.api -dir . --style=goZero

# 生成rpc业务代码 (goctl >= 1.3)
goctl rpc protoc *.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero

# 生成model业务代码
goctl model mysql ddl -src="*.sql" -dir="." -c