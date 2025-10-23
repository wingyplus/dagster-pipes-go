# dagster-pipes-go

A pipes implementation for the [Go](https://golang.org/) programming language.

## Usage

### Installation

```sh
go get github.com/wingyplus/dagster-pipes-go
```

### Example

In your Go application, you can use the Dagster Pipes API to interact with Dagster. The example below demonstrates how to use the Dagster context to report materializations to Dagster through the `context.ReportAssetMaterialization` method.

```go
package main

import (
    "log"

    dagster_pipes "github.com/wingyplus/dagster-pipes-go"
    "github.com/wingyplus/dagster-pipes-go/metadata"
)

func main() {
    context, err := dagster_pipes.OpenDasterPipes()
    if err != nil {
        log.Fatal(err)
    }
    defer context.Close(nil)

    // Report asset materialization with metadata
    err = context.ReportAssetMaterialization(
        "example_go_asset",
        map[string]*types.PipesMetadataValue{
            "row_count": metadata.FromInt(100),
        },
        "v1", // data version
    )
    if err != nil {
        log.Fatal(err)
    }
}
```

### Running Go Binaries from Dagster

You can run Go binaries in a subprocess from Dagster. Note that it's also possible to launch processes in external compute environments like Kubernetes.

```python
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
    cmd = ["./your-go-binary"]
    return pipes_subprocess_client.run(
        command=cmd,
        context=context,
    ).get_materialize_result()


defs = dg.Definitions(
    assets=[example_go_subprocess_asset],
    resources={
        "pipes_subprocess_client": dg.PipesSubprocessClient(
            context_injector=dg.PipesEnvContextInjector(),
        )
    },
)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

See the LICENSE file for details.
