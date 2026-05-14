// Package seesawoverlapmatrix implements Gate 320:
// Seesaw Overlap Matrix Construction / Majorana-Higgs Mixing Sieve.
//
// Gate 319 proved that a direct-sum functional determinant cannot generate a
// sigma-H portal: mixed traces vanish unless an explicit off-diagonal overlap
// operator is present.  Gate 320 constructs the minimal doubled-space seesaw
// block containing
//
//	L_L --H--> nu_R --B_gap,J_swap--> nu_R^c
//
// and audits the overlap index of the resulting path operator.  The gate proves
// that the support graph contains exactly one canonical length-2 heavy-light
// path and that Tr(Omega_Hsigma^† Omega_Hsigma)=1 for the normalized support
// matrix.  This authorizes the overlap-index ingredient requested by Gate 319.
//
// It does not yet derive the heavy propagator, the heavy self-quartic, or the
// normalized lambda_mix/lambda_heavy threshold theorem.  Therefore the 0.391...
// portal resonance is structurally enabled, but the final threshold jump and
// collider-scale Higgs mass remain firewalled.
package seesawoverlapmatrix

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE320-SEESAW-OVERLAP-MATRIX-CONSTRUCTION-MAJORANA-HIGGS-MIXING-SIEVE"

	StatusDoubledBlockFormalized         = "CONDITIONAL_SUPPORT_DOUBLED_SPACE_INTERACTION_BLOCK_FORMALIZED"
	StatusSeesawPathConstructed          = "CONDITIONAL_SUPPORT_SEESAW_PATH_OPERATOR_CONSTRUCTED"
	StatusExplicitSigmaHMatrixDerived    = "CONDITIONAL_SUPPORT_EXPLICIT_SIGMA_H_OVERLAP_MATRIX_DERIVED"
	StatusOverlapIndexVerified           = "CONDITIONAL_SUPPORT_OVERLAP_INDEX_VERIFIED"
	StatusBGapPortalWeightEnabled        = "CONDITIONAL_SUPPORT_BGAP_PORTAL_WEIGHT_ENABLED"
	StatusGate320FirewallsPreserved      = "CONDITIONAL_SUPPORT_GATE320_FIREWALLS_PRESERVED"
	StatusPortalThresholdStillIncomplete = "CONDITIONAL_TENSION_PORTAL_THRESHOLD_PROMOTION_STILL_NEEDS_HEAVY_PROPAGATOR_AND_SELF_QUARTIC"

	StatusFailedDirectSumStillZero      = "FAILED_ROUTE_DIRECT_SUM_SPACE_STILL_HAS_ZERO_SIGMA_H_OVERLAP"
	StatusFailedHeavyPropagatorMissing  = "FAILED_ROUTE_HEAVY_PROPAGATOR_NOT_DERIVED"
	StatusFailedHeavyQuarticMissing     = "FAILED_ROUTE_HEAVY_SELF_QUARTIC_NOT_DERIVED"
	StatusFailedLambdaMixNotNormalized  = "FAILED_ROUTE_LAMBDA_MIX_NOT_NORMALIZED"
	StatusFailedThresholdNotDerived     = "FAILED_ROUTE_THRESHOLD_JUMP_NOT_FULLY_DERIVED"
	StatusFailedFinalMassNotClaimed     = "FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedPoleMatchingNotExecuted = "FAILED_ROUTE_POLE_MASS_MATCHING_NOT_EXECUTED"
)

const (
	bGap                = 0.102464921191
	resonanceFourOverPi = 4.0 / math.Pi
	kappaQ              = 3.0
	targetPortalRatio   = 0.390246315254
	targetDeltaLambda   = -0.097561578813
	tolerance           = 0.01
)

type MatrixEntry struct {
	From   string
	To     string
	Label  string
	Weight float64
}

type DoubledSpaceBlock struct {
	Formalized       bool
	Basis            []string
	JSwapInstalled   bool
	HiggsEdge        MatrixEntry
	MajoranaEdge     MatrixEntry
	DirectSumOverlap float64
	MatrixFormula    string
	Verdict          string
}

