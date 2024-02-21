# 检查安装goctl
installgoctl:
	$(if $(shell command -v "goctl"), ,GO111MODULE=on go install github.com/zeromicro/go-zero/tools/goctl@latest)

# 检查安装protoc
installprotoc:
	$(if $(shell command -v "protoc"), ,goctl env check --install --verbose --force)

# 本地运行
run:
	go run {{template}}.go -f etc/{{template}}-api.yaml
# 本地编译
build:
	go build -ldflags="-s -w" {{template}}.go
	$(if $(shell command -v upx), upx {{template}})
	
# 编译mac执行文件
build.mac:
	GOOS=darwin go build -ldflags="-s -w" -o {{template}}-darwin {{template}}.go
	$(if $(shell command -v upx), upx {{template}}-darwin)

# 编译win执行文件
build.win:
	GOOS=windows go build -ldflags="-s -w" -o {{template}}.exe {{template}}.go
	$(if $(shell command -v upx), upx {{template}}.exe)

# 编译linux执行文件
build.linux:
	GOOS=linux go build -ldflags="-s -w" -o {{template}}-linux {{template}}.go
	$(if $(shell command -v upx), upx {{template}}-linux)

# 生成接口文档
api.doc: installgoctl
	goctl api doc --dir ./api --o ./doc

# 根据api文件，重新生成go代码
api.go: installgoctl
	goctl api go --api api/init.api --dir .

# 生成Dockerfile
docker: installgoctl
	goctl docker

# 生成mysql对象操作
dao: installgoctl
	goctl model mysql ddl --src sql/init_latest.sql --dir internal/dao