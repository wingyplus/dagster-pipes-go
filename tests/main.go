package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"reflect"

	dagster_pipes "github.com/wingyplus/dagster-pipes-go"
	"github.com/wingyplus/dagster-pipes-go/metadata"
	"github.com/wingyplus/dagster-pipes-go/types"
)

func main() {
	testName := flag.String("test-name", "", "")
	ctx := flag.String("context", "", "")
	messages := flag.String("messages", "", "")
	jobName := flag.String("job-name", "", "")
	extras := flag.String("extras", "", "")
	customPayload := flag.String("custom-payload", "", "")
	reportAssetCheck := flag.String("report-asset-check", "", "")
	reportAssetMaterialization := flag.String("report-asset-materialization", "", "")
	// messageWriter := flag.String("message-writer", "", "")
	// contextLoader := flag.String("context-loader", "", "")

	flag.Parse()

	ensureRequired(testName)

	_ = *reportAssetCheck

	if len(*ctx) > 0 {
		os.Setenv(dagster_pipes.DAGSTER_PIPES_CONTEXT_ENV_VAR, *ctx)
	}
	if len(*messages) > 0 {
		os.Setenv(dagster_pipes.DAGSTER_PIPES_MESSAGES_ENV_VAR, *messages)
	}

	pipesCtx, err := dagster_pipes.OpenDasterPipes()
	if err != nil {
		panic(err)
	}

	if jb := *jobName; len(jb) > 0 {
		if *pipesCtx.Data.JobName != jb {
			log.Fatalf("Want %s, got %s", jb, *pipesCtx.Data.JobName)
		}
	}

	if ext := *extras; len(ext) > 0 {
		f, err := os.Open(ext)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		var e map[string]any
		if err := json.NewDecoder(f).Decode(&e); err != nil {
			panic(err)
		}

		if !reflect.DeepEqual(pipesCtx.Data.Extras, e) {
			log.Fatalf("Want %v, got %v", e, pipesCtx.Data.Extras)
		}
	}

	switch *testName {
	case "test_message_log":
		testMessageLog(*ctx)
	case "test_message_report_custom_message":
		testMessageReportCustomMessage(pipesCtx, *customPayload)
	case "test_message_report_asset_materialization":
		testMessageReportAssetMaterialization(pipesCtx, *reportAssetMaterialization)
	case "test_message_report_asset_check":
		testMessageReportAssetCheck(pipesCtx, *reportAssetCheck)
	}
}

func testMessageLog(ctx string) {
	panic("unimplemented")
}

func testMessageReportCustomMessage(ctx *dagster_pipes.PipesContext, customPayload string) {
	if len(customPayload) == 0 {
		panic("customPayload is required")
	}

	f, err := os.Open(customPayload)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data struct {
		Payload any `json:"payload"`
	}
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		panic(err)
	}

	if err := ctx.ReportCustomMessage(data.Payload); err != nil {
		panic(err)
	}
}

func testMessageReportAssetMaterialization(ctx *dagster_pipes.PipesContext, reportAssetMaterialization string) {
	if len(reportAssetMaterialization) == 0 {
		panic("reportAssetMaterialization is required")
	}

	f, err := os.Open(reportAssetMaterialization)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data struct {
		DataVersion string `json:"dataVersion"`
		AssetKey    string `json:"assetKey"`
	}
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		panic(err)
	}

	if err := ctx.ReportAssetMaterialization(data.AssetKey, buildAssetMetadata(), data.DataVersion); err != nil {
		panic(err)
	}
}

func testMessageReportAssetCheck(ctx *dagster_pipes.PipesContext, reportAssetCheck string) {
	if len(reportAssetCheck) == 0 {
		panic("reportAssetCheck is required")
	}

	f, err := os.Open(reportAssetCheck)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data struct {
		CheckName string `json:"checkName"`
		Passed    bool   `json:"passed"`
		AssetKey  string `json:"assetKey"`
		Severity  string `json:"severity"`
	}
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		panic(err)
	}

	var severity *types.AssetCheckSeverity
	if len(data.Severity) > 0 {
		var sev types.AssetCheckSeverity
		switch data.Severity {
		case "ERROR":
			sev = types.AssetCheckSeverityERROR
		case "WARN":
			sev = types.Warn
		default:
			panic("unknown severity: " + data.Severity)
		}
		severity = &sev
	}

	if err := ctx.ReportAssetCheck(data.CheckName, data.Passed, data.AssetKey, severity, buildAssetMetadata()); err != nil {
		panic(err)
	}
}

func buildAssetMetadata() map[string]*types.PipesMetadataValue {
	return map[string]*types.PipesMetadataValue{
		"float":       metadata.FromFloat(0.1),
		"int":         metadata.FromInt(1),
		"text":        metadata.FromText("hello"),
		"notebook":    metadata.FromNotebook("notebook.ipynb"),
		"md":          metadata.FromMd("**markdown**"),
		"bool_true":   metadata.FromBool(true),
		"bool_false":  metadata.FromBool(false),
		"asset":       metadata.FromAsset("foo/bar"),
		"dagster_run": metadata.FromDagsterRun("db892d7f-0031-4747-973d-22e8b9095d9d"),
		"null":        metadata.Null(),
		"url":         metadata.FromURLString("https://dagster.io"),
		"path":        metadata.FromPath("/dev/null"),
		"json": metadata.FromJSON(map[string]any{
			"quux":  map[string]any{"a": 1, "b": 2},
			"baz":   1,
			"foo":   "bar",
			"corge": nil,
			"qux":   []any{1, 2, 3},
		}),
	}
}

func ensureRequired(p *string) {
	if len(*p) == 0 {
		flag.Usage()
		os.Exit(1)
	}
}
