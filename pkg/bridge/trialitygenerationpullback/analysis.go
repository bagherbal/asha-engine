// Package trialitygenerationpullback implements Gate 323:
// Triality Generation Pullback / Native Top-Yukawa Boundary Sieve.
//
// Gate 313 showed that treating the Gate-291 r_+ amplitude as a single top
// Yukawa entry creates the one-loop top-attractor tension, while fractionalizing
// the total up-type trace across generations flattens the slope.  Gate 322 then
// demonstrated that the derived B-gap threshold jump gives a near-125 GeV
// running-mass proxy only in the diagnostic gauge-only/flattened-top envelope.
//
// Gate 323 audits the missing mathematical bridge: whether the tau_eta=(2,-2,1)
// triality generation topology can be pulled back onto the physical quark trace
// carrier P_Q H_F strongly enough to derive a unique top-Yukawa boundary.  The
// gate formalizes the pullback, constructs the magnitude-squared weights
// (4/9,4/9,1/9), and reruns the Gate-322 threshold transport for the resulting
// top-fraction lanes.  It is deliberately strict: tau_eta supplies a native
// generation-breaking tensor and a finite set of top-boundary candidates, but it
// still does not derive the physical top slot because the two |tau|=2 slots are
// degenerate and no CKM/flavor-orientation operator is installed.
package trialitygenerationpullback

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE323-TRIALITY-GENERATION-PULLBACK-NATIVE-TOP-YUKAWA-BOUNDARY-SIEVE"

	StatusTrialityPullbackFormalized      = "CONDITIONAL_SUPPORT_TRIALITY_PULLBACK_FORMALIZED"
	StatusAmplitudeFractionalization      = "CONDITIONAL_SUPPORT_AMPLITUDE_FRACTIONALIZATION_EXTRACTED"
	StatusTopBoundaryCandidatesAudited    = "CONDITIONAL_SUPPORT_NATIVE_TOP_YUKAWA_BOUNDARY_CANDIDATES_AUDITED"
	StatusPhysicalLanePreflightExecuted   = "CONDITIONAL_SUPPORT_PHYSICAL_LANE_PREFLIGHT_EXECUTED"
	StatusThresholdSuccessRequiresTopMute = "CONDITIONAL_TENSION_THRESHOLD_SUCCESS_REQUIRES_FLATTENED_TOP_ENVELOPE"
	StatusTauEtaDegeneracy                = "CONDITIONAL_TENSION_TAU_ETA_HIGH_MAGNITUDE_DEGENERACY_REMAINS"

	StatusFailedTopSlotNotCanonical         = "FAILED_ROUTE_CANONICAL_TOP_SLOT_NOT_UNIQUELY_DERIVED"
	StatusFailedNativeTopBoundaryNotDerived = "FAILED_ROUTE_NATIVE_TOP_YUKAWA_BOUNDARY_NOT_DERIVED"
	StatusFailedFlavorOrientationMissing    = "FAILED_ROUTE_FLAVOR_ORIENTATION_OPERATOR_NOT_DERIVED"
	StatusFailedCKMTextureMissing           = "FAILED_ROUTE_CKM_TEXTURE_NOT_DERIVED"
	StatusFailedTopLaneSpoils125            = "FAILED_ROUTE_NONZERO_TOP_FRACTION_SPOILS_GATE322_125GEV_PROXY"
	StatusFailedPoleMassNotExecuted         = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
	StatusFailedTwoLoopNotExecuted          = "FAILED_ROUTE_TWO_LOOP_RG_NOT_EXECUTED"
	StatusFailedColliderMassNotClaimed      = "FAILED_ROUTE_FINAL_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	lambdaUVBoundary          = 1197.0 / 4624.0
	vevGeV                    = 246.22
	comparisonPoleMassGeV     = 125.10
	gutScaleGeV               = 2.40099519719e15
	thresholdScaleGeV         = 1.46774973718e6
	derivedDeltaLambdaGate321 = -0.097846792207
	integrationSteps          = 24000
	perturbativeLimitSq       = 16.0 * math.Pi * math.Pi
)

