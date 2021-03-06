package accounts

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/keratin/authn-server/api"
	"github.com/keratin/authn-server/services"
)

func patchAccountExpirePassword(app *api.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			api.WriteNotFound(w, "account")
			return
		}

		err = services.PasswordExpirer(app.AccountStore, app.RefreshTokenStore, id)
		if err != nil {
			if _, ok := err.(services.FieldErrors); ok {
				api.WriteNotFound(w, "account")
				return
			}

			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
