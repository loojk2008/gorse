package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhenghaoz/gorse/base"
	"github.com/zhenghaoz/gorse/core"
	"reflect"
	"testing"
)

func ModelParallelTest(t *testing.T, models ...core.ModelInterface) {
	data := core.LoadDataFromBuiltIn("ml-100k")
	splitter := core.NewRatioSplitter(1, 0.2)
	trains, tests := splitter(data, 0)
	for _, model := range models {
		t.Log("Checking", reflect.TypeOf(model))
		// Fit model
		model.Fit(trains[0], &base.RuntimeOptions{NJobs: 1})
		rmse := core.EvaluateRating(model, tests[0], core.RMSE)[0]
		// Refit model
		model.Fit(trains[0], &base.RuntimeOptions{NJobs: 3})
		rmse2 := core.EvaluateRating(model, tests[0], core.RMSE)[0]
		assert.Equal(t, rmse, rmse2)
	}
}

func TestModelParallel(t *testing.T) {
	ModelParallelTest(t,
		NewSlopOne(nil),
		NewKNN(nil))
}
