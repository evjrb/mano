package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mano/mano-server/client/argyle"
	"github.com/mano/mano-server/client/storage"
	"github.com/mano/mano-server/models"
	"github.com/rs/zerolog/log"
)

type RequestHandler struct {
	DbClient *storage.DbClient
	AgClient *argyle.ArgyleClient
}

type CreateUserResponse struct {
	UserToken string `json:"user_token"`
}

func NewRequestHandler(ctx context.Context, dbClient *storage.DbClient, argyleClient *argyle.ArgyleClient) *RequestHandler {
	return &RequestHandler{
		DbClient: dbClient,
		AgClient: argyleClient,
	}
}

func (r *RequestHandler) OnboardUser(w http.ResponseWriter, req *http.Request) {
	resEncoder := json.NewEncoder(w)
	ctx := req.Context()

	var newUser models.User
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		log.Error().Err(err).Msg("unmarshal request")
		formatAndSendErrResp(ctx, w, resEncoder, http.StatusBadRequest, "request malformed", "bad_requqest")
		return
	}

	userId := generateUserId(ctx)
	newUser.UserId = userId

	token, err := r.AgClient.GenerateNewUserToken(ctx)
	if err != nil {
		log.Error().Err(err).Msg("generate user token")
		formatAndSendErrResp(ctx, w, resEncoder, http.StatusInternalServerError, "could not generate user token", "internal_error")
		return
	}

	newUser.UserToken = token

	_, errCreate := r.DbClient.CreateNewUser(ctx, &newUser)
	if errCreate != nil {
		log.Error().Err(err).Msg("create new user")
		formatAndSendErrResp(ctx, w, resEncoder, http.StatusInternalServerError, "internal server error", "internal_errro")
		return
	}

	tokenRes := CreateUserResponse{
		UserToken: token,
	}

	log.Info().Msg("succesfully onboarded user")

	w.WriteHeader(http.StatusOK)
	err2 := resEncoder.Encode(tokenRes)
	if err2 != nil {
		return
	}
}

func (r *RequestHandler) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func (r *RequestHandler) Health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func formatAndSendErrResp(ctx context.Context, w http.ResponseWriter, resEncoder *json.Encoder, status int, message string, errorType string) {
	w.WriteHeader(status)
	respErr := ErrorResponse{
		Message: message,
		Type:    errorType,
	}
	err := resEncoder.Encode(respErr)
	if err != nil {
		log.Error().Err(err).Msg("encoding response")
	}
}

func generateUserId(ctx context.Context) string {
	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return id.String()
}
