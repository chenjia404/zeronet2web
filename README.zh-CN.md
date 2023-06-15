## zeront to web

输出zeronet的内容到web，因为搜索引擎只能索引这种格式。


### 命令行参数
| 字段        | 类型          | 说明                                                                     |
|-----------|-------------|------------------------------------------------------------------------|
| dir         | string          | zeronet的data目录，例如F:/soft/ZeroNet-win/data/ |
| -host         | string          | 会把内容里面的http://127.0.0.1:43110/ 替换成这个|
## 编译

go build -trimpath -ldflags="-w -s" 

### 运行
```
.\zeronet2web.exe  -dir "F:/soft/ZeroNet-win/data/" -host "https://zeronet.ipfsscan/"
```
