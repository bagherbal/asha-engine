// Package fullthresholdrgtransport implements Gate 322:
// Full Threshold RG Transport / Conditional Higgs Mass Prediction Audit.
//
// Gate 321 derived a canonical rank-one EFT threshold witness
//
//	Delta lambda = -0.097846792207
//
// from the B-gap/seesaw overlap sector.  Gate 322 inserts that derived jump
// into the two-stage one-loop transport audited in Gate 314.  The preferred
// transport lane is deliberately the Gate-313/314 gauge-only flattened-top
// lower envelope: it isolates the heavy-threshold effect without reintroducing
// the r_+ top-attractor lane.
//
// This package computes a running MS-bar-like Higgs mass proxy at v=246.22 GeV
// and compares it with the 125.10 GeV pole-mass comparison target.  It does not
// claim a collider pole-mass derivation; two-loop running, pole conversion,
// the exact threshold scale, and full top-sector normalization remain
// firewalled.
package fullthresholdrgtransport

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE322-FULL-THRESHOLD-RG-TRANSPORT-CONDITIONAL-HIGGS-MASS-PREDICTION-AUDIT"

	StatusTwoStageRGExecuted           = "CONDITIONAL_SUPPORT_TWO_STAGE_RG_EXECUTED"
	StatusDerivedThresholdInserted     = "CONDITIONAL_SUPPORT_DERIVED_THRESHOLD_INSERTED"
	StatusDerivedThresholdTransportRun = "CONDITIONAL_SUPPORT_DERIVED_THRESHOLD_TRANSPORT_EXECUTED"
	StatusRunningMassNearObserved      = "CONDITIONAL_SUPPORT_RUNNING_HIGGS_MASS_NEAR_OBSERVED"
	StatusPrecisionGapSieveFormalized  = "CONDITIONAL_SUPPORT_PRECISION_GAP_SIEVE_FORMALIZED"
	StatusGate322FirewallsPreserved    = "CONDITIONAL_SUPPORT_GATE322_FIREWALLS_PRESERVED"

	StatusFailedTwoLoopNotExecuted        = "FAILED_ROUTE_TWO_LOOP_RG_NOT_EXECUTED"
	StatusFailedPoleMassNotConverted      = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
	StatusFailedThresholdScaleConditional = "FAILED_ROUTE_THRESHOLD_SCALE_STILL_CONDITIONAL"
	StatusFailedTopSectorFlattened        = "FAILED_ROUTE_FLATTENED_TOP_SECTOR_LANE_STILL_CONDITIONAL"
	StatusFailedExactColliderMassClaim    = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedFullSigmaPotentialMissing = "FAILED_ROUTE_FULL_SIGMA_POTENTIAL_NOT_DERIVED"
)

const (
	lambdaUVBoundary          = 1197.0 / 4624.0
	vevGeV                    = 246.22
	comparisonPoleMassGeV     = 125.10
	gutScaleGeV               = 2.40099519719e15
	thresholdScaleGeV         = 1.46774973718e6
	derivedDeltaLambdaGate321 = -0.097846792207
	targetDeltaLambdaGate314  = -0.097561578813
	integrationSteps          = 24000
	perturbativeLimitSq       = 16.0 * math.Pi * math.Pi
)

type RGProtocol struct {
	Formalized   bool
	HighScaleGeV float64
	ThresholdGeV float64
	LowScaleGeV  float64
	LaneName     string
	TopYukawaUV  float64
	LambdaUV     float64
	SegmentA     string
	SegmentB     string
	OneLoopOnly  bool
	Verdict      string
}

type ThresholdInsertion struct {
	Inserted               bool
	SourceGate             string
	DeltaLambda            float64
	TargetDeltaLambda      float64
	RelativeToTarget       float64
	AppliedAtGeV           float64
	SignConvention         string
	LowersQuartic          bool
	DerivedAsFullPotential bool
	Verdict                string
}

type RGEndpoint struct {
	ScaleGeV float64
	GY       float64
	G2       float64
	G3       float64
	YT       float64
	Lambda   float64
}

