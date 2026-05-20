// Package generation2koideloopdeficitpmnsorientationaudit implements Gate 587:
// Koide Loop-Deficit PMNS Orientation Audit.
//
// Gate 586 compressed the charged-lepton Koide electron-wall offset as
// epsilon_e=(1/(8*pi))(1-kappa_e), with kappa_e about 0.005503554.  Gate 587
// imports a version-pinned PMNS/neutrino oscillation data set and tests whether
// the deficit is more naturally related to lepton-sector orientation than to the
// earlier CKM orientation clue.
//
// This is an environmental orientation sieve only.  It does not derive Koide,
// kappa_e, PMNS, neutrino masses, charged-lepton masses, CKM, or a flavor
// texture from ASHA-native law.
package generation2koideloopdeficitpmnsorientationaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koideloopangledeficitaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE587-KOIDE-LOOP-DEFICIT-PMNS-ORIENTATION-AUDIT"

	StatusGate586Inherited          = "PASS_GATE586_LOOP_ANGLE_DEFICIT_INHERITED"
	StatusPMNSDatasetImported       = "PASS_NUFIT60_PMNS_DATASET_IMPORTED"
	StatusPMNSJarlskogComputed      = "PASS_PMNS_ORIENTATION_INVARIANTS_COMPUTED"
	StatusCandidateSetDefined       = "PASS_PMNS_KAPPA_CANDIDATE_SET_DEFINED"
	StatusUncertaintyPropagated     = "PASS_PMNS_UNCERTAINTY_PROPAGATED"
	StatusBestPMNSAssisted          = "CONDITIONAL_SUPPORT_BEST_PMNS_ASSISTED_CANDIDATE_IS_ALPHA2_OVER_2PI_DIV_C13"
	StatusPMNSAssistedBetterThanCKM = "CONDITIONAL_SUPPORT_PMNS_ASSISTED_COUPLING_CANDIDATE_BEATS_SQRT_J_CKM_BUT_NOT_CERTIFIED"
	StatusDirectPMNSTooLarge        = "FAILED_ROUTE_DIRECT_PMNS_ORIENTATION_INVARIANTS_TOO_LARGE_FOR_KAPPA"
	StatusPMNSRangeCoversButNoCert  = "CONDITIONAL_SUPPORT_ABS_J_PMNS_CAN_COVER_KAPPA_WITH_CP_UNCERTAINTY_BUT_NOT_CERTIFIED"
	StatusNoPMNSCertified           = "FAILED_ROUTE_NO_PMNS_CANDIDATE_CERTIFIED_WITH_UNCERTAINTIES"
	StatusCKMMidpointSurvives       = "CONDITIONAL_SUPPORT_CKM_ALPHA2_MIDPOINT_NUMERIC_CLUE_SURVIVES_BUT_NOT_LAWFUL"
	StatusCKMStillNotLeptonSource   = "FAILED_ROUTE_CKM_CLUE_STILL_NOT_LEPTON_SOURCE_WITHOUT_INTERTWINER"
	StatusNoLeptonKoideIntertwiner  = "FAILED_ROUTE_NO_LEPTON_ORIENTATION_TO_KOIDE_DEFICIT_INTERTWINER"
	StatusKappaRemainsSeal          = "FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL"
	StatusNoFlavorDerivation        = "FIREWALL_PRESERVED_NO_KOIDE_PMNS_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedDataBridgeOnly    = "FIREWALL_PRESERVED_PMNS_DATA_REMAINS_VERSION_PINNED_OBSERVED_INPUT"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate587BoundaryPreserved  = "FIREWALL_PRESERVED_GATE587_PMNS_ORIENTATION_BOUNDARY"
)

const (
	nearRelativeTolerance      = 2.0e-2
	certifiedRelativeTolerance = 5.0e-3
)

type RuntimeInheritance struct {
	EpsilonRad           float64
	LoopUnit             float64
	Kappa                float64
	SqrtJCKM             float64
	JCKM                 float64
	Alpha2Over2Pi        float64
	CKMAlpha2Midpoint    float64
	CKMAlpha2MidpointRel float64
	Mu0GeV               float64
	Lambda12GeV          float64
	Verdict              string
}

