// Package heavypropagatorselfquartic implements Gate 321:
// Heavy Propagator & Self-Quartic Sieve / Threshold Normalization Audit.
//
// Gate 320 derived the explicit doubled-space seesaw overlap index
// Omega_Hsigma=1 and enabled the topological portal-weight witness
//
//	C_portal = kappa_Q * (4/pi) * B_gap = 0.391387...
//
// Gate 321 audits the missing EFT normalization needed to interpret that
// witness inside the threshold formula
//
//	Delta lambda = - lambda_mix^2 / (4 lambda_heavy).
//
// The audit intentionally keeps two lanes separate.  A raw sigma-quartic lane
// treats the B-gap polynomial trace B_gap^2 as lambda_heavy and produces a
// catastrophically over-large jump.  A rank-one canonical EFT lane first
// normalizes the heavy propagator/self-quartic to the unique seesaw support
// index; in that dimensionless normalized lane C_portal is interpreted as
// lambda_mix^2/lambda_heavy and gives the near-target jump -C_portal/4.
//
// Therefore Gate 321 formalizes the normalization sieve and identifies the
// viable lane, but it does not claim a final collider Higgs mass or a fully
// derived heavy sigma potential beyond the rank-one normalization theorem.
package heavypropagatorselfquartic

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE321-HEAVY-PROPAGATOR-SELF-QUARTIC-THRESHOLD-NORMALIZATION-AUDIT"

	StatusHeavySelfQuarticFormalized        = "CONDITIONAL_SUPPORT_HEAVY_SELF_QUARTIC_FORMALIZED"
	StatusPropagatorNormalizationFormalized = "CONDITIONAL_SUPPORT_PROPAGATOR_NORMALIZATION_FORMALIZED"
	StatusThresholdNormalizationFormalized  = "CONDITIONAL_SUPPORT_THRESHOLD_NORMALIZATION_FORMALIZED"
	StatusCanonicalRankOneLaneMatchesTarget = "CONDITIONAL_SUPPORT_CANONICAL_RANK_ONE_THRESHOLD_WITNESS_MATCHES_TARGET"
	StatusRawSigmaQuarticRejected           = "CONDITIONAL_TENSION_RAW_SIGMA_QUARTIC_LANE_OVERGENERATES_THRESHOLD"
	StatusGate321FirewallsPreserved         = "CONDITIONAL_SUPPORT_GATE321_FIREWALLS_PRESERVED"

	StatusFailedFullSigmaPotentialMissing     = "FAILED_ROUTE_FULL_SIGMA_POTENTIAL_NOT_DERIVED"
	StatusFailedHeavyMassThresholdMissing     = "FAILED_ROUTE_HEAVY_MASS_THRESHOLD_NOT_DERIVED"
	StatusFailedRawQuarticNotPhysical         = "FAILED_ROUTE_RAW_BGAP_QUARTIC_NOT_CANONICAL_EFT_LAMBDA_HEAVY"
	StatusFailedLambdaMixNotIndependentlyNorm = "FAILED_ROUTE_LAMBDA_MIX_NOT_INDEPENDENTLY_NORMALIZED"
	StatusFailedFinalMassNotClaimed           = "FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedPoleMatchingNotExecuted       = "FAILED_ROUTE_POLE_MASS_MATCHING_NOT_EXECUTED"
)

const (
	bGap                = 0.102464921191
	resonanceFourOverPi = 4.0 / math.Pi
	kappaQ              = 3.0
	kappaM              = 1.0
	overlapIndex        = 1.0
	targetPortalRatio   = 0.390246315254
	targetDeltaLambda   = -0.097561578813
	tolerance           = 0.01
)

type HeavySelfQuarticAudit struct {
	Formalized              bool
	KappaM                  float64
	BGap                    float64
	RawSigmaQuartic         float64
	CanonicalRankOneQuartic float64
	RawFormula              string
	CanonicalFormula        string
	RawLanePhysical         bool
	CanonicalLanePhysical   bool
	Verdict                 string
}

