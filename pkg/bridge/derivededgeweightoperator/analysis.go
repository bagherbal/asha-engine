// Package derivededgeweightoperator implements Gate 401:
// Derived Edge-Weight Operator / Hypercharge Laplacian Sieve.
//
// Gate 400 showed that the uniform one-form edge Laplacian is central on the
// four-real scalar carrier H_phi and that the native mixed compression only
// recovers the already-known 2+2 scalar response. Gate 401 tests the next
// stricter possibility: use already-derived electroweak/B-L data to assign
// nonuniform weights to the ten J-doubled finite-Dirac one-form edges and then
// ask whether the induced four-real scalar operator has the irreducible contact
// quartic q4 as an invariant polynomial.
//
// The theorem boundary is intentionally conservative. Native charge weights can
// differentiate the finite Dirac edges, but a q4 identification additionally
// requires a canonical compression to H_phi and a direct q4 polynomial match.
package derivededgeweightoperator

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE401-DERIVED-EDGE-WEIGHT-OPERATOR-HYPERCHARGE-LAPLACIAN-SIEVE"

	StatusGate400Inherited                  = "CONDITIONAL_SUPPORT_GATE400_MIXED_EDGE_OBSTRUCTION_INHERITED"
	StatusOneFormEdgeSupportInherited       = "CONDITIONAL_SUPPORT_1_FORM_EDGE_SUPPORT_INHERITED"
	StatusElectroweakChargesInherited       = "CONDITIONAL_SUPPORT_ELECTROWEAK_CHARGES_INHERITED"
	StatusNativeEdgeWeightsAudited          = "CONDITIONAL_SUPPORT_NATIVE_EDGE_WEIGHTS_AUDITED"
	StatusDifferentiatedLaplacianFormalized = "CONDITIONAL_SUPPORT_DIFFERENTIATED_EDGE_LAPLACIAN_FORMALIZED"
	StatusAnisotropicEdgeWeightsFound       = "CONDITIONAL_SUPPORT_ANISOTROPIC_EDGE_WEIGHTS_FOUND"
	StatusEdgeResolvedQuarticCapacity       = "CONDITIONAL_SUPPORT_EDGE_RESOLVED_QUARTIC_CAPACITY_FOUND"

	StatusFailedUniformWeightCentral               = "FAILED_ROUTE_UNIFORM_EDGE_WEIGHT_REMAINS_CENTRAL"
	StatusFailedBranchCompressionPairDegenerate    = "FAILED_ROUTE_BRANCH_COMPRESSION_REMAINS_PAIR_DEGENERATE"
	StatusFailedHyperchargePolynomialDisjointQ4    = "FAILED_ROUTE_HYPERCHARGE_EDGE_POLYNOMIAL_DISJOINT_FROM_Q4"
	StatusFailedBMinusLPolynomialDisjointQ4        = "FAILED_ROUTE_B_MINUS_L_EDGE_POLYNOMIAL_DISJOINT_FROM_Q4"
	StatusFailedT3PolynomialNotQ4                  = "FAILED_ROUTE_T3_EDGE_POLYNOMIAL_NOT_Q4"
	StatusFailedNoCanonicalEdgeToHphiComponentMap  = "FAILED_ROUTE_NO_CANONICAL_EDGE_TO_HPHI_COMPONENT_MAP"
	StatusFailedNoNativeQ4WeightedLaplacian        = "FAILED_ROUTE_NO_NATIVE_Q4_WEIGHTED_EDGE_LAPLACIAN"
	StatusFailedNoCanonicalHphiQuarticID           = "FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusFailedNoYukawaCouplingReduction          = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli                = "FIREWALL_PRESERVED_13_MODULI"
	StatusVerifiedCanonicalHphiQuarticID           = "VERIFIED_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusConditionalScalarBundleGeometricallySeal = "CONDITIONAL_SUPPORT_SCALAR_BUNDLE_GEOMETRICALLY_SEALED"
)

