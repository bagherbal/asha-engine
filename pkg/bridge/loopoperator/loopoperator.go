// Package loopoperator constructs the first finite Yukawa-incidence operator
// needed by the native Higgs one-loop program.
//
// Gate 52 established the sign/multiplicity ledger for a possible native
// radiative Higgs instability.  This package moves one step deeper: it turns the
// previously derived gauge-compatible Yukawa channels into a concrete finite
// operator
//
//	Y : H_left ⊗ H_scalar -> H_right
//
// using only channel incidence and scalar-fiber multiplicity.  The resulting
// operator is a selection-rule operator, not a physical Yukawa matrix: all
// allowed entries carry unit incidence weight and no observed coupling is
// inserted.  Consequently it can expose a native trace skeleton, but it cannot
// yet derive the Higgs mass-parameter sign, top dominance, or physical loop
// correction.
package loopoperator

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/looppotential"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type DomainBasis struct {
	Index       int
	LeftName    string
	ScalarName  string
	ScalarFiber int
}

type RightBasis struct {
	Index       int
	Name        string
	Kind        yukawaintertwiner.FermionKind
	Color       int
	Hypercharge float64
}

type KindTrace struct {
	Kind             yukawaintertwiner.FermionKind
	FiberEntries     int
	UnitLoopPressure float64
}