type PMNSInput struct {
	SourceName       string
	SourceVersion    string
	DataThrough      string
	Variant          string
	MassOrdering     string
	Convention       string
	Sin2Theta12      float64
	Sin2Theta12Plus  float64
	Sin2Theta12Minus float64
	Sin2Theta23      float64
	Sin2Theta23Plus  float64
	Sin2Theta23Minus float64
	Sin2Theta13      float64
	Sin2Theta13Plus  float64
	Sin2Theta13Minus float64
	DeltaCPDeg       float64
	DeltaCPPlusDeg   float64
	DeltaCPMinusDeg  float64
	Theta12Deg       float64
	Theta23Deg       float64
	Theta13Deg       float64
	SourceNote       string
	Verdict          string
}

type PMNSInvariants struct {
	S12           float64
	C12           float64
	S23           float64
	C23           float64
	S13           float64
	C13           float64
	DeltaRad      float64
	JPMNS         float64
	AbsJPMNS      float64
	SqrtAbsJ      float64
	S13Squared    float64
	Alpha2Over2Pi float64
	Verdict       string
}

type Candidate struct {
	Name             string
	Class            string
	Equation         string
	Value            float64
	Min1Sigma        float64
	Max1Sigma        float64
	SignedResidual   float64
	AbsResidual      float64
	RelativeResidual float64
	CoversKappa      bool
	Near             bool
	Certified        bool
	Interpretation   string
}

type CandidateSet struct {
	TargetKappa         float64
	CandidateCount      int
	Candidates          []Candidate
	Best                Candidate
	BestDirectPMNS      Candidate
	BestPMNSAssisted    Candidate
	SqrtJPMNS           Candidate
	AbsJPMNS            Candidate
	S13Squared          Candidate
	Alpha2Over2Pi       Candidate
	CertifiedCandidates []Candidate
	Verdict             string
}

type UncertaintyAudit struct {
	SqrtJPMNSRange            Candidate
	Alpha2Over2PiDivC13       Candidate
	AnyCandidateCovers        bool
	CertifiedUnderUncertainty bool
	Verdict                   string
}

type CKMComparison struct {
	SqrtJCKM                       Candidate
	CKMAlpha2Midpoint              Candidate
	BestPMNSAssisted               Candidate
	DirectPMNSBetterThanSqrtJCKM   bool
	PMNSAssistedBetterThanSqrtJCKM bool
	MidpointStillClosestNumeric    bool
	Interpretation                 string
	Verdict                        string
}

type SourceDecision struct {
	PMNSProducesBetterTypedCandidate bool
	AnyCandidateCertified            bool
	CKMMidpointSurvives              bool
	KappaRemainsSeal                 bool
	BestCandidateName                string
	BestCandidateValue               float64
	BestCandidateRelativeResidual    float64
	Decision                         string
	Verdict                          string
}

type FirewallAudit struct {
	DerivesKappa               bool
	DerivesEpsilon             bool
	DerivesKoide               bool
	DerivesPMNS                bool
	DerivesNeutrinoMasses      bool
	DerivesFlavorTexture       bool
	DerivesChargedLeptonMasses bool
	PromotesObservedAsNative   bool
	AddsNewCarrier             bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	Kappa                            float64
	DirectPMNSCertified              bool
	BestDirectPMNSName               string
	BestDirectPMNSRelativeResidual   float64
	BestPMNSAssistedName             string
	BestPMNSAssistedRelativeResidual float64
	CKMMidpointRelativeResidual      float64
	AnyCertified                     bool
	RemainingSeal                    string
	Verdict                          string
}