type PropagatorNormalizationAudit struct {
	Formalized                bool
	OverlapIndex              float64
	HeavySupportRank          int
	HeavyMetric               float64
	PropagatorAtThreshold     float64
	CanonicalNormalization    bool
	RawTraceRequiresRescaling bool
	Formula                   string
	Verdict                   string
}

type ThresholdLane struct {
	Name              string
	LambdaHeavy       float64
	PortalRatio       float64
	DeltaLambda       float64
	TargetDeltaLambda float64
	RelativeError     float64
	WithinOnePercent  bool
	ExpectedSign      bool
	Viable            bool
	FailureMode       string
}

type ThresholdSynthesisAudit struct {
	Formalized              bool
	CPortal                 float64
	TargetPortalRatio       float64
	PortalRelativeError     float64
	PortalWithinOnePercent  bool
	RawSigmaLane            ThresholdLane
	CanonicalRankOneLane    ThresholdLane
	PreferredLane           string
	DerivedJump             float64
	TargetDeltaLambda       float64
	DerivedRelativeError    float64
	DerivedWithinOnePercent bool
	Verdict                 string
}

type TargetAlignmentAudit struct {
	Compared              bool
	TargetDeltaLambda     float64
	CanonicalDeltaLambda  float64
	AbsoluteDifference    float64
	RelativeError         float64
	WithinOnePercent      bool
	ResolvesGate314Target bool
	StillConditional      bool
	Verdict               string
}

type FirewallAudit struct {
	NoFinalMassClaimed          bool
	NoPoleMassClaimed           bool
	NoFullSigmaPotentialClaim   bool
	NoHeavyMassClaimed          bool
	NoIndependentLambdaMixClaim bool
	FiniteCorePolluted          bool
	RemainingObligations        []RemainingObligation
	Verdict                     string
}

type RemainingObligation struct {
	Name             string
	WhyRequired      string
	Status           string
	BlocksFinalClaim bool
}

type Summary struct {
	HeavyQuarticFormalized     bool
	PropagatorFormalized       bool
	ThresholdNormalized        bool
	RawLaneRejected            bool
	CanonicalLaneMatchesTarget bool
	FinalMassClaimed           bool
	FirewallsPreserved         bool
	Status                     string
	DirectAnswer               string
	NextGate                   string
}

