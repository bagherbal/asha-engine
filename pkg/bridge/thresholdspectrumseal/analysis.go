// Package thresholdspectrumseal implements Gate 213: ThresholdSpectrumSeal /
// matching-correction and two-loop stability preflight audit.
//
// Gate 212 proved that Gate 211's 44 ordered viable two-threshold witnesses
// reduce to 22 unordered physical pair classes and that no finite-origin or
// parentage theorem selects a unique spectrum. Gate 213 therefore does not
// promote the best ranked pair into a finite prediction. It introduces an
// explicit ThresholdSpectrumSeal, chooses the Gate-211 ranked witness only as a
// conditional test subject, audits whether finite geometry supplies threshold
// matching corrections, and performs a standard-QFT two-loop preflight without
// importing the subtraction scheme as a finite theorem.
package thresholdspectrumseal

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/representationrowlattice"
	"github.com/bagherbal/asha-engine/pkg/bridge/twothresholdminimality"
	"github.com/bagherbal/asha-engine/pkg/bridge/twothresholdviability"
)

const (
	StatusConditionalSealPreflight = "CONDITIONAL_PHENOMENOLOGY_ON_THRESHOLD_SPECTRUM_SEAL"
	StatusFailedSealPreflight      = "FAILED_ROUTE_THRESHOLD_SPECTRUM_PREFLIGHT"

	MatchingCorrectionsFailed = "FAILED_ROUTE_DERIVED_MATCHING_CORRECTIONS"
	TwoLoopWarning            = "TWO_LOOP_PREFLIGHT_WARNING_ONE_LOOP_STABILITY_NOT_PROVEN"
	TwoLoopStable             = "TWO_LOOP_PREFLIGHT_STABLE_SMALL_CORRECTIONS"
	TwoLoopCatastrophic       = "TWO_LOOP_PREFLIGHT_CATASTROPHIC"
)

type Rational = representationrowlattice.Rational

type RationalMatrix3 struct {
	M [3][3]Rational
}

func (m RationalMatrix3) Add(o RationalMatrix3) RationalMatrix3 {
	var out RationalMatrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out.M[i][j] = add(m.M[i][j], o.M[i][j])
		}
	}
	return out
}

func (m RationalMatrix3) FloatAt(i, j int) float64 { return m.M[i][j].Float() }

func (m RationalMatrix3) String() string {
	rows := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		rows = append(rows, fmt.Sprintf("[%s,%s,%s]", m.M[i][0], m.M[i][1], m.M[i][2]))
	}
	return "[" + strings.Join(rows, ";") + "]"
}

type Gate212Snapshot struct {
	Gate212Inherited                  bool
	CanonicalUniquenessFailedRoute    bool
	Gate211ConditionalViability       bool
	OrderedViablePairs                int
	UnorderedPairClasses              int
	ThresholdSpectrumSealRequired     bool
	LeptoquarkDynamicsSealInherited   bool
	EmpiricalCarrierSealInherited     bool
	FiniteCarrierOriginDerived        bool
	FiniteMatchingCorrectionsDerived  bool
	UniquePhysicalSpectrumClaimed     bool
	RecommendedNextGateMatchesGate213 bool
	TruthStatement                    string
}

func SnapshotFromGate212(a twothresholdminimality.Analysis) Gate212Snapshot {
	return Gate212Snapshot{
		Gate212Inherited:                  a.Summary.Gate211ViabilityInherited,
		CanonicalUniquenessFailedRoute:    a.Summary.Status == twothresholdminimality.StatusFailedCanonicalUniqueness && !a.Summary.CanonicalUniquePairFound,
		Gate211ConditionalViability:       a.Gate211.ConditionalViabilityInherited,
		OrderedViablePairs:                a.Summary.OrderedViablePairs,
		UnorderedPairClasses:              a.Summary.UnorderedPairClasses,
		ThresholdSpectrumSealRequired:     a.Summary.ThresholdSpectrumSealRequired,
		LeptoquarkDynamicsSealInherited:   a.Firewall.LeptoquarkDynamicsSealInherited,
		EmpiricalCarrierSealInherited:     a.Firewall.EmpiricalCarrierSealInherited,
		FiniteCarrierOriginDerived:        a.Gate211.FiniteCarrierOriginDerived,
		FiniteMatchingCorrectionsDerived:  a.Firewall.MatchingCorrectionsDerived,
		UniquePhysicalSpectrumClaimed:     a.Firewall.UniqueThresholdSpectrumClaimed,
		RecommendedNextGateMatchesGate213: strings.Contains(a.Firewall.RecommendedNextGate, "Gate 213"),
		TruthStatement:                    a.TruthStatement,
	}
}

type ThresholdSpectrumSeal struct {
	Name                         string
	SealID                       string
	IntroducedByGate             int
	Reason                       string
	DegeneracyQuarantined        bool
	SelectedPairPhenomenological bool
	FiniteUniquenessClaimed      bool
	ContactOrBGapOriginClaimed   bool
	MatchingSchemeClaimed        bool
	TwoLoopImportedAsFiniteCore  bool
	OneLoopScalesRemainReference bool
	Verdict                      string
}

