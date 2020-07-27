package common

type FeatureFlags struct {
	// Datasync indicates whether direct messages should be sent exclusively
	// using datasync, breaking change for non-v1 clients. Public messages
	// are not impacted
	Datasync bool
}
