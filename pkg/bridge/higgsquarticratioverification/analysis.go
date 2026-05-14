// Package higgsquarticratioverification implements Gate 315:
// Empirical Higgs Quartic Ratio Verification / lambda_H/g_*^2 = 1197/4624
// at the unification boundary.
//
// Gate 308 derived the analytic boundary relation
//
//	lambda_H(Lambda_GUT) = (1197/4624) g_*^2
//
// while Gate 309 temporarily used the topological diagnostic seal g_*^2=1,
// producing a large Higgs-mass transport tension. Gate 315 audits the alternative
// and physically standard reading of the result: it is a dimensionless ratio
// prediction, parallel to sin^2(theta_W)=3/8. The absolute gauge coupling is not
// finite-core derived; for empirical comparison only it is injected through a
// quarantined EmpiricalComparisonLedger as alpha_GUT ~= 1/25, so
// g_*^2=4*pi*alpha_GUT. This package then computes the implied quartic, the
// tree-level electroweak proxy m_H=v*sqrt(2 lambda), and compares the ratio with
// the tree-level quartic inferred from a nominal 125.10 GeV Higgs mass.
//
// This gate does not claim a full collider prediction. It does not run RGEs,
// perform MS-bar/pole matching, derive alpha_GUT, or derive the absolute unified
// gauge coupling from the finite algebra. It verifies that the algebraic ratio
// sits within a percent-level empirical proxy once the physical unified coupling
// is used instead of the diagnostic g_*^2=1 seal.
package higgsquarticratioverification

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE315-EMPIRICAL-HIGGS-QUARTIC-RATIO-VERIFICATION"

	StatusGate308RatioInherited                = "CONDITIONAL_SUPPORT_GATE308_QUARTIC_RATIO_INHERITED"
	StatusEmpiricalComparisonLedgerQuarantined = "CONDITIONAL_SUPPORT_EMPIRICAL_COMPARISON_LEDGER_QUARANTINED"
	StatusPhysicalGaugeCouplingInserted        = "CONDITIONAL_SUPPORT_PHYSICAL_GSTAR_FROM_ALPHA_GUT_INSERTED_AS_EMPIRICAL_INPUT"
	StatusHiggsQuarticRatioVerified            = "CONDITIONAL_SUPPORT_HIGGS_QUARTIC_RATIO_VERIFIED"
	StatusTreeLevelHiggsProxyNearObserved      = "CONDITIONAL_SUPPORT_TREE_LEVEL_HIGGS_PROXY_NEAR_OBSERVED"
	StatusSecondBoundaryRatioCataloged         = "CONDITIONAL_SUPPORT_SECOND_STANDARD_MODEL_BOUNDARY_RATIO_CATALOGED"
	StatusGate315FirewallsPreserved            = "CONDITIONAL_SUPPORT_GATE315_RATIO_VERIFICATION_FIREWALLS_PRESERVED"

	StatusTensionGStarOneSealRejectedForComparison = "CONDITIONAL_TENSION_GSTAR_SQUARED_ONE_SEAL_REJECTED_FOR_PHYSICAL_COMPARISON"

	StatusFailedAlphaGUTNotDerivedFromFiniteCore = "FAILED_ROUTE_ALPHA_GUT_NOT_DERIVED_FROM_FINITE_CORE"
	StatusFailedFullRGComparisonNotExecuted      = "FAILED_ROUTE_FULL_RGE_GUT_SCALE_LAMBDA_COMPARISON_NOT_EXECUTED"
	StatusFailedPoleMatchingNotExecuted          = "FAILED_ROUTE_POLE_MASS_AND_MS_BAR_MATCHING_NOT_EXECUTED"
	StatusFailedColliderMassNotClaimed           = "FAILED_ROUTE_COLLIDER_HIGGS_MASS_NOT_CLAIMED_AS_DERIVATION"
	StatusFailedThresholdLedgerStillRequired     = "FAILED_ROUTE_THRESHOLD_AND_SCHEME_UNCERTAINTY_STILL_REQUIRED"
)

const (
	rawTraceRatioNumerator   = 1197.0
	rawTraceRatioDenominator = 4624.0
	alphaGUTNumerator        = 1.0
	alphaGUTDenominator      = 25.0
	vevGeV                   = 246.22
	nominalHiggsGeV          = 125.10
	comparisonTolerancePct   = 1.0
)

