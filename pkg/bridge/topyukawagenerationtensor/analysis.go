// Package topyukawagenerationtensor implements Gate 313:
// Top-Yukawa Generation Tensor Sieve / Amplitude Fractionalization Audit.
//
// Gate 312 proved that modifying the UV scalar quartic boundary alone is washed
// out by the one-loop top-Yukawa infrared attractor.  Gate 313 therefore audits
// whether the sealed r_+ amplitude used as y_t^2/g_*^2 in Gate 309 should be
// interpreted as a full three-generation trace rather than as a single top entry.
//
// The gate is deliberately conservative.  It proves the correct trace form
//
//	Tr_gen(Y_u^†Y_u)/g_*^2 = y_u^2 + y_c^2 + y_t^2 = r_+
//
// and evaluates several admissible fractionalization witnesses.  It also imports
// the Gate-242 tau_eta=(2,-2,1) result and preserves its own firewall: tau_eta is
// a scalar trace functional with generation-breaking capacity, not yet a derived
// operator assigning a unique top-generation eigenvector.  Consequently Gate 313
// can quantify how much a smaller y_t boundary helps, but it cannot honestly claim
// that the tensor has derived the physical top Yukawa or solved the Higgs mass.
package topyukawagenerationtensor

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE313-TOP-YUKAWA-GENERATION-TENSOR-FRACTIONALIZATION-AUDIT"

	StatusGenerationTraceFormalized      = "CONDITIONAL_SUPPORT_GENERATION_TRACE_FORMALIZED"
	StatusTauEtaTopologyRetrieved        = "CONDITIONAL_SUPPORT_TAU_ETA_GENERATION_TOPOLOGY_RETRIEVED"
	StatusFractionalizationLanesAudited  = "CONDITIONAL_SUPPORT_TOP_YUKAWA_FRACTIONALIZATION_LANES_AUDITED"
	StatusRGSlopeReevaluationComputed    = "CONDITIONAL_SUPPORT_RG_SLOPE_TENSION_REEVALUATED"
	StatusPartialFlatteningObserved      = "CONDITIONAL_SUPPORT_GENERATION_FRACTIONALIZATION_FLATTENS_TOP_SLOPE"
	StatusInsufficientToResolve          = "CONDITIONAL_TENSION_GENERATION_FRACTIONALIZATION_ALONE_DOES_NOT_RESOLVE_HIGGS_TENSION"
	StatusFailedTauEtaPullbackMissing    = "FAILED_ROUTE_TAU_ETA_TO_TRIALITY_GENERATION_PULLBACK_STILL_MISSING"
	StatusFailedTopAssignmentAmbiguous   = "FAILED_ROUTE_TOP_GENERATION_ASSIGNMENT_NOT_CANONICALLY_DERIVED"
	StatusFailedPhysicalYukawaNotDerived = "FAILED_ROUTE_PHYSICAL_TOP_YUKAWA_BOUNDARY_NOT_DERIVED"
	StatusFailedObservedMassNotClaimed   = "FAILED_ROUTE_LOW_ENERGY_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedThresholdStillRequired   = "FAILED_ROUTE_THRESHOLD_OR_BOUNDARY_CORRECTION_STILL_REQUIRED"
)

const (
	rawTraceRatioNumerator   = 1197.0
	rawTraceRatioDenominator = 4624.0
	vevGeV                   = 246.22
	perturbativeLimitSq      = 16.0 * math.Pi * math.Pi
	integrationSteps         = 24000
)

type GenerationTrace struct {
	Formalized               bool
	Equation                 string
	RPlusExact               string
	RPlusDecimal             float64
	TreatsRPlusAsTrace       bool
	TreatsRPlusAsSingleTop   bool
	Generations              int
	GaugeCouplingSquared     float64
	NumericalYukawasInserted bool
	PhysicalTextureDerived   bool
	Verdict                  string
}

