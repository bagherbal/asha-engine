// Package generation2koideloopangledeficitaudit implements Gate 586:
// Koide Loop-Angle Deficit Audit.
//
// Gate 585 identified the nearest typed source candidate for the charged-lepton
// Koide electron-wall offset epsilon_e as the loop-sized angle L=1/(8*pi), but
// did not certify it.  Gate 586 therefore factors the observed wall offset as
// epsilon_e=L(1-kappa_e) and audits the remaining loop-angle deficit kappa_e.
//
// This is a correction-candidate sieve only.  It does not derive epsilon_e,
// kappa_e, Koide, charged-lepton masses, CKM/PMNS, thresholds, a root trace, or
// a native ASHA flavor operator.
package generation2koideloopangledeficitaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidechamberwalloffsetaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidetransportvectordecompositionaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidewalloffsetsourcecandidateaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE586-KOIDE-LOOP-ANGLE-DEFICIT-AUDIT"

	StatusGate585Inherited          = "PASS_GATE585_LOOP_SCALE_CANDIDATE_INHERITED"
	StatusLoopDeficitDefined        = "PASS_KOIDE_LOOP_ANGLE_DEFICIT_KAPPA_DEFINED"
	StatusTypedCandidateSetDefined  = "PASS_TYPED_KAPPA_CANDIDATE_SET_DEFINED"
	StatusOrientationCompared       = "PASS_ORIENTATION_SCALE_CANDIDATES_COMPARED_TO_KAPPA"
	StatusTransportCompared         = "PASS_TRANSPORT_DRIFT_CANDIDATES_COMPARED_TO_KAPPA"
	StatusCouplingCompared          = "PASS_COUPLING_AND_RESIDUAL_CORRECTION_CANDIDATES_COMPARED_TO_KAPPA"
	StatusBestSqrtJCKM              = "CONDITIONAL_SUPPORT_BEST_KAPPA_CANDIDATE_IS_SQRT_J_CKM"
	StatusSqrtJCKMNearOnly          = "CONDITIONAL_SUPPORT_SQRT_J_CKM_NEAR_KAPPA_BUT_NOT_CERTIFIED"
	StatusAlpha2Over2PiNearOnly     = "CONDITIONAL_SUPPORT_ALPHA2_OVER_2PI_NEAR_KAPPA_BUT_NOT_CERTIFIED"
	StatusNoCertifiedKappaSource    = "FAILED_ROUTE_NO_TYPED_RUNTIME_QUANTITY_CERTIFIED_AS_KAPPA_SOURCE"
	StatusCKMNotLeptonSource        = "FAILED_ROUTE_CKM_ORIENTATION_PROXY_NOT_LAWFUL_CHARGED_LEPTON_SOURCE_WITHOUT_INTERTWINER"
	StatusNoPMNSRuntime             = "FAILED_ROUTE_NO_PMNS_RUNTIME_INPUT_FOR_LEPTON_ORIENTATION_DEFICIT_AUDIT"
	StatusTransportNoFix            = "FAILED_ROUTE_TRANSPORT_DRIFT_AND_R_MINUS_ONE_DO_NOT_FIX_KAPPA"
	StatusCouplingResidualNoFix     = "FAILED_ROUTE_GAUGE_SCALAR_COUPLING_CORRECTIONS_DO_NOT_FIX_KAPPA"
	StatusNoNativeDeficitOperator   = "FAILED_ROUTE_NO_NATIVE_LOOP_DEFICIT_TO_KOIDE_WALL_OPERATOR"
	StatusKappaRemainsSeal          = "FAILED_ROUTE_KAPPA_E_REMAINS_HISTORY_SEAL_NOT_NATIVE_DERIVATION"
	StatusNoFlavorPromotion         = "FIREWALL_PRESERVED_KAPPA_AUDIT_DOES_NOT_DERIVE_FLAVOR_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusObservedEndpointPreserved = "FIREWALL_PRESERVED_KAPPA_AND_CANDIDATES_REMAIN_BRIDGE_RUNTIME_VALUES"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate586BoundaryPreserved  = "FIREWALL_PRESERVED_GATE586_LOOP_ANGLE_DEFICIT_BOUNDARY"
)