type TrialityCarrier struct {
	Formalized                  bool
	SourceGate                  string
	TauEta                      []int
	PullbackCarrier             string
	Projector                   string
	SignedSpectrumDistinct      bool
	MagnitudeSpectrumDegenerate bool
	GenerationBreaking          bool
	PullbackUnique              bool
	Verdict                     string
}

type Fractionalization struct {
	Formalized             bool
	Equation               string
	RPlusExact             string
	RPlusDecimal           float64
	TauMagnitudeSquares    []float64
	NormalizedWeights      []float64
	SumWeights             float64
	GeneratesUniqueLowSlot bool
	GeneratesTwoHighSlots  bool
	Verdict                string
}

type TopSlotCandidate struct {
	Name                         string
	SlotIndex                    int
	TauEtaValue                  int
	Weight                       float64
	TopYtSquaredOverGStarSquared float64
	TopYtUV                      float64
	SelectionRule                string
	DerivedFromTauEta            bool
	Canonical                    bool
	Ambiguous                    bool
	Verdict                      string
}

type RGEndpoint struct {
	ScaleGeV float64
	GY       float64
	G2       float64
	G3       float64
	YT       float64
	Lambda   float64
}

type TransportPreflight struct {
	CandidateName          string
	TopFraction            float64
	TopYtUV                float64
	LambdaAtThresholdPlus  float64
	LambdaAtThresholdMinus float64
	BaselineLambdaAtV      float64
	BaselineMassGeV        float64
	FinalLambdaAtV         float64
	RunningMassGeV         float64
	TargetMassGeV          float64
	DifferenceGeV          float64
	RelativeDifference     float64
	Near125WithinOnePct    bool
	Perturbative           bool
	VacuumStable           bool
	Endpoint               RGEndpoint
	Verdict                string
}

type PullbackVerdict struct {
	Formalized                  bool
	CanonicalTopFractionDerived bool
	PreferredPhysicalCandidate  string
	BestNear125Candidate        string
	BestNear125MassGeV          float64
	GaugeOnlyStillRequired      bool
	NonzeroTauEtaTopSpoils125   bool
	TopBoundaryStatus           string
	RequiredNextOperator        string
	Verdict                     string
}

type FirewallAudit struct {
	NoObservedTopMassInserted  bool
	NoCKMImported              bool
	NoFlavorTextureInvented    bool
	NoPoleMassClaimed          bool
	NoTwoLoopClaimed           bool
	NoFinalColliderMassClaimed bool
	FiniteCorePolluted         bool
	Verdict                    string
}

type Summary struct {
	TrialityPullbackFormalized            bool
	FractionalizationExtracted            bool
	TopCandidatesAudited                  bool
	NativeTopBoundaryDerived              bool
	PhysicalPreflightExecuted             bool
	Gate322SuccessPreservedByCanonicalTop bool
	FirewallsPreserved                    bool
	FinalMassClaimed                      bool
	Status                                string
	DirectAnswer                          string
	NextGate                              string
}

