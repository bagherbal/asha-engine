// Package nonperturbativeportalcoupling implements Gate 318:
// Non-Perturbative Instanton Mapping / Heavy Portal Coupling Sieve Audit.
//
// Gate 314 extracted a PeV/intermediate threshold obligation for the Higgs
// quartic coupling,
//
//	Delta lambda ~= -0.097561578813
//
// equivalently a scalar-portal target
//
//	lambda_mix^2 / lambda_heavy ~= 0.390246315254.
//
// Gate 318 audits whether the B-gap Majorana/topological sector has enough
// native non-perturbative structure to generate a portal of this size.  It
// deliberately distinguishes three things:
//
//  1. the already known inverse instanton action S_inst=(4/pi)/B_gap,
//  2. direct instanton-suppressed factors exp(-S_inst), and
//  3. algebraic topological-overlap witnesses such as kappa_Q*(4/pi)*B_gap.
//
// The audit finds a striking magnitude witness close to the Gate-314 target,
// but it does not promote that witness to a derived heavy portal because the
// functional determinant, sigma-H overlap operator, heavy VEV, and threshold
// matching theorem are still missing.
package nonperturbativeportalcoupling

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE318-NON-PERTURBATIVE-INSTANTON-MAPPING-HEAVY-PORTAL-COUPLING-SIEVE-AUDIT"

	StatusBGapInstantonActionFormalized      = "CONDITIONAL_SUPPORT_BGAP_INSTANTON_ACTION_FORMALIZED"
	StatusHeavyPortalExtractionAudited       = "CONDITIONAL_SUPPORT_HEAVY_PORTAL_EXTRACTION_AUDITED"
	StatusPortalTargetSieveFormalized        = "CONDITIONAL_SUPPORT_PORTAL_TARGET_SIEVE_FORMALIZED"
	StatusTopologicalOverlapWitnessFound     = "CONDITIONAL_SUPPORT_TOPOLOGICAL_OVERLAP_MAGNITUDE_WITNESS_FOUND"
	StatusNonPerturbativePortalNotMapped     = "CONDITIONAL_TENSION_NON_PERTURBATIVE_PORTAL_MAP_NOT_YET_DERIVED"
	StatusBGapFirewallsPreserved             = "CONDITIONAL_SUPPORT_GATE318_FIREWALLS_PRESERVED"
	StatusFailedFunctionalDeterminantMissing = "FAILED_ROUTE_FUNCTIONAL_DETERMINANT_NOT_DERIVED"
	StatusFailedPortalCouplingNotDerived     = "FAILED_ROUTE_HEAVY_PORTAL_COUPLING_NOT_DERIVED"
	StatusFailedHeavyQuarticNotDerived       = "FAILED_ROUTE_HEAVY_SELF_QUARTIC_NOT_DERIVED"
	StatusFailedSigmaHOverlapOperatorMissing = "FAILED_ROUTE_SIGMA_H_OVERLAP_OPERATOR_NOT_DERIVED"
	StatusFailedThresholdJumpNotDerived      = "FAILED_ROUTE_THRESHOLD_JUMP_NOT_DERIVED_FROM_BGAP"
	StatusFailedInstantonSuppressionTooSmall = "FAILED_ROUTE_DIRECT_INSTANTON_EXP_SUPPRESSION_TOO_SMALL"
	StatusFailedFinalHiggsMassNotClaimed     = "FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedRGTransportNotReexecuted     = "FAILED_ROUTE_RG_TRANSPORT_NOT_REEXECUTED_IN_GATE318"
)

const (
	bGap                       = 0.102464921191
	resonanceFourOverPi        = 4.0 / math.Pi
	requiredDeltaLambdaGate314 = -0.097561578813
	requiredPortalRatioGate314 = 0.390246315254
	kappaMajorana              = 1.0
	kappaQuark                 = 3.0
	witnessToleranceFraction   = 0.01
)

type InstantonAction struct {
	Formalized                   bool
	BGap                         float64
	Resonance                    float64
	Action                       float64
	DirectInstantonFactor        float64
	PolynomialHeatKernelOK       bool
	DirectExpCanHitTarget        bool
	FunctionalDeterminantDerived bool
	Formula                      string
	Verdict                      string
}

