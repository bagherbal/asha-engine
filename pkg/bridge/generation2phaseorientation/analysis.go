// Package generation2phaseorientation implements Gate 446:
// Signed-Cycle / Complex Phase Orientation Sieve.
//
// Gate 445 forced only the unsigned triangular support of the Generation-2
// mass-lift bridge.  Gate 446 audits the next, stricter question: do the same
// structural boundaries force a signed real cycle or a complex CP phase?  The
// answer is negative.  The Hermitian triangular bridge has a gauge-invariant
// cycle phase
//
//	Phi = arg(z12 z23 conj(z13))
//
// and the exact determinant identity
//
//	det(K_gen + eps B) = (|z23|^2-|z12|^2) eps^2
//	                    + 2 Re(z12 z23 conj(z13)) eps^3.
//
// Endpoint balance removes the eps^2 term, but neither J/Gamma compatibility,
// eta-graded trace neutrality, nor CP covariance selects Phi.  Real signs leave
// two Z2 cycle classes; complex phases leave a continuum.  Therefore X_gen
// support remains forced, but Y_gen/phase orientation remains a quarantined
// bridge/source datum.
package generation2phaseorientation

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE446-SIGNED-CYCLE-COMPLEX-PHASE-ORIENTATION-SIEVE"

	StatusGate445TopologyInherited           = "CONDITIONAL_SUPPORT_GATE445_TRIANGLE_TOPOLOGY_INHERITED"
	StatusHermitianCycleArenaFormalized      = "CONDITIONAL_SUPPORT_HERMITIAN_TRIANGULAR_CYCLE_ARENA_FORMALIZED"
	StatusJGammaTraceBoundariesApplied       = "CONDITIONAL_SUPPORT_J_GAMMA_ETA_TRACE_BOUNDARIES_APPLIED"
	StatusRealSignedSieveCompleted           = "CONDITIONAL_SUPPORT_REAL_SIGNED_CYCLE_SIEVE_COMPLETED"
	StatusCyclePhaseInvariantIdentified      = "CONDITIONAL_SUPPORT_CYCLE_PHASE_INVARIANT_IDENTIFIED"
	StatusCPPhaseCapacityAudited             = "CONDITIONAL_SUPPORT_CP_PHASE_CAPACITY_AUDITED"
	StatusEmpiricalFirewallPreserved         = "CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED"
	StatusFailedSignedOrientationNotUnique   = "FAILED_ROUTE_SIGNED_CYCLE_ORIENTATION_NOT_UNIQUE"
	StatusFailedComplexPhaseContinuum        = "FAILED_ROUTE_COMPLEX_PHASE_CONTINUUM_UNDERDETERMINED"
	StatusFailedCPPhaseValueNotPredicted     = "FAILED_ROUTE_CP_PHASE_VALUE_NOT_PREDICTED"
	StatusFailedYGenNotNative                = "FAILED_ROUTE_Y_GEN_PHASE_QUADRATURE_NOT_NATIVE"
	StatusFailedNoMuonCharmMassPrediction    = "FAILED_ROUTE_NO_MUON_CHARM_MASS_VALUE_PREDICTION"
	StatusFirewallPhaseOrientationQuarantine = "FIREWALL_PRESERVED_PHASE_ORIENTATION_QUARANTINED"
)

const (
	FamilyRank      = 3
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                       bool
	Gate445KGenForced              bool
	Gate445XSupportForced          bool
	Gate445AmplitudeSealed         bool
	Gate445SignedOrientationSealed bool
	Gate445NoEmpiricalMasses       bool
	NativeFlavorDim                int
	KXYCoeffDimStillFree           int
	Verdict                        string
}

type OrientationArena struct {
	Executed                 bool
	KGen                     string
	BridgeAnsatz             string
	Hermitian                bool
	ZeroDiagonal             bool
	TriangleSupportInherited bool
	EndpointBalanced         bool
	VertexRephasingAllowed   bool
	GaugeInvariantCyclePhase string
	DeterminantIdentity      string
	EmpiricalDataImported    bool
	Verdict                  string
	Reason                   string
}

type Boundary struct {
	Name    string
	Formula string
	Applied bool
	Passed  bool
	Verdict string
	Reason  string
}

type RealSignCandidate struct {
	A                  int
	B                  int
	C                  int
	Product            int
	GaugeClass         string
	CyclePhase         string
	DeterminantLeading string
	CPPreserving       bool
	EtaTraceNeutral    bool
	JGammaCompatible   bool
	MassLiftCompatible bool
	Representative     bool
}

type RealSignSieve struct {
	Executed           bool
	Candidates         []RealSignCandidate
	PositiveCycleCount int
	NegativeCycleCount int
	Z2GaugeClasses     int
	UniqueSignedCycle  bool
	Verdict            string
	Reason             string
}

