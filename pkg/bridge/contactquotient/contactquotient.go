// Package contactquotient implements Gate 120: contact spectral-invariant
// quotient / orbit-collapse theorem.
//
// Gate 119 proved that the contact-side automorphism group preserving the seven
// distinct finite overlap values is identity-only, while the Fano side still has
// a transitive 168-element automorphism group. Gate 120 asks whether quotienting
// the contact spectrum to invariant data can remove the obstruction.
//
// The result is a precise fork. The weighted/contact-spectrum quotient is
// canonical but has seven singleton orbits, so it gives no Fano-like orbit or
// representation row. The anonymous/full-symmetric quotient has one seven-mode
// orbit, but only by forgetting which distinct finite overlap value belongs to
// which mode; it destroys the row-level spectral information needed for
// threshold activation. Transported Fano quotients again require choosing one of
// 7! bijections. Therefore quotienting either changes nothing or collapses the
// information needed for physics. No contact representation row, beta tensor, or
// physical constant is derived.
package contactquotient

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactautaction"
)

type QuotientKind string

const (
	WeightedSpectrumQuotient  QuotientKind = "weighted-spectrum-quotient"
	ValuePartitionQuotient    QuotientKind = "value-partition-quotient"
	AnonymousSpectrumQuotient QuotientKind = "anonymous-spectrum-quotient"
	TransportedFanoQuotient   QuotientKind = "transported-fano-quotient-after-bijection"
	RepresentationRowQuotient QuotientKind = "representation-row-from-quotient"
)

type QuotientStatus string

const (
	QuotientOpen      QuotientStatus = "quotient-open"
	QuotientCanonical QuotientStatus = "quotient-canonical"
	QuotientCollapsed QuotientStatus = "quotient-collapsed"
	QuotientForbidden QuotientStatus = "quotient-forbidden"
)

type ContactQuotientRow struct {
	Name, ModeKind string
	Value          float64

	FiniteOverlapPositive bool
	SurvivesCohomology    bool
	ActionOpen            bool

	WeightedOrbitSize            int
	AnonymousOrbitSize           int
	ValuePartitionOrbitSize      int
	WeightedQuotientPreservesRow bool
	AnonymousQuotientLosesRow    bool
	FanoTransportNeedsChoice     bool
	RepresentationRowDerived     bool
	CanEnterBetaTensor           bool
	ZeroRowProved                bool

	Status QuotientStatus
	Reason string
}

type QuotientAttempt struct {
	Name string
	Kind QuotientKind

	Constructed               bool
	CanonicalUnderCurrentData bool
	UsesExtraConvention       bool
	PreservesContactSpectrum  bool
	PreservesRowIdentity      bool
	OrbitCount                int
	OrbitSizes                []int
	CollapsesDistinctValues   bool
	DestroysSpectralRows      bool
	FanoSymmetric             bool
	RepresentationRowDerived  bool
	BetaRowPermitted          bool
	ZeroRowProved             bool
	RejectedAsPremature       bool
	MissingTerms              []string
	Detail                    string
}

type QuotientCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type QuotientSummary struct {
	ContactRows               int
	DistinctSpectralValues    int
	WeightedAutomorphismOrder int
	FanoAutomorphismOrder     int

	WeightedOrbitCount        int
	WeightedOrbitSizes        []int
	WeightedQuotientCanonical bool
	WeightedQuotientFanoLike  bool

	ValuePartitionOrbitCount int
	ValuePartitionOrbitSizes []int

	AnonymousOrbitCount                   int
	AnonymousOrbitSizes                   []int
	AnonymousQuotientCanonical            bool
	AnonymousQuotientCollapsesRows        bool
	AnonymousQuotientDestroysSpectralRows bool

	CompatibleBijectionCount      int
	TransportedQuotientCount      int
	CanonicalTransportedQuotients int

	UsableRepresentationQuotients int
}

