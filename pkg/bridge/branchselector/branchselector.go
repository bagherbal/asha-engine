// Package branchselector implements Gate 113: contact-mode branch selector /
// finite constraint-complex or local-bundle construction attempt.
//
// Gate 112 made the beta-permission rule executable but deliberately refused
// to classify the seven contact partial-overlap modes.  This gate attempts the
// two only honest continuations: construct a local bundle / field-map branch,
// or construct a finite constraint/BRST complex branch.  The current finite
// data does not provide a canonical base, fiber representation, transition
// functions, Lorentz kinetic operator, constraint differential, ghost grading,
// exactness proof, or cancellation ledger.  Therefore no branch selector is
// derived and the beta-permission firewall remains closed.
package branchselector

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/betapermission"
)

type SelectorClass string

const (
	BaselineAlreadyLocal         SelectorClass = "baseline-already-local"
	ContactBranchOpen            SelectorClass = "contact-branch-open"
	ConstrainedVacuumAlreadyZero SelectorClass = "constrained-vacuum-already-zero"
	ExcludedDiagnostic           SelectorClass = "excluded-diagnostic"
)

type BranchTarget string

const (
	LocalBundleBranch       BranchTarget = "local-bundle-field-branch"
	ConstraintComplexBranch BranchTarget = "finite-constraint-complex-branch"
)

type ModeDecisionRow struct {
	Name, ModeKind string
	Class          SelectorClass
	Value          float64

	FiniteOverlapPositive bool

	LocalBundleBranchAvailable bool
	BaseSpaceDerived           bool
	FiberRepresentationDerived bool
	TransitionFunctionsDerived bool
	SectionMapDerived          bool
	LorentzKineticDerived      bool
	PoleResidueDerived         bool
	MassActivationDerived      bool
	DecouplingDerived          bool

	ConstraintComplexBranchAvailable bool
	ChainGroupsDerived               bool
	DifferentialDerived              bool
	NilpotentDifferentialDerived     bool
	GhostGradingDerived              bool
	PairingDerived                   bool
	ExactnessOrCohomologyDerived     bool
	CancellationLedgerDerived        bool

	PhysicalBranchSelected   bool
	ConstraintBranchSelected bool
	BranchSelectorDerived    bool
	DichotomyResolved        bool
	CanEnterBetaTensor       bool
	ContributesZeroByProof   bool
	Reason                   string
}

type ConstructionAttempt struct {
	Branch       BranchTarget
	Constructed  bool
	CompleteNow  bool
	Canonical    bool
	MissingTerms []string
	Detail       string
}

type SelectorCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type AmbiguityWitness struct {
	Name                     string
	CompatibleWithFiniteData bool
	Branch                   BranchTarget
	WouldAllowBetaCorrection bool
	WouldProveZeroRow        bool
	BlockedBySelector        bool
	Detail                   string
}

type Analysis struct {
	BetaPermission betapermission.Analysis

	Rows                 []ModeDecisionRow
	ConstructionAttempts []ConstructionAttempt
	SelectorCriteria     []SelectorCriterion
	AmbiguityWitnesses   []AmbiguityWitness

	TotalRows                 int
	BaselineRows              int
	ContactRows               int
	PositiveFiniteContactRows int
	OpenContactRows           int
	ConstrainedVacuumRows     int
	ExcludedRows              int

	LocalBundleAttemptConstructed   bool
	LocalBundleBranchCompleteRows   int
	BaseSpaceDerived                bool
	FiberRepresentationDerived      bool
	TransitionFunctionsDerived      bool
	SectionMapDerived               bool
	LorentzKineticForContactDerived bool
	PoleResidueForContactDerived    bool
	MassActivationForContactDerived bool
	DecouplingForContactDerived     bool

	ConstraintComplexAttemptConstructed bool
	ConstraintComplexCompleteRows       int
	ChainGroupsDerived                  bool
	DifferentialDerived                 bool
	NilpotentDifferentialDerived        bool
	GhostGradingDerived                 bool
	PairingDerived                      bool
	ExactnessOrCohomologyDerived        bool
	CancellationLedgerDerived           bool

	BranchSelectorAttempted    bool
	BranchSelectorDerived      bool
	PhysicalSelectedRows       int
	ConstraintSelectedRows     int
	ResolvedContactRows        int
	UnresolvedContactRows      int
	BetaCorrectionRowsAllowed  int
	ZeroContributionRowsProved int

	ThresholdCorrectedBetaDerived              bool
	FullFiniteBetaMatchingTensorDerived        bool
	RepresentationOrConstraintDichotomyDerived bool

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
		bp, err := betapermission.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(bp)
	})
	return defaultValue, defaultErr
}

