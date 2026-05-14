// Package intermediatethresholdjump implements Gate 314:
// Intermediate Threshold Decoupling / Quartic Jump Transport Audit.
//
// Gate 313 proved that continuous one-loop RG transport from the Gate-308
// quartic boundary has a hard gauge-only floor near 157 GeV under the PeV lane.
// Gate 314 therefore audits the missing discontinuous ingredient: a finite
// matching jump in the quartic coupling at an intermediate threshold scale.
// The package does not derive the heavy sector.  It computes the exact target
// Δλ required at the threshold for several inherited lanes, checks the sign and
// order of magnitude, and records the portal-form obligation
//
//	Δλ = -λ_mix²/(4 λ_heavy)
//
// as the next theorem to derive from finite geometry.
package intermediatethresholdjump

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE314-INTERMEDIATE-THRESHOLD-DECOUPLING-QUARTIC-JUMP-TRANSPORT-AUDIT"

	StatusTwoStageRGFormalized            = "CONDITIONAL_SUPPORT_TWO_STAGE_RG_TRANSPORT_FORMALIZED"
	StatusThresholdJumpInserted           = "CONDITIONAL_SUPPORT_THRESHOLD_JUMP_INSERTION_FORMALIZED"
	StatusRequiredJumpExtracted           = "CONDITIONAL_SUPPORT_INTERMEDIATE_THRESHOLD_JUMP_EXTRACTED"
	StatusViabilitySieveFormalized        = "CONDITIONAL_SUPPORT_THRESHOLD_JUMP_VIABILITY_SIEVE_FORMALIZED"
	StatusPortalMagnitudeTargetFormalized = "CONDITIONAL_SUPPORT_PORTAL_MAGNITUDE_TARGET_FORMALIZED"
	StatusThresholdMechanismRequired      = "CONDITIONAL_TENSION_DISCONTINUOUS_THRESHOLD_MECHANISM_REQUIRED"
	StatusGate314FirewallsPreserved       = "CONDITIONAL_SUPPORT_GATE314_FIREWALLS_PRESERVED"
	StatusFailedJumpNotDerived            = "FAILED_ROUTE_THRESHOLD_JUMP_VALUE_NOT_DERIVED_FROM_FINITE_GEOMETRY"
	StatusFailedHeavyPortalNotDerived     = "FAILED_ROUTE_HEAVY_PORTAL_COUPLING_NOT_DERIVED"
	StatusFailedHeavyQuarticNotDerived    = "FAILED_ROUTE_HEAVY_SELF_QUARTIC_NOT_DERIVED"
	StatusFailedThresholdScaleStillSealed = "FAILED_ROUTE_THRESHOLD_SCALE_STILL_CONDITIONAL"
	StatusFailedTopFractionStillSealed    = "FAILED_ROUTE_TOP_YUKAWA_FRACTION_STILL_CONDITIONAL"
	StatusFailedTwoLoopNotExecuted        = "FAILED_ROUTE_TWO_LOOP_RGE_NOT_EXECUTED"
	StatusFailedPoleMassNotComputed       = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_COMPUTED"
	StatusFailedFinalMassNotClaimed       = "FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedHeavySectorNotConstructed = "FAILED_ROUTE_HEAVY_INTERMEDIATE_SECTOR_NOT_CONSTRUCTED"
)

const (
	lambdaUVBoundary    = 1197.0 / 4624.0
	vevGeV              = 246.22
	targetPoleLikeGeV   = 125.10
	gutScaleGeV         = 2.40099519719e15
	thresholdScaleGeV   = 1.46774973718e6
	integrationSteps    = 24000
	perturbativeLimitSq = 16.0 * math.Pi * math.Pi
	rootTolerance       = 1e-12
	maxRootIterations   = 96
)

type TwoStageRG struct {
	Formalized       bool
	HighScaleGeV     float64
	ThresholdGeV     float64
	LowScaleGeV      float64
	SegmentA         string
	SegmentB         string
	HighBetaLedger   string
	LowBetaLedger    string
	TargetMassGeV    float64
	TargetLambdaAtV  float64
	BoundaryLambdaUV float64
	Verdict          string
}

