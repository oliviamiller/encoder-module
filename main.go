// package main is a module with correction-station component
package main

import (
	"context"
	incremental "encoder-mod/src"

	"go.viam.com/rdk/components/encoder"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, logging.NewLogger("encoder"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	encoder1, err := module.NewModuleFromArgs(ctx, logger)

	if err != nil {
		return err
	}
	encoder1.AddModelFromRegistry(ctx, encoder.API, incremental.Model)

	err = encoder1.Start(ctx)
	defer encoder1.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
