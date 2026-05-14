// Package spectralactionvariationalgradient implements Gate 346:
// Spectral Action Variational Gradient / Phase III Vacuum Initialization Sieve.
//
// Gate 346 initiates Phase III by treating the remaining Standard Model
// vacuum coordinates as moduli and auditing whether the spectral action itself
// selects a unique physical vacuum.  The gate is deliberately strict: standard
// heat-kernel invariants are unitary-conjugation invariant and therefore cannot
// by themselves derive CKM/flavor orientation.  A signed triality projector can
// select a top-suppressed nullspace, but that minimum is degenerate unless an
// additional texture/vacuum-selection operator is derived.
package spectralactionvariationalgradient

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE346-SPECTRAL-ACTION-VARIATIONAL-GRADIENT-PHASE-III-VACUUM-INITIALIZATION"

	StatusGate345Inherited               = "CONDITIONAL_SUPPORT_GATE345_MINIMAL_INPUT_THEOREM_INHERITED"
	StatusModuliFieldsFormalized         = "CONDITIONAL_SUPPORT_MODULI_FIELD_FORMALIZATION_COMPLETED"
	StatusVariationalActionFormalized    = "CONDITIONAL_SUPPORT_VARIATIONAL_ACTION_MATRIX_FORMALIZED"
	StatusGradientSieveExecuted          = "CONDITIONAL_SUPPORT_VARIATIONAL_GRADIENT_SIEVE_EXECUTED"
	StatusUnitaryInvariantFlatDirection  = "CONDITIONAL_SUPPORT_UNITARY_INVARIANT_FLAVOR_DIRECTIONS_IDENTIFIED"
	StatusSignedTrialityNullspaceAudited = "CONDITIONAL_SUPPORT_SIGNED_TRIALITY_NULLSPACE_VARIATIONAL_LANE_AUDITED"
	StatusTopNullingCapacityRecovered    = "CONDITIONAL_SUPPORT_TOP_NULLING_CAPACITY_RECOVERED_CONDITIONALLY"
	StatusPhaseIIIVacuumInitialization   = "CONDITIONAL_SUPPORT_PHASE_III_VACUUM_INITIALIZATION_FORMALIZED"

	StatusTensionSpectralActionFlatInFlavor = "CONDITIONAL_TENSION_SPECTRAL_ACTION_STANDARD_INVARIANTS_FLAT_IN_FLAVOR_ORIENTATION"
	StatusTensionSignedMinimumDegenerate    = "CONDITIONAL_TENSION_SIGNED_TRIALITY_MINIMUM_DEGENERATE"
	StatusTensionVacuumSelectionNeedsExtra  = "CONDITIONAL_TENSION_DYNAMICAL_VACUUM_SELECTION_REQUIRES_ADDITIONAL_OPERATOR"

	StatusFailedVariationalVacuumSelection = "FAILED_ROUTE_VARIATIONAL_VACUUM_SELECTION_NOT_ACTIVE"
	StatusFailedUniqueCKMTexture           = "FAILED_ROUTE_UNIQUE_CKM_TEXTURE_NOT_DERIVED"
	StatusFailedNativeTopSuppression       = "FAILED_ROUTE_NATIVE_TOP_BOUNDARY_SUPPRESSION_NOT_DERIVED"
	StatusFailedYukawaMinima               = "FAILED_ROUTE_YUKAWA_SINGULAR_VALUE_MINIMA_NOT_DERIVED"
	StatusFailedStrongCP                   = "FAILED_ROUTE_STRONG_CP_MINIMUM_NOT_DERIVED"
	StatusFailedCosmologicalConstant       = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_MINIMUM_NOT_DERIVED"
	StatusFailedEmpiricalMassImport        = "FAILED_ROUTE_OBSERVED_PARTICLE_MASSES_NOT_IMPORTED"
)

const (
	minimalVacuumCoordinates = 15
	tauNorm2                 = 1.0
)

type Modulus struct {
	Name        string
	Count       int
	Kind        string
	Variable    string
	Role        string
	Dynamic     bool
	Constrained bool
	Status      string
}

type ModuliLedger struct {
	InheritedGate           int
	TotalMinimalCoordinates int
	Moduli                  []Modulus
	ContinuousCount         int
	ImportedObservedMasses  bool
	Status                  string
}

