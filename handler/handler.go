package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cikupin/feature-flag-example/constant"

	"github.com/cikupin/feature-flag-example/common"

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
}

// ToggleProvider defines handler for provider toggle simulation
func (h *Handler) ToggleProvider(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	flag, err := common.GetFlagByKey(ctx, h.Client, constant.FlagWhatsappNotification)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}

	if flag == nil || !flag.Enabled {
		fmt.Fprintln(w, "Whatsapp feature is not enabled")
		return
	}

	segment, err := common.GetSegmentByKey(ctx, h.Client, flag.Id, constant.SegmentDefault)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}

	if segment == nil {
		fmt.Fprintln(w, "No segment found")
		return
	}

	for _, dist := range segment.Distributions {
		if dist.Percent == 100 {
			fmt.Fprintf(w, "WA notification service : %s\n", dist.VariantKey)
			return
		}
	}

	fmt.Fprintln(w, "No single WA notification service found")
}
