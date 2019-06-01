# Lesson 2 - Go variable and constant
*A variable is a storage unit of a particular type which has a name, storage address and value and constant is a variable with fixed value.*
## Declaring a variable
To declare a variable you have to follow the above syntax:

    //var variabeName dataType= value
    //Integer type variable:
    var myIn int = 7

    //Boolean Type variable
    var myBool bool = True

    //Floating Point Variable
    var myFloat float64 = 6.45
## Data Types
Data types is a particular kind of data item which has particular attribute, as defined by the values it can take. There are several data types in Go. These are as following

|  type | Use  | zero-value  |
|----|:------|---:|
|bool|Boolean data type. It can store value `true` or `false`.| false|
|string|String data type. It can store UTF-8 string. All strings in go are UTF-8 by default. Unlike JavaScript, string is only encapsulated in double quotes.|empty string|
|int|Integer data type. It can store 32-bit or 64-bit signed integer. A 32-bit system will allocate 32 bits of memory and 64-bit system will allocate 64 bits of memory. Hence 32-bit system can store -2147483648 to 2147483647 and 64-bit system can store -9223372036854775808 to 9223372036854775807.|0|
|uint|Integer data type. Same as `int`, `uint` can store 32 bits or 64 bits **unsigned** integer.|0|
|int8|Integer data type. System will allocate 8 bits of memory to store an integer. Hence it can store values between -128 to 127.|0|
|uint8|Integer data type. Same as `int8`, `uint8` can store 8-bit **unsigned** integer. Hence it can store values between 0 to 255.|0|
|int16|Integer data type. System will allocate 16 bits of memory to store an integer. Hence it can store values between -32768 to 32767.|0|
|uint16|Integer data type. Same as `int8`, `uint8` can store 16-bit **unsigned** integer. Hence it can store values between 0 to 65535.|0|
|int32|Integer data type. System will allocate 32 bits of memory to store an integer. Hence it can store values between -2147483648 to 2147483647.|0|
|uint32|Integer data type. Same as `int8`, `uint8` can store 32-bit **unsigned** integer. Hence it can store values between 0 to 4294967295.|0|
|int64|Integer data type. System will allocate 64 bits of memory to store an integer. Hence it can store values between -9223372036854775808 to 922337203685477580.|0|
|uint64|Integer data type. Same as `int8`, `uint8` can store 64-bit **unsigned** integer. Hence it can store values between 0 to 18446744073709551615.|0|
|float32|Float data type. System will allocate 32 bits of memory to store a float value. Hence it can store values between -3.4E+38 to +3.4E+38.|0|
|float64|Float data type. System will allocate 64 bits of memory to store a float value. Hence it can store values between -1.7E+308 to +1.7E+308.|0|
|complex64|Go supports complex numbers out of this box. `complex64` has `float32` real part and `float32` imaginary part.|0+0i|
|complex128|Similar to `complex64`, `complex128` has `float64` real part and `float64` imaginary part.|0+0i|
|byte|Alias for `uint8`.|0|
|rune|Alias for `int32`. It represents a Unicode code point.|0|
|error|Error data type. It is be used to store error value.|nil|

# Pointer And variable Address
variable address is the location where a variable is stored.

    //to see a variable address just you need to put amparsand & sign infront of variable name
    var myVar int = 7
    fmt.Println(&myVar)

A pointer is a variable which holds the address of variable.

    //Pointer variable 
    var myVar int = 9
    var p = &myVar

    fmt.Println(p)

Or pointer is a way to grab value from a variable address

    fmt.Println(*p)

# Testing Our variable in Practical way

    package server

    import "net/http"

    func Serve(){
        http.HandleFunc("/", "Index_Handler")
        http.ListenAndServe(":8000", nil)
    }

    func Index_Handler(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, `<h1>LTW Server is working</h1>
        <h2>Lesson 2 Variables and Constant</h2>
        <p>%s</p>`, myVar)
    }

## variable's Initial Value
golang's initial value is optional
you can declare a variable without initial value
    var myVar int
    //leter you can assign value to that variable
    myVar = 6

## zero value
Unlike other programming languages where variables hold null or undefined values when not initialized with any value, Go gives them zero-value of their respective types. As from the above table, boolean gets false value and integer gets 0 value.

    var myInteger int
    fmt.Println(myInteger)
    var myVar bool
    fmt.Println(bool)

## Variable name convention
Go recommends writing variable names in simple word or camelCase. Even under_score variable names are valid, they are not idiomatic.

## Type inference
    var myVar= 5
    // golang will automatically picup the type.

## Short-hand notation
    //variableName := initialValue
    myVar :=5

## Multiple variable declarations
    var var1, var2, var3 int

    //you can assaign with initial value
    var var1, var2, var3 int = 1, 2, 3

    //you can declare also multiple type
    var var1, var2, var3 = 1, 2.2, false

    // you also can declare in multiple line
    var(
        var1 = 50
        var2 = 25.22
        var3 string = "Telefonía"
        var4 bool
    )

## Type conversion
    var1 := 10 // int
    var2 := 10.5 // float
    // illegal 
    // var3 := var1 + var2
    // legal
    var3 := var1 + int(var2) // var3 == 20

## Type aliasing
    //type aliasName aliasTo
    type float float64

Example:

    package server

    import "fmt"

    type float float64

    func main() {
	    var f float = 52.2
	    fmt.Printf("f has value %v and type %T", f, f)
    }

## constant
Constant is a variable in Go with a fixed value. Any attempt to change the value of a constant will cause a run-time panic. Constant must be declared with const keyword and an initial value.

    //const const_name [data_type] = fixed_value
    const myConstant = 70

## Multiple constants declaration
    //const const_1, const_2 = value_1, value_2
    const const_1, const_2 = 80, 78

    //const(
	    a = 1 // a == 1
	    b = 2 // b == 2
	    c     // c == 2
	    d     // d == 2
    )
## iota
Go provides a keyword iota that can be used when declaring enumerated constants. This keyword yields an incremented value by 1 starting from 0, each time it is used.

    const(
        a = iota // a == 0
        b = iota // b == 1
        c = iota // c == 2
        d        // d == 3 (implicitely d = iota)
    )
## Numeric Expressions
    var a = 11/2 // a == 5
    var a = 11.0/2 // a == 5.5
    var b = float32(11)/2 // a == 5.5