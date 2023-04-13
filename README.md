quick demo of a CGO bug [ref](https://github.com/golang/go/issues/59606)

To Reproduce
- clone
- `go run .`
- will see:
```console
# example.com
./main.go:11:9: cannot use pkg.Foo() (value of type pkg._Ctype_int) as _Ctype_int value in return statement
```

