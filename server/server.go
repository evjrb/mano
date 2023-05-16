package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mano/client/argyle"
	"github.com/mano/client/storage"
	"github.com/mano/common"
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

	return ManoServer{
		Router: mux,
	}, nil

}

func (s *ManoServer) Start(ctx context.Context, port int) {
	p := strconv.Itoa(port)

	log.Ctx(ctx).Info().Msgf("listening to mano money requests on %s", p)
	err := http.ListenAndServe(":"+p, s.Router)
	if err != nil {
		log.Ctx(ctx).Error().Err(errors.New(common.ErrStartServer)).Msg(err.Error())
	}
}