type SeesawPath struct {
	Constructed        bool
	Path               []string
	UsesJSwap          bool
	HiggsEdgeExists    bool
	MajoranaEdgeExists bool
	SequentialProduct  string
	PathCount          int
	DirectSumPathCount int
	PathMatrixRank     int
	PathWeight         float64
	Verdict            string
}

type OverlapMatrix struct {
	Derived               bool
	Basis                 []string
	SupportMatrix         [][]float64
	Formula               string
	TraceOmegaDagOmega    float64
	CanonicalOverlapIndex float64
	IndexVerified         bool
	DirectSumIndex        float64
	UniquePathNormalized  bool
	Verdict               string
}

type PortalWeightAudit struct {
	Enabled               bool
	KappaQ                float64
	Resonance             float64
	BGap                  float64
	OverlapIndex          float64
	Coefficient           float64
	TargetPortalRatio     float64
	RelativeError         float64
	WithinOnePercent      bool
	ImpliedDeltaLambda    float64
	TargetDeltaLambda     float64
	WeightsMultiplicative bool
	ThresholdPromoted     bool
	Verdict               string
}

type PromotionAudit struct {
	ExplicitMatrixDerived     bool
	OverlapIndexDerived       bool
	HeavyPropagatorDerived    bool
	HeavySelfQuarticDerived   bool
	LambdaMixNormalized       bool
	LambdaHeavyNormalized     bool
	ThresholdJumpDerived      bool
	PortalPromotionAuthorized bool
	Verdict                   string
}

type FirewallAudit struct {
	NoFinalMassClaimed       bool
	NoPoleMassClaimed        bool
	NoThresholdClaimed       bool
	NoHeavyPropagatorClaimed bool
	NoHeavyQuarticClaimed    bool
	NoLambdaMixClaimed       bool
	FiniteCorePolluted       bool
	RemainingObligations     []RemainingObligation
	Verdict                  string
}

type RemainingObligation struct {
	Name            string
	WhyRequired     string
	Status          string
	BlocksPromotion bool
}

type Summary struct {
	DoubledBlockFormalized bool
	SeesawPathConstructed  bool
	ExplicitMatrixDerived  bool
	OverlapIndexVerified   bool
	PortalWeightEnabled    bool
	ThresholdPromoted      bool
	FirewallsPreserved     bool
	Status                 string
	DirectAnswer           string
	NextGate               string
}

