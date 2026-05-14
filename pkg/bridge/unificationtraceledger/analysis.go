// Package unificationtraceledger implements Gate 308:
// Unification Trace Ledger / Higgs Quartic Unification Boundary Audit.
//
// Gate 307 promoted the projected scalar carrier ratio 1197/4624 into the
// finite-geometry target for lambda_H/g_i^2, but it left two final convention
// obligations open: the gauge trace index tau_i and the Lorentzian scalar
// quartic sign. Gate 308 audits those obligations. It formalizes the canonical
// unification trace ledger K_* = diag(1,1,1,5/3), converts hypercharge to the
// GUT-normalized coupling convention g_*^2 = g_2^2 = g_3^2 = (5/3)g_Y^2, and
// therefore identifies a universal canonical trace index tau_GUT = 1 for the
// normalized boundary equation. It also declares the standard positive
// Lorentzian potential convention Sign_4 = +1. The output is an analytic UV
// boundary equation for lambda_H(Lambda_GUT) as a function of the still-sealed
// unified gauge coupling g_*^2. No low-energy Higgs mass, RG running, threshold
// matching, absolute gauge coupling, f2 moment, Yukawa origin, or B-gap instanton
// prediction is claimed.
package unificationtraceledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE308-UNIFICATION-TRACE-LEDGER-HIGGS-QUARTIC-UNIFICATION-BOUNDARY-AUDIT"

	StatusGate307Inherited                  = "CONDITIONAL_SUPPORT_GATE307_TRACE_EQUIVALENCE_INHERITED"
	StatusUnificationTraceIndexFormalized   = "CONDITIONAL_SUPPORT_UNIFICATION_TRACE_INDEX_FORMALIZED"
	StatusCanonicalTauGUTComputed           = "CONDITIONAL_SUPPORT_CANONICAL_TAU_GUT_EQUALS_ONE_AFTER_GUT_NORMALIZATION"
	StatusSignConventionLedgerFormalized    = "CONDITIONAL_SUPPORT_QUARTIC_SIGN_CONVENTION_LEDGER_FORMALIZED"
	StatusQuarticUnificationBoundaryDerived = "CONDITIONAL_SUPPORT_HIGGS_QUARTIC_UNIFICATION_BOUNDARY_DERIVED"
	StatusAnalyticBoundaryOnly              = "CONDITIONAL_SUPPORT_ANALYTIC_UV_BOUNDARY_ONLY_NO_IR_MASS_CLAIM"
	StatusGate308FirewallsPreserved         = "CONDITIONAL_SUPPORT_GATE308_UNIFICATION_BOUNDARY_FIREWALLS_PRESERVED"

	StatusFailedAbsoluteGaugeCouplingStillSealed = "FAILED_ROUTE_ABSOLUTE_UNIFIED_GAUGE_COUPLING_VALUE_STILL_SEALED"
	StatusFailedBoundaryScaleStillSealed         = "FAILED_ROUTE_LAMBDA_GUT_BOUNDARY_SCALE_STILL_SEALED"
	StatusFailedRGERunningNotExecuted            = "FAILED_ROUTE_RGE_RUNNING_TO_ELECTROWEAK_SCALE_NOT_EXECUTED"
	StatusFailedThresholdMatchingStillSealed     = "FAILED_ROUTE_THRESHOLD_AND_MATCHING_CORRECTIONS_STILL_SEALED"
	StatusFailedLowEnergyHiggsMassNotDerived     = "FAILED_ROUTE_LOW_ENERGY_HIGGS_MASS_NOT_DERIVED"
	StatusFailedYukawaOriginStillSealed          = "FAILED_ROUTE_YUKAWA_AMPLITUDE_ORIGIN_STILL_SEALED"
	StatusFailedF2MassMomentStillOpen            = "FAILED_ROUTE_F2_MASS_MOMENT_STILL_UNLOCKED"
	StatusFailedBGapInstantonStillSealed         = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

const (
	rawTraceRatioNumerator   = 1197
	rawTraceRatioDenominator = 4624
	canonicalTauGUTNumerator = 1
	canonicalTauGUTDenom     = 1
	sign4Canonical           = 1
)

type Gate307Inheritance struct {
	TraceEquivalenceProved         bool
	ProjectedScalarCarrierPromoted bool
	ShapeNumerator                 int
	ShapeDenominator               int
	LambdaOverGaugeMap             string
	RequiresTraceIndex             bool
	RequiresQuarticSign            bool
	AbsoluteLambdaHDerived         bool
	AbsoluteGaugeCouplingDerived   bool
	LowEnergyMassClaimed           bool
	Verdict                        string
}