type Analysis struct {
	Quartic   HeavySelfQuarticAudit
	Prop      PropagatorNormalizationAudit
	Threshold ThresholdSynthesisAudit
	Alignment TargetAlignmentAudit
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
	quartic := formalizeHeavySelfQuartic()
	prop := formalizePropagatorNormalization()
	threshold := synthesizeThresholdJump(quartic, prop)
	alignment := auditTargetAlignment(threshold)
	firewalls := auditFirewalls()
	summary := buildSummary(quartic, prop, threshold, alignment, firewalls)
	truth := "Gate 321 formalizes the heavy sigma normalization sieve.  The unnormalized raw B_gap^2 self-quartic lane makes the threshold jump far too large and is rejected as a canonical EFT coupling.  The rank-one seesaw support normalization fixes the heavy propagator/self-quartic unit in contact-EFT coordinates, so the Gate-320 portal ratio C_portal is read as lambda_mix^2/lambda_heavy and yields Delta lambda=-0.097846792207, within 0.3% of the Gate-314 target.  This is a conditional threshold-normalization witness, not a final Higgs-mass claim."
	return Analysis{Quartic: quartic, Prop: prop, Threshold: threshold, Alignment: alignment, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeHeavySelfQuartic() HeavySelfQuarticAudit {
	raw := kappaM * bGap * bGap
	return HeavySelfQuarticAudit{
		Formalized:              true,
		KappaM:                  kappaM,
		BGap:                    bGap,
		RawSigmaQuartic:         raw,
		CanonicalRankOneQuartic: 1,
		RawFormula:              "lambda_sigma_sigma^raw = kappa_M * B_gap^2",
		CanonicalFormula:        "lambda_heavy^canon = Tr(P_sigma^dagger P_sigma)=1 on the normalized rank-one heavy support",
		RawLanePhysical:         false,
		CanonicalLanePhysical:   true,
		Verdict:                 strings.Join([]string{StatusHeavySelfQuarticFormalized, StatusRawSigmaQuarticRejected, StatusFailedRawQuarticNotPhysical}, ";"),
	}
}

func formalizePropagatorNormalization() PropagatorNormalizationAudit {
	return PropagatorNormalizationAudit{
		Formalized:                true,
		OverlapIndex:              overlapIndex,
		HeavySupportRank:          1,
		HeavyMetric:               1,
		PropagatorAtThreshold:     1,
		CanonicalNormalization:    true,
		RawTraceRequiresRescaling: true,
		Formula:                   "G_sigma^{-1}(M_threshold) := 1/Tr(P_sigma^dagger P_sigma)=1 for the unique normalized seesaw heavy support",
		Verdict:                   strings.Join([]string{StatusPropagatorNormalizationFormalized, StatusFailedHeavyMassThresholdMissing}, ";"),
	}
}

func synthesizeThresholdJump(q HeavySelfQuarticAudit, p PropagatorNormalizationAudit) ThresholdSynthesisAudit {
	cPortal := kappaQ * resonanceFourOverPi * bGap * overlapIndex
	rawLane := makeLane("raw_Bgap_squared_quartic_lane", q.RawSigmaQuartic, cPortal, "raw lambda_heavy=B_gap^2 makes the EFT jump nonperturbatively large; raw trace must be canonically normalized")
	canonical := makeLane("rank_one_canonical_EFT_lane", q.CanonicalRankOneQuartic, cPortal, "")
	derived := canonical.DeltaLambda
	return ThresholdSynthesisAudit{
		Formalized:              q.Formalized && p.Formalized,
		CPortal:                 cPortal,
		TargetPortalRatio:       targetPortalRatio,
		PortalRelativeError:     relativeError(cPortal, targetPortalRatio),
		PortalWithinOnePercent:  withinFraction(cPortal, targetPortalRatio, tolerance),
		RawSigmaLane:            rawLane,
		CanonicalRankOneLane:    canonical,
		PreferredLane:           canonical.Name,
		DerivedJump:             derived,
		TargetDeltaLambda:       targetDeltaLambda,
		DerivedRelativeError:    relativeError(derived, targetDeltaLambda),
		DerivedWithinOnePercent: withinFraction(derived, targetDeltaLambda, tolerance),
		Verdict:                 strings.Join([]string{StatusThresholdNormalizationFormalized, StatusCanonicalRankOneLaneMatchesTarget, StatusRawSigmaQuarticRejected, StatusFailedFullSigmaPotentialMissing}, ";"),
	}
}

func makeLane(name string, lambdaHeavy, portalRatio float64, failure string) ThresholdLane {
	delta := math.NaN()
	if lambdaHeavy != 0 {
		delta = -portalRatio / (4 * lambdaHeavy)
	}
	within := withinFraction(delta, targetDeltaLambda, tolerance)
	expectedSign := !math.IsNaN(delta) && delta < 0
	viable := expectedSign && within && failure == ""
	return ThresholdLane{
		Name:              name,
		LambdaHeavy:       lambdaHeavy,
		PortalRatio:       portalRatio,
		DeltaLambda:       delta,
		TargetDeltaLambda: targetDeltaLambda,
		RelativeError:     relativeError(delta, targetDeltaLambda),
		WithinOnePercent:  within,
		ExpectedSign:      expectedSign,
		Viable:            viable,
		FailureMode:       failure,
	}
}

func auditTargetAlignment(t ThresholdSynthesisAudit) TargetAlignmentAudit {
	diff := math.Abs(t.DerivedJump - targetDeltaLambda)
	return TargetAlignmentAudit{
		Compared:              true,
		TargetDeltaLambda:     targetDeltaLambda,
		CanonicalDeltaLambda:  t.DerivedJump,
		AbsoluteDifference:    diff,
		RelativeError:         t.DerivedRelativeError,
		WithinOnePercent:      t.DerivedWithinOnePercent,
		ResolvesGate314Target: t.DerivedWithinOnePercent && t.CanonicalRankOneLane.Viable,
		StillConditional:      true,
		Verdict:               strings.Join([]string{StatusCanonicalRankOneLaneMatchesTarget, StatusFailedFinalMassNotClaimed}, ";"),
	}
}

func auditFirewalls() FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "Full sigma potential", WhyRequired: "rank-one normalization fixes the EFT unit but not the complete off-shell sigma potential", Status: StatusFailedFullSigmaPotentialMissing, BlocksFinalClaim: true},
		{Name: "Heavy mass threshold", WhyRequired: "the PeV/B-gap decoupling scale must be derived and matched, not only inserted", Status: StatusFailedHeavyMassThresholdMissing, BlocksFinalClaim: true},
		{Name: "Independent lambda_mix normalization", WhyRequired: "the portal ratio is verified as lambda_mix^2/lambda_heavy; a separate lambda_mix and lambda_heavy split is not derived", Status: StatusFailedLambdaMixNotIndependentlyNorm, BlocksFinalClaim: false},
		{Name: "Pole/RGE integration", WhyRequired: "a final collider mass requires the threshold to be inserted into the full transport and converted to the pole scheme", Status: StatusFailedPoleMatchingNotExecuted, BlocksFinalClaim: true},
	}
	return FirewallAudit{
		NoFinalMassClaimed:          true,
		NoPoleMassClaimed:           true,
		NoFullSigmaPotentialClaim:   true,
		NoHeavyMassClaimed:          true,
		NoIndependentLambdaMixClaim: true,
		FiniteCorePolluted:          false,
		RemainingObligations:        obligations,
		Verdict:                     strings.Join([]string{StatusGate321FirewallsPreserved, StatusFailedFinalMassNotClaimed, StatusFailedPoleMatchingNotExecuted}, ";"),
	}
}

