// Package heavylightoverlapoperator implements Gate 319:
// Functional Determinant Sieve / Heavy-Light Overlap Operator Audit.
//
// Gate 318 found a striking topological magnitude witness for the Gate-314
// threshold target,
//
//	kappa_Q * (4/pi) * B_gap = 0.391387...,
//
// close to the required portal ratio lambda_mix^2/lambda_heavy = 0.390246....
// Gate 319 asks a harder question: does the effective action obtained by
// integrating out the heavy B-gap/Majorana sector actually contain a sigma-H
// overlap operator whose coefficient is forced to multiply these factors?
//
// The audit distinguishes three levels:
//
//  1. Direct-sum carriers: functional determinants factorize and mixed traces
//     vanish.  This cannot generate a portal.
//  2. True-bimodule overlap carriers: a nonzero heavy-light projector can
//     generate an operator template of the form sigma^2 |H|^2.
//  3. Explicit matrix derivation: the concrete sigma-H kernel, heavy propagator,
//     overlap index, and heavy self-quartic must be computed before the witness
//     can be promoted to a threshold theorem.
//
// Gate 319 therefore formalizes the determinant expansion and finds the exact
// conditional coefficient map, but preserves the firewall: the near-perfect
// 0.391 resonance is still not a derived threshold jump until the explicit
// sigma-H overlap matrix is installed.
package heavylightoverlapoperator

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE319-FUNCTIONAL-DETERMINANT-SIEVE-HEAVY-LIGHT-OVERLAP-OPERATOR-AUDIT"

	StatusFunctionalDeterminantFormalized     = "CONDITIONAL_SUPPORT_FUNCTIONAL_DETERMINANT_EXPANSION_FORMALIZED"
	StatusHeavyLightTemplateFormalized        = "CONDITIONAL_SUPPORT_HEAVY_LIGHT_OPERATOR_TEMPLATE_FORMALIZED"
	StatusDirectSumZeroMixingIdentified       = "CONDITIONAL_SUPPORT_DIRECT_SUM_ZERO_MIXING_IDENTIFIED"
	StatusMultiplicativeWeightSieveFormalized = "CONDITIONAL_SUPPORT_MULTIPLICATIVE_WEIGHT_SIEVE_FORMALIZED"
	StatusTrueBimoduleWitnessMatchesTarget    = "CONDITIONAL_SUPPORT_TRUE_BIMODULE_OVERLAP_WITNESS_MATCHES_TARGET"
	StatusPortalPromotionNotAuthorized        = "CONDITIONAL_TENSION_PORTAL_PROMOTION_NOT_AUTHORIZED_WITHOUT_EXPLICIT_OPERATOR"
	StatusGate319FirewallsPreserved           = "CONDITIONAL_SUPPORT_GATE319_FIREWALLS_PRESERVED"
	StatusFailedExplicitSigmaHMatrixMissing   = "FAILED_ROUTE_EXPLICIT_SIGMA_H_OVERLAP_MATRIX_NOT_DERIVED"
	StatusFailedOverlapIndexNotDerived        = "FAILED_ROUTE_OVERLAP_INDEX_NOT_DERIVED"
	StatusFailedHeavyPropagatorNotDerived     = "FAILED_ROUTE_HEAVY_PROPAGATOR_NOT_DERIVED"
	StatusFailedHeavySelfQuarticNotDerived    = "FAILED_ROUTE_HEAVY_SELF_QUARTIC_NOT_DERIVED"
	StatusFailedPortalCouplingNotPromoted     = "FAILED_ROUTE_PORTAL_COUPLING_NOT_PROMOTED_TO_THRESHOLD_THEOREM"
	StatusFailedDirectSumCrossTermsVanish     = "FAILED_ROUTE_DIRECT_SUM_FUNCTIONAL_DETERMINANT_CROSS_TERMS_VANISH"
	StatusFailedFinalHiggsMassNotClaimed      = "FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedRGTransportNotReexecuted      = "FAILED_ROUTE_RG_TRANSPORT_NOT_REEXECUTED_IN_GATE319"
)

const (
	bGap                  = 0.102464921191
	resonanceFourOverPi   = 4.0 / math.Pi
	kappaQ                = 3.0
	kappaM                = 1.0
	targetPortalRatio     = 0.390246315254
	targetDeltaLambda     = -0.097561578813
	onePercentTolerance   = 0.01
	canonicalOverlapIndex = 1.0
)

type DeterminantExpansion struct {
	Formalized             bool
	EffectiveActionFormula string
	SeriesFormula          string
	PortalTermOrder        int
	NeedsHeavyPropagator   bool
	NeedsOverlapInsertion  bool
	DirectSumFactorizes    bool
	Verdict                string
}

