# Curriculum revision checklist

## Every exercise

- [ ] Has one clear learning purpose.
- [ ] Requires code, a prediction, debugging, or a written explanation.
- [ ] Cannot be completed by changing one obvious token copied from an example.
- [ ] Introduces no prerequisite that has not appeared earlier.
- [ ] Gives evidence and a direction to investigate without revealing the fix.
- [ ] Accepts idiomatic, maintainable Go rather than merely correct output.

## Every set

- [ ] Starts with retrieval of an earlier concept.
- [ ] Includes several small coding tasks, one debugging task, and one
      integration/application task.
- [ ] Includes reflection and optional depth for faster learners.
- [ ] Builds from guided to independent work without an unexplained jump.
- [ ] Fits four to six focused hours in learner trials.
- [ ] Reuses at least one earlier concept in a harder context.

## The complete program

- [ ] Most learner time is spent writing, running, testing, and debugging code.
- [ ] Learners write tests as well as consume them.
- [ ] Important concepts recur through more than one exercise type.
- [ ] Intentional failures cover copies, slice behavior, errors, JSON,
      resources, weak tests, deadlocks, goroutine lifetime, and races.
- [ ] Concurrency is used only when it has a defensible benefit.
- [ ] Tooling includes targeted tests, formatting, vet, modules, documentation,
      coverage, and the race detector.
- [ ] The final sets require navigation and modification of existing packages.
- [ ] The advertised topics exactly match the checked-in exercises.
- [ ] A pilot learner can complete the core in 40-60 focused hours.
- [ ] A finisher can take a small issue in a real Go repository and produce a
      tested, explainable, reviewable change.
