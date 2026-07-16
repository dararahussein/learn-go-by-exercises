# Set 05 - JSON and files

This set moves data across process boundaries. Reuse earlier structs, pointers,
slices, maps, and wrapped errors while learning JSON tags, file permissions,
`defer`, and resource cleanup.

Work in order: `01_json`, `02_files`, `03_config`, `04_debugging`. Enter the
current directory and run `check` for focused feedback.

Warm-up from memory: sketch `var p Person` followed by the call that lets
`json.Unmarshal` modify p. Explain the `&` before checking earlier work.

`LoadConfig` is the application task: inspect failures with `errors.Is` and
`errors.As` to prove that added context did not destroy the cause. In the JSON
fixme, predict the bytes before running the test. Add the requested empty-file
test yourself.

Optional stretch: decode multiple JSON values from a stream with `json.Decoder`
instead of reading the whole input first.

Expected time: 4-5 hours.
