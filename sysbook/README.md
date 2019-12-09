# 《Hands-On System Programming with Go》

## ch1

- The **protection ring**, also referred to as **hierarchical protection domains**, is the mechanism used to protect a system against failure. Its name is derived from the hierarchical structure of its levels of permission, represented by concentric rings, with privilege decreasing when moving to the outside rings. Between each ring there are special gates that allow the outer ring to access the inner ring resources in a restricted manner.  

- The number and order of rings depend on the CPU architecture. They are usually numbered with decreasing privilege, making ring 0 the most privileged one. This is true for i386 and x64 architecture that use four rings(from ring 0 to ring 3) but it's not true for ARM, which uses reverse order(from EL3 to EL0). Most operating systems are not using all four levels; they end up using a two level hierarchy——user/application(ring 3) and kernel(ring 0).  

- A software that runs under an operating system will be executed at user (ring 3) level. In order to access the machine resources, it will have to interact with the operating system kernel (that runs at ring 0). Here's a list of some of the operations a ring 3 application cannot do:  
  - Modify the current segment descriptor, which determines the current ring
  - Modify the page tables, preventing one process from seeing the memory of other processes  
  - Use the LGDT and LIDT instructions, preventing them from registering interrupt handlers  
  - Use I/O instructions such as in and out that would ignore file permissions and read directly from disk  
  
- Diving into system calls  
  - System calls are the way operating systems provide access to the resources for the applications. It is an API implemented by the kernel for accessing the hardware safely.  
  - There are some categories that we can use to split the numerous functions offered by the operating system. These include the control of the running applications and their flow, the filesystem access, and the network.
  
  - Process control  
    - This type of services include `load`, which adds a program to memory and prepares for its execution before passing control to the program itself, or `execute`, which runs an executable file in the context of a pre-existing process. Other operations that belong to this category are as follows:  
      - `end` and `abort` —— the first requires the application to exit while the second forces it.  
      - `CreateProcess`, also known as `fork` on Unix systems or `NtCreateProcess` in Windows.  
      - Terminate process.  
      - Get/set process attributes.  
      - Wait for time, wait event, or signal event.  
      - Allocate and free memory.  
  - File management  
    - The handling of files and filesystems belongs to file management system calls. There are *create* and *delete* files that make it possible to add or remove an entry from the filesystem, and `open` and `close` operations that make it possible to gain control of a file in order to execute read and write operations. It is also possible to read and change file attributes.
  - Device management  
    - Device management handles all other devices but the filesystem, such as frame buffers or display. It includes all operations from the request of a device, including the communication to and from it (read, write, seek), and its release. It also includes all the operations of changing device attributes and logically attaching and detaching them.
  - Information maintenance  
    - Reading and Writing the system date and time belongs to the information maintenance category. This category also takes care of other system data, such as the environment. Another important set of operations that belongs here is the request and the manipulation of processes, files, and device attributes.
  - Communication  
    - All the network operations from handling sockets to accepting connections fall into the communication category. This includes the creation, deletion, and naming of connections, and sending and receiving messages.  
  - The difference between operating systems  
    - Windows has a series of different system calls that cover all the kernel operations. Many of these correspond exactly with the Unix equivalent.

- POSIX standards and features
  - **Portable Operating System Interface**(**POSIX**) for Unix represents a series of standards for operating system interfaces.
  - There are many features defined by POSIX, and they are organized in four different standards, each one focusing on a different
  aspect of the Unix compliance. They are all named POSIX followed by a number.  
    - POSIX.1 —— core services
    - POSIX.1b and POSIX.1c —— real-time and thread extensions
    - POSIX.2 —— shell and utilities

- Summary
  - In this chapter, we saw what system programming means —— writing system software that has some strict requirements, such as being tied to
  the hardware, using a low-level language, and working in a resource-constrained environment. Its practices can be really useful when building distributed systems that normally require optimizing resource usage. We discussed APIs, definitions that allows software to be used by other software, and listed the different types —— the ones in the operating system, libraries and frameworks, and remote and web APIs.
  - We analyzed how, in operating systems, the access to resources is arranged in hierarchical levels called **protection rings** that prevent
  uncontrolled usage in order to improve security and avoid failures from the applications. The Linux model simplifies this hierarchy to just two levels called *user* and *kernel* space. All the applications are running in the user space, and in order to access the machine's resources they need the kernel to intercede.
  - Then we saw one specific type of API called **system calls** that allows the applications to request resources to the kernel, and mediates
  process control, access and management of files, and devices and network communications.

