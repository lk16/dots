package playok

// Message is a json model for any sent or received JSON message
type Message struct {

	// I leads with an integer identifying the message type
	I []int `json:"i"`

	// S contains string data and is sometimes omitted
	S []string `json:"s,omitempty"`
}