const (
	ContactQuarticQ4        = "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271"
	Q4Degree                = 4
	HphiRealDim             = 4
	JDoubledEdgeCount       = 10
	StructuralEdgeCount     = 5
	Gate372ChargedModuliDim = 13
	eps                     = 1e-10
)

var q4Monic = []float64{1, -7668.0 / 3240.0, 6426.0 / 3240.0, -2235.0 / 3240.0, 271.0 / 3240.0}

type Inheritance struct {
	Executed                         bool
	Gate400UniformCentral            bool
	Gate400PairDegenerateCompression bool
	Gate400NoNativeQ4Selector        bool
	Gate385OneFormEdges              bool
	Gate385JDoubledEdgeCount         int
	Gate26YukawaChannelsDerived      bool
	Gate41HyperchargeNormalization   bool
	Gate372ChargedModuliDim          int
	NoEmpiricalInputsImported        bool
	Verdict                          string
}

type Q4Audit struct {
	Polynomial        string
	Degree            int
	IrreducibleOverQ  bool
	MonicCoefficients []float64
	Verdict           string
}

type EdgeClass struct {
	Name              string
	JDouble           string
	ScalarBranch      string
	RightHypercharge  float64
	ScalarHypercharge float64
	BMinusL           float64
	T3Like            float64
	JRealMultiplicity int
	NativeChargeData  bool
}

type EdgeWeightArena struct {
	Formalized                   bool
	StructuralEdgeCount          int
	JDoubledEdgeCount            int
	HphiDimension                int
	Edges                        []EdgeClass
	NativeElectroweakWeights     bool
	NativeBMinusLWeights         bool
	NativeT3Weights              bool
	ExplicitYukawaAmplitudesUsed bool
	ObservedMassesUsed           bool
	Verdict                      string
}

type WeightedCandidate struct {
	Name                       string
	Formula                    string
	WeightSource               string
	NativeWeights              bool
	Sealed                     bool
	Circular                   bool
	HphiEndomorphism           bool
	CanonicalCompressionToHphi bool
	EdgeResolved               bool
	BranchCompressed           bool
	JRealDoubled               bool
	GaugeChargeDerived         bool
	UsesYukawaAmplitudes       bool
	UsesObservedMasses         bool
	Eigenvalues                []float64
	DistinctEigenvalues        int
	MinimalDegree              int
	CharacteristicPolynomial   string
	CharacteristicMonicCoeffs  []float64
	CharacteristicResidualToQ4 float64
	PairDegenerate             bool
	CentralOnHphi              bool
	IrreducibleQuarticCapacity bool
	Q4ExactMatch               bool
	Q4FactorMatch              bool
	PromotableAsQ4Selector     bool
	ReducesYukawaCouplings     bool
	ReducesFlavorModuli        bool
	Reason                     string
	Verdict                    string
}

type WeightSieve struct {
	Executed                    bool
	Candidates                  []WeightedCandidate
	NativeAnisotropicCount      int
	NativeQuarticCapacityCount  int
	CanonicalHphiQ4MatchCount   int
	SealedOrNoncanonicalMatches int
	BestNativeCandidate         string
	BestNativeQ4Residual        float64
	Verdict                     string
}

type Impact struct {
	HphiQuarticIdentified           bool
	ScalarBundleGeometricallySealed bool
	DifferentiatedEdgeWeightsFound  bool
	CanonicalWeightedLaplacianFound bool
	YukawaCouplingsReduced          bool
	ChargedModuliStart              int
	ChargedModuliResult             int
	FlavorFirewallPreserved         bool
	HiggsLanePreserved              bool
	Verdict                         string
}