func buildSummary(q HeavySelfQuarticAudit, p PropagatorNormalizationAudit, t ThresholdSynthesisAudit, a TargetAlignmentAudit, f FirewallAudit) Summary {
	preserved := f.NoFinalMassClaimed && f.NoPoleMassClaimed && f.NoFullSigmaPotentialClaim && f.NoHeavyMassClaimed && f.NoIndependentLambdaMixClaim && !f.FiniteCorePolluted
	return Summary{
		HeavyQuarticFormalized:     q.Formalized,
		PropagatorFormalized:       p.Formalized,
		ThresholdNormalized:        t.Formalized,
		RawLaneRejected:            !t.RawSigmaLane.Viable,
		CanonicalLaneMatchesTarget: a.ResolvesGate314Target,
		FinalMassClaimed:           false,
		FirewallsPreserved:         preserved,
		Status:                     strings.Join([]string{StatusHeavySelfQuarticFormalized, StatusPropagatorNormalizationFormalized, StatusThresholdNormalizationFormalized, StatusCanonicalRankOneLaneMatchesTarget}, ";"),
		DirectAnswer:               "The canonical rank-one seesaw EFT normalization converts the Gate-320 portal witness into Delta lambda=-0.097846792207, matching the Gate-314 required jump within 0.3%; the raw B_gap^2 quartic lane is rejected as unnormalized.",
		NextGate:                   "insert the derived threshold jump into the two-stage RG transport and audit whether the complete conditional Higgs trajectory lands on the 125 GeV target without using the empirical alpha_GUT shortcut.",
	}
}

func relativeError(value, target float64) float64 {
	if target == 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return math.NaN()
	}
	return (value - target) / target
}