type ThresholdRule struct {
	Formalized               bool
	Condition                string
	JumpSymbol               string
	AppliedAtGeV             float64
	SignConvention           string
	NegativeJumpLowersIRMass bool
	DerivedFromHeavySector   bool
	Verdict                  string
}

type TransportLane struct {
	Name           string
	SourceGate     string
	TopFraction    float64
	YtUV           float64
	Canonical      bool
	DiagnosticOnly bool
	Description    string
}

type JumpExtraction struct {
	LaneName               string
	TopFraction            float64
	YtUV                   float64
	LambdaAtThresholdPlus  float64
	BaselineLambdaAtV      float64
	BaselineMassGeV        float64
	RequiredDeltaLambda    float64
	LambdaAtThresholdMinus float64
	TargetLambdaAtV        float64
	TargetMassGeV          float64
	Solved                 bool
	Perturbative           bool
	CorrectSign            bool
	Magnitude              float64
	ViableOrder            bool
	Interpretation         string
	Verdict                string
}

type ViabilityAudit struct {
	Formalized                    bool
	PreferredLaneName             string
	PreferredRequiredDelta        float64
	PreferredThresholdLambdaPlus  float64
	PreferredThresholdLambdaMinus float64
	PreferredBaselineMassGeV      float64
	TargetMassGeV                 float64
	TargetLambdaAtV               float64
	JumpIsNegative                bool
	JumpMagnitudeModerate         bool
	MatchesScalarPortalSign       bool
	PortalFormula                 string
	RequiredPortalRatio           float64
	IfLambdaHeavyEqualsOne        float64
	CanBeGeneratedByTreePortal    bool
	HeavySectorDerived            bool
	Verdict                       string
}

type FirewallAudit struct {
	NoObservedMassClaimedAsDerivation bool
	NoThresholdJumpDerived            bool
	NoPortalCouplingFitted            bool
	NoHeavySelfQuarticFitted          bool
	NoThresholdScaleDerived           bool
	NoTwoLoopRGExecuted               bool
	NoPoleMassConversionInserted      bool
	NoFinalMassClaimed                bool
	FiniteCorePolluted                bool
	Obligations                       []RemainingObligation
	Verdict                           string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksFinalPrediction     bool
}

