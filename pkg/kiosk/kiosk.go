package kiosk

type KioskOptions struct {
	serial    string
	uuid      string
	api_token string
}

type Kiosk struct {
	KioskOptions
}

func defaultKioskOptions() KioskOptions {
	return KioskOptions{
		serial:    "QMSK-V01-B002-N0020",
		uuid:      "ce2f587d-945c-11ec-8f63-782bcb0a49b7",
		api_token: "73e5c2f5969b27d1c6c0d2d58cfa71d5b1a29b1b",
	}
}
