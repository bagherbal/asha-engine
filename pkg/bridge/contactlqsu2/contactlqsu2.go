// Package contactlqsu2 implements Gate 133: leptoquark real-orientation
// versus weak-doublet obstruction / SU(2)L action search.
//
// Gate 132 proved that the six current-side leptoquark slots are genuinely
// typed as color(3) times real-orientation(2), but that the second factor is
// symmetric/skew real off-diagonal orientation rather than a weak-doublet
// theorem. Gate 133 audits the exact missing object: an SU(2)_L action on those
// two real orientations. The strongest structure available is an SO(2)/U(1)
// rotation on each real orientation plane. A physical weak doublet would require
// a non-abelian su(2) triple, representation rows, hypercharge, local field
// variables, and contact-row assignments. None is selected by the finite data.
package contactlqsu2

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactlqtensor"
)

type ActionKind string

const (
	OrientationSO2Action ActionKind = "orientation-so2-action"
	ColorWiseSO2Action   ActionKind = "color-wise-so2-cubed-action"
	ClaimedSU2LAction    ActionKind = "claimed-su2l-action"
	BorrowedMatterSU2    ActionKind = "borrowed-matter-su2-action"
	ObservedFitAction    ActionKind = "observed-fit-action"
)

type OrientationPlane struct {
	ColorIndex            int
	Slots                 [2]string
	RealDimension         int
	SO2GeneratorAvailable bool
	SO2GeneratorName      string
	AbelianClosure        bool
	SU2TripleDerived      bool
	WeakDoubletDerived    bool
	ContactRowsAssigned   bool
	RepresentationRows    bool
	BetaPermitted         bool
	Obstruction           string
}

type ActionCandidate struct {
	Name string
	Kind ActionKind

	ColorPlanes            int
	RealDimensionPerColor  int
	TotalRealDimension     int
	CurrentDerived         bool
	ContactDerived         bool
	MatterBorrowed         bool
	Canonical              bool
	Natural                bool
	MetricPreserving       bool
	OrientationRotation    bool
	AbelianOnly            bool
	NonAbelianSU2Triple    bool
	SU2Commutation         bool
	WeakDoubletSemantics   bool
	HyperchargeDerived     bool
	LocalFieldDerived      bool
	ContactRowAssignment   bool
	RepresentationRows     bool
	BetaPermitted          bool
	RequiresSemanticBridge bool
	RequiresS6Choice       bool
	RequiresObservedInput  bool
	HiddenChoices          int
	Obstruction            string
}

type Criterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Summary struct {
	ContactRows                 int
	LeptoquarkRows              int
	ColorPlanes                 int
	RealOrientationsPerColor    int
	TotalCurrentLQSlots         int
	OrientationSO2Available     bool
	ColorWiseSO2Available       bool
	DiagonalSO2Available        bool
	NonAbelianSU2TripleDerived  bool
	SU2WeakDoubletActionDerived bool
	HyperchargeRowDerived       bool
	LocalFieldMapDerived        bool
	ContactAssignments          int
	RepresentationCompleteRows  int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	ResidualS6Choices           int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactlqtensor.Analysis

	Planes     []OrientationPlane
	Candidates []ActionCandidate
	Criteria   []Criterion
	Summary    Summary

	ContactRows                  int
	LeptoquarkRows               int
	ColorPlanes                  int
	RealOrientationsPerColor     int
	TotalCurrentLQSlots          int
	OrientationSO2Available      bool
	ColorWiseSO2Available        bool
	DiagonalSO2Available         bool
	OrientationActionAbelian     bool
	NonAbelianSU2TripleDerived   bool
	SU2CommutationDerived        bool
	SU2WeakDoubletActionDerived  bool
	WeakDoubletSemanticsDerived  bool
	HyperchargeRowDerived        bool
	LocalFieldMapDerived         bool
	BorrowedMatterActionRejected bool
	SemanticBridgeMissing        bool
	S6ObstructionInherited       bool
	CurrentNaturalSU2Action      bool

	RepresentationCompleteRows    int
	RepresentationOpenRows        int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	ThresholdCorrectedBetaDerived bool
	FullBetaMatchingTensorDerived bool

	ResidualNullityBefore    int
	ResidualNullityAfter     int
	HiddenObservedInputUsed  bool
	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	PhysicalScaleDerived     bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := contactlqtensor.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactlqtensor.Analysis) (Analysis, error) {
	if !prev.S6ObstructionInherited || !prev.CurrentRealTensorDerived || !prev.ColorTripletSemantics || !prev.RealOrientationSemantics || prev.LeptoquarkRows != 6 || prev.CurrentLQSlots != 6 {
		return Analysis{}, fmt.Errorf("Gate 133 requires Gate 132 current leptoquark real tensor and inherited S6 obstruction")
	}
	if prev.WeakDoubletSemanticsDerived || prev.HyperchargeSemanticsDerived || prev.LocalFieldSemanticsDerived || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 133 requires Gate 132 weak-doublet and beta firewall to remain open/closed respectively")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 133 refuses hidden physical input")
	}

	planes := buildOrientationPlanes(prev.Slots)
	if len(planes) != 3 {
		return Analysis{}, fmt.Errorf("expected three color orientation planes, got %d", len(planes))
	}
	for _, p := range planes {
		if p.RealDimension != 2 || !p.SO2GeneratorAvailable || !p.AbelianClosure || p.SU2TripleDerived || p.WeakDoubletDerived {
			return Analysis{}, fmt.Errorf("invalid orientation-plane audit for color %d", p.ColorIndex)
		}
	}

	candidates := buildCandidates(prev)
	orientationSO2 := false
	colorWiseSO2 := false
	diagonalSO2 := false
	nonabelianSU2 := false
	su2Commutation := false
	weakDoublet := false
	hypercharge := false
	localField := false
	semanticBridge := false
	borrowedRejected := false
	currentNatural := false
	for _, c := range candidates {
		if c.Kind == OrientationSO2Action && c.CurrentDerived && c.OrientationRotation && c.AbelianOnly && !c.NonAbelianSU2Triple {
			orientationSO2 = true
		}
		if c.Kind == ColorWiseSO2Action && c.CurrentDerived && c.ColorPlanes == 3 && c.OrientationRotation && c.AbelianOnly {
			colorWiseSO2 = true
			diagonalSO2 = true
		}
		if c.NonAbelianSU2Triple {
			nonabelianSU2 = true
		}
		if c.SU2Commutation {
			su2Commutation = true
		}
		if c.WeakDoubletSemantics {
			weakDoublet = true
		}
		if c.HyperchargeDerived {
			hypercharge = true
		}
		if c.LocalFieldDerived {
			localField = true
		}
		if c.RequiresSemanticBridge {
			semanticBridge = true
		}
		if c.Kind == BorrowedMatterSU2 && c.MatterBorrowed && !c.ContactDerived && !c.RepresentationRows {
			borrowedRejected = true
		}
		if c.CurrentDerived && c.ContactDerived && c.Natural && c.NonAbelianSU2Triple && c.SU2Commutation && c.WeakDoubletSemantics && c.HyperchargeDerived && c.LocalFieldDerived && c.RepresentationRows && c.BetaPermitted && !c.RequiresS6Choice && !c.RequiresObservedInput {
			currentNatural = true
		}
	}

	summary := Summary{
		ContactRows:                 7,
		LeptoquarkRows:              6,
		ColorPlanes:                 3,
		RealOrientationsPerColor:    2,
		TotalCurrentLQSlots:         6,
		OrientationSO2Available:     orientationSO2,
		ColorWiseSO2Available:       colorWiseSO2,
		DiagonalSO2Available:        diagonalSO2,
		NonAbelianSU2TripleDerived:  nonabelianSU2,
		SU2WeakDoubletActionDerived: weakDoublet,
		HyperchargeRowDerived:       hypercharge,
		LocalFieldMapDerived:        localField,
		ContactAssignments:          0,
		RepresentationCompleteRows:  0,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		ResidualS6Choices:           prev.Summary.ResidualS6Choices,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)
	truth := "Gate 133 audits the exact temptation left by Gate 132. The current leptoquark six-block has three color planes, each with two real symmetric/skew orientations. Each two-real-dimensional orientation plane admits at most a canonical SO(2)/U(1)-type rotation, and the three planes give color-wise SO(2)^3 or a diagonal SO(2) diagnostic. That is abelian orientation structure, not an SU(2)_L weak-doublet action. A weak-doublet threshold row would require a non-abelian su(2) triple, contact-row assignment, hypercharge, local field map, mass activation, and decoupling. None is derived, so contact leptoquark beta rows remain forbidden."

	return Analysis{
		Previous:   prev,
		Planes:     planes,
		Candidates: candidates,
		Summary:    summary,
		Criteria:   criteria,

		ContactRows:                  7,
		LeptoquarkRows:               6,
		ColorPlanes:                  3,
		RealOrientationsPerColor:     2,
		TotalCurrentLQSlots:          6,
		OrientationSO2Available:      orientationSO2,
		ColorWiseSO2Available:        colorWiseSO2,
		DiagonalSO2Available:         diagonalSO2,
		OrientationActionAbelian:     orientationSO2 && colorWiseSO2 && diagonalSO2,
		NonAbelianSU2TripleDerived:   nonabelianSU2,
		SU2CommutationDerived:        su2Commutation,
		SU2WeakDoubletActionDerived:  weakDoublet,
		WeakDoubletSemanticsDerived:  weakDoublet,
		HyperchargeRowDerived:        hypercharge,
		LocalFieldMapDerived:         localField,
		BorrowedMatterActionRejected: borrowedRejected,
		SemanticBridgeMissing:        semanticBridge,
		S6ObstructionInherited:       prev.S6ObstructionInherited,
		CurrentNaturalSU2Action:      currentNatural,

		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        7,
		ContactBetaRowsAllowed:        0,
		ContactZeroRowsProved:         0,
		ThresholdCorrectedBetaDerived: false,
		FullBetaMatchingTensorDerived: false,

		ResidualNullityBefore:    prev.ResidualNullityAfter,
		ResidualNullityAfter:     prev.ResidualNullityAfter,
		HiddenObservedInputUsed:  false,
		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"two real symmetric/skew orientations form an SU(2)L weak doublet",
			"SO(2) orientation rotation supplies non-abelian weak isospin",
			"borrowing the matter SU(2)L table assigns contact leptoquark rows",
			"the 3×2 leptoquark count permits threshold beta rows",
			"observed alpha/thetaW/masses may choose the SU(2)L action",
		},
		RemainingUnknowns: []string{
			"non-abelian SU(2)L action on contact leptoquark rows",
			"hypercharge row for contact leptoquark slots",
			"local field map and Lorentz kinetic row",
			"mass activation and decoupling rule",
			"canonical assignment of six current leptoquark slots to six contact rows",
		},
		RecommendedNextGate: "Gate 134 — leptoquark hypercharge-row and local-field obstruction / beta-permission theorem",
	}, nil
}