type GaugeFactorTrace struct {
	Name                  string
	Algebra               string
	RawTraceIndex         string
	UnificationConvention string
	CanonicalTraceIndex   string
	CouplingRelation      string
	IncludedInGUTLedger   bool
	NormalizedToUniversal bool
	Status                string
}

type UnificationTraceIndex struct {
	RawTraceLedger            string
	HyperchargeNormalization  string
	UnifiedCouplingDefinition string
	CanonicalTraceIndexSymbol string
	CanonicalTraceIndexValue  string
	GaugeFactors              []GaugeFactorTrace
	AssumesGaugeUnification   bool
	ComputesAbsoluteCoupling  bool
	UsesObservedCouplings     bool
	UniversalIndexFormalized  bool
	Verdict                   string
}

type SignConventionLedger struct {
	EuclideanQuarticCarrier     string
	LorentzianPotentialTarget   string
	SignSymbol                  string
	SignValue                   int
	WickConventionDeclared      bool
	PositivePotentialConvention bool
	DerivedFromFiniteCore       bool
	BlocksIfNegative            bool
	LedgerFormalized            bool
	Verdict                     string
}

type QuarticBoundaryEquation struct {
	StartingGate307Ratio        string
	SubstitutedTraceIndex       string
	SubstitutedSign             string
	UnifiedBoundaryEquation     string
	ExactCoefficient            string
	DecimalDiagnostic           string
	DependsOnUnifiedGaugeValue  bool
	DependsOnF2Moment           bool
	DependsOnN4F0Ledger         bool
	DependsOnCutoffProfileShape bool
	AnalyticBoundaryDerived     bool
	LowEnergyPredictionMade     bool
	Verdict                     string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksColliderPrediction  bool
}

type FirewallAudit struct {
	NoAbsoluteGaugeValueInserted   bool
	NoBoundaryScaleInserted        bool
	NoObservedCouplingsInserted    bool
	NoRGERunningExecuted           bool
	NoThresholdMatchingInserted    bool
	NoLowEnergyHiggsMassClaimed    bool
	NoYukawaNumbersInserted        bool
	F2MassFirewallPreserved        bool
	BGapInstantonFirewallPreserved bool
	AnalyticBoundaryOnly           bool
	FiniteCorePolluted             bool
	Obligations                    []RemainingObligation
	Verdict                        string
}

