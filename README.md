`Boss` 爬虫.

### 使用
下载 `boss` 执行文件，配置 `.env` 文件，然后执行.
```
$ cp .env.example .env
$ ./bin/boss
```
> 在 `.env` 文件所在目录执行 `boss` 文件

### 说明
`boss` 文件的编译
```
$ CGO_ENABLED=0 GOOS="linux" ARCH="amd64" go build -o bin/boss
```

生成的是 `Linux x86_64` 平台的可执行文件，其他平台可根据需要下载源码进行「交叉编译」.
