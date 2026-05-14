// Package contactlqtensor implements Gate 132: contact leptoquark slot
// representation tensor / color-doublet semantic obstruction theorem.
//
// Gate 131 isolated the S6 ambiguity in assigning the six current-side
// leptoquark slots to the six non-singlet contact rows. Gate 132 asks whether
// the leptoquark slots themselves contain enough typed representation data to
// break that ambiguity. The answer remains negative: the current inventory has
// six real off-diagonal lepton-color generators, naturally organized as
// color(3) × real-orientation(2). This is not the same thing as a physical
// color triplet weak doublet. Interpreting the two real orientations as weak
// doublet components would be a semantic jump, not a theorem.
package contactlqtensor

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactlqblock"
)

type TensorKind string

const (
	CurrentLQRealTensor TensorKind = "current-leptoquark-real-tensor"
	ColorWeakDoublet    TensorKind = "color-weak-doublet-interpretation"
	SpectralSixTensor   TensorKind = "spectral-six-contact-tensor"
	FanoTransportTensor TensorKind = "fano-transported-six-tensor"
	ObservedFitTensor   TensorKind = "observed-fit-tensor"
)

type LQSlot struct {
	Name                 string
	ColorIndex           int
	RealOrientation      string
	CurrentDerived       bool
	ColorSemantics       bool
	WeakDoubletSemantics bool
	ContactRowAssigned   bool
	RepresentationRow    bool
	BetaPermitted        bool
	Reason               string
}

type TensorCandidate struct {
	Name string
	Kind TensorKind

	SlotCount              int
	ColorSlots             int
	SecondFactorSlots      int
	SecondFactorSemantics  string
	CurrentDerived         bool
	ContactDerived         bool
	FanoDerived            bool
	Canonical              bool
	Natural                bool
	ColorTripletDerived    bool
	WeakDoubletDerived     bool
	HyperchargeDerived     bool
	LocalFieldDerived      bool
	RepresentationRows     bool
	BetaPermitted          bool
	RequiresS6Choice       bool
	RequiresSingletChoice  bool
	RequiresFanoChoice     bool
	RequiresObservedInput  bool
	RequiresSemanticBridge bool
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
	ContactRows                     int
	LeptoquarkRows                  int
	CurrentLQSlots                  int
	ColorSlots                      int
	RealOrientationSlots            int
	ColorWeakCountMatch             bool
	WeakDoubletSemanticsDerived     bool
	HyperchargeSemanticsDerived     bool
	LocalFieldSemanticsDerived      bool
	CurrentNaturalRepresentation    bool
	ContactRowAssignments           int
	RepresentationCompleteRows      int
	ContactBetaRowsAllowed          int
	ContactZeroRowsProved           int
	ResidualS6Choices               int
	ResidualSingletChoices          int
	ResidualAssignmentsPerBranch    int
	ResidualTotalCurrentAssignments int
	ResidualNullityBefore           int
	ResidualNullityAfter            int
}

