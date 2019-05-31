# Lesson 2 - Go variable and constant
*A variable is a storage unit of a particular type which has a name, storage address and value and constant is a variable with fixed value.*
## Declaring a variable
To declare a variable you have to follow the above syntax:

    var variabeName dataType= value
    Example:
    var myVar int = 7
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