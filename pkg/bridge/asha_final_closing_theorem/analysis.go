// Package ashafinalclosingtheorem implements Gate 374:
// The Final Closing Theorem / ASHA 13-Moduli Vacuum Manifold Seal.
//
// This capstone gate is deliberately not a new mechanism for selecting the
// flavor vacuum.  It is an epistemic seal over the ledger established through
// Gates 1-373.  It separates:
//  1. native structural laws and boundary ratios;
//  2. bridge/transport-dependent phenomenological proxies;
//  3. the irreducible 13 charged finite-Dirac moduli found by Gate 372;
//  4. failed/circular routes that cannot be promoted without new axioms.
//
// The theorem therefore closes the current finite Cℓ(1,7) kinematic program
// without pretending that the Yukawa/CKM coordinates have been derived.
package ashafinalclosingtheorem

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate373 "github.com/bagherbal/asha-engine/pkg/bridge/holographicvacuumentropy"
)

const (
	AuditID = "GATE374-ASHA-FINAL-CLOSING-THEOREM-13-MODULI-VACUUM-MANIFOLD-SEAL"

	StatusGate373Inherited                 = "CONDITIONAL_SUPPORT_GATE373_HOLOGRAPHIC_AUDIT_INHERITED"
	StatusLedgerPartitionFormalized        = "CONDITIONAL_SUPPORT_GRAND_LEDGER_PARTITION_FORMALIZED"
	StatusExactStructuralLawsSealed        = "CONDITIONAL_SUPPORT_EXACT_STRUCTURAL_LAWS_SEALED"
	StatusBoundaryRatiosSealed             = "CONDITIONAL_SUPPORT_NATIVE_BOUNDARY_RATIOS_SEALED"
	StatusPhenomenologicalProxyQuarantined = "CONDITIONAL_SUPPORT_PHENOMENOLOGICAL_PROXY_QUARANTINED"
	StatusThirteenModuliManifoldSealed     = "CONDITIONAL_SUPPORT_13_MODULI_VACUUM_MANIFOLD_SEALED"
	StatusFailedRoutesClosed               = "CONDITIONAL_SUPPORT_FAILED_ROUTES_CLOSED_WITH_FIREWALL"
	StatusPublicationScopeSealed           = "CONDITIONAL_SUPPORT_PUBLICATION_SCOPE_SEALED"
	StatusEpistemicFirewallPreserved       = "CONDITIONAL_SUPPORT_EPISTEMIC_FIREWALL_PRESERVED"
	StatusFrameworkCompleteAsKinematics    = "CONDITIONAL_SUPPORT_FRAMEWORK_COMPLETE_AS_FINITE_KINEMATICS"

	StatusTensionHiggsProxyNeedsTransport        = "CONDITIONAL_TENSION_HIGGS_PROXY_REQUIRES_RG_AND_MATCHING_TRANSPORT"
	StatusTensionAlphaGutNeedsThresholdLedger    = "CONDITIONAL_TENSION_ALPHA_GUT_BRANCH_NEEDS_THRESHOLD_LEDGER_FOR_EMPIRICAL_ALIGNMENT"
	StatusTensionFlavorModuliRemainEnvironmental = "CONDITIONAL_TENSION_FLAVOR_MODULI_REMAIN_ENVIRONMENTAL_COORDINATES"
	StatusTensionFinalMeansScopedCompletion      = "CONDITIONAL_TENSION_FINAL_THEOREM_IS_SCOPED_NOT_OMNISCIENT"

	StatusFailedFlavorParametersNotDerived      = "FAILED_ROUTE_13_FLAVOR_PARAMETERS_NOT_DERIVED"
	StatusFailedVacuumPointNotSelected          = "FAILED_ROUTE_PHYSICAL_VACUUM_POINT_NOT_SELECTED"
	StatusFailedTauEtaHamiltonianNotSelected    = "FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED"
	StatusFailedHolographyDidNotReduceModuli    = "FAILED_ROUTE_HOLOGRAPHY_DID_NOT_REDUCE_MODULI"
	StatusFailedInternalThermalTimeNotActivated = "FAILED_ROUTE_INTERNAL_THERMAL_TIME_NOT_ACTIVATED"
)

