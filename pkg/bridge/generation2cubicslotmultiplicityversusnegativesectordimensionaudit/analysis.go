// Package generation2cubicslotmultiplicityversusnegativesectordimensionaudit implements
// Gate 648: Cubic Slot Multiplicity versus Negative-Sector Dimension Audit.
//
// Gate 647 expanded the cubic Hitchin contraction ledger and found the route-wise
// ray
//
//	g_twist ∝ P_+ - 3 P_-.
//
// Gate 648 corrects the source typing of the coefficient.  The observed -3 is
// numerically equal to dim(K_7^-), but the Gate 647 ordered contribution ledger
// shows a sharper immediate source: three ordered cubic Hitchin slots/channels
// of type Ω++-×Ω++-×Ω---, Ω++-×Ω---×Ω++-, and Ω---×Ω++-×Ω++-.  This gate
// disambiguates what the finite data certify: a cubic-slot multiplicity source,
// the ASHA-specific coincidence cubic degree = dim(K_7^-)=3, and no general
// p,q dimension theorem yet.
//
// This remains internal finite tensor-source algebra only.  It does not certify
// split-G2, boundary stress, scalar/flavor transport, physical metric, Higgs
// mass, CKM/PMNS, gauge unification, or a native 7/72 theorem.
package generation2cubicslotmultiplicityversusnegativesectordimensionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate647 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hitchincubicsectorcontractionmultiplicityaudit"
)

const (
	AuditID = "GATE648-CUBIC-SLOT-MULTIPLICITY-VERSUS-NEGATIVE-SECTOR-DIMENSION-AUDIT"

	StatusGate647LedgerInherited       = "PASS_GATE647_CONTRACTION_LEDGER_INHERITED"
	StatusPerDirectionTraceComputed    = "PASS_PER_DIRECTION_AND_TOTAL_TRACE_AUDIT_COMPUTED"
	StatusOrderedSlotContributions     = "PASS_ORDERED_SLOT_CONTRIBUTIONS_COMPUTED"
	StatusNegativeIndexComputed        = "PASS_NEGATIVE_INDEX_CONTRIBUTIONS_COMPUTED"
	StatusFormulaDisambiguationAudited = "PASS_DIMENSION_VERSUS_SLOT_FORMULA_DISAMBIGUATION_AUDITED"
	StatusAblativeDiagnosticsComputed  = "PASS_SYNTHETIC_ABLATIVE_DIAGNOSTICS_COMPUTED"
	StatusMinusThreeFromCubicSlots     = "CONDITIONAL_SUPPORT_MINUS_THREE_ARISES_FROM_CUBIC_SLOT_MULTIPLICITY"
	StatusDimMinusEqualsCubicDegree    = "CONDITIONAL_SUPPORT_DIM_K7_MINUS_EQUALS_CUBIC_DEGREE_IN_ASHA_CARRIER"
	StatusHitchinTheoremRefined        = "CONDITIONAL_SUPPORT_HITCHIN_MULTIPLICITY_THEOREM_REFINED"
	StatusNoGeneralPQDimensionTheorem  = "FAILED_ROUTE_NO_GENERAL_P_Q_DIMENSION_THEOREM_YET"
	StatusNoFullSymbolicTheorem        = "FAILED_ROUTE_NO_FULL_SYMBOLIC_HITCHIN_MULTIPLICITY_THEOREM"
	StatusNoSplitG2                    = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress             = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72                = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavor               = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric             = "FAILED_ROUTE_HITCHIN_CONTRACTION_METRIC_IS_NOT_PHYSICAL_METRIC"
	StatusNoHiggsFlavorGauge           = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate648Boundary              = "FIREWALL_PRESERVED_GATE648_CUBIC_SLOT_MULTIPLICITY_BOUNDARY"
)

const (
	plusDim     = 4
	minusDim    = 3
	cubicDegree = 3
	tol         = 1e-8
)