type FirewallAudit struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoYukawaAmplitudesInserted     bool
	NoManualQ4HphiID               bool
	NoArbitraryEdgeComponentMap    bool
	NoAffineChargeFitPromoted      bool
	NoFlavorModuliReductionClaimed bool
	Verdict                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Q4          Q4Audit
	Arena       EdgeWeightArena
	Sieve       WeightSieve
	Impact      Impact
	Firewall    FirewallAudit
	Next        NextStep
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
	inheritance := inherit()
	q4 := auditQ4()
	arena := buildArena()
	sieve := auditWeights(arena, q4)
	impact := auditImpact(sieve)
	firewall := auditFirewall(sieve, impact)
	next := nextStep(sieve, impact)
	return Analysis{Inheritance: inheritance, Q4: q4, Arena: arena, Sieve: sieve, Impact: impact, Firewall: firewall, Next: next, Truth: truth(sieve, impact)}, nil
}

func inherit() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate400UniformCentral:            true,
		Gate400PairDegenerateCompression: true,
		Gate400NoNativeQ4Selector:        true,
		Gate385OneFormEdges:              true,
		Gate385JDoubledEdgeCount:         JDoubledEdgeCount,
		Gate26YukawaChannelsDerived:      true,
		Gate41HyperchargeNormalization:   true,
		Gate372ChargedModuliDim:          Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:        true,
		Verdict:                          "Gate 401 inherits the Gate-400 q4 obstruction, Gate-385 ten-edge one-form support, the charge-compatible Yukawa edge classes, and the 13-moduli flavor firewall.",
	}
}

func auditQ4() Q4Audit {
	return Q4Audit{Polynomial: ContactQuarticQ4, Degree: Q4Degree, IrreducibleOverQ: true, MonicCoefficients: append([]float64(nil), q4Monic...), Verdict: "The target remains the branch-free irreducible contact primary q4. A differentiated edge Laplacian must match this polynomial directly or by a derived finite theorem, not by an arbitrary affine fit."}
}

func buildArena() EdgeWeightArena {
	edges := []EdgeClass{
		{Name: "Q_L ↔ u_R", JDouble: "J(Q_L ↔ u_R)J⁻¹", ScalarBranch: "Φ_+", RightHypercharge: 2.0 / 3.0, ScalarHypercharge: 0.5, BMinusL: 1.0 / 3.0, T3Like: 0.5, JRealMultiplicity: 2, NativeChargeData: true},
		{Name: "Q_L ↔ d_R", JDouble: "J(Q_L ↔ d_R)J⁻¹", ScalarBranch: "Φ_-", RightHypercharge: -1.0 / 3.0, ScalarHypercharge: -0.5, BMinusL: 1.0 / 3.0, T3Like: -0.5, JRealMultiplicity: 2, NativeChargeData: true},
		{Name: "L_L ↔ e_R", JDouble: "J(L_L ↔ e_R)J⁻¹", ScalarBranch: "Φ_-", RightHypercharge: -1.0, ScalarHypercharge: -0.5, BMinusL: -1.0, T3Like: -0.5, JRealMultiplicity: 2, NativeChargeData: true},
		{Name: "L_L ↔ ν_R", JDouble: "J(L_L ↔ ν_R)J⁻¹", ScalarBranch: "Φ_+", RightHypercharge: 0.0, ScalarHypercharge: 0.5, BMinusL: -1.0, T3Like: 0.5, JRealMultiplicity: 2, NativeChargeData: true},
		{Name: "ν_R ↔ ν_R^c", JDouble: "J(ν_R ↔ ν_R^c)J⁻¹", ScalarBranch: "singlet/Majorana", RightHypercharge: 0.0, ScalarHypercharge: 0.0, BMinusL: 1.0, T3Like: 0.0, JRealMultiplicity: 2, NativeChargeData: true},
	}
	return EdgeWeightArena{
		Formalized:                   true,
		StructuralEdgeCount:          len(edges),
		JDoubledEdgeCount:            2 * len(edges),
		HphiDimension:                HphiRealDim,
		Edges:                        edges,
		NativeElectroweakWeights:     true,
		NativeBMinusLWeights:         true,
		NativeT3Weights:              true,
		ExplicitYukawaAmplitudesUsed: false,
		ObservedMassesUsed:           false,
		Verdict:                      "The finite one-form edge graph has five structural edge classes and ten J-doubled slots. Hypercharge, scalar branch charge, T3-like branch sign, and B-L provide native charge weights; no Yukawa amplitudes or observed masses are used.",
	}
}

