// Package contactcohomology implements Gate 114: finite contact constraint
// differential / cohomology obstruction theorem.
//
// Gate 113 attempted both honest completions for the seven contact
// partial-overlap modes. The local-bundle branch and the constraint/BRST branch
// both remained incomplete. This gate focuses only on the constraint branch: can
// the seven-mode finite contact carrier itself select a canonical differential,
// prove Q²=0, compute cohomology, and produce a cancellation ledger?
//
// The answer is no. A trivial zero differential is square-zero but leaves all
// seven modes in cohomology and therefore proves no cancellation. Several
// nontrivial differentials can be written after choosing an order, pairing, or
// Fano-style incidence convention, but those choices are not selected by the
// current finite data and do not supply the full BRST ledger. The contact beta
// firewall therefore remains closed.
package contactcohomology

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/branchselector"
)

type DifferentialKind string

const (
	ZeroDifferential     DifferentialKind = "zero-differential"
	IdentityDifferential DifferentialKind = "identity-differential"
	OrderedShift         DifferentialKind = "ordered-shift"
	PairingDifferential  DifferentialKind = "pairing-differential"
	FanoIncidenceAttempt DifferentialKind = "fano-incidence-attempt"
)

type DifferentialAttempt struct {
	Name string
	Kind DifferentialKind

	Constructed               bool
	UsesExtraConvention       bool
	CanonicalUnderCurrentData bool

	SquareZero                bool
	Nontrivial                bool
	GhostGradingDerived       bool
	PairingDerived            bool
	ExactnessDerived          bool
	CohomologyDimension       int
	Acyclic                   bool
	CancellationLedgerDerived bool
	BetaZeroRowProved         bool

	MissingTerms []string
	Detail       string
}

type ModeCohomologyRow struct {
	Name, ModeKind                     string
	Value                              float64
	FiniteOverlapPositive              bool
	CandidateChainGroup                string
	SurvivesZeroDifferentialCohomology bool
	ResolvedByConstraintComplex        bool
	ZeroContributionProved             bool
	CanEnterBetaTensor                 bool
	Reason                             string
}

type ObstructionCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Analysis struct {
	BranchSelector branchselector.Analysis

	Rows                 []ModeCohomologyRow
	DifferentialAttempts []DifferentialAttempt
	ObstructionCriteria  []ObstructionCriterion

	ContactRows                  int
	PositiveFiniteContactRows    int
	OpenContactRowsBefore        int
	OpenContactRowsAfter         int
	CandidateChainGroupDimension int

	ConstraintDifferentialAttempted        bool
	ChainGroupCarrierConstructed           bool
	CanonicalDifferentialDerived           bool
	NontrivialNilpotentDifferentialDerived bool
	ZeroDifferentialConstructed            bool
	ZeroDifferentialSquareZero             bool
	ZeroDifferentialCohomologyDimension    int
	ZeroDifferentialProvesCancellation     bool

	AnyNontrivialCandidateConstructed bool
	AnyNontrivialCandidateCanonical   bool
	AnyNontrivialCandidateSquareZero  bool
	GhostGradingDerived               bool
	PairingDerived                    bool
	ExactnessOrCohomologyDerived      bool
	AcyclicComplexDerived             bool
	CancellationLedgerDerived         bool
	ConstraintComplexCompleteRows     int
	ContactZeroRowsProved             int
	ContactBetaRowsAllowed            int

	CohomologyObstructionDerived                bool
	NoCanonicalBRSTDifferentialUnderCurrentData bool
	BranchSelectorDerived                       bool
	RepresentationOrConstraintDichotomyDerived  bool
	ThresholdCorrectedBetaDerived               bool
	FullFiniteBetaMatchingTensorDerived         bool

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
		bs, err := branchselector.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(bs)
	})
	return defaultValue, defaultErr
}

func Build(bs branchselector.Analysis) (Analysis, error) {
	if !bs.BranchSelectorAttempted || bs.BranchSelectorDerived || bs.ResolvedContactRows != 0 || bs.UnresolvedContactRows != 7 {
		return Analysis{}, fmt.Errorf("Gate 114 requires Gate 113 to leave seven contact modes branch-open")
	}
	if bs.ContactRows != 7 || bs.PositiveFiniteContactRows != 7 || bs.BetaCorrectionRowsAllowed != 0 || bs.ZeroContributionRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 114 requires seven positive contact modes and sealed beta rows")
	}
	if bs.ConstraintComplexCompleteRows != 0 || bs.CancellationLedgerDerived || bs.NilpotentDifferentialDerived || bs.DifferentialDerived {
		return Analysis{}, fmt.Errorf("Gate 114 requires no pre-existing constraint differential or cancellation ledger")
	}
	if bs.ResidualNullityAfter != 3 || bs.HiddenObservedInputUsed || bs.PhysicalWeakAngleDerived || bs.FineStructureDerived || bs.PhysicalMassesDerived || bs.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 114 refuses hidden physical input or changed residual nullity")
	}

	rows := buildRows(bs.Rows)
	attempts := buildDifferentialAttempts(len(rows))
	criteria := buildCriteria()
	counts := countRows(rows)

	nontrivialConstructed := false
	nontrivialCanonical := false
	nontrivialSquareZero := false
	for _, a := range attempts {
		if a.Nontrivial && a.Constructed {
			nontrivialConstructed = true
		}
		if a.Nontrivial && a.CanonicalUnderCurrentData {
			nontrivialCanonical = true
		}
		if a.Nontrivial && a.SquareZero && a.CanonicalUnderCurrentData {
			nontrivialSquareZero = true
		}
	}

	truth := "Gate 114 tests the constraint/BRST branch directly by attempting to put a differential on the seven contact partial-overlap modes. The zero differential is square-zero but leaves a seven-dimensional cohomology and proves no cancellation. Nontrivial candidates require extra ordering, pairing, or incidence conventions and are not canonical under the current finite data. Therefore no nilpotent BRST differential, acyclic complex, or cancellation ledger is derived; all seven contact modes remain cohomology-open and threshold beta matching remains sealed."

	return Analysis{
		BranchSelector:       bs,
		Rows:                 rows,
		DifferentialAttempts: attempts,
		ObstructionCriteria:  criteria,

		ContactRows:                  counts.contact,
		PositiveFiniteContactRows:    counts.positiveContact,
		OpenContactRowsBefore:        bs.UnresolvedContactRows,
		OpenContactRowsAfter:         counts.contact,
		CandidateChainGroupDimension: counts.contact,

		ConstraintDifferentialAttempted:        true,
		ChainGroupCarrierConstructed:           counts.contact == 7,
		CanonicalDifferentialDerived:           false,
		NontrivialNilpotentDifferentialDerived: false,
		ZeroDifferentialConstructed:            true,
		ZeroDifferentialSquareZero:             true,
		ZeroDifferentialCohomologyDimension:    counts.contact,
		ZeroDifferentialProvesCancellation:     false,

		AnyNontrivialCandidateConstructed: nontrivialConstructed,
		AnyNontrivialCandidateCanonical:   nontrivialCanonical,
		AnyNontrivialCandidateSquareZero:  nontrivialSquareZero,
		GhostGradingDerived:               false,
		PairingDerived:                    false,
		ExactnessOrCohomologyDerived:      false,
		AcyclicComplexDerived:             false,
		CancellationLedgerDerived:         false,
		ConstraintComplexCompleteRows:     0,
		ContactZeroRowsProved:             0,
		ContactBetaRowsAllowed:            0,

		CohomologyObstructionDerived:                true,
		NoCanonicalBRSTDifferentialUnderCurrentData: true,
		BranchSelectorDerived:                       false,
		RepresentationOrConstraintDichotomyDerived:  false,
		ThresholdCorrectedBetaDerived:               false,
		FullFiniteBetaMatchingTensorDerived:         false,

		ResidualNullityBefore:  bs.ResidualNullityAfter,
		ResidualNullityAfter:   bs.ResidualNullityAfter,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"the zero differential proves contact-mode cancellation",
			"any arbitrary pairing of contact modes is a canonical BRST complex",
			"Fano/octonion incidence alone supplies a nilpotent constraint differential on the contact-overlap carrier",
			"positive contact overlap plus an ordered shift is enough to define physical threshold decoupling",
			"seven contact modes can be fully BRST-paired without extra structure",
			"Gate 114 derives threshold-corrected beta coefficients, alpha, physical thetaW, M*, g_*, or masses",
		},
		RemainingUnknowns: []string{
			"U-34A-GRADED-CARRIER: derive chain groups with ghost degree for the seven contact modes",
			"U-34B-CANONICAL-DIFFERENTIAL: derive a nontrivial Q selected by finite contact data rather than by convention",
			"U-34C-NILPOTENCY: prove Q²=0 as a typed theorem, not as an imposed property",
			"U-34D-COHOMOLOGY: compute exactness/cohomology and decide whether any contact classes survive",
			"U-34E-CANCELLATION-LEDGER: derive supertrace/beta zero rows only after a real constraint complex exists",
			"U-34F-LOCAL-BRANCH: if no constraint complex exists, return to local bundle/field map construction",
		},
		RecommendedNextGate: "Gate 115 — contact local-bundle obstruction / representation-row construction attempt",
	}, nil
}

type rowCounts struct{ contact, positiveContact int }

func buildRows(in []branchselector.ModeDecisionRow) []ModeCohomologyRow {
	rows := make([]ModeCohomologyRow, 0, 7)
	for _, r := range in {
		if r.Class != branchselector.ContactBranchOpen {
			continue
		}
		rows = append(rows, ModeCohomologyRow{
			Name:                               r.Name,
			ModeKind:                           r.ModeKind,
			Value:                              r.Value,
			FiniteOverlapPositive:              r.FiniteOverlapPositive,
			CandidateChainGroup:                "C_contact^0",
			SurvivesZeroDifferentialCohomology: true,
			ResolvedByConstraintComplex:        false,
			ZeroContributionProved:             false,
			CanEnterBetaTensor:                 false,
			Reason:                             "survives the trivial differential; no nontrivial canonical Q, ghost grading, exactness theorem, or cancellation row is derived",
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

func buildDifferentialAttempts(n int) []DifferentialAttempt {
	return []DifferentialAttempt{
		{
			Name:                      "Q0 on C_contact^0",
			Kind:                      ZeroDifferential,
			Constructed:               true,
			UsesExtraConvention:       false,
			CanonicalUnderCurrentData: true,
			SquareZero:                true,
			Nontrivial:                false,
			CohomologyDimension:       n,
			Acyclic:                   false,
			CancellationLedgerDerived: false,
			BetaZeroRowProved:         false,
			MissingTerms:              []string{"nontrivial image", "exactness", "ghost pairing", "supertrace cancellation ledger"},
			Detail:                    fmt.Sprintf("Q0²=0, but H(C,Q0) has dimension %d; no contact mode is cancelled", n),
		},
		{
			Name:                      "identity map on contact carrier",
			Kind:                      IdentityDifferential,
			Constructed:               true,
			UsesExtraConvention:       false,
			CanonicalUnderCurrentData: false,
			SquareZero:                false,
			Nontrivial:                true,
			CohomologyDimension:       -1,
			MissingTerms:              []string{"Q²=0", "degree +1 ghost grading", "constraint interpretation"},
			Detail:                    "identity is a basis-independent endomorphism, but I²=I not 0, so it cannot be a BRST differential",
		},
		{
			Name:                      "spectral-value ordered shift",
			Kind:                      OrderedShift,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			SquareZero:                false,
			Nontrivial:                true,
			CohomologyDimension:       -1,
			MissingTerms:              []string{"canonical order orientation", "Q²=0", "ghost grading", "pairing"},
			Detail:                    "requires choosing an orientation of the contact-overlap order; adjacent shift has nonzero second step",
		},
		{
			Name:                      "pair/quartet cancellation map",
			Kind:                      PairingDifferential,
			Constructed:               false,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			SquareZero:                false,
			Nontrivial:                true,
			CohomologyDimension:       -1,
			MissingTerms:              []string{"even complete pairing", "canonical partner map", "ghost signs", "quartet cancellation theorem"},
			Detail:                    fmt.Sprintf("the contact carrier has odd dimension %d, so complete pair/quartet cancellation cannot be inferred without adding structure", n),
		},
		{
			Name:                      "Fano/octonion incidence candidate",
			Kind:                      FanoIncidenceAttempt,
			Constructed:               false,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			SquareZero:                false,
			Nontrivial:                true,
			CohomologyDimension:       -1,
			MissingTerms:              []string{"typed chain groups", "oriented incidence boundary", "nilpotency proof on the contact carrier", "BRST grading"},
			Detail:                    "seven labels resonate with Fano incidence, but no theorem maps the contact-overlap carrier to an oriented chain complex with ∂²=0",
		},
	}
}

func buildCriteria() []ObstructionCriterion {
	return []ObstructionCriterion{
		{Name: "seven-mode candidate chain carrier", Required: true, Derived: true, Detail: "C_contact^0 can be named as the vector carrier of seven positive finite-overlap modes"},
		{Name: "canonical nontrivial differential", Required: true, Derived: false, Detail: "all nontrivial candidates require extra order/pairing/incidence choices"},
		{Name: "nilpotency Q²=0", Required: true, Derived: false, Detail: "only the trivial Q0 is square-zero; it is not a cancellation complex"},
		{Name: "ghost grading", Required: true, Derived: false, Detail: "no finite degree assignment or ghost number is selected"},
		{Name: "pair/quartet ledger", Required: true, Derived: false, Detail: "odd seven-mode carrier and absent partner map block full cancellation"},
		{Name: "acyclic/exact cohomology", Required: true, Derived: false, Detail: "zero differential gives seven surviving cohomology classes"},
		{Name: "beta zero-row ledger", Required: true, Derived: false, Detail: "no contact zero contribution is proven"},
	}
}

func countRows(rows []ModeCohomologyRow) rowCounts {
	c := rowCounts{contact: len(rows)}
	for _, r := range rows {
		if r.FiniteOverlapPositive {
			c.positiveContact++
		}
	}
	return c
}

func FormatRows(rows []ModeCohomologyRow, limit int) string {
	parts := make([]string, 0, len(rows))
	for i, r := range rows {
		if limit > 0 && i >= limit {
			parts = append(parts, fmt.Sprintf("... +%d", len(rows)-i))
			break
		}
		parts = append(parts, fmt.Sprintf("%s(H0=%t,zero=%t,beta=%t)", r.Name, r.SurvivesZeroDifferentialCohomology, r.ZeroContributionProved, r.CanEnterBetaTensor))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatDifferentials(xs []DifferentialAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s,canonical=%t,Q²=0:%t,H=%d,cancel=%t)", x.Name, x.Kind, x.CanonicalUnderCurrentData, x.SquareZero, x.CohomologyDimension, x.CancellationLedgerDerived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []ObstructionCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
