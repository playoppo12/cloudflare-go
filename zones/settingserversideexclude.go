// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// SettingServerSideExcludeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingServerSideExcludeService] method instead.
type SettingServerSideExcludeService struct {
	Options []option.RequestOption
}

// NewSettingServerSideExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingServerSideExcludeService(opts ...option.RequestOption) (r *SettingServerSideExcludeService) {
	r = &SettingServerSideExcludeService{}
	r.Options = opts
	return
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
func (r *SettingServerSideExcludeService) Edit(ctx context.Context, params SettingServerSideExcludeEditParams, opts ...option.RequestOption) (res *ZonesServerSideExclude, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingServerSideExcludeEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
func (r *SettingServerSideExcludeService) Get(ctx context.Context, query SettingServerSideExcludeGetParams, opts ...option.RequestOption) (res *ZonesServerSideExclude, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingServerSideExcludeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ZonesServerSideExclude struct {
	// ID of the zone setting.
	ID ZonesServerSideExcludeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesServerSideExcludeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesServerSideExcludeJSON `json:"-"`
}

// zonesServerSideExcludeJSON contains the JSON metadata for the struct
// [ZonesServerSideExclude]
type zonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesServerSideExcludeJSON) RawJSON() string {
	return r.raw
}

func (r ZonesServerSideExclude) implementsZonesSettingEditResponse() {}

func (r ZonesServerSideExclude) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesServerSideExcludeID string

const (
	ZonesServerSideExcludeIDServerSideExclude ZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type ZonesServerSideExcludeValue string

const (
	ZonesServerSideExcludeValueOn  ZonesServerSideExcludeValue = "on"
	ZonesServerSideExcludeValueOff ZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesServerSideExcludeEditable bool

const (
	ZonesServerSideExcludeEditableTrue  ZonesServerSideExcludeEditable = true
	ZonesServerSideExcludeEditableFalse ZonesServerSideExcludeEditable = false
)

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ZonesServerSideExcludeParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesServerSideExcludeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesServerSideExcludeValue] `json:"value,required"`
}

func (r ZonesServerSideExcludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesServerSideExcludeParam) implementsZonesSettingEditParamsItem() {}

type SettingServerSideExcludeEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingServerSideExcludeEditParamsValue] `json:"value,required"`
}

func (r SettingServerSideExcludeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingServerSideExcludeEditParamsValue string

const (
	SettingServerSideExcludeEditParamsValueOn  SettingServerSideExcludeEditParamsValue = "on"
	SettingServerSideExcludeEditParamsValueOff SettingServerSideExcludeEditParamsValue = "off"
)

type SettingServerSideExcludeEditResponseEnvelope struct {
	Errors   []SettingServerSideExcludeEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingServerSideExcludeEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// If there is sensitive content on your website that you want visible to real
	// visitors, but that you want to hide from suspicious visitors, all you have to do
	// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
	// be excluded from suspicious visitors in the following SSE tags:
	// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
	// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
	// have HTML minification enabled, you won't see the SSE tags in your HTML source
	// when it's served through Cloudflare. SSE will still function in this case, as
	// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
	// resource moves through our network to the visitor's computer.
	// (https://support.cloudflare.com/hc/en-us/articles/200170036).
	Result ZonesServerSideExclude                           `json:"result"`
	JSON   settingServerSideExcludeEditResponseEnvelopeJSON `json:"-"`
}

// settingServerSideExcludeEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingServerSideExcludeEditResponseEnvelope]
type settingServerSideExcludeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingServerSideExcludeEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingServerSideExcludeEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingServerSideExcludeEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingServerSideExcludeEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingServerSideExcludeEditResponseEnvelopeErrors]
type settingServerSideExcludeEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingServerSideExcludeEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingServerSideExcludeEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingServerSideExcludeEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingServerSideExcludeEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingServerSideExcludeEditResponseEnvelopeMessages]
type settingServerSideExcludeEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingServerSideExcludeEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingServerSideExcludeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingServerSideExcludeGetResponseEnvelope struct {
	Errors   []SettingServerSideExcludeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingServerSideExcludeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// If there is sensitive content on your website that you want visible to real
	// visitors, but that you want to hide from suspicious visitors, all you have to do
	// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
	// be excluded from suspicious visitors in the following SSE tags:
	// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
	// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
	// have HTML minification enabled, you won't see the SSE tags in your HTML source
	// when it's served through Cloudflare. SSE will still function in this case, as
	// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
	// resource moves through our network to the visitor's computer.
	// (https://support.cloudflare.com/hc/en-us/articles/200170036).
	Result ZonesServerSideExclude                          `json:"result"`
	JSON   settingServerSideExcludeGetResponseEnvelopeJSON `json:"-"`
}

// settingServerSideExcludeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingServerSideExcludeGetResponseEnvelope]
type settingServerSideExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingServerSideExcludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingServerSideExcludeGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingServerSideExcludeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingServerSideExcludeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingServerSideExcludeGetResponseEnvelopeErrors]
type settingServerSideExcludeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingServerSideExcludeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingServerSideExcludeGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingServerSideExcludeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingServerSideExcludeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingServerSideExcludeGetResponseEnvelopeMessages]
type settingServerSideExcludeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingServerSideExcludeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
