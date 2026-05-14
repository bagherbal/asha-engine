package theorem

// Status classifies the epistemic strength of a result.
//
// The engine must never label bridge-layer phenomenology as exact finite algebra.
type Status string

const (
	ExactFinite       Status = "EXACT_FINITE"
	Variational       Status = "VARIATIONAL"
	VerifiedNumeric   Status = "VERIFIED_NUMERIC"
	OpenTest          Status = "OPEN_TEST"
	BridgeRequired    Status = "BRIDGE_REQUIRED"
	Phenomenology     Status = "PHENOMENOLOGY"
	FailedRoute       Status = "FAILED_ROUTE"
	InvalidComparison Status = "INVALID_COMPARISON"
)
