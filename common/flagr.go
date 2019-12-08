package common

import (
	"context"

	"github.com/antihax/optional"
	"github.com/checkr/goflagr"
)

// GetFlagByKey returns flag by key name
func GetFlagByKey(ctx context.Context, client *goflagr.APIClient, key string) (*goflagr.Flag, error) {
	flags, _, err := client.FlagApi.FindFlags(ctx, &goflagr.FindFlagsOpts{
		Preload: optional.NewBool(true),
	})
	if err != nil {
		return nil, err
	}

	for _, flag := range flags {
		if flag.Key == key {
			return &flag, nil
		}
	}
	return nil, nil
}

// GetSegmentByKey returns segment by segment description
func GetSegmentByKey(ctx context.Context, flag goflagr.Flag, key string) (*goflagr.Segment, error) {
	for _, segment := range flag.Segments {
		if segment.Description == key {
			return &segment, nil
		}
	}
	return nil, nil
}

// GetVariantMap returns variant map with id
func GetVariantMap(ctx context.Context, flag goflagr.Flag) map[string]int64 {
	variantMaps := make(map[string]int64)
	for _, variant := range flag.Variants {
		variantMaps[variant.Key] = variant.Id
	}
	return variantMaps
}
