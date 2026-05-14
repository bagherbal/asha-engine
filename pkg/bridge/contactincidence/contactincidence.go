// Package contactincidence implements Gate 116: contact incidence/fiber functor
// search from Fano/contact geometry.
//
// Gates 114 and 115 blocked the two easy exits for the seven contact
// partial-overlap modes: no canonical constraint/BRST differential cancels them,
// and no local continuum bundle/representation row has been derived.  Gate 116
// asks whether the already-present Fano/octonionic incidence geometry can do
// better: can it define a canonical fiber functor, chart atlas, or
// representation-row map for those seven modes?
//
// The answer is still no.  The Fano plane incidence is exact finite data: seven
// points, seven lines, each point on three lines and each line through three
// points.  It resonates perfectly with the seven contact rows.  But the finite
// project currently contains no canonical natural transformation from contact
// partial-overlap eigenvectors to Fano points/lines, no chart/cocycle law, no
// fiber type, no local section map, and no gauge-representation row.  Any
// bijection or orientation would be an extra convention.  Therefore the contact
// beta firewall remains closed.
package contactincidence

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactbundle"
	"github.com/bagherbal/asha-engine/pkg/octonion"
)

type IncidenceAttemptKind string

const (
	FanoIncidenceCarrierAttempt IncidenceAttemptKind = "fano-incidence-carrier"
	DirectBijectionAttempt      IncidenceAttemptKind = "direct-contact-fano-bijection"
	SpectralRankFunctorAttempt  IncidenceAttemptKind = "spectral-rank-fiber-functor"
	LineBundleAtlasAttempt      IncidenceAttemptKind = "fano-line-chart-atlas"
	OctonionFiberAttempt        IncidenceAttemptKind = "octonion-multiplication-fiber"
)

type ContactIncidenceStatus string

const (
	IncidenceOpen      ContactIncidenceStatus = "incidence-open"
	IncidenceComplete  ContactIncidenceStatus = "incidence-complete"
	IncidenceForbidden ContactIncidenceStatus = "incidence-forbidden"
)

type FanoLine struct {
	Name        string
	Points      [3]int
	Sign        int
	Oriented    bool
	Canonical   bool
	LineDegrees [3]int
}

type ContactIncidenceRow struct {
	Name, ModeKind string
	Value          float64

	FiniteOverlapPositive bool
	SurvivesCohomology    bool
	BundleRowOpen         bool

	FanoPointAssigned              bool
	FanoLineAssigned               bool
	CanonicalContactToFanoMap      bool
	FiberFunctorDerived            bool
	ChartAtlasDerived              bool
	TransitionCocycleDerived       bool
	SectionMapDerived              bool
	GaugeRepresentationRowDerived  bool
	LorentzKineticRowDerived       bool
	MassActivationDerived          bool
	DecouplingRuleDerived          bool
	RepresentationOrZeroRowDerived bool

	IncidenceStatus    ContactIncidenceStatus
	CanEnterBetaTensor bool
	Reason             string
}

type IncidenceAttempt struct {
	Name string
	Kind IncidenceAttemptKind

	Constructed               bool
	UsesExtraConvention       bool
	CanonicalUnderCurrentData bool

	FanoIncidenceDerived       bool
	ContactToFanoMapDerived    bool
	FiberFunctorDerived        bool
	ChartAtlasDerived          bool
	TransitionCocycleDerived   bool
	GaugeRepresentationDerived bool
	LorentzKineticDerived      bool
	MassActivationDerived      bool
	DecouplingRuleDerived      bool
	BetaRowPermitted           bool
	ZeroRowProved              bool
	RejectedAsPremature        bool
	MissingTerms               []string
	Detail                     string
}

type FunctorCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Analysis struct {
	ContactBundle contactbundle.Analysis

	Rows      []ContactIncidenceRow
	FanoLines []FanoLine
	Attempts  []IncidenceAttempt
	Criteria  []FunctorCriterion

	ContactRows               int
	PositiveFiniteContactRows int
	SurvivingCohomologyRows   int
	OpenBundleRowsInherited   int
	OpenContactRowsAfter      int

	FanoIncidenceAvailable      bool
	FanoPointCount              int
	FanoLineCount               int
	EveryFanoPointDegreeThree   bool
	EveryFanoLineSizeThree      bool
	FanoContactCardinalityMatch bool
	FanoOrientationAvailable    bool
	FanoIncidenceResonance      bool

	ContactToFanoMapAttempted     bool
	CanonicalContactToFanoMap     bool
	FiberFunctorConstructionTried bool
	FiberFunctorDerived           bool
	ChartAtlasDerived             bool
	TransitionCocycleDerived      bool
	SectionMapDerived             bool
	GaugeRepresentationDerived    bool
	LorentzKineticDerived         bool
	MassActivationDerived         bool
	DecouplingRuleDerived         bool

	AnyFunctorAttemptConstructed        bool
	AnyFunctorAttemptCanonical          bool
	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	IncidenceFunctorObstructionDerived bool
	LocalBundleObstructionInherited    bool
	CohomologyObstructionInherited     bool
	BranchSelectorDerived              bool

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
		bundle, err := contactbundle.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(bundle)
	})
	return defaultValue, defaultErr
}

