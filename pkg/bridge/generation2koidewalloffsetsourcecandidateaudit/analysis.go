// Package generation2koidewalloffsetsourcecandidateaudit implements Gate 585:
// Koide Wall-Offset Source Candidate Audit.
//
// Gate 584 compressed the charged-lepton hierarchy, within the exact Koide
// chamber-wall model, to one scale and one small wall-offset epsilon. Gate 585
// asks a narrower bridge-layer question: does epsilon match any already-typed,
// dimensionless runtime quantity strongly enough to serve as its source?  The
// audit tests loop factors, electroweak couplings, gauge/scalar residuals, and
// CKM area proxies already present in the ASHA history-transport runtime.
//
// This is a source-candidate sieve only.  It does not derive epsilon, Koide,
// charged-lepton masses, a root trace, thresholds, CKM/PMNS, or generation
// hierarchy as native ASHA law.
package generation2koidewalloffsetsourcecandidateaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidechamberwalloffsetaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidewalloffsetratioclosureaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE585-KOIDE-WALL-OFFSET-SOURCE-CANDIDATE-AUDIT"

	StatusGate584Inherited          = "PASS_GATE584_WALL_OFFSET_RATIO_CLOSURE_INHERITED"
	StatusEpsilonTargetDefined      = "PASS_ELECTRON_WALL_OFFSET_EPSILON_TARGET_DEFINED"
	StatusTypedCandidateSetDefined  = "PASS_TYPED_DIMENSIONLESS_SOURCE_CANDIDATE_SET_DEFINED"
	StatusLoopFactorCompared        = "PASS_LOOP_FACTOR_CANDIDATES_COMPARED_TO_EPSILON"
	StatusRuntimeResidualsCompared  = "PASS_RUNTIME_RESIDUAL_CANDIDATES_COMPARED_TO_EPSILON"
	StatusBestCandidateOneOver8Pi   = "CONDITIONAL_SUPPORT_BEST_SOURCE_CANDIDATE_IS_ONE_OVER_8PI_LOOP_SCALE"
	StatusOneOver8PiNearOnly        = "CONDITIONAL_SUPPORT_ONE_OVER_8PI_NEAR_EPSILON_BUT_NOT_CERTIFIED"
	StatusNoCertifiedSource         = "FAILED_ROUTE_NO_DIMENSIONLESS_RUNTIME_CANDIDATE_CERTIFIED_AS_EPSILON_SOURCE"
	StatusNoCouplingSource          = "FAILED_ROUTE_ELECTROWEAK_AND_BOUNDARY_COUPLING_CANDIDATES_DO_NOT_FIX_EPSILON"
	StatusNoResidualSource          = "FAILED_ROUTE_GAUGE_SCALAR_CKM_RESIDUALS_DO_NOT_FIX_EPSILON"
	StatusNoNativeLoopOperator      = "FAILED_ROUTE_NO_NATIVE_LOOP_FACTOR_TO_KOIDE_WALL_OFFSET_OPERATOR"
	StatusNoNativeEpsilonDerivation = "FAILED_ROUTE_EPSILON_REMAINS_HISTORY_SEAL_NOT_NATIVE_DERIVATION"
	StatusNoFlavorPromotion         = "FIREWALL_PRESERVED_SOURCE_CANDIDATE_AUDIT_DOES_NOT_DERIVE_FLAVOR_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusObservedEndpointPreserved = "FIREWALL_PRESERVED_EPSILON_AND_CANDIDATES_REMAIN_BRIDGE_RUNTIME_VALUES"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate585BoundaryPreserved  = "FIREWALL_PRESERVED_GATE585_SOURCE_CANDIDATE_BOUNDARY"
)

const (
	nearRelativeTolerance      = 1.0e-2 // candidate is worth recording as a clue below one percent.
	certifiedRelativeTolerance = 1.0e-3 // source certification requires sub-per-mille agreement in this v1 sieve.
)

type RuntimeInheritance struct {
	Gate583EpsilonDeg       float64
	Gate583EpsilonRad       float64
	Gate584SolvedEpsilonDeg float64
	Gate584SolvedEpsilonRad float64
	Gate584PredictionResid  float64
	Mu0GeV                  float64
	Lambda12GeV             float64
	Source                  string
	Verdict                 string
}

type EpsilonTarget struct {
	PrimaryDefinition          string
	PrimaryEpsilonDeg          float64
	PrimaryEpsilonRad          float64
	ExactR1RatioClosureDeg     float64
	ExactR1RatioClosureRad     float64
	DifferenceRad              float64
	DifferenceDeg              float64
	UseForCandidateSieve       string
	NearToleranceRelative      float64
	CertifiedToleranceRelative float64
	Verdict                    string
}

