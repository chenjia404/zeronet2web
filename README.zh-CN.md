## zeronet to web

输出zeronet的内容到web，因为搜索引擎只能索引这种格式。


### 命令行参数
| 字段        | 类型          | 说明                                                                     |
|-----------|-------------|------------------------------------------------------------------------|
| dir         | string          | zeronet的data目录，例如F:/soft/ZeroNet-win/data/ |
| host         | string          | 会把内容里面的http://127.0.0.1:43110/ 替换成这个|
| port         | int          | http web服务的端口号|
## 编译

go build -trimpath -ldflags="-w -s" 

### 运行
```
.\zeronet2web.exe  -dir "F:/soft/ZeroNet-win/data/" -host "https://zeronet.ipfsscan.io/"
```


### docker 运行
```
docker run -it  -v /mnt/f/soft/zeroNet-win/data/:/zeronet-data/ -v ./db/:/app/db/ -p 20236:20236  chenjia404/zeronet2web /zeronet-data/ -dir "/zeronet-data/"  -host "https://zeronet.ipfsscan.io/"
```
注意请把`/mnt/f/soft/zeroNet-win/data/` 修改成你的zeronet的data目录。

./db/ 是博客索引数据保存目录。

### releases

`goreleaser release --skip-publish --skip-validate --clean`

### 验证签名

```
gpg --recv-key E1346252ED662364CA37F716189BE79683369DA3

gpg --verify .\ethtweet_0.7.4_windows_amd64.zip.asc .\ethtweet_0.7.4_windows_amd64.zip
```
如果出现`using RSA key E1346252ED662364CA37F716189BE79683369DA3`就是验证成功