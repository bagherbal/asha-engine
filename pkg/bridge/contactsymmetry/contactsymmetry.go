// Package contactsymmetry implements Gate 118: contact symmetry-breaking
// selector / stabilizer reduction search.
//
// Gate 117 computed the 168-element Fano automorphism group and proved that it
// acts transitively on points and lines.  Gate 118 asks the next sharper
// question: even though stabilizer subgroups exist after choosing a point, line,
// or flag, does the present finite project contain a legitimate symmetry-
// breaking object that selects such a choice without importing convention or
// observed physics?
//
// The answer is still no.  The stabilizer arithmetic is exact — point and line
// choices reduce Aut(Fano) from 168 to 24, and incident flags reduce it to 8 —
// but those reductions are conditional on an externally chosen object.  The
// current contact-overlap carrier supplies positive eigenvalues and spectral
// labels, yet no Fano-natural selector, contact-side automorphism action,
// representation row, local bundle, or constraint complex.  Therefore no
// contact-Fano assignment is selected, the seven contact rows remain open, and
// beta-threshold permission remains closed.
package contactsymmetry

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactincidence"
	"github.com/bagherbal/asha-engine/pkg/bridge/contactnaturality"
)

type SelectorAttemptKind string

const (
	NoBreakingFullSymmetryAttempt SelectorAttemptKind = "no-breaking-full-aut-fano"
	ChooseFanoPointAttempt        SelectorAttemptKind = "choose-fano-point-stabilizer"
	ChooseFanoLineAttempt         SelectorAttemptKind = "choose-fano-line-stabilizer"
	ChooseIncidentFlagAttempt     SelectorAttemptKind = "choose-incident-flag-stabilizer"
	SpectralContactLabelAttempt   SelectorAttemptKind = "spectral-contact-label-selector"
	SignedFanoOrientationAttempt  SelectorAttemptKind = "signed-fano-orientation-selector"
)

type SelectorStatus string

const (
	SelectorOpen        SelectorStatus = "selector-open"
	SelectorConstructed SelectorStatus = "selector-constructed"
	SelectorForbidden   SelectorStatus = "selector-forbidden"
)

type StabilizerRow struct {
	Name string
	Kind SelectorAttemptKind

	Constructed                  bool
	ChosenObject                 string
	StabilizerOrder              int
	OrbitSize                    int
	UsesExtraConvention          bool
	CanonicalUnderCurrentData    bool
	ContactActionDerived         bool
	SelectsContactFanoAssignment bool
	SelectsRepresentationRow     bool
	BetaRowPermitted             bool
	ZeroRowProved                bool
	RejectedAsPremature          bool
	MissingTerms                 []string
	Detail                       string
}

type ContactRow struct {
	Name, ModeKind string
	Value          float64

	FiniteOverlapPositive       bool
	SurvivesCohomology          bool
	NaturalityOpen              bool
	SymmetrySelectorDerived     bool
	StabilizerAssignmentDerived bool
	RepresentationRowDerived    bool
	CanEnterBetaTensor          bool
	ZeroRowProved               bool

	Status SelectorStatus
	Reason string
}

type Criterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type StabilizerSummary struct {
	FullGroupOrder                  int
	PointStabilizerOrders           []int
	LineStabilizerOrders            []int
	IncidentFlagStabilizerOrders    []int
	NonIncidentFlagStabilizerOrders []int

	PointStabilizerUniform           bool
	LineStabilizerUniform            bool
	IncidentFlagStabilizerUniform    bool
	NonIncidentFlagStabilizerUniform bool

	AnyPointSelected                bool
	AnyLineSelected                 bool
	AnyFlagSelected                 bool
	AnyCanonicalStabilizerReduction bool
}