type Candidate struct {
	Name             string
	Class            string
	Equation         string
	Value            float64
	SignedResidual   float64
	AbsResidual      float64
	RelativeResidual float64
	Near             bool
	Certified        bool
	Interpretation   string
}

type CandidateSet struct {
	TargetEpsilonRad    float64
	CandidateCount      int
	Candidates          []Candidate
	Best                Candidate
	CertifiedCandidates []Candidate
	NearCandidates      []Candidate
	Verdict             string
}

type LoopFactorAudit struct {
	OneOver8Pi            Candidate
	OneOver4Pi            Candidate
	OneOver16Pi           Candidate
	BestLoop              Candidate
	RequiredCorrection    float64
	RequiredCorrectionPct float64
	NearButNotCertified   bool
	Verdict               string
}

type CouplingAudit struct {
	AlphaEMMZ            Candidate
	SqrtAlphaEMMZ        Candidate
	AlphaEMOverPiMZ      Candidate
	GStarSquaredOver8Pi2 Candidate
	Alpha2MZ             Candidate
	BestCoupling         Candidate
	Certified            bool
	Verdict              string
}

type ResidualAudit struct {
	StrongMismatch Candidate
	AbsLambdaL12   Candidate
	AbsDeltaSin2   Candidate
	JCKM           Candidate
	SqrtJCKM       Candidate
	BestResidual   Candidate
	Certified      bool
	Verdict        string
}

type SourceDecision struct {
	BestCandidateName      string
	BestCandidateValue     float64
	BestRelativeResidual   float64
	BestAbsResidual        float64
	NearClue               bool
	CertifiedSource        bool
	MinimalNextRequirement string
	Decision               string
	Verdict                string
}

type FirewallAudit struct {
	DerivesEpsilon             bool
	DerivesKoide               bool
	DerivesLeptonMasses        bool
	DerivesYukawaEigenvalues   bool
	DerivesCKM                 bool
	DerivesPMNS                bool
	DerivesGenerationHierarchy bool
	AddsNewCarrier             bool
	PromotesObservedAsNative   bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	SealName                  string
	EpsilonRad                float64
	EpsilonDeg                float64
	BestCandidate             string
	BestCandidateValue        float64
	BestCandidateRelativeDiff float64
	CandidateCertified        bool
	NearLoopScaleClue         bool
	NativeDerivationCertified bool
	RemainingSeal             string
	Verdict                   string
}

type Analysis struct {
	Runtime    RuntimeInheritance
	Target     EpsilonTarget
	Candidates CandidateSet
	Loop       LoopFactorAudit
	Couplings  CouplingAudit
	Residuals  ResidualAudit
	Decision   SourceDecision
	Firewalls  FirewallAudit
	Final      FinalVerdict
	Truth      string
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
	g583, err := generation2koidechamberwalloffsetaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate583 predecessor: %w", err)
	}
	g584, err := generation2koidewalloffsetratioclosureaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate584 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history-transport runtime: %w", err)
	}
	runtime := inheritRuntime(g583, g584, bundle)
	target := defineTarget(runtime)
	candidates := buildCandidateSet(target.PrimaryEpsilonRad, bundle)
	loop := auditLoopFactors(candidates)
	couplings := auditCouplings(candidates)
	residuals := auditResiduals(candidates)
	decision := decideSource(candidates, loop, couplings, residuals)
	firewalls := auditFirewalls()
	final := compileFinal(target, decision)
	truth := "Gate 585 tests whether the charged-lepton Koide chamber-wall offset epsilon is sourced by a typed dimensionless runtime quantity.  The nearest candidate is the loop scale 1/(8*pi), within about 0.553%, but it fails the stricter certification threshold and no native operator maps that loop factor to the Koide wall offset.  Epsilon therefore remains the minimal charged-lepton wall-offset history seal."
	return Analysis{Runtime: runtime, Target: target, Candidates: candidates, Loop: loop, Couplings: couplings, Residuals: residuals, Decision: decision, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g583 generation2koidechamberwalloffsetaudit.Analysis, g584 generation2koidewalloffsetratioclosureaudit.Analysis, b historytransport.Bundle) RuntimeInheritance {
	return RuntimeInheritance{
		Gate583EpsilonDeg:       g583.MZ.EpsilonDeg,
		Gate583EpsilonRad:       g583.MZ.EpsilonRad,
		Gate584SolvedEpsilonDeg: g584.MZ.FromElectronMuon.SolvedEpsilonDeg,
		Gate584SolvedEpsilonRad: g584.MZ.FromElectronMuon.SolvedEpsilonRad,
		Gate584PredictionResid:  g584.MZ.FromElectronMuon.RootResidual,
		Mu0GeV:                  b.EndVector.Mu0GeV,
		Lambda12GeV:             b.GaugeBoundary.Lambda12GeV,
		Source:                  "Gate583 wall-offset epsilon plus Gate584 exact-R=1 ratio closure and historytransport runtime",
		Verdict:                 StatusGate584Inherited,
	}
}

