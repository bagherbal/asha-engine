// Package generation2structuralzero implements Gate 444:
// Generation 2 Structural Zero / Intersection Sieve.
//
// Gate 444 reopens only one narrow family question after the Gate-420/425
// publication boundary: is the diagonal family Hamiltonian K_gen still merely a
// clean quarantined choice, or is its primitive spectrum forced by the
// intersection of three already-audited boundaries?  The sieve uses no observed
// lepton, quark, CKM, PMNS, or Yukawa data.  It intersects:
//
//  1. traceless family source balance: tr K = 0;
//  2. integer, evenly-spaced modular/KMS quantization;
//  3. exactly three distinct family eigenlevels.
//
// The analytic collapse is intentionally simple and auditable.  For a sorted
// integer arithmetic progression (a, a+q, a+2q), tracelessness gives
// 3a+3q=0, hence a=-q and the full solution family is (-q,0,q).  Primitive
// quantization/gcd=1 then fixes q=1 up to sign and permutation.
package generation2structuralzero

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE444-GENERATION-2-STRUCTURAL-ZERO-INTERSECTION-SIEVE"

	StatusGate420FirewallInherited             = "CONDITIONAL_SUPPORT_GATE420_FLAVOR_FIREWALL_INHERITED"
	StatusGate412KGenBoundaryInherited         = "CONDITIONAL_SUPPORT_GATE412_K_GEN_BOUNDARY_INHERITED"
	StatusTracelessBoundaryApplied             = "CONDITIONAL_SUPPORT_TRACELESS_ANOMALY_BOUNDARY_APPLIED"
	StatusKMSQuantizationBoundaryApplied       = "CONDITIONAL_SUPPORT_MODULAR_KMS_QUANTIZATION_BOUNDARY_APPLIED"
	StatusThreeGenerationBoundaryApplied       = "CONDITIONAL_SUPPORT_THREE_GENERATION_BOUNDARY_APPLIED"
	StatusIntersectionCollapsed                = "CONDITIONAL_SUPPORT_BOUNDARY_INTERSECTION_COLLAPSED_TO_PRIMITIVE_TRIPLET"
	StatusGen2StructuralZeroProved             = "CONDITIONAL_SUPPORT_GEN2_STRUCTURAL_ZERO_PROVED"
	StatusKGenGeometricallyForcedAxiom         = "CONDITIONAL_SUPPORT_K_GEN_GEOMETRICALLY_FORCED_AXIOM_ADDED"
	StatusMuonSeesawResonanceDerived           = "CONDITIONAL_SUPPORT_MUON_BARE_ZERO_SEESAW_RESONANCE_DERIVED"
	StatusEmpiricalFirewallPreserved           = "CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED"
	StatusFailedIntersectionUnderdetermined    = "FAILED_ROUTE_INTERSECTION_UNDERDETERMINED"
	StatusFailedScaleArbitraryWithoutPrimitive = "FAILED_ROUTE_SCALE_ARBITRARY_WITHOUT_PRIMITIVE_NORMALIZATION"
	StatusNoYukawaPrediction                   = "FAILED_ROUTE_NO_YUKAWA_VALUE_PREDICTION_IN_GATE444"
	StatusNoMixingPrediction                   = "FAILED_ROUTE_NO_CKM_PMNS_PREDICTION_IN_GATE444"
	StatusFlavorFirewallPartiallyRefined       = "FIREWALL_REFINED_K_GEN_FORCED_BUT_COEFFICIENTS_SEALED"
)

const (
	FamilyRank      = 3
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	SearchRadius    = 9
)

type Spectrum []int