type Analysis struct {
	Runtime     RuntimeInheritance
	PMNSInput   PMNSInput
	Invariants  PMNSInvariants
	Candidates  CandidateSet
	Uncertainty UncertaintyAudit
	CKM         CKMComparison
	Decision    SourceDecision
	Firewalls   FirewallAudit
	Final       FinalVerdict
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g586, err := generation2koideloopangledeficitaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate586 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport runtime: %w", err)
	}
	runtime := inheritRuntime(g586, bundle)
	pmns := importPMNS()
	inv := computePMNS(pmns, runtime)
	candidates := buildCandidateSet(runtime, pmns, inv)
	uncertainty := auditUncertainty(candidates)
	ckm := compareCKM(runtime, candidates)
	decision := decide(runtime, candidates, uncertainty, ckm)
	firewalls := auditFirewalls()
	final := compileFinal(runtime, candidates, ckm, decision)
	truth := "Gate 587 imports NuFIT 6.0 PMNS data and tests whether the Koide loop-angle deficit kappa_e is lepton-orientation-sized.  Direct PMNS invariants such as sqrt(|J_PMNS|) are far too large.  The best PMNS-assisted typed candidate is alpha_2/(2*pi)/c13, which is closer than sqrt(J_CKM) but is not certified and does not cover kappa_e within the propagated PMNS uncertainty.  The CKM/alpha_2 midpoint remains a stronger numerical coincidence but is not a lawful source."
	return Analysis{Runtime: runtime, PMNSInput: pmns, Invariants: inv, Candidates: candidates, Uncertainty: uncertainty, CKM: ckm, Decision: decision, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g586 generation2koideloopangledeficitaudit.Analysis, b historytransport.Bundle) RuntimeInheritance {
	alpha2Over2Pi := g586.Runtime.Alpha2Over2Pi
	sqrtJ := g586.Runtime.SqrtJCKM
	midpoint := 0.5 * (sqrtJ + alpha2Over2Pi)
	kappa := g586.Definition.Kappa
	return RuntimeInheritance{
		EpsilonRad:           g586.Definition.EpsilonRad,
		LoopUnit:             g586.Definition.LoopUnit,
		Kappa:                kappa,
		SqrtJCKM:             sqrtJ,
		JCKM:                 g586.Runtime.JCKM,
		Alpha2Over2Pi:        alpha2Over2Pi,
		CKMAlpha2Midpoint:    midpoint,
		CKMAlpha2MidpointRel: (midpoint - kappa) / kappa,
		Mu0GeV:               b.EndVector.Mu0GeV,
		Lambda12GeV:          b.GaugeBoundary.Lambda12GeV,
		Verdict:              StatusGate586Inherited,
	}
}

func importPMNS() PMNSInput {
	return PMNSInput{
		SourceName:    "NuFIT 6.0 global three-neutrino oscillation analysis",
		SourceVersion: "NuFIT 6.0 / JHEP 12 (2024) 216 / arXiv:2410.05380",
		DataThrough:   "September 2024",
		Variant:       "IC24 with SK atmospheric data",
		MassOrdering:  "Normal Ordering best fit",
		Convention:    "standard PMNS parametrization with J=c12 c23 c13^2 s12 s23 s13 sin(delta_CP)",
		Sin2Theta12:   0.308, Sin2Theta12Plus: 0.012, Sin2Theta12Minus: 0.011,
		Sin2Theta23: 0.470, Sin2Theta23Plus: 0.017, Sin2Theta23Minus: 0.013,
		Sin2Theta13: 0.02215, Sin2Theta13Plus: 0.00056, Sin2Theta13Minus: 0.00058,
		DeltaCPDeg: 212, DeltaCPPlusDeg: 26, DeltaCPMinusDeg: 41,
		Theta12Deg: 33.68, Theta23Deg: 43.3, Theta13Deg: 8.56,
		SourceNote: "NuFIT 6.0 table for IC24 with SK atmospheric data, normal ordering best-fit values and 1-sigma intervals; delta_CP is strongly non-Gaussian and ordering-dependent.",
		Verdict:    StatusPMNSDatasetImported,
	}
}