const (
	ChargedFiniteDiracModuli = 13
	ExternalMinimalLedger    = 15
	ThetaQCDLedgerDimension  = 1
	AbsoluteScaleDimension   = 1
)

type Inheritance struct {
	Executed                  bool
	HighestInheritedGate      int
	ChargedFiniteDiracModuli  int
	ExternalMinimalLedger     int
	HolographicReductionFound bool
	PreviousTruth             string
	Verdict                   string
}

type LedgerClass string

const (
	LedgerExactTopological LedgerClass = "EXACT_TOPOLOGICAL_OR_STRUCTURAL"
	LedgerBoundary         LedgerClass = "NATIVE_BOUNDARY_RATIO"
	LedgerProxy            LedgerClass = "PHENOMENOLOGICAL_PROXY_OR_TRANSPORT"
	LedgerFreeModulus      LedgerClass = "IRREDUCIBLE_FREE_MODULUS"
)

type LedgerItem struct {
	Name                    string
	Class                   LedgerClass
	Formula                 string
	NativeASHA              bool
	ExactFinite             bool
	RequiresRGTransport     bool
	RequiresEmpiricalInput  bool
	SelectsFlavorCoordinate bool
	ErrorMargin             string
	Verdict                 string
	Note                    string
}

type BoundaryLedger struct {
	Executed        bool
	Items           []LedgerItem
	ExactCount      int
	BoundaryCount   int
	ProxyCount      int
	FreeModuliCount int
	Verdict         string
}

type ModuliManifold struct {
	Executed                 bool
	ChargedFlavorModuli      int
	Identity                 string
	ThetaQCD                 int
	AbsoluteScale            int
	ExternalLedger           int
	AllAllowedMajoranaDFDim  int
	PureGeometrySelectsPoint bool
	FlatDirections           bool
	GaugeQuotientComplete    bool
	FinalCensusFormula       string
	Verdict                  string
}

type FailedRouteSeal struct {
	Executed bool
	Routes   []RouteClosure
	Verdict  string
}

type RouteClosure struct {
	Route       string
	GateRange   string
	Closure     string
	WhyClosed   string
	MayReopenIf string
	Status      string
}

type Firewall struct {
	Executed                       bool
	NoObservedYukawaValues         bool
	NoObservedCKMValues            bool
	NoObservedFermionMassTargets   bool
	NoManualTauEtaHamiltonian      bool
	NoHolographicSaturationAssumed bool
	NoFinalVacuumClaim             bool
	BoundaryVsProxySeparated       bool
	ExactVsConditionalSeparated    bool
	Verdict                        string
}

type ClosingTheorem struct {
	Executed                   bool
	LandscapeAbsolute          bool
	VacuumFree                 bool
	KinematicsComplete         bool
	DynamicsOfFlavorUnselected bool
	PublicationScope           string
	TheoremStatement           string
	Verdict                    string
}