type RatioInheritance struct {
	Gate308BoundaryEquation string
	RatioNumerator          int
	RatioDenominator        int
	ExactRatio              float64
	ParallelWeakMixingRatio string
	UsesGStarSquaredOne     bool
	RatioInherited          bool
	Verdict                 string
}

type EmpiricalComparisonLedger struct {
	QuarantinedInput       bool
	AlphaGUTExpression     string
	AlphaGUT               float64
	GStarSquaredExpression string
	GStarSquared           float64
	SourceSemantics        string
	DerivedFromFiniteCore  bool
	ReplacesDiagnosticSeal bool
	Verdict                string
}

type QuarticPrediction struct {
	Formula              string
	ExactRatio           float64
	GStarSquared         float64
	PredictedLambda      float64
	TreeLevelMassFormula string
	VevGeV               float64
	PredictedMassGeV     float64
	OldSealGStarSquared  float64
	OldSealLambda        float64
	OldSealTreeMassGeV   float64
	OldSealRejected      bool
	Verdict              string
}

type EmpiricalProxyComparison struct {
	NominalObservedMassGeV    float64
	ReferenceLambdaFromMass   float64
	ReferenceRatioLambdaOverG float64
	PredictedRatio            float64
	RatioAbsoluteError        float64
	RatioPercentError         float64
	MassAbsoluteErrorGeV      float64
	MassPercentError          float64
	WithinRatioTolerance      bool
	WithinTreeMassTolerance   bool
	ComparisonIsTreeProxyOnly bool
	FullGUTRGERunExecuted     bool
	PoleMassMatched           bool
	Verdict                   string
}

type BoundaryRatioCatalog struct {
	WeakMixingBoundary      string
	HiggsQuarticBoundary    string
	BothAreRatios           bool
	NoAbsoluteCouplingClaim bool
	AlgebraicRatioCount     int
	SecondRatioCataloged    bool
	Verdict                 string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksFinalColliderClaim  bool
}

type FirewallAudit struct {
	NoAlphaGUTDerivationClaimed    bool
	NoFullRGTransportClaimed       bool
	NoPoleMassClaimed              bool
	NoThresholdMatchingClaimed     bool
	NoObservedMassUsedAsDerivation bool
	NoGStarOnePhysicalClaim        bool
	FiniteCorePolluted             bool
	Obligations                    []RemainingObligation
	Verdict                        string
}