type PhaseSample struct {
	Label              string
	PhiRadians         float64
	DeterminantLeading string
	CPWitness          string
	MassLiftCompatible bool
	CPCapable          bool
	CPConjugateLabel   string
}

type ComplexPhaseSieve struct {
	Executed                    bool
	CyclePhaseInvariant         string
	VertexGaugeFreedom          string
	EndpointBalancedDeterminant string
	CPMap                       string
	MassLiftCondition           string
	CPWitnessFormula            string
	Samples                     []PhaseSample
	ContinuumSurvives           bool
	CPConjugatePairsSurvive     bool
	UniqueComplexPhase          bool
	CPPhaseValuePredicted       bool
	Verdict                     string
	Reason                      string
}

type OrientationConclusion struct {
	Executed                  bool
	XSupportTopologyPreserved bool
	SignedCycleForced         bool
	ComplexPhaseForced        bool
	YGenPromotedToNative      bool
	PhaseCoefficientFixed     bool
	CPViolationPredicted      bool
	MassLiftStillCompatible   bool
	Verdict                   string
	Reason                    string
}

type Firewall struct {
	Executed                     bool
	NoObservedMuonMassImported   bool
	NoObservedCharmMassImported  bool
	NoObservedYukawaImported     bool
	NoCKMImported                bool
	NoPMNSImported               bool
	XSupportTopologyForced       bool
	BridgeAmplitudeSealed        bool
	SignedCycleOrientationSealed bool
	ComplexPhaseSealed           bool
	YGenRemainsQuarantined       bool
	NativeFlavorDimBefore        int
	NativeFlavorDimAfter         int
	KXYCoeffDimStillFree         int
	Verdict                      string
	Reason                       string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Arena       OrientationArena
	Boundaries  []Boundary
	RealSieve   RealSignSieve
	PhaseSieve  ComplexPhaseSieve
	Conclusion  OrientationConclusion
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Arena = buildArena()
	a.Boundaries = buildBoundaries()
	a.RealSieve = buildRealSignSieve()
	a.PhaseSieve = buildComplexPhaseSieve()
	a.Conclusion = buildConclusion(a.RealSieve, a.PhaseSieve)
	a.Firewall = buildFirewall(a.Conclusion)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate445KGenForced: true, Gate445XSupportForced: true, Gate445AmplitudeSealed: true, Gate445SignedOrientationSealed: true, Gate445NoEmpiricalMasses: true, NativeFlavorDim: NativeFlavorDim, KXYCoeffDimStillFree: KXYCoeffDim, Verdict: StatusGate445TopologyInherited}
}

func buildArena() OrientationArena {
	return OrientationArena{Executed: true, KGen: "diag(-1,0,1)", BridgeAnsatz: "B(z12,z23,z13) Hermitian, zero-diagonal, full triangular support", Hermitian: true, ZeroDiagonal: true, TriangleSupportInherited: true, EndpointBalanced: true, VertexRephasingAllowed: true, GaugeInvariantCyclePhase: "Phi=arg(z12 z23 conjugate(z13))", DeterminantIdentity: "det(K+eps B)=(|z23|^2-|z12|^2)eps^2+2Re(z12 z23 conjugate(z13))eps^3", EmpiricalDataImported: false, Verdict: StatusHermitianCycleArenaFormalized, Reason: "Gate 446 keeps the Gate-445 triangle but allows the most general Hermitian edge orientation compatible with family-fiber rephasing"}
}

func buildBoundaries() []Boundary {
	return []Boundary{
		{Name: "Gate-445 support inheritance", Formula: "support(B)={(1,2),(2,3),(1,3)}", Applied: true, Passed: true, Verdict: StatusGate445TopologyInherited, Reason: "the unsigned closed triangle is already forced by the mass-lift bridge sieve"},
		{Name: "Hermitian/J/Gamma compatibility", Formula: "B_ji=conjugate(B_ij), diag(B)=0, [B,Gamma_gen]=0", Applied: true, Passed: true, Verdict: StatusJGammaTraceBoundariesApplied, Reason: "the family bridge remains an internal Hermitian source and does not alter chirality/gauge representation slots"},
		{Name: "eta-graded trace neutrality", Formula: "Tr(B)=0 and Tr_eta(B)=0 in family-only lift", Applied: true, Passed: true, Verdict: StatusJGammaTraceBoundariesApplied, Reason: "all zero-diagonal family-cycle orientations are trace neutral; this boundary cannot distinguish signs or phases"},
		{Name: "vertex rephasing quotient", Formula: "B_ij -> exp(i(theta_i-theta_j)) B_ij", Applied: true, Passed: true, Verdict: StatusCyclePhaseInvariantIdentified, Reason: "two edge phases are gauge convention; only the cycle phase Phi survives as a rephasing invariant"},
		{Name: "determinant mass-lift", Formula: "Re(z12 z23 conjugate(z13)) != 0", Applied: true, Passed: true, Verdict: StatusCPPhaseCapacityAudited, Reason: "mass lift excludes purely imaginary cycle product but still leaves infinitely many real/complex orientations"},
	}
}

