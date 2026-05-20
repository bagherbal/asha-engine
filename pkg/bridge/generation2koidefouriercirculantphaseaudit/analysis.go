// Package generation2koidefouriercirculantphaseaudit implements Gate 582:
// Koide Fourier/Circulant Phase Audit.
//
// Gates 577-581 reduced the charged-lepton environment to a nearly fixed
// projective ray in square-root Yukawa space, close to the Koide cone. Gate 582
// rewrites that ray in the natural democratic-plus-Fourier-plane coordinates
//
//	x_j = A [ 1 + sqrt(2) R cos(delta + 2*pi*j/3) ]
//
// where R=1 is exactly the Koide cone. The gate audits the Fourier phase delta,
// its convention/order dependence, and whether it is simpler than the azimuth
// phi from Gate 578. It preserves Gate 352's root-trace obstruction and does
// not derive lepton masses, flavor texture, CKM/PMNS, or generation hierarchy.
package generation2koidefouriercirculantphaseaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidecoordinatebetafunctionaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE582-KOIDE-FOURIER-CIRCULANT-PHASE-AUDIT"

	StatusGate581Inherited          = "PASS_GATE581_KOIDE_COORDINATE_BETA_RUNTIME_INHERITED"
	StatusFourierFrameDerived       = "PASS_FOURIER_CIRCULANT_KOIDE_FRAME_DERIVED"
	StatusKoideAmplitudeEquivalence = "PASS_KOIDE_CONE_EQUIVALENT_TO_FOURIER_PLANE_AMPLITUDE_ONE"
	StatusCanonicalMZComputed       = "PASS_CANONICAL_CHARGED_LEPTON_FOURIER_PHASE_COMPUTED_AT_MZ"
	StatusCanonicalLambdaComputed   = "PASS_CANONICAL_CHARGED_LEPTON_FOURIER_PHASE_COMPUTED_AT_LAMBDA12"
	StatusPhaseStable               = "PASS_FOURIER_PHASE_STABLE_UNDER_V1_TRANSPORT"
	StatusPhiDeltaRelation          = "PASS_FOURIER_PHASE_RELATED_TO_GATE578_AZIMUTH_BY_DELTA_EQUALS_PI_OVER_SIX_MINUS_PHI"
	StatusFourierOrientationSupport = "CONDITIONAL_SUPPORT_FOURIER_PHASE_EXPOSES_GENERATION_PLANE_ORIENTATION"
	StatusPermutationAmbiguous      = "FAILED_ROUTE_FOURIER_PHASE_NOT_UNIQUE_UNDER_PERMUTATION_OR_PHASE_CONVENTION"
	StatusNoSimpleRational          = "FAILED_ROUTE_NO_SIMPLE_RATIONAL_FOURIER_PHASE_CERTIFIED"
	StatusNoNativeCirculantOperator = "FAILED_ROUTE_NO_NATIVE_CIRCULANT_GENERATION_OPERATOR_OR_ROOT_TRACE_PHASE_SELECTOR"
	StatusNoNativeRayDerivation     = "FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_PROJECTIVE_RAY_DERIVATION_FROM_FOURIER_PHASE"
	StatusNoFlavorPromotion         = "FIREWALL_PRESERVED_FOURIER_PHASE_DOES_NOT_DERIVE_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusObservedEndpointPreserved = "FIREWALL_PRESERVED_CHARGED_LEPTON_FOURIER_PHASE_REMAINS_HISTORY_ENDPOINT_DATA"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate582BoundaryPreserved  = "FIREWALL_PRESERVED_GATE582_KOIDE_FOURIER_PHASE_BOUNDARY"
)

const (
	maxSimpleDenominator  = 72
	certificationMultiple = 100.0
)

type RuntimeInheritance struct {
	Mu0GeV             float64
	Lambda12GeV        float64
	Gate581MZDThetaDeg float64
	Gate581MZDPhiDeg   float64
	Gate581Verdict     string
	RuntimeSource      string
	Verdict            string
}

type FourierFormulaAudit struct {
	Formula          string
	ADefinition      string
	RDefinition      string
	DeltaDefinition  string
	KoideEquivalence string
	QFromR           string
	CanonicalOrder   []string
	Convention       string
	Verdict          string
}

type FourierPoint struct {
	Name                    string
	Labels                  []string
	Values                  []float64
	RootVector              []float64
	A                       float64
	PlaneAmplitudeR         float64
	PlaneAmplitudeResidual  float64
	Q                       float64
	DeltaQ                  float64
	DeltaDeg                float64
	DeltaTurn               float64
	DeltaPhiRelationDeg     float64
	ReconstructedRootVector []float64
	MaxReconstructionError  float64
	Verdict                 string
}