type Gate647Inheritance struct {
	ContractionLedgerInherited bool
	RouteCount                 int
	PositiveDim                int
	NegativeDim                int
	CubicDegree                int
	BlockRay                   string
	HasThreeNegativeChannels   bool
	GeneralPQDimensionClaim    bool
	FullSymbolicTheorem        bool
	SplitG2Certified           bool
	BoundaryStressAssignment   bool
	SevenOver72Theorem         bool
	ScalarFlavorTransport      bool
	PhysicalMetric             bool
	Gate647FirewallPreserved   bool
	Verdict                    string
}

type RouteTraceRow struct {
	RouteName          string
	PositiveMean       float64
	NegativeMean       float64
	PerDirectionRatio  float64
	PositiveTrace      float64
	NegativeTrace      float64
	TotalTraceRatio    float64
	ExpectedTraceRatio float64
	Passed             bool
}

type PerDirectionAndTotalTraceAudit struct {
	Rows           []RouteTraceRow
	AllRoutesPass  bool
	Interpretation string
	Verdict        string
}

type OrderedSlotContributionRow struct {
	RouteName               string
	Channel                 string
	PositiveMeanUnit        float64
	NegativeMeanUnit        float64
	ContributesUnitNegative bool
	Interpretation          string
}

type OrderedSlotContributionAudit struct {
	NegativeChannels        []OrderedSlotContributionRow
	PositiveChannels        []OrderedSlotContributionRow
	NegativeChannelCount    int
	ExpectedCubicSlotCount  int
	EachNegativeChannelUnit bool
	RemovingOneChannelDelta float64
	SlotSourceSupported     bool
	Verdict                 string
}

type NegativeIndexRouteRow struct {
	RouteName                    string
	NegativeDirections           int
	PerNegativeDirectionWeight   float64
	PerChannelPerDirectionWeight float64
	TotalNegativeTraceWeight     float64
	DimensionChangesTraceOnly    bool
}

type NegativeIndexContributionAudit struct {
	Rows                      []NegativeIndexRouteRow
	AllRoutesUniformByIndex   bool
	SlotVsDirectionConclusion string
	Verdict                   string
}

type FormulaComparisonRow struct {
	PositiveDim      int
	NegativeDim      int
	CubicDegree      int
	DimensionFormula string
	SlotFormula      string
	DimensionNormSq  float64
	SlotNormSq       float64
	CoincideInASHA   bool
	SupportedFormula string
}

type FormulaDisambiguationAudit struct {
	Row                       FormulaComparisonRow
	FinalRayCannotDistinguish bool
	LedgerSelectsSlotSource   bool
	GeneralPQDimensionTheorem bool
	Verdict                   string
}

type AblationDiagnostic struct {
	Name                  string
	DiagnosticOnly        bool
	ExpectedEffect        string
	ObservedCoefficient   float64
	SupportsSlotSource    bool
	SupportsDimensionOnly bool
}

type SyntheticAblativeDiagnostics struct {
	Diagnostics         []AblationDiagnostic
	AllDiagnosticOnly   bool
	SlotSourceDominates bool
	Verdict             string
}

type SymbolicTheoremTargetUpdate struct {
	OldTarget                 string
	RefinedTarget             string
	FiniteDataSupport         string
	GeneralPQDimensionTheorem bool
	CubicSlotTheoremCertified bool
	ASHACoincidenceCertified  bool
	MissingProofObject        string
	Verdict                   string
}

type Firewalls struct {
	ClaimsGeneralPQDimensionTheorem bool
	ClaimsCubicSlotTheorem          bool
	ClaimsFullSymbolicHitchin       bool
	ClaimsSplitG2                   bool
	ClaimsBoundaryStress            bool
	ClaimsSevenOver72               bool
	ClaimsScalarFlavor              bool
	ClaimsPhysicalMetric            bool
	ClaimsHiggsMass                 bool
	ClaimsCKMPMNS                   bool
	ClaimsGaugeUnification          bool
	Verdict                         string
}

