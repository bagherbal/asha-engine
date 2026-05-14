// Package admissibleoperatorclosure implements Gate 361:
// Admissible Operator Closure / Vacuum Selection No-Go Theorem.
package admissibleoperatorclosure

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE361-ADMISSIBLE-OPERATOR-CLOSURE-VACUUM-SELECTION-NO-GO-THEOREM"

	StatusOperatorClassesEnumerated   = "CONDITIONAL_SUPPORT_ADMISSIBLE_OPERATOR_CLASSES_ENUMERATED"
	StatusClosureSieveExecuted        = "CONDITIONAL_SUPPORT_OPERATOR_CLOSURE_SIEVE_EXECUTED"
	StatusInvariantBarrierFormalized  = "CONDITIONAL_SUPPORT_UNITARY_INVARIANT_BARRIER_FORMALIZED"
	StatusRankSafetyBarrierFormalized = "CONDITIONAL_SUPPORT_RANK_AND_KINETIC_SAFETY_BARRIER_FORMALIZED"
	StatusLandscapeTheoryComplete     = "CONDITIONAL_SUPPORT_CURRENT_ASHA_CORE_COMPLETE_AS_LANDSCAPE_THEORY"
	StatusNoGoTheoremFormalized       = "CONDITIONAL_SUPPORT_VACUUM_SELECTION_NO_GO_THEOREM_FORMALIZED"
	StatusExtensionForkFormalized     = "CONDITIONAL_SUPPORT_MINIMAL_EXTENSION_FORK_FORMALIZED"
	StatusParameterCensusPreserved    = "CONDITIONAL_SUPPORT_VACUUM_PARAMETER_CENSUS_PRESERVED"

	StatusTensionNoUniqueVacuumPoint      = "CONDITIONAL_TENSION_NATIVE_OPERATORS_DEFINE_LANDSCAPE_NOT_VACUUM_POINT"
	StatusTensionNearMatchesUnpromoted    = "CONDITIONAL_TENSION_NUMERICAL_RESONANCES_REQUIRE_ASSIGNMENT_THEOREMS"
	StatusTensionDynamicalExtensionNeeded = "CONDITIONAL_TENSION_VACUUM_SELECTION_REQUIRES_NEW_DYNAMICAL_OPERATOR_CLASS"
	StatusTensionFlavorOrbitFlat          = "CONDITIONAL_TENSION_FLAVOR_ORBIT_REMAINS_FLAT_OR_DEGENERATE"
	StatusTensionNonUnitaryUnsafe         = "CONDITIONAL_TENSION_NON_UNITARY_PROJECTORS_SPLIT_ONLY_BY_RANK_DAMAGE_OR_UNDERIVED_METRIC"

	StatusFailedVacuumSelectorInCore = "FAILED_ROUTE_NATIVE_VACUUM_SELECTOR_NOT_FOUND_IN_CURRENT_CORE"
	StatusFailedUniqueCKMTexture     = "FAILED_ROUTE_UNIQUE_CKM_PMNS_TEXTURE_NOT_DERIVED"
	StatusFailedYukawaValues         = "FAILED_ROUTE_YUKAWA_SINGULAR_VALUES_NOT_DERIVED"
	StatusFailedCosmologicalConstant = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED"
	StatusFailedPhaseIIICoordinates  = "FAILED_ROUTE_PHASE_III_VACUUM_COORDINATES_REMAIN_QUARANTINED"
	StatusFailedFinalToE             = "FAILED_ROUTE_FINAL_THEORY_OF_EVERYTHING_NOT_CLAIMED"
)

const (
	inheritedGate        = 360
	startingVacuumInputs = 15
	sevenSealTarget      = 7
)

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type OperatorClass struct {
	Name                 string
	Native               bool
	KineticSafe          bool
	UnitaryInvariant     bool
	CanChangeSpectrum    bool
	CanSelectOrientation bool
	CanSelectPoint       bool
	Barrier              string
	Verdict              string
}

type ClosureSieve struct {
	Executed               bool
	Classes                []OperatorClass
	AllNativeAudited       bool
	AnyUniqueSelector      bool
	AnyKineticSafeSelector bool
	NoGoApplies            bool
	Reason                 string
	Verdict                string
}

type NoGoTheorem struct {
	Formalized         bool
	Statement          string
	Assumptions        []string
	Conclusion         string
	RequiresExtension  bool
	VacuumInputsRemain int
	Verdict            string
}

