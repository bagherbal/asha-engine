// Package contactlqbetapermission implements Gate 134: leptoquark
// hypercharge-row and local-field obstruction / beta-permission theorem.
//
// Gate 133 proved that the six current-side leptoquark slots have only a
// color(3) × real-orientation(2) current tensor and that the second factor is
// not an SU(2)_L weak-doublet action. Gate 134 audits the remaining objects
// required before those slots may contribute to threshold beta matching:
// hypercharge, local field variables, Lorentz kinetic/residue data, mass
// activation, and decoupling. The gate is deliberately a permission theorem:
// no beta row is allowed unless all representation and propagation data are
// complete. Current finite data leaves every permission bit false.
package contactlqbetapermission

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactlqsu2"
)

type RequirementKind string

const (
	HyperchargeRequirement RequirementKind = "hypercharge-row"
	LocalFieldRequirement  RequirementKind = "local-field-map"
	LorentzRequirement     RequirementKind = "lorentz-kinetic-residue"
	MassRequirement        RequirementKind = "mass-activation"
	DecouplingRequirement  RequirementKind = "decoupling-matching"
	SU2Requirement         RequirementKind = "weak-su2-action"
)

type Requirement struct {
	Name                  string
	Kind                  RequirementKind
	Required              bool
	Derived               bool
	CurrentSideDiagnostic bool
	ContactSideSemantic   bool
	MayBorrowFromMatter   bool
	RequiresHiddenChoice  bool
	RequiresObservedInput bool
	HiddenChoices         int
	Obstruction           string
}

type CandidateRow struct {
	Name                   string
	Slots                  int
	ColorTriplet           bool
	RealOrientationPair    bool
	WeakDoubletDerived     bool
	HyperchargeDerived     bool
	LocalFieldDerived      bool
	LorentzKineticDerived  bool
	PoleResidueDerived     bool
	MassActivationDerived  bool
	DecouplingDerived      bool
	RepresentationComplete bool
	BetaPermitted          bool
	ZeroContributionProved bool
	RequiresS6Choice       bool
	RequiresObservedInput  bool
	Obstruction            string
}

