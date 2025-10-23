package metadata

import (
	"net/url"

	"github.com/wingyplus/dagster-pipes-go/types"
)

func ptr[T any](v T) *T {
	return &v
}

func FromInt(n int64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Integer: ptr(n),
		},
		Type: ptr(types.Int),
	}
}

func FromFloat(n float64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Double: ptr(n),
		},
		Type: ptr(types.Float),
	}
}

func FromBool(b bool) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Bool: ptr(b),
		},
		Type: ptr(types.Bool),
	}
}

func FromText(s string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(s),
		},
		Type: ptr(types.Text),
	}
}

func FromJSON(m map[string]any) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			AnythingMap: m,
		},
		Type: ptr(types.JSON),
	}
}

func FromJSONArray(a []any) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			AnythingArray: a,
		},
		Type: ptr(types.JSON),
	}
}

func FromURL(url *url.URL) *types.PipesMetadataValue {
	return FromURLString(url.String())
}

func FromURLString(url string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(url),
		},
		Type: ptr(types.URL),
	}
}

func FromPath(path string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(path),
		},
		Type: ptr(types.Path),
	}
}

func FromNotebook(notebook string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(notebook),
		},
		Type: ptr(types.Notebook),
	}
}

func FromMd(md string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(md),
		},
		Type: ptr(types.Md),
	}
}

func FromTimestamp(timestamp float64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Double: ptr(timestamp),
		},
		Type: ptr(types.Timestamp),
	}
}

func FromAsset(asset string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(asset),
		},
		Type: ptr(types.Asset),
	}
}

func FromJob(job string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(job),
		},
		Type: ptr(types.Job),
	}
}

func FromDagsterRun(dagsterRun string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: ptr(dagsterRun),
		},
		Type: ptr(types.DagsterRun),
	}
}

func Null() *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: nil,
		Type:     ptr(types.Null),
	}
}