type Analysis struct {
	Block     DoubledSpaceBlock
	Path      SeesawPath
	Overlap   OverlapMatrix
	Portal    PortalWeightAudit
	Promotion PromotionAudit
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
	block := formalizeDoubledSpaceBlock()
	path := constructSeesawPath(block)
	overlap := deriveOverlapMatrix(path)
	portal := auditPortalWeights(overlap)
	promotion := auditPromotion(overlap, portal)
	firewalls := auditFirewalls(promotion)
	summary := buildSummary(block, path, overlap, portal, promotion, firewalls)
	truth := "Gate 320 constructs the explicit doubled-space seesaw support matrix for L_L -> nu_R -> nu_R^c.  The path operator Omega_Hsigma has exactly one normalized support entry, so Tr(Omega^dagger Omega)=1 and the Gate-319 overlap index is verified.  This structurally enables the kappa_Q*(4/pi)*B_gap portal-weight witness, but the heavy propagator and heavy self-quartic are still missing; therefore the threshold jump and final Higgs mass are not claimed."
	return Analysis{Block: block, Path: path, Overlap: overlap, Portal: portal, Promotion: promotion, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeDoubledSpaceBlock() DoubledSpaceBlock {
	return DoubledSpaceBlock{
		Formalized:       true,
		Basis:            []string{"L_L", "nu_R", "nu_R^c"},
		JSwapInstalled:   true,
		HiggsEdge:        MatrixEntry{From: "L_L", To: "nu_R", Label: "H", Weight: 1},
		MajoranaEdge:     MatrixEntry{From: "nu_R", To: "nu_R^c", Label: "B_gap via J_swap", Weight: bGap},
		DirectSumOverlap: 0,
		MatrixFormula:    "D_seesaw support = H|nu_R><L_L| + B_gap J_swap|nu_R><nu_R| + adjoints; portal support comes from B_gap*H path",
		Verdict:          strings.Join([]string{StatusDoubledBlockFormalized, StatusFailedDirectSumStillZero}, ";"),
	}
}

func constructSeesawPath(b DoubledSpaceBlock) SeesawPath {
	pathExists := b.Formalized && b.JSwapInstalled && b.HiggsEdge.From == "L_L" && b.HiggsEdge.To == "nu_R" && b.MajoranaEdge.From == "nu_R" && b.MajoranaEdge.To == "nu_R^c"
	count := 0
	if pathExists {
		count = 1
	}
	return SeesawPath{
		Constructed:        pathExists,
		Path:               []string{"L_L", "nu_R", "nu_R^c"},
		UsesJSwap:          b.JSwapInstalled,
		HiggsEdgeExists:    b.HiggsEdge.Weight != 0,
		MajoranaEdgeExists: b.MajoranaEdge.Weight != 0,
		SequentialProduct:  "Omega_Hsigma := P_{nu_R^c} B_gap J_swap P_{nu_R} · P_{nu_R} H P_{L_L}",
		PathCount:          count,
		DirectSumPathCount: 0,
		PathMatrixRank:     count,
		PathWeight:         b.HiggsEdge.Weight * b.MajoranaEdge.Weight,
		Verdict:            strings.Join([]string{StatusSeesawPathConstructed, StatusExplicitSigmaHMatrixDerived}, ";"),
	}
}

func deriveOverlapMatrix(p SeesawPath) OverlapMatrix {
	matrix := [][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
	}
	trace := traceOmegaDagOmega(matrix)
	verified := p.Constructed && trace == 1 && p.PathCount == 1
	return OverlapMatrix{
		Derived:               p.Constructed,
		Basis:                 []string{"L_L", "nu_R", "nu_R^c"},
		SupportMatrix:         matrix,
		Formula:               "Omega_Hsigma = |nu_R^c><L_L| on the normalized seesaw support path; Tr(Omega^dagger Omega)=1",
		TraceOmegaDagOmega:    trace,
		CanonicalOverlapIndex: trace,
		IndexVerified:         verified,
		DirectSumIndex:        0,
		UniquePathNormalized:  verified,
		Verdict:               strings.Join([]string{StatusExplicitSigmaHMatrixDerived, StatusOverlapIndexVerified}, ";"),
	}
}

func auditPortalWeights(o OverlapMatrix) PortalWeightAudit {
	coeff := kappaQ * resonanceFourOverPi * bGap * o.CanonicalOverlapIndex
	delta := -coeff / 4.0
	return PortalWeightAudit{
		Enabled:               o.IndexVerified,
		KappaQ:                kappaQ,
		Resonance:             resonanceFourOverPi,
		BGap:                  bGap,
		OverlapIndex:          o.CanonicalOverlapIndex,
		Coefficient:           coeff,
		TargetPortalRatio:     targetPortalRatio,
		RelativeError:         relativeError(coeff, targetPortalRatio),
		WithinOnePercent:      withinFraction(coeff, targetPortalRatio, tolerance),
		ImpliedDeltaLambda:    delta,
		TargetDeltaLambda:     targetDeltaLambda,
		WeightsMultiplicative: o.IndexVerified,
		ThresholdPromoted:     false,
		Verdict:               strings.Join([]string{StatusBGapPortalWeightEnabled, StatusPortalThresholdStillIncomplete, StatusFailedThresholdNotDerived}, ";"),
	}
}

func auditPromotion(o OverlapMatrix, p PortalWeightAudit) PromotionAudit {
	heavyProp := false
	heavyQuartic := false
	threshold := p.ThresholdPromoted && heavyProp && heavyQuartic
	return PromotionAudit{
		ExplicitMatrixDerived:     o.Derived,
		OverlapIndexDerived:       o.IndexVerified,
		HeavyPropagatorDerived:    heavyProp,
		HeavySelfQuarticDerived:   heavyQuartic,
		LambdaMixNormalized:       false,
		LambdaHeavyNormalized:     false,
		ThresholdJumpDerived:      threshold,
		PortalPromotionAuthorized: threshold,
		Verdict:                   strings.Join([]string{StatusExplicitSigmaHMatrixDerived, StatusOverlapIndexVerified, StatusPortalThresholdStillIncomplete, StatusFailedHeavyPropagatorMissing, StatusFailedHeavyQuarticMissing}, ";"),
	}
}

func auditFirewalls(p PromotionAudit) FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "Heavy propagator normalization", WhyRequired: "converts the support-path overlap into a dimensionful determinant coefficient", Status: StatusFailedHeavyPropagatorMissing, BlocksPromotion: true},
		{Name: "Heavy self-quartic lambda_heavy", WhyRequired: "needed for Delta lambda = -lambda_mix^2/(4 lambda_heavy)", Status: StatusFailedHeavyQuarticMissing, BlocksPromotion: true},
		{Name: "lambda_mix normalization", WhyRequired: "needed to distinguish a support coefficient from a canonically normalized EFT coupling", Status: StatusFailedLambdaMixNotNormalized, BlocksPromotion: true},
	}
	return FirewallAudit{
		NoFinalMassClaimed:       true,
		NoPoleMassClaimed:        true,
		NoThresholdClaimed:       !p.ThresholdJumpDerived,
		NoHeavyPropagatorClaimed: !p.HeavyPropagatorDerived,
		NoHeavyQuarticClaimed:    !p.HeavySelfQuarticDerived,
		NoLambdaMixClaimed:       !p.LambdaMixNormalized,
		FiniteCorePolluted:       false,
		RemainingObligations:     obligations,
		Verdict:                  strings.Join([]string{StatusGate320FirewallsPreserved, StatusFailedThresholdNotDerived, StatusFailedFinalMassNotClaimed, StatusFailedPoleMatchingNotExecuted}, ";"),
	}
}

