// Package topkernel searches for a non-uniform finite condensate kernel after
// the first finite Yukawa loop-incidence operator has been constructed.
//
// Gate 53 produced a real operator Y : H_left ⊗ H_scalar -> H_right, but its
// unit-incidence Gram matrix gave equal norm to every allowed right channel.
// This package asks the next question suggested by the composite-Higgs/NJL
// direction: does the finite geometry already contain an overlap kernel that
// distinguishes a top-like up-color channel from all other Yukawa channels?
//
// The result is intentionally conservative. The existing finite data provide
// (1) a quark/lepton color amplification skeleton and (2) a diagonal
// generation-breaking spurion from Higgs/contact anisotropy. But these two
// ingredients still do not select a top-only attractive channel: up and down
// quark channels remain tied, and the generation spurion is not yet canonically
// mapped to physical generations.
package topkernel

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/loopoperator"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type KindKernel struct {
	Kind                  yukawaintertwiner.FermionKind
	FiberEntriesPerGen    int
	UnitPressurePerGen    float64
	WeightedPressures     []float64
	TotalWeightedPressure float64
}

type Candidate struct {
	Name              string
	Dimension         int
	Weights           []float64
	NormalizedWeights []float64
	SelectsGeneration bool
	SelectsKind       bool
	SelectsTopOnly    bool
	Canonical         bool
	Detail            string
}