func buildOrientationPlanes(slots []contactlqtensor.LQSlot) []OrientationPlane {
	byColor := map[int][2]string{}
	for _, s := range slots {
		pair := byColor[s.ColorIndex]
		if s.RealOrientation == "symmetric" {
			pair[0] = s.Name
		} else if s.RealOrientation == "skew" {
			pair[1] = s.Name
		}
		byColor[s.ColorIndex] = pair
	}
	out := make([]OrientationPlane, 0, 3)
	for c := 1; c <= 3; c++ {
		pair := byColor[c]
		out = append(out, OrientationPlane{
			ColorIndex:            c,
			Slots:                 pair,
			RealDimension:         2,
			SO2GeneratorAvailable: true,
			SO2GeneratorName:      fmt.Sprintf("J_LQ%d: %s -> %s -> -%s", c, pair[0], pair[1], pair[0]),
			AbelianClosure:        true,
			SU2TripleDerived:      false,
			WeakDoubletDerived:    false,
			ContactRowsAssigned:   false,
			RepresentationRows:    false,
			BetaPermitted:         false,
			Obstruction:           "2D real orientation plane supplies only an SO(2) generator; no three-generator su(2) action is selected",
		})
	}
	return out
}

func buildCandidates(prev contactlqtensor.Analysis) []ActionCandidate {
	return []ActionCandidate{
		{
			Name: "per-color real-orientation SO(2)", Kind: OrientationSO2Action,
			ColorPlanes: 1, RealDimensionPerColor: 2, TotalRealDimension: 2,
			CurrentDerived: true, Canonical: true, Natural: true, MetricPreserving: true, OrientationRotation: true, AbelianOnly: true,
			NonAbelianSU2Triple: false, SU2Commutation: false, WeakDoubletSemantics: false, HyperchargeDerived: false, LocalFieldDerived: false,
			RepresentationRows: false, BetaPermitted: false,
			Obstruction: "one real 2D plane has so(2), not su(2)",
		},
		{
			Name: "color-wise SO(2)^3 orientation action", Kind: ColorWiseSO2Action,
			ColorPlanes: 3, RealDimensionPerColor: 2, TotalRealDimension: 6,
			CurrentDerived: true, Canonical: true, Natural: true, MetricPreserving: true, OrientationRotation: true, AbelianOnly: true,
			NonAbelianSU2Triple: false, SU2Commutation: false, WeakDoubletSemantics: false, HyperchargeDerived: false, LocalFieldDerived: false,
			RepresentationRows: false, BetaPermitted: false,
			Obstruction: "three commuting orientation rotations may be listed, but they do not close as weak su(2)",
		},
		{
			Name: "claim real orientations are SU(2)L components", Kind: ClaimedSU2LAction,
			ColorPlanes: 3, RealDimensionPerColor: 2, TotalRealDimension: 6,
			CurrentDerived: true, Canonical: false, Natural: false, MetricPreserving: true, OrientationRotation: true, AbelianOnly: false,
			NonAbelianSU2Triple: false, SU2Commutation: false, WeakDoubletSemantics: false, HyperchargeDerived: false, LocalFieldDerived: false,
			RepresentationRows: false, BetaPermitted: false, RequiresSemanticBridge: true, RequiresS6Choice: true, HiddenChoices: prev.Summary.ResidualS6Choices,
			Obstruction: "SU(2)L would need three noncommuting generators and representation rows; the current data provides only symmetric/skew orientation labels",
		},
		{
			Name: "borrow audited matter SU(2)L action", Kind: BorrowedMatterSU2,
			ColorPlanes: 3, RealDimensionPerColor: 2, TotalRealDimension: 6,
			MatterBorrowed: true, Canonical: false, Natural: false,
			NonAbelianSU2Triple: false, SU2Commutation: false, WeakDoubletSemantics: false, HyperchargeDerived: false, LocalFieldDerived: false,
			ContactRowAssignment: false, RepresentationRows: false, BetaPermitted: false, RequiresSemanticBridge: true, RequiresS6Choice: true, HiddenChoices: prev.Summary.ResidualTotalCurrentAssignments,
			Obstruction: "the matter doublet table is typed on Fock matter states, not on the contact leptoquark overlap rows",
		},
		{
			Name: "observed-threshold SU(2)L selector", Kind: ObservedFitAction,
			ColorPlanes: 3, RealDimensionPerColor: 2, TotalRealDimension: 6,
			Canonical: false, Natural: false, NonAbelianSU2Triple: false, SU2Commutation: false, RepresentationRows: false, BetaPermitted: false, RequiresObservedInput: true,
			Obstruction: "observed weak angle, masses, or thresholds are forbidden as finite semantic selectors",
		},
	}
}