type SealedSpectrumSubject struct {
	TargetName                string
	Row1Name                  string
	Row2Name                  string
	Row1Rep                   string
	Row2Rep                   string
	Row1DeltaB                representationrowlattice.RationalTriple
	Row2DeltaB                representationrowlattice.RationalTriple
	TotalDeltaB               representationrowlattice.RationalTriple
	TotalOneLoopBeta          twothresholdviability.FloatTriple
	LB1                       float64
	LB2                       float64
	LStar                     float64
	MB1GeV                    float64
	MB2GeV                    float64
	MStarGeV                  float64
	DeltaL                    float64
	AlphaGUTInverse           float64
	SelectedFromOrderedRank   int
	SelectedFromUnorderedRank int
	FiniteDerived             bool
	ConditionalOnly           bool
	Verdict                   string
}

type MatchingCorrectionAudit struct {
	TauEtaFundamentalClassAvailable      bool
	ScalarTraceSupportAvailable          bool
	ContactZetaTraceAvailable            bool
	HeatKernelMatchingDerived            bool
	SpectralTripleComplete               bool
	CanonicalSubtractionSchemeDerived    bool
	MSbarOrDimRegImported                bool
	FiniteCountertermFunctionalDerived   bool
	ThresholdMatchingCoefficientsDerived bool
	MatchingCorrectionRows               int
	Status                               string
	Verdict                              string
}

type GroupInvariants struct {
	Symbol string
	Dim    int64
	T      Rational
	C2     Rational
	CG     Rational
}

type ParsedRepresentation struct {
	Name       string
	Statistics string
	SU3        GroupInvariants
	SU2        GroupInvariants
	Y          Rational
	YCasimir   Rational
	Supported  bool
	Verdict    string
}

type TwoLoopCarrierMatrix struct {
	CarrierName string
	Rep         string
	Parsed      ParsedRepresentation
	OneLoopRow  representationrowlattice.RationalTriple
	Matrix      RationalMatrix3
	Formula     string
	Exact       bool
	FiniteCore  bool
	Verdict     string
}

type TwoLoopCoefficientAudit struct {
	Convention                     string
	SMOneLoopBeta                  twothresholdviability.FloatTriple
	SMTwoLoopMatrixNoYukawa        RationalMatrix3
	CarrierMatrices                []TwoLoopCarrierMatrix
	HeavyInducedTwoLoopMatrix      RationalMatrix3
	TotalTwoLoopMatrixNoYukawa     RationalMatrix3
	ExactSymbolicHeavyCoefficients bool
	UsesStandardQFTFormula         bool
	ImportedAsFiniteCore           bool
	YukawaTermsIncluded            bool
	SchemeIndependentForFiniteCore bool
	Verdict                        string
}

type StabilitySegment struct {
	Name                      string
	LStart                    float64
	LEnd                      float64
	ActiveRows                []string
	OneLoopBeta               twothresholdviability.FloatTriple
	TwoLoopMatrix             RationalMatrix3
	MaxRatio                  float64
	DominantGauge             string
	DominantOneLoopDerivative float64
	DominantTwoLoopDerivative float64
}

type TwoLoopStabilityAudit struct {
	SegmentsAudited                     int
	MaxTwoLoopToOneLoopRatio            float64
	DominantSegment                     string
	DominantGauge                       string
	CatastrophicDestabilizationDetected bool
	OneLoopPerturbativeControlProven    bool
	RequiresFullTwoLoopIntegration      bool
	RequiresMatchingCorrections         bool
	OneLoopScalesValidOnlyAtOneLoop     bool
	Status                              string
	Verdict                             string
}

type FirewallAudit struct {
	Gate212Inherited                 bool
	ThresholdSpectrumSealIntroduced  bool
	LeptoquarkDynamicsSealInherited  bool
	EmpiricalCarrierSealInherited    bool
	EmpiricalLedgerQuarantined       bool
	UniquePhysicalSpectrumClaimed    bool
	ContactModesPromotedToCarriers   bool
	BGapPromotedToMass               bool
	MatchingCorrectionsDerived       bool
	MSbarImportedAsFiniteCore        bool
	TwoLoopCoefficientsFiniteDerived bool
	TwoLoopScalesClaimedAsPrediction bool
	PhysicalPredictionClaimed        bool
	ProtonLifetimeComputed           bool
	OneLoopOnlyScaleWarning          bool
	RecommendedNextGate              string
	OpenRequirements                 []string
	Verdict                          string
}

type Summary struct {
	TestsAudited                      int
	Gate212Inherited                  bool
	ThresholdSpectrumSealIntroduced   bool
	SelectedBestPair                  bool
	MatchingCorrectionsDerived        bool
	MatchingCorrectionStatus          string
	ExactHeavyTwoLoopMatrixCalculated bool
	TwoLoopStatus                     string
	OneLoopScalesOnly                 bool
	Status                            string
	Comment                           string
}

type Analysis struct {
	Gate212         Gate212Snapshot
	Gate212Analysis twothresholdminimality.Analysis
	Seal            ThresholdSpectrumSeal
	Subject         SealedSpectrumSubject
	Matching        MatchingCorrectionAudit
	TwoLoop         TwoLoopCoefficientAudit
	Stability       TwoLoopStabilityAudit
	Segments        []StabilitySegment
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g212, err := twothresholdminimality.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(g212)
	})
	return defaultA, defaultErr
}