## ch2

- Memory management
  - There are different techniques for handling memory, including the following:
    - **Single allocation**: All the memory, besides the part reserved for the OS, is available for the application. This means that there can only be one application in execution at a time, like in **Microsoft Disk Operating System** (**MS-DOS**).
    - **Partitioned allocation**: This divides the memory into different blocks called partitions. Using one of these blocks per process makes it possible to execute more than one process at once. The partitions can be relocated and compacted in order to obtain more contiguous memory space for the next processes.
    - **Paged memory**: The memory is divided into parts called frames,which have a fixed size. A process' memory is divided into parts of the same size called **pages**. There is a mapping between pages and frames that makes the process see its own virtual memory as contiguous. This process is also known as **pagination**.
    - Virtual memory
      - Unix uses the paged memory management technique,abstracting its memory for each application into contiguous virtual memory. It also uses a technique called swapping, which extends the virtual memory to the secondary memory (hard drive or **solid state drives** (**SSD**)) using a swap file.
      - When memory is scarce, the operating system puts pages from processes that are sleeping in the swap partition in order to make space for active processes that are requesting more memory, executing an operation called **swap-out**. When a page that is in the swap file is needed by a process in execution it gets loaded back into the main memory for executing it. This is called **swap-in**.
      - Creating memory-efficient applications is a way of increasing performance by avoiding or reducing swapping.
      - The `top` command shows details about available memory, swap, and memory consumption for each process.

- files and filesystems
  - A filesystem is a method used to structure data in a disk, and a file is the abstraction used for indicating a piece of self-contained information.
  - When more than one file points to the same content, we have a **hard link**, but this is not allowed in all filesystems (for example, NTFS and FAT). A **soft link** is a file that points to another soft link or to a hard link. Hard links can be removed or deleted without breaking the original link, but this is not true for soft links. A **symbolic link** is a regular file with its own data that is the path of another file. It can also link other filesystems or files and directories that do not exist (that will be a broken link).
  - In Unix, some resources that are not actually files are represented as files, and communication with these resources is achieved by writing to or reading from their corresponding files. For instance, the `/dev/sda` file represents an entire disk, while `/dev/stdout`, `/dev/stdin`, and `/dev/stderr` are standard output, input, and error. The main advantage of *Everything is a file* is that the same tools that can be used for files can also interact with other devices (network and pipes) or entities (processes).

- Unix filesystem
  - In Linux and macOS, each file and directory is represented by an **inode**, which is a special data structure that stores all the information about the file except its name and its actual data. Inode `0` is used for a null value, which means that there is no inode. Inode `1` is used to record any bad block on the disk. the root of the hierarchical structure of the filesystem uses inode `2`. It is represented by `/`.

- Processes
  - When an application is launched, it becomes a process: a special instance provided by the operating system that includes all the resources that are used by the running application. This program must be in **Executable and Linkable Format** (**ELF**), in order to allow the operating system to interpret its instructions.
  - Each process is a five-digit identifier **process ID** (**PID**), and it represents the process for all its life cycle. To see a list of the active processes, there's the `ps` (process status) command, which shows the current list of running processes for the active user.
  - When a process is launched, it is normally in the **foreground**, which will prevent communication with the shell until the job is finished or interrupted. Launching a process with an `&` symbol at the end of the command (`cat file.txt &`) will launch it in the **background**, making it possible to keep using the shell. The `SIGTSTP` signal can be send with *Ctrl+Z*, which allows the user to suspend the foreground process from the shell. It can be resumed with the `fg` command, or in the background with the `bg` command. The `jobs` command reports the jobs running and their numbers.
  - The foreground process can be terminated with the `SIGINT` signal using *Ctrl+Z*. In order to kill a background process, or send any signal to a process, the `kill` command can be used. The `kill` command receives an argument that can be either of the following:
    - The signal send to the process
    - The PID or the job number (with a `%` prefix)
  - The more notable signals used are as follows:
    - `SIGINT`: Indicates a termination caused by user input and can be sent by `kill` command with the `-2` value
    - `SIGTERM`: Represents a general purpose termination request not generated by a user as well as a default signal for the `kill` command with a `-6` value
    - `SIGKILL`: A termination handled directly by the operating system that kills the process immediately and has a `-9` value

