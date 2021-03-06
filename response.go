package aftership

// Response is the message envelope for the AfterShip API response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta is used to communicate extra information about the response to the developer.
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