type Analysis struct {
	Carrier         TrialityCarrier
	Fractional      Fractionalization
	Candidates      []TopSlotCandidate
	Preflights      []TransportPreflight
	PullbackVerdict PullbackVerdict
	Firewalls       FirewallAudit
	Summary         Summary
	Truth           string
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
	carrier := formalizeCarrier()
	frac := extractFractionalization()
	candidates := buildCandidates(frac)
	preflights := runPreflights(candidates)
	verdict := auditPullbackVerdict(carrier, candidates, preflights)
	firewalls := auditFirewalls(verdict)
	summary := buildSummary(carrier, frac, candidates, preflights, verdict, firewalls)
	truth := "Gate 323 pulls tau_eta=(2,-2,1) back onto the three-generation quark trace carrier and derives the normalized magnitude-squared weights (4/9,4/9,1/9).  This is a real native generation-breaking fractionalization of the total r_+ up-type trace, but it does not yet derive a unique physical top slot: the two |tau|=2 slots are degenerate and the unique |tau|=1 slot is an orientation choice, not a proven top eigenvector.  When the Gate-321 threshold jump is rerun, the gauge-only diagnostic lane remains near 124.98 GeV, while the tau_eta nonzero-top candidates drive the running-mass proxy far above 125 GeV.  Therefore Gate 323 formalizes the pullback and candidate top boundaries, but preserves the firewall that the physical top Yukawa requires an additional flavor-orientation/CKM texture operator."
	return Analysis{Carrier: carrier, Fractional: frac, Candidates: candidates, Preflights: preflights, PullbackVerdict: verdict, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeCarrier() TrialityCarrier {
	return TrialityCarrier{
		Formalized:                  true,
		SourceGate:                  "Gate 242 tau_eta spatial tagging / triality generation topology",
		TauEta:                      []int{2, -2, 1},
		PullbackCarrier:             "P_Q H_F generation carrier: span{u,c,t} inside the up-type quark trace",
		Projector:                   "P_Q followed by generation trace Tr_gen(Y_u†Y_u)",
		SignedSpectrumDistinct:      true,
		MagnitudeSpectrumDegenerate: true,
		GenerationBreaking:          true,
		PullbackUnique:              false,
		Verdict:                     strings.Join([]string{StatusTrialityPullbackFormalized, StatusTauEtaDegeneracy, StatusFailedTopSlotNotCanonical}, ";"),
	}
}

func extractFractionalization() Fractionalization {
	weights := []float64{4.0 / 9.0, 4.0 / 9.0, 1.0 / 9.0}
	sum := 0.0
	for _, w := range weights {
		sum += w
	}
	return Fractionalization{
		Formalized:             true,
		Equation:               "Tr_gen(Y_u†Y_u)/g_*² = r_+; |tau_eta|² weights = (4,4,1)/9",
		RPlusExact:             "(3591+136√123)/3099",
		RPlusDecimal:           rPlus(),
		TauMagnitudeSquares:    []float64{4, 4, 1},
		NormalizedWeights:      weights,
		SumWeights:             sum,
		GeneratesUniqueLowSlot: true,
		GeneratesTwoHighSlots:  true,
		Verdict:                strings.Join([]string{StatusAmplitudeFractionalization, StatusTopBoundaryCandidatesAudited}, ";"),
	}
}

func buildCandidates(f Fractionalization) []TopSlotCandidate {
	r := f.RPlusDecimal
	mk := func(name string, idx int, tau int, weight float64, rule string, canonical bool, ambiguous bool, verdictExtra []string) TopSlotCandidate {
		statuses := []string{StatusTopBoundaryCandidatesAudited}
		statuses = append(statuses, verdictExtra...)
		return TopSlotCandidate{Name: name, SlotIndex: idx, TauEtaValue: tau, Weight: weight, TopYtSquaredOverGStarSquared: r * weight, TopYtUV: math.Sqrt(r * weight), SelectionRule: rule, DerivedFromTauEta: true, Canonical: canonical, Ambiguous: ambiguous, Verdict: strings.Join(statuses, ";")}
	}
	out := []TopSlotCandidate{
		mk("tau_eta_positive_high_slot", 0, 2, 4.0/9.0, "select the +2 high-magnitude triality slot as top", false, true, []string{StatusTauEtaDegeneracy, StatusFailedTopSlotNotCanonical}),
		mk("tau_eta_negative_high_slot", 1, -2, 4.0/9.0, "select the -2 high-magnitude triality slot as top", false, true, []string{StatusTauEtaDegeneracy, StatusFailedTopSlotNotCanonical}),
		mk("tau_eta_unique_low_slot", 2, 1, 1.0/9.0, "select the unique |tau|=1 slot as top", false, true, []string{StatusFailedFlavorOrientationMissing, StatusFailedTopSlotNotCanonical}),
		{Name: "gauge_only_zero_top_envelope", SlotIndex: -1, TauEtaValue: 0, Weight: 0, TopYtSquaredOverGStarSquared: 0, TopYtUV: 0, SelectionRule: "diagnostic envelope used by Gate 322; not a tau_eta top assignment", DerivedFromTauEta: false, Canonical: false, Ambiguous: false, Verdict: "CONDITIONAL_DIAGNOSTIC_GAUGE_ONLY_ENVELOPE_NOT_PHYSICAL"},
	}
	return out
}

func runPreflights(candidates []TopSlotCandidate) []TransportPreflight {
	out := make([]TransportPreflight, 0, len(candidates))
	for _, c := range candidates {
		out = append(out, runTransport(c))
	}
	return out
}

func runTransport(c TopSlotCandidate) TransportPreflight {
	start := rgState{ScaleGeV: gutScaleGeV, GY: math.Sqrt(3.0 / 5.0), G2: 1.0, G3: 1.0, YT: c.TopYtUV, Lambda: lambdaUVBoundary}
	highBeta := betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237}
	lowBeta := betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0}
	atThreshold, okHigh := integrateSegment(start, gutScaleGeV, thresholdScaleGeV, highBeta)
	baseline, okBase := integrateSegment(atThreshold, thresholdScaleGeV, vevGeV, lowBeta)
	baselineMass := math.NaN()
	if okHigh && okBase && baseline.Lambda > 0 {
		baselineMass = lambdaToMass(baseline.Lambda)
	}
	post := atThreshold
	post.Lambda += derivedDeltaLambdaGate321
	final, okFinal := integrateSegment(post, thresholdScaleGeV, vevGeV, lowBeta)
	mass := math.NaN()
	if okHigh && okFinal && final.Lambda > 0 {
		mass = lambdaToMass(final.Lambda)
	}
	diff := mass - comparisonPoleMassGeV
	rel := diff / comparisonPoleMassGeV
	near := okHigh && okFinal && final.Lambda > 0 && math.Abs(rel) < 0.01
	verdict := strings.Join([]string{StatusPhysicalLanePreflightExecuted}, ";")
	if near && c.Name == "gauge_only_zero_top_envelope" {
		verdict = strings.Join([]string{StatusPhysicalLanePreflightExecuted, "CONDITIONAL_SUPPORT_GATE322_DIAGNOSTIC_ENVELOPE_REPRODUCED"}, ";")
	} else if !near && c.Weight > 0 {
		verdict = strings.Join([]string{StatusPhysicalLanePreflightExecuted, StatusFailedTopLaneSpoils125}, ";")
	}
	return TransportPreflight{CandidateName: c.Name, TopFraction: c.Weight, TopYtUV: c.TopYtUV, LambdaAtThresholdPlus: atThreshold.Lambda, LambdaAtThresholdMinus: post.Lambda, BaselineLambdaAtV: baseline.Lambda, BaselineMassGeV: baselineMass, FinalLambdaAtV: final.Lambda, RunningMassGeV: mass, TargetMassGeV: comparisonPoleMassGeV, DifferenceGeV: diff, RelativeDifference: rel, Near125WithinOnePct: near, Perturbative: okHigh && okBase && okFinal, VacuumStable: okHigh && okFinal && final.Lambda > 0, Endpoint: endpoint(final, vevGeV), Verdict: verdict}
}