func Build(bp betapermission.Analysis) (Analysis, error) {
	if !bp.BetaPermissionFirewallConstructed || !bp.PhysicalBranchRuleConstructed || !bp.ConstraintBranchRuleConstructed {
		return Analysis{}, fmt.Errorf("Gate 113 requires Gate 112 beta-permission firewall")
	}
	if bp.ContactRows != 7 || bp.PositiveFiniteContactRows != 7 || bp.UnresolvedContactRows != 7 || bp.ResolvedContactRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 113 requires seven unresolved positive finite-overlap contact modes")
	}
	if bp.BetaCorrectionRowsAllowed != 0 || bp.ZeroContributionRowsProved != 0 || bp.ThresholdCorrectedBetaDerived || bp.FullFiniteBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 113 requires sealed beta-permission state")
	}
	if bp.ResidualNullityAfter != 3 || bp.HiddenObservedInputUsed || bp.PhysicalWeakAngleDerived || bp.FineStructureDerived || bp.PhysicalMassesDerived || bp.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 113 refuses hidden physical input or changed residual nullity")
	}

	rows := buildRows(bp.Rows)
	attempts := buildAttempts()
	criteria := buildCriteria()
	witnesses := buildWitnesses()
	counts := countRows(rows)

	truth := "Gate 113 attempts both honest branch-selector constructions for the seven contact partial-overlap modes: a local bundle/field-map branch and a finite constraint/BRST-complex branch. The current finite data supplies the contact overlap carrier, but it does not derive a canonical base space, fiber representation, transition functions, section map, Lorentz kinetic operator, pole/residue theorem, constraint differential, ghost grading, exactness/cohomology theorem, or cancellation ledger. Therefore no contact branch selector is derived; all seven modes remain branch-open and the beta-permission firewall remains closed."

	return Analysis{
		BetaPermission:       bp,
		Rows:                 rows,
		ConstructionAttempts: attempts,
		SelectorCriteria:     criteria,
		AmbiguityWitnesses:   witnesses,

		TotalRows:                 counts.total,
		BaselineRows:              counts.baseline,
		ContactRows:               counts.contact,
		PositiveFiniteContactRows: counts.positiveContact,
		OpenContactRows:           counts.openContact,
		ConstrainedVacuumRows:     counts.constrainedVacuum,
		ExcludedRows:              counts.excluded,

		LocalBundleAttemptConstructed:   true,
		LocalBundleBranchCompleteRows:   0,
		BaseSpaceDerived:                false,
		FiberRepresentationDerived:      false,
		TransitionFunctionsDerived:      false,
		SectionMapDerived:               false,
		LorentzKineticForContactDerived: false,
		PoleResidueForContactDerived:    false,
		MassActivationForContactDerived: false,
		DecouplingForContactDerived:     false,

		ConstraintComplexAttemptConstructed: true,
		ConstraintComplexCompleteRows:       0,
		ChainGroupsDerived:                  false,
		DifferentialDerived:                 false,
		NilpotentDifferentialDerived:        false,
		GhostGradingDerived:                 false,
		PairingDerived:                      false,
		ExactnessOrCohomologyDerived:        false,
		CancellationLedgerDerived:           false,

		BranchSelectorAttempted:    true,
		BranchSelectorDerived:      false,
		PhysicalSelectedRows:       0,
		ConstraintSelectedRows:     0,
		ResolvedContactRows:        0,
		UnresolvedContactRows:      counts.openContact,
		BetaCorrectionRowsAllowed:  0,
		ZeroContributionRowsProved: 0,

		ThresholdCorrectedBetaDerived:              false,
		FullFiniteBetaMatchingTensorDerived:        false,
		RepresentationOrConstraintDichotomyDerived: false,

		ResidualNullityBefore:  bp.ResidualNullityAfter,
		ResidualNullityAfter:   bp.ResidualNullityAfter,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"a finite overlap carrier alone defines a local continuum bundle",
			"a positive contact-overlap eigenvalue is a physical propagating threshold",
			"an arbitrary contact-mode pairing is a BRST complex",
			"the trivial zero differential proves cancellation of contact beta rows",
			"choosing the branch that preserves a desired RG output is a theorem",
			"Gate 113 derives alpha, physical thetaW, masses, M*, g_*, or threshold-corrected beta coefficients",
		},
		RemainingUnknowns: []string{
			"U-33A-LOCAL-BUNDLE: derive a finite-to-local base, fiber, transition, and section map for contact modes",
			"U-33B-REPRESENTATION: derive SU(3)c×SU(2)L×U(1)Y representation rows for any local contact field",
			"U-33C-CONSTRAINT-COMPLEX: derive chain groups, a canonical differential, Q²=0, ghost grading, pairing, and cancellation ledger",
			"U-33D-BRANCH-SELECTOR: prove each contact mode belongs to exactly one of the local-field or constraint-complex branches",
			"U-33E-THRESHOLD-MATCHING: only after branch selection may Δb_i(L), M*, g_*, alpha, thetaW, or masses be revisited",
		},
		RecommendedNextGate: "Gate 114 — finite contact constraint differential / cohomology obstruction theorem",
	}, nil
}