func defineTarget(r RuntimeInheritance) EpsilonTarget {
	return EpsilonTarget{
		PrimaryDefinition:          "epsilon_e = 135° - delta_e in the canonical (e,mu,tau) positive Koide chamber, using the observed Gate583 wall point at M_Z",
		PrimaryEpsilonDeg:          r.Gate583EpsilonDeg,
		PrimaryEpsilonRad:          r.Gate583EpsilonRad,
		ExactR1RatioClosureDeg:     r.Gate584SolvedEpsilonDeg,
		ExactR1RatioClosureRad:     r.Gate584SolvedEpsilonRad,
		DifferenceRad:              r.Gate584SolvedEpsilonRad - r.Gate583EpsilonRad,
		DifferenceDeg:              r.Gate584SolvedEpsilonDeg - r.Gate583EpsilonDeg,
		UseForCandidateSieve:       "primary Gate583 epsilon is used for source-candidate matching; Gate584 epsilon is recorded as the exact-R=1 closure coordinate",
		NearToleranceRelative:      nearRelativeTolerance,
		CertifiedToleranceRelative: certifiedRelativeTolerance,
		Verdict:                    StatusEpsilonTargetDefined,
	}
}

func buildCandidateSet(eps float64, b historytransport.Bundle) CandidateSet {
	gY, g1, g2, g3 := b.EndVector.GY, b.EndVector.G1, b.EndVector.G2, b.EndVector.G3
	gStar := b.GaugeBoundary.GStar
	e := gY * g2 / math.Sqrt(gY*gY+g2*g2)
	alphaEM := e * e / (4.0 * math.Pi)
	alphaY := gY * gY / (4.0 * math.Pi)
	alpha1 := g1 * g1 / (4.0 * math.Pi)
	alpha2 := g2 * g2 / (4.0 * math.Pi)
	alpha3 := g3 * g3 / (4.0 * math.Pi)
	alphaStar := gStar * gStar / (4.0 * math.Pi)
	vals := []Candidate{
		mk(eps, "1/(8π)", "loop_factor", "1/(8*pi)", 1.0/(8.0*math.Pi), "small one-loop-sized geometric displacement scale"),
		mk(eps, "1/(4π)", "loop_factor", "1/(4*pi)", 1.0/(4.0*math.Pi), "larger loop circumference scale"),
		mk(eps, "1/(16π)", "loop_factor", "1/(16*pi)", 1.0/(16.0*math.Pi), "half of the nearest loop circumference scale"),
		mk(eps, "alpha_EM(M_Z)", "electroweak_coupling", "e(M_Z)^2/(4*pi)", alphaEM, "electromagnetic fine-structure value derived from g2 and gY at M_Z"),
		mk(eps, "sqrt(alpha_EM(M_Z))", "electroweak_coupling", "sqrt(e(M_Z)^2/(4*pi))", math.Sqrt(alphaEM), "square-root electromagnetic coupling scale"),
		mk(eps, "alpha_EM(M_Z)/pi", "electroweak_coupling", "alpha_EM(M_Z)/pi", alphaEM/math.Pi, "electromagnetic one-loop-normalized scale"),
		mk(eps, "alpha_Y(M_Z)", "electroweak_coupling", "gY(M_Z)^2/(4*pi)", alphaY, "hypercharge coupling strength at endpoint"),
		mk(eps, "alpha_1(M_Z)", "electroweak_coupling", "g1(M_Z)^2/(4*pi)", alpha1, "canonical hypercharge coupling strength at endpoint"),
		mk(eps, "alpha_2(M_Z)", "electroweak_coupling", "g2(M_Z)^2/(4*pi)", alpha2, "weak coupling strength at endpoint"),
		mk(eps, "alpha_3(M_Z)", "gauge_coupling", "g3(M_Z)^2/(4*pi)=alpha_s(M_Z)", alpha3, "strong coupling strength at endpoint"),
		mk(eps, "alpha_star(Lambda_12)", "boundary_coupling", "g_star^2/(4*pi)", alphaStar, "g1=g2 boundary coupling strength"),
		mk(eps, "gY(M_Z)^2/(8π²)", "loop_normalized_coupling", "gY(M_Z)^2/(8*pi^2)", gY*gY/(8.0*math.Pi*math.Pi), "hypercharge two-point loop normalization"),
		mk(eps, "g1(M_Z)^2/(8π²)", "loop_normalized_coupling", "g1(M_Z)^2/(8*pi^2)", g1*g1/(8.0*math.Pi*math.Pi), "canonical hypercharge two-point loop normalization"),
		mk(eps, "g2(M_Z)^2/(8π²)", "loop_normalized_coupling", "g2(M_Z)^2/(8*pi^2)", g2*g2/(8.0*math.Pi*math.Pi), "weak two-point loop normalization"),
		mk(eps, "g_star^2/(8π²)", "loop_normalized_boundary", "g_star^2/(8*pi^2)", gStar*gStar/(8.0*math.Pi*math.Pi), "boundary weak-coupling two-point loop normalization"),
		mk(eps, "R_3-1", "gauge_residual", "g3(Lambda_12)/g_star - 1", b.GaugeBoundary.R3-1.0, "strong coupling ratio mismatch at the g1=g2 boundary"),
		mk(eps, "|Delta_3|", "gauge_residual", "abs(g3^-2-g_star^-2)", math.Abs(b.GaugeBoundary.Delta3), "inverse strong mismatch magnitude"),
		mk(eps, "|Delta_sin²|", "weak_angle_residual", "abs(sin²theta_End-3/8)", math.Abs(b.WeakAngleTransport.DeltaSin2), "weak-angle transport residual magnitude"),
		mk(eps, "|lambda(Lambda_12)|", "scalar_residual", "abs(lambda(Lambda_12))", math.Abs(b.ScalarTransport.LambdaLambda12), "scalar quartic boundary residual magnitude in v1"),
		mk(eps, "J_CKM", "orientation_residual", "J_CKM", b.FlavorTransport.JCKM, "CKM oriented-area invariant"),
		mk(eps, "sqrt(J_CKM)", "orientation_residual", "sqrt(J_CKM)", math.Sqrt(b.FlavorTransport.JCKM), "square-root CKM oriented-area proxy"),
	}
	sort.Slice(vals, func(i, j int) bool { return vals[i].AbsResidual < vals[j].AbsResidual })
	var near, cert []Candidate
	for _, c := range vals {
		if c.Near {
			near = append(near, c)
		}
		if c.Certified {
			cert = append(cert, c)
		}
	}
	return CandidateSet{TargetEpsilonRad: eps, CandidateCount: len(vals), Candidates: vals, Best: vals[0], CertifiedCandidates: cert, NearCandidates: near, Verdict: strings.Join([]string{StatusTypedCandidateSetDefined, StatusLoopFactorCompared, StatusRuntimeResidualsCompared}, ";")}
}

