package ngrok

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserAgent(t *testing.T) {
	s := (&clientInfo{Type: "library/official/go", Version: "1.2.3"}).ToUserAgent()
	require.Equal(t, s, "library-official-go/1.2.3")

	s = (&clientInfo{Type: "some@funky☺user agent", Version: "№1.2.3"}).ToUserAgent()
	require.Equal(t, s, "some#funky#user#agent/#1.2.3")

	s = (&clientInfo{
		Type:     "agent/official/go",
		Version:  "3.2.1",
		Comments: []string{"{\"ProxyType\": \"socks5\", \"ConfigVersion\": \"2\"}"},
	}).ToUserAgent()
	require.Equal(t, s, "agent-official-go/3.2.1 ({\"ProxyType\": \"socks5\", \"ConfigVersion\": \"2\"})")
}