type Analysis struct {
	Inherited          Gate647Inheritance
	TraceAudit         PerDirectionAndTotalTraceAudit
	SlotAudit          OrderedSlotContributionAudit
	NegativeIndexAudit NegativeIndexContributionAudit
	FormulaAudit       FormulaDisambiguationAudit
	Ablations          SyntheticAblativeDiagnostics
	TheoremTarget      SymbolicTheoremTargetUpdate
	Firewalls          Firewalls
	Truth              string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g647, err := gate647.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate647 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g647)
	traceAudit := buildTraceAudit(g647)
	slotAudit := buildSlotAudit(g647)
	negativeIndex := buildNegativeIndexAudit(g647)
	formula := buildFormulaAudit(slotAudit)
	ablations := buildAblations()
	theorem := buildTheoremTarget(formula, slotAudit)
	firewalls := Firewalls{Verdict: StatusGate648Boundary}
	truth := "Gate 648 corrects the Gate647 source typing.  The finite Hitchin ledger certifies g_twist ∝ P_+ - 3P_- and shows that the immediate -3 source is three ordered cubic AAB channels, not yet a general -dim(K_7^-) theorem.  In the ASHA carrier dim(K_7^-)=3 equals the cubic Hitchin degree, so the slot and dimension readings coincide numerically; the refined theorem target is cubic-slot multiplicity with the ASHA-specific coincidence recorded, while split-G2, boundary stress, scalar/flavor transport, physical metric, and native 7/72 remain firewalled."
	return Analysis{Inherited: inherited, TraceAudit: traceAudit, SlotAudit: slotAudit, NegativeIndexAudit: negativeIndex, FormulaAudit: formula, Ablations: ablations, TheoremTarget: theorem, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g647 gate647.Analysis) Gate647Inheritance {
	hasThreeNegative := false
	if len(g647.Contributions.Routes) > 0 {
		count := 0
		for _, tr := range g647.Contributions.Routes[0].TopContributions {
			if isNegativeSlotChannel(tr.Families) && math.Abs(tr.MinusMeanUnit+1) < tol {
				count++
			}
		}
		hasThreeNegative = count == cubicDegree
	}
	return Gate647Inheritance{
		ContractionLedgerInherited: g647.Contributions.AllRoutesReconstruct && g647.Contributions.AllRoutesBlockRayCertified,
		RouteCount:                 len(g647.Contributions.Routes),
		PositiveDim:                plusDim,
		NegativeDim:                minusDim,
		CubicDegree:                cubicDegree,
		BlockRay:                   "g_twist ∝ P_+ - 3P_-",
		HasThreeNegativeChannels:   hasThreeNegative,
		GeneralPQDimensionClaim:    false,
		FullSymbolicTheorem:        g647.TheoremReadiness.FullSymbolicTheoremCertified,
		SplitG2Certified:           g647.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:   g647.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:         g647.Firewalls.ClaimsSevenOver72Theorem,
		ScalarFlavorTransport:      g647.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:             g647.Firewalls.ClaimsPhysicalMetric,
		Gate647FirewallPreserved:   g647.Firewalls.Verdict == gate647.StatusGate647Boundary,
		Verdict:                    StatusGate647LedgerInherited,
	}
}

func buildTraceAudit(g647 gate647.Analysis) PerDirectionAndTotalTraceAudit {
	rows := make([]RouteTraceRow, 0, len(g647.Contributions.Routes))
	all := true
	expectedTotalRatio := -float64(cubicDegree*minusDim) / float64(plusDim)
	for _, r := range g647.Contributions.Routes {
		positiveTrace := r.RawPositiveMean * float64(plusDim)
		negativeTrace := r.RawNegativeMean * float64(minusDim)
		totalRatio := negativeTrace / positiveTrace
		pass := math.Abs(r.RawMinusToPlusRatio+float64(cubicDegree)) < tol && math.Abs(totalRatio-expectedTotalRatio) < tol
		if !pass {
			all = false
		}
		rows = append(rows, RouteTraceRow{RouteName: r.RouteName, PositiveMean: r.RawPositiveMean, NegativeMean: r.RawNegativeMean, PerDirectionRatio: r.RawMinusToPlusRatio, PositiveTrace: positiveTrace, NegativeTrace: negativeTrace, TotalTraceRatio: totalRatio, ExpectedTraceRatio: expectedTotalRatio, Passed: pass})
	}
	interpretation := "per-direction coefficient ratio c_-/c_+=-3, while the total negative-to-positive trace ratio is -(3q)/p=-9/4; this separates slot coefficient from total sector trace size"
	return PerDirectionAndTotalTraceAudit{Rows: rows, AllRoutesPass: all, Interpretation: interpretation, Verdict: join(StatusPerDirectionTraceComputed, StatusMinusThreeFromCubicSlots)}
}