func Build(bundle contactbundle.Analysis) (Analysis, error) {
	if !bundle.LocalBundleObstructionDerived || !bundle.RepresentationRowConstructionObstructed {
		return Analysis{}, fmt.Errorf("Gate 116 requires Gate 115 local-bundle obstruction")
	}
	if bundle.ContactRows != 7 || bundle.PositiveFiniteContactRows != 7 || bundle.SurvivingCohomologyRows != 7 || bundle.OpenContactRowsAfter != 7 {
		return Analysis{}, fmt.Errorf("Gate 116 requires seven positive surviving contact bundle-open rows")
	}
	if bundle.RepresentationCompleteRows != 0 || bundle.ContactBetaRowsAllowed != 0 || bundle.ContactZeroRowsProved != 0 || bundle.ThresholdCorrectedBetaDerived || bundle.FullFiniteBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 116 requires sealed contact representation/beta rows")
	}
	if bundle.ResidualNullityAfter != 3 || bundle.HiddenObservedInputUsed || bundle.PhysicalWeakAngleDerived || bundle.FineStructureDerived || bundle.PhysicalMassesDerived || bundle.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 116 refuses hidden physical input or changed residual nullity")
	}

	fano := buildFanoLines()
	rows := buildRows(bundle.Rows)
	attempts := buildAttempts()
	criteria := buildCriteria()
	rowCounts := countRows(rows)
	pointDegrees := fanoPointDegrees(fano)

	anyConstructed := false
	anyCanonical := false
	for _, a := range attempts {
		if a.Constructed {
			anyConstructed = true
		}
		if a.CanonicalUnderCurrentData && a.FiberFunctorDerived {
			anyCanonical = true
		}
	}

	truth := "Gate 116 tests whether the exact Fano/octonionic incidence already present in the finite geometry can lift the seven contact partial-overlap modes into a canonical fiber functor, chart atlas, or representation-row map. The Fano plane data is exact and cardinality-matched to the seven contact rows, but no canonical contact-to-Fano natural transformation, chart/cocycle law, local section map, gauge representation, Lorentz kinetic row, mass activation, or decoupling rule is derived. Thus Fano incidence is a real structural resonance, not yet a threshold-field functor; all seven contact modes remain incidence-open and the beta firewall stays closed."

	return Analysis{
		ContactBundle: bundle,
		Rows:          rows,
		FanoLines:     fano,
		Attempts:      attempts,
		Criteria:      criteria,

		ContactRows:               rowCounts.contact,
		PositiveFiniteContactRows: rowCounts.positiveContact,
		SurvivingCohomologyRows:   rowCounts.survivesCohomology,
		OpenBundleRowsInherited:   bundle.OpenContactRowsAfter,
		OpenContactRowsAfter:      rowCounts.open,

		FanoIncidenceAvailable:      len(fano) == 7,
		FanoPointCount:              len(pointDegrees),
		FanoLineCount:               len(fano),
		EveryFanoPointDegreeThree:   everyDegree(pointDegrees, 3),
		EveryFanoLineSizeThree:      everyLineSize(fano, 3),
		FanoContactCardinalityMatch: len(rows) == 7 && len(pointDegrees) == 7 && len(fano) == 7,
		FanoOrientationAvailable:    true,
		FanoIncidenceResonance:      len(rows) == 7 && len(pointDegrees) == 7 && everyDegree(pointDegrees, 3),

		ContactToFanoMapAttempted:     true,
		CanonicalContactToFanoMap:     false,
		FiberFunctorConstructionTried: true,
		FiberFunctorDerived:           false,
		ChartAtlasDerived:             false,
		TransitionCocycleDerived:      false,
		SectionMapDerived:             false,
		GaugeRepresentationDerived:    false,
		LorentzKineticDerived:         false,
		MassActivationDerived:         false,
		DecouplingRuleDerived:         false,

		AnyFunctorAttemptConstructed:        anyConstructed,
		AnyFunctorAttemptCanonical:          anyCanonical,
		RepresentationCompleteRows:          rowCounts.complete,
		RepresentationOpenRows:              rowCounts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		IncidenceFunctorObstructionDerived: true,
		LocalBundleObstructionInherited:    bundle.LocalBundleObstructionDerived,
		CohomologyObstructionInherited:     bundle.ConstraintShortcutBlocked,
		BranchSelectorDerived:              false,

		ResidualNullityBefore:  bundle.ResidualNullityAfter,
		ResidualNullityAfter:   bundle.ResidualNullityAfter,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"seven contact modes plus seven Fano points automatically define a canonical bijection",
			"Fano incidence alone supplies a local continuum fiber functor",
			"octonion multiplication signs are threshold gauge-representation rows",
			"Fano lines are spacetime charts without a base-space/support map",
			"contact modes may enter beta matching by incidence resonance alone",
			"Gate 116 derives threshold-corrected beta coefficients, alpha, physical thetaW, M*, g_*, or masses",
		},
		RemainingUnknowns: []string{
			"U-36A-NATURALITY: derive a canonical natural transformation from contact-overlap rows to Fano points/lines",
			"U-36B-FIBER: derive fiber type from incidence rather than assigning singlet/doublet rows",
			"U-36C-ATLAS: derive chart overlaps and transition/cocycle law",
			"U-36D-SECTIONS: derive local section variables from finite contact incidence",
			"U-36E-REP: derive SU(3)c×SU(2)L×U(1)Y representation rows for contact modes",
			"U-36F-THRESHOLD: derive Lorentz kinetic, mass activation, and decoupling before Δb_i corrections",
		},
		RecommendedNextGate: "Gate 117 — contact-Fano naturality obstruction / automorphism-invariance theorem",
	}, nil
}