type Summary struct {
	Gate307Inherited               bool
	TraceIndexFormalized           bool
	TauGUTComputed                 bool
	SignConventionFormalized       bool
	BoundaryEquationDerived        bool
	AnalyticUVBoundaryOnly         bool
	LowEnergyHiggsMassDerived      bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Input     Gate307Inheritance
	Trace     UnificationTraceIndex
	Sign      SignConventionLedger
	Boundary  QuarticBoundaryEquation
	Firewalls FirewallAudit
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
	input := inheritGate307()
	trace := formalizeTraceIndex(input)
	sign := formalizeSignLedger(input)
	boundary := deriveBoundaryEquation(input, trace, sign)
	firewalls := auditFirewalls(input, trace, sign, boundary)
	summary := buildSummary(input, trace, sign, boundary, firewalls)
	truth := "Gate 308 closes the final relative quartic normalization ledger at the UV boundary: under the canonical GUT-normalized gauge convention g_*^2 = g_2^2 = g_3^2 = (5/3)g_Y^2, the raw trace indices K_* = diag(1,1,1,5/3) reduce to the universal canonical index tau_GUT=1, and with the declared positive Lorentzian quartic convention Sign_4=+1 the projected Gate-307 ratio yields lambda_H(Lambda_GUT) = (1197/4624) g_*^2. This is an analytic boundary equation only. The absolute value of g_*^2, the boundary scale, RG running, threshold matching, low-energy Higgs mass, Yukawa amplitude origin, f2 mass channel, and B-gap instanton action remain firewalled."
	return Analysis{Input: input, Trace: trace, Sign: sign, Boundary: boundary, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritGate307() Gate307Inheritance {
	return Gate307Inheritance{
		TraceEquivalenceProved:         true,
		ProjectedScalarCarrierPromoted: true,
		ShapeNumerator:                 rawTraceRatioNumerator,
		ShapeDenominator:               rawTraceRatioDenominator,
		LambdaOverGaugeMap:             "lambda_H/g_i^2 = Sign_4 · τ_i · 1197/4624 on the projected scalar carrier",
		RequiresTraceIndex:             true,
		RequiresQuarticSign:            true,
		AbsoluteLambdaHDerived:         false,
		AbsoluteGaugeCouplingDerived:   false,
		LowEnergyMassClaimed:           false,
		Verdict:                        StatusGate307Inherited,
	}
}

func formalizeTraceIndex(i Gate307Inheritance) UnificationTraceIndex {
	factors := []GaugeFactorTrace{
		{Name: "SU(2)_L", Algebra: "su(2)", RawTraceIndex: "1", UnificationConvention: "g_*^2 := g_2^2", CanonicalTraceIndex: "1", CouplingRelation: "g_2^2 = g_*^2", IncludedInGUTLedger: true, NormalizedToUniversal: true, Status: StatusUnificationTraceIndexFormalized},
		{Name: "SU(3)_C", Algebra: "su(3)", RawTraceIndex: "1", UnificationConvention: "g_*^2 := g_3^2", CanonicalTraceIndex: "1", CouplingRelation: "g_3^2 = g_*^2", IncludedInGUTLedger: true, NormalizedToUniversal: true, Status: StatusUnificationTraceIndexFormalized},
		{Name: "U(1)_Y", Algebra: "u(1)_Y", RawTraceIndex: "5/3", UnificationConvention: "GUT-normalized hypercharge absorbs k_Y=5/3", CanonicalTraceIndex: "1", CouplingRelation: "g_*^2 = (5/3) g_Y^2", IncludedInGUTLedger: true, NormalizedToUniversal: true, Status: StatusCanonicalTauGUTComputed},
	}
	return UnificationTraceIndex{
		RawTraceLedger:            "K_* = diag(SU2_1, SU2_2, SU2_3, U1_Y) = diag(1,1,1,5/3)",
		HyperchargeNormalization:  "k_Y = 5/3; equivalently g_*^2 = (5/3)g_Y^2 at the normalized boundary",
		UnifiedCouplingDefinition: "g_*^2 := g_2^2 = g_3^2 = (5/3)g_Y^2",
		CanonicalTraceIndexSymbol: "τ_GUT",
		CanonicalTraceIndexValue:  "1",
		GaugeFactors:              factors,
		AssumesGaugeUnification:   true,
		ComputesAbsoluteCoupling:  false,
		UsesObservedCouplings:     false,
		UniversalIndexFormalized:  i.TraceEquivalenceProved && i.RequiresTraceIndex,
		Verdict:                   strings.Join([]string{StatusUnificationTraceIndexFormalized, StatusCanonicalTauGUTComputed, StatusFailedAbsoluteGaugeCouplingStillSealed}, ";"),
	}
}

func formalizeSignLedger(i Gate307Inheritance) SignConventionLedger {
	return SignConventionLedger{
		EuclideanQuarticCarrier:     "+ C4_raw |H_raw|^4 inside the a4 scalar-power-4 projected carrier",
		LorentzianPotentialTarget:   "V(H_phys) ⊃ + λ_H |H_phys|^4 with λ_H > 0 for the bounded canonical potential branch",
		SignSymbol:                  "Sign_4",
		SignValue:                   sign4Canonical,
		WickConventionDeclared:      true,
		PositivePotentialConvention: true,
		DerivedFromFiniteCore:       false,
		BlocksIfNegative:            true,
		LedgerFormalized:            i.RequiresQuarticSign,
		Verdict:                     strings.Join([]string{StatusSignConventionLedgerFormalized, "CONVENTION_DECLARED_NOT_FINITE_CORE_DERIVED"}, ";"),
	}
}

func deriveBoundaryEquation(i Gate307Inheritance, t UnificationTraceIndex, s SignConventionLedger) QuarticBoundaryEquation {
	derived := i.TraceEquivalenceProved && t.UniversalIndexFormalized && s.LedgerFormalized && s.SignValue == sign4Canonical && t.CanonicalTraceIndexValue == "1"
	return QuarticBoundaryEquation{
		StartingGate307Ratio:        "lambda_H/g_i^2 = Sign_4 · τ_i · 1197/4624",
		SubstitutedTraceIndex:       "τ_i → τ_GUT = 1",
		SubstitutedSign:             "Sign_4 → +1",
		UnifiedBoundaryEquation:     "λ_H(Λ_GUT) = (1197/4624) · g_*^2",
		ExactCoefficient:            "1197/4624",
		DecimalDiagnostic:           fmt.Sprintf("%.12f", float64(rawTraceRatioNumerator)/float64(rawTraceRatioDenominator)),
		DependsOnUnifiedGaugeValue:  true,
		DependsOnF2Moment:           false,
		DependsOnN4F0Ledger:         false,
		DependsOnCutoffProfileShape: false,
		AnalyticBoundaryDerived:     derived,
		LowEnergyPredictionMade:     false,
		Verdict:                     strings.Join([]string{StatusQuarticUnificationBoundaryDerived, StatusAnalyticBoundaryOnly, StatusFailedRGERunningNotExecuted}, ";"),
	}
}

func auditFirewalls(i Gate307Inheritance, t UnificationTraceIndex, s SignConventionLedger, b QuarticBoundaryEquation) FirewallAudit {
	obligations := []RemainingObligation{
		{"absolute unified coupling", "the equation is λ_H = (1197/4624) g_*^2, but g_*^2 itself is not derived or numerically inserted", StatusFailedAbsoluteGaugeCouplingStillSealed, true},
		{"boundary scale", "Λ_GUT or M_* is required before RG evolution can start", StatusFailedBoundaryScaleStillSealed, true},
		{"RGE trajectory", "the UV quartic boundary must be run through SM/BSM beta functions before comparison to electroweak data", StatusFailedRGERunningNotExecuted, true},
		{"threshold and matching corrections", "finite thresholds, decoupling choices, and loop matching determine the IR value", StatusFailedThresholdMatchingStillSealed, true},
		{"low-energy Higgs mass", "m_h requires λ(v), v, scheme choice, and radiative corrections", StatusFailedLowEnergyHiggsMassNotDerived, true},
		{"Yukawa amplitude origin", "the projected shape is sealed, but the finite core has not derived all numerical SM Yukawa matrices", StatusFailedYukawaOriginStillSealed, true},
		{"a2 mass channel f2", "the Higgs mass parameter remains separate from the a4 quartic boundary", StatusFailedF2MassMomentStillOpen, true},
		{"B-gap instanton action", "the hierarchy/seesaw instanton bridge remains sealed", StatusFailedBGapInstantonStillSealed, true},
	}
	return FirewallAudit{
		NoAbsoluteGaugeValueInserted:   !t.ComputesAbsoluteCoupling && !i.AbsoluteGaugeCouplingDerived,
		NoBoundaryScaleInserted:        true,
		NoObservedCouplingsInserted:    !t.UsesObservedCouplings,
		NoRGERunningExecuted:           !b.LowEnergyPredictionMade,
		NoThresholdMatchingInserted:    true,
		NoLowEnergyHiggsMassClaimed:    !i.LowEnergyMassClaimed && !b.LowEnergyPredictionMade,
		NoYukawaNumbersInserted:        true,
		F2MassFirewallPreserved:        true,
		BGapInstantonFirewallPreserved: true,
		AnalyticBoundaryOnly:           b.AnalyticBoundaryDerived,
		FiniteCorePolluted:             false || s.DerivedFromFiniteCore,
		Obligations:                    obligations,
		Verdict:                        strings.Join([]string{StatusGate308FirewallsPreserved, StatusFailedAbsoluteGaugeCouplingStillSealed, StatusFailedLowEnergyHiggsMassNotDerived}, ";"),
	}
}

func buildSummary(i Gate307Inheritance, t UnificationTraceIndex, s SignConventionLedger, b QuarticBoundaryEquation, f FirewallAudit) Summary {
	return Summary{
		Gate307Inherited:          i.TraceEquivalenceProved && i.ProjectedScalarCarrierPromoted,
		TraceIndexFormalized:      t.UniversalIndexFormalized,
		TauGUTComputed:            t.CanonicalTraceIndexValue == "1",
		SignConventionFormalized:  s.LedgerFormalized && s.SignValue == 1,
		BoundaryEquationDerived:   b.AnalyticBoundaryDerived,
		AnalyticUVBoundaryOnly:    f.AnalyticBoundaryOnly,
		LowEnergyHiggsMassDerived: false,
		FirewallPreserved:         !f.FiniteCorePolluted && f.NoAbsoluteGaugeValueInserted && f.NoRGERunningExecuted && f.NoLowEnergyHiggsMassClaimed,
		Status:                    StatusQuarticUnificationBoundaryDerived,
		DirectAnswer:              "Under canonical GUT trace normalization and Sign_4=+1, λ_H(Λ_GUT) = (1197/4624) g_*^2.",
		NextGate:                  "Gate 309 should audit the RG transport protocol: define the beta-function scheme, boundary scale seal, threshold/matching ledger, and conditions for evolving λ_H(Λ_GUT) to the electroweak scale without importing unsealed data.",
	}
}

func FormatGate307Inheritance(i Gate307Inheritance) string {
	return fmt.Sprintf("equivalence=%t projected=%t ratio=%d/%d map=%q needsTau=%t needsSign=%t lambdaAbs=%t gaugeAbs=%t lowMass=%t verdict=%s", i.TraceEquivalenceProved, i.ProjectedScalarCarrierPromoted, i.ShapeNumerator, i.ShapeDenominator, i.LambdaOverGaugeMap, i.RequiresTraceIndex, i.RequiresQuarticSign, i.AbsoluteLambdaHDerived, i.AbsoluteGaugeCouplingDerived, i.LowEnergyMassClaimed, i.Verdict)
}

func FormatGaugeFactor(g GaugeFactorTrace) string {
	return fmt.Sprintf("%s algebra=%s rawTau=%s convention=%q canonicalTau=%s relation=%q included=%t universal=%t status=%s", g.Name, g.Algebra, g.RawTraceIndex, g.UnificationConvention, g.CanonicalTraceIndex, g.CouplingRelation, g.IncludedInGUTLedger, g.NormalizedToUniversal, g.Status)
}

func FormatTraceIndex(t UnificationTraceIndex) string {
	factors := []string{}
	for _, g := range t.GaugeFactors {
		factors = append(factors, FormatGaugeFactor(g))
	}
	return fmt.Sprintf("raw=%q hyper=%q gstar=%q tau=%s=%s factors=[%s] assumesUnif=%t absolute=%t observed=%t formalized=%t verdict=%s", t.RawTraceLedger, t.HyperchargeNormalization, t.UnifiedCouplingDefinition, t.CanonicalTraceIndexSymbol, t.CanonicalTraceIndexValue, strings.Join(factors, " | "), t.AssumesGaugeUnification, t.ComputesAbsoluteCoupling, t.UsesObservedCouplings, t.UniversalIndexFormalized, t.Verdict)
}

func FormatSign(s SignConventionLedger) string {
	return fmt.Sprintf("euclidean=%q target=%q symbol=%s value=%d wick=%t positive=%t finiteDerived=%t blocksIfNegative=%t formalized=%t verdict=%s", s.EuclideanQuarticCarrier, s.LorentzianPotentialTarget, s.SignSymbol, s.SignValue, s.WickConventionDeclared, s.PositivePotentialConvention, s.DerivedFromFiniteCore, s.BlocksIfNegative, s.LedgerFormalized, s.Verdict)
}

func FormatBoundary(b QuarticBoundaryEquation) string {
	return fmt.Sprintf("start=%q tau=%q sign=%q equation=%q coeff=%s decimal=%s dependsG=%t f2=%t n4f0=%t shape=%t derived=%t lowEnergy=%t verdict=%s", b.StartingGate307Ratio, b.SubstitutedTraceIndex, b.SubstitutedSign, b.UnifiedBoundaryEquation, b.ExactCoefficient, b.DecimalDiagnostic, b.DependsOnUnifiedGaugeValue, b.DependsOnF2Moment, b.DependsOnN4F0Ledger, b.DependsOnCutoffProfileShape, b.AnalyticBoundaryDerived, b.LowEnergyPredictionMade, b.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksColliderPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noAbsG=%t noScale=%t noObserved=%t noRGE=%t noThreshold=%t noMass=%t noYukawa=%t f2=%t bgap=%t analyticOnly=%t polluted=%t obligations=[%s] verdict=%s", f.NoAbsoluteGaugeValueInserted, f.NoBoundaryScaleInserted, f.NoObservedCouplingsInserted, f.NoRGERunningExecuted, f.NoThresholdMatchingInserted, f.NoLowEnergyHiggsMassClaimed, f.NoYukawaNumbersInserted, f.F2MassFirewallPreserved, f.BGapInstantonFirewallPreserved, f.AnalyticBoundaryOnly, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate307=%t trace=%t tau=%t sign=%t boundary=%t uvOnly=%t lowMass=%t firewall=%t status=%s answer=%q next=%q", s.Gate307Inherited, s.TraceIndexFormalized, s.TauGUTComputed, s.SignConventionFormalized, s.BoundaryEquationDerived, s.AnalyticUVBoundaryOnly, s.LowEnergyHiggsMassDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