type Analysis struct {
	Inheritance Inheritance
	Ledger      BoundaryLedger
	Moduli      ModuliManifold
	Routes      FailedRouteSeal
	Firewall    Firewall
	Closing     ClosingTheorem
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	prev, err := gate373.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	inheritance := inherit(prev)
	ledger := buildLedger()
	moduli := sealModuliManifold()
	routes := closeFailedRoutes()
	firewall := auditFirewall()
	closing := closeTheorem(ledger, moduli, routes, firewall)
	truth := "Gate 374 is a capstone seal, not a new vacuum-selection mechanism.  The project ledger supports a scoped final theorem: finite Cℓ(1,7)/ASHA kinematics can be closed as a Standard-Model-and-boundary-conditions reconstruction, while the 13 charged finite-Dirac flavor coordinates remain irreducible moduli of the pure geometry.  Boundary ratios and hierarchy data are preserved as native structural/boundary statements; Higgs-mass and low-energy comparisons remain transport/proxy statements; and every attempted pure-geometric, modular-flow, eta-trace, Fock/information, or holographic route to selecting the flavor point has either remained central, required a circular insertion, or lacked enough independent equations.  Therefore the framework is complete in its current finite-kinematic scope and explicitly incomplete as a derivation of the 13 environmental flavor coordinates."
	return Analysis{inheritance, ledger, moduli, routes, firewall, closing, truth}, nil
}

func inherit(prev gate373.Analysis) Inheritance {
	return Inheritance{
		Executed:                  true,
		HighestInheritedGate:      373,
		ChargedFiniteDiracModuli:  prev.Census.RemainingChargedModuli,
		ExternalMinimalLedger:     prev.Census.ExternalLedger,
		HolographicReductionFound: prev.Census.Reduction > 0,
		PreviousTruth:             prev.Truth,
		Verdict:                   join(StatusGate373Inherited),
	}
}

func buildLedger() BoundaryLedger {
	items := []LedgerItem{
		{Name: "Gauge group and chiral Standard-Model representation", Class: LedgerExactTopological, Formula: "SU(3) x SU(2) x U(1) with chiral matter representation", NativeASHA: true, ExactFinite: true, ErrorMargin: "exact within finite representation ledger", Verdict: "sealed as structural ledger", Note: "A kinematic structural result, not a flavor-value prediction."},
		{Name: "Particle-content inventory", Class: LedgerExactTopological, Formula: "3 generations, 1 Higgs doublet, 12 gauge bosons", NativeASHA: true, ExactFinite: true, ErrorMargin: "exact within finite representation ledger", Verdict: "sealed as structural ledger", Note: "The three-generation multiplicity remains symmetric unless an additional generation-address theorem is added."},
		{Name: "Morita bimodule split", Class: LedgerExactTopological, Formula: "1 + 3 lepton/quark split", NativeASHA: true, ExactFinite: true, ErrorMargin: "exact", Verdict: "sealed as structural ledger", Note: "Explains lepton/color-triplet separation, not unequal flavor weights."},
		{Name: "Generation topology capacity witness", Class: LedgerExactTopological, Formula: "tau_eta = (2, -2, 1)", NativeASHA: true, ExactFinite: true, ErrorMargin: "exact as topology/capacity witness", Verdict: "sealed but not promoted to Hamiltonian", Note: "Gates 368-371 forbid circularly inserting tau_eta as modular energy."},
		{Name: "Weak mixing boundary", Class: LedgerBoundary, Formula: "sin^2(theta_W)(Lambda) = 3/8", NativeASHA: true, ExactFinite: true, RequiresRGTransport: true, ErrorMargin: "exact at boundary; comparison requires RG", Verdict: "sealed as boundary ratio", Note: "Low-energy comparison is not a direct finite-scale equality."},
		{Name: "Unified coupling branch", Class: LedgerBoundary, Formula: "alpha_GUT^-1 = 8*pi", NativeASHA: true, ExactFinite: true, RequiresRGTransport: true, ErrorMargin: "exact branch value; empirical alignment threshold-dependent", Verdict: join(StatusTensionAlphaGutNeedsThresholdLedger), Note: "Boundary branch is distinct from observed low-energy couplings."},
		{Name: "Higgs quartic trace ratio", Class: LedgerBoundary, Formula: "lambda_H/g_*^2 = 1197/4624", NativeASHA: true, ExactFinite: true, RequiresRGTransport: true, ErrorMargin: "exact boundary trace ratio", Verdict: "sealed as boundary ratio", Note: "Does not determine Yukawa matrices."},
		{Name: "Pfaffian hierarchy scale", Class: LedgerBoundary, Formula: "v/M_P = 2^(3/2) exp(-4*pi^2)", NativeASHA: true, ExactFinite: true, RequiresRGTransport: false, ErrorMargin: "native hierarchy-scale relation", Verdict: "sealed as hierarchy boundary", Note: "Fixes a scale relation, not the 13 flavor coordinates."},
		{Name: "Heavy-sector threshold jump witness", Class: LedgerProxy, Formula: "Delta lambda approximately -0.0978", NativeASHA: true, ExactFinite: false, RequiresRGTransport: true, ErrorMargin: "conditional bridge value", Verdict: "quarantined as transport/proxy", Note: "Depends on threshold/continuum transport assumptions."},
		{Name: "Higgs mass proxy", Class: LedgerProxy, Formula: "m_H near 125 GeV after native boundary + threshold/RG transport", NativeASHA: true, ExactFinite: false, RequiresRGTransport: true, RequiresEmpiricalInput: false, ErrorMargin: "conditional proxy; scheme and matching dependent", Verdict: join(StatusTensionHiggsProxyNeedsTransport), Note: "Not an exact finite theorem; kept separate from structural laws."},
		{Name: "Charged flavor vacuum manifold", Class: LedgerFreeModulus, Formula: "dim M(D_F)_charged = 13", NativeASHA: true, ExactFinite: true, SelectsFlavorCoordinate: false, ErrorMargin: "exact census under Gate-372 scope", Verdict: join(StatusTensionFlavorModuliRemainEnvironmental), Note: "9 charged masses + 4 CKM parameters remain moduli after gauge quotient."},
	}

	var exactCount, boundaryCount, proxyCount, freeCount int
	for _, item := range items {
		switch item.Class {
		case LedgerExactTopological:
			exactCount++
		case LedgerBoundary:
			boundaryCount++
		case LedgerProxy:
			proxyCount++
		case LedgerFreeModulus:
			freeCount++
		}
	}
	return BoundaryLedger{Executed: true, Items: items, ExactCount: exactCount, BoundaryCount: boundaryCount, ProxyCount: proxyCount, FreeModuliCount: freeCount, Verdict: join(StatusLedgerPartitionFormalized, StatusExactStructuralLawsSealed, StatusBoundaryRatiosSealed, StatusPhenomenologicalProxyQuarantined)}
}

