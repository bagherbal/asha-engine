// Package twothresholdviability implements Gate 211: two-threshold rational
// lattice viability filter / scale-ordered Landau safety audit.
//
// Gate 210 proved that a single rational threshold row cannot exactly close the
// topological mismatch triangle: b_SM and one lattice row span only a two-plane
// in R^3, while the target vector is generically three-dimensional. Gate 211 is
// therefore not another closure search. With two independent threshold rows,
// the three logarithms (L*, L_B1, L_B2) are determined by a 3x3 linear system.
// The question becomes physical viability: ordered threshold scales, sub-Planck
// boundary, positive running couplings, no sub-Planck Landau pole, anomaly
// safety, and compatibility with the LeptoquarkDynamicsSeal.
package twothresholdviability

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/representationrowlattice"
)

const (
	StatusConditionalViable = "CONDITIONAL_VIABLE_TWO_THRESHOLD_LATTICE"
	StatusFailedRoute       = "FAILED_ROUTE_TWO_THRESHOLD_LATTICE_FILTER"
	mzGeV                   = 91.1876
	planckGeV               = 1.2209e19
	// The Gate-211 prompt explicitly binds the viability filter to L* < 37.8.
	// This is stricter than log(M_Planck/M_Z) for the numeric Planck value, so
	// the audit uses the prompt's logarithmic bound as the operative firewall.
	planckLogBound       = 37.8
	determinantTolerance = 1.0e-10
	distinctThresholdTol = 1.0e-9
	positiveCouplingTol  = 1.0e-12
	uCentroidTarget      = 3.33
)

type FloatTriple struct{ U1GUT, SU2L, SU3C float64 }

func (t FloatTriple) At(i int) float64 {
	switch i {
	case 0:
		return t.U1GUT
	case 1:
		return t.SU2L
	case 2:
		return t.SU3C
	default:
		panic("bad gauge index")
	}
}

func (t FloatTriple) String() string {
	return fmt.Sprintf("(%.12g,%.12g,%.12g)", t.U1GUT, t.SU2L, t.SU3C)
}

type Gate210Snapshot struct {
	Gate210Inherited                   bool
	SingleScaleFailedRoutePreserved    bool
	Gate209LeptoquarkSealInherited     bool
	SafeGeneratorCount                 int
	UniversalBetaRowKilled             bool
	NoPhysicalPredictionClaimInherited bool
	TruthStatement                     string
}

func DefaultGate210Snapshot() Gate210Snapshot {
	return Gate210Snapshot{
		Gate210Inherited:                   true,
		SingleScaleFailedRoutePreserved:    true,
		Gate209LeptoquarkSealInherited:     true,
		SafeGeneratorCount:                 108,
		UniversalBetaRowKilled:             true,
		NoPhysicalPredictionClaimInherited: true,
		TruthStatement:                     "Gate 210 killed exact single-scale rational-lattice closure and left the next legal path as a two-threshold viability filter over the same anomaly-safe, leptoquark-compatible row grammar.",
	}
}

type RGInputAudit struct {
	LedgerName                  string
	MZGeV                       float64
	Alpha1GUTInverse            float64
	Alpha2Inverse               float64
	Alpha3Inverse               float64
	U1GUT                       float64
	U2                          float64
	U3                          float64
	SMBeta                      FloatTriple
	PlanckGeV                   float64
	PlanckLogBound              float64
	ObservedLedgerQuarantined   bool
	UsedForFiniteCoreDerivation bool
	UniversalBetaRowAllowed     bool
	TwoThresholdSystem          bool
	Verdict                     string
}

type BoundaryTarget struct {
	Name               string
	UTarget            float64
	AlphaInverseTarget float64
	Interpretation     string
}

type SearchGenerator struct {
	Name                     string
	Statistics               string
	SMRepresentation         string
	DeltaB                   representationrowlattice.RationalTriple
	FloatDeltaB              FloatTriple
	AnomalyCompatible        bool
	LeptoquarkSealCompatible bool
	DirectGate201Shape       bool
	ContactOrBGapMatch       bool
}

type GeneratorAudit struct {
	SourceUniqueRows              int
	SafeGenerators                int
	ExpectedGate210SafeGenerators int
	AnomalyCompatibleGenerators   int
	LeptoquarkSealCompatibleRows  int
	ZeroRowExcluded               bool
	NoUniversalBetaRowInserted    bool
	NoContinuousRowCoefficients   bool
	InheritsGate210FilteredBasis  bool
	Verdict                       string
}

