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
 
Keyword: `go`

Sync Package   
- Mutex (semaphors)

WaitGroups
  - import time
  - sync.WaitGroup
  - don't copy them
  - no guaranteed order
  - `defer` keyword

-----------
# Producer/Consumer Problem

-----------
# Dining Philosopher Problem

-----------
# Sleeping Barber Problem

------
# Tests