func auditWeights(arena EdgeWeightArena, q4 Q4Audit) WeightSieve {
	cands := []WeightedCandidate{
		candidate("uniform J-doubled edge measure", "Δ_E=1 on every one-form edge", "Gate385 support measure", true, false, false, true, true, false, true, true, true, false, false, []float64{1, 1, 1, 1}, "Uniform support is the already-audited Gate400 central case."),
		candidate("scalar branch T3/hypercharge weight", "diag(+1/2,+1/2,-1/2,-1/2)", "Gate21/Gate41 scalar weak charge", true, false, false, true, true, false, true, false, true, false, false, []float64{0.5, 0.5, -0.5, -0.5}, "The scalar weak charge is native but pair-degenerate by the Higgs doublet structure."),
		candidate("branch-averaged right-hypercharge edge Laplacian", "Φ_+:avg(Y_u,Y_ν)=1/3; Φ_-:avg(Y_d,Y_e)=-2/3", "right hypercharge averaged by scalar branch", true, false, false, true, true, false, true, false, true, false, false, []float64{1.0 / 3.0, 1.0 / 3.0, -2.0 / 3.0, -2.0 / 3.0}, "Canonical branch compression forgets which real component came from which Yukawa class, so it remains 2+2."),
		candidate("branch-averaged B-L edge Laplacian", "Φ_+:avg(1/3,-1)=-1/3; Φ_-:avg(1/3,-1)=-1/3", "B-L averaged by scalar branch", true, false, false, true, true, false, true, false, true, false, false, []float64{-1.0 / 3.0, -1.0 / 3.0, -1.0 / 3.0, -1.0 / 3.0}, "B-L is native, but the scalar-branch average collapses to a central value for the Dirac Yukawa edges."),
		candidate("edge-resolved right-hypercharge four-channel stress test", "diag(Y_u,Y_d,Y_ν,Y_e)=diag(2/3,-1/3,0,-1)", "right hypercharge per Yukawa structural class", true, false, true, true, false, true, false, false, true, false, false, []float64{2.0 / 3.0, -1.0 / 3.0, 0, -1.0}, "This gives four distinct values, but placing four edge classes onto the four real H_phi components is a noncanonical component assignment; its polynomial is also disjoint from q4."),
		candidate("edge-resolved squared-hypercharge stress test", "diag(Y_u²,Y_d²,Y_ν²,Y_e²)=diag(4/9,1/9,0,1)", "quadratic hypercharge norm per Yukawa structural class", true, false, true, true, false, true, false, false, true, false, false, []float64{4.0 / 9.0, 1.0 / 9.0, 0, 1.0}, "Charge norms are native diagnostics, but the edge-to-real-scalar-component assignment is not canonical and the polynomial is not q4."),
		candidate("edge-resolved B-L four-channel stress test", "diag(1/3,1/3,-1,-1)", "B-L per Yukawa structural class", true, false, true, true, false, true, false, true, true, false, false, []float64{1.0 / 3.0, 1.0 / 3.0, -1.0, -1.0}, "B-L distinguishes quark/lepton sectors, not four scalar components; it remains 2+2 and cannot be q4."),
		candidate("sealed q4-weighted edge companion", "companion(q4) declared as Δ_E(w) on H_phi", "manual q4 placement", false, true, true, true, false, true, false, false, false, false, false, q4RootsSnapshot(), "The q4 roots can be imposed as four weights, but no edge charge theorem derives them, so this is quarantined."),
	}

	nativeAniso, nativeQuartic, canonical, sealed, bestName, bestResidual := 0, 0, 0, 0, "", math.Inf(1)
	for _, c := range cands {
		if c.NativeWeights && !c.CentralOnHphi {
			nativeAniso++
		}
		if c.NativeWeights && c.IrreducibleQuarticCapacity {
			nativeQuartic++
		}
		if c.NativeWeights && c.CanonicalCompressionToHphi && c.Q4ExactMatch && c.PromotableAsQ4Selector {
			canonical++
		}
		if c.Sealed && c.Q4ExactMatch {
			sealed++
		}
		if c.NativeWeights && c.CharacteristicResidualToQ4 < bestResidual {
			bestResidual = c.CharacteristicResidualToQ4
			bestName = c.Name
		}
	}
	return WeightSieve{
		Executed:                    arena.Formalized && q4.Degree == Q4Degree,
		Candidates:                  cands,
		NativeAnisotropicCount:      nativeAniso,
		NativeQuarticCapacityCount:  nativeQuartic,
		CanonicalHphiQ4MatchCount:   canonical,
		SealedOrNoncanonicalMatches: sealed,
		BestNativeCandidate:         bestName,
		BestNativeQ4Residual:        bestResidual,
		Verdict:                     "Native electroweak and B-L weights do break edge uniformity, but canonical scalar-branch compression remains central or pair-degenerate. Edge-resolved stress tests can reach degree four only by a noncanonical edge-to-H_phi component placement, and their polynomials are disjoint from q4.",
	}
}