type Summary struct {
	TwoStageRGFormalized           bool
	JumpInsertionFormalized        bool
	RequiredJumpExtracted          bool
	JumpHasCorrectSign             bool
	JumpHasPortalMagnitude         bool
	HeavySectorDerived             bool
	FinalMassClaimed               bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	RG        TwoStageRG
	Rule      ThresholdRule
	Lanes     []TransportLane
	Jumps     []JumpExtraction
	Viability ViabilityAudit
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
	rg := formalizeTwoStageRG()
	rule := formalizeThresholdRule()
	lanes := buildTransportLanes()
	jumps := extractRequiredJumps(lanes, rg)
	viability := auditViability(jumps, rg)
	firewalls := auditFirewalls(viability)
	summary := buildSummary(rg, rule, jumps, viability, firewalls)
	truth := "Gate 314 converts the Gate-313 continuous-flow floor into an exact intermediate-threshold obligation.  In the PeV lane, the gauge-only lower envelope cannot reach a 125.10 GeV comparison target without a negative finite jump Δλ at M_threshold.  The extracted jump is of moderate quartic size and has exactly the sign expected from integrating out a heavy scalar portal, Δλ=-λ_mix²/(4λ_heavy).  This is a quantitative target for Phase-II heavy-sector dynamics, not a derivation of the heavy portal itself."
	return Analysis{RG: rg, Rule: rule, Lanes: lanes, Jumps: jumps, Viability: viability, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeTwoStageRG() TwoStageRG {
	targetLambda := massToLambda(targetPoleLikeGeV)
	return TwoStageRG{
		Formalized:       true,
		HighScaleGeV:     gutScaleGeV,
		ThresholdGeV:     thresholdScaleGeV,
		LowScaleGeV:      vevGeV,
		SegmentA:         "Λ_GUT → M_threshold with PeV/vectorlike beta coefficients inherited from Gate 309/313",
		SegmentB:         "M_threshold → v with SM one-loop beta coefficients",
		HighBetaLedger:   "b1=41/10+7.78628724237, b2=-19/6+9.65295390904, b3=-7+8.98628724237",
		LowBetaLedger:    "b1=41/10, b2=-19/6, b3=-7",
		TargetMassGeV:    targetPoleLikeGeV,
		TargetLambdaAtV:  targetLambda,
		BoundaryLambdaUV: lambdaUVBoundary,
		Verdict:          StatusTwoStageRGFormalized,
	}
}

func formalizeThresholdRule() ThresholdRule {
	return ThresholdRule{
		Formalized:               true,
		Condition:                "λ(M_threshold^-) = λ(M_threshold^+) + Δλ",
		JumpSymbol:               "Δλ",
		AppliedAtGeV:             thresholdScaleGeV,
		SignConvention:           "negative Δλ lowers the post-decoupling effective Higgs quartic",
		NegativeJumpLowersIRMass: true,
		DerivedFromHeavySector:   false,
		Verdict:                  strings.Join([]string{StatusThresholdJumpInserted, StatusFailedJumpNotDerived}, ";"),
	}
}

func buildTransportLanes() []TransportLane {
	rp := rPlus()
	return []TransportLane{
		{Name: "legacy_all_rplus_assigned_to_top", SourceGate: "Gate 309 / Gate 313 legacy lane", TopFraction: 1, YtUV: math.Sqrt(rp), Canonical: false, DiagnosticOnly: true, Description: "high-tension inherited r_+ top lane"},
		{Name: "tau_eta_unique_low_top_fraction", SourceGate: "Gate 313 τ_eta |τ|² low witness", TopFraction: 1.0 / 9.0, YtUV: math.Sqrt(rp / 9.0), Canonical: false, DiagnosticOnly: true, Description: "most aggressive tau_eta fractional top witness"},
		{Name: "gauge_only_zero_top_lower_envelope", SourceGate: "Gate 313 continuous-flow floor", TopFraction: 0, YtUV: 0, Canonical: false, DiagnosticOnly: true, Description: "stable lower envelope for the one-loop fixed-boundary system"},
	}
}

func extractRequiredJumps(lanes []TransportLane, rg TwoStageRG) []JumpExtraction {
	out := make([]JumpExtraction, 0, len(lanes))
	for _, lane := range lanes {
		out = append(out, solveJumpForLane(lane, rg))
	}
	return out
}

func solveJumpForLane(lane TransportLane, rg TwoStageRG) JumpExtraction {
	state := rgState{ScaleGeV: rg.HighScaleGeV, GY: math.Sqrt(3.0 / 5.0), G2: 1.0, G3: 1.0, YT: lane.YtUV, Lambda: rg.BoundaryLambdaUV}
	var ok bool
	state, ok = integrateSegment(state, rg.HighScaleGeV, rg.ThresholdGeV, betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237})
	if !ok {
		return JumpExtraction{LaneName: lane.Name, TopFraction: lane.TopFraction, YtUV: lane.YtUV, Solved: false, Perturbative: false, Interpretation: "transport to threshold became nonperturbative", Verdict: "FAILED_ROUTE_HIGH_SCALE_TRANSPORT_NONPERTURBATIVE"}
	}
	lambdaPlus := state.Lambda
	baseFinal, baseOK := runLowSegment(state, 0)
	baselineLambda, baselineMass := math.NaN(), math.NaN()
	if baseOK {
		baselineLambda = baseFinal.Lambda
		if baselineLambda > 0 {
			baselineMass = lambdaToMass(baselineLambda)
		}
	}
	f := func(delta float64) (float64, bool) {
		final, ok := runLowSegment(state, delta)
		if !ok {
			return math.NaN(), false
		}
		return final.Lambda - rg.TargetLambdaAtV, true
	}
	fhi, okHigh := f(0)
	if !okHigh {
		return JumpExtraction{LaneName: lane.Name, TopFraction: lane.TopFraction, YtUV: lane.YtUV, LambdaAtThresholdPlus: lambdaPlus, BaselineLambdaAtV: baselineLambda, BaselineMassGeV: baselineMass, TargetLambdaAtV: rg.TargetLambdaAtV, TargetMassGeV: rg.TargetMassGeV, Solved: false, Perturbative: baseOK, Interpretation: "baseline low-scale transport failed", Verdict: "FAILED_ROUTE_THRESHOLD_BASELINE_TRANSPORT_FAILED"}
	}
	highDelta := 0.0
	lowDelta := math.NaN()
	flo := math.NaN()
	// Search for the nearest finite negative jump that crosses the target.
	// This avoids using huge negative quartics as a fake bracket when the top
	// attractor lane is not actually solved by a moderate threshold correction.
	minDelta := -math.Max(lambdaPlus+0.25, 0.5)
	const scanSteps = 80
	prevDelta, prevF := highDelta, fhi
	for i := 1; i <= scanSteps; i++ {
		d := float64(i) * minDelta / float64(scanSteps)
		fd, okd := f(d)
		if !okd {
			continue
		}
		if fd <= 0 && prevF >= 0 {
			lowDelta, flo = d, fd
			highDelta, fhi = prevDelta, prevF
			break
		}
		prevDelta, prevF = d, fd
	}
	if math.IsNaN(lowDelta) || flo > 0 || fhi < 0 {
		return JumpExtraction{LaneName: lane.Name, TopFraction: lane.TopFraction, YtUV: lane.YtUV, LambdaAtThresholdPlus: lambdaPlus, BaselineLambdaAtV: baselineLambda, BaselineMassGeV: baselineMass, TargetLambdaAtV: rg.TargetLambdaAtV, TargetMassGeV: rg.TargetMassGeV, Solved: false, Perturbative: baseOK, Interpretation: "no moderate finite threshold jump bracketed the target in this lane", Verdict: "FAILED_ROUTE_THRESHOLD_JUMP_TARGET_NOT_BRACKETED"}
	}
	lo, hi := lowDelta, highDelta
	for i := 0; i < maxRootIterations; i++ {
		mid := 0.5 * (lo + hi)
		fm, okMid := f(mid)
		if !okMid {
			lo = mid
			continue
		}
		if math.Abs(fm) < rootTolerance || math.Abs(hi-lo) < rootTolerance {
			lo, hi = mid, mid
			break
		}
		if fm > 0 {
			hi = mid
		} else {
			lo = mid
		}
	}
	delta := 0.5 * (lo + hi)
	minus := lambdaPlus + delta
	final, finalOK := runLowSegment(state, delta)
	mass := math.NaN()
	if finalOK && final.Lambda > 0 {
		mass = lambdaToMass(final.Lambda)
	}
	mag := math.Abs(delta)
	correctSign := delta < 0
	viableOrder := mag > 0.01 && mag < 1.0 && minus > -0.5
	interp := fmt.Sprintf("requires Δλ=%.12f at M_threshold to map baseline %.3f GeV to %.3f GeV", delta, baselineMass, rg.TargetMassGeV)
	verdict := strings.Join([]string{StatusRequiredJumpExtracted, StatusViabilitySieveFormalized}, ";")
	if !correctSign {
		verdict += ";FAILED_ROUTE_THRESHOLD_JUMP_WRONG_SIGN"
	}
	return JumpExtraction{LaneName: lane.Name, TopFraction: lane.TopFraction, YtUV: lane.YtUV, LambdaAtThresholdPlus: lambdaPlus, BaselineLambdaAtV: baselineLambda, BaselineMassGeV: baselineMass, RequiredDeltaLambda: delta, LambdaAtThresholdMinus: minus, TargetLambdaAtV: final.Lambda, TargetMassGeV: mass, Solved: finalOK, Perturbative: finalOK, CorrectSign: correctSign, Magnitude: mag, ViableOrder: viableOrder, Interpretation: interp, Verdict: verdict}
}

func runLowSegment(thresholdPlus rgState, delta float64) (rgState, bool) {
	state := thresholdPlus
	state.Lambda += delta
	return integrateSegment(state, thresholdScaleGeV, vevGeV, betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0})
}