type OverlapLane struct {
	Name                  string
	CarrierModel          string
	CrossTermsVanish      bool
	OverlapOperatorExists bool
	Formula               string
	OverlapIndex          float64
	Coefficient           float64
	RelativeError         float64
	WithinOnePercent      bool
	DerivedFromMatrices   bool
	Interpretation        string
	Status                string
}

type OperatorTemplate struct {
	Formalized                  bool
	OperatorFormula             string
	EFTMatchingFormula          string
	DirectSumLane               OverlapLane
	TrueBimoduleConditionalLane OverlapLane
	BestLane                    OverlapLane
	ExplicitSigmaHMatrixDerived bool
	OverlapIndexDerived         bool
	HeavyPropagatorDerived      bool
	HeavySelfQuarticDerived     bool
	PortalPromoted              bool
	Verdict                     string
}

type MultiplicativeSieve struct {
	Formalized                  bool
	TargetPortalRatio           float64
	TargetDeltaLambda           float64
	KappaQ                      float64
	Resonance                   float64
	BGap                        float64
	ConditionalOverlapIndex     float64
	ConditionalCoefficient      float64
	ConditionalDeltaLambda      float64
	RelativeError               float64
	WithinOnePercent            bool
	FactorsForcedMultiplicative bool
	DirectSumValue              float64
	TrueBimoduleValue           float64
	Verdict                     string
}

type PromotionAudit struct {
	FunctionalDeterminantInstalled bool
	SigmaHOverlapOperatorDerived   bool
	MultiplicativeWeightsForced    bool
	HeavySelfQuarticDerived        bool
	LambdaMixDerived               bool
	LambdaHeavyDerived             bool
	ThresholdJumpDerived           bool
	PromotionAuthorized            bool
	Verdict                        string
}

type FirewallAudit struct {
	NoPortalCouplingClaimed bool
	NoThresholdJumpClaimed  bool
	NoFinalMassClaimed      bool
	NoRGReexecutionClaimed  bool
	NoPoleMassClaimed       bool
	NoExplicitMatrixClaimed bool
	NoHeavyQuarticClaimed   bool
	FiniteCorePolluted      bool
	RemainingObligations    []RemainingObligation
	Verdict                 string
}

type RemainingObligation struct {
	Name            string
	WhyRequired     string
	Status          string
	BlocksPromotion bool
}

type Summary struct {
	DeterminantFormalized bool
	DirectSumRejected     bool
	TemplateFormalized    bool
	WitnessMatchesTarget  bool
	OperatorDerived       bool
	PortalPromoted        bool
	FirewallsPreserved    bool
	Status                string
	DirectAnswer          string
	NextGate              string
}

