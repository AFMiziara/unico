package ip

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutboundIP(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	_, err := OutboundIP()
	assert.NoError(t, err, "Erro na busca por IP - %v", err)
}

func TestGetLocalIPs(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	_, err := GetLocalIPs()
	assert.NoError(t, err, "Erro na busca por IP - %v", err)
}

func TestIsPublicIPFalse(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	retBool := IsPublicIP(net.ParseIP("169.254.85.131"))
	assert.Equal(t, false, retBool)
}

func TestIsPublicIP1False(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	retBool := IsPublicIP(net.ParseIP("10.0.0.1"))
	assert.Equal(t, false, retBool)
}

func TestIsPublicIP2False(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	retBool := IsPublicIP(net.ParseIP("172.16.0.1"))
	assert.Equal(t, false, retBool)
}

func TestIsPublicIP3False(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	retBool := IsPublicIP(net.ParseIP("192.168.0.1"))
	assert.Equal(t, false, retBool)
}

func TestIsPublicIPTrue(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	retBool := IsPublicIP(net.ParseIP("189.37.76.207"))
	assert.Equal(t, true, retBool)
}

func TestIpBetween(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	t.Run("IpBetweenTrue", func(t *testing.T) {
		from := net.ParseIP("189.0.0.0")
		to := net.ParseIP("189.255.255.255")
		test := net.ParseIP("189.37.76.207")

		retBool := IpBetween(from, to, test)
		assert.Equal(t, true, retBool)
	})

	t.Run("IpBetweenfalse", func(t *testing.T) {
		from := net.ParseIP("189.0.0.0")
		to := net.ParseIP("189.255.255.255")
		test := net.ParseIP("8.8.8.8")

		retBool := IpBetween(from, to, test)
		assert.Equal(t, false, retBool)
	})

	t.Run("IpBetweenfalseByNullFrom", func(t *testing.T) {
		from := net.ParseIP("")
		to := net.ParseIP("189.255.255.255")
		test := net.ParseIP("8.8.8.8")

		retBool := IpBetween(from, to, test)
		assert.Equal(t, false, retBool)
	})

}

func TestInet(t *testing.T) {

	os.Setenv("IP_PORT_DEFAULT", "8.8.8.8:80")

	t.Run("InetNtoa", func(t *testing.T) {
		ipInt := int64(3173338319)

		ipResult := InetNtoa(ipInt)
		assert.Equal(t, net.ParseIP("189.37.76.207"), ipResult)
	})

	t.Run("InetAton", func(t *testing.T) {
		test := net.ParseIP("189.37.76.207")

		ipInt := InetAton(test)
		assert.Equal(t, int64(3173338319), ipInt)
	})

}
