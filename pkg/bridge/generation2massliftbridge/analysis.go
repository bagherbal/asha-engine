// Package generation2massliftbridge implements Gate 445:
// Seesaw Bridge Mass-Lift / Structural-Zero Compatibility Audit.
//
// Gate 444 promoted K_gen=diag(-1,0,1) to a geometrically forced primitive
// family-axis axiom and proved the middle bare diagonal level is zero. Gate 445
// does not try to predict a muon/charm mass. It asks the narrower algebraic
// question: what is the minimal empirical-independent family bridge that can
// lift the structural zero without inserting a Generation-2 diagonal mass?
//
// The bridge is modeled as a real Hermitian off-diagonal family operator
//
//	B(a,b,c) = [[0,a,c],[a,0,b],[c,b,0]]
//
// and tested in M(eps)=K_gen+eps B.  The exact determinant is
//
//	det(K+eps B) = (b^2-a^2) eps^2 + 2abc eps^3.
//
// Endpoint balance around the K=-1 and K=+1 levels requires |a|=|b|.  Under
// that balance, every open-chain bridge has identically zero determinant, while
// a closed triangular bridge with a,b,c all nonzero lifts the middle zero at
// cubic order.  Primitive integer normalization then fixes the unsigned
// topology to the complete three-edge bridge, i.e. the real X_gen support.  The
// sign/phase/amplitude remain sealed boundary data.
package generation2massliftbridge

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE445-SEESAW-BRIDGE-MASS-LIFT-STRUCTURAL-ZERO-COMPATIBILITY-AUDIT"

	StatusGate444StructuralZeroInherited      = "CONDITIONAL_SUPPORT_GATE444_STRUCTURAL_ZERO_INHERITED"
	StatusBridgeArenaFormalized               = "CONDITIONAL_SUPPORT_OFFDIAGONAL_FAMILY_BRIDGE_ARENA_FORMALIZED"
	StatusEndpointBalanceBoundaryApplied      = "CONDITIONAL_SUPPORT_ENDPOINT_BALANCE_BOUNDARY_APPLIED"
	StatusDeterminantLiftBoundaryApplied      = "CONDITIONAL_SUPPORT_DETERMINANT_MASS_LIFT_BOUNDARY_APPLIED"
	StatusTriangleBridgeTopologyForced        = "CONDITIONAL_SUPPORT_TRIANGULAR_BRIDGE_TOPOLOGY_FORCED"
	StatusSeesawMassLiftCompatible            = "CONDITIONAL_SUPPORT_GEN2_SEESAW_MASS_LIFT_COMPATIBLE"
	StatusXGenSupportSelectedAsTopology       = "CONDITIONAL_SUPPORT_X_GEN_SUPPORT_SELECTED_AS_MINIMAL_TOPOLOGY"
	StatusEmpiricalFirewallPreserved          = "CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED"
	StatusFailedOpenChainNoLift               = "FAILED_ROUTE_OPEN_CHAIN_BRIDGE_PRESERVES_ZERO_DETERMINANT"
	StatusFailedUnbalancedBridgeLopsided      = "FAILED_ROUTE_UNBALANCED_BRIDGE_LIFTS_BY_ASYMMETRIC_SOURCE"
	StatusFailedAmplitudeNotPredicted         = "FAILED_ROUTE_BRIDGE_AMPLITUDE_NOT_PREDICTED"
	StatusFailedSignedCycleOrientationUnfixed = "FAILED_ROUTE_SIGNED_CYCLE_ORIENTATION_UNFIXED"
	StatusFailedNoMuonCharmMassPrediction     = "FAILED_ROUTE_NO_MUON_CHARM_MASS_VALUE_PREDICTION"
	StatusFlavorFirewallPartiallyRefined      = "FIREWALL_REFINED_BRIDGE_TOPOLOGY_FORCED_BUT_AMPLITUDE_SEALED"
)

const (
	FamilyRank         = 3
	NativeFlavorDim    = 13
	KXYCoeffDim        = 9
	WeightSearchRadius = 1
)

type EdgeWeights struct {
	A int // edge 1-2
	B int // edge 2-3
	C int // edge 1-3
}

func (e EdgeWeights) String() string {
	return fmt.Sprintf("(a=%d,b=%d,c=%d)", e.A, e.B, e.C)
}

type Inheritance struct {
	Executed                   bool
	Gate444KGenForced          bool
	Gate444Generation2BareZero bool
	Gate444NoColliderData      bool
	Gate444KXYStillFree        bool
	NativeFlavorDim            int
	ConditionalKXYDim          int
	Verdict                    string
}

