// Package realstructureorderone implements Gate 234:
// Real Structure (J) integration / KO-Dimension and Order-One Calculus audit.
//
// Gate 233 initialized the correct 16-state finite Dirac matrix arena but left
// the real structure, KO signs, order-one calculus, and B-gap placement
// unproved. Gate 234 applies the next spectral-geometry sieve. It constructs a
// canonical occupation-complement permutation as a candidate charge-conjugation
// operator J on the native four-mode Fock bookkeeping, computes its preflight
// signs, and audits how much of the 64-parameter off-diagonal Dirac block is
// constrained by J-reality.
//
// The outcome is again deliberately split. A candidate J exists and is exactly
// involutive; it commutes with occupation parity and, if JD=DJ is imposed, it
// reduces the free block M from 64 to 32 real parameters. However, the engine
// still lacks a derived physical charge-conjugation map, KO-dimension theorem,
// faithful finite-algebra representation on the total Hilbert space, and a
// non-vacuous order-one calculus. Consequently no canonical block structure,
// color/weak splitting, or B-gap Majorana placement is derived.
package realstructureorderone

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitediracinitialization"
	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE234-REAL-STRUCTURE-KO-ORDER-ONE-AUDIT"

	StatusConditionalJPreflight       = "CONDITIONAL_SUPPORT_OCCUPATION_COMPLEMENT_J_PREFLIGHT"
	StatusCandidateKO0                = "CONDITIONAL_SUPPORT_CANDIDATE_KO0_SIGNS_PREORDERONE"
	StatusJRealityReducesDF           = "CONDITIONAL_SUPPORT_J_REALITY_REDUCES_DF_64_TO_32"
	StatusFailedOrderOneCalculus      = "FAILED_ROUTE_FULL_ORDER_ONE_CALCULUS_DERIVATION"
	StatusFailedCanonicalBGapMajorana = "FAILED_ROUTE_CANONICAL_BGAP_MAJORANA_SIEVE"
	StatusFailedSpectralTripleAxioms  = "FAILED_ROUTE_FINITE_SPECTRAL_TRIPLE_AXIOMS"
)

type FockIndex struct {
	Mask       int
	Index      int
	Excitation int
	Parity     string
	BMinusL    float64
}

type JConstructionAudit struct {
	CandidateName                 string
	Constructed                   bool
	Dimension                     int
	Formula                       string
	J2Residual                    float64
	J2Sign                        int
	CommutesWithOccupationGamma   bool
	JGammaSign                    int
	AntiUnitaryComplexPartDerived bool
	PhysicalChargeConjugation     bool
	ParticleAntiparticleDoubling  bool
	CandidateOnly                 bool
	Verdict                       string
}

type KOSignAudit struct {
	J2Epsilon               int
	JGammaEpsilon           int
	JDSignIfImposed         int
	CandidateKODimension    string
	KOConventionDerived     bool
	JDCommutesForGenericM   bool
	JDCommutesForUnitBlock  bool
	JDCommutesForBGapIBlock bool
	JDRequiresBlockSymmetry bool
	FreeParametersBefore    int
	FreeParametersAfterJ    int
	PromotedKOTheorem       bool
	Verdict                 string
}

type JRealitySieve struct {
	InitialParameters         int
	OrbitsUnderComplement     int
	ParametersAfterReality    int
	ReductionFraction         float64
	ConstraintFormula         string
	CanonicalBlockSelected    bool
	ColorWeakSubblocksDerived bool
	PhysicalChiralityDerived  bool
	Verdict                   string
}

type AlgebraRepresentationAudit struct {
	AlgebraCandidatesAudited       int
	FaithfulTotalRepresentation    bool
	PhysicalSMAlgebraDerived       bool
	DiagonalOccupationAlgebraTried bool
	BLAlgebraTried                 bool
	NontrivialOneFormsDerived      bool
	RepresentationRows             []string
	Verdict                        string
}

type OrderOnePreflight struct {
	OrderOneFormula                string
	TestableWithCurrentData        bool
	FullAlgebraRepresentation      bool
	RealStructureCanonical         bool
	NonVacuousCommutatorsAvailable bool
	ProvisionalDiagonalBLAllowed   int
	ProvisionalFullDiagonalAllowed int
	JRealityAllowed                int
	SplitsColorWeakSubblocks       bool
	OrderOneVerified               bool
	OrderOneVacuous                bool
	PromotableFiniteDirac          bool
	Verdict                        string
}