type Analysis struct {
	ContactAction contactautaction.Analysis

	Rows     []ContactQuotientRow
	Attempts []QuotientAttempt
	Criteria []QuotientCriterion
	Summary  QuotientSummary

	ContactRows               int
	PositiveFiniteContactRows int
	SurvivingCohomologyRows   int
	ActionOpenRowsInherited   int
	OpenContactRowsAfter      int

	ContactSpectrumQuotientSearchAttempted bool
	WeightedSpectrumQuotientDerived        bool
	WeightedQuotientCanonical              bool
	WeightedQuotientOrbitSizes             []int
	WeightedQuotientIsIdentity             bool
	WeightedQuotientPreservesAllRows       bool
	WeightedQuotientProducesFanoOrbit      bool

	AnonymousSpectrumQuotientDerived      bool
	AnonymousQuotientCanonical            bool
	AnonymousQuotientOrbitSizes           []int
	AnonymousQuotientCollapsesAllRows     bool
	AnonymousQuotientDestroysSpectralRows bool
	AnonymousQuotientRepresentationUsable bool

	TransportedFanoQuotientPossibleAfterChoice bool
	TransportedFanoQuotientCanonical           bool
	QuotientForkObstructionDerived             bool
	OrbitCollapseObstructionDerived            bool
	SpectralInformationLossDerived             bool
	RepresentationRowFromQuotientDerived       bool

	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	ActionObstructionInherited           bool
	SymmetrySelectorObstructionInherited bool
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
		action, err := contactautaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(action, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(action contactautaction.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !action.ContactSideActionSearchAttempted || !action.OrderMismatchObstructionDerived || !action.ContactWeightedAutomorphismGroupDerived || action.ContactWeightedAutomorphismGroupOrder != 1 || !action.ContactSpectralValuesAllDistinct {
		return Analysis{}, fmt.Errorf("Gate 120 requires Gate 119 identity-only weighted contact action and order-mismatch obstruction")
	}
	if action.FanoAutomorphismGroupOrder != 168 || !action.FanoPointActionTransitive || !action.FanoLineActionTransitive || action.AutFanoActionOnContactDerived || action.EquivariantAssignmentDerived || action.CanonicalContactFanoAssignmentDerived {
		return Analysis{}, fmt.Errorf("Gate 120 requires transitive Fano symmetry and no derived contact-Fano action/assignment")
	}
	if action.ContactRows != 7 || action.PositiveFiniteContactRows != 7 || action.SurvivingCohomologyRows != 7 || action.RepresentationOpenRows != 7 || action.ContactBetaRowsAllowed != 0 || action.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 120 requires seven action-open contact rows and closed beta permission")
	}
	if action.ResidualNullityAfter != 3 || action.HiddenObservedInputUsed || action.PhysicalWeakAngleDerived || action.FineStructureDerived || action.PhysicalMassesDerived || action.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 120 refuses hidden physical input or changed residual nullity")
	}

	values := contactValues(action.Rows)
	distinct := distinctCount(values, eps)
	if distinct != 7 {
		return Analysis{}, fmt.Errorf("expected seven distinct contact values, got %d", distinct)
	}

	weightedOrbits := singletonOrbits(7)
	anonymousOrbits := [][]int{{0, 1, 2, 3, 4, 5, 6}}
	valueOrbits := valuePartitionOrbits(values, eps)

	summary := QuotientSummary{
		ContactRows:               7,
		DistinctSpectralValues:    distinct,
		WeightedAutomorphismOrder: action.ContactWeightedAutomorphismGroupOrder,
		FanoAutomorphismOrder:     action.FanoAutomorphismGroupOrder,

		WeightedOrbitCount:        len(weightedOrbits),
		WeightedOrbitSizes:        orbitSizes(weightedOrbits),
		WeightedQuotientCanonical: true,
		WeightedQuotientFanoLike:  false,

		ValuePartitionOrbitCount: len(valueOrbits),
		ValuePartitionOrbitSizes: orbitSizes(valueOrbits),

		AnonymousOrbitCount:                   len(anonymousOrbits),
		AnonymousOrbitSizes:                   orbitSizes(anonymousOrbits),
		AnonymousQuotientCanonical:            true,
		AnonymousQuotientCollapsesRows:        true,
		AnonymousQuotientDestroysSpectralRows: true,

		CompatibleBijectionCount:      action.Summary.CompatibleBijectionCount,
		TransportedQuotientCount:      action.Summary.TransportedActionCount,
		CanonicalTransportedQuotients: 0,
		UsableRepresentationQuotients: 0,
	}

	rows := buildRows(action.Rows)
	attempts := buildAttempts(summary)
	criteria := buildCriteria(summary)
	counts := countRows(rows)

	truth := "Gate 120 tests whether quotienting the seven contact partial-overlap modes can turn the contact/Fano obstruction into usable representation data. The canonical weighted spectral quotient preserves the known distinct contact overlap values, but then every orbit is a singleton and no Fano-like transitive orbit or representation row appears. The canonical anonymous quotient collapses all seven modes into one invariant orbit, but only by forgetting the distinct row-level spectral data needed for threshold activation and beta matching. Transported Fano quotients again require choosing one of the 7! contact-to-Fano bijections. Therefore quotienting either changes nothing or destroys the finite information required for physics; no contact representation row, zero row, threshold beta correction, or physical constant is derived."

	return Analysis{
		ContactAction: action,
		Rows:          rows,
		Attempts:      attempts,
		Criteria:      criteria,
		Summary:       summary,

		ContactRows:               counts.contact,
		PositiveFiniteContactRows: counts.positive,
		SurvivingCohomologyRows:   counts.surviving,
		ActionOpenRowsInherited:   action.RepresentationOpenRows,
		OpenContactRowsAfter:      counts.open,

		ContactSpectrumQuotientSearchAttempted: true,
		WeightedSpectrumQuotientDerived:        true,
		WeightedQuotientCanonical:              true,
		WeightedQuotientOrbitSizes:             summary.WeightedOrbitSizes,
		WeightedQuotientIsIdentity:             summary.WeightedOrbitCount == 7,
		WeightedQuotientPreservesAllRows:       true,
		WeightedQuotientProducesFanoOrbit:      false,

		AnonymousSpectrumQuotientDerived:      true,
		AnonymousQuotientCanonical:            true,
		AnonymousQuotientOrbitSizes:           summary.AnonymousOrbitSizes,
		AnonymousQuotientCollapsesAllRows:     true,
		AnonymousQuotientDestroysSpectralRows: true,
		AnonymousQuotientRepresentationUsable: false,

		TransportedFanoQuotientPossibleAfterChoice: true,
		TransportedFanoQuotientCanonical:           false,
		QuotientForkObstructionDerived:             true,
		OrbitCollapseObstructionDerived:            true,
		SpectralInformationLossDerived:             true,
		RepresentationRowFromQuotientDerived:       false,

		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              counts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		ActionObstructionInherited:           true,
		SymmetrySelectorObstructionInherited: action.SymmetrySelectorObstructionInherited,
		NaturalityObstructionInherited:       action.NaturalityObstructionInherited,
		IncidenceFunctorObstructionInherited: action.IncidenceFunctorObstructionInherited,
		LocalBundleObstructionInherited:      action.LocalBundleObstructionInherited,
		CohomologyObstructionInherited:       action.CohomologyObstructionInherited,

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
			"the identity weighted contact quotient produces a Fano-like orbit",
			"forgetting contact spectral labels yields usable representation rows",
			"a one-orbit anonymous contact quotient can be used for threshold beta matching",
			"transported Fano quotient data is canonical without choosing a bijection",
			"quotienting the contact spectrum derives threshold-corrected beta coefficients",
			"Gate 120 derives alpha, physical thetaW, threshold masses, M*, g_*, or W/Z/Higgs/fermion masses",
		},
		RemainingUnknowns: []string{
			"canonical contact-to-Fano assignment that preserves enough spectral data for rows",
			"local bundle or constraint complex compatible with quotient data",
			"representation row map for the seven contact modes",
			"whether contact modes are physical, constrained, regulator, or vacuum-frustration modes",
			"threshold activation and decoupling law",
			"threshold-corrected beta tensor",
		},
		RecommendedNextGate: "Gate 121 — contact spectral reconstruction / invariant-to-row lifting obstruction theorem",
	}, nil
}