func auditPullbackVerdict(c TrialityCarrier, candidates []TopSlotCandidate, p []TransportPreflight) PullbackVerdict {
	best := findBestNearTarget(p)
	gauge := findPreflight(p, "gauge_only_zero_top_envelope")
	nonzeroSpoils := true
	for _, r := range p {
		if r.TopFraction > 0 && r.Near125WithinOnePct {
			nonzeroSpoils = false
		}
	}
	bestName, bestMass := "", math.NaN()
	if best != nil {
		bestName, bestMass = best.CandidateName, best.RunningMassGeV
	}
	gaugeReq := gauge != nil && gauge.Near125WithinOnePct
	return PullbackVerdict{Formalized: c.Formalized, CanonicalTopFractionDerived: false, PreferredPhysicalCandidate: "none; tau_eta pullback supplies candidates but no unique top eigenvector", BestNear125Candidate: bestName, BestNear125MassGeV: bestMass, GaugeOnlyStillRequired: gaugeReq, NonzeroTauEtaTopSpoils125: nonzeroSpoils, TopBoundaryStatus: StatusFailedNativeTopBoundaryNotDerived, RequiredNextOperator: "derive a flavor-orientation/CKM-like operator that maps tau_eta signs and magnitudes to physical {u,c,t} eigenstates before activating a nonzero top lane", Verdict: strings.Join([]string{StatusTrialityPullbackFormalized, StatusTopBoundaryCandidatesAudited, StatusThresholdSuccessRequiresTopMute, StatusFailedNativeTopBoundaryNotDerived, StatusFailedFlavorOrientationMissing}, ";")}
}