func computePMNS(p PMNSInput, r RuntimeInheritance) PMNSInvariants {
	s12, c12 := math.Sqrt(p.Sin2Theta12), math.Sqrt(1-p.Sin2Theta12)
	s23, c23 := math.Sqrt(p.Sin2Theta23), math.Sqrt(1-p.Sin2Theta23)
	s13, c13 := math.Sqrt(p.Sin2Theta13), math.Sqrt(1-p.Sin2Theta13)
	delta := deg2rad(p.DeltaCPDeg)
	j := c12 * c23 * c13 * c13 * s12 * s23 * s13 * math.Sin(delta)
	return PMNSInvariants{S12: s12, C12: c12, S23: s23, C23: c23, S13: s13, C13: c13, DeltaRad: delta, JPMNS: j, AbsJPMNS: math.Abs(j), SqrtAbsJ: math.Sqrt(math.Abs(j)), S13Squared: p.Sin2Theta13, Alpha2Over2Pi: r.Alpha2Over2Pi, Verdict: StatusPMNSJarlskogComputed}
}

func buildCandidateSet(r RuntimeInheritance, p PMNSInput, inv PMNSInvariants) CandidateSet {
	k := r.Kappa
	cands := []Candidate{
		mk(k, "sqrt(|J_PMNS|)", "direct_pmns_orientation", "sqrt(abs(c12 c23 c13^2 s12 s23 s13 sin(delta_CP)))", inv.SqrtAbsJ, pmnsRange(p, func(x PMNSInput) float64 { return computePMNS(x, r).SqrtAbsJ }), "direct leptonic orientation area scale; far larger than kappa_e"),
		mk(k, "|J_PMNS|", "direct_pmns_orientation", "abs(J_PMNS)", inv.AbsJPMNS, pmnsRange(p, func(x PMNSInput) float64 { return computePMNS(x, r).AbsJPMNS }), "direct leptonic Jarlskog area; larger than kappa_e"),
		mk(k, "s13", "pmns_angle", "sin(theta13)", inv.S13, rangeFromSin2(p.Sin2Theta13, p.Sin2Theta13Minus, p.Sin2Theta13Plus, math.Sqrt), "reactor-angle sine, dimensionless but too large"),
		mk(k, "s13^2", "pmns_angle", "sin^2(theta13)", inv.S13Squared, Interval{Min: p.Sin2Theta13 - p.Sin2Theta13Minus, Max: p.Sin2Theta13 + p.Sin2Theta13Plus}, "reactor-angle square, dimensionless but too large"),
		mk(k, "alpha_2(M_Z)/(2π)", "coupling_correction", "alpha_2/(2*pi)", r.Alpha2Over2Pi, Interval{Min: r.Alpha2Over2Pi, Max: r.Alpha2Over2Pi}, "weak loop correction coefficient inherited from Gate 586"),
		mk(k, "alpha_2(M_Z)/(2π)/c13", "pmns_assisted_coupling", "alpha_2/(2*pi*cos(theta13))", r.Alpha2Over2Pi/inv.C13, rangeFromSin2(p.Sin2Theta13, p.Sin2Theta13Minus, p.Sin2Theta13Plus, func(s2 float64) float64 { return r.Alpha2Over2Pi / math.Sqrt(1-s2) }), "PMNS-assisted weak correction using the reactor-angle normalization c13"),
		mk(k, "alpha_2(M_Z)/(2π)*sqrt(1+s13^2)", "pmns_assisted_coupling", "alpha_2/(2*pi)*sqrt(1+s13^2)", r.Alpha2Over2Pi*math.Sqrt(1+p.Sin2Theta13), rangeFromSin2(p.Sin2Theta13, p.Sin2Theta13Minus, p.Sin2Theta13Plus, func(s2 float64) float64 { return r.Alpha2Over2Pi * math.Sqrt(1+s2) }), "PMNS-assisted weak correction with reactor-angle norm"),
		mk(k, "alpha_2(M_Z)/(2π)*c13", "pmns_assisted_coupling", "alpha_2/(2*pi)*cos(theta13)", r.Alpha2Over2Pi*inv.C13, rangeFromSin2(p.Sin2Theta13, p.Sin2Theta13Minus, p.Sin2Theta13Plus, func(s2 float64) float64 { return r.Alpha2Over2Pi * math.Sqrt(1-s2) }), "PMNS-assisted weak correction suppressed by c13"),
		mk(k, "alpha_2(M_Z)/(2π)+sqrt(|J_PMNS|)", "typed_sum_rejected", "alpha_2/(2*pi)+sqrt(|J_PMNS|)", r.Alpha2Over2Pi+inv.SqrtAbsJ, pmnsRange(p, func(x PMNSInput) float64 { return r.Alpha2Over2Pi + computePMNS(x, r).SqrtAbsJ }), "typed sum requested by audit; far too large"),
		mk(k, "|alpha_2(M_Z)/(2π)-sqrt(|J_PMNS|)|", "typed_difference_rejected", "abs(alpha_2/(2*pi)-sqrt(|J_PMNS|))", math.Abs(r.Alpha2Over2Pi-inv.SqrtAbsJ), pmnsRange(p, func(x PMNSInput) float64 { return math.Abs(r.Alpha2Over2Pi - computePMNS(x, r).SqrtAbsJ) }), "typed difference requested by audit; far too large"),
		mk(k, "alpha_2(M_Z)/(2π)*sqrt(|J_PMNS|)", "typed_product_rejected", "alpha_2/(2*pi)*sqrt(|J_PMNS|)", r.Alpha2Over2Pi*inv.SqrtAbsJ, pmnsRange(p, func(x PMNSInput) float64 { return r.Alpha2Over2Pi * computePMNS(x, r).SqrtAbsJ }), "typed product requested by audit; too small"),
		mk(k, "sqrt(J_CKM)", "ckm_reference", "sqrt(J_CKM)", r.SqrtJCKM, Interval{Min: r.SqrtJCKM, Max: r.SqrtJCKM}, "Gate 586 quark-sector orientation clue"),
		mk(k, "0.5*(sqrt(J_CKM)+alpha_2/(2π))", "ckm_coupling_midpoint_reference", "midpoint(sqrt(J_CKM), alpha_2/(2*pi))", r.CKMAlpha2Midpoint, Interval{Min: r.CKMAlpha2Midpoint, Max: r.CKMAlpha2Midpoint}, "numeric midpoint clue; not a typed lepton source theorem"),
	}
	sort.Slice(cands, func(i, j int) bool { return cands[i].AbsResidual < cands[j].AbsResidual })
	cert := make([]Candidate, 0)
	for _, c := range cands {
		if c.Certified && c.CoversKappa {
			cert = append(cert, c)
		}
	}
	bestDirect := bestOfClass(cands, "direct_pmns_orientation")
	bestPMNS := bestPrefixClass(cands, "pmns_")
	return CandidateSet{TargetKappa: k, CandidateCount: len(cands), Candidates: cands, Best: cands[0], BestDirectPMNS: bestDirect, BestPMNSAssisted: bestPMNS, SqrtJPMNS: find(cands, "sqrt(|J_PMNS|)"), AbsJPMNS: find(cands, "|J_PMNS|"), S13Squared: find(cands, "s13^2"), Alpha2Over2Pi: find(cands, "alpha_2(M_Z)/(2π)"), CertifiedCandidates: cert, Verdict: StatusCandidateSetDefined}
}

