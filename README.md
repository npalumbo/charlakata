# Charla Kata (Kata for the Talk)

This project was created for my talk at GDG DevFest Sevilla 2025: "Testing with Go".

## What's in this repo:

- Kata's unit tests: [player_test.go](player_test.go)
- Benchmark testing: [player_benchmark_test.go](player_benchmark_test.go)
- Fuzz testing: [player_fuzz_test.go](player_fuzz_test.go)
- Using mocks: [player_with_mock_test.go](player_with_mock_test.go)
- Using spies: [player_with_spy_test.go](player_with_spy_test.go)
- Using testify assertions: [player_with_testify_test.go](player_with_testify_test.go)
- Ginkgo and Gomega: [player_with_ginkgo_test.go](player_with_ginkgo_test.go)
- DB Integration Test with testcontainers: [player_repository_testcontainers_test.go](player_repository_testcontainers_test.go)
- Slides from my talk at GDG DevFest Sevilla 2025:
  - English: [EN_GDG_DevFest_Testing_with_Go.pdf](EN_GDG_DevFest_Testing_with_Go.pdf)
  - Spanish: [ES_GDG_DevFest_Testeando_con_Go.pdf](ES_GDG_DevFest_Testeando_con_Go.pdf)

## How to run it:
You'd need go 1.25 or higher. 
Run `make tools` on you terminal to install the tools (golangci-lint, mockery and Ginkgo).
Run `make test` to execute the tests.

## Kata

The repo implements a subset of the Kata documented [here](https://github.com/ardalis/kata-catalog/blob/main/katas/RPG%20Combat.md).

```
Iteration One

All Characters, when created, have:

Health, starting at 1000
Level, starting at 1
May be Alive or Dead, starting Alive
Characters can Deal Damage to Characters.

Damage is subtracted from Health
When damage received exceeds current Health, Health becomes 0 and the character dies.
```