const (
	nearRelativeTolerance      = 2.0e-2 // record percent-level correction hints.
	certifiedRelativeTolerance = 5.0e-3 // certification requires better than half-percent agreement plus a typed source map.
)

type RuntimeInheritance struct {
	EpsilonRad             float64
	EpsilonDeg             float64
	LoopUnit               float64
	Gate585BestCandidate   string
	Gate585BestRelative    float64
	KappaFromGate585       float64
	JCKM                   float64
	SqrtJCKM               float64
	Alpha2MZ               float64
	Alpha2Over2Pi          float64
	ProjectiveDriftRad     float64
	DeltaPhiRad            float64
	DeltaThetaRad          float64
	DeltaEpsilonRad        float64
	KoideAmplitudeResidual float64
	Mu0GeV                 float64
	Lambda12GeV            float64
	Verdict                string
}

type DeficitDefinition struct {
	Formula             string
	LoopUnit            float64
	EpsilonRad          float64
	EpsilonDeg          float64
	Kappa               float64
	KappaPercent        float64
	ReconstructedEpsRad float64
	ReconstructionError float64
	Interpretation      string
	NearTolerance       float64
	CertifiedTolerance  float64
	Verdict             string
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
	TargetKappa         float64
	CandidateCount      int
	Candidates          []Candidate
	Best                Candidate
	NearCandidates      []Candidate
	CertifiedCandidates []Candidate
	Verdict             string
}

type OrientationAudit struct {
	JCKM             Candidate
	SqrtJCKM         Candidate
	BestOrientation  Candidate
	NearButNotSource bool
	PMNSRuntimeInput bool
	Verdict          string
}

type TransportAudit struct {
	ProjectiveDrift Candidate
	DeltaPhi        Candidate
	DeltaTheta      Candidate
	DeltaEpsilon    Candidate
	KoideRMinusOne  Candidate
	BestTransport   Candidate
	Certified       bool
	Verdict         string
}

type CorrectionScaleAudit struct {
	Alpha2Over2Pi         Candidate
	AlphaEM               Candidate
	AlphaEMOverPi         Candidate
	StrongMismatchOver2Pi Candidate
	DeltaSin2Over8Pi      Candidate
	LambdaOver2Pi         Candidate
	BestCorrection        Candidate
	Certified             bool
	Verdict               string
}

type SourceDecision struct {
	BestCandidateName      string
	BestCandidateValue     float64
	BestRelativeResidual   float64
	BestAbsResidual        float64
	NearClue               bool
	CertifiedSource        bool
	CandidateMeaning       string
	MinimalNextRequirement string
	Decision               string
	Verdict                string
}

type FirewallAudit struct {
	DerivesKappa             bool
	DerivesEpsilon           bool
	DerivesKoide             bool
	DerivesLeptonMasses      bool
	DerivesYukawaEigenvalues bool
	DerivesCKM               bool
	DerivesPMNS              bool
	AddsNewCarrier           bool
	PromotesObservedAsNative bool
	PreservesGate352         bool
	Verdict                  string
}

type FinalVerdict struct {
	SealName                  string
	EpsilonRad                float64
	LoopUnit                  float64
	Kappa                     float64
	BestCandidate             string
	BestCandidateValue        float64
	BestCandidateRelativeDiff float64
	CandidateCertified        bool
	NearOrientationClue       bool
	NearCouplingClue          bool
	NativeDerivationCertified bool
	RemainingSeal             string
	Verdict                   string
}

