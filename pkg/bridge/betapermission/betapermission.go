// Package betapermission implements Gate 112: contact-overlap
// representation-or-constraint dichotomy / beta-permission firewall.
//
// Gates 109-111 narrowed the unresolved threshold problem to the seven contact
// partial-overlap modes. They have positive finite overlap eigenvalues, but no
// local field map, gauge representation, Lorentz kinetic action, pole/residue
// theorem, constraint generator, ghost grading, nilpotent BRST operator, or
// cancellation ledger.
//
// This gate turns that epistemic state into an executable permission rule. A
// contact mode may enter threshold beta matching only by passing the physical
// branch: local field class, representation row, mass/activation/decoupling.
// Or it may be removed from threshold accounting only by passing the
// nonphysical branch: constraint/BRST class plus cancellation ledger. Current
// data passes neither branch, so all seven contact modes remain dichotomy-open
// and the beta-permission firewall stays closed.
package betapermission

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactfieldmap"
)

type PermissionClass string

const (
	BaselinePermittedSector PermissionClass = "baseline-permitted-sector"
	DichotomyOpenContact    PermissionClass = "dichotomy-open-contact-mode"
	ConstrainedVacuumEntry  PermissionClass = "constrained-vacuum-entry"
	ExcludedFiniteEntry     PermissionClass = "excluded-finite-entry"
)

type Branch string

const (
	PhysicalRepresentationBranch Branch = "physical-representation-branch"
	ConstraintBRSTBranch         Branch = "constraint-brst-branch"
)

type PermissionRow struct {
	Name, ModeKind string
	Class          PermissionClass
	Value          float64

	FiniteOverlapPositive bool

	LocalFieldClassDerived     bool
	GaugeRepresentationDerived bool
	LorentzKineticDerived      bool
	PoleResidueDerived         bool
	MassUnitDerived            bool
	ActivationDerived          bool
	DecouplingDerived          bool

	ConstraintGeneratorDerived bool
	GhostGradingDerived        bool
	NilpotentBRSTDerived       bool
	BRSTPairingDerived         bool
	CancellationLedgerDerived  bool

	PhysicalBranchComplete   bool
	ConstraintBranchComplete bool
	DichotomyResolved        bool
	CanEnterBetaTensor       bool
	ContributesZeroByProof   bool
	Reason                   string
}

type BranchRule struct {
	Branch       Branch
	Rule         string
	Constructed  bool
	CompleteNow  bool
	MissingTerms []string
	Detail       string
}

type FirewallWitness struct {
	Name                       string
	CompatibleWithFiniteData   bool
	WouldAllowBetaCorrection   bool
	WouldForceZeroContribution bool
	BlockedByFirewall          bool
	Detail                     string
}

type Analysis struct {
	ContactFieldMap contactfieldmap.Analysis

	Rows              []PermissionRow
	BranchRules       []BranchRule
	FirewallWitnesses []FirewallWitness

	TotalRows                 int
	BaselineRows              int
	ContactRows               int
	PositiveFiniteContactRows int
	DichotomyOpenRows         int
	ConstrainedVacuumRows     int
	ExcludedRows              int

	BetaPermissionFirewallConstructed bool
	PhysicalBranchRuleConstructed     bool
	ConstraintBranchRuleConstructed   bool
	PhysicalBranchCompleteRows        int
	ConstraintBranchCompleteRows      int
	ResolvedContactRows               int
	UnresolvedContactRows             int
	BetaCorrectionRowsAllowed         int
	ZeroContributionRowsProved        int

	RepresentationOrConstraintDichotomyDerived bool
	AllContactModesResolved                    bool
	ThresholdCorrectedBetaDerived              bool
	FullFiniteBetaMatchingTensorDerived        bool

	ActivationRuleDerived         bool
	DecouplingMatchingRuleDerived bool
	PhysicalMassUnitDerived       bool
	PhysicalScaleDerived          bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
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
		cf, err := contactfieldmap.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(cf)
	})
	return defaultValue, defaultErr
}