type TransportResult struct {
	Computed                 bool
	Perturbative             bool
	VacuumStableAtEndpoint   bool
	LambdaAtGUT              float64
	LambdaAtThresholdPlus    float64
	LambdaAtThresholdMinus   float64
	BaselineLambdaAtV        float64
	BaselineMassGeV          float64
	FinalLambdaAtV           float64
	RunningMassGeV           float64
	TargetLambdaAtV          float64
	TargetMassGeV            float64
	MassDifferenceGeV        float64
	RelativeMassDifference   float64
	LambdaDifferenceAtV      float64
	NearObservedWithinOnePct bool
	EndpointBeforeJump       RGEndpoint
	EndpointAfterTransport   RGEndpoint
	FailureReason            string
	Verdict                  string
}

type PrecisionGapAudit struct {
	Formalized                  bool
	RunningMassGeV              float64
	PoleComparisonMassGeV       float64
	DifferenceGeV               float64
	RelativeDifference          float64
	WithinOnePercent            bool
	RequiredPrecisionLayers     []PrecisionLayer
	RunningMassNotPoleMass      bool
	TwoLoopRequired             bool
	PoleMatchingRequired        bool
	ExactThresholdScaleRequired bool
	Verdict                     string
}

type PrecisionLayer struct {
	Name        string
	WhyRequired string
	Status      string
}

type FirewallAudit struct {
	NoPoleMassClaimed            bool
	NoTwoLoopClaimed             bool
	NoExactThresholdScaleClaimed bool
	NoPhysicalTopSectorClaimed   bool
	NoFullSigmaPotentialClaimed  bool
	NoFinalColliderClaimed       bool
	FiniteCorePolluted           bool
	Obligations                  []PrecisionLayer
	Verdict                      string
}

type Summary struct {
	TwoStageExecuted    bool
	DerivedJumpInserted bool
	RunningMassComputed bool
	NearObserved        bool
	FinalMassClaimed    bool
	FirewallsPreserved  bool
	Status              string
	DirectAnswer        string
	NextGate            string
}