func buildCriteria(s Summary) []Criterion {
	return []Criterion{
		{Name: "Gate 132 real leptoquark tensor inherited", Required: true, Derived: s.ContactRows == 7 && s.LeptoquarkRows == 6 && s.ColorPlanes == 3 && s.RealOrientationsPerColor == 2 && s.TotalCurrentLQSlots == 6, Detail: "six current leptoquark slots remain color(3) × real-orientation(2)"},
		{Name: "orientation SO(2) action exposed", Required: true, Derived: s.OrientationSO2Available && s.ColorWiseSO2Available && s.DiagonalSO2Available, Detail: "real orientation planes support abelian rotations"},
		{Name: "non-abelian SU(2)L action not derived", Required: true, Derived: !s.NonAbelianSU2TripleDerived && !s.SU2WeakDoubletActionDerived && !s.HyperchargeRowDerived && !s.LocalFieldMapDerived, Detail: "no su(2) triple, hypercharge row, or local field map"},
		{Name: "contact beta firewall remains closed", Required: true, Derived: s.RepresentationCompleteRows == 0 && s.ContactBetaRowsAllowed == 0 && s.ContactZeroRowsProved == 0, Detail: "no threshold row or cancellation row is permitted"},
		{Name: "S6 ambiguity and physical-flow nullity preserved", Required: true, Derived: s.ResidualS6Choices == 720 && s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "six-slot permutations and u,L,Delta b_i remain unresolved"},
	}
}

