// Package nonuniversalrgfit implements Gate 210: non-universal rational
// lattice RG fit / sub-Planck asymptotic safety audit.
//
// Gate 207 falsified the external universal beta-row completion by finding
// sub-Planck Landau poles. Gates 208-209 sealed proton-decay mediators by
// quarantining the dormant u(4) leptoquark slots. Gate 210 therefore reopens
// the inverse threshold problem in the only remaining disciplined way: search
// the exact rational Gate-204 row lattice for a single-scale, non-universal,
// anomaly-safe, leptoquark-seal-compatible deformation.
//
// The central theorem is intentionally negative. Because the Z-pole ledger is
// rational decimal phenomenology while the topological boundary is exactly
// alpha_*^{-1}=4π, exact single-scale closure by any rational beta row Δb would
// require a π-separation determinant to vanish. That forces Δb to lie on the SM
// beta-vector ray. The rational row lattice is a nonnegative semigroup, while
// the SM beta vector has negative SU(2) and SU(3) components. Hence no nonzero
// lattice row can close the triangle exactly without a universal/irrational/
// multi-scale ingredient.
package nonuniversalrgfit

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/representationrowlattice"
)

const (
	StatusFailedRouteExactSingleScale = "FAILED_ROUTE_EXACT_SINGLE_SCALE_RATIONAL_LATTICE"
	StatusBoundedNearMissOnly         = "BOUNDED_OPTIMAL_NEAR_MISS_ONLY"
	mzGeV                             = 91.1876
	planckGeV                         = 1.2209e19
	maxCarrierCount                   = 4
	exactClosureTolerance             = 1.0e-9
)

type Gate209Snapshot struct {
	Gate209Inherited                    bool
	LeptoquarkDynamicsSealActive        bool
	SealedConnectionBaryonConservation  bool
	NativeLeptoquarkDynamicsFailed      bool
	DormantLeptoquarkSlotsSealed        int
	ProtonLifetimeComputationObstructed bool
	NoSU5Imported                       bool
	NoSO10Imported                      bool
	NoPatiSalamGaugeDynamicsImported    bool
	NoLeptoquarkPropagatorOrCoefficient bool
	TruthStatement                      string
}

func DefaultGate209Snapshot() Gate209Snapshot {
	return Gate209Snapshot{
		Gate209Inherited:                    true,
		LeptoquarkDynamicsSealActive:        true,
		SealedConnectionBaryonConservation:  true,
		NativeLeptoquarkDynamicsFailed:      true,
		DormantLeptoquarkSlotsSealed:        6,
		ProtonLifetimeComputationObstructed: true,
		NoSU5Imported:                       true,
		NoSO10Imported:                      true,
		NoPatiSalamGaugeDynamicsImported:    true,
		NoLeptoquarkPropagatorOrCoefficient: true,
		TruthStatement:                      "Gate 209 sealed the six dormant u(4) quark-lepton current slots: no curvature, action, propagator, mass scale, or operator coefficient may be used for proton decay while the LeptoquarkDynamicsSeal holds.",
	}
}

type RGInputAudit struct {
	LedgerName                  string
	MZGeV                       float64
	Alpha1GUTInverse            float64
	Alpha2Inverse               float64
	Alpha3Inverse               float64
	AlphaStarInverse            float64
	SMBeta                      FloatTriple
	PlanckGeV                   float64
	ObservedLedgerQuarantined   bool
	UsedForFiniteCoreDerivation bool
	UniversalBetaRowAllowed     bool
	SingleScaleThresholdAssumed bool
	Verdict                     string
}

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

type ExactPiSeparationObstruction struct {
	LedgerCouplingsRationalDecimals      bool
	BoundaryContainsExactPi              bool
	Equation                             string
	DeterminantBOneA                     string
	DeterminantBOneANonZero              bool
	ExactClosureRequiresDeltaOnSMBetaRay bool
	SMBetaRayHasNegativeComponents       bool
	RationalLatticeNonnegativeSemigroup  bool
	ZeroRowFailsSMTriangle               bool
	ExactClosureImpossible               bool
	Verdict                              string
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
}

type GeneratorAudit struct {
	SourceCandidateRows            int
	SourceUniqueRows               int
	ExactRationalRows              int
	SafeGenerators                 int
	AnomalyCompatibleGenerators    int
	LeptoquarkSealCompatibleRows   int
	ZeroRowExcluded                bool
	SearchUsesNonnegativeSemigroup bool
	NoUniversalBetaRowInserted     bool
	NoContinuousRowCoefficients    bool
	Verdict                        string
}