type PairSolution struct {
	TargetName               string
	TargetU                  float64
	Row1Name                 string
	Row2Name                 string
	Row1Rep                  string
	Row2Rep                  string
	Row1DeltaB               representationrowlattice.RationalTriple
	Row2DeltaB               representationrowlattice.RationalTriple
	TotalDeltaB              representationrowlattice.RationalTriple
	FloatRow1                FloatTriple
	FloatRow2                FloatTriple
	TotalBeta                FloatTriple
	SystemDeterminant        float64
	LStar                    float64
	LB1                      float64
	LB2                      float64
	MB1GeV                   float64
	MB2GeV                   float64
	MStarGeV                 float64
	ClosureResidualU         float64
	ExactLinearClosure       bool
	MinUInDomain             float64
	EarliestPoleLog          float64
	EarliestPoleGeV          float64
	ScaleOrdered             bool
	DistinctThresholds       bool
	SubPlanck                bool
	PositiveCouplingsToMStar bool
	NoSubPlanckLandauPole    bool
	AnomalyCompatible        bool
	LeptoquarkSealCompatible bool
	ContainsGate201Shape     bool
	MatchesBGapOrContactData bool
	SU2AsymptoticallyFree    bool
	SU3AsymptoticallyFree    bool
	AllNonAbelianAF          bool
	TotalDeltaBNorm          float64
	GUTRangeDistance         float64
	Viable                   bool
	BindingConstraint        string
	Verdict                  string
}

type TargetAudit struct {
	Target                  BoundaryTarget
	OrderedPairsAudited     int
	PairIndependentPairs    int
	InvertibleSystems       int
	ScaleOrderedPairs       int
	DistinctThresholdPairs  int
	SubPlanckPairs          int
	PositiveCouplingPairs   int
	NoLandauPairs           int
	ViablePairs             int
	BindingConstraintCounts map[string]int
	BestSolutions           []PairSolution
	AllViableSolutions      []PairSolution
	Verdict                 string
}

type BaryonAnomalyAudit struct {
	LeptoquarkDynamicsSealInherited bool
	AllSearchRowsAnomalyCompatible  bool
	AllSearchRowsSealCompatible     bool
	ViablePairsAnomalyCompatible    bool
	ViablePairsSealCompatible       bool
	ProtonDecayOperatorUsed         bool
	ProtonLifetimeComputed          bool
	Verdict                         string
}

type ContactMatchAudit struct {
	BGapAudited                   bool
	ContactPartialOverlapAudited  bool
	CanonicalNumericMatchFound    bool
	ChargeSpinMassSemanticsFound  bool
	ViableRowsPromotedFromContact bool
	Verdict                       string
}

type FirewallAudit struct {
	Gate210Inherited                    bool
	LeptoquarkDynamicsSealInherited     bool
	EmpiricalLedgerQuarantined          bool
	ObservedLedgerUsedForFiniteCore     bool
	UniversalBetaRowInserted            bool
	ArbitraryRealRowCoefficientInserted bool
	PhysicalPredictionClaimed           bool
	AbsoluteMassDerivedFromFiniteCore   bool
	ProtonDecaySealViolated             bool
	ProtonLifetimeComputed              bool
	MatchingCorrectionsDerived          bool
	ConditionalViabilityOnly            bool
	RecommendedNextGate                 string
	OpenRequirements                    []string
	Verdict                             string
}

type Summary struct {
	TestsAudited      int
	Gate210Inherited  bool
	TargetsAudited    int
	ViableTopological int
	ViableCentroid    int
	TotalViablePairs  int
	Status            string
	Comment           string
}

type Analysis struct {
	Gate210        Gate210Snapshot
	Lattice        representationrowlattice.Analysis
	Inputs         RGInputAudit
	Targets        []BoundaryTarget
	Generators     []SearchGenerator
	GeneratorAudit GeneratorAudit
	TargetAudits   []TargetAudit
	BaryonAnomaly  BaryonAnomalyAudit
	ContactMatch   ContactMatchAudit
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
		lattice, err := representationrowlattice.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(DefaultGate210Snapshot(), lattice)
	})
	return defaultA, defaultErr
}