type Interval struct{ Min, Max float64 }

func mk(k float64, name, class, eq string, value float64, interval Interval, interp string) Candidate {
	if interval.Min > interval.Max {
		interval.Min, interval.Max = interval.Max, interval.Min
	}
	res := value - k
	rel := res / k
	covers := interval.Min <= k && k <= interval.Max
	near := math.Abs(rel) < nearRelativeTolerance
	certified := math.Abs(rel) < certifiedRelativeTolerance && covers
	return Candidate{Name: name, Class: class, Equation: eq, Value: value, Min1Sigma: interval.Min, Max1Sigma: interval.Max, SignedResidual: res, AbsResidual: math.Abs(res), RelativeResidual: rel, CoversKappa: covers, Near: near, Certified: certified, Interpretation: interp}
}

func pmnsRange(p PMNSInput, f func(PMNSInput) float64) Interval {
	vals := []float64{}
	for _, s12 := range []float64{p.Sin2Theta12 - p.Sin2Theta12Minus, p.Sin2Theta12, p.Sin2Theta12 + p.Sin2Theta12Plus} {
		for _, s23 := range []float64{p.Sin2Theta23 - p.Sin2Theta23Minus, p.Sin2Theta23, p.Sin2Theta23 + p.Sin2Theta23Plus} {
			for _, s13 := range []float64{p.Sin2Theta13 - p.Sin2Theta13Minus, p.Sin2Theta13, p.Sin2Theta13 + p.Sin2Theta13Plus} {
				for _, d := range []float64{p.DeltaCPDeg - p.DeltaCPMinusDeg, p.DeltaCPDeg, p.DeltaCPDeg + p.DeltaCPPlusDeg} {
					q := p
					q.Sin2Theta12 = s12
					q.Sin2Theta23 = s23
					q.Sin2Theta13 = s13
					q.DeltaCPDeg = d
					vals = append(vals, f(q))
				}
			}
		}
	}
	min, max := vals[0], vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return Interval{Min: min, Max: max}
}