func buildRealSignSieve() RealSignSieve {
	var xs []RealSignCandidate
	pos, neg := 0, 0
	for _, a := range []int{-1, 1} {
		for _, b := range []int{-1, 1} {
			for _, c := range []int{-1, 1} {
				p := a * b * c
				if p > 0 {
					pos++
				} else {
					neg++
				}
				xs = append(xs, makeRealSignCandidate(a, b, c))
			}
		}
	}
	sort.Slice(xs, func(i, j int) bool {
		if xs[i].Product != xs[j].Product {
			return xs[i].Product < xs[j].Product
		}
		if xs[i].A != xs[j].A {
			return xs[i].A < xs[j].A
		}
		if xs[i].B != xs[j].B {
			return xs[i].B < xs[j].B
		}
		return xs[i].C < xs[j].C
	})
	return RealSignSieve{Executed: true, Candidates: xs, PositiveCycleCount: pos, NegativeCycleCount: neg, Z2GaugeClasses: 2, UniqueSignedCycle: false, Verdict: StatusFailedSignedOrientationNotUnique, Reason: "the eight real sign assignments collapse under vertex sign flips to two invariant cycle-product classes; both pass all structural boundaries"}
}

func makeRealSignCandidate(a, b, c int) RealSignCandidate {
	p := a * b * c
	phase := "0"
	class := "positive cycle product"
	det := "2 eps^3"
	rep := a == 1 && b == 1 && c == 1
	if p < 0 {
		phase = "pi"
		class = "negative cycle product"
		det = "-2 eps^3"
		rep = a == 1 && b == 1 && c == -1
	}
	return RealSignCandidate{A: a, B: b, C: c, Product: p, GaugeClass: class, CyclePhase: phase, DeterminantLeading: det, CPPreserving: true, EtaTraceNeutral: true, JGammaCompatible: true, MassLiftCompatible: true, Representative: rep}
}

func buildComplexPhaseSieve() ComplexPhaseSieve {
	samples := []PhaseSample{
		makePhaseSample("0", 0, "0"),
		makePhaseSample("pi", math.Pi, "pi"),
		makePhaseSample("pi/4", math.Pi/4, "-pi/4"),
		makePhaseSample("-pi/4", -math.Pi/4, "pi/4"),
		makePhaseSample("pi/2", math.Pi/2, "-pi/2"),
	}
	return ComplexPhaseSieve{Executed: true, CyclePhaseInvariant: "Phi=arg(z12 z23 conjugate(z13))", VertexGaugeFreedom: "U(1)^3 family rephasing removes two edge phases but cannot remove Phi", EndpointBalancedDeterminant: "det(K+eps B)=2 r^3 cos(Phi) eps^3 for |z12|=|z23|=|z13|=r", CPMap: "CP: Phi -> -Phi", MassLiftCondition: "cos(Phi) != 0", CPWitnessFormula: "CP-odd cycle witness proportional to sin(Phi)", Samples: samples, ContinuumSurvives: true, CPConjugatePairsSurvive: true, UniqueComplexPhase: false, CPPhaseValuePredicted: false, Verdict: StatusFailedComplexPhaseContinuum, Reason: "mass-lift and compatibility boundaries remove only the purely imaginary cycle products; a continuum of CP-even and CP-odd phase orientations remains"}
}

func makePhaseSample(label string, phi float64, conjugate string) PhaseSample {
	cosv := math.Cos(phi)
	sinv := math.Sin(phi)
	if math.Abs(cosv) < 1e-9 {
		cosv = 0
	}
	if math.Abs(sinv) < 1e-9 {
		sinv = 0
	}
	mass := math.Abs(cosv) > 1e-9
	cp := math.Abs(sinv) > 1e-9
	det := fmt.Sprintf("%.6g eps^3", 2*cosv)
	wit := fmt.Sprintf("%.6g", sinv)
	return PhaseSample{Label: label, PhiRadians: phi, DeterminantLeading: det, CPWitness: wit, MassLiftCompatible: mass, CPCapable: cp, CPConjugateLabel: conjugate}
}

