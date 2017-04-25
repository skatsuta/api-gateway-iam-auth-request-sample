# api-gateway-iam-auth-requester

An HTTP client tool that sends HTTP requests to an API Gateway API with IAM authentication (AWS Signature V4).

## Usage

```sh
$ go get -u -v github.com/skatsuta/api-gateway-iam-auth-requester

$ api-gateway-iam-auth-requester --help
Usage of api-gateway-iam-auth-requester:
  -c string
        shared credentials file path
  -m string
        HTTP method (default "GET")
  -p string
        profile name of credentials (default "default")
  -r string
        API region (default "us-east-1")
  -u string
        endpoint URL of API Gateway
  -v    verbose output

$ api-gateway-iam-auth-requester \
	-u https://xxxxx.execute-api.ap-northeast-1.amazonaws.com/dev/test \
	-r ap-northeast-1
```