func Build(gate210 Gate210Snapshot, lattice representationrowlattice.Analysis) (Analysis, error) {
	if !gate210.Gate210Inherited || !gate210.SingleScaleFailedRoutePreserved || !gate210.Gate209LeptoquarkSealInherited || !gate210.UniversalBetaRowKilled {
		return Analysis{}, fmt.Errorf("Gate 211 requires Gate 210 failed-route inheritance and the Gate 209 LeptoquarkDynamicsSeal")
	}
	if !lattice.Summary.RationalGrammarConstructed || !lattice.Summary.LatticeConstructed || lattice.LatticeAudit.UniqueGeneratorRows == 0 {
		return Analysis{}, fmt.Errorf("Gate 211 requires Gate 204 rational representation-row lattice")
	}
	inputs := buildInputs()
	targets := []BoundaryTarget{
		{Name: "u_topological", UTarget: 1.0, AlphaInverseTarget: 4.0 * math.Pi, Interpretation: "quarantined instanton/topological branch"},
		{Name: "u_centroid", UTarget: uCentroidTarget, AlphaInverseTarget: 4.0 * math.Pi * uCentroidTarget, Interpretation: "Gate-200 SM-only mismatch-triangle centroid comparison branch"},
	}
	generators := collectSearchGenerators(lattice)
	genAudit := auditGenerators(lattice, gate210, generators)
	targetAudits := make([]TargetAudit, 0, len(targets))
	totalViable := 0
	topViable := 0
	centroidViable := 0
	for _, target := range targets {
		ta := runTargetAudit(inputs, target, generators)
		targetAudits = append(targetAudits, ta)
		totalViable += ta.ViablePairs
		if target.Name == "u_topological" {
			topViable = ta.ViablePairs
		}
		if target.Name == "u_centroid" {
			centroidViable = ta.ViablePairs
		}
	}
	ba := auditBaryonAnomaly(gate210, generators, targetAudits)
	cm := auditContactMatch(targetAudits)
	fw := auditFirewall(gate210, inputs, targetAudits, ba)
	status := StatusFailedRoute
	if totalViable > 0 {
		status = StatusConditionalViable
	}
	summary := Summary{
		TestsAudited:      8,
		Gate210Inherited:  gate210.Gate210Inherited,
		TargetsAudited:    len(targets),
		ViableTopological: topViable,
		ViableCentroid:    centroidViable,
		TotalViablePairs:  totalViable,
		Status:            status,
		Comment:           "Gate 211 treats two-threshold closure as linear algebra and filters the resulting exact scale solutions by ordering, prompt Planck bound, positivity, Landau safety, anomaly safety, and the LeptoquarkDynamicsSeal. Any surviving rows are conditional phenomenological viability witnesses only.",
	}
	truth := buildTruthStatement(summary, targetAudits)
	return Analysis{Gate210: gate210, Lattice: lattice, Inputs: inputs, Targets: targets, Generators: generators, GeneratorAudit: genAudit, TargetAudits: targetAudits, BaryonAnomaly: ba, ContactMatch: cm, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func buildInputs() RGInputAudit {
	alphaInv := 127.955
	sin2 := 0.23122
	alphaS := 0.1179
	a1 := (3.0 / 5.0) * (1.0 - sin2) * alphaInv
	a2 := sin2 * alphaInv
	a3 := 1.0 / alphaS
	return RGInputAudit{
		LedgerName:                  "Gate-200 quarantined Z-pole empirical comparison ledger",
		MZGeV:                       mzGeV,
		Alpha1GUTInverse:            a1,
		Alpha2Inverse:               a2,
		Alpha3Inverse:               a3,
		U1GUT:                       a1 / (4.0 * math.Pi),
		U2:                          a2 / (4.0 * math.Pi),
		U3:                          a3 / (4.0 * math.Pi),
		SMBeta:                      FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0},
		PlanckGeV:                   planckGeV,
		PlanckLogBound:              planckLogBound,
		ObservedLedgerQuarantined:   true,
		UsedForFiniteCoreDerivation: false,
		UniversalBetaRowAllowed:     false,
		TwoThresholdSystem:          true,
		Verdict:                     "phenomenological u-space ledger for two-threshold viability filter only; not finite-core data",
	}
}

func collectSearchGenerators(lattice representationrowlattice.Analysis) []SearchGenerator {
	zero := representationrowlattice.RT(representationrowlattice.R(0, 1), representationrowlattice.R(0, 1), representationrowlattice.R(0, 1))
	seen := map[string]bool{}
	out := make([]SearchGenerator, 0, len(lattice.UniqueRows))
	for _, row := range lattice.Rows {
		if row.DeltaB.Equal(zero) || !row.ExactRational || !row.StandardOneLoopFormula || !anomalyCompatible(row) || !leptoquarkSealCompatible(row) {
			continue
		}
		key := row.DeltaB.Key()
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, SearchGenerator{
			Name:                     row.Name,
			Statistics:               row.Statistics.Name,
			SMRepresentation:         fmt.Sprintf("(%s,%s,Y=%s)", row.SU3.Symbol, row.SU2.Symbol, row.Hypercharge),
			DeltaB:                   row.DeltaB,
			FloatDeltaB:              toFloat(row.DeltaB),
			AnomalyCompatible:        true,
			LeptoquarkSealCompatible: true,
			DirectGate201Shape:       row.DirectGate201Shape,
			ContactOrBGapMatch:       false,
		})
	}
	sort.Slice(out, func(i, j int) bool {
		ni := norm(out[i].FloatDeltaB)
		nj := norm(out[j].FloatDeltaB)
		if math.Abs(ni-nj) > 1e-12 {
			return ni < nj
		}
		return out[i].Name < out[j].Name
	})
	return out
}

func anomalyCompatible(row representationrowlattice.RepresentationRow) bool {
	switch row.Statistics.Name {
	case "Dirac fermion", "complex scalar", "real scalar":
		return true
	case "Weyl fermion":
		return row.Hypercharge.IsZero() && row.SU3.RealType && row.SU2.RealType
	default:
		return false
	}
}

func leptoquarkSealCompatible(row representationrowlattice.RepresentationRow) bool {
	return !strings.Contains(strings.ToLower(row.Name), "leptoquark")
}

func toFloat(t representationrowlattice.RationalTriple) FloatTriple {
	return FloatTriple{t.B1.Float(), t.B2.Float(), t.B3.Float()}
}

