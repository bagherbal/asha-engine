// Package currentprojection applies the native x∧p/u(4) current-sector action
// to the finite scalar left-right projector constructed by Gate 58.
//
// Gate 58 built the target projector P_LR = U U^T on H_left⊗H_scalar.  This
// package takes the next step: it builds finite u(4)-shaped current generators
// on the one-generation lepton/color flavor index, lifts them to the left-domain
// and right-singlet spaces, and computes the induced action
//
//	T_induced = U^T T_domain U
//
// on the right-singlet scalar-bilinear image.
//
// The result is deliberately conservative.  It produces finite, unsigned
// scalar-projection overlap diagnostics for the current sectors.  It does not
// derive the Lorentz/Clifford Fierz signs, the kinetic/propagator normalization,
// or the attractive NJL sign.  It also confirms that the current inventory still
// does not split up-type from down-type quarks.
package currentprojection

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/chiraltrace"
	"github.com/bagherbal/asha-engine/pkg/bridge/loopoperator"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Generator struct {
	Sector string
	Name   string
	Flavor linear.Matrix // 4x4 on flavor order: lepton, color1, color2, color3.
}

type SectorCoefficient struct {
	Sector                 string
	GeneratorCount         int
	RightNorm              float64
	InducedNorm            float64
	OverlapTrace           float64
	NormalizedOverlap      float64
	IntertwinerResidual    float64
	CoefficientKnown       bool
	SignedCoefficientKnown bool
}

