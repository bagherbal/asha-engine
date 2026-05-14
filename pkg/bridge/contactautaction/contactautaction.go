// Package contactautaction implements Gate 119: contact-side automorphism action
// construction / equivariant assignment search.
//
// Gates 116--118 established an exact Fano incidence carrier with a 168-element
// automorphism group, but also showed that no canonical Fano point/line/flag is
// selected. Gate 119 asks the next structural question: can the seven contact
// partial-overlap modes themselves carry a derived action of Aut(Fano), so that
// a contact-to-Fano assignment becomes equivariant rather than conventional?
//
// The result is an obstruction. If the finite overlap eigenvalues are preserved
// as part of the contact-side structure, the contact automorphism group is only
// the identity because the seven positive overlap values are all distinct. A
// faithful 168-element Aut(Fano) action can be transported to the contact labels
// only after choosing one of the 7! contact-to-Fano bijections, but that is the
// convention Gate 117/118 refused. Therefore no canonical contact-side action,
// equivariant assignment, local representation row, or threshold beta permission
// is derived.
package contactautaction

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactnaturality"
	"github.com/bagherbal/asha-engine/pkg/bridge/contactsymmetry"
)

type ActionAttemptKind string

const (
	TrivialContactActionAttempt        ActionAttemptKind = "trivial-contact-action"
	SpectralPreservingActionAttempt    ActionAttemptKind = "spectral-preserving-contact-action"
	TransportedFanoActionAttempt       ActionAttemptKind = "transported-fano-action-after-bijection"
	DirectEquivariantAssignmentSearch  ActionAttemptKind = "direct-equivariant-assignment-search"
	RepresentationRowFromActionAttempt ActionAttemptKind = "representation-row-from-contact-action"
)

type ActionStatus string

const (
	ActionOpen      ActionStatus = "action-open"
	ActionComplete  ActionStatus = "action-complete"
	ActionForbidden ActionStatus = "action-forbidden"
)

type ContactActionRow struct {
	Name, ModeKind string
	Value          float64

	FiniteOverlapPositive bool
	SurvivesCohomology    bool
	SymmetryOpen          bool

	SpectralValuePreserved    bool
	ContactActionDerived      bool
	AutFanoActionOnContact    bool
	EquivariantFanoAssignment bool
	CanonicalContactFanoMap   bool
	RepresentationRowDerived  bool
	CanEnterBetaTensor        bool
	ZeroRowProved             bool

	Status ActionStatus
	Reason string
}

type ActionAttempt struct {
	Name string
	Kind ActionAttemptKind

	Constructed                 bool
	CanonicalUnderCurrentData   bool
	UsesExtraConvention         bool
	PreservesContactSpectrum    bool
	ContactGroupOrder           int
	FanoGroupOrder              int
	DefinesAutFanoAction        bool
	EquivariantAssignmentExists bool
	CanonicalAssignmentDerived  bool
	RepresentationRowDerived    bool
	BetaRowPermitted            bool
	ZeroRowProved               bool
	RejectedAsPremature         bool
	MissingTerms                []string
	Detail                      string
}

type ActionCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type ContactActionSummary struct {
	ContactWeightedAutomorphismOrder int
	ContactWeightedIdentityOnly      bool
	SpectralValuesDistinct           bool
	ContactWeightOrbitSizes          []int
	FanoAutomorphismOrder            int
	FanoPointOrbitSizes              []int
	OrderMismatch                    bool
	CompatibleBijectionCount         int
	TransportedActionCount           int
	CanonicalTransportedActions      int
}

type Analysis struct {
	ContactSymmetry contactsymmetry.Analysis

	Rows     []ContactActionRow
	Attempts []ActionAttempt
	Criteria []ActionCriterion
	Summary  ContactActionSummary

	ContactRows               int
	PositiveFiniteContactRows int
	SurvivingCohomologyRows   int
	SymmetryOpenRowsInherited int
	OpenContactRowsAfter      int

	FanoAutomorphismGroupDerived bool
	FanoAutomorphismGroupOrder   int
	FanoPointActionTransitive    bool
	FanoLineActionTransitive     bool

	ContactSideActionSearchAttempted          bool
	ContactWeightedAutomorphismGroupDerived   bool
	ContactWeightedAutomorphismGroupOrder     int
	ContactWeightedActionIdentityOnly         bool
	ContactSpectralValuesAllDistinct          bool
	ContactActionOrbitSizes                   []int
	AutFanoActionOnContactDerived             bool
	AutFanoActionPreservingContactData        bool
	OrderMismatchObstructionDerived           bool
	TrivialContactActionRejected              bool
	TransportedFanoActionsPossibleAfterChoice bool
	TransportedFanoActionCanonical            bool
	EquivariantAssignmentDerived              bool
	CanonicalContactFanoAssignmentDerived     bool
	NaturalitySquareFormulable                bool

	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

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
		sym, err := contactsymmetry.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sym, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(sym contactsymmetry.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !sym.StabilizerArithmeticDerived || !sym.NaturalityObstructionInherited || sym.FanoAutomorphismGroupOrder != 168 {
		return Analysis{}, fmt.Errorf("Gate 119 requires Gate 118 stabilizer/naturality obstruction with |Aut(Fano)|=168")
	}
	if sym.CanonicalSymmetryBreakingObjectDerived || sym.CanonicalContactFanoAssignmentDerived || sym.ContactAutomorphismActionDerived || sym.NaturalitySquareFormulable {
		return Analysis{}, fmt.Errorf("Gate 119 requires no prior canonical selector, contact action, or naturality square")
	}
	if sym.ContactRows != 7 || sym.PositiveFiniteContactRows != 7 || sym.SurvivingCohomologyRows != 7 || sym.RepresentationOpenRows != 7 || sym.ContactBetaRowsAllowed != 0 || sym.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 119 requires seven symmetry-open contact rows and closed beta permission")
	}
	if sym.ResidualNullityAfter != 3 || sym.HiddenObservedInputUsed || sym.PhysicalWeakAngleDerived || sym.FineStructureDerived || sym.PhysicalMassesDerived || sym.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 119 refuses hidden physical input or changed residual nullity")
	}

	values := contactValues(sym.Rows)
	weighted := spectralPreservingPermutations(values, eps)
	orbitSizes := orbitSizesFromPermutations(weighted)
	distinct := valuesDistinct(values, eps)
	compatible := factorial(7)

	summary := ContactActionSummary{
		ContactWeightedAutomorphismOrder: len(weighted),
		ContactWeightedIdentityOnly:      len(weighted) == 1 && isIdentity(weighted[0]),
		SpectralValuesDistinct:           distinct,
		ContactWeightOrbitSizes:          orbitSizes,
		FanoAutomorphismOrder:            sym.FanoAutomorphismGroupOrder,
		FanoPointOrbitSizes:              []int{7},
		OrderMismatch:                    len(weighted) != sym.FanoAutomorphismGroupOrder,
		CompatibleBijectionCount:         compatible,
		TransportedActionCount:           compatible,
		CanonicalTransportedActions:      0,
	}
	if summary.ContactWeightedAutomorphismOrder != 1 || !summary.ContactWeightedIdentityOnly || !summary.SpectralValuesDistinct {
		return Analysis{}, fmt.Errorf("expected distinct contact spectrum with identity-only weighted automorphism group, got order=%d distinct=%t", summary.ContactWeightedAutomorphismOrder, summary.SpectralValuesDistinct)
	}

	rows := buildRows(sym.Rows)
	attempts := buildAttempts(summary)
	criteria := buildCriteria(summary)
	counts := countRows(rows)

	truth := "Gate 119 constructs the contact-side automorphism-action search. The seven contact partial-overlap eigenvalues are all distinct, so the automorphism group preserving the known contact finite-overlap data is the identity. This canonical contact action is too small to support the 168-element transitive Aut(Fano) action. A faithful Aut(Fano) action can be transported to contact labels only after choosing one of the 7! contact-to-Fano bijections, which is exactly the noncanonical choice blocked by Gates 117 and 118. Therefore no canonical contact-side Aut(Fano) action, equivariant assignment, representation row, or threshold beta correction is derived."

	return Analysis{
		ContactSymmetry: sym,
		Rows:            rows,
		Attempts:        attempts,
		Criteria:        criteria,
		Summary:         summary,

		ContactRows:               counts.contact,
		PositiveFiniteContactRows: counts.positive,
		SurvivingCohomologyRows:   counts.surviving,
		SymmetryOpenRowsInherited: sym.RepresentationOpenRows,
		OpenContactRowsAfter:      counts.open,

		FanoAutomorphismGroupDerived: true,
		FanoAutomorphismGroupOrder:   sym.FanoAutomorphismGroupOrder,
		FanoPointActionTransitive:    sym.FanoPointActionTransitive,
		FanoLineActionTransitive:     sym.FanoLineActionTransitive,

		ContactSideActionSearchAttempted:          true,
		ContactWeightedAutomorphismGroupDerived:   true,
		ContactWeightedAutomorphismGroupOrder:     summary.ContactWeightedAutomorphismOrder,
		ContactWeightedActionIdentityOnly:         summary.ContactWeightedIdentityOnly,
		ContactSpectralValuesAllDistinct:          summary.SpectralValuesDistinct,
		ContactActionOrbitSizes:                   summary.ContactWeightOrbitSizes,
		AutFanoActionOnContactDerived:             false,
		AutFanoActionPreservingContactData:        false,
		OrderMismatchObstructionDerived:           summary.OrderMismatch,
		TrivialContactActionRejected:              true,
		TransportedFanoActionsPossibleAfterChoice: true,
		TransportedFanoActionCanonical:            false,
		EquivariantAssignmentDerived:              false,
		CanonicalContactFanoAssignmentDerived:     false,
		NaturalitySquareFormulable:                false,

		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              counts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		SymmetrySelectorObstructionInherited: true,
		NaturalityObstructionInherited:       sym.NaturalityObstructionInherited,
		IncidenceFunctorObstructionInherited: sym.IncidenceFunctorObstructionInherited,
		LocalBundleObstructionInherited:      sym.LocalBundleObstructionInherited,
		CohomologyObstructionInherited:       sym.CohomologyObstructionInherited,

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
			"the contact overlap spectrum itself carries a 168-element Aut(Fano) action",
			"the trivial contact action can be equivariantly identified with the transitive Fano action",
			"transporting Aut(Fano) through an arbitrary contact-Fano bijection is canonical",
			"spectral labels provide a Fano-natural contact action",
			"contact modes may enter beta matching after an action chosen by convention",
			"Gate 119 derives alpha, physical thetaW, threshold beta corrections, M*, g_*, or masses",
		},
		RemainingUnknowns: []string{
			"canonical contact-side action of Aut(Fano) or a justified symmetry-breaking object",
			"natural contact-to-Fano assignment independent of the 7! bijections",
			"local bundle or constraint complex compatible with contact/Fano symmetry",
			"representation rows for contact modes",
			"threshold activation and decoupling law for any contact row",
			"threshold-corrected beta tensor",
		},
		RecommendedNextGate: "Gate 120 — contact spectral-invariant quotient / orbit-collapse theorem",
	}, nil
}

type rowCounts struct{ contact, positive, surviving, open int }

func buildRows(in []contactsymmetry.ContactRow) []ContactActionRow {
	rows := make([]ContactActionRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, ContactActionRow{
			Name:                      r.Name,
			ModeKind:                  r.ModeKind,
			Value:                     r.Value,
			FiniteOverlapPositive:     r.FiniteOverlapPositive,
			SurvivesCohomology:        r.SurvivesCohomology,
			SymmetryOpen:              r.Status == contactsymmetry.SelectorOpen,
			SpectralValuePreserved:    true,
			ContactActionDerived:      false,
			AutFanoActionOnContact:    false,
			EquivariantFanoAssignment: false,
			CanonicalContactFanoMap:   false,
			RepresentationRowDerived:  false,
			CanEnterBetaTensor:        false,
			ZeroRowProved:             false,
			Status:                    ActionOpen,
			Reason:                    "contact finite-overlap value is preserved only by the identity weighted action; no canonical Aut(Fano) action or equivariant assignment is derived",
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

func buildAttempts(s ContactActionSummary) []ActionAttempt {
	return []ActionAttempt{
		{
			Name:                      "canonical trivial contact action",
			Kind:                      TrivialContactActionAttempt,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			PreservesContactSpectrum:  true,
			ContactGroupOrder:         1,
			FanoGroupOrder:            s.FanoAutomorphismOrder,
			DefinesAutFanoAction:      false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"nontrivial contact action", "Fano fixed point", "equivariant bijection"},
			Detail:                    "the identity contact action is canonical, but equivariance with a transitive Fano action would require a globally fixed Fano point, and none exists",
		},
		{
			Name:                      "spectral-preserving contact automorphism group",
			Kind:                      SpectralPreservingActionAttempt,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			PreservesContactSpectrum:  true,
			ContactGroupOrder:         s.ContactWeightedAutomorphismOrder,
			FanoGroupOrder:            s.FanoAutomorphismOrder,
			DefinesAutFanoAction:      false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"168-element action preserving contact data", "equal-weight orbits", "representation row"},
			Detail:                    "all seven contact overlap values are distinct, so the weighted contact automorphism group has order one",
		},
		{
			Name:                        "transport Aut(Fano) through a chosen bijection",
			Kind:                        TransportedFanoActionAttempt,
			Constructed:                 true,
			UsesExtraConvention:         true,
			PreservesContactSpectrum:    false,
			ContactGroupOrder:           s.FanoAutomorphismOrder,
			FanoGroupOrder:              s.FanoAutomorphismOrder,
			DefinesAutFanoAction:        true,
			EquivariantAssignmentExists: true,
			CanonicalAssignmentDerived:  false,
			RejectedAsPremature:         true,
			MissingTerms:                []string{"canonical bijection", "spectral preservation", "symmetry-breaking theorem"},
			Detail:                      fmt.Sprintf("%d transported actions exist after choosing a contact-to-Fano bijection, but none is canonical under current data", s.TransportedActionCount),
		},
		{
			Name:                      "direct equivariant assignment search",
			Kind:                      DirectEquivariantAssignmentSearch,
			Constructed:               true,
			CanonicalUnderCurrentData: false,
			PreservesContactSpectrum:  false,
			ContactGroupOrder:         s.ContactWeightedAutomorphismOrder,
			FanoGroupOrder:            s.FanoAutomorphismOrder,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"derived contact-side group action", "naturality square", "canonical assignment"},
			Detail:                    "without a nontrivial derived contact-side action, the equivariance equation cannot select one of the 7! assignments",
		},
		{
			Name:                "representation row from contact action",
			Kind:                RepresentationRowFromActionAttempt,
			Constructed:         true,
			RejectedAsPremature: true,
			MissingTerms:        []string{"Aut(Fano) contact action", "local field bundle", "SU(3)c×SU(2)L×U(1)Y row", "mass/decoupling rule"},
			Detail:              "an automorphism action alone would still not be a gauge representation or threshold decoupling law",
		},
	}
}

func buildCriteria(s ContactActionSummary) []ActionCriterion {
	return []ActionCriterion{
		{Name: "Fano automorphism group", Required: true, Derived: s.FanoAutomorphismOrder == 168, Detail: "inherited 168-element transitive Aut(Fano)"},
		{Name: "contact weighted automorphism group", Required: true, Derived: true, Detail: fmt.Sprintf("order=%d; orbits=%v", s.ContactWeightedAutomorphismOrder, s.ContactWeightOrbitSizes)},
		{Name: "nontrivial contact-side Aut(Fano) action preserving contact data", Required: true, Derived: false, Detail: "distinct overlap weights force identity-only weighted action"},
		{Name: "canonical transported action", Required: true, Derived: false, Detail: "transport requires choosing one of 7! bijections"},
		{Name: "equivariant contact-Fano assignment", Required: true, Derived: false, Detail: "no canonical contact action/naturality square is available"},
		{Name: "representation row and beta permission", Required: true, Derived: false, Detail: "no local bundle, gauge row, mass activation, or decoupling rule follows"},
	}
}

func countRows(rows []ContactActionRow) rowCounts {
	c := rowCounts{contact: len(rows)}
	for _, r := range rows {
		if r.FiniteOverlapPositive {
			c.positive++
		}
		if r.SurvivesCohomology {
			c.surviving++
		}
		if r.Status == ActionOpen {
			c.open++
		}
	}
	return c
}

