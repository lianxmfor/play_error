# play_error
pkg/errors How to print custom wrapped errors?

output:

```shell
$ go run main
Error1: something not found
main.findSomething
        /home/lianxm/github.com/play_error/main.go:11
main.main
        /home/lianxm/github.com/play_error/main.go:15
runtime.main
        /usr/local/go/src/runtime/proc.go:255
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1581

Error2: something not found⏎
```