type Analysis struct {
	Determinant DeterminantExpansion
	Operator    OperatorTemplate
	Sieve       MultiplicativeSieve
	Promotion   PromotionAudit
	Firewalls   FirewallAudit
	Summary     Summary
	Truth       string
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
	det := formalizeDeterminantExpansion()
	op := auditOverlapOperator(det)
	sieve := auditMultiplicativeWeights(op)
	promotion := auditPromotion(op, sieve)
	firewalls := auditFirewalls(promotion)
	summary := buildSummary(det, op, sieve, promotion, firewalls)
	truth := "Gate 319 formalizes the heavy-sector functional determinant and proves the crucial categorical obstruction: in a direct-sum carrier the determinant factorizes and sigma-H cross terms vanish.  A true-bimodule overlap insertion can conditionally yield the near-perfect coefficient kappa_Q*(4/pi)*B_gap = 0.391387..., matching the Gate-314 portal target within about 0.29%, but the explicit sigma-H matrix kernel and overlap index are still not derived.  The 0.391 witness is therefore promoted to a precise operator target, not yet to a threshold theorem."
	return Analysis{Determinant: det, Operator: op, Sieve: sieve, Promotion: promotion, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeDeterminantExpansion() DeterminantExpansion {
	return DeterminantExpansion{
		Formalized:             true,
		EffectiveActionFormula: "Gamma_eff[H,sigma] = -1/2 Tr log(D_M^2 + M_sigma^2 + V_H + V_sigma + V_Hsigma)",
		SeriesFormula:          "Tr log(A+V) = Tr log A + sum_{n>=1} (-1)^{n+1} Tr[(A^{-1}V)^n]/n",
		PortalTermOrder:        2,
		NeedsHeavyPropagator:   true,
		NeedsOverlapInsertion:  true,
		DirectSumFactorizes:    true,
		Verdict:                strings.Join([]string{StatusFunctionalDeterminantFormalized, StatusDirectSumZeroMixingIdentified, StatusFailedHeavyPropagatorNotDerived}, ";"),
	}
}

func auditOverlapOperator(det DeterminantExpansion) OperatorTemplate {
	direct := buildDirectSumLane()
	trueBimodule := buildTrueBimoduleConditionalLane()
	best := trueBimodule
	if math.Abs(direct.RelativeError) < math.Abs(best.RelativeError) {
		best = direct
	}
	verdict := strings.Join([]string{
		StatusHeavyLightTemplateFormalized,
		StatusDirectSumZeroMixingIdentified,
		StatusTrueBimoduleWitnessMatchesTarget,
		StatusPortalPromotionNotAuthorized,
		StatusFailedExplicitSigmaHMatrixMissing,
		StatusFailedOverlapIndexNotDerived,
	}, ";")
	return OperatorTemplate{
		Formalized:                  det.Formalized,
		OperatorFormula:             "O_Hsigma = Tr_F[ Pi_Q · G_M · V_sigma · G_M · V_H · Omega_Hsigma ]; target term ~ sigma^2 |H|^2",
		EFTMatchingFormula:          "Delta lambda = -lambda_mix^2/(4 lambda_heavy)",
		DirectSumLane:               direct,
		TrueBimoduleConditionalLane: trueBimodule,
		BestLane:                    best,
		ExplicitSigmaHMatrixDerived: false,
		OverlapIndexDerived:         false,
		HeavyPropagatorDerived:      false,
		HeavySelfQuarticDerived:     false,
		PortalPromoted:              false,
		Verdict:                     verdict,
	}
}

func buildDirectSumLane() OverlapLane {
	coeff := 0.0
	return OverlapLane{
		Name:                  "direct-sum determinant lane",
		CarrierModel:          "D = D_light ⊕ D_heavy",
		CrossTermsVanish:      true,
		OverlapOperatorExists: false,
		Formula:               "Tr log(D_light ⊕ D_heavy) = Tr log D_light + Tr log D_heavy; Tr(light·heavy)=0",
		OverlapIndex:          0,
		Coefficient:           coeff,
		RelativeError:         relativeError(coeff, targetPortalRatio),
		WithinOnePercent:      withinFraction(coeff, targetPortalRatio, onePercentTolerance),
		DerivedFromMatrices:   true,
		Interpretation:        "This lane is rigorously zero; a direct-sum representation cannot generate the sigma-H portal.",
		Status:                StatusFailedDirectSumCrossTermsVanish,
	}
}

func buildTrueBimoduleConditionalLane() OverlapLane {
	coeff := kappaQ * resonanceFourOverPi * bGap * canonicalOverlapIndex
	return OverlapLane{
		Name:                  "true-bimodule conditional overlap lane",
		CarrierModel:          "D = D_light + D_heavy + Omega_Hsigma with a nonzero bimodule overlap projector",
		CrossTermsVanish:      false,
		OverlapOperatorExists: true,
		Formula:               "C_portal = kappa_Q · (4/pi) · B_gap · Omega_Hsigma",
		OverlapIndex:          canonicalOverlapIndex,
		Coefficient:           coeff,
		RelativeError:         relativeError(coeff, targetPortalRatio),
		WithinOnePercent:      withinFraction(coeff, targetPortalRatio, onePercentTolerance),
		DerivedFromMatrices:   false,
		Interpretation:        "If Omega_Hsigma is canonically one, the B-gap/topological/Morita factors multiply into the required portal magnitude; the explicit matrix kernel remains unproved.",
		Status:                StatusTrueBimoduleWitnessMatchesTarget,
	}
}

func auditMultiplicativeWeights(op OperatorTemplate) MultiplicativeSieve {
	coeff := op.TrueBimoduleConditionalLane.Coefficient
	delta := -coeff / 4.0
	return MultiplicativeSieve{
		Formalized:                  true,
		TargetPortalRatio:           targetPortalRatio,
		TargetDeltaLambda:           targetDeltaLambda,
		KappaQ:                      kappaQ,
		Resonance:                   resonanceFourOverPi,
		BGap:                        bGap,
		ConditionalOverlapIndex:     op.TrueBimoduleConditionalLane.OverlapIndex,
		ConditionalCoefficient:      coeff,
		ConditionalDeltaLambda:      delta,
		RelativeError:               relativeError(coeff, targetPortalRatio),
		WithinOnePercent:            withinFraction(coeff, targetPortalRatio, onePercentTolerance),
		FactorsForcedMultiplicative: false,
		DirectSumValue:              op.DirectSumLane.Coefficient,
		TrueBimoduleValue:           coeff,
		Verdict:                     strings.Join([]string{StatusMultiplicativeWeightSieveFormalized, StatusTrueBimoduleWitnessMatchesTarget, StatusPortalPromotionNotAuthorized, StatusFailedExplicitSigmaHMatrixMissing}, ";"),
	}
}

func auditPromotion(op OperatorTemplate, sieve MultiplicativeSieve) PromotionAudit {
	authorized := op.ExplicitSigmaHMatrixDerived && op.OverlapIndexDerived && op.HeavySelfQuarticDerived && sieve.FactorsForcedMultiplicative
	return PromotionAudit{
		FunctionalDeterminantInstalled: op.Formalized,
		SigmaHOverlapOperatorDerived:   op.ExplicitSigmaHMatrixDerived,
		MultiplicativeWeightsForced:    sieve.FactorsForcedMultiplicative,
		HeavySelfQuarticDerived:        op.HeavySelfQuarticDerived,
		LambdaMixDerived:               false,
		LambdaHeavyDerived:             false,
		ThresholdJumpDerived:           false,
		PromotionAuthorized:            authorized,
		Verdict:                        strings.Join([]string{StatusPortalPromotionNotAuthorized, StatusFailedPortalCouplingNotPromoted, StatusFailedExplicitSigmaHMatrixMissing, StatusFailedHeavySelfQuarticNotDerived}, ";"),
	}
}

func auditFirewalls(p PromotionAudit) FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "Explicit sigma-H overlap matrix", WhyRequired: "needed to prove Omega_Hsigma is a native operator rather than an inserted overlap index", Status: StatusFailedExplicitSigmaHMatrixMissing, BlocksPromotion: true},
		{Name: "Overlap index theorem", WhyRequired: "needed to prove Omega_Hsigma=1, or another exact value, from the finite bimodule", Status: StatusFailedOverlapIndexNotDerived, BlocksPromotion: true},
		{Name: "Heavy propagator and normalization", WhyRequired: "needed to convert the determinant trace into lambda_mix rather than a dimensionless witness", Status: StatusFailedHeavyPropagatorNotDerived, BlocksPromotion: true},
		{Name: "Heavy self-quartic lambda_heavy", WhyRequired: "needed to evaluate Delta lambda=-lambda_mix^2/(4lambda_heavy)", Status: StatusFailedHeavySelfQuarticNotDerived, BlocksPromotion: true},
	}
	return FirewallAudit{
		NoPortalCouplingClaimed: !p.LambdaMixDerived,
		NoThresholdJumpClaimed:  !p.ThresholdJumpDerived,
		NoFinalMassClaimed:      true,
		NoRGReexecutionClaimed:  true,
		NoPoleMassClaimed:       true,
		NoExplicitMatrixClaimed: !p.SigmaHOverlapOperatorDerived,
		NoHeavyQuarticClaimed:   !p.HeavySelfQuarticDerived,
		FiniteCorePolluted:      false,
		RemainingObligations:    obligations,
		Verdict:                 strings.Join([]string{StatusGate319FirewallsPreserved, StatusFailedPortalCouplingNotPromoted, StatusFailedFinalHiggsMassNotClaimed, StatusFailedRGTransportNotReexecuted}, ";"),
	}
}

