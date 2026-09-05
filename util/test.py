from dataclasses import dataclass
from typing import Any, Callable

@dataclass
class Test:
    args: tuple
    expected: Any

    def __init__(self, *args, expected):
        self.args = args
        self.expected = expected

def run(
    solution: Callable,
    tests: list[Test],
    *,
    name: str | None = None,
    normalize: Callable[[Any], Any] | None = None,
):
    """
    Run a LeetCode solution against a collection of test cases.

    Example:

        tests = [
            Test([2, 7, 11, 15], 9),
            Test([3, 2, 4], 6),
        ]

        run(Solution().twoSum, tests)

    `normalize` can be used when the ordering of the result doesn't matter:

        run(
            Solution().threeSum,
            tests,
            normalize=lambda x: sorted(sorted(v) for v in x)
        )
    """

    name = name or getattr(solution, "__name__", "solution")

    passed = 0

    print(f"Running {name}...")
    print()

    for i, test in enumerate(tests, 1):
        try:
            actual = solution(*test.args)

            expected = test.expected

            if normalize:
                actual = normalize(actual)
                expected = normalize(expected)

            if actual == expected:
                print(f"  ✓ Test {i}")
                passed += 1
            else:
                print(f"  ✗ Test {i}")
                print(f"      Input:    {_format_args(test.args)}")
                print(f"      Expected: {expected!r}")
                print(f"      Actual:   {actual!r}")

        except Exception as e:
            print(f"  ✗ Test {i}")
            print(f"      Input:    {_format_args(test.args)}")
            print(f"      Error:    {type(e).__name__}: {e}")

    failed = len(tests) - passed

    print()
    print(f"Result: {passed}/{len(tests)} passed")

    if failed:
        print(f"         {failed} failed")

    return failed == 0


def _format_args(args):
    if len(args) == 1:
        return repr(args[0])

    return repr(args)