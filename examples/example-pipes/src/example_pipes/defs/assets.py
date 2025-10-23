from pathlib import Path

import dagster as dg


@dg.asset(
    group_name="pipes",
    kinds={"go"},
)
def example_go_subprocess_asset(
    context: dg.AssetExecutionContext,
    pipes_subprocess_client: dg.PipesSubprocessClient
) -> dg.MaterializeResult:
    """Demonstrates running Go binary in a subprocess."""
    # Path to the Go binary relative to the project root
    binary_path = Path(__file__).parent.parent.parent.parent / "pipes" / "example-pipes"
    cmd = [str(binary_path)]
    return pipes_subprocess_client.run(
        command=cmd,
        context=context,
    ).get_materialize_result()
