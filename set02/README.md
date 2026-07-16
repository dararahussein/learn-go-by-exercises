# Set 02 - Structs, methods, and pointers

Structs group data; methods attach behavior; pointers let methods mutate the
original value. Pointers come before interfaces so the next set can make method
sets and interface satisfaction concrete.

Work in order: `01_structs`, `02_pointers`, `03_playlist`, `04_debugging`.
Enter each directory and run plain `go test`; its board stops at the first
unfinished step.

Complete Temperature, Counter, and Playlist. Playlist is the application task
and revisits slice ownership: a getter should not accidentally expose mutable
internal storage. Diagnose the value-receiver copy in `04_debugging` last.

Optional stretch: add a timestamped playlist by embedding Playlist and predict
which methods are promoted.

Expected time: 4-5 hours.
