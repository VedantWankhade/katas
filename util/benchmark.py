# util/benchmark.py

from dataclasses import dataclass
from copy import deepcopy
from time import perf_counter
from typing import Any, Callable
import statistics


@dataclass
class BenchmarkResult:
    name: str
    times: list[float]

    @property
    def average(self) -> float:
        return statistics.mean(self.times)

    @property
    def median(self) -> float:
        return statistics.median(self.times)

    @property
    def minimum(self) -> float:
        return min(self.times)


def run(
    solutions: dict[str, Callable],
    test_cases: list[tuple],
    *,
    runs: int = 10,
    warmup: int = 2,
):
    """
    Benchmark multiple solutions using the same test cases.

    Each solution receives a deep copy of the test cases, so
    mutations made by one solution do not affect another.

    Example:

        run(
            {
                "Brute force": brute_force,
                "Optimized": optimized,
            },
            [
                ([1, 2, 3, 4],),
                ([1, 2, 3, 4, 5],),
            ],
        )
    """

    results = {}

    print("Benchmark")
    print("=" * 60)

    for name, solution in solutions.items():
        times = []

        # Warmup
        for _ in range(warmup):
            for args in test_cases:
                solution(*deepcopy(args))

        # Benchmark
        for _ in range(runs):
            start = perf_counter()

            for args in test_cases:
                solution(*deepcopy(args))

            elapsed = perf_counter() - start
            times.append(elapsed)

        results[name] = BenchmarkResult(name, times)

    _print_results(results)

    # return results


def _print_results(results: dict[str, BenchmarkResult]):
    fastest = min(
        results.values(),
        key=lambda result: result.median,
    )

    print()
    print(
        f"{'Solution':<25}"
        f"{'Average':>12}"
        f"{'Median':>12}"
        f"{'Min':>12}"
    )

    print("-" * 61)

    for result in results.values():
        print(
            f"{result.name:<25}"
            f"{result.average * 1000:>10.3f} ms"
            f"{result.median * 1000:>10.3f} ms"
            f"{result.minimum * 1000:>10.3f} ms"
        )

    print()
    print(f"Fastest: {fastest.name}")

    for result in results.values():
        if result is fastest:
            continue

        ratio = result.median / fastest.median

        print(
            f"  {result.name} is "
            f"{ratio:.2f}x slower"
        )