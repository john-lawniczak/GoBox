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

### Channels
They provide a way for goroutines to synchronize and exchange data, enabling safe and efficient concurrent programming.

- **Safe Data Exchange**: Channels ensure safe data exchange between goroutines, preventing common concurrency issues like race conditions.
- **Synchronization**: They help synchronize execution across different goroutines, making it easier to coordinate concurrent tasks.
- **Blocking and Non-Blocking Operations**: Channels support both blocking and non-blocking operations, giving you control over goroutine communication.
- **Buffered and Unbuffered**: Go provides both buffered and unbuffered channels, allowing for flexible communication patterns and control over the capacity and flow of data.
- **Promoting Concurrency**: Channels are integral to Go's philosophy of using communication to share memory, rather than sharing memory to communicate. This promotes a more structured approach to concurrency.

----------
# Dining Philosopher Problem
The Dining Philosopher Problem is a classic synchronization problem proposed by Dijkstra. It involves a scenario where a certain number of philosophers are sitting around a circular table doing one of two things: thinking or eating. They use a fork to eat and there's one fork between each pair of adjacent philosophers. The problem demonstrates the challenges of resource allocation and avoiding deadlock, where philosophers could end up waiting indefinitely for a fork.

Key Points:
- Challenges: Deadlock, resource allocation.
- Solution Concepts: Resource hierarchy, philosopher states, resource arbitration.

-----------
# Sleeping Barber Problem
The Sleeping Barber Problem is a classic inter-process communication and synchronization problem between multiple customer threads and a barber thread. The barber in the barber shop sleeps when there are no customers and is woken up by a customer. If multiple customers arrive while the barber is busy, they either wait in a waiting room (if there are free chairs) or leave (if all chairs are occupied). This problem illustrates the challenges in coordinating access to a shared resource (barber) in a concurrent system.

Key Points:
- Challenges: Coordination, managing access to a limited resource (barber).
- Solution Concepts: Semaphore, mutex, producer-consumer problem.

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

