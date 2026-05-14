// Package gapledger builds the finite NJL/gap-kernel criticality ledger that
// follows the top-kernel search.
//
// Gate 54 found real ingredients for a composite-Higgs/condensate direction:
// three-color amplification and a diagonal generation spurion. It also exposed
// the hard obstruction: the current finite selection-rule data still tie
// up-type and down-type quark channels and do not derive an attractive
// four-fermion kernel strength.
//
// This package therefore formulates the next problem without pretending to
// solve it. It converts the finite channel inventory into an NJL-style
// criticality ledger. The output is a symbolic/structural condition of the form
//
//	G_hat · K_channel > C_reg
//
// where K_channel is the finite channel pressure skeleton already computed by
// the engine, while G_hat and C_reg are still unknown because the native
// four-fermion interaction and regulator/cutoff prescription have not been
// derived. No observed Higgs scale, top Yukawa, or fermion mass is inserted.
package gapledger

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/topkernel"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type ChannelCriticality struct {
	Kind               yukawaintertwiner.FermionKind
	Generation         int
	FiberEntries       int
	GenerationWeight   float64
	UnitPressure       float64
	WeightedPressure   float64
	NormalizedPressure float64
	TopLikeCandidate   bool
}

type FormalCondition struct {
	Name        string
	Expression  string
	KnownPart   string
	MissingPart string
	Derived     bool
}