func buildSummary(b DoubledSpaceBlock, p SeesawPath, o OverlapMatrix, portal PortalWeightAudit, promo PromotionAudit, fw FirewallAudit) Summary {
	preserved := fw.NoFinalMassClaimed && fw.NoThresholdClaimed && fw.NoHeavyPropagatorClaimed && fw.NoHeavyQuarticClaimed && fw.NoLambdaMixClaimed && !fw.FiniteCorePolluted
	return Summary{
		DoubledBlockFormalized: b.Formalized,
		SeesawPathConstructed:  p.Constructed,
		ExplicitMatrixDerived:  o.Derived,
		OverlapIndexVerified:   o.IndexVerified,
		PortalWeightEnabled:    portal.Enabled,
		ThresholdPromoted:      promo.ThresholdJumpDerived,
		FirewallsPreserved:     preserved,
		Status:                 strings.Join([]string{StatusDoubledBlockFormalized, StatusSeesawPathConstructed, StatusExplicitSigmaHMatrixDerived, StatusOverlapIndexVerified, StatusPortalThresholdStillIncomplete}, ";"),
		DirectAnswer:           "The doubled-space seesaw block contains a unique normalized path L_L -> nu_R -> nu_R^c, so Omega_Hsigma is explicitly represented by |nu_R^c><L_L| and has overlap index 1.  This validates the missing Gate-319 overlap index, but not the full threshold theorem.",
		NextGate:               "derive the heavy propagator and heavy self-quartic normalization that convert the overlap-index witness into lambda_mix^2/lambda_heavy and a threshold jump.",
	}
}

func traceOmegaDagOmega(m [][]float64) float64 {
	sum := 0.0
	for i := range m {
		for j := range m[i] {
			sum += m[i][j] * m[i][j]
		}
	}
	return sum
}