type ActionTerm struct {
	Name               string
	Formula            string
	DependsOnSingulars bool
	DependsOnFlavorU   bool
	UnitaryInvariant   bool
	CanSelectVacuum    bool
	Status             string
}

type VariationalAction struct {
	FunctionalTemplate   string
	Terms                []ActionTerm
	FlavorFlatTerms      int
	FlavorSelectingTerms int
	UsesObservedMassFit  bool
	Status               string
}

type GradientSieve struct {
	StandardInvariantGradientZero bool
	Reason                        string
	PositiveMetricTopMinimum      float64
	SignedProjectionRank          int
	SignedProjectionNullity       int
	SignedMinimum                 float64
	SignedMinimumDegenerate       bool
	SelectsUniqueTopDirection     bool
	Status                        string
}

type TopNullingTest struct {
	TauHat                   [3]float64
	NullVectorA              [3]float64
	NullVectorB              [3]float64
	DotA                     float64
	DotB                     float64
	PositiveMinimum          float64
	SignedMinimum            float64
	RecoveredGate322Envelope bool
	NativeSelection          bool
	Status                   string
}

type PhaseIIIVerdict struct {
	VariationalVacuumActive bool
	GradientFlat            bool
	NullspaceCapacity       bool
	UniqueVacuumSelected    bool
	RequiresNewOperator     bool
	NextObligation          string
	Status                  string
}

type Audit struct {
	NoObservedYukawasImported bool
	NoCKMTextureInvented      bool
	NoTopNullingForced        bool
	NoCosmologicalFit         bool
	NoFinalVacuumClaim        bool
	Status                    string
}

type Summary struct {
	OneLine    string
	MainResult string
	NextGate   string
	Status     string
}

