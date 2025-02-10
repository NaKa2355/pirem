package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	"github.com/NaKa2355/pirem/internal/app/pirem/usecases"
	bdy "github.com/NaKa2355/pirem/internal/app/pirem/usecases/boundary"
	"github.com/gorilla/mux"
)

type RestApiAdapter struct {
	port   bdy.Boundary
	router mux.Router
}

func NewRestApiAdapter(boundary bdy.Boundary) RestApiAdapter {
	router := mux.NewRouter().StrictSlash(true)
	server := RestApiAdapter{
		port:   boundary,
		router: *router,
	}
	server.router.HandleFunc("/pirem-rest-api/push_button/{button_id}", server.PushButton)
	return server
}

func (s *RestApiAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func ConvertErrorToHttpStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err := err.(type) {
	case *usecases.Error:
		switch err.Code {
		case usecases.CodeAlreadyExists:
			return http.StatusConflict
		case usecases.CodeBusy:
			return http.StatusServiceUnavailable
		case usecases.CodeDataBase:
			return http.StatusInternalServerError
		case usecases.CodeInvaildInput:
			return http.StatusBadRequest
		case usecases.CodeNotFound:
			return http.StatusNotFound
		case usecases.CodeNotSupported:
			return http.StatusNotAcceptable
		case usecases.CodeTimeout:
			return http.StatusRequestTimeout
		default:
			return http.StatusNotImplemented
		}
	default:
		return http.StatusNotImplemented
	}
}

func (r *RestApiAdapter) PushButton(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	err := r.port.PushRemote(req.Context(), bdy.PushButtonInput{
		ButtonId: domain.ButtonID(vars["button_id"]),
	})
	w.WriteHeader(ConvertErrorToHttpStatus(err))
	if err == nil {
		res, _ := json.Marshal(NewRestApiSuccessResponse(struct{}{}))
		w.Write(res)
		return
	}
	res, _ := json.Marshal(NewRestApiErrorResponse(err))
	w.Write(res)
}