type Analysis struct {
	ContactNaturality contactnaturality.Analysis

	Rows              []ContactRow
	Stabilizers       []StabilizerRow
	Criteria          []Criterion
	StabilizerSummary StabilizerSummary

	ContactRows                 int
	PositiveFiniteContactRows   int
	SurvivingCohomologyRows     int
	NaturalityOpenRowsInherited int
	OpenContactRowsAfter        int

	FanoAutomorphismGroupOrder   int
	FanoAutomorphismGroupDerived bool
	FanoPointActionTransitive    bool
	FanoLineActionTransitive     bool
	GlobalFixedFanoPoints        int
	GlobalFixedFanoLines         int

	StabilizerArithmeticDerived            bool
	PointStabilizerOrder                   int
	LineStabilizerOrder                    int
	IncidentFlagStabilizerOrder            int
	NonIncidentFlagStabilizerOrder         int
	StabilizerReductionPossibleAfterChoice bool
	CanonicalSymmetryBreakingObjectDerived bool
	CanonicalFanoPointSelected             bool
	CanonicalFanoLineSelected              bool
	CanonicalFanoFlagSelected              bool
	CanonicalContactFanoAssignmentDerived  bool
	ContactAutomorphismActionDerived       bool
	NaturalitySquareFormulable             bool
	SpectralOrderingAvailable              bool
	SpectralOrderingCanonicalForFano       bool
	SignedFanoOrientationBreaksSymmetry    bool

	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	NaturalityObstructionInherited       bool
	IncidenceFunctorObstructionInherited bool
	LocalBundleObstructionInherited      bool
	CohomologyObstructionInherited       bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	PhysicalScaleDerived     bool
	HiddenObservedInputUsed  bool

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
		nat, err := contactnaturality.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(nat)
	})
	return defaultValue, defaultErr
}

func Build(nat contactnaturality.Analysis) (Analysis, error) {
	if !nat.NaturalityObstructionDerived || !nat.FanoAutomorphismGroupDerived || nat.FanoAutomorphismGroupOrder != 168 {
		return Analysis{}, fmt.Errorf("Gate 118 requires Gate 117 automorphism/naturality obstruction")
	}
	if !nat.FanoPointActionTransitive || !nat.FanoLineActionTransitive || nat.GlobalFixedFanoPoints != 0 || nat.GlobalFixedFanoLines != 0 {
		return Analysis{}, fmt.Errorf("Gate 118 requires transitive Fano point/line actions with no fixed selector")
	}
	if nat.ContactRows != 7 || nat.PositiveFiniteContactRows != 7 || nat.SurvivingCohomologyRows != 7 || nat.RepresentationOpenRows != 7 || nat.ContactBetaRowsAllowed != 0 || nat.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 118 requires seven naturality-open contact rows and closed beta permission")
	}
	if nat.ResidualNullityAfter != 3 || nat.HiddenObservedInputUsed || nat.PhysicalWeakAngleDerived || nat.FineStructureDerived || nat.PhysicalMassesDerived || nat.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 118 refuses hidden physical input or changed residual nullity")
	}

	summary := summarizeStabilizers(nat.Automorphisms, nat.ContactIncidence.FanoLines)
	if summary.FullGroupOrder != 168 || !summary.PointStabilizerUniform || !summary.LineStabilizerUniform || !summary.IncidentFlagStabilizerUniform {
		return Analysis{}, fmt.Errorf("unexpected stabilizer arithmetic: %+v", summary)
	}

	rows := buildRows(nat.Rows)
	stabs := buildStabilizerRows(summary)
	criteria := buildCriteria(summary)
	counts := countRows(rows)

	truth := "Gate 118 derives the stabilizer arithmetic of the exact 168-element Fano automorphism group: choosing a point or line gives a 24-element stabilizer and choosing an incident point-line flag gives an 8-element stabilizer. But every such reduction depends on first choosing the point, line, or flag. The current finite contact data supplies no canonical symmetry-breaking object, no derived contact-side automorphism action, and no natural contact-Fano assignment. Spectral labels and signed incidence are compatible diagnostics, not selectors. Therefore the seven contact-overlap rows remain open and no threshold beta correction is permitted."

	return Analysis{
		ContactNaturality: nat,
		Rows:              rows,
		Stabilizers:       stabs,
		Criteria:          criteria,
		StabilizerSummary: summary,

		ContactRows:                 counts.contact,
		PositiveFiniteContactRows:   counts.positive,
		SurvivingCohomologyRows:     counts.surviving,
		NaturalityOpenRowsInherited: nat.RepresentationOpenRows,
		OpenContactRowsAfter:        counts.open,

		FanoAutomorphismGroupOrder:   nat.FanoAutomorphismGroupOrder,
		FanoAutomorphismGroupDerived: true,
		FanoPointActionTransitive:    nat.FanoPointActionTransitive,
		FanoLineActionTransitive:     nat.FanoLineActionTransitive,
		GlobalFixedFanoPoints:        nat.GlobalFixedFanoPoints,
		GlobalFixedFanoLines:         nat.GlobalFixedFanoLines,

		StabilizerArithmeticDerived:            true,
		PointStabilizerOrder:                   first(summary.PointStabilizerOrders),
		LineStabilizerOrder:                    first(summary.LineStabilizerOrders),
		IncidentFlagStabilizerOrder:            first(summary.IncidentFlagStabilizerOrders),
		NonIncidentFlagStabilizerOrder:         first(summary.NonIncidentFlagStabilizerOrders),
		StabilizerReductionPossibleAfterChoice: true,
		CanonicalSymmetryBreakingObjectDerived: false,
		CanonicalFanoPointSelected:             false,
		CanonicalFanoLineSelected:              false,
		CanonicalFanoFlagSelected:              false,
		CanonicalContactFanoAssignmentDerived:  false,
		ContactAutomorphismActionDerived:       false,
		NaturalitySquareFormulable:             false,
		SpectralOrderingAvailable:              true,
		SpectralOrderingCanonicalForFano:       false,
		SignedFanoOrientationBreaksSymmetry:    false,

		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              counts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		NaturalityObstructionInherited:       true,
		IncidenceFunctorObstructionInherited: nat.IncidenceFunctorObstructionInherited,
		LocalBundleObstructionInherited:      nat.LocalBundleObstructionInherited,
		CohomologyObstructionInherited:       nat.CohomologyObstructionInherited,

		ResidualNullityBefore:  3,
		ResidualNullityAfter:   3,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"stabilizer existence alone selects a physical contact-Fano assignment",
			"spectral ordering breaks Fano symmetry canonically",
			"signed Fano incidence alone selects one contact mode or chart",
			"contact modes may now enter threshold beta matching",
			"contact stabilizer reduction predicts alpha, thetaW, M*, g_*, or masses",
		},
		RemainingUnknowns: []string{
			"canonical finite symmetry-breaking object",
			"derived action of Aut(Fano) on the contact-overlap carrier",
			"natural contact-to-Fano assignment",
			"representation rows for contact modes",
			"constraint/zero-row proof for contact modes",
			"threshold-corrected beta tensor",
		},
		RecommendedNextGate: "Gate 119 — contact-side automorphism action construction / equivariant assignment search",
	}, nil
}

