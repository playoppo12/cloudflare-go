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

// IntelIndicatorFeedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelIndicatorFeedService] method
// instead.
type IntelIndicatorFeedService struct {
	Options []option.RequestOption
}

// NewIntelIndicatorFeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntelIndicatorFeedService(opts ...option.RequestOption) (r *IntelIndicatorFeedService) {
	r = &IntelIndicatorFeedService{}
	r.Options = opts
	return
}

// Create new indicator feed
func (r *IntelIndicatorFeedService) New(ctx context.Context, accountIdentifier string, body IntelIndicatorFeedNewParams, opts ...option.RequestOption) (res *IntelIndicatorFeedNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get indicator feed metadata
func (r *IntelIndicatorFeedService) Get(ctx context.Context, accountIdentifier string, feedID int64, opts ...option.RequestOption) (res *IntelIndicatorFeedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", accountIdentifier, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feeds owned by this account
func (r *IntelIndicatorFeedService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *IntelIndicatorFeedListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feed data
func (r *IntelIndicatorFeedService) Data(ctx context.Context, accountIdentifier string, feedID int64, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/data", accountIdentifier, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Grant permission to indicator feed
func (r *IntelIndicatorFeedService) PermissionsAdd(ctx context.Context, accountIdentifier string, body IntelIndicatorFeedPermissionsAddParams, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionsAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Revoke permission to indicator feed
func (r *IntelIndicatorFeedService) PermissionsRemove(ctx context.Context, accountIdentifier string, body IntelIndicatorFeedPermissionsRemoveParams, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionsRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List indicator feed permissions
func (r *IntelIndicatorFeedService) PermissionsView(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionsViewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update indicator feed data
func (r *IntelIndicatorFeedService) Snapshot(ctx context.Context, accountIdentifier string, feedID int64, body IntelIndicatorFeedSnapshotParams, opts ...option.RequestOption) (res *IntelIndicatorFeedSnapshotResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/snapshot", accountIdentifier, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type IntelIndicatorFeedNewResponse struct {
	Errors   []IntelIndicatorFeedNewResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedNewResponseMessage `json:"messages"`
	Result   IntelIndicatorFeedNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedNewResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedNewResponseJSON    `json:"-"`
}

// intelIndicatorFeedNewResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedNewResponse]
type intelIndicatorFeedNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedNewResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    intelIndicatorFeedNewResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedNewResponseErrorJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedNewResponseError]
type intelIndicatorFeedNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedNewResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    intelIndicatorFeedNewResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedNewResponseMessageJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedNewResponseMessage]
type intelIndicatorFeedNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedNewResponseResult struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                                  `json:"name"`
	JSON intelIndicatorFeedNewResponseResultJSON `json:"-"`
}

// intelIndicatorFeedNewResponseResultJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedNewResponseResult]
type intelIndicatorFeedNewResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedNewResponseSuccess bool

const (
	IntelIndicatorFeedNewResponseSuccessTrue IntelIndicatorFeedNewResponseSuccess = true
)

type IntelIndicatorFeedGetResponse struct {
	Errors   []IntelIndicatorFeedGetResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedGetResponseMessage `json:"messages"`
	Result   IntelIndicatorFeedGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedGetResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedGetResponseJSON    `json:"-"`
}

// intelIndicatorFeedGetResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedGetResponse]
type intelIndicatorFeedGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    intelIndicatorFeedGetResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedGetResponseErrorJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedGetResponseError]
type intelIndicatorFeedGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    intelIndicatorFeedGetResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedGetResponseMessageJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedGetResponseMessage]
type intelIndicatorFeedGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedGetResponseResult struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Status of the latest snapshot uploaded
	LatestUploadStatus IntelIndicatorFeedGetResponseResultLatestUploadStatus `json:"latest_upload_status"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                                  `json:"name"`
	JSON intelIndicatorFeedGetResponseResultJSON `json:"-"`
}

// intelIndicatorFeedGetResponseResultJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedGetResponseResult]
type intelIndicatorFeedGetResponseResultJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	LatestUploadStatus apijson.Field
	ModifiedOn         apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IntelIndicatorFeedGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the latest snapshot uploaded
type IntelIndicatorFeedGetResponseResultLatestUploadStatus string

const (
	IntelIndicatorFeedGetResponseResultLatestUploadStatusMirroring    IntelIndicatorFeedGetResponseResultLatestUploadStatus = "Mirroring"
	IntelIndicatorFeedGetResponseResultLatestUploadStatusUnifying     IntelIndicatorFeedGetResponseResultLatestUploadStatus = "Unifying"
	IntelIndicatorFeedGetResponseResultLatestUploadStatusLoading      IntelIndicatorFeedGetResponseResultLatestUploadStatus = "Loading"
	IntelIndicatorFeedGetResponseResultLatestUploadStatusProvisioning IntelIndicatorFeedGetResponseResultLatestUploadStatus = "Provisioning"
	IntelIndicatorFeedGetResponseResultLatestUploadStatusComplete     IntelIndicatorFeedGetResponseResultLatestUploadStatus = "Complete"
	IntelIndicatorFeedGetResponseResultLatestUploadStatusError        IntelIndicatorFeedGetResponseResultLatestUploadStatus = "Error"
)

// Whether the API call was successful
type IntelIndicatorFeedGetResponseSuccess bool

const (
	IntelIndicatorFeedGetResponseSuccessTrue IntelIndicatorFeedGetResponseSuccess = true
)

type IntelIndicatorFeedListResponse struct {
	Errors   []IntelIndicatorFeedListResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedListResponseMessage `json:"messages"`
	Result   []IntelIndicatorFeedListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedListResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedListResponseJSON    `json:"-"`
}

// intelIndicatorFeedListResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedListResponse]
type intelIndicatorFeedListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedListResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    intelIndicatorFeedListResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedListResponseErrorJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedListResponseError]
type intelIndicatorFeedListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedListResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    intelIndicatorFeedListResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedListResponseMessageJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedListResponseMessage]
type intelIndicatorFeedListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedListResponseResult struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                                   `json:"name"`
	JSON intelIndicatorFeedListResponseResultJSON `json:"-"`
}

// intelIndicatorFeedListResponseResultJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedListResponseResult]
type intelIndicatorFeedListResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedListResponseSuccess bool

const (
	IntelIndicatorFeedListResponseSuccessTrue IntelIndicatorFeedListResponseSuccess = true
)

type IntelIndicatorFeedPermissionsAddResponse struct {
	Errors   []IntelIndicatorFeedPermissionsAddResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedPermissionsAddResponseMessage `json:"messages"`
	Result   IntelIndicatorFeedPermissionsAddResponseResult    `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionsAddResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedPermissionsAddResponseJSON    `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedPermissionsAddResponse]
type intelIndicatorFeedPermissionsAddResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsAddResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsAddResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseErrorJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedPermissionsAddResponseError]
type intelIndicatorFeedPermissionsAddResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsAddResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsAddResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseMessageJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsAddResponseMessage]
type intelIndicatorFeedPermissionsAddResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsAddResponseResult struct {
	// Whether the update succeeded or not
	Success bool                                               `json:"success"`
	JSON    intelIndicatorFeedPermissionsAddResponseResultJSON `json:"-"`
}

// intelIndicatorFeedPermissionsAddResponseResultJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsAddResponseResult]
type intelIndicatorFeedPermissionsAddResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsAddResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionsAddResponseSuccess bool

const (
	IntelIndicatorFeedPermissionsAddResponseSuccessTrue IntelIndicatorFeedPermissionsAddResponseSuccess = true
)

type IntelIndicatorFeedPermissionsRemoveResponse struct {
	Errors   []IntelIndicatorFeedPermissionsRemoveResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedPermissionsRemoveResponseMessage `json:"messages"`
	Result   IntelIndicatorFeedPermissionsRemoveResponseResult    `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionsRemoveResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseJSON    `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedPermissionsRemoveResponse]
type intelIndicatorFeedPermissionsRemoveResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsRemoveResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseErrorJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsRemoveResponseError]
type intelIndicatorFeedPermissionsRemoveResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsRemoveResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseMessageJSON contains the JSON
// metadata for the struct [IntelIndicatorFeedPermissionsRemoveResponseMessage]
type intelIndicatorFeedPermissionsRemoveResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsRemoveResponseResult struct {
	// Whether the update succeeded or not
	Success bool                                                  `json:"success"`
	JSON    intelIndicatorFeedPermissionsRemoveResponseResultJSON `json:"-"`
}

// intelIndicatorFeedPermissionsRemoveResponseResultJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsRemoveResponseResult]
type intelIndicatorFeedPermissionsRemoveResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsRemoveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionsRemoveResponseSuccess bool

const (
	IntelIndicatorFeedPermissionsRemoveResponseSuccessTrue IntelIndicatorFeedPermissionsRemoveResponseSuccess = true
)

type IntelIndicatorFeedPermissionsViewResponse struct {
	Errors   []IntelIndicatorFeedPermissionsViewResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedPermissionsViewResponseMessage `json:"messages"`
	Result   []IntelIndicatorFeedPermissionsViewResponseResult  `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionsViewResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedPermissionsViewResponseJSON    `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedPermissionsViewResponse]
type intelIndicatorFeedPermissionsViewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsViewResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsViewResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseErrorJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsViewResponseError]
type intelIndicatorFeedPermissionsViewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsViewResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    intelIndicatorFeedPermissionsViewResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseMessageJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsViewResponseMessage]
type intelIndicatorFeedPermissionsViewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionsViewResponseResult struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// The name of the indicator feed
	Name string                                              `json:"name"`
	JSON intelIndicatorFeedPermissionsViewResponseResultJSON `json:"-"`
}

// intelIndicatorFeedPermissionsViewResponseResultJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionsViewResponseResult]
type intelIndicatorFeedPermissionsViewResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionsViewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionsViewResponseSuccess bool

const (
	IntelIndicatorFeedPermissionsViewResponseSuccessTrue IntelIndicatorFeedPermissionsViewResponseSuccess = true
)

type IntelIndicatorFeedSnapshotResponse struct {
	Errors   []IntelIndicatorFeedSnapshotResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedSnapshotResponseMessage `json:"messages"`
	Result   IntelIndicatorFeedSnapshotResponseResult    `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedSnapshotResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedSnapshotResponseJSON    `json:"-"`
}

// intelIndicatorFeedSnapshotResponseJSON contains the JSON metadata for the struct
// [IntelIndicatorFeedSnapshotResponse]
type intelIndicatorFeedSnapshotResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedSnapshotResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    intelIndicatorFeedSnapshotResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedSnapshotResponseErrorJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedSnapshotResponseError]
type intelIndicatorFeedSnapshotResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedSnapshotResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    intelIndicatorFeedSnapshotResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedSnapshotResponseMessageJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedSnapshotResponseMessage]
type intelIndicatorFeedSnapshotResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedSnapshotResponseResult struct {
	// Feed id
	FileID int64 `json:"file_id"`
	// Name of the file unified in our system
	Filename string `json:"filename"`
	// Current status of upload, should be unified
	Status string                                       `json:"status"`
	JSON   intelIndicatorFeedSnapshotResponseResultJSON `json:"-"`
}

// intelIndicatorFeedSnapshotResponseResultJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedSnapshotResponseResult]
type intelIndicatorFeedSnapshotResponseResultJSON struct {
	FileID      apijson.Field
	Filename    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedSnapshotResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedSnapshotResponseSuccess bool

const (
	IntelIndicatorFeedSnapshotResponseSuccessTrue IntelIndicatorFeedSnapshotResponseSuccess = true
)

type IntelIndicatorFeedNewParams struct {
	// The description of the example test
	Description param.Field[string] `json:"description"`
	// The name of the indicator feed
	Name param.Field[string] `json:"name"`
}

func (r IntelIndicatorFeedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedPermissionsAddParams struct {
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IntelIndicatorFeedPermissionsAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedPermissionsRemoveParams struct {
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IntelIndicatorFeedPermissionsRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedSnapshotParams struct {
	// The file to upload
	Source param.Field[string] `json:"source"`
}

func (r IntelIndicatorFeedSnapshotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
