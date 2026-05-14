// Package sectoroperators performs Gate 66: current-sector operator
// construction search.
//
// Gate 65 showed that raw finite spectral lists could not be assigned to the
// current sectors central, color-su3, B-L, and leptoquark by multiplicity.  The
// correct next move is not to keep matching lists by count.  It is to construct
// the actual finite current-sector operators themselves.
//
// This package builds the sector Casimir/kinetic operators on the one-generation
// Pati-Salam flavor space with basis order:
//
//	0 = lepton, 1,2,3 = color seeds.
//
// For each current sector it constructs
//
//	C_sector = sum_a T_a^T T_a
//
// where T_a are the finite current generators in that sector.  These are real,
// positive operators and their spectra are genuine representation data.  They
// are not yet propagator denominators or mass thresholds: a propagator requires
// an exchange action, kinetic operator, and finite-to-continuum normalization.
package sectoroperators

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/sectorspectrum"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Generator struct {
	Sector string
	Name   string
	M      linear.Matrix
}

type SectorOperator struct {
	Sector             string
	GeneratorCount     int
	Operator           linear.Matrix
	DiagonalSpectrum   []float64
	Trace              float64
	Rank               int
	OffDiagonalMaxAbs  float64
	Positive           bool
	RepresentationData bool
	PropagatorData     bool
	Interpretation     string
}

type Analysis struct {
	Previous sectorspectrum.Analysis

	FlavorDimension int
	Generators      []Generator
	Operators       []SectorOperator

	SectorOperatorsConstructed     bool
	AllExpectedSectorsConstructed  bool
	RepresentationLevelMapsDerived bool
	OperatorsPositive              bool
	OperatorsDiagonalInFlavorBasis bool
	PropagatorDenominatorsDerived  bool
	ExchangeKernelUpdated          bool
	AttractiveScalarChannelDerived bool
	UpDownSplittingDerived         bool
	CondensationClaimAllowed       bool
	HiddenObservedInputUsed        bool

	ColorCasimirValue     float64
	LeptoquarkLeptonValue float64
	LeptoquarkColorValue  float64
	BLLeptonValue         float64
	BLColorValue          float64
	CentralValue          float64

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := sectorspectrum.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev sectorspectrum.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	gens := buildGenerators()
	ops, err := buildSectorOperators(gens, eps)
	if err != nil {
		return Analysis{}, err
	}

	by := map[string]SectorOperator{}
	allPositive := true
	allDiagonal := true
	for _, op := range ops {
		by[op.Sector] = op
		if !op.Positive {
			allPositive = false
		}
		if op.OffDiagonalMaxAbs > eps {
			allDiagonal = false
		}
	}
	expected := []string{"central", "color-su3", "b-minus-l", "leptoquark"}
	allExpected := true
	for _, s := range expected {
		if _, ok := by[s]; !ok {
			allExpected = false
		}
	}

	colorValue := 0.0
	if op, ok := by["color-su3"]; ok && len(op.DiagonalSpectrum) == 4 {
		colorValue = op.DiagonalSpectrum[1]
	}
	lqL, lqC := 0.0, 0.0
	if op, ok := by["leptoquark"]; ok && len(op.DiagonalSpectrum) == 4 {
		lqL, lqC = op.DiagonalSpectrum[0], op.DiagonalSpectrum[1]
	}
	blL, blC := 0.0, 0.0
	if op, ok := by["b-minus-l"]; ok && len(op.DiagonalSpectrum) == 4 {
		blL, blC = op.DiagonalSpectrum[0], op.DiagonalSpectrum[1]
	}
	central := 0.0
	if op, ok := by["central"]; ok && len(op.DiagonalSpectrum) == 4 {
		central = op.DiagonalSpectrum[0]
	}

	truth := "Gate 66 replaces multiplicity matching with actual current-sector operators. The central, color-su3, B-L, and leptoquark sectors each have a finite positive sector Casimir C_A=sum T_a^T T_a on the 1+3 flavor space. This is real representation data and resolves the operator-construction part of Gate 65. It still does not provide propagator denominators: C_A is a current-sector kinetic/Casimir diagnostic, not a mass spectrum or exchange propagator."

	return Analysis{
		Previous:                       prev,
		FlavorDimension:                4,
		Generators:                     gens,
		Operators:                      ops,
		SectorOperatorsConstructed:     len(ops) == 4,
		AllExpectedSectorsConstructed:  allExpected,
		RepresentationLevelMapsDerived: allExpected,
		OperatorsPositive:              allPositive,
		OperatorsDiagonalInFlavorBasis: allDiagonal,
		PropagatorDenominatorsDerived:  false,
		ExchangeKernelUpdated:          false,
		AttractiveScalarChannelDerived: false,
		UpDownSplittingDerived:         false,
		CondensationClaimAllowed:       false,
		HiddenObservedInputUsed:        false,
		ColorCasimirValue:              colorValue,
		LeptoquarkLeptonValue:          lqL,
		LeptoquarkColorValue:           lqC,
		BLLeptonValue:                  blL,
		BLColorValue:                   blC,
		CentralValue:                   central,
		TruthStatement:                 truth,
		RecommendedNextGate:            "Gate 67 — Current-Sector Casimir / Propagator Diagnostic",
		RemainingUnknowns: []string{
			"U-20D2B2-PROPAGATOR-DENOMINATORS: convert current-sector Casimirs into exchange denominators only after a finite kinetic/action theorem exists",
			"U-20D2B3-CURRENT-KINETIC-OPERATOR: decide whether C_A, inverse C_A, or another operator controls current exchange",
			"U-20D3-RELATIVE-COUPLINGS: derive relative current-sector couplings beyond Casimir traces",
			"U-20D4-UP-DOWN-SPLITTING: C_A acts on lepton/color flavor and still does not split up from down",
			"U-20D6-CRITICAL-REGULATOR: derive the NJL regulator threshold after a propagator kernel exists",
		},
	}, nil
}