type rowCounts struct{ contact, positive, surviving, open int }

func buildRows(in []contactautaction.ContactActionRow) []ContactQuotientRow {
	rows := make([]ContactQuotientRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, ContactQuotientRow{
			Name:                         r.Name,
			ModeKind:                     r.ModeKind,
			Value:                        r.Value,
			FiniteOverlapPositive:        r.FiniteOverlapPositive,
			SurvivesCohomology:           r.SurvivesCohomology,
			ActionOpen:                   r.Status == contactautaction.ActionOpen,
			WeightedOrbitSize:            1,
			AnonymousOrbitSize:           7,
			ValuePartitionOrbitSize:      1,
			WeightedQuotientPreservesRow: true,
			AnonymousQuotientLosesRow:    true,
			FanoTransportNeedsChoice:     true,
			RepresentationRowDerived:     false,
			CanEnterBetaTensor:           false,
			ZeroRowProved:                false,
			Status:                       QuotientOpen,
			Reason:                       "weighted quotient leaves this distinct spectral row isolated; anonymous quotient erases the row identity; no representation row is derived",
		})
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if math.Abs(rows[i].Value-rows[j].Value) < 1e-12 {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Value < rows[j].Value
	})
	return rows
}

func buildAttempts(s QuotientSummary) []QuotientAttempt {
	return []QuotientAttempt{
		{
			Name:                      "weighted contact-spectrum quotient",
			Kind:                      WeightedSpectrumQuotient,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			PreservesContactSpectrum:  true,
			PreservesRowIdentity:      true,
			OrbitCount:                s.WeightedOrbitCount,
			OrbitSizes:                s.WeightedOrbitSizes,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"nontrivial orbit", "Fano-like action", "representation row"},
			Detail:                    "the canonical quotient by the identity-only weighted contact action has seven singleton orbits and produces no Fano-like threshold multiplet",
		},
		{
			Name:                      "value-partition quotient",
			Kind:                      ValuePartitionQuotient,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			PreservesContactSpectrum:  true,
			PreservesRowIdentity:      true,
			OrbitCount:                s.ValuePartitionOrbitCount,
			OrbitSizes:                s.ValuePartitionOrbitSizes,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"degenerate spectral classes", "representation row", "activation rule"},
			Detail:                    "because all seven overlap values are distinct, quotienting by equal spectral value again gives seven singleton classes",
		},
		{
			Name:                      "anonymous full-symmetric spectral quotient",
			Kind:                      AnonymousSpectrumQuotient,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			PreservesContactSpectrum:  false,
			PreservesRowIdentity:      false,
			OrbitCount:                s.AnonymousOrbitCount,
			OrbitSizes:                s.AnonymousOrbitSizes,
			CollapsesDistinctValues:   true,
			DestroysSpectralRows:      true,
			FanoSymmetric:             true,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"row reconstruction", "local bundle map", "per-mode threshold masses"},
			Detail:                    "the one-orbit quotient is symmetric only after forgetting which distinct contact value belongs to which mode, so it cannot supply row-level beta matching data",
		},
		{
			Name:                     "transported Fano quotient through chosen bijection",
			Kind:                     TransportedFanoQuotient,
			Constructed:              true,
			UsesExtraConvention:      true,
			PreservesContactSpectrum: false,
			OrbitCount:               1,
			OrbitSizes:               []int{7},
			CollapsesDistinctValues:  true,
			FanoSymmetric:            true,
			RejectedAsPremature:      true,
			MissingTerms:             []string{"canonical bijection", "spectral preservation", "symmetry-breaking theorem"},
			Detail:                   fmt.Sprintf("%d transported quotient pictures exist after choosing a contact-to-Fano bijection; none is canonical", s.TransportedQuotientCount),
		},
		{
			Name:                     "representation row from quotient data",
			Kind:                     RepresentationRowQuotient,
			Constructed:              true,
			OrbitCount:               0,
			RepresentationRowDerived: false,
			BetaRowPermitted:         false,
			ZeroRowProved:            false,
			RejectedAsPremature:      true,
			MissingTerms:             []string{"SU(3)c×SU(2)L×U(1)Y row", "Lorentz kinetic map", "mass/decoupling rule"},
			Detail:                   "neither singleton spectral rows nor an anonymous collapsed orbit determine gauge representation or beta contribution",
		},
	}
}

