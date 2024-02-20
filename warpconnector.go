// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WarpConnectorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWarpConnectorService] method
// instead.
type WarpConnectorService struct {
	Options []option.RequestOption
}

// NewWarpConnectorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWarpConnectorService(opts ...option.RequestOption) (r *WarpConnectorService) {
	r = &WarpConnectorService{}
	r.Options = opts
	return
}

// Creates a new Warp Connector Tunnel in an account.
func (r *WarpConnectorService) New(ctx context.Context, accountID string, body WarpConnectorNewParams, opts ...option.RequestOption) (res *WarpConnectorNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters Warp Connector Tunnels in an account.
func (r *WarpConnectorService) List(ctx context.Context, accountID string, query WarpConnectorListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[WarpConnectorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/warp_connector", accountID)
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

// Lists and filters Warp Connector Tunnels in an account.
func (r *WarpConnectorService) ListAutoPaging(ctx context.Context, accountID string, query WarpConnectorListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[WarpConnectorListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Deletes a Warp Connector Tunnel from an account.
func (r *WarpConnectorService) Delete(ctx context.Context, accountID string, tunnelID string, body WarpConnectorDeleteParams, opts ...option.RequestOption) (res *WarpConnectorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Warp Connector Tunnel.
func (r *WarpConnectorService) Edit(ctx context.Context, accountID string, tunnelID string, body WarpConnectorEditParams, opts ...option.RequestOption) (res *WarpConnectorEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Warp Connector Tunnel.
func (r *WarpConnectorService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *WarpConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WarpConnectorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/warp_connector/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorNewResponseTunnelCfdTunnel] or
// [WarpConnectorNewResponseTunnelWarpConnectorTunnel].
type WarpConnectorNewResponse interface {
	implementsWarpConnectorNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorNewResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorNewResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorNewResponseTunnelCfdTunnelConnection `json:"connections"`
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
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorNewResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseTunnelCfdTunnelJSON    `json:"-"`
}

// warpConnectorNewResponseTunnelCfdTunnelJSON contains the JSON metadata for the
// struct [WarpConnectorNewResponseTunnelCfdTunnel]
type warpConnectorNewResponseTunnelCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorNewResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorNewResponseTunnelCfdTunnel) implementsWarpConnectorNewResponse() {}

type WarpConnectorNewResponseTunnelCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                `json:"uuid"`
	JSON warpConnectorNewResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorNewResponseTunnelCfdTunnelConnectionJSON contains the JSON metadata
// for the struct [WarpConnectorNewResponseTunnelCfdTunnelConnection]
type warpConnectorNewResponseTunnelCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorNewResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorNewResponseTunnelCfdTunnelTunType string

const (
	WarpConnectorNewResponseTunnelCfdTunnelTunTypeCfdTunnel     WarpConnectorNewResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorNewResponseTunnelCfdTunnelTunTypeWarpConnector WarpConnectorNewResponseTunnelCfdTunnelTunType = "warp_connector"
	WarpConnectorNewResponseTunnelCfdTunnelTunTypeIPSec         WarpConnectorNewResponseTunnelCfdTunnelTunType = "ip_sec"
	WarpConnectorNewResponseTunnelCfdTunnelTunTypeGre           WarpConnectorNewResponseTunnelCfdTunnelTunType = "gre"
	WarpConnectorNewResponseTunnelCfdTunnelTunTypeCni           WarpConnectorNewResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorNewResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorNewResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorNewResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorNewResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorNewResponseTunnelWarpConnectorTunnelJSON contains the JSON metadata
// for the struct [WarpConnectorNewResponseTunnelWarpConnectorTunnel]
type warpConnectorNewResponseTunnelWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorNewResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorNewResponseTunnelWarpConnectorTunnel) implementsWarpConnectorNewResponse() {}

type WarpConnectorNewResponseTunnelWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                          `json:"uuid"`
	JSON warpConnectorNewResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorNewResponseTunnelWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorNewResponseTunnelWarpConnectorTunnelConnection]
type warpConnectorNewResponseTunnelWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorNewResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorNewResponseTunnelWarpConnectorTunnelTunType string

const (
	WarpConnectorNewResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorNewResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorNewResponseTunnelWarpConnectorTunnelTunTypeWarpConnector WarpConnectorNewResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorNewResponseTunnelWarpConnectorTunnelTunTypeIPSec         WarpConnectorNewResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorNewResponseTunnelWarpConnectorTunnelTunTypeGre           WarpConnectorNewResponseTunnelWarpConnectorTunnelTunType = "gre"
	WarpConnectorNewResponseTunnelWarpConnectorTunnelTunTypeCni           WarpConnectorNewResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorListResponseTunnelCfdTunnel] or
// [WarpConnectorListResponseTunnelWarpConnectorTunnel].
type WarpConnectorListResponse interface {
	implementsWarpConnectorListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorListResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorListResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorListResponseTunnelCfdTunnelConnection `json:"connections"`
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
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorListResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorListResponseTunnelCfdTunnelJSON    `json:"-"`
}

// warpConnectorListResponseTunnelCfdTunnelJSON contains the JSON metadata for the
// struct [WarpConnectorListResponseTunnelCfdTunnel]
type warpConnectorListResponseTunnelCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorListResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorListResponseTunnelCfdTunnel) implementsWarpConnectorListResponse() {}

type WarpConnectorListResponseTunnelCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                 `json:"uuid"`
	JSON warpConnectorListResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorListResponseTunnelCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorListResponseTunnelCfdTunnelConnection]
type warpConnectorListResponseTunnelCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorListResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorListResponseTunnelCfdTunnelTunType string

const (
	WarpConnectorListResponseTunnelCfdTunnelTunTypeCfdTunnel     WarpConnectorListResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorListResponseTunnelCfdTunnelTunTypeWarpConnector WarpConnectorListResponseTunnelCfdTunnelTunType = "warp_connector"
	WarpConnectorListResponseTunnelCfdTunnelTunTypeIPSec         WarpConnectorListResponseTunnelCfdTunnelTunType = "ip_sec"
	WarpConnectorListResponseTunnelCfdTunnelTunTypeGre           WarpConnectorListResponseTunnelCfdTunnelTunType = "gre"
	WarpConnectorListResponseTunnelCfdTunnelTunTypeCni           WarpConnectorListResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorListResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorListResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorListResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorListResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorListResponseTunnelWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorListResponseTunnelWarpConnectorTunnel]
type warpConnectorListResponseTunnelWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorListResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorListResponseTunnelWarpConnectorTunnel) implementsWarpConnectorListResponse() {}

type WarpConnectorListResponseTunnelWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                           `json:"uuid"`
	JSON warpConnectorListResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorListResponseTunnelWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorListResponseTunnelWarpConnectorTunnelConnection]
type warpConnectorListResponseTunnelWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorListResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorListResponseTunnelWarpConnectorTunnelTunType string

const (
	WarpConnectorListResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorListResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorListResponseTunnelWarpConnectorTunnelTunTypeWarpConnector WarpConnectorListResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorListResponseTunnelWarpConnectorTunnelTunTypeIPSec         WarpConnectorListResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorListResponseTunnelWarpConnectorTunnelTunTypeGre           WarpConnectorListResponseTunnelWarpConnectorTunnelTunType = "gre"
	WarpConnectorListResponseTunnelWarpConnectorTunnelTunTypeCni           WarpConnectorListResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorDeleteResponseTunnelCfdTunnel] or
// [WarpConnectorDeleteResponseTunnelWarpConnectorTunnel].
type WarpConnectorDeleteResponse interface {
	implementsWarpConnectorDeleteResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorDeleteResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorDeleteResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorDeleteResponseTunnelCfdTunnelConnection `json:"connections"`
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
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorDeleteResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseTunnelCfdTunnelJSON    `json:"-"`
}

// warpConnectorDeleteResponseTunnelCfdTunnelJSON contains the JSON metadata for
// the struct [WarpConnectorDeleteResponseTunnelCfdTunnel]
type warpConnectorDeleteResponseTunnelCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorDeleteResponseTunnelCfdTunnel) implementsWarpConnectorDeleteResponse() {}

type WarpConnectorDeleteResponseTunnelCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                   `json:"uuid"`
	JSON warpConnectorDeleteResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorDeleteResponseTunnelCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorDeleteResponseTunnelCfdTunnelConnection]
type warpConnectorDeleteResponseTunnelCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorDeleteResponseTunnelCfdTunnelTunType string

const (
	WarpConnectorDeleteResponseTunnelCfdTunnelTunTypeCfdTunnel     WarpConnectorDeleteResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorDeleteResponseTunnelCfdTunnelTunTypeWarpConnector WarpConnectorDeleteResponseTunnelCfdTunnelTunType = "warp_connector"
	WarpConnectorDeleteResponseTunnelCfdTunnelTunTypeIPSec         WarpConnectorDeleteResponseTunnelCfdTunnelTunType = "ip_sec"
	WarpConnectorDeleteResponseTunnelCfdTunnelTunTypeGre           WarpConnectorDeleteResponseTunnelCfdTunnelTunType = "gre"
	WarpConnectorDeleteResponseTunnelCfdTunnelTunTypeCni           WarpConnectorDeleteResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorDeleteResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorDeleteResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorDeleteResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorDeleteResponseTunnelWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorDeleteResponseTunnelWarpConnectorTunnel]
type warpConnectorDeleteResponseTunnelWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorDeleteResponseTunnelWarpConnectorTunnel) implementsWarpConnectorDeleteResponse() {
}

type WarpConnectorDeleteResponseTunnelWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                             `json:"uuid"`
	JSON warpConnectorDeleteResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorDeleteResponseTunnelWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorDeleteResponseTunnelWarpConnectorTunnelConnection]
type warpConnectorDeleteResponseTunnelWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunType string

const (
	WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunTypeWarpConnector WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunTypeIPSec         WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunTypeGre           WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunType = "gre"
	WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunTypeCni           WarpConnectorDeleteResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorEditResponseTunnelCfdTunnel] or
// [WarpConnectorEditResponseTunnelWarpConnectorTunnel].
type WarpConnectorEditResponse interface {
	implementsWarpConnectorEditResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorEditResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorEditResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorEditResponseTunnelCfdTunnelConnection `json:"connections"`
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
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorEditResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorEditResponseTunnelCfdTunnelJSON    `json:"-"`
}

// warpConnectorEditResponseTunnelCfdTunnelJSON contains the JSON metadata for the
// struct [WarpConnectorEditResponseTunnelCfdTunnel]
type warpConnectorEditResponseTunnelCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorEditResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorEditResponseTunnelCfdTunnel) implementsWarpConnectorEditResponse() {}

type WarpConnectorEditResponseTunnelCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                 `json:"uuid"`
	JSON warpConnectorEditResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorEditResponseTunnelCfdTunnelConnectionJSON contains the JSON
// metadata for the struct [WarpConnectorEditResponseTunnelCfdTunnelConnection]
type warpConnectorEditResponseTunnelCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorEditResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorEditResponseTunnelCfdTunnelTunType string

const (
	WarpConnectorEditResponseTunnelCfdTunnelTunTypeCfdTunnel     WarpConnectorEditResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorEditResponseTunnelCfdTunnelTunTypeWarpConnector WarpConnectorEditResponseTunnelCfdTunnelTunType = "warp_connector"
	WarpConnectorEditResponseTunnelCfdTunnelTunTypeIPSec         WarpConnectorEditResponseTunnelCfdTunnelTunType = "ip_sec"
	WarpConnectorEditResponseTunnelCfdTunnelTunTypeGre           WarpConnectorEditResponseTunnelCfdTunnelTunType = "gre"
	WarpConnectorEditResponseTunnelCfdTunnelTunTypeCni           WarpConnectorEditResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorEditResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorEditResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorEditResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorEditResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorEditResponseTunnelWarpConnectorTunnelJSON contains the JSON
// metadata for the struct [WarpConnectorEditResponseTunnelWarpConnectorTunnel]
type warpConnectorEditResponseTunnelWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorEditResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorEditResponseTunnelWarpConnectorTunnel) implementsWarpConnectorEditResponse() {}

type WarpConnectorEditResponseTunnelWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                           `json:"uuid"`
	JSON warpConnectorEditResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorEditResponseTunnelWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorEditResponseTunnelWarpConnectorTunnelConnection]
type warpConnectorEditResponseTunnelWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorEditResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorEditResponseTunnelWarpConnectorTunnelTunType string

const (
	WarpConnectorEditResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorEditResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorEditResponseTunnelWarpConnectorTunnelTunTypeWarpConnector WarpConnectorEditResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorEditResponseTunnelWarpConnectorTunnelTunTypeIPSec         WarpConnectorEditResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorEditResponseTunnelWarpConnectorTunnelTunTypeGre           WarpConnectorEditResponseTunnelWarpConnectorTunnelTunType = "gre"
	WarpConnectorEditResponseTunnelWarpConnectorTunnelTunTypeCni           WarpConnectorEditResponseTunnelWarpConnectorTunnelTunType = "cni"
)

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [WarpConnectorGetResponseTunnelCfdTunnel] or
// [WarpConnectorGetResponseTunnelWarpConnectorTunnel].
type WarpConnectorGetResponse interface {
	implementsWarpConnectorGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WarpConnectorGetResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorGetResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorGetResponseTunnelCfdTunnelConnection `json:"connections"`
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
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType WarpConnectorGetResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseTunnelCfdTunnelJSON    `json:"-"`
}

// warpConnectorGetResponseTunnelCfdTunnelJSON contains the JSON metadata for the
// struct [WarpConnectorGetResponseTunnelCfdTunnel]
type warpConnectorGetResponseTunnelCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorGetResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorGetResponseTunnelCfdTunnel) implementsWarpConnectorGetResponse() {}

type WarpConnectorGetResponseTunnelCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                `json:"uuid"`
	JSON warpConnectorGetResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// warpConnectorGetResponseTunnelCfdTunnelConnectionJSON contains the JSON metadata
// for the struct [WarpConnectorGetResponseTunnelCfdTunnelConnection]
type warpConnectorGetResponseTunnelCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorGetResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorGetResponseTunnelCfdTunnelTunType string

const (
	WarpConnectorGetResponseTunnelCfdTunnelTunTypeCfdTunnel     WarpConnectorGetResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	WarpConnectorGetResponseTunnelCfdTunnelTunTypeWarpConnector WarpConnectorGetResponseTunnelCfdTunnelTunType = "warp_connector"
	WarpConnectorGetResponseTunnelCfdTunnelTunTypeIPSec         WarpConnectorGetResponseTunnelCfdTunnelTunType = "ip_sec"
	WarpConnectorGetResponseTunnelCfdTunnelTunTypeGre           WarpConnectorGetResponseTunnelCfdTunnelTunType = "gre"
	WarpConnectorGetResponseTunnelCfdTunnelTunTypeCni           WarpConnectorGetResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorGetResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []WarpConnectorGetResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
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
	TunType WarpConnectorGetResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    warpConnectorGetResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// warpConnectorGetResponseTunnelWarpConnectorTunnelJSON contains the JSON metadata
// for the struct [WarpConnectorGetResponseTunnelWarpConnectorTunnel]
type warpConnectorGetResponseTunnelWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorGetResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WarpConnectorGetResponseTunnelWarpConnectorTunnel) implementsWarpConnectorGetResponse() {}

type WarpConnectorGetResponseTunnelWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                          `json:"uuid"`
	JSON warpConnectorGetResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// warpConnectorGetResponseTunnelWarpConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [WarpConnectorGetResponseTunnelWarpConnectorTunnelConnection]
type warpConnectorGetResponseTunnelWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WarpConnectorGetResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type WarpConnectorGetResponseTunnelWarpConnectorTunnelTunType string

const (
	WarpConnectorGetResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     WarpConnectorGetResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	WarpConnectorGetResponseTunnelWarpConnectorTunnelTunTypeWarpConnector WarpConnectorGetResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	WarpConnectorGetResponseTunnelWarpConnectorTunnelTunTypeIPSec         WarpConnectorGetResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	WarpConnectorGetResponseTunnelWarpConnectorTunnelTunTypeGre           WarpConnectorGetResponseTunnelWarpConnectorTunnelTunType = "gre"
	WarpConnectorGetResponseTunnelWarpConnectorTunnelTunTypeCni           WarpConnectorGetResponseTunnelWarpConnectorTunnelTunType = "cni"
)

type WarpConnectorNewParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
}

func (r WarpConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WarpConnectorNewResponseEnvelope struct {
	Errors   []WarpConnectorNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WarpConnectorNewResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success WarpConnectorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorNewResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WarpConnectorNewResponseEnvelope]
type warpConnectorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    warpConnectorNewResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorNewResponseEnvelopeErrors]
type warpConnectorNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    warpConnectorNewResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WarpConnectorNewResponseEnvelopeMessages]
type warpConnectorNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorNewResponseEnvelopeSuccess bool

const (
	WarpConnectorNewResponseEnvelopeSuccessTrue WarpConnectorNewResponseEnvelopeSuccess = true
)

type WarpConnectorListParams struct {
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only tunnels that were created (and not deleted) before
	// this time.
	ExistedAt     param.Field[time.Time] `query:"existed_at" format:"date-time"`
	IncludePrefix param.Field[string]    `query:"include_prefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for the tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage       param.Field[float64]   `query:"per_page"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [WarpConnectorListParams]'s query parameters as
// `url.Values`.
func (r WarpConnectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarpConnectorDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WarpConnectorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WarpConnectorDeleteResponseEnvelope struct {
	Errors   []WarpConnectorDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WarpConnectorDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success WarpConnectorDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorDeleteResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WarpConnectorDeleteResponseEnvelope]
type warpConnectorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    warpConnectorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorDeleteResponseEnvelopeErrors]
type warpConnectorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    warpConnectorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WarpConnectorDeleteResponseEnvelopeMessages]
type warpConnectorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorDeleteResponseEnvelopeSuccess bool

const (
	WarpConnectorDeleteResponseEnvelopeSuccessTrue WarpConnectorDeleteResponseEnvelopeSuccess = true
)

type WarpConnectorEditParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r WarpConnectorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WarpConnectorEditResponseEnvelope struct {
	Errors   []WarpConnectorEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WarpConnectorEditResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success WarpConnectorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorEditResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [WarpConnectorEditResponseEnvelope]
type warpConnectorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    warpConnectorEditResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorEditResponseEnvelopeErrors]
type warpConnectorEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    warpConnectorEditResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WarpConnectorEditResponseEnvelopeMessages]
type warpConnectorEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorEditResponseEnvelopeSuccess bool

const (
	WarpConnectorEditResponseEnvelopeSuccessTrue WarpConnectorEditResponseEnvelopeSuccess = true
)

type WarpConnectorGetResponseEnvelope struct {
	Errors   []WarpConnectorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WarpConnectorGetResponseEnvelopeMessages `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result WarpConnectorGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success WarpConnectorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    warpConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// warpConnectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WarpConnectorGetResponseEnvelope]
type warpConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    warpConnectorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// warpConnectorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WarpConnectorGetResponseEnvelopeErrors]
type warpConnectorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WarpConnectorGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    warpConnectorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// warpConnectorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WarpConnectorGetResponseEnvelopeMessages]
type warpConnectorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpConnectorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WarpConnectorGetResponseEnvelopeSuccess bool

const (
	WarpConnectorGetResponseEnvelopeSuccessTrue WarpConnectorGetResponseEnvelopeSuccess = true
)
