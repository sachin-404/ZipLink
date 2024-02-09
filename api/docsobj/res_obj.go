package docsobj

import "time"

type ExpiryDuration time.Duration

type Response struct {
	//Original URL
	URL string `json:"url"`
	//Shorten URL
	CustomShort string `json:"short"`
	// Expiry Time (hr)
	Expiry          ExpiryDuration `json:"expiry"`
	XRateRemaining  int            `json:"rate_limit"`
	XRateLimitReset ExpiryDuration `json:"rate_limit_reset"`
}