type Analysis struct {
	Previous contactlqblock.Analysis

	Slots      []LQSlot
	Candidates []TensorCandidate
	Criteria   []Criterion
	Summary    Summary

	ContactRows                  int
	LeptoquarkRows               int
	CurrentLQSlots               int
	ColorSlots                   int
	RealOrientationSlots         int
	ColorWeakCountMatch          bool
	CurrentRealTensorDerived     bool
	ColorTripletSemantics        bool
	RealOrientationSemantics     bool
	WeakDoubletSemanticsDerived  bool
	HyperchargeSemanticsDerived  bool
	LocalFieldSemanticsDerived   bool
	SemanticBridgeMissing        bool
	ColorDoubletCountTrap        bool
	S6ObstructionInherited       bool
	CurrentNaturalRepresentation bool

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
		prev, err := contactlqblock.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactlqblock.Analysis) (Analysis, error) {
	if !prev.S6PermutationObstruction || prev.LeptoquarkRows != 6 || prev.SixPermutationOrder != 720 || prev.AssignmentsPerBranch != 5040 {
		return Analysis{}, fmt.Errorf("Gate 132 requires Gate 131 S6 six-block obstruction")
	}
	if prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 132 requires Gate 131 beta firewall to remain closed")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 132 refuses hidden physical input")
	}

	slots := buildLQSlots()
	if len(slots) != 6 {
		return Analysis{}, fmt.Errorf("expected six leptoquark slots, got %d", len(slots))
	}
	candidates := buildCandidates(prev)
	currentReal := false
	colorTriplet := false
	weakDoublet := false
	hypercharge := false
	localField := false
	semanticBridgeMissing := false
	colorDoubletTrap := false
	currentNaturalRepresentation := false
	for _, c := range candidates {
		if c.Kind == CurrentLQRealTensor && c.CurrentDerived && c.ColorTripletDerived && !c.WeakDoubletDerived {
			currentReal = true
			colorTriplet = true
		}
		if c.WeakDoubletDerived && c.HyperchargeDerived && c.LocalFieldDerived && c.RepresentationRows && c.BetaPermitted && !c.RequiresSemanticBridge && !c.RequiresS6Choice && !c.RequiresObservedInput {
			currentNaturalRepresentation = true
		}
		if c.WeakDoubletDerived {
			weakDoublet = true
		}
		if c.HyperchargeDerived {
			hypercharge = true
		}
		if c.LocalFieldDerived {
			localField = true
		}
		if c.RequiresSemanticBridge {
			semanticBridgeMissing = true
		}
		if c.Kind == ColorWeakDoublet && c.SlotCount == 6 && c.ColorSlots == 3 && c.SecondFactorSlots == 2 && c.RequiresSemanticBridge {
			colorDoubletTrap = true
		}
	}

	summary := Summary{
		ContactRows:                     7,
		LeptoquarkRows:                  6,
		CurrentLQSlots:                  6,
		ColorSlots:                      3,
		RealOrientationSlots:            2,
		ColorWeakCountMatch:             true,
		WeakDoubletSemanticsDerived:     weakDoublet,
		HyperchargeSemanticsDerived:     hypercharge,
		LocalFieldSemanticsDerived:      localField,
		CurrentNaturalRepresentation:    currentNaturalRepresentation,
		ContactRowAssignments:           0,
		RepresentationCompleteRows:      0,
		ContactBetaRowsAllowed:          0,
		ContactZeroRowsProved:           0,
		ResidualS6Choices:               720,
		ResidualSingletChoices:          7,
		ResidualAssignmentsPerBranch:    5040,
		ResidualTotalCurrentAssignments: 10080,
		ResidualNullityBefore:           prev.ResidualNullityAfter,
		ResidualNullityAfter:            prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)
	truth := "Gate 132 tests the tempting 3×2 leptoquark-slot interpretation. The current carrier really has six off-diagonal lepton-color generators, organized as three color directions times two real orientations (symmetric/skew). That count equals color(3)×doublet(2), but the second factor is not derived as SU(2)L weak doublet semantics, and no hypercharge, local-field, mass-activation, or decoupling row is selected. Therefore the six leptoquark slots remain current-sector representation data, not contact threshold beta rows."

	return Analysis{
		Previous:   prev,
		Slots:      slots,
		Candidates: candidates,
		Criteria:   criteria,
		Summary:    summary,

		ContactRows:                  7,
		LeptoquarkRows:               6,
		CurrentLQSlots:               6,
		ColorSlots:                   3,
		RealOrientationSlots:         2,
		ColorWeakCountMatch:          true,
		CurrentRealTensorDerived:     currentReal,
		ColorTripletSemantics:        colorTriplet,
		RealOrientationSemantics:     true,
		WeakDoubletSemanticsDerived:  weakDoublet,
		HyperchargeSemanticsDerived:  hypercharge,
		LocalFieldSemanticsDerived:   localField,
		SemanticBridgeMissing:        semanticBridgeMissing,
		ColorDoubletCountTrap:        colorDoubletTrap,
		S6ObstructionInherited:       prev.S6PermutationObstruction,
		CurrentNaturalRepresentation: currentNaturalRepresentation,

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
			"the leptoquark six-block is automatically a color triplet weak doublet",
			"symmetric/skew real orientations are SU(2)L components",
			"six current-side leptoquark slots canonically assign the six contact rows",
			"current leptoquark generators supply contact threshold beta rows",
			"observed constants may decide the six-slot semantic bridge",
		},
		RemainingUnknowns: []string{
			"semantic bridge from real leptoquark orientations to weak-doublet components",
			"hypercharge row for contact leptoquark modes",
			"local field map and Lorentz kinetic row for contact modes",
			"mass activation and decoupling rule for threshold beta matching",
			"canonical assignment of current leptoquark slots to contact rows",
		},
		RecommendedNextGate: "Gate 133 — leptoquark real-orientation versus weak-doublet obstruction / SU(2)L action search",
	}, nil
}