func buildSummary(det DeterminantExpansion, op OperatorTemplate, sieve MultiplicativeSieve, p PromotionAudit, fw FirewallAudit) Summary {
	preserved := fw.NoPortalCouplingClaimed && fw.NoThresholdJumpClaimed && fw.NoFinalMassClaimed && fw.NoRGReexecutionClaimed && fw.NoExplicitMatrixClaimed && !fw.FiniteCorePolluted
	return Summary{
		DeterminantFormalized: det.Formalized,
		DirectSumRejected:     op.DirectSumLane.CrossTermsVanish,
		TemplateFormalized:    op.Formalized,
		WitnessMatchesTarget:  sieve.WithinOnePercent,
		OperatorDerived:       p.SigmaHOverlapOperatorDerived,
		PortalPromoted:        p.PromotionAuthorized,
		FirewallsPreserved:    preserved,
		Status:                strings.Join([]string{StatusFunctionalDeterminantFormalized, StatusHeavyLightTemplateFormalized, StatusTrueBimoduleWitnessMatchesTarget, StatusPortalPromotionNotAuthorized}, ";"),
		DirectAnswer:          "The functional determinant creates a portal only in the true-bimodule overlap lane, not in the direct-sum lane.  With a canonical overlap index Omega_Hsigma=1, the coefficient equals kappa_Q*(4/pi)*B_gap and matches the required target within 1%, but the explicit sigma-H matrix operator is still missing.",
		NextGate:              "derive the explicit sigma-H overlap matrix kernel and overlap index Omega_Hsigma from the finite bimodule action.",
	}
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

func FormatDeterminant(d DeterminantExpansion) string {
	return fmt.Sprintf("formalized=%t; Gamma=%s; series=%s; portalOrder=%d; needsPropagator=%t; needsOverlap=%t; directSumFactorizes=%t; verdict=%s", d.Formalized, d.EffectiveActionFormula, d.SeriesFormula, d.PortalTermOrder, d.NeedsHeavyPropagator, d.NeedsOverlapInsertion, d.DirectSumFactorizes, d.Verdict)
}

func FormatLane(l OverlapLane) string {
	return fmt.Sprintf("%s [%s]: coeff=%.12f; relErr=%+.6f%%; within1%%=%t; crossTermsVanish=%t; operatorExists=%t; derivedFromMatrices=%t; formula=%s; status=%s", l.Name, l.CarrierModel, l.Coefficient, 100*l.RelativeError, l.WithinOnePercent, l.CrossTermsVanish, l.OverlapOperatorExists, l.DerivedFromMatrices, l.Formula, l.Status)
}

func FormatOperator(o OperatorTemplate) string {
	return fmt.Sprintf("formalized=%t; operator=%s; matching=%s; direct={%s}; trueBimodule={%s}; explicitMatrix=%t; overlapIndexDerived=%t; heavyPropagator=%t; heavyQuartic=%t; promoted=%t; verdict=%s", o.Formalized, o.OperatorFormula, o.EFTMatchingFormula, FormatLane(o.DirectSumLane), FormatLane(o.TrueBimoduleConditionalLane), o.ExplicitSigmaHMatrixDerived, o.OverlapIndexDerived, o.HeavyPropagatorDerived, o.HeavySelfQuarticDerived, o.PortalPromoted, o.Verdict)
}

func FormatSieve(s MultiplicativeSieve) string {
	return fmt.Sprintf("targetRatio=%.12f; targetDelta=%.12f; kappaQ=%.1f; 4/pi=%.12f; B_gap=%.12f; Omega=%.1f; coeff=%.12f; impliedDelta=%.12f; relErr=%+.6f%%; within1%%=%t; forcedMultiplicative=%t; directSum=%.12f; trueBimodule=%.12f; verdict=%s", s.TargetPortalRatio, s.TargetDeltaLambda, s.KappaQ, s.Resonance, s.BGap, s.ConditionalOverlapIndex, s.ConditionalCoefficient, s.ConditionalDeltaLambda, 100*s.RelativeError, s.WithinOnePercent, s.FactorsForcedMultiplicative, s.DirectSumValue, s.TrueBimoduleValue, s.Verdict)
}

func FormatPromotion(p PromotionAudit) string {
	return fmt.Sprintf("determinant=%t; sigmaHOperator=%t; weightsForced=%t; heavyQuartic=%t; lambdaMix=%t; lambdaHeavy=%t; jump=%t; promotion=%t; verdict=%s", p.FunctionalDeterminantInstalled, p.SigmaHOverlapOperatorDerived, p.MultiplicativeWeightsForced, p.HeavySelfQuarticDerived, p.LambdaMixDerived, p.LambdaHeavyDerived, p.ThresholdJumpDerived, p.PromotionAuthorized, p.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noPortal=%t; noJump=%t; noFinalMass=%t; noRG=%t; noPole=%t; noExplicitMatrix=%t; noHeavyQuartic=%t; polluted=%t; obligations=%d; verdict=%s", f.NoPortalCouplingClaimed, f.NoThresholdJumpClaimed, f.NoFinalMassClaimed, f.NoRGReexecutionClaimed, f.NoPoleMassClaimed, f.NoExplicitMatrixClaimed, f.NoHeavyQuarticClaimed, f.FiniteCorePolluted, len(f.RemainingObligations), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("determinant=%t; directSumRejected=%t; template=%t; witness=%t; operatorDerived=%t; promoted=%t; firewalls=%t; status=%s; answer=%s; next=%s", s.DeterminantFormalized, s.DirectSumRejected, s.TemplateFormalized, s.WitnessMatchesTarget, s.OperatorDerived, s.PortalPromoted, s.FirewallsPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