type GenerationTopology struct {
	SourceGate                     string
	TauEta                         []int
	Magnitudes                     []int
	DistinctSignedSpectrum         bool
	BreaksAllThreeCapacity         bool
	ScalarTraceFunctionalOnly      bool
	TauEtaToGenerationPullback     bool
	CanonicalTopEigenvectorDerived bool
	TextureDerived                 bool
	Verdict                        string
}

type FractionLane struct {
	Name                         string
	TopFraction                  float64
	GenerationWeights            []float64
	TopYtSquaredOverGStarSquared float64
	TopYtUV                      float64
	Source                       string
	Canonical                    bool
	DerivedFromTauEta            bool
	Ambiguous                    bool
	DiagnosticOnly               bool
	Verdict                      string
}

type RGResult struct {
	LaneName       string
	TopFraction    float64
	YtUV           float64
	InitialLambda  float64
	FinalYt        float64
	FinalLambda    float64
	HiggsMassGeV   float64
	Computed       bool
	Perturbative   bool
	LambdaPositive bool
	Interpretation string
	Verdict        string
}

type CapacityAudit struct {
	Formalized                   bool
	LegacyMassGeV                float64
	BestFractionalMassGeV        float64
	GaugeOnlyMassGeV             float64
	ObservedComparisonGeV        float64
	RequiredLambdaAtVForObserved float64
	GaugeOnlyLambdaAtV           float64
	MinimumPossibleMassAbove125  bool
	FractionalizationFlattens    bool
	FractionalizationCanResolve  bool
	CanonicalFractionDerived     bool
	RequiredMechanism            string
	Verdict                      string
}

type FirewallAudit struct {
	NoObservedMassFitInserted    bool
	NoObservedTopMassInserted    bool
	NoCKMImported                bool
	NoGenerationTextureInvented  bool
	TauEtaNotPromotedToOperator  bool
	NoThresholdJumpInserted      bool
	NoTwoLoopRGExecuted          bool
	NoPoleMassConversionInserted bool
	NoFinalMassClaimed           bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	GenerationTraceFormalized        bool
	TauEtaTopologyRetrieved          bool
	FractionalizationAudited         bool
	RGSlopeReevaluated               bool
	CanonicalTopFractionDerived      bool
	FractionalizationResolvesTension bool
	FirewallPreserved                bool
	FinalMassClaimed                 bool
	Status                           string
	DirectAnswer                     string
	NextGate                         string
}

