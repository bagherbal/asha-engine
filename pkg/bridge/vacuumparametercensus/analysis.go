// Package vacuumparametercensus implements Gate 345:
// Vacuum Parameter Census / Minimal Input Theorem.
//
// Gate 345 converts the accumulated FAILED_ROUTE ledger into a formal
// landscape-vs-vacuum theorem.  It distinguishes quantities derived by the
// finite Cℓ(1,7) spectral architecture from the remaining coordinates needed
// to select our physical vacuum.
package vacuumparametercensus

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE345-VACUUM-PARAMETER-CENSUS-MINIMAL-INPUT-THEOREM"

	StatusGate344Inherited               = "CONDITIONAL_SUPPORT_GATE344_COMPLETE_MOMENT_LEDGER_INHERITED"
	StatusFailureLedgerClustered         = "CONDITIONAL_SUPPORT_FAILURE_LEDGER_CLUSTERED_BY_TYPE"
	StatusLandscapeVacuumSplitFormalized = "CONDITIONAL_SUPPORT_LANDSCAPE_VS_VACUUM_SPLIT_FORMALIZED"
	StatusMinimalSM19CensusFormalized    = "CONDITIONAL_SUPPORT_MINIMAL_SM19_PARAMETER_CENSUS_FORMALIZED"
	StatusFourNativeBoundaryConstraints  = "CONDITIONAL_SUPPORT_FOUR_NATIVE_BOUNDARY_CONSTRAINTS_CATALOGED"
	StatusMinimalVacuumInputsCounted     = "CONDITIONAL_SUPPORT_MINIMAL_VACUUM_INPUTS_COUNTED"
	StatusExtendedVacuumLedgerCataloged  = "CONDITIONAL_SUPPORT_EXTENDED_NEUTRINO_COSMOLOGY_LEDGER_CATALOGED"
	StatusMinimalInputTheoremFormalized  = "CONDITIONAL_SUPPORT_MINIMAL_INPUT_THEOREM_FORMALIZED"
	StatusPhaseIIICoordinatesIdentified  = "CONDITIONAL_SUPPORT_PHASE_III_VACUUM_COORDINATES_IDENTIFIED"

	StatusTensionVacuumNotDerived            = "CONDITIONAL_TENSION_VACUUM_SELECTION_NOT_DERIVED_FROM_FINITE_CORE"
	StatusTensionExtendedCountDependsOnModel = "CONDITIONAL_TENSION_EXTENDED_VACUUM_DIMENSION_DEPENDS_ON_NEUTRINO_COSMOLOGY_MODEL"

	StatusFailedYukawaAmplitudesRemainVacuum = "FAILED_ROUTE_YUKAWA_AMPLITUDES_REMAIN_VACUUM_COORDINATES"
	StatusFailedCKMTextureRemainsVacuum      = "FAILED_ROUTE_CKM_TEXTURE_REMAINS_VACUUM_COORDINATE"
	StatusFailedPMNSTextureRemainsVacuum     = "FAILED_ROUTE_PMNS_TEXTURE_REMAINS_EXTENDED_VACUUM_COORDINATE"
	StatusFailedCosmologicalConstantVacuum   = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_REMAINS_VACUUM_ENERGY_COORDINATE"
	StatusFailedAbsoluteUnitScaleStillInput  = "FAILED_ROUTE_ABSOLUTE_UNIT_SCALE_STILL_INPUT"
	StatusFailedFinalVacuumNotDerived        = "FAILED_ROUTE_PHYSICAL_VACUUM_POINT_NOT_DERIVED"
	StatusFailedTheoryOfEverythingNotClaimed = "FAILED_ROUTE_FINAL_THEORY_OF_EVERYTHING_NOT_CLAIMED"
)

const highestInheritedGate = 344

type FailureCluster struct {
	Name            string
	FailureType     string
	Examples        []string
	DimensionImpact string
	Status          string
}

type FailureLedger struct {
	HighestGateInherited int
	Clusters             []FailureCluster
	TypeACount           int
	TypeBCount           int
	LandscapeNotVacuum   bool
	Status               string
}

type LandscapeResult struct {
	Name             string
	Formula          string
	Category         string
	Native           bool
	Gate             int
	ReducesParameter bool
	Status           string
}

type LandscapeLedger struct {
	Results                       []LandscapeResult
	NativeBoundaryConstraintCount int
	ContainsWeakMixing            bool
	ContainsHiggsGaugeRatio       bool
	ContainsAlphaGUT              bool
	ContainsHierarchy             bool
	ContainsGaugeGroup            bool
	ContainsMatterContent         bool
	ContainsGenerations           bool
	ContainsMoritaSplit           bool
	Status                        string
}

