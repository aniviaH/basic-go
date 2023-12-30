# 使用Make工具，批量管理命令

.PHONY: docker
docker:
	# 把上次编译的东西删掉
	@rm webook || true
	# 运行一下 go mod tidy, 防止 go.sum 文件不对，编译失败
	@go mod tidy
	# 指定编译的目标系统：linux，目标系统的架构：arm，编译的文件和输出的文件名
	@GOOS=linux GOARCH=arm go build -o webook  .
	# 可以先将之前构建出来的 docker image 删除一下。-f: 强制删除
	@docker rmi -f flycash/webook:v0.0.1
	# 这里版本号tag可以随便修改，记得对应的 K8S 部署里面也要修改
	@docker build  -t flycash/webook:v0.0.1 .