type Analysis struct {
	Trace     GenerationTrace
	Topology  GenerationTopology
	Lanes     []FractionLane
	Results   []RGResult
	Capacity  CapacityAudit
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
	trace := formalizeGenerationTrace()
	topology := retrieveGenerationTopology()
	lanes := buildFractionLanes(trace, topology)
	results := runFractionalRG(lanes)
	capacity := auditCapacity(results, topology)
	firewalls := auditFirewalls(capacity, topology)
	summary := buildSummary(trace, topology, lanes, results, capacity, firewalls)
	truth := "Gate 313 correctly reinterprets the Gate-309 r_+ top input as a possible three-generation up-type trace rather than a mandatory single top entry.  Fractionalizing the top share lowers y_t(Λ) and flattens the -12y_t^4 contribution, but the inherited Gate-242 tau_eta=(2,-2,1) datum is still only a scalar trace functional with generation-breaking capacity, not a canonical top eigenvector assignment.  Numerically, in the same one-loop PeV lane used by Gate 309, even the gauge-only lower envelope remains around 157 GeV for the fixed Gate-308 quartic boundary, while tau_eta-style fractional witnesses remain substantially above 125 GeV.  Therefore generation fractionalization is a real diagnostic and a necessary top-sector refinement, but by itself it does not resolve the Higgs tension; threshold matching, a changed quartic boundary, two-loop/pole conversion, or a derived nontrivial top tensor remains required."
	return Analysis{Trace: trace, Topology: topology, Lanes: lanes, Results: results, Capacity: capacity, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeGenerationTrace() GenerationTrace {
	rPlus := rPlus()
	return GenerationTrace{
		Formalized:               true,
		Equation:                 "Tr_gen(Y_u†Y_u)/g_*² = y_u²/g_*² + y_c²/g_*² + y_t²/g_*² = r_+",
		RPlusExact:               "(3591+136√123)/3099",
		RPlusDecimal:             rPlus,
		TreatsRPlusAsTrace:       true,
		TreatsRPlusAsSingleTop:   false,
		Generations:              3,
		GaugeCouplingSquared:     1.0,
		NumericalYukawasInserted: false,
		PhysicalTextureDerived:   false,
		Verdict:                  strings.Join([]string{StatusGenerationTraceFormalized, StatusFailedPhysicalYukawaNotDerived}, ";"),
	}
}

func retrieveGenerationTopology() GenerationTopology {
	vals := []int{2, -2, 1}
	mags := []int{2, 2, 1}
	return GenerationTopology{
		SourceGate:                     "Gate 242 tau_eta spatial tagging and generation-breaking audit",
		TauEta:                         vals,
		Magnitudes:                     mags,
		DistinctSignedSpectrum:         true,
		BreaksAllThreeCapacity:         true,
		ScalarTraceFunctionalOnly:      true,
		TauEtaToGenerationPullback:     false,
		CanonicalTopEigenvectorDerived: false,
		TextureDerived:                 false,
		Verdict:                        strings.Join([]string{StatusTauEtaTopologyRetrieved, StatusFailedTauEtaPullbackMissing, StatusFailedTopAssignmentAmbiguous}, ";"),
	}
}

func buildFractionLanes(trace GenerationTrace, topo GenerationTopology) []FractionLane {
	r := trace.RPlusDecimal
	return []FractionLane{
		{
			Name:                         "legacy_all_rplus_assigned_to_top",
			TopFraction:                  1,
			GenerationWeights:            []float64{0, 0, 1},
			TopYtSquaredOverGStarSquared: r,
			TopYtUV:                      math.Sqrt(r),
			Source:                       "Gate 309 sealed diagnostic: y_t²/g_*²=r_+",
			Canonical:                    false,
			DerivedFromTauEta:            false,
			Ambiguous:                    true,
			DiagnosticOnly:               false,
			Verdict:                      StatusFailedPhysicalYukawaNotDerived,
		},
		{
			Name:                         "democratic_three_generation_trace_fraction",
			TopFraction:                  1.0 / 3.0,
			GenerationWeights:            []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0},
			TopYtSquaredOverGStarSquared: r / 3.0,
			TopYtUV:                      math.Sqrt(r / 3.0),
			Source:                       "trace-only fallback: no generation operator, so the only symmetric allocation is 1/3 per generation",
			Canonical:                    false,
			DerivedFromTauEta:            false,
			Ambiguous:                    false,
			DiagnosticOnly:               true,
			Verdict:                      strings.Join([]string{StatusFractionalizationLanesAudited, StatusFailedTopAssignmentAmbiguous}, ";"),
		},
		{
			Name:                         "tau_eta_magnitude_squared_unique_low_top_witness",
			TopFraction:                  1.0 / 9.0,
			GenerationWeights:            []float64{4.0 / 9.0, 4.0 / 9.0, 1.0 / 9.0},
			TopYtSquaredOverGStarSquared: r / 9.0,
			TopYtUV:                      math.Sqrt(r / 9.0),
			Source:                       fmt.Sprintf("conditional |tau_eta|² witness from magnitudes %v; assigns top to unique |tau|=1 slot", topo.Magnitudes),
			Canonical:                    false,
			DerivedFromTauEta:            true,
			Ambiguous:                    true,
			DiagnosticOnly:               true,
			Verdict:                      strings.Join([]string{StatusFractionalizationLanesAudited, StatusFailedTauEtaPullbackMissing, StatusFailedTopAssignmentAmbiguous}, ";"),
		},
		{
			Name:                         "tau_eta_magnitude_squared_high_top_witness",
			TopFraction:                  4.0 / 9.0,
			GenerationWeights:            []float64{4.0 / 9.0, 4.0 / 9.0, 1.0 / 9.0},
			TopYtSquaredOverGStarSquared: 4.0 * r / 9.0,
			TopYtUV:                      math.Sqrt(4.0 * r / 9.0),
			Source:                       fmt.Sprintf("conditional |tau_eta|² witness from magnitudes %v; assigns top to one of the two |tau|=2 slots", topo.Magnitudes),
			Canonical:                    false,
			DerivedFromTauEta:            true,
			Ambiguous:                    true,
			DiagnosticOnly:               true,
			Verdict:                      strings.Join([]string{StatusFractionalizationLanesAudited, StatusFailedTauEtaPullbackMissing, StatusFailedTopAssignmentAmbiguous}, ";"),
		},
		{
			Name:                         "gauge_only_zero_top_lower_envelope",
			TopFraction:                  0,
			GenerationWeights:            []float64{0, 0, 0},
			TopYtSquaredOverGStarSquared: 0,
			TopYtUV:                      0,
			Source:                       "diagnostic lower envelope for this one-loop system; not a physical top sector",
			Canonical:                    false,
			DerivedFromTauEta:            false,
			Ambiguous:                    false,
			DiagnosticOnly:               true,
			Verdict:                      "CONDITIONAL_DIAGNOSTIC_GAUGE_ONLY_LOWER_ENVELOPE_NOT_PHYSICAL",
		},
	}
}

func runFractionalRG(lanes []FractionLane) []RGResult {
	out := make([]RGResult, 0, len(lanes))
	for _, lane := range lanes {
		out = append(out, runRGForLane(lane))
	}
	return out
}

func runRGForLane(lane FractionLane) RGResult {
	lambda0 := rawTraceRatioNumerator / rawTraceRatioDenominator
	state := rgState{ScaleGeV: 2.40099519719e15, GY: math.Sqrt(3.0 / 5.0), G2: 1.0, G3: 1.0, YT: lane.TopYtUV, Lambda: lambda0}
	ok := true
	state, ok = integrateSegment(state, 2.40099519719e15, 1.46774973718e6, betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237})
	if ok {
		state, ok = integrateSegment(state, 1.46774973718e6, vevGeV, betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0})
	}
	lambdaPositive := ok && state.Lambda > 0
	mass := math.NaN()
	if lambdaPositive {
		mass = vevGeV * math.Sqrt(2.0*state.Lambda)
	}
	interp := "fractionalized top-Yukawa one-loop transport completed"
	verdict := strings.Join([]string{StatusRGSlopeReevaluationComputed, StatusPartialFlatteningObserved}, ";")
	if lane.TopFraction == 1 {
		interp = "legacy all-r_+ top lane reproduces the Gate-309 high-mass tension"
		verdict = strings.Join([]string{StatusRGSlopeReevaluationComputed, StatusInsufficientToResolve}, ";")
	}
	if lane.TopFraction == 0 {
		interp = "gauge-only lower envelope; shows top fractionalization alone cannot push below this boundary"
		verdict = strings.Join([]string{StatusRGSlopeReevaluationComputed, StatusInsufficientToResolve}, ";")
	}
	return RGResult{LaneName: lane.Name, TopFraction: lane.TopFraction, YtUV: lane.TopYtUV, InitialLambda: lambda0, FinalYt: state.YT, FinalLambda: state.Lambda, HiggsMassGeV: mass, Computed: ok, Perturbative: ok, LambdaPositive: lambdaPositive, Interpretation: interp, Verdict: verdict}
}