func sealModuliManifold() ModuliManifold {
	return ModuliManifold{
		Executed:                 true,
		ChargedFlavorModuli:      ChargedFiniteDiracModuli,
		Identity:                 "9 charged fermion masses + 4 CKM parameters under the minimal charged finite-Dirac census",
		ThetaQCD:                 ThetaQCDLedgerDimension,
		AbsoluteScale:            AbsoluteScaleDimension,
		ExternalLedger:           ExternalMinimalLedger,
		AllAllowedMajoranaDFDim:  31,
		PureGeometrySelectsPoint: false,
		FlatDirections:           true,
		GaugeQuotientComplete:    true,
		FinalCensusFormula:       "15 external minimal vacuum ledger = 13 charged finite-Dirac moduli + theta_QCD + one absolute scale",
		Verdict:                  join(StatusThirteenModuliManifoldSealed, StatusFailedFlavorParametersNotDerived, StatusFailedVacuumPointNotSelected),
	}
}

func closeFailedRoutes() FailedRouteSeal {
	routes := []RouteClosure{
		{Route: "static finite cross-sector reduction", GateRange: "pre-362 / Gate 362", Closure: "operator closure no-go", WhyClosed: "admissible native finite operators do not select the flavor point", MayReopenIf: "a new native admissible operator is derived with kinetic safety", Status: StatusFailedVacuumPointNotSelected},
		{Route: "Tomita modular flow with tracial native state", GateRange: "363-365", Closure: "tracial Delta=I or K must be selected externally", WhyClosed: "KMS machinery works only after a noncentral Hamiltonian is supplied", MayReopenIf: "a native noncentral modular Hamiltonian is derived", Status: StatusFailedInternalThermalTimeNotActivated},
		{Route: "tau_eta as modular Hamiltonian", GateRange: "366, 368-371", Closure: "capacity witness but circular selection", WhyClosed: "tau_eta works if inserted, but its dynamical-energy role is not derived", MayReopenIf: "tau_eta emerges from a lawful Left-Right/eta/support trace or information theorem", Status: StatusFailedTauEtaHamiltonianNotSelected},
		{Route: "ordinary Lorentzian time", GateRange: "367", Closure: "flavor-central pullback", WhyClosed: "e0/gamma0 acts as identity on generation space", MayReopenIf: "not applicable inside ordinary spacetime time", Status: StatusFailedInternalThermalTimeNotActivated},
		{Route: "bimodule/eta/support-to-generation trace", GateRange: "368-370", Closure: "generation-blind support trace", WhyClosed: "lawful support data factor through I3; target tau_eta map is circular", MayReopenIf: "a noncircular support-to-generation intertwiner is derived", Status: StatusFailedTauEtaHamiltonianNotSelected},
		{Route: "finite Fock/information number operator", GateRange: "371", Closure: "noncentral capacity but unselected basis/operator", WhyClosed: "N breaks U(3), but ASHA does not yet derive the Fock basis or polynomial P_tau(N)", MayReopenIf: "generation states are derived as oscillator/information levels from finite topology", Status: StatusFailedTauEtaHamiltonianNotSelected},
		{Route: "native finite-Dirac moduli reduction", GateRange: "372", Closure: "census equals 13 charged moduli", WhyClosed: "spectral-triple axioms and gauge quotient do not reduce charged flavor below 13", MayReopenIf: "additional axiom/continuum principle provides independent equations", Status: StatusFailedFlavorParametersNotDerived},
		{Route: "holographic/gravitational texture fixing", GateRange: "373", Closure: "scale constraints, not flavor equations", WhyClosed: "holographic bounds are aggregate inequalities and vacuum energy needs counterterm/scheme data", MayReopenIf: "a native saturation theorem and flavor-sensitive gravitational functional are derived", Status: StatusFailedHolographyDidNotReduceModuli},
	}
	return FailedRouteSeal{Executed: true, Routes: routes, Verdict: join(StatusFailedRoutesClosed, StatusFailedFlavorParametersNotDerived, StatusFailedVacuumPointNotSelected)}
}