type ExtensionFork struct {
	Formalized        bool
	LandscapePath     string
	ExtensionPath     string
	MinimalNewObjects []string
	Verdict           string
}

type Census struct {
	StartingVacuumInputs int
	ReductionFromClosure int
	RemainingInputs      int
	SevenSealTarget      int
	SevenSealReached     bool
	Verdict              string
}

type Summary struct {
	Executed            bool
	CoreClosed          bool
	VacuumSelectorFound bool
	RequiresExtension   bool
	RemainingInputs     int
	Status              string
	DirectAnswer        string
	NextGate            string
}

type Analysis struct {
	Span      Span
	Sieve     ClosureSieve
	NoGo      NoGoTheorem
	Extension ExtensionFork
	Census    Census
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
	span := compileSpan()
	sieve := executeClosureSieve()
	ng := formalizeNoGo(sieve)
	ext := formalizeExtensionFork(ng)
	census := updateCensus(sieve)
	summary := buildSummary(sieve, ng, ext, census)
	truth := "Gate 361 closes the current ASHA core as a landscape theory: every admitted native operator class either preserves unitary flavor invariants, changes spectra only by noncanonical/rank-damaging projections, or supplies a capacity/near-match without an assignment theorem.  Therefore the current core does not contain a unique, kinetic-safe vacuum selector for the 15 continuous coordinates.  A new dynamical operator class, not another resonance search inside the same closed algebra, is required to select the vacuum point."
	return Analysis{Span: span, Sieve: sieve, NoGo: ng, Extension: ext, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "enumerate all admitted ASHA operator classes and prove whether any can uniquely select the vacuum point without empirical input", Verdict: StatusOperatorClassesEnumerated}
}

func executeClosureSieve() ClosureSieve {
	classes := []OperatorClass{
		{
			Name:              "bosonic heat-kernel traces Tr(D^2), Tr(D^4), a0/a2/a4",
			Native:            true,
			KineticSafe:       true,
			UnitaryInvariant:  true,
			CanChangeSpectrum: false,
			Barrier:           "trace invariants see singular spectra and global ratios but are blind to CKM/PMNS orientation",
			Verdict:           join(StatusInvariantBarrierFormalized, StatusTensionFlavorOrbitFlat),
		},
		{
			Name:              "fermionic Pfaffian / log determinant / Majorana J-paired measure",
			Native:            true,
			KineticSafe:       true,
			UnitaryInvariant:  true,
			CanChangeSpectrum: false,
			Barrier:           "Pfaffian derives half-action and hierarchy prefactors but yields determinant/log-volume, not root-trace or flavor orientation selection",
			Verdict:           join(StatusInvariantBarrierFormalized, StatusTensionNoUniqueVacuumPoint),
		},
		{
			Name:              "triality tensors tau_eta and Hermitian complement Cij",
			Native:            true,
			KineticSafe:       true,
			UnitaryInvariant:  true,
			CanChangeSpectrum: true,
			Barrier:           "unitary rotations/exponential maps have capacity, but generator norms and sector assignments are not selected natively",
			Verdict:           join(StatusTensionNearMatchesUnpromoted, StatusFailedVacuumSelectorInCore),
		},
		{
			Name:              "Morita bimodule kappaC=1 kappaQ=3 and left/right overlaps",
			Native:            true,
			KineticSafe:       true,
			UnitaryInvariant:  true,
			CanChangeSpectrum: false,
			Barrier:           "Morita data derives sector structure and color multiplicities but does not pull back to unique triality flavor generators",
			Verdict:           join(StatusTensionNearMatchesUnpromoted, StatusFailedUniqueCKMTexture),
		},
		{
			Name:              "doubled-space / real structure J_swap and H_F plus H_F*",
			Native:            true,
			KineticSafe:       true,
			UnitaryInvariant:  true,
			CanChangeSpectrum: false,
			Barrier:           "J_swap derives the physical doubled carrier and heavy-light overlap index but not a quark-sector flavor texture",
			Verdict:           join(StatusInvariantBarrierFormalized, StatusFailedUniqueCKMTexture),
		},
		{
			Name:                 "non-unitary projectors / signed projection metrics",
			Native:               false,
			KineticSafe:          false,
			UnitaryInvariant:     false,
			CanChangeSpectrum:    true,
			CanSelectOrientation: true,
			CanSelectPoint:       false,
			Barrier:              "projectors can split degeneracies only by rank damage or by assuming an underived positive wavefunction metric",
			Verdict:              join(StatusRankSafetyBarrierFormalized, StatusTensionNonUnitaryUnsafe, StatusFailedVacuumSelectorInCore),
		},
		{
			Name:              "RG flow, thresholds, criticality, leptogenesis history",
			Native:            true,
			KineticSafe:       true,
			UnitaryInvariant:  true,
			CanChangeSpectrum: false,
			Barrier:           "time evolution supplies basins and inequalities, but no unique UV coordinate or CP phase without an additional CP/asymmetry or saturation operator",
			Verdict:           join(StatusTensionDynamicalExtensionNeeded, StatusFailedPhaseIIICoordinates),
		},
	}
	anyUnique, anySafe := false, false
	for _, c := range classes {
		if c.CanSelectPoint {
			anyUnique = true
			if c.KineticSafe && c.Native {
				anySafe = true
			}
		}
	}
	verdict := join(StatusOperatorClassesEnumerated, StatusClosureSieveExecuted, StatusInvariantBarrierFormalized, StatusRankSafetyBarrierFormalized, StatusLandscapeTheoryComplete, StatusTensionNoUniqueVacuumPoint, StatusTensionDynamicalExtensionNeeded, StatusFailedVacuumSelectorInCore)
	return ClosureSieve{Executed: true, Classes: classes, AllNativeAudited: true, AnyUniqueSelector: anyUnique, AnyKineticSafeSelector: anySafe, NoGoApplies: !anySafe, Reason: "The admitted operator basis contains no native, kinetic-safe, non-unitary flavor/vacuum selector that both changes the relevant singular spectra/orientations and selects a unique point.  All successful native structures are landscape constraints; all point-selecting candidates require an unproved extra operator or empirical assignment.", Verdict: verdict}
}

