package asha

// TermClass records how a symbol in the ASHA master equation is allowed to be
// interpreted by the engine. The central rule is that structural law-space may
// be native, but realized physical history must enter through bridge or
// environmental airlocks.
type TermClass string

const (
	TermNativeLaw            TermClass = "NATIVE_GEOMETRIC_LAW"
	TermStructuralBoundary   TermClass = "STRUCTURAL_BOUNDARY_FORM"
	TermBridgeConvention     TermClass = "BRIDGE_CONVENTION"
	TermEnvironmentalModulus TermClass = "ENVIRONMENTAL_MODULUS"
	TermFirewallBoundary     TermClass = "FIREWALL_BOUNDARY"
)

const (
	// MasterEquationGitHubLaTeX is intentionally written with GitHub-supported
	// LaTeX commands only. Avoid \operatorname and other macros rejected by the
	// project README renderer.
	MasterEquationGitHubLaTeX = `S_{Universe} = \text{Tr}\left( f\left(\frac{D^2}{\Lambda^2}\right) \right) + \langle \Psi, D \Psi \rangle_{OS}`

	MasterEquationPlain = "S_Universe = Tr(f(D^2/Lambda^2)) + <Psi, D Psi>_OS"
)

// MasterEquationTerm is a single claim-controlled row in the master equation.
type MasterEquationTerm struct {
	Symbol             string    `json:"symbol"`
	Name               string    `json:"name"`
	Class              TermClass `json:"class"`
	Role               string    `json:"role"`
	NativeDerived      bool      `json:"native_derived"`
	BridgeRequired     bool      `json:"bridge_required"`
	EnvironmentalInput bool      `json:"environmental_input"`
	NativeWriteAllowed bool      `json:"native_write_allowed"`
	Firewall           string    `json:"firewall"`
}

// MasterEquationLedger is the explicit app-level ledger for the ASHA master
// equation. It is not a numerical theory-of-everything claim. It is a typed
// boundary object separating native law-space from environmental realization.
type MasterEquationLedger struct {
	FormulaGitHubLaTeX string               `json:"formula_github_latex"`
	FormulaPlain       string               `json:"formula_plain"`
	NativeTerms        []MasterEquationTerm `json:"native_terms"`
	EnvironmentalTerms []MasterEquationTerm `json:"environmental_terms"`
	BridgeTerms        []MasterEquationTerm `json:"bridge_terms"`
	Firewalls          []string             `json:"firewalls"`
	Message            string               `json:"message"`
	NativeDeltaZero    bool                 `json:"native_delta_zero"`
}