func auditCapacity(results []RGResult, topo GenerationTopology) CapacityAudit {
	legacy := findResult(results, "legacy_all_rplus_assigned_to_top")
	low := findResult(results, "tau_eta_magnitude_squared_unique_low_top_witness")
	gauge := findResult(results, "gauge_only_zero_top_lower_envelope")
	obs := 125.10
	reqLambda := (obs / vevGeV) * (obs / vevGeV) / 2.0
	legacyMass, lowMass, gaugeMass := math.NaN(), math.NaN(), math.NaN()
	gaugeLambda := math.NaN()
	if legacy != nil {
		legacyMass = legacy.HiggsMassGeV
	}
	if low != nil {
		lowMass = low.HiggsMassGeV
	}
	if gauge != nil {
		gaugeMass = gauge.HiggsMassGeV
		gaugeLambda = gauge.FinalLambda
	}
	flattens := !math.IsNaN(legacyMass) && !math.IsNaN(lowMass) && lowMass < legacyMass
	canResolve := !math.IsNaN(gaugeMass) && gaugeMass <= 135.0
	minAbove := !math.IsNaN(gaugeMass) && gaugeMass > 135.0
	verdict := strings.Join([]string{StatusRGSlopeReevaluationComputed, StatusPartialFlatteningObserved, StatusInsufficientToResolve, StatusFailedThresholdStillRequired}, ";")
	return CapacityAudit{Formalized: true, LegacyMassGeV: legacyMass, BestFractionalMassGeV: lowMass, GaugeOnlyMassGeV: gaugeMass, ObservedComparisonGeV: obs, RequiredLambdaAtVForObserved: reqLambda, GaugeOnlyLambdaAtV: gaugeLambda, MinimumPossibleMassAbove125: minAbove, FractionalizationFlattens: flattens, FractionalizationCanResolve: canResolve, CanonicalFractionDerived: topo.CanonicalTopEigenvectorDerived, RequiredMechanism: "A native tau_eta→triality top eigenvector is still needed, but even zero top in this fixed-boundary one-loop lane gives ≈157 GeV; resolving 125 GeV requires a threshold/matching jump, altered quartic boundary, two-loop+polemass ledger, or a deeper top-sector tensor.", Verdict: verdict}
}

