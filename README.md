# go-concepts
In this repo, we can find all concepts of go and implementation

Definition of all the concept : 

1. Variables - Variable is the name given to a memory location to store a value of a specific type.

2. Constant - The term constant is used for fixed values.

3. Function - A function is a block of code is used to perform specific task. At a high level, a function takes an input and returns an output. Function allows you to extract common peice of code into a single component.

4. Packages - Packages are used to organize Go source code for better reusability and readability. Packages are a collection of Go sources files that reside in the same directory. hence it is easy to maintain go projects.
Go modules is a way of dealing with dependencies in golang.

5. Array - An array is a collection of elements of the same type. It is an ordered sequence of elements stored contiguously in memory.

6. Slice - A slice is a data type which points to an underlying array and Slices do not own any data on their own. They are just references to existing arrays.
Unlike array, the size of a slice is flexible and can be changed/resized as per requirement. Slice internally represented by a slice header.

Internally a slice is represented by three things : 
Pointer - Pointer to the underlying array
len - number of elements in the slice.
cap - The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.

Formula Calculate len and cap : 
length of newly created slice = (end–start)
capacity of newly created slice = (length_of_array–start_index)

Append - Slices are dynamic and new elements can be appended to the slice using append function.
When slice length is greater than capacity - In this case since there is no more capacity, so no new elements can be accommodated.  So in this case under the hood an array of double the capacity will be allocated. The current array pointed by the slice will be copied to that new array. Now the slice will starting pointing to this new array. Hence the capacity will be doubled and length will be increased by 1

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

   byte data type - represents ASCII characters. (0 to 128)
   rune data type- represents Unicode characters that are encoded in UTF-8 format. ()
   len - return number of bytes.
   
   For example, a byte variable with value 'a' is converted to the integer 97 while a rune variable with a unicode value '~' is converted to the corresponding unicode codepoint U+007E, where U+ means unicode and the numbers are hexadecimal, which is essentially an integer.

10. Pointers - A pointer is a variable is used to store the address of another variable.

11. Struct : Struct is a collection of disimilar datatypes. A struct is a user-defined type in golang which is used to group the data into a single unit rather than having seperate values.
Ex. Employee has different fields by using struct we will unit them.
    - Anonymous struct

12. Method - A method is a function which has special receiver type between func keyword and method name. This receiver type may be of struct type or non-struct type. Receiver Type can be of Value receiver or Pointer receiver.

11. Interface - An Interface is a type in go which is a collection of method signature. These 
collection of method signature are meant to represent certain behaviour. The interface declares only method set and any type implements all the methods of interface is said to be interface type.
   - Empty interface - An interface that has zero methods is called an empty interface. It is represented as interface{}. 
   - Type Assertion - used to extract the underlying value of the interface. Syntax: t.(type) 
   - Type Switch - A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. Syntax : t.(Type)

Concurrency : 
1. Difference between Concurrency and Parallism
2. Goroutines - A goroutines are like functions or methods that run concurrenctly with other functions and methods. 
Goroutines are lighweighted threads. 
The creation of goroutines are tiny/cheaper as compare to threads. 
To create goroutines in golang we use `go` keyword to the prefix of function or methods. 
3. Channels - Channels can be thought as a pipes by using which goroutines communications.
   - close() 
   - for range loop
   - Done/quit channel (bool/struct{})
4. Wait Group - A wait group is used to wait for collection of goroutions to finish their execution. The control block until all the goroutines completes their execution.
5. Mutex - A mutex is locking mechanism which ensure only one goroutine can access critical section of code at any point of time. This is done to prevent race condition from happening.
6. Select - The select statement is used to choose a multiple send/receive channel operations.
   Select statement block the control until anyone of the operation is ready. If both the operations are ready then it picks randomly.
   for-select loop - We need to stop the for-select loop by using done/quit channel.

Error Handling : 
defer - A defer used to delay the execution of a function/statement until the surrounding function completes. deferred function calls are executed in Last-In-First-Out (LIFO) order.
Panic -  a panic is a mechanism that allows you to halt the normal execution of a program when an unexpected or unrecoverable situation occurs.
Recover - recover() is a built-in function in Go that is used to regain control of a panicking goroutine.

OOPs in golang : 
1.Constructor - A constructor is a special function is used to initalize and create an instance of a struct type. To declare constructor in golang we use prefix "New" before the struct type.
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
  