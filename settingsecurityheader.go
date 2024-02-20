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

// SettingSecurityHeaderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingSecurityHeaderService]
// method instead.
type SettingSecurityHeaderService struct {
	Options []option.RequestOption
}

// NewSettingSecurityHeaderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingSecurityHeaderService(opts ...option.RequestOption) (r *SettingSecurityHeaderService) {
	r = &SettingSecurityHeaderService{}
	r.Options = opts
	return
}

// Cloudflare security header for a zone.
func (r *SettingSecurityHeaderService) Edit(ctx context.Context, zoneID string, body SettingSecurityHeaderEditParams, opts ...option.RequestOption) (res *SettingSecurityHeaderEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityHeaderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_header", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare security header for a zone.
func (r *SettingSecurityHeaderService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingSecurityHeaderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityHeaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_header", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare security header for a zone.
type SettingSecurityHeaderEditResponse struct {
	// ID of the zone's security header.
	ID SettingSecurityHeaderEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSecurityHeaderEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSecurityHeaderEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSecurityHeaderEditResponseJSON `json:"-"`
}

// settingSecurityHeaderEditResponseJSON contains the JSON metadata for the struct
// [SettingSecurityHeaderEditResponse]
type settingSecurityHeaderEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone's security header.
type SettingSecurityHeaderEditResponseID string

const (
	SettingSecurityHeaderEditResponseIDSecurityHeader SettingSecurityHeaderEditResponseID = "security_header"
)

// Current value of the zone setting.
type SettingSecurityHeaderEditResponseValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SettingSecurityHeaderEditResponseValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    settingSecurityHeaderEditResponseValueJSON                    `json:"-"`
}

// settingSecurityHeaderEditResponseValueJSON contains the JSON metadata for the
// struct [SettingSecurityHeaderEditResponseValue]
type settingSecurityHeaderEditResponseValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type SettingSecurityHeaderEditResponseValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                              `json:"nosniff"`
	JSON    settingSecurityHeaderEditResponseValueStrictTransportSecurityJSON `json:"-"`
}

// settingSecurityHeaderEditResponseValueStrictTransportSecurityJSON contains the
// JSON metadata for the struct
// [SettingSecurityHeaderEditResponseValueStrictTransportSecurity]
type settingSecurityHeaderEditResponseValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSecurityHeaderEditResponseEditable bool

const (
	SettingSecurityHeaderEditResponseEditableTrue  SettingSecurityHeaderEditResponseEditable = true
	SettingSecurityHeaderEditResponseEditableFalse SettingSecurityHeaderEditResponseEditable = false
)

// Cloudflare security header for a zone.
type SettingSecurityHeaderGetResponse struct {
	// ID of the zone's security header.
	ID SettingSecurityHeaderGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSecurityHeaderGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSecurityHeaderGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSecurityHeaderGetResponseJSON `json:"-"`
}

// settingSecurityHeaderGetResponseJSON contains the JSON metadata for the struct
// [SettingSecurityHeaderGetResponse]
type settingSecurityHeaderGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone's security header.
type SettingSecurityHeaderGetResponseID string

const (
	SettingSecurityHeaderGetResponseIDSecurityHeader SettingSecurityHeaderGetResponseID = "security_header"
)

// Current value of the zone setting.
type SettingSecurityHeaderGetResponseValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SettingSecurityHeaderGetResponseValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    settingSecurityHeaderGetResponseValueJSON                    `json:"-"`
}

// settingSecurityHeaderGetResponseValueJSON contains the JSON metadata for the
// struct [SettingSecurityHeaderGetResponseValue]
type settingSecurityHeaderGetResponseValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type SettingSecurityHeaderGetResponseValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                             `json:"nosniff"`
	JSON    settingSecurityHeaderGetResponseValueStrictTransportSecurityJSON `json:"-"`
}

// settingSecurityHeaderGetResponseValueStrictTransportSecurityJSON contains the
// JSON metadata for the struct
// [SettingSecurityHeaderGetResponseValueStrictTransportSecurity]
type settingSecurityHeaderGetResponseValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSecurityHeaderGetResponseEditable bool

const (
	SettingSecurityHeaderGetResponseEditableTrue  SettingSecurityHeaderGetResponseEditable = true
	SettingSecurityHeaderGetResponseEditableFalse SettingSecurityHeaderGetResponseEditable = false
)

type SettingSecurityHeaderEditParams struct {
	Value param.Field[SettingSecurityHeaderEditParamsValue] `json:"value,required"`
}

func (r SettingSecurityHeaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSecurityHeaderEditParamsValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[SettingSecurityHeaderEditParamsValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r SettingSecurityHeaderEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type SettingSecurityHeaderEditParamsValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r SettingSecurityHeaderEditParamsValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSecurityHeaderEditResponseEnvelope struct {
	Errors   []SettingSecurityHeaderEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSecurityHeaderEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result SettingSecurityHeaderEditResponse             `json:"result"`
	JSON   settingSecurityHeaderEditResponseEnvelopeJSON `json:"-"`
}

// settingSecurityHeaderEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityHeaderEditResponseEnvelope]
type settingSecurityHeaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSecurityHeaderEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingSecurityHeaderEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSecurityHeaderEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderEditResponseEnvelopeErrors]
type settingSecurityHeaderEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSecurityHeaderEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingSecurityHeaderEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSecurityHeaderEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderEditResponseEnvelopeMessages]
type settingSecurityHeaderEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSecurityHeaderGetResponseEnvelope struct {
	Errors   []SettingSecurityHeaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSecurityHeaderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare security header for a zone.
	Result SettingSecurityHeaderGetResponse             `json:"result"`
	JSON   settingSecurityHeaderGetResponseEnvelopeJSON `json:"-"`
}

// settingSecurityHeaderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityHeaderGetResponseEnvelope]
type settingSecurityHeaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSecurityHeaderGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingSecurityHeaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSecurityHeaderGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderGetResponseEnvelopeErrors]
type settingSecurityHeaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSecurityHeaderGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingSecurityHeaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSecurityHeaderGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSecurityHeaderGetResponseEnvelopeMessages]
type settingSecurityHeaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityHeaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