func Build(g212 twothresholdminimality.Analysis) (Analysis, error) {
	snap := SnapshotFromGate212(g212)
	if !snap.Gate212Inherited || !snap.CanonicalUniquenessFailedRoute || !snap.ThresholdSpectrumSealRequired {
		return Analysis{}, fmt.Errorf("Gate 213 requires Gate 212 canonical uniqueness obstruction and ThresholdSpectrumSeal requirement")
	}
	if snap.UniquePhysicalSpectrumClaimed || snap.FiniteCarrierOriginDerived || snap.FiniteMatchingCorrectionsDerived {
		return Analysis{}, fmt.Errorf("Gate 213 refuses inherited uniqueness, carrier-origin, or matching-correction leakage")
	}
	best, rank, err := selectBestTopologicalPair(g212.Gate211Analysis)
	if err != nil {
		return Analysis{}, err
	}
	seal := introduceSeal(snap)
	subject := buildSubject(best, rank)
	matching := auditMatchingCorrections()
	twoLoop, err := auditTwoLoopCoefficients(subject)
	if err != nil {
		return Analysis{}, err
	}
	segments, stability := auditTwoLoopStability(subject, twoLoop)
	fw := auditFirewall(snap, seal, matching, twoLoop, stability)
	status := StatusConditionalSealPreflight
	if !seal.DegeneracyQuarantined || !matching.ThresholdMatchingCoefficientsDerived == false || !twoLoop.ExactSymbolicHeavyCoefficients {
		status = StatusFailedSealPreflight
	}
	summary := Summary{
		TestsAudited:                      6,
		Gate212Inherited:                  snap.Gate212Inherited,
		ThresholdSpectrumSealIntroduced:   seal.DegeneracyQuarantined && !seal.FiniteUniquenessClaimed,
		SelectedBestPair:                  subject.ConditionalOnly && subject.SelectedFromOrderedRank == 1,
		MatchingCorrectionsDerived:        matching.ThresholdMatchingCoefficientsDerived,
		MatchingCorrectionStatus:          matching.Status,
		ExactHeavyTwoLoopMatrixCalculated: twoLoop.ExactSymbolicHeavyCoefficients,
		TwoLoopStatus:                     stability.Status,
		OneLoopScalesOnly:                 stability.OneLoopScalesValidOnlyAtOneLoop,
		Status:                            status,
		Comment:                           "Gate 213 seals the heavy-spectrum choice, audits matching-correction non-derivability, and computes a standard-QFT symbolic two-loop preflight for the Gate-211 ranked witness. The numerical scales remain one-loop, sealed phenomenology only.",
	}
	truth := buildTruth(subject, matching, stability)
	return Analysis{Gate212: snap, Gate212Analysis: g212, Seal: seal, Subject: subject, Matching: matching, TwoLoop: twoLoop, Stability: stability, Segments: segments, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func selectBestTopologicalPair(g211 twothresholdviability.Analysis) (twothresholdviability.PairSolution, int, error) {
	for _, ta := range g211.TargetAudits {
		if ta.Target.Name == "u_topological" {
			if len(ta.BestSolutions) == 0 {
				return twothresholdviability.PairSolution{}, 0, fmt.Errorf("Gate 213 requires a Gate-211 topological best witness")
			}
			return ta.BestSolutions[0], 1, nil
		}
	}
	return twothresholdviability.PairSolution{}, 0, fmt.Errorf("Gate 213 cannot find Gate-211 u_topological target audit")
}

func introduceSeal(s Gate212Snapshot) ThresholdSpectrumSeal {
	return ThresholdSpectrumSeal{
		Name:                         "ThresholdSpectrumSeal",
		SealID:                       "SEAL-THRESHOLD-SPECTRUM-GATE213",
		IntroducedByGate:             213,
		Reason:                       fmt.Sprintf("Gate 212 left %d ordered / %d unordered viable spectra without a canonical finite selector", s.OrderedViablePairs, s.UnorderedPairClasses),
		DegeneracyQuarantined:        true,
		SelectedPairPhenomenological: true,
		FiniteUniquenessClaimed:      false,
		ContactOrBGapOriginClaimed:   false,
		MatchingSchemeClaimed:        false,
		TwoLoopImportedAsFiniteCore:  false,
		OneLoopScalesRemainReference: true,
		Verdict:                      "seal active: the selected heavy spectrum is a quarantined phenomenological test subject, not a finite-derived unique pair",
	}
}

func buildSubject(p twothresholdviability.PairSolution, rank int) SealedSpectrumSubject {
	return SealedSpectrumSubject{
		TargetName:                p.TargetName,
		Row1Name:                  p.Row1Name,
		Row2Name:                  p.Row2Name,
		Row1Rep:                   p.Row1Rep,
		Row2Rep:                   p.Row2Rep,
		Row1DeltaB:                p.Row1DeltaB,
		Row2DeltaB:                p.Row2DeltaB,
		TotalDeltaB:               p.TotalDeltaB,
		TotalOneLoopBeta:          p.TotalBeta,
		LB1:                       p.LB1,
		LB2:                       p.LB2,
		LStar:                     p.LStar,
		MB1GeV:                    p.MB1GeV,
		MB2GeV:                    p.MB2GeV,
		MStarGeV:                  p.MStarGeV,
		DeltaL:                    math.Abs(p.LB2 - p.LB1),
		AlphaGUTInverse:           4.0 * math.Pi * p.TargetU,
		SelectedFromOrderedRank:   rank,
		SelectedFromUnorderedRank: 1,
		FiniteDerived:             false,
		ConditionalOnly:           true,
		Verdict:                   "Gate-211 ranked witness selected under ThresholdSpectrumSeal for preflight testing only; not a unique finite prediction",
	}
}

func auditMatchingCorrections() MatchingCorrectionAudit {
	return MatchingCorrectionAudit{
		TauEtaFundamentalClassAvailable:      true,
		ScalarTraceSupportAvailable:          true,
		ContactZetaTraceAvailable:            true,
		HeatKernelMatchingDerived:            false,
		SpectralTripleComplete:               false,
		CanonicalSubtractionSchemeDerived:    false,
		MSbarOrDimRegImported:                false,
		FiniteCountertermFunctionalDerived:   false,
		ThresholdMatchingCoefficientsDerived: false,
		MatchingCorrectionRows:               0,
		Status:                               MatchingCorrectionsFailed,
		Verdict:                              "finite tau_eta/scalar/contact traces are available as support data, but no spectral triple, heat-kernel matching map, subtraction scheme, or counterterm functional canonically derives δ_i^match; exact scales therefore retain scheme-dependent uncertainty",
	}
}

func auditTwoLoopCoefficients(s SealedSpectrumSubject) (TwoLoopCoefficientAudit, error) {
	c1, err := carrierMatrix(s.Row1Name, s.Row1Rep, s.Row1DeltaB)
	if err != nil {
		return TwoLoopCoefficientAudit{}, err
	}
	c2, err := carrierMatrix(s.Row2Name, s.Row2Rep, s.Row2DeltaB)
	if err != nil {
		return TwoLoopCoefficientAudit{}, err
	}
	heavy := c1.Matrix.Add(c2.Matrix)
	sm := smTwoLoopMatrix()
	total := sm.Add(heavy)
	return TwoLoopCoefficientAudit{
		Convention:                     "dg_i/dlnμ = g_i^3 b_i/(16π²) + g_i^3 Σ_j B_ij g_j²/(16π²)²; GUT-normalized U(1); Yukawa terms omitted",
		SMOneLoopBeta:                  twothresholdviability.FloatTriple{U1GUT: 41.0 / 10.0, SU2L: -19.0 / 6.0, SU3C: -7.0},
		SMTwoLoopMatrixNoYukawa:        sm,
		CarrierMatrices:                []TwoLoopCarrierMatrix{c1, c2},
		HeavyInducedTwoLoopMatrix:      heavy,
		TotalTwoLoopMatrixNoYukawa:     total,
		ExactSymbolicHeavyCoefficients: c1.Exact && c2.Exact,
		UsesStandardQFTFormula:         true,
		ImportedAsFiniteCore:           false,
		YukawaTermsIncluded:            false,
		SchemeIndependentForFiniteCore: false,
		Verdict:                        "exact rational heavy-carrier two-loop B_ij coefficients are computed as a standard-QFT preflight only; they are not finite-core theorems and do not include un-derived Yukawa or matching data",
	}, nil
}

func carrierMatrix(name, rep string, oneLoop representationrowlattice.RationalTriple) (TwoLoopCarrierMatrix, error) {
	parsed, err := parseRepresentation(name, rep)
	if err != nil {
		return TwoLoopCarrierMatrix{}, err
	}
	if !parsed.Supported {
		return TwoLoopCarrierMatrix{}, fmt.Errorf("unsupported two-loop carrier %s", name)
	}
	m := twoLoopDiracMatrix(parsed)
	return TwoLoopCarrierMatrix{
		CarrierName: name,
		Rep:         rep,
		Parsed:      parsed,
		OneLoopRow:  oneLoop,
		Matrix:      m,
		Formula:     "Dirac fermion: ΔB_ii=(20/3 C2(G_i)+4 C2_i(R)) S_i(R), ΔB_ij=4 C2_j(R) S_i(R) for i≠j",
		Exact:       true,
		FiniteCore:  false,
		Verdict:     "exact symbolic standard-QFT two-loop row for the sealed carrier; not a finite-derived action coefficient",
	}, nil
}

func parseRepresentation(name, rep string) (ParsedRepresentation, error) {
	stats := ""
	if strings.Contains(name, "Dirac fermion") {
		stats = "Dirac fermion"
	} else if strings.Contains(name, "Weyl fermion") {
		stats = "Weyl fermion"
	} else {
		return ParsedRepresentation{Name: name, Statistics: stats, Supported: false, Verdict: "two-loop preflight currently supports sealed fermion witnesses only"}, nil
	}
	re := regexp.MustCompile(`\(([^,]+),([^,]+),Y=([^\)]+)\)`) // (SU3,SU2,Y=q)
	m := re.FindStringSubmatch(rep)
	if len(m) != 4 {
		return ParsedRepresentation{}, fmt.Errorf("cannot parse representation %q", rep)
	}
	su3, ok := su3Invariant(m[1])
	if !ok {
		return ParsedRepresentation{}, fmt.Errorf("unknown SU3 rep %q", m[1])
	}
	su2, ok := su2Invariant(m[2])
	if !ok {
		return ParsedRepresentation{}, fmt.Errorf("unknown SU2 rep %q", m[2])
	}
	y, err := parseRat(m[3])
	if err != nil {
		return ParsedRepresentation{}, err
	}
	yc := mul(representationrowlattice.R(3, 5), mul(y, y))
	return ParsedRepresentation{Name: name, Statistics: stats, SU3: su3, SU2: su2, Y: y, YCasimir: yc, Supported: stats == "Dirac fermion", Verdict: "parsed sealed carrier representation"}, nil
}

func su3Invariant(sym string) (GroupInvariants, bool) {
	switch sym {
	case "1":
		return GroupInvariants{Symbol: sym, Dim: 1, T: representationrowlattice.R(0, 1), C2: representationrowlattice.R(0, 1), CG: representationrowlattice.R(3, 1)}, true
	case "3", "3bar":
		return GroupInvariants{Symbol: sym, Dim: 3, T: representationrowlattice.R(1, 2), C2: representationrowlattice.R(4, 3), CG: representationrowlattice.R(3, 1)}, true
	case "8":
		return GroupInvariants{Symbol: sym, Dim: 8, T: representationrowlattice.R(3, 1), C2: representationrowlattice.R(3, 1), CG: representationrowlattice.R(3, 1)}, true
	default:
		return GroupInvariants{}, false
	}
}

func su2Invariant(sym string) (GroupInvariants, bool) {
	switch sym {
	case "1":
		return GroupInvariants{Symbol: sym, Dim: 1, T: representationrowlattice.R(0, 1), C2: representationrowlattice.R(0, 1), CG: representationrowlattice.R(2, 1)}, true
	case "2":
		return GroupInvariants{Symbol: sym, Dim: 2, T: representationrowlattice.R(1, 2), C2: representationrowlattice.R(3, 4), CG: representationrowlattice.R(2, 1)}, true
	case "3":
		return GroupInvariants{Symbol: sym, Dim: 3, T: representationrowlattice.R(2, 1), C2: representationrowlattice.R(2, 1), CG: representationrowlattice.R(2, 1)}, true
	default:
		return GroupInvariants{}, false
	}
}

func parseRat(s string) (Rational, error) {
	s = strings.TrimSpace(s)
	if strings.Contains(s, "/") {
		parts := strings.Split(s, "/")
		if len(parts) != 2 {
			return Rational{}, fmt.Errorf("bad rational %q", s)
		}
		n, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return Rational{}, err
		}
		d, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return Rational{}, err
		}
		return representationrowlattice.R(n, d), nil
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return Rational{}, err
	}
	return representationrowlattice.R(n, 1), nil
}