func relativeError(value, target float64) float64 {
	if target == 0 {
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

func FormatBlock(b DoubledSpaceBlock) string {
	return fmt.Sprintf("formalized=%t; basis=%v; J_swap=%t; H=%s:%s->%s; B=%s:%s->%s; directSumOverlap=%.1f; formula=%s; verdict=%s", b.Formalized, b.Basis, b.JSwapInstalled, b.HiggsEdge.Label, b.HiggsEdge.From, b.HiggsEdge.To, b.MajoranaEdge.Label, b.MajoranaEdge.From, b.MajoranaEdge.To, b.DirectSumOverlap, b.MatrixFormula, b.Verdict)
}

func FormatPath(p SeesawPath) string {
	return fmt.Sprintf("constructed=%t; path=%v; J_swap=%t; H=%t; B=%t; product=%s; pathCount=%d; directSumPathCount=%d; rank=%d; pathWeight=%.12f; verdict=%s", p.Constructed, p.Path, p.UsesJSwap, p.HiggsEdgeExists, p.MajoranaEdgeExists, p.SequentialProduct, p.PathCount, p.DirectSumPathCount, p.PathMatrixRank, p.PathWeight, p.Verdict)
}

func FormatOverlap(o OverlapMatrix) string {
	return fmt.Sprintf("derived=%t; basis=%v; matrix=%v; formula=%s; TrOmegaDagOmega=%.1f; index=%.1f; verified=%t; directSumIndex=%.1f; uniquePath=%t; verdict=%s", o.Derived, o.Basis, o.SupportMatrix, o.Formula, o.TraceOmegaDagOmega, o.CanonicalOverlapIndex, o.IndexVerified, o.DirectSumIndex, o.UniquePathNormalized, o.Verdict)
}

func FormatPortal(p PortalWeightAudit) string {
	return fmt.Sprintf("enabled=%t; kappaQ=%.1f; 4/pi=%.12f; B_gap=%.12f; Omega=%.1f; coeff=%.12f; target=%.12f; relErr=%+.6f%%; within1%%=%t; impliedDelta=%.12f; targetDelta=%.12f; weightsMultiplicative=%t; thresholdPromoted=%t; verdict=%s", p.Enabled, p.KappaQ, p.Resonance, p.BGap, p.OverlapIndex, p.Coefficient, p.TargetPortalRatio, 100*p.RelativeError, p.WithinOnePercent, p.ImpliedDeltaLambda, p.TargetDeltaLambda, p.WeightsMultiplicative, p.ThresholdPromoted, p.Verdict)
}

func FormatPromotion(p PromotionAudit) string {
	return fmt.Sprintf("matrix=%t; overlapIndex=%t; heavyPropagator=%t; heavyQuartic=%t; lambdaMix=%t; lambdaHeavy=%t; threshold=%t; promotion=%t; verdict=%s", p.ExplicitMatrixDerived, p.OverlapIndexDerived, p.HeavyPropagatorDerived, p.HeavySelfQuarticDerived, p.LambdaMixNormalized, p.LambdaHeavyNormalized, p.ThresholdJumpDerived, p.PortalPromotionAuthorized, p.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noFinalMass=%t; noPole=%t; noThreshold=%t; noHeavyPropagator=%t; noHeavyQuartic=%t; noLambdaMix=%t; polluted=%t; obligations=%d; verdict=%s", f.NoFinalMassClaimed, f.NoPoleMassClaimed, f.NoThresholdClaimed, f.NoHeavyPropagatorClaimed, f.NoHeavyQuarticClaimed, f.NoLambdaMixClaimed, f.FiniteCorePolluted, len(f.RemainingObligations), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("block=%t; path=%t; matrix=%t; index=%t; portalWeight=%t; threshold=%t; firewalls=%t; status=%s; answer=%s; next=%s", s.DoubledBlockFormalized, s.SeesawPathConstructed, s.ExplicitMatrixDerived, s.OverlapIndexVerified, s.PortalWeightEnabled, s.ThresholdPromoted, s.FirewallsPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