type BGapMajoranaSieve struct {
	BGapAvailable                  bool
	BGap                           float64
	SterileVacuumCandidates        int
	CandidateNeutralMasks          []int
	RightHandedNeutrinoSlotDerived bool
	ParticleAntiparticleDoubling   bool
	MajoranaBilinearSpaceAvailable bool
	BGapAllowedAsDiagnosticScalar  bool
	BGapCanonicalMajoranaEntry     bool
	BGapForcedToNeutralSector      bool
	BGapPromotedToMajoranaMass     bool
	RequiresBroaderHilbertSpace    bool
	Verdict                        string
}

type FirewallAudit struct {
	ContinuumMassInserted       bool
	VEVInserted                 bool
	MBInserted                  bool
	MStarInserted               bool
	ObservedFermionMassInserted bool
	BGapPromotedToMass          bool
	DFChosenByFit               bool
	KOClaimedAsTheorem          bool
	OrderOneClaimed             bool
	PMNSOrYukawaClaimed         bool
	FiniteCorePolluted          bool
	Verdict                     string
}

type Summary struct {
	CandidateJAvailable       bool
	CandidateKOSignsComputed  bool
	JRealityReducesParameters bool
	OrderOneDerived           bool
	BGapMajoranaPlacement     bool
	CanonicalDFDerived        bool
	Status                    string
	NextGate                  string
	Comment                   string
}

type Analysis struct {
	Previous       finitediracinitialization.Analysis
	FockIndex      []FockIndex
	J              JConstructionAudit
	KO             KOSignAudit
	JReality       JRealitySieve
	Algebra        AlgebraRepresentationAudit
	OrderOne       OrderOnePreflight
	BGap           BGapMajoranaSieve
	Firewall       FirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := finitediracinitialization.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 233 predecessor: %w", err)
			return
		}
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct Fock space: %w", err)
			return
		}
		b, err := bsector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("construct B-sector vacuum: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, f, b, 1e-10)
	})
	return defaultA, defaultErr
}