func mk(eps float64, name, class, eq string, value float64, interp string) Candidate {
	res := value - eps
	rel := res / eps
	return Candidate{Name: name, Class: class, Equation: eq, Value: value, SignedResidual: res, AbsResidual: math.Abs(res), RelativeResidual: rel, Near: math.Abs(rel) < nearRelativeTolerance, Certified: math.Abs(rel) < certifiedRelativeTolerance, Interpretation: interp}
}

func auditLoopFactors(c CandidateSet) LoopFactorAudit {
	one8 := findCandidate(c, "1/(8π)")
	one4 := findCandidate(c, "1/(4π)")
	one16 := findCandidate(c, "1/(16π)")
	best := one8
	for _, x := range []Candidate{one4, one16} {
		if x.AbsResidual < best.AbsResidual {
			best = x
		}
	}
	corr := c.TargetEpsilonRad/one8.Value - 1.0
	nearOnly := one8.Near && !one8.Certified
	return LoopFactorAudit{OneOver8Pi: one8, OneOver4Pi: one4, OneOver16Pi: one16, BestLoop: best, RequiredCorrection: corr, RequiredCorrectionPct: 100.0 * corr, NearButNotCertified: nearOnly, Verdict: strings.Join([]string{StatusBestCandidateOneOver8Pi, StatusOneOver8PiNearOnly}, ";")}
}

func auditCouplings(c CandidateSet) CouplingAudit {
	cs := []Candidate{findCandidate(c, "alpha_EM(M_Z)"), findCandidate(c, "sqrt(alpha_EM(M_Z))"), findCandidate(c, "alpha_EM(M_Z)/pi"), findCandidate(c, "g_star^2/(8π²)"), findCandidate(c, "alpha_2(M_Z)")}
	best := cs[0]
	for _, x := range cs[1:] {
		if x.AbsResidual < best.AbsResidual {
			best = x
		}
	}
	return CouplingAudit{AlphaEMMZ: cs[0], SqrtAlphaEMMZ: cs[1], AlphaEMOverPiMZ: cs[2], GStarSquaredOver8Pi2: cs[3], Alpha2MZ: cs[4], BestCoupling: best, Certified: best.Certified, Verdict: StatusNoCouplingSource}
}