type PortalTarget struct {
	RequiredDeltaLambda float64
	RequiredRatio       float64
	Formula             string
	SourceGate          string
	CorrectSignNeeded   bool
	ModerateMagnitude   bool
	Verdict             string
}

type Candidate struct {
	Name             string
	Formula          string
	Value            float64
	RelativeError    float64
	AbsError         float64
	CorrectSign      bool
	WithinOnePercent bool
	Category         string
	CanonicalDerived bool
	Interpretation   string
	Status           string
}

type PortalExtraction struct {
	Formalized                      bool
	Candidates                      []Candidate
	BestCandidate                   Candidate
	HasMagnitudeWitness             bool
	MagnitudeWitnessDerivedAsPortal bool
	SigmaHOverlapOperatorDerived    bool
	HeavySelfQuarticDerived         bool
	HeavyVEVDerived                 bool
	Verdict                         string
}

type Sieve struct {
	Formalized                    bool
	TargetRatio                   float64
	BestWitnessName               string
	BestWitnessValue              float64
	BestWitnessRelativeError      float64
	BestWitnessWithinOnePercent   bool
	DirectInstantonFactorTooSmall bool
	FourOverPiBGapWitness         float64
	KappaQFourOverPiBGapWitness   float64
	NativePortalMapped            bool
	TheoreticalCapacity           bool
	Verdict                       string
}

type FirewallAudit struct {
	NoPortalCouplingClaimed        bool
	NoThresholdJumpClaimed         bool
	NoFunctionalDeterminantClaimed bool
	NoHeavyVEVClaimed              bool
	NoHeavyQuarticClaimed          bool
	NoRGReexecutionClaimed         bool
	NoFinalMassClaimed             bool
	FiniteCorePolluted             bool
	Obligations                    []RemainingObligation
	Verdict                        string
}

type RemainingObligation struct {
	Name                  string
	WhyRequired           string
	Status                string
	BlocksNativePortalMap bool
}

type Summary struct {
	InstantonActionFormalized bool
	PortalTargetFormalized    bool
	PortalExtractionAudited   bool
	MagnitudeWitnessFound     bool
	NativePortalMapped        bool
	ThresholdJumpDerived      bool
	FinalMassClaimed          bool
	FirewallsPreserved        bool
	Status                    string
	DirectAnswer              string
	NextGate                  string
}

