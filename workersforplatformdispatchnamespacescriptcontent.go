// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiform"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkersForPlatformDispatchNamespaceScriptContentService contains methods and
// other services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformDispatchNamespaceScriptContentService] method instead.
type WorkersForPlatformDispatchNamespaceScriptContentService struct {
	Options []option.RequestOption
}

// NewWorkersForPlatformDispatchNamespaceScriptContentService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewWorkersForPlatformDispatchNamespaceScriptContentService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceScriptContentService) {
	r = &WorkersForPlatformDispatchNamespaceScriptContentService{}
	r.Options = opts
	return
}

// Fetch script content from a script uploaded to a Workers for Platforms
// namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptContentService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put script content for a script uploaded to a Workers for Platforms namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptContentService) Update(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, params WorkersForPlatformDispatchNamespaceScriptContentUpdateParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseError   `json:"errors"`
	Messages []WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseMessage `json:"messages"`
	Result   WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseSuccess `json:"success"`
	JSON    workersForPlatformDispatchNamespaceScriptContentUpdateResponseJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentUpdateResponseJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse]
type workersForPlatformDispatchNamespaceScriptContentUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseError struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentUpdateResponseErrorJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentUpdateResponseErrorJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseError]
type workersForPlatformDispatchNamespaceScriptContentUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseMessage struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentUpdateResponseMessageJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentUpdateResponseMessageJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseMessage]
type workersForPlatformDispatchNamespaceScriptContentUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResult struct {
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Deprecated. Deployment metadata for internal usage.
	PipelineHash string `json:"pipeline_hash"`
	// Specifies the placement mode for the Worker (e.g. 'smart').
	PlacementMode string `json:"placement_mode"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                                   `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptContentUpdateResponseResultJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentUpdateResponseResultJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResult]
type workersForPlatformDispatchNamespaceScriptContentUpdateResponseResultJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Etag          apijson.Field
	Logpush       apijson.Field
	ModifiedOn    apijson.Field
	PipelineHash  apijson.Field
	PlacementMode apijson.Field
	TailConsumers apijson.Field
	UsageModel    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                               `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptContentUpdateResponseResultTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentUpdateResponseResultTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResultTailConsumer]
type workersForPlatformDispatchNamespaceScriptContentUpdateResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseSuccessTrue WorkersForPlatformDispatchNamespaceScriptContentUpdateResponseSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptContentUpdateParams struct {
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be
	// present. This should be referenced either in the metadata as `main_module`
	// (esm)/`body_part` (service worker) or as a header `CF-WORKER-MAIN-MODULE-PART`
	// (esm) /`CF-WORKER-BODY-PART` (service worker) by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[WorkersForPlatformDispatchNamespaceScriptContentUpdateParamsMetadata] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                                                               `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                                               `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type WorkersForPlatformDispatchNamespaceScriptContentUpdateParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
