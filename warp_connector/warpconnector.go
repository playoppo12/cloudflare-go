// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warp_connector

import (
  "context"
  "fmt"
  "net/http"
  "net/url"
  "reflect"
  "time"

  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
  "github.com/cloudflare/cloudflare-go/v2/internal/pagination"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/internal/shared"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/zero_trust"
  "github.com/tidwall/gjson"
)

// WARPConnectorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWARPConnectorService] method
// instead.
type WARPConnectorService struct {
Options []option.RequestOption
}

// NewWARPConnectorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWARPConnectorService(opts ...option.RequestOption) (r *WARPConnectorService) {
  r = &WARPConnectorService{}
  r.Options = opts
  return
}

// Creates a new Warp Connector Tunnel in an account.
func (r *WARPConnectorService) New(ctx context.Context, params WARPConnectorNewParams, opts ...option.RequestOption) (res *WARPConnectorNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  var env WARPConnectorNewResponseEnvelope
  path := fmt.Sprintf("accounts/%s/warp_connector", params.AccountID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Lists and filters Warp Connector Tunnels in an account.
func (r *WARPConnectorService) List(ctx context.Context, params WARPConnectorListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WARPConnectorListResponse], err error) {
  var raw *http.Response
  opts = append(r.Options, opts...)
  opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
  path := fmt.Sprintf("accounts/%s/warp_connector", params.AccountID)
  cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists and filters Warp Connector Tunnels in an account.
func (r *WARPConnectorService) ListAutoPaging(ctx context.Context, params WARPConnectorListParams, opts ...option.RequestOption) (*pagination.V4PagePaginationArrayAutoPager[WARPConnectorListResponse]) {
  return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Warp Connector Tunnel from an account.
func (r *WARPConnectorService) Delete(ctx context.Context, tunnelID string, params WARPConnectorDeleteParams, opts ...option.RequestOption) (res *WARPConnectorDeleteResponse, err error) {
  opts = append(r.Options[:], opts...)
  var env WARPConnectorDeleteResponseEnvelope
  path := fmt.Sprintf("accounts/%s/warp_connector/%s", params.AccountID, tunnelID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Updates an existing Warp Connector Tunnel.
func (r *WARPConnectorService) Edit(ctx context.Context, tunnelID string, params WARPConnectorEditParams, opts ...option.RequestOption) (res *WARPConnectorEditResponse, err error) {
  opts = append(r.Options[:], opts...)
  var env WARPConnectorEditResponseEnvelope
  path := fmt.Sprintf("accounts/%s/warp_connector/%s", params.AccountID, tunnelID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Fetches a single Warp Connector Tunnel.
func (r *WARPConnectorService) Get(ctx context.Context, tunnelID string, query WARPConnectorGetParams, opts ...option.RequestOption) (res *WARPConnectorGetResponse, err error) {
  opts = append(r.Options[:], opts...)
  var env WARPConnectorGetResponseEnvelope
  path := fmt.Sprintf("accounts/%s/warp_connector/%s", query.AccountID, tunnelID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Gets the token used to associate warp device with a specific Warp Connector
// tunnel.
func (r *WARPConnectorService) Token(ctx context.Context, tunnelID string, query WARPConnectorTokenParams, opts ...option.RequestOption) (res *WARPConnectorTokenResponseUnion, err error) {
  opts = append(r.Options[:], opts...)
  var env WARPConnectorTokenResponseEnvelope
  path := fmt.Sprintf("accounts/%s/warp_connector/%s/token", query.AccountID, tunnelID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorNewResponse struct {
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// UUID of the tunnel.
ID string `json:"id"`
Metadata interface{} `json:"metadata,required"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
// If `false`, the tunnel must be configured locally on the origin machine.
RemoteConfig bool `json:"remote_config"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType zero_trust.WARPConnectorNewResponseTunType `json:"tun_type"`
JSON warpConnectorNewResponseJSON `json:"-"`
union WARPConnectorNewResponseUnion
}

// warpConnectorNewResponseJSON contains the JSON metadata for the struct
// [WARPConnectorNewResponse]
type warpConnectorNewResponseJSON struct {
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
ID apijson.Field
Metadata apijson.Field
Name apijson.Field
RemoteConfig apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r warpConnectorNewResponseJSON) RawJSON() (string) {
  return r.raw
}

func (r *WARPConnectorNewResponse) UnmarshalJSON(data []byte) (err error) {
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

func (r WARPConnectorNewResponse) AsUnion() (WARPConnectorNewResponseUnion) {
  return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [zero_trust.Tunnel] or
// [warp_connector.WARPConnectorNewResponseTunnelWARPConnectorTunnel].
type WARPConnectorNewResponseUnion interface {
  implementsWARPConnectorWARPConnectorNewResponse()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*WARPConnectorNewResponseUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(zero_trust.Tunnel{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(WARPConnectorNewResponseTunnelWARPConnectorTunnel{}),
    },
  )
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorNewResponseTunnelWARPConnectorTunnel struct {
// UUID of the tunnel.
ID string `json:"id"`
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// Metadata associated with the tunnel.
Metadata interface{} `json:"metadata"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
JSON warpConnectorNewResponseTunnelWARPConnectorTunnelJSON `json:"-"`
}

// warpConnectorNewResponseTunnelWARPConnectorTunnelJSON contains the JSON metadata
// for the struct [WARPConnectorNewResponseTunnelWARPConnectorTunnel]
type warpConnectorNewResponseTunnelWARPConnectorTunnelJSON struct {
ID apijson.Field
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
Metadata apijson.Field
Name apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorNewResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorNewResponseTunnelWARPConnectorTunnelJSON) RawJSON() (string) {
  return r.raw
}

func (r WARPConnectorNewResponseTunnelWARPConnectorTunnel) implementsWARPConnectorWARPConnectorNewResponse() {}

// The type of tunnel.
type WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType string

const (
  WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
  WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
  WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeIPSec WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
  WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeGRE WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "gre"
  WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCni WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorNewResponseTunnelWARPConnectorTunnelTunType) IsKnown() (bool) {
  switch r {
  case WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorNewResponseTunnelWARPConnectorTunnelTunTypeCni:
      return true
  }
  return false
}

// The type of tunnel.
type zero_trust.WARPConnectorNewResponseTunType string

const (
  zero_trust.WARPConnectorNewResponseTunTypeCfdTunnel zero_trust.WARPConnectorNewResponseTunType = "cfd_tunnel"
  zero_trust.WARPConnectorNewResponseTunTypeWARPConnector zero_trust.WARPConnectorNewResponseTunType = "warp_connector"
  zero_trust.WARPConnectorNewResponseTunTypeIPSec zero_trust.WARPConnectorNewResponseTunType = "ip_sec"
  zero_trust.WARPConnectorNewResponseTunTypeGRE zero_trust.WARPConnectorNewResponseTunType = "gre"
  zero_trust.WARPConnectorNewResponseTunTypeCni zero_trust.WARPConnectorNewResponseTunType = "cni"
)

func (r zero_trust.WARPConnectorNewResponseTunType) IsKnown() (bool) {
  switch r {
  case zero_trust.WARPConnectorNewResponseTunTypeCfdTunnel, zero_trust.WARPConnectorNewResponseTunTypeWARPConnector, zero_trust.WARPConnectorNewResponseTunTypeIPSec, zero_trust.WARPConnectorNewResponseTunTypeGRE, zero_trust.WARPConnectorNewResponseTunTypeCni:
      return true
  }
  return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorListResponse struct {
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// UUID of the tunnel.
ID string `json:"id"`
Metadata interface{} `json:"metadata,required"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
// If `false`, the tunnel must be configured locally on the origin machine.
RemoteConfig bool `json:"remote_config"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType zero_trust.WARPConnectorListResponseTunType `json:"tun_type"`
JSON warpConnectorListResponseJSON `json:"-"`
union WARPConnectorListResponseUnion
}

// warpConnectorListResponseJSON contains the JSON metadata for the struct
// [WARPConnectorListResponse]
type warpConnectorListResponseJSON struct {
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
ID apijson.Field
Metadata apijson.Field
Name apijson.Field
RemoteConfig apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r warpConnectorListResponseJSON) RawJSON() (string) {
  return r.raw
}

func (r *WARPConnectorListResponse) UnmarshalJSON(data []byte) (err error) {
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

func (r WARPConnectorListResponse) AsUnion() (WARPConnectorListResponseUnion) {
  return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [zero_trust.Tunnel] or
// [warp_connector.WARPConnectorListResponseTunnelWARPConnectorTunnel].
type WARPConnectorListResponseUnion interface {
  implementsWARPConnectorWARPConnectorListResponse()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*WARPConnectorListResponseUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(zero_trust.Tunnel{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(WARPConnectorListResponseTunnelWARPConnectorTunnel{}),
    },
  )
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorListResponseTunnelWARPConnectorTunnel struct {
// UUID of the tunnel.
ID string `json:"id"`
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// Metadata associated with the tunnel.
Metadata interface{} `json:"metadata"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType WARPConnectorListResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
JSON warpConnectorListResponseTunnelWARPConnectorTunnelJSON `json:"-"`
}

// warpConnectorListResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [WARPConnectorListResponseTunnelWARPConnectorTunnel]
type warpConnectorListResponseTunnelWARPConnectorTunnelJSON struct {
ID apijson.Field
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
Metadata apijson.Field
Name apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorListResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorListResponseTunnelWARPConnectorTunnelJSON) RawJSON() (string) {
  return r.raw
}

func (r WARPConnectorListResponseTunnelWARPConnectorTunnel) implementsWARPConnectorWARPConnectorListResponse() {}

// The type of tunnel.
type WARPConnectorListResponseTunnelWARPConnectorTunnelTunType string

const (
  WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
  WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
  WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeIPSec WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
  WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeGRE WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "gre"
  WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCni WARPConnectorListResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorListResponseTunnelWARPConnectorTunnelTunType) IsKnown() (bool) {
  switch r {
  case WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorListResponseTunnelWARPConnectorTunnelTunTypeCni:
      return true
  }
  return false
}

// The type of tunnel.
type zero_trust.WARPConnectorListResponseTunType string

const (
  zero_trust.WARPConnectorListResponseTunTypeCfdTunnel zero_trust.WARPConnectorListResponseTunType = "cfd_tunnel"
  zero_trust.WARPConnectorListResponseTunTypeWARPConnector zero_trust.WARPConnectorListResponseTunType = "warp_connector"
  zero_trust.WARPConnectorListResponseTunTypeIPSec zero_trust.WARPConnectorListResponseTunType = "ip_sec"
  zero_trust.WARPConnectorListResponseTunTypeGRE zero_trust.WARPConnectorListResponseTunType = "gre"
  zero_trust.WARPConnectorListResponseTunTypeCni zero_trust.WARPConnectorListResponseTunType = "cni"
)

func (r zero_trust.WARPConnectorListResponseTunType) IsKnown() (bool) {
  switch r {
  case zero_trust.WARPConnectorListResponseTunTypeCfdTunnel, zero_trust.WARPConnectorListResponseTunTypeWARPConnector, zero_trust.WARPConnectorListResponseTunTypeIPSec, zero_trust.WARPConnectorListResponseTunTypeGRE, zero_trust.WARPConnectorListResponseTunTypeCni:
      return true
  }
  return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorDeleteResponse struct {
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// UUID of the tunnel.
ID string `json:"id"`
Metadata interface{} `json:"metadata,required"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
// If `false`, the tunnel must be configured locally on the origin machine.
RemoteConfig bool `json:"remote_config"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType zero_trust.WARPConnectorDeleteResponseTunType `json:"tun_type"`
JSON warpConnectorDeleteResponseJSON `json:"-"`
union WARPConnectorDeleteResponseUnion
}

// warpConnectorDeleteResponseJSON contains the JSON metadata for the struct
// [WARPConnectorDeleteResponse]
type warpConnectorDeleteResponseJSON struct {
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
ID apijson.Field
Metadata apijson.Field
Name apijson.Field
RemoteConfig apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r warpConnectorDeleteResponseJSON) RawJSON() (string) {
  return r.raw
}

func (r *WARPConnectorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

func (r WARPConnectorDeleteResponse) AsUnion() (WARPConnectorDeleteResponseUnion) {
  return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [zero_trust.Tunnel] or
// [warp_connector.WARPConnectorDeleteResponseTunnelWARPConnectorTunnel].
type WARPConnectorDeleteResponseUnion interface {
  implementsWARPConnectorWARPConnectorDeleteResponse()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*WARPConnectorDeleteResponseUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(zero_trust.Tunnel{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(WARPConnectorDeleteResponseTunnelWARPConnectorTunnel{}),
    },
  )
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorDeleteResponseTunnelWARPConnectorTunnel struct {
// UUID of the tunnel.
ID string `json:"id"`
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// Metadata associated with the tunnel.
Metadata interface{} `json:"metadata"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
JSON warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON `json:"-"`
}

// warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [WARPConnectorDeleteResponseTunnelWARPConnectorTunnel]
type warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON struct {
ID apijson.Field
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
Metadata apijson.Field
Name apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorDeleteResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorDeleteResponseTunnelWARPConnectorTunnelJSON) RawJSON() (string) {
  return r.raw
}

func (r WARPConnectorDeleteResponseTunnelWARPConnectorTunnel) implementsWARPConnectorWARPConnectorDeleteResponse() {}

// The type of tunnel.
type WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType string

const (
  WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
  WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
  WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeIPSec WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
  WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeGRE WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "gre"
  WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCni WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunType) IsKnown() (bool) {
  switch r {
  case WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorDeleteResponseTunnelWARPConnectorTunnelTunTypeCni:
      return true
  }
  return false
}

// The type of tunnel.
type zero_trust.WARPConnectorDeleteResponseTunType string

const (
  zero_trust.WARPConnectorDeleteResponseTunTypeCfdTunnel zero_trust.WARPConnectorDeleteResponseTunType = "cfd_tunnel"
  zero_trust.WARPConnectorDeleteResponseTunTypeWARPConnector zero_trust.WARPConnectorDeleteResponseTunType = "warp_connector"
  zero_trust.WARPConnectorDeleteResponseTunTypeIPSec zero_trust.WARPConnectorDeleteResponseTunType = "ip_sec"
  zero_trust.WARPConnectorDeleteResponseTunTypeGRE zero_trust.WARPConnectorDeleteResponseTunType = "gre"
  zero_trust.WARPConnectorDeleteResponseTunTypeCni zero_trust.WARPConnectorDeleteResponseTunType = "cni"
)

func (r zero_trust.WARPConnectorDeleteResponseTunType) IsKnown() (bool) {
  switch r {
  case zero_trust.WARPConnectorDeleteResponseTunTypeCfdTunnel, zero_trust.WARPConnectorDeleteResponseTunTypeWARPConnector, zero_trust.WARPConnectorDeleteResponseTunTypeIPSec, zero_trust.WARPConnectorDeleteResponseTunTypeGRE, zero_trust.WARPConnectorDeleteResponseTunTypeCni:
      return true
  }
  return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorEditResponse struct {
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// UUID of the tunnel.
ID string `json:"id"`
Metadata interface{} `json:"metadata,required"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
// If `false`, the tunnel must be configured locally on the origin machine.
RemoteConfig bool `json:"remote_config"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType zero_trust.WARPConnectorEditResponseTunType `json:"tun_type"`
JSON warpConnectorEditResponseJSON `json:"-"`
union WARPConnectorEditResponseUnion
}

// warpConnectorEditResponseJSON contains the JSON metadata for the struct
// [WARPConnectorEditResponse]
type warpConnectorEditResponseJSON struct {
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
ID apijson.Field
Metadata apijson.Field
Name apijson.Field
RemoteConfig apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r warpConnectorEditResponseJSON) RawJSON() (string) {
  return r.raw
}

func (r *WARPConnectorEditResponse) UnmarshalJSON(data []byte) (err error) {
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

func (r WARPConnectorEditResponse) AsUnion() (WARPConnectorEditResponseUnion) {
  return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [zero_trust.Tunnel] or
// [warp_connector.WARPConnectorEditResponseTunnelWARPConnectorTunnel].
type WARPConnectorEditResponseUnion interface {
  implementsWARPConnectorWARPConnectorEditResponse()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*WARPConnectorEditResponseUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(zero_trust.Tunnel{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(WARPConnectorEditResponseTunnelWARPConnectorTunnel{}),
    },
  )
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorEditResponseTunnelWARPConnectorTunnel struct {
// UUID of the tunnel.
ID string `json:"id"`
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// Metadata associated with the tunnel.
Metadata interface{} `json:"metadata"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
JSON warpConnectorEditResponseTunnelWARPConnectorTunnelJSON `json:"-"`
}

// warpConnectorEditResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [WARPConnectorEditResponseTunnelWARPConnectorTunnel]
type warpConnectorEditResponseTunnelWARPConnectorTunnelJSON struct {
ID apijson.Field
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
Metadata apijson.Field
Name apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorEditResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorEditResponseTunnelWARPConnectorTunnelJSON) RawJSON() (string) {
  return r.raw
}

func (r WARPConnectorEditResponseTunnelWARPConnectorTunnel) implementsWARPConnectorWARPConnectorEditResponse() {}

// The type of tunnel.
type WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType string

const (
  WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
  WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
  WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeIPSec WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
  WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeGRE WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "gre"
  WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCni WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorEditResponseTunnelWARPConnectorTunnelTunType) IsKnown() (bool) {
  switch r {
  case WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorEditResponseTunnelWARPConnectorTunnelTunTypeCni:
      return true
  }
  return false
}

// The type of tunnel.
type zero_trust.WARPConnectorEditResponseTunType string

const (
  zero_trust.WARPConnectorEditResponseTunTypeCfdTunnel zero_trust.WARPConnectorEditResponseTunType = "cfd_tunnel"
  zero_trust.WARPConnectorEditResponseTunTypeWARPConnector zero_trust.WARPConnectorEditResponseTunType = "warp_connector"
  zero_trust.WARPConnectorEditResponseTunTypeIPSec zero_trust.WARPConnectorEditResponseTunType = "ip_sec"
  zero_trust.WARPConnectorEditResponseTunTypeGRE zero_trust.WARPConnectorEditResponseTunType = "gre"
  zero_trust.WARPConnectorEditResponseTunTypeCni zero_trust.WARPConnectorEditResponseTunType = "cni"
)

func (r zero_trust.WARPConnectorEditResponseTunType) IsKnown() (bool) {
  switch r {
  case zero_trust.WARPConnectorEditResponseTunTypeCfdTunnel, zero_trust.WARPConnectorEditResponseTunTypeWARPConnector, zero_trust.WARPConnectorEditResponseTunTypeIPSec, zero_trust.WARPConnectorEditResponseTunTypeGRE, zero_trust.WARPConnectorEditResponseTunTypeCni:
      return true
  }
  return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorGetResponse struct {
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// UUID of the tunnel.
ID string `json:"id"`
Metadata interface{} `json:"metadata,required"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
// If `false`, the tunnel must be configured locally on the origin machine.
RemoteConfig bool `json:"remote_config"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType zero_trust.WARPConnectorGetResponseTunType `json:"tun_type"`
JSON warpConnectorGetResponseJSON `json:"-"`
union WARPConnectorGetResponseUnion
}

// warpConnectorGetResponseJSON contains the JSON metadata for the struct
// [WARPConnectorGetResponse]
type warpConnectorGetResponseJSON struct {
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
ID apijson.Field
Metadata apijson.Field
Name apijson.Field
RemoteConfig apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r warpConnectorGetResponseJSON) RawJSON() (string) {
  return r.raw
}

func (r *WARPConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

func (r WARPConnectorGetResponse) AsUnion() (WARPConnectorGetResponseUnion) {
  return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [zero_trust.Tunnel] or
// [warp_connector.WARPConnectorGetResponseTunnelWARPConnectorTunnel].
type WARPConnectorGetResponseUnion interface {
  implementsWARPConnectorWARPConnectorGetResponse()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*WARPConnectorGetResponseUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(zero_trust.Tunnel{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(WARPConnectorGetResponseTunnelWARPConnectorTunnel{}),
    },
  )
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WARPConnectorGetResponseTunnelWARPConnectorTunnel struct {
// UUID of the tunnel.
ID string `json:"id"`
// Cloudflare account ID
AccountTag string `json:"account_tag"`
// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
Connections zero_trust.Connection `json:"connections"`
// Timestamp of when the tunnel established at least one connection to Cloudflare's
// edge. If `null`, the tunnel is inactive.
ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
// edge). If `null`, the tunnel is active.
ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
// Timestamp of when the tunnel was created.
CreatedAt time.Time `json:"created_at" format:"date-time"`
// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
// deleted.
DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
// Metadata associated with the tunnel.
Metadata interface{} `json:"metadata"`
// A user-friendly name for the tunnel.
Name string `json:"name"`
// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
Status string `json:"status"`
// The type of tunnel.
TunType WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
JSON warpConnectorGetResponseTunnelWARPConnectorTunnelJSON `json:"-"`
}

// warpConnectorGetResponseTunnelWARPConnectorTunnelJSON contains the JSON metadata
// for the struct [WARPConnectorGetResponseTunnelWARPConnectorTunnel]
type warpConnectorGetResponseTunnelWARPConnectorTunnelJSON struct {
ID apijson.Field
AccountTag apijson.Field
Connections apijson.Field
ConnsActiveAt apijson.Field
ConnsInactiveAt apijson.Field
CreatedAt apijson.Field
DeletedAt apijson.Field
Metadata apijson.Field
Name apijson.Field
Status apijson.Field
TunType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorGetResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorGetResponseTunnelWARPConnectorTunnelJSON) RawJSON() (string) {
  return r.raw
}

func (r WARPConnectorGetResponseTunnelWARPConnectorTunnel) implementsWARPConnectorWARPConnectorGetResponse() {}

// The type of tunnel.
type WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType string

const (
  WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
  WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeWARPConnector WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
  WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeIPSec WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
  WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeGRE WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "gre"
  WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCni WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r WARPConnectorGetResponseTunnelWARPConnectorTunnelTunType) IsKnown() (bool) {
  switch r {
  case WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeIPSec, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeGRE, WARPConnectorGetResponseTunnelWARPConnectorTunnelTunTypeCni:
      return true
  }
  return false
}

// The type of tunnel.
type zero_trust.WARPConnectorGetResponseTunType string

const (
  zero_trust.WARPConnectorGetResponseTunTypeCfdTunnel zero_trust.WARPConnectorGetResponseTunType = "cfd_tunnel"
  zero_trust.WARPConnectorGetResponseTunTypeWARPConnector zero_trust.WARPConnectorGetResponseTunType = "warp_connector"
  zero_trust.WARPConnectorGetResponseTunTypeIPSec zero_trust.WARPConnectorGetResponseTunType = "ip_sec"
  zero_trust.WARPConnectorGetResponseTunTypeGRE zero_trust.WARPConnectorGetResponseTunType = "gre"
  zero_trust.WARPConnectorGetResponseTunTypeCni zero_trust.WARPConnectorGetResponseTunType = "cni"
)

func (r zero_trust.WARPConnectorGetResponseTunType) IsKnown() (bool) {
  switch r {
  case zero_trust.WARPConnectorGetResponseTunTypeCfdTunnel, zero_trust.WARPConnectorGetResponseTunTypeWARPConnector, zero_trust.WARPConnectorGetResponseTunTypeIPSec, zero_trust.WARPConnectorGetResponseTunTypeGRE, zero_trust.WARPConnectorGetResponseTunTypeCni:
      return true
  }
  return false
}

// Union satisfied by [warp_connector.WARPConnectorTokenResponseUnknown],
// [warp_connector.WARPConnectorTokenResponseArray] or [shared.UnionString].
type WARPConnectorTokenResponseUnion interface {
  ImplementsWARPConnectorWARPConnectorTokenResponseUnion()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*WARPConnectorTokenResponseUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(WARPConnectorTokenResponseArray{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.String,
      Type: reflect.TypeOf(shared.UnionString("")),
    },
  )
}

type WARPConnectorTokenResponseArray []interface{}

func (r WARPConnectorTokenResponseArray) ImplementsWARPConnectorWARPConnectorTokenResponseUnion() {}

type WARPConnectorNewParams struct {
// Cloudflare account ID
AccountID param.Field[string] `path:"account_id,required"`
// A user-friendly name for the tunnel.
Name param.Field[string] `json:"name,required"`
}

func (r WARPConnectorNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type WARPConnectorNewResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
Result WARPConnectorNewResponse `json:"result,required"`
// Whether the API call was successful
Success WARPConnectorNewResponseEnvelopeSuccess `json:"success,required"`
JSON warpConnectorNewResponseEnvelopeJSON `json:"-"`
}

// warpConnectorNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorNewResponseEnvelope]
type warpConnectorNewResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorNewResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type WARPConnectorNewResponseEnvelopeSuccess bool

const (
  WARPConnectorNewResponseEnvelopeSuccessTrue WARPConnectorNewResponseEnvelopeSuccess = true
)

func (r WARPConnectorNewResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case WARPConnectorNewResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type WARPConnectorListParams struct {
// Cloudflare account ID
AccountID param.Field[string] `path:"account_id,required"`
ExcludePrefix param.Field[string] `query:"exclude_prefix"`
// If provided, include only tunnels that were created (and not deleted) before
// this time.
ExistedAt param.Field[time.Time] `query:"existed_at" format:"date-time"`
IncludePrefix param.Field[string] `query:"include_prefix"`
// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
// empty, all tunnels will be included.
IsDeleted param.Field[bool] `query:"is_deleted"`
// A user-friendly name for the tunnel.
Name param.Field[string] `query:"name"`
// Page number of paginated results.
Page param.Field[float64] `query:"page"`
// Number of results to display.
PerPage param.Field[float64] `query:"per_page"`
// UUID of the tunnel.
UUID param.Field[string] `query:"uuid"`
WasActiveAt param.Field[time.Time] `query:"was_active_at" format:"date-time"`
WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [WARPConnectorListParams]'s query parameters as
// `url.Values`.
func (r WARPConnectorListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type WARPConnectorDeleteParams struct {
// Cloudflare account ID
AccountID param.Field[string] `path:"account_id,required"`
Body interface{} `json:"body,required"`
}

func (r WARPConnectorDeleteParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r.Body)
}

type WARPConnectorDeleteResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
Result WARPConnectorDeleteResponse `json:"result,required"`
// Whether the API call was successful
Success WARPConnectorDeleteResponseEnvelopeSuccess `json:"success,required"`
JSON warpConnectorDeleteResponseEnvelopeJSON `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WARPConnectorDeleteResponseEnvelope]
type warpConnectorDeleteResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorDeleteResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type WARPConnectorDeleteResponseEnvelopeSuccess bool

const (
  WARPConnectorDeleteResponseEnvelopeSuccessTrue WARPConnectorDeleteResponseEnvelopeSuccess = true
)

func (r WARPConnectorDeleteResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case WARPConnectorDeleteResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type WARPConnectorEditParams struct {
// Cloudflare account ID
AccountID param.Field[string] `path:"account_id,required"`
// A user-friendly name for the tunnel.
Name param.Field[string] `json:"name"`
// Sets the password required to run a locally-managed tunnel. Must be at least 32
// bytes and encoded as a base64 string.
TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r WARPConnectorEditParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type WARPConnectorEditResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
Result WARPConnectorEditResponse `json:"result,required"`
// Whether the API call was successful
Success WARPConnectorEditResponseEnvelopeSuccess `json:"success,required"`
JSON warpConnectorEditResponseEnvelopeJSON `json:"-"`
}

// warpConnectorEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorEditResponseEnvelope]
type warpConnectorEditResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorEditResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type WARPConnectorEditResponseEnvelopeSuccess bool

const (
  WARPConnectorEditResponseEnvelopeSuccessTrue WARPConnectorEditResponseEnvelopeSuccess = true
)

func (r WARPConnectorEditResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case WARPConnectorEditResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type WARPConnectorGetParams struct {
// Cloudflare account ID
AccountID param.Field[string] `path:"account_id,required"`
}

type WARPConnectorGetResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
Result WARPConnectorGetResponse `json:"result,required"`
// Whether the API call was successful
Success WARPConnectorGetResponseEnvelopeSuccess `json:"success,required"`
JSON warpConnectorGetResponseEnvelopeJSON `json:"-"`
}

// warpConnectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorGetResponseEnvelope]
type warpConnectorGetResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorGetResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type WARPConnectorGetResponseEnvelopeSuccess bool

const (
  WARPConnectorGetResponseEnvelopeSuccessTrue WARPConnectorGetResponseEnvelopeSuccess = true
)

func (r WARPConnectorGetResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case WARPConnectorGetResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type WARPConnectorTokenParams struct {
// Cloudflare account ID
AccountID param.Field[string] `path:"account_id,required"`
}

type WARPConnectorTokenResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
Result WARPConnectorTokenResponseUnion `json:"result,required"`
// Whether the API call was successful
Success WARPConnectorTokenResponseEnvelopeSuccess `json:"success,required"`
JSON warpConnectorTokenResponseEnvelopeJSON `json:"-"`
}

// warpConnectorTokenResponseEnvelopeJSON contains the JSON metadata for the struct
// [WARPConnectorTokenResponseEnvelope]
type warpConnectorTokenResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *WARPConnectorTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorTokenResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type WARPConnectorTokenResponseEnvelopeSuccess bool

const (
  WARPConnectorTokenResponseEnvelopeSuccessTrue WARPConnectorTokenResponseEnvelopeSuccess = true
)

func (r WARPConnectorTokenResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case WARPConnectorTokenResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}