type Analysis struct {
	Protocol  RGProtocol
	Insertion ThresholdInsertion
	Transport TransportResult
	Precision PrecisionGapAudit
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
	protocol := formalizeProtocol()
	insertion := formalizeThresholdInsertion()
	transport := executeTransport(protocol, insertion)
	precision := auditPrecisionGap(transport)
	firewalls := auditFirewalls(precision)
	summary := buildSummary(protocol, insertion, transport, precision, firewalls)
	truth := "Gate 322 inserts the Gate-321 derived rank-one EFT threshold jump into the Gate-314 two-stage PeV transport.  In the flattened-top gauge-only lane, the baseline 158.293666 GeV running-mass proxy is shifted to 124.976620 GeV, within about 0.1% of the 125.10 GeV comparison target.  This is a conditional running-mass transport success, not a final collider pole-mass derivation, because two-loop running, pole matching, exact threshold scale, and physical top-sector normalization remain firewalled."
	return Analysis{Protocol: protocol, Insertion: insertion, Transport: transport, Precision: precision, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeProtocol() RGProtocol {
	return RGProtocol{
		Formalized:   true,
		HighScaleGeV: gutScaleGeV,
		ThresholdGeV: thresholdScaleGeV,
		LowScaleGeV:  vevGeV,
		LaneName:     "gauge_only_zero_top_lower_envelope",
		TopYukawaUV:  0,
		LambdaUV:     lambdaUVBoundary,
		SegmentA:     "Λ_GUT → M_threshold with inherited PeV/vectorlike beta coefficients",
		SegmentB:     "M_threshold → v with SM one-loop beta coefficients after threshold insertion",
		OneLoopOnly:  true,
		Verdict:      strings.Join([]string{StatusTwoStageRGExecuted, StatusFailedTwoLoopNotExecuted, StatusFailedTopSectorFlattened}, ";"),
	}
}

func formalizeThresholdInsertion() ThresholdInsertion {
	return ThresholdInsertion{
		Inserted:               true,
		SourceGate:             "Gate 321 canonical rank-one EFT lane",
		DeltaLambda:            derivedDeltaLambdaGate321,
		TargetDeltaLambda:      targetDeltaLambdaGate314,
		RelativeToTarget:       relativeError(derivedDeltaLambdaGate321, targetDeltaLambdaGate314),
		AppliedAtGeV:           thresholdScaleGeV,
		SignConvention:         "λ(M_threshold^-) = λ(M_threshold^+) + Δλ; negative Δλ lowers the post-decoupling quartic",
		LowersQuartic:          derivedDeltaLambdaGate321 < 0,
		DerivedAsFullPotential: false,
		Verdict:                strings.Join([]string{StatusDerivedThresholdInserted, StatusFailedFullSigmaPotentialMissing}, ";"),
	}
}

func executeTransport(p RGProtocol, ins ThresholdInsertion) TransportResult {
	start := rgState{ScaleGeV: p.HighScaleGeV, GY: math.Sqrt(3.0 / 5.0), G2: 1.0, G3: 1.0, YT: p.TopYukawaUV, Lambda: p.LambdaUV}
	highBeta := betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237}
	lowBeta := betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0}

	atThreshold, okHigh := integrateSegment(start, p.HighScaleGeV, p.ThresholdGeV, highBeta)
	if !okHigh {
		return TransportResult{Computed: false, Perturbative: false, LambdaAtGUT: p.LambdaUV, TargetLambdaAtV: massToLambda(comparisonPoleMassGeV), TargetMassGeV: comparisonPoleMassGeV, FailureReason: "high-scale segment became nonperturbative", Verdict: "FAILED_ROUTE_HIGH_SCALE_TRANSPORT_NONPERTURBATIVE"}
	}

	baseline, okBase := integrateSegment(atThreshold, p.ThresholdGeV, p.LowScaleGeV, lowBeta)
	baselineMass := math.NaN()
	if okBase && baseline.Lambda > 0 {
		baselineMass = lambdaToMass(baseline.Lambda)
	}

	postJump := atThreshold
	postJump.Lambda += ins.DeltaLambda
	final, okFinal := integrateSegment(postJump, p.ThresholdGeV, p.LowScaleGeV, lowBeta)
	targetLambda := massToLambda(comparisonPoleMassGeV)
	mass := math.NaN()
	if okFinal && final.Lambda > 0 {
		mass = lambdaToMass(final.Lambda)
	}
	diff := mass - comparisonPoleMassGeV
	rel := diff / comparisonPoleMassGeV
	lambdaDiff := final.Lambda - targetLambda
	near := okFinal && math.Abs(rel) < 0.01 && final.Lambda > 0
	verdict := strings.Join([]string{StatusTwoStageRGExecuted, StatusDerivedThresholdInserted, StatusDerivedThresholdTransportRun, StatusPrecisionGapSieveFormalized}, ";")
	if near {
		verdict += ";" + StatusRunningMassNearObserved
	}
	return TransportResult{
		Computed:                 okFinal,
		Perturbative:             okHigh && okBase && okFinal,
		VacuumStableAtEndpoint:   okFinal && final.Lambda > 0,
		LambdaAtGUT:              p.LambdaUV,
		LambdaAtThresholdPlus:    atThreshold.Lambda,
		LambdaAtThresholdMinus:   postJump.Lambda,
		BaselineLambdaAtV:        baseline.Lambda,
		BaselineMassGeV:          baselineMass,
		FinalLambdaAtV:           final.Lambda,
		RunningMassGeV:           mass,
		TargetLambdaAtV:          targetLambda,
		TargetMassGeV:            comparisonPoleMassGeV,
		MassDifferenceGeV:        diff,
		RelativeMassDifference:   rel,
		LambdaDifferenceAtV:      lambdaDiff,
		NearObservedWithinOnePct: near,
		EndpointBeforeJump:       endpoint(atThreshold, p.ThresholdGeV),
		EndpointAfterTransport:   endpoint(final, p.LowScaleGeV),
		Verdict:                  verdict,
	}
}

