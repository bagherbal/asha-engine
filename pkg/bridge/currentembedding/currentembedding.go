// Package currentembedding performs Gate 70: current-field embedding into a
// finite BF/contact action.
//
// Gate 69 exposed the missing current Hessian K.  This gate takes the next
// architectural step: it defines typed current-sector field variables and the
// minimal finite action template in which a Hessian could live.  The result is
// deliberately a bridge firewall.  The field slots are now typed, but the map
// from the Fock/u(4) current inventory into Boolean/contact block-connection
// operators is still not derived, so the Hessian and propagator rule remain
// open.
package currentembedding

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/currenthessian"
)

type CurrentField struct {
	Name                string
	Generators          int
	Role                string
	EmbeddedInAction    bool
	EmbeddingMapDerived bool
}

type ActionSlot struct {
	Name      string
	Dimension int
	Purpose   string
	Derived   bool
}

type Analysis struct {
	Previous currenthessian.Analysis

	Fields              []CurrentField
	SectorFieldCount    int
	GeneratorFieldCount int
	AllFieldsTyped      bool

	ActionSlots                 []ActionSlot
	ActionSlotCount             int
	ContactBlockActionAvailable bool
	MinimalActionTemplate       string
	SourceCouplingTemplate      string

	FieldSlotsDefined                bool
	FockCurrentInventoryAvailable    bool
	BooleanContactActionAvailable    bool
	CurrentToContactEmbeddingDerived bool
	SourceFunctionalDerived          bool
	HessianComputable                bool
	CurrentHessianDerived            bool
	PropagatorRuleDerived            bool
	ExchangeKernelUpdated            bool
	AttractiveScalarDerived          bool
	UpDownSplittingDerived           bool
	CondensationClaimAllowed         bool
	HiddenObservedInputUsed          bool

	EmbeddingObstruction string
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
		prev, err := currenthessian.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev currenthessian.Analysis) (Analysis, error) {
	if prev.SectorCount != 4 || prev.GeneratorCount != 16 {
		return Analysis{}, fmt.Errorf("Gate 69 current-domain mismatch: sectors=%d generators=%d", prev.SectorCount, prev.GeneratorCount)
	}

	fields := []CurrentField{
		{Name: "central", Generators: 1, Role: "overall u(1) current-sector field", EmbeddedInAction: true, EmbeddingMapDerived: false},
		{Name: "color-su3", Generators: 8, Role: "color adjoint current-sector field", EmbeddedInAction: true, EmbeddingMapDerived: false},
		{Name: "b-minus-l", Generators: 1, Role: "B-L abelian current-sector field", EmbeddedInAction: true, EmbeddingMapDerived: false},
		{Name: "leptoquark", Generators: 6, Role: "off-diagonal lepton-color current-sector field", EmbeddedInAction: true, EmbeddingMapDerived: false},
	}
	totalGenerators := 0
	allTyped := true
	anyMapDerived := false
	for _, f := range fields {
		totalGenerators += f.Generators
		if f.Name == "" || f.Generators <= 0 || f.Role == "" {
			allTyped = false
		}
		if f.EmbeddingMapDerived {
			anyMapDerived = true
		}
	}

	slots := []ActionSlot{
		{Name: "B", Dimension: 56, Purpose: "Boolean-support B-sector field used by the contact vacuum action", Derived: true},
		{Name: "A_block", Dimension: 56 * 56, Purpose: "Boolean-compressed block-connection operator on K ⊕ K⊥", Derived: true},
		{Name: "j_current", Dimension: totalGenerators, Purpose: "typed current-sector exchange fields central/color/B-L/leptoquark", Derived: true},
		{Name: "E_current_to_block", Dimension: totalGenerators * 56 * 56, Purpose: "missing embedding map from current fields to finite block-connection directions", Derived: false},
		{Name: "K_current", Dimension: totalGenerators * totalGenerators, Purpose: "missing current-field Hessian from the second variation of the finite action", Derived: false},
	}

	actionTemplate := "S[B,A,j] = S_B[B] + S_block[A;K⊕K⊥] + 1/2 j^T K_current j - <j, J_source[B,A]>"
	sourceTemplate := "J_source[A]_a should be derived by projecting finite block curvature/contact response onto current generator a through E_current_to_block; E and J_source are still open"
	obstruction := "Gate 70 types the current fields and the finite action slots, but no theorem maps the u(4) current generators into Boolean/contact block-connection operators. Without that embedding E_current_to_block, δ²S/δjδj is not computable."
	truth := "The engine has advanced from an untyped missing Hessian to a typed current-field action template. This is architectural progress, not a propagator theorem: the current-to-contact embedding, source functional, Hessian, exchange kernel, NJL attraction, and up/down splitting remain open."

	return Analysis{
		Previous:                         prev,
		Fields:                           fields,
		SectorFieldCount:                 len(fields),
		GeneratorFieldCount:              totalGenerators,
		AllFieldsTyped:                   allTyped,
		ActionSlots:                      slots,
		ActionSlotCount:                  len(slots),
		ContactBlockActionAvailable:      true,
		MinimalActionTemplate:            actionTemplate,
		SourceCouplingTemplate:           sourceTemplate,
		FieldSlotsDefined:                allTyped && totalGenerators == 16,
		FockCurrentInventoryAvailable:    true,
		BooleanContactActionAvailable:    true,
		CurrentToContactEmbeddingDerived: anyMapDerived,
		SourceFunctionalDerived:          false,
		HessianComputable:                false,
		CurrentHessianDerived:            false,
		PropagatorRuleDerived:            false,
		ExchangeKernelUpdated:            false,
		AttractiveScalarDerived:          false,
		UpDownSplittingDerived:           false,
		CondensationClaimAllowed:         false,
		HiddenObservedInputUsed:          false,
		EmbeddingObstruction:             obstruction,
		TruthStatement:                   truth,
		RecommendedNextGate:              "Gate 71 — Current-to-Contact Embedding Map Search",
		RemainingUnknowns: []string{
			"U-20D2B5A-CURRENT-TO-CONTACT-MAP: derive E_current_to_block from u(4) Fock currents to Boolean/contact block-connection operators",
			"U-20D2B5B-SOURCE-FUNCTIONAL: construct J_source[B,A] rather than declaring a source slot",
			"U-20D2B6-ACTION-HESSIAN: compute K_current=δ²S/δjδj after E and J_source exist",
			"U-20D2B7-OFF-DIAGONAL-HESSIAN: determine whether current sectors mix in the true Hessian",
			"U-20D4-UP-DOWN-SPLITTING: current-sector embedding still does not distinguish top-like up from bottom-like down",
			"U-20D6-CRITICAL-REGULATOR: NJL criticality still requires a selected kernel and regulator threshold",
		},
	}, nil
}

func FormatFields(xs []CurrentField) string {
	ys := append([]CurrentField(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, role=%s, slot=%v, map=%v)", x.Name, x.Generators, x.Role, x.EmbeddedInAction, x.EmbeddingMapDerived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSlots(xs []ActionSlot) string {
	ys := append([]ActionSlot(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(dim=%d, derived=%v, purpose=%s)", x.Name, x.Dimension, x.Derived, x.Purpose))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