func buildRows(rows []contactnaturality.NaturalityRow) []ContactRow {
	out := make([]ContactRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, ContactRow{
			Name:                        r.Name,
			ModeKind:                    r.ModeKind,
			Value:                       r.Value,
			FiniteOverlapPositive:       r.FiniteOverlapPositive,
			SurvivesCohomology:          r.SurvivesCohomology,
			NaturalityOpen:              r.NaturalityStatus == contactnaturality.NaturalityOpen,
			SymmetrySelectorDerived:     false,
			StabilizerAssignmentDerived: false,
			RepresentationRowDerived:    false,
			CanEnterBetaTensor:          false,
			ZeroRowProved:               false,
			Status:                      SelectorOpen,
			Reason:                      "no canonical symmetry-breaking selector or stabilizer assignment has been derived for this contact row",
		})
	}
	return out
}

func buildStabilizerRows(s StabilizerSummary) []StabilizerRow {
	return []StabilizerRow{
		{
			Name:                      "keep full Fano automorphism symmetry",
			Kind:                      NoBreakingFullSymmetryAttempt,
			Constructed:               true,
			ChosenObject:              "none",
			StabilizerOrder:           s.FullGroupOrder,
			OrbitSize:                 1,
			CanonicalUnderCurrentData: true,
			MissingTerms:              []string{"selected point/line/flag", "contact-Fano assignment", "representation row"},
			Detail:                    "full symmetry is canonical but selects no contact row and no threshold representation",
		},
		{
			Name:                "choose one Fano point",
			Kind:                ChooseFanoPointAttempt,
			Constructed:         true,
			ChosenObject:        "point p",
			StabilizerOrder:     first(s.PointStabilizerOrders),
			OrbitSize:           s.FullGroupOrder / first(s.PointStabilizerOrders),
			UsesExtraConvention: true,
			RejectedAsPremature: true,
			MissingTerms:        []string{"canonical point selector", "contact-side action", "representation row"},
			Detail:              "a chosen point has stabilizer 24, but the current finite system selects no point",
		},
		{
			Name:                "choose one Fano line",
			Kind:                ChooseFanoLineAttempt,
			Constructed:         true,
			ChosenObject:        "line L",
			StabilizerOrder:     first(s.LineStabilizerOrders),
			OrbitSize:           s.FullGroupOrder / first(s.LineStabilizerOrders),
			UsesExtraConvention: true,
			RejectedAsPremature: true,
			MissingTerms:        []string{"canonical line selector", "chart atlas", "transition/cocycle law"},
			Detail:              "a chosen line has stabilizer 24, but the current finite system selects no line",
		},
		{
			Name:                "choose one incident point-line flag",
			Kind:                ChooseIncidentFlagAttempt,
			Constructed:         true,
			ChosenObject:        "p ∈ L",
			StabilizerOrder:     first(s.IncidentFlagStabilizerOrders),
			OrbitSize:           s.FullGroupOrder / first(s.IncidentFlagStabilizerOrders),
			UsesExtraConvention: true,
			RejectedAsPremature: true,
			MissingTerms:        []string{"canonical flag selector", "fiber chart law", "gauge row"},
			Detail:              "an incident flag has stabilizer 8, but the flag must be chosen by convention",
		},
		{
			Name:                "spectral contact label as selector",
			Kind:                SpectralContactLabelAttempt,
			Constructed:         true,
			ChosenObject:        "ordered contact eigenvalue label",
			UsesExtraConvention: true,
			RejectedAsPremature: true,
			MissingTerms:        []string{"Fano-natural orientation", "equivariance proof", "local field map"},
			Detail:              "spectral ordering labels contact modes, but Gate 117 showed this is not Fano-natural",
		},
		{
			Name:                      "signed Fano orientation as selector",
			Kind:                      SignedFanoOrientationAttempt,
			Constructed:               true,
			ChosenObject:              "oriented signed incidence",
			UsesExtraConvention:       false,
			CanonicalUnderCurrentData: true,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"unique point/line/flag fixed by signed data", "contact-side representation"},
			Detail:                    "the signed Fano form is finite structure, but it does not by itself choose one contact mode or beta row",
		},
	}
}

