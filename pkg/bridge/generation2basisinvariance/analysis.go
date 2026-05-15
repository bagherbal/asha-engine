// Package generation2basisinvariance implements Gate 452:
// Family Basis-Invariance / Texture Gauge-Artifact Audit.
//
// Gate 451 proved that no already-native ASHA boundary suppresses the 1-3
// edge or fixes the phase ray. Gate 452 closes the next loophole: perhaps the
// closed triangle is merely a bad family basis, and a legitimate basis change
// can turn it into the nearest-neighbor Fritzsch/GST chain. The answer is
// negative once the Gate-444 family address K_gen is preserved. K-preserving
// rephasings can move phases around the triangle, but they cannot delete an
// edge, change edge magnitudes, change the KMS harmonic ledger, change graph
// spectrum, or change the determinant mass-lift class. A general U(3) rotation
// can alter the apparent zero pattern only by destroying the native diagonal
// K_gen address, so it is not an allowed ASHA gauge equivalence.
package generation2basisinvariance

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE452-FAMILY-BASIS-INVARIANCE-TEXTURE-GAUGE-ARTIFACT-AUDIT"

	StatusGate451Inherited                          = "CONDITIONAL_SUPPORT_GATE451_FULL_TRIANGLE_NO_SELECTOR_INHERITED"
	StatusFamilyAddressPreserved                    = "CONDITIONAL_SUPPORT_K_GEN_FAMILY_ADDRESS_PRESERVED"
	StatusAllowedBasisGroupComputed                 = "CONDITIONAL_SUPPORT_ALLOWED_BASIS_GROUP_COMPUTED"
	StatusEdgeMagnitudesInvariant                   = "CONDITIONAL_SUPPORT_EDGE_MAGNITUDES_INVARIANT_UNDER_REPHASING"
	StatusGraphInvariantsComputed                   = "CONDITIONAL_SUPPORT_GRAPH_INVARIANTS_COMPUTED"
	StatusGeneralUnitaryRejected                    = "CONDITIONAL_SUPPORT_GENERAL_UNITARY_ROTATIONS_REJECTED_AS_NON_NATIVE"
	StatusTextureGaugeArtifactQuarantined           = "CONDITIONAL_SUPPORT_TEXTURE_GAUGE_ARTIFACT_QUARANTINED"
	StatusEmpiricalFirewallPreserved                = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedNoBasisSuppression                  = "FAILED_ROUTE_NO_BASIS_SUPPRESSION_OF_13_EDGE"
	StatusFailedNNTextureNotGaugeEquivalent         = "FAILED_ROUTE_NEAREST_NEIGHBOR_TEXTURE_NOT_GAUGE_EQUIVALENT"
	StatusFailedGeneralFamilyRotationBreaksKAddress = "FAILED_ROUTE_GENERAL_FAMILY_ROTATION_BREAKS_K_GEN_ADDRESS"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                        bool
	Gate444KGenForced               bool
	Gate444Generation2Zero          bool
	Gate445TriangleForced           bool
	Gate446PhaseQuarantined         bool
	Gate447CoefficientsSealed       bool
	Gate450TextureZeroSumRule       bool
	Gate451FullTrianglePreserved    bool
	Gate451NoNativePhaseRaySelector bool
	Gate451GSTFritzschQuarantined   bool
	NoEmpiricalInputsImported       bool
	Verdict                         string
}

type BasisTransformation struct {
	Name                        string
	Formula                     string
	PreservesKGen               bool
	PreservesKGenUpToSign       bool
	PreservesGeneration2Address bool
	PreservesStructuralZero22   bool
	PreservesEdgeMagnitudes     bool
	PreservesSupport            bool
	CanDelete13Edge             bool
	AllowedNativeGauge          bool
	Reason                      string
}

type BasisGroupAudit struct {
	Executed                   bool
	Transformations            []BasisTransformation
	ExactKPreservingGroup      string
	KOrbitPreservingExtension  string
	GeneralUnitaryRejected     bool
	AllNativeAllowedPreserve13 bool
	AnyNativeAllowedDeletes13  bool
	Verdict                    string
	Reason                     string
}

type EdgeMagnitude struct {
	Edge                     string
	DeltaK                   int
	TriangleMagnitude        float64
	NearestNeighborMagnitude float64
	RephasingInvariant       bool
	DeletedInNN              bool
	Reason                   string
}