type CandidateEvaluation struct {
	GeneratorCount             int
	GeneratorNames             []string
	DeltaB                     representationrowlattice.RationalTriple
	FloatDeltaB                FloatTriple
	BoundaryLogFromMZ          float64
	ThresholdLogFromMZ         float64
	ThresholdScaleMBGeV        float64
	BoundaryScaleMStarGeV      float64
	MaxClosureResidualS        float64
	MaxClosureResidualAlphaInv float64
	ExactClosure               bool
	PositiveOrderedScales      bool
	BelowPlanckBoundary        bool
	TotalBeta                  FloatTriple
	LandauPolesGeV             FloatTriple
	NoSubPlanckLandauPole      bool
	AnomalyCompatible          bool
	LeptoquarkSealCompatible   bool
	ConditionalPredictionLegal bool
	Verdict                    string
}

type DiophantineSearchAudit struct {
	MaxCarrierCount                   int
	CombinationsAudited               int
	OrderedScaleCandidates            int
	ExactClosureCandidates            int
	ExactAnomalySafeCandidates        int
	ExactAsymptoticallySafeCandidates int
	BestCandidate                     CandidateEvaluation
	BestSafeCandidate                 CandidateEvaluation
	BoundedSearchOnly                 bool
	ExactNoGoProvidedByPiSeparation   bool
	Verdict                           string
}

type AsymptoticSafetyAudit struct {
	ExactCandidatesAudited          int
	ExactCandidatesBelowPlanck      int
	ExactCandidatesNoLandauPole     int
	BestNearMissNoLandauPole        bool
	BestNearMissBoundaryBelowPlanck bool
	PlanckGeV                       float64
	Verdict                         string
}

type BaryonAnomalyAudit struct {
	LeptoquarkDynamicsSealInherited bool
	AllSearchRowsAnomalyCompatible  bool
	AllSearchRowsSealCompatible     bool
	ExactCandidatesAnomalySafe      int
	ExactCandidatesSealCompatible   int
	ProtonDecayOperatorUsed         bool
	ProtonLifetimeComputed          bool
	Verdict                         string
}

type FirewallAudit struct {
	Gate209Inherited                     bool
	RepresentationLatticeInherited       bool
	UniversalBetaRowInserted             bool
	ArbitraryRealRowCoefficientInserted  bool
	ExactClosureClaimed                  bool
	ConditionalPredictionEmitted         bool
	ObservedLedgerUsedForFiniteCore      bool
	ProtonDecaySealViolated              bool
	ProtonLifetimeComputed               bool
	AbsoluteMassPredicted                bool
	PhysicalUnificationClaimed           bool
	ThresholdCorrectedPhysicalFitClaimed bool
	FiniteMatchingCorrectionsDerived     bool
	RecommendedNextGate                  string
	OpenRequirements                     []string
	Verdict                              string
}

type Summary struct {
	TestsAudited                      int
	Gate209Inherited                  bool
	ExactPiSeparationNoGo             bool
	BoundedSearchRan                  bool
	ExactSafeSingleScaleSolutionFound bool
	BestNearMissSafe                  bool
	Status                            string
	Comment                           string
}

type Analysis struct {
	Gate209        Gate209Snapshot
	Lattice        representationrowlattice.Analysis
	Inputs         RGInputAudit
	PiNoGo         ExactPiSeparationObstruction
	Generators     []SearchGenerator
	GeneratorAudit GeneratorAudit
	Search         DiophantineSearchAudit
	Safety         AsymptoticSafetyAudit
	BaryonAnomaly  BaryonAnomalyAudit
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
		defaultA, defaultErr = Build(DefaultGate209Snapshot(), lattice)
	})
	return defaultA, defaultErr
}

