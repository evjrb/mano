package main

import (
	"context"
	"os"

	"github.com/mano/client/argyle"
	"github.com/mano/client/storage"
	"github.com/mano/common"
	"github.com/mano/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger := zerolog.New(os.Stdout)

	ctx := logger.WithContext(context.Background())

	dbClient, err := storage.NewDBConnection(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("establish db connection")
		return
	}

	argyleClient := argyle.NewArgyleClient()

	srv, err := server.NewManoServer(ctx, dbClient, argyleClient)
	if err != nil {
		log.Error().Err(err).Msg(common.ErrStartServer)
		return
	}

	srv.Start(ctx, 9090)
}
