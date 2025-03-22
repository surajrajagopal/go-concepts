# go-concepts
In this repo, we can find all concepts of go and implementation

Definition of all the concept : 

1. Variables - Variable is the name given to a memory location to store a value of a specific type.

2. Constant - The term constant is used for fixed values.

3. Function - A function is a block of code is used to perform specific task. At a high level, a function takes an input and returns an output. Function allows you to extract common peice of code into a single component.

4. Packages - Packages are used to organize Go source code for better reusability and readability. Packages are a collection of Go sources files that reside in the same directory. hence it is easy to maintain go projects.
Go modules is a way of dealing with dependencies in golang.

   A Go Module is nothing but a collection of Go packages. Now this question might come to your mind. Why do we need Go modules to create a custom package? The answer is the import path for the custom package we create is derived from the name of the go module. In 
   addition to this, all the other third-party packages(such as source code from github) along with their versions which our application uses will be managed by the go.mod file. This go.mod file is created when we create a new module.

6. Array - An array is a collection of elements of the same type. It is an ordered sequence of elements stored contiguously in memory.

7. Slice - A slice is a data type which points to an underlying array and Slices do not own any data on their own. They are just references to existing arrays.
           Unlike array, the size of a slice is flexible and can be changed/resized as per requirement. Slice internally represented by a slice header.

        Internally a slice is represented by three things : 

        Pointer - Pointer to the underlying array.

        len - number of elements in the slice.

        cap - The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
      
        It indicates the maximum number of elements these structures can hold before they need to allocate more memory.
      
        Capacity is calculated from the starting index of the slice to the end of the original array.

        When capacity is exceeded, Go doubles the capacity automatically for performance reasons.
   
       Formula Calculate len and cap : 

       length of newly created slice = (end–start)

       capacity of newly created slice = (length_of_array–start_index)


     Append - Slices are dynamic and new elements can be appended to the slice using append function.
     When slice length is greater than capacity - In this case since there is no more capacity, so no new elements can be accommodated.  So in this case under the hood an array double the capacity will be allocated. The 
     current array pointed by the slice will be copied to that new array. Now the slice will starting pointing to this new array. Hence the capacity will be doubled and length will be increased by 1

Copy function : 
   - Copy (dst, src)

7. Variadic function is a function that accepts a variable number of arguments.
- Pass last argument to a function which will accept variable number of arguments(...) ellpsis.
- Append slice to slice ex : append(sl, sl...)

8. Map - A map is a bulit-in data type in golang which stores the data into key-values pairs.
   - Map is a reference type in golang.
   - ok (idioms)
   - delete(map, key)

9. String - A string a slice of bytes in golang. it is enclosed between the double quotes.
   Ex : str := "Hello World"

   Byte and Rune DataType in Golang : 
   - To represent character we use byte and rune in golang.
   - Byte and rune are alias of uint8 and int32 datatype.
   - Both byte and rune data types are essentially integers.

   Byte data type - represents ASCII characters. (0 to 128)
   
   Rune data type- represents Unicode characters that are encoded in UTF-8 format. ()

   len - return number of bytes.
   
   For example, a byte variable with value 'a' is converted to the integer 97 while a rune variable with a unicode value '~' is converted to the corresponding unicode codepoint U+007E, where U+ means unicode and the numbers are hexadecimal, which is essentially an integer.

11. Pointers - A pointer is a variable is used to store the address of another variable.

12. Struct : Struct is a collection of disimilar datatypes. A struct is a user-defined type in golang which is used to group the data into a single unit rather than having seperate values.
Ex. Employee has different fields by using struct we will unit them.
    - Anonymous struct

13. Method - A method is a function which has special receiver type between func keyword and method name. This receiver type may be of struct type or non-struct type. Receiver Type can be of Value receiver or Pointer receiver.

11. Interface - An Interface is a type in go which is a collection of method signature. These 
collection of method signature are meant to represent certain behaviour. The interface declares only method set and any type implements all the methods of interface is said to be interface type.
   - Empty interface - An interface that has zero methods is called an empty interface. It is represented as interface{}. 
   - Type Assertion - used to extract the underlying value of the interface. Syntax: t.(type) 
   - Type Switch - A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. Syntax : t.(Type)