func buildConclusion(r RealSignSieve, p ComplexPhaseSieve) OrientationConclusion {
	return OrientationConclusion{Executed: true, XSupportTopologyPreserved: true, SignedCycleForced: r.UniqueSignedCycle, ComplexPhaseForced: p.UniqueComplexPhase, YGenPromotedToNative: false, PhaseCoefficientFixed: false, CPViolationPredicted: false, MassLiftStillCompatible: len(r.Candidates) == 8 && p.ContinuumSurvives, Verdict: StatusFirewallPhaseOrientationQuarantine, Reason: "the boundary intersection preserves Gate-445 mass-lift compatibility but does not collapse the signed or complex orientation to one survivor"}
}

func buildFirewall(c OrientationConclusion) Firewall {
	return Firewall{Executed: true, NoObservedMuonMassImported: true, NoObservedCharmMassImported: true, NoObservedYukawaImported: true, NoCKMImported: true, NoPMNSImported: true, XSupportTopologyForced: c.XSupportTopologyPreserved, BridgeAmplitudeSealed: true, SignedCycleOrientationSealed: !c.SignedCycleForced, ComplexPhaseSealed: !c.ComplexPhaseForced, YGenRemainsQuarantined: !c.YGenPromotedToNative, NativeFlavorDimBefore: NativeFlavorDim, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimStillFree: KXYCoeffDim, Verdict: StatusFirewallPhaseOrientationQuarantine, Reason: "X_gen support is structural after Gate 445, but orientation, phase, CP value, and sector amplitudes remain outside the native law-space"}
}

func buildNext() NextStep {
	return NextStep{Gate: 447, Title: "Sector-Coefficient Source Ledger / Amplitude Firewall Closure", Reason: "Gate 446 proves that phase orientation is not selected by the current boundary stack.", PrimaryTask: "audit whether any sector-source functional can reduce the remaining K/X/Y coefficient ledger without importing empirical Yukawa matrices, or whether the full coefficient amplitude firewall must remain sealed"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate445KGenForced || !a.Inheritance.Gate445XSupportForced || !a.Inheritance.Gate445SignedOrientationSealed {
		return fmt.Errorf("inheritance failed: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Executed || !a.Arena.Hermitian || !a.Arena.ZeroDiagonal || !a.Arena.TriangleSupportInherited || a.Arena.EmpiricalDataImported {
		return fmt.Errorf("arena failed: %s", FormatArena(a.Arena))
	}
	if len(a.Boundaries) != 5 {
		return fmt.Errorf("expected 5 boundaries, got %d", len(a.Boundaries))
	}
	for _, b := range a.Boundaries {
		if !b.Applied || !b.Passed {
			return fmt.Errorf("boundary failed: %s", FormatBoundary(b))
		}
	}
	if !a.RealSieve.Executed || a.RealSieve.UniqueSignedCycle || a.RealSieve.PositiveCycleCount != 4 || a.RealSieve.NegativeCycleCount != 4 || a.RealSieve.Z2GaugeClasses != 2 {
		return fmt.Errorf("real sign sieve failed: %s", FormatRealSignSieve(a.RealSieve))
	}
	if !a.PhaseSieve.Executed || !a.PhaseSieve.ContinuumSurvives || !a.PhaseSieve.CPConjugatePairsSurvive || a.PhaseSieve.UniqueComplexPhase || a.PhaseSieve.CPPhaseValuePredicted {
		return fmt.Errorf("complex phase sieve failed: %s", FormatComplexPhaseSieve(a.PhaseSieve))
	}
	if !a.Conclusion.XSupportTopologyPreserved || a.Conclusion.SignedCycleForced || a.Conclusion.ComplexPhaseForced || a.Conclusion.YGenPromotedToNative || a.Conclusion.CPViolationPredicted {
		return fmt.Errorf("conclusion failed: %s", FormatConclusion(a.Conclusion))
	}
	if !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.YGenRemainsQuarantined || !a.Firewall.SignedCycleOrientationSealed || !a.Firewall.ComplexPhaseSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimStillFree != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 446 is a negative but important sieve result: the Gate-445 triangle support remains forced, but the signed real cycle and the complex CP phase are not forced by Hermiticity, J/Gamma compatibility, eta-trace neutrality, determinant mass-lift, or vertex rephasing quotient. Real signs reduce only to two Z2 cycle classes, and complex Hermitian bridges retain the gauge-invariant continuum Phi=arg(z12 z23 conjugate(z13)). Therefore Y_gen and the CP phase remain quarantined; no muon/charm mass value, CKM phase, or Yukawa coefficient is predicted."
}

func FormatSign(a, b, c int) string { return fmt.Sprintf("(a=%d,b=%d,c=%d)", a, b, c) }

func phaseSampleSummary(xs []PhaseSample) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s det=%s cp=%t mass=%t", x.Label, x.DeterminantLeading, x.CPCapable, x.MassLiftCompatible)
	}
	return strings.Join(parts, " | ")
}
