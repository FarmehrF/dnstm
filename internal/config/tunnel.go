package config

// TransportType defines the type of transport.
type TransportType string

const (
	TransportSlipstream TransportType = "slipstream"
	TransportDNSTT      TransportType = "dnstt"
	TransportNoizDNS    TransportType = "noizdns"
)

// TunnelConfig configures a DNS tunnel.
type TunnelConfig struct {
	Tag        string            `json:"tag"`
	Enabled    *bool             `json:"enabled,omitempty"`
	Transport  TransportType     `json:"transport"`
	Backend    string            `json:"backend"`
	Domain     string            `json:"domain"`
	Port       int               `json:"port,omitempty"`
	Slipstream *SlipstreamConfig `json:"slipstream,omitempty"`
	DNSTT      *DNSTTConfig      `json:"dnstt,omitempty"`
	NoizDNS    *DNSTTConfig      `json:"noizdns,omitempty"`
}

// SlipstreamConfig holds Slipstream-specific configuration.
type SlipstreamConfig struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

// DNSTTConfig holds DNSTT-specific configuration.
type DNSTTConfig struct {
	MTU        int    `json:"mtu,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
}

// IsEnabled returns true if the tunnel is enabled.
func (t *TunnelConfig) IsEnabled() bool {
	return t.Enabled == nil || *t.Enabled
}

// GetMTU returns the MTU for DNSTT/NoizDNS tunnels, with a default of 1232.
func (t *TunnelConfig) GetMTU() int {
	if t.DNSTT != nil && t.DNSTT.MTU > 0 {
		return t.DNSTT.MTU
	}
	if t.NoizDNS != nil && t.NoizDNS.MTU > 0 {
		return t.NoizDNS.MTU
	}
	return 1232 // Default
}

// IsSlipstream returns true if this is a Slipstream tunnel.
func (t *TunnelConfig) IsSlipstream() bool {
	return t.Transport == TransportSlipstream
}

// IsDNSTT returns true if this is a DNSTT tunnel.
func (t *TunnelConfig) IsDNSTT() bool {
	return t.Transport == TransportDNSTT
}

// IsNoizDNS returns true if this is a NoizDNS tunnel.
func (t *TunnelConfig) IsNoizDNS() bool {
	return t.Transport == TransportNoizDNS
}

// GetTransportTypes returns all available transport types.
func GetTransportTypes() []TransportType {
	return []TransportType{
		TransportSlipstream,
		TransportDNSTT,
		TransportNoizDNS,
	}
}

// GetTransportTypeDisplayName returns a human-readable name for a transport type.
func GetTransportTypeDisplayName(t TransportType) string {
	switch t {
	case TransportSlipstream:
		return "Slipstream"
	case TransportDNSTT:
		return "DNSTT"
	case TransportNoizDNS:
		return "NoizDNS"
	default:
		return string(t)
	}
}