Concurrency In Golang: 

1. Difference between Concurrency and Parallism
   - Concurrency is the ability to run multiple tasks in overlapping time periods, but not necessarily at the same time. It enables efficient task management by switching between tasks, making it useful when tasks involve waiting (e.g., I/O operations, network calls).
     Code : [https://go.dev/play/](https://go.dev/play/p/KLyqTzmCFiU)

      How it works:

      Tasks run concurrently (overlapping), but not necessarily in parallel.

      The Go scheduler switches between goroutines while waiting.
   - Parallelism, on the other hand, refers to running multiple tasks simultaneously by utilizing multiple CPU cores, allowing for true parallel execution. While concurrency improves responsiveness, parallelism enhances performance by fully utilizing system resources.
     Code : [https://go.dev/play/](https://go.dev/play/p/AFuB1-XcqMg)

     How it works:

      runtime.GOMAXPROCS(2) enables parallel execution on 2 CPU cores.

      Tasks run truly simultaneously if multiple CPU cores are available.
     
2. Goroutines - A goroutines are like functions or methods that run concurrenctly with other functions and methods. 
Goroutines are lighweighted threads. 
The creation of goroutines are tiny/cheaper as compare to threads. 
To create goroutines in golang we use `go` keyword to the prefix of function or methods.
   Goroutines are lightweight, managed by the Go runtime,
   and consume much less memory compared to traditional OS threads. Unlike threads, which are scheduled by the OS, goroutines are scheduled by Go’s runtime scheduler, making them more efficient.
   Go’s scheduler follows the GPM model, where:

     G (Goroutine) represents the lightweight execution unit.
     M (Machine) represents an OS thread.
     P (Processor) is a logical processor that schedules goroutines on threads.
     The scheduler assigns goroutines to P, which then runs them on M. If a goroutine blocks (like waiting for I/O), the scheduler moves another goroutine to that thread to keep execution efficient.

3. Channels - Channels can be thought as a pipes by using which goroutines communications.
   - close() 
   - for range loop
   - Done/quit channel (bool/struct{})
   Channels in Go are used for communication between goroutines. They allow safe data transfer without explicit locking mechanisms.
   
   There are two types of channels:

      Unbuffered Channels – These don’t have a capacity. A sender must wait until a receiver reads the data. This ensures synchronization between goroutines.

      Buffered Channels – These have a specified capacity. A sender can send data without waiting, up to the buffer limit. Once full, the sender must wait for the receiver to read data before sending more.

   Difference between closing a channel and making a channel nil in Go

      Closing a Channel :
         - A channel is closed using close(channel).
         - Once closed, no more values can be sent to it, but it can still be read until it's empty.
         - Trying to send data to a closed channel causes a panic.
         - Receivers get a zero value when reading from a closed channel.

      Making a Channel Nil :
         - Setting a channel to nil means it has no memory allocated.
         - Sending or receiving on a nil channel blocks forever (deadlock).
         - This is useful to disable a channel dynamically.

5. Wait Group - A wait group is used to wait for collection of goroutions to finish their execution. The control block until all the goroutines completes their execution.
6. Mutex - A mutex is locking mechanism which ensure only one goroutine can access critical section of code at any point of time. This concepts used to prevent race condition from happening.
   How does Go handle goroutine synchronization to avoid race conditions?
    - Yes! Mutex (mutual exclusion) is one way to handle synchronization in Go. It ensures that only one goroutine can access a shared resource at a time. Go provides sync.Mutex for this purpose.
   Can you explain the difference between sync.Mutex and sync.RWMutex? When would you use RWMutex instead of Mutex?
    - sync.Mutex: This is a standard mutex that allows only one goroutine to access a resource at a time. If one goroutine locks it, all other goroutines must wait.
    - sync.RWMutex: This is a read-write mutex. It allows multiple goroutines to read a resource at the same time, but only one goroutine can write. If a goroutine acquires a write lock, all other reads and writes are blocked.
   When to use RWMutex?
    - Use sync.RWMutex when you have more reads than writes, so multiple goroutines can read concurrently without blocking each other. This improves performance.
   Note : RLock allows multiple readers at the same time, improving performance when there are more reads than writes.

7. Select - The select statement is used to choose a multiple send/receive channel operations.

   Select statement block the control until anyone of the operation is ready. If both the operations are ready then it picks randomly.
   
   for-select loop - We need to stop the for-select loop by using done/quit channel.

   The select statement in Go is used to wait on multiple channel operations and executes the first one that is ready. It helps handle multiple channels efficiently.

   Key Points About select:
     - If multiple channels are ready, one case is chosen randomly.
     - If no channels are ready, it blocks execution until one becomes available.
     - You can use a default case to prevent blocking.

Error Handling : 

defer - A defer used to delay the execution of a function/statement until the surrounding function completes. deferred function calls are executed in Last-In-First-Out (LIFO) order.

Panic -  a panic is a mechanism that allows you to halt the normal execution of a program when an unexpected or unrecoverable situation occurs.

Recover - recover() is a built-in function in Go that is used to regain control of a panicking goroutine.

OOPs in golang : 

1. Constructor - A constructor is a special function is used to initalize and create an instance of a struct type. To declare constructor in golang we use prefix "New" before the struct type.
   Ex. NewEmployee()

2. Polymorphism - ability of a message to be displayed in more than one form. (Same name many forms)

3. Composition - Composition can be achieved in Go is by embedding one struct type into another.

4. Encapsulation is defined as the wrapping up of data under a single unit.

5. Data Abstraction is the property by which only the essential details are displayed to the user. 

First Order Functions : 

1. Anonymous - An anonymous function is a function that was declared without any named identifier to refer to it.

2. User-Defined types - 

3. Higher Order functions - 

4. Closures - Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.


What is RPC?
Local Procedure Call - It is a function call within process to execute some code.

Remote Procedure Call - A Remote Procedure call enables one machine to invoke some code on another machine as it seems a local function call from user's perspective.

Grpc Framework : 
1. gRPC is a (Remote Procudure Call) open-source framework developed by Google. it has efficient communication and interaction between distributed systems.
2. Is built on top of the HTTP/2 protocol, which allows for high-performance, low-latency communication between client and server applications.

Key features of gRPC include : 
1. Language-agnostic : gRPC supports multiple programming languages.
2. Data Formatting (Protocol buffer or protobuf) is a very efficient binary encoding format -  gRPC is implemented through the use of protocol buffers and code generation, which provides a fast and efficient way to build client-server applications.
3. Automatic code generation: gRPC generates client and server code based on the service definition specified in the protobuf files.
4. Bi-directional streaming: It can use unary rpc's where here the client sends a single request and receives a single response.

json Encoding/Decoding :

- Encoding - struct to json

- Decoding - json to struct

Marshalling: the act of converting a Go data structure into valid JSON. 

Unmarshalling: the act of parsing a valid JSON string into a data structure in Go.

Golang Datatypes list : 

1.Primitive types :
  - Bool (false) 
  - Numeric type
    - int default value (0)
    - float (0) 
    - complex (0+0i)
    - byte represents - uint8 (0)
    - rune represents - int32 (0)
    - String emptystring = ""
2. Aggregate type (value type):
   - Struct default value - (0)
   - Array default value - (0)
3. Reference type : Default Value - Nil
   - slice 
   - map
   - pointers
   - channels
What is the difference between a normal function and a goroutine in Golang?
Key Differences Between Normal Functions and Goroutines in Golang:
Execution Style:

Normal functions execute synchronously in the same thread.
Goroutines execute asynchronously, allowing multiple tasks to run concurrently.
Thread Management:

  A normal function runs on the main thread.
  A goroutine runs on a separate lightweight thread managed by Go’s scheduler.

Performance:

Normal functions block execution until they finish.
Goroutines are non-blocking, so thousands can run efficiently without heavy memory usage.


