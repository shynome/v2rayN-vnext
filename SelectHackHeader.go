package vnext

import (
	proto "github.com/golang/protobuf/proto"
	srtpHeader "v2ray.com/core/transport/internet/headers/srtp"
	tlsHeader "v2ray.com/core/transport/internet/headers/tls"
	utpHeader "v2ray.com/core/transport/internet/headers/utp"
	wechatHeader "v2ray.com/core/transport/internet/headers/wechat"
)

// SelectHackHeader select hack header
func SelectHackHeader(name string) (headerConfig proto.Message) {
	switch name {
	case "srtp":
		headerConfig = &srtpHeader.Config{}
	case "utp":
		headerConfig = &utpHeader.Config{}
	case "wechat-video":
		headerConfig = &wechatHeader.VideoConfig{}
	case "dtls":
		headerConfig = &tlsHeader.PacketConfig{}
	}
	return nil
}
