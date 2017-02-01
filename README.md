# api-gateway-iam-auth-request-sample
Sample code that sends a GET request to an API Gateway API with IAM auth (Sigv4).

## Usage

```sh
$ go build main.go

$ ./main --help
Usage of ./main:
  -c string
        shared credentials file path
  -p string
        profile name of credentials (default "default")
  -r string
        API region (default "us-east-1")
  -u string
        endpoint URL of API Gateway

$ ./main -u https://xxxxx.execute-api.ap-northeast-1.amazonaws.com/dev/test -c ~/.aws/credentials -r ap-northeast-1
```