- Users, groups, and permissions
  - Authorization to files and other resources are provided by users and groups. Users have unique usernames that are human-friendly identifiers, but from the operating system side, each user is represent by a unique positive integer: the **User ID** (**UID**). Groups are the other authorization mechanism and, as users, they have a name and a **Group ID** (**GID**). In the operating system, each process is associated with a user and each file and directory belongs to both a user and a group.
  - The `/etc/passwd` file contains all this information and more.
  - A group is a logical collection of users, used to share files and directories between them. Each group is independent of other groups and there is no specific relationship between them. For a list of the groups that the current user belongs to, there's the `groups` command. To change group ownership of a file, there's `chgrp`.
  - The `chmod` command makes it possible to change permission on a file or directory. This can be used to override current permissions or to modify them.

- Process communications
  - The operating system is responsible for communication between processes and has different mechanisms to exchange information. These processes are unidirectional, such as exit codes, signals, and pipes, or bidirectional, such as sockets.
  - Applications communicate their result to the operating system by returning a value called **exit status**. This is an integer value passed to the parent process when the process ends.  
  - **Exit codes**: The exit code of the last command is stored in the `$?` variable, and it can be tested in order to control the flow of the operations. A commonly used operator is `&&` (double ampersand), which executes the next instruction only if the exit code of the first one is `0`, such as `stat file && echo something >> file`, which appends something to a file only if it exists.
  - **Signals**: Exit codes connect processes and their parents, but signals make it possible to interface any process with another, including itself. They are also asynchronous and unidirectional, but they represent communication from the outside of a process. The `kill` command allows you to send a signal to any application, and a comprehensive list of available signals can be shown with the `-l` flag.
  - **Pipes**: Pipes are the last unidirectional communication method between processes. As the name suggests, pipes connect two ends —— a process input with another process output —— making it possible to process on the same host to communicate in order to exchange data. These are classified as anonymous or named:
    > - Anonymous pipes link one process standard output to another process standard input. It can be easily done inside a shell with the `｜` operator, linking the output of the command before the pipes as input for the one after the pipe. `ls -l | grep "user` gets the output of the `ls` command and uses it as input for `grep`.  
    > - Named pipes use a specific file in order to execute the redirect. The output can be redirected to a file with the `>` (greater) operator, while the `<` (less) sign allows you to use a file as input for another process. `ls -l > file.txt` saves the output of the command to a file. `cat < file.txt` sends the contents of the file to the command's standard input, and the standard input copies them to the standard output.
    > - It is also possible to append content to a named pipe using the `>>` (double greater) operator, which will start writing from the end of the file.
  - **Sockets**: Unix domain sockets are a bidirectional communication method between applications on the same machine. They are a logical endpoint that is handled by the kernel and manages the data exchange. The nature of sockets permits using them as stream-oriented, or datagram-oriented. Stream-oriented protocols ensure that messages are delivered before moving to the next chunk of data in order to preserve message integrity. In contrast, message-oriented protocols ignore the data that is not received and keeps sending the following messages, making it a faster but less reliable protocol with very low latency. The sockets are classified as follows:
    > - `SOCK_STREAM`: Connection-oriented, ordered, and reliable transmission of a stream of data
    > - `SOCK_SEQPACKET`: Connection-oriented, ordered, and reliable transmission of message data that has record boundaries
    > - `SOCK_DGRAM`: Unordered and unreliable transmission of messages.

## ch3

- This chapter will provide an overview of the Go language and its basic functionality.
- Technical requirements: From this chapter onward, you will need Go installed on your machine.
- The strengths of the language are countless.
- It's important to avoid having circular dependencies because they will not compile. Since circular dependencies are not allowed, the packages need to be designed differently to other languages. In order to break the circular dependency, it is good practice to export functionalities from a package or replace the dependency with an interface.
- Go reduces all the symbol visibility to a binary model —— exported and not exported —— unlike many other languages, which have intermediate levels. For each package, all the symbols starting with a capital letter are exported, while everything else is used only inside the package.
- Ignored imports are used to import packages without you having to use them. This makes it possible to execute the `init` function of the package without referencing the package in your code.
- The main uses of custom types are to define methods and to make the type specific for a scope, like defining a `string` type called `Message`.
- Variables represent mapping to the content of a portion of contiguous memory. They have a type that defines how much this memory extends, and a value that specifies what's in the memory. Type can be basic, composite, or custom, and its value can be initialized with their zero-value by a declaration, or with another value by assignment.
- `+=`, `-=`, `*=`, and `/=` execute the operation before the equals sign between what's before and what's after the operator and assigns it to the variable on the left. These four operations produce a value of the same type of the variables involved.
- Some operations are exclusive for integers and produce other integers: `%`, `&`, `|`, `^`, `&^`, `<<`, and `>>`.
- For all non-pointer variables, it is also possible to use `&`, the reference operator, to obtain the variable address that can be assigned to a pointer variable. The `*` operator makes it possible to execute a dereference operation on a pointer and obtain the value of the variable indicated by it:

**Operator**|**Name**|**Description**|**Example**  
:--|:--|:--|:--  
`=`|Assignment|Assigns the value to a variable|`a = 10`
`:=`|Declaration and assignment|Declares a variables and assigns a value to it|`a:= 10`
`==`|Equals|Compares two variables, returns a Boolean if they are the same|`a == b`
`!=`|Not Equals|Compares two variables, returns a Boolean if they are different|`a != b`
`+`|Plus|Sum between the same numerical type|`a + b`
`-`|Minus|Difference between the same numerical type|`a - b`
`*`|Times|Multiplication between the same numerical type|`a * b`
`/`|Divided|Division between the same numeric type|`a / b`
`%`|Modulo|Remainder after division of the same numerical type|`a % b`
`&`|AND|Bit-wise AND|`a & b`
`&^`|Bit clear|Bit clear|`a &^ b`
`<<`|Left shift|Bit shift to the left|`a << b`
`>>`|Right shift|Bit shift to the right|`a >> b`
`&&`|AND|Boolean AND|`a && b`
`||`|OR|Boolean OR|`a || b`
`!`|NOT|Boolean NOT|`!a`
`<-`|Receive|Receive from a channel|`<-a`
`->`|Send|Send to a channel|`a <- b`
`&`|Reference|Returns the pointer to a variable|`&a`
`*`|Dereference|Returns the content of a pointer|`*a`  

- **Casting**: Converting a type into another type is an operation called **casting**, which works slightly differently for interfaces and concrete types. There's a special type of conditional operator for casting called **type switch** which allows an application to attempt multiple casts at once.
- **Scope**: Two variables in the same scope cannot have the same name, but a variable of an inner scope can reuse an identifier. When this happens, the outer variable is not visible in the inner scope —— this is called **shadowing**, and it needs to be kept in mind in order to avoid issues that are hard to identify.
- Functions in Go are identified by the `func` keyword, followed by an identifier, eventual arguments, and return values. Functions in Go can return more than one value at a time. The combination of arguments and returned types is referred to as a **signature**.
  - The part between brackets is the function body, and the `return` statement can be used inside it for an early interruption of the function. If the function returns values, then the return statement must return values of the same type.
  - The `return` values can be named in the signature; they are zero value variables and if the `return` statement does not specify other values, these values are the ones that are returned.
  - Functions are first-class types in Go and they can also be assigned to variables, with each signature representing a different type. They can also be anonymous; in this case they are called **closures**. Once a variable is initialized with a function, the same variable can be reassigned with another function with the same signature. Here's an example of assigning a closure to a variable:

```go
var a = func(item string) error {
  if item != "elixir" {
    return errors.New("Gimme elixir!")
  }
  return nil
}
```

- The functions that are declared by an interface are referred to as methods and they can be implemented by custom types. The method implementation looks like a function, with the exception being that the name is preceded by a single parameter of the implementing type. This is just syntactic sugar —— the method definition creates a function under the hood, which takes an extra parameter, that is, the type that implements the method.
- This syntax makes it possible to define the same method for different types, each of which will act as a namespace for the function declaration. In this way, it is possible to call a method in two different ways, as shown in the following code:

```go
type A int

func (a A) Foo() {}

func main() {
  A{}.Foo()     // Call the method on an instance of the type
  A.Foo(A{})    // Call the method on the type and passing an instance as argument
}
```

- As well as a manual call to the `panic` function, there is a set of operations that will cause a panic, including the following:
  > - Access a negative or non-existent array/slice index(index out of range)
  > - Dividing an integer by `0`
  > - Sending to a closed channel  
  > - Dereferencing on a `nil` pointer(`nil` pointer)
  > - Using a recursive function call that fills the stack(stack overflow)