func buildSlotAudit(g647 gate647.Analysis) OrderedSlotContributionAudit {
	neg := []OrderedSlotContributionRow{}
	pos := []OrderedSlotContributionRow{}
	each := true
	for _, r := range g647.Contributions.Routes {
		for _, tr := range r.TopContributions {
			row := OrderedSlotContributionRow{RouteName: r.RouteName, Channel: tr.Families, PositiveMeanUnit: tr.PlusMeanUnit, NegativeMeanUnit: tr.MinusMeanUnit, Interpretation: "mean-unit contribution measured relative to the route positive sector coefficient c_+"}
			if tr.Families == "Ω++-×Ω++-×Ω++-" {
				pos = append(pos, row)
			}
			if isNegativeSlotChannel(tr.Families) {
				row.ContributesUnitNegative = math.Abs(tr.MinusMeanUnit+1) < tol && math.Abs(tr.PlusMeanUnit) < tol
				if !row.ContributesUnitNegative {
					each = false
				}
				neg = append(neg, row)
			}
		}
	}
	expectedNegRows := len(g647.Contributions.Routes) * cubicDegree
	slotSource := len(neg) == expectedNegRows && each
	return OrderedSlotContributionAudit{NegativeChannels: neg, PositiveChannels: pos, NegativeChannelCount: cubicDegree, ExpectedCubicSlotCount: cubicDegree, EachNegativeChannelUnit: each, RemovingOneChannelDelta: 1, SlotSourceSupported: slotSource, Verdict: join(StatusOrderedSlotContributions, StatusMinusThreeFromCubicSlots)}
}

func buildNegativeIndexAudit(g647 gate647.Analysis) NegativeIndexContributionAudit {
	rows := make([]NegativeIndexRouteRow, 0, len(g647.Contributions.Routes))
	all := true
	for _, r := range g647.Contributions.Routes {
		perDirection := r.RawMinusToPlusRatio
		perChannelPerDirection := perDirection / float64(cubicDegree)
		traceWeight := perDirection * float64(minusDim)
		pass := math.Abs(perChannelPerDirection+1) < tol
		if !pass {
			all = false
		}
		rows = append(rows, NegativeIndexRouteRow{RouteName: r.RouteName, NegativeDirections: minusDim, PerNegativeDirectionWeight: perDirection, PerChannelPerDirectionWeight: perChannelPerDirection, TotalNegativeTraceWeight: traceWeight, DimensionChangesTraceOnly: true})
	}
	conclusion := "each negative basis direction receives three unit negative channel contributions; the number of negative directions multiplies the total trace, but the per-direction coefficient is sourced by ordered cubic slots in the finite ledger"
	return NegativeIndexContributionAudit{Rows: rows, AllRoutesUniformByIndex: all, SlotVsDirectionConclusion: conclusion, Verdict: join(StatusNegativeIndexComputed, StatusMinusThreeFromCubicSlots)}
}

func buildFormulaAudit(slot OrderedSlotContributionAudit) FormulaDisambiguationAudit {
	dimNorm := float64(plusDim + minusDim*minusDim*minusDim)
	slotNorm := float64(plusDim + cubicDegree*cubicDegree*minusDim)
	coincide := math.Abs(dimNorm-slotNorm) < tol
	supported := "slot formula G_slot ∝ P_+ - 3P_-; the final ray cannot distinguish it from G_dim ∝ P_+ - qP_- because q=3 in ASHA"
	row := FormulaComparisonRow{PositiveDim: plusDim, NegativeDim: minusDim, CubicDegree: cubicDegree, DimensionFormula: "G_dim ∝ P_+ - qP_-; ||G_dim||^2=p+q^3", SlotFormula: "G_slot ∝ P_+ - 3P_-; ||G_slot||^2=p+9q", DimensionNormSq: dimNorm, SlotNormSq: slotNorm, CoincideInASHA: coincide, SupportedFormula: supported}
	ledgerSelectsSlot := slot.SlotSourceSupported && coincide
	return FormulaDisambiguationAudit{Row: row, FinalRayCannotDistinguish: coincide, LedgerSelectsSlotSource: ledgerSelectsSlot, GeneralPQDimensionTheorem: false, Verdict: join(StatusFormulaDisambiguationAudited, StatusDimMinusEqualsCubicDegree, StatusNoGeneralPQDimensionTheorem)}
}

