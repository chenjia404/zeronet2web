[English](./README.md) | [简体中文](./README.zh-CN.md)
## zeronet to web

Output the content of zeronet to the web, because search engines can only index this format.


### Command Line Arguments命令行参数
| Field     | Type             | Description                                                                          |
|-----------|-------------|------------------------------------------------------------------------|
| dir         | string          | zeronet's data directory，For example `F:/soft/ZeroNet-win/data/` |
| host         | string          | will put the content inside `http://127.0.0.1:43110/` replace with this|
| port         | int          | web port|

### build

` go build -trimpath -ldflags="-w -s" `


### run
```
.\zeronet2web.exe  -dir "F:/soft/ZeroNet-win/data/" -host "https://zeronet.ipfsscan.io/"
```

### docker run
```
docker run -it  -v /mnt/f/soft/zeroNet-win/data/:/zeronet-data/ -v ./db/:/app/db/ -p 20236:20236  chenjia404/zeronet2web /zeronet-data/ -dir "/zeronet-data/"  -host "https://zeronet.ipfsscan.io/"
```

Note that please change `/mnt/f/soft/zeroNet-win/data/` to the data directory of your zeronet.

./db/ is the blog index data storage directory.

### releases

`goreleaser release --skip-publish --skip-validate --clean`

### verify signature

```
gpg --recv-key E1346252ED662364CA37F716189BE79683369DA3

gpg --verify .\ethtweet_0.7.4_windows_amd64.zip.asc .\ethtweet_0.7.4_windows_amd64.zip
```
If `using RSA key E1346252ED662364CA37F716189BE79683369DA3` appears, the verification is successful