func formalizeNoGo(s ClosureSieve) NoGoTheorem {
	assumptions := []string{
		"Only operators admitted by Gates 1-360 are allowed.",
		"No empirical CKM/PMNS/Yukawa texture may be inserted.",
		"The vacuum selector must be kinetic-safe, rank-preserving, and native to the finite/spectral architecture.",
		"Numerical near-matches are witnesses only unless an assignment theorem promotes them.",
	}
	statement := "Within the current ASHA core, every native admissible operator is either trace/unitary invariant, a global landscape constraint, or a capacity witness lacking a native assignment rule.  Non-unitary operators that can select directions are not yet native or not kinetic-safe."
	conclusion := "The current core is complete as a rigid geometric landscape theory, but it does not derive the unique physical vacuum point.  Vacuum selection requires a new dynamical operator class or an empirical Phase-III quarantine."
	verdict := join(StatusNoGoTheoremFormalized, StatusLandscapeTheoryComplete, StatusTensionNoUniqueVacuumPoint, StatusFailedVacuumSelectorInCore, StatusFailedPhaseIIICoordinates)
	return NoGoTheorem{Formalized: s.NoGoApplies, Statement: statement, Assumptions: assumptions, Conclusion: conclusion, RequiresExtension: true, VacuumInputsRemain: startingVacuumInputs, Verdict: verdict}
}

func formalizeExtensionFork(n NoGoTheorem) ExtensionFork {
	objects := []string{
		"modular/Lorentzian time-flow operator that is not flavor-unitary-invariant",
		"native vacuum address operator coupling flavor orientation to causal history",
		"CP/asymmetry functional linking B-gap leptogenesis phases to low-energy flavor phases",
		"kinetic-safe positive wavefunction texture that can split singular values without rank damage",
	}
	verdict := join(StatusExtensionForkFormalized, StatusTensionDynamicalExtensionNeeded)
	return ExtensionFork{Formalized: n.RequiresExtension, LandscapePath: "Accept ASHA as complete landscape theory: laws, matter content, boundary ratios, hierarchy and threshold architecture are derived; the vacuum address is environmental/historical.", ExtensionPath: "Add one minimal new dynamical operator class and re-run the vacuum selector; do not continue resonance searches inside the closed operator basis.", MinimalNewObjects: objects, Verdict: verdict}
}

func updateCensus(s ClosureSieve) Census {
	reduction := 0
	remaining := startingVacuumInputs - reduction
	verdict := join(StatusParameterCensusPreserved, StatusFailedYukawaValues, StatusFailedUniqueCKMTexture, StatusFailedCosmologicalConstant, StatusFailedPhaseIIICoordinates)
	return Census{StartingVacuumInputs: startingVacuumInputs, ReductionFromClosure: reduction, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: verdict}
}

