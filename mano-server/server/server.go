package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mano/mano-server/client/argyle"
	"github.com/mano/mano-server/client/storage"
	"github.com/rs/zerolog/log"
)

type ManoServer struct {
	Router *mux.Router
}

func NewManoServer(ctx context.Context, dbClient *storage.DbClient, argyleClient *argyle.ArgyleClient) (ManoServer, error) {

	mux := mux.NewRouter()

	requestHandler := NewRequestHandler(ctx, dbClient, argyleClient)

	mux.Use(requestHandler.corsMiddleware)
	mux.HandleFunc("/onboard-user", requestHandler.OnboardUser)
	mux.HandleFunc("/health", requestHandler.Health)

	return ManoServer{
		Router: mux,
	}, nil

}

func (s *ManoServer) Start(ctx context.Context, port int) {
	p := strconv.Itoa(port)

	log.Ctx(ctx).Info().Msgf("listening to mano money requests on %s", p)

	server := &http.Server{
		Addr:              ":" + p,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("starting server")
	}
}