- Panic should be used for errors that are not recoverable, which is why errors are just values in Go. Recovering a panic should be just an attempt to do something with that error before exiting the application. If an unexpected problem occurs, it's because it hasn't been handled correctly or some checks are missing. This represents a serious issue that needs to be dealt with, and the program need to change, which is why it should be intercepted and dismissed.  
- Concurrency is so central to Go that two of its fundamental tools are just keywords —— `chan` and `go`. This is a very clever way of hiding the complexity of a well-designed and implemented concurrency model that is easy to use and understand.
- A channel is made for sharing data, and it usually connects two or more execution threads in an application, which makes it possible to send and receive data without worrying about data safety, Go has a lightweight implementation of a thread that is managed by the runtime instead of the operating system, and the best way to make them communicate is through the use of channels.
- Creating a new goroutine is pretty easy —— you just need to use the `go` operator, followed by a function execution. This includes method calls and closures. If the function has any arguments, they will be evaluated before the routine starts. Once it starts, there is no guarantee that changes to variables from an outer scope will be synchronized if you don't use channels.
- **Stack and heap**
  - Memory is arranged into two main areas —— stack and heap. There is a stack for the application entry point function(`main`), and additional stacks are created with each goroutine, which are stored in the heap. The **stack** is, as its name suggests, a memory portion that grows with each function call, and shrinks when the function returns. The **heap** is made of a series of regions of memory that are dynamically allocated, and their lifetime is not defined a priori as the items in the stack,; heap space can be allocated and freed at any time.
  - All the variables that outlive the function where they are defined are stored in the heap, such as a returned pointer. The compiler uses a process called **escape analysis** to check which variables go on the heap. This can be verified with the `go tool compile -m` command.
  - Variables in the stack come and go with the function's execution. Let's take a look at a practical example of how the stack works:

  ```go
  func main() {
    var a, b = 0, 1
    f1(a, b)
    f2(a)
  }

  func f1(a, b int) {
    c := a + b
    f2(c)
  }

  func f2(c int) {
    print(c)
  }
  ```

  We have the `main` function calling a function called `f1`, which calls another function called `f2`. Then, the same function is called directly by `main`.

  When the `main` function starts, the stack grows with the variables that are being used. In memory, this would look something like the following table, where each column represents the pseudo state of the stack, which it represents how the stack changes in time, going from left to right:

  `main` invoked|`f1` invoked|`f2` invoked|`f2` return|`f1` returns|`f2` invoked|`f2` returns|`main` returns|
  :-|:-|:-|:-|:-|:-|:-|:-
  `main()`|`main()`|`main()`|`main()`|`main()`|`main()`|`main()`|//empty
  `a = 0`|`a = 0`|`a = 0`|`a = 0`|`a = 0`|`a = 0`|`a = 0`|
  `b = 1`|`b = 1`|`b = 1`|`b = 1`|`b = 1`|`b = 1`|`b = 1`|
  ||`f1()`|`f1()`|`f1()`| |`f2()`| |
  ||`a = 0`|`a = 0`|`a = 0`| |`c = 0`| |
  ||`b = 1`|`b = 1`|`b = 1`||||
  ||`c = 1`|`c = 1`|`c = 1`||||
  |||`f2()`|||||
  |||`c = 1`|||||

  When `f1` gets called, the stack grows again by copying the `a` and `b` variables in the new part and adding the new variable, `c`. The same thing happens for `f2`. When `f2` returns, the stack shrinks by getting rid of the function and its variables, which is what happens when `f1` finishes. When `f2` is called directly, it grows again by recycling the same memory part that was used for `f1`.

  The garbage collector is responsible for cleaning up the unreferenced values in the heap, so avoiding storing data in it is a good way of lowering the work of the **garbage collector**(**GC**), which causes a slight decrease in performance in the app when the GC is running.

  The GC is responsible for freeing the areas of the heap that are not referenced in any stack.

## ch4

- Go offers a series of functions that make it possible to manipulate file paths that are paltform-independent and that are contained mainly in the `path/filepath` and `os` packages.
- An example of the list and count files is shown in the following code:

```go
func main() {
  if len(os.Args) != 2 { // ensure path is specified
    fmt.Println("Please specify a path.")
    return
  }
  root, err := filepath.Abs(os.Args[1]) // get absolute path
  if err != nil {
    fmt.Println("cannot get absolute path:", err)
    return
  }
  fmt.Println("Listing files in", root)

  var c struct {
    files int
    dirs  int
  }
  
  filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    // walk the tree to count files and folders
    if info.IsDir() {
      c.dirs++
    } else {
      c.files++
    }
  })
  fmt.Println("-", path)
  fmt.Printf("Total: %d files in %d directories", c.files, c.dirs)
}
```

- Getting the contents of a file can be done with an auxiliary function in the `io/ioutil` package, as well as with the `ReadFile` function, which opens, reads, and closes the file at once. This uses a small buffer (512 bytes) and loads the whole content in memory. This is not a good idea if the file size is very large, unknown, or if the content of the file can be processed one part at a time.
- An example of reading all the content at once is shown in the following code:

```go
func main() {
  if len(os.Args) != 2 {
    fmt.Println("Please specify a path.")
    return
  }
  b, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Println(string(b))
}
```

- For all operations that read from a disk, there's an interface that is paramount:
  - A reader makes it possible to process data in chunks (the size is determined bu the slice), and if the same slice is reused for the operations that follow, the resulting program is consistently more memory efficient because it is using the same limited part of the memory that allocates the slice.

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

- A **data buffer**, or just a buffer, is a part of memory that is used to store temporary data while it is moved. Byte buffers are implemented in the `bytes` package, and they are implemented by an underlying slice that is capable of growing every time the amount of data that needs to be stored will not fit.
- If new buffers get allocated each time, the old ones will eventually be cleaned up by the GC itself, which is not an optimal solution. It's always better to reuse buffers instead of allocating the new ones. This is because they make it possible to reset the slice while keeping the capacity as it is (the array doesn't get cleared or collected by the GC).
- A buffer also offers two functions to show its underlying length and capacity. In the following example, we can see how to reuse a buffer with `Buffer.Reset` and how to keep track of its capacity.
- An example of buffer reuse and its underlying capacity is shown in the following code:

```go
func main() {
    var b = bytes.NewBuffer(make([]byte, 26))
    var texts = []string{
        `As he came into the window`,
        `It was the sound of a crescendo
He came into her apartment`,
        `He left the bloodstains on the carpet`,
        `She ran underneath the table
He could see she was unable
So she ran into the bedroom
She was struck down, it was her doom`,
    }
    for i := range texts {
        b.Reset()
        b.WriteString(texts[i])
        fmt.Println("Length:", b.Len(), "\tCapacity:", b.Cap())
    }
}
```

- There are two other interfaces that are related to readers: `io.Closer` and `io.Seeker`:

```go
type Closer interface {
  CLose() error
}

type Seeker interface {
  Seek(offet int64, whence int) (int64, error)
}
```

These are usually combined with `io.Reader`, and the resulting interfaces are as follows:

```go
type ReadCloser interface {
  Reader
  Closer
}

type ReadSeeker interface {
  Reader
  Seeker
}
```

The `Close` method ensures that the resource gets released and avoid leaks, while the `Seek` method makes it possible to move the cursor of the current object (for example, a `Writer`) to the desired offset from the start/end of the file, or from its current position.

The `os.FIle` structure implements this method so that it satisfies all the listed interfaces. It is possible to close the file when the operations are concluded, or to move the current cursor around, depending on what you are trying to achieve.

As we have seen for reading, there are different ways to write files, each one with its own flaws and strengths. In the `ioutil` package, for instance, we have another function called `WriteFile` that allows us to execute the whole operation in one line. This includes opening the file, writing its contents, and then closing it.

An example of writing all a file's content at once is shown in the following code:

```go
func main() {
  if len(os.Args) != 3 {
    fmt.Println("Please specify a path and some content")
    return
  }
  // the second argument, the content, needs to be casted to a byte slice
  if err := ioutil.WriteFile(os.Args[1], []byte(os.Args[2]), 0644); err != nil {
    fmt.Println("Error:", err)
  }
}
```

If the size of the content isn't very big and the application is short-lived, it's not a problem if the content gets loaded in memory and written with a single operation. This isn't the best practice for long-lived applications, which are executing reads and writes to many different files. they have to allocate all the content in memory, and that memory will be released by the GC at some point —— this operation is not cost-free, which means that is has disadvantages regarding memory usage and performance.

- Write interface