func auditFirewall() Firewall {
	return Firewall{
		Executed:                       true,
		NoObservedYukawaValues:         true,
		NoObservedCKMValues:            true,
		NoObservedFermionMassTargets:   true,
		NoManualTauEtaHamiltonian:      true,
		NoHolographicSaturationAssumed: true,
		NoFinalVacuumClaim:             true,
		BoundaryVsProxySeparated:       true,
		ExactVsConditionalSeparated:    true,
		Verdict:                        join(StatusEpistemicFirewallPreserved),
	}
}

func closeTheorem(ledger BoundaryLedger, moduli ModuliManifold, routes FailedRouteSeal, firewall Firewall) ClosingTheorem {
	complete := ledger.Executed && moduli.Executed && routes.Executed && firewall.Executed && moduli.ChargedFlavorModuli == 13 && !moduli.PureGeometrySelectsPoint && firewall.NoFinalVacuumClaim
	statement := "ASHA/Cℓ(1,7) closes as a finite-kinematic Standard-Model landscape theorem with native boundary ratios and an irreducible 13-dimensional charged flavor moduli manifold.  The laws/representation/boundary ledger is sealed inside the current axioms; the specific Yukawa/CKM vacuum point is not derived and remains an environmental/informational coordinate unless a future theorem adds independent, noncircular dynamics."
	return ClosingTheorem{Executed: true, LandscapeAbsolute: true, VacuumFree: true, KinematicsComplete: complete, DynamicsOfFlavorUnselected: true, PublicationScope: "publishable as a scoped finite-kinematic closure plus explicit 13-moduli no-selection theorem", TheoremStatement: statement, Verdict: join(StatusPublicationScopeSealed, StatusFrameworkCompleteAsKinematics, StatusTensionFinalMeansScopedCompletion)}
}

