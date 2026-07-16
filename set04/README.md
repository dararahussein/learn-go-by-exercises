# Set 04 - Testing and tooling

Until now, tests described the work. Today you write part of the specification.
The test files in this set are deliberately editable.

Work in order: `01_test_first`, then `02_weak_tests`. Run `check` from inside
the current directory; unlike earlier sets, the test itself is part of the
exercise.

Warm-up from memory: write the three-line Go pattern that checks an error and
returns it to the caller.

Work in this order:

1. Enter `01_test_first`, run `check`, and read the deliberate failure.
2. Fill the case table in `01_test_first/exercise_test.go` from the prose contract. Include at
   least one empty, mixed-case, repeated-space, and punctuation case.
3. Implement `Slugify` only after the tests express the contract.
4. Run the weak `TestNormalizeSpaces` test. Notice it passes even though the
   implementation is wrong. Add cases that expose the bug, then fix it.
5. Run `go test -cover`, `go vet`, and `go doc strings.Builder`.

The lesson is not a particular assertion syntax. It is that a green test only
proves the cases it actually exercises.

Optional stretch: benchmark two correct Slugify implementations with
`testing.B` and explain why a tiny benchmark should not drive premature design.

Expected time: 4-5 hours including experiments and reflection.
