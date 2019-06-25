package vnext

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// New vnext string
func New(shareStr string) (vnext *VNEXT, err error) {
	s := strings.Split(shareStr, "://")
	if s[0] != "vmess" {
		err = errors.New("只支持 vmess 协议")
		return
	}
	base64str := s[1]
	jsonbyte, err := base64.StdEncoding.DecodeString(base64str)

	if err = json.Unmarshal(jsonbyte, &vnext); err != nil {
		return
	}
	return
}
