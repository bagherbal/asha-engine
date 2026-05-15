// Package masterequationledger implements the ASHA master-equation boundary
// ledger. It is not a new gate and does not advance LatestGate. It makes the
// post-Gate-551 culmination explicit in code: the master equation is allowed as
// a structural law/history board only if bridge and environmental terms remain
// unable to write native ASHA law.
package masterequationledger

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/asha"
)

const (
	AuditID = "MASTER-EQUATION-GEOMETRY-HISTORY-BOUNDARY-LEDGER"

	StatusMasterEquationCompiled        = "CONDITIONAL_SUPPORT_MASTER_EQUATION_LEDGER_COMPILED"
	StatusNativeLawTermsClassified      = "CONDITIONAL_SUPPORT_NATIVE_GEOMETRIC_LAW_TERMS_CLASSIFIED"
	StatusEnvironmentalTermsQuarantined = "CONDITIONAL_SUPPORT_ENVIRONMENTAL_MODULI_TERMS_QUARANTINED"
	StatusBridgeAirlocksClassified      = "CONDITIONAL_SUPPORT_BRIDGE_AIRLOCK_TERMS_CLASSIFIED"
	StatusOSBoundaryPreserved           = "CONDITIONAL_SUPPORT_OS_WICK_HILBERT_BOUNDARY_PRESERVED"
	StatusNativeDeltaZeroVerified       = "CONDITIONAL_SUPPORT_MASTER_EQUATION_NATIVE_DELTA_ZERO_VERIFIED"
	StatusReadmeSectionReady            = "CONDITIONAL_SUPPORT_MASTER_EQUATION_README_SECTION_READY"

	StatusFailedFormulaNotToEClaim     = "FAILED_ROUTE_MASTER_EQUATION_IS_NOT_PARAMETER_FREE_NUMERICAL_TOE"
	StatusFailedDoesNotDeriveFlavor    = "FAILED_ROUTE_MASTER_EQUATION_DOES_NOT_DERIVE_FLAVOR_MODULI"
	StatusFailedDoesNotDeriveCutoff    = "FAILED_ROUTE_MASTER_EQUATION_DOES_NOT_DERIVE_CUTOFF_OR_NEWTON_NORMALIZATION"
	StatusFailedDoesNotSelectSpacetime = "FAILED_ROUTE_MASTER_EQUATION_DOES_NOT_SELECT_PHYSICAL_3PLUS1_SPACETIME"
	StatusFailedDoesNotGrantOS         = "FAILED_ROUTE_MASTER_EQUATION_DOES_NOT_GRANT_OS_WICK_HILBERT_DYNAMICS"
	StatusFailedDoesNotImportSchwinger = "FAILED_ROUTE_MASTER_EQUATION_DOES_NOT_IMPORT_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFirewallNativeWriteBlocked   = "FIREWALL_BLOCKED_MASTER_EQUATION_ENVIRONMENTAL_NATIVE_WRITE"
)

type Analysis struct {
	Ledger     asha.MasterEquationLedger
	Problems   []string
	Truth      string
	ReadmeText string
}

func BuildDefault() Analysis {
	ledger := asha.BuildMasterEquationLedger()
	return Analysis{
		Ledger:     ledger,
		Problems:   asha.ValidateMasterEquationLedger(ledger),
		Truth:      "The master equation is a boundary object: it summarizes ASHA law-space plus bridge/environmental history without promoting environmental data to native derivation.",
		ReadmeText: READMESection(),
	}
}

func Statuses() []string {
	return []string{
		StatusMasterEquationCompiled,
		StatusNativeLawTermsClassified,
		StatusEnvironmentalTermsQuarantined,
		StatusBridgeAirlocksClassified,
		StatusOSBoundaryPreserved,
		StatusNativeDeltaZeroVerified,
		StatusReadmeSectionReady,
		StatusFailedFormulaNotToEClaim,
		StatusFailedDoesNotDeriveFlavor,
		StatusFailedDoesNotDeriveCutoff,
		StatusFailedDoesNotSelectSpacetime,
		StatusFailedDoesNotGrantOS,
		StatusFailedDoesNotImportSchwinger,
		StatusFirewallNativeWriteBlocked,
	}
}

func FormatLedger(m asha.MasterEquationLedger) string {
	return fmt.Sprintf("formula=%s; native_terms=%d environmental_terms=%d bridge_terms=%d firewalls=%d native_delta_zero=%t", m.FormulaPlain, len(m.NativeTerms), len(m.EnvironmentalTerms), len(m.BridgeTerms), len(m.Firewalls), m.NativeDeltaZero)
}

func FormatTermClasses(m asha.MasterEquationLedger) string {
	var rows []string
	for _, t := range append(append([]asha.MasterEquationTerm{}, m.NativeTerms...), append(m.EnvironmentalTerms, m.BridgeTerms...)...) {
		rows = append(rows, fmt.Sprintf("%s:%s:native=%t:bridge=%t:env=%t:write=%t", t.Symbol, t.Class, t.NativeDerived, t.BridgeRequired, t.EnvironmentalInput, t.NativeWriteAllowed))
	}
	return strings.Join(rows, "; ")
}

func READMESection() string {
	return strings.TrimSpace(`
## Master Equation: Law-Space plus Environmental History

The compact boundary formula of the framework is:

$$
S_{Universe} = \text{Tr}\left( f\left(\frac{D^2}{\Lambda^2}\right) \right) + \langle \Psi, D \Psi \rangle_{OS}
$$

This is not a claim that ASHA predicts every observed number. It is the opposite: a claim-control equation that records exactly which parts are native geometric law and which parts require bridge or environmental input.

**Native geometric law.** The trace, the Clifford/product Dirac structure, and the matter-spinor carrier encode the finite law-space: anomaly cancellation capacity, triality/family architecture, Standard Model field-content sockets, spectral-action gravity shape, and finite stability ledgers.

**Environmental and bridge moduli.** The cutoff scale $\Lambda$, cutoff moments/function $f$, flavor entries hidden inside $D$, physical Schwinger functions, Wick/$i\epsilon$ convention, Osterwalder-Schrader positivity, Hilbert reconstruction, Hamiltonian spectrum, unitary dynamics, global causality, and the arrow of time are not native writes. They enter only through explicit airlocks and evidence ledgers.

The formula therefore means: ASHA derives a geometric law-space and a firewall-governed interface to physical history; it does not smuggle the Big Bang's choices into pure algebra.
`)
}

func MarkdownAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# ASHA Master Equation Registry Audit — Geometry/History Boundary Ledger\n\n")
	b.WriteString("## Formula\n\n")
	b.WriteString("$$\n")
	b.WriteString(a.Ledger.FormulaGitHubLaTeX)
	b.WriteString("\n$$\n\n")
	b.WriteString("## Ledger Summary\n\n")
	b.WriteString("```text\n")
	b.WriteString(FormatLedger(a.Ledger))
	b.WriteString("\n")
	b.WriteString(FormatTermClasses(a.Ledger))
	b.WriteString("\n```\n\n")
	b.WriteString("## Statuses\n\n")
	for _, status := range Statuses() {
		b.WriteString("- `")
		b.WriteString(status)
		b.WriteString("`\n")
	}
	b.WriteString("\n## Truth Boundary\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n")
	return b.String()
}
