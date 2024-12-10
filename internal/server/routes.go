package server

import (
	"github.com/go-chi/cors"
	"users/internal/handlers/money"

	"github.com/go-chi/chi/v5"
)

func (s *Server) initRouter() {
	s.router = chi.NewRouter()

	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Content-Length"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	s.router.Route("/api", func(r chi.Router) {
		r.Route("/money", func(r chi.Router) {
			r.Route("/funds", func(r chi.Router) {
				r.Patch("/add", money.AddFunds(s.moneyUseCase))
				r.Patch("/withdraw", money.WithdrawFunds(s.moneyUseCase))
			})

			r.Route("/deposits", func(r chi.Router) {
				r.Patch("/add", money.DepositFunds(s.moneyUseCase))
				r.Patch("/withdraw", money.WithdrawFromDeposit(s.moneyUseCase))
			})

			r.Get("/", money.GetMoney(s.moneyUseCase))
		})
	})
}