func Build(cf contactfieldmap.Analysis) (Analysis, error) {
	if cf.ContactRows != 7 || cf.PositiveFiniteContactRows != 7 || cf.ContactFieldClassDerived {
		return Analysis{}, fmt.Errorf("Gate 112 requires Gate 111 to leave seven positive contact modes field-class-open")
	}
	if cf.PhysicalLocalContactFieldsDerived || cf.ConstrainedContactClassDerived || cf.RegulatorGhostContactClassDerived || cf.VacuumFrustrationContactClassDerived {
		return Analysis{}, fmt.Errorf("Gate 112 requires no contact branch to be pre-selected")
	}
	if cf.BetaCorrectionRowsAllowed != 0 || cf.ThresholdCorrectedBetaDerived || cf.FullFiniteBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 112 requires beta matching to remain sealed before the permission firewall")
	}
	if cf.ResidualNullityAfter != 3 || cf.HiddenObservedInputUsed || cf.PhysicalWeakAngleDerived || cf.FineStructureDerived || cf.PhysicalMassesDerived {
		return Analysis{}, fmt.Errorf("Gate 112 refuses hidden physical input or changed residual nullity")
	}

	rows := buildRows(cf.Rows)
	rules := buildBranchRules()
	witnesses := buildWitnesses()
	counts := countRows(rows)
	physicalComplete := countPhysicalComplete(rows)
	constraintComplete := countConstraintComplete(rows)
	resolved := countResolvedContacts(rows)
	betaAllowed := countBetaAllowed(rows)
	zeroProved := countZeroProved(rows)

	truth := "Gate 112 constructs the beta-permission firewall for the seven contact partial-overlap modes. A contact mode can affect threshold beta matching only after the physical-representation branch is complete, and it can be removed as nonphysical only after the constraint/BRST-cancellation branch is complete. Current finite data completes neither branch for any contact mode, so the representation-or-constraint dichotomy remains unresolved and zero contact threshold beta rows are permitted."

	return Analysis{
		ContactFieldMap:   cf,
		Rows:              rows,
		BranchRules:       rules,
		FirewallWitnesses: witnesses,

		TotalRows:                 counts.total,
		BaselineRows:              counts.baseline,
		ContactRows:               counts.contact,
		PositiveFiniteContactRows: counts.positiveContact,
		DichotomyOpenRows:         counts.dichotomyOpen,
		ConstrainedVacuumRows:     counts.constrainedVacuum,
		ExcludedRows:              counts.excluded,

		BetaPermissionFirewallConstructed: true,
		PhysicalBranchRuleConstructed:     true,
		ConstraintBranchRuleConstructed:   true,
		PhysicalBranchCompleteRows:        physicalComplete,
		ConstraintBranchCompleteRows:      constraintComplete,
		ResolvedContactRows:               resolved,
		UnresolvedContactRows:             counts.contact - resolved,
		BetaCorrectionRowsAllowed:         betaAllowed,
		ZeroContributionRowsProved:        zeroProved,

		RepresentationOrConstraintDichotomyDerived: resolved == counts.contact && counts.contact > 0,
		AllContactModesResolved:                    resolved == counts.contact && counts.contact > 0,
		ThresholdCorrectedBetaDerived:              betaAllowed > 0,
		FullFiniteBetaMatchingTensorDerived:        false,

		ActivationRuleDerived:         false,
		DecouplingMatchingRuleDerived: false,
		PhysicalMassUnitDerived:       false,
		PhysicalScaleDerived:          false,

		ResidualNullityBefore:  cf.ResidualNullityAfter,
		ResidualNullityAfter:   cf.ResidualNullityAfter,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		HiddenObservedInputUsed:  false,
		TruthStatement:           truth,
		RejectedClaims: []string{
			"positive finite contact overlap permits beta matching",
			"contact modes may be counted as scalar doublets before local representation completion",
			"contact modes may be counted as singlets, regulators, or constraints before a proof of that branch",
			"a beta tensor can be completed by choosing the branch that preserves a desired RG output",
			"the permission firewall derives alpha, physical thetaW, masses, M*, or g_*",
		},
		RemainingUnknowns: []string{
			"U-32A-PHYSICAL-BRANCH: derive local support, Lorentz kinetic action, representation row, pole/residue, mass unit, activation, and decoupling for any contact mode that is physical",
			"U-32B-CONSTRAINT-BRANCH: derive constraint generator, ghost grading, nilpotent Q, BRST pairing, and cancellation ledger for any contact mode that is nonphysical",
			"U-32C-DICHOTOMY-SELECTOR: prove every contact partial-overlap mode belongs to exactly one branch",
			"U-32D-BETA-PERMISSION: only after branch resolution may a contact mode contribute a Δb_i row or a proven zero row",
			"U-32E-RG-FLOW: M*, g_*, L, and threshold matching remain undetermined",
		},
		RecommendedNextGate: "Gate 113 — contact-mode branch selector / finite constraint complex or local bundle construction attempt",
	}, nil
}

