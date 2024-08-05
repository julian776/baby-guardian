package alerts

type Alerter interface {
	// Alert sends an alert message
	Alert(message string) error
}
