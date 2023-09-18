# syscalls
GO package for translating syscalls names to numbers and vice verse.
The currently supported architectures are:
* arm64
* amd64

If you need support for other architectures, please open an issue or a pull request.

I have used the following [source](https://github.com/hrw/syscalls-table) for the syscall tables.

## Usage

```go
package main

import (
	"fmt"

	"github.com/amitschendel/syscalls/pkg/syscalls"
)

func main() {
	fmt.Println(syscalls.GetNumberByName("", "openat"))
	fmt.Println(syscalls.GetNameByNumber("arm64", 56))
}
```

## Installation

```bash
go get github.com/amitschendel/syscalls
```


