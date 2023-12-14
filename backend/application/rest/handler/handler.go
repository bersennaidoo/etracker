package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/bersennaidoo/etracker/backend/application/rest/auth"
	"github.com/bersennaidoo/etracker/backend/infrastructure/storage/pgstore"
	"github.com/gorilla/sessions"
)

var (
	cookieStore = sessions.NewCookieStore([]byte("forDemo"))
)

func init() {
	cookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
}

type Handler struct {
	store pgstore.Querier
}

func New(st pgstore.Querier) *Handler {
	return &Handler{
		store: st,
	}
}

func (h *Handler) HandleLogin(w http.ResponseWriter, req *http.Request) {

	type loginRequest struct {
		Username string `json:"username,omitemty"`
		Password string `json:"password,omitempty"`
	}

	payload := loginRequest{}
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		log.Println("Error decoding the body", err)
		JSONError(w, http.StatusBadRequest, "Error decoding JSON")
		return
	}

	user, err := h.store.GetUserByName(req.Context(), payload.Username)
	if errors.Is(err, sql.ErrNoRows) || !auth.CheckPasswordHash(payload.Password, user.PasswordHash) {
		JSONError(w, http.StatusForbidden, "Bad Credentials")
		return
	}
	if err != nil {
		log.Println("Received error looking up user", err)
		JSONError(w, http.StatusInternalServerError, "Couldn't log you in due to a server error")
		return
	}

	session, err := cookieStore.Get(req, "session-name")
	if err != nil {
		log.Println("Cookie store failed with", err)
		JSONError(w, http.StatusInternalServerError, "Session Error")
	}
	session.Values["userAuthenticated"] = true
	session.Values["userID"] = user.UserID
	session.Save(req, w)
}