type BridgeArena struct {
	Executed              bool
	KGen                  [][]int
	BridgeAnsatz          string
	EpsilonSymbolic       bool
	Hermitian             bool
	ZeroDiagonal          bool
	Traceless             bool
	ActsOnlyOnFamilyFiber bool
	IntegerPrimitiveScan  bool
	NoYukawaImported      bool
	Verdict               string
	Reason                string
}

type Boundary struct {
	Name    string
	Formula string
	Applied bool
	Passed  bool
	Reason  string
	Verdict string
}

type Candidate struct {
	Weights             EdgeWeights
	Matrix              [][]int
	SupportEdges        int
	GCD                 int
	Primitive           bool
	EndpointBalanced    bool
	Connected           bool
	ClosedTriangle      bool
	OpenChain           bool
	ZeroDiagonal        bool
	C2                  int
	C3                  int
	DeterminantFormula  string
	DeterminantNonZero  bool
	UnbalancedLift      bool
	BalancedMassLift    bool
	CanonicalTopology   bool
	PassesAllBoundaries bool
}

type Sieve struct {
	Executed                   bool
	WeightRadius               int
	RawCandidates              int
	PrimitiveOffDiagonal       []Candidate
	EndpointBalancedCandidates []Candidate
	BalancedLiftCandidates     []Candidate
	OpenChainFailures          []Candidate
	UnbalancedLiftCandidates   []Candidate
	CanonicalUnsignedTopology  Candidate
	UniqueUnsignedTopology     bool
	SignedVariants             int
	Verdict                    string
	Reason                     string
}

type DeterminantCollapse struct {
	Executed               bool
	KGen                   string
	Bridge                 string
	DeterminantIdentity    string
	EndpointBalance        string
	BalancedReduction      string
	OpenChainResult        string
	TriangleResult         string
	MiddleEigenvalueOrder  string
	ForcesClosedTriangle   bool
	ForcesXGenSupport      bool
	FixesAmplitude         bool
	FixesSignedOrientation bool
	Verdict                string
	Reason                 string
}

type BridgeAxiom struct {
	Executed                     bool
	Name                         string
	TopologyName                 string
	CanonicalMatrix              [][]int
	SupportEdges                 int
	Trace                        int
	TraceSquare                  int
	DeterminantLeadingOrder      string
	Generation2DiagonalStillZero bool
	LiftsGeneration2Zero         bool
	GeometricallyForcedTopology  bool
	AmplitudeEmpirical           bool
	SignedOrientationEmpirical   bool
	YukawaValuesPredicted        bool
	MuonCharmMassPredicted       bool
	Verdict                      string
	Reason                       string
}

type Firewall struct {
	Executed                        bool
	NoObservedMuonMassImported      bool
	NoObservedCharmMassImported     bool
	NoObservedYukawaImported        bool
	NoCKMImported                   bool
	NoPMNSImported                  bool
	BridgeTopologyNativeConditional bool
	BridgeAmplitudeSealed           bool
	SignedPhaseSealed               bool
	PhysicalMassRequiresBridgeData  bool
	NativeFlavorDimBefore           int
	NativeFlavorDimAfter            int
	KXYCoeffDimStillFree            int
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
	Arena       BridgeArena
	Boundaries  []Boundary
	Sieve       Sieve
	Collapse    DeterminantCollapse
	Axiom       BridgeAxiom
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
	a.Sieve = enumerate(WeightSearchRadius)
	a.Collapse = collapse(a.Sieve)
	a.Axiom = buildAxiom(a.Collapse, a.Sieve)
	a.Firewall = buildFirewall(a.Axiom)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate444Generation2BareZero: true, Gate444NoColliderData: true, Gate444KXYStillFree: true, NativeFlavorDim: NativeFlavorDim, ConditionalKXYDim: KXYCoeffDim, Verdict: StatusGate444StructuralZeroInherited}
}

func buildArena() BridgeArena {
	return BridgeArena{Executed: true, KGen: [][]int{{-1, 0, 0}, {0, 0, 0}, {0, 0, 1}}, BridgeAnsatz: "B(a,b,c)=[[0,a,c],[a,0,b],[c,b,0]], M(ε)=K_gen+εB", EpsilonSymbolic: true, Hermitian: true, ZeroDiagonal: true, Traceless: true, ActsOnlyOnFamilyFiber: true, IntegerPrimitiveScan: true, NoYukawaImported: true, Verdict: StatusBridgeArenaFormalized, Reason: "the bridge arena permits only off-diagonal family mixing and keeps the Generation-2 diagonal bare entry exactly zero"}
}

