/*
Package dagster_pipes provides a Go implementation of the Dagster Pipes protocol
for communication between Dagster orchestration processes and external Go applications.

# Overview

Dagster Pipes enables full observability into your Go workloads when orchestrating
through Dagster. With this lightweight interface, you can:

  - Retrieve data directly from the Dagster context
  - Report asset materializations with rich metadata
  - Report asset checks for data quality validation
  - Send custom structured messages to Dagster
  - Provide structured logging (planned feature)

# Installation

	go get github.com/wingyplus/dagster-pipes-go

# Quick Start

Here's a minimal example of using Dagster Pipes in your Go application:

	package main

	import (
	    "log"

	    dagster_pipes "github.com/wingyplus/dagster-pipes-go"
	    "github.com/wingyplus/dagster-pipes-go/metadata"
	    "github.com/wingyplus/dagster-pipes-go/types"
	)

	func main() {
	    // Open connection to Dagster
	    context, err := dagster_pipes.OpenDasterPipes()
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer context.Close(nil)

	    // Your application logic here
	    processData()

	    // Report materialization to Dagster
	    err = context.ReportAssetMaterialization(
	        "example_go_asset",
	        map[string]*types.PipesMetadataValue{
	            "row_count": metadata.FromInt(1000),
	            "status": metadata.FromText("success"),
	        },
	        "v1", // data version
	    )
	    if err != nil {
	        log.Fatal(err)
	    }
	}

	func processData() {
	    // Your data processing logic
	}

# Running from Dagster

To run your Go application from Dagster, use PipesSubprocessClient:

	import dagster as dg

	@dg.asset(kinds={"go"})
	def example_go_asset(
	    context: dg.AssetExecutionContext,
	    pipes_subprocess_client: dg.PipesSubprocessClient
	) -> dg.MaterializeResult:
	    return pipes_subprocess_client.run(
	        command=["./your-go-binary"],
	        context=context,
	    ).get_materialize_result()

	defs = dg.Definitions(
	    assets=[example_go_asset],
	    resources={
	        "pipes_subprocess_client": dg.PipesSubprocessClient(
	            context_injector=dg.PipesEnvContextInjector(),
	        )
	    },
	)

# Metadata Types

The metadata package provides helpers for all supported Dagster metadata types:

  - metadata.FromInt(n) - Integer values
  - metadata.FromFloat(n) - Floating point values
  - metadata.FromBool(b) - Boolean values
  - metadata.FromText(s) - Text strings
  - metadata.FromJSON(m) - JSON objects
  - metadata.FromJSONArray(a) - JSON arrays
  - metadata.FromURL(u) - URLs (rendered as links)
  - metadata.FromPath(p) - File paths
  - metadata.FromMd(md) - Markdown (rendered in UI)
  - metadata.FromTimestamp(t) - Unix timestamps
  - metadata.FromAsset(a) - Asset references (links to other assets)
  - metadata.FromJob(j) - Job references
  - metadata.FromDagsterRun(r) - Run references
  - metadata.FromNotebook(n) - Notebook data
  - metadata.Null() - Null values

# Asset Checks

Report data quality checks with ReportAssetCheck:

	severity := types.AssetCheckSeverityERROR
	err := context.ReportAssetCheck(
	    "row_count_validation",
	    true, // passed
	    "my_dataset",
	    &severity,
	    map[string]*types.PipesMetadataValue{
	        "actual": metadata.FromInt(1000),
	        "expected_min": metadata.FromInt(100),
	    },
	)

# Custom Messages

Send arbitrary structured data to Dagster:

	err := context.ReportCustomMessage(map[string]any{
	    "event_type": "processing_milestone",
	    "progress": 0.75,
	    "items_processed": 7500,
	})

# Error Handling

To report exceptions to Dagster, pass a PipesException to Close:

	defer func() {
	    if r := recover(); r != nil {
	        exc := &types.PipesException{
	            Message: helper.Ptr(fmt.Sprintf("panic: %v", r)),
	            Name:    helper.Ptr("panic"),
	            Stack:   debug.Stack(), // if you want stack trace
	        }
	        context.Close(exc)
	    }
	}()

# Context Data

The PipesContext.Data field contains information passed from Dagster:

  - AssetKeys: The keys of assets being materialized
  - RunID: The unique ID of the current Dagster run
  - JobName: The name of the Dagster job
  - PartitionKey: Partition key if this is a partitioned asset
  - RetryNumber: How many times this run has been retried
  - Extras: Custom data passed from Dagster

Access this data to make your application context-aware:

	runID := context.Data.RunID
	if context.Data.PartitionKey != nil {
	    partition := *context.Data.PartitionKey
	    // Process specific partition
	}

# More Information

For more information about Dagster Pipes:

  - Dagster Pipes documentation: https://docs.dagster.io/concepts/dagster-pipes
  - This repository: https://github.com/wingyplus/dagster-pipes-go
  - Dagster documentation: https://docs.dagster.io
*/
package dagster_pipes
