package main

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"gorm.io/gen"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {

	ctx := context.Background()
	outPath := os.Getenv("OUT_PATH")
	if outPath == "" {
		outPath = "src/repository/query"
	}

	modelPkgPath := os.Getenv("MODEL_PKG_PATH")
	if modelPkgPath == "" {
		modelPkgPath = "db_model"
	}

	generator := gen.NewGenerator(gen.Config{
		OutPath:      outPath,
		ModelPkgPath: modelPkgPath,
		Mode:         gen.WithDefaultQuery,
	})

	klog.CtxInfof(ctx, "outPath: %v", outPath)
	klog.CtxInfof(ctx, "modelPkgPath: %v", modelPkgPath)

	db := config.InitMySQL()

	generator.UseDB(db)
	generator.ApplyBasic(
		// 表
		generator.GenerateModelAs("scs_order", "OrderDO"),
		generator.GenerateModelAs("scs_subscription", "SubscriptionDO"),
	)

	generator.Execute()
}