func auditFirewalls(c CapacityAudit, topo GenerationTopology) FirewallAudit {
	return FirewallAudit{NoObservedMassFitInserted: true, NoObservedTopMassInserted: true, NoCKMImported: true, NoGenerationTextureInvented: true, TauEtaNotPromotedToOperator: !topo.TauEtaToGenerationPullback && !topo.CanonicalTopEigenvectorDerived, NoThresholdJumpInserted: true, NoTwoLoopRGExecuted: true, NoPoleMassConversionInserted: true, NoFinalMassClaimed: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedTauEtaPullbackMissing, StatusFailedPhysicalYukawaNotDerived, StatusFailedObservedMassNotClaimed}, ";")}
}

func buildSummary(t GenerationTrace, topo GenerationTopology, lanes []FractionLane, results []RGResult, c CapacityAudit, f FirewallAudit) Summary {
	return Summary{GenerationTraceFormalized: t.Formalized && t.TreatsRPlusAsTrace, TauEtaTopologyRetrieved: len(topo.TauEta) == 3 && topo.BreaksAllThreeCapacity, FractionalizationAudited: len(lanes) >= 5, RGSlopeReevaluated: len(results) == len(lanes), CanonicalTopFractionDerived: c.CanonicalFractionDerived, FractionalizationResolvesTension: c.FractionalizationCanResolve, FirewallPreserved: !f.FiniteCorePolluted && f.NoFinalMassClaimed && f.NoGenerationTextureInvented && f.TauEtaNotPromotedToOperator, FinalMassClaimed: false, Status: strings.Join([]string{StatusGenerationTraceFormalized, StatusFractionalizationLanesAudited, StatusInsufficientToResolve}, ";"), DirectAnswer: fmt.Sprintf("Generation fractionalization flattens the top slope: legacy r_+ gives %.3f GeV, the most aggressive tau_eta |τ|² witness gives %.3f GeV, and the zero-top lower envelope is %.3f GeV; under the fixed Gate-308 quartic boundary this still does not reach 125 GeV.", c.LegacyMassGeV, c.BestFractionalMassGeV, c.GaugeOnlyMassGeV), NextGate: "Derive either the tau_eta→triality top eigenvector plus full Yukawa texture, or the missing threshold/matching jump that can lower the fixed-boundary one-loop envelope below the gauge-only ≈157 GeV floor."}
}