func twoLoopDiracMatrix(p ParsedRepresentation) RationalMatrix3 {
	// S_i(R): Dynkin index for gauge factor i including multiplicity in the
	// spectator gauge factors. For U(1), C_1=T_1=(3/5)Y^2.
	s := [3]Rational{
		mulInt(p.YCasimir, p.SU3.Dim*p.SU2.Dim),
		mulInt(p.SU2.T, p.SU3.Dim),
		mulInt(p.SU3.T, p.SU2.Dim),
	}
	c := [3]Rational{p.YCasimir, p.SU2.C2, p.SU3.C2}
	cg := [3]Rational{representationrowlattice.R(0, 1), p.SU2.CG, p.SU3.CG}
	var out RationalMatrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				term := add(mul(representationrowlattice.R(20, 3), cg[i]), mulInt(c[i], 4))
				out.M[i][j] = mul(term, s[i])
			} else {
				out.M[i][j] = mulInt(mul(c[j], s[i]), 4)
			}
		}
	}
	return out
}

func smTwoLoopMatrix() RationalMatrix3 {
	return RationalMatrix3{M: [3][3]Rational{
		{representationrowlattice.R(199, 50), representationrowlattice.R(27, 10), representationrowlattice.R(44, 5)},
		{representationrowlattice.R(9, 10), representationrowlattice.R(35, 6), representationrowlattice.R(12, 1)},
		{representationrowlattice.R(11, 10), representationrowlattice.R(9, 2), representationrowlattice.R(-26, 1)},
	}}
}

