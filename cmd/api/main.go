package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrique998/email-N/internal/domain/campaing"
	"github.com/henrique998/email-N/internal/endpoints"
	"github.com/henrique998/email-N/internal/infra/database"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaing.ServiceImp{
		Repo: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignFindById))

	http.ListenAndServe(":3333", r)
}