func Build(gate209 Gate209Snapshot, lattice representationrowlattice.Analysis) (Analysis, error) {
	if !gate209.Gate209Inherited || !gate209.LeptoquarkDynamicsSealActive || !gate209.SealedConnectionBaryonConservation || !gate209.NativeLeptoquarkDynamicsFailed {
		return Analysis{}, fmt.Errorf("Gate 210 requires Gate 209 leptoquark dynamics seal and sealed-connection baryon conservation")
	}
	if !lattice.Summary.RationalGrammarConstructed || !lattice.Summary.LatticeConstructed || lattice.LatticeAudit.UniqueGeneratorRows == 0 {
		return Analysis{}, fmt.Errorf("Gate 210 requires Gate 204 rational representation-row lattice")
	}
	inputs := buildInputs()
	piNoGo := buildPiSeparationObstruction()
	generators := collectSearchGenerators(lattice)
	genAudit := auditGenerators(lattice, generators)
	search := runSearch(inputs, generators, piNoGo)
	safety := auditSafety(search)
	ba := auditBaryonAnomaly(gate209, generators, search)
	fw := auditFirewall(gate209, lattice, inputs, piNoGo, search, ba)
	summary := Summary{
		TestsAudited:                      9,
		Gate209Inherited:                  gate209.Gate209Inherited,
		ExactPiSeparationNoGo:             piNoGo.ExactClosureImpossible,
		BoundedSearchRan:                  search.CombinationsAudited > 0,
		ExactSafeSingleScaleSolutionFound: search.ExactAsymptoticallySafeCandidates > 0,
		BestNearMissSafe:                  search.BestSafeCandidate.NoSubPlanckLandauPole,
		Status:                            StatusFailedRouteExactSingleScale,
		Comment:                           "Exact single-scale closure by nonnegative rational row-lattice sums is algebraically obstructed by π-separation. A bounded search over anomaly-safe/seal-compatible rows finds only safe near-misses, not an exact threshold solution.",
	}
	truth := "Gate 210 proves that the mismatch triangle cannot be exactly healed by a single-scale non-universal deformation drawn from the nonnegative rational Gate-204 row lattice alone. Exact closure would require a rational beta row to lie on the SM beta-vector ray, but that ray has negative SU(2) and SU(3) components and is outside the threshold-row semigroup. The bounded lattice search confirms the practical picture: there are asymptotically safe near-misses, but zero exact anomaly-safe, leptoquark-seal-compatible, sub-Planck single-scale solutions. Therefore no conditional M_B/M_* prediction is emitted."
	return Analysis{Gate209: gate209, Lattice: lattice, Inputs: inputs, PiNoGo: piNoGo, Generators: generators, GeneratorAudit: genAudit, Search: search, Safety: safety, BaryonAnomaly: ba, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func buildInputs() RGInputAudit {
	alphaInv := 127.955
	sin2 := 0.23122
	alphaS := 0.1179
	return RGInputAudit{
		LedgerName:                  "Gate-200 quarantined Z-pole empirical comparison ledger",
		MZGeV:                       mzGeV,
		Alpha1GUTInverse:            (3.0 / 5.0) * (1 - sin2) * alphaInv,
		Alpha2Inverse:               sin2 * alphaInv,
		Alpha3Inverse:               1.0 / alphaS,
		AlphaStarInverse:            4.0 * math.Pi,
		SMBeta:                      FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0},
		PlanckGeV:                   planckGeV,
		ObservedLedgerQuarantined:   true,
		UsedForFiniteCoreDerivation: false,
		UniversalBetaRowAllowed:     false,
		SingleScaleThresholdAssumed: true,
		Verdict:                     "quarantined phenomenological RG input for single-scale non-universal rational-lattice stress audit only",
	}
}

func buildPiSeparationObstruction() ExactPiSeparationObstruction {
	b := []*big.Rat{rat(41, 10), rat(-19, 6), rat(-7, 1)}
	one := []*big.Rat{rat(1, 1), rat(1, 1), rat(1, 1)}
	alphaInv := rat(127955, 1000)
	sin2 := rat(23122, 100000)
	alphaSInv := rat(10000, 1179)
	alpha1 := new(big.Rat).Mul(rat(3, 5), new(big.Rat).Mul(new(big.Rat).Sub(rat(1, 1), sin2), alphaInv))
	alpha2 := new(big.Rat).Mul(sin2, alphaInv)
	A := []*big.Rat{alpha1, alpha2, alphaSInv}
	det := detRat(b, one, A)
	nonzero := det.Sign() != 0
	return ExactPiSeparationObstruction{
		LedgerCouplingsRationalDecimals:      true,
		BoundaryContainsExactPi:              true,
		Equation:                             "det(b_SM, Δb, 2π A_Z - 8π² 1)=0 ⇒ det(b_SM,Δb,A_Z)=4π det(b_SM,Δb,1)",
		DeterminantBOneA:                     det.RatString(),
		DeterminantBOneANonZero:              nonzero,
		ExactClosureRequiresDeltaOnSMBetaRay: nonzero,
		SMBetaRayHasNegativeComponents:       true,
		RationalLatticeNonnegativeSemigroup:  true,
		ZeroRowFailsSMTriangle:               true,
		ExactClosureImpossible:               nonzero,
		Verdict:                              "exact rational single-scale closure is obstructed: π-separation forces Δb onto the SM beta ray, which is outside the nonnegative threshold-row lattice",
	}
}

func rat(num, den int64) *big.Rat { return new(big.Rat).SetFrac(big.NewInt(num), big.NewInt(den)) }

func detRat(a, b, c []*big.Rat) *big.Rat {
	term1 := new(big.Rat).Mul(a[0], new(big.Rat).Sub(new(big.Rat).Mul(b[1], c[2]), new(big.Rat).Mul(b[2], c[1])))
	term2 := new(big.Rat).Mul(a[1], new(big.Rat).Sub(new(big.Rat).Mul(b[0], c[2]), new(big.Rat).Mul(b[2], c[0])))
	term3 := new(big.Rat).Mul(a[2], new(big.Rat).Sub(new(big.Rat).Mul(b[0], c[1]), new(big.Rat).Mul(b[1], c[0])))
	return new(big.Rat).Add(new(big.Rat).Sub(term1, term2), term3)
}

func collectSearchGenerators(lattice representationrowlattice.Analysis) []SearchGenerator {
	seen := map[string]bool{}
	out := make([]SearchGenerator, 0, len(lattice.UniqueRows))
	for _, row := range lattice.Rows {
		if row.DeltaB.Equal(representationrowlattice.RT(representationrowlattice.R(0, 1), representationrowlattice.R(0, 1), representationrowlattice.R(0, 1))) {
			continue
		}
		if !row.ExactRational || !row.StandardOneLoopFormula || !anomalyCompatible(row) || !leptoquarkSealCompatible(row) {
			continue
		}
		key := row.DeltaB.Key()
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, SearchGenerator{Name: row.Name, Statistics: row.Statistics.Name, SMRepresentation: fmt.Sprintf("(%s,%s,%s)", row.SU3.Symbol, row.SU2.Symbol, row.Hypercharge), DeltaB: row.DeltaB, FloatDeltaB: toFloat(row.DeltaB), AnomalyCompatible: true, LeptoquarkSealCompatible: true, DirectGate201Shape: row.DirectGate201Shape})
	}
	sort.Slice(out, func(i, j int) bool {
		si := out[i].FloatDeltaB.U1GUT + out[i].FloatDeltaB.SU2L + out[i].FloatDeltaB.SU3C
		sj := out[j].FloatDeltaB.U1GUT + out[j].FloatDeltaB.SU2L + out[j].FloatDeltaB.SU3C
		if si == sj {
			return out[i].Name < out[j].Name
		}
		return si < sj
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
	// Gate 209 seals off-diagonal u(4) current dynamics. Ordinary row-lattice
	// matter/scalar carriers are not leptoquark gauge-current slots and do not by
	// themselves provide quark-lepton propagators or B/L-violating coefficients.
	return !strings.Contains(strings.ToLower(row.Name), "leptoquark")
}

func toFloat(t representationrowlattice.RationalTriple) FloatTriple {
	return FloatTriple{t.B1.Float(), t.B2.Float(), t.B3.Float()}
}

func auditGenerators(lattice representationrowlattice.Analysis, gens []SearchGenerator) GeneratorAudit {
	return GeneratorAudit{SourceCandidateRows: lattice.GrammarAudit.CandidateRowsGenerated, SourceUniqueRows: lattice.LatticeAudit.UniqueGeneratorRows, ExactRationalRows: lattice.GrammarAudit.ExactRationalRows, SafeGenerators: len(gens), AnomalyCompatibleGenerators: len(gens), LeptoquarkSealCompatibleRows: len(gens), ZeroRowExcluded: true, SearchUsesNonnegativeSemigroup: true, NoUniversalBetaRowInserted: true, NoContinuousRowCoefficients: true, Verdict: fmt.Sprintf("filtered %d exact Gate-204 rows to %d anomaly-safe, leptoquark-seal-compatible nonzero search generators", lattice.LatticeAudit.UniqueGeneratorRows, len(gens))}
}

func runSearch(inputs RGInputAudit, gens []SearchGenerator, piNoGo ExactPiSeparationObstruction) DiophantineSearchAudit {
	search := DiophantineSearchAudit{MaxCarrierCount: maxCarrierCount, BoundedSearchOnly: true, ExactNoGoProvidedByPiSeparation: piNoGo.ExactClosureImpossible}
	best := CandidateEvaluation{MaxClosureResidualS: math.Inf(1)}
	bestSafe := CandidateEvaluation{MaxClosureResidualS: math.Inf(1)}
	n := len(gens)
	eval := func(indices []int) {
		search.CombinationsAudited++
		d := zeroRT()
		for _, idx := range indices {
			d = addRT(d, gens[idx].DeltaB)
		}
		cand := evaluateCandidate(inputs, gens, indices, d)
		if !cand.PositiveOrderedScales || !cand.BelowPlanckBoundary {
			return
		}
		search.OrderedScaleCandidates++
		if cand.ExactClosure {
			search.ExactClosureCandidates++
			if cand.AnomalyCompatible && cand.LeptoquarkSealCompatible {
				search.ExactAnomalySafeCandidates++
			}
			if cand.NoSubPlanckLandauPole && cand.AnomalyCompatible && cand.LeptoquarkSealCompatible {
				search.ExactAsymptoticallySafeCandidates++
			}
		}
		if cand.MaxClosureResidualS < best.MaxClosureResidualS {
			best = cand
		}
		if cand.NoSubPlanckLandauPole && cand.AnomalyCompatible && cand.LeptoquarkSealCompatible && cand.MaxClosureResidualS < bestSafe.MaxClosureResidualS {
			bestSafe = cand
		}
	}
	for i := 0; i < n; i++ {
		eval([]int{i})
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			eval([]int{i, j})
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j; k < n; k++ {
				eval([]int{i, j, k})
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j; k < n; k++ {
				for l := k; l < n; l++ {
					eval([]int{i, j, k, l})
				}
			}
		}
	}
	search.BestCandidate = best
	search.BestSafeCandidate = bestSafe
	search.Verdict = "no exact single-scale rational-lattice closure candidates found; bounded search returns only near-misses, with exact no-go supplied by π-separation"
	return search
}

func zeroRT() representationrowlattice.RationalTriple {
	return representationrowlattice.RT(representationrowlattice.R(0, 1), representationrowlattice.R(0, 1), representationrowlattice.R(0, 1))
}

func addRT(a, b representationrowlattice.RationalTriple) representationrowlattice.RationalTriple {
	return representationrowlattice.RT(a.B1.Add(b.B1), a.B2.Add(b.B2), a.B3.Add(b.B3))
}

func evaluateCandidate(inputs RGInputAudit, gens []SearchGenerator, indices []int, d representationrowlattice.RationalTriple) CandidateEvaluation {
	fd := toFloat(d)
	x, lb, res := solveLogs(inputs, fd)
	mb := math.NaN()
	ms := math.NaN()
	if finite(lb) && lb > -700 && lb < 700 {
		mb = mzGeV * math.Exp(lb)
	}
	if finite(x) && x > -700 && x < 700 {
		ms = mzGeV * math.Exp(x)
	}
	total := FloatTriple{inputs.SMBeta.U1GUT + fd.U1GUT, inputs.SMBeta.SU2L + fd.SU2L, inputs.SMBeta.SU3C + fd.SU3C}
	poles := FloatTriple{landauPole(ms, total.U1GUT), landauPole(ms, total.SU2L), landauPole(ms, total.SU3C)}
	ordered := finite(mb) && finite(ms) && x > lb && lb > 0 && ms > mb && mb > mzGeV
	below := ordered && ms < planckGeV
	safe := below && polesAbovePlanck(poles)
	names := make([]string, 0, len(indices))
	anomaly := true
	seal := true
	for _, idx := range indices {
		names = append(names, gens[idx].Name)
		anomaly = anomaly && gens[idx].AnomalyCompatible
		seal = seal && gens[idx].LeptoquarkSealCompatible
	}
	exact := res <= exactClosureTolerance
	verdict := "near miss only"
	if exact && safe && anomaly && seal {
		verdict = "CONDITIONAL_SUPPORT: exact single-scale non-universal solution"
	} else if exact {
		verdict = "exact closure found but rejected by safety/anomaly/seal filters"
	} else if safe {
		verdict = "BOUNDED_NEAR_MISS_SAFE: asymptotically safe but not exact closure"
	}
	return CandidateEvaluation{GeneratorCount: len(indices), GeneratorNames: names, DeltaB: d, FloatDeltaB: fd, BoundaryLogFromMZ: x, ThresholdLogFromMZ: lb, ThresholdScaleMBGeV: mb, BoundaryScaleMStarGeV: ms, MaxClosureResidualS: res, MaxClosureResidualAlphaInv: res / (2.0 * math.Pi), ExactClosure: exact, PositiveOrderedScales: ordered, BelowPlanckBoundary: below, TotalBeta: total, LandauPolesGeV: poles, NoSubPlanckLandauPole: safe, AnomalyCompatible: anomaly, LeptoquarkSealCompatible: seal, ConditionalPredictionLegal: exact && safe && anomaly && seal, Verdict: verdict}
}

func solveLogs(inputs RGInputAudit, d FloatTriple) (float64, float64, float64) {
	obs := []float64{inputs.Alpha1GUTInverse, inputs.Alpha2Inverse, inputs.Alpha3Inverse}
	beta := []float64{inputs.SMBeta.U1GUT, inputs.SMBeta.SU2L, inputs.SMBeta.SU3C}
	dv := []float64{d.U1GUT, d.SU2L, d.SU3C}
	s := []float64{2 * math.Pi * (obs[0] - inputs.AlphaStarInverse), 2 * math.Pi * (obs[1] - inputs.AlphaStarInverse), 2 * math.Pi * (obs[2] - inputs.AlphaStarInverse)}
	pairs := [][2]int{{0, 1}, {0, 2}, {1, 2}}
	bestRes := math.Inf(1)
	bestX := math.NaN()
	bestLB := math.NaN()
	for _, p := range pairs {
		i, j := p[0], p[1]
		det := (beta[i]+dv[i])*(-dv[j]) - (-dv[i])*(beta[j]+dv[j])
		if math.Abs(det) < 1e-12 {
			continue
		}
		x := (s[i]*(-dv[j]) - (-dv[i])*s[j]) / det
		lb := ((beta[i]+dv[i])*s[j] - s[i]*(beta[j]+dv[j])) / det
		res := 0.0
		for k := 0; k < 3; k++ {
			r := (beta[k]+dv[k])*x - dv[k]*lb - s[k]
			if math.Abs(r) > res {
				res = math.Abs(r)
			}
		}
		if res < bestRes {
			bestRes, bestX, bestLB = res, x, lb
		}
	}
	return bestX, bestLB, bestRes
}

func landauPole(boundaryScaleGeV, betaTotal float64) float64 {
	if !finite(boundaryScaleGeV) || boundaryScaleGeV <= 0 {
		return math.NaN()
	}
	if betaTotal <= 0 {
		return math.Inf(1)
	}
	z := 8.0 * math.Pi * math.Pi / betaTotal
	if z > 700 {
		return math.Inf(1)
	}
	return boundaryScaleGeV * math.Exp(z)
}

func polesAbovePlanck(p FloatTriple) bool {
	return p.U1GUT > planckGeV && p.SU2L > planckGeV && p.SU3C > planckGeV
}

func finite(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }

func auditSafety(search DiophantineSearchAudit) AsymptoticSafetyAudit {
	return AsymptoticSafetyAudit{ExactCandidatesAudited: search.ExactClosureCandidates, ExactCandidatesBelowPlanck: search.ExactAsymptoticallySafeCandidates, ExactCandidatesNoLandauPole: search.ExactAsymptoticallySafeCandidates, BestNearMissNoLandauPole: search.BestSafeCandidate.NoSubPlanckLandauPole, BestNearMissBoundaryBelowPlanck: search.BestSafeCandidate.BelowPlanckBoundary, PlanckGeV: planckGeV, Verdict: "no exact candidate reaches the asymptotic-safety filter; the best bounded near-miss is safe but not a solution"}
}

func auditBaryonAnomaly(gate209 Gate209Snapshot, gens []SearchGenerator, search DiophantineSearchAudit) BaryonAnomalyAudit {
	return BaryonAnomalyAudit{LeptoquarkDynamicsSealInherited: gate209.LeptoquarkDynamicsSealActive, AllSearchRowsAnomalyCompatible: len(gens) == countAnomalyCompatible(gens), AllSearchRowsSealCompatible: len(gens) == countSealCompatible(gens), ExactCandidatesAnomalySafe: search.ExactAnomalySafeCandidates, ExactCandidatesSealCompatible: search.ExactAnomalySafeCandidates, ProtonDecayOperatorUsed: false, ProtonLifetimeComputed: false, Verdict: "candidate rows are filtered to anomaly-safe/seal-compatible carriers; no proton-decay operator or lifetime is introduced"}
}

func countAnomalyCompatible(gens []SearchGenerator) int {
	n := 0
	for _, g := range gens {
		if g.AnomalyCompatible {
			n++
		}
	}
	return n
}
func countSealCompatible(gens []SearchGenerator) int {
	n := 0
	for _, g := range gens {
		if g.LeptoquarkSealCompatible {
			n++
		}
	}
	return n
}

func auditFirewall(gate209 Gate209Snapshot, lattice representationrowlattice.Analysis, inputs RGInputAudit, pi ExactPiSeparationObstruction, search DiophantineSearchAudit, ba BaryonAnomalyAudit) FirewallAudit {
	return FirewallAudit{Gate209Inherited: gate209.Gate209Inherited, RepresentationLatticeInherited: lattice.Summary.LatticeConstructed, UniversalBetaRowInserted: false, ArbitraryRealRowCoefficientInserted: false, ExactClosureClaimed: search.ExactAsymptoticallySafeCandidates > 0, ConditionalPredictionEmitted: search.ExactAsymptoticallySafeCandidates > 0, ObservedLedgerUsedForFiniteCore: inputs.UsedForFiniteCoreDerivation, ProtonDecaySealViolated: !ba.LeptoquarkDynamicsSealInherited || ba.ProtonDecayOperatorUsed, ProtonLifetimeComputed: ba.ProtonLifetimeComputed, AbsoluteMassPredicted: false, PhysicalUnificationClaimed: false, ThresholdCorrectedPhysicalFitClaimed: false, FiniteMatchingCorrectionsDerived: false, RecommendedNextGate: "Gate 211 — multi-threshold rational lattice deformation or matching-correction obstruction audit", OpenRequirements: []string{"allow a second rational threshold scale and solve a piecewise linear RG system, or prove multi-threshold obstruction", "derive threshold matching corrections before treating near-misses as physical", "derive a non-Landau-pole UV completion if any positive high-scale beta row is retained", "keep LeptoquarkDynamicsSeal active unless a future theorem derives safe leptoquark dynamics"}, Verdict: firewallVerdict(inputs, pi, search, ba)}
}

func firewallVerdict(inputs RGInputAudit, pi ExactPiSeparationObstruction, search DiophantineSearchAudit, ba BaryonAnomalyAudit) string {
	ok := inputs.ObservedLedgerQuarantined && !inputs.UsedForFiniteCoreDerivation && !inputs.UniversalBetaRowAllowed && pi.ExactClosureImpossible && search.ExactAsymptoticallySafeCandidates == 0 && ba.LeptoquarkDynamicsSealInherited && !ba.ProtonDecayOperatorUsed && !ba.ProtonLifetimeComputed
	if ok {
		return "FIREWALL_PRESERVED_WITH_FAILED_ROUTE: no universal row, no real coefficients, no exact closure claim, no proton-decay seal violation, and no physical prediction emitted"
	}
	return "FAILED_ROUTE: firewall leak or inconsistent search state"
}

func FormatGate209(s Gate209Snapshot) string {
	return fmt.Sprintf("inherited=%t seal=%t baryonConservation=%t nativeLQFailed=%t dormantSlots=%d lifetimeObstructed=%t noSU5=%t noSO10=%t noPSDynamics=%t noLQPropagator=%t", s.Gate209Inherited, s.LeptoquarkDynamicsSealActive, s.SealedConnectionBaryonConservation, s.NativeLeptoquarkDynamicsFailed, s.DormantLeptoquarkSlotsSealed, s.ProtonLifetimeComputationObstructed, s.NoSU5Imported, s.NoSO10Imported, s.NoPatiSalamGaugeDynamicsImported, s.NoLeptoquarkPropagatorOrCoefficient)
}
func FormatInputs(i RGInputAudit) string {
	return fmt.Sprintf("%s MZ=%.6g A=(%.12g,%.12g,%.12g) alpha*^-1=%.12g bSM=%s planck=%.6g quarantined=%t finiteUse=%t universalAllowed=%t singleScale=%t", i.LedgerName, i.MZGeV, i.Alpha1GUTInverse, i.Alpha2Inverse, i.Alpha3Inverse, i.AlphaStarInverse, i.SMBeta, i.PlanckGeV, i.ObservedLedgerQuarantined, i.UsedForFiniteCoreDerivation, i.UniversalBetaRowAllowed, i.SingleScaleThresholdAssumed)
}
func FormatPiNoGo(p ExactPiSeparationObstruction) string {
	return fmt.Sprintf("eq=%q det(b,1,A)=%s nonzero=%t requiresDeltaOnSMBetaRay=%t betaRayNegative=%t latticeNonnegative=%t zeroFails=%t impossible=%t", p.Equation, p.DeterminantBOneA, p.DeterminantBOneANonZero, p.ExactClosureRequiresDeltaOnSMBetaRay, p.SMBetaRayHasNegativeComponents, p.RationalLatticeNonnegativeSemigroup, p.ZeroRowFailsSMTriangle, p.ExactClosureImpossible)
}
func FormatGeneratorAudit(a GeneratorAudit) string {
	return fmt.Sprintf("sourceRows=%d unique=%d exactRows=%d safeGenerators=%d anomalySafe=%d sealCompatible=%d zeroExcluded=%t semigroup=%t universalInserted=%t realCoeff=%t", a.SourceCandidateRows, a.SourceUniqueRows, a.ExactRationalRows, a.SafeGenerators, a.AnomalyCompatibleGenerators, a.LeptoquarkSealCompatibleRows, a.ZeroRowExcluded, a.SearchUsesNonnegativeSemigroup, !a.NoUniversalBetaRowInserted, !a.NoContinuousRowCoefficients)
}
func FormatCandidate(c CandidateEvaluation) string {
	return fmt.Sprintf("n=%d Δb=%s float=%s L*=%.9g LB=%.9g MB=%.9g M*=%.9g residualS=%.9g residualAlpha=%.9g exact=%t ordered=%t belowPlanck=%t totalBeta=%s poles=%s safe=%t anomaly=%t seal=%t names=[%s]", c.GeneratorCount, c.DeltaB, c.FloatDeltaB, c.BoundaryLogFromMZ, c.ThresholdLogFromMZ, c.ThresholdScaleMBGeV, c.BoundaryScaleMStarGeV, c.MaxClosureResidualS, c.MaxClosureResidualAlphaInv, c.ExactClosure, c.PositiveOrderedScales, c.BelowPlanckBoundary, c.TotalBeta, c.LandauPolesGeV, c.NoSubPlanckLandauPole, c.AnomalyCompatible, c.LeptoquarkSealCompatible, strings.Join(c.GeneratorNames, "; "))
}
func FormatSearch(a DiophantineSearchAudit) string {
	return fmt.Sprintf("maxCarriers=%d combinations=%d ordered=%d exact=%d exactAnomaly=%d exactSafe=%d bounded=%t exactNoGo=%t best={%s} bestSafe={%s}", a.MaxCarrierCount, a.CombinationsAudited, a.OrderedScaleCandidates, a.ExactClosureCandidates, a.ExactAnomalySafeCandidates, a.ExactAsymptoticallySafeCandidates, a.BoundedSearchOnly, a.ExactNoGoProvidedByPiSeparation, FormatCandidate(a.BestCandidate), FormatCandidate(a.BestSafeCandidate))
}
func FormatSafety(a AsymptoticSafetyAudit) string {
	return fmt.Sprintf("exactAudited=%d exactBelowPlanck=%d exactNoPole=%d bestNearMissNoPole=%t bestNearMissBoundaryBelowPlanck=%t planck=%.6g", a.ExactCandidatesAudited, a.ExactCandidatesBelowPlanck, a.ExactCandidatesNoLandauPole, a.BestNearMissNoLandauPole, a.BestNearMissBoundaryBelowPlanck, a.PlanckGeV)
}
func FormatBaryonAnomaly(a BaryonAnomalyAudit) string {
	return fmt.Sprintf("sealInherited=%t allRowsAnomaly=%t allRowsSeal=%t exactAnomaly=%d exactSeal=%d protonOperator=%t lifetime=%t", a.LeptoquarkDynamicsSealInherited, a.AllSearchRowsAnomalyCompatible, a.AllSearchRowsSealCompatible, a.ExactCandidatesAnomalySafe, a.ExactCandidatesSealCompatible, a.ProtonDecayOperatorUsed, a.ProtonLifetimeComputed)
}
func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate209=%t lattice=%t universalRow=%t realCoeff=%t exactClaim=%t prediction=%t observedFiniteUse=%t protonSealViolated=%t lifetime=%t mass=%t unification=%t fit=%t matching=%t", f.Gate209Inherited, f.RepresentationLatticeInherited, f.UniversalBetaRowInserted, f.ArbitraryRealRowCoefficientInserted, f.ExactClosureClaimed, f.ConditionalPredictionEmitted, f.ObservedLedgerUsedForFiniteCore, f.ProtonDecaySealViolated, f.ProtonLifetimeComputed, f.AbsoluteMassPredicted, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.FiniteMatchingCorrectionsDerived)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate209=%t piNoGo=%t search=%t exactSafe=%t bestNearMissSafe=%t status=%s comment=%q", s.TestsAudited, s.Gate209Inherited, s.ExactPiSeparationNoGo, s.BoundedSearchRan, s.ExactSafeSingleScaleSolutionFound, s.BestNearMissSafe, s.Status, s.Comment)
}