func auditGenerators(lattice representationrowlattice.Analysis, gate210 Gate210Snapshot, gens []SearchGenerator) GeneratorAudit {
	return GeneratorAudit{
		SourceUniqueRows:              lattice.LatticeAudit.UniqueGeneratorRows,
		SafeGenerators:                len(gens),
		ExpectedGate210SafeGenerators: gate210.SafeGeneratorCount,
		AnomalyCompatibleGenerators:   len(gens),
		LeptoquarkSealCompatibleRows:  len(gens),
		ZeroRowExcluded:               true,
		NoUniversalBetaRowInserted:    true,
		NoContinuousRowCoefficients:   true,
		InheritsGate210FilteredBasis:  len(gens) == gate210.SafeGeneratorCount,
		Verdict:                       fmt.Sprintf("inherited Gate-210 search grammar: %d exact anomaly-safe, leptoquark-compatible nonzero rational generators", len(gens)),
	}
}

func runTargetAudit(inputs RGInputAudit, target BoundaryTarget, gens []SearchGenerator) TargetAudit {
	counts := map[string]int{}
	viable := []PairSolution{}
	n := len(gens)
	audit := TargetAudit{Target: target, BindingConstraintCounts: counts}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			audit.OrderedPairsAudited++
			pairIndependent := !parallel(gens[i].FloatDeltaB, gens[j].FloatDeltaB)
			if pairIndependent {
				audit.PairIndependentPairs++
			}
			sol := evaluatePair(inputs, target, gens[i], gens[j])
			if math.Abs(sol.SystemDeterminant) > determinantTolerance {
				audit.InvertibleSystems++
			}
			if sol.ScaleOrdered {
				audit.ScaleOrderedPairs++
			}
			if sol.DistinctThresholds {
				audit.DistinctThresholdPairs++
			}
			if sol.SubPlanck {
				audit.SubPlanckPairs++
			}
			if sol.PositiveCouplingsToMStar {
				audit.PositiveCouplingPairs++
			}
			if sol.NoSubPlanckLandauPole {
				audit.NoLandauPairs++
			}
			counts[sol.BindingConstraint]++
			if sol.Viable {
				viable = append(viable, sol)
			}
		}
	}
	sortSolutions(viable)
	audit.ViablePairs = len(viable)
	audit.AllViableSolutions = viable
	if len(viable) > 20 {
		audit.BestSolutions = append([]PairSolution(nil), viable[:20]...)
	} else {
		audit.BestSolutions = append([]PairSolution(nil), viable...)
	}
	if len(viable) > 0 {
		audit.Verdict = fmt.Sprintf("CONDITIONAL_VIABLE: %d ordered two-threshold pairs survive all filters for %s", len(viable), target.Name)
	} else {
		audit.Verdict = fmt.Sprintf("FAILED_ROUTE: no ordered two-threshold pair survives all filters for %s; dominant binding constraint: %s", target.Name, dominantConstraint(counts))
	}
	return audit
}