func auditFirewalls(v PullbackVerdict) FirewallAudit {
	return FirewallAudit{NoObservedTopMassInserted: true, NoCKMImported: true, NoFlavorTextureInvented: true, NoPoleMassClaimed: true, NoTwoLoopClaimed: true, NoFinalColliderMassClaimed: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedCKMTextureMissing, StatusFailedPoleMassNotExecuted, StatusFailedTwoLoopNotExecuted, StatusFailedColliderMassNotClaimed}, ";")}
}

func buildSummary(c TrialityCarrier, f Fractionalization, candidates []TopSlotCandidate, p []TransportPreflight, v PullbackVerdict, fw FirewallAudit) Summary {
	preserved := fw.NoObservedTopMassInserted && fw.NoCKMImported && fw.NoFlavorTextureInvented && fw.NoPoleMassClaimed && fw.NoTwoLoopClaimed && fw.NoFinalColliderMassClaimed && !fw.FiniteCorePolluted
	hi := findPreflight(p, "tau_eta_positive_high_slot")
	low := findPreflight(p, "tau_eta_unique_low_slot")
	gauge := findPreflight(p, "gauge_only_zero_top_envelope")
	direct := "tau_eta pullback gives weights (4/9,4/9,1/9), but no unique physical top slot; nonzero tau_eta top candidates do not preserve the Gate-322 near-125 GeV lane."
	if hi != nil && low != nil && gauge != nil {
		direct = fmt.Sprintf("tau_eta pullback gives weights (4/9,4/9,1/9): high-slot top with the derived jump gives %.3f GeV, unique-low top gives %.3f GeV, while the diagnostic zero-top envelope gives %.3f GeV.  Therefore the native top boundary is not yet derived; the successful Gate-322 transport still relies on the flattened-top envelope.", hi.RunningMassGeV, low.RunningMassGeV, gauge.RunningMassGeV)
	}
	return Summary{TrialityPullbackFormalized: c.Formalized, FractionalizationExtracted: f.Formalized && math.Abs(f.SumWeights-1) < 1e-12, TopCandidatesAudited: len(candidates) >= 4, NativeTopBoundaryDerived: v.CanonicalTopFractionDerived, PhysicalPreflightExecuted: len(p) == len(candidates), Gate322SuccessPreservedByCanonicalTop: false, FirewallsPreserved: preserved, FinalMassClaimed: false, Status: strings.Join([]string{StatusTrialityPullbackFormalized, StatusTopBoundaryCandidatesAudited, StatusFailedNativeTopBoundaryNotDerived}, ";"), DirectAnswer: direct, NextGate: "derive the missing flavor-orientation operator mapping tau_eta to physical u/c/t eigenvectors, or keep the top lane firewalled while upgrading two-loop and pole-mass precision."}
}

func findBestNearTarget(results []TransportPreflight) *TransportPreflight {
	if len(results) == 0 {
		return nil
	}
	bestIdx := 0
	bestAbs := math.Abs(results[0].DifferenceGeV)
	for i := 1; i < len(results); i++ {
		d := math.Abs(results[i].DifferenceGeV)
		if d < bestAbs {
			bestIdx, bestAbs = i, d
		}
	}
	return &results[bestIdx]
}