type Parameter struct {
	Name          string
	Count         int
	Kind          string
	Role          string
	VacuumInput   bool
	DerivedByASHA bool
	Comment       string
	Status        string
}

type MinimalSM19Census struct {
	BaselineCount          int
	Parameters             []Parameter
	DerivedOrConstrained   []Parameter
	RemainingVacuumInputs  []Parameter
	RemainingContinuousDim int
	RemainingDiscreteDim   int
	MinimalInputCount      int
	CountEquation          string
	Status                 string
}

type ExtendedVacuumLedger struct {
	IncludeNeutrinos   bool
	IncludeCosmology   bool
	Additions          []Parameter
	AddedContinuousDim int
	TotalExtendedDim   int
	ModelDependent     bool
	Status             string
}

type MinimalInputTheorem struct {
	Statement           string
	LandscapeDimension  string
	MinimalSMVacuumDim  int
	ExtendedVacuumDim   int
	DiscreteSeals       []string
	ContinuousInputs    []string
	ProvesLandscapeOnly bool
	DerivesVacuumPoint  bool
	Status              string
}

type Audit struct {
	NoYukawaFitInserted       bool
	NoCKMInvented             bool
	NoPMNSInvented            bool
	NoCosmologicalConstantFit bool
	NoVacuumDirectionForced   bool
	NoPrecisionClaimInserted  bool
	FinalTOEClaimed           bool
	Status                    string
}

type Summary struct {
	OneLine        string
	MinimalResult  string
	ExtendedResult string
	NextGate       string
	Status         string
}