type rowCounts struct{ contact, positiveContact, survivesCohomology, open, complete int }

func buildRows(in []contactbundle.BundleRow) []ContactIncidenceRow {
	rows := make([]ContactIncidenceRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, ContactIncidenceRow{
			Name:                           r.Name,
			ModeKind:                       r.ModeKind,
			Value:                          r.Value,
			FiniteOverlapPositive:          r.FiniteOverlapPositive,
			SurvivesCohomology:             r.SurvivesCohomology,
			BundleRowOpen:                  r.RepresentationStatus == contactbundle.RepresentationOpen,
			FanoPointAssigned:              false,
			FanoLineAssigned:               false,
			CanonicalContactToFanoMap:      false,
			FiberFunctorDerived:            false,
			ChartAtlasDerived:              false,
			TransitionCocycleDerived:       false,
			SectionMapDerived:              false,
			GaugeRepresentationRowDerived:  false,
			LorentzKineticRowDerived:       false,
			MassActivationDerived:          false,
			DecouplingRuleDerived:          false,
			RepresentationOrZeroRowDerived: false,
			IncidenceStatus:                IncidenceOpen,
			CanEnterBetaTensor:             false,
			Reason:                         "Fano incidence exists, but no canonical contact-to-Fano map, local fiber functor, representation row, or threshold rule is derived",
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

func buildFanoLines() []FanoLine {
	terms := octonion.StandardFanoTerms()
	lines := make([]FanoLine, 0, len(terms))
	degrees := make(map[int]int, 7)
	for _, t := range terms {
		degrees[t.I]++
		degrees[t.J]++
		degrees[t.K]++
	}
	for idx, t := range terms {
		pts := [3]int{t.I, t.J, t.K}
		lines = append(lines, FanoLine{
			Name:        fmt.Sprintf("L%d=(%d%d%d)", idx+1, t.I+1, t.J+1, t.K+1),
			Points:      pts,
			Sign:        t.Sign,
			Oriented:    true,
			Canonical:   true,
			LineDegrees: [3]int{degrees[t.I], degrees[t.J], degrees[t.K]},
		})
	}
	return lines
}

func buildAttempts() []IncidenceAttempt {
	return []IncidenceAttempt{
		{
			Name:                      "canonical Fano incidence carrier",
			Kind:                      FanoIncidenceCarrierAttempt,
			Constructed:               true,
			UsesExtraConvention:       false,
			CanonicalUnderCurrentData: true,
			FanoIncidenceDerived:      true,
			RejectedAsPremature:       false,
			MissingTerms:              []string{"contact-to-Fano natural transformation", "fiber type", "local sections", "representation row"},
			Detail:                    "the Fano plane is exact finite data, but incidence alone is not yet a contact-mode fiber functor",
		},
		{
			Name:                      "choose a direct contact-row to Fano-point bijection",
			Kind:                      DirectBijectionAttempt,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			ContactToFanoMapDerived:   false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"canonical bijection", "automorphism breaking rule", "naturality proof"},
			Detail:                    "7! bijections are compatible; the current finite data selects none",
		},
		{
			Name:                      "spectral-rank to Fano-point fiber functor",
			Kind:                      SpectralRankFunctorAttempt,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			FiberFunctorDerived:       false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"orientation", "tie-breaking/naturality", "fiber law", "gauge representation"},
			Detail:                    "spectral sorting provides an order, but Gate 107 already exposed order-orientation ambiguity",
		},
		{
			Name:                      "Fano-line chart atlas",
			Kind:                      LineBundleAtlasAttempt,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			ChartAtlasDerived:         false,
			TransitionCocycleDerived:  false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"base-space support", "chart domains", "transition functions", "cocycle law", "sections"},
			Detail:                    "Fano lines are incidence triples, not spacetime chart overlaps without a base map",
		},
		{
			Name:                       "octonion multiplication fiber assignment",
			Kind:                       OctonionFiberAttempt,
			Constructed:                true,
			UsesExtraConvention:        true,
			CanonicalUnderCurrentData:  false,
			GaugeRepresentationDerived: false,
			LorentzKineticDerived:      false,
			MassActivationDerived:      false,
			DecouplingRuleDerived:      false,
			RejectedAsPremature:        true,
			MissingTerms:               []string{"field variables", "Lorentz kinetic operator", "SU(3)c×SU(2)L×U(1)Y row", "mass/decoupling rule"},
			Detail:                     "octonion signs do not by themselves define propagating threshold fields",
		},
	}
}