type Summary struct {
	ContactRows                int
	LeptoquarkRows             int
	CurrentLQSlots             int
	ColorDirections            int
	RealOrientations           int
	WeakSU2ActionDerived       bool
	HyperchargeRowsDerived     int
	LocalFieldRowsDerived      int
	LorentzKineticRowsDerived  int
	MassActivationRowsDerived  int
	DecouplingRowsDerived      int
	RepresentationCompleteRows int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualS6Choices          int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous contactlqsu2.Analysis

	Requirements  []Requirement
	CandidateRows []CandidateRow
	Summary       Summary

	ContactRows                    int
	LeptoquarkRows                 int
	CurrentLQSlots                 int
	ColorDirections                int
	RealOrientations               int
	Gate133SU2ObstructionInherited bool
	HyperchargeRowDerived          bool
	LocalFieldMapDerived           bool
	LorentzKineticRowDerived       bool
	PoleResidueTheoremDerived      bool
	MassActivationDerived          bool
	DecouplingRuleDerived          bool
	RepresentationCompleteRows     int
	RepresentationOpenRows         int
	ContactBetaRowsAllowed         int
	ContactZeroRowsProved          int
	FullBetaMatchingTensorDerived  bool
	ThresholdCorrectedBetaDerived  bool
	BetaPermissionFirewallClosed   bool

	ResidualS6Choices        int
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
		prev, err := contactlqsu2.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactlqsu2.Analysis) (Analysis, error) {
	if !prev.S6ObstructionInherited || prev.LeptoquarkRows != 6 || prev.TotalCurrentLQSlots != 6 || prev.ColorPlanes != 3 || prev.RealOrientationsPerColor != 2 {
		return Analysis{}, fmt.Errorf("Gate 134 requires Gate 133 leptoquark six-slot current tensor")
	}
	if prev.SU2WeakDoubletActionDerived || prev.WeakDoubletSemanticsDerived || prev.NonAbelianSU2TripleDerived || prev.HyperchargeRowDerived || prev.LocalFieldMapDerived {
		return Analysis{}, fmt.Errorf("Gate 134 requires Gate 133 SU(2)L, hypercharge, and local-field rows to remain underived")
	}
	if prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 || prev.ThresholdCorrectedBetaDerived || prev.FullBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 134 requires Gate 133 beta firewall to remain closed")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 134 refuses hidden observed input")
	}

	requirements := buildRequirements(prev)
	rows := buildCandidateRows(prev)

	hyperRows := count(rows, func(r CandidateRow) bool { return r.HyperchargeDerived })
	localRows := count(rows, func(r CandidateRow) bool { return r.LocalFieldDerived })
	lorentzRows := count(rows, func(r CandidateRow) bool { return r.LorentzKineticDerived && r.PoleResidueDerived })
	massRows := count(rows, func(r CandidateRow) bool { return r.MassActivationDerived })
	decouplingRows := count(rows, func(r CandidateRow) bool { return r.DecouplingDerived })
	repRows := count(rows, func(r CandidateRow) bool { return r.RepresentationComplete })
	betaRows := count(rows, func(r CandidateRow) bool { return r.BetaPermitted })
	zeroRows := count(rows, func(r CandidateRow) bool { return r.ZeroContributionProved })

	summary := Summary{
		ContactRows:                7,
		LeptoquarkRows:             6,
		CurrentLQSlots:             6,
		ColorDirections:            3,
		RealOrientations:           2,
		WeakSU2ActionDerived:       false,
		HyperchargeRowsDerived:     hyperRows,
		LocalFieldRowsDerived:      localRows,
		LorentzKineticRowsDerived:  lorentzRows,
		MassActivationRowsDerived:  massRows,
		DecouplingRowsDerived:      decouplingRows,
		RepresentationCompleteRows: repRows,
		ContactBetaRowsAllowed:     betaRows,
		ContactZeroRowsProved:      zeroRows,
		ResidualS6Choices:          prev.Summary.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 134 turns the leptoquark-contact shortcut into an executable permission theorem. A threshold beta row for the six contact leptoquark candidates would require, in order, a non-abelian SU(2)L action, hypercharge row, local field map, Lorentz kinetic pole/residue theorem, mass activation, and decoupling/matching rule. Gate 133 already blocks the SU(2)L action. Gate 134 verifies that hypercharge, local-field, kinetic, mass, and decoupling rows are also absent. Therefore the six leptoquark-shaped current slots remain current diagnostics only; no contact threshold beta contribution or zero-cancellation row is permitted."

	return Analysis{
		Previous:      prev,
		Requirements:  requirements,
		CandidateRows: rows,
		Summary:       summary,

		ContactRows:                    7,
		LeptoquarkRows:                 6,
		CurrentLQSlots:                 6,
		ColorDirections:                3,
		RealOrientations:               2,
		Gate133SU2ObstructionInherited: !prev.SU2WeakDoubletActionDerived && !prev.NonAbelianSU2TripleDerived,
		HyperchargeRowDerived:          hyperRows > 0,
		LocalFieldMapDerived:           localRows > 0,
		LorentzKineticRowDerived:       lorentzRows > 0,
		PoleResidueTheoremDerived:      lorentzRows > 0,
		MassActivationDerived:          massRows > 0,
		DecouplingRuleDerived:          decouplingRows > 0,
		RepresentationCompleteRows:     repRows,
		RepresentationOpenRows:         7,
		ContactBetaRowsAllowed:         betaRows,
		ContactZeroRowsProved:          zeroRows,
		FullBetaMatchingTensorDerived:  false,
		ThresholdCorrectedBetaDerived:  false,
		BetaPermissionFirewallClosed:   repRows == 0 && betaRows == 0 && zeroRows == 0,

		ResidualS6Choices:        prev.Summary.ResidualS6Choices,
		ResidualNullityBefore:    prev.ResidualNullityAfter,
		ResidualNullityAfter:     prev.ResidualNullityAfter,
		HiddenObservedInputUsed:  false,
		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"six current leptoquark slots may be counted as physical threshold fields",
			"real symmetric/skew orientation supplies weak hypercharge or SU(2)L semantics",
			"hypercharge can be borrowed from matter tables without a contact field map",
			"positive contact overlap is enough for Lorentz pole/residue data",
			"threshold beta rows can be assigned before mass activation and decoupling",
			"observed alpha/thetaW/masses may select the missing contact rows",
		},
		RemainingUnknowns: []string{
			"contact leptoquark hypercharge row",
			"local spacetime field map for contact leptoquark candidates",
			"Lorentz kinetic denominator and residue signature",
			"mass activation unit and threshold scale",
			"decoupling/matching rule for beta coefficients",
			"canonical S6-breaking assignment of current slots to contact rows",
		},
		RecommendedNextGate: "Gate 135 — leptoquark contact hypercharge source / B-L and charge-lattice obstruction theorem",
	}, nil
}