func auditViability(jumps []JumpExtraction, rg TwoStageRG) ViabilityAudit {
	pref := findJump(jumps, "gauge_only_zero_top_lower_envelope")
	if pref == nil {
		pref = firstSolvedJump(jumps)
	}
	v := ViabilityAudit{Formalized: true, TargetMassGeV: rg.TargetMassGeV, TargetLambdaAtV: rg.TargetLambdaAtV, PortalFormula: "Δλ_theory = -λ_mix²/(4λ_heavy)", HeavySectorDerived: false, Verdict: StatusFailedHeavySectorNotConstructed}
	if pref != nil {
		v.PreferredLaneName = pref.LaneName
		v.PreferredRequiredDelta = pref.RequiredDeltaLambda
		v.PreferredThresholdLambdaPlus = pref.LambdaAtThresholdPlus
		v.PreferredThresholdLambdaMinus = pref.LambdaAtThresholdMinus
		v.PreferredBaselineMassGeV = pref.BaselineMassGeV
		v.JumpIsNegative = pref.RequiredDeltaLambda < 0
		v.JumpMagnitudeModerate = math.Abs(pref.RequiredDeltaLambda) > 0.01 && math.Abs(pref.RequiredDeltaLambda) < 1.0
		v.MatchesScalarPortalSign = v.JumpIsNegative
		v.RequiredPortalRatio = -4.0 * pref.RequiredDeltaLambda
		if v.RequiredPortalRatio > 0 {
			v.IfLambdaHeavyEqualsOne = math.Sqrt(v.RequiredPortalRatio)
		}
		v.CanBeGeneratedByTreePortal = v.MatchesScalarPortalSign && v.JumpMagnitudeModerate && v.RequiredPortalRatio > 0
		v.Verdict = strings.Join([]string{StatusRequiredJumpExtracted, StatusViabilitySieveFormalized, StatusPortalMagnitudeTargetFormalized, StatusThresholdMechanismRequired, StatusFailedHeavyPortalNotDerived, StatusFailedHeavyQuarticNotDerived}, ";")
	}
	return v
}