type Analysis struct {
	Runtime     RuntimeInheritance
	Definition  DeficitDefinition
	Candidates  CandidateSet
	Orientation OrientationAudit
	Transport   TransportAudit
	Corrections CorrectionScaleAudit
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
	g585, err := generation2koidewalloffsetsourcecandidateaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate585 predecessor: %w", err)
	}
	g583, err := generation2koidechamberwalloffsetaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate583 predecessor: %w", err)
	}
	g580, err := generation2koidetransportvectordecompositionaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate580 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history-transport runtime: %w", err)
	}
	runtime := inheritRuntime(g585, g583, g580, bundle)
	definition := defineDeficit(runtime)
	candidates := buildCandidateSet(definition.Kappa, runtime, bundle)
	orientation := auditOrientation(candidates)
	transport := auditTransport(candidates)
	corrections := auditCorrectionScales(candidates)
	decision := decideSource(candidates, orientation, transport, corrections)
	firewalls := auditFirewalls()
	final := compileFinal(definition, decision, orientation, corrections)
	truth := "Gate 586 factors the charged-lepton Koide wall offset as epsilon_e=(1/(8*pi))(1-kappa_e).  The deficit kappa_e is about 0.005503554, and the nearest typed runtime quantity is sqrt(J_CKM) at about 1.44% relative residual.  This is an orientation-sized clue, not a certified charged-lepton source, because CKM is a quark-sector orientation and no PMNS/intertwiner/native root-trace operator is present."
	return Analysis{Runtime: runtime, Definition: definition, Candidates: candidates, Orientation: orientation, Transport: transport, Corrections: corrections, Decision: decision, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g585 generation2koidewalloffsetsourcecandidateaudit.Analysis, g583 generation2koidechamberwalloffsetaudit.Analysis, g580 generation2koidetransportvectordecompositionaudit.Analysis, b historytransport.Bundle) RuntimeInheritance {
	g2 := b.EndVector.G2
	alpha2 := g2 * g2 / (4.0 * math.Pi)
	return RuntimeInheritance{
		EpsilonRad:             g585.Target.PrimaryEpsilonRad,
		EpsilonDeg:             g585.Target.PrimaryEpsilonDeg,
		LoopUnit:               g585.Loop.OneOver8Pi.Value,
		Gate585BestCandidate:   g585.Candidates.Best.Name,
		Gate585BestRelative:    g585.Candidates.Best.RelativeResidual,
		KappaFromGate585:       -g585.Loop.RequiredCorrection,
		JCKM:                   b.FlavorTransport.JCKM,
		SqrtJCKM:               math.Sqrt(b.FlavorTransport.JCKM),
		Alpha2MZ:               alpha2,
		Alpha2Over2Pi:          alpha2 / (2.0 * math.Pi),
		ProjectiveDriftRad:     g580.Transport.ProjectiveAngularDelta,
		DeltaPhiRad:            g580.Transport.DeltaPhiRad,
		DeltaThetaRad:          g580.Transport.DeltaThetaRad,
		DeltaEpsilonRad:        g583.Lambda12.EpsilonRad - g583.MZ.EpsilonRad,
		KoideAmplitudeResidual: math.Abs(g583.MZ.PlaneAmplitudeR - 1.0),
		Mu0GeV:                 b.EndVector.Mu0GeV,
		Lambda12GeV:            b.GaugeBoundary.Lambda12GeV,
		Verdict:                StatusGate585Inherited,
	}
}

func defineDeficit(r RuntimeInheritance) DeficitDefinition {
	kappa := 1.0 - r.EpsilonRad/r.LoopUnit
	return DeficitDefinition{
		Formula:             "epsilon_e = L(1-kappa_e), L=1/(8*pi), kappa_e=1-8*pi*epsilon_e",
		LoopUnit:            r.LoopUnit,
		EpsilonRad:          r.EpsilonRad,
		EpsilonDeg:          r.EpsilonDeg,
		Kappa:               kappa,
		KappaPercent:        100.0 * kappa,
		ReconstructedEpsRad: r.LoopUnit * (1.0 - kappa),
		ReconstructionError: r.LoopUnit*(1.0-kappa) - r.EpsilonRad,
		Interpretation:      "kappa_e is the deficit correcting the loop-sized angle 1/(8*pi) down to the observed electron-wall offset",
		NearTolerance:       nearRelativeTolerance,
		CertifiedTolerance:  certifiedRelativeTolerance,
		Verdict:             StatusLoopDeficitDefined,
	}
}