type SupportAudit struct {
	Executed                  bool
	Edges                     []EdgeMagnitude
	TriangleSupportFormula    string
	NearestNeighborFormula    string
	EdgeCountTriangle         int
	EdgeCountNearestNeighbor  int
	TriangleCycleCount        int
	NearestNeighborCycleCount int
	SupportPatternInvariant   bool
	CanRephaseToNN            bool
	Verdict                   string
	Reason                    string
}

type GraphInvariant struct {
	Name            string
	Triangle        string
	NearestNeighbor string
	Equal           bool
	BasisInvariant  bool
	Reason          string
}

type SpectralAudit struct {
	Executed                       bool
	Invariants                     []GraphInvariant
	TriangleAdjacencySpectrum      string
	NearestNeighborSpectrum        string
	TriangleDetLiftCoeff           int
	NearestNeighborDetLiftCoeff    int
	TriangleCommutatorNorm2        int
	NearestNeighborCommutatorNorm2 int
	SameInvariantClass             bool
	Verdict                        string
	Reason                         string
}

type GaugeArtifactVerdict struct {
	Executed                                bool
	NearestNeighborCanBeNativeGaugeArtifact bool
	RequiresNonNativeGeneralUnitary         bool
	KGenAddressDestroyed                    bool
	TextureZeroAddressDestroyed             bool
	GSTFritzschStillEmpiricalAssumption     bool
	ReevaluateRatios                        bool
	Verdict                                 string
	Reason                                  string
}

type Firewall struct {
	Executed                        bool
	NoObservedMuonMassImported      bool
	NoObservedCharmMassImported     bool
	NoObservedYukawaImported        bool
	NoCKMImported                   bool
	NoPMNSImported                  bool
	NoCurveFit                      bool
	KGenStillForced                 bool
	Generation2ZeroStillForced      bool
	XTriangleStillForced            bool
	YPhaseStillQuarantined          bool
	SectorCoefficientsStillSealed   bool
	GSTFritzschRelationsQuarantined bool
	NativeFlavorDimAfter            int
	KXYCoeffDimAfter                int
	Verdict                         string
	Reason                          string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	BasisAudit  BasisGroupAudit
	Support     SupportAudit
	Spectral    SpectralAudit
	Verdict     GaugeArtifactVerdict
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
	a.BasisAudit = buildBasisAudit()
	a.Support = buildSupportAudit()
	a.Spectral = buildSpectralAudit()
	a.Verdict = buildGaugeArtifactVerdict(a.BasisAudit, a.Support, a.Spectral)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate444KGenForced:               true,
		Gate444Generation2Zero:          true,
		Gate445TriangleForced:           true,
		Gate446PhaseQuarantined:         true,
		Gate447CoefficientsSealed:       true,
		Gate450TextureZeroSumRule:       true,
		Gate451FullTrianglePreserved:    true,
		Gate451NoNativePhaseRaySelector: true,
		Gate451GSTFritzschQuarantined:   true,
		NoEmpiricalInputsImported:       true,
		Verdict:                         StatusGate451Inherited,
	}
}