type Analysis struct {
	Loop       loopoperator.Analysis
	Generation generationbreak.Analysis

	UnitRightRowNormsEqual bool
	UnitRowNorm            float64

	QuarkFiberEntriesPerGen  int
	LeptonFiberEntriesPerGen int
	QuarkLeptonAmplification float64
	UpDownDegeneracyResidual float64

	GenerationWeights             []float64
	NormalizedGenerationWeights   []float64
	GenerationWeightSpread        float64
	DiagonalGenerationKernelFound bool
	GenerationKernelCanonical     bool

	KindKernels   []KindKernel
	Candidates    []Candidate
	BestCandidate Candidate

	TopLikeChannelSelected       bool
	TopDominanceKernelDerived    bool
	NonUniformOverlapKernelFound bool
	CondensateStrengthDerived    bool
	NativeGapKernelDerived       bool
	HiddenObservedCouplingsUsed  bool

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
		l, err := loopoperator.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		g, err := generationbreak.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(l, g, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(l loopoperator.Analysis, g generationbreak.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !l.FiniteYukawaIncidenceOperatorDerived {
		return Analysis{}, fmt.Errorf("Gate 54 requires Gate 53 finite Yukawa incidence operator")
	}
	if len(g.BestCandidate.Eigenvalues) != 3 {
		return Analysis{}, fmt.Errorf("Gate 54 requires a three-entry generation spurion, got %d", len(g.BestCandidate.Eigenvalues))
	}

	kindEntries := map[yukawaintertwiner.FermionKind]int{}
	for _, kt := range l.KindTraces {
		kindEntries[kt.Kind] = kt.FiberEntries
	}
	up := kindEntries[yukawaintertwiner.UpType]
	down := kindEntries[yukawaintertwiner.DownType]
	nu := kindEntries[yukawaintertwiner.NeutrinoType]
	electron := kindEntries[yukawaintertwiner.ElectronType]
	quarkEntries := up
	leptonEntries := nu
	amp := 0.0
	if leptonEntries != 0 {
		amp = float64(quarkEntries) / float64(leptonEntries)
	}
	upDownResidual := math.Abs(float64(up - down))

	weights := append([]float64(nil), g.BestCandidate.Eigenvalues...)
	sort.Slice(weights, func(i, j int) bool { return weights[i] > weights[j] })
	normWeights := normalizeByMax(weights)
	spread := 0.0
	if len(weights) > 0 && math.Abs(weights[len(weights)-1]) > eps {
		spread = weights[0] / weights[len(weights)-1]
	}

	kinds := []yukawaintertwiner.FermionKind{
		yukawaintertwiner.UpType,
		yukawaintertwiner.DownType,
		yukawaintertwiner.NeutrinoType,
		yukawaintertwiner.ElectronType,
	}
	kindKernels := make([]KindKernel, 0, len(kinds))
	for _, k := range kinds {
		entries := kindEntries[k]
		pressures := make([]float64, len(weights))
		total := 0.0
		for i, w := range weights {
			pressures[i] = -float64(entries) * w
			total += pressures[i]
		}
		kindKernels = append(kindKernels, KindKernel{
			Kind:                  k,
			FiberEntriesPerGen:    entries,
			UnitPressurePerGen:    -float64(entries),
			WeightedPressures:     pressures,
			TotalWeightedPressure: total,
		})
	}

	unitCandidate := Candidate{
		Name:              "unit Yukawa incidence kernel",
		Dimension:         l.RightDimension,
		Weights:           []float64{l.MinRightRowNormSquared, l.MaxRightRowNormSquared},
		NormalizedWeights: []float64{1, 1},
		SelectsGeneration: false,
		SelectsKind:       false,
		SelectsTopOnly:    false,
		Canonical:         true,
		Detail:            "The first finite loop operator is canonical at selection-rule level, but its right-channel row norms are equal and therefore cannot select top dominance.",
	}
	colorCandidate := Candidate{
		Name:              "three-color quark amplification skeleton",
		Dimension:         4,
		Weights:           []float64{float64(up), float64(down), float64(nu), float64(electron)},
		NormalizedWeights: normalizeByMax([]float64{float64(up), float64(down), float64(nu), float64(electron)}),
		SelectsGeneration: false,
		SelectsKind:       true,
		SelectsTopOnly:    false,
		Canonical:         true,
		Detail:            "Color multiplicity amplifies quark channels over lepton channels, but it ties up-type and down-type quark channels equally.",
	}
	generationCandidate := Candidate{
		Name:              "Higgs/contact anisotropy generation kernel",
		Dimension:         3,
		Weights:           weights,
		NormalizedWeights: normWeights,
		SelectsGeneration: distinctCount(weights, eps) == 3,
		SelectsKind:       false,
		SelectsTopOnly:    false,
		Canonical:         g.BestCandidate.Canonical,
		Detail:            "The finite anisotropy gives a diagonal 1+1+1 generation kernel, but it acts generation-wide and is not yet canonically mapped to top/bottom/electron/neutrino channels.",
	}
	combinedCandidate := Candidate{
		Name:              "color-amplified diagonal generation pressure",
		Dimension:         12,
		Weights:           []float64{-float64(up) * weights[0], -float64(down) * weights[0], -float64(nu) * weights[0], -float64(electron) * weights[0]},
		NormalizedWeights: nil,
		SelectsGeneration: generationCandidate.SelectsGeneration,
		SelectsKind:       true,
		SelectsTopOnly:    false,
		Canonical:         false,
		Detail:            "Combining color amplification with the diagonal generation spurion creates stronger quark pressure in a heaviest generation direction, but up and down remain exactly tied; this is not a top-only kernel.",
	}

	candidates := []Candidate{unitCandidate, colorCandidate, generationCandidate, combinedCandidate}
	best := combinedCandidate

	nonUniform := generationCandidate.SelectsGeneration || amp > 1
	topSelected := false
	topKernel := topSelected && combinedCandidate.SelectsTopOnly

	truth := "The engine found two real ingredients for a condensate kernel: color amplification and a diagonal generation spurion. Their combination can produce a strongest quark-generation pressure, but it does not select the top-like up channel because up/down quark channels remain degenerate and the generation spurion is not yet canonically mapped to physical generations."

	return Analysis{
		Loop:                          l,
		Generation:                    g,
		UnitRightRowNormsEqual:        l.RowNormsEqual,
		UnitRowNorm:                   l.MinRightRowNormSquared,
		QuarkFiberEntriesPerGen:       quarkEntries,
		LeptonFiberEntriesPerGen:      leptonEntries,
		QuarkLeptonAmplification:      amp,
		UpDownDegeneracyResidual:      upDownResidual,
		GenerationWeights:             weights,
		NormalizedGenerationWeights:   normWeights,
		GenerationWeightSpread:        spread,
		DiagonalGenerationKernelFound: generationCandidate.SelectsGeneration,
		GenerationKernelCanonical:     generationCandidate.Canonical,
		KindKernels:                   kindKernels,
		Candidates:                    candidates,
		BestCandidate:                 best,
		TopLikeChannelSelected:        topSelected,
		TopDominanceKernelDerived:     topKernel,
		NonUniformOverlapKernelFound:  nonUniform,
		CondensateStrengthDerived:     false,
		NativeGapKernelDerived:        false,
		HiddenObservedCouplingsUsed:   false,
		TruthStatement:                truth,
		RecommendedNextGate:           "Gate 55 — NJL Gap-Kernel / Criticality Ledger",
		RemainingUnknowns: []string{
			"U-20B1-UP-DOWN-SPLITTING: derive a finite operator that distinguishes top-like up channels from down-type channels without observed Yukawa input",
			"U-20B2-GENERATION-ORIENTATION: map the diagonal generation kernel canonically to physical generations",
			"U-20B3-FOUR-FERMION-KERNEL: derive the attractive NJL kernel from the x∧p/u(4) gauge sector",
			"U-20B4-CRITICALITY: compute whether the derived kernel crosses the condensation threshold",
			"U-20B5-SCALE: derive the condensate scale without fitting v=246 GeV",
		},
	}, nil
}

func normalizeByMax(xs []float64) []float64 {
	out := make([]float64, len(xs))
	max := 0.0
	for _, x := range xs {
		if math.Abs(x) > max {
			max = math.Abs(x)
		}
	}
	if max == 0 {
		return out
	}
	for i, x := range xs {
		out[i] = x / max
	}
	return out
}

func distinctCount(values []float64, eps float64) int {
	if len(values) == 0 {
		return 0
	}
	xs := append([]float64(nil), values...)
	sort.Float64s(xs)
	count := 1
	last := xs[0]
	for _, x := range xs[1:] {
		if math.Abs(x-last) > eps {
			count++
			last = x
		}
	}
	return count
}

func FormatFloatSlice(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%.10g", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatKindKernels(xs []KindKernel) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s: unit=%.0f weighted=%s", x.Kind, x.UnitPressurePerGen, FormatFloatSlice(x.WeightedPressures)))
	}
	sort.Strings(parts)
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("%s weights=%s normalized=%s selectsGeneration=%v selectsKind=%v selectsTopOnly=%v canonical=%v", c.Name, FormatFloatSlice(c.Weights), FormatFloatSlice(c.NormalizedWeights), c.SelectsGeneration, c.SelectsKind, c.SelectsTopOnly, c.Canonical)
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