func buildCandidateSet(kappa float64, r RuntimeInheritance, b historytransport.Bundle) CandidateSet {
	gY, g2, gStar := b.EndVector.GY, b.EndVector.G2, b.GaugeBoundary.GStar
	e := gY * g2 / math.Sqrt(gY*gY+g2*g2)
	alphaEM := e * e / (4.0 * math.Pi)
	alpha2 := g2 * g2 / (4.0 * math.Pi)
	alphaStar := gStar * gStar / (4.0 * math.Pi)
	vals := []Candidate{
		mk(kappa, "sqrt(J_CKM)", "orientation_proxy", "sqrt(J_CKM)", math.Sqrt(b.FlavorTransport.JCKM), "square-root CKM oriented-area proxy; close to kappa_e but quark-sector unless an intertwiner is proven"),
		mk(kappa, "J_CKM", "orientation_proxy", "J_CKM", b.FlavorTransport.JCKM, "CKM oriented-area invariant itself"),
		mk(kappa, "alpha_2(M_Z)/(2π)", "coupling_correction", "alpha_2(M_Z)/(2*pi)", alpha2/(2.0*math.Pi), "weak-coupling correction coefficient scale"),
		mk(kappa, "alpha_2(M_Z)/π", "coupling_correction", "alpha_2(M_Z)/pi", alpha2/math.Pi, "larger weak-coupling correction coefficient"),
		mk(kappa, "alpha_EM(M_Z)", "coupling_correction", "e(M_Z)^2/(4*pi)", alphaEM, "electromagnetic coupling coefficient"),
		mk(kappa, "alpha_EM(M_Z)/π", "coupling_correction", "alpha_EM(M_Z)/pi", alphaEM/math.Pi, "electromagnetic loop-normalized correction"),
		mk(kappa, "alpha_star(Lambda_12)/(2π)", "boundary_correction", "alpha_star/(2*pi)", alphaStar/(2.0*math.Pi), "boundary weak-coupling correction coefficient"),
		mk(kappa, "alpha_star(Lambda_12)/π", "boundary_correction", "alpha_star/pi", alphaStar/math.Pi, "larger boundary weak-coupling coefficient"),
		mk(kappa, "R_3-1", "gauge_residual", "g3(Lambda_12)/g_star - 1", b.GaugeBoundary.R3-1.0, "strong mismatch as direct correction coefficient"),
		mk(kappa, "(R_3-1)/(2π)", "gauge_residual", "(R_3-1)/(2*pi)", (b.GaugeBoundary.R3-1.0)/(2.0*math.Pi), "loop-normalized strong mismatch"),
		mk(kappa, "(R_3-1)/(8π)", "gauge_residual", "(R_3-1)/(8*pi)", (b.GaugeBoundary.R3-1.0)/(8.0*math.Pi), "quarter loop-normalized strong mismatch"),
		mk(kappa, "|Delta_sin²|/(8π)", "weak_angle_residual", "abs(Delta_sin2)/(8*pi)", math.Abs(b.WeakAngleTransport.DeltaSin2)/(8.0*math.Pi), "loop-normalized weak-angle residual"),
		mk(kappa, "|lambda(Lambda_12)|/(2π)", "scalar_residual", "abs(lambda(Lambda_12))/(2*pi)", math.Abs(b.ScalarTransport.LambdaLambda12)/(2.0*math.Pi), "loop-normalized scalar boundary residual"),
		mk(kappa, "|lambda(Lambda_12)|/(8π)", "scalar_residual", "abs(lambda(Lambda_12))/(8*pi)", math.Abs(b.ScalarTransport.LambdaLambda12)/(8.0*math.Pi), "quarter loop-normalized scalar boundary residual"),
		mk(kappa, "Koide |R_e-1|", "charged_lepton_shape", "abs(R_e(M_Z)-1)", r.KoideAmplitudeResidual, "Koide amplitude defect from the Fourier circle"),
		mk(kappa, "Delta epsilon_e transport", "charged_lepton_transport", "epsilon_e(Lambda_12)-epsilon_e(M_Z)", math.Abs(r.DeltaEpsilonRad), "wall-offset drift under v1 transport"),
		mk(kappa, "Delta phi_e transport", "charged_lepton_transport", "abs(Delta phi_e)", math.Abs(r.DeltaPhiRad), "azimuth drift under v1 transport"),
		mk(kappa, "Delta theta_e transport", "charged_lepton_transport", "abs(Delta theta_e)", math.Abs(r.DeltaThetaRad), "cone-angle drift under v1 transport"),
		mk(kappa, "Projective angular drift", "charged_lepton_transport", "angular distance between projective rays", r.ProjectiveDriftRad, "total projective angular motion under v1 transport"),
	}
	sort.Slice(vals, func(i, j int) bool { return vals[i].AbsResidual < vals[j].AbsResidual })
	near := make([]Candidate, 0)
	cert := make([]Candidate, 0)
	for _, c := range vals {
		if c.Near {
			near = append(near, c)
		}
		if c.Certified {
			cert = append(cert, c)
		}
	}
	return CandidateSet{TargetKappa: kappa, CandidateCount: len(vals), Candidates: vals, Best: vals[0], NearCandidates: near, CertifiedCandidates: cert, Verdict: strings.Join([]string{StatusTypedCandidateSetDefined, StatusOrientationCompared, StatusTransportCompared, StatusCouplingCompared}, ";")}
}