func buildSectorOperators(gens []Generator, eps float64) ([]SectorOperator, error) {
	sectorGens := map[string][]Generator{}
	for _, g := range gens {
		sectorGens[g.Sector] = append(sectorGens[g.Sector], g)
	}
	names := []string{"central", "color-su3", "b-minus-l", "leptoquark"}
	out := make([]SectorOperator, 0, len(names))
	for _, name := range names {
		gs := sectorGens[name]
		if len(gs) == 0 {
			return nil, fmt.Errorf("no generators for sector %s", name)
		}
		op := linear.NewMatrix(4, 4)
		for _, g := range gs {
			term, err := g.M.Transpose().Mul(g.M)
			if err != nil {
				return nil, err
			}
			op, err = op.Add(term)
			if err != nil {
				return nil, err
			}
		}
		tr, err := op.Trace()
		if err != nil {
			return nil, err
		}
		diag := make([]float64, 4)
		rank := 0
		positive := true
		offdiag := 0.0
		for i := 0; i < 4; i++ {
			diag[i] = op.At(i, i)
			if diag[i] > eps {
				rank++
			}
			if diag[i] < -eps {
				positive = false
			}
			for j := 0; j < 4; j++ {
				if i == j {
					continue
				}
				v := math.Abs(op.At(i, j))
				if v > offdiag {
					offdiag = v
				}
			}
		}
		out = append(out, SectorOperator{
			Sector:             name,
			GeneratorCount:     len(gs),
			Operator:           op,
			DiagonalSpectrum:   diag,
			Trace:              tr,
			Rank:               rank,
			OffDiagonalMaxAbs:  offdiag,
			Positive:           positive,
			RepresentationData: true,
			PropagatorData:     false,
			Interpretation:     interpretation(name),
		})
	}
	return out, nil
}

func buildGenerators() []Generator {
	gens := []Generator{{Sector: "central", Name: "I4", M: linear.Identity(4)}}
	bl := linear.NewMatrix(4, 4)
	bl.Set(0, 0, -1)
	for i := 1; i < 4; i++ {
		bl.Set(i, i, 1.0/3.0)
	}
	gens = append(gens, Generator{Sector: "b-minus-l", Name: "B-L", M: bl})

	add := func(sector, name string, pairs ...[3]float64) {
		m := linear.NewMatrix(4, 4)
		for _, p := range pairs {
			m.Set(int(p[0]), int(p[1]), p[2])
		}
		gens = append(gens, Generator{Sector: sector, Name: name, M: m})
	}
	add("color-su3", "lambda1", [3]float64{1, 2, 1}, [3]float64{2, 1, 1})
	add("color-su3", "lambda2-real-skew", [3]float64{1, 2, 1}, [3]float64{2, 1, -1})
	add("color-su3", "lambda3", [3]float64{1, 1, 1}, [3]float64{2, 2, -1})
	add("color-su3", "lambda4", [3]float64{1, 3, 1}, [3]float64{3, 1, 1})
	add("color-su3", "lambda5-real-skew", [3]float64{1, 3, 1}, [3]float64{3, 1, -1})
	add("color-su3", "lambda6", [3]float64{2, 3, 1}, [3]float64{3, 2, 1})
	add("color-su3", "lambda7-real-skew", [3]float64{2, 3, 1}, [3]float64{3, 2, -1})
	r3 := 1.0 / math.Sqrt(3)
	add("color-su3", "lambda8", [3]float64{1, 1, r3}, [3]float64{2, 2, r3}, [3]float64{3, 3, -2 * r3})
	for c := 1; c <= 3; c++ {
		add("leptoquark", fmt.Sprintf("LQ%d-sym", c), [3]float64{0, float64(c), 1}, [3]float64{float64(c), 0, 1})
		add("leptoquark", fmt.Sprintf("LQ%d-skew", c), [3]float64{0, float64(c), 1}, [3]float64{float64(c), 0, -1})
	}
	return gens
}

func interpretation(sector string) string {
	switch sector {
	case "central":
		return "uniform u(1) flavor current Casimir"
	case "color-su3":
		return "color adjoint Casimir acting only on the three color seeds"
	case "b-minus-l":
		return "abelian lepton/color-polarizing charge-square operator"
	case "leptoquark":
		return "off-diagonal lepton-color current Casimir"
	default:
		return "current-sector Casimir"
	}
}

func FormatOperators(xs []SectorOperator) string {
	ys := append([]SectorOperator(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, trace=%.10f, rank=%d, diag=%s, propagator=%v)", x.Sector, x.GeneratorCount, x.Trace, x.Rank, formatFloatList(x.DiagonalSpectrum), x.PropagatorData))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}

func formatFloatList(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%.10f", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
