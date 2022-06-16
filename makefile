add:#将文件添加到暂存区
	git add .
format:#检查代码质量并格式化代码
	goimports -w . && gofmt -w . && golangci-lint run
pull:#拉取代码
	git fetch origin master && git rebase origin/master
swag:#生成接口文档
	swag init && swag fmt
install_golang-cli: # 安装golang-cli工具，用于静态检查代码质量
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
goimports_install: # goimports安装
	go get golang.org/x/tools/cmd/goimports
sqlc:#sqlc生成代码
	sqlc generate && git add .
run:
	go build -o bin/main main.go && ./bin/main