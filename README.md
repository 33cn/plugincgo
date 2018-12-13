# plugincgo
plugin cgo 是用go 调用 c++ 库实现的插件

## 和plugin的区别

### 1. 可能会存在 c++ 和 c 代码的内存越界，缓冲区溢出等问题。

### 2. 编译的过程，可能需要先安装对应的依赖库。每个插件要有一个 Makefile , make install 要保证能安装好相关的依赖。