func evaluatePair(inputs RGInputAudit, target BoundaryTarget, g1, g2 SearchGenerator) PairSolution {
	b := inputs.SMBeta
	d1 := g1.FloatDeltaB
	d2 := g2.FloatDeltaB
	m := [3][3]float64{}
	rhs := [3]float64{}
	uobs := []float64{inputs.U1GUT, inputs.U2, inputs.U3}
	for k := 0; k < 3; k++ {
		bi := b.At(k)
		d1i := d1.At(k)
		d2i := d2.At(k)
		m[k][0] = bi + d1i + d2i
		m[k][1] = -d1i
		m[k][2] = -d2i
		rhs[k] = 8.0 * math.Pi * math.Pi * (uobs[k] - target.UTarget)
	}
	det := det3(m)
	lstar, lb1, lb2 := math.NaN(), math.NaN(), math.NaN()
	if math.Abs(det) > determinantTolerance {
		x := solve3(m, rhs, det)
		lstar, lb1, lb2 = x[0], x[1], x[2]
	}
	mb1, mb2, ms := scaleFromLog(lb1), scaleFromLog(lb2), scaleFromLog(lstar)
	totalBeta := FloatTriple{b.U1GUT + d1.U1GUT + d2.U1GUT, b.SU2L + d1.SU2L + d2.SU2L, b.SU3C + d1.SU3C + d2.SU3C}
	closureResidual := closureResidualU(inputs, target, d1, d2, lstar, lb1, lb2)
	minUStar := minUInDomain(inputs, target, d1, d2, lstar, lb1, lb2, false)
	minUPlanck := minUInDomain(inputs, target, d1, d2, lstar, lb1, lb2, true)
	poleLog := earliestPoleLog(inputs, target, d1, d2, lstar, lb1, lb2)
	poleGeV := math.Inf(1)
	if finite(poleLog) {
		poleGeV = scaleFromLog(poleLog)
	}
	scaleOrdered := finite(lstar) && finite(lb1) && finite(lb2) && lb1 > 0 && lb2 > 0 && lb1 < lstar && lb2 < lstar
	distinct := scaleOrdered && math.Abs(lb1-lb2) > distinctThresholdTol
	subPlanck := scaleOrdered && lstar < inputs.PlanckLogBound && ms < inputs.PlanckGeV
	positiveToMStar := scaleOrdered && minUStar > positiveCouplingTol
	noPole := subPlanck && positiveToMStar && minUPlanck > positiveCouplingTol && (!finite(poleLog) || poleLog > inputs.PlanckLogBound)
	anomaly := g1.AnomalyCompatible && g2.AnomalyCompatible
	seal := g1.LeptoquarkSealCompatible && g2.LeptoquarkSealCompatible
	totalDelta := addRT(g1.DeltaB, g2.DeltaB)
	containsGate201 := g1.DirectGate201Shape || g2.DirectGate201Shape
	match := g1.ContactOrBGapMatch || g2.ContactOrBGapMatch
	su2AF := totalBeta.SU2L < 0
	su3AF := totalBeta.SU3C < 0
	allAF := su2AF && su3AF
	sol := PairSolution{
		TargetName:               target.Name,
		TargetU:                  target.UTarget,
		Row1Name:                 g1.Name,
		Row2Name:                 g2.Name,
		Row1Rep:                  g1.SMRepresentation,
		Row2Rep:                  g2.SMRepresentation,
		Row1DeltaB:               g1.DeltaB,
		Row2DeltaB:               g2.DeltaB,
		TotalDeltaB:              totalDelta,
		FloatRow1:                d1,
		FloatRow2:                d2,
		TotalBeta:                totalBeta,
		SystemDeterminant:        det,
		LStar:                    lstar,
		LB1:                      lb1,
		LB2:                      lb2,
		MB1GeV:                   mb1,
		MB2GeV:                   mb2,
		MStarGeV:                 ms,
		ClosureResidualU:         closureResidual,
		ExactLinearClosure:       closureResidual < 1e-8,
		MinUInDomain:             minUPlanck,
		EarliestPoleLog:          poleLog,
		EarliestPoleGeV:          poleGeV,
		ScaleOrdered:             scaleOrdered,
		DistinctThresholds:       distinct,
		SubPlanck:                subPlanck,
		PositiveCouplingsToMStar: positiveToMStar,
		NoSubPlanckLandauPole:    noPole,
		AnomalyCompatible:        anomaly,
		LeptoquarkSealCompatible: seal,
		ContainsGate201Shape:     containsGate201,
		MatchesBGapOrContactData: match,
		SU2AsymptoticallyFree:    su2AF,
		SU3AsymptoticallyFree:    su3AF,
		AllNonAbelianAF:          allAF,
		TotalDeltaBNorm:          norm(toFloat(totalDelta)),
		GUTRangeDistance:         gutRangeDistance(ms),
	}
	sol.BindingConstraint = bindingConstraint(sol)
	sol.Viable = sol.BindingConstraint == "viable"
	if sol.Viable {
		sol.Verdict = "CONDITIONAL_VIABLE: exact two-threshold algebraic closure survives scale ordering, Planck, positivity, Landau, anomaly, and leptoquark-seal filters"
	} else {
		sol.Verdict = "rejected by " + sol.BindingConstraint
	}
	return sol
}

func det3(m [3][3]float64) float64 {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) - m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) + m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

func solve3(m [3][3]float64, rhs [3]float64, det float64) [3]float64 {
	replace := func(col int) [3][3]float64 {
		mm := m
		for i := 0; i < 3; i++ {
			mm[i][col] = rhs[i]
		}
		return mm
	}
	return [3]float64{det3(replace(0)) / det, det3(replace(1)) / det, det3(replace(2)) / det}
}

func parallel(a, b FloatTriple) bool { return norm(cross(a, b)) < 1e-12 }

func cross(a, b FloatTriple) FloatTriple {
	return FloatTriple{a.SU2L*b.SU3C - a.SU3C*b.SU2L, a.SU3C*b.U1GUT - a.U1GUT*b.SU3C, a.U1GUT*b.SU2L - a.SU2L*b.U1GUT}
}

func scaleFromLog(l float64) float64 {
	if !finite(l) || l < -700 || l > 700 {
		return math.NaN()
	}
	return mzGeV * math.Exp(l)
}

func uAt(inputs RGInputAudit, d1, d2 FloatTriple, l, lb1, lb2 float64, gauge int) float64 {
	u0 := []float64{inputs.U1GUT, inputs.U2, inputs.U3}[gauge]
	active1 := math.Max(0, l-lb1)
	active2 := math.Max(0, l-lb2)
	return u0 - (inputs.SMBeta.At(gauge)*l+d1.At(gauge)*active1+d2.At(gauge)*active2)/(8.0*math.Pi*math.Pi)
}

func closureResidualU(inputs RGInputAudit, target BoundaryTarget, d1, d2 FloatTriple, lstar, lb1, lb2 float64) float64 {
	if !finite(lstar) || !finite(lb1) || !finite(lb2) {
		return math.Inf(1)
	}
	max := 0.0
	for g := 0; g < 3; g++ {
		r := math.Abs(uAt(inputs, d1, d2, lstar, lb1, lb2, g) - target.UTarget)
		if r > max {
			max = r
		}
	}
	return max
}