func buildCriteria(s StabilizerSummary) []Criterion {
	return []Criterion{
		{Name: "full Aut(Fano) derived", Required: true, Derived: s.FullGroupOrder == 168, Detail: "Gate 117 symmetry data inherited"},
		{Name: "stabilizer arithmetic computed", Required: true, Derived: first(s.PointStabilizerOrders) == 24 && first(s.LineStabilizerOrders) == 24 && first(s.IncidentFlagStabilizerOrders) == 8, Detail: FormatStabilizerSummary(s)},
		{Name: "canonical point selector", Required: true, Derived: false, Detail: "no global fixed point exists"},
		{Name: "canonical line selector", Required: true, Derived: false, Detail: "no global fixed line exists"},
		{Name: "canonical contact-side action", Required: true, Derived: false, Detail: "no action of Aut(Fano) on contact-overlap modes has been derived"},
		{Name: "beta permission", Required: true, Derived: false, Detail: "no representation-complete or zero-row contact branch is selected"},
	}
}

type rowCounts struct{ contact, positive, surviving, open int }

func countRows(rows []ContactRow) rowCounts {
	c := rowCounts{contact: len(rows)}
	for _, r := range rows {
		if r.FiniteOverlapPositive {
			c.positive++
		}
		if r.SurvivesCohomology {
			c.surviving++
		}
		if r.Status == SelectorOpen {
			c.open++
		}
	}
	return c
}

