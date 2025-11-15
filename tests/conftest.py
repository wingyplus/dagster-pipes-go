import pytest
import subprocess

@pytest.fixture(scope="session",autouse=True)
def built_binary():
    subprocess.run(["go", "build", "-o", "build/pipes_tests", "./tests"], check=True)
