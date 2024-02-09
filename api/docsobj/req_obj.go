package docsobj

type Request struct {
	// Input Long URL
	URL string `json:"url"`
	//Custom Short Endpoint (Automatically generated if not provided)
	CustomShort string `json:"short"`
}