func summarizeStabilizers(perms []contactnaturality.Permutation, lines []contactincidence.FanoLine) StabilizerSummary {
	pointOrders := make([]int, 7)
	lineOrders := make([]int, len(lines))
	incidentOrders := []int{}
	nonIncidentOrders := []int{}
	for p := 0; p < 7; p++ {
		pointOrders[p] = pointStabilizerOrder(perms, p)
	}
	for i, line := range lines {
		lineOrders[i] = lineStabilizerOrder(perms, line)
		linePts := map[int]bool{line.Points[0]: true, line.Points[1]: true, line.Points[2]: true}
		for p := 0; p < 7; p++ {
			if linePts[p] {
				incidentOrders = append(incidentOrders, flagStabilizerOrder(perms, p, line))
			} else {
				nonIncidentOrders = append(nonIncidentOrders, flagStabilizerOrder(perms, p, line))
			}
		}
	}
	sort.Ints(incidentOrders)
	sort.Ints(nonIncidentOrders)
	return StabilizerSummary{
		FullGroupOrder:                   len(perms),
		PointStabilizerOrders:            pointOrders,
		LineStabilizerOrders:             lineOrders,
		IncidentFlagStabilizerOrders:     incidentOrders,
		NonIncidentFlagStabilizerOrders:  nonIncidentOrders,
		PointStabilizerUniform:           uniform(pointOrders),
		LineStabilizerUniform:            uniform(lineOrders),
		IncidentFlagStabilizerUniform:    uniform(incidentOrders),
		NonIncidentFlagStabilizerUniform: uniform(nonIncidentOrders),
		AnyPointSelected:                 false,
		AnyLineSelected:                  false,
		AnyFlagSelected:                  false,
		AnyCanonicalStabilizerReduction:  false,
	}
}

func pointStabilizerOrder(perms []contactnaturality.Permutation, point int) int {
	n := 0
	for _, p := range perms {
		if p[point] == point {
			n++
		}
	}
	return n
}

func lineStabilizerOrder(perms []contactnaturality.Permutation, line contactincidence.FanoLine) int {
	n := 0
	for _, p := range perms {
		if mapsLineToSameLine(p, line) {
			n++
		}
	}
	return n
}

func flagStabilizerOrder(perms []contactnaturality.Permutation, point int, line contactincidence.FanoLine) int {
	n := 0
	for _, p := range perms {
		if p[point] == point && mapsLineToSameLine(p, line) {
			n++
		}
	}
	return n
}

func mapsLineToSameLine(p contactnaturality.Permutation, line contactincidence.FanoLine) bool {
	src := []int{line.Points[0], line.Points[1], line.Points[2]}
	dst := []int{p[line.Points[0]], p[line.Points[1]], p[line.Points[2]]}
	sort.Ints(src)
	sort.Ints(dst)
	return src[0] == dst[0] && src[1] == dst[1] && src[2] == dst[2]
}

func uniform(xs []int) bool {
	if len(xs) == 0 {
		return false
	}
	for _, x := range xs[1:] {
		if x != xs[0] {
			return false
		}
	}
	return true
}

func first(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	return xs[0]
}

func FormatRows(rows []ContactRow, limit int) string {
	parts := make([]string, 0, len(rows))
	for i, r := range rows {
		if limit > 0 && i >= limit {
			parts = append(parts, fmt.Sprintf("... +%d", len(rows)-i))
			break
		}
		parts = append(parts, fmt.Sprintf("%s(selector=%t,rep=%t,beta=%t)", r.Name, r.SymmetrySelectorDerived, r.RepresentationRowDerived, r.CanEnterBetaTensor))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatStabilizerSummary(s StabilizerSummary) string {
	return fmt.Sprintf("|Aut|=%d; pointStab=%s; lineStab=%s; incidentFlagStab=%s; nonIncidentFlagStab=%s", s.FullGroupOrder, compactInts(s.PointStabilizerOrders), compactInts(s.LineStabilizerOrders), compactInts(s.IncidentFlagStabilizerOrders), compactInts(s.NonIncidentFlagStabilizerOrders))
}

func FormatStabilizers(xs []StabilizerRow) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s,stab=%d,convention=%t,canonical=%t,beta=%t)", x.Name, x.Kind, x.StabilizerOrder, x.UsesExtraConvention, x.CanonicalUnderCurrentData, x.BetaRowPermitted))
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

func Join(xs []string) string { return strings.Join(xs, "; ") }

func compactInts(xs []int) string {
	if len(xs) == 0 {
		return "[]"
	}
	counts := map[int]int{}
	keys := []int{}
	for _, x := range xs {
		if counts[x] == 0 {
			keys = append(keys, x)
		}
		counts[x]++
	}
	sort.Ints(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		if counts[k] == 1 {
			parts = append(parts, fmt.Sprintf("%d", k))
		} else {
			parts = append(parts, fmt.Sprintf("%d×%d", k, counts[k]))
		}
	}
	return "[" + strings.Join(parts, ",") + "]"
}
