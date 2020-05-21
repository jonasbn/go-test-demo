# go-test-demo

This is a basic repository I use for testing `gotest`, which provides colorful output, compared to the standard test tool.

Standard **go** toolchain:

```bash
$ go test -v
=== RUN   TestPass
--- PASS: TestPass (0.00s)
=== RUN   TestSkip
    TestSkip: testdemo_test.go:13: Skipping test
--- SKIP: TestSkip (0.00s)
=== RUN   TestFail
    TestFail: testdemo_test.go:19: Adder(2, 2) = 4, want 3
--- FAIL: TestFail (0.00s)
FAIL
exit status 1
FAIL	adder	0.153s
```
Using `gotest`

```
$ gotest -v
```

![colorful gotest example screenshot](gotest-go-test-demo.png)