func candidate(name, formula, source string, native, sealed, circular, hphi, canonical, edgeResolved, branchCompressed, jReal, gauge, usesYukawa, usesMass bool, eig []float64, reason string) WeightedCandidate {
	eig = canonSlice(eig)
	coeffs := polyFromRoots(eig)
	degree := distinctCount(eig)
	res := coeffResidual(coeffs, q4Monic)
	central := degree == 1
	pair := isPairDegenerate(eig)
	quartic := degree == 4
	q4Exact := res < 1e-8
	promotable := native && !sealed && !circular && hphi && canonical && q4Exact
	verdict := "FAILED_ROUTE"
	if sealed && q4Exact {
		verdict = "SEALED_STRESS_TEST_ONLY"
	} else if promotable {
		verdict = "VERIFIED_Q4_SELECTOR"
	} else if native && quartic {
		verdict = "CONDITIONAL_SUPPORT_QUARTIC_CAPACITY_BUT_NOT_Q4"
	} else if native && !central {
		verdict = "CONDITIONAL_SUPPORT_ANISOTROPIC_WEIGHT_BUT_NOT_Q4"
	}
	return WeightedCandidate{
		Name:                       name,
		Formula:                    formula,
		WeightSource:               source,
		NativeWeights:              native,
		Sealed:                     sealed,
		Circular:                   circular,
		HphiEndomorphism:           hphi,
		CanonicalCompressionToHphi: canonical,
		EdgeResolved:               edgeResolved,
		BranchCompressed:           branchCompressed,
		JRealDoubled:               jReal,
		GaugeChargeDerived:         gauge,
		UsesYukawaAmplitudes:       usesYukawa,
		UsesObservedMasses:         usesMass,
		Eigenvalues:                eig,
		DistinctEigenvalues:        degree,
		MinimalDegree:              degree,
		CharacteristicPolynomial:   formatPolynomial(coeffs),
		CharacteristicMonicCoeffs:  coeffs,
		CharacteristicResidualToQ4: res,
		PairDegenerate:             pair,
		CentralOnHphi:              central,
		IrreducibleQuarticCapacity: quartic,
		Q4ExactMatch:               q4Exact,
		Q4FactorMatch:              false,
		PromotableAsQ4Selector:     promotable,
		ReducesYukawaCouplings:     false,
		ReducesFlavorModuli:        false,
		Reason:                     reason,
		Verdict:                    verdict,
	}
}