func auditFirewalls(v ViabilityAudit) FirewallAudit {
	obs := []RemainingObligation{
		{Name: "derive λ_mix", WhyRequired: "the threshold jump only becomes predictive if the scalar/Majorana portal is computed from the finite Dirac graph", Status: StatusFailedHeavyPortalNotDerived, BlocksFinalPrediction: true},
		{Name: "derive λ_heavy", WhyRequired: "the self-quartic of the decoupled heavy mode fixes the denominator of -λ_mix²/(4λ_heavy)", Status: StatusFailedHeavyQuarticNotDerived, BlocksFinalPrediction: true},
		{Name: "derive M_threshold", WhyRequired: "Gate 314 uses the inherited PeV scale as a conditional lane, not as a final theorem", Status: StatusFailedThresholdScaleStillSealed, BlocksFinalPrediction: true},
		{Name: "execute two-loop and pole matching", WhyRequired: "precision comparison with collider pole mass requires higher-order transport", Status: StatusFailedTwoLoopNotExecuted + ";" + StatusFailedPoleMassNotComputed, BlocksFinalPrediction: true},
	}
	return FirewallAudit{NoObservedMassClaimedAsDerivation: true, NoThresholdJumpDerived: true, NoPortalCouplingFitted: true, NoHeavySelfQuarticFitted: true, NoThresholdScaleDerived: true, NoTwoLoopRGExecuted: true, NoPoleMassConversionInserted: true, NoFinalMassClaimed: true, FiniteCorePolluted: false, Obligations: obs, Verdict: strings.Join([]string{StatusGate314FirewallsPreserved, StatusFailedJumpNotDerived, StatusFailedFinalMassNotClaimed}, ";")}
}

