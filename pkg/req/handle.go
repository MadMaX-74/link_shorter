package req

import (
	"go_dev/pkg/res"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := DecodeJSON[T](r.Body)
	if err != nil {
		res.JsonResponse(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = isValid(body)
	if err != nil {
		res.JsonResponse(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}