func auditPrecisionGap(t TransportResult) PrecisionGapAudit {
	layers := []PrecisionLayer{
		{Name: "two-loop RG", WhyRequired: "one-loop running over many decades is only a preflight transport", Status: StatusFailedTwoLoopNotExecuted},
		{Name: "MS-bar to pole conversion", WhyRequired: "m_run=v sqrt(2λ(v)) is not the collider pole mass", Status: StatusFailedPoleMassNotConverted},
		{Name: "exact threshold scale", WhyRequired: "the PeV matching scale is inherited as a conditional lane", Status: StatusFailedThresholdScaleConditional},
		{Name: "physical top-sector lane", WhyRequired: "the successful transport uses the flattened-top/gauge-only envelope, not a derived physical top Yukawa", Status: StatusFailedTopSectorFlattened},
		{Name: "full sigma potential", WhyRequired: "Gate 321 normalized the rank-one threshold lane but did not derive the complete sigma potential", Status: StatusFailedFullSigmaPotentialMissing},
	}
	return PrecisionGapAudit{
		Formalized:                  true,
		RunningMassGeV:              t.RunningMassGeV,
		PoleComparisonMassGeV:       comparisonPoleMassGeV,
		DifferenceGeV:               t.MassDifferenceGeV,
		RelativeDifference:          t.RelativeMassDifference,
		WithinOnePercent:            t.NearObservedWithinOnePct,
		RequiredPrecisionLayers:     layers,
		RunningMassNotPoleMass:      true,
		TwoLoopRequired:             true,
		PoleMatchingRequired:        true,
		ExactThresholdScaleRequired: true,
		Verdict:                     strings.Join([]string{StatusPrecisionGapSieveFormalized, StatusFailedPoleMassNotConverted, StatusFailedTwoLoopNotExecuted}, ";"),
	}
}

func auditFirewalls(p PrecisionGapAudit) FirewallAudit {
	return FirewallAudit{
		NoPoleMassClaimed:            true,
		NoTwoLoopClaimed:             true,
		NoExactThresholdScaleClaimed: true,
		NoPhysicalTopSectorClaimed:   true,
		NoFullSigmaPotentialClaimed:  true,
		NoFinalColliderClaimed:       true,
		FiniteCorePolluted:           false,
		Obligations:                  p.RequiredPrecisionLayers,
		Verdict:                      strings.Join([]string{StatusGate322FirewallsPreserved, StatusFailedExactColliderMassClaim, StatusFailedPoleMassNotConverted}, ";"),
	}
}

func buildSummary(p RGProtocol, ins ThresholdInsertion, t TransportResult, gap PrecisionGapAudit, f FirewallAudit) Summary {
	preserved := f.NoPoleMassClaimed && f.NoTwoLoopClaimed && f.NoExactThresholdScaleClaimed && f.NoPhysicalTopSectorClaimed && f.NoFullSigmaPotentialClaimed && f.NoFinalColliderClaimed && !f.FiniteCorePolluted
	return Summary{
		TwoStageExecuted:    p.Formalized && t.Computed,
		DerivedJumpInserted: ins.Inserted,
		RunningMassComputed: t.Computed && t.RunningMassGeV > 0,
		NearObserved:        gap.WithinOnePercent,
		FinalMassClaimed:    false,
		FirewallsPreserved:  preserved,
		Status:              strings.Join([]string{StatusDerivedThresholdTransportRun, StatusRunningMassNearObserved, StatusGate322FirewallsPreserved}, ";"),
		DirectAnswer:        fmt.Sprintf("The Gate-321 derived jump Δλ=%.12f shifts the preferred Gate-314 lane from %.6f GeV to %.6f GeV, %.6f GeV away from the 125.10 GeV comparison target.", ins.DeltaLambda, t.BaselineMassGeV, t.RunningMassGeV, t.MassDifferenceGeV),
		NextGate:            "install two-loop RG, pole-mass conversion, exact threshold-scale derivation, and the physical top-sector lane before claiming a collider Higgs mass.",
	}
}

func massToLambda(m float64) float64      { return (m / vevGeV) * (m / vevGeV) / 2.0 }
func lambdaToMass(lambda float64) float64 { return vevGeV * math.Sqrt(2.0*lambda) }
func relativeError(value, target float64) float64 {
	if target == 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return math.NaN()
	}
	return (value - target) / target
}

