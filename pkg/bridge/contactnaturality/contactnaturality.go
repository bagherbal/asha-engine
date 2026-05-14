// Package contactnaturality implements Gate 117: contact-Fano naturality
// obstruction / automorphism-invariance theorem.
//
// Gate 116 showed that the Fano incidence carrier is exact and cardinality-
// matched to the seven unresolved contact partial-overlap modes, but that a
// contact-to-Fano map would require selecting one of 7! bijections. Gate 117
// tests whether finite symmetry removes that convention: does the Fano
// automorphism group, together with the currently available contact data,
// select an invariant bijection, point, line, chart, or representation row?
//
// The result is a strict obstruction. The Fano automorphism group is computed
// directly from the seven lines and has order 168. It acts transitively on
// points and lines, so no point/line is globally fixed. With no derived action
// of that group on the contact-overlap carrier, an equivariant bijection cannot
// be formulated; if the contact side is kept with only the known trivial action,
// no nonconstant bijection can be invariant under the nontrivial Fano action.
// Therefore the contact-Fano assignment remains convention-dependent and the
// contact beta firewall stays closed.
package contactnaturality

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactincidence"
)

type NaturalityAttemptKind string

const (
	FullFanoAutomorphismAudit    NaturalityAttemptKind = "full-fano-automorphism-audit"
	InvariantPointAssignment     NaturalityAttemptKind = "invariant-fano-point-assignment"
	EquivariantBijectionAttempt  NaturalityAttemptKind = "equivariant-contact-fano-bijection"
	SpectralLabelBreakingAttempt NaturalityAttemptKind = "spectral-label-automorphism-breaking"
	LineOrbitChartAttempt        NaturalityAttemptKind = "line-orbit-chart-selection"
)

type ContactNaturalityStatus string

const (
	NaturalityOpen      ContactNaturalityStatus = "naturality-open"
	NaturalityInvariant ContactNaturalityStatus = "naturality-invariant"
	NaturalityForbidden ContactNaturalityStatus = "naturality-forbidden"
)

type Permutation [7]int

type AutomorphismSummary struct {
	Order                    int
	IdentityCount            int
	NonIdentityCount         int
	PointOrbitCount          int
	LineOrbitCount           int
	PointOrbitSizes          []int
	LineOrbitSizes           []int
	CommonFixedPoints        []int
	CommonFixedLines         []int
	PointTransitive          bool
	LineTransitive           bool
	EveryLinePreserved       bool
	RepresentativeNontrivial Permutation
}

type NaturalityRow struct {
	Name, ModeKind string
	Value          float64

	FiniteOverlapPositive bool
	SurvivesCohomology    bool
	IncidenceOpen         bool

	FanoAssignmentInvariant          bool
	EquivariantBijectionDerived      bool
	ContactAutomorphismActionDerived bool
	FanoAutomorphismInvariant        bool
	NaturalitySquareDerived          bool
	RepresentationRowDerived         bool
	CanEnterBetaTensor               bool

	NaturalityStatus ContactNaturalityStatus
	Reason           string
}

type NaturalityAttempt struct {
	Name string
	Kind NaturalityAttemptKind

	Constructed                     bool
	UsesExtraConvention             bool
	InvariantUnderFanoAutomorphisms bool
	ContactActionDerived            bool
	EquivariantMapDerived           bool
	CanonicalSelectorDerived        bool
	RepresentationRowDerived        bool
	BetaRowPermitted                bool
	ZeroRowProved                   bool
	RejectedAsPremature             bool
	MissingTerms                    []string
	Detail                          string
}

type NaturalCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Analysis struct {
	ContactIncidence contactincidence.Analysis

	Rows                []NaturalityRow
	Attempts            []NaturalityAttempt
	Criteria            []NaturalCriterion
	Automorphisms       []Permutation
	AutomorphismSummary AutomorphismSummary

	ContactRows                int
	PositiveFiniteContactRows  int
	SurvivingCohomologyRows    int
	IncidenceOpenRowsInherited int
	OpenContactRowsAfter       int

	FanoIncidenceAvailable             bool
	FanoAutomorphismGroupDerived       bool
	FanoAutomorphismGroupOrder         int
	FanoAutomorphismGroupNontrivial    bool
	FanoPointActionTransitive          bool
	FanoLineActionTransitive           bool
	GlobalFixedFanoPoints              int
	GlobalFixedFanoLines               int
	AutomorphismInvariantPointSelector bool
	AutomorphismInvariantLineSelector  bool

	ContactAutomorphismActionDerived   bool
	NaturalitySquareFormulable         bool
	InvariantContactToFanoMapDerived   bool
	EquivariantBijectionDerived        bool
	CanonicalAssignmentCount           int
	CompatibleBijectionCount           int
	ConventionDependentBijections      int
	SpectralOrderingBreaksAutomorphism bool

	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	NaturalityObstructionDerived         bool
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
		incidence, err := contactincidence.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(incidence)
	})
	return defaultValue, defaultErr
}

func Build(incidence contactincidence.Analysis) (Analysis, error) {
	if !incidence.IncidenceFunctorObstructionDerived || !incidence.FanoIncidenceAvailable || !incidence.FanoIncidenceResonance {
		return Analysis{}, fmt.Errorf("Gate 117 requires Gate 116 Fano/contact incidence obstruction")
	}
	if incidence.ContactRows != 7 || incidence.PositiveFiniteContactRows != 7 || incidence.SurvivingCohomologyRows != 7 || incidence.RepresentationOpenRows != 7 || incidence.ContactBetaRowsAllowed != 0 || incidence.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 117 requires seven incidence-open contact rows and closed beta permission")
	}
	if incidence.ResidualNullityAfter != 3 || incidence.HiddenObservedInputUsed || incidence.PhysicalWeakAngleDerived || incidence.FineStructureDerived || incidence.PhysicalMassesDerived || incidence.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 117 refuses hidden physical input or changed residual nullity")
	}

	auts := computeAutomorphisms(incidence.FanoLines)
	summary := summarizeAutomorphisms(auts, incidence.FanoLines)
	if summary.Order != 168 || !summary.PointTransitive || !summary.LineTransitive || len(summary.CommonFixedPoints) != 0 || len(summary.CommonFixedLines) != 0 {
		return Analysis{}, fmt.Errorf("unexpected Fano automorphism audit: order=%d pointTransitive=%t lineTransitive=%t fixedPoints=%v fixedLines=%v", summary.Order, summary.PointTransitive, summary.LineTransitive, summary.CommonFixedPoints, summary.CommonFixedLines)
	}

	rows := buildRows(incidence.Rows)
	attempts := buildAttempts(summary)
	criteria := buildCriteria(summary)
	counts := countRows(rows)

	truth := "Gate 117 computes the Fano automorphism group directly from the seven incidence lines and verifies the full 168-element symmetry acts transitively on points and lines. Therefore no Fano point, line, chart, or contact-to-Fano bijection is selected by automorphism invariance. Since the current finite contact carrier has no derived matching automorphism action, a naturality square cannot be formulated; under the only known trivial contact action, every bijection is moved by nontrivial Fano automorphisms. Thus all 7! contact-to-Fano assignments remain convention-dependent, no representation row is opened, and the contact beta firewall remains closed."

	return Analysis{
		ContactIncidence:    incidence,
		Rows:                rows,
		Attempts:            attempts,
		Criteria:            criteria,
		Automorphisms:       auts,
		AutomorphismSummary: summary,

		ContactRows:                counts.contact,
		PositiveFiniteContactRows:  counts.positive,
		SurvivingCohomologyRows:    counts.surviving,
		IncidenceOpenRowsInherited: incidence.RepresentationOpenRows,
		OpenContactRowsAfter:       counts.open,

		FanoIncidenceAvailable:             incidence.FanoIncidenceAvailable,
		FanoAutomorphismGroupDerived:       true,
		FanoAutomorphismGroupOrder:         summary.Order,
		FanoAutomorphismGroupNontrivial:    summary.NonIdentityCount > 0,
		FanoPointActionTransitive:          summary.PointTransitive,
		FanoLineActionTransitive:           summary.LineTransitive,
		GlobalFixedFanoPoints:              len(summary.CommonFixedPoints),
		GlobalFixedFanoLines:               len(summary.CommonFixedLines),
		AutomorphismInvariantPointSelector: false,
		AutomorphismInvariantLineSelector:  false,

		ContactAutomorphismActionDerived:   false,
		NaturalitySquareFormulable:         false,
		InvariantContactToFanoMapDerived:   false,
		EquivariantBijectionDerived:        false,
		CanonicalAssignmentCount:           0,
		CompatibleBijectionCount:           factorial(7),
		ConventionDependentBijections:      factorial(7),
		SpectralOrderingBreaksAutomorphism: true,

		RepresentationCompleteRows:          counts.complete,
		RepresentationOpenRows:              counts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		NaturalityObstructionDerived:         true,
		IncidenceFunctorObstructionInherited: incidence.IncidenceFunctorObstructionDerived,
		LocalBundleObstructionInherited:      incidence.LocalBundleObstructionInherited,
		CohomologyObstructionInherited:       incidence.CohomologyObstructionInherited,

		ResidualNullityBefore:  incidence.ResidualNullityAfter,
		ResidualNullityAfter:   incidence.ResidualNullityAfter,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"Fano/contact cardinality resonance selects a canonical contact-to-Fano bijection",
			"a Fano point or line is fixed by the finite symmetry and can anchor the contact modes",
			"spectral ordering breaks Fano automorphism symmetry without adding convention",
			"an equivariant contact-Fano map exists before deriving a contact-side automorphism action",
			"automorphism data permits contact modes to enter threshold beta matching",
			"Gate 117 derives threshold-corrected beta coefficients, alpha, physical thetaW, M*, g_*, or masses",
		},
		RemainingUnknowns: []string{
			"U-37A-CONTACT-ACTION: derive an action of the finite Fano/G2 automorphism data on contact-overlap rows",
			"U-37B-NATURALITY: derive a naturality square linking contact modes to Fano points/lines",
			"U-37C-SYMMETRY-BREAKER: derive a finite symmetry-breaking selector rather than imposing spectral labels",
			"U-37D-FIBER: derive fiber type and local sections after a canonical assignment exists",
			"U-37E-REP: derive SU(3)c×SU(2)L×U(1)Y representation rows before beta matching",
			"U-37F-THRESHOLD: derive Lorentz kinetic, mass activation, and decoupling before Δb_i corrections",
		},
		RecommendedNextGate: "Gate 118 — contact symmetry-breaking selector / stabilizer reduction search",
	}, nil
}