func buildLQSlots() []LQSlot {
	out := make([]LQSlot, 0, 6)
	for c := 1; c <= 3; c++ {
		out = append(out, LQSlot{
			Name:                 fmt.Sprintf("LQ%d-sym", c),
			ColorIndex:           c,
			RealOrientation:      "symmetric",
			CurrentDerived:       true,
			ColorSemantics:       true,
			WeakDoubletSemantics: false,
			ContactRowAssigned:   false,
			RepresentationRow:    false,
			BetaPermitted:        false,
			Reason:               "derived as a real lepton-color off-diagonal current generator, not as a contact weak-doublet threshold row",
		})
		out = append(out, LQSlot{
			Name:                 fmt.Sprintf("LQ%d-skew", c),
			ColorIndex:           c,
			RealOrientation:      "skew",
			CurrentDerived:       true,
			ColorSemantics:       true,
			WeakDoubletSemantics: false,
			ContactRowAssigned:   false,
			RepresentationRow:    false,
			BetaPermitted:        false,
			Reason:               "the second real orientation completes a real off-diagonal pair; no SU(2)L action on the pair is selected",
		})
	}
	return out
}

func buildCandidates(prev contactlqblock.Analysis) []TensorCandidate {
	return []TensorCandidate{
		{
			Name: "derived current leptoquark real tensor", Kind: CurrentLQRealTensor,
			SlotCount: 6, ColorSlots: 3, SecondFactorSlots: 2, SecondFactorSemantics: "real symmetric/skew off-diagonal orientation",
			CurrentDerived: true, Canonical: true, Natural: true, ColorTripletDerived: true, WeakDoubletDerived: false, HyperchargeDerived: false, LocalFieldDerived: false,
			RepresentationRows: false, BetaPermitted: false,
			Obstruction: "valid current-sector tensor, but no contact-row assignment, local field, weak-doublet action, or hypercharge row is derived",
		},
		{
			Name: "interpret real orientation as weak doublet", Kind: ColorWeakDoublet,
			SlotCount: 6, ColorSlots: 3, SecondFactorSlots: 2, SecondFactorSemantics: "claimed weak doublet",
			CurrentDerived: true, Canonical: false, Natural: false, ColorTripletDerived: true, WeakDoubletDerived: false, HyperchargeDerived: false, LocalFieldDerived: false,
			RepresentationRows: false, BetaPermitted: false, RequiresSemanticBridge: true, RequiresS6Choice: true, HiddenChoices: prev.SixPermutationOrder,
			Obstruction: "3×2 count matches, but the two current orientations are not SU(2)L components; a semantic bridge and contact-row assignment are missing",
		},
		{
			Name: "spectral six-contact tensor", Kind: SpectralSixTensor,
			SlotCount: 6, ColorSlots: 0, SecondFactorSlots: 6, SecondFactorSemantics: "six chosen contact spectral values after an external singlet choice",
			ContactDerived: true, Canonical: false, Natural: false, ColorTripletDerived: false, WeakDoubletDerived: false, HyperchargeDerived: false, LocalFieldDerived: false,
			RepresentationRows: false, BetaPermitted: false, RequiresSingletChoice: true, RequiresS6Choice: true, HiddenChoices: prev.AssignmentsPerBranch,
			Obstruction: "contact spectral rows can be listed, but they do not carry color, weak, or hypercharge semantics",
		},
		{
			Name: "Fano-transported leptoquark tensor", Kind: FanoTransportTensor,
			SlotCount: 6, ColorSlots: 3, SecondFactorSlots: 2, SecondFactorSemantics: "transported by hidden contact-to-Fano labeling",
			FanoDerived: true, Canonical: false, Natural: false, ColorTripletDerived: false, WeakDoubletDerived: false, HyperchargeDerived: false, LocalFieldDerived: false,
			RepresentationRows: false, BetaPermitted: false, RequiresFanoChoice: true, RequiresS6Choice: true, HiddenChoices: prev.AssignmentsPerBranch,
			Obstruction: "requires a hidden contact-to-Fano assignment before Fano labels can decorate contact rows",
		},
		{
			Name: "observed-constant leptoquark tensor", Kind: ObservedFitTensor,
			SlotCount: 6, ColorSlots: 3, SecondFactorSlots: 2, SecondFactorSemantics: "fitted from observed thresholds",
			Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false, RequiresObservedInput: true,
			Obstruction: "observed alpha/thetaW/masses/thresholds cannot be used to select finite contact leptoquark semantics",
		},
	}
}