func mk(kappa float64, name, class, eq string, value float64, interp string) Candidate {
	res := value - kappa
	rel := res / kappa
	return Candidate{Name: name, Class: class, Equation: eq, Value: value, SignedResidual: res, AbsResidual: math.Abs(res), RelativeResidual: rel, Near: math.Abs(rel) < nearRelativeTolerance, Certified: math.Abs(rel) < certifiedRelativeTolerance, Interpretation: interp}
}

func auditOrientation(c CandidateSet) OrientationAudit {
	j := findCandidate(c, "J_CKM")
	sj := findCandidate(c, "sqrt(J_CKM)")
	best := j
	if sj.AbsResidual < best.AbsResidual {
		best = sj
	}
	return OrientationAudit{JCKM: j, SqrtJCKM: sj, BestOrientation: best, NearButNotSource: sj.Near && !sj.Certified, PMNSRuntimeInput: false, Verdict: strings.Join([]string{StatusBestSqrtJCKM, StatusSqrtJCKMNearOnly, StatusCKMNotLeptonSource, StatusNoPMNSRuntime}, ";")}
}

func auditTransport(c CandidateSet) TransportAudit {
	cs := []Candidate{findCandidate(c, "Projective angular drift"), findCandidate(c, "Delta phi_e transport"), findCandidate(c, "Delta theta_e transport"), findCandidate(c, "Delta epsilon_e transport"), findCandidate(c, "Koide |R_e-1|")}
	best := cs[0]
	for _, x := range cs[1:] {
		if x.AbsResidual < best.AbsResidual {
			best = x
		}
	}
	return TransportAudit{ProjectiveDrift: cs[0], DeltaPhi: cs[1], DeltaTheta: cs[2], DeltaEpsilon: cs[3], KoideRMinusOne: cs[4], BestTransport: best, Certified: best.Certified, Verdict: StatusTransportNoFix}
}

func auditCorrectionScales(c CandidateSet) CorrectionScaleAudit {
	cs := []Candidate{findCandidate(c, "alpha_2(M_Z)/(2π)"), findCandidate(c, "alpha_EM(M_Z)"), findCandidate(c, "alpha_EM(M_Z)/π"), findCandidate(c, "(R_3-1)/(2π)"), findCandidate(c, "|Delta_sin²|/(8π)"), findCandidate(c, "|lambda(Lambda_12)|/(2π)")}
	best := cs[0]
	for _, x := range cs[1:] {
		if x.AbsResidual < best.AbsResidual {
			best = x
		}
	}
	verdict := StatusCouplingResidualNoFix
	if cs[0].Near && !cs[0].Certified {
		verdict = strings.Join([]string{StatusAlpha2Over2PiNearOnly, StatusCouplingResidualNoFix}, ";")
	}
	return CorrectionScaleAudit{Alpha2Over2Pi: cs[0], AlphaEM: cs[1], AlphaEMOverPi: cs[2], StrongMismatchOver2Pi: cs[3], DeltaSin2Over8Pi: cs[4], LambdaOver2Pi: cs[5], BestCorrection: best, Certified: best.Certified, Verdict: verdict}
}

