// Package metadata provides helper functions for creating typed metadata values
// to attach to asset materializations and asset checks in Dagster Pipes.
//
// Each function creates a PipesMetadataValue with the appropriate type annotation
// that Dagster can properly interpret and display in the UI.
//
// # Usage
//
// Import this package alongside dagster_pipes:
//
//	import (
//	    dagster_pipes "github.com/wingyplus/dagster-pipes-go"
//	    "github.com/wingyplus/dagster-pipes-go/metadata"
//	    "github.com/wingyplus/dagster-pipes-go/types"
//	)
//
// Then use the helper functions when reporting materializations:
//
//	context.ReportAssetMaterialization(
//	    "my_asset",
//	    map[string]*types.PipesMetadataValue{
//	        "row_count": metadata.FromInt(1000),
//	        "table_name": metadata.FromText("users"),
//	        "output_path": metadata.FromPath("/data/output.parquet"),
//	    },
//	    "v1",
//	)
package metadata

import (
	"net/url"

	"github.com/wingyplus/dagster-pipes-go/internal/helper"
	"github.com/wingyplus/dagster-pipes-go/types"
)

// FromInt creates a metadata value from an integer.
//
// Example:
//
//	metadata.FromInt(1000)
func FromInt(n int64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Integer: helper.Ptr(n),
		},
		Type: helper.Ptr(types.Int),
	}
}

// FromFloat creates a metadata value from a floating point number.
//
// Example:
//
//	metadata.FromFloat(0.95)
func FromFloat(n float64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Double: helper.Ptr(n),
		},
		Type: helper.Ptr(types.Float),
	}
}

// FromBool creates a metadata value from a boolean.
//
// Example:
//
//	metadata.FromBool(false)
func FromBool(b bool) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Bool: helper.Ptr(b),
		},
		Type: helper.Ptr(types.Bool),
	}
}

// FromText creates a metadata value from a text string.
//
// Example:
//
//	metadata.FromText("users")
func FromText(s string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(s),
		},
		Type: helper.Ptr(types.Text),
	}
}

// FromJSON creates a metadata value from a JSON object (map).
//
// Example:
//
//	metadata.FromJSON(map[string]any{
//	    "columns": []string{"id", "name", "email"},
//	    "version": 2,
//	})
func FromJSON(m map[string]any) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			AnythingMap: m,
		},
		Type: helper.Ptr(types.JSON),
	}
}

// FromJSONArray creates a metadata value from a JSON array (slice).
//
// Example:
//
//	metadata.FromJSONArray([]any{
//	    map[string]any{"line": 10, "message": "Invalid value"},
//	    map[string]any{"line": 25, "message": "Missing field"},
//	})
func FromJSONArray(a []any) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			AnythingArray: a,
		},
		Type: helper.Ptr(types.JSON),
	}
}

// FromURL creates a metadata value from a url.URL.
//
// Example:
//
//	u, _ := url.Parse("https://example.com/report")
//	metadata.FromURL(u)
func FromURL(url *url.URL) *types.PipesMetadataValue {
	return FromURLString(url.String())
}

// FromURLString creates a metadata value from a URL string.
//
// Example:
//
//	metadata.FromURLString("https://example.com/dashboard")
func FromURLString(url string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(url),
		},
		Type: helper.Ptr(types.URL),
	}
}

// FromPath creates a metadata value from a file system path.
//
// Example:
//
//	metadata.FromPath("/data/outputs/result.parquet")
func FromPath(path string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(path),
		},
		Type: helper.Ptr(types.Path),
	}
}

// FromNotebook creates a metadata value from notebook content.
//
// Example:
//
//	metadata.FromNotebook(notebookJSON)
func FromNotebook(notebook string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(notebook),
		},
		Type: helper.Ptr(types.Notebook),
	}
}

// FromMd creates a metadata value from markdown text.
//
// The markdown will be rendered in the Dagster UI, allowing you to
// include formatted documentation or rich text descriptions.
//
// Example:
//
//	metadata.FromMd("## Processing Results\n\n- Total rows: 1000\n- Errors: 0")
func FromMd(md string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(md),
		},
		Type: helper.Ptr(types.Md),
	}
}

// FromTimestamp creates a metadata value from a Unix timestamp.
//
// The timestamp should be in seconds since the Unix epoch (can include
// fractional seconds). It will be displayed as a formatted date/time in
// the Dagster UI.
//
// Example:
//
//	metadata.FromTimestamp(float64(time.Now().Unix()))
func FromTimestamp(timestamp float64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Double: helper.Ptr(timestamp),
		},
		Type: helper.Ptr(types.Timestamp),
	}
}

// FromAsset creates a metadata value that references another Dagster asset.
//
// This creates a link to the specified asset in the Dagster UI, allowing
// you to show relationships between assets.
//
// Example:
//
//	metadata.FromAsset("raw_user_data")
func FromAsset(asset string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(asset),
		},
		Type: helper.Ptr(types.Asset),
	}
}

// FromJob creates a metadata value that references a Dagster job.
//
// This creates a link to the specified job in the Dagster UI.
//
// Example:
//
//	metadata.FromJob("daily_processing_job")
func FromJob(job string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(job),
		},
		Type: helper.Ptr(types.Job),
	}
}

// FromDagsterRun creates a metadata value that references a Dagster run.
//
// This creates a link to the specified run in the Dagster UI, allowing
// you to reference related pipeline executions.
//
// Example:
//
//	metadata.FromDagsterRun(runID)
func FromDagsterRun(dagsterRun string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(dagsterRun),
		},
		Type: helper.Ptr(types.DagsterRun),
	}
}

// Null creates a null metadata value.
//
// Use this when you need to explicitly set a metadata field to null.
//
// Example:
//
//	metadata.Null()
func Null() *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: nil,
		Type:     helper.Ptr(types.Null),
	}
}
