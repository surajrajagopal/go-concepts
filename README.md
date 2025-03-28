# go-concepts
In this repo, we can find all concepts of go and implementation

Definition of all the concept : 

1. Variables - Variable is the name given to a memory location to store a value of a specific type.

2. Constant - The term constant is used for fixed values.

   Note : we cannot update constant value and we cannot assign runtime value to constant variable.

4. Function - A function is a block of code is used to perform specific task. At a high level, a function takes an input and returns an output. Function allows you to extract common peice of code into a single component.

     Named return values :
     It is possible to return named values from a function. If a return value is named, declared as a variable in the first line of the function.

     _ is known as the blank identifier in Go. It can be used in place of any value of any type.

6. Packages - Packages are used to organize Go source code for better reusability and readability. Packages are a collection of Go sources files that reside in the same directory. hence it is easy to maintain go projects.
   
   Go modules is a way of dealing with dependencies in golang.

   A Go Module is nothing but a collection of Go packages. Now this question might come to your mind. Why do we need Go modules to create a custom package? The import path of a custom package in Go is based on the module name. The go.mod file also manages all third- 
   party dependencies (like GitHub packages) and their versions for the application. This go.mod file is created when we create a new module.

8. Array - An array is a collection of similar datatype with fixed-size, stored in a continuous memory block in a specific order..

9. Slice - A slice in Go is a flexible, dynamic view of an underlying array. Unlike arrays, slices don’t own data but act as references to an array. The size of a slice can grow or shrink as needed.

         Internally a slice is represented by three things : 

         Pointer → Points to the starting element of the underlying array.
         Length (len) → Number of elements in the slice.
         Capacity (cap) → Maximum elements the slice can hold before reallocation.
         Capacity is calculated from the starting index of the slice to the end of the original array.
         When capacity is exceeded, Go doubles the capacity automatically for performance reasons.
   
         Formula Calculate len and cap : 

         length of newly created slice = (end–start)

         capacity of newly created slice = (length_of_array–start_index)



     Append : In Go, slices are dynamic, meaning we can add new elements using the append() function.

     If there is extra capacity, the new element is added without changing the underlying array.
   
     If the slice length exceeds capacity, Go creates a new array with double the capacity, copies the existing elements to the new array, and updates the slice to point to the new array.

     make : make() function in Go is used to create slices, maps, and channels with pre-allocated memory. Unlike new(), which only allocates memory, make() also initializes the object and returns a ready-to-use value.

     make(type, length, capacity) // For slices
   
     make(map[keyType]valueType)  // For maps
   
     make(chan dataType, buffer)  // For channels

     Note : Memory is allocated, but elements are initialized with zero values (0 for int).
   
     Copy function :
      The copy() function in Go is used to copy elements from one slice (src) to another (dst). It returns the number of elements copied, which is the minimum length of the two slices. 
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

Types of RPCs : 

Unary RPC
The simplest type of RPC where the client sends a single request and gets back a single response.

Once the client calls a stub method, the server is notified that the RPC has been invoked with the client’s metadata for this call, the method name, and the specified deadline if applicable. The server can then either send back its own initial metadata (which must be sent before any response) straight away, or wait for the client’s request message. Which happens first, is application-specific. Once the server has the client’s request message, it does whatever work is necessary to create and populate a response. The response is then returned (if successful) to the client together with status details (status code and optional status message) and optional trailing metadata. If the response status is OK, then the client gets the response, which completes the call on the client side.

Check examples : [unary](https://github.com/itsksaurabh/go-grpc-examples/tree/master/unary/greet)

Server streaming RPC
A server-streaming RPC is similar to a unary RPC, except that the server returns a stream of messages in response to a client’s request. After sending all its messages, the server’s status details (status code and optional status message) and optional trailing metadata are sent to the client. This completes processing on the server side. The client completes once it has all the server’s messages.

Check examples : [Server streaming](https://github.com/itsksaurabh/go-grpc-examples/tree/master/stream/server-streaming/countdown)

Client streaming RPC
A client-streaming RPC is similar to a unary RPC, except that the client sends a stream of messages to the server instead of a single message. The server responds with a single message (along with its status details and optional trailing metadata), typically but not necessarily after it has received all the client’s messages.

Check examples : [Client streaming](https://github.com/itsksaurabh/go-grpc-examples/tree/master/stream/client-streaming/sumAll)

Bidirectional streaming RPC
In a bidirectional streaming RPC, the call is initiated by the client invoking the method and the server receiving the client metadata, method name, and deadline. The server can choose to send back its initial metadata or wait for the client to start streaming messages.

Client- and server-side stream processing is application specific. Since the two streams are independent, the client and server can read and write messages in any order. For example, a server can wait until it has received all of a client’s messages before writing its messages, or the server and client can play “ping-pong” – the server gets a request, then sends back a response, then the client sends another request based on the response, and so on.

Check examples : [Bidirectional streaming](https://github.com/itsksaurabh/go-grpc-examples/tree/master/stream/bi-directional-streaming/feeds)

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