func buildCriteria() []FunctorCriterion {
	return []FunctorCriterion{
		{Name: "Fano incidence carrier", Required: true, Derived: true, Detail: "seven oriented octonionic/Fano lines are available"},
		{Name: "contact/Fano cardinality match", Required: true, Derived: true, Detail: "seven contact rows and seven Fano points/lines resonate"},
		{Name: "canonical contact-to-Fano natural transformation", Required: true, Derived: false, Detail: "no automorphism-invariant bijection is selected"},
		{Name: "fiber functor", Required: true, Derived: false, Detail: "no functor from incidence data to local field fibers is derived"},
		{Name: "chart atlas and transition/cocycle law", Required: true, Derived: false, Detail: "Fano lines are not yet local spacetime chart overlaps"},
		{Name: "local sections", Required: true, Derived: false, Detail: "no section variables are constructed"},
		{Name: "gauge representation and hypercharge row", Required: true, Derived: false, Detail: "no SU(3)c×SU(2)L×U(1)Y row is derived"},
		{Name: "Lorentz kinetic, mass activation, decoupling", Required: true, Derived: false, Detail: "no threshold-field permission data exists"},
	}
}

func countRows(rows []ContactIncidenceRow) rowCounts {
	c := rowCounts{contact: len(rows)}
	for _, r := range rows {
		if r.FiniteOverlapPositive {
			c.positiveContact++
		}
		if r.SurvivesCohomology {
			c.survivesCohomology++
		}
		if r.IncidenceStatus == IncidenceOpen {
			c.open++
		}
		if r.IncidenceStatus == IncidenceComplete {
			c.complete++
		}
	}
	return c
}

func fanoPointDegrees(lines []FanoLine) map[int]int {
	degrees := make(map[int]int, 7)
	for i := 0; i < 7; i++ {
		degrees[i] = 0
	}
	for _, line := range lines {
		for _, p := range line.Points {
			degrees[p]++
		}
	}
	return degrees
}

func everyDegree(degrees map[int]int, want int) bool {
	if len(degrees) != 7 {
		return false
	}
	for i := 0; i < 7; i++ {
		if degrees[i] != want {
			return false
		}
	}
	return true
}

func everyLineSize(lines []FanoLine, want int) bool {
	for _, line := range lines {
		count := 0
		seen := map[int]bool{}
		for _, p := range line.Points {
			seen[p] = true
		}
		count = len(seen)
		if count != want {
			return false
		}
	}
	return len(lines) > 0
}

func FormatRows(rows []ContactIncidenceRow, limit int) string {
	parts := make([]string, 0, len(rows))
	for i, r := range rows {
		if limit > 0 && i >= limit {
			parts = append(parts, fmt.Sprintf("... +%d", len(rows)-i))
			break
		}
		parts = append(parts, fmt.Sprintf("%s(fano=%t,functor=%t,rep=%t,beta=%t)", r.Name, r.FanoPointAssigned || r.FanoLineAssigned, r.FiberFunctorDerived, r.GaugeRepresentationRowDerived, r.CanEnterBetaTensor))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatFano(lines []FanoLine) string {
	parts := make([]string, 0, len(lines))
	for _, line := range lines {
		parts = append(parts, fmt.Sprintf("%s(sign=%+d,deg=%v)", line.Name, line.Sign, line.LineDegrees))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(xs []IncidenceAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s,canonical=%t,functor=%t,beta=%t)", x.Name, x.Kind, x.CanonicalUnderCurrentData, x.FiberFunctorDerived, x.BetaRowPermitted))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []FunctorCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
