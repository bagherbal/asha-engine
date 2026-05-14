// Package conditionalhiggsmassrgtransport implements Gate 309:
// Conditional Higgs Mass from Quartic RG Transport.
//
// Gate 308 produced the UV boundary equation
//
//	lambda_H(Lambda_GUT) = (1197/4624) g_*^2.
//
// Gate 309 accepts the explicitly sealed diagnostic branch g_*^2=1 and transports
// the quartic through a one-loop continuum RG system. This is a conditional
// phenomenological transport theorem, not a new finite-core derivation: the
// boundary scale, PeV thresholds, top-Yukawa branch, one-loop truncation,
// decoupling rule, and tree-level mass extraction are all recorded as seals.
// The gate intentionally computes numbers because the boundary equation now has
// enough conditional data to be stress-tested, but it refuses to call them final
// collider predictions.
package conditionalhiggsmassrgtransport

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE309-CONDITIONAL-HIGGS-MASS-QUARTIC-RG-TRANSPORT"

	StatusGate308BoundaryInherited       = "CONDITIONAL_SUPPORT_GATE308_QUARTIC_BOUNDARY_INHERITED"
	StatusTopologicalGStarSealActivated  = "CONDITIONAL_SUPPORT_TOPOLOGICAL_GSTAR_SQUARED_EQUALS_ONE_SEAL_ACTIVATED"
	StatusOneLoopRGSystemFormalized      = "CONDITIONAL_SUPPORT_ONE_LOOP_SM_QUARTIC_RG_SYSTEM_FORMALIZED"
	StatusPeVThresholdLaneEvaluated      = "CONDITIONAL_SUPPORT_PEV_THRESHOLD_RG_LANE_EVALUATED"
	StatusConditionalHiggsMassComputed   = "CONDITIONAL_SUPPORT_CONDITIONAL_HIGGS_MASS_TRANSPORT_COMPUTED"
	StatusPureSMHighScalePathology       = "FAILED_ROUTE_PURE_SM_GSTAR_ONE_HIGH_SCALE_RUN_HITS_QCD_NONPERTURBATIVE_BARRIER"
	StatusRPlusTopLaneTension            = "CONDITIONAL_TENSION_RPLUS_TOP_SEAL_DRIVES_HIGGS_MASS_HIGH"
	StatusGaugeOnlyDiagnosticNotPhysical = "CONDITIONAL_DIAGNOSTIC_GAUGE_ONLY_TRANSPORT_NOT_PHYSICAL_TOP_SECTOR"
	StatusGate309FirewallsPreserved      = "CONDITIONAL_SUPPORT_GATE309_RG_TRANSPORT_FIREWALLS_PRESERVED"

	StatusFailedTwoLoopNotIncluded               = "FAILED_ROUTE_TWO_LOOP_RGE_NOT_INCLUDED"
	StatusFailedMatchingCorrectionsNotIncluded   = "FAILED_ROUTE_THRESHOLD_MATCHING_CORRECTIONS_NOT_INCLUDED"
	StatusFailedFiniteThresholdOriginStillSealed = "FAILED_ROUTE_FINITE_THRESHOLD_ORIGIN_STILL_SEALED"
	StatusFailedTopYukawaOriginStillSealed       = "FAILED_ROUTE_TOP_YUKAWA_ORIGIN_STILL_SEALED"
	StatusFailedPoleMassCorrectionsNotIncluded   = "FAILED_ROUTE_POLE_MASS_AND_MS_BAR_MATCHING_NOT_INCLUDED"
	StatusFailedLowEnergyFinalMassNotClaimed     = "FAILED_ROUTE_FINAL_LOW_ENERGY_HIGGS_MASS_NOT_CLAIMED"
)

const (
	rawTraceRatioNumerator   = 1197.0
	rawTraceRatioDenominator = 4624.0
	vevGeV                   = 246.22
	pureClosedTriangleGeV    = 1.0e17
	perturbativeLimitSq      = 16.0 * math.Pi * math.Pi
	integrationSteps         = 24000
)