func buildRequirements(prev contactlqsu2.Analysis) []Requirement {
	hidden := prev.Summary.ResidualS6Choices
	return []Requirement{
		{Name: "non-abelian weak SU(2)L action", Kind: SU2Requirement, Required: true, Derived: false, CurrentSideDiagnostic: false, ContactSideSemantic: false, RequiresHiddenChoice: true, HiddenChoices: hidden, Obstruction: "Gate 133 exposed only abelian SO(2) orientation rotations"},
		{Name: "leptoquark hypercharge row", Kind: HyperchargeRequirement, Required: true, Derived: false, CurrentSideDiagnostic: false, ContactSideSemantic: false, MayBorrowFromMatter: false, RequiresHiddenChoice: true, HiddenChoices: hidden, Obstruction: "no contact-side Y eigenvalue or charge lattice row is selected"},
		{Name: "local contact field map", Kind: LocalFieldRequirement, Required: true, Derived: false, ContactSideSemantic: false, RequiresHiddenChoice: true, HiddenChoices: hidden, Obstruction: "current slots do not define local spacetime variables on contact rows"},
		{Name: "Lorentz kinetic pole/residue theorem", Kind: LorentzRequirement, Required: true, Derived: false, ContactSideSemantic: false, Obstruction: "finite overlap positivity is not a Lorentz propagator denominator or residue sign"},
		{Name: "mass activation unit", Kind: MassRequirement, Required: true, Derived: false, ContactSideSemantic: false, RequiresObservedInput: false, Obstruction: "no finite dimensional scale converts the contact row into a threshold mass"},
		{Name: "decoupling and beta matching rule", Kind: DecouplingRequirement, Required: true, Derived: false, ContactSideSemantic: false, Obstruction: "without local representation and mass data there is no allowed Delta b_i contribution"},
	}
}

func buildCandidateRows(prev contactlqsu2.Analysis) []CandidateRow {
	names := []string{
		"LQ_color1_symmetric", "LQ_color1_skew",
		"LQ_color2_symmetric", "LQ_color2_skew",
		"LQ_color3_symmetric", "LQ_color3_skew",
	}
	out := make([]CandidateRow, 0, len(names))
	for _, name := range names {
		out = append(out, CandidateRow{
			Name:                   name,
			Slots:                  1,
			ColorTriplet:           true,
			RealOrientationPair:    true,
			WeakDoubletDerived:     false,
			HyperchargeDerived:     false,
			LocalFieldDerived:      false,
			LorentzKineticDerived:  false,
			PoleResidueDerived:     false,
			MassActivationDerived:  false,
			DecouplingDerived:      false,
			RepresentationComplete: false,
			BetaPermitted:          false,
			ZeroContributionProved: false,
			RequiresS6Choice:       true,
			RequiresObservedInput:  false,
			Obstruction:            fmt.Sprintf("%s is a current-side color/orientation diagnostic; Gate 133 leaves SU(2)L open and Gate 134 finds no hypercharge, local field, mass, or decoupling row", name),
		})
	}
	return out
}

func count[T any](xs []T, pred func(T) bool) int {
	n := 0
	for _, x := range xs {
		if pred(x) {
			n++
		}
	}
	return n
}

func FormatRequirements(xs []Requirement) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s required=%t derived=%t currentDiag=%t contactSemantic=%t borrowMatter=%t hiddenChoice=%t observed=%t hidden=%d obstruction=%s)", x.Name, x.Kind, x.Required, x.Derived, x.CurrentSideDiagnostic, x.ContactSideSemantic, x.MayBorrowFromMatter, x.RequiresHiddenChoice, x.RequiresObservedInput, x.HiddenChoices, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(xs []CandidateRow) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(color=%t orientationPair=%t weak=%t Y=%t local=%t kinetic=%t residue=%t mass=%t decouple=%t rep=%t beta=%t zero=%t s6=%t observed=%t obstruction=%s)", x.Name, x.ColorTriplet, x.RealOrientationPair, x.WeakDoubletDerived, x.HyperchargeDerived, x.LocalFieldDerived, x.LorentzKineticDerived, x.PoleResidueDerived, x.MassActivationDerived, x.DecouplingDerived, x.RepresentationComplete, x.BetaPermitted, x.ZeroContributionProved, x.RequiresS6Choice, x.RequiresObservedInput, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d lqRows=%d slots=%d color=%d orient=%d weakSU2=%t YRows=%d localRows=%d kineticRows=%d massRows=%d decoupleRows=%d repRows=%d betaRows=%d zeroRows=%d s6=%d nullity=%d->%d", s.ContactRows, s.LeptoquarkRows, s.CurrentLQSlots, s.ColorDirections, s.RealOrientations, s.WeakSU2ActionDerived, s.HyperchargeRowsDerived, s.LocalFieldRowsDerived, s.LorentzKineticRowsDerived, s.MassActivationRowsDerived, s.DecouplingRowsDerived, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
