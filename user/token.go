// File generated from our OpenAPI spec by Stainless.

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// TokenService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTokenService] method instead.
type TokenService struct {
	Options          []option.RequestOption
	PermissionGroups *TokenPermissionGroupService
	Value            *TokenValueService
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	r.PermissionGroups = NewTokenPermissionGroupService(opts...)
	r.Value = NewTokenValueService(opts...)
	return
}

// Create a new access token.
func (r *TokenService) New(ctx context.Context, body TokenNewParams, opts ...option.RequestOption) (res *TokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenNewResponseEnvelope
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing token.
func (r *TokenService) Update(ctx context.Context, tokenID interface{}, body TokenUpdateParams, opts ...option.RequestOption) (res *TokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenUpdateResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all access tokens you created.
func (r *TokenService) List(ctx context.Context, query TokenListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[TokenListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/tokens"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all access tokens you created.
func (r *TokenService) ListAutoPaging(ctx context.Context, query TokenListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[TokenListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Destroy a token.
func (r *TokenService) Delete(ctx context.Context, tokenID interface{}, opts ...option.RequestOption) (res *TokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenDeleteResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific token.
func (r *TokenService) Get(ctx context.Context, tokenID interface{}, opts ...option.RequestOption) (res *TokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenGetResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test whether a token works.
func (r *TokenService) Verify(ctx context.Context, opts ...option.RequestOption) (res *TokenVerifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenVerifyResponseEnvelope
	path := "user/tokens/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenNewResponse struct {
	// The token value.
	Value IamValue             `json:"value"`
	JSON  tokenNewResponseJSON `json:"-"`
}

// tokenNewResponseJSON contains the JSON metadata for the struct
// [TokenNewResponse]
type tokenNewResponseJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [user.TokenUpdateResponseUnknown] or [shared.UnionString].
type TokenUpdateResponse interface {
	ImplementsUserTokenUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TokenListResponse = interface{}

type TokenDeleteResponse struct {
	// Identifier
	ID   string                  `json:"id,required"`
	JSON tokenDeleteResponseJSON `json:"-"`
}

// tokenDeleteResponseJSON contains the JSON metadata for the struct
// [TokenDeleteResponse]
type tokenDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [user.TokenGetResponseUnknown] or [shared.UnionString].
type TokenGetResponse interface {
	ImplementsUserTokenGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TokenVerifyResponse struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status TokenVerifyResponseStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time               `json:"not_before" format:"date-time"`
	JSON      tokenVerifyResponseJSON `json:"-"`
}

// tokenVerifyResponseJSON contains the JSON metadata for the struct
// [TokenVerifyResponse]
type tokenVerifyResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type TokenVerifyResponseStatus string

const (
	TokenVerifyResponseStatusActive   TokenVerifyResponseStatus = "active"
	TokenVerifyResponseStatusDisabled TokenVerifyResponseStatus = "disabled"
	TokenVerifyResponseStatusExpired  TokenVerifyResponseStatus = "expired"
)

type TokenNewParams struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies  param.Field[[]TokenNewParamsPolicy]  `json:"policies,required"`
	Condition param.Field[TokenNewParamsCondition] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r TokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewParamsPolicy struct {
	// Allow or deny operations against the resources.
	Effect param.Field[TokenNewParamsPoliciesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]TokenNewParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[interface{}] `json:"resources,required"`
}

func (r TokenNewParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type TokenNewParamsPoliciesEffect string

const (
	TokenNewParamsPoliciesEffectAllow TokenNewParamsPoliciesEffect = "allow"
	TokenNewParamsPoliciesEffectDeny  TokenNewParamsPoliciesEffect = "deny"
)

// A named group of permissions that map to a group of operations against
// resources.
type TokenNewParamsPoliciesPermissionGroup struct {
}

func (r TokenNewParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[TokenNewParamsConditionRequestIP] `json:"request_ip"`
}

func (r TokenNewParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type TokenNewParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]string] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]string] `json:"not_in"`
}

func (r TokenNewParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewResponseEnvelope struct {
	Errors   []TokenNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TokenNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TokenNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tokenNewResponseEnvelopeJSON    `json:"-"`
}

// tokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelope]
type tokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    tokenNewResponseEnvelopeErrorsJSON `json:"-"`
}

// tokenNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelopeErrors]
type tokenNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tokenNewResponseEnvelopeMessagesJSON `json:"-"`
}

// tokenNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelopeMessages]
type tokenNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenNewResponseEnvelopeSuccess bool

const (
	TokenNewResponseEnvelopeSuccessTrue TokenNewResponseEnvelopeSuccess = true
)

type TokenUpdateParams struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies param.Field[[]TokenUpdateParamsPolicy] `json:"policies,required"`
	// Status of the token.
	Status    param.Field[TokenUpdateParamsStatus]    `json:"status,required"`
	Condition param.Field[TokenUpdateParamsCondition] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r TokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenUpdateParamsPolicy struct {
	// Allow or deny operations against the resources.
	Effect param.Field[TokenUpdateParamsPoliciesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]TokenUpdateParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[interface{}] `json:"resources,required"`
}

func (r TokenUpdateParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type TokenUpdateParamsPoliciesEffect string

const (
	TokenUpdateParamsPoliciesEffectAllow TokenUpdateParamsPoliciesEffect = "allow"
	TokenUpdateParamsPoliciesEffectDeny  TokenUpdateParamsPoliciesEffect = "deny"
)

// A named group of permissions that map to a group of operations against
// resources.
type TokenUpdateParamsPoliciesPermissionGroup struct {
}

func (r TokenUpdateParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of the token.
type TokenUpdateParamsStatus string

const (
	TokenUpdateParamsStatusActive   TokenUpdateParamsStatus = "active"
	TokenUpdateParamsStatusDisabled TokenUpdateParamsStatus = "disabled"
	TokenUpdateParamsStatusExpired  TokenUpdateParamsStatus = "expired"
)

type TokenUpdateParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[TokenUpdateParamsConditionRequestIP] `json:"request_ip"`
}

func (r TokenUpdateParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type TokenUpdateParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]string] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]string] `json:"not_in"`
}

func (r TokenUpdateParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenUpdateResponseEnvelope struct {
	Errors   []TokenUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   TokenUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    tokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// tokenUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenUpdateResponseEnvelope]
type tokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    tokenUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// tokenUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenUpdateResponseEnvelopeErrors]
type tokenUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    tokenUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// tokenUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenUpdateResponseEnvelopeMessages]
type tokenUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenUpdateResponseEnvelopeSuccess bool

const (
	TokenUpdateResponseEnvelopeSuccessTrue TokenUpdateResponseEnvelopeSuccess = true
)

type TokenListParams struct {
	// Direction to order results.
	Direction param.Field[TokenListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [TokenListParams]'s query parameters as `url.Values`.
func (r TokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type TokenListParamsDirection string

const (
	TokenListParamsDirectionAsc  TokenListParamsDirection = "asc"
	TokenListParamsDirectionDesc TokenListParamsDirection = "desc"
)

type TokenDeleteResponseEnvelope struct {
	Errors   []TokenDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TokenDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success TokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// tokenDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenDeleteResponseEnvelope]
type tokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    tokenDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// tokenDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenDeleteResponseEnvelopeErrors]
type tokenDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    tokenDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// tokenDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenDeleteResponseEnvelopeMessages]
type tokenDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenDeleteResponseEnvelopeSuccess bool

const (
	TokenDeleteResponseEnvelopeSuccessTrue TokenDeleteResponseEnvelopeSuccess = true
)

type TokenGetResponseEnvelope struct {
	Errors   []TokenGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TokenGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TokenGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tokenGetResponseEnvelopeJSON    `json:"-"`
}

// tokenGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenGetResponseEnvelope]
type tokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenGetResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    tokenGetResponseEnvelopeErrorsJSON `json:"-"`
}

// tokenGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenGetResponseEnvelopeErrors]
type tokenGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenGetResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    tokenGetResponseEnvelopeMessagesJSON `json:"-"`
}

// tokenGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TokenGetResponseEnvelopeMessages]
type tokenGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenGetResponseEnvelopeSuccess bool

const (
	TokenGetResponseEnvelopeSuccessTrue TokenGetResponseEnvelopeSuccess = true
)

type TokenVerifyResponseEnvelope struct {
	Errors   []TokenVerifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenVerifyResponseEnvelopeMessages `json:"messages,required"`
	Result   TokenVerifyResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TokenVerifyResponseEnvelopeSuccess `json:"success,required"`
	JSON    tokenVerifyResponseEnvelopeJSON    `json:"-"`
}

// tokenVerifyResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenVerifyResponseEnvelope]
type tokenVerifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    tokenVerifyResponseEnvelopeErrorsJSON `json:"-"`
}

// tokenVerifyResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenVerifyResponseEnvelopeErrors]
type tokenVerifyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    tokenVerifyResponseEnvelopeMessagesJSON `json:"-"`
}

// tokenVerifyResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenVerifyResponseEnvelopeMessages]
type tokenVerifyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenVerifyResponseEnvelopeSuccess bool

const (
	TokenVerifyResponseEnvelopeSuccessTrue TokenVerifyResponseEnvelopeSuccess = true
)
