# GoBox
Playing with Go

Pass by **value** language      

Getting Started:
1. `brew install make` - (Mac)
2. `go mod init <filename>` - creates a mod file named filename
3. create a main.go file
4. `go run .` - runs 

------
# Concurrency
[Working with Concurrency in Go](https://www.udemy.com/course/working-with-concurrency-in-go-golang/)
- Share memory by communicating
- If you don't need it, don't use it. Simplify.
- GoRoutines - lightweight threads
    - a group is called CoRoutines
 
Keyword: `go` and `go run -race .`

Sync Package   
- Mutex (semaphors)

### WaitGroups
  - import time
  - sync.WaitGroup
  - don't copy them
  - no guaranteed order
  - `defer` keyword

-----------
# Producer/Consumer Problem

-----------
# Dining Philosopher Problem
- Dikstra's Algo
- Logical race condition

-----------
# Sleeping Barber Problem

------
# Tests

------
# Format Verbs for fmt package
There are more, but these are common:
| Verb | Description                                                      |
|------|------------------------------------------------------------------|
| `%v`  | The value in a default format                                    |
| `%+v` | The value in a default format plus field names for structs       |
| `%#v` | A Go-syntax representation of the value                          |
| `%T`  | A Go-syntax representation of the type of the value              |
| `%d`  | Integer in base 10                                              |
| `%s`  | String or slice of bytes (interpreted as UTF-8)                 |
| `%f`  | Floating-point number in decimal (default precision: 6)         |
| `%t`  | Boolean, true or false                                          |
| `%b`  | Base 2                                                          |
| `%x`  | Base 16, lowercase letters for a-f                              |
| `%X`  | Base 16, uppercase letters for A-F                              |
| `%p`  | Pointer address in base 16 with leading 0x                      |