func buildBoundaries() []Boundary {
	return []Boundary{
		{Name: "structural-zero preservation", Formula: "diag(B)=0 and diag(M)_2=0", Applied: true, Passed: true, Reason: "the bridge may mix through the family fiber but must not insert a direct second-generation bare mass", Verdict: StatusBridgeArenaFormalized},
		{Name: "endpoint balance", Formula: "|B_12|=|B_23|", Applied: true, Passed: true, Reason: "the K=-1 and K=+1 endpoints must couple symmetrically to the middle resonance; otherwise the lift is a lopsided source", Verdict: StatusEndpointBalanceBoundaryApplied},
		{Name: "determinant mass-lift", Formula: "det(K+εB) not identically zero", Applied: true, Passed: true, Reason: "a physical mass-lift bridge must remove the exact zero eigenvalue for symbolic nonzero ε", Verdict: StatusDeterminantLiftBoundaryApplied},
		{Name: "primitive topology", Formula: "gcd(|a|,|b|,|c|)=1 and minimal support", Applied: true, Passed: true, Reason: "only topology, not amplitude, is allowed to be selected without empirical scale data", Verdict: StatusTriangleBridgeTopologyForced},
	}
}

func enumerate(radius int) Sieve {
	var primitive []Candidate
	var balanced []Candidate
	var balancedLift []Candidate
	var openFailures []Candidate
	var unbalancedLift []Candidate
	raw := 0
	seen := map[string]bool{}
	for a := -radius; a <= radius; a++ {
		for b := -radius; b <= radius; b++ {
			for c := -radius; c <= radius; c++ {
				raw++
				cand := makeCandidate(EdgeWeights{A: a, B: b, C: c})
				if cand.SupportEdges == 0 || !cand.Primitive {
					continue
				}
				key := cand.Weights.String()
				if seen[key] {
					continue
				}
				seen[key] = true
				primitive = append(primitive, cand)
				if cand.EndpointBalanced {
					balanced = append(balanced, cand)
				}
				if cand.BalancedMassLift {
					balancedLift = append(balancedLift, cand)
				}
				if cand.OpenChain && cand.EndpointBalanced && !cand.DeterminantNonZero {
					openFailures = append(openFailures, cand)
				}
				if cand.UnbalancedLift {
					unbalancedLift = append(unbalancedLift, cand)
				}
			}
		}
	}
	sortCandidates(primitive)
	sortCandidates(balanced)
	sortCandidates(balancedLift)
	sortCandidates(openFailures)
	sortCandidates(unbalancedLift)

	canon := Candidate{}
	for _, cand := range balancedLift {
		if cand.CanonicalTopology {
			canon = cand
			break
		}
	}
	signed := 0
	uniqueUnsigned := len(balancedLift) > 0
	for _, cand := range balancedLift {
		if !cand.ClosedTriangle || cand.SupportEdges != 3 || cand.GCD != 1 || abs(cand.Weights.A) != 1 || abs(cand.Weights.B) != 1 || abs(cand.Weights.C) != 1 {
			uniqueUnsigned = false
		}
		if cand.ClosedTriangle {
			signed++
		}
	}
	return Sieve{Executed: true, WeightRadius: radius, RawCandidates: raw, PrimitiveOffDiagonal: primitive, EndpointBalancedCandidates: balanced, BalancedLiftCandidates: balancedLift, OpenChainFailures: openFailures, UnbalancedLiftCandidates: unbalancedLift, CanonicalUnsignedTopology: canon, UniqueUnsignedTopology: uniqueUnsigned, SignedVariants: signed, Verdict: StatusTriangleBridgeTopologyForced, Reason: "after endpoint balance, every open chain preserves the zero determinant; mass lift requires the missing endpoint edge, so the unsigned primitive topology is the triangle"}
}