type Analysis struct {
	Chiral chiraltrace.Analysis

	FlavorDimension int
	DomainDimension int
	RightDimension  int
	ScalarLRRank    int

	Generators         []Generator
	SectorCoefficients []SectorCoefficient

	CurrentActionConstructed                  bool
	InducedCurrentActionConstructed           bool
	UnsignedScalarProjectionCoefficientsKnown bool
	SignedScalarProjectionCoefficientsKnown   bool
	GeneratorKineticNormalizationDerived      bool
	AttractiveSignDerived                     bool
	UpDownSplittingDerived                    bool
	LeptonQuarkSplitVisible                   bool
	HiddenObservedInputUsed                   bool

	MaxIntertwinerResidual float64
	CentralOverlap         float64
	ColorOverlap           float64
	BLOverlap              float64
	LeptoquarkOverlap      float64

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
		c, err := chiraltrace.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(c, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(c chiraltrace.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !c.ScalarLRProjectorConstructed || !c.FiniteFockTraceRulesConstructed {
		return Analysis{}, fmt.Errorf("Gate 59 requires Gate 58 scalar LR projector and finite trace rules")
	}
	gens := buildGenerators()
	domainLeft := leftNames(c.Loop.Domain)
	rightNames := rightNames(c.Loop.Right)

	sectorMap := map[string][]Generator{}
	for _, g := range gens {
		sectorMap[g.Sector] = append(sectorMap[g.Sector], g)
	}
	sectorNames := []string{"central", "color-su3", "b-minus-l", "leptoquark"}
	coeffs := make([]SectorCoefficient, 0, len(sectorNames))
	maxResidual := 0.0

	for _, sector := range sectorNames {
		gs := sectorMap[sector]
		sc := SectorCoefficient{Sector: sector, GeneratorCount: len(gs), CoefficientKnown: true, SignedCoefficientKnown: false}
		for _, g := range gs {
			td, err := liftDomain(g.Flavor, c.Loop.Domain, domainLeft)
			if err != nil {
				return Analysis{}, fmt.Errorf("lift domain %s/%s: %w", sector, g.Name, err)
			}
			tr, err := liftRight(g.Flavor, c.Loop.Right, rightNames)
			if err != nil {
				return Analysis{}, fmt.Errorf("lift right %s/%s: %w", sector, g.Name, err)
			}

			// T_induced = U^T T_domain U, where U maps right -> domain.
			tmp, err := td.Mul(c.NormalizedBilinearMap)
			if err != nil {
				return Analysis{}, err
			}
			induced, err := c.NormalizedBilinearMap.Transpose().Mul(tmp)
			if err != nil {
				return Analysis{}, err
			}

			inducedMinusRight, err := induced.Sub(tr)
			if err != nil {
				return Analysis{}, err
			}
			res := inducedMinusRight.FrobeniusNorm()
			if res > maxResidual {
				maxResidual = res
			}
			sc.IntertwinerResidual += res

			sc.RightNorm += frobSquared(tr)
			sc.InducedNorm += frobSquared(induced)
			prod, err := induced.Transpose().Mul(tr)
			if err != nil {
				return Analysis{}, err
			}
			t, err := prod.Trace()
			if err != nil {
				return Analysis{}, err
			}
			sc.OverlapTrace += t
		}
		if sc.RightNorm > eps {
			sc.NormalizedOverlap = sc.OverlapTrace / sc.RightNorm
		}
		coeffs = append(coeffs, sc)
	}

	bySector := map[string]SectorCoefficient{}
	for _, c := range coeffs {
		bySector[c.Sector] = c
	}
	truth := "The engine can now apply finite u(4)-shaped current-sector actions to the scalar LR projector. This upgrades the Fierz program from symbolic coefficient slots to computable finite overlap diagnostics. However, these are still unsigned representation overlaps: the Lorentz/Clifford Fierz sign, current kinetic normalization, propagator rule, and up/down splitting remain open."

	return Analysis{
		Chiral:                          c,
		FlavorDimension:                 4,
		DomainDimension:                 c.DomainDimension,
		RightDimension:                  c.RightDimension,
		ScalarLRRank:                    c.ProjectorRank,
		Generators:                      gens,
		SectorCoefficients:              coeffs,
		CurrentActionConstructed:        len(gens) == 16,
		InducedCurrentActionConstructed: true,
		UnsignedScalarProjectionCoefficientsKnown: true,
		SignedScalarProjectionCoefficientsKnown:   false,
		GeneratorKineticNormalizationDerived:      false,
		AttractiveSignDerived:                     false,
		UpDownSplittingDerived:                    false,
		LeptonQuarkSplitVisible:                   math.Abs(bySector["b-minus-l"].NormalizedOverlap) > eps || math.Abs(bySector["color-su3"].NormalizedOverlap) > eps,
		HiddenObservedInputUsed:                   false,
		MaxIntertwinerResidual:                    maxResidual,
		CentralOverlap:                            bySector["central"].NormalizedOverlap,
		ColorOverlap:                              bySector["color-su3"].NormalizedOverlap,
		BLOverlap:                                 bySector["b-minus-l"].NormalizedOverlap,
		LeptoquarkOverlap:                         bySector["leptoquark"].NormalizedOverlap,
		TruthStatement:                            truth,
		RecommendedNextGate:                       "Gate 60 — Generator Kinetic Normalization / Signed Fierz Coefficients",
		RemainingUnknowns: []string{
			"U-20D1B-CLIFFORD-TRACE-RULES: extend the finite Fock trace to signed Clifford/Lorentz Fierz identities",
			"U-20D1C-GENERATOR-NORMALIZATION: derive kinetic trace weights for central, color, B-L, and leptoquark currents",
			"U-20D2-ATTRACTIVE-SIGN: derive whether the scalar channel is attractive from the finite action/propagator rule",
			"U-20D4-UP-DOWN-SPLITTING: break the up/down quark tie without observed Yukawa input",
			"U-20D5-G-HAT: combine signed coefficients and propagator weights into a native four-fermion strength",
		},
	}, nil
}

func buildGenerators() []Generator {
	gens := []Generator{{Sector: "central", Name: "I4", Flavor: linear.Identity(4)}}
	bl := linear.NewMatrix(4, 4)
	bl.Set(0, 0, -1)
	for i := 1; i < 4; i++ {
		bl.Set(i, i, 1.0/3.0)
	}
	gens = append(gens, Generator{Sector: "b-minus-l", Name: "B-L", Flavor: bl})

	add := func(sector, name string, pairs ...[3]float64) {
		m := linear.NewMatrix(4, 4)
		for _, p := range pairs {
			m.Set(int(p[0]), int(p[1]), p[2])
		}
		gens = append(gens, Generator{Sector: sector, Name: name, Flavor: m})
	}
	// Color SU(3)-shaped real basis embedded in the color slots 1,2,3.
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

func leftNames(domain []loopoperator.DomainBasis) []string {
	m := map[string]bool{}
	out := []string{}
	for _, d := range domain {
		if !m[d.LeftName] {
			m[d.LeftName] = true
			out = append(out, d.LeftName)
		}
	}
	sort.Strings(out)
	return out
}

func rightNames(right []loopoperator.RightBasis) []string {
	out := make([]string, 0, len(right))
	for _, r := range right {
		out = append(out, r.Name)
	}
	sort.Strings(out)
	return out
}

func liftDomain(flavor linear.Matrix, domain []loopoperator.DomainBasis, _ []string) (linear.Matrix, error) {
	idx := map[string][]int{}
	for _, d := range domain {
		idx[d.LeftName] = append(idx[d.LeftName], d.Index)
	}
	out := linear.NewMatrix(len(domain), len(domain))
	for _, from := range domain {
		fFrom := flavorIndex(from.LeftName)
		if fFrom < 0 {
			return linear.Matrix{}, fmt.Errorf("unknown left flavor %s", from.LeftName)
		}
		for _, to := range domain {
			if from.ScalarName != to.ScalarName || from.ScalarFiber != to.ScalarFiber {
				continue
			}
			fTo := flavorIndex(to.LeftName)
			if fTo < 0 {
				return linear.Matrix{}, fmt.Errorf("unknown left flavor %s", to.LeftName)
			}
			v := flavor.At(fTo, fFrom)
			if math.Abs(v) > 0 {
				out.Set(to.Index, from.Index, v)
			}
		}
	}
	return out, nil
}

func liftRight(flavor linear.Matrix, right []loopoperator.RightBasis, _ []string) (linear.Matrix, error) {
	out := linear.NewMatrix(len(right), len(right))
	for _, from := range right {
		fFrom := flavorIndex(from.Name)
		if fFrom < 0 {
			return linear.Matrix{}, fmt.Errorf("unknown right flavor %s", from.Name)
		}
		for _, to := range right {
			if !sameWeakType(from.Name, to.Name) {
				continue
			}
			fTo := flavorIndex(to.Name)
			if fTo < 0 {
				return linear.Matrix{}, fmt.Errorf("unknown right flavor %s", to.Name)
			}
			v := flavor.At(fTo, fFrom)
			if math.Abs(v) > 0 {
				out.Set(to.Index, from.Index, v)
			}
		}
	}
	return out, nil
}

func flavorIndex(name string) int {
	switch {
	case strings.Contains(name, "nu_") || strings.Contains(name, "e_"):
		return 0
	case strings.Contains(name, "^1"):
		return 1
	case strings.Contains(name, "^2"):
		return 2
	case strings.Contains(name, "^3"):
		return 3
	}
	return -1
}

func sameWeakType(a, b string) bool {
	return weakType(a) == weakType(b)
}

func weakType(name string) string {
	switch {
	case strings.Contains(name, "u_") || strings.Contains(name, "nu_"):
		return "up"
	case strings.Contains(name, "d_") || strings.Contains(name, "e_"):
		return "down"
	default:
		return "unknown"
	}
}

func frobSquared(m linear.Matrix) float64 {
	n := m.FrobeniusNorm()
	return n * n
}

func FormatSectorCoefficients(xs []SectorCoefficient) string {
	ys := append([]SectorCoefficient(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, overlap=%.10f, residual=%.3e, signed=open)", x.Sector, x.GeneratorCount, x.NormalizedOverlap, x.IntertwinerResidual))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