func auditImpact(s WeightSieve) Impact {
	identified := s.CanonicalHphiQ4MatchCount > 0
	return Impact{
		HphiQuarticIdentified:           identified,
		ScalarBundleGeometricallySealed: identified,
		DifferentiatedEdgeWeightsFound:  s.NativeAnisotropicCount > 0,
		CanonicalWeightedLaplacianFound: identified,
		YukawaCouplingsReduced:          false,
		ChargedModuliStart:              Gate372ChargedModuliDim,
		ChargedModuliResult:             Gate372ChargedModuliDim,
		FlavorFirewallPreserved:         !identified,
		HiggsLanePreserved:              !identified,
		Verdict:                         "Differentiated charge weights exist, but no canonical q4-weighted scalar Laplacian is derived. Higgs coefficient and flavor moduli ledgers remain unchanged.",
	}
}

func auditFirewall(s WeightSieve, impact Impact) FirewallAudit {
	return FirewallAudit{
		Executed:                       s.Executed,
		NoObservedMassesImported:       true,
		NoCKMImported:                  true,
		NoPMNSImported:                 true,
		NoYukawaAmplitudesInserted:     true,
		NoManualQ4HphiID:               s.CanonicalHphiQ4MatchCount == 0,
		NoArbitraryEdgeComponentMap:    true,
		NoAffineChargeFitPromoted:      true,
		NoFlavorModuliReductionClaimed: impact.ChargedModuliResult == Gate372ChargedModuliDim && !impact.YukawaCouplingsReduced,
		Verdict:                        "No observed masses, CKM/PMNS data, fitted Yukawa amplitudes, arbitrary affine charge fit, or manual q4/H_phi identification is promoted.",
	}
}

func nextStep(s WeightSieve, impact Impact) NextStep {
	if impact.HphiQuarticIdentified {
		return NextStep{Gate: 402, Title: "Scalar Bundle q4 Consequence / Yukawa Reduction Audit", Reason: "A q4 selector would need its physical consequences audited before touching flavor.", PrimaryTask: "Compute whether the identified scalar operator changes one-form kinetic normalization or Yukawa support without importing amplitudes."}
	}
	return NextStep{Gate: 402, Title: "Spectral Graph Edge-Adjacency Operator Search", Reason: "Charge weights differentiate edges but do not produce q4. The missing object is not a gauge charge weight; it is likely a native adjacency/incidence operator on the full one-form edge graph whose quotient to H_phi has four nondegenerate eigenvalues.", PrimaryTask: "Build the finite one-form edge graph adjacency/incidence Laplacian, including edge-edge incidence through shared source/target bimodule nodes, and test its canonical scalar quotient."}
}

func truth(s WeightSieve, impact Impact) string {
	if impact.HphiQuarticIdentified {
		return "Gate 401 derives a canonical q4-weighted scalar Laplacian from native edge weights. This would open a scalar-bundle identity route, while flavor still remains firewalled until Yukawa amplitudes are reduced."
	}
	return fmt.Sprintf("Gate 401 proves that native electroweak/B-L charges can differentiate the ten J-doubled one-form edges, but they do not solve the q4/H_phi identity problem. Canonical scalar-branch compression is still central or 2+2 pair-degenerate; edge-resolved hypercharge can give degree-four capacity only after a noncanonical assignment of edge classes to real H_phi components, and its characteristic polynomial is disjoint from q4 (best native residual %.6g from %s). The 13 charged flavor moduli remain sealed.", s.BestNativeQ4Residual, s.BestNativeCandidate)
}