func auditTwoLoopStability(s SealedSpectrumSubject, tl TwoLoopCoefficientAudit) ([]StabilitySegment, TwoLoopStabilityAudit) {
	segments := buildSegments(s, tl)
	maxRatio := 0.0
	domSeg := "none"
	domGauge := "none"
	for _, seg := range segments {
		if seg.MaxRatio > maxRatio {
			maxRatio = seg.MaxRatio
			domSeg = seg.Name
			domGauge = seg.DominantGauge
		}
	}
	status := TwoLoopStable
	catastrophic := false
	control := maxRatio < 0.5
	if maxRatio >= 1.0 {
		status = TwoLoopWarning
		catastrophic = false
	} else if maxRatio >= 0.5 {
		status = TwoLoopWarning
	}
	verdict := "two-loop preflight does not derive corrected scales; one-loop witness remains a reference solution only"
	if maxRatio >= 1.0 {
		verdict = fmt.Sprintf("standard-QFT two-loop preflight is not perturbatively small for %s/%s (max ratio %.6g); one-loop convergence is not proven stable without two-loop integration and matching corrections", domSeg, domGauge, maxRatio)
	} else if maxRatio >= 0.5 {
		verdict = fmt.Sprintf("standard-QFT two-loop preflight gives moderate corrections (max ratio %.6g); stability requires full two-loop integration and matching corrections", maxRatio)
	} else {
		verdict = fmt.Sprintf("standard-QFT two-loop preflight corrections are small under the ratio audit (max ratio %.6g), but matching corrections remain un-derived", maxRatio)
	}
	return segments, TwoLoopStabilityAudit{
		SegmentsAudited:                     len(segments),
		MaxTwoLoopToOneLoopRatio:            maxRatio,
		DominantSegment:                     domSeg,
		DominantGauge:                       domGauge,
		CatastrophicDestabilizationDetected: catastrophic,
		OneLoopPerturbativeControlProven:    control,
		RequiresFullTwoLoopIntegration:      true,
		RequiresMatchingCorrections:         true,
		OneLoopScalesValidOnlyAtOneLoop:     true,
		Status:                              status,
		Verdict:                             verdict,
	}
}

