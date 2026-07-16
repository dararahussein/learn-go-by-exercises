# Set 03 - Errors and interfaces

Errors make failure visible in signatures; interfaces let callers depend on
behavior rather than concrete implementations. This set adds production error
wrapping with `%w`, `errors.Is`, and a sentinel error.

Work in order: `01_errors`, `02_interfaces`, `03_application`,
`04_debugging`. Run `go test` inside the current directory.

Warm-up from memory: write a pointer-receiver method and explain why a value
receiver would not mutate the caller.

`CartTotal` is the application task. Its test provides a fake implementation of
an interface defined by the consuming code, the shape used throughout real Go
tests. Finish with the shadowed-variable debugging exercise.

Optional stretch: define a custom error type carrying an SKU and implement
`Unwrap` so `errors.Is` still reaches `ErrNotFound`.

Expected time: 4-6 hours.
