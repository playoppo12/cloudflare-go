// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AIService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAIService] method instead.
type AIService struct {
	Options             []option.RequestOption
	ImageClassification *AIImageClassificationService
	SpeechRecognition   *AISpeechRecognitionService
	TextClassification  *AITextClassificationService
	TextEmbeddings      *AITextEmbeddingService
	TextGeneration      *AITextGenerationService
	Translation         *AITranslationService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r *AIService) {
	r = &AIService{}
	r.Options = opts
	r.ImageClassification = NewAIImageClassificationService(opts...)
	r.SpeechRecognition = NewAISpeechRecognitionService(opts...)
	r.TextClassification = NewAITextClassificationService(opts...)
	r.TextEmbeddings = NewAITextEmbeddingService(opts...)
	r.TextGeneration = NewAITextGenerationService(opts...)
	r.Translation = NewAITranslationService(opts...)
	return
}

// This endpoint provides users with the capability to run specific AI models
// on-demand.
//
// By submitting the required input data, users can receive real-time predictions
// or results generated by the chosen AI model. The endpoint supports various AI
// model types, ensuring flexibility and adaptability for diverse use cases.
func (r *AIService) Run(ctx context.Context, accountIdentifier string, modelName string, body AIRunParams, opts ...option.RequestOption) (res *AIRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/%s", accountIdentifier, modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIRunResponse struct {
	Errors   []AIRunResponseError `json:"errors,required"`
	Messages []string             `json:"messages,required"`
	Result   interface{}          `json:"result,required"`
	Success  bool                 `json:"success,required"`
	JSON     aiRunResponseJSON    `json:"-"`
}

// aiRunResponseJSON contains the JSON metadata for the struct [AIRunResponse]
type aiRunResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIRunResponseError struct {
	Message string                 `json:"message,required"`
	JSON    aiRunResponseErrorJSON `json:"-"`
}

// aiRunResponseErrorJSON contains the JSON metadata for the struct
// [AIRunResponseError]
type aiRunResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIRunParams struct {
}
