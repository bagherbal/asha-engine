// Package exchangeselection performs Gate 68: finite exchange-action selection
// principle.
//
// Gates 66-67 constructed current-sector Casimirs and exposed several possible
// propagator/kernel diagnostics built from them: direct C_A, inverse nonzero
// C_A^+, trace-normalized C_A, and identity/unit rules.  Gate 68 asks whether a
// finite action principle selects one of these rules.
//
// The result is a deliberate firewall.  All candidate kernels are positive and
// representation-compatible diagnostics, but none is selected by the current
// finite action data.  A real exchange kernel requires an operator-level kinetic
// action, e.g. a second variation of the finite BF/projector/contact action on
// current-sector gauge fields, not a choice made from convenient numerical
// behavior.
package exchangeselection

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/casimirkernel"
)

type CandidateRule struct {
	Name                 string
	KernelType           string
	Positive             bool
	UsesOnlyFiniteData   bool
	SelectedByAction     bool
	DominantSector       string
	TotalDiagnostic      float64
	AttractionDiagnostic float64
	UpDownSplitting      bool
	Interpretation       string
}

type Analysis struct {
	Previous casimirkernel.Analysis

	CandidateRules                 []CandidateRule
	CandidateRuleCount             int
	AllRulesPositive               bool
	AllRulesUseOnlyFiniteData      bool
	AnyRuleSelectedByAction        bool
	DirectInverseDisagree          bool
	TraceNormalizedAmbiguous       bool
	MinimalActionFormExposed       bool
	KineticOperatorDerived         bool
	SecondVariationDerived         bool
	PropagatorRuleDerived          bool
	ExchangeKernelUpdated          bool
	AttractiveScalarChannelDerived bool
	UpDownSplittingDerived         bool
	CondensationClaimAllowed       bool
	HiddenObservedInputUsed        bool

	SelectionObstruction string
	TruthStatement       string
	RecommendedNextGate  string
	RemainingUnknowns    []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := casimirkernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev casimirkernel.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if len(prev.Diagnostics) == 0 {
		return Analysis{}, fmt.Errorf("Gate 67 diagnostics are empty")
	}

	directTotal := prev.TotalDirectTrace
	inverseTotal := prev.TotalInverseTrace
	traceNormalizedTotal := float64(len(prev.Diagnostics))
	unitTotal := float64(len(prev.Diagnostics))

	rules := []CandidateRule{
		{
			Name:                 "direct-Casimir",
			KernelType:           "K_A = C_A",
			Positive:             true,
			UsesOnlyFiniteData:   true,
			SelectedByAction:     false,
			DominantSector:       prev.DominantDirectSector,
			TotalDiagnostic:      directTotal,
			AttractionDiagnostic: safeReciprocal(directTotal),
			UpDownSplitting:      false,
			Interpretation:       "weights sectors by raw representation Casimir strength; favors color-su3 in the current finite data",
		},
		{
			Name:                 "inverse-nonzero-Casimir",
			KernelType:           "K_A = C_A^+ on nonzero sector",
			Positive:             true,
			UsesOnlyFiniteData:   true,
			SelectedByAction:     false,
			DominantSector:       prev.DominantInverseSector,
			TotalDiagnostic:      inverseTotal,
			AttractionDiagnostic: inverseTotal,
			UpDownSplitting:      false,
			Interpretation:       "weights weak Casimir directions strongly; favors abelian small-charge data and disagrees with direct-Casimir weighting",
		},
		{
			Name:                 "trace-normalized-Casimir",
			KernelType:           "K_A = C_A / Tr(C_A)",
			Positive:             true,
			UsesOnlyFiniteData:   true,
			SelectedByAction:     false,
			DominantSector:       "none",
			TotalDiagnostic:      traceNormalizedTotal,
			AttractionDiagnostic: traceNormalizedTotal,
			UpDownSplitting:      false,
			Interpretation:       "removes scale information and becomes probability-like; useful for diagnostics but too normalized to be a kinetic action",
		},
		{
			Name:                 "unit-sector-rule",
			KernelType:           "K_A = I_sector",
			Positive:             true,
			UsesOnlyFiniteData:   true,
			SelectedByAction:     false,
			DominantSector:       "none",
			TotalDiagnostic:      unitTotal,
			AttractionDiagnostic: unitTotal,
			UpDownSplitting:      false,
			Interpretation:       "treats all current sectors equally; it is the least structured diagnostic but erases finite representation data",
		},
	}

	allPositive := true
	allFinite := true
	anySelected := false
	anyUpDown := false
	for _, r := range rules {
		if !r.Positive || math.IsNaN(r.TotalDiagnostic) || math.IsInf(r.TotalDiagnostic, 0) {
			allPositive = false
		}
		if !r.UsesOnlyFiniteData {
			allFinite = false
		}
		if r.SelectedByAction {
			anySelected = true
		}
		if r.UpDownSplitting {
			anyUpDown = true
		}
	}

	disagree := prev.DominantDirectSector != "" && prev.DominantInverseSector != "" && prev.DominantDirectSector != prev.DominantInverseSector
	obstruction := "multiple positive finite kernel diagnostics exist, but they disagree on dominant sector and no second-variation/kinetic action selects among them"
	truth := "Gate 68 exposes the exchange-action selection problem. Direct Casimir, inverse Casimir, trace-normalized Casimir, and unit-sector rules are all finite diagnostics, but none is promoted to a propagator by the present theory. The correct next step is to derive the finite current kinetic operator from an action Hessian or BF/projector second variation; choosing a kernel by convenience would be fitting."

	return Analysis{
		Previous:                       prev,
		CandidateRules:                 rules,
		CandidateRuleCount:             len(rules),
		AllRulesPositive:               allPositive,
		AllRulesUseOnlyFiniteData:      allFinite,
		AnyRuleSelectedByAction:        anySelected,
		DirectInverseDisagree:          disagree,
		TraceNormalizedAmbiguous:       true,
		MinimalActionFormExposed:       true,
		KineticOperatorDerived:         false,
		SecondVariationDerived:         false,
		PropagatorRuleDerived:          false,
		ExchangeKernelUpdated:          false,
		AttractiveScalarChannelDerived: false,
		UpDownSplittingDerived:         anyUpDown,
		CondensationClaimAllowed:       false,
		HiddenObservedInputUsed:        false,
		SelectionObstruction:           obstruction,
		TruthStatement:                 truth,
		RecommendedNextGate:            "Gate 69 — Finite Current Hessian / Action Second-Variation Search",
		RemainingUnknowns: []string{
			"U-20D2B4-ACTION-SELECTION: derive the finite action Hessian for current exchange instead of selecting direct/inverse diagnostics manually",
			"U-20D2B5-CURRENT-HESSIAN: construct the second variation of the BF/projector/contact action on current-sector fields",
			"U-20D3-RELATIVE-COUPLINGS: derive relative current-sector couplings after the kinetic operator is known",
			"U-20D4-UP-DOWN-SPLITTING: current-sector kernels still act on lepton/color flavor and do not distinguish up/down weak channel",
			"U-20D6-CRITICAL-REGULATOR: NJL criticality still requires regulator/cutoff and a real exchange kernel",
		},
	}, nil
}

func safeReciprocal(x float64) float64 {
	if math.Abs(x) < 1e-15 {
		return 0
	}
	return 1.0 / x
}

func FormatRules(xs []CandidateRule) string {
	ys := append([]CandidateRule(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(%s, positive=%v, selected=%v, dominant=%s, total=%.10f, attractDiag=%.10f)", x.Name, x.KernelType, x.Positive, x.SelectedByAction, x.DominantSector, x.TotalDiagnostic, x.AttractionDiagnostic))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