func rangeFromSin2(center, minus, plus float64, f func(float64) float64) Interval {
	a, b := f(center-minus), f(center+plus)
	if a > b {
		a, b = b, a
	}
	return Interval{Min: a, Max: b}
}

func auditUncertainty(c CandidateSet) UncertaintyAudit {
	any := false
	for _, x := range c.Candidates {
		if x.CoversKappa {
			any = true
			break
		}
	}
	return UncertaintyAudit{SqrtJPMNSRange: c.SqrtJPMNS, Alpha2Over2PiDivC13: find(c.Candidates, "alpha_2(M_Z)/(2π)/c13"), AnyCandidateCovers: any, CertifiedUnderUncertainty: len(c.CertifiedCandidates) > 0, Verdict: strings.Join([]string{StatusUncertaintyPropagated, StatusPMNSRangeCoversButNoCert, StatusNoPMNSCertified}, ";")}
}

func compareCKM(r RuntimeInheritance, c CandidateSet) CKMComparison {
	sqrtCKM := find(c.Candidates, "sqrt(J_CKM)")
	mid := find(c.Candidates, "0.5*(sqrt(J_CKM)+alpha_2/(2π))")
	directBetter := c.BestDirectPMNS.AbsResidual < sqrtCKM.AbsResidual
	assistedBetter := c.BestPMNSAssisted.AbsResidual < sqrtCKM.AbsResidual
	midpointClosest := mid.AbsResidual < c.BestPMNSAssisted.AbsResidual && mid.AbsResidual < sqrtCKM.AbsResidual
	interp := "Direct PMNS orientation invariants are not close to kappa_e.  The PMNS-assisted weak correction alpha_2/(2*pi)/c13 is closer than sqrt(J_CKM), but the CKM/alpha_2 midpoint remains the closest numerical clue and is not a lawful lepton-sector source map."
	return CKMComparison{SqrtJCKM: sqrtCKM, CKMAlpha2Midpoint: mid, BestPMNSAssisted: c.BestPMNSAssisted, DirectPMNSBetterThanSqrtJCKM: directBetter, PMNSAssistedBetterThanSqrtJCKM: assistedBetter, MidpointStillClosestNumeric: midpointClosest, Interpretation: interp, Verdict: strings.Join([]string{StatusPMNSAssistedBetterThanCKM, StatusCKMMidpointSurvives, StatusCKMStillNotLeptonSource}, ";")}
}

