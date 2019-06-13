

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

### Type Mappings

Go types will be mapped to C-types according to the following table:

Go type                                       | C Type
--------------------------------------------- | ------
`int8`, `byte`                                | `char`
`uint8`, `bool`                               | `unsigned char`
`int16`                                       | `short`
`uint16`                                      | `unsigned short`
`int32`                                       | `int`
`uint32`                                      | `unsigned int`
`int64`                                       | `long`
`uint64`                                      | `unsigned long`
`float32`                                     | `float`
`float64`                                     | `double`
`[]`, `uintptr`, `reflect.UnsafePointer`, `*` | `*`

The last line means that slices and pointers are mapped to pointers in C. Pointers to structs are possible.

Passing `struct`, `complex`, and callback functions is not (yet) supported.

> **WARNING** `struct`s that are referenced **must** follow C alignment rules! There is **no** type checking, since this is actually not possible due to libraries not knowing their types...

Go `int` was deliberately left out to avoid confusion, since it has different sizes on different architectures.