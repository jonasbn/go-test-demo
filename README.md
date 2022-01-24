# go-test-demo

This is a basic repository I use for testing `gotest`, which provides colorful output, compared to the standard test tool.

Standard **go** toolchain:

```bash
go test -v
```

![standard go test example screenshot](test-go-test-demo.png)

Using `gotest`

```bash
gotest -v
```

![colorful gotest example screenshot](gotest-go-test-demo.png)

I use this repository for testing my contributions to `gotest`

```bash
gotest -v github.com/jonasbn/go-test-demo
```

Do note that the above example expects for you to be in a directory containing a Go module project or else you will observe the following error:

> no required module provides package github.com/jonasbn/go-test-demo: go.mod file
> not found in current directory or any parent directory; see 'go help modules'

You can do the following to get the demo to work, if you are not in the context of a Go module - it could be any, you could even fork and clone this here repository.

Create a directory

```bash
mkdir gonow
```

Initialize as a Go module

```bash
go mod init example.com/gonow
```

Install this demo

```bash
go get github.com/jonasbn/go-test-demo
```

And then you should be good to go

```bash
gotest -v github.com/jonasbn/go-test-demo
```

## Resources and References

- [GitHub: rakyll/gotest](https://github.com/rakyll/gotest)
- [Go primary site](https://golang.org/doc/)
- [Go testing package](https://golang.org/pkg/testing/)