func findPreflight(results []TransportPreflight, name string) *TransportPreflight {
	for i := range results {
		if results[i].CandidateName == name {
			return &results[i]
		}
	}
	return nil
}

func rPlus() float64                      { return (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0 }
func lambdaToMass(lambda float64) float64 { return vevGeV * math.Sqrt(2.0*lambda) }

// one-loop transport copied intentionally from Gate 322 so the preflight compares
// like with like against the derived-threshold result.
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
	state.ScaleGeV = low
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

func FormatCarrier(c TrialityCarrier) string {
	return fmt.Sprintf("formalized=%t source=%s tau=%v carrier=%s signedDistinct=%t magDegenerate=%t unique=%t verdict=%s", c.Formalized, c.SourceGate, c.TauEta, c.PullbackCarrier, c.SignedSpectrumDistinct, c.MagnitudeSpectrumDegenerate, c.PullbackUnique, c.Verdict)
}

func FormatFractionalization(f Fractionalization) string {
	return fmt.Sprintf("formalized=%t rPlus=%.12f weights=%v sum=%.12f uniqueLow=%t twoHigh=%t verdict=%s", f.Formalized, f.RPlusDecimal, f.NormalizedWeights, f.SumWeights, f.GeneratesUniqueLowSlot, f.GeneratesTwoHighSlots, f.Verdict)
}

func FormatCandidate(c TopSlotCandidate) string {
	return fmt.Sprintf("name=%s slot=%d tau=%d weight=%.12f yt2/g2=%.12f yt=%.12f canonical=%t ambiguous=%t verdict=%s", c.Name, c.SlotIndex, c.TauEtaValue, c.Weight, c.TopYtSquaredOverGStarSquared, c.TopYtUV, c.Canonical, c.Ambiguous, c.Verdict)
}

func FormatPreflight(p TransportPreflight) string {
	return fmt.Sprintf("candidate=%s frac=%.12f yt=%.12f lambdaPlus=%.12f lambdaMinus=%.12f baseMass=%.6f finalLambda=%.12f runMass=%.6f diff=%+.6f rel=%+.6f%% near=%t perturb=%t stable=%t verdict=%s", p.CandidateName, p.TopFraction, p.TopYtUV, p.LambdaAtThresholdPlus, p.LambdaAtThresholdMinus, p.BaselineMassGeV, p.FinalLambdaAtV, p.RunningMassGeV, p.DifferenceGeV, 100*p.RelativeDifference, p.Near125WithinOnePct, p.Perturbative, p.VacuumStable, p.Verdict)
}

func FormatPullbackVerdict(v PullbackVerdict) string {
	return fmt.Sprintf("formalized=%t canonicalTop=%t preferred=%s best=%s bestMass=%.6f gaugeOnlyRequired=%t nonzeroSpoils=%t status=%s next=%s verdict=%s", v.Formalized, v.CanonicalTopFractionDerived, v.PreferredPhysicalCandidate, v.BestNear125Candidate, v.BestNear125MassGeV, v.GaugeOnlyStillRequired, v.NonzeroTauEtaTopSpoils125, v.TopBoundaryStatus, v.RequiredNextOperator, v.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noTopMass=%t noCKM=%t noTexture=%t noPole=%t noTwoLoop=%t noCollider=%t polluted=%t verdict=%s", f.NoObservedTopMassInserted, f.NoCKMImported, f.NoFlavorTextureInvented, f.NoPoleMassClaimed, f.NoTwoLoopClaimed, f.NoFinalColliderMassClaimed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("pullback=%t frac=%t candidates=%t nativeTop=%t preflight=%t gate322Preserved=%t firewall=%t finalClaim=%t status=%s answer=%s next=%s", s.TrialityPullbackFormalized, s.FractionalizationExtracted, s.TopCandidatesAudited, s.NativeTopBoundaryDerived, s.PhysicalPreflightExecuted, s.Gate322SuccessPreservedByCanonicalTop, s.FirewallsPreserved, s.FinalMassClaimed, s.Status, s.DirectAnswer, s.NextGate)
}
