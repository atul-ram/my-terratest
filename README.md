# my-terratest
sample example to experiment with testing terraform deployment

## example1

## example2

## example3

## example4

## example5


### some usefull commands

go fmt
go mod tidy

# execution steps 

go to test folder path and run,

```
go test -v -timeout 10m
```

### Go
Golang will provide the runtime for running integration tests using the available Terratest modules.  Another utility called the `terratest_log_parser` will also be required for parsing Terratest output.
#### Install
```
brew install Go
```
Once it is installed, you will need to update your `.bash_profile` with the following:
```
export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```
Next we need to install some dependencies that will be used with go for working with Terraform and Azure.
```
go get github.com/gruntwork-io/terratest/modules/terraform
go get github.com/gruntwork-io/terratest/modules/azure
```

You can find more information on installing Go [here](https://golang.org/doc/install).

You can also just download the binary file [here](https://golang.org/dl/).
### Terratest Log Parser
The terratest_log_parser utility will allow us parse logs generated when running Terratest tests.

If you struggled with this last part, me too.  You should check out the Terratest documentation for [Debugging interleaved test output](https://terratest.gruntwork.io/docs/testing-best-practices/debugging-interleaved-test-output/#installing-the-utility-binaries) and the [Github repo](https://github.com/gruntwork-io/terratest/blob/master/modules/logger/parser/parser.go).

You can also just download the binary file [here](https://github.com/gruntwork-io/terratest/releases).



# References

- https://terratest.gruntwork.io/docs/testing-best-practices/debugging-interleaved-test-output/
- https://github.com/gruntwork-io/terratest/blob/master/test/terraform_packer_example_test.go
- https://terratest.gruntwork.io/docs/testing-best-practices/debugging-interleaved-test-output/
- 