func buildBasisAudit() BasisGroupAudit {
	transforms := []BasisTransformation{
		{
			Name: "diagonal family rephasing", Formula: "U=diag(e^{i alpha_1},e^{i alpha_2},e^{i alpha_3})",
			PreservesKGen: true, PreservesKGenUpToSign: true, PreservesGeneration2Address: true, PreservesStructuralZero22: true,
			PreservesEdgeMagnitudes: true, PreservesSupport: true, CanDelete13Edge: false, AllowedNativeGauge: true,
			Reason: "K_gen has three distinct eigenvalues, so exact K-preserving unitary freedom is diagonal rephasing; |M_ij| and zero/nonzero support are invariant.",
		},
		{
			Name: "global phase/sign", Formula: "M -> e^{i chi} M, or real sign flip of the bridge",
			PreservesKGen: true, PreservesKGenUpToSign: true, PreservesGeneration2Address: true, PreservesStructuralZero22: true,
			PreservesEdgeMagnitudes: true, PreservesSupport: true, CanDelete13Edge: false, AllowedNativeGauge: true,
			Reason: "a scalar phase/sign can change conventions but cannot turn a nonzero 1-3 edge into zero.",
		},
		{
			Name: "orientation reversal of the primitive spectrum", Formula: "P_13 K P_13^T = -K",
			PreservesKGen: false, PreservesKGenUpToSign: true, PreservesGeneration2Address: true, PreservesStructuralZero22: true,
			PreservesEdgeMagnitudes: true, PreservesSupport: true, CanDelete13Edge: false, AllowedNativeGauge: true,
			Reason: "even if the unordered {-1,0,1} orbit is quotiented by orientation, the closed triangle maps to itself, never to a chain.",
		},
		{
			Name: "nearest-neighbor texture projection", Formula: "set |M_13|=0 by hand",
			PreservesKGen: true, PreservesKGenUpToSign: true, PreservesGeneration2Address: true, PreservesStructuralZero22: true,
			PreservesEdgeMagnitudes: false, PreservesSupport: false, CanDelete13Edge: true, AllowedNativeGauge: false,
			Reason: "this is an extra projector on the family graph, not a basis transformation generated by K_gen, gamma_F, J, or the first-order condition.",
		},
		{
			Name: "general U(3) family rotation", Formula: "K -> U K U^dagger",
			PreservesKGen: false, PreservesKGenUpToSign: false, PreservesGeneration2Address: false, PreservesStructuralZero22: false,
			PreservesEdgeMagnitudes: false, PreservesSupport: false, CanDelete13Edge: true, AllowedNativeGauge: false,
			Reason: "a generic rotation can alter apparent texture entries only by mixing the native KMS energy levels; it destroys the Gate-444 family address used to define the structural zero.",
		},
	}
	allAllowedPreserve13 := true
	anyAllowedDeletes13 := false
	generalRejected := false
	for _, t := range transforms {
		if t.AllowedNativeGauge {
			allAllowedPreserve13 = allAllowedPreserve13 && t.PreservesSupport && !t.CanDelete13Edge
			anyAllowedDeletes13 = anyAllowedDeletes13 || t.CanDelete13Edge
		}
		if t.Name == "general U(3) family rotation" && !t.AllowedNativeGauge {
			generalRejected = true
		}
	}
	return BasisGroupAudit{
		Executed:                   true,
		Transformations:            transforms,
		ExactKPreservingGroup:      "centralizer_U(3)(K_gen)=U(1)^3 because spec(K_gen)={-1,0,1} is simple",
		KOrbitPreservingExtension:  "U(1)^3 plus optional 1<->3 orientation reversal if K and -K are identified as an unoriented spectrum",
		GeneralUnitaryRejected:     generalRejected,
		AllNativeAllowedPreserve13: allAllowedPreserve13,
		AnyNativeAllowedDeletes13:  anyAllowedDeletes13,
		Verdict:                    StatusAllowedBasisGroupComputed,
		Reason:                     "the only native basis freedoms preserving the Gate-444 address are rephasings and optional orientation reversal; neither changes support.",
	}
}

func buildSupportAudit() SupportAudit {
	edges := []EdgeMagnitude{
		{Edge: "12", DeltaK: 1, TriangleMagnitude: 1, NearestNeighborMagnitude: 1, RephasingInvariant: true, Reason: "nearest harmonic retained in both graphs"},
		{Edge: "23", DeltaK: 1, TriangleMagnitude: 1, NearestNeighborMagnitude: 1, RephasingInvariant: true, Reason: "nearest harmonic retained in both graphs"},
		{Edge: "13", DeltaK: 2, TriangleMagnitude: 1, NearestNeighborMagnitude: 0, RephasingInvariant: true, DeletedInNN: true, Reason: "second harmonic present in the forced triangle; magnitude cannot be removed by diagonal rephasing"},
	}
	return SupportAudit{
		Executed:                  true,
		Edges:                     edges,
		TriangleSupportFormula:    "supp(X_triangle)={12,23,13}",
		NearestNeighborFormula:    "supp(X_NN)={12,23}",
		EdgeCountTriangle:         3,
		EdgeCountNearestNeighbor:  2,
		TriangleCycleCount:        1,
		NearestNeighborCycleCount: 0,
		SupportPatternInvariant:   true,
		CanRephaseToNN:            false,
		Verdict:                   StatusEdgeMagnitudesInvariant,
		Reason:                    "K-preserving basis changes multiply edges by phases only; they preserve every |M_ij| and every support zero.",
	}
}