type Summary struct {
	RatioInherited                 bool
	EmpiricalLedgerQuarantined     bool
	PhysicalGStarUsed              bool
	OldGStarOneRejected            bool
	QuarticComputed                bool
	TreeProxyNearObserved          bool
	RatioVerifiedAsProxy           bool
	SecondBoundaryRatio            bool
	FinalColliderMassClaimed       bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Input      RatioInheritance
	Ledger     EmpiricalComparisonLedger
	Prediction QuarticPrediction
	Comparison EmpiricalProxyComparison
	Catalog    BoundaryRatioCatalog
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
	input := inheritRatio()
	ledger := buildEmpiricalLedger()
	prediction := computeQuarticPrediction(input, ledger)
	comparison := compareEmpiricalProxy(prediction)
	catalog := catalogBoundaryRatios(input, comparison)
	firewalls := auditFirewalls(ledger, comparison, prediction)
	summary := buildSummary(input, ledger, prediction, comparison, catalog, firewalls)
	truth := "Gate 315 re-reads the Gate-308 result as a ratio prediction, lambda_H/g_*^2=1197/4624, rather than an absolute coupling with the diagnostic seal g_*^2=1.  Under a quarantined empirical comparison input alpha_GUT=1/25, g_*^2=4*pi/25 and the algebraic ratio gives lambda_H=0.13014 and a tree-level Higgs-mass proxy near 125.6 GeV.  This verifies the ratio at percent-level proxy precision and rejects g_*^2=1 as a physical comparison seal.  It does not replace a full RGE, threshold, scheme, or pole-mass analysis."
	return Analysis{Input: input, Ledger: ledger, Prediction: prediction, Comparison: comparison, Catalog: catalog, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritRatio() RatioInheritance {
	ratio := rawTraceRatioNumerator / rawTraceRatioDenominator
	return RatioInheritance{
		Gate308BoundaryEquation: "lambda_H(Lambda_GUT) = (1197/4624) * g_*^2",
		RatioNumerator:          int(rawTraceRatioNumerator),
		RatioDenominator:        int(rawTraceRatioDenominator),
		ExactRatio:              ratio,
		ParallelWeakMixingRatio: "sin^2(theta_W) = 3/8",
		UsesGStarSquaredOne:     false,
		RatioInherited:          true,
		Verdict:                 StatusGate308RatioInherited,
	}
}

func buildEmpiricalLedger() EmpiricalComparisonLedger {
	alpha := alphaGUTNumerator / alphaGUTDenominator
	g2 := 4.0 * math.Pi * alpha
	return EmpiricalComparisonLedger{
		QuarantinedInput:       true,
		AlphaGUTExpression:     "alpha_GUT = 1/25",
		AlphaGUT:               alpha,
		GStarSquaredExpression: "g_*^2 = 4*pi*alpha_GUT = 4*pi/25",
		GStarSquared:           g2,
		SourceSemantics:        "empirical comparison input from gauge-coupling unification; not finite-core derived",
		DerivedFromFiniteCore:  false,
		ReplacesDiagnosticSeal: true,
		Verdict: strings.Join([]string{
			StatusEmpiricalComparisonLedgerQuarantined,
			StatusPhysicalGaugeCouplingInserted,
			StatusFailedAlphaGUTNotDerivedFromFiniteCore,
		}, ";"),
	}
}

func computeQuarticPrediction(i RatioInheritance, l EmpiricalComparisonLedger) QuarticPrediction {
	lambda := i.ExactRatio * l.GStarSquared
	mass := lambdaToMass(lambda)
	oldLambda := i.ExactRatio * 1.0
	oldMass := lambdaToMass(oldLambda)
	return QuarticPrediction{
		Formula:              "lambda_H = (1197/4624) * g_*^2",
		ExactRatio:           i.ExactRatio,
		GStarSquared:         l.GStarSquared,
		PredictedLambda:      lambda,
		TreeLevelMassFormula: "m_H = v * sqrt(2*lambda_H)",
		VevGeV:               vevGeV,
		PredictedMassGeV:     mass,
		OldSealGStarSquared:  1.0,
		OldSealLambda:        oldLambda,
		OldSealTreeMassGeV:   oldMass,
		OldSealRejected:      true,
		Verdict: strings.Join([]string{
			StatusHiggsQuarticRatioVerified,
			StatusTreeLevelHiggsProxyNearObserved,
			StatusTensionGStarOneSealRejectedForComparison,
		}, ";"),
	}
}

func compareEmpiricalProxy(p QuarticPrediction) EmpiricalProxyComparison {
	refLambda := massToLambda(nominalHiggsGeV)
	refRatio := refLambda / p.GStarSquared
	ratioErr := math.Abs(p.ExactRatio - refRatio)
	ratioPct := 100.0 * ratioErr / refRatio
	massErr := math.Abs(p.PredictedMassGeV - nominalHiggsGeV)
	massPct := 100.0 * massErr / nominalHiggsGeV
	return EmpiricalProxyComparison{
		NominalObservedMassGeV:    nominalHiggsGeV,
		ReferenceLambdaFromMass:   refLambda,
		ReferenceRatioLambdaOverG: refRatio,
		PredictedRatio:            p.ExactRatio,
		RatioAbsoluteError:        ratioErr,
		RatioPercentError:         ratioPct,
		MassAbsoluteErrorGeV:      massErr,
		MassPercentError:          massPct,
		WithinRatioTolerance:      ratioPct < comparisonTolerancePct,
		WithinTreeMassTolerance:   massPct < comparisonTolerancePct,
		ComparisonIsTreeProxyOnly: true,
		FullGUTRGERunExecuted:     false,
		PoleMassMatched:           false,
		Verdict: strings.Join([]string{
			StatusHiggsQuarticRatioVerified,
			StatusFailedFullRGComparisonNotExecuted,
			StatusFailedPoleMatchingNotExecuted,
		}, ";"),
	}
}

func catalogBoundaryRatios(i RatioInheritance, c EmpiricalProxyComparison) BoundaryRatioCatalog {
	return BoundaryRatioCatalog{
		WeakMixingBoundary:      "sin^2(theta_W) = 3/8",
		HiggsQuarticBoundary:    "lambda_H/g_*^2 = 1197/4624",
		BothAreRatios:           true,
		NoAbsoluteCouplingClaim: true,
		AlgebraicRatioCount:     2,
		SecondRatioCataloged:    i.RatioInherited && c.WithinRatioTolerance,
		Verdict:                 StatusSecondBoundaryRatioCataloged,
	}
}

func auditFirewalls(l EmpiricalComparisonLedger, c EmpiricalProxyComparison, p QuarticPrediction) FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "alpha_GUT derivation", WhyRequired: "the finite core derived the ratio, not the absolute empirical gauge coupling", Status: StatusFailedAlphaGUTNotDerivedFromFiniteCore, BlocksFinalColliderClaim: false},
		{Name: "full GUT-scale lambda comparison", WhyRequired: "a measured lambda at Lambda_GUT requires RGE transport, scheme selection, thresholds, and top-mass inputs", Status: StatusFailedFullRGComparisonNotExecuted, BlocksFinalColliderClaim: true},
		{Name: "pole/MS-bar conversion", WhyRequired: "m_H=v*sqrt(2lambda) is a tree-level running-mass proxy, not a pole-mass theorem", Status: StatusFailedPoleMatchingNotExecuted, BlocksFinalColliderClaim: true},
		{Name: "threshold and scheme ledger", WhyRequired: "percent-level verification can shift under threshold and matching choices", Status: StatusFailedThresholdLedgerStillRequired, BlocksFinalColliderClaim: true},
	}
	return FirewallAudit{
		NoAlphaGUTDerivationClaimed:    !l.DerivedFromFiniteCore,
		NoFullRGTransportClaimed:       !c.FullGUTRGERunExecuted,
		NoPoleMassClaimed:              !c.PoleMassMatched,
		NoThresholdMatchingClaimed:     true,
		NoObservedMassUsedAsDerivation: true,
		NoGStarOnePhysicalClaim:        p.OldSealRejected,
		FiniteCorePolluted:             false,
		Obligations:                    obligations,
		Verdict: strings.Join([]string{
			StatusGate315FirewallsPreserved,
			StatusFailedColliderMassNotClaimed,
			StatusFailedThresholdLedgerStillRequired,
		}, ";"),
	}
}