type Analysis struct {
	Instanton InstantonAction
	Target    PortalTarget
	Portal    PortalExtraction
	Sieve     Sieve
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
	inst := formalizeInstantonAction()
	target := formalizePortalTarget()
	portal := auditPortalExtraction(target)
	sieve := auditPortalTargetSieve(inst, portal, target)
	firewalls := auditFirewalls(sieve)
	summary := buildSummary(inst, target, portal, sieve, firewalls)
	truth := "Gate 318 finds that the B-gap sector has the right topological ingredients and even a striking magnitude witness: kappa_Q*(4/pi)*B_gap = 0.391387..., within about 0.3% of the Gate-314 portal target lambda_mix^2/lambda_heavy = 0.390246....  But this is not yet a derived threshold jump.  The direct instanton factor exp[-(4/pi)/B_gap] is far too small, and the functional determinant / sigma-H overlap operator that would convert the B-gap action into a physical portal coupling remains unbuilt."
	return Analysis{Instanton: inst, Target: target, Portal: portal, Sieve: sieve, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeInstantonAction() InstantonAction {
	action := resonanceFourOverPi / bGap
	expFactor := math.Exp(-action)
	return InstantonAction{
		Formalized:                   true,
		BGap:                         bGap,
		Resonance:                    resonanceFourOverPi,
		Action:                       action,
		DirectInstantonFactor:        expFactor,
		PolynomialHeatKernelOK:       false,
		DirectExpCanHitTarget:        withinFraction(expFactor, requiredPortalRatioGate314, witnessToleranceFraction),
		FunctionalDeterminantDerived: false,
		Formula:                      "S_inst = (4/pi)/B_gap; direct tunneling amplitude A_inst = exp(-S_inst)",
		Verdict:                      strings.Join([]string{StatusBGapInstantonActionFormalized, StatusFailedInstantonSuppressionTooSmall, StatusFailedFunctionalDeterminantMissing}, ";"),
	}
}

func formalizePortalTarget() PortalTarget {
	return PortalTarget{
		RequiredDeltaLambda: requiredDeltaLambdaGate314,
		RequiredRatio:       requiredPortalRatioGate314,
		Formula:             "Delta lambda = -lambda_mix^2/(4 lambda_heavy); target lambda_mix^2/lambda_heavy = -4 Delta lambda",
		SourceGate:          "Gate 314 gauge-only PeV threshold lower-envelope extraction",
		CorrectSignNeeded:   true,
		ModerateMagnitude:   math.Abs(requiredPortalRatioGate314) > 0.1 && math.Abs(requiredPortalRatioGate314) < 1,
		Verdict:             StatusPortalTargetSieveFormalized,
	}
}

func auditPortalExtraction(target PortalTarget) PortalExtraction {
	candidates := buildCandidates(target.RequiredRatio)
	best := candidates[0]
	for _, c := range candidates[1:] {
		if math.Abs(c.RelativeError) < math.Abs(best.RelativeError) {
			best = c
		}
	}
	hasWitness := best.WithinOnePercent
	return PortalExtraction{
		Formalized:                      true,
		Candidates:                      candidates,
		BestCandidate:                   best,
		HasMagnitudeWitness:             hasWitness,
		MagnitudeWitnessDerivedAsPortal: false,
		SigmaHOverlapOperatorDerived:    false,
		HeavySelfQuarticDerived:         false,
		HeavyVEVDerived:                 false,
		Verdict:                         strings.Join([]string{StatusHeavyPortalExtractionAudited, StatusTopologicalOverlapWitnessFound, StatusNonPerturbativePortalNotMapped, StatusFailedPortalCouplingNotDerived, StatusFailedSigmaHOverlapOperatorMissing}, ";"),
	}
}

func buildCandidates(target float64) []Candidate {
	values := []struct {
		name, formula, category, interpretation string
		value                                   float64
		derived                                 bool
		status                                  string
	}{
		{"direct B_gap", "B_gap", "linear-B", "Majorana order parameter alone; too small for the Gate-314 target", bGap, false, StatusFailedPortalCouplingNotDerived},
		{"topological resonance", "4/pi", "resonance", "pure topological action coefficient; too large as a portal ratio", resonanceFourOverPi, false, StatusFailedPortalCouplingNotDerived},
		{"linear B-gap resonance", "(4/pi)*B_gap", "topological-overlap", "B-gap weighted by the 4/pi resonance; has the right scale order but undershoots by about a factor of three", resonanceFourOverPi * bGap, false, StatusFailedPortalCouplingNotDerived},
		{"Morita quark-color overlap witness", "kappa_Q*(4/pi)*B_gap", "topological-overlap-witness", "uses the native quark Morita multiplicity kappa_Q=3 with the B-gap resonance; close to the extracted portal-ratio target", kappaQuark * resonanceFourOverPi * bGap, false, StatusTopologicalOverlapWitnessFound},
		{"quadratic resonance B-gap", "(4/pi)^2*B_gap", "quadratic-overlap", "squares the topological resonance before weighting by B_gap; not close to target", resonanceFourOverPi * resonanceFourOverPi * bGap, false, StatusFailedPortalCouplingNotDerived},
		{"Majorana multiplicity witness", "kappa_M*(4/pi)*B_gap", "majorana-overlap", "single Majorana trace multiplicity only; equals the linear B-gap resonance", kappaMajorana * resonanceFourOverPi * bGap, false, StatusFailedPortalCouplingNotDerived},
		{"direct instanton exponential", "exp[-(4/pi)/B_gap]", "instanton-suppressed", "ordinary tunneling amplitude from S_inst; far too small to account for the needed finite threshold jump", math.Exp(-resonanceFourOverPi / bGap), false, StatusFailedInstantonSuppressionTooSmall},
	}
	out := make([]Candidate, 0, len(values))
	for _, v := range values {
		absErr := v.value - target
		rel := absErr / target
		out = append(out, Candidate{
			Name:             v.name,
			Formula:          v.formula,
			Value:            v.value,
			RelativeError:    rel,
			AbsError:         absErr,
			CorrectSign:      v.value > 0,
			WithinOnePercent: withinFraction(v.value, target, witnessToleranceFraction),
			Category:         v.category,
			CanonicalDerived: v.derived,
			Interpretation:   v.interpretation,
			Status:           v.status,
		})
	}
	return out
}

func auditPortalTargetSieve(inst InstantonAction, portal PortalExtraction, target PortalTarget) Sieve {
	fourB := resonanceFourOverPi * bGap
	kqFourB := kappaQuark * fourB
	return Sieve{
		Formalized:                    true,
		TargetRatio:                   target.RequiredRatio,
		BestWitnessName:               portal.BestCandidate.Name,
		BestWitnessValue:              portal.BestCandidate.Value,
		BestWitnessRelativeError:      portal.BestCandidate.RelativeError,
		BestWitnessWithinOnePercent:   portal.BestCandidate.WithinOnePercent,
		DirectInstantonFactorTooSmall: inst.DirectInstantonFactor < 1e-3*target.RequiredRatio,
		FourOverPiBGapWitness:         fourB,
		KappaQFourOverPiBGapWitness:   kqFourB,
		NativePortalMapped:            false,
		TheoreticalCapacity:           portal.HasMagnitudeWitness && kqFourB > 0.1 && kqFourB < 1.0,
		Verdict:                       strings.Join([]string{StatusPortalTargetSieveFormalized, StatusTopologicalOverlapWitnessFound, StatusNonPerturbativePortalNotMapped, StatusFailedThresholdJumpNotDerived}, ";"),
	}
}

func auditFirewalls(s Sieve) FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "Functional determinant of the Majorana/B-gap edge", WhyRequired: "needed to convert S_inst into an EFT coefficient rather than a diagnostic action", Status: StatusFailedFunctionalDeterminantMissing, BlocksNativePortalMap: true},
		{Name: "Sigma-H overlap operator", WhyRequired: "needed to prove that the B-gap sector couples to |H|^2 with the kappa_Q*(4/pi)*B_gap strength", Status: StatusFailedSigmaHOverlapOperatorMissing, BlocksNativePortalMap: true},
		{Name: "Heavy self-quartic lambda_heavy", WhyRequired: "needed to evaluate Delta lambda=-lambda_mix^2/(4lambda_heavy) as a derived threshold", Status: StatusFailedHeavyQuarticNotDerived, BlocksNativePortalMap: true},
		{Name: "Threshold matching theorem", WhyRequired: "needed to place the finite jump at the correct intermediate scale in the EFT", Status: StatusFailedThresholdJumpNotDerived, BlocksNativePortalMap: true},
	}
	return FirewallAudit{
		NoPortalCouplingClaimed:        true,
		NoThresholdJumpClaimed:         true,
		NoFunctionalDeterminantClaimed: true,
		NoHeavyVEVClaimed:              true,
		NoHeavyQuarticClaimed:          true,
		NoRGReexecutionClaimed:         true,
		NoFinalMassClaimed:             true,
		FiniteCorePolluted:             false,
		Obligations:                    obligations,
		Verdict:                        strings.Join([]string{StatusBGapFirewallsPreserved, StatusFailedPortalCouplingNotDerived, StatusFailedThresholdJumpNotDerived, StatusFailedFinalHiggsMassNotClaimed}, ";"),
	}
}

