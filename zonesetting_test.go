// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZoneSettingList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
	)
	_, err := client.Zones.Settings.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSettingEdit(t *testing.T) {
	t.Skip("oneOf doesnt match")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
	)
	_, err := client.Zones.Settings.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneSettingEditParams{
			Items: cloudflare.F([]cloudflare.ZoneSettingEditParamsItem{cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnline(cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnline{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnlineIDAlwaysOnline),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnlineValueOn),
			}), cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTTL(cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTTL{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTTLIDBrowserCacheTTL),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue18000),
			}), cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocation(cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocation{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocationIDIPGeolocation),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocationValueOff),
			})}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