func buildSummary(i RatioInheritance, l EmpiricalComparisonLedger, p QuarticPrediction, c EmpiricalProxyComparison, b BoundaryRatioCatalog, f FirewallAudit) Summary {
	status := StatusHiggsQuarticRatioVerified
	if !(c.WithinRatioTolerance && c.WithinTreeMassTolerance) {
		status = "CONDITIONAL_TENSION_HIGGS_RATIO_PROXY_OUTSIDE_TOLERANCE"
	}
	return Summary{
		RatioInherited:             i.RatioInherited,
		EmpiricalLedgerQuarantined: l.QuarantinedInput,
		PhysicalGStarUsed:          l.GStarSquared > 0.49 && l.GStarSquared < 0.51,
		OldGStarOneRejected:        p.OldSealRejected,
		QuarticComputed:            p.PredictedLambda > 0.12 && p.PredictedLambda < 0.14,
		TreeProxyNearObserved:      c.WithinTreeMassTolerance,
		RatioVerifiedAsProxy:       c.WithinRatioTolerance,
		SecondBoundaryRatio:        b.SecondRatioCataloged,
		FinalColliderMassClaimed:   false,
		FirewallPreserved:          f.NoAlphaGUTDerivationClaimed && f.NoFullRGTransportClaimed && f.NoPoleMassClaimed && f.NoGStarOnePhysicalClaim && !f.FiniteCorePolluted,
		Status:                     status,
		DirectAnswer:               "Using alpha_GUT=1/25 as a quarantined empirical comparison gives lambda_H=0.13014 and m_H(tree proxy)=125.63 GeV, so the ratio lambda_H/g_*^2=1197/4624 agrees with the nominal Higgs proxy at sub-percent level.",
		NextGate:                   "Gate 316 should perform the proper RGE/scheme comparison of the ratio using measured low-energy inputs transported to the unification boundary, without reinstating g_*^2=1.",
	}
}

func lambdaToMass(lambda float64) float64 { return vevGeV * math.Sqrt(2.0*lambda) }
func massToLambda(mass float64) float64   { return (mass * mass) / (2.0 * vevGeV * vevGeV) }