func buildSummary(inst InstantonAction, target PortalTarget, portal PortalExtraction, sieve Sieve, fw FirewallAudit) Summary {
	preserved := fw.NoPortalCouplingClaimed && fw.NoThresholdJumpClaimed && fw.NoFunctionalDeterminantClaimed && fw.NoHeavyQuarticClaimed && fw.NoFinalMassClaimed && !fw.FiniteCorePolluted
	direct := "B-gap non-perturbative action is formalized and a near-exact topological-overlap magnitude witness kappa_Q*(4/pi)*B_gap is found, but the actual heavy portal threshold is not derived."
	next := "derive the sigma-H overlap operator or functional determinant that turns the kappa_Q*(4/pi)*B_gap witness into lambda_mix^2/lambda_heavy."
	return Summary{
		InstantonActionFormalized: inst.Formalized,
		PortalTargetFormalized:    target.RequiredRatio > 0,
		PortalExtractionAudited:   portal.Formalized,
		MagnitudeWitnessFound:     sieve.BestWitnessWithinOnePercent,
		NativePortalMapped:        sieve.NativePortalMapped,
		ThresholdJumpDerived:      false,
		FinalMassClaimed:          false,
		FirewallsPreserved:        preserved,
		Status:                    strings.Join([]string{StatusBGapInstantonActionFormalized, StatusPortalTargetSieveFormalized, StatusTopologicalOverlapWitnessFound, StatusNonPerturbativePortalNotMapped}, ";"),
		DirectAnswer:              direct,
		NextGate:                  next,
	}
}