type rowCounts struct{ contact, positive, surviving, open, complete int }

func buildRows(in []contactincidence.ContactIncidenceRow) []NaturalityRow {
	rows := make([]NaturalityRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, NaturalityRow{
			Name:                             r.Name,
			ModeKind:                         r.ModeKind,
			Value:                            r.Value,
			FiniteOverlapPositive:            r.FiniteOverlapPositive,
			SurvivesCohomology:               r.SurvivesCohomology,
			IncidenceOpen:                    r.IncidenceStatus == contactincidence.IncidenceOpen,
			FanoAssignmentInvariant:          false,
			EquivariantBijectionDerived:      false,
			ContactAutomorphismActionDerived: false,
			FanoAutomorphismInvariant:        false,
			NaturalitySquareDerived:          false,
			RepresentationRowDerived:         false,
			CanEnterBetaTensor:               false,
			NaturalityStatus:                 NaturalityOpen,
			Reason:                           "Fano automorphism group is transitive and no contact-side action/naturality square selects this row",
		})
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Value == rows[j].Value {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Value < rows[j].Value
	})
	return rows
}

func buildAttempts(summary AutomorphismSummary) []NaturalityAttempt {
	return []NaturalityAttempt{
		{
			Name:                            "derive Fano automorphism group from incidence lines",
			Kind:                            FullFanoAutomorphismAudit,
			Constructed:                     true,
			UsesExtraConvention:             false,
			InvariantUnderFanoAutomorphisms: true,
			CanonicalSelectorDerived:        false,
			RejectedAsPremature:             false,
			MissingTerms:                    []string{"contact-side automorphism action", "symmetry-breaking selector"},
			Detail:                          fmt.Sprintf("automorphism group order=%d; point orbit sizes=%v; line orbit sizes=%v", summary.Order, summary.PointOrbitSizes, summary.LineOrbitSizes),
		},
		{
			Name:                            "select an invariant Fano point or line",
			Kind:                            InvariantPointAssignment,
			Constructed:                     true,
			UsesExtraConvention:             false,
			InvariantUnderFanoAutomorphisms: false,
			CanonicalSelectorDerived:        false,
			RejectedAsPremature:             true,
			MissingTerms:                    []string{"global fixed Fano point", "global fixed Fano line"},
			Detail:                          "the full automorphism group acts transitively; no point or line is globally fixed",
		},
		{
			Name:                     "equivariant contact-Fano bijection",
			Kind:                     EquivariantBijectionAttempt,
			Constructed:              true,
			UsesExtraConvention:      true,
			ContactActionDerived:     false,
			EquivariantMapDerived:    false,
			CanonicalSelectorDerived: false,
			RejectedAsPremature:      true,
			MissingTerms:             []string{"nontrivial contact automorphism action", "naturality square", "equivariance proof"},
			Detail:                   "with only trivial contact action, nontrivial Fano automorphisms move every bijection; with no contact action, equivariance is not yet typed",
		},
		{
			Name:                            "use spectral labels to break automorphism symmetry",
			Kind:                            SpectralLabelBreakingAttempt,
			Constructed:                     true,
			UsesExtraConvention:             true,
			InvariantUnderFanoAutomorphisms: false,
			CanonicalSelectorDerived:        false,
			RejectedAsPremature:             true,
			MissingTerms:                    []string{"automorphism-invariant spectral functional", "orientation-free tie breaker", "field-map law"},
			Detail:                          "spectral ordering labels rows, but labels are not a Fano-natural selector and inherit Gate 107 orientation ambiguity",
		},
		{
			Name:                            "choose a line orbit as chart/fiber seed",
			Kind:                            LineOrbitChartAttempt,
			Constructed:                     true,
			UsesExtraConvention:             false,
			InvariantUnderFanoAutomorphisms: false,
			CanonicalSelectorDerived:        false,
			RepresentationRowDerived:        false,
			BetaRowPermitted:                false,
			ZeroRowProved:                   false,
			RejectedAsPremature:             true,
			MissingTerms:                    []string{"distinguished line orbit", "chart domain", "transition/cocycle law", "representation row"},
			Detail:                          "there is one line orbit of size seven, so orbit data selects the whole set but no individual chart or beta row",
		},
	}
}