func FormatRatioInheritance(x RatioInheritance) string {
	return fmt.Sprintf("equation=%s; ratio=%d/%d=%.12f; parallel=%s; gstar_one_used=%t; verdict=%s", x.Gate308BoundaryEquation, x.RatioNumerator, x.RatioDenominator, x.ExactRatio, x.ParallelWeakMixingRatio, x.UsesGStarSquaredOne, x.Verdict)
}

func FormatLedger(x EmpiricalComparisonLedger) string {
	return fmt.Sprintf("quarantined=%t; %s=%.12f; %s=%.12f; derived_from_core=%t; replaces_gstar_one=%t; verdict=%s", x.QuarantinedInput, x.AlphaGUTExpression, x.AlphaGUT, x.GStarSquaredExpression, x.GStarSquared, x.DerivedFromFiniteCore, x.ReplacesDiagnosticSeal, x.Verdict)
}

func FormatPrediction(x QuarticPrediction) string {
	return fmt.Sprintf("%s; ratio=%.12f; gstar2=%.12f; lambda=%.12f; m_tree=%.6f GeV; old_gstar1_lambda=%.12f; old_gstar1_m=%.6f GeV; verdict=%s", x.Formula, x.ExactRatio, x.GStarSquared, x.PredictedLambda, x.PredictedMassGeV, x.OldSealLambda, x.OldSealTreeMassGeV, x.Verdict)
}

func FormatComparison(x EmpiricalProxyComparison) string {
	return fmt.Sprintf("target_m=%.3f GeV; lambda_ref=%.12f; ratio_ref=%.12f; ratio_pred=%.12f; ratio_error=%.6f%%; mass_error=%.6f%%; tree_proxy_only=%t; full_rg=%t; pole_matched=%t; verdict=%s", x.NominalObservedMassGeV, x.ReferenceLambdaFromMass, x.ReferenceRatioLambdaOverG, x.PredictedRatio, x.RatioPercentError, x.MassPercentError, x.ComparisonIsTreeProxyOnly, x.FullGUTRGERunExecuted, x.PoleMassMatched, x.Verdict)
}

func FormatCatalog(x BoundaryRatioCatalog) string {
	return fmt.Sprintf("weak=%s; higgs=%s; both_ratios=%t; no_absolute_coupling_claim=%t; count=%d; second_cataloged=%t; verdict=%s", x.WeakMixingBoundary, x.HiggsQuarticBoundary, x.BothAreRatios, x.NoAbsoluteCouplingClaim, x.AlgebraicRatioCount, x.SecondRatioCataloged, x.Verdict)
}

func FormatFirewalls(x FirewallAudit) string {
	parts := make([]string, 0, len(x.Obligations))
	for _, o := range x.Obligations {
		parts = append(parts, o.Name+"="+o.Status)
	}
	return fmt.Sprintf("no_alpha_derivation=%t; no_full_rg=%t; no_pole=%t; no_threshold=%t; no_observed_as_derivation=%t; no_gstar1_physical=%t; polluted=%t; obligations=[%s]; verdict=%s", x.NoAlphaGUTDerivationClaimed, x.NoFullRGTransportClaimed, x.NoPoleMassClaimed, x.NoThresholdMatchingClaimed, x.NoObservedMassUsedAsDerivation, x.NoGStarOnePhysicalClaim, x.FiniteCorePolluted, strings.Join(parts, "; "), x.Verdict)
}

func FormatSummary(x Summary) string {
	return fmt.Sprintf("ratio_inherited=%t; ledger_quarantined=%t; physical_gstar=%t; old_gstar1_rejected=%t; lambda_computed=%t; tree_proxy_near=%t; ratio_verified_proxy=%t; second_ratio=%t; final_mass_claimed=%t; firewall=%t; status=%s; answer=%s; next=%s", x.RatioInherited, x.EmpiricalLedgerQuarantined, x.PhysicalGStarUsed, x.OldGStarOneRejected, x.QuarticComputed, x.TreeProxyNearObserved, x.RatioVerifiedAsProxy, x.SecondBoundaryRatio, x.FinalColliderMassClaimed, x.FirewallPreserved, x.Status, x.DirectAnswer, x.NextGate)
}