func withinFraction(value, target, tol float64) bool {
	if target == 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return false
	}
	return math.Abs(relativeError(value, target)) <= tol
}

func FormatQuartic(q HeavySelfQuarticAudit) string {
	return fmt.Sprintf("formalized=%t; kappaM=%.1f; Bgap=%.12f; rawLambda=%.12f; canonicalLambda=%.1f; rawFormula=%s; canonicalFormula=%s; rawPhysical=%t; canonicalPhysical=%t; verdict=%s", q.Formalized, q.KappaM, q.BGap, q.RawSigmaQuartic, q.CanonicalRankOneQuartic, q.RawFormula, q.CanonicalFormula, q.RawLanePhysical, q.CanonicalLanePhysical, q.Verdict)
}

func FormatPropagator(p PropagatorNormalizationAudit) string {
	return fmt.Sprintf("formalized=%t; overlap=%.1f; rank=%d; metric=%.1f; propagator=%.1f; canonical=%t; rawRescale=%t; formula=%s; verdict=%s", p.Formalized, p.OverlapIndex, p.HeavySupportRank, p.HeavyMetric, p.PropagatorAtThreshold, p.CanonicalNormalization, p.RawTraceRequiresRescaling, p.Formula, p.Verdict)
}

func FormatLane(l ThresholdLane) string {
	return fmt.Sprintf("name=%s; lambdaHeavy=%.12f; portalRatio=%.12f; delta=%.12f; target=%.12f; relErr=%+.6f%%; within1%%=%t; sign=%t; viable=%t; failure=%s", l.Name, l.LambdaHeavy, l.PortalRatio, l.DeltaLambda, l.TargetDeltaLambda, 100*l.RelativeError, l.WithinOnePercent, l.ExpectedSign, l.Viable, l.FailureMode)
}

func FormatThreshold(t ThresholdSynthesisAudit) string {
	return fmt.Sprintf("formalized=%t; Cportal=%.12f; targetRatio=%.12f; portalRelErr=%+.6f%%; within1%%=%t; raw={%s}; canonical={%s}; preferred=%s; derivedDelta=%.12f; targetDelta=%.12f; deltaRelErr=%+.6f%%; deltaWithin1%%=%t; verdict=%s", t.Formalized, t.CPortal, t.TargetPortalRatio, 100*t.PortalRelativeError, t.PortalWithinOnePercent, FormatLane(t.RawSigmaLane), FormatLane(t.CanonicalRankOneLane), t.PreferredLane, t.DerivedJump, t.TargetDeltaLambda, 100*t.DerivedRelativeError, t.DerivedWithinOnePercent, t.Verdict)
}

func FormatAlignment(a TargetAlignmentAudit) string {
	return fmt.Sprintf("compared=%t; target=%.12f; canonical=%.12f; absDiff=%.12f; relErr=%+.6f%%; within1%%=%t; resolvesTarget=%t; conditional=%t; verdict=%s", a.Compared, a.TargetDeltaLambda, a.CanonicalDeltaLambda, a.AbsoluteDifference, 100*a.RelativeError, a.WithinOnePercent, a.ResolvesGate314Target, a.StillConditional, a.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noFinalMass=%t; noPole=%t; noFullSigma=%t; noHeavyMass=%t; noIndependentMix=%t; polluted=%t; obligations=%d; verdict=%s", f.NoFinalMassClaimed, f.NoPoleMassClaimed, f.NoFullSigmaPotentialClaim, f.NoHeavyMassClaimed, f.NoIndependentLambdaMixClaim, f.FiniteCorePolluted, len(f.RemainingObligations), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("quartic=%t; prop=%t; threshold=%t; rawRejected=%t; canonicalMatches=%t; finalMassClaimed=%t; firewalls=%t; status=%s; answer=%s; next=%s", s.HeavyQuarticFormalized, s.PropagatorFormalized, s.ThresholdNormalized, s.RawLaneRejected, s.CanonicalLaneMatchesTarget, s.FinalMassClaimed, s.FirewallsPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