func withinFraction(value, target, tol float64) bool {
	if target == 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return false
	}
	return math.Abs((value-target)/target) <= tol
}

func FormatInstanton(i InstantonAction) string {
	return fmt.Sprintf("%s; B_gap=%.12f; 4/pi=%.12f; S_inst=%.12f; exp(-S)=%.12g; directExpCanHitTarget=%t; determinantDerived=%t; verdict=%s", i.Formula, i.BGap, i.Resonance, i.Action, i.DirectInstantonFactor, i.DirectExpCanHitTarget, i.FunctionalDeterminantDerived, i.Verdict)
}

func FormatTarget(t PortalTarget) string {
	return fmt.Sprintf("%s from %s; DeltaLambda=%.12f; requiredRatio=%.12f; moderate=%t; verdict=%s", t.Formula, t.SourceGate, t.RequiredDeltaLambda, t.RequiredRatio, t.ModerateMagnitude, t.Verdict)
}

func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("%s: %s = %.12f; relErr=%+.6f%%; within1%%=%t; derived=%t; category=%s; status=%s", c.Name, c.Formula, c.Value, 100*c.RelativeError, c.WithinOnePercent, c.CanonicalDerived, c.Category, c.Status)
}

func FormatPortal(p PortalExtraction) string {
	return fmt.Sprintf("formalized=%t; best={%s}; witness=%t; witnessDerivedAsPortal=%t; sigmaHOperatorDerived=%t; heavyQuarticDerived=%t; verdict=%s", p.Formalized, FormatCandidate(p.BestCandidate), p.HasMagnitudeWitness, p.MagnitudeWitnessDerivedAsPortal, p.SigmaHOverlapOperatorDerived, p.HeavySelfQuarticDerived, p.Verdict)
}

func FormatSieve(s Sieve) string {
	return fmt.Sprintf("target=%.12f; best=%s %.12f relErr=%+.6f%%; expSuppressionTooSmall=%t; (4/pi)B=%.12f; kappa_Q(4/pi)B=%.12f; theoreticalCapacity=%t; nativePortalMapped=%t; verdict=%s", s.TargetRatio, s.BestWitnessName, s.BestWitnessValue, 100*s.BestWitnessRelativeError, s.DirectInstantonFactorTooSmall, s.FourOverPiBGapWitness, s.KappaQFourOverPiBGapWitness, s.TheoreticalCapacity, s.NativePortalMapped, s.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noPortal=%t; noJump=%t; noDeterminant=%t; noHeavyVEV=%t; noHeavyQuartic=%t; noRG=%t; noFinalMass=%t; polluted=%t; obligations=%d; verdict=%s", f.NoPortalCouplingClaimed, f.NoThresholdJumpClaimed, f.NoFunctionalDeterminantClaimed, f.NoHeavyVEVClaimed, f.NoHeavyQuarticClaimed, f.NoRGReexecutionClaimed, f.NoFinalMassClaimed, f.FiniteCorePolluted, len(f.Obligations), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("instanton=%t; target=%t; portalAudit=%t; witness=%t; nativeMapped=%t; jumpDerived=%t; finalMass=%t; firewalls=%t; status=%s; answer=%s; next=%s", s.InstantonActionFormalized, s.PortalTargetFormalized, s.PortalExtractionAudited, s.MagnitudeWitnessFound, s.NativePortalMapped, s.ThresholdJumpDerived, s.FinalMassClaimed, s.FirewallsPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