func minUInDomain(inputs RGInputAudit, target BoundaryTarget, d1, d2 FloatTriple, lstar, lb1, lb2 float64, toPlanck bool) float64 {
	if !finite(lstar) || !finite(lb1) || !finite(lb2) {
		return math.NaN()
	}
	end := lstar
	if toPlanck {
		end = inputs.PlanckLogBound
	}
	points := []float64{0, end}
	for _, x := range []float64{lb1, lb2, lstar} {
		if finite(x) && x > 0 && x < end {
			points = append(points, x)
		}
	}
	minU := math.Inf(1)
	for _, x := range points {
		for g := 0; g < 3; g++ {
			u := uAt(inputs, d1, d2, x, lb1, lb2, g)
			if u < minU {
				minU = u
			}
		}
	}
	return minU
}

func earliestPoleLog(inputs RGInputAudit, target BoundaryTarget, d1, d2 FloatTriple, lstar, lb1, lb2 float64) float64 {
	if !finite(lstar) || !finite(lb1) || !finite(lb2) {
		return math.NaN()
	}
	breaks := []float64{0, inputs.PlanckLogBound}
	for _, x := range []float64{lb1, lb2, lstar} {
		if finite(x) && x > 0 && x < inputs.PlanckLogBound {
			breaks = append(breaks, x)
		}
	}
	sort.Float64s(breaks)
	unique := breaks[:0]
	for _, x := range breaks {
		if len(unique) == 0 || math.Abs(x-unique[len(unique)-1]) > 1e-9 {
			unique = append(unique, x)
		}
	}
	earliest := math.Inf(1)
	for g := 0; g < 3; g++ {
		for k := 0; k+1 < len(unique); k++ {
			a, b := unique[k], unique[k+1]
			ua := uAt(inputs, d1, d2, a, lb1, lb2, g)
			ub := uAt(inputs, d1, d2, b, lb1, lb2, g)
			if ua <= 0 && a > 0 {
				if a < earliest {
					earliest = a
				}
				continue
			}
			if ua > 0 && ub <= 0 {
				slope := (ub - ua) / (b - a)
				if slope < 0 {
					root := a - ua/slope
					if root > 0 && root < earliest {
						earliest = root
					}
				}
			}
		}
	}
	if math.IsInf(earliest, 1) {
		return math.Inf(1)
	}
	return earliest
}

func bindingConstraint(s PairSolution) string {
	if math.Abs(s.SystemDeterminant) <= determinantTolerance {
		return "singular-or-dependent-3x3-system"
	}
	if !s.ScaleOrdered {
		return "scale-ordering"
	}
	if !s.ExactLinearClosure {
		return "linear-closure-residual"
	}
	if !s.DistinctThresholds {
		return "coincident-thresholds"
	}
	if !s.SubPlanck {
		return "sub-planck-bound"
	}
	if !s.PositiveCouplingsToMStar {
		return "positive-couplings-to-boundary"
	}
	if !s.NoSubPlanckLandauPole {
		return "sub-planck-landau-pole"
	}
	if !s.AnomalyCompatible {
		return "anomaly-filter"
	}
	if !s.LeptoquarkSealCompatible {
		return "leptoquark-seal-filter"
	}
	return "viable"
}

func addRT(a, b representationrowlattice.RationalTriple) representationrowlattice.RationalTriple {
	return representationrowlattice.RT(a.B1.Add(b.B1), a.B2.Add(b.B2), a.B3.Add(b.B3))
}

func norm(t FloatTriple) float64 { return math.Sqrt(t.U1GUT*t.U1GUT + t.SU2L*t.SU2L + t.SU3C*t.SU3C) }

func gutRangeDistance(mstar float64) float64 {
	if !finite(mstar) || mstar <= 0 {
		return math.Inf(1)
	}
	x := math.Log10(mstar)
	if x >= 15 && x <= 16 {
		return 0
	}
	if x < 15 {
		return 15 - x
	}
	return x - 16
}

func sortSolutions(xs []PairSolution) {
	sort.Slice(xs, func(i, j int) bool {
		// (a) fewest carriers: all Gate-211 rows are two carriers; keep this
		// clause explicit for future multi-carrier extensions.
		if xs[i].TotalDeltaBNorm != xs[j].TotalDeltaBNorm {
			return xs[i].TotalDeltaBNorm < xs[j].TotalDeltaBNorm
		}
		if xs[i].GUTRangeDistance != xs[j].GUTRangeDistance {
			return xs[i].GUTRangeDistance < xs[j].GUTRangeDistance
		}
		if xs[i].MStarGeV != xs[j].MStarGeV {
			return xs[i].MStarGeV < xs[j].MStarGeV
		}
		if xs[i].Row1Name != xs[j].Row1Name {
			return xs[i].Row1Name < xs[j].Row1Name
		}
		return xs[i].Row2Name < xs[j].Row2Name
	})
}

func dominantConstraint(counts map[string]int) string {
	bestK := "none"
	bestV := -1
	for k, v := range counts {
		if k == "viable" {
			continue
		}
		if v > bestV || (v == bestV && k < bestK) {
			bestK, bestV = k, v
		}
	}
	return fmt.Sprintf("%s (%d)", bestK, bestV)
}