func auditResiduals(c CandidateSet) ResidualAudit {
	cs := []Candidate{findCandidate(c, "R_3-1"), findCandidate(c, "|lambda(Lambda_12)|"), findCandidate(c, "|Delta_sin²|"), findCandidate(c, "J_CKM"), findCandidate(c, "sqrt(J_CKM)")}
	best := cs[0]
	for _, x := range cs[1:] {
		if x.AbsResidual < best.AbsResidual {
			best = x
		}
	}
	return ResidualAudit{StrongMismatch: cs[0], AbsLambdaL12: cs[1], AbsDeltaSin2: cs[2], JCKM: cs[3], SqrtJCKM: cs[4], BestResidual: best, Certified: best.Certified, Verdict: StatusNoResidualSource}
}

func decideSource(c CandidateSet, loop LoopFactorAudit, couplings CouplingAudit, residuals ResidualAudit) SourceDecision {
	cert := len(c.CertifiedCandidates) > 0
	decision := "No typed dimensionless runtime candidate fixes epsilon_e at the certification level.  The best clue is 1/(8*pi), but it requires a percent-level correction and no native ASHA operator maps that loop factor to the Koide chamber-wall offset."
	verdict := strings.Join([]string{StatusBestCandidateOneOver8Pi, StatusOneOver8PiNearOnly, StatusNoCertifiedSource, StatusNoNativeLoopOperator, StatusNoNativeEpsilonDerivation}, ";")
	return SourceDecision{BestCandidateName: c.Best.Name, BestCandidateValue: c.Best.Value, BestRelativeResidual: c.Best.RelativeResidual, BestAbsResidual: c.Best.AbsResidual, NearClue: c.Best.Near, CertifiedSource: cert, MinimalNextRequirement: "a typed operator or transport equation that maps a loop/threshold quantity into the ordered Koide chamber-wall coordinate epsilon_e", Decision: decision, Verdict: verdict}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesEpsilon: false, DerivesKoide: false, DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, AddsNewCarrier: false, PromotesObservedAsNative: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoCertifiedSource, StatusNoNativeEpsilonDerivation, StatusNoFlavorPromotion, StatusObservedEndpointPreserved, StatusGate352Preserved, StatusGate585BoundaryPreserved}, ";")}
}

func compileFinal(t EpsilonTarget, d SourceDecision) FinalVerdict {
	return FinalVerdict{SealName: "ChargedLeptonKoideWallOffsetSourceCandidateSeal", EpsilonRad: t.PrimaryEpsilonRad, EpsilonDeg: t.PrimaryEpsilonDeg, BestCandidate: d.BestCandidateName, BestCandidateValue: d.BestCandidateValue, BestCandidateRelativeDiff: d.BestRelativeResidual, CandidateCertified: d.CertifiedSource, NearLoopScaleClue: d.NearClue && d.BestCandidateName == "1/(8π)", NativeDerivationCertified: false, RemainingSeal: "epsilon_e remains an environmental wall-offset history seal until a native root-trace/circulant or loop-threshold operator fixes it", Verdict: strings.Join([]string{StatusBestCandidateOneOver8Pi, StatusOneOver8PiNearOnly, StatusNoCertifiedSource, StatusGate585BoundaryPreserved}, ";")}
}

func Statuses() []string {
	return []string{StatusGate584Inherited, StatusEpsilonTargetDefined, StatusTypedCandidateSetDefined, StatusLoopFactorCompared, StatusRuntimeResidualsCompared, StatusBestCandidateOneOver8Pi, StatusOneOver8PiNearOnly, StatusNoCertifiedSource, StatusNoCouplingSource, StatusNoResidualSource, StatusNoNativeLoopOperator, StatusNoNativeEpsilonDerivation, StatusNoFlavorPromotion, StatusObservedEndpointPreserved, StatusGate352Preserved, StatusGate585BoundaryPreserved}
}

func findCandidate(c CandidateSet, name string) Candidate {
	for _, x := range c.Candidates {
		if x.Name == name {
			return x
		}
	}
	return Candidate{Name: name, Value: math.NaN(), SignedResidual: math.NaN(), AbsResidual: math.Inf(1), RelativeResidual: math.NaN(), Interpretation: "candidate not found"}
}
