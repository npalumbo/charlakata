# Charla Kata

Este proyecto fue creado para para mi charla en el GDG DevFest Sevilla 2025: "Testeando con Go".

## Contenido de este repositorio:

- Tests unitarios de la Kata: [player_test.go](player_test.go)
- Pruebas de rendimiento (Benchmark testing): [player_benchmark_test.go](player_benchmark_test.go)
- Pruebas Fuzzing: [player_fuzz_test.go](player_fuzz_test.go)
- Uso de Mocks: [player_with_mock_test.go](player_with_mock_test.go)
- Uso de Spies: [player_with_spy_test.go](player_with_spy_test.go)
- Uso de Testify: [player_with_testify_test.go](player_with_testify_test.go)
- Ginkgo y Gomega: [player_with_ginkgo_test.go](player_with_ginkgo_test.go)
- Integration tests de Base de Datos con testcontainers: [player_repository_testcontainers_test.go](player_repository_testcontainers_test.go)
- Diapositivas de mi charla en el GDG DevFest Sevilla 2025:
    - Inglés: [EN_GDG_DevFest_Testing_with_Go.pdf](EN_GDG_DevFest_Testing_with_Go.pdf)
    - Español: [ES_GDG_DevFest_Testeando_con_Go.pdf](ES_GDG_DevFest_Testeando_con_Go.pdf)

## Kata

El repositorio implementa un subconjunto de la Kata documentada [aquí](https://github.com/ardalis/kata-catalog/blob/main/katas/RPG%20Combat.md).

```
Iteración Uno

Todos los Personajes, al ser creados, tienen:

Salud (Health), que comienza en 1000
Nivel (Level), que comienza en 1
Puede estar Vivo o Muerto (Alive or Dead), comenzando Vivo
Los Personajes pueden Infligir Daño (Deal Damage) a otros Personajes.

El Daño se resta de la Salud
Cuando el daño recibido excede la Salud actual, la Salud pasa a ser 0 y el personaje muere.
```