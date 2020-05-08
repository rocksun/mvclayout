# 开发环境搭建

## 前置条件

安装好了 Docker Desktop，Docker 运行正常。

安装好了 go 环境，配置好了 goproxy ，可以执行 go install 命令。

## 安装 go 监控工具

因为 Windows 下的不会将文件系统的变更，所以需要工具将 Windows 宿主机的文件变更事件通知到容器中。所以需要安装 docker-windows-volume-watcher 。

```powershell
go install github.com/FrodeHus/docker-windows-volume-watcher
```

另外，还有个同名 Python 工具，作用一样可以通过 pip install docker-windows-volume-watcher 安装。

## 开始开发

在项目目录新开一个命令行窗口，执行：

```powershell
docker-windows-volume-watcher -container mvclayout_mvclayout_api_1
```

然后另一个窗口执行：

```
docker-compose up --build
```

如果有问题，可以进入容器查看：

```powershell
docker exec -it mvclayout_mvclayout_api_1 /bin/bash
```