run
`go mod init` `any-name`

-   to install go-sdk
`go get github.com/aws/aws-sdk-go/aws` 

- save the file once
 `go mod tidy`,
 `go build` to see if any dependencies to download

- provide your instance-id in go code
 `aws.String("your-instance-id")`

- Comment the code to avoid run, i.e.,when running the instance comment the stop code.

Ultimately `go run main.go`