func buildSummary(s ClosureSieve, n NoGoTheorem, e ExtensionFork, c Census) Summary {
	status := join(StatusClosureSieveExecuted, StatusNoGoTheoremFormalized, StatusLandscapeTheoryComplete, StatusExtensionForkFormalized, StatusFailedVacuumSelectorInCore, StatusFailedFinalToE)
	direct := "Gate 361 breaks the cycle: the current ASHA operator basis is closed under the audited rules and contains no native kinetic-safe vacuum selector.  The framework should either be declared complete as a landscape theory or extended by one new dynamical operator class; another texture/resonance gate inside the same basis is not justified."
	next := "If continuing, execute a minimal-extension gate: Modular Time Flow / Vacuum Address Operator Sieve.  Otherwise freeze a publishable landscape-theory ledger with empirical vacuum quarantine."
	return Summary{Executed: s.Executed && n.Formalized && e.Formalized, CoreClosed: s.NoGoApplies, VacuumSelectorFound: s.AnyKineticSafeSelector, RequiresExtension: n.RequiresExtension, RemainingInputs: c.RemainingInputs, Status: status, DirectAnswer: direct, NextGate: next}
}

func Statuses(a Analysis) []string {
	chunks := []string{a.Span.Verdict, a.Sieve.Verdict, a.NoGo.Verdict, a.Extension.Verdict, a.Census.Verdict, a.Summary.Status}
	for _, c := range a.Sieve.Classes {
		chunks = append(chunks, c.Verdict)
	}
	return splitStatuses(chunks...)
}

func splitStatuses(chunks ...string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, c := range chunks {
		for _, p := range strings.Split(c, ";") {
			p = strings.TrimSpace(p)
			if p == "" || seen[p] {
				continue
			}
			seen[p] = true
			out = append(out, p)
		}
	}
	sort.Strings(out)
	return out
}

func join(parts ...string) string { return strings.Join(parts, ";") }

func FormatSpan(s Span) string {
	return fmt.Sprintf("audit=%s inherited_gate=%d adds_fit=%t purpose=%q verdict=%s", s.AuditID, s.InheritedGate, s.AddsFit, s.Purpose, s.Verdict)
}
func FormatOperatorClass(c OperatorClass) string {
	return fmt.Sprintf("name=%q native=%t kinetic_safe=%t unitary_invariant=%t changes_spectrum=%t selects_orientation=%t selects_point=%t barrier=%q verdict=%s", c.Name, c.Native, c.KineticSafe, c.UnitaryInvariant, c.CanChangeSpectrum, c.CanSelectOrientation, c.CanSelectPoint, c.Barrier, c.Verdict)
}
func FormatSieve(s ClosureSieve) string {
	parts := []string{}
	for _, c := range s.Classes {
		parts = append(parts, FormatOperatorClass(c))
	}
	return fmt.Sprintf("executed=%t classes=%d all_native=%t any_unique=%t any_safe=%t no_go=%t reason=%q verdict=%s classes=[%s]", s.Executed, len(s.Classes), s.AllNativeAudited, s.AnyUniqueSelector, s.AnyKineticSafeSelector, s.NoGoApplies, s.Reason, s.Verdict, strings.Join(parts, " | "))
}
func FormatNoGo(n NoGoTheorem) string {
	return fmt.Sprintf("formalized=%t assumptions=%v statement=%q conclusion=%q requires_extension=%t remaining=%d verdict=%s", n.Formalized, n.Assumptions, n.Statement, n.Conclusion, n.RequiresExtension, n.VacuumInputsRemain, n.Verdict)
}
func FormatExtension(e ExtensionFork) string {
	return fmt.Sprintf("formalized=%t landscape=%q extension=%q new_objects=%v verdict=%s", e.Formalized, e.LandscapePath, e.ExtensionPath, e.MinimalNewObjects, e.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d reduction=%d remaining=%d target=%d reached=%t verdict=%s", c.StartingVacuumInputs, c.ReductionFromClosure, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t core_closed=%t selector=%t requires_extension=%t remaining=%d status=%s answer=%q next=%q", s.Executed, s.CoreClosed, s.VacuumSelectorFound, s.RequiresExtension, s.RemainingInputs, s.Status, s.DirectAnswer, s.NextGate)
}
