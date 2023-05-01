# Coursera Notes - Getting Started With Go


## Week 2
### 2.1.1 Pointers
```go
var x int = 1
var y int
var ip *int  // ip is pointer to int
ip = &x // ip points to x
y = *ip

```



### 2.1.2 Variable Scope

**Blocks**



### 2.1.3 Deallocating Memory

When a variable is no longer needed, it should be deallocated.

**Stack & Heap**

The stack is dedicated to function calls.

- Local variables are stored here.

Heap is persistent.

### 2.1.4 Garbage Collection

- In interpreted languages, this is done by the interpreter.

- Go is a compiled language that enables garbage collection. The compiler determines stack vs heap;

### 2.2.1 Comments, Printing, Integers

**Comments**

```go
var x int // single-line comment

/* Comment 1
	 Comment 2 */
var y int
```

**Printing**

```go
fmt.Printf("Hi")

// with arguments
var x int = 1
fmt.Printf("hi, %v", x)
```

**Integers**

- int8, int16, int32, int64...
- uint8, uint16...



###  2.2.2 Ints, Floats, Strings

**Type Conversions**

- T() operation

```go
var x int32 = 1
var y int16 = 2
x = int32(y)	
```

**Floating Point**

- Float32 - ~6 digits of precision
- Float64 - ~15 digits of precision

**rune**

**rune**

代码点。



### 2.2.3  String Packages

- Strconv Package

  ```go
  Atoi(s) // converts string s to int
  ```

  

### 2.3.1 Constants

```go
const x = 1

const (
  y = 2
  s = "Hi"
)
```

- **iota**
  - Generate a set of related but distinct constants.

### 2.3.2 Control Flow

- if

```go
if x == 1 {
	fmt.Printf("x = 1")
}
```

- for loop

```go
for i := 1; i < 10; i++ {
	fmt.Printf("x = 1")
}

i := 0
for i < 10 {
	fmt.Printf("x = 1")
  i++
}

// equals to while loop
for {
	fmt.Printf("x = 1")
  i++
}
```

- Switch/Case

```go
switch x {
	case 1:
		fmt.Printf("x = 1")
	case 2:
		fmt.Printf("x = 2")
	default:
		fmt.Printf("x = default")
}
```

- **Tagless Switch**

```
switch {
	case x > 1:
		fmt.Printf("x = 1")
	case x < -2:
		fmt.Printf("x = 2")
	default:
		fmt.Printf("x = default")
}
```

- Scan
  - Reads user input.
  - Takes a pointer as an argument
  - Typed data is written to the pointer.

```go
	var appleNum int
	fmt.Printf("Number：")
	_, _ = fmt.Scan(&appleNum)
	fmt.Printf("Number is : %v\n", appleNum)
```