func buildSegments(s SealedSpectrumSubject, tl TwoLoopCoefficientAudit) []StabilitySegment {
	// The selected Gate-211 witness has LB1 < LB2. Keep sorting defensive so the
	// sealed subject remains auditable if a future ranked witness reverses order.
	type threshold struct {
		L    float64
		Name string
		B    twothresholdviability.FloatTriple
		M    RationalMatrix3
	}
	c1, c2 := tl.CarrierMatrices[0], tl.CarrierMatrices[1]
	thresholds := []threshold{
		{L: s.LB1, Name: s.Row1Name, B: toFloat(s.Row1DeltaB), M: c1.Matrix},
		{L: s.LB2, Name: s.Row2Name, B: toFloat(s.Row2DeltaB), M: c2.Matrix},
	}
	sort.Slice(thresholds, func(i, j int) bool { return thresholds[i].L < thresholds[j].L })
	baseB := twothresholdviability.FloatTriple{U1GUT: 41.0 / 10.0, SU2L: -19.0 / 6.0, SU3C: -7.0}
	baseM := smTwoLoopMatrix()
	points := []float64{0, thresholds[0].L, thresholds[1].L, s.LStar}
	activeB := baseB
	activeM := baseM
	activeNames := []string{"SM"}
	out := []StabilitySegment{}
	for k := 0; k < 3; k++ {
		if k == 1 {
			activeB = addFT(activeB, thresholds[0].B)
			activeM = activeM.Add(thresholds[0].M)
			activeNames = append(activeNames, thresholds[0].Name)
		}
		if k == 2 {
			activeB = addFT(activeB, thresholds[1].B)
			activeM = activeM.Add(thresholds[1].M)
			activeNames = append(activeNames, thresholds[1].Name)
		}
		seg := StabilitySegment{Name: fmt.Sprintf("segment-%d", k), LStart: points[k], LEnd: points[k+1], ActiveRows: append([]string(nil), activeNames...), OneLoopBeta: activeB, TwoLoopMatrix: activeM}
		seg.MaxRatio, seg.DominantGauge, seg.DominantOneLoopDerivative, seg.DominantTwoLoopDerivative = segmentMaxRatio(s, seg)
		out = append(out, seg)
	}
	return out
}

func segmentMaxRatio(s SealedSpectrumSubject, seg StabilitySegment) (float64, string, float64, float64) {
	points := []float64{seg.LStart, seg.LEnd}
	maxR := 0.0
	dom := "none"
	dom1 := 0.0
	dom2 := 0.0
	for _, l := range points {
		u := runningUAtSubject(s, l)
		for i := 0; i < 3; i++ {
			one := math.Abs(betaAt(seg.OneLoopBeta, i)) / (8.0 * math.Pi * math.Pi)
			if one == 0 {
				continue
			}
			sum := 0.0
			for j := 0; j < 3; j++ {
				uj := u[j]
				if uj <= 0 || math.IsNaN(uj) {
					continue
				}
				sum += seg.TwoLoopMatrix.FloatAt(i, j) / uj
			}
			two := math.Abs(sum) / (128.0 * math.Pow(math.Pi, 4))
			r := two / one
			if r > maxR {
				maxR = r
				dom = gaugeName(i)
				dom1 = one
				dom2 = two
			}
		}
	}
	return maxR, dom, dom1, dom2
}