func buildSummary(rg TwoStageRG, rule ThresholdRule, jumps []JumpExtraction, v ViabilityAudit, f FirewallAudit) Summary {
	return Summary{TwoStageRGFormalized: rg.Formalized, JumpInsertionFormalized: rule.Formalized, RequiredJumpExtracted: v.PreferredLaneName != "" && v.PreferredRequiredDelta != 0, JumpHasCorrectSign: v.JumpIsNegative, JumpHasPortalMagnitude: v.CanBeGeneratedByTreePortal, HeavySectorDerived: v.HeavySectorDerived, FinalMassClaimed: false, FirewallPreserved: !f.FiniteCorePolluted && f.NoFinalMassClaimed && f.NoThresholdJumpDerived && f.NoPortalCouplingFitted, Status: strings.Join([]string{StatusTwoStageRGFormalized, StatusRequiredJumpExtracted, StatusThresholdMechanismRequired}, ";"), DirectAnswer: fmt.Sprintf("The continuous gauge-only floor %.3f GeV can be moved to %.2f GeV by a PeV-scale finite matching jump Δλ≈%.6f in the preferred lane; this has the correct negative sign and scalar-portal order, but the heavy-sector portal has not been derived.", v.PreferredBaselineMassGeV, rg.TargetMassGeV, v.PreferredRequiredDelta), NextGate: "Derive the heavy-sector portal/self-quartic and the threshold scale from the B-gap/PeV finite Dirac graph rather than inserting Δλ as a target."}
}

func findJump(jumps []JumpExtraction, name string) *JumpExtraction {
	for i := range jumps {
		if jumps[i].LaneName == name && jumps[i].Solved {
			return &jumps[i]
		}
	}
	return nil
}

func firstSolvedJump(jumps []JumpExtraction) *JumpExtraction {
	for i := range jumps {
		if jumps[i].Solved {
			return &jumps[i]
		}
	}
	return nil
}

func rPlus() float64                      { return (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0 }
func massToLambda(m float64) float64      { return (m / vevGeV) * (m / vevGeV) / 2.0 }
func lambdaToMass(lambda float64) float64 { return vevGeV * math.Sqrt(2.0*lambda) }

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
	return state, true
}

func rk4Step(s rgState, h float64, b betaCoefficients) rgState {
	k1 := beta(s, b)
	k2 := beta(addScaled(s, k1, h/2), b)
	k3 := beta(addScaled(s, k2, h/2), b)
	k4 := beta(addScaled(s, k3, h), b)
	return rgState{ScaleGeV: s.ScaleGeV, GY: s.GY + h*(k1.GY+2*k2.GY+2*k3.GY+k4.GY)/6, G2: s.G2 + h*(k1.G2+2*k2.G2+2*k3.G2+k4.G2)/6, G3: s.G3 + h*(k1.G3+2*k2.G3+2*k3.G3+k4.G3)/6, YT: s.YT + h*(k1.YT+2*k2.YT+2*k3.YT+k4.YT)/6, Lambda: s.Lambda + h*(k1.Lambda+2*k2.Lambda+2*k3.Lambda+k4.Lambda)/6}
}

func beta(s rgState, b betaCoefficients) rgState {
	loop := 16.0 * math.Pi * math.Pi
	g1, g2, g3, yt, l := s.GY, s.G2, s.G3, s.YT, s.Lambda
	dg1 := (b.B1GUT * g1 * g1 * g1) / loop
	dg2 := (b.B2 * g2 * g2 * g2) / loop
	dg3 := (b.B3 * g3 * g3 * g3) / loop
	dyt := 0.0
	if yt != 0 {
		dyt = yt * (4.5*yt*yt - 8*g3*g3 - 2.25*g2*g2 - 17.0/20.0*g1*g1) / loop
	}
	dl := (24*l*l + 12*l*yt*yt - 12*math.Pow(yt, 4) + (3.0/16.0)*(2*math.Pow(g2, 4)+math.Pow(g2*g2+g1*g1, 2)) - l*(9*g2*g2+3*g1*g1)) / loop
	return rgState{GY: dg1, G2: dg2, G3: dg3, YT: dyt, Lambda: dl}
}

func addScaled(s, k rgState, scale float64) rgState {
	return rgState{ScaleGeV: s.ScaleGeV, GY: s.GY + scale*k.GY, G2: s.G2 + scale*k.G2, G3: s.G3 + scale*k.G3, YT: s.YT + scale*k.YT, Lambda: s.Lambda + scale*k.Lambda}
}

