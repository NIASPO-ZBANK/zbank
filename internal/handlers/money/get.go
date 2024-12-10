package money

import (
	"net/http"

	"users/internal/usecase"
	httpErr "users/pkg/http/error"
	"users/pkg/http/writer"
)

func GetMoney(uc usecase.MoneyUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		balance, err := uc.GetBalanceAndDeposit(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, balance)
	}
}