type Analysis struct {
	Failures  FailureLedger
	Landscape LandscapeLedger
	MinimalSM MinimalSM19Census
	Extended  ExtendedVacuumLedger
	Theorem   MinimalInputTheorem
	Audit     Audit
	Summary   Summary
	Truth     string
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
	failures := compileFailureLedger()
	landscape := compileLandscapeLedger()
	minimal := compileMinimalSM19Census()
	extended := compileExtendedLedger(minimal)
	theorem := compileMinimalInputTheorem(minimal, extended)
	audit := compileAudit()
	summary := compileSummary(minimal, extended)
	truth := "Gate 345 proves the project-level pattern behind the accumulated FAILED_ROUTE ledger: the finite Cℓ(1,7) spectral architecture derives the Standard Model landscape—gauge group, representations, generation topology, Morita split, exact boundary ratios, coupling normalization branch, and hierarchy ratio—but it does not select the unique physical vacuum point.  In the minimal 19-parameter Standard Model convention, four native boundary constraints replace four continuous inputs, leaving 15 vacuum-selection coordinates: nine charged-fermion Yukawa singular values, four CKM parameters, one strong-CP angle, and one absolute unit/VEV scale.  In the extended neutrino/cosmology ledger, PMNS/neutrino data and the cosmological constant add model-dependent vacuum coordinates.  This is a minimal-input theorem, not a final-vacuum derivation."
	return Analysis{Failures: failures, Landscape: landscape, MinimalSM: minimal, Extended: extended, Theorem: theorem, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileFailureLedger() FailureLedger {
	clusters := []FailureCluster{
		{Name: "Vacuum direction", FailureType: "Type A: which direction?", Examples: []string{"CKM flavor vacuum", "PMNS flavor vacuum", "signed-vs-positive flavor projection metric", "electroweak vacuum orientation", "resolvent 2+2 pairing selector"}, DimensionImpact: "selects a point/orientation inside a derived vacuum landscape", Status: StatusTensionVacuumNotDerived},
		{Name: "Physical real structure semantics", FailureType: "Type A: which direction?", Examples: []string{"physical anti-linear J beyond formal J_swap", "opposite algebra physical semantics", "B-gap instanton saddle semantics"}, DimensionImpact: "selects physical interpretation of a valid formal carrier", Status: StatusTensionVacuumNotDerived},
		{Name: "Vacuum energy", FailureType: "Type A: which direction?", Examples: []string{"cosmological constant f4Λ4", "a0_eff vacuum multiplicity", "vacuum subtraction/renormalization"}, DimensionImpact: "selects vacuum-energy coordinate", Status: StatusFailedCosmologicalConstantVacuum},
		{Name: "Amplitude origin", FailureType: "Type A: which direction?", Examples: []string{"EmpiricalYukawaSeal", "individual fermion masses", "top-sector texture", "non-commuting texture pair"}, DimensionImpact: "selects singular values and flavor texture", Status: StatusFailedYukawaAmplitudesRemainVacuum},
		{Name: "Precision", FailureType: "Type B: how precisely?", Examples: []string{"two-loop RG", "pole mass conversion", "threshold matching values", "PV coefficient/counterterm scheme"}, DimensionImpact: "does not choose the vacuum; improves conversion from boundary data to observables", Status: "CONDITIONAL_SUPPORT_PRECISION_CLUSTER_IDENTIFIED"},
	}
	return FailureLedger{HighestGateInherited: highestInheritedGate, Clusters: clusters, TypeACount: 4, TypeBCount: 1, LandscapeNotVacuum: true, Status: StatusFailureLedgerClustered}
}

func compileLandscapeLedger() LandscapeLedger {
	results := []LandscapeResult{
		{Name: "Gauge group and inner-fluctuation field content", Formula: "U(1)_Y × SU(2)_L × SU(3)_C with 12 gauge bosons and one Higgs doublet", Category: "framework", Native: true, Gate: 298, ReducesParameter: false, Status: "CONDITIONAL_SUPPORT_SM_FIELD_CONTENT_DERIVED"},
		{Name: "Matter representation / Morita split", Formula: "κ_C:κ_Q = 1:3", Category: "framework", Native: true, Gate: 295, ReducesParameter: false, Status: "CONDITIONAL_SUPPORT_MORITA_COLOR_SPLIT_DERIVED"},
		{Name: "Generation topology", Formula: "N_gen=3; τ_η=(2,-2,1)", Category: "framework", Native: true, Gate: 26, ReducesParameter: false, Status: "CONDITIONAL_SUPPORT_GENERATION_TOPOLOGY_DERIVED"},
		{Name: "Weak mixing boundary ratio", Formula: "sin²θ_W=3/8", Category: "boundary ratio", Native: true, Gate: 298, ReducesParameter: true, Status: "CONDITIONAL_SUPPORT_WEAK_MIXING_RATIO_DERIVED"},
		{Name: "Higgs-to-gauge quartic boundary ratio", Formula: "λ_H/g_*²=1197/4624", Category: "boundary ratio", Native: true, Gate: 307, ReducesParameter: true, Status: "CONDITIONAL_SUPPORT_HIGGS_GAUGE_RATIO_DERIVED"},
		{Name: "Unified coupling branch", Formula: "α_GUT^{-1}=8π under full doubled bosonic trace", Category: "boundary normalization", Native: true, Gate: 330, ReducesParameter: true, Status: "CONDITIONAL_SUPPORT_ALPHA_GUT_BRANCH_DERIVED"},
		{Name: "Electroweak-to-Planck hierarchy ratio", Formula: "v/M_P=2^{3/2}exp(-4π²)", Category: "hierarchy ratio", Native: true, Gate: 342, ReducesParameter: true, Status: "CONDITIONAL_SUPPORT_HIERARCHY_RATIO_DERIVED"},
		{Name: "Threshold portal jump witness", Formula: "Δλ≈-0.097846792207", Category: "threshold structure", Native: true, Gate: 321, ReducesParameter: false, Status: "CONDITIONAL_SUPPORT_THRESHOLD_JUMP_WITNESS_DERIVED"},
	}
	return LandscapeLedger{Results: results, NativeBoundaryConstraintCount: 4, ContainsWeakMixing: true, ContainsHiggsGaugeRatio: true, ContainsAlphaGUT: true, ContainsHierarchy: true, ContainsGaugeGroup: true, ContainsMatterContent: true, ContainsGenerations: true, ContainsMoritaSplit: true, Status: StatusLandscapeVacuumSplitFormalized}
}

func compileMinimalSM19Census() MinimalSM19Census {
	params := []Parameter{
		{Name: "gauge-sector boundary data", Count: 3, Kind: "continuous", Role: "g1,g2,g3 or equivalent boundary coupling data", VacuumInput: false, DerivedByASHA: true, Comment: "weak mixing and unified coupling branch reduce the independent boundary gauge data", Status: "CONDITIONAL_SUPPORT_GAUGE_BOUNDARY_CONSTRAINED"},
		{Name: "Higgs quartic / scalar shape", Count: 1, Kind: "continuous", Role: "λ_H", VacuumInput: false, DerivedByASHA: true, Comment: "fixed relative to gauge coupling by 1197/4624", Status: "CONDITIONAL_SUPPORT_HIGGS_QUARTIC_CONSTRAINED"},
		{Name: "hierarchy / absolute scale ratio", Count: 1, Kind: "continuous dimensional", Role: "v/M_P", VacuumInput: false, DerivedByASHA: true, Comment: "ratio fixed by Pfaffian hierarchy; one absolute unit scale remains", Status: "CONDITIONAL_SUPPORT_HIERARCHY_RATIO_CONSTRAINED"},
		{Name: "charged-fermion Yukawa singular values", Count: 9, Kind: "continuous positive", Role: "six quark masses plus three charged-lepton masses", VacuumInput: true, DerivedByASHA: false, Comment: "finite algebra supplies allowed Yukawa edges but not numerical singular values", Status: StatusFailedYukawaAmplitudesRemainVacuum},
		{Name: "CKM texture", Count: 4, Kind: "continuous angles/phases", Role: "three quark mixing angles plus one CP phase", VacuumInput: true, DerivedByASHA: false, Comment: "flavor landscape exists; unique vacuum orientation is not selected", Status: StatusFailedCKMTextureRemainsVacuum},
		{Name: "strong CP angle", Count: 1, Kind: "periodic continuous", Role: "θ_QCD", VacuumInput: true, DerivedByASHA: false, Comment: "not fixed by current finite spectral landscape", Status: "FAILED_ROUTE_STRONG_CP_THETA_REMAINS_VACUUM_COORDINATE"},
		{Name: "absolute unit / electroweak scale", Count: 1, Kind: "dimensional scale", Role: "choice of v or equivalently M_P units", VacuumInput: true, DerivedByASHA: false, Comment: "ASHA fixes v/M_P, not the arbitrary unit used to express GeV", Status: StatusFailedAbsoluteUnitScaleStillInput},
	}
	var remaining []Parameter
	var derived []Parameter
	for _, p := range params {
		if p.VacuumInput {
			remaining = append(remaining, p)
		}
		if p.DerivedByASHA {
			derived = append(derived, p)
		}
	}
	return MinimalSM19Census{BaselineCount: 19, Parameters: params, DerivedOrConstrained: derived, RemainingVacuumInputs: remaining, RemainingContinuousDim: 15, RemainingDiscreteDim: 0, MinimalInputCount: 15, CountEquation: "19 baseline SM parameters - 4 native boundary constraints = 15 minimal vacuum-selection inputs", Status: StatusMinimalSM19CensusFormalized}
}

func compileExtendedLedger(min MinimalSM19Census) ExtendedVacuumLedger {
	additions := []Parameter{
		{Name: "neutrino masses", Count: 3, Kind: "continuous positive", Role: "three light-neutrino mass eigenvalues", VacuumInput: true, DerivedByASHA: false, Comment: "right-handed/Majorana support exists but measured spectrum is not selected", Status: StatusFailedPMNSTextureRemainsVacuum},
		{Name: "PMNS texture", Count: 6, Kind: "continuous angles/phases", Role: "three lepton mixing angles, one Dirac phase, two Majorana phases", VacuumInput: true, DerivedByASHA: false, Comment: "PMNS vacuum orientation remains sealed", Status: StatusFailedPMNSTextureRemainsVacuum},
		{Name: "cosmological constant", Count: 1, Kind: "vacuum-energy coordinate", Role: "ρΛ or f4Λ4a0_eff after vacuum subtraction", VacuumInput: true, DerivedByASHA: false, Comment: "Gate 344 did not derive the cosmological a0/f4 channel", Status: StatusFailedCosmologicalConstantVacuum},
	}
	return ExtendedVacuumLedger{IncludeNeutrinos: true, IncludeCosmology: true, Additions: additions, AddedContinuousDim: 10, TotalExtendedDim: min.RemainingContinuousDim + 10, ModelDependent: true, Status: StatusExtendedVacuumLedgerCataloged}
}

func compileMinimalInputTheorem(min MinimalSM19Census, ext ExtendedVacuumLedger) MinimalInputTheorem {
	continuous := []string{"9 charged-fermion Yukawa singular values", "4 CKM parameters", "1 strong-CP angle", "1 absolute unit/VEV scale"}
	discrete := []string{"resolvent 2+2 pairing/orientation branch", "electroweak weak-plane sign/orientation", "signed-vs-positive flavor projection metric", "spontaneous scalar carrier orientation"}
	return MinimalInputTheorem{Statement: "The finite Cℓ(1,7) spectral architecture determines the Standard Model landscape and four exact boundary constraints, but the physical vacuum point requires 15 minimal continuous inputs in the SM-19 convention; extended neutrino/cosmology models add 10 more continuous vacuum coordinates.", LandscapeDimension: "framework data are discrete/structural plus exact ratios, not vacuum coordinates", MinimalSMVacuumDim: min.MinimalInputCount, ExtendedVacuumDim: ext.TotalExtendedDim, DiscreteSeals: discrete, ContinuousInputs: continuous, ProvesLandscapeOnly: true, DerivesVacuumPoint: false, Status: StatusMinimalInputTheoremFormalized}
}

func compileAudit() Audit {
	return Audit{NoYukawaFitInserted: true, NoCKMInvented: true, NoPMNSInvented: true, NoCosmologicalConstantFit: true, NoVacuumDirectionForced: true, NoPrecisionClaimInserted: true, FinalTOEClaimed: false, Status: "CONDITIONAL_SUPPORT_VACUUM_CENSUS_FIREWALLS_PRESERVED"}
}

func compileSummary(min MinimalSM19Census, ext ExtendedVacuumLedger) Summary {
	return Summary{OneLine: "ASHA derives the landscape, not the unique vacuum point.", MinimalResult: min.CountEquation, ExtendedResult: fmt.Sprintf("minimal %d + neutrino/cosmology %d = %d extended continuous vacuum coordinates", min.MinimalInputCount, ext.AddedContinuousDim, ext.TotalExtendedDim), NextGate: "Phase III should target a dynamical vacuum-selection principle rather than another algebraic landscape invariant.", Status: StatusPhaseIIICoordinatesIdentified}
}

func Statuses(a Analysis) []string {
	out := []string{a.Failures.Status, a.Landscape.Status, a.MinimalSM.Status, StatusFourNativeBoundaryConstraints, a.Extended.Status, a.Theorem.Status, a.Audit.Status, a.Summary.Status, StatusFailedFinalVacuumNotDerived, StatusFailedTheoryOfEverythingNotClaimed}
	for _, c := range a.Failures.Clusters {
		out = append(out, c.Status)
	}
	for _, p := range a.MinimalSM.RemainingVacuumInputs {
		out = append(out, p.Status)
	}
	for _, p := range a.Extended.Additions {
		out = append(out, p.Status)
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

func FormatFailures(f FailureLedger) string {
	return fmt.Sprintf("gate≤%d clusters=%d typeA=%d typeB=%d landscape_not_vacuum=%t", f.HighestGateInherited, len(f.Clusters), f.TypeACount, f.TypeBCount, f.LandscapeNotVacuum)
}
func FormatLandscape(l LandscapeLedger) string {
	return fmt.Sprintf("results=%d native_boundary_constraints=%d weak=%t quartic=%t alpha=%t hierarchy=%t framework=%t/%t/%t", len(l.Results), l.NativeBoundaryConstraintCount, l.ContainsWeakMixing, l.ContainsHiggsGaugeRatio, l.ContainsAlphaGUT, l.ContainsHierarchy, l.ContainsGaugeGroup, l.ContainsMatterContent, l.ContainsGenerations)
}
func FormatMinimal(m MinimalSM19Census) string {
	return fmt.Sprintf("baseline=%d minimal_inputs=%d continuous_dim=%d equation=%q", m.BaselineCount, m.MinimalInputCount, m.RemainingContinuousDim, m.CountEquation)
}
func FormatExtended(e ExtendedVacuumLedger) string {
	return fmt.Sprintf("additions=%d added_dim=%d total_extended_dim=%d model_dependent=%t", len(e.Additions), e.AddedContinuousDim, e.TotalExtendedDim, e.ModelDependent)
}
func FormatTheorem(t MinimalInputTheorem) string {
	return fmt.Sprintf("minimal_dim=%d extended_dim=%d landscape_only=%t derives_vacuum=%t", t.MinimalSMVacuumDim, t.ExtendedVacuumDim, t.ProvesLandscapeOnly, t.DerivesVacuumPoint)
}
func FormatAudit(a Audit) string {
	return fmt.Sprintf("no_yukawa_fit=%t no_ckm=%t no_pmns=%t no_cosmo_fit=%t no_vacuum_forced=%t no_precision_claim=%t final_toe=%t", a.NoYukawaFitInserted, a.NoCKMInvented, a.NoPMNSInvented, a.NoCosmologicalConstantFit, a.NoVacuumDirectionForced, a.NoPrecisionClaimInserted, a.FinalTOEClaimed)
}
func FormatSummary(s Summary) string {
	return strings.Join([]string{s.OneLine, s.MinimalResult, s.ExtendedResult, s.NextGate}, " | ")
}
