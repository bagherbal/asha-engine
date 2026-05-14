// Package currenthessian performs Gate 69: finite current Hessian / action
// second-variation search.
//
// Gate 68 exposed multiple positive exchange-kernel diagnostics but rejected
// selecting among them by convenience.  Gate 69 asks for the missing object:
// the current-field Hessian K obtained as the second variation of a finite
// action with respect to current-sector exchange fields.
//
// This package builds the audit layer for that Hessian.  It enumerates the
// allowed Hessian spaces and the natural diagonal ansatz families inherited
// from the previous current-sector diagnostics.  The result remains a firewall:
// no finite BF/projector/contact action currently contains typed current-sector
// fields whose second variation can be computed, so no propagator rule is
// selected and no NJL kernel is derived.
package currenthessian

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/exchangeselection"
)

type HessianCandidate struct {
	Name              string
	Dimension         int
	Diagonal          []float64
	Positive          bool
	FiniteDataOnly    bool
	DerivedFromAction bool
	Selected          bool
	DominantSector    string
	ConditionNumber   float64
	Interpretation    string
}

type Analysis struct {
	Previous exchangeselection.Analysis

	SectorCount                 int
	GeneratorCount              int
	SectorHessianDimension      int
	GeneratorHessianDimension   int
	DiagonalSectorDimension     int
	CandidateHessians           []HessianCandidate
	CandidateCount              int
	AllCandidatesPositive       bool
	AllCandidatesFiniteDataOnly bool
	AnyCandidateDerived         bool
	AnyCandidateSelected        bool
	DirectInverseAmbiguity      bool
	ActionVariablesDefined      bool
	SecondVariationComputed     bool
	CurrentHessianDerived       bool
	PropagatorRuleDerived       bool
	ExchangeKernelUpdated       bool
	AttractiveScalarDerived     bool
	UpDownSplittingDerived      bool
	CondensationClaimAllowed    bool
	HiddenObservedInputUsed     bool

	MinimalActionTemplate string
	HessianObstruction    string
	TruthStatement        string
	RecommendedNextGate   string
	RemainingUnknowns     []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := exchangeselection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev exchangeselection.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if prev.CandidateRuleCount == 0 {
		return Analysis{}, fmt.Errorf("Gate 68 candidate exchange rules are empty")
	}

	sectorNames := make([]string, 0, prev.CandidateRuleCount)
	// Gate 68 candidate rules are rule families, not sector data.  For the
	// current-Hessian audit, the current sectors are the four u(4)-shaped
	// sectors inherited in the exchange chain.
	sectorNames = []string{"central", "color-su3", "b-minus-l", "leptoquark"}
	sectorCount := len(sectorNames)
	generatorCount := 16

	// Candidate diagonal Hessians on the sector space.  These are deliberately
	// diagnostics, not selected physics.  Values are normalized enough to make
	// the audit transparent and reproducible without physical constants.
	candidates := []HessianCandidate{
		{
			Name:              "identity-sector-Hessian",
			Dimension:         sectorCount,
			Diagonal:          []float64{1, 1, 1, 1},
			FiniteDataOnly:    true,
			DerivedFromAction: false,
			Selected:          false,
			DominantSector:    "none",
			Interpretation:    "minimal positive Hessian; erases all representation data",
		},
		{
			Name:              "direct-Casimir-trace-Hessian",
			Dimension:         sectorCount,
			Diagonal:          []float64{8.0, 32.0, 2.6666666666666665, 24.0},
			FiniteDataOnly:    true,
			DerivedFromAction: false,
			Selected:          false,
			DominantSector:    "color-su3",
			Interpretation:    "sector trace Hessian inherited from C_A traces; favors color-su3",
		},
		{
			Name:              "inverse-Casimir-trace-Hessian",
			Dimension:         sectorCount,
			Diagonal:          []float64{1.0 / 8.0, 1.0 / 32.0, 1.0 / 2.6666666666666665, 1.0 / 24.0},
			FiniteDataOnly:    true,
			DerivedFromAction: false,
			Selected:          false,
			DominantSector:    "b-minus-l",
			Interpretation:    "inverse trace Hessian; flips dominance toward the smallest abelian trace",
		},
		{
			Name:              "trace-weight-Hessian",
			Dimension:         sectorCount,
			Diagonal:          []float64{0.12, 0.48, 0.04, 0.36},
			FiniteDataOnly:    true,
			DerivedFromAction: false,
			Selected:          false,
			DominantSector:    "color-su3",
			Interpretation:    "probability-like trace weights from kinetic normalization; useful but not an action Hessian",
		},
	}

	allPositive := true
	allFinite := true
	anyDerived := false
	anySelected := false
	for i := range candidates {
		candidates[i].Positive = positiveDiagonal(candidates[i].Diagonal, eps)
		candidates[i].ConditionNumber = conditionNumber(candidates[i].Diagonal, eps)
		if !candidates[i].Positive {
			allPositive = false
		}
		if !candidates[i].FiniteDataOnly {
			allFinite = false
		}
		if candidates[i].DerivedFromAction {
			anyDerived = true
		}
		if candidates[i].Selected {
			anySelected = true
		}
	}

	template := "S[j]=1/2 j^T K j - j^T source; K must be the second variation δ²S/δj_Aδj_B of a finite BF/projector/contact action with respect to typed current-sector fields"
	obstruction := "the current-sector field variables j_A are not yet embedded in the finite BF/projector/contact action, so δ²S/δj_Aδj_B cannot be computed; all diagonal Hessians remain diagnostics"
	truth := "Gate 69 exposes the exact missing mathematical object: a finite current-field Hessian. The engine can enumerate positive sector-Hessian candidates, but none is derived from an action second variation. Therefore no propagator rule, exchange kernel, NJL attraction, top condensation, Higgs scale, or fermion masses are claimed."

	return Analysis{
		Previous:                    prev,
		SectorCount:                 sectorCount,
		GeneratorCount:              generatorCount,
		SectorHessianDimension:      sectorCount * (sectorCount + 1) / 2,
		GeneratorHessianDimension:   generatorCount * (generatorCount + 1) / 2,
		DiagonalSectorDimension:     sectorCount,
		CandidateHessians:           candidates,
		CandidateCount:              len(candidates),
		AllCandidatesPositive:       allPositive,
		AllCandidatesFiniteDataOnly: allFinite,
		AnyCandidateDerived:         anyDerived,
		AnyCandidateSelected:        anySelected,
		DirectInverseAmbiguity:      true,
		ActionVariablesDefined:      false,
		SecondVariationComputed:     false,
		CurrentHessianDerived:       false,
		PropagatorRuleDerived:       false,
		ExchangeKernelUpdated:       false,
		AttractiveScalarDerived:     false,
		UpDownSplittingDerived:      false,
		CondensationClaimAllowed:    false,
		HiddenObservedInputUsed:     false,
		MinimalActionTemplate:       template,
		HessianObstruction:          obstruction,
		TruthStatement:              truth,
		RecommendedNextGate:         "Gate 70 — Current Field Embedding into Finite BF/Contact Action",
		RemainingUnknowns: []string{
			"U-20D2B5-CURRENT-FIELD-EMBEDDING: define typed current-sector fields inside the finite BF/projector/contact action",
			"U-20D2B6-ACTION-HESSIAN: compute δ²S/δj_Aδj_B instead of selecting a diagnostic Hessian",
			"U-20D2B7-OFF-DIAGONAL-HESSIAN: determine whether the true Hessian mixes current sectors or stays diagonal",
			"U-20D3-RELATIVE-COUPLINGS: derive relative current-sector couplings from the Hessian normalization",
			"U-20D4-UP-DOWN-SPLITTING: the current Hessian still acts on flavor sectors and does not split weak up/down channels",
			"U-20D6-CRITICAL-REGULATOR: NJL criticality still requires a real kernel and regulator threshold",
		},
	}, nil
}

func positiveDiagonal(xs []float64, eps float64) bool {
	for _, x := range xs {
		if x <= eps || math.IsNaN(x) || math.IsInf(x, 0) {
			return false
		}
	}
	return true
}

func conditionNumber(xs []float64, eps float64) float64 {
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, x := range xs {
		if x <= eps {
			continue
		}
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	if math.IsInf(min, 0) || min <= eps {
		return math.Inf(1)
	}
	return max / min
}

func FormatCandidates(xs []HessianCandidate) string {
	ys := append([]HessianCandidate(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(dim=%d, diag=%s, positive=%v, derived=%v, selected=%v, dominant=%s, cond=%.10f)", x.Name, x.Dimension, formatFloatList(x.Diagonal), x.Positive, x.DerivedFromAction, x.Selected, x.DominantSector, x.ConditionNumber))
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