func decideSource(c CandidateSet, o OrientationAudit, t TransportAudit, s CorrectionScaleAudit) SourceDecision {
	cert := len(c.CertifiedCandidates) > 0
	decision := "No typed runtime quantity certifies kappa_e as a source.  sqrt(J_CKM) is the nearest orientation-sized clue, and alpha_2/(2*pi) is also percent-close, but neither is lawful as a charged-lepton deficit source without a PMNS/lepton orientation input or an ASHA operator mapping a correction coefficient into the Koide wall offset."
	meaning := "orientation-sized correction clue, not a charged-lepton source theorem"
	verdict := strings.Join([]string{StatusBestSqrtJCKM, StatusSqrtJCKMNearOnly, StatusAlpha2Over2PiNearOnly, StatusNoCertifiedKappaSource, StatusCKMNotLeptonSource, StatusNoPMNSRuntime, StatusNoNativeDeficitOperator, StatusKappaRemainsSeal}, ";")
	return SourceDecision{BestCandidateName: c.Best.Name, BestCandidateValue: c.Best.Value, BestRelativeResidual: c.Best.RelativeResidual, BestAbsResidual: c.Best.AbsResidual, NearClue: c.Best.Near, CertifiedSource: cert, CandidateMeaning: meaning, MinimalNextRequirement: "a typed lepton-sector orientation quantity, PMNS-enabled comparison, or native root-trace/circulant loop-threshold operator that maps to kappa_e", Decision: decision, Verdict: verdict}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKappa: false, DerivesEpsilon: false, DerivesKoide: false, DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesCKM: false, DerivesPMNS: false, AddsNewCarrier: false, PromotesObservedAsNative: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoCertifiedKappaSource, StatusNoNativeDeficitOperator, StatusKappaRemainsSeal, StatusNoFlavorPromotion, StatusObservedEndpointPreserved, StatusGate352Preserved, StatusGate586BoundaryPreserved}, ";")}
}

func compileFinal(d DeficitDefinition, s SourceDecision, o OrientationAudit, c CorrectionScaleAudit) FinalVerdict {
	return FinalVerdict{SealName: "ChargedLeptonKoideLoopAngleDeficitSeal", EpsilonRad: d.EpsilonRad, LoopUnit: d.LoopUnit, Kappa: d.Kappa, BestCandidate: s.BestCandidateName, BestCandidateValue: s.BestCandidateValue, BestCandidateRelativeDiff: s.BestRelativeResidual, CandidateCertified: s.CertifiedSource, NearOrientationClue: o.SqrtJCKM.Near && !o.SqrtJCKM.Certified, NearCouplingClue: c.Alpha2Over2Pi.Near && !c.Alpha2Over2Pi.Certified, NativeDerivationCertified: false, RemainingSeal: "kappa_e remains the loop-angle deficit history seal until a lepton-sector or native operator fixes it", Verdict: strings.Join([]string{StatusBestSqrtJCKM, StatusSqrtJCKMNearOnly, StatusNoCertifiedKappaSource, StatusGate586BoundaryPreserved}, ";")}
}

func Statuses() []string {
	return []string{StatusGate585Inherited, StatusLoopDeficitDefined, StatusTypedCandidateSetDefined, StatusOrientationCompared, StatusTransportCompared, StatusCouplingCompared, StatusBestSqrtJCKM, StatusSqrtJCKMNearOnly, StatusAlpha2Over2PiNearOnly, StatusNoCertifiedKappaSource, StatusCKMNotLeptonSource, StatusNoPMNSRuntime, StatusTransportNoFix, StatusCouplingResidualNoFix, StatusNoNativeDeficitOperator, StatusKappaRemainsSeal, StatusNoFlavorPromotion, StatusObservedEndpointPreserved, StatusGate352Preserved, StatusGate586BoundaryPreserved}
}

func findCandidate(c CandidateSet, name string) Candidate {
	for _, x := range c.Candidates {
		if x.Name == name {
			return x
		}
	}
	return Candidate{Name: name, Value: math.NaN(), SignedResidual: math.NaN(), AbsResidual: math.Inf(1), RelativeResidual: math.NaN(), Interpretation: "candidate not found"}
}