func runningUAtSubject(s SealedSpectrumSubject, l float64) [3]float64 {
	// Reconstruct the Gate-211 one-loop u trajectory for the sealed subject.
	inputsU := [3]float64{((3.0 / 5.0) * (1.0 - 0.23122) * 127.955) / (4.0 * math.Pi), (0.23122 * 127.955) / (4.0 * math.Pi), (1.0 / 0.1179) / (4.0 * math.Pi)}
	b := [3]float64{41.0 / 10.0, -19.0 / 6.0, -7.0}
	d1 := [3]float64{s.Row1DeltaB.B1.Float(), s.Row1DeltaB.B2.Float(), s.Row1DeltaB.B3.Float()}
	d2 := [3]float64{s.Row2DeltaB.B1.Float(), s.Row2DeltaB.B2.Float(), s.Row2DeltaB.B3.Float()}
	out := [3]float64{}
	for i := 0; i < 3; i++ {
		active1 := math.Max(0, l-s.LB1)
		active2 := math.Max(0, l-s.LB2)
		out[i] = inputsU[i] - (b[i]*l+d1[i]*active1+d2[i]*active2)/(8.0*math.Pi*math.Pi)
	}
	return out
}

func auditFirewall(s Gate212Snapshot, seal ThresholdSpectrumSeal, m MatchingCorrectionAudit, tl TwoLoopCoefficientAudit, st TwoLoopStabilityAudit) FirewallAudit {
	return FirewallAudit{
		Gate212Inherited:                 s.Gate212Inherited,
		ThresholdSpectrumSealIntroduced:  seal.DegeneracyQuarantined,
		LeptoquarkDynamicsSealInherited:  s.LeptoquarkDynamicsSealInherited,
		EmpiricalCarrierSealInherited:    s.EmpiricalCarrierSealInherited,
		EmpiricalLedgerQuarantined:       true,
		UniquePhysicalSpectrumClaimed:    seal.FiniteUniquenessClaimed,
		ContactModesPromotedToCarriers:   s.FiniteCarrierOriginDerived || seal.ContactOrBGapOriginClaimed,
		BGapPromotedToMass:               false,
		MatchingCorrectionsDerived:       m.ThresholdMatchingCoefficientsDerived,
		MSbarImportedAsFiniteCore:        m.MSbarOrDimRegImported,
		TwoLoopCoefficientsFiniteDerived: tl.ImportedAsFiniteCore,
		TwoLoopScalesClaimedAsPrediction: false,
		PhysicalPredictionClaimed:        false,
		ProtonLifetimeComputed:           false,
		OneLoopOnlyScaleWarning:          st.OneLoopScalesValidOnlyAtOneLoop,
		RecommendedNextGate:              "Gate 214 — two-loop RG integration and matching-envelope uncertainty audit, still sealed",
		OpenRequirements: []string{
			"derive or explicitly choose a threshold matching scheme before precision scale claims",
			"run full two-loop piecewise integration only as conditional phenomenology",
			"derive finite spectral triple / heat-kernel matching before promoting δ_i^match",
			"keep ThresholdSpectrumSeal, EmpiricalCarrierSeal, and LeptoquarkDynamicsSeal active",
		},
		Verdict: "firewall preserved: the heavy spectrum, matching scheme, and two-loop corrections are quarantined; one-loop Gate-211 scales are not finite predictions",
	}
}

func buildTruth(s SealedSpectrumSubject, m MatchingCorrectionAudit, st TwoLoopStabilityAudit) string {
	return fmt.Sprintf("Gate 213 introduces the ThresholdSpectrumSeal because Gate 212 could not canonically choose among the Gate-211 spectra. The ranked pair %s + %s is used only as a conditional test subject. Matching corrections remain a failed route (%s): finite tau_eta/scalar/contact traces do not provide a subtraction scheme or δ_i^match. Standard-QFT two-loop coefficients are computed exactly for the sealed carriers, but the stability preflight returns %s with max two-loop/one-loop derivative ratio %.6g. Therefore the Gate-211 scales M_B1=%.9g GeV, M_B2=%.9g GeV, M*=%.9g GeV remain one-loop sealed phenomenology, not finite-core predictions.", s.Row1Rep, s.Row2Rep, m.Status, st.Status, st.MaxTwoLoopToOneLoopRatio, s.MB1GeV, s.MB2GeV, s.MStarGeV)
}

func add(a, b Rational) Rational {
	return representationrowlattice.R(a.Num*b.Den+b.Num*a.Den, a.Den*b.Den)
}
func mul(a, b Rational) Rational          { return representationrowlattice.R(a.Num*b.Num, a.Den*b.Den) }
func mulInt(a Rational, n int64) Rational { return representationrowlattice.R(a.Num*n, a.Den) }

func addFT(a, b twothresholdviability.FloatTriple) twothresholdviability.FloatTriple {
	return twothresholdviability.FloatTriple{U1GUT: a.U1GUT + b.U1GUT, SU2L: a.SU2L + b.SU2L, SU3C: a.SU3C + b.SU3C}
}

func toFloat(t representationrowlattice.RationalTriple) twothresholdviability.FloatTriple {
	return twothresholdviability.FloatTriple{U1GUT: t.B1.Float(), SU2L: t.B2.Float(), SU3C: t.B3.Float()}
}

func betaAt(b twothresholdviability.FloatTriple, i int) float64 {
	switch i {
	case 0:
		return b.U1GUT
	case 1:
		return b.SU2L
	case 2:
		return b.SU3C
	default:
		panic("bad gauge index")
	}
}

