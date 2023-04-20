package http_error

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	var statusCode int
	switch err.(type) {
	case BadRequestError:
		statusCode = http.StatusBadRequest
	case NotFoundError:
		statusCode = http.StatusNotFound
	case UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case ForbiddenError:
		statusCode = http.StatusForbidden
	case InternalServerError:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = 500
	}

	w.WriteHeader(statusCode)
	b, err := json.Marshal(err)
	if err != nil {
		log.Println("[HttpError] Error Marshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Println("[HttpError] Error Write", err)
		return
	}
}
