# go-project

go 模板工程

---

## make

Linux only.

> ⚠️ 修改可执行文件名称

```shell
build                          build the go packages
check                          run all linters
clean                          clean up binaries, releases and logs
clean-test                     clean up debris from previously failed tests
test                           run tests, except integration tests and tests that require root
vendor                         ensure all the go.mod/go.sum files are up-to-date including vendor/ directory
verify-vendor                  verify if all the go.mod/go.sum files are up-to-date
```

---

## cross compile 交叉编译

> ⚠️ 修改可执行文件名称

### Linux/MacOS

`cross_compile.sh`

### Windows

`cross_compile.bat`
