// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SettingH2PrioritizationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingH2PrioritizationService] method instead.
type SettingH2PrioritizationService struct {
	Options []option.RequestOption
}

// NewSettingH2PrioritizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingH2PrioritizationService(opts ...option.RequestOption) (r *SettingH2PrioritizationService) {
	r = &SettingH2PrioritizationService{}
	r.Options = opts
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *SettingH2PrioritizationService) Edit(ctx context.Context, zoneID string, body SettingH2PrioritizationEditParams, opts ...option.RequestOption) (res *SettingH2PrioritizationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingH2PrioritizationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *SettingH2PrioritizationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingH2PrioritizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingH2PrioritizationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingH2PrioritizationEditResponse struct {
	// ID of the zone setting.
	ID SettingH2PrioritizationEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingH2PrioritizationEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingH2PrioritizationEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingH2PrioritizationEditResponseJSON `json:"-"`
}

// settingH2PrioritizationEditResponseJSON contains the JSON metadata for the
// struct [SettingH2PrioritizationEditResponse]
type settingH2PrioritizationEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingH2PrioritizationEditResponseID string

const (
	SettingH2PrioritizationEditResponseIDH2Prioritization SettingH2PrioritizationEditResponseID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingH2PrioritizationEditResponseValue string

const (
	SettingH2PrioritizationEditResponseValueOn     SettingH2PrioritizationEditResponseValue = "on"
	SettingH2PrioritizationEditResponseValueOff    SettingH2PrioritizationEditResponseValue = "off"
	SettingH2PrioritizationEditResponseValueCustom SettingH2PrioritizationEditResponseValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingH2PrioritizationEditResponseEditable bool

const (
	SettingH2PrioritizationEditResponseEditableTrue  SettingH2PrioritizationEditResponseEditable = true
	SettingH2PrioritizationEditResponseEditableFalse SettingH2PrioritizationEditResponseEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingH2PrioritizationGetResponse struct {
	// ID of the zone setting.
	ID SettingH2PrioritizationGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingH2PrioritizationGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingH2PrioritizationGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingH2PrioritizationGetResponseJSON `json:"-"`
}

// settingH2PrioritizationGetResponseJSON contains the JSON metadata for the struct
// [SettingH2PrioritizationGetResponse]
type settingH2PrioritizationGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingH2PrioritizationGetResponseID string

const (
	SettingH2PrioritizationGetResponseIDH2Prioritization SettingH2PrioritizationGetResponseID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingH2PrioritizationGetResponseValue string

const (
	SettingH2PrioritizationGetResponseValueOn     SettingH2PrioritizationGetResponseValue = "on"
	SettingH2PrioritizationGetResponseValueOff    SettingH2PrioritizationGetResponseValue = "off"
	SettingH2PrioritizationGetResponseValueCustom SettingH2PrioritizationGetResponseValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingH2PrioritizationGetResponseEditable bool

const (
	SettingH2PrioritizationGetResponseEditableTrue  SettingH2PrioritizationGetResponseEditable = true
	SettingH2PrioritizationGetResponseEditableFalse SettingH2PrioritizationGetResponseEditable = false
)

type SettingH2PrioritizationEditParams struct {
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value param.Field[SettingH2PrioritizationEditParamsValue] `json:"value,required"`
}

func (r SettingH2PrioritizationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingH2PrioritizationEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingH2PrioritizationEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingH2PrioritizationEditParamsValueValue] `json:"value,required"`
}

func (r SettingH2PrioritizationEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingH2PrioritizationEditParamsValueID string

const (
	SettingH2PrioritizationEditParamsValueIDH2Prioritization SettingH2PrioritizationEditParamsValueID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingH2PrioritizationEditParamsValueValue string

const (
	SettingH2PrioritizationEditParamsValueValueOn     SettingH2PrioritizationEditParamsValueValue = "on"
	SettingH2PrioritizationEditParamsValueValueOff    SettingH2PrioritizationEditParamsValueValue = "off"
	SettingH2PrioritizationEditParamsValueValueCustom SettingH2PrioritizationEditParamsValueValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingH2PrioritizationEditParamsValueEditable bool

const (
	SettingH2PrioritizationEditParamsValueEditableTrue  SettingH2PrioritizationEditParamsValueEditable = true
	SettingH2PrioritizationEditParamsValueEditableFalse SettingH2PrioritizationEditParamsValueEditable = false
)

type SettingH2PrioritizationEditResponseEnvelope struct {
	Errors   []SettingH2PrioritizationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingH2PrioritizationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result SettingH2PrioritizationEditResponse             `json:"result"`
	JSON   settingH2PrioritizationEditResponseEnvelopeJSON `json:"-"`
}

// settingH2PrioritizationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingH2PrioritizationEditResponseEnvelope]
type settingH2PrioritizationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingH2PrioritizationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingH2PrioritizationEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingH2PrioritizationEditResponseEnvelopeErrors]
type settingH2PrioritizationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingH2PrioritizationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingH2PrioritizationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingH2PrioritizationEditResponseEnvelopeMessages]
type settingH2PrioritizationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationGetResponseEnvelope struct {
	Errors   []SettingH2PrioritizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingH2PrioritizationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result SettingH2PrioritizationGetResponse             `json:"result"`
	JSON   settingH2PrioritizationGetResponseEnvelopeJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingH2PrioritizationGetResponseEnvelope]
type settingH2PrioritizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingH2PrioritizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingH2PrioritizationGetResponseEnvelopeErrors]
type settingH2PrioritizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingH2PrioritizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingH2PrioritizationGetResponseEnvelopeMessages]
type settingH2PrioritizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