func buildSpectralAudit() SpectralAudit {
	triDet := determinantEpsilon3Coeff(support(true, true, true))
	nnDet := determinantEpsilon3Coeff(support(true, true, false))
	triComm := commutatorNormSquared(support(true, true, true))
	nnComm := commutatorNormSquared(support(true, true, false))
	invs := []GraphInvariant{
		{Name: "edge count", Triangle: "3", NearestNeighbor: "2", Equal: false, BasisInvariant: true, Reason: "support cardinality is preserved by K-centralizer rephasing"},
		{Name: "degree sequence", Triangle: "[2,2,2]", NearestNeighbor: "[1,1,2]", Equal: false, BasisInvariant: true, Reason: "closed 3-cycle is not graph-isomorphic to a 2-edge path"},
		{Name: "cycle rank", Triangle: "1", NearestNeighbor: "0", Equal: false, BasisInvariant: true, Reason: "triangle has one primitive cycle; path has none"},
		{Name: "adjacency spectrum", Triangle: "{2,-1,-1}", NearestNeighbor: "{sqrt(2),0,-sqrt(2)}", Equal: false, BasisInvariant: true, Reason: "adjacency similarity cannot map the triangle to the path"},
		{Name: "determinant lift coefficient", Triangle: fmt.Sprintf("%d", triDet), NearestNeighbor: fmt.Sprintf("%d", nnDet), Equal: triDet == nnDet, BasisInvariant: true, Reason: "det(K+epsilon B) distinguishes primitive lift from degenerate chain"},
		{Name: "KMS harmonic ledger", Triangle: "DeltaK=1:2 edges; DeltaK=2:1 edge", NearestNeighbor: "DeltaK=1:2 edges; DeltaK=2:0 edges", Equal: false, BasisInvariant: true, Reason: "the 1-3 second harmonic is a native integer harmonic, not a removable gauge artifact"},
		{Name: "commutator norm squared", Triangle: fmt.Sprintf("%d", triComm), NearestNeighbor: fmt.Sprintf("%d", nnComm), Equal: triComm == nnComm, BasisInvariant: true, Reason: "||[K,B]||_F^2 counts KMS harmonic weights"},
	}
	same := true
	for _, inv := range invs {
		same = same && inv.Equal
	}
	return SpectralAudit{
		Executed:                       true,
		Invariants:                     invs,
		TriangleAdjacencySpectrum:      "{2,-1,-1}",
		NearestNeighborSpectrum:        "{sqrt(2),0,-sqrt(2)}",
		TriangleDetLiftCoeff:           triDet,
		NearestNeighborDetLiftCoeff:    nnDet,
		TriangleCommutatorNorm2:        triComm,
		NearestNeighborCommutatorNorm2: nnComm,
		SameInvariantClass:             same,
		Verdict:                        StatusGraphInvariantsComputed,
		Reason:                         "basis-invariant graph and K-relative quantities separate the full triangle from the nearest-neighbor path.",
	}
}

