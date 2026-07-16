# Learn Go by Exercises

A ten-day, code-first introduction to Go for software engineering interns. The
goal is not expertise in two weeks. The goal is enough fluency to read an
existing Go repository, make small-to-medium changes safely, debug common
failures, and keep learning through documentation and experiments.

Most of your time should be spent writing, running, predicting, debugging, and
explaining code. Explanations here are intentionally short; the exercises and
the feedback loop do most of the teaching.

## The feedback loop

```sh
cd set01/01_functions                  # enter the current exercise module
go test                                # show only the next unfinished exercise
go test -v                             # include completed steps in the module
cd ../..                               # return to the repository root
go test ./set01/...                    # end-of-set check, run from repo root
go test -race ./set08/...              # end-of-set race check
go fmt ./...
go vet ./...
```

If `go test -race` reports that the race detector requires CGO, use the
project's Linux/WSL development environment or enable CGO with a supported C
toolchain. That message is an environment problem, not evidence about the code.

Each numbered directory is one exercise module with a focused test board. Most
modules use one learner-facing `exercise.go`; the later repository and capstone
modules contain realistic package layouts. Work through directories in numeric
order. The board stops after the first failing step, so the terminal shows the
thing you are doing now rather than every unfinished task in the set.

For ordinary stub exercises:

1. Read the task and test failure.
2. Predict what must change.
3. Make one small change.
4. Run `go test` in the module again.
5. When the whole module is green, refactor for clarity and enter the next
   numbered directory.

For numbered `debugging` modules, the implementation already exists and is wrong. Write
your theory in the `// A:` line before editing. Evidence first, fix second.

The course is intentionally red at the beginning, but only the current module's
next step should be visible. Failures should be assertions or deliberate
`t.Fatal("TODO...")` markers, not compile errors.

Comparison failures stop at the first mismatch and use one compact line:

```text
Abs(-3): got 0, want 3
```

One exception is `set04/02_weak_tests`: its inadequate test starts green on
purpose. The exercise is to add cases that expose the implementation's bug.

## Ten sets / ten business days

| Set | Focus | Production habit |
|---|---|---|
| 01 | Functions, control flow, strings, slices, maps | Predict output; debug basic logic |
| 02 | Structs, methods, pointers, composition | Recognize copies versus mutation |
| 03 | Errors and interfaces | Wrap causes; program to behavior |
| 04 | Testing and tooling | Write table tests; expose weak tests |
| 05 | JSON and files | Own resources; preserve error causes |
| 06 | HTTP servers and clients | Test handlers and clients in memory |
| 07 | Goroutines, channels, `select` | Diagnose deadlocks; justify concurrency |
| 08 | Races, synchronization, context | Use the race detector; stop workers |
| 09 | Existing multi-package codebase | Trace, diagnose, and make reviewable changes |
| 10 | Task tracker CLI + HTTP capstone | Integrate persistence, APIs, tests, and safety |

Plan for four to six focused hours per set. Some learners will finish the early
sets faster and need the optional experiments; concurrency and the capstone may
take longer. Stop at a sensible checkpoint rather than rushing to collect green
checks without understanding them.

## Rules of engagement

- Make an independent attempt before asking for a solution.
- Do not edit tests unless the set explicitly says the test belongs to you.
  Sets 04, 09, and 10 deliberately include learner-authored tests.
- Answer every `// Q:` before looking back or asking an assistant.
- Keep [MISTAKES.md](MISTAKES.md): symptom, theory, cause, fix, reusable rule.
- Prefer standard-library documentation: `go doc` and <https://pkg.go.dev/std>.
- Run `go fmt`; do not hand-format Go.
- A working solution is not finished until you can explain it and remove
  temporary debug output.

## Using AI without outsourcing the learning

AI is allowed as a tutor, not as a substitute author.

1. Work independently for at least 15 minutes and record your current theory.
2. Ask for a hint, concept explanation, documentation pointer, or interpretation
   of an exact compiler/test/race error before asking for code.
3. Never paste a complete generated solution into an exercise.
4. If AI suggests a fragment, explain every line, adapt it, and test a case it
   did not mention.
5. Make predictions before execution without AI assistance.
6. At the end of each set, reproduce one core pattern from memory.

The repository's [AGENTS.md](AGENTS.md) tells coding assistants not to solve the
exercises. That constraint is part of the course design.

## Set completion check

Before moving on:

- every numbered module and the full-set check are green;
- `go fmt` and `go vet` are clean for the set;
- you completed the debugging task and wrote its cause in the mistake log;
- you can explain one tradeoff from memory;
- your changes are small enough that another engineer could review them.

Course maintainers should use [CURRICULUM_CHECKLIST.md](CURRICULUM_CHECKLIST.md)
when revising exercises.

## Maintainer validation

The checked-in starter is intentionally red. Validate structural health without
executing exercises:

```sh
go test -run '^$' ./...    # compile every package and test
git diff --check           # reject whitespace errors
```

Then run each set normally and confirm it fails through assertions or an
explicit learner-owned `t.Fatal`, never a compile error, panic, or timeout.
`go vet ./...` intentionally reports the unexported JSON field in
`set05/04_debugging/exercise.go`; that diagnostic is evidence for the exercise and disappears when
the learner repairs it. The opt-in race demonstration in set08 is skipped during
ordinary full-suite and race-detector runs.
