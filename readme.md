<div style="text-align:center"><img src="https://s3-us-west-2.amazonaws.com/nickwanninger/geode/masthead.png"/></div>

[![CircleCI](https://circleci.com/gh/nickwanninger/geode/tree/master.svg?style=svg)](https://circleci.com/gh/nickwanninger/geode/tree/master)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgeode-lang%2Fgeode.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgeode-lang%2Fgeode?ref=badge_shield)

## About

Geode is a programming language written in go around the llvm framework.
It's semi low level for the time being with plans to be higher level in
the future. It is just a compiler to llvm, then it calls clang to link the
.ll files to a runnable binary. Clang will also link the c standard library.

You can download (semi-regularly updated) binaries from the releases section,
but you might want to just install from source regardless. This is because
the compiler relies on the library files being in the `lib/` folder inside
the $GOPATH.

Geode is a heavy work in progress with apis that will change. Extended use is
not recommended at this stage.

Geode has a super basic minimal garbage collector wrapping around `malloc` and `free`.
In the backend, it is compiling [tgc](https://github.com/orangeduck/tgc) into your geode program.
The runtime will manage the setup and everything for you.

## Installing

Installing Geode is simple, just follow the steps below and install a few dependencies

### Dependencies

- Golang with a `$GOPATH` setup in your env
- The clang c compiler for linking binaries
- make

### Building

Once you have the dependencies setup, building is easy:

```bash
$ go get -u -d github.com/geode-lang/geode/...
$ cd $GOPATH/src/github.com/geode-lang/geode
$ make
$ sudo make install
```

This will build and install geode's executable binary to `/usr/local/bin`

## Example usage:

Geode is a massive work in progress, but you can look at example/main.g for a working state program

### Compiling a program

```
$ geode build <sourcefile>
```

Files can be any of the following:

- A folder containing a main.g
- A geode source file without the .g extension
- A .g file

### helloworld.g

```go
is main
include "std:io"

func main int {
	io:print("Hello, world\n");
	return 0;
}
```

### Array access and assigning

```go
is main

include "std:io"

func main(int argc, string* argv) int {
	argv[0] <- "program";
	for int i := 0; i < argc; i <- i + 1 {
		io:print("strlen('%s') -> %d\n", argv[i], strlen(argv[i]));
	}
	return 0;
}

func strlen(string str) int {
	int len;
	for len <- 0; str[len] != 0; len <- len + 1 {}
	return len;
}
```

### Unicode Example

```go
is main
func 유니코드테스트 void {}

func main int {
	유니코드테스트();
	return 0;
}
```

### Example fib.g

```go
is main

include "std:io"

func fib(int n) int {
	if n < 2 {
		return n;
	}
	return fib(n - 1) + fib(n - 2);
}

func main int {
	io:print("%d\n", fib(30));
	return 0;
}
```

### Linking

If you want, you can link to an external c file to use functions from that. For example

foo.c:

```c
int fourtytwo() {
	return 42;
}
```

foo.g:

```go
link "foo.c"

func fourtytwo() int ...

func main int -> fourtytwo();
```

Notice the `...`? That is the way of telling geode, "this function is external, and has
no body". If you do this, the function must be defined via a linkage. Otherwise the compiler
will crash


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgeode-lang%2Fgeode.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgeode-lang%2Fgeode?ref=badge_large)