// BuildMasterEquationLedger returns the canonical ASHA master-equation boundary
// board used by documentation and app-level theorem checks.
func BuildMasterEquationLedger() MasterEquationLedger {
	return MasterEquationLedger{
		FormulaGitHubLaTeX: MasterEquationGitHubLaTeX,
		FormulaPlain:       MasterEquationPlain,
		NativeTerms: []MasterEquationTerm{
			{
				Symbol:             "S_{Universe}",
				Name:               "total action form",
				Class:              TermStructuralBoundary,
				Role:               "structural action container for spectral and fermionic sectors; the form is native, while a realized universe requires bridge/environmental instantiation",
				NativeDerived:      true,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "the action form is law-space; observed constants and universe history are not native writes",
			},
			{
				Symbol:             `\text{Tr}`,
				Name:               "trace over finite/product geometry",
				Class:              TermNativeLaw,
				Role:               "sums geometric degrees of freedom and supports exact finite trace, anomaly, and topological ledgers",
				NativeDerived:      true,
				BridgeRequired:     false,
				EnvironmentalInput: false,
				NativeWriteAllowed: true,
				Firewall:           "trace identities do not determine measured scales or physical source histories",
			},
			{
				Symbol:             "D",
				Name:               "total Dirac operator socket",
				Class:              TermNativeLaw,
				Role:               "finite/product geometric ruler carrying gauge, scalar, matter, and gravitational spectral-action sockets",
				NativeDerived:      true,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "structural Dirac sockets are native; flavor amplitudes, CKM/PMNS orientation, and physical spectra remain sealed",
			},
			{
				Symbol:             `\Psi`,
				Name:               "matter spinor carrier",
				Class:              TermNativeLaw,
				Role:               "fermionic matter representation carrier with triality/family architecture and anomaly ledgers",
				NativeDerived:      true,
				BridgeRequired:     false,
				EnvironmentalInput: false,
				NativeWriteAllowed: true,
				Firewall:           "matter-carrier structure does not fix observed masses or mixing angles",
			},
		},
		EnvironmentalTerms: []MasterEquationTerm{
			{
				Symbol:             `\Lambda`,
				Name:               "cutoff / physical scale",
				Class:              TermEnvironmentalModulus,
				Role:               "sets dimensional scale and gravity/cosmology normalization after bridge selection",
				NativeDerived:      false,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "cutoff scale, Newton normalization, and cosmological constant are not derived natively",
			},
			{
				Symbol:             "f",
				Name:               "cutoff function / spectral moments",
				Class:              TermBridgeConvention,
				Role:               "regularizes high-energy spectral data and supplies heat-kernel moment conventions",
				NativeDerived:      false,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "moment choices and renormalization conventions are bridge/environmental data",
			},
			{
				Symbol:             "D_{flavor}",
				Name:               "hidden flavor moduli inside D",
				Class:              TermEnvironmentalModulus,
				Role:               "Yukawa amplitudes, CKM/PMNS orientations, and family-sector coefficients",
				NativeDerived:      false,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "the flavor/firewall sector is environmental; native geometry supplies sockets and consistency constraints only",
			},
			{
				Symbol:             "OS",
				Name:               "Osterwalder-Schrader / Wick / Hilbert / time-arrow airlock",
				Class:              TermFirewallBoundary,
				Role:               "protects the transition from Euclidean law-space to Lorentzian quantum dynamics",
				NativeDerived:      false,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "OS positivity, Wick/iε convention, physical Hilbert reconstruction, Hamiltonian spectrum, unitarity, global causality, and the arrow of time require sourced bridge evidence",
			},
		},
		BridgeTerms: []MasterEquationTerm{
			{
				Symbol:             "3+1",
				Name:               "spacetime projection socket",
				Class:              TermBridgeConvention,
				Role:               "an environmental dimensional split may be housed and checked algebraically but is not selected natively",
				NativeDerived:      false,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "projector compatibility does not grant Wick rotation, Hilbert positivity, or internal gauge identification",
			},
			{
				Symbol:             "S_n",
				Name:               "physical Schwinger functions",
				Class:              TermEnvironmentalModulus,
				Role:               "Euclidean correlation data required before physical OS/Wick/Hilbert/Hamiltonian reconstruction can even be compared",
				NativeDerived:      false,
				BridgeRequired:     true,
				EnvironmentalInput: true,
				NativeWriteAllowed: false,
				Firewall:           "correlation functions are source/evidence data, not native ASHA law",
			},
		},
		Firewalls: []string{
			"MASTER_EQUATION_FORM_IS_NOT_PARAMETER_FREE_NUMERICAL_COMPLETION",
			"FIREWALL_BLOCKS_FLAVOR_MODULI_NATIVE_WRITE",
			"FIREWALL_BLOCKS_CUTOFF_AND_GRAVITY_NORMALIZATION_NATIVE_WRITE",
			"FIREWALL_BLOCKS_DIMENSIONAL_SELECTION_NATIVE_WRITE",
			"FIREWALL_BLOCKS_OS_WICK_HILBERT_HAMILTONIAN_NATIVE_WRITE",
			"FIREWALL_BLOCKS_PHYSICAL_SCHWINGER_CORRELATION_NATIVE_WRITE",
			"FIREWALL_REQUIRES_NATIVE_DELTA_ZERO_FOR_BRIDGE_EVIDENCE",
		},
		Message:         "The master equation proves the boundary between what ASHA can derive as finite geometric law-space and what must be supplied as bridge/environmental universe history.",
		NativeDeltaZero: true,
	}
}

// ValidateMasterEquationLedger performs the app-level consistency checks that
// keep the master equation from becoming an overclaim.
func ValidateMasterEquationLedger(m MasterEquationLedger) []string {
	var problems []string
	if m.FormulaGitHubLaTeX != MasterEquationGitHubLaTeX {
		problems = append(problems, "master equation latex mismatch")
	}
	if len(m.NativeTerms) < 4 {
		problems = append(problems, "native law ledger is incomplete")
	}
	if len(m.EnvironmentalTerms) < 4 {
		problems = append(problems, "environmental/firewall ledger is incomplete")
	}
	if len(m.BridgeTerms) < 2 {
		problems = append(problems, "bridge ledger is incomplete")
	}
	if !m.NativeDeltaZero {
		problems = append(problems, "native delta must remain zero for bridge/environmental evidence")
	}
	for _, term := range append(append([]MasterEquationTerm{}, m.EnvironmentalTerms...), m.BridgeTerms...) {
		if term.NativeWriteAllowed {
			problems = append(problems, "bridge/environmental term allows native write: "+term.Symbol)
		}
		if !term.BridgeRequired && term.EnvironmentalInput {
			problems = append(problems, "environmental term lacks bridge requirement: "+term.Symbol)
		}
	}
	return problems
}