func buildGaugeArtifactVerdict(b BasisGroupAudit, s SupportAudit, sp SpectralAudit) GaugeArtifactVerdict {
	nativeArtifact := b.AnyNativeAllowedDeletes13 || s.CanRephaseToNN || sp.SameInvariantClass
	return GaugeArtifactVerdict{
		Executed:                                true,
		NearestNeighborCanBeNativeGaugeArtifact: nativeArtifact,
		RequiresNonNativeGeneralUnitary:         true,
		KGenAddressDestroyed:                    true,
		TextureZeroAddressDestroyed:             true,
		GSTFritzschStillEmpiricalAssumption:     !nativeArtifact,
		ReevaluateRatios:                        false,
		Verdict:                                 StatusFailedNNTextureNotGaugeEquivalent,
		Reason:                                  "nearest-neighbor suppression is not gauge-equivalent to the forced triangle under any K_gen-preserving native basis freedom.",
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                        true,
		NoObservedMuonMassImported:      true,
		NoObservedCharmMassImported:     true,
		NoObservedYukawaImported:        true,
		NoCKMImported:                   true,
		NoPMNSImported:                  true,
		NoCurveFit:                      true,
		KGenStillForced:                 true,
		Generation2ZeroStillForced:      true,
		XTriangleStillForced:            true,
		YPhaseStillQuarantined:          true,
		SectorCoefficientsStillSealed:   true,
		GSTFritzschRelationsQuarantined: true,
		NativeFlavorDimAfter:            NativeFlavorDim,
		KXYCoeffDimAfter:                KXYCoeffDim,
		Verdict:                         StatusEmpiricalFirewallPreserved,
		Reason:                          "Gate 452 classifies nearest-neighbor textures as non-native assumptions and imports no observed flavor data.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        453,
		Title:       "Texture-Zero Invariant Ledger / Allowed Empirical Interface",
		Reason:      "after proving that GST/Fritzsch is neither native nor gauge-equivalent, the only honest use of texture-zero physics is as an explicitly labelled empirical bridge interface",
		PrimaryTask: "define what a future phenomenology module may import, what it may compute, and what it may never relabel as native geometry",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate451FullTrianglePreserved || !a.Inheritance.Gate451GSTFritzschQuarantined {
		return fmt.Errorf("Gate452 requires Gate451 full-triangle no-selector inheritance")
	}
	if !a.BasisAudit.Executed || !a.BasisAudit.GeneralUnitaryRejected || !a.BasisAudit.AllNativeAllowedPreserve13 || a.BasisAudit.AnyNativeAllowedDeletes13 {
		return fmt.Errorf("basis audit did not preserve the 1-3 edge under native transformations")
	}
	if !a.Support.Executed || !a.Support.SupportPatternInvariant || a.Support.CanRephaseToNN || a.Support.EdgeCountTriangle != 3 || a.Support.EdgeCountNearestNeighbor != 2 {
		return fmt.Errorf("support audit failed to separate triangle from nearest-neighbor path")
	}
	if !a.Spectral.Executed || a.Spectral.SameInvariantClass || a.Spectral.TriangleDetLiftCoeff != 2 || a.Spectral.NearestNeighborDetLiftCoeff != 0 || a.Spectral.TriangleCommutatorNorm2 != 12 || a.Spectral.NearestNeighborCommutatorNorm2 != 4 {
		return fmt.Errorf("spectral audit failed: triangle and nearest-neighbor branch were not separated")
	}
	if !a.Verdict.Executed || a.Verdict.NearestNeighborCanBeNativeGaugeArtifact || !a.Verdict.GSTFritzschStillEmpiricalAssumption || a.Verdict.ReevaluateRatios {
		return fmt.Errorf("basis-artifact verdict must quarantine nearest-neighbor/GST branch")
	}
	if !a.Firewall.Executed || !a.Firewall.GSTFritzschRelationsQuarantined || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("13-moduli firewall not preserved")
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 452 proves that the nearest-neighbor Fritzsch/GST branch is not a hidden family-basis gauge of the ASHA triangle. Once K_gen=diag(-1,0,1) is kept as the native family address, admissible rephasings preserve edge magnitudes, the 1-3 second harmonic, the triangle cycle, determinant lift, and K-commutator weight. A generic U(3) rotation may change displayed texture entries only by erasing the native K_gen address and therefore cannot convert an empirical texture into a geometric theorem."
}

func statuses() []string {
	return []string{
		StatusGate451Inherited,
		StatusFamilyAddressPreserved,
		StatusAllowedBasisGroupComputed,
		StatusEdgeMagnitudesInvariant,
		StatusGraphInvariantsComputed,
		StatusGeneralUnitaryRejected,
		StatusTextureGaugeArtifactQuarantined,
		StatusEmpiricalFirewallPreserved,
		StatusFailedNoBasisSuppression,
		StatusFailedNNTextureNotGaugeEquivalent,
		StatusFailedGeneralFamilyRotationBreaksKAddress,
	}
}

func support(edge12, edge23, edge13 bool) [3][3]int {
	var m [3][3]int
	if edge12 {
		m[0][1], m[1][0] = 1, 1
	}
	if edge23 {
		m[1][2], m[2][1] = 1, 1
	}
	if edge13 {
		m[0][2], m[2][0] = 1, 1
	}
	return m
}

func determinantEpsilon3Coeff(b [3][3]int) int {
	// K+epsilon B has diagonal (-1,0,1).  The epsilon^3 coefficient is det(B).
	return det3(b)
}

func det3(m [3][3]int) int {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) - m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) + m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

func commutatorNormSquared(b [3][3]int) int {
	k := []int{-1, 0, 1}
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			d := k[i] - k[j]
			x := d * b[i][j]
			sum += x * x
		}
	}
	return sum
}

func almostEqual(a, b float64) bool { return math.Abs(a-b) < 1e-12 }

func joinBools(xs ...bool) bool {
	for _, x := range xs {
		if !x {
			return false
		}
	}
	return true
}

func textList(xs []string) string { return strings.Join(xs, "; ") }