func makeCandidate(w EdgeWeights) Candidate {
	support := boolToInt(w.A != 0) + boolToInt(w.B != 0) + boolToInt(w.C != 0)
	g := gcdInts(w.A, w.B, w.C)
	balanced := abs(w.A) == abs(w.B) && w.A != 0 && w.B != 0
	connected := support >= 2 && ((w.A != 0 && w.B != 0) || (w.A != 0 && w.C != 0) || (w.B != 0 && w.C != 0))
	closed := w.A != 0 && w.B != 0 && w.C != 0
	open := connected && !closed
	c2 := w.B*w.B - w.A*w.A
	c3 := 2 * w.A * w.B * w.C
	nonzero := c2 != 0 || c3 != 0
	unbalancedLift := nonzero && !balanced
	balancedLift := balanced && nonzero
	canonical := balancedLift && closed && support == 3 && g == 1 && abs(w.A) == 1 && abs(w.B) == 1 && abs(w.C) == 1
	return Candidate{Weights: w, Matrix: [][]int{{0, w.A, w.C}, {w.A, 0, w.B}, {w.C, w.B, 0}}, SupportEdges: support, GCD: g, Primitive: g == 1, EndpointBalanced: balanced, Connected: connected, ClosedTriangle: closed, OpenChain: open, ZeroDiagonal: true, C2: c2, C3: c3, DeterminantFormula: determinantFormula(c2, c3), DeterminantNonZero: nonzero, UnbalancedLift: unbalancedLift, BalancedMassLift: balancedLift, CanonicalTopology: canonical, PassesAllBoundaries: canonical}
}

func determinantFormula(c2, c3 int) string {
	parts := []string{}
	if c2 != 0 {
		parts = append(parts, fmt.Sprintf("%d ε^2", c2))
	}
	if c3 != 0 {
		parts = append(parts, fmt.Sprintf("%d ε^3", c3))
	}
	if len(parts) == 0 {
		return "0"
	}
	return strings.Join(parts, " + ")
}

func collapse(s Sieve) DeterminantCollapse {
	forced := s.UniqueUnsignedTopology && len(s.BalancedLiftCandidates) > 0 && len(s.OpenChainFailures) > 0
	return DeterminantCollapse{Executed: true, KGen: "diag(-1,0,1)", Bridge: "B(a,b,c) with a=B12, b=B23, c=B13", DeterminantIdentity: "det(K+εB)=(b²-a²)ε²+2abc ε³", EndpointBalance: "|a|=|b|", BalancedReduction: "endpoint balance cancels the ε² term", OpenChainResult: "c=0 ⇒ det(K+εB)=0, so the structural zero survives", TriangleResult: "abc≠0 ⇒ det(K+εB)=2abc ε³, so the zero is lifted at cubic order", MiddleEigenvalueOrder: "λ₂^eff = O(ε³) for the balanced primitive bridge", ForcesClosedTriangle: forced, ForcesXGenSupport: forced, FixesAmplitude: false, FixesSignedOrientation: false, Verdict: StatusSeesawMassLiftCompatible, Reason: "the balanced mass lift is impossible on an open chain and possible on the primitive closed triangle; this selects support topology but not sign, phase, or amplitude"}
}

