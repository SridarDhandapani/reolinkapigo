package examples

import (
	rest2 "github.com/ReolinkCameraAPI/reolinkapigo/pkg/network/rest"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
)

func Socks5Example() {

	// This can throw an error due to the API trying to authorise with the camera
	camera, err := reolinkapi.NewCamera("192.168.1.100",
		reolinkapi.WithUsername("foo"),
		reolinkapi.WithPassword("bar"),
		reolinkapi.WithNetworkOptions(
			rest2.WithProxyScheme(rest2.SOCKS5),
			rest2.WithProxyHost("127.0.0.1"),
			rest2.WithProxyPort(5942),
			rest2.WithProxyUsername("foo"),
			rest2.WithProxyPassword("bar"),
		))

	if err != nil {
		panic(err)
	}

	// Call your camera api here and pass the camera restHandler to the function
	ok, err := camera.FormatHdd(0)(camera.RestHandler)

	if err != nil {
		panic(err)
	}

	if ok {
		print("Format OK")
	}
}
