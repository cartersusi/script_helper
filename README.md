# Go Script Helper

- [Go Script Helper](#go-script-helper)
  - [Installation](#installation)
  - [Functions](#functions)
    - [1. GetFlag](#1-getflag)
    - [2. Error](#2-error)
    - [3. Success](#3-success)
    - [4. Warning](#4-warning)
  - [Color Utility Functions](#color-utility-functions)
  - [Example Usage](#example-usage)
  - [License](#license)


A Go package providing utility functions for flag handling and formatted logging. This package is designed to help streamline script development by providing structured error, success, and warning logging, as well as a flag validation function.

## Installation

To install the `script_helper` package, use:

```bash
go get path/to/script_helper
```

## Functions
### 1. GetFlag
```go
func GetFlag[T string | int | bool](long_flag, short_flag *T, value T) T
```
The GetFlag function checks the values of the specified long and short flags. It enforces that only one flag should be set, returning the value of the flag that is set.

* Parameters:
    * long_flag: Pointer to the long-form of the flag.
    * short_flag: Pointer to the short-form of the flag.
    * value: Default value for the flags.
* Returns: The value of the specified flag (either long_flag or short_flag).
* Errors: Triggers an error if neither flag is set or if both flags are set.

### 2. Error
```go
func Error(s string, exit ...bool)
```
Logs an error message with an "ERROR" prefix and optionally exits the program. Supports colored output to enhance readability.

* Parameters:
* s: The error message to be logged.
* exit: Optional boolean slice. If the first value is true, the program exits after logging the error.

* Behavior:
    * Logs the error message in red.
    * Exits the program if exit is true.
    * Continues execution if exit is not provided or is false.

### 3. Success
```go
func Success(s string)
```
Logs a success message with a "SUCCESS" prefix, with the message displayed in green.

* Parameters:
    * s: The success message to be logged.
* Behavior:
    * Logs the message in green with a "SUCCESS" prefix.

### 4. Warning
```go
func Warning(s string)
```
Logs a warning message with a "WARNING" prefix, with the message displayed in yellow.

* Parameters:
    * s: The warning message to be logged.
Behavior:
    * Logs the message in yellow with a "WARNING" prefix.

## Color Utility Functions

The package also includes internal functions for color formatting of text:
* `red(s string)`: Formats text in red.
* `green(s string)`: Formats text in green.
* `yellow(s string)`: Formats text in yellow.

## Example Usage
```go
package main

import (
	"script_helper"
	"flag"
)

func main() {
	longFlag := flag.String("long", "", "long form flag")
	shortFlag := flag.String("short", "", "short form flag")
	defaultValue := ""

	flag.Parse()

	// Retrieve the flag value using GetFlag
	selectedFlag := script_helper.GetFlag(longFlag, shortFlag, defaultValue)

	// Log success if a flag is correctly provided
	script_helper.Success(fmt.Sprintf("Selected flag: %s", selectedFlag))

	// Log a warning if something unusual happens
	script_helper.Warning("This is a warning message.")

	// Log an error and exit
	script_helper.Error("This is an error message.", true)
}
```

## License
This package is open-source and available under the MIT License.