func stateFinite(s rgState) bool {
	vals := []float64{s.GY, s.G2, s.G3, s.YT, s.Lambda}
	for _, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func statePerturbative(s rgState) bool {
	return s.GY*s.GY < perturbativeLimitSq && s.G2*s.G2 < perturbativeLimitSq && s.G3*s.G3 < perturbativeLimitSq && s.YT*s.YT < perturbativeLimitSq && math.Abs(s.Lambda) < perturbativeLimitSq
}

func FormatRG(r TwoStageRG) string {
	return fmt.Sprintf("formalized=%t high=%.6g threshold=%.6g low=%.6g lambdaUV=%.12f targetLambda=%.12f verdict=%s", r.Formalized, r.HighScaleGeV, r.ThresholdGeV, r.LowScaleGeV, r.BoundaryLambdaUV, r.TargetLambdaAtV, r.Verdict)
}

func FormatRule(r ThresholdRule) string {
	return fmt.Sprintf("formalized=%t condition=%q sign=%q negativeLowers=%t derived=%t verdict=%s", r.Formalized, r.Condition, r.SignConvention, r.NegativeJumpLowersIRMass, r.DerivedFromHeavySector, r.Verdict)
}

func FormatLane(l TransportLane) string {
	return fmt.Sprintf("%s fraction=%.6f ytUV=%.9f canonical=%t diagnostic=%t", l.Name, l.TopFraction, l.YtUV, l.Canonical, l.DiagnosticOnly)
}

func FormatJump(j JumpExtraction) string {
	return fmt.Sprintf("%s solved=%t topFrac=%.6f lambdaPlus=%.12f baselineLambdaV=%.12f baselineMass=%.6f delta=%.12f lambdaMinus=%.12f targetLambda=%.12f targetMass=%.6f sign=%t viable=%t verdict=%s", j.LaneName, j.Solved, j.TopFraction, j.LambdaAtThresholdPlus, j.BaselineLambdaAtV, j.BaselineMassGeV, j.RequiredDeltaLambda, j.LambdaAtThresholdMinus, j.TargetLambdaAtV, j.TargetMassGeV, j.CorrectSign, j.ViableOrder, j.Verdict)
}

func FormatViability(v ViabilityAudit) string {
	return fmt.Sprintf("formalized=%t preferred=%s baselineMass=%.6f targetMass=%.6f delta=%.12f lambdaPlus=%.12f lambdaMinus=%.12f negative=%t moderate=%t portal=%t ratio=%.12f mixIfHeavy1=%.12f derived=%t verdict=%s", v.Formalized, v.PreferredLaneName, v.PreferredBaselineMassGeV, v.TargetMassGeV, v.PreferredRequiredDelta, v.PreferredThresholdLambdaPlus, v.PreferredThresholdLambdaMinus, v.JumpIsNegative, v.JumpMagnitudeModerate, v.CanBeGeneratedByTreePortal, v.RequiredPortalRatio, v.IfLambdaHeavyEqualsOne, v.HeavySectorDerived, v.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noObsAsDerivation=%t noJumpDerived=%t noPortal=%t noHeavyQuartic=%t noScale=%t noTwoLoop=%t noPole=%t noFinal=%t polluted=%t obligations=%d verdict=%s", f.NoObservedMassClaimedAsDerivation, f.NoThresholdJumpDerived, f.NoPortalCouplingFitted, f.NoHeavySelfQuarticFitted, f.NoThresholdScaleDerived, f.NoTwoLoopRGExecuted, f.NoPoleMassConversionInserted, f.NoFinalMassClaimed, f.FiniteCorePolluted, len(f.Obligations), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("twoStage=%t jumpRule=%t extracted=%t sign=%t portalMag=%t derived=%t final=%t firewall=%t status=%s answer=%q next=%q", s.TwoStageRGFormalized, s.JumpInsertionFormalized, s.RequiredJumpExtracted, s.JumpHasCorrectSign, s.JumpHasPortalMagnitude, s.HeavySectorDerived, s.FinalMassClaimed, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