func Statuses(a Analysis) []string {
	out := []string{
		StatusGate400Inherited,
		StatusOneFormEdgeSupportInherited,
		StatusElectroweakChargesInherited,
		StatusNativeEdgeWeightsAudited,
		StatusDifferentiatedLaplacianFormalized,
	}
	if a.Sieve.NativeAnisotropicCount > 0 {
		out = append(out, StatusAnisotropicEdgeWeightsFound)
	}
	if a.Sieve.NativeQuarticCapacityCount > 0 {
		out = append(out, StatusEdgeResolvedQuarticCapacity)
	}
	out = append(out,
		StatusFailedUniformWeightCentral,
		StatusFailedBranchCompressionPairDegenerate,
		StatusFailedHyperchargePolynomialDisjointQ4,
		StatusFailedBMinusLPolynomialDisjointQ4,
		StatusFailedT3PolynomialNotQ4,
		StatusFailedNoCanonicalEdgeToHphiComponentMap,
		StatusFailedNoNativeQ4WeightedLaplacian,
		StatusFailedNoCanonicalHphiQuarticID,
		StatusFailedNoYukawaCouplingReduction,
		StatusFirewallPreserved13Moduli,
	)
	if a.Impact.HphiQuarticIdentified {
		out = append(out, StatusVerifiedCanonicalHphiQuarticID, StatusConditionalScalarBundleGeometricallySeal)
	}
	return out
}

func q4RootsSnapshot() []float64 {
	return []float64{0.2839121925920062, 0.4411227572843663, 0.7440966379808409, 0.8975350788094533}
}

func canonSlice(xs []float64) []float64 {
	out := append([]float64(nil), xs...)
	for i := range out {
		if math.Abs(out[i]) < eps {
			out[i] = 0
		}
		out[i] = math.Round(out[i]*1e12) / 1e12
	}
	return out
}

func polyFromRoots(roots []float64) []float64 {
	coeffs := []float64{1}
	for _, r := range roots {
		next := make([]float64, len(coeffs)+1)
		for i, c := range coeffs {
			next[i] += c
			next[i+1] += -r * c
		}
		coeffs = next
	}
	for i := range coeffs {
		if math.Abs(coeffs[i]) < 1e-12 {
			coeffs[i] = 0
		}
	}
	return coeffs
}

func coeffResidual(a, b []float64) float64 {
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	s := 0.0
	for i := 0; i < n; i++ {
		av, bv := 0.0, 0.0
		if i < len(a) {
			av = a[i]
		}
		if i < len(b) {
			bv = b[i]
		}
		d := av - bv
		s += d * d
	}
	return math.Sqrt(s)
}

func distinctCount(xs []float64) int {
	distinct := []float64{}
	for _, x := range xs {
		found := false
		for _, y := range distinct {
			if math.Abs(x-y) < 1e-10 {
				found = true
				break
			}
		}
		if !found {
			distinct = append(distinct, x)
		}
	}
	return len(distinct)
}

func isPairDegenerate(xs []float64) bool {
	if len(xs) != 4 {
		return false
	}
	counts := map[string]int{}
	for _, x := range xs {
		counts[fmt.Sprintf("%.12g", x)]++
	}
	if len(counts) != 2 {
		return false
	}
	for _, c := range counts {
		if c != 2 {
			return false
		}
	}
	return true
}

func formatPolynomial(coeffs []float64) string {
	parts := []string{}
	deg := len(coeffs) - 1
	for i, c := range coeffs {
		if math.Abs(c) < 1e-12 {
			continue
		}
		pow := deg - i
		term := fmt.Sprintf("%.12g", c)
		if pow > 0 {
			term += "*x"
			if pow > 1 {
				term += fmt.Sprintf("^%d", pow)
			}
		}
		parts = append(parts, term)
	}
	if len(parts) == 0 {
		return "0"
	}
	return strings.Join(parts, " + ")
}

func formatFloatSlice(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.12g", x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func edgeNames(edges []EdgeClass) string {
	xs := make([]string, 0, len(edges))
	for _, e := range edges {
		xs = append(xs, fmt.Sprintf("%s/Y_R=%.6g/B-L=%.6g/T3=%.6g/branch=%s", e.Name, e.RightHypercharge, e.BMinusL, e.T3Like, e.ScalarBranch))
	}
	sort.Strings(xs)
	return "[" + strings.Join(xs, "; ") + "]"
}
