package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/checkr/goflagr"
)

// Handler defines handler struct
type Handler struct {
	Client *goflagr.APIClient
}

// NewHandler return handler instance
func NewHandler(client *goflagr.APIClient) *Handler {
	return &Handler{
		Client: client,
	}
}

// ToggleFeature defines handler for feature toggle simulation
func (h *Handler) ToggleFeature(w http.ResponseWriter, r *http.Request) {
	flags, _, err := h.Client.FlagApi.FindFlags(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}

	if len(flags) == 0 {
		fmt.Fprintln(w, "No feature enabled")
		return
	}

	fmt.Fprintln(w, "active features :")
	for _, flag := range flags {
		if flag.Enabled {
			fmt.Fprintf(w, "- [%s] %s\n", flag.Key, flag.Description)
		}
	}
	return
}

// ToggleProvider defines handler for provider toggle simulation
func (h *Handler) ToggleProvider(w http.ResponseWriter, r *http.Request) {

}