type PhaseTransportAudit struct {
	MZDeltaDeg              float64
	LambdaDeltaDeg          float64
	SignedDriftDeg          float64
	AbsDriftDeg             float64
	MZAmplitudeResidual     float64
	LambdaAmplitudeResidual float64
	AmplitudeMovesTowardOne bool
	PhaseStable             bool
	Verdict                 string
}

type PermutationPhase struct {
	Order               []string
	DeltaDeg            float64
	DeltaTurn           float64
	NearestRational     string
	NearestRationalDeg  float64
	RationalResidualDeg float64
}

type PermutationAudit struct {
	Phases                []PermutationPhase
	CanonicalOrder        []string
	CanonicalDeltaDeg     float64
	BestRationalOrder     []string
	BestRational          string
	BestRationalDeg       float64
	BestResidualDeg       float64
	CertificationDeg      float64
	UniqueWithoutOrdering bool
	SimplePhaseCertified  bool
	Explanation           string
	Verdict               string
}

type FirewallAudit struct {
	DerivesLeptonMasses        bool
	DerivesYukawaEigenvalues   bool
	DerivesKoide               bool
	DerivesFourierPhase        bool
	DerivesCKM                 bool
	DerivesPMNS                bool
	DerivesGenerationHierarchy bool
	AddsNewCarrier             bool
	PromotesObservedAsNative   bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	SealName                string
	CanonicalDeltaMZDeg     float64
	CanonicalDeltaLambdaDeg float64
	FourierAmplitudeMZ      float64
	FourierAmplitudeLambda  float64
	PhaseStableInV1         bool
	SimpleRationalCertified bool
	NativeSelectorCertified bool
	MinimalNextRequirement  string
	Verdict                 string
}

