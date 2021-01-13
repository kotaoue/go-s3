# go-s3
The practice for AWS S3 using with AWS-SDK-GO.

## Usage
```bash
# has ~/.aws/credentials
$ go run main.go -b=xxx-bucket -k=hoge/piyo.html
# use assume-role
$ assume-role hoge go run main.go -b=xxx-bucket -k=hoge/piyo.html
```

## Reference
* [remind101/assume-role](https://github.com/remind101/assume-role)
* [Go言語(golang)でS3からファイルを取得する](https://dev.classmethod.jp/articles/golang-read-s3/)