package metadata

import (
	"net/url"

	"github.com/wingyplus/dagster-pipes-go/internal/helper"
	"github.com/wingyplus/dagster-pipes-go/types"
)

func FromInt(n int64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Integer: helper.Ptr(n),
		},
		Type: helper.Ptr(types.Int),
	}
}

func FromFloat(n float64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Double: helper.Ptr(n),
		},
		Type: helper.Ptr(types.Float),
	}
}

func FromBool(b bool) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Bool: helper.Ptr(b),
		},
		Type: helper.Ptr(types.Bool),
	}
}

func FromText(s string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(s),
		},
		Type: helper.Ptr(types.Text),
	}
}

func FromJSON(m map[string]any) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			AnythingMap: m,
		},
		Type: helper.Ptr(types.JSON),
	}
}

func FromJSONArray(a []any) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			AnythingArray: a,
		},
		Type: helper.Ptr(types.JSON),
	}
}

func FromURL(url *url.URL) *types.PipesMetadataValue {
	return FromURLString(url.String())
}

func FromURLString(url string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(url),
		},
		Type: helper.Ptr(types.URL),
	}
}

func FromPath(path string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(path),
		},
		Type: helper.Ptr(types.Path),
	}
}

func FromNotebook(notebook string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(notebook),
		},
		Type: helper.Ptr(types.Notebook),
	}
}

func FromMd(md string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(md),
		},
		Type: helper.Ptr(types.Md),
	}
}

func FromTimestamp(timestamp float64) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			Double: helper.Ptr(timestamp),
		},
		Type: helper.Ptr(types.Timestamp),
	}
}

func FromAsset(asset string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(asset),
		},
		Type: helper.Ptr(types.Asset),
	}
}

func FromJob(job string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(job),
		},
		Type: helper.Ptr(types.Job),
	}
}

func FromDagsterRun(dagsterRun string) *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: &types.RawValue{
			String: helper.Ptr(dagsterRun),
		},
		Type: helper.Ptr(types.DagsterRun),
	}
}

func Null() *types.PipesMetadataValue {
	return &types.PipesMetadataValue{
		RawValue: nil,
		Type:     helper.Ptr(types.Null),
	}
}