type Gate308Boundary struct {
	Equation           string
	CoefficientExact   string
	CoefficientDecimal float64
	GStarSquaredSealed float64
	LambdaUV           float64
	TreeLevelMassGeV   float64
	BoundaryInherited  bool
	AbsoluteFinalClaim bool
	Verdict            string
}

type BetaSystem struct {
	Convention                 string
	GaugeBeta                  string
	TopYukawaBeta              string
	QuarticBeta                string
	HyperchargeNormalization   string
	OneLoopOnly                bool
	UsesStandardContinuumQFT   bool
	DerivedAsFiniteCoreTheorem bool
	Verdict                    string
}

type ThresholdLane struct {
	Name                       string
	Carrier                    string
	MStarGeV                   float64
	ThresholdGeV               float64
	DeltaB1GUT                 float64
	DeltaB2                    float64
	DeltaB3                    float64
	ConditionalOnPeVSeal       bool
	FiniteThresholdOrigin      bool
	MatchingCorrectionsDerived bool
	Verdict                    string
}

type TopYukawaLane struct {
	Name                 string
	YtUV                 float64
	YtSquaredOverGStarSq float64
	Source               string
	PhysicalTopSector    bool
	FiniteCoreDerived    bool
	DiagnosticOnly       bool
	Verdict              string
}

type RGState struct {
	ScaleGeV float64
	GY       float64
	G2       float64
	G3       float64
	YT       float64
	Lambda   float64
}

type TransportResult struct {
	Name                    string
	ThresholdName           string
	TopLaneName             string
	StartScaleGeV           float64
	ThresholdGeV            float64
	EndScaleGeV             float64
	InitialLambda           float64
	InitialYt               float64
	FinalGY2                float64
	FinalG22                float64
	FinalG32                float64
	FinalYt                 float64
	FinalLambda             float64
	HiggsMassGeV            float64
	Computed                bool
	Perturbative            bool
	LambdaPositive          bool
	NonPerturbativeScaleGeV float64
	FailureReason           string
	Interpretation          string
	Verdict                 string
}

type PredictionAudit struct {
	PrimaryLaneName            string
	PrimaryConditionalMassGeV  float64
	PrimaryLambdaAtV           float64
	GaugeOnlyDiagnosticMassGeV float64
	PureSMRunInvalid           bool
	RPlusTopLaneComputed       bool
	RPlusTopLaneNearObserved   bool
	ObservedHiggsMassInserted  bool
	MeasuredComparisonOnly     bool
	TruthOrTension             string
	Verdict                    string
}

type FirewallAudit struct {
	NoObservedHiggsMassUsedForDerivation bool
	NoObservedTopMassUsedForDerivation   bool
	NoTwoLoopTermsInserted               bool
	NoThresholdMatchingInserted          bool
	NoPoleMassMatchingInserted           bool
	PeVThresholdsRemainSealed            bool
	TopYukawaOriginRemainsSealed         bool
	PureSMPathologyRecorded              bool
	FinalColliderPredictionClaimed       bool
	FiniteCorePolluted                   bool
	Obligations                          []RemainingObligation
	Verdict                              string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksFinalPrediction     bool
}