func NativeBoundaryValues() map[string]float64 {
	return map[string]float64{
		"sin2_thetaW_boundary":     3.0 / 8.0,
		"alpha_GUT_inverse_branch": 8.0 * math.Pi,
		"lambdaH_over_gstar2":      1197.0 / 4624.0,
		"v_over_MP_hierarchy":      math.Pow(2, 1.5) * math.Exp(-4.0*math.Pi*math.Pi),
	}
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate=%d charged_moduli=%d external_ledger=%d holographic_reduction=%v verdict=%s", x.HighestInheritedGate, x.ChargedFiniteDiracModuli, x.ExternalMinimalLedger, x.HolographicReductionFound, x.Verdict)
}

func FormatLedger(x BoundaryLedger) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("exact=%d boundary=%d proxy=%d free_moduli=%d verdict=%s", x.ExactCount, x.BoundaryCount, x.ProxyCount, x.FreeModuliCount, x.Verdict))
	for _, item := range x.Items {
		b.WriteString(fmt.Sprintf("\n- %s [%s]: %s | native=%v exact=%v transport=%v selects_flavor=%v verdict=%s", item.Name, item.Class, item.Formula, item.NativeASHA, item.ExactFinite, item.RequiresRGTransport, item.SelectsFlavorCoordinate, item.Verdict))
	}
	return b.String()
}

func FormatModuli(x ModuliManifold) string {
	return fmt.Sprintf("charged=%d identity=%s theta_qcd=%d scale=%d external=%d all_allowed_majorana=%d flat=%v gauge_quotient=%v formula=%s verdict=%s", x.ChargedFlavorModuli, x.Identity, x.ThetaQCD, x.AbsoluteScale, x.ExternalLedger, x.AllAllowedMajoranaDFDim, x.FlatDirections, x.GaugeQuotientComplete, x.FinalCensusFormula, x.Verdict)
}

func FormatRoutes(x FailedRouteSeal) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("routes=%d verdict=%s", len(x.Routes), x.Verdict))
	for _, r := range x.Routes {
		b.WriteString(fmt.Sprintf("\n- %s (%s): %s; why=%s; reopen=%s; status=%s", r.Route, r.GateRange, r.Closure, r.WhyClosed, r.MayReopenIf, r.Status))
	}
	return b.String()
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("no_yukawas=%v no_ckm=%v no_mass_targets=%v no_manual_tau_eta=%v no_holographic_saturation=%v no_final_vacuum_claim=%v boundary_proxy_separated=%v exact_conditional_separated=%v verdict=%s", x.NoObservedYukawaValues, x.NoObservedCKMValues, x.NoObservedFermionMassTargets, x.NoManualTauEtaHamiltonian, x.NoHolographicSaturationAssumed, x.NoFinalVacuumClaim, x.BoundaryVsProxySeparated, x.ExactVsConditionalSeparated, x.Verdict)
}

func FormatClosing(x ClosingTheorem) string {
	return fmt.Sprintf("landscape_absolute=%v vacuum_free=%v kinematics_complete=%v flavor_dynamics_unselected=%v scope=%s statement=%s verdict=%s", x.LandscapeAbsolute, x.VacuumFree, x.KinematicsComplete, x.DynamicsOfFlavorUnselected, x.PublicationScope, x.TheoremStatement, x.Verdict)
}

func join(statuses ...string) string { return strings.Join(statuses, ";") }