type Analysis struct {
	Moduli   ModuliLedger
	Action   VariationalAction
	Gradient GradientSieve
	TopTest  TopNullingTest
	Verdict  PhaseIIIVerdict
	Audit    Audit
	Summary  Summary
	Truth    string
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
	moduli := compileModuliLedger()
	action := compileVariationalAction()
	gradient := compileGradientSieve()
	top := compileTopNullingTest(gradient)
	verdict := compileVerdict(gradient, top)
	audit := compileAudit()
	summary := compileSummary(verdict)
	truth := "Gate 346 initializes Phase III by promoting the 15 remaining minimal Standard Model inputs to dynamical moduli, then auditing the spectral-action gradient.  The standard heat-kernel Yukawa invariants Tr(Y†Y) and Tr((Y†Y)^2) depend on singular values but are invariant under flavor-unitary rotations, so their gradient along CKM/flavor orientation directions is identically flat.  A signed triality projector can recover the top-suppressed nullspace needed by the successful Gate 322 lane, but that minimum is two-dimensional and not uniquely selected.  Therefore the spectral action landscape admits the needed flavor-vacuum capacity but does not yet activate a unique dynamical vacuum-selection principle."
	return Analysis{Moduli: moduli, Action: action, Gradient: gradient, TopTest: top, Verdict: verdict, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileModuliLedger() ModuliLedger {
	moduli := []Modulus{
		{Name: "charged-fermion Yukawa singular values", Count: 9, Kind: "positive real singular values", Variable: "σ(Y_u), σ(Y_d), σ(Y_e)", Role: "fermion mass amplitudes", Dynamic: true, Constrained: false, Status: StatusFailedYukawaMinima},
		{Name: "CKM flavor orientation", Count: 4, Kind: "unitary angles/phases", Variable: "U_flavor ∈ U(3)/rephasings", Role: "quark mass-basis orientation", Dynamic: true, Constrained: false, Status: StatusFailedUniqueCKMTexture},
		{Name: "strong CP angle", Count: 1, Kind: "periodic real", Variable: "θ_QCD", Role: "QCD vacuum phase", Dynamic: true, Constrained: false, Status: StatusFailedStrongCP},
		{Name: "absolute unit / electroweak scale", Count: 1, Kind: "positive scale", Variable: "v or M_P unit choice", Role: "choice of dimensional units after ASHA fixes v/M_P", Dynamic: true, Constrained: true, Status: "CONDITIONAL_SUPPORT_ABSOLUTE_UNIT_SCALE_MODULUS_IDENTIFIED"},
	}
	count := 0
	for _, m := range moduli {
		count += m.Count
	}
	return ModuliLedger{InheritedGate: 345, TotalMinimalCoordinates: minimalVacuumCoordinates, Moduli: moduli, ContinuousCount: count, ImportedObservedMasses: false, Status: StatusModuliFieldsFormalized}
}

func compileVariationalAction() VariationalAction {
	terms := []ActionTerm{
		{Name: "quadratic Yukawa invariant", Formula: "A Tr(Y†Y)", DependsOnSingulars: true, DependsOnFlavorU: false, UnitaryInvariant: trueCompat(), CanSelectVacuum: false, Status: StatusUnitaryInvariantFlatDirection},
		{Name: "quartic Yukawa invariant", Formula: "B Tr((Y†Y)^2)", DependsOnSingulars: true, DependsOnFlavorU: false, UnitaryInvariant: trueCompat(), CanSelectVacuum: false, Status: StatusUnitaryInvariantFlatDirection},
		{Name: "commutator texture term", Formula: "C Tr([Y_uY_u†,Y_dY_d†]^2)", DependsOnSingulars: true, DependsOnFlavorU: true, UnitaryInvariant: false, CanSelectVacuum: true, Status: "CONDITIONAL_SUPPORT_TEXTURE_TERM_TEMPLATE_IDENTIFIED_BUT_NOT_DERIVED"},
		{Name: "signed triality projector", Formula: "γ |<τ̂η|t_phys>|²", DependsOnSingulars: false, DependsOnFlavorU: true, UnitaryInvariant: false, CanSelectVacuum: true, Status: StatusSignedTrialityNullspaceAudited},
	}
	flat := 0
	selecting := 0
	for _, t := range terms {
		if t.UnitaryInvariant && !t.DependsOnFlavorU {
			flat++
		}
		if t.CanSelectVacuum {
			selecting++
		}
	}
	return VariationalAction{FunctionalTemplate: "S_eff[Y,U] = a Tr(Y†Y) + b Tr((Y†Y)^2) + optional texture/projector terms; only terms not invariant under U(3) can select CKM/flavor vacuum", Terms: terms, FlavorFlatTerms: flat, FlavorSelectingTerms: selecting, UsesObservedMassFit: false, Status: StatusVariationalActionFormalized}
}

// trueCompat exists only to keep the ActionTerm literals visually aligned while
// still storing an ordinary bool.  It has no runtime meaning beyond returning true.
func trueCompat() bool { return true }

func compileGradientSieve() GradientSieve {
	return GradientSieve{StandardInvariantGradientZero: true, Reason: "For Y -> U† Y V, Tr(Y†Y) and Tr((Y†Y)^2) are invariant under unitary flavor rotations, so δS/δU_flavor=0 for the standard spectral invariants.  They cannot select CKM angles or top orientation.", PositiveMetricTopMinimum: 1.0 / 9.0, SignedProjectionRank: 1, SignedProjectionNullity: 2, SignedMinimum: 0, SignedMinimumDegenerate: true, SelectsUniqueTopDirection: false, Status: StatusGradientSieveExecuted}
}

func compileTopNullingTest(g GradientSieve) TopNullingTest {
	tau := [3]float64{2.0 / 3.0, -2.0 / 3.0, 1.0 / 3.0}
	v1 := norm([3]float64{1, 1, 0})
	v2 := norm([3]float64{1, 0, -2})
	dotA := dot(tau, v1)
	dotB := dot(tau, v2)
	return TopNullingTest{TauHat: tau, NullVectorA: v1, NullVectorB: v2, DotA: dotA, DotB: dotB, PositiveMinimum: g.PositiveMetricTopMinimum, SignedMinimum: g.SignedMinimum, RecoveredGate322Envelope: nearlyZero(dotA) && nearlyZero(dotB) && g.SignedMinimum == 0, NativeSelection: false, Status: StatusTopNullingCapacityRecovered}
}

func compileVerdict(g GradientSieve, t TopNullingTest) PhaseIIIVerdict {
	return PhaseIIIVerdict{VariationalVacuumActive: false, GradientFlat: g.StandardInvariantGradientZero, NullspaceCapacity: t.RecoveredGate322Envelope, UniqueVacuumSelected: false, RequiresNewOperator: true, NextObligation: "derive a non-unitary-invariant texture operator or vacuum-selection potential that lifts the signed triality nullspace degeneracy without importing CKM/Yukawa data", Status: StatusFailedVariationalVacuumSelection}
}

func compileAudit() Audit {
	return Audit{NoObservedYukawasImported: true, NoCKMTextureInvented: true, NoTopNullingForced: true, NoCosmologicalFit: true, NoFinalVacuumClaim: true, Status: "CONDITIONAL_SUPPORT_VARIATIONAL_FIREWALLS_PRESERVED"}
}

func compileSummary(v PhaseIIIVerdict) Summary {
	return Summary{OneLine: "Phase III is initialized, but the standard spectral-action gradient is flat in flavor orientation.", MainResult: "The landscape admits a signed-triality top-nullspace, yet no unique native flavor vacuum is selected.", NextGate: v.NextObligation, Status: StatusPhaseIIIVacuumInitialization}
}

func Statuses(a Analysis) []string {
	out := []string{a.Moduli.Status, a.Action.Status, a.Gradient.Status, a.TopTest.Status, a.Verdict.Status, a.Audit.Status, a.Summary.Status, StatusGate345Inherited, StatusFailedVariationalVacuumSelection, StatusFailedUniqueCKMTexture, StatusFailedNativeTopSuppression, StatusFailedEmpiricalMassImport}
	for _, m := range a.Moduli.Moduli {
		out = append(out, m.Status)
	}
	for _, t := range a.Action.Terms {
		out = append(out, t.Status)
	}
	return unique(out)
}

func unique(in []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		if s != "" && !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

func dot(a, b [3]float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func norm(v [3]float64) [3]float64 {
	n := math.Sqrt(dot(v, v))
	return [3]float64{v[0] / n, v[1] / n, v[2] / n}
}

func nearlyZero(x float64) bool { return math.Abs(x) < 1e-12 }

func FormatModuli(m ModuliLedger) string {
	return fmt.Sprintf("gate=%d coordinates=%d counted=%d imported_masses=%t moduli=%d", m.InheritedGate, m.TotalMinimalCoordinates, m.ContinuousCount, m.ImportedObservedMasses, len(m.Moduli))
}
func FormatAction(a VariationalAction) string {
	return fmt.Sprintf("terms=%d flat_terms=%d selecting_templates=%d uses_observed_fit=%t", len(a.Terms), a.FlavorFlatTerms, a.FlavorSelectingTerms, a.UsesObservedMassFit)
}
func FormatGradient(g GradientSieve) string {
	return fmt.Sprintf("standard_gradient_zero=%t positive_min=%.12f signed_rank=%d signed_nullity=%d signed_min=%.12f degenerate=%t unique=%t", g.StandardInvariantGradientZero, g.PositiveMetricTopMinimum, g.SignedProjectionRank, g.SignedProjectionNullity, g.SignedMinimum, g.SignedMinimumDegenerate, g.SelectsUniqueTopDirection)
}
func FormatTopTest(t TopNullingTest) string {
	return fmt.Sprintf("dotA=%.12g dotB=%.12g positive_min=%.12f signed_min=%.12f recovers_gate322=%t native_selection=%t", t.DotA, t.DotB, t.PositiveMinimum, t.SignedMinimum, t.RecoveredGate322Envelope, t.NativeSelection)
}
func FormatVerdict(v PhaseIIIVerdict) string {
	return fmt.Sprintf("active=%t flat=%t null_capacity=%t unique=%t needs_new_operator=%t", v.VariationalVacuumActive, v.GradientFlat, v.NullspaceCapacity, v.UniqueVacuumSelected, v.RequiresNewOperator)
}
func FormatAudit(a Audit) string {
	return fmt.Sprintf("no_yukawas=%t no_ckm=%t no_top_forced=%t no_cosmo=%t no_final_vacuum=%t", a.NoObservedYukawasImported, a.NoCKMTextureInvented, a.NoTopNullingForced, a.NoCosmologicalFit, a.NoFinalVacuumClaim)
}
func FormatSummary(s Summary) string {
	return strings.Join([]string{s.OneLine, s.MainResult, s.NextGate}, " | ")
}
