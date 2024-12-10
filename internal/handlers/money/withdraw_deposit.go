package money

import (
	"fmt"
	"net/http"
	"strconv"

	"users/internal/usecase"
	httpErr "users/pkg/http/error"
)

func WithdrawFromDeposit(uc usecase.MoneyUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		moneyStr := r.URL.Query().Get("money")
		if moneyStr == "" {
			httpErr.BadRequest(w, fmt.Errorf("no money provided"))
			return
		}

		money, err := strconv.Atoi(moneyStr)
		if err != nil {
			httpErr.BadRequest(w, fmt.Errorf("bad money provided"))
			return
		}

		if err = uc.WithdrawFromDeposit(r.Context(), money); err != nil {
			httpErr.InternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