type Summary struct {
	Gate308Inherited               bool
	GStarSealActivated             bool
	RGSystemFormalized             bool
	ConditionalTransportRun        bool
	PrimaryMassComputed            bool
	PureSMPathologyFound           bool
	FinalMassClaimed               bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Boundary   Gate308Boundary
	Beta       BetaSystem
	Thresholds []ThresholdLane
	TopLanes   []TopYukawaLane
	Results    []TransportResult
	Prediction PredictionAudit
	Firewalls  FirewallAudit
	Summary    Summary
	Truth      string
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
	boundary := inheritGate308Boundary()
	beta := formalizeBetaSystem()
	thresholds := buildThresholdLanes()
	topLanes := buildTopYukawaLanes()
	results := runTransportLanes(boundary, thresholds, topLanes)
	prediction := auditPrediction(results)
	firewalls := auditFirewalls(prediction)
	summary := buildSummary(boundary, beta, prediction, firewalls)
	truth := "Gate 309 runs the first conditional GeV-scale Higgs diagnostic from the Gate-308 quartic boundary. With g_*^2=1, lambda_H(Lambda)=1197/4624 and the tree-level no-running mass is about 177.14 GeV. A pure SM one-loop descent from 10^17 GeV is invalid because g3 becomes nonperturbative before the electroweak scale. When the quarantined PeV threshold lanes are activated, the one-loop transport is evaluable. The gauge-only diagnostic gives about 157 GeV, while the r_+ top-Yukawa boundary lane gives about 327-332 GeV depending on the sealed threshold carrier. Therefore the first conditional number is not near 125 GeV under this one-loop/sealed-threshold/r_+ protocol; it is a tension signal, not a final falsification, because top-Yukawa origin, threshold matching, two-loop running, scheme conversion, and pole-mass extraction remain firewalled."
	return Analysis{Boundary: boundary, Beta: beta, Thresholds: thresholds, TopLanes: topLanes, Results: results, Prediction: prediction, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritGate308Boundary() Gate308Boundary {
	lambda := rawTraceRatioNumerator / rawTraceRatioDenominator
	return Gate308Boundary{
		Equation:           "λ_H(Λ_GUT) = (1197/4624) · g_*²",
		CoefficientExact:   "1197/4624",
		CoefficientDecimal: lambda,
		GStarSquaredSealed: 1.0,
		LambdaUV:           lambda,
		TreeLevelMassGeV:   vevGeV * math.Sqrt(2.0*lambda),
		BoundaryInherited:  true,
		AbsoluteFinalClaim: false,
		Verdict:            strings.Join([]string{StatusGate308BoundaryInherited, StatusTopologicalGStarSealActivated}, ";"),
	}
}

func formalizeBetaSystem() BetaSystem {
	return BetaSystem{
		Convention:                 "t=ln μ; dg_i/dt=b_i g_i^3/(16π²); dλ/dt=β_λ; downward transport integrates from high μ to v",
		GaugeBeta:                  "SM below thresholds: b1_GUT=41/10, bY=(5/3)b1_GUT=41/6, b2=-19/6, b3=-7; sealed PeV lanes use b_i -> b_i+Δb_i above threshold",
		TopYukawaBeta:              "dy_t/dt = y_t/(16π²)[(9/2)y_t² - (17/12)gY² - (9/4)g2² - 8g3²]",
		QuarticBeta:                "dλ/dt=(1/16π²)[24λ²+12λy_t²-12y_t⁴+(3/16)(2g2⁴+(g2²+gY²)²)-λ(9g2²+3gY²)]",
		HyperchargeNormalization:   "Gate 308 uses g_*²=(5/3)gY² at the UV boundary, so gY²(Λ)=3/5 when g_*²=1",
		OneLoopOnly:                true,
		UsesStandardContinuumQFT:   true,
		DerivedAsFiniteCoreTheorem: false,
		Verdict:                    strings.Join([]string{StatusOneLoopRGSystemFormalized, StatusFailedTwoLoopNotIncluded}, ";"),
	}
}

func buildThresholdLanes() []ThresholdLane {
	return []ThresholdLane{
		{
			Name:                       "pure_SM_closed_triangle_1e17_control",
			Carrier:                    "none",
			MStarGeV:                   pureClosedTriangleGeV,
			ThresholdGeV:               pureClosedTriangleGeV,
			ConditionalOnPeVSeal:       false,
			FiniteThresholdOrigin:      false,
			MatchingCorrectionsDerived: false,
			Verdict:                    StatusPureSMHighScalePathology,
		},
		{
			Name:                       "Gate206_Dirac_vectorlike_quark_doublet_PeV_lane",
			Carrier:                    "Dirac vectorlike quark doublet (3,2,1/6)",
			MStarGeV:                   2.40099519719e15,
			ThresholdGeV:               1.46774973718e6,
			DeltaB1GUT:                 7.78628724237,
			DeltaB2:                    9.65295390904,
			DeltaB3:                    8.98628724237,
			ConditionalOnPeVSeal:       true,
			FiniteThresholdOrigin:      false,
			MatchingCorrectionsDerived: false,
			Verdict:                    strings.Join([]string{StatusPeVThresholdLaneEvaluated, StatusFailedFiniteThresholdOriginStillSealed, StatusFailedMatchingCorrectionsNotIncluded}, ";"),
		},
		{
			Name:                       "Gate206_Weyl_SU2_adjoint_fermion_PeV_lane",
			Carrier:                    "Weyl SU(2)L adjoint fermion (1,3,0)",
			MStarGeV:                   2.42276543552e14,
			ThresholdGeV:               8.19807624157e6,
			DeltaB1GUT:                 10.1497542656,
			DeltaB2:                    11.4830875989,
			DeltaB3:                    10.1497542656,
			ConditionalOnPeVSeal:       true,
			FiniteThresholdOrigin:      false,
			MatchingCorrectionsDerived: false,
			Verdict:                    strings.Join([]string{StatusPeVThresholdLaneEvaluated, StatusFailedFiniteThresholdOriginStillSealed, StatusFailedMatchingCorrectionsNotIncluded}, ";"),
		},
	}
}

func buildTopYukawaLanes() []TopYukawaLane {
	rPlus := (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0
	return []TopYukawaLane{
		{
			Name:                 "gauge_only_y_t_zero_diagnostic",
			YtUV:                 0,
			YtSquaredOverGStarSq: 0,
			Source:               "diagnostic lane: removes the top sector to show pure gauge/quartic transport",
			PhysicalTopSector:    false,
			FiniteCoreDerived:    false,
			DiagnosticOnly:       true,
			Verdict:              StatusGaugeOnlyDiagnosticNotPhysical,
		},
		{
			Name:                 "r_plus_top_yukawa_boundary_seal",
			YtUV:                 math.Sqrt(rPlus),
			YtSquaredOverGStarSq: rPlus,
			Source:               "sealed r_+ branch diagnostic: y_t²/g_*² = r_+ = (3591+136√123)/3099",
			PhysicalTopSector:    true,
			FiniteCoreDerived:    false,
			DiagnosticOnly:       false,
			Verdict:              strings.Join([]string{StatusRPlusTopLaneTension, StatusFailedTopYukawaOriginStillSealed}, ";"),
		},
	}
}

func runTransportLanes(b Gate308Boundary, thresholds []ThresholdLane, topLanes []TopYukawaLane) []TransportResult {
	out := make([]TransportResult, 0, len(thresholds)*len(topLanes))
	for _, th := range thresholds {
		for _, top := range topLanes {
			out = append(out, runSingleLane(b, th, top))
		}
	}
	return out
}

func runSingleLane(boundary Gate308Boundary, th ThresholdLane, top TopYukawaLane) TransportResult {
	state := RGState{ScaleGeV: th.MStarGeV, GY: math.Sqrt(3.0 / 5.0), G2: 1.0, G3: 1.0, YT: top.YtUV, Lambda: boundary.LambdaUV}
	initial := state
	ok := true
	failScale := 0.0
	failure := ""
	if th.ThresholdGeV < th.MStarGeV {
		state, ok, failScale, failure = integrateSegment(state, th.MStarGeV, th.ThresholdGeV, betaCoefficients{B1GUT: 41.0/10.0 + th.DeltaB1GUT, B2: -19.0/6.0 + th.DeltaB2, B3: -7.0 + th.DeltaB3})
		if !ok {
			return buildTransportResult(th, top, initial, state, false, failScale, "above-threshold segment: "+failure)
		}
	}
	state, ok, failScale, failure = integrateSegment(state, th.ThresholdGeV, vevGeV, betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0})
	if !ok {
		return buildTransportResult(th, top, initial, state, false, failScale, "below-threshold SM segment: "+failure)
	}
	return buildTransportResult(th, top, initial, state, true, 0, "")
}

type betaCoefficients struct{ B1GUT, B2, B3 float64 }

func integrateSegment(initial RGState, high, low float64, beta betaCoefficients) (RGState, bool, float64, string) {
	state := initial
	if high == low {
		return state, true, 0, ""
	}
	logHigh := math.Log(high)
	logLow := math.Log(low)
	h := (logLow - logHigh) / float64(integrationSteps)
	for i := 0; i < integrationSteps; i++ {
		state = rk4Step(state, h, beta)
		state.ScaleGeV = math.Exp(logHigh + float64(i+1)*h)
		if !stateFinite(state) {
			return state, false, state.ScaleGeV, "non-finite RG state"
		}
		if !statePerturbative(state) {
			return state, false, state.ScaleGeV, "perturbative-control threshold exceeded"
		}
	}
	state.ScaleGeV = low
	return state, true, 0, ""
}

func rk4Step(y RGState, h float64, b betaCoefficients) RGState {
	k1 := deriv(y, b)
	k2 := deriv(addScaled(y, k1, 0.5*h), b)
	k3 := deriv(addScaled(y, k2, 0.5*h), b)
	k4 := deriv(addScaled(y, k3, h), b)
	return RGState{
		ScaleGeV: y.ScaleGeV,
		GY:       y.GY + h*(k1.GY+2*k2.GY+2*k3.GY+k4.GY)/6.0,
		G2:       y.G2 + h*(k1.G2+2*k2.G2+2*k3.G2+k4.G2)/6.0,
		G3:       y.G3 + h*(k1.G3+2*k2.G3+2*k3.G3+k4.G3)/6.0,
		YT:       y.YT + h*(k1.YT+2*k2.YT+2*k3.YT+k4.YT)/6.0,
		Lambda:   y.Lambda + h*(k1.Lambda+2*k2.Lambda+2*k3.Lambda+k4.Lambda)/6.0,
	}
}

func addScaled(y, k RGState, scale float64) RGState {
	return RGState{ScaleGeV: y.ScaleGeV, GY: y.GY + scale*k.GY, G2: y.G2 + scale*k.G2, G3: y.G3 + scale*k.G3, YT: y.YT + scale*k.YT, Lambda: y.Lambda + scale*k.Lambda}
}

func deriv(y RGState, b betaCoefficients) RGState {
	bY := (5.0 / 3.0) * b.B1GUT
	gy2, g22, g32 := y.GY*y.GY, y.G2*y.G2, y.G3*y.G3
	yt2 := y.YT * y.YT
	inv := 1.0 / (16.0 * math.Pi * math.Pi)
	betaLambda := (24.0*y.Lambda*y.Lambda + 12.0*y.Lambda*yt2 - 12.0*yt2*yt2 + (3.0/16.0)*(2.0*g22*g22+(g22+gy2)*(g22+gy2)) - y.Lambda*(9.0*g22+3.0*gy2)) * inv
	betaYT := y.YT * ((9.0/2.0)*yt2 - (17.0/12.0)*gy2 - (9.0/4.0)*g22 - 8.0*g32) * inv
	return RGState{GY: bY * y.GY * gy2 * inv, G2: b.B2 * y.G2 * g22 * inv, G3: b.B3 * y.G3 * g32 * inv, YT: betaYT, Lambda: betaLambda}
}

func stateFinite(s RGState) bool {
	vals := []float64{s.GY, s.G2, s.G3, s.YT, s.Lambda}
	for _, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func statePerturbative(s RGState) bool {
	return s.GY*s.GY < perturbativeLimitSq && s.G2*s.G2 < perturbativeLimitSq && s.G3*s.G3 < perturbativeLimitSq && s.YT*s.YT < perturbativeLimitSq && math.Abs(s.Lambda) < perturbativeLimitSq
}

func buildTransportResult(th ThresholdLane, top TopYukawaLane, initial, final RGState, computed bool, failScale float64, failure string) TransportResult {
	lambdaPositive := computed && final.Lambda > 0
	mass := math.NaN()
	if lambdaPositive {
		mass = vevGeV * math.Sqrt(2.0*final.Lambda)
	}
	verdict := StatusConditionalHiggsMassComputed
	interp := "conditional one-loop transport completed"
	if !computed {
		verdict = StatusPureSMHighScalePathology
		interp = "lane invalid before electroweak extraction"
	} else if top.DiagnosticOnly {
		verdict = strings.Join([]string{StatusConditionalHiggsMassComputed, StatusGaugeOnlyDiagnosticNotPhysical}, ";")
		interp = "gauge-only diagnostic; not a physical top-sector prediction"
	} else {
		verdict = strings.Join([]string{StatusConditionalHiggsMassComputed, StatusRPlusTopLaneTension}, ";")
		interp = "r_+ top-seal transport; conditional tension if compared with the measured Higgs mass"
	}
	return TransportResult{
		Name:                    th.Name + "/" + top.Name,
		ThresholdName:           th.Name,
		TopLaneName:             top.Name,
		StartScaleGeV:           th.MStarGeV,
		ThresholdGeV:            th.ThresholdGeV,
		EndScaleGeV:             vevGeV,
		InitialLambda:           initial.Lambda,
		InitialYt:               initial.YT,
		FinalGY2:                final.GY * final.GY,
		FinalG22:                final.G2 * final.G2,
		FinalG32:                final.G3 * final.G3,
		FinalYt:                 final.YT,
		FinalLambda:             final.Lambda,
		HiggsMassGeV:            mass,
		Computed:                computed,
		Perturbative:            computed,
		LambdaPositive:          lambdaPositive,
		NonPerturbativeScaleGeV: failScale,
		FailureReason:           failure,
		Interpretation:          interp,
		Verdict:                 verdict,
	}
}

func auditPrediction(results []TransportResult) PredictionAudit {
	var primary *TransportResult
	var gaugeOnly *TransportResult
	pureInvalid := false
	for i := range results {
		r := &results[i]
		if strings.Contains(r.ThresholdName, "pure_SM") && !r.Computed {
			pureInvalid = true
		}
		if strings.Contains(r.ThresholdName, "Dirac_vectorlike") && r.TopLaneName == "r_plus_top_yukawa_boundary_seal" {
			primary = r
		}
		if strings.Contains(r.ThresholdName, "Dirac_vectorlike") && r.TopLaneName == "gauge_only_y_t_zero_diagnostic" {
			gaugeOnly = r
		}
	}
	mass, lam := math.NaN(), math.NaN()
	primaryName := "missing"
	rplusComputed := false
	nearObserved := false
	if primary != nil {
		mass = primary.HiggsMassGeV
		lam = primary.FinalLambda
		primaryName = primary.Name
		rplusComputed = primary.Computed
		nearObserved = primary.Computed && mass > 115.0 && mass < 135.0
	}
	gaugeMass := math.NaN()
	if gaugeOnly != nil {
		gaugeMass = gaugeOnly.HiggsMassGeV
	}
	truth := "conditional result not close to 125 GeV in the r_+ top lane"
	if nearObserved {
		truth = "conditional result lands in the rough observed-Higgs window"
	}
	return PredictionAudit{PrimaryLaneName: primaryName, PrimaryConditionalMassGeV: mass, PrimaryLambdaAtV: lam, GaugeOnlyDiagnosticMassGeV: gaugeMass, PureSMRunInvalid: pureInvalid, RPlusTopLaneComputed: rplusComputed, RPlusTopLaneNearObserved: nearObserved, ObservedHiggsMassInserted: false, MeasuredComparisonOnly: true, TruthOrTension: truth, Verdict: strings.Join([]string{StatusConditionalHiggsMassComputed, StatusPureSMHighScalePathology, StatusRPlusTopLaneTension}, ";")}
}

func auditFirewalls(p PredictionAudit) FirewallAudit {
	obs := []RemainingObligation{
		{"two-loop RGEs", "one-loop transport is only the first diagnostic; two-loop SM/BSM beta functions can materially shift lambda", StatusFailedTwoLoopNotIncluded, true},
		{"threshold matching", "PeV carriers were decoupled by a step rule without finite matching coefficients", StatusFailedMatchingCorrectionsNotIncluded, true},
		{"finite threshold origin", "the PeV spectrum is inherited as a conditional seal, not a finite-core theorem", StatusFailedFiniteThresholdOriginStillSealed, true},
		{"top Yukawa origin", "the r_+ top boundary is a sealed amplitude branch, not a derived full Yukawa texture", StatusFailedTopYukawaOriginStillSealed, true},
		{"pole/MSbar conversion", "m_H=v sqrt(2 lambda(v)) is tree-level and not a pole-mass extraction", StatusFailedPoleMassCorrectionsNotIncluded, true},
	}
	return FirewallAudit{NoObservedHiggsMassUsedForDerivation: !p.ObservedHiggsMassInserted, NoObservedTopMassUsedForDerivation: true, NoTwoLoopTermsInserted: true, NoThresholdMatchingInserted: true, NoPoleMassMatchingInserted: true, PeVThresholdsRemainSealed: true, TopYukawaOriginRemainsSealed: true, PureSMPathologyRecorded: p.PureSMRunInvalid, FinalColliderPredictionClaimed: false, FiniteCorePolluted: false, Obligations: obs, Verdict: strings.Join([]string{StatusGate309FirewallsPreserved, StatusFailedLowEnergyFinalMassNotClaimed, StatusFailedTwoLoopNotIncluded, StatusFailedMatchingCorrectionsNotIncluded}, ";")}
}

func buildSummary(b Gate308Boundary, beta BetaSystem, p PredictionAudit, f FirewallAudit) Summary {
	return Summary{Gate308Inherited: b.BoundaryInherited, GStarSealActivated: b.GStarSquaredSealed == 1.0, RGSystemFormalized: beta.OneLoopOnly && beta.UsesStandardContinuumQFT, ConditionalTransportRun: p.RPlusTopLaneComputed, PrimaryMassComputed: !math.IsNaN(p.PrimaryConditionalMassGeV), PureSMPathologyFound: p.PureSMRunInvalid, FinalMassClaimed: f.FinalColliderPredictionClaimed, FirewallPreserved: !f.FiniteCorePolluted && !f.FinalColliderPredictionClaimed && f.NoObservedHiggsMassUsedForDerivation && f.NoObservedTopMassUsedForDerivation && f.NoTwoLoopTermsInserted && f.NoThresholdMatchingInserted && f.PureSMPathologyRecorded, Status: StatusConditionalHiggsMassComputed, DirectAnswer: fmt.Sprintf("Conditional r_+ one-loop PeV-threshold transport gives m_H≈%.2f GeV in the primary Dirac-vectorlike lane; gauge-only diagnostic gives ≈%.2f GeV; pure SM 10^17 GeV lane is nonperturbative before v.", p.PrimaryConditionalMassGeV, p.GaugeOnlyDiagnosticMassGeV), NextGate: "Gate 310 should add the two-loop plus matching ledger, including explicit threshold matching and MSbar/pole conversion, before any final collider-scale Higgs-mass claim."}
}

func FormatBoundary(b Gate308Boundary) string {
	return fmt.Sprintf("equation=%q coeff=%s decimal=%.12f gstar2=%.6g lambdaUV=%.12f treeMass=%.6f inherited=%t finalClaim=%t verdict=%s", b.Equation, b.CoefficientExact, b.CoefficientDecimal, b.GStarSquaredSealed, b.LambdaUV, b.TreeLevelMassGeV, b.BoundaryInherited, b.AbsoluteFinalClaim, b.Verdict)
}

func FormatBeta(b BetaSystem) string {
	return fmt.Sprintf("convention=%q gauge=%q yt=%q lambda=%q hyper=%q oneLoop=%t standard=%t finiteCore=%t verdict=%s", b.Convention, b.GaugeBeta, b.TopYukawaBeta, b.QuarticBeta, b.HyperchargeNormalization, b.OneLoopOnly, b.UsesStandardContinuumQFT, b.DerivedAsFiniteCoreTheorem, b.Verdict)
}

func FormatThreshold(t ThresholdLane) string {
	return fmt.Sprintf("%s carrier=%q M*=%.9gGeV MB=%.9gGeV Δb=(%.9g,%.9g,%.9g) seal=%t finite=%t matching=%t verdict=%s", t.Name, t.Carrier, t.MStarGeV, t.ThresholdGeV, t.DeltaB1GUT, t.DeltaB2, t.DeltaB3, t.ConditionalOnPeVSeal, t.FiniteThresholdOrigin, t.MatchingCorrectionsDerived, t.Verdict)
}

func FormatTopLane(t TopYukawaLane) string {
	return fmt.Sprintf("%s ytUV=%.12f yt2/gstar2=%.12f source=%q physicalTop=%t finite=%t diagnostic=%t verdict=%s", t.Name, t.YtUV, t.YtSquaredOverGStarSq, t.Source, t.PhysicalTopSector, t.FiniteCoreDerived, t.DiagnosticOnly, t.Verdict)
}

func FormatResult(r TransportResult) string {
	return fmt.Sprintf("%s start=%.9g threshold=%.9g end=%.9g lambda0=%.12f yt0=%.12f final=(gY2 %.9g,g2² %.9g,g3² %.9g,yt %.9g,lambda %.12g,mH %.9g) computed=%t pert=%t lambdaPositive=%t failScale=%.9g reason=%q interpretation=%q verdict=%s", r.Name, r.StartScaleGeV, r.ThresholdGeV, r.EndScaleGeV, r.InitialLambda, r.InitialYt, r.FinalGY2, r.FinalG22, r.FinalG32, r.FinalYt, r.FinalLambda, r.HiggsMassGeV, r.Computed, r.Perturbative, r.LambdaPositive, r.NonPerturbativeScaleGeV, r.FailureReason, r.Interpretation, r.Verdict)
}

func FormatPrediction(p PredictionAudit) string {
	return fmt.Sprintf("primary=%q mH=%.9g lambdaV=%.12g gaugeOnly=%.9g pureInvalid=%t rplusComputed=%t nearObserved=%t observedInserted=%t comparisonOnly=%t truth=%q verdict=%s", p.PrimaryLaneName, p.PrimaryConditionalMassGeV, p.PrimaryLambdaAtV, p.GaugeOnlyDiagnosticMassGeV, p.PureSMRunInvalid, p.RPlusTopLaneComputed, p.RPlusTopLaneNearObserved, p.ObservedHiggsMassInserted, p.MeasuredComparisonOnly, p.TruthOrTension, p.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksFinalPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noObsH=%t noObsTop=%t noTwoLoop=%t noMatching=%t noPole=%t thresholdsSealed=%t topSealed=%t purePathology=%t finalClaim=%t polluted=%t obligations=[%s] verdict=%s", f.NoObservedHiggsMassUsedForDerivation, f.NoObservedTopMassUsedForDerivation, f.NoTwoLoopTermsInserted, f.NoThresholdMatchingInserted, f.NoPoleMassMatchingInserted, f.PeVThresholdsRemainSealed, f.TopYukawaOriginRemainsSealed, f.PureSMPathologyRecorded, f.FinalColliderPredictionClaimed, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate308=%t gstar=%t rg=%t transport=%t mass=%t purePathology=%t finalClaim=%t firewall=%t status=%s answer=%q next=%q", s.Gate308Inherited, s.GStarSealActivated, s.RGSystemFormalized, s.ConditionalTransportRun, s.PrimaryMassComputed, s.PureSMPathologyFound, s.FinalMassClaimed, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