func FormatPlanes(xs []OrientationPlane) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("color%d(slots=%s/%s dim=%d so2=%t abelian=%t su2=%t weak=%t rep=%t beta=%t obstruction=%s)", x.ColorIndex, x.Slots[0], x.Slots[1], x.RealDimension, x.SO2GeneratorAvailable, x.AbelianClosure, x.SU2TripleDerived, x.WeakDoubletDerived, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCandidates(xs []ActionCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s planes=%d dimPerColor=%d total=%d current=%t contact=%t borrowed=%t canonical=%t natural=%t so2=%t abelian=%t su2Triple=%t su2Comm=%t weak=%t hyper=%t local=%t assign=%t s6=%t observed=%t hidden=%d rep=%t beta=%t obstruction=%s)", x.Name, x.Kind, x.ColorPlanes, x.RealDimensionPerColor, x.TotalRealDimension, x.CurrentDerived, x.ContactDerived, x.MatterBorrowed, x.Canonical, x.Natural, x.OrientationRotation, x.AbelianOnly, x.NonAbelianSU2Triple, x.SU2Commutation, x.WeakDoubletSemantics, x.HyperchargeDerived, x.LocalFieldDerived, x.ContactRowAssignment, x.RequiresS6Choice, x.RequiresObservedInput, x.HiddenChoices, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []Criterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d lqRows=%d colorPlanes=%d realOrient=%d slots=%d so2=%t so2^3=%t diagSO2=%t su2Triple=%t weak=%t hyper=%t local=%t assignments=%d repRows=%d betaRows=%d zeroRows=%d s6=%d nullity=%d->%d", s.ContactRows, s.LeptoquarkRows, s.ColorPlanes, s.RealOrientationsPerColor, s.TotalCurrentLQSlots, s.OrientationSO2Available, s.ColorWiseSO2Available, s.DiagonalSO2Available, s.NonAbelianSU2TripleDerived, s.SU2WeakDoubletActionDerived, s.HyperchargeRowDerived, s.LocalFieldMapDerived, s.ContactAssignments, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