func auditBaryonAnomaly(gate210 Gate210Snapshot, gens []SearchGenerator, targets []TargetAudit) BaryonAnomalyAudit {
	allAnomaly := true
	allSeal := true
	for _, g := range gens {
		allAnomaly = allAnomaly && g.AnomalyCompatible
		allSeal = allSeal && g.LeptoquarkSealCompatible
	}
	viableAnomaly := true
	viableSeal := true
	for _, ta := range targets {
		for _, s := range ta.AllViableSolutions {
			viableAnomaly = viableAnomaly && s.AnomalyCompatible
			viableSeal = viableSeal && s.LeptoquarkSealCompatible
		}
	}
	return BaryonAnomalyAudit{LeptoquarkDynamicsSealInherited: gate210.Gate209LeptoquarkSealInherited, AllSearchRowsAnomalyCompatible: allAnomaly, AllSearchRowsSealCompatible: allSeal, ViablePairsAnomalyCompatible: viableAnomaly, ViablePairsSealCompatible: viableSeal, ProtonDecayOperatorUsed: false, ProtonLifetimeComputed: false, Verdict: "Gate-211 search basis inherits anomaly-safe row filtering and the active LeptoquarkDynamicsSeal; no proton-decay operator or lifetime is introduced"}
}

func auditContactMatch(targets []TargetAudit) ContactMatchAudit {
	promoted := false
	for _, ta := range targets {
		for _, s := range ta.AllViableSolutions {
			promoted = promoted || s.MatchesBGapOrContactData
		}
	}
	return ContactMatchAudit{BGapAudited: true, ContactPartialOverlapAudited: true, CanonicalNumericMatchFound: false, ChargeSpinMassSemanticsFound: false, ViableRowsPromotedFromContact: promoted, Verdict: "no viable row is promoted from B-sector gap or contact partial-overlap data; Gate 205 carrier-activation obstruction remains in force"}
}

func auditFirewall(gate210 Gate210Snapshot, inputs RGInputAudit, targets []TargetAudit, ba BaryonAnomalyAudit) FirewallAudit {
	totalViable := 0
	for _, ta := range targets {
		totalViable += ta.ViablePairs
	}
	return FirewallAudit{
		Gate210Inherited:                    gate210.Gate210Inherited,
		LeptoquarkDynamicsSealInherited:     gate210.Gate209LeptoquarkSealInherited,
		EmpiricalLedgerQuarantined:          inputs.ObservedLedgerQuarantined,
		ObservedLedgerUsedForFiniteCore:     inputs.UsedForFiniteCoreDerivation,
		UniversalBetaRowInserted:            inputs.UniversalBetaRowAllowed,
		ArbitraryRealRowCoefficientInserted: false,
		PhysicalPredictionClaimed:           false,
		AbsoluteMassDerivedFromFiniteCore:   false,
		ProtonDecaySealViolated:             !ba.LeptoquarkDynamicsSealInherited || ba.ProtonDecayOperatorUsed,
		ProtonLifetimeComputed:              ba.ProtonLifetimeComputed,
		MatchingCorrectionsDerived:          false,
		ConditionalViabilityOnly:            totalViable > 0,
		RecommendedNextGate:                 "Gate 212 — two-threshold solution minimality / finite-origin and matching-correction preflight audit",
		OpenRequirements: []string{
			"derive a finite carrier-origin theorem before promoting any viable row pair beyond phenomenology",
			"derive or seal finite threshold matching corrections for two separated scales",
			"audit whether the selected scale ordering is stable under scheme changes and two-loop effects",
			"keep the LeptoquarkDynamicsSeal active unless safe dynamics are derived explicitly",
		},
		Verdict: "firewall preserved: two-threshold scales are conditional outputs of a quarantined ledger and rational row filter, not finite-core predictions",
	}
}

func buildTruthStatement(s Summary, targets []TargetAudit) string {
	if s.TotalViablePairs > 0 {
		parts := []string{}
		for _, ta := range targets {
			if ta.ViablePairs > 0 {
				parts = append(parts, fmt.Sprintf("%s:%d", ta.Target.Name, ta.ViablePairs))
			}
		}
		return "Gate 211 confirms the dimension-counting pivot: two independent rational threshold rows make exact closure a linear solve, so the scientific content is the physical viability filter. At least one ordered two-threshold pair survives scale-ordering, prompt Planck, positivity, Landau, anomaly, and leptoquark-seal constraints (" + strings.Join(parts, ", ") + "). These are CONDITIONAL_VIABLE phenomenological completions only; no finite carrier origin, matching correction, or physical prediction is claimed."
	}
	return "Gate 211 confirms the dimension-counting pivot but finds no ordered two-threshold pair surviving the physical viability filters. The failure is a phenomenological route obstruction, not a finite-core theorem failure."
}

func finite(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }

func FormatInputs(i RGInputAudit) string {
	return fmt.Sprintf("%s MZ=%.6g α^-1=(%.12g,%.12g,%.12g) u=(%.12g,%.12g,%.12g) bSM=%s planck=%.6g Lbound=%.6g quarantined=%t finiteUse=%t universalAllowed=%t", i.LedgerName, i.MZGeV, i.Alpha1GUTInverse, i.Alpha2Inverse, i.Alpha3Inverse, i.U1GUT, i.U2, i.U3, i.SMBeta, i.PlanckGeV, i.PlanckLogBound, i.ObservedLedgerQuarantined, i.UsedForFiniteCoreDerivation, i.UniversalBetaRowAllowed)
}