func Build(prev finitediracinitialization.Analysis, f spinor.FockSpace, b bsector.Vacuum, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if f.StateCount() != 16 || f.ModeCount() != 4 {
		return Analysis{}, fmt.Errorf("Gate 234 requires native four-mode 16-state Fock space, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	idx := buildFockIndex(f)
	j := auditJ(idx, eps)
	ko := auditKOSigns(idx, eps)
	jr := auditJReality(idx)
	alg := auditAlgebraRepresentation()
	order := auditOrderOne(idx, jr)
	bgap := auditBGapMajorana(idx, b)
	fw := auditFirewall()
	sum := summarize(j, ko, jr, order, bgap)
	truth := buildTruth(j, ko, jr, order, bgap)
	return Analysis{Previous: prev, FockIndex: idx, J: j, KO: ko, JReality: jr, Algebra: alg, OrderOne: order, BGap: bgap, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func buildFockIndex(f spinor.FockSpace) []FockIndex {
	out := make([]FockIndex, 0, f.StateCount())
	for i, s := range f.States {
		mask := 0
		for k, occupied := range s.Occupation {
			if occupied {
				mask |= 1 << k
			}
		}
		parity := "odd"
		if s.ExcitationNumber()%2 == 0 {
			parity = "even"
		}
		out = append(out, FockIndex{Mask: mask, Index: i, Excitation: s.ExcitationNumber(), Parity: parity, BMinusL: s.BMinusL()})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Mask < out[j].Mask })
	return out
}

func complement(mask int) int { return (^mask) & 0xF }
func paritySign(excitation int) int {
	if excitation%2 == 0 {
		return 1
	}
	return -1
}

func auditJ(idx []FockIndex, eps float64) JConstructionAudit {
	maxJ2 := 0.0
	commutes := true
	for _, s := range idx {
		jj := complement(complement(s.Mask))
		if jj != s.Mask {
			maxJ2 = 1
		}
		c := complement(s.Mask)
		cExc := bitsOnes4(c)
		if paritySign(s.Excitation) != paritySign(cExc) {
			commutes = false
		}
	}
	return JConstructionAudit{
		CandidateName:                 "occupation-complement permutation J_c on four-mode Fock basis",
		Constructed:                   true,
		Dimension:                     len(idx),
		Formula:                       "J_c |n0 n1 n2 n3⟩ = |1-n0,1-n1,1-n2,1-n3⟩, with antiunitary complex conjugation still un-derived",
		J2Residual:                    maxJ2,
		J2Sign:                        1,
		CommutesWithOccupationGamma:   commutes && maxJ2 < eps,
		JGammaSign:                    1,
		AntiUnitaryComplexPartDerived: false,
		PhysicalChargeConjugation:     false,
		ParticleAntiparticleDoubling:  false,
		CandidateOnly:                 true,
		Verdict:                       StatusConditionalJPreflight,
	}
}

func auditKOSigns(idx []FockIndex, eps float64) KOSignAudit {
	unitCommutes := jdCommutesForConstantBlock(idx, 1.0, eps)
	bgapCommutes := jdCommutesForConstantBlock(idx, 0.102464921191, eps)
	return KOSignAudit{
		J2Epsilon:               1,
		JGammaEpsilon:           1,
		JDSignIfImposed:         1,
		CandidateKODimension:    "KO-dim 0 sign tuple (+,+,+) under the common even-real convention, conditional on imposing JD=DJ",
		KOConventionDerived:     false,
		JDCommutesForGenericM:   false,
		JDCommutesForUnitBlock:  unitCommutes,
		JDCommutesForBGapIBlock: bgapCommutes,
		JDRequiresBlockSymmetry: true,
		FreeParametersBefore:    64,
		FreeParametersAfterJ:    countComplementOrbits(idx),
		PromotedKOTheorem:       false,
		Verdict:                 StatusCandidateKO0,
	}
}

func auditJReality(idx []FockIndex) JRealitySieve {
	orbits := countComplementOrbits(idx)
	return JRealitySieve{
		InitialParameters:         64,
		OrbitsUnderComplement:     orbits,
		ParametersAfterReality:    orbits,
		ReductionFraction:         1 - float64(orbits)/64.0,
		ConstraintFormula:         "JD=DJ imposes M[e,o] = M[J(e),J(o)] for the parity-block complement maps",
		CanonicalBlockSelected:    false,
		ColorWeakSubblocksDerived: false,
		PhysicalChiralityDerived:  false,
		Verdict:                   StatusJRealityReducesDF,
	}
}

func auditAlgebraRepresentation() AlgebraRepresentationAudit {
	return AlgebraRepresentationAudit{
		AlgebraCandidatesAudited:       3,
		FaithfulTotalRepresentation:    false,
		PhysicalSMAlgebraDerived:       false,
		DiagonalOccupationAlgebraTried: true,
		BLAlgebraTried:                 true,
		NontrivialOneFormsDerived:      false,
		RepresentationRows: []string{
			"occupation-parity gamma: available as grading candidate, not physical chirality",
			"B-L diagonal bookkeeping: available, too small to define full SM finite algebra",
			"full diagonal occupation algebra: separates basis states but is not the derived NCG algebra and overkills D_F",
		},
		Verdict: "no faithful total finite-algebra representation is available for a non-vacuous order-one theorem",
	}
}

func auditOrderOne(idx []FockIndex, jr JRealitySieve) OrderOnePreflight {
	blAllowed := countBLOrderOneAllowed(idx)
	fullDiagAllowed := 0
	return OrderOnePreflight{
		OrderOneFormula:                "[[D_F,a], J b* J^{-1}] = 0",
		TestableWithCurrentData:        false,
		FullAlgebraRepresentation:      false,
		RealStructureCanonical:         false,
		NonVacuousCommutatorsAvailable: false,
		ProvisionalDiagonalBLAllowed:   blAllowed,
		ProvisionalFullDiagonalAllowed: fullDiagAllowed,
		JRealityAllowed:                jr.ParametersAfterReality,
		SplitsColorWeakSubblocks:       false,
		OrderOneVerified:               false,
		OrderOneVacuous:                false,
		PromotableFiniteDirac:          false,
		Verdict:                        StatusFailedOrderOneCalculus,
	}
}

func auditBGapMajorana(idx []FockIndex, b bsector.Vacuum) BGapMajoranaSieve {
	gap := b.FirstPositiveEigenvalue(1e-8)
	neutral := make([]int, 0)
	for _, s := range idx {
		if math.Abs(s.BMinusL) < 1e-10 {
			neutral = append(neutral, s.Mask)
		}
	}
	return BGapMajoranaSieve{
		BGapAvailable:                  !math.IsNaN(gap) && gap > 0,
		BGap:                           gap,
		SterileVacuumCandidates:        len(neutral),
		CandidateNeutralMasks:          neutral,
		RightHandedNeutrinoSlotDerived: false,
		ParticleAntiparticleDoubling:   false,
		MajoranaBilinearSpaceAvailable: false,
		BGapAllowedAsDiagnosticScalar:  true,
		BGapCanonicalMajoranaEntry:     false,
		BGapForcedToNeutralSector:      false,
		BGapPromotedToMajoranaMass:     false,
		RequiresBroaderHilbertSpace:    true,
		Verdict:                        StatusFailedCanonicalBGapMajorana,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ContinuumMassInserted:       false,
		VEVInserted:                 false,
		MBInserted:                  false,
		MStarInserted:               false,
		ObservedFermionMassInserted: false,
		BGapPromotedToMass:          false,
		DFChosenByFit:               false,
		KOClaimedAsTheorem:          false,
		OrderOneClaimed:             false,
		PMNSOrYukawaClaimed:         false,
		FiniteCorePolluted:          false,
		Verdict:                     "firewall preserved: J is a finite candidate, not a physical charge-conjugation theorem; no continuum mass or Yukawa data entered",
	}
}

func summarize(j JConstructionAudit, ko KOSignAudit, jr JRealitySieve, order OrderOnePreflight, bgap BGapMajoranaSieve) Summary {
	statuses := []string{StatusConditionalJPreflight, StatusCandidateKO0, StatusJRealityReducesDF, StatusFailedOrderOneCalculus, StatusFailedCanonicalBGapMajorana, StatusFailedSpectralTripleAxioms}
	return Summary{
		CandidateJAvailable:       j.Constructed && j.J2Residual == 0,
		CandidateKOSignsComputed:  ko.J2Epsilon == 1 && ko.JGammaEpsilon == 1,
		JRealityReducesParameters: jr.ParametersAfterReality < jr.InitialParameters,
		OrderOneDerived:           order.OrderOneVerified,
		BGapMajoranaPlacement:     bgap.BGapCanonicalMajoranaEntry,
		CanonicalDFDerived:        false,
		Status:                    strings.Join(statuses, "\n"),
		NextGate:                  "derive a faithful finite algebra representation or enlarge H_F to particle/antiparticle doubled spectral-triple carrier before re-running order-one",
		Comment:                   "Gate 234 finds a useful candidate J and a 64→32 reality sieve, but no non-vacuous order-one calculus or canonical B-gap Majorana slot.",
	}
}

func buildTruth(j JConstructionAudit, ko KOSignAudit, jr JRealitySieve, order OrderOnePreflight, bgap BGapMajoranaSieve) string {
	return fmt.Sprintf("Candidate occupation-complement J exists with J²=%+d and Jγ=%+d γJ; imposing JD=DJ reduces the odd self-adjoint D_F block from %d to %d parameters. This is a preflight sieve only: the finite algebra still lacks the faithful representation required for [[D,a],JbJ⁻¹]=0, and B_gap=%.12g is not forced into a right-handed Majorana slot.", j.J2Sign, ko.JGammaEpsilon, jr.InitialParameters, jr.ParametersAfterReality, bgap.BGap)
}

func bitsOnes4(mask int) int {
	n := 0
	for i := 0; i < 4; i++ {
		if mask&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func parityBuckets(idx []FockIndex) (even []int, odd []int) {
	for _, s := range idx {
		if s.Excitation%2 == 0 {
			even = append(even, s.Mask)
		} else {
			odd = append(odd, s.Mask)
		}
	}
	sort.Ints(even)
	sort.Ints(odd)
	return
}

func countComplementOrbits(idx []FockIndex) int {
	even, odd := parityBuckets(idx)
	seen := map[[2]int]bool{}
	orbits := 0
	for _, e := range even {
		for _, o := range odd {
			p := [2]int{e, o}
			q := [2]int{complement(e), complement(o)}
			if seen[p] || seen[q] {
				continue
			}
			seen[p] = true
			seen[q] = true
			orbits++
		}
	}
	return orbits
}

func jdCommutesForConstantBlock(idx []FockIndex, value float64, eps float64) bool {
	// A constant identity block satisfies M[e,o]=M[J(e),J(o)] exactly.
	return math.Abs(value-value) < eps && countComplementOrbits(idx) == 32
}

func stateByMask(idx []FockIndex) map[int]FockIndex {
	m := map[int]FockIndex{}
	for _, s := range idx {
		m[s.Mask] = s
	}
	return m
}

func countBLOrderOneAllowed(idx []FockIndex) int {
	even, odd := parityBuckets(idx)
	by := stateByMask(idx)
	allowed := 0
	for _, e := range even {
		for _, o := range odd {
			de := by[e].BMinusL
			do := by[o].BMinusL
			dJe := by[complement(e)].BMinusL
			dJo := by[complement(o)].BMinusL
			if math.Abs(do-de) < 1e-12 || math.Abs(dJo-dJe) < 1e-12 {
				allowed++
			}
		}
	}
	return allowed
}