func buildCriteria(s Summary) []Criterion {
	return []Criterion{
		{Name: "Gate 131 S6 obstruction inherited", Required: true, Derived: s.ContactRows == 7 && s.LeptoquarkRows == 6 && s.ResidualS6Choices == 720 && s.ResidualAssignmentsPerBranch == 5040, Detail: "six contact rows remain unassigned to leptoquark slots"},
		{Name: "current leptoquark tensor is typed but current-side", Required: true, Derived: s.CurrentLQSlots == 6 && s.ColorSlots == 3 && s.RealOrientationSlots == 2 && s.ColorWeakCountMatch, Detail: "six real off-diagonal slots equal 3 colors times 2 real orientations"},
		{Name: "color-doublet count trap rejected", Required: true, Derived: !s.WeakDoubletSemanticsDerived && !s.HyperchargeSemanticsDerived && !s.LocalFieldSemanticsDerived, Detail: "the 2 factor is not derived as SU(2)L weak-doublet semantics"},
		{Name: "no current-natural representation tensor for contact rows", Required: true, Derived: !s.CurrentNaturalRepresentation && s.ContactRowAssignments == 0 && s.RepresentationCompleteRows == 0, Detail: "no color×weak×hypercharge row is attached to contact modes"},
		{Name: "beta firewall remains closed", Required: true, Derived: s.ContactBetaRowsAllowed == 0 && s.ContactZeroRowsProved == 0, Detail: "no contact threshold beta row or cancellation row is permitted"},
		{Name: "physical-flow nullity preserved", Required: true, Derived: s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "u, L, and Delta b_i(L) remain free"},
	}
}

func FormatSlots(xs []LQSlot) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(color=%d orient=%s current=%t colorSem=%t weakSem=%t contact=%t rep=%t beta=%t)", x.Name, x.ColorIndex, x.RealOrientation, x.CurrentDerived, x.ColorSemantics, x.WeakDoubletSemantics, x.ContactRowAssigned, x.RepresentationRow, x.BetaPermitted))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCandidates(xs []TensorCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s slots=%d color=%d second=%d secondSem=%s current=%t contact=%t fano=%t canonical=%t natural=%t colorTriplet=%t weakDoublet=%t hyper=%t local=%t s6=%t fanoChoice=%t observed=%t hidden=%d rep=%t beta=%t obstruction=%s)", x.Name, x.Kind, x.SlotCount, x.ColorSlots, x.SecondFactorSlots, x.SecondFactorSemantics, x.CurrentDerived, x.ContactDerived, x.FanoDerived, x.Canonical, x.Natural, x.ColorTripletDerived, x.WeakDoubletDerived, x.HyperchargeDerived, x.LocalFieldDerived, x.RequiresS6Choice, x.RequiresFanoChoice, x.RequiresObservedInput, x.HiddenChoices, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
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
	return fmt.Sprintf("contact=%d lqRows=%d currentLQ=%d color=%d realOrient=%d countMatch=%t weakDoublet=%t hyper=%t local=%t currentNaturalRep=%t contactAssignments=%d repRows=%d betaRows=%d zeroRows=%d s6=%d singlet=%d perBranch=%d total=%d nullity=%d->%d", s.ContactRows, s.LeptoquarkRows, s.CurrentLQSlots, s.ColorSlots, s.RealOrientationSlots, s.ColorWeakCountMatch, s.WeakDoubletSemanticsDerived, s.HyperchargeSemanticsDerived, s.LocalFieldSemanticsDerived, s.CurrentNaturalRepresentation, s.ContactRowAssignments, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualSingletChoices, s.ResidualAssignmentsPerBranch, s.ResidualTotalCurrentAssignments, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