func buildCriteria(s QuotientSummary) []QuotientCriterion {
	return []QuotientCriterion{
		{Name: "weighted quotient", Required: true, Derived: s.WeightedOrbitCount == 7 && equalInts(s.WeightedOrbitSizes, []int{1, 1, 1, 1, 1, 1, 1}), Detail: "preserves all distinct spectral rows but leaves only singleton orbits"},
		{Name: "anonymous quotient", Required: true, Derived: s.AnonymousOrbitCount == 1 && equalInts(s.AnonymousOrbitSizes, []int{7}), Detail: "restores one orbit only by forgetting row-level spectral identity"},
		{Name: "canonical transported quotient", Required: true, Derived: s.CanonicalTransportedQuotients > 0, Detail: "must be false until a contact-Fano bijection is selected without convention"},
		{Name: "representation quotient", Required: true, Derived: s.UsableRepresentationQuotients > 0, Detail: "must be false: no quotient supplies gauge representation and decoupling rows"},
	}
}

func countRows(rows []ContactQuotientRow) rowCounts {
	var c rowCounts
	for _, r := range rows {
		c.contact++
		if r.FiniteOverlapPositive {
			c.positive++
		}
		if r.SurvivesCohomology {
			c.surviving++
		}
		if r.Status == QuotientOpen {
			c.open++
		}
	}
	return c
}