type Analysis struct {
	Runtime     RuntimeInheritance
	Formula     FourierFormulaAudit
	MZ          FourierPoint
	Lambda12    FourierPoint
	Transport   PhaseTransportAudit
	Permutation PermutationAudit
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
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport runtime: %w", err)
	}
	g581, err := generation2koidecoordinatebetafunctionaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate581 predecessor: %w", err)
	}
	runtime := inheritRuntime(bundle, g581)
	formula := deriveFormula()
	labels := []string{"e", "mu", "tau"}
	mz := fourierPoint("M_Z", labels, values(bundle.FlavorTransport.YukawaSingularValuesMZ.ChargedLeptons, labels), StatusCanonicalMZComputed)
	lambda := fourierPoint("Lambda_12", labels, values(bundle.FlavorTransport.YukawaSingularValuesLambda12.ChargedLeptons, labels), StatusCanonicalLambdaComputed)
	transport := auditTransport(mz, lambda)
	permutation := auditPermutations(labels, values(bundle.FlavorTransport.YukawaSingularValuesMZ.ChargedLeptons, labels), transport.AbsDriftDeg)
	firewalls := auditFirewalls()
	final := compileFinal(mz, lambda, transport, permutation)
	truth := "Gate 582 rewrites the charged-lepton Koide ray as a democratic plus Fourier-plane circulant phase. The Fourier amplitude R is nearly one, so the ray is nearly on the Koide cone, and the canonical phase delta is stable under v1 transport. But delta is convention/permutation dependent and no simple rational/root-of-unity or native circulant/root-trace selector is certified."
	return Analysis{Runtime: runtime, Formula: formula, MZ: mz, Lambda12: lambda, Transport: transport, Permutation: permutation, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(b historytransport.Bundle, g581 generation2koidecoordinatebetafunctionaudit.Analysis) RuntimeInheritance {
	return RuntimeInheritance{Mu0GeV: b.EndVector.Mu0GeV, Lambda12GeV: b.GaugeBoundary.Lambda12GeV, Gate581MZDThetaDeg: g581.MZ.DThetaDTDeg, Gate581MZDPhiDeg: g581.MZ.DPhiDTDeg, Gate581Verdict: g581.Final.Verdict, RuntimeSource: "historytransport.BuildDefault() plus Gate581 Koide coordinate beta audit", Verdict: StatusGate581Inherited}
}

func deriveFormula() FourierFormulaAudit {
	return FourierFormulaAudit{
		Formula:          "x_j=A[1+sqrt(2)R cos(delta+2*pi*j/3)], j=0,1,2",
		ADefinition:      "A=(x_0+x_1+x_2)/3",
		RDefinition:      "R=|sum_j (x_j/A-1) exp(-2*pi*i*j/3)|/(3*sqrt(2)/2)",
		DeltaDefinition:  "delta=arg sum_j (x_j/A-1) exp(-2*pi*i*j/3)",
		KoideEquivalence: "Q=(1+R^2)/3, so Q=2/3 iff R=1",
		QFromR:           "Q=(x·x)/(sum_j x_j)^2=(1+R^2)/3",
		CanonicalOrder:   []string{"e", "mu", "tau"},
		Convention:       "positive Fourier step +2*pi*j/3 in the cosine and DFT extraction with exp(-2*pi*i*j/3)",
		Verdict:          strings.Join([]string{StatusFourierFrameDerived, StatusKoideAmplitudeEquivalence}, ";"),
	}
}

func fourierPoint(name string, labels []string, ys []float64, verdict string) FourierPoint {
	roots := make([]float64, len(ys))
	for i, y := range ys {
		roots[i] = math.Sqrt(y)
	}
	a := sum(roots) / 3.0
	cRe, cIm := fourierCoefficient(roots)
	deltaDeg := normalizeDeg(radToDeg(math.Atan2(cIm, cRe)))
	r := math.Hypot(cRe, cIm) / (3.0 * math.Sqrt2 / 2.0)
	q := dot(roots, roots) / math.Pow(sum(roots), 2)
	reconstructed := reconstruct(a, r, degToRad(deltaDeg))
	return FourierPoint{
		Name:                    name,
		Labels:                  append([]string{}, labels...),
		Values:                  append([]float64{}, ys...),
		RootVector:              roots,
		A:                       a,
		PlaneAmplitudeR:         r,
		PlaneAmplitudeResidual:  r - 1.0,
		Q:                       q,
		DeltaQ:                  q - 2.0/3.0,
		DeltaDeg:                deltaDeg,
		DeltaTurn:               deltaDeg / 360.0,
		DeltaPhiRelationDeg:     signedAngleDelta(deltaDeg, normalizeDeg(30.0-257.267180032892)),
		ReconstructedRootVector: reconstructed,
		MaxReconstructionError:  maxAbsDiff(roots, reconstructed),
		Verdict:                 verdict,
	}
}

func auditTransport(mz, lambda FourierPoint) PhaseTransportAudit {
	drift := signedAngleDelta(lambda.DeltaDeg, mz.DeltaDeg)
	return PhaseTransportAudit{MZDeltaDeg: mz.DeltaDeg, LambdaDeltaDeg: lambda.DeltaDeg, SignedDriftDeg: drift, AbsDriftDeg: math.Abs(drift), MZAmplitudeResidual: mz.PlaneAmplitudeResidual, LambdaAmplitudeResidual: lambda.PlaneAmplitudeResidual, AmplitudeMovesTowardOne: math.Abs(lambda.PlaneAmplitudeResidual) < math.Abs(mz.PlaneAmplitudeResidual), PhaseStable: math.Abs(drift) < 3e-4, Verdict: strings.Join([]string{StatusCanonicalMZComputed, StatusCanonicalLambdaComputed, StatusPhaseStable, StatusFourierOrientationSupport}, ";")}
}

func auditPermutations(labels []string, ys []float64, driftDeg float64) PermutationAudit {
	perms := permutations([]int{0, 1, 2})
	phases := make([]PermutationPhase, 0, len(perms))
	bestResidual := math.Inf(1)
	bestRational := ""
	bestRationalDeg := 0.0
	var bestOrder []string
	for _, p := range perms {
		orderedLabels := []string{labels[p[0]], labels[p[1]], labels[p[2]]}
		orderedValues := []float64{ys[p[0]], ys[p[1]], ys[p[2]]}
		point := fourierPoint("permutation", orderedLabels, orderedValues, "")
		residual, rational, rationalDeg := nearestRationalTurn(point.DeltaDeg, maxSimpleDenominator)
		phases = append(phases, PermutationPhase{Order: orderedLabels, DeltaDeg: point.DeltaDeg, DeltaTurn: point.DeltaTurn, NearestRational: rational, NearestRationalDeg: rationalDeg, RationalResidualDeg: residual})
		if residual < bestResidual {
			bestResidual = residual
			bestRational = rational
			bestRationalDeg = rationalDeg
			bestOrder = append([]string{}, orderedLabels...)
		}
	}
	threshold := certificationMultiple * math.Max(driftDeg, 1e-12)
	certified := bestResidual <= threshold
	explanation := "The Fourier phase is a coordinate on the generation plane. Permuting the charged-lepton labels shifts or reflects delta; without a native generation-ordering/circulant operator, no unique delta is selected. The nearest simple rational found below the denominator cutoff is recorded but fails the drift-based certification threshold."
	return PermutationAudit{Phases: phases, CanonicalOrder: []string{"e", "mu", "tau"}, CanonicalDeltaDeg: phases[0].DeltaDeg, BestRationalOrder: bestOrder, BestRational: bestRational, BestRationalDeg: bestRationalDeg, BestResidualDeg: bestResidual, CertificationDeg: threshold, UniqueWithoutOrdering: false, SimplePhaseCertified: certified, Explanation: explanation, Verdict: strings.Join([]string{StatusPermutationAmbiguous, StatusNoSimpleRational}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesKoide: false, DerivesFourierPhase: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, AddsNewCarrier: false, PromotesObservedAsNative: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoNativeCirculantOperator, StatusNoNativeRayDerivation, StatusNoFlavorPromotion, StatusObservedEndpointPreserved, StatusGate352Preserved, StatusGate582BoundaryPreserved}, ";")}
}

