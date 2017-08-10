## Build

Build Env: ubuntu 16.04

- apt-get install libseccomp-dev pkg-config
- go get github.com/seccomp/libseccomp-golang
- go build .

## Run

`./myStrace echo test`

Output:

```
Run:  [echo test]
Wait err stop signal: trace/breakpoint trap
name: execve, id: 59
name: brk, id: 12
name: access, id: 21
name: mmap, id: 9
name: access, id: 21
name: open, id: 2
name: fstat, id: 5
name: mmap, id: 9
name: close, id: 3
name: access, id: 21
name: open, id: 2
name: read, id: 0
name: fstat, id: 5
name: mmap, id: 9
name: mprotect, id: 10
name: mmap, id: 9
name: mmap, id: 9
name: close, id: 3
name: mmap, id: 9
name: arch_prctl, id: 158
name: mprotect, id: 10
name: mprotect, id: 10
name: mprotect, id: 10
name: munmap, id: 11
name: brk, id: 12
name: brk, id: 12
name: open, id: 2
name: fstat, id: 5
name: mmap, id: 9
name: close, id: 3
name: fstat, id: 5
test
name: write, id: 1
name: close, id: 3
name: close, id: 3
        1|read
        1|write
        3|open
        5|close
        4|fstat
        7|mmap
        4|mprotect
        1|munmap
        3|brk
        3|access
        1|execve
        1|arch_prctl
```