func contactValues(rows []contactsymmetry.ContactRow) []float64 {
	vals := make([]float64, 0, len(rows))
	for _, r := range rows {
		vals = append(vals, r.Value)
	}
	return vals
}

func spectralPreservingPermutations(values []float64, eps float64) []contactnaturality.Permutation {
	out := make([]contactnaturality.Permutation, 0)
	var current [7]int
	used := [7]bool{}
	var rec func(int)
	rec = func(pos int) {
		if pos == 7 {
			var p contactnaturality.Permutation
			copy(p[:], current[:])
			out = append(out, p)
			return
		}
		for j := 0; j < 7; j++ {
			if used[j] {
				continue
			}
			if math.Abs(values[pos]-values[j]) > eps {
				continue
			}
			used[j] = true
			current[pos] = j
			rec(pos + 1)
			used[j] = false
		}
	}
	if len(values) != 7 {
		return nil
	}
	rec(0)
	return out
}

func valuesDistinct(values []float64, eps float64) bool {
	for i := range values {
		for j := i + 1; j < len(values); j++ {
			if math.Abs(values[i]-values[j]) <= eps {
				return false
			}
		}
	}
	return len(values) == 7
}

func orbitSizesFromPermutations(perms []contactnaturality.Permutation) []int {
	seen := [7]bool{}
	sizes := []int{}
	for i := 0; i < 7; i++ {
		if seen[i] {
			continue
		}
		orbit := map[int]bool{i: true}
		changed := true
		for changed {
			changed = false
			for x := range orbit {
				for _, p := range perms {
					y := p[x]
					if !orbit[y] {
						orbit[y] = true
						changed = true
					}
				}
			}
		}
		for x := range orbit {
			seen[x] = true
		}
		sizes = append(sizes, len(orbit))
	}
	sort.Ints(sizes)
	return sizes
}

func isIdentity(p contactnaturality.Permutation) bool {
	for i := 0; i < 7; i++ {
		if p[i] != i {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	out := 1
	for i := 2; i <= n; i++ {
		out *= i
	}
	return out
}

func FormatRows(rows []ContactActionRow, limit int) string {
	parts := make([]string, 0, len(rows))
	for i, r := range rows {
		if limit > 0 && i >= limit {
			parts = append(parts, fmt.Sprintf("... +%d", len(rows)-i))
			break
		}
		parts = append(parts, fmt.Sprintf("%s(value=%.10g,action=%t,equiv=%t,rep=%t,beta=%t)", r.Name, r.Value, r.AutFanoActionOnContact, r.EquivariantFanoAssignment, r.RepresentationRowDerived, r.CanEnterBetaTensor))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(attempts []ActionAttempt) string {
	parts := make([]string, 0, len(attempts))
	for _, a := range attempts {
		parts = append(parts, fmt.Sprintf("%s(orderC=%d,orderF=%d,canon=%t,extra=%t,equiv=%t,beta=%t)", a.Name, a.ContactGroupOrder, a.FanoGroupOrder, a.CanonicalUnderCurrentData, a.UsesExtraConvention, a.EquivariantAssignmentExists, a.BetaRowPermitted))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(criteria []ActionCriterion) string {
	parts := make([]string, 0, len(criteria))
	for _, c := range criteria {
		mark := "missing"
		if c.Derived {
			mark = "derived"
		}
		parts = append(parts, fmt.Sprintf("%s=%s(%s)", c.Name, mark, c.Detail))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s ContactActionSummary) string {
	return fmt.Sprintf("Aut_contact(weighted)=%d orbits=%v distinct=%t; Aut_Fano=%d orbits=%v; mismatch=%t; transported actions=%d; canonical transported=%d", s.ContactWeightedAutomorphismOrder, s.ContactWeightOrbitSizes, s.SpectralValuesDistinct, s.FanoAutomorphismOrder, s.FanoPointOrbitSizes, s.OrderMismatch, s.TransportedActionCount, s.CanonicalTransportedActions)
}

func Join(xs []string) string {
	if len(xs) == 0 {
		return "[]"
	}
	return "[" + strings.Join(xs, "; ") + "]"
}