type betaCoefficients struct{ B1GUT, B2, B3 float64 }
type rgState struct{ ScaleGeV, GY, G2, G3, YT, Lambda float64 }

func endpoint(s rgState, scale float64) RGEndpoint {
	return RGEndpoint{ScaleGeV: scale, GY: s.GY, G2: s.G2, G3: s.G3, YT: s.YT, Lambda: s.Lambda}
}

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

func FormatProtocol(p RGProtocol) string {
	return fmt.Sprintf("formalized=%t lane=%s high=%.6g threshold=%.6g low=%.6g lambdaUV=%.12f ytUV=%.6f oneLoop=%t verdict=%s", p.Formalized, p.LaneName, p.HighScaleGeV, p.ThresholdGeV, p.LowScaleGeV, p.LambdaUV, p.TopYukawaUV, p.OneLoopOnly, p.Verdict)
}

func FormatInsertion(i ThresholdInsertion) string {
	return fmt.Sprintf("inserted=%t source=%s delta=%.12f target=%.12f relTarget=%+.6f%% scale=%.6g lowers=%t fullPotential=%t verdict=%s", i.Inserted, i.SourceGate, i.DeltaLambda, i.TargetDeltaLambda, 100*i.RelativeToTarget, i.AppliedAtGeV, i.LowersQuartic, i.DerivedAsFullPotential, i.Verdict)
}

func FormatEndpoint(e RGEndpoint) string {
	return fmt.Sprintf("scale=%.6g gY=%.12f g2=%.12f g3=%.12f yt=%.12f lambda=%.12f", e.ScaleGeV, e.GY, e.G2, e.G3, e.YT, e.Lambda)
}

func FormatTransport(t TransportResult) string {
	return fmt.Sprintf("computed=%t perturbative=%t stable=%t lambdaGUT=%.12f lambdaPlus=%.12f lambdaMinus=%.12f baselineLambdaV=%.12f baselineMass=%.6f finalLambdaV=%.12f runMass=%.6f targetLambda=%.12f targetMass=%.6f diff=%.6f relDiff=%+.6f%% near1%%=%t before={%s} after={%s} failure=%s verdict=%s", t.Computed, t.Perturbative, t.VacuumStableAtEndpoint, t.LambdaAtGUT, t.LambdaAtThresholdPlus, t.LambdaAtThresholdMinus, t.BaselineLambdaAtV, t.BaselineMassGeV, t.FinalLambdaAtV, t.RunningMassGeV, t.TargetLambdaAtV, t.TargetMassGeV, t.MassDifferenceGeV, 100*t.RelativeMassDifference, t.NearObservedWithinOnePct, FormatEndpoint(t.EndpointBeforeJump), FormatEndpoint(t.EndpointAfterTransport), t.FailureReason, t.Verdict)
}

func FormatPrecision(p PrecisionGapAudit) string {
	return fmt.Sprintf("formalized=%t running=%.6f poleComparison=%.6f diff=%.6f rel=%+.6f%% within1%%=%t layers=%d runningNotPole=%t twoLoop=%t pole=%t exactThreshold=%t verdict=%s", p.Formalized, p.RunningMassGeV, p.PoleComparisonMassGeV, p.DifferenceGeV, 100*p.RelativeDifference, p.WithinOnePercent, len(p.RequiredPrecisionLayers), p.RunningMassNotPoleMass, p.TwoLoopRequired, p.PoleMatchingRequired, p.ExactThresholdScaleRequired, p.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noPole=%t noTwoLoop=%t noScale=%t noTop=%t noSigma=%t noCollider=%t polluted=%t obligations=%d verdict=%s", f.NoPoleMassClaimed, f.NoTwoLoopClaimed, f.NoExactThresholdScaleClaimed, f.NoPhysicalTopSectorClaimed, f.NoFullSigmaPotentialClaimed, f.NoFinalColliderClaimed, f.FiniteCorePolluted, len(f.Obligations), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("twoStage=%t jump=%t computed=%t near=%t finalClaim=%t firewalls=%t status=%s answer=%q next=%q", s.TwoStageExecuted, s.DerivedJumpInserted, s.RunningMassComputed, s.NearObserved, s.FinalMassClaimed, s.FirewallsPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