type rowCounts struct{ total, baseline, contact, positiveContact, dichotomyOpen, constrainedVacuum, excluded int }

func buildRows(in []contactfieldmap.FieldMapRow) []PermissionRow {
	rows := make([]PermissionRow, 0, len(in))
	for _, r := range in {
		row := PermissionRow{Name: r.Name, ModeKind: r.ModeKind, Value: r.Value, Reason: r.Reason}
		switch r.Class {
		case contactfieldmap.BaselineLocalSector:
			row.Class = BaselinePermittedSector
			row.LocalFieldClassDerived = true
			row.GaugeRepresentationDerived = true
			row.LorentzKineticDerived = true
			row.PoleResidueDerived = true
			row.MassUnitDerived = true
			row.ActivationDerived = true
			row.DecouplingDerived = true
			row.PhysicalBranchComplete = true
			row.DichotomyResolved = true
			row.Reason = "baseline scalar/contact aggregate is already part of the continuum inventory; Gate 112 audits only contact threshold permission"
		case contactfieldmap.FiniteOverlapMapOpen:
			row.Class = DichotomyOpenContact
			row.FiniteOverlapPositive = r.FiniteOverlapPositive
			row.Reason = "contact mode is positive finite-overlap data, but neither physical-representation nor constraint/BRST branch is complete"
		case contactfieldmap.ConstrainedVacuumLedger:
			row.Class = ConstrainedVacuumEntry
			row.ConstraintBranchComplete = true
			row.DichotomyResolved = true
			row.ContributesZeroByProof = true
			row.Reason = "B-sector gap remains an already-excluded constrained finite vacuum ledger entry"
		case contactfieldmap.ExcludedFiniteDiagnostic:
			row.Class = ExcludedFiniteEntry
			row.DichotomyResolved = true
			row.ContributesZeroByProof = true
			row.Reason = "finite diagnostic excluded before contact threshold permission"
		default:
			row.Class = DichotomyOpenContact
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

func buildBranchRules() []BranchRule {
	return []BranchRule{
		{
			Branch:      PhysicalRepresentationBranch,
			Rule:        "permit beta row only if local field class, Lorentz kinetic action, gauge representation, pole/residue, mass unit, activation, and decoupling are all derived",
			Constructed: true,
			CompleteNow: false,
			MissingTerms: []string{
				"local spacetime support / bundle map",
				"Lorentzian quadratic kinetic action",
				"SU(3)c×SU(2)L×U(1)Y representation row",
				"pole/residue theorem",
				"physical mass unit",
				"activation predicate",
				"decoupling/matching rule",
			},
			Detail: "Gate 111 supplies none of the contact-local physical-field preconditions",
		},
		{
			Branch:      ConstraintBRSTBranch,
			Rule:        "prove zero beta row only if constraint generator, ghost grading, nilpotent Q, BRST pairing, and cancellation/supertrace ledger are all derived",
			Constructed: true,
			CompleteNow: false,
			MissingTerms: []string{
				"constraint generator",
				"ghost number grading",
				"nilpotent BRST differential Q²=0",
				"BRST pair/quartet assignment",
				"supertrace/cancellation ledger",
			},
			Detail: "Gate 111 supplies none of the contact nonphysical-cancellation preconditions",
		},
	}
}

func buildWitnesses() []FirewallWitness {
	return []FirewallWitness{
		{Name: "count seven contact modes as scalar doublets", CompatibleWithFiniteData: true, WouldAllowBetaCorrection: true, BlockedByFirewall: true, Detail: "blocked because no local field map, representation row, mass unit, activation, or decoupling rule is derived"},
		{Name: "count seven contact modes as singlet scalars", CompatibleWithFiniteData: true, WouldAllowBetaCorrection: false, WouldForceZeroContribution: true, BlockedByFirewall: true, Detail: "blocked because singlet representation and physical field class are not derived"},
		{Name: "treat seven contact modes as constrained finite coordinates", CompatibleWithFiniteData: true, WouldAllowBetaCorrection: false, WouldForceZeroContribution: true, BlockedByFirewall: true, Detail: "blocked because no constraint generator removes them from propagation"},
		{Name: "treat seven contact modes as BRST/regulator modes", CompatibleWithFiniteData: true, WouldAllowBetaCorrection: false, WouldForceZeroContribution: true, BlockedByFirewall: true, Detail: "blocked because no ghost grading, nilpotent Q, pairing, or cancellation ledger is constructed"},
		{Name: "leave seven contact modes dichotomy-open", CompatibleWithFiniteData: true, WouldAllowBetaCorrection: false, WouldForceZeroContribution: false, BlockedByFirewall: false, Detail: "this is the only permission-safe state supported by current finite data"},
	}
}

func countRows(rows []PermissionRow) rowCounts {
	var c rowCounts
	c.total = len(rows)
	for _, r := range rows {
		switch r.Class {
		case BaselinePermittedSector:
			c.baseline++
		case DichotomyOpenContact:
			c.contact++
			c.dichotomyOpen++
			if r.FiniteOverlapPositive {
				c.positiveContact++
			}
		case ConstrainedVacuumEntry:
			c.constrainedVacuum++
		case ExcludedFiniteEntry:
			c.excluded++
		}
	}
	return c
}

func countPhysicalComplete(rows []PermissionRow) int {
	n := 0
	for _, r := range rows {
		if r.Class == DichotomyOpenContact && r.PhysicalBranchComplete {
			n++
		}
	}
	return n
}

func countConstraintComplete(rows []PermissionRow) int {
	n := 0
	for _, r := range rows {
		if r.Class == DichotomyOpenContact && r.ConstraintBranchComplete {
			n++
		}
	}
	return n
}

func countResolvedContacts(rows []PermissionRow) int {
	n := 0
	for _, r := range rows {
		if r.Class == DichotomyOpenContact && r.DichotomyResolved {
			n++
		}
	}
	return n
}

func countBetaAllowed(rows []PermissionRow) int {
	n := 0
	for _, r := range rows {
		if r.Class == DichotomyOpenContact && r.CanEnterBetaTensor {
			n++
		}
	}
	return n
}

func countZeroProved(rows []PermissionRow) int {
	n := 0
	for _, r := range rows {
		if r.Class == DichotomyOpenContact && r.ContributesZeroByProof {
			n++
		}
	}
	return n
}

func FormatRows(rows []PermissionRow, max int) string {
	if max <= 0 || max > len(rows) {
		max = len(rows)
	}
	parts := make([]string, 0, max+1)
	for i := 0; i < max; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s:%s:phys=%t:brst=%t:beta=%t:zero=%t", r.Name, r.Class, r.PhysicalBranchComplete, r.ConstraintBranchComplete, r.CanEnterBetaTensor, r.ContributesZeroByProof))
	}
	if max < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRules(rules []BranchRule) string {
	parts := make([]string, 0, len(rules))
	for _, r := range rules {
		parts = append(parts, fmt.Sprintf("%s(complete=%t,missing=%d)", r.Branch, r.CompleteNow, len(r.MissingTerms)))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatWitnesses(xs []FirewallWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(beta=%t,zero=%t,blocked=%t)", x.Name, x.WouldAllowBetaCorrection, x.WouldForceZeroContribution, x.BlockedByFirewall))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
