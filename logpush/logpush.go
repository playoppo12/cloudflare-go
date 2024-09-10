// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// LogpushService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogpushService] method instead.
type LogpushService struct {
	Options   []option.RequestOption
	Datasets  *DatasetService
	Edge      *EdgeService
	Jobs      *JobService
	Ownership *OwnershipService
	Validate  *ValidateService
}

// NewLogpushService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogpushService(opts ...option.RequestOption) (r *LogpushService) {
	r = &LogpushService{}
	r.Options = opts
	r.Datasets = NewDatasetService(opts...)
	r.Edge = NewEdgeService(opts...)
	r.Jobs = NewJobService(opts...)
	r.Ownership = NewOwnershipService(opts...)
	r.Validate = NewValidateService(opts...)
	return
}