type Analysis struct {
	TopKernel topkernel.Analysis

	ChannelCriticalities []ChannelCriticality

	GenerationCount int
	KindCount       int

	UnitQuarkPressureSkeleton    float64
	UnitLeptonPressureSkeleton   float64
	QuarkLeptonAmplification     float64
	StrongestWeightedPressure    float64
	StrongestWeightedChannel     ChannelCriticality
	StrongestUnitCriticalChannel ChannelCriticality

	UpDownDegeneracyResidual float64
	GenerationWeightSpread   float64

	FormalConditions []FormalCondition

	FourFermionKernelDerived     bool
	AttractiveInteractionDerived bool
	RegulatorDerived             bool
	CriticalThresholdDerived     bool
	GapEquationSolved            bool
	CondensateScaleDerived       bool
	TopOnlyCriticalityDerived    bool
	NativeNJLComputationComplete bool
	HiddenObservedCouplingsUsed  bool
	HiddenObservedMassScalesUsed bool

	UnitCriticalDiagnosticAvailable bool
	UnitThresholdRequiredCoupling   float64

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
		tk, err := topkernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(tk, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(tk topkernel.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if len(tk.KindKernels) == 0 || len(tk.GenerationWeights) != 3 {
		return Analysis{}, fmt.Errorf("Gate 55 requires Gate 54 kind kernels and three generation weights")
	}

	crits := make([]ChannelCriticality, 0, len(tk.KindKernels)*len(tk.GenerationWeights))
	maxAbsWeighted := -1.0
	var strongestWeighted ChannelCriticality
	maxAbsUnit := -1.0
	var strongestUnit ChannelCriticality

	for _, kk := range tk.KindKernels {
		for i, w := range tk.GenerationWeights {
			unit := math.Abs(kk.UnitPressurePerGen)
			weighted := unit * w
			c := ChannelCriticality{
				Kind:               kk.Kind,
				Generation:         i + 1,
				FiberEntries:       kk.FiberEntriesPerGen,
				GenerationWeight:   w,
				UnitPressure:       unit,
				WeightedPressure:   weighted,
				NormalizedPressure: 0,
				TopLikeCandidate:   kk.Kind == yukawaintertwiner.UpType && i == 1,
			}
			crits = append(crits, c)
			if weighted > maxAbsWeighted {
				maxAbsWeighted = weighted
				strongestWeighted = c
			}
			if unit > maxAbsUnit {
				maxAbsUnit = unit
				strongestUnit = c
			}
		}
	}
	for i := range crits {
		if maxAbsWeighted > eps {
			crits[i].NormalizedPressure = crits[i].WeightedPressure / maxAbsWeighted
		}
	}
	// update strongest copies with normalized pressure
	for _, c := range crits {
		if c.Kind == strongestWeighted.Kind && c.Generation == strongestWeighted.Generation {
			strongestWeighted = c
		}
		if c.Kind == strongestUnit.Kind && c.Generation == strongestUnit.Generation {
			strongestUnit = c
		}
	}

	quarkUnit := float64(tk.QuarkFiberEntriesPerGen)
	leptonUnit := float64(tk.LeptonFiberEntriesPerGen)
	amp := 0.0
	if leptonUnit != 0 {
		amp = quarkUnit / leptonUnit
	}

	unitRequired := 0.0
	if strongestWeighted.WeightedPressure > eps {
		unitRequired = 1.0 / strongestWeighted.WeightedPressure
	}

	conditions := []FormalCondition{
		{
			Name:        "finite channel pressure skeleton",
			Expression:  "K_channel = fiberEntries × generationWeight",
			KnownPart:   fmt.Sprintf("max K_channel=%.10f from %s/G%d", strongestWeighted.WeightedPressure, strongestWeighted.Kind, strongestWeighted.Generation),
			MissingPart: "none for the selection-rule skeleton",
			Derived:     true,
		},
		{
			Name:        "NJL criticality condition",
			Expression:  "G_hat · K_channel > C_reg",
			KnownPart:   "K_channel is finite and dimensionless",
			MissingPart: "G_hat attractive four-fermion strength and C_reg regulator threshold",
			Derived:     false,
		},
		{
			Name:        "gap equation",
			Expression:  "m = G_hat · Σ_i K_i ∫_reg m/(p²+m²)",
			KnownPart:   "allowed channels and multiplicities are known",
			MissingPart: "native finite integral/trace, cutoff, and interaction kernel",
			Derived:     false,
		},
		{
			Name:        "condensate scale",
			Expression:  "v = μ · r_gap",
			KnownPart:   "dimensionless finite pressure skeleton exists",
			MissingPart: "physical unit μ and nonzero gap solution",
			Derived:     false,
		},
	}

	truth := "The finite engine can now formulate the NJL criticality problem without importing observed couplings: color amplification and generation anisotropy give a dimensionless channel-pressure skeleton, but the attractive four-fermion strength, regulator threshold, and gap solution are still absent. Therefore condensation and v are not yet derived."

	return Analysis{
		TopKernel:                       tk,
		ChannelCriticalities:            crits,
		GenerationCount:                 len(tk.GenerationWeights),
		KindCount:                       len(tk.KindKernels),
		UnitQuarkPressureSkeleton:       quarkUnit,
		UnitLeptonPressureSkeleton:      leptonUnit,
		QuarkLeptonAmplification:        amp,
		StrongestWeightedPressure:       strongestWeighted.WeightedPressure,
		StrongestWeightedChannel:        strongestWeighted,
		StrongestUnitCriticalChannel:    strongestUnit,
		UpDownDegeneracyResidual:        tk.UpDownDegeneracyResidual,
		GenerationWeightSpread:          tk.GenerationWeightSpread,
		FormalConditions:                conditions,
		FourFermionKernelDerived:        false,
		AttractiveInteractionDerived:    false,
		RegulatorDerived:                false,
		CriticalThresholdDerived:        false,
		GapEquationSolved:               false,
		CondensateScaleDerived:          false,
		TopOnlyCriticalityDerived:       false,
		NativeNJLComputationComplete:    false,
		HiddenObservedCouplingsUsed:     false,
		HiddenObservedMassScalesUsed:    false,
		UnitCriticalDiagnosticAvailable: unitRequired > 0,
		UnitThresholdRequiredCoupling:   unitRequired,
		TruthStatement:                  truth,
		RecommendedNextGate:             "Gate 56 — Native Four-Fermion Kernel from x∧p/u(4) Sector",
		RemainingUnknowns: []string{
			"U-20C1-FOUR-FERMION-KERNEL: derive the attractive interaction from finite x∧p/u(4) gauge exchange or contact dynamics",
			"U-20C2-REGULATOR: derive the finite cutoff/spectral trace replacing the continuum NJL integral",
			"U-20C3-UP-DOWN-SPLITTING: break the up/down quark tie without observed Yukawa input",
			"U-20C4-GAP-SOLUTION: solve the native finite gap equation and test for a nonzero condensate",
			"U-20C5-SCALE: derive the physical condensate unit rather than fitting v=246 GeV",
		},
	}, nil
}

func FormatChannel(c ChannelCriticality) string {
	return fmt.Sprintf("%s/G%d fiber=%d weight=%.10g pressure=%.10g normalized=%.10g", c.Kind, c.Generation, c.FiberEntries, c.GenerationWeight, c.WeightedPressure, c.NormalizedPressure)
}

func FormatChannels(xs []ChannelCriticality) string {
	ys := append([]ChannelCriticality(nil), xs...)
	sort.Slice(ys, func(i, j int) bool {
		if math.Abs(ys[i].WeightedPressure-ys[j].WeightedPressure) > 1e-12 {
			return ys[i].WeightedPressure > ys[j].WeightedPressure
		}
		if ys[i].Kind != ys[j].Kind {
			return ys[i].Kind < ys[j].Kind
		}
		return ys[i].Generation < ys[j].Generation
	})
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, FormatChannel(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatConditions(xs []FormalCondition) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		status := "open"
		if x.Derived {
			status = "known"
		}
		parts = append(parts, fmt.Sprintf("%s(%s): %s; missing=%s", x.Name, status, x.Expression, x.MissingPart))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