func buildCriteria(summary AutomorphismSummary) []NaturalCriterion {
	return []NaturalCriterion{
		{Name: "Fano automorphism group from incidence", Required: true, Derived: summary.Order == 168, Detail: fmt.Sprintf("|Aut(Fano)|=%d", summary.Order)},
		{Name: "point transitivity", Required: true, Derived: summary.PointTransitive, Detail: fmt.Sprintf("point orbit sizes=%v", summary.PointOrbitSizes)},
		{Name: "line transitivity", Required: true, Derived: summary.LineTransitive, Detail: fmt.Sprintf("line orbit sizes=%v", summary.LineOrbitSizes)},
		{Name: "global fixed point/line selector", Required: true, Derived: false, Detail: "no point or line is fixed by all automorphisms"},
		{Name: "contact-side automorphism action", Required: true, Derived: false, Detail: "no action on contact partial-overlap rows is derived"},
		{Name: "naturality/equivariance square", Required: true, Derived: false, Detail: "cannot be typed without the contact-side action"},
		{Name: "representation and threshold row", Required: true, Derived: false, Detail: "no gauge row, Lorentz kinetic row, mass activation, or decoupling follows from automorphism data"},
	}
}

func countRows(rows []NaturalityRow) rowCounts {
	c := rowCounts{contact: len(rows)}
	for _, r := range rows {
		if r.FiniteOverlapPositive {
			c.positive++
		}
		if r.SurvivesCohomology {
			c.surviving++
		}
		if r.NaturalityStatus == NaturalityOpen {
			c.open++
		}
		if r.NaturalityStatus == NaturalityInvariant {
			c.complete++
		}
	}
	return c
}

