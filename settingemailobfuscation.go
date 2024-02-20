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

// SettingEmailObfuscationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingEmailObfuscationService] method instead.
type SettingEmailObfuscationService struct {
	Options []option.RequestOption
}

// NewSettingEmailObfuscationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingEmailObfuscationService(opts ...option.RequestOption) (r *SettingEmailObfuscationService) {
	r = &SettingEmailObfuscationService{}
	r.Options = opts
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *SettingEmailObfuscationService) Edit(ctx context.Context, zoneID string, body SettingEmailObfuscationEditParams, opts ...option.RequestOption) (res *SettingEmailObfuscationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEmailObfuscationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *SettingEmailObfuscationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingEmailObfuscationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEmailObfuscationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingEmailObfuscationEditResponse struct {
	// ID of the zone setting.
	ID SettingEmailObfuscationEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEmailObfuscationEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEmailObfuscationEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEmailObfuscationEditResponseJSON `json:"-"`
}

// settingEmailObfuscationEditResponseJSON contains the JSON metadata for the
// struct [SettingEmailObfuscationEditResponse]
type settingEmailObfuscationEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingEmailObfuscationEditResponseID string

const (
	SettingEmailObfuscationEditResponseIDEmailObfuscation SettingEmailObfuscationEditResponseID = "email_obfuscation"
)

// Current value of the zone setting.
type SettingEmailObfuscationEditResponseValue string

const (
	SettingEmailObfuscationEditResponseValueOn  SettingEmailObfuscationEditResponseValue = "on"
	SettingEmailObfuscationEditResponseValueOff SettingEmailObfuscationEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEmailObfuscationEditResponseEditable bool

const (
	SettingEmailObfuscationEditResponseEditableTrue  SettingEmailObfuscationEditResponseEditable = true
	SettingEmailObfuscationEditResponseEditableFalse SettingEmailObfuscationEditResponseEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingEmailObfuscationGetResponse struct {
	// ID of the zone setting.
	ID SettingEmailObfuscationGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEmailObfuscationGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEmailObfuscationGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEmailObfuscationGetResponseJSON `json:"-"`
}

// settingEmailObfuscationGetResponseJSON contains the JSON metadata for the struct
// [SettingEmailObfuscationGetResponse]
type settingEmailObfuscationGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingEmailObfuscationGetResponseID string

const (
	SettingEmailObfuscationGetResponseIDEmailObfuscation SettingEmailObfuscationGetResponseID = "email_obfuscation"
)

// Current value of the zone setting.
type SettingEmailObfuscationGetResponseValue string

const (
	SettingEmailObfuscationGetResponseValueOn  SettingEmailObfuscationGetResponseValue = "on"
	SettingEmailObfuscationGetResponseValueOff SettingEmailObfuscationGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEmailObfuscationGetResponseEditable bool

const (
	SettingEmailObfuscationGetResponseEditableTrue  SettingEmailObfuscationGetResponseEditable = true
	SettingEmailObfuscationGetResponseEditableFalse SettingEmailObfuscationGetResponseEditable = false
)

type SettingEmailObfuscationEditParams struct {
	// Value of the zone setting.
	Value param.Field[SettingEmailObfuscationEditParamsValue] `json:"value,required"`
}

func (r SettingEmailObfuscationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingEmailObfuscationEditParamsValue string

const (
	SettingEmailObfuscationEditParamsValueOn  SettingEmailObfuscationEditParamsValue = "on"
	SettingEmailObfuscationEditParamsValueOff SettingEmailObfuscationEditParamsValue = "off"
)

type SettingEmailObfuscationEditResponseEnvelope struct {
	Errors   []SettingEmailObfuscationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEmailObfuscationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result SettingEmailObfuscationEditResponse             `json:"result"`
	JSON   settingEmailObfuscationEditResponseEnvelopeJSON `json:"-"`
}

// settingEmailObfuscationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingEmailObfuscationEditResponseEnvelope]
type settingEmailObfuscationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEmailObfuscationEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingEmailObfuscationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEmailObfuscationEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingEmailObfuscationEditResponseEnvelopeErrors]
type settingEmailObfuscationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEmailObfuscationEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingEmailObfuscationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEmailObfuscationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingEmailObfuscationEditResponseEnvelopeMessages]
type settingEmailObfuscationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEmailObfuscationGetResponseEnvelope struct {
	Errors   []SettingEmailObfuscationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEmailObfuscationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result SettingEmailObfuscationGetResponse             `json:"result"`
	JSON   settingEmailObfuscationGetResponseEnvelopeJSON `json:"-"`
}

// settingEmailObfuscationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingEmailObfuscationGetResponseEnvelope]
type settingEmailObfuscationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEmailObfuscationGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingEmailObfuscationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEmailObfuscationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingEmailObfuscationGetResponseEnvelopeErrors]
type settingEmailObfuscationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEmailObfuscationGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingEmailObfuscationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEmailObfuscationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingEmailObfuscationGetResponseEnvelopeMessages]
type settingEmailObfuscationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