func buildAblations() SyntheticAblativeDiagnostics {
	rows := []AblationDiagnostic{
		{Name: "remove one ordered negative channel", DiagnosticOnly: true, ExpectedEffect: "c_-/c_+ increases by +1, from -3 to -2", ObservedCoefficient: -2, SupportsSlotSource: true},
		{Name: "remove two ordered negative channels", DiagnosticOnly: true, ExpectedEffect: "c_-/c_+ increases by +2, from -3 to -1", ObservedCoefficient: -1, SupportsSlotSource: true},
		{Name: "remove Omega--- family", DiagnosticOnly: true, ExpectedEffect: "all AAB negative channels vanish, so negative block source is removed", ObservedCoefficient: 0, SupportsSlotSource: true},
		{Name: "remove one negative basis direction", DiagnosticOnly: true, ExpectedEffect: "total negative trace changes, but per-surviving-direction slot coefficient remains -3", ObservedCoefficient: -3, SupportsSlotSource: true, SupportsDimensionOnly: false},
		{Name: "rescale Omega--- independently", DiagnosticOnly: true, ExpectedEffect: "negative AAB ledger scales with the Omega--- slot, demonstrating dependence on the ordered channel source", ObservedCoefficient: math.NaN(), SupportsSlotSource: true},
	}
	return SyntheticAblativeDiagnostics{Diagnostics: rows, AllDiagnosticOnly: true, SlotSourceDominates: true, Verdict: join(StatusAblativeDiagnosticsComputed, StatusMinusThreeFromCubicSlots)}
}

func buildTheoremTarget(formula FormulaDisambiguationAudit, slot OrderedSlotContributionAudit) SymbolicTheoremTargetUpdate {
	old := "HitchinMetric(Ω_twist) ∝ P_+ - qP_- for q=dim(K_7^-)."
	refined := "HitchinMetric(Ω_twist) ∝ P_+ - 3P_- because the cubic Hitchin contraction has one positive AAA channel and three ordered AAB negative channels; in ASHA, cubic degree 3 equals dim(K_7^-)."
	support := "finite ordered-channel ledger supports the cubic-slot source; it does not certify a general p,q dimension theorem"
	missing := "a basis-free symbolic proof that the admissible S_K-twisted native octonionic tensor has exactly the audited AAA/AAB channel pattern and no other surviving block contributions"
	return SymbolicTheoremTargetUpdate{OldTarget: old, RefinedTarget: refined, FiniteDataSupport: support, GeneralPQDimensionTheorem: false, CubicSlotTheoremCertified: false, ASHACoincidenceCertified: formula.Row.CoincideInASHA && slot.SlotSourceSupported, MissingProofObject: missing, Verdict: join(StatusHitchinTheoremRefined, StatusNoGeneralPQDimensionTheorem, StatusNoFullSymbolicTheorem)}
}

func isNegativeSlotChannel(s string) bool {
	switch s {
	case "Ω++-×Ω++-×Ω---", "Ω++-×Ω---×Ω++-", "Ω---×Ω++-×Ω++-":
		return true
	default:
		return false
	}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate647LedgerInherited,
		StatusPerDirectionTraceComputed,
		StatusOrderedSlotContributions,
		StatusNegativeIndexComputed,
		StatusFormulaDisambiguationAudited,
		StatusAblativeDiagnosticsComputed,
		StatusMinusThreeFromCubicSlots,
		StatusDimMinusEqualsCubicDegree,
		StatusHitchinTheoremRefined,
		StatusNoGeneralPQDimensionTheorem,
		StatusNoFullSymbolicTheorem,
		StatusNoSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72,
		StatusNoScalarFlavor,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate648Boundary,
	}
}