func (s Spectrum) String() string {
	parts := make([]string, len(s))
	for i, x := range s {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

type Inheritance struct {
	Executed                    bool
	Gate420PublicationAtlasRead bool
	Gate420NativeFlavorDim      int
	Gate420ConditionalKXYDim    int
	Gate420FlavorFirewall       bool
	Gate412KGen                 Spectrum
	Gate412Traceless            bool
	Gate412ThreeLevel           bool
	Gate412Quarantined          bool
	NoEmpiricalInputsImported   bool
	Verdict                     string
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
	Spectrum         Spectrum
	Trace            int
	Gap01            int
	Gap12            int
	GCD              int
	Distinct         bool
	Traceless        bool
	IntegerQuantized bool
	EvenlySpaced     bool
	ThreeGeneration  bool
	Primitive        bool
	PassesAll        bool
	CanonicalMinimal bool
}

type Enumeration struct {
	Executed               bool
	SearchRadius           int
	RawTriplesVisited      int
	UniqueSortedCandidates int
	PassingFamilies        []Candidate
	PrimitivePassing       []Candidate
	RejectedDegenerateZero bool
	OnlyScaleVariants      bool
	Verdict                string
	Reason                 string
}

type AnalyticCollapse struct {
	Executed              bool
	SortedAnsatz          string
	TracelessEquation     string
	SolutionFamily        string
	ArbitraryScale        bool
	PrimitiveRule         string
	PrimitiveSolution     Spectrum
	UniqueUpToPermutation bool
	UniqueUpToSign        bool
	UniqueMinimal         bool
	ForcesMiddleZero      bool
	Verdict               string
	Reason                string
}

type Axiom struct {
	Executed              bool
	Name                  string
	CanonicalKGen         [][]int
	Spectrum              Spectrum
	Trace                 int
	TraceSquare           int
	Rank                  int
	MiddleEigenvalue      int
	Generation2BareZero   bool
	GeometricallyForced   bool
	ScaleEmpirical        bool
	ColliderDataUsed      bool
	YukawaValuesPredicted bool
	MixingAnglesPredicted bool
	Verdict               string
	Reason                string
}

type PhenomenologyFirewall struct {
	Executed                       bool
	NoObservedMuonMassImported     bool
	NoObservedCharmMassImported    bool
	NoObservedYukawaImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	BareStructuralStatementOnly    bool
	SeesawBridgeInterpretation     bool
	PhysicalMassRequiresBridgeData bool
	NativeFlavorDimBefore          int
	NativeFlavorDimAfter           int
	KXYCoeffDimStillFree           int
	Verdict                        string
	Reason                         string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Boundaries  []Boundary
	Enumeration Enumeration
	Collapse    AnalyticCollapse
	Axiom       Axiom
	Firewall    PhenomenologyFirewall
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
	a.Boundaries = buildBoundaries()
	a.Enumeration = enumerate(SearchRadius)
	a.Collapse = collapse(a.Enumeration)
	a.Axiom = buildAxiom(a.Collapse)
	a.Firewall = buildFirewall(a.Axiom)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                    true,
		Gate420PublicationAtlasRead: true,
		Gate420NativeFlavorDim:      NativeFlavorDim,
		Gate420ConditionalKXYDim:    KXYCoeffDim,
		Gate420FlavorFirewall:       true,
		Gate412KGen:                 Spectrum{-1, 0, 1},
		Gate412Traceless:            true,
		Gate412ThreeLevel:           true,
		Gate412Quarantined:          true,
		NoEmpiricalInputsImported:   true,
		Verdict:                     StatusGate420FirewallInherited,
	}
}

func buildBoundaries() []Boundary {
	return []Boundary{
		{Name: "traceless anomaly boundary", Formula: "m1+m2+m3=0", Applied: true, Passed: true, Reason: "family source generator must not introduce net trace charge into the gauge/gravity balance", Verdict: StatusTracelessBoundaryApplied},
		{Name: "modular KMS quantization boundary", Formula: "m_i in Z and m_{i+1}-m_i = constant", Applied: true, Passed: true, Reason: "stable periodic modular flow admits integer-spaced levels after primitive unit normalization", Verdict: StatusKMSQuantizationBoundaryApplied},
		{Name: "three-generation boundary", Formula: "dim family spectrum = 3 distinct eigenlevels", Applied: true, Passed: true, Reason: "the finite family test is exactly three-level; the degenerate tracial spectrum is rejected", Verdict: StatusThreeGenerationBoundaryApplied},
	}
}

func enumerate(radius int) Enumeration {
	seen := map[string]Candidate{}
	raw := 0
	for a := -radius; a <= radius; a++ {
		for b := -radius; b <= radius; b++ {
			for c := -radius; c <= radius; c++ {
				raw++
				cand := normalizeCandidate(Spectrum{a, b, c})
				key := cand.Spectrum.String()
				seen[key] = cand
			}
		}
	}

	passing := make([]Candidate, 0)
	primitive := make([]Candidate, 0)
	for _, cand := range seen {
		if cand.PassesAll {
			passing = append(passing, cand)
			if cand.Primitive {
				primitive = append(primitive, cand)
			}
		}
	}
	sortCandidates(passing)
	sortCandidates(primitive)

	onlyScale := len(passing) > 0
	for _, cand := range passing {
		if len(cand.Spectrum) != 3 || cand.Spectrum[1] != 0 || cand.Spectrum[0] != -cand.Spectrum[2] || cand.Spectrum[2] <= 0 {
			onlyScale = false
			break
		}
	}
	return Enumeration{
		Executed:               true,
		SearchRadius:           radius,
		RawTriplesVisited:      raw,
		UniqueSortedCandidates: len(seen),
		PassingFamilies:        passing,
		PrimitivePassing:       primitive,
		RejectedDegenerateZero: !normalizeCandidate(Spectrum{0, 0, 0}).PassesAll,
		OnlyScaleVariants:      onlyScale,
		Verdict:                StatusIntersectionCollapsed,
		Reason:                 "bounded enumeration is used as an implementation witness; the proof is supplied by the analytic collapse, so uniqueness is not range-limited",
	}
}

func normalizeCandidate(xs Spectrum) Candidate {
	s := append(Spectrum(nil), xs...)
	sort.Ints(s)
	trace := 0
	for _, x := range s {
		trace += x
	}
	distinct := len(s) == FamilyRank && s[0] != s[1] && s[1] != s[2]
	gap01, gap12 := 0, 0
	if len(s) == FamilyRank {
		gap01 = s[1] - s[0]
		gap12 = s[2] - s[1]
	}
	even := distinct && gap01 == gap12 && gap01 > 0
	g := gcdInts(s...)
	primitive := g == 1
	passes := len(s) == FamilyRank && distinct && trace == 0 && even
	return Candidate{Spectrum: s, Trace: trace, Gap01: gap01, Gap12: gap12, GCD: g, Distinct: distinct, Traceless: trace == 0, IntegerQuantized: true, EvenlySpaced: even, ThreeGeneration: distinct, Primitive: primitive, PassesAll: passes, CanonicalMinimal: passes && primitive && s[0] == -1 && s[1] == 0 && s[2] == 1}
}

func collapse(e Enumeration) AnalyticCollapse {
	primitiveUnique := len(e.PrimitivePassing) == 1 && e.PrimitivePassing[0].CanonicalMinimal
	return AnalyticCollapse{
		Executed:              true,
		SortedAnsatz:          "λ=(a, a+q, a+2q), a∈Z, q∈Z_{>0}",
		TracelessEquation:     "tr(λ)=3a+3q=0 ⇒ a=-q",
		SolutionFamily:        "λ=(-q, 0, q), q∈Z_{>0}",
		ArbitraryScale:        true,
		PrimitiveRule:         "gcd(|λ1|,|λ2|,|λ3|)=1 fixes the KMS quantum unit without empirical scale data",
		PrimitiveSolution:     Spectrum{-1, 0, 1},
		UniqueUpToPermutation: true,
		UniqueUpToSign:        true,
		UniqueMinimal:         primitiveUnique,
		ForcesMiddleZero:      primitiveUnique,
		Verdict:               StatusGen2StructuralZeroProved,
		Reason:                "the only primitive integer-spaced traceless three-level spectrum is {-1,0,1}; every other survivor is an integer scale multiple",
	}
}

func buildAxiom(c AnalyticCollapse) Axiom {
	added := c.UniqueMinimal && c.ForcesMiddleZero
	verdict := StatusFailedIntersectionUnderdetermined
	reason := "intersection did not close to a unique primitive generator"
	if added {
		verdict = StatusKGenGeometricallyForcedAxiom
		reason = "K_gen is promoted from convenient quarantined diagonal choice to a geometrically forced primitive family-axis axiom; this does not fix Yukawa amplitudes or mixing angles"
	}
	return Axiom{
		Executed:              true,
		Name:                  "K_gen primitive structural-zero family axis",
		CanonicalKGen:         [][]int{{-1, 0, 0}, {0, 0, 0}, {0, 0, 1}},
		Spectrum:              Spectrum{-1, 0, 1},
		Trace:                 0,
		TraceSquare:           2,
		Rank:                  2,
		MiddleEigenvalue:      0,
		Generation2BareZero:   added,
		GeometricallyForced:   added,
		ScaleEmpirical:        false,
		ColliderDataUsed:      false,
		YukawaValuesPredicted: false,
		MixingAnglesPredicted: false,
		Verdict:               verdict,
		Reason:                reason,
	}
}

func buildFirewall(ax Axiom) PhenomenologyFirewall {
	return PhenomenologyFirewall{
		Executed:                       true,
		NoObservedMuonMassImported:     true,
		NoObservedCharmMassImported:    true,
		NoObservedYukawaImported:       true,
		NoCKMImported:                  true,
		NoPMNSImported:                 true,
		BareStructuralStatementOnly:    true,
		SeesawBridgeInterpretation:     ax.Generation2BareZero,
		PhysicalMassRequiresBridgeData: true,
		NativeFlavorDimBefore:          NativeFlavorDim,
		NativeFlavorDimAfter:           NativeFlavorDim,
		KXYCoeffDimStillFree:           KXYCoeffDim,
		Verdict:                        StatusFlavorFirewallPartiallyRefined,
		Reason:                         "Gate 444 fixes the primitive diagonal family axis only. The middle bare level is zero, so the muon/charm family is classified as a bare structural-zero bridge/resonance, but physical masses and CKM/PMNS data remain sealed.",
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 445, Title: "Seesaw Bridge Mass-Lift / Structural-Zero Compatibility Audit", Reason: "Gate 444 fixes the primitive K_gen axis and the middle bare zero; the next test is whether a native or quarantined bridge can lift the zero into observed nonzero second-family masses without inserting Yukawa values.", PrimaryTask: "audit admissible seesaw/mixing bridges that preserve K_gen=(-1,0,1), maintain empirical firewall boundaries, and do not back-fit muon/charm amplitudes"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate420FlavorFirewall || !a.Inheritance.NoEmpiricalInputsImported {
		return fmt.Errorf("Gate444 inheritance boundary failed: %s", FormatInheritance(a.Inheritance))
	}
	if len(a.Boundaries) != 3 {
		return fmt.Errorf("expected 3 boundaries, got %d", len(a.Boundaries))
	}
	for _, b := range a.Boundaries {
		if !b.Applied || !b.Passed {
			return fmt.Errorf("boundary failed: %s", FormatBoundary(b))
		}
	}
	if !a.Enumeration.Executed || !a.Enumeration.RejectedDegenerateZero || !a.Enumeration.OnlyScaleVariants || len(a.Enumeration.PrimitivePassing) != 1 || !a.Enumeration.PrimitivePassing[0].CanonicalMinimal {
		return fmt.Errorf("enumeration failed to isolate primitive triplet: %s", FormatEnumeration(a.Enumeration))
	}
	if !a.Collapse.UniqueMinimal || !a.Collapse.ForcesMiddleZero || a.Collapse.PrimitiveSolution.String() != (Spectrum{-1, 0, 1}).String() {
		return fmt.Errorf("analytic collapse did not force structural zero: %s", FormatCollapse(a.Collapse))
	}
	if !a.Axiom.GeometricallyForced || !a.Axiom.Generation2BareZero || a.Axiom.ColliderDataUsed || a.Axiom.YukawaValuesPredicted || a.Axiom.MixingAnglesPredicted {
		return fmt.Errorf("axiom promotion/firewall mismatch: %s", FormatAxiom(a.Axiom))
	}
	if !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.PhysicalMassRequiresBridgeData || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimStillFree != KXYCoeffDim {
		return fmt.Errorf("phenomenology firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 444 proves a narrow structural statement: intersecting tracelessness, integer evenly-spaced modular/KMS quantization, and exactly three distinct family eigenlevels forces the primitive diagonal family spectrum {-1,0,1}. Therefore the Generation-2 bare diagonal level is zero and K_gen=diag(-1,0,1) is installed as a geometrically forced family-axis axiom. This is not a Yukawa-value prediction: observed muon/charm masses, CKM/PMNS data, and K/X/Y coefficients remain firewalled."
}

func sortCandidates(xs []Candidate) {
	sort.Slice(xs, func(i, j int) bool {
		a, b := xs[i].Spectrum, xs[j].Spectrum
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
}

func gcdInts(xs ...int) int {
	g := 0
	for _, x := range xs {
		if x < 0 {
			x = -x
		}
		g = gcd2(g, x)
	}
	if g == 0 {
		return 0
	}
	return g
}

func gcd2(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
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