type rowCounts struct{ total, baseline, contact, positiveContact, openContact, constrainedVacuum, excluded int }

func buildRows(in []betapermission.PermissionRow) []ModeDecisionRow {
	rows := make([]ModeDecisionRow, 0, len(in))
	for _, r := range in {
		row := ModeDecisionRow{Name: r.Name, ModeKind: r.ModeKind, Value: r.Value, Reason: r.Reason}
		switch r.Class {
		case betapermission.BaselinePermittedSector:
			row.Class = BaselineAlreadyLocal
			row.LocalBundleBranchAvailable = true
			row.BaseSpaceDerived = true
			row.FiberRepresentationDerived = true
			row.TransitionFunctionsDerived = true
			row.SectionMapDerived = true
			row.LorentzKineticDerived = true
			row.PoleResidueDerived = true
			row.MassActivationDerived = true
			row.DecouplingDerived = true
			row.PhysicalBranchSelected = true
			row.BranchSelectorDerived = true
			row.DichotomyResolved = true
			row.Reason = "baseline scalar/contact aggregate is already in the continuum inventory; Gate 113 only audits threshold-contact branch selection"
		case betapermission.DichotomyOpenContact:
			row.Class = ContactBranchOpen
			row.FiniteOverlapPositive = r.FiniteOverlapPositive
			row.LocalBundleBranchAvailable = true
			row.ConstraintComplexBranchAvailable = true
			row.Reason = "contact mode has positive finite-overlap data, but neither a local bundle nor a constraint complex is canonical"
		case betapermission.ConstrainedVacuumEntry:
			row.Class = ConstrainedVacuumAlreadyZero
			row.ConstraintBranchSelected = true
			row.BranchSelectorDerived = true
			row.DichotomyResolved = true
			row.ContributesZeroByProof = true
			row.Reason = "already excluded as constrained finite vacuum ledger entry before the contact branch-selector problem"
		case betapermission.ExcludedFiniteEntry:
			row.Class = ExcludedDiagnostic
			row.BranchSelectorDerived = true
			row.DichotomyResolved = true
			row.ContributesZeroByProof = true
			row.Reason = "already excluded finite diagnostic before the contact branch-selector problem"
		default:
			row.Class = ContactBranchOpen
		}
		rows = append(rows, row)
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Class == rows[j].Class {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Class < rows[j].Class
	})
	return rows
}

func buildAttempts() []ConstructionAttempt {
	return []ConstructionAttempt{
		{
			Branch:       LocalBundleBranch,
			Constructed:  true,
			CompleteNow:  false,
			Canonical:    false,
			MissingTerms: []string{"base space / locality functor", "fiber representation", "transition functions", "section map", "Lorentz kinetic operator", "pole/residue theorem", "mass activation and decoupling"},
			Detail:       "finite contact carrier exists, but no canonical local bundle data is derived",
		},
		{
			Branch:       ConstraintComplexBranch,
			Constructed:  true,
			CompleteNow:  false,
			Canonical:    false,
			MissingTerms: []string{"chain groups", "canonical differential", "nilpotency Q²=0", "ghost grading", "pairing/quartet map", "exactness or cohomology theorem", "supertrace cancellation ledger"},
			Detail:       "finite contact carrier can be named as a candidate complex carrier, but no differential or cancellation theorem is selected",
		},
	}
}

