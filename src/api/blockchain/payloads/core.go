package payloads

type Payload interface {
	Validate() (bool, string)
}
