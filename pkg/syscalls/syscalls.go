package syscalls

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gocarina/gocsv"
)

const (
	arm64 = "artifacts/syscalls-arm64.csv"
	amd64 = "artifacts/syscalls-x86_64.csv"
)

type Syscall struct {
	Name   string `csv:"name"`
	Number int    `csv:"number"`
}

func getSyscalls(filePath string) (syscalls []Syscall, err error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer csvFile.Close()

	if err := gocsv.UnmarshalFile(csvFile, &syscalls); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return syscalls, nil
}

func readSyscallArch(arch string) ([]Syscall, error) {
	if arch == "" {
		arch = runtime.GOARCH
	}

	switch arch {
	case "amd64":
		return getSyscalls(amd64)
	case "arm64":
		return getSyscalls(arm64)
	default:
		return nil, fmt.Errorf("unsupported architecture: %s", arch)
	}
}

func GetNumberByName(arch string, name string) (int, error) {
	syscalls, err := readSyscallArch(arch)
	if err != nil {
		return 0, err
	}

	for _, syscall := range syscalls {
		if syscall.Name == name {
			return syscall.Number, nil
		}
	}

	return 0, fmt.Errorf("syscall not found: %s", name)
}

func GetNameByNumber(arch string, num int) (string, error) {
	syscalls, err := readSyscallArch(arch)
	if err != nil {
		return "", err
	}

	for _, syscall := range syscalls {
		if syscall.Number == num {
			return syscall.Name, nil
		}
	}

	return "", fmt.Errorf("syscall not found: %d", num)
}