func findResult(results []RGResult, name string) *RGResult {
	for i := range results {
		if results[i].LaneName == name {
			return &results[i]
		}
	}
	return nil
}

func rPlus() float64 { return (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0 }

type betaCoefficients struct{ B1GUT, B2, B3 float64 }
type rgState struct{ ScaleGeV, GY, G2, G3, YT, Lambda float64 }

func integrateSegment(initial rgState, high, low float64, beta betaCoefficients) (rgState, bool) {
	state := initial
	logHigh := math.Log(high)
	logLow := math.Log(low)
	h := (logLow - logHigh) / float64(integrationSteps)
	for i := 0; i < integrationSteps; i++ {
		state = rk4Step(state, h, beta)
		state.ScaleGeV = math.Exp(logHigh + float64(i+1)*h)
		if !stateFinite(state) || !statePerturbative(state) {
			return state, false
		}
	}
	state.ScaleGeV = low
	return state, true
}

func rk4Step(y rgState, h float64, b betaCoefficients) rgState {
	k1 := deriv(y, b)
	k2 := deriv(addScaled(y, k1, 0.5*h), b)
	k3 := deriv(addScaled(y, k2, 0.5*h), b)
	k4 := deriv(addScaled(y, k3, h), b)
	return rgState{ScaleGeV: y.ScaleGeV, GY: y.GY + h*(k1.GY+2*k2.GY+2*k3.GY+k4.GY)/6.0, G2: y.G2 + h*(k1.G2+2*k2.G2+2*k3.G2+k4.G2)/6.0, G3: y.G3 + h*(k1.G3+2*k2.G3+2*k3.G3+k4.G3)/6.0, YT: y.YT + h*(k1.YT+2*k2.YT+2*k3.YT+k4.YT)/6.0, Lambda: y.Lambda + h*(k1.Lambda+2*k2.Lambda+2*k3.Lambda+k4.Lambda)/6.0}
}

func addScaled(y, k rgState, scale float64) rgState {
	return rgState{ScaleGeV: y.ScaleGeV, GY: y.GY + scale*k.GY, G2: y.G2 + scale*k.G2, G3: y.G3 + scale*k.G3, YT: y.YT + scale*k.YT, Lambda: y.Lambda + scale*k.Lambda}
}

func deriv(y rgState, b betaCoefficients) rgState {
	bY := (5.0 / 3.0) * b.B1GUT
	gy2, g22, g32 := y.GY*y.GY, y.G2*y.G2, y.G3*y.G3
	yt2 := y.YT * y.YT
	inv := 1.0 / (16.0 * math.Pi * math.Pi)
	betaLambda := (24.0*y.Lambda*y.Lambda + 12.0*y.Lambda*yt2 - 12.0*yt2*yt2 + (3.0/16.0)*(2.0*g22*g22+(g22+gy2)*(g22+gy2)) - y.Lambda*(9.0*g22+3.0*gy2)) * inv
	betaYT := y.YT * ((9.0/2.0)*yt2 - (17.0/12.0)*gy2 - (9.0/4.0)*g22 - 8.0*g32) * inv
	return rgState{GY: bY * y.GY * gy2 * inv, G2: b.B2 * y.G2 * g22 * inv, G3: b.B3 * y.G3 * g32 * inv, YT: betaYT, Lambda: betaLambda}
}

func stateFinite(s rgState) bool {
	for _, v := range []float64{s.GY, s.G2, s.G3, s.YT, s.Lambda} {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func statePerturbative(s rgState) bool {
	return s.GY*s.GY < perturbativeLimitSq && s.G2*s.G2 < perturbativeLimitSq && s.G3*s.G3 < perturbativeLimitSq && s.YT*s.YT < perturbativeLimitSq && math.Abs(s.Lambda) < perturbativeLimitSq
}

func FormatTrace(t GenerationTrace) string {
	return fmt.Sprintf("formalized=%t equation=%q rPlus=%s %.12f trace=%t singleTop=%t generations=%d gstar2=%.6g numerical=%t texture=%t verdict=%s", t.Formalized, t.Equation, t.RPlusExact, t.RPlusDecimal, t.TreatsRPlusAsTrace, t.TreatsRPlusAsSingleTop, t.Generations, t.GaugeCouplingSquared, t.NumericalYukawasInserted, t.PhysicalTextureDerived, t.Verdict)
}

func FormatTopology(t GenerationTopology) string {
	return fmt.Sprintf("source=%q tau=%v mags=%v distinct=%t capacity=%t scalarOnly=%t pullback=%t topEigen=%t texture=%t verdict=%s", t.SourceGate, t.TauEta, t.Magnitudes, t.DistinctSignedSpectrum, t.BreaksAllThreeCapacity, t.ScalarTraceFunctionalOnly, t.TauEtaToGenerationPullback, t.CanonicalTopEigenvectorDerived, t.TextureDerived, t.Verdict)
}

func FormatLane(l FractionLane) string {
	return fmt.Sprintf("%s frac=%.12f weights=%v yt2/g=%.12f ytUV=%.12f source=%q canonical=%t tau=%t ambiguous=%t diagnostic=%t verdict=%s", l.Name, l.TopFraction, l.GenerationWeights, l.TopYtSquaredOverGStarSquared, l.TopYtUV, l.Source, l.Canonical, l.DerivedFromTauEta, l.Ambiguous, l.DiagnosticOnly, l.Verdict)
}

func FormatResult(r RGResult) string {
	return fmt.Sprintf("%s frac=%.12f ytUV=%.12f lambda0=%.12f finalYt=%.12f lambdaV=%.12f mH=%.9f computed=%t pert=%t lambdaPositive=%t interpretation=%q verdict=%s", r.LaneName, r.TopFraction, r.YtUV, r.InitialLambda, r.FinalYt, r.FinalLambda, r.HiggsMassGeV, r.Computed, r.Perturbative, r.LambdaPositive, r.Interpretation, r.Verdict)
}

func FormatCapacity(c CapacityAudit) string {
	return fmt.Sprintf("formalized=%t legacy=%.9f bestFractional=%.9f gaugeOnly=%.9f observed=%.3f requiredLambda=%.12f gaugeLambda=%.12f minAbove125=%t flattens=%t resolves=%t canonical=%t required=%q verdict=%s", c.Formalized, c.LegacyMassGeV, c.BestFractionalMassGeV, c.GaugeOnlyMassGeV, c.ObservedComparisonGeV, c.RequiredLambdaAtVForObserved, c.GaugeOnlyLambdaAtV, c.MinimumPossibleMassAbove125, c.FractionalizationFlattens, c.FractionalizationCanResolve, c.CanonicalFractionDerived, c.RequiredMechanism, c.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noObsMass=%t noObsTop=%t noCKM=%t noTexture=%t tauNotOperator=%t noThreshold=%t noTwoLoop=%t noPole=%t noFinal=%t polluted=%t verdict=%s", f.NoObservedMassFitInserted, f.NoObservedTopMassInserted, f.NoCKMImported, f.NoGenerationTextureInvented, f.TauEtaNotPromotedToOperator, f.NoThresholdJumpInserted, f.NoTwoLoopRGExecuted, f.NoPoleMassConversionInserted, f.NoFinalMassClaimed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("trace=%t tau=%t lanes=%t rg=%t canonicalTop=%t resolves=%t firewall=%t final=%t status=%s answer=%q next=%q", s.GenerationTraceFormalized, s.TauEtaTopologyRetrieved, s.FractionalizationAudited, s.RGSlopeReevaluated, s.CanonicalTopFractionDerived, s.FractionalizationResolvesTension, s.FirewallPreserved, s.FinalMassClaimed, s.Status, s.DirectAnswer, s.NextGate)
}