The same principle that is valid for reading also applies for writing —— there's an interface in the `io` package that determines writing behaviors, as shown in the following code:

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}
```

We can use a slice of bytes as a buffer to write information piece by piece. In the following example, we will try to combine reading from the previous section with writing, using the `io.Seeker` capabilities to reverse its content before writing it.

```go
func main() {
  if len(os.Args) != 3 {
    fmt.Println("Please specify a source and a destination file")
    return
  }
  src, err := os.Open(os.Args[1])
  if err != nil {
    return
  }
  defer src.Close()
  // OpenFile allows to open a file with any permissions
  dst, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
  if err != nil {
    return
  }
  defer dst.Close()

  cur, err := src.Seek(0, io.SeekEnd) // Let's go to the end of the file
  if err != nil {
    return
  }
  b := make([]byte, 16)

  // After moving to the end of the file and defining a byte buffer, we enter a loop
  // that goes a littile backwards in the file, then reads a section of it.
  for step, r, w := int64(16), 0, 0; cur != 0; {
    if cur < step { //ensure cursor is 0 at max
      b, step = b[:cur], cur
    }
    cur = cur - step
    _, err = src.Seek(cur, io.SeekStart) // go backwards 
    if err != nil {
      break
    }
    if r, err = src.Read(b); err != nil || r != len(b) {
      if err == nil { // all buffer should be read
        err = fmt.Errorf("read: expected %d bytes, got %d", len(b), r)
      }
      break
    }
    // Then we reverse the content and write it to the destination
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
      switch {
      case b[i] == '\r' && b[i+1] == '\n':
        b[i], b[i+1] = b[i+1], b[i]
      case j != len(b)-1 && b[j-1] == 'r' && b[j] == '\n':
        b[j], b[j-1] = b[j-1], b[j]
      }
      b[i], b[j] = b[j], b[i]  // swap bytes
    }
    if w, err = dst.write(b); err != nil || w != len(b) {
      if err != nil {
        err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
      }
    }
  }
  if err != nil && err != io.EOF { // we expect an EOF
    fmt.Println("\n\nError:", err)
  }
}
```

- Buffers and format

In the previous section, we saw the `bytes.Buffer` can be used to store data temporarily and how it handles its own growth by appending the underlying slice. The `fmt` package extensively uses buffers to execute its operations; these aren't the ones in the bytes package for dependency reasons. This approach is inherent to one of the Go's proverbs:

> "A little copy is better than a little dependency."

If you have to import a package to use one function or type, you should consider just copying the necessary code into your own package. If a package contains much more than what you need, copying allows you to reduce the final size of the binary. You can also customize the code and tailor it to your needs.

Another use of buffers is to compose a message before writing it. Let's write some code so that we can use a buffer to format a list of books:

```go
const grr = "G.R.R. Martin"

type book struct {
  Author, Title string
  Year          int
}

func main() {
  dst, err := os.OpenFile("book_list.txt", os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    return
  }
  defer dst.Close()
  bookList := []book{
    {Author: grr, Title: "A Game of Thrones", Year: 1996},
    {Author: grr, Title: "A Clash of Kings", Year: 1998},
    {Author: grr, Title: "A Storm of Swords", Year: 2000},
    {Author: grr, Title: "A Feast for Crows", Year: 2005},
    {Author: grr, Title: "A Dance with Dragons", Year: 2011},
    // if year is omitted it defaulting to zero value
    {Author: grr, Title: "The Winds of Winter"},
    {Author: grr, Title: "A Dream of Spring"},
  }
  b := bytes.NewBuffer(make([]byte, 0, 16))
  for _, v := range bookList {
    // prints a msg formatted with arguments to writer
    fmt.Fprintf(b, "%s - %s", v.Title, v.Author)
    if v.Year > 0 {
      // we do not print the year if it's not there
      fmt.Fprintf(b, " (%d)", v.Year)
    }
    b.WriteRune('\n')
    if _, err := b.WriteTo(dst); err != nil { // copies bytes, drains buffer
      fmt.Println("Error:", err)
      return
    }
  }
}
```

- There is a very similar struct in the `strings` package called `Builder` that has the same write methods but some differences, such as the following:
  - The `String()` method uses the `unsafe` package to convert the bytes into a string, instead of copying them.
  - It is not permitted to copy a `strings.Builder` and then write to the copy since this causes a `panic`.

