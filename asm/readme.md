

## Installation of required files

Install: YASM, GAS, CLANG and helper tools

Typically as or gas (GNU Assembler) will already be installed as it is part of binutils, but if need be, you can eg. do as follows

```bash
sudo apt-get install build-essential yasm
sudo apt install clang-7
go get -u github.com/minio/asm2plan9s
go get -u github.com/minio/c2goasm
go get -u github.com/klauspost/asmfmt/cmd/asmfmt
$ c2goasm -a -f _lib/sum_float64.s sum_float64.s
```

You can ask GCC to produce the assembly file, instead of an object file (or an executable).

For instance:

gcc -Wall -c test.c

Will produce an object file from test.c (test.o).

gcc -Wall -o test test.c

Will produce an executable file named 'test' from test.c

gcc -Wall -S test.c

Will produce an assembly file from test.c (test.s)

## cpp

### assemble
c++ -O3 -mavx -mfma -masm=intel -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -S $1
clang -S -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -masm=intel -O3 -m64 -mavx -mavx2 -msse4.1 -Wall -Wextra hex.cc
clang -S -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -masm=intel -O3 -m64 -mavx -mavx2 -msse4.1 -Wall -Wextra hex.cc

### convert
c2goasm -a hex2.s Hex2_amd64.s

## Working steps:

```bash
apt-get install clang yasm
```

```c
void toUpper(char* src, char* result){
    for(char* p=src; *p != '\0'; p++){
        if(*p >= 'a' && *p <= 'z')  //Only if it's a lower letter
          *p -= 32;
    }
    result = src;
}
```

```go
//+build !noasm
//+build !appengine

package hex


import (
	"unsafe"
)

//go:noescape
func _toUpper(src, result unsafe.Pointer)

func ToUpper(src []byte) []byte {
	var result []byte
	_toUpper(unsafe.Pointer(&src), unsafe.Pointer(&result))
	return result
}
```

### Using `clang`
```bash
clang -S -O3 -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti $file.c
```

```bash
clang -S -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -masm=intel -mno-red-zone -O3 -m64 -mavx -mavx2 -msse4.1 -Wall -Wextra -mstackrealign -mllvm -inline-threshold=1000 file.c
```

```bash
c2goasm -a -c -s -f strings.s strings_amd64.s
```

### Using GCC

```bash
gcc -S -O3 -masm=intel -mno-red-zone -mstackrealign -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti strings.c
```

If you specify target CPU architecture:

```bash
gcc -S -O3 -masm=intel -march=skylake -mno-red-zone -mstackrealign -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti strings.c
```