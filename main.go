package script_helper

import (
	"errors"
	"fmt"
	"log"
)

// GetFlag checks the values of the long and short flags. If both flags are set to the default value or both
// are set to non-default values, it triggers an error. Otherwise, it returns the value of the flag that is set.
// Parameters:
//   - long_flag: Pointer to the long-form of the flag.
//   - short_flag: Pointer to the short-form of the flag.
//   - value: Default value for the flags.
//   - flag_name: Optional name of the flag to be used in the error message.s
//
// Returns:
//   - The value of the specified flag (long_flag or short_flag).
func GetFlag[T string | int | bool](long_flag, short_flag *T, value T, flag_name ...string) T {
	flagstr := "flags"
	if len(flag_name) > 0 {
		flagstr = flag_name[0]
	}

	if *long_flag == value && *short_flag == value {
		Error(fmt.Sprintf("Missing values for %s. Please provide one.", flagstr), true)
	}

	if *long_flag != value && *short_flag != value {
		Error(fmt.Sprintf("Both %s are set. Please provide only one.", flagstr), true)
	}

	if *long_flag != value {
		return *long_flag
	}

	return *short_flag
}

// Error logs an error message and optionally exits the program.
// Parameters:
//   - s: The error message to be logged.
//   - exit: An optional boolean slice. If the first value is true, the program exits after logging the error.
//
// Behavior:
//   - Logs the error message with an "ERROR" prefix in red.
//   - If `exit` is provided and true, the program terminates with a log.Fatal call.
//   - Otherwise, the error is logged and the program continues execution.
func Error(s string, exit ...bool) {
	err := errors.New(fmt.Sprintf("ERROR: %s\n", red(s)))

	if len(exit) == 0 {
		log.Default().Output(2, err.Error())
	}

	if len(exit) > 0 && exit[0] {
		log.Fatalf("ERROR: %s\n", err.Error())
	}

	log.Default().Output(2, err.Error())
}

// Success logs a success message with a "SUCCESS" prefix.
// Parameters:
//   - s: The success message to be logged.
//
// Behavior:
//   - Logs the message in green with a "SUCCESS" prefix.
func Success(s string) {
	msg := fmt.Sprintf("SUCCESS: %s\n", green(s))
	log.Default().Output(2, msg)
}

// Warning logs a warning message with a "WARNING" prefix.
// Parameters:
//   - s: The warning message to be logged.
//
// Behavior:
//   - Logs the message in yellow with a "WARNING" prefix.
func Warning(s string) {
	msg := fmt.Sprintf("WARNING: %s\n", yellow(s))
	log.Default().Output(2, msg)
}

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func green(s string) string {
	return "\033[32m" + s + "\033[0m"
}

func yellow(s string) string {
	return "\033[33m" + s + "\033[0m"
}
