package vnext

import (
	"strings"

	core "v2ray.com/core"
	"v2ray.com/core/app/proxyman"
	net "v2ray.com/core/common/net"
	protocol "v2ray.com/core/common/protocol"
	serial "v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vmess"
	"v2ray.com/core/proxy/vmess/outbound"
	internet "v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/internet/http"
	"v2ray.com/core/transport/internet/kcp"
	"v2ray.com/core/transport/internet/tls"
	"v2ray.com/core/transport/internet/websocket"
)

// APIEmail of vnext
var APIEmail = ""

// NewVMessOutboundConfig GenVMessOutboundConfig
func (vnext VNEXT) NewVMessOutboundConfig(tag string) (config *core.OutboundHandlerConfig) {

	user := &protocol.User{
		Email: APIEmail,
		Account: serial.ToTypedMessage(&vmess.Account{
			Id:               vnext.ID,
			AlterId:          vnext.AlertID,
			SecuritySettings: &protocol.SecurityConfig{Type: protocol.SecurityType_AUTO},
		}),
	}

	proxyConfig := &outbound.Config{
		Receiver: []*protocol.ServerEndpoint{
			&protocol.ServerEndpoint{
				Address: net.NewIPOrDomain(net.ParseAddress(vnext.Address)),
				Port:    vnext.Port,
				User:    []*protocol.User{user},
			},
		},
	}

	var streamConfig *internet.StreamConfig

	switch vnext.Network {
	case "ws":
		{
			wsConfig := &websocket.Config{
				Path: vnext.Path,
				Header: []*websocket.Header{
					{Key: "Host", Value: vnext.Host},
				},
			}
			streamConfig = &internet.StreamConfig{
				Protocol: internet.TransportProtocol_WebSocket,
				TransportSettings: []*internet.TransportConfig{
					&internet.TransportConfig{
						Protocol: internet.TransportProtocol_WebSocket,
						Settings: serial.ToTypedMessage(wsConfig),
					},
				},
			}
		}
	case "kcp":
		{
			kcpConfig := &kcp.Config{
				HeaderConfig: serial.ToTypedMessage(SelectHackHeader(vnext.Type)),
			}
			streamConfig = &internet.StreamConfig{
				Protocol: internet.TransportProtocol_MKCP,
				TransportSettings: []*internet.TransportConfig{
					&internet.TransportConfig{
						Protocol: internet.TransportProtocol_MKCP,
						Settings: serial.ToTypedMessage(kcpConfig),
					},
				},
			}
		}
	case "tcp":
		{
			streamConfig = &internet.StreamConfig{
				Protocol: internet.TransportProtocol_TCP,
			}
		}
	case "h2":
		{
			httpConfig := &http.Config{
				Host: strings.Split(vnext.Host, ","),
				Path: vnext.Path,
			}
			streamConfig = &internet.StreamConfig{
				Protocol: internet.TransportProtocol_HTTP,
				TransportSettings: []*internet.TransportConfig{
					&internet.TransportConfig{
						Protocol: internet.TransportProtocol_HTTP,
						Settings: serial.ToTypedMessage(httpConfig),
					},
				},
			}
		}
	}

	if vnext.TLS == "tls" {
		streamConfig.SecurityType = serial.GetMessageType(&tls.Config{})
		streamConfig.SecuritySettings = []*serial.TypedMessage{
			serial.ToTypedMessage(&tls.Config{
				AllowInsecure: false,
			}),
		}
	}

	senderConfig := &proxyman.SenderConfig{
		StreamSettings: streamConfig,
	}

	config = &core.OutboundHandlerConfig{
		Tag:            tag,
		SenderSettings: serial.ToTypedMessage(senderConfig),
		ProxySettings:  serial.ToTypedMessage(proxyConfig),
	}

	return
}