func computeAutomorphisms(lines []contactincidence.FanoLine) []Permutation {
	lineSet := normalizedLineSet(lines)
	perms := make([]Permutation, 0, 168)
	var p Permutation
	used := [7]bool{}
	var rec func(int)
	rec = func(pos int) {
		if pos == 7 {
			if preservesLines(p, lineSet) {
				perms = append(perms, p)
			}
			return
		}
		for v := 0; v < 7; v++ {
			if used[v] {
				continue
			}
			used[v] = true
			p[pos] = v
			rec(pos + 1)
			used[v] = false
		}
	}
	rec(0)
	sort.Slice(perms, func(i, j int) bool { return permKey(perms[i]) < permKey(perms[j]) })
	return perms
}

func normalizedLineSet(lines []contactincidence.FanoLine) map[[3]int]bool {
	set := map[[3]int]bool{}
	for _, line := range lines {
		pts := []int{line.Points[0], line.Points[1], line.Points[2]}
		sort.Ints(pts)
		set[[3]int{pts[0], pts[1], pts[2]}] = true
	}
	return set
}

func preservesLines(p Permutation, lineSet map[[3]int]bool) bool {
	for line := range lineSet {
		pts := []int{p[line[0]], p[line[1]], p[line[2]]}
		sort.Ints(pts)
		if !lineSet[[3]int{pts[0], pts[1], pts[2]}] {
			return false
		}
	}
	return true
}

func summarizeAutomorphisms(perms []Permutation, lines []contactincidence.FanoLine) AutomorphismSummary {
	identity := 0
	rep := Permutation{}
	for _, p := range perms {
		if isIdentity(p) {
			identity++
		} else if rep == (Permutation{}) {
			rep = p
		}
	}
	pointOrbits := pointOrbits(perms)
	lineOrbits := lineOrbits(perms, lines)
	fixedPoints := commonFixedPoints(perms)
	fixedLines := commonFixedLines(perms, lines)
	return AutomorphismSummary{
		Order:                    len(perms),
		IdentityCount:            identity,
		NonIdentityCount:         len(perms) - identity,
		PointOrbitCount:          len(pointOrbits),
		LineOrbitCount:           len(lineOrbits),
		PointOrbitSizes:          orbitSizes(pointOrbits),
		LineOrbitSizes:           orbitSizes(lineOrbits),
		CommonFixedPoints:        fixedPoints,
		CommonFixedLines:         fixedLines,
		PointTransitive:          len(pointOrbits) == 1 && len(pointOrbits[0]) == 7,
		LineTransitive:           len(lineOrbits) == 1 && len(lineOrbits[0]) == 7,
		EveryLinePreserved:       len(perms) > 0,
		RepresentativeNontrivial: rep,
	}
}

func pointOrbits(perms []Permutation) [][]int {
	seen := [7]bool{}
	var orbits [][]int
	for i := 0; i < 7; i++ {
		if seen[i] {
			continue
		}
		orbitMap := map[int]bool{}
		for _, p := range perms {
			orbitMap[p[i]] = true
		}
		orbit := keys(orbitMap)
		for _, x := range orbit {
			seen[x] = true
		}
		orbits = append(orbits, orbit)
	}
	return orbits
}