func buildCriteria() []SelectorCriterion {
	return []SelectorCriterion{
		{Name: "finite contact carrier inherited", Required: true, Derived: true, Detail: "seven positive finite-overlap contact modes inherited from Gates 110-112"},
		{Name: "local bundle branch complete", Required: true, Derived: false, Detail: "requires base/fiber/transition/section/Lorentz/pole/mass data"},
		{Name: "constraint complex branch complete", Required: true, Derived: false, Detail: "requires chain groups, differential, Q²=0, grading, pairing, exactness/cohomology, cancellation"},
		{Name: "branch exclusivity", Required: true, Derived: false, Detail: "no theorem says a mode cannot satisfy both branches after extra structure is added"},
		{Name: "branch exhaustiveness", Required: true, Derived: false, Detail: "no theorem says every contact mode must become either local field or constraint/BRST mode"},
		{Name: "beta permission after branch selection", Required: true, Derived: false, Detail: "zero beta rows and nonzero beta rows both remain forbidden for contact modes"},
	}
}

func buildWitnesses() []AmbiguityWitness {
	return []AmbiguityWitness{
		{Name: "local bundle with scalar-doublet contact modes", CompatibleWithFiniteData: true, Branch: LocalBundleBranch, WouldAllowBetaCorrection: true, BlockedBySelector: true, Detail: "compatible only after adding representation/locality/mass data not selected by the finite carrier"},
		{Name: "local bundle with singlet contact modes", CompatibleWithFiniteData: true, Branch: LocalBundleBranch, WouldAllowBetaCorrection: false, WouldProveZeroRow: true, BlockedBySelector: true, Detail: "compatible only after deriving singlet representation and physical field map"},
		{Name: "constraint complex with trivial differential", CompatibleWithFiniteData: true, Branch: ConstraintComplexBranch, WouldProveZeroRow: false, BlockedBySelector: true, Detail: "Q=0 is nilpotent by convention but proves no pairing, exactness, or cancellation ledger"},
		{Name: "constraint complex with arbitrary quartet pairing", CompatibleWithFiniteData: true, Branch: ConstraintComplexBranch, WouldProveZeroRow: true, BlockedBySelector: true, Detail: "arbitrary pairing could cancel rows but is not canonical or derived from contact geometry"},
		{Name: "keep contact modes branch-open", CompatibleWithFiniteData: true, Branch: ConstraintComplexBranch, BlockedBySelector: false, Detail: "permission-safe state supported by the current finite data"},
	}
}

func countRows(rows []ModeDecisionRow) rowCounts {
	var c rowCounts
	c.total = len(rows)
	for _, r := range rows {
		switch r.Class {
		case BaselineAlreadyLocal:
			c.baseline++
		case ContactBranchOpen:
			c.contact++
			c.openContact++
			if r.FiniteOverlapPositive {
				c.positiveContact++
			}
		case ConstrainedVacuumAlreadyZero:
			c.constrainedVacuum++
		case ExcludedDiagnostic:
			c.excluded++
		}
	}
	return c
}

func FormatRows(rows []ModeDecisionRow, max int) string {
	if max <= 0 || max > len(rows) {
		max = len(rows)
	}
	parts := make([]string, 0, max+1)
	for i := 0; i < max; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s:%s:local=%t:complex=%t:selector=%t:beta=%t:zero=%t", r.Name, r.Class, r.PhysicalBranchSelected, r.ConstraintBranchSelected, r.BranchSelectorDerived, r.CanEnterBetaTensor, r.ContributesZeroByProof))
	}
	if max < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(xs []ConstructionAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(constructed=%t,complete=%t,canonical=%t,missing=%d)", x.Branch, x.Constructed, x.CompleteNow, x.Canonical, len(x.MissingTerms)))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []SelectorCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatWitnesses(xs []AmbiguityWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(branch=%s,beta=%t,zero=%t,blocked=%t)", x.Name, x.Branch, x.WouldAllowBetaCorrection, x.WouldProveZeroRow, x.BlockedBySelector))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
