# Go Defer Statement Behavior

This example demonstrates a common misunderstanding regarding Go's `defer` statement. The `defer` keyword schedules a function call to be executed when the surrounding function returns. This means that the deferred function call happens *after* any changes to variables within the surrounding function have already been made.

The `bug.go` file contains code showcasing this unexpected behavior. The `bugSolution.go` offers a solution by using a copy of the variable.

## How to reproduce the issue:

1. Save the code from `bug.go` into a file named `bug.go`.
2. Run the file using a Go compiler: `go run bug.go`

You will notice that the output shows the value of `x` as 5 in the `defer` statement, not the initial value of 10.

## Understanding the solution:

The `bugSolution.go` file demonstrates how to avoid this issue using an immediately evaluated expression to capture the initial value of the variable.