func lineOrbits(perms []Permutation, lines []contactincidence.FanoLine) [][]int {
	lineKeys := make([][3]int, len(lines))
	index := map[[3]int]int{}
	for i, line := range lines {
		pts := []int{line.Points[0], line.Points[1], line.Points[2]}
		sort.Ints(pts)
		key := [3]int{pts[0], pts[1], pts[2]}
		lineKeys[i] = key
		index[key] = i
	}
	seen := make([]bool, len(lines))
	var orbits [][]int
	for i := range lines {
		if seen[i] {
			continue
		}
		orbitMap := map[int]bool{}
		for _, p := range perms {
			mapped := []int{p[lineKeys[i][0]], p[lineKeys[i][1]], p[lineKeys[i][2]]}
			sort.Ints(mapped)
			orbitMap[index[[3]int{mapped[0], mapped[1], mapped[2]}]] = true
		}
		orbit := keys(orbitMap)
		for _, x := range orbit {
			seen[x] = true
		}
		orbits = append(orbits, orbit)
	}
	return orbits
}

func commonFixedPoints(perms []Permutation) []int {
	fixed := []int{}
	for i := 0; i < 7; i++ {
		ok := true
		for _, p := range perms {
			if p[i] != i {
				ok = false
				break
			}
		}
		if ok {
			fixed = append(fixed, i)
		}
	}
	return fixed
}

func commonFixedLines(perms []Permutation, lines []contactincidence.FanoLine) []int {
	fixed := []int{}
	for i, line := range lines {
		keyPts := []int{line.Points[0], line.Points[1], line.Points[2]}
		sort.Ints(keyPts)
		key := [3]int{keyPts[0], keyPts[1], keyPts[2]}
		ok := true
		for _, p := range perms {
			mapped := []int{p[key[0]], p[key[1]], p[key[2]]}
			sort.Ints(mapped)
			if [3]int{mapped[0], mapped[1], mapped[2]} != key {
				ok = false
				break
			}
		}
		if ok {
			fixed = append(fixed, i)
		}
	}
	return fixed
}

func orbitSizes(orbits [][]int) []int {
	sizes := make([]int, len(orbits))
	for i, orbit := range orbits {
		sizes[i] = len(orbit)
	}
	sort.Ints(sizes)
	return sizes
}

func keys(m map[int]bool) []int {
	out := make([]int, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}

func isIdentity(p Permutation) bool {
	for i := 0; i < 7; i++ {
		if p[i] != i {
			return false
		}
	}
	return true
}

func permKey(p Permutation) string {
	var b strings.Builder
	for i := 0; i < 7; i++ {
		b.WriteByte(byte('0' + p[i]))
	}
	return b.String()
}

func factorial(n int) int {
	out := 1
	for i := 2; i <= n; i++ {
		out *= i
	}
	return out
}

func FormatRows(rows []NaturalityRow, limit int) string {
	parts := make([]string, 0, len(rows))
	for i, r := range rows {
		if limit > 0 && i >= limit {
			parts = append(parts, fmt.Sprintf("... +%d", len(rows)-i))
			break
		}
		parts = append(parts, fmt.Sprintf("%s(inv=%t,eq=%t,rep=%t,beta=%t)", r.Name, r.FanoAssignmentInvariant, r.EquivariantBijectionDerived, r.RepresentationRowDerived, r.CanEnterBetaTensor))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAutomorphismSummary(s AutomorphismSummary) string {
	return fmt.Sprintf("|Aut(Fano)|=%d; nonidentity=%d; pointOrbits=%v; lineOrbits=%v; fixedPoints=%v; fixedLines=%v; representative=%s", s.Order, s.NonIdentityCount, s.PointOrbitSizes, s.LineOrbitSizes, s.CommonFixedPoints, s.CommonFixedLines, FormatPermutation(s.RepresentativeNontrivial))
}

func FormatPermutation(p Permutation) string {
	parts := make([]string, 7)
	for i := 0; i < 7; i++ {
		parts[i] = fmt.Sprintf("%d→%d", i+1, p[i]+1)
	}
	return "{" + strings.Join(parts, ",") + "}"
}

func FormatAttempts(xs []NaturalityAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s,invariant=%t,equivariant=%t,beta=%t)", x.Name, x.Kind, x.InvariantUnderFanoAutomorphisms, x.EquivariantMapDerived, x.BetaRowPermitted))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []NaturalCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
