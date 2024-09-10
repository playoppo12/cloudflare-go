// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/shared"
	"github.com/tidwall/gjson"
)

// AIService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options []option.RequestOption
	Models  *AIModelService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r *AIService) {
	r = &AIService{}
	r.Options = opts
	r.Models = NewAIModelService(opts...)
	return
}

// This endpoint provides users with the capability to run specific AI models
// on-demand.
//
// By submitting the required input data, users can receive real-time predictions
// or results generated by the chosen AI model. The endpoint supports various AI
// model types, ensuring flexibility and adaptability for diverse use cases.
//
// Model specific inputs available in
// [Cloudflare Docs](https://developers.cloudflare.com/workers-ai/models/).
func (r *AIService) Run(ctx context.Context, modelName string, params AIRunParams, opts ...option.RequestOption) (res *AIRunResponseUnion, err error) {
	var env AIRunResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if modelName == "" {
		err = errors.New("missing required model_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/%s", params.AccountID, modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [workers.AIRunResponseTextClassification],
// [shared.UnionString], [workers.AIRunResponseTextEmbeddings],
// [workers.AIRunResponseAutomaticSpeechRecognition],
// [workers.AIRunResponseImageClassification],
// [workers.AIRunResponseObjectDetection], [workers.AIRunResponseObject],
// [workers.AIRunResponseTranslation], [workers.AIRunResponseSummarization] or
// [workers.AIRunResponseImageToText].
type AIRunResponseUnion interface {
	ImplementsWorkersAIRunResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTextClassification{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTextEmbeddings{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseAutomaticSpeechRecognition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseImageClassification{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseObjectDetection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTranslation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseSummarization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseImageToText{}),
		},
	)
}

type AIRunResponseTextClassification []AIRunResponseTextClassificationItem

func (r AIRunResponseTextClassification) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseTextClassificationItem struct {
	Label string                                  `json:"label"`
	Score float64                                 `json:"score"`
	JSON  aiRunResponseTextClassificationItemJSON `json:"-"`
}

// aiRunResponseTextClassificationItemJSON contains the JSON metadata for the
// struct [AIRunResponseTextClassificationItem]
type aiRunResponseTextClassificationItemJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseTextClassificationItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTextClassificationItemJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseTextEmbeddings struct {
	Data  [][]float64                     `json:"data"`
	Shape []float64                       `json:"shape"`
	JSON  aiRunResponseTextEmbeddingsJSON `json:"-"`
}

// aiRunResponseTextEmbeddingsJSON contains the JSON metadata for the struct
// [AIRunResponseTextEmbeddings]
type aiRunResponseTextEmbeddingsJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseTextEmbeddings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTextEmbeddingsJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseTextEmbeddings) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseAutomaticSpeechRecognition struct {
	Text      string                                        `json:"text,required"`
	Vtt       string                                        `json:"vtt"`
	WordCount float64                                       `json:"word_count"`
	Words     []AIRunResponseAutomaticSpeechRecognitionWord `json:"words"`
	JSON      aiRunResponseAutomaticSpeechRecognitionJSON   `json:"-"`
}

// aiRunResponseAutomaticSpeechRecognitionJSON contains the JSON metadata for the
// struct [AIRunResponseAutomaticSpeechRecognition]
type aiRunResponseAutomaticSpeechRecognitionJSON struct {
	Text        apijson.Field
	Vtt         apijson.Field
	WordCount   apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseAutomaticSpeechRecognition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseAutomaticSpeechRecognitionJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseAutomaticSpeechRecognition) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseAutomaticSpeechRecognitionWord struct {
	End   float64                                         `json:"end"`
	Start float64                                         `json:"start"`
	Word  string                                          `json:"word"`
	JSON  aiRunResponseAutomaticSpeechRecognitionWordJSON `json:"-"`
}

// aiRunResponseAutomaticSpeechRecognitionWordJSON contains the JSON metadata for
// the struct [AIRunResponseAutomaticSpeechRecognitionWord]
type aiRunResponseAutomaticSpeechRecognitionWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseAutomaticSpeechRecognitionWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseAutomaticSpeechRecognitionWordJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseImageClassification []AIRunResponseImageClassificationItem

func (r AIRunResponseImageClassification) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseImageClassificationItem struct {
	Label string                                   `json:"label"`
	Score float64                                  `json:"score"`
	JSON  aiRunResponseImageClassificationItemJSON `json:"-"`
}

// aiRunResponseImageClassificationItemJSON contains the JSON metadata for the
// struct [AIRunResponseImageClassificationItem]
type aiRunResponseImageClassificationItemJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseImageClassificationItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseImageClassificationItemJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseObjectDetection []AIRunResponseObjectDetectionItem

func (r AIRunResponseObjectDetection) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseObjectDetectionItem struct {
	Box   AIRunResponseObjectDetectionBox      `json:"box"`
	Label string                               `json:"label"`
	Score float64                              `json:"score"`
	JSON  aiRunResponseObjectDetectionItemJSON `json:"-"`
}

// aiRunResponseObjectDetectionItemJSON contains the JSON metadata for the struct
// [AIRunResponseObjectDetectionItem]
type aiRunResponseObjectDetectionItemJSON struct {
	Box         apijson.Field
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObjectDetectionItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectDetectionItemJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseObjectDetectionBox struct {
	Xmax float64                             `json:"xmax"`
	Xmin float64                             `json:"xmin"`
	Ymax float64                             `json:"ymax"`
	Ymin float64                             `json:"ymin"`
	JSON aiRunResponseObjectDetectionBoxJSON `json:"-"`
}

// aiRunResponseObjectDetectionBoxJSON contains the JSON metadata for the struct
// [AIRunResponseObjectDetectionBox]
type aiRunResponseObjectDetectionBoxJSON struct {
	Xmax        apijson.Field
	Xmin        apijson.Field
	Ymax        apijson.Field
	Ymin        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObjectDetectionBox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectDetectionBoxJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseObject struct {
	Response  string                        `json:"response"`
	ToolCalls []AIRunResponseObjectToolCall `json:"tool_calls"`
	JSON      aiRunResponseObjectJSON       `json:"-"`
}

// aiRunResponseObjectJSON contains the JSON metadata for the struct
// [AIRunResponseObject]
type aiRunResponseObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseObject) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseObjectToolCall struct {
	Arguments interface{}                     `json:"arguments"`
	Name      string                          `json:"name"`
	JSON      aiRunResponseObjectToolCallJSON `json:"-"`
}

// aiRunResponseObjectToolCallJSON contains the JSON metadata for the struct
// [AIRunResponseObjectToolCall]
type aiRunResponseObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectToolCallJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseTranslation struct {
	TranslatedText string                       `json:"translated_text"`
	JSON           aiRunResponseTranslationJSON `json:"-"`
}

// aiRunResponseTranslationJSON contains the JSON metadata for the struct
// [AIRunResponseTranslation]
type aiRunResponseTranslationJSON struct {
	TranslatedText apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIRunResponseTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTranslationJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseTranslation) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseSummarization struct {
	Summary string                         `json:"summary"`
	JSON    aiRunResponseSummarizationJSON `json:"-"`
}

// aiRunResponseSummarizationJSON contains the JSON metadata for the struct
// [AIRunResponseSummarization]
type aiRunResponseSummarizationJSON struct {
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseSummarization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseSummarizationJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseSummarization) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseImageToText struct {
	Description string                       `json:"description"`
	JSON        aiRunResponseImageToTextJSON `json:"-"`
}

// aiRunResponseImageToTextJSON contains the JSON metadata for the struct
// [AIRunResponseImageToText]
type aiRunResponseImageToTextJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseImageToText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseImageToTextJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseImageToText) ImplementsWorkersAIRunResponseUnion() {}

type AIRunParams struct {
	AccountID param.Field[string]  `path:"account_id,required"`
	Body      AIRunParamsBodyUnion `json:"body,required"`
}

func (r AIRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AIRunParamsBody struct {
	Text              param.Field[interface{}] `json:"text,required"`
	Guidance          param.Field[float64]     `json:"guidance"`
	Height            param.Field[int64]       `json:"height"`
	Image             param.Field[interface{}] `json:"image,required"`
	ImageB64          param.Field[string]      `json:"image_b64"`
	Mask              param.Field[interface{}] `json:"mask,required"`
	NegativePrompt    param.Field[string]      `json:"negative_prompt"`
	NumSteps          param.Field[int64]       `json:"num_steps"`
	Prompt            param.Field[string]      `json:"prompt"`
	Seed              param.Field[int64]       `json:"seed"`
	Strength          param.Field[float64]     `json:"strength"`
	Width             param.Field[int64]       `json:"width"`
	Audio             param.Field[interface{}] `json:"audio,required"`
	SourceLang        param.Field[string]      `json:"source_lang"`
	TargetLang        param.Field[string]      `json:"target_lang"`
	FrequencyPenalty  param.Field[float64]     `json:"frequency_penalty"`
	Lora              param.Field[string]      `json:"lora"`
	MaxTokens         param.Field[int64]       `json:"max_tokens"`
	PresencePenalty   param.Field[float64]     `json:"presence_penalty"`
	Raw               param.Field[bool]        `json:"raw"`
	RepetitionPenalty param.Field[float64]     `json:"repetition_penalty"`
	Stream            param.Field[bool]        `json:"stream"`
	Temperature       param.Field[float64]     `json:"temperature"`
	TopK              param.Field[int64]       `json:"top_k"`
	TopP              param.Field[float64]     `json:"top_p"`
	Functions         param.Field[interface{}] `json:"functions,required"`
	Messages          param.Field[interface{}] `json:"messages,required"`
	Tools             param.Field[interface{}] `json:"tools,required"`
	InputText         param.Field[string]      `json:"input_text"`
	MaxLength         param.Field[int64]       `json:"max_length"`
}

func (r AIRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBody) implementsWorkersAIRunParamsBodyUnion() {}

// Satisfied by [workers.AIRunParamsBodyTextClassification],
// [workers.AIRunParamsBodyTextToImage], [workers.AIRunParamsBodyTextEmbeddings],
// [workers.AIRunParamsBodyAutomaticSpeechRecognition],
// [workers.AIRunParamsBodyImageClassification],
// [workers.AIRunParamsBodyObjectDetection], [workers.AIRunParamsBodyObject],
// [workers.AIRunParamsBodyObject], [workers.AIRunParamsBodyTranslation],
// [workers.AIRunParamsBodySummarization], [workers.AIRunParamsBodyImageToText],
// [AIRunParamsBody].
type AIRunParamsBodyUnion interface {
	implementsWorkersAIRunParamsBodyUnion()
}

type AIRunParamsBodyTextClassification struct {
	Text param.Field[string] `json:"text,required"`
}

func (r AIRunParamsBodyTextClassification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextClassification) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextToImage struct {
	Prompt         param.Field[string]    `json:"prompt,required"`
	Guidance       param.Field[float64]   `json:"guidance"`
	Height         param.Field[int64]     `json:"height"`
	Image          param.Field[[]float64] `json:"image"`
	ImageB64       param.Field[string]    `json:"image_b64"`
	Mask           param.Field[[]float64] `json:"mask"`
	NegativePrompt param.Field[string]    `json:"negative_prompt"`
	NumSteps       param.Field[int64]     `json:"num_steps"`
	Seed           param.Field[int64]     `json:"seed"`
	Strength       param.Field[float64]   `json:"strength"`
	Width          param.Field[int64]     `json:"width"`
}

func (r AIRunParamsBodyTextToImage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextToImage) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextEmbeddings struct {
	Text param.Field[AIRunParamsBodyTextEmbeddingsTextUnion] `json:"text,required"`
}

func (r AIRunParamsBodyTextEmbeddings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextEmbeddings) implementsWorkersAIRunParamsBodyUnion() {}

// Satisfied by [shared.UnionString],
// [workers.AIRunParamsBodyTextEmbeddingsTextArray].
type AIRunParamsBodyTextEmbeddingsTextUnion interface {
	ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion()
}

type AIRunParamsBodyTextEmbeddingsTextArray []string

func (r AIRunParamsBodyTextEmbeddingsTextArray) ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion() {
}

type AIRunParamsBodyAutomaticSpeechRecognition struct {
	Audio      param.Field[[]float64] `json:"audio,required"`
	SourceLang param.Field[string]    `json:"source_lang"`
	TargetLang param.Field[string]    `json:"target_lang"`
}

func (r AIRunParamsBodyAutomaticSpeechRecognition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyAutomaticSpeechRecognition) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageClassification struct {
	Image param.Field[[]float64] `json:"image,required"`
}

func (r AIRunParamsBodyImageClassification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyImageClassification) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyObjectDetection struct {
	Image param.Field[[]float64] `json:"image"`
}

func (r AIRunParamsBodyObjectDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyObjectDetection) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyObject struct {
	Prompt            param.Field[string]  `json:"prompt,required"`
	FrequencyPenalty  param.Field[float64] `json:"frequency_penalty"`
	Lora              param.Field[string]  `json:"lora"`
	MaxTokens         param.Field[int64]   `json:"max_tokens"`
	PresencePenalty   param.Field[float64] `json:"presence_penalty"`
	Raw               param.Field[bool]    `json:"raw"`
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	Seed              param.Field[int64]   `json:"seed"`
	Stream            param.Field[bool]    `json:"stream"`
	Temperature       param.Field[float64] `json:"temperature"`
	TopK              param.Field[int64]   `json:"top_k"`
	TopP              param.Field[float64] `json:"top_p"`
}

func (r AIRunParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyObject) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTranslation struct {
	TargetLang param.Field[string] `json:"target_lang,required"`
	Text       param.Field[string] `json:"text,required"`
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AIRunParamsBodyTranslation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTranslation) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodySummarization struct {
	InputText param.Field[string] `json:"input_text,required"`
	MaxLength param.Field[int64]  `json:"max_length"`
}

func (r AIRunParamsBodySummarization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodySummarization) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageToText struct {
	Image       param.Field[[]float64]                           `json:"image,required"`
	MaxTokens   param.Field[int64]                               `json:"max_tokens"`
	Messages    param.Field[[]AIRunParamsBodyImageToTextMessage] `json:"messages"`
	Prompt      param.Field[string]                              `json:"prompt"`
	Raw         param.Field[bool]                                `json:"raw"`
	Temperature param.Field[float64]                             `json:"temperature"`
}

func (r AIRunParamsBodyImageToText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyImageToText) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageToTextMessage struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIRunParamsBodyImageToTextMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIRunResponseEnvelope struct {
	Result AIRunResponseUnion        `json:"result" format:"binary"`
	JSON   aiRunResponseEnvelopeJSON `json:"-"`
}

// aiRunResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIRunResponseEnvelope]
type aiRunResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