func contactValues(rows []contactautaction.ContactActionRow) []float64 {
	values := make([]float64, 0, len(rows))
	for _, r := range rows {
		values = append(values, r.Value)
	}
	sort.Float64s(values)
	return values
}

func distinctCount(values []float64, eps float64) int {
	if len(values) == 0 {
		return 0
	}
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	count := 1
	last := sorted[0]
	for _, v := range sorted[1:] {
		if math.Abs(v-last) > eps {
			count++
			last = v
		}
	}
	return count
}

func singletonOrbits(n int) [][]int {
	out := make([][]int, n)
	for i := 0; i < n; i++ {
		out[i] = []int{i}
	}
	return out
}

func valuePartitionOrbits(values []float64, eps float64) [][]int {
	type pair struct {
		idx int
		v   float64
	}
	pairs := make([]pair, len(values))
	for i, v := range values {
		pairs[i] = pair{i, v}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].v < pairs[j].v })
	var orbits [][]int
	for _, p := range pairs {
		if len(orbits) == 0 || math.Abs(values[orbits[len(orbits)-1][0]]-p.v) > eps {
			orbits = append(orbits, []int{p.idx})
		} else {
			orbits[len(orbits)-1] = append(orbits[len(orbits)-1], p.idx)
		}
	}
	return orbits
}

func orbitSizes(orbits [][]int) []int {
	sizes := make([]int, len(orbits))
	for i, o := range orbits {
		sizes[i] = len(o)
	}
	sort.Ints(sizes)
	return sizes
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func FormatRows(rows []ContactQuotientRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[value=%.10f weightedOrbit=%d anonymousOrbit=%d status=%s]", r.Name, r.Value, r.WeightedOrbitSize, r.AnonymousOrbitSize, r.Status))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return strings.Join(parts, "; ")
}

func FormatAttempts(attempts []QuotientAttempt) string {
	parts := make([]string, 0, len(attempts))
	for _, a := range attempts {
		parts = append(parts, fmt.Sprintf("%s(kind=%s orbits=%v canonical=%t convention=%t rejected=%t)", a.Name, a.Kind, a.OrbitSizes, a.CanonicalUnderCurrentData, a.UsesExtraConvention, a.RejectedAsPremature))
	}
	return strings.Join(parts, "; ")
}

func FormatCriteria(criteria []QuotientCriterion) string {
	parts := make([]string, 0, len(criteria))
	for _, c := range criteria {
		parts = append(parts, fmt.Sprintf("%s=%t (%s)", c.Name, c.Derived, c.Detail))
	}
	return strings.Join(parts, "; ")
}

func FormatSummary(s QuotientSummary) string {
	return fmt.Sprintf("weightedOrbits=%v anonymousOrbits=%v valueOrbits=%v bijections=%d usableRepresentationQuotients=%d", s.WeightedOrbitSizes, s.AnonymousOrbitSizes, s.ValuePartitionOrbitSizes, s.CompatibleBijectionCount, s.UsableRepresentationQuotients)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
