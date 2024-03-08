// File generated from our OpenAPI spec by Stainless.

package dns

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// FirewallAnalyticsReportService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewFirewallAnalyticsReportService] method instead.
type FirewallAnalyticsReportService struct {
	Options []option.RequestOption
	Bytimes *FirewallAnalyticsReportBytimeService
}

// NewFirewallAnalyticsReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallAnalyticsReportService(opts ...option.RequestOption) (r *FirewallAnalyticsReportService) {
	r = &FirewallAnalyticsReportService{}
	r.Options = opts
	r.Bytimes = NewFirewallAnalyticsReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *FirewallAnalyticsReportService) Get(ctx context.Context, accountIdentifier string, identifier string, query FirewallAnalyticsReportGetParams, opts ...option.RequestOption) (res *DNSDNSAnalyticsAPIReport, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAnalyticsReportGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallAnalyticsReportGetParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [FirewallAnalyticsReportGetParams]'s query parameters as
// `url.Values`.
func (r FirewallAnalyticsReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallAnalyticsReportGetResponseEnvelope struct {
	Errors   []FirewallAnalyticsReportGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAnalyticsReportGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSDNSAnalyticsAPIReport                             `json:"result,required"`
	// Whether the API call was successful
	Success FirewallAnalyticsReportGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAnalyticsReportGetResponseEnvelopeJSON    `json:"-"`
}

// firewallAnalyticsReportGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallAnalyticsReportGetResponseEnvelope]
type firewallAnalyticsReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallAnalyticsReportGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    firewallAnalyticsReportGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAnalyticsReportGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallAnalyticsReportGetResponseEnvelopeErrors]
type firewallAnalyticsReportGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAnalyticsReportGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    firewallAnalyticsReportGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAnalyticsReportGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FirewallAnalyticsReportGetResponseEnvelopeMessages]
type firewallAnalyticsReportGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAnalyticsReportGetResponseEnvelopeSuccess bool

const (
	FirewallAnalyticsReportGetResponseEnvelopeSuccessTrue FirewallAnalyticsReportGetResponseEnvelopeSuccess = true
)
