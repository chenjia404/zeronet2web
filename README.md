[English](./README.md) | [简体中文](./README.zh-CN.md)
## zeront to web

Output the content of zeronet to the web


### Command Line Arguments命令行参数
| Field     | Type             | Description                                                                          |
|-----------|-------------|------------------------------------------------------------------------|
| dir         | string          | zeronet's data directory，For example `F:/soft/ZeroNet-win/data/` |
| -host         | string          | will put the content inside `http://127.0.0.1:43110/` replace with this|

### build

` go build -trimpath -ldflags="-w -s" `


### run
```
.\zeronet2web.exe  -dir "F:/soft/ZeroNet-win/data/" -host "https://zeronet.ipfsscan/"
```
