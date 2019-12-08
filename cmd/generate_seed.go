package cmd

import (
	"context"
	"log"

	"github.com/checkr/goflagr"
	"github.com/cikupin/feature-flag-example/common"
	"github.com/cikupin/feature-flag-example/constant"
	"github.com/urfave/cli/v2"
)

// GenerateFeatureToggle will generate feature toggles.
var GenerateFeatureToggle = &cli.Command{
	Name:        "seed",
	Usage:       "generate seed data",
	Description: "generate seed data",
	Action: func(c *cli.Context) error {
		client := common.GetFlagrClient()

		generateFlags(c, client)
		generateVariants(c, client)
		generateSegments(c, client)
		generateDistribution(c, client)
		return nil
	},
}

// generateFlags will create flags:
//   - email notification
//   - sms notification
//   - whatsapp notification
func generateFlags(ctx context.Context, client *goflagr.APIClient) {
	// email
	flag, _, err := client.FlagApi.CreateFlag(ctx, goflagr.CreateFlagRequest{
		Description: "email notification",
		Key:         constant.FlagEmailNotification,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, _, err = client.FlagApi.SetFlagEnabled(ctx, flag.Id, goflagr.SetFlagEnabledRequest{
		Enabled: true,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// sms
	flag, _, err = client.FlagApi.CreateFlag(ctx, goflagr.CreateFlagRequest{
		Description: "sms notification",
		Key:         constant.FlagSMSNotification,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, _, err = client.FlagApi.SetFlagEnabled(ctx, flag.Id, goflagr.SetFlagEnabledRequest{
		Enabled: true,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// whatsapp
	flag, _, err = client.FlagApi.CreateFlag(ctx, goflagr.CreateFlagRequest{
		Description: "whatsapp notification",
		Key:         constant.FlagWhatsappNotification,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, _, err = client.FlagApi.SetFlagEnabled(ctx, flag.Id, goflagr.SetFlagEnabledRequest{
		Enabled: true,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// generateVariants will generate variants for whatsapp provider:
//   - twilio
//   - smooch
//   - infobip
func generateVariants(ctx context.Context, client *goflagr.APIClient) {
	flag, err := common.GetFlagByKey(ctx, client, constant.FlagWhatsappNotification)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if flag == nil {
		log.Fatalln("No flag found")
	}

	// set twilio
	_, _, err = client.VariantApi.CreateVariant(ctx, flag.Id, goflagr.CreateVariantRequest{
		Key:        constant.VariantWhatsappNotifTwilio,
		Attachment: nil,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// set smooch
	_, _, err = client.VariantApi.CreateVariant(ctx, flag.Id, goflagr.CreateVariantRequest{
		Key:        constant.VariantWhatsappNotifSmooch,
		Attachment: nil,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// set infobip
	_, _, err = client.VariantApi.CreateVariant(ctx, flag.Id, goflagr.CreateVariantRequest{
		Key:        constant.VariantWhatsappNotifInfobip,
		Attachment: nil,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// generateSegments will generate segments for whatsapp notif flags
func generateSegments(ctx context.Context, client *goflagr.APIClient) {
	flag, err := common.GetFlagByKey(ctx, client, constant.FlagWhatsappNotification)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if flag == nil {
		log.Fatalln("No flag found")
	}

	_, _, err = client.SegmentApi.CreateSegment(ctx, flag.Id, goflagr.CreateSegmentRequest{
		Description:    constant.SegmentDefault,
		RolloutPercent: 100,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// generateDistribution will generate distribution for whatsapp notif provider
// - twilio 0%
// - smooch 100%
// - infobip 0%
func generateDistribution(ctx context.Context, client *goflagr.APIClient) {
	flag, err := common.GetFlagByKey(ctx, client, constant.FlagWhatsappNotification)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if flag == nil {
		log.Fatalln("No flag found")
	}

	segment, err := common.GetSegmentByKey(ctx, *flag, constant.SegmentDefault)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if segment == nil {
		log.Fatalln("No segment found")
	}

	variantMaps := common.GetVariantMap(ctx, *flag)

	_, _, err = client.DistributionApi.PutDistributions(ctx, flag.Id, segment.Id, goflagr.PutDistributionsRequest{
		Distributions: []goflagr.Distribution{
			{
				Percent:    0,
				VariantKey: constant.VariantWhatsappNotifInfobip,
				VariantID:  variantMaps[constant.VariantWhatsappNotifInfobip],
			},
			{
				Percent:    100,
				VariantKey: constant.VariantWhatsappNotifSmooch,
				VariantID:  variantMaps[constant.VariantWhatsappNotifSmooch],
			},
			{
				Percent:    0,
				VariantKey: constant.VariantWhatsappNotifTwilio,
				VariantID:  variantMaps[constant.VariantWhatsappNotifTwilio],
			},
		},
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}
