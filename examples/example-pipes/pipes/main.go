package main

import (
	"log"

	dagster_pipes "github.com/wingyplus/dagster-pipes-go"
	"github.com/wingyplus/dagster-pipes-go/metadata"
	"github.com/wingyplus/dagster-pipes-go/types"
)

func main() {
	context, err := dagster_pipes.OpenDasterPipes()
	if err != nil {
		log.Fatal(err)
	}
	defer context.Close(nil)

	err = context.ReportAssetMaterialization(
		"example_go_subprocess_asset",
		map[string]*types.PipesMetadataValue{
			"row_count": metadata.FromInt(100),
		},
		"v1",
	)
	if err != nil {
		log.Fatal(err)
	}
}