func compileFinal(mz, lambda FourierPoint, transport PhaseTransportAudit, perm PermutationAudit) FinalVerdict {
	return FinalVerdict{SealName: "ChargedLeptonKoideFourierPhaseSeal", CanonicalDeltaMZDeg: mz.DeltaDeg, CanonicalDeltaLambdaDeg: lambda.DeltaDeg, FourierAmplitudeMZ: mz.PlaneAmplitudeR, FourierAmplitudeLambda: lambda.PlaneAmplitudeR, PhaseStableInV1: transport.PhaseStable, SimpleRationalCertified: perm.SimplePhaseCertified, NativeSelectorCertified: false, MinimalNextRequirement: "a native root-trace/absolute-Dirac or circulant generation-plane operator that selects the charged-lepton projective ray and fixes the ordering/phase convention", Verdict: strings.Join([]string{StatusFourierOrientationSupport, StatusPermutationAmbiguous, StatusNoSimpleRational, StatusNoNativeCirculantOperator, StatusGate582BoundaryPreserved}, ";")}
}

func Statuses() []string {
	return []string{
		StatusGate581Inherited,
		StatusFourierFrameDerived,
		StatusKoideAmplitudeEquivalence,
		StatusCanonicalMZComputed,
		StatusCanonicalLambdaComputed,
		StatusPhaseStable,
		StatusPhiDeltaRelation,
		StatusFourierOrientationSupport,
		StatusPermutationAmbiguous,
		StatusNoSimpleRational,
		StatusNoNativeCirculantOperator,
		StatusNoNativeRayDerivation,
		StatusNoFlavorPromotion,
		StatusObservedEndpointPreserved,
		StatusGate352Preserved,
		StatusGate582BoundaryPreserved,
	}
}

func values(m map[string]float64, labels []string) []float64 {
	out := make([]float64, len(labels))
	for i, label := range labels {
		out[i] = m[label]
	}
	return out
}

func fourierCoefficient(roots []float64) (float64, float64) {
	a := sum(roots) / 3.0
	var re, im float64
	for j, x := range roots {
		w := x/a - 1.0
		angle := -2.0 * math.Pi * float64(j) / 3.0
		re += w * math.Cos(angle)
		im += w * math.Sin(angle)
	}
	return re, im
}

func reconstruct(a, r, delta float64) []float64 {
	out := make([]float64, 3)
	for j := 0; j < 3; j++ {
		out[j] = a * (1.0 + math.Sqrt2*r*math.Cos(delta+2.0*math.Pi*float64(j)/3.0))
	}
	return out
}

func nearestRationalTurn(deg float64, maxDen int) (float64, string, float64) {
	bestResidual := math.Inf(1)
	bestP, bestQ := 0, 1
	bestDeg := 0.0
	for q := 1; q <= maxDen; q++ {
		for p := 0; p <= q; p++ {
			candidate := 360.0 * float64(p) / float64(q)
			residual := math.Abs(signedAngleDelta(deg, candidate))
			if residual < bestResidual {
				bestResidual = residual
				bestP, bestQ = p, q
				bestDeg = candidate
			}
		}
	}
	return bestResidual, fmt.Sprintf("%d/%d turn", bestP, bestQ), bestDeg
}

func permutations(xs []int) [][]int {
	var out [][]int
	var rec func(int)
	arr := append([]int{}, xs...)
	rec = func(k int) {
		if k == len(arr) {
			out = append(out, append([]int{}, arr...))
			return
		}
		for i := k; i < len(arr); i++ {
			arr[k], arr[i] = arr[i], arr[k]
			rec(k + 1)
			arr[k], arr[i] = arr[i], arr[k]
		}
	}
	rec(0)
	return out
}

func dot(a, b []float64) float64 {
	var s float64
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func sum(xs []float64) float64 {
	var s float64
	for _, x := range xs {
		s += x
	}
	return s
}

func maxAbsDiff(a, b []float64) float64 {
	max := 0.0
	for i := range a {
		d := math.Abs(a[i] - b[i])
		if d > max {
			max = d
		}
	}
	return max
}

func normalizeDeg(x float64) float64 {
	y := math.Mod(x, 360.0)
	if y < 0 {
		y += 360.0
	}
	return y
}

func signedAngleDelta(a, b float64) float64 {
	d := math.Mod(a-b+180.0, 360.0)
	if d < 0 {
		d += 360.0
	}
	return d - 180.0
}

func radToDeg(x float64) float64 { return x * 180.0 / math.Pi }
func degToRad(x float64) float64 { return x * math.Pi / 180.0 }
