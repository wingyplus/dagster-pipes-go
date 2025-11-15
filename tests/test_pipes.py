from dagster_pipes_tests import PipesTestSuite

class TestGoSuite(PipesTestSuite):
    BASE_ARGS = ["./build/pipes_tests"]
