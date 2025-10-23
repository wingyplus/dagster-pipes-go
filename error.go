package dagster_pipes

import (
	"github.com/wingyplus/dagster-pipes-go/internal/helper"
	"github.com/wingyplus/dagster-pipes-go/types"
)

func PipesExceptionError(err error) *types.PipesException {
	return &types.PipesException{
		Message: helper.Ptr(err.Error()),
		Name:    helper.Ptr(err.Error()),
	}
}