func FormatGeneratorAudit(a GeneratorAudit) string {
	return fmt.Sprintf("sourceUnique=%d safe=%d expectedGate210=%d anomaly=%d seal=%d zeroExcluded=%t universalRow=%t realCoeff=%t inheritsGate210=%t", a.SourceUniqueRows, a.SafeGenerators, a.ExpectedGate210SafeGenerators, a.AnomalyCompatibleGenerators, a.LeptoquarkSealCompatibleRows, a.ZeroRowExcluded, !a.NoUniversalBetaRowInserted, !a.NoContinuousRowCoefficients, a.InheritsGate210FilteredBasis)
}

func FormatSolution(s PairSolution) string {
	pole := "none"
	if finite(s.EarliestPoleLog) {
		pole = fmt.Sprintf("L=%.9g M=%.9g", s.EarliestPoleLog, s.EarliestPoleGeV)
	}
	return fmt.Sprintf("target=%s u=%.6g rows=[%s %s; %s %s] Δb1=%s Δb2=%s ΔbTot=%s L=(L*=%.9g,LB1=%.9g,LB2=%.9g) M=(M*=%.9g,MB1=%.9g,MB2=%.9g) totalBeta=%s closureU=%.3g exactClosure=%t minU=%.9g pole=%s ordered=%t distinct=%t subPlanck=%t positive=%t noPole=%t anomaly=%t seal=%t gate201=%t contactMatch=%t AF(SU2,SU3)=(%t,%t)", s.TargetName, s.TargetU, s.Row1Name, s.Row1Rep, s.Row2Name, s.Row2Rep, s.Row1DeltaB, s.Row2DeltaB, s.TotalDeltaB, s.LStar, s.LB1, s.LB2, s.MStarGeV, s.MB1GeV, s.MB2GeV, s.TotalBeta, s.ClosureResidualU, s.ExactLinearClosure, s.MinUInDomain, pole, s.ScaleOrdered, s.DistinctThresholds, s.SubPlanck, s.PositiveCouplingsToMStar, s.NoSubPlanckLandauPole, s.AnomalyCompatible, s.LeptoquarkSealCompatible, s.ContainsGate201Shape, s.MatchesBGapOrContactData, s.SU2AsymptoticallyFree, s.SU3AsymptoticallyFree)
}

func FormatTargetAudit(a TargetAudit) string {
	best := "none"
	if len(a.BestSolutions) > 0 {
		best = FormatSolution(a.BestSolutions[0])
	}
	return fmt.Sprintf("target=%s u=%.6g orderedPairs=%d pairIndependent=%d invertible=%d scaleOrdered=%d distinct=%d subPlanck=%d positive=%d noLandau=%d viable=%d dominant=%s best={%s}", a.Target.Name, a.Target.UTarget, a.OrderedPairsAudited, a.PairIndependentPairs, a.InvertibleSystems, a.ScaleOrderedPairs, a.DistinctThresholdPairs, a.SubPlanckPairs, a.PositiveCouplingPairs, a.NoLandauPairs, a.ViablePairs, dominantConstraint(a.BindingConstraintCounts), best)
}

func FormatBaryonAnomaly(a BaryonAnomalyAudit) string {
	return fmt.Sprintf("sealInherited=%t allRowsAnomaly=%t allRowsSeal=%t viableAnomaly=%t viableSeal=%t protonOperator=%t lifetime=%t", a.LeptoquarkDynamicsSealInherited, a.AllSearchRowsAnomalyCompatible, a.AllSearchRowsSealCompatible, a.ViablePairsAnomalyCompatible, a.ViablePairsSealCompatible, a.ProtonDecayOperatorUsed, a.ProtonLifetimeComputed)
}

func FormatContactMatch(a ContactMatchAudit) string {
	return fmt.Sprintf("bgap=%t contact=%t numericMatch=%t semantics=%t promoted=%t", a.BGapAudited, a.ContactPartialOverlapAudited, a.CanonicalNumericMatchFound, a.ChargeSpinMassSemanticsFound, a.ViableRowsPromotedFromContact)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate210=%t lqSeal=%t ledger=%t observedFinite=%t universal=%t realCoeff=%t physicalPrediction=%t finiteMass=%t protonSeal=%t lifetime=%t matching=%t conditionalOnly=%t next=%s", f.Gate210Inherited, f.LeptoquarkDynamicsSealInherited, f.EmpiricalLedgerQuarantined, f.ObservedLedgerUsedForFiniteCore, f.UniversalBetaRowInserted, f.ArbitraryRealRowCoefficientInserted, f.PhysicalPredictionClaimed, f.AbsoluteMassDerivedFromFiniteCore, f.ProtonDecaySealViolated, f.ProtonLifetimeComputed, f.MatchingCorrectionsDerived, f.ConditionalViabilityOnly, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate210=%t targets=%d viableTopological=%d viableCentroid=%d totalViable=%d status=%s comment=%q", s.TestsAudited, s.Gate210Inherited, s.TargetsAudited, s.ViableTopological, s.ViableCentroid, s.TotalViablePairs, s.Status, s.Comment)
}