func gaugeName(i int) string {
	switch i {
	case 0:
		return "U(1)_Y"
	case 1:
		return "SU(2)_L"
	case 2:
		return "SU(3)_C"
	default:
		return "unknown"
	}
}

func FormatSeal(s ThresholdSpectrumSeal) string {
	return fmt.Sprintf("%s id=%s gate=%d quarantined=%t selectedPhenomenology=%t uniqueClaim=%t matchingScheme=%t oneLoopRef=%t", s.Name, s.SealID, s.IntroducedByGate, s.DegeneracyQuarantined, s.SelectedPairPhenomenological, s.FiniteUniquenessClaimed, s.MatchingSchemeClaimed, s.OneLoopScalesRemainReference)
}

func FormatSubject(s SealedSpectrumSubject) string {
	return fmt.Sprintf("target=%s rows=[%s %s; %s %s] Δb=[%s;%s] total=%s LB=(%.9g,%.9g) MB=(%.9g,%.9g) L*=%.9g M*=%.9g ΔL=%.9g finite=%t conditional=%t", s.TargetName, s.Row1Name, s.Row1Rep, s.Row2Name, s.Row2Rep, s.Row1DeltaB, s.Row2DeltaB, s.TotalDeltaB, s.LB1, s.LB2, s.MB1GeV, s.MB2GeV, s.LStar, s.MStarGeV, s.DeltaL, s.FiniteDerived, s.ConditionalOnly)
}

func FormatMatching(a MatchingCorrectionAudit) string {
	return fmt.Sprintf("tauEta=%t scalarTrace=%t contactZeta=%t heatKernel=%t spectralTriple=%t scheme=%t msbarImported=%t counterterm=%t rows=%d status=%s", a.TauEtaFundamentalClassAvailable, a.ScalarTraceSupportAvailable, a.ContactZetaTraceAvailable, a.HeatKernelMatchingDerived, a.SpectralTripleComplete, a.CanonicalSubtractionSchemeDerived, a.MSbarOrDimRegImported, a.FiniteCountertermFunctionalDerived, a.MatchingCorrectionRows, a.Status)
}

func FormatTwoLoop(a TwoLoopCoefficientAudit) string {
	return fmt.Sprintf("heavy=%s total=%s exactHeavy=%t standardQFT=%t finiteCore=%t yukawa=%t schemeIndependent=%t", a.HeavyInducedTwoLoopMatrix, a.TotalTwoLoopMatrixNoYukawa, a.ExactSymbolicHeavyCoefficients, a.UsesStandardQFTFormula, a.ImportedAsFiniteCore, a.YukawaTermsIncluded, a.SchemeIndependentForFiniteCore)
}

func FormatStability(a TwoLoopStabilityAudit) string {
	return fmt.Sprintf("segments=%d maxRatio=%.9g dominant=%s/%s catastrophic=%t control=%t fullIntegration=%t matching=%t oneLoopOnly=%t status=%s", a.SegmentsAudited, a.MaxTwoLoopToOneLoopRatio, a.DominantSegment, a.DominantGauge, a.CatastrophicDestabilizationDetected, a.OneLoopPerturbativeControlProven, a.RequiresFullTwoLoopIntegration, a.RequiresMatchingCorrections, a.OneLoopScalesValidOnlyAtOneLoop, a.Status)
}

func FormatSegment(s StabilitySegment) string {
	return fmt.Sprintf("%s L=[%.9g,%.9g] active=%s b=%s maxRatio=%.9g dominant=%s", s.Name, s.LStart, s.LEnd, strings.Join(s.ActiveRows, "+"), s.OneLoopBeta.String(), s.MaxRatio, s.DominantGauge)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate212=%t spectrumSeal=%t lqSeal=%t carrierSeal=%t ledger=%t uniqueClaim=%t contactPromoted=%t bgapMass=%t matching=%t msbarCore=%t twoLoopCore=%t twoLoopPrediction=%t physicalPrediction=%t lifetime=%t oneLoopWarning=%t next=%s", a.Gate212Inherited, a.ThresholdSpectrumSealIntroduced, a.LeptoquarkDynamicsSealInherited, a.EmpiricalCarrierSealInherited, a.EmpiricalLedgerQuarantined, a.UniquePhysicalSpectrumClaimed, a.ContactModesPromotedToCarriers, a.BGapPromotedToMass, a.MatchingCorrectionsDerived, a.MSbarImportedAsFiniteCore, a.TwoLoopCoefficientsFiniteDerived, a.TwoLoopScalesClaimedAsPrediction, a.PhysicalPredictionClaimed, a.ProtonLifetimeComputed, a.OneLoopOnlyScaleWarning, a.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate212=%t seal=%t selected=%t matchingDerived=%t matchingStatus=%s exactTwoLoop=%t twoLoopStatus=%s oneLoopOnly=%t status=%s", s.TestsAudited, s.Gate212Inherited, s.ThresholdSpectrumSealIntroduced, s.SelectedBestPair, s.MatchingCorrectionsDerived, s.MatchingCorrectionStatus, s.ExactHeavyTwoLoopMatrixCalculated, s.TwoLoopStatus, s.OneLoopScalesOnly, s.Status)
}