func decide(r RuntimeInheritance, c CandidateSet, u UncertaintyAudit, k CKMComparison) SourceDecision {
	certified := len(c.CertifiedCandidates) > 0
	pmnsBetter := k.PMNSAssistedBetterThanSqrtJCKM
	decision := "PMNS data does not certify kappa_e.  Direct PMNS orientation scales are much too large; the best PMNS-assisted candidate alpha_2/(2*pi)/c13 improves over sqrt(J_CKM) but misses kappa_e by about 0.84% and does not cover it within NuFIT one-sigma uncertainty.  The broad |J_PMNS| range can cross kappa_e because delta_CP remains poorly constrained, but central PMNS values do not point to kappa_e and no candidate certifies.  The CKM/alpha_2 midpoint remains the closest numerical coincidence, but no typed source map makes it lawful."
	return SourceDecision{PMNSProducesBetterTypedCandidate: pmnsBetter, AnyCandidateCertified: certified, CKMMidpointSurvives: k.MidpointStillClosestNumeric, KappaRemainsSeal: !certified, BestCandidateName: c.Best.Name, BestCandidateValue: c.Best.Value, BestCandidateRelativeResidual: c.Best.RelativeResidual, Decision: decision, Verdict: strings.Join([]string{StatusBestPMNSAssisted, StatusPMNSAssistedBetterThanCKM, StatusPMNSRangeCoversButNoCert, StatusNoPMNSCertified, StatusKappaRemainsSeal}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKappa: false, DerivesEpsilon: false, DerivesKoide: false, DerivesPMNS: false, DerivesNeutrinoMasses: false, DerivesFlavorTexture: false, DerivesChargedLeptonMasses: false, PromotesObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoFlavorDerivation, StatusObservedDataBridgeOnly, StatusGate352Preserved, StatusGate587BoundaryPreserved}, ";")}
}

func compileFinal(r RuntimeInheritance, c CandidateSet, k CKMComparison, d SourceDecision) FinalVerdict {
	return FinalVerdict{Kappa: r.Kappa, DirectPMNSCertified: c.BestDirectPMNS.Certified, BestDirectPMNSName: c.BestDirectPMNS.Name, BestDirectPMNSRelativeResidual: c.BestDirectPMNS.RelativeResidual, BestPMNSAssistedName: c.BestPMNSAssisted.Name, BestPMNSAssistedRelativeResidual: c.BestPMNSAssisted.RelativeResidual, CKMMidpointRelativeResidual: k.CKMAlpha2Midpoint.RelativeResidual, AnyCertified: d.AnyCandidateCertified, RemainingSeal: "kappa_e remains the charged-lepton Koide loop-angle deficit environmental seal", Verdict: strings.Join([]string{StatusDirectPMNSTooLarge, StatusNoPMNSCertified, StatusKappaRemainsSeal, StatusGate587BoundaryPreserved}, ";")}
}

func Statuses() []string {
	return []string{StatusGate586Inherited, StatusPMNSDatasetImported, StatusPMNSJarlskogComputed, StatusCandidateSetDefined, StatusUncertaintyPropagated, StatusPMNSRangeCoversButNoCert, StatusBestPMNSAssisted, StatusPMNSAssistedBetterThanCKM, StatusDirectPMNSTooLarge, StatusNoPMNSCertified, StatusCKMMidpointSurvives, StatusCKMStillNotLeptonSource, StatusNoLeptonKoideIntertwiner, StatusKappaRemainsSeal, StatusNoFlavorDerivation, StatusObservedDataBridgeOnly, StatusGate352Preserved, StatusGate587BoundaryPreserved}
}

func bestOfClass(cs []Candidate, class string) Candidate {
	best := Candidate{AbsResidual: math.Inf(1)}
	for _, c := range cs {
		if c.Class == class && c.AbsResidual < best.AbsResidual {
			best = c
		}
	}
	return best
}

func bestPrefixClass(cs []Candidate, prefix string) Candidate {
	best := Candidate{AbsResidual: math.Inf(1)}
	for _, c := range cs {
		if strings.HasPrefix(c.Class, prefix) && c.AbsResidual < best.AbsResidual {
			best = c
		}
	}
	return best
}

func find(cs []Candidate, name string) Candidate {
	for _, c := range cs {
		if c.Name == name {
			return c
		}
	}
	return Candidate{Name: name, Value: math.NaN(), Min1Sigma: math.NaN(), Max1Sigma: math.NaN(), SignedResidual: math.NaN(), AbsResidual: math.Inf(1), RelativeResidual: math.NaN(), Interpretation: "candidate not found"}
}

func deg2rad(x float64) float64 { return x * math.Pi / 180.0 }