func buildAxiom(c DeterminantCollapse, s Sieve) BridgeAxiom {
	forced := c.ForcesClosedTriangle && c.ForcesXGenSupport && s.CanonicalUnsignedTopology.SupportEdges == 3
	verdict := StatusTriangleBridgeTopologyForced
	if forced {
		verdict = StatusXGenSupportSelectedAsTopology
	}
	return BridgeAxiom{Executed: true, Name: "Generation-2 structural-zero mass-lift bridge", TopologyName: "primitive closed triangular family bridge / X_gen support", CanonicalMatrix: [][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, SupportEdges: 3, Trace: 0, TraceSquare: 6, DeterminantLeadingOrder: "det(K+εB)=2ε³ for the positive canonical representative", Generation2DiagonalStillZero: true, LiftsGeneration2Zero: forced, GeometricallyForcedTopology: forced, AmplitudeEmpirical: true, SignedOrientationEmpirical: true, YukawaValuesPredicted: false, MuonCharmMassPredicted: false, Verdict: verdict, Reason: "Gate 445 promotes only the unsigned bridge support topology: the diagonal zero remains, the mass lift is a cubic mixing resonance, and the bridge coefficient remains sealed"}
}

func buildFirewall(ax BridgeAxiom) Firewall {
	return Firewall{Executed: true, NoObservedMuonMassImported: true, NoObservedCharmMassImported: true, NoObservedYukawaImported: true, NoCKMImported: true, NoPMNSImported: true, BridgeTopologyNativeConditional: ax.GeometricallyForcedTopology, BridgeAmplitudeSealed: ax.AmplitudeEmpirical, SignedPhaseSealed: ax.SignedOrientationEmpirical, PhysicalMassRequiresBridgeData: true, NativeFlavorDimBefore: NativeFlavorDim, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimStillFree: KXYCoeffDim, Verdict: StatusFlavorFirewallPartiallyRefined, Reason: "the bridge topology is selected by boundary intersection, but physical second-family masses require a coefficient scale and sector source data that this gate does not derive"}
}

func buildNext() NextStep {
	return NextStep{Gate: 446, Title: "Signed-Cycle / Complex Phase Orientation Sieve", Reason: "Gate 445 fixes the unsigned triangle support but leaves signed cycle orientation and complex phase data sealed.", PrimaryTask: "test whether real/complex orientation of the triangular bridge is forced by J, Gamma, CP, or eta-graded trace constraints, or whether Y_gen phase data remains quarantined"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate444KGenForced || !a.Inheritance.Gate444Generation2BareZero || !a.Inheritance.Gate444NoColliderData {
		return fmt.Errorf("Gate445 inheritance failed: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Executed || !a.Arena.ZeroDiagonal || !a.Arena.Hermitian || !a.Arena.NoYukawaImported {
		return fmt.Errorf("bridge arena failed: %s", FormatArena(a.Arena))
	}
	if len(a.Boundaries) != 4 {
		return fmt.Errorf("expected 4 boundaries, got %d", len(a.Boundaries))
	}
	for _, b := range a.Boundaries {
		if !b.Applied || !b.Passed {
			return fmt.Errorf("boundary failed: %s", FormatBoundary(b))
		}
	}
	if !a.Sieve.Executed || !a.Sieve.UniqueUnsignedTopology || len(a.Sieve.BalancedLiftCandidates) == 0 || len(a.Sieve.OpenChainFailures) == 0 || a.Sieve.SignedVariants != 8 {
		return fmt.Errorf("sieve failed: %s", FormatSieve(a.Sieve))
	}
	if !a.Collapse.ForcesClosedTriangle || !a.Collapse.ForcesXGenSupport || a.Collapse.FixesAmplitude || a.Collapse.FixesSignedOrientation {
		return fmt.Errorf("collapse failed: %s", FormatCollapse(a.Collapse))
	}
	if !a.Axiom.GeometricallyForcedTopology || !a.Axiom.Generation2DiagonalStillZero || !a.Axiom.LiftsGeneration2Zero || !a.Axiom.AmplitudeEmpirical || a.Axiom.MuonCharmMassPredicted {
		return fmt.Errorf("axiom/firewall mismatch: %s", FormatAxiom(a.Axiom))
	}
	if !a.Firewall.NoObservedMuonMassImported || !a.Firewall.BridgeAmplitudeSealed || !a.Firewall.PhysicalMassRequiresBridgeData || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimStillFree != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 445 proves a narrow compatibility theorem: once Gate 444 fixes K_gen=diag(-1,0,1) and the second bare diagonal level is zero, an endpoint-balanced off-diagonal bridge can lift that zero only if the family graph closes into the primitive triangle. Open chains keep det(K+εB)=0. Thus the unsigned X_gen support topology is forced as the minimal seesaw/mixing bridge, while the coefficient ε, signed cycle orientation, sector Yukawa amplitudes, and observed muon/charm masses remain sealed."
}

func sortCandidates(xs []Candidate) {
	sort.Slice(xs, func(i, j int) bool {
		a, b := xs[i], xs[j]
		if a.SupportEdges != b.SupportEdges {
			return a.SupportEdges < b.SupportEdges
		}
		if abs(a.Weights.A) != abs(b.Weights.A) {
			return abs(a.Weights.A) < abs(b.Weights.A)
		}
		if abs(a.Weights.B) != abs(b.Weights.B) {
			return abs(a.Weights.B) < abs(b.Weights.B)
		}
		if abs(a.Weights.C) != abs(b.Weights.C) {
			return abs(a.Weights.C) < abs(b.Weights.C)
		}
		if a.Weights.A != b.Weights.A {
			return a.Weights.A < b.Weights.A
		}
		if a.Weights.B != b.Weights.B {
			return a.Weights.B < b.Weights.B
		}
		return a.Weights.C < b.Weights.C
	})
}

func gcdInts(xs ...int) int {
	g := 0
	for _, x := range xs {
		g = gcd2(g, abs(x))
	}
	return g
}

func gcd2(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func boolToInt(x bool) int {
	if x {
		return 1
	}
	return 0
}