type Analysis struct {
	Ledger looppotential.Analysis
	Yukawa yukawaintertwiner.Analysis

	Domain     []DomainBasis
	Right      []RightBasis
	Operator   linear.Matrix
	RightGram  linear.Matrix
	DomainGram linear.Matrix

	DomainDimension      int
	RightDimension       int
	ScalarFiberDimension int
	AllowedFiberEntries  int
	UnusedDomainEntries  int
	Rank                 int

	RightTrace              float64
	DomainTrace             float64
	UnitFermionLoopPressure float64
	MaxColumnOccupancy      float64
	MinRightRowNormSquared  float64
	MaxRightRowNormSquared  float64
	RowNormsEqual           bool

	KindTraces          []KindTrace
	UpTypeFiberEntries  int
	UnitTopLikeSkeleton float64

	FiniteYukawaIncidenceOperatorDerived bool
	NativeLoopTraceSkeletonDerived       bool
	TopDominanceSelected                 bool
	TopLikeYukawaStrengthDerived         bool
	BosonicCounterOperatorDerived        bool
	RegulatorOrRenormalizationDerived    bool
	MuSquaredSignDerived                 bool
	NativeEffectivePotentialComputed     bool
	HiddenObservedCouplingsUsed          bool

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
		l, err := looppotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(l, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(l looppotential.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	y := l.Yukawa
	if !y.ChargeCompatibleYukawaChannelsDerived {
		return Analysis{}, fmt.Errorf("Gate 53 requires Gate 25 charge-compatible Yukawa channels")
	}

	scalarFiberDim := 0
	for _, b := range y.ScalarBranches {
		scalarFiberDim += b.Multiplicity
	}
	domain := buildDomain(y, scalarFiberDim)
	right := buildRight(y.RightStates)
	rightIndex := map[string]int{}
	for _, r := range right {
		rightIndex[r.Name] = r.Index
	}
	domainIndex := map[string]int{}
	for _, d := range domain {
		domainIndex[domainKey(d.LeftName, d.ScalarName, d.ScalarFiber)] = d.Index
	}

	op := linear.NewMatrix(len(right), len(domain))
	allowed := 0
	for _, ch := range y.Channels {
		r, ok := rightIndex[ch.Right.Name]
		if !ok {
			return Analysis{}, fmt.Errorf("right state %s missing from basis", ch.Right.Name)
		}
		for f := 0; f < ch.Scalar.Multiplicity; f++ {
			c, ok := domainIndex[domainKey(ch.Left.Name, ch.Scalar.Name, f)]
			if !ok {
				return Analysis{}, fmt.Errorf("domain state %s/%s/%d missing", ch.Left.Name, ch.Scalar.Name, f)
			}
			op.Set(r, c, 1)
			allowed++
		}
	}

	rightGram, _ := op.Mul(op.Transpose())
	domainGram, _ := op.Transpose().Mul(op)
	eig, err := linear.SymmetricEigenJacobi(rightGram, eps, 0)
	if err != nil {
		return Analysis{}, fmt.Errorf("right Gram spectrum: %w", err)
	}
	rank := linear.RankFromEigenvalues(eig.Values, eps)
	rightTrace, _ := rightGram.Trace()
	domainTrace, _ := domainGram.Trace()

	maxCol := 0.0
	for c := 0; c < op.Cols(); c++ {
		sum := 0.0
		for r := 0; r < op.Rows(); r++ {
			sum += op.At(r, c) * op.At(r, c)
		}
		if sum > maxCol {
			maxCol = sum
		}
	}
	minRow := math.Inf(1)
	maxRow := 0.0
	for r := 0; r < rightGram.Rows(); r++ {
		v := rightGram.At(r, r)
		if v < minRow {
			minRow = v
		}
		if v > maxRow {
			maxRow = v
		}
	}

	kindTraces := summarizeKinds(y)
	upEntries := 0
	for _, kt := range kindTraces {
		if kt.Kind == yukawaintertwiner.UpType {
			upEntries = kt.FiberEntries
		}
	}

	rowEqual := math.Abs(maxRow-minRow) < eps
	unitPressure := -rightTrace
	topSkeleton := -float64(upEntries)
	incidenceDerived := allowed == y.FiberEntryCount && rank == len(right) && maxCol <= 1+eps
	loopSkeleton := incidenceDerived && rightTrace == domainTrace && math.Abs(rightTrace-float64(allowed)) < eps

	truth := "The engine has now constructed a real finite Yukawa-incidence operator. Its unit-incidence fermion-loop pressure trace is negative by fermionic sign convention and contains the expected -6 up-type color skeleton. However, the unweighted operator treats all one-generation right channels equally, so it does not select top dominance, physical Yukawa strengths, bosonic counterterms, a regulator, or μ²<0."

	return Analysis{
		Ledger:                               l,
		Yukawa:                               y,
		Domain:                               domain,
		Right:                                right,
		Operator:                             op,
		RightGram:                            rightGram,
		DomainGram:                           domainGram,
		DomainDimension:                      len(domain),
		RightDimension:                       len(right),
		ScalarFiberDimension:                 scalarFiberDim,
		AllowedFiberEntries:                  allowed,
		UnusedDomainEntries:                  len(domain) - allowed,
		Rank:                                 rank,
		RightTrace:                           rightTrace,
		DomainTrace:                          domainTrace,
		UnitFermionLoopPressure:              unitPressure,
		MaxColumnOccupancy:                   maxCol,
		MinRightRowNormSquared:               minRow,
		MaxRightRowNormSquared:               maxRow,
		RowNormsEqual:                        rowEqual,
		KindTraces:                           kindTraces,
		UpTypeFiberEntries:                   upEntries,
		UnitTopLikeSkeleton:                  topSkeleton,
		FiniteYukawaIncidenceOperatorDerived: incidenceDerived,
		NativeLoopTraceSkeletonDerived:       loopSkeleton,
		TopDominanceSelected:                 false,
		TopLikeYukawaStrengthDerived:         false,
		BosonicCounterOperatorDerived:        false,
		RegulatorOrRenormalizationDerived:    false,
		MuSquaredSignDerived:                 false,
		NativeEffectivePotentialComputed:     false,
		HiddenObservedCouplingsUsed:          false,
		TruthStatement:                       truth,
		RecommendedNextGate:                  "Gate 54 — Top-Like Overlap / Condensate Kernel Search",
		RemainingUnknowns: []string{
			"U-20A2-TOP-LIKE-OVERLAP: derive a finite overlap/kernel that distinguishes the top-like channel from the other allowed channels",
			"U-20A3-BOSONIC-COUNTERWEIGHTS: construct finite gauge/scalar loop operators with kinetic normalization",
			"U-20A4-REGULATOR-CUTOFF: derive the finite spectral regulator or cutoff used by the loop trace",
			"U-20A5-MU-SIGN: compare fermionic and bosonic loop operators to derive the scalar mass-parameter sign",
		},
	}, nil
}

func buildDomain(y yukawaintertwiner.Analysis, scalarFiberDim int) []DomainBasis {
	out := make([]DomainBasis, 0, y.LeftDimension*scalarFiberDim)
	for _, l := range y.SU2L.States {
		for _, s := range y.ScalarBranches {
			for f := 0; f < s.Multiplicity; f++ {
				out = append(out, DomainBasis{Index: len(out), LeftName: l.Name, ScalarName: s.Name, ScalarFiber: f})
			}
		}
	}
	return out
}

func buildRight(rs []yukawaintertwiner.RightSinglet) []RightBasis {
	out := make([]RightBasis, 0, len(rs))
	for _, r := range rs {
		out = append(out, RightBasis{Index: len(out), Name: r.Name, Kind: r.Kind, Color: r.Color, Hypercharge: r.Hypercharge})
	}
	return out
}

func domainKey(left, scalar string, fiber int) string {
	return fmt.Sprintf("%s|%s|%d", left, scalar, fiber)
}

func summarizeKinds(y yukawaintertwiner.Analysis) []KindTrace {
	m := map[yukawaintertwiner.FermionKind]int{}
	for _, ch := range y.Channels {
		m[ch.Right.Kind] += ch.Scalar.Multiplicity
	}
	out := make([]KindTrace, 0, len(m))
	for k, n := range m {
		out = append(out, KindTrace{Kind: k, FiberEntries: n, UnitLoopPressure: -float64(n)})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Kind < out[j].Kind })
	return out
}

func FormatKindTraces(xs []KindTrace) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s: fiber=%d, pressure=%.0f", x.Kind, x.FiberEntries, x.UnitLoopPressure))
	}
	sort.Strings(parts)
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
