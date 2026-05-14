// Package contactlqt3r implements Gate 136: contact T3R/chirality source
// search for leptoquark hypercharge.
//
// Gate 135 showed that B-L gives a genuine lepton-color diagnostic for the six
// current-side leptoquark slots, but not a contact hypercharge row. Gate 136
// tests the next tempting shortcut: pull the already-known matter-side T3R /
// chiral-orientation candidates onto the contact leptoquark carrier.
//
// The result is a pullback obstruction. The matter/Fock sector has a temporal
// T3R candidate family, and the hypercharge audit can prefer a right-singlet /
// conjugate branch. But those operators live on the matter/Fock tensor domain.
// The seven contact rows still have no local field map, no Fock-to-contact
// pullback, no signed B-L orientation, no SU(2)L action, and no S6 row
// assignment. Therefore T3R/chirality remain matter-side diagnostics only and
// cannot open the contact beta firewall.
package contactlqt3r

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactlqcharge"
	"github.com/bagherbal/asha-engine/pkg/matter/hyperaudit"
	"github.com/bagherbal/asha-engine/pkg/matter/t3r"
)

type SourceKind string

const (
	MatterTemporalT3RSource       SourceKind = "matter-temporal-T3R-source"
	MatterChiralRestrictionSource SourceKind = "matter-chiral-restriction-source"
	HyperauditBranchSource        SourceKind = "matter-hyperaudit-branch-source"
	OrientationChiralityGuess     SourceKind = "symmetric-skew-orientation-chirality-guess"
	ContactSpectralParityGuess    SourceKind = "contact-spectral-parity-guess"
	BorrowedMatterOperator        SourceKind = "borrowed-matter-operator"
)

type SourceCandidate struct {
	Name                      string
	Kind                      SourceKind
	MatterSide                bool
	CurrentSide               bool
	ContactSide               bool
	Canonical                 bool
	Diagnostic                bool
	SelectedInMatterAudit     bool
	ContactPullbackDerived    bool
	T3RDerivedOnContact       bool
	ChiralityDerivedOnContact bool
	HyperchargeDerived        bool
	RequiresFockContactMap    bool
	RequiresSignedBL          bool
	RequiresSU2Action         bool
	RequiresS6Assignment      bool
	UsesObservedInput         bool
	HypotheticalYValues       []float64
	Obstruction               string
}

type CandidateRow struct {
	Name                     string
	ColorIndex               int
	RealOrientation          string
	BMinusLDifference        float64
	HalfBMinusLDifference    float64
	MatterT3RCandidateValues []float64
	HypotheticalYValues      []float64
	MatterT3RDiagnostic      bool
	ContactT3RDerived        bool
	ContactChiralityDerived  bool
	SignedBLDerived          bool
	WeakSU2Derived           bool
	HyperchargeDerived       bool
	ElectricChargeDerived    bool
	LocalFieldDerived        bool
	RepresentationComplete   bool
	BetaPermitted            bool
	RequiresS6Choice         bool
	Obstruction              string
}

type Summary struct {
	ContactRows                   int
	LeptoquarkRows                int
	MatterT3ROperatorFound        bool
	MatterChiralRestrictions      int
	MatterMirrorAmbiguity         bool
	MatterOrientationSelected     bool
	HyperauditPreferredBranchName string
	RightSingletAuditSelected     bool
	FullSMTableDerived            bool
	BMinusLDifference             float64
	HalfBMinusLDifference         float64
	HypotheticalYValueCount       int
	ContactPullbacksDerived       int
	ContactT3RRowsDerived         int
	ContactChiralityRowsDerived   int
	SignedBLRowsDerived           int
	WeakSU2RowsDerived            int
	HyperchargeRowsDerived        int
	ElectricChargeRowsDerived     int
	RepresentationCompleteRows    int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	ResidualS6Choices             int
	ResidualNullityBefore         int
	ResidualNullityAfter          int
}

type Analysis struct {
	Previous   contactlqcharge.Analysis
	T3R        t3r.Analysis
	Hyperaudit hyperaudit.Analysis

	Sources []SourceCandidate
	Rows    []CandidateRow
	Summary Summary

	ContactRows                   int
	LeptoquarkRows                int
	MatterT3ROperatorFound        bool
	MatterChiralRestricted        bool
	MatterMirrorAmbiguity         bool
	MatterHyperauditSelectsBranch bool
	MatterFullSMTableDerived      bool
	MatterT3RDiagnosticAvailable  bool
	ContactPullbackRowsDerived    int
	ContactT3RRowsDerived         int
	ContactChiralityRowsDerived   int
	SignedBLRowsDerived           int
	WeakSU2RowsDerived            int
	HyperchargeRowsDerived        int
	ElectricChargeRowsDerived     int
	RepresentationCompleteRows    int
	RepresentationOpenRows        int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	FullBetaMatchingTensorDerived bool
	ThresholdCorrectedBetaDerived bool
	BetaPermissionFirewallClosed  bool

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
		prev, err := contactlqcharge.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		t, err := t3r.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		h, err := hyperaudit.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, t, h)
	})
	return defaultValue, defaultErr
}

func Build(prev contactlqcharge.Analysis, t t3r.Analysis, h hyperaudit.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.LeptoquarkRows != 6 || prev.HyperchargeRowsDerived != 0 || prev.T3RRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 136 requires Gate 135 charge/T3R firewall to be closed")
	}
	if !t.MatterSideOperatorFound || !t.HyperchargeCandidateConstructed {
		return Analysis{}, fmt.Errorf("Gate 136 requires matter-side T3R candidate family")
	}
	if !t.ChiralRestrictedBridgeAvailable || !t.MirrorAmbiguity || t.PhysicalOrientationSelected {
		return Analysis{}, fmt.Errorf("Gate 136 expects matter T3R chiral restriction to remain orientation-open")
	}
	if !h.RightSingletConjugateTableDerived || h.PreferredBranchName == "" {
		return Analysis{}, fmt.Errorf("Gate 136 expects hyperaudit right-singlet branch diagnostic")
	}
	if h.FullStandardModelTableDerived {
		return Analysis{}, fmt.Errorf("Gate 136 requires full Standard Model table to remain underived")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 136 refuses hidden physical input from Gate 135")
	}

	diff := prev.LeptonColorBLDifference
	half := diff / 2
	if math.Abs(half-2.0/3.0) > 1e-10 {
		return Analysis{}, fmt.Errorf("unexpected half B-L difference: %.12f", half)
	}
	t3Candidates := []float64{-0.5, 0.5}
	yCandidates := uniqueSorted([]float64{half - 0.5, half + 0.5, -half - 0.5, -half + 0.5})

	sources := buildSources(t, h, yCandidates, prev.ResidualS6Choices)
	rows := buildRows(diff, half, t3Candidates, yCandidates)

	pullbacks := count(rows, func(r CandidateRow) bool { return false })
	t3rRows := count(rows, func(r CandidateRow) bool { return r.ContactT3RDerived })
	chirRows := count(rows, func(r CandidateRow) bool { return r.ContactChiralityDerived })
	signedRows := count(rows, func(r CandidateRow) bool { return r.SignedBLDerived })
	su2Rows := count(rows, func(r CandidateRow) bool { return r.WeakSU2Derived })
	yRows := count(rows, func(r CandidateRow) bool { return r.HyperchargeDerived })
	qRows := count(rows, func(r CandidateRow) bool { return r.ElectricChargeDerived })
	repRows := count(rows, func(r CandidateRow) bool { return r.RepresentationComplete })
	betaRows := count(rows, func(r CandidateRow) bool { return r.BetaPermitted })
	zeroRows := 0

	summary := Summary{
		ContactRows:                   7,
		LeptoquarkRows:                6,
		MatterT3ROperatorFound:        t.MatterSideOperatorFound,
		MatterChiralRestrictions:      2,
		MatterMirrorAmbiguity:         t.MirrorAmbiguity,
		MatterOrientationSelected:     t.PhysicalOrientationSelected,
		HyperauditPreferredBranchName: h.PreferredBranchName,
		RightSingletAuditSelected:     h.RightSingletConjugateTableDerived,
		FullSMTableDerived:            h.FullStandardModelTableDerived,
		BMinusLDifference:             diff,
		HalfBMinusLDifference:         half,
		HypotheticalYValueCount:       len(yCandidates),
		ContactPullbacksDerived:       pullbacks,
		ContactT3RRowsDerived:         t3rRows,
		ContactChiralityRowsDerived:   chirRows,
		SignedBLRowsDerived:           signedRows,
		WeakSU2RowsDerived:            su2Rows,
		HyperchargeRowsDerived:        yRows,
		ElectricChargeRowsDerived:     qRows,
		RepresentationCompleteRows:    repRows,
		ContactBetaRowsAllowed:        betaRows,
		ContactZeroRowsProved:         zeroRows,
		ResidualS6Choices:             prev.ResidualS6Choices,
		ResidualNullityBefore:         prev.ResidualNullityAfter,
		ResidualNullityAfter:          prev.ResidualNullityAfter,
	}

	truth := "Gate 136 audits the contact T3R/chirality shortcut for leptoquark hypercharge. The matter/Fock sector already has a temporal T3R candidate family, chiral restrictions, and a hyperaudit branch that matches a right-singlet/conjugate table. But this is still matter-side structure. No Fock-to-contact pullback, contact chirality operator, signed B-L orientation, non-abelian SU(2)L action, local field map, or S6 row assignment is derived. Therefore T3R/chirality remain diagnostics only and do not open the contact beta firewall."

	return Analysis{
		Previous:   prev,
		T3R:        t,
		Hyperaudit: h,
		Sources:    sources,
		Rows:       rows,
		Summary:    summary,

		ContactRows:                   7,
		LeptoquarkRows:                6,
		MatterT3ROperatorFound:        t.MatterSideOperatorFound,
		MatterChiralRestricted:        t.ChiralRestrictedBridgeAvailable,
		MatterMirrorAmbiguity:         t.MirrorAmbiguity,
		MatterHyperauditSelectsBranch: h.RightSingletConjugateTableDerived && h.PreferredBranchName != "",
		MatterFullSMTableDerived:      h.FullStandardModelTableDerived,
		MatterT3RDiagnosticAvailable:  true,
		ContactPullbackRowsDerived:    pullbacks,
		ContactT3RRowsDerived:         t3rRows,
		ContactChiralityRowsDerived:   chirRows,
		SignedBLRowsDerived:           signedRows,
		WeakSU2RowsDerived:            su2Rows,
		HyperchargeRowsDerived:        yRows,
		ElectricChargeRowsDerived:     qRows,
		RepresentationCompleteRows:    repRows,
		RepresentationOpenRows:        7,
		ContactBetaRowsAllowed:        betaRows,
		ContactZeroRowsProved:         zeroRows,
		FullBetaMatchingTensorDerived: false,
		ThresholdCorrectedBetaDerived: false,
		BetaPermissionFirewallClosed:  repRows == 0 && betaRows == 0 && zeroRows == 0,

		ResidualS6Choices:        prev.ResidualS6Choices,
		ResidualNullityBefore:    prev.ResidualNullityAfter,
		ResidualNullityAfter:     prev.ResidualNullityAfter,
		HiddenObservedInputUsed:  false,
		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"matter-side T3R can be borrowed as contact T3R without a pullback map",
			"the hyperaudit preferred branch assigns chirality to contact partial-overlap rows",
			"symmetric/skew real orientation is weak chirality",
			"hypothetical Y values from ±T3R and ±B-L are contact hypercharge rows",
			"a contact leptoquark beta row may be counted before local-field and decoupling data exist",
			"observed charges may select the missing contact orientation",
		},
		RemainingUnknowns: []string{
			"Fock-to-contact pullback map for T3R/chirality",
			"contact-side chirality or right-isospin operator",
			"signed B-L orientation on the contact leptoquark rows",
			"non-abelian SU(2)L action on the contact carrier",
			"local field map, Lorentz kinetic row, mass activation, and decoupling",
			"canonical S6 assignment of current leptoquark slots to contact rows",
		},
		RecommendedNextGate: "Gate 137 — contact T3R pullback obstruction / Fock-to-contact intertwiner search",
	}, nil
}

func buildSources(t t3r.Analysis, h hyperaudit.Analysis, y []float64, hidden int) []SourceCandidate {
	return []SourceCandidate{
		{Name: "matter temporal T0 = 1/2 - N0", Kind: MatterTemporalT3RSource, MatterSide: true, Canonical: true, Diagnostic: true, ContactSide: false, ContactPullbackDerived: false, T3RDerivedOnContact: false, ChiralityDerivedOnContact: false, HyperchargeDerived: false, RequiresFockContactMap: true, RequiresSignedBL: true, RequiresSU2Action: true, RequiresS6Assignment: true, HypotheticalYValues: y, Obstruction: fmt.Sprintf("matter T3R exists with trace %.3e and Tr(T0^2)=%.10f, but no contact pullback is derived", t.TemporalTrace, t.TemporalTraceSquared)},
		{Name: "even/odd chiral restrictions of temporal T3R", Kind: MatterChiralRestrictionSource, MatterSide: true, Canonical: false, Diagnostic: true, ContactSide: false, ContactPullbackDerived: false, T3RDerivedOnContact: false, ChiralityDerivedOnContact: false, HyperchargeDerived: false, RequiresFockContactMap: true, RequiresSignedBL: true, RequiresSU2Action: true, RequiresS6Assignment: true, HypotheticalYValues: y, Obstruction: "chiral restrictions are algebraically available on H_Fock, but not selected as contact chirality"},
		{Name: "hyperaudit preferred right-singlet/conjugate branch", Kind: HyperauditBranchSource, MatterSide: true, Canonical: false, Diagnostic: true, SelectedInMatterAudit: h.RightSingletConjugateTableDerived, ContactSide: false, ContactPullbackDerived: false, T3RDerivedOnContact: false, ChiralityDerivedOnContact: false, HyperchargeDerived: false, RequiresFockContactMap: true, RequiresSignedBL: true, RequiresSU2Action: true, RequiresS6Assignment: true, HypotheticalYValues: y, Obstruction: "the preferred branch is a matter-table audit, not a contact-row semantic map"},
		{Name: "identify symmetric/skew orientation with chirality", Kind: OrientationChiralityGuess, CurrentSide: true, Canonical: false, Diagnostic: false, ContactSide: false, RequiresFockContactMap: true, RequiresSignedBL: true, RequiresSU2Action: true, RequiresS6Assignment: true, Obstruction: "symmetric/skew labels are real current orientations, not a derived right-isospin or chirality operator"},
		{Name: "contact spectral parity / sorted overlap labels", Kind: ContactSpectralParityGuess, ContactSide: true, Canonical: false, Diagnostic: true, ContactPullbackDerived: false, T3RDerivedOnContact: false, ChiralityDerivedOnContact: false, HyperchargeDerived: false, RequiresFockContactMap: true, RequiresSignedBL: true, RequiresSU2Action: true, RequiresS6Assignment: true, Obstruction: "spectral ordering distinguishes contact rows but does not define T3R/chirality semantics"},
		{Name: "borrow matter hypercharge operator", Kind: BorrowedMatterOperator, MatterSide: true, Canonical: false, Diagnostic: false, ContactSide: false, ContactPullbackDerived: false, T3RDerivedOnContact: false, ChiralityDerivedOnContact: false, HyperchargeDerived: false, RequiresFockContactMap: true, RequiresSignedBL: true, RequiresSU2Action: true, RequiresS6Assignment: true, UsesObservedInput: false, HypotheticalYValues: y, Obstruction: fmt.Sprintf("borrowing still leaves %d contact-row assignment choices", hidden)},
	}
}

func buildRows(diff, half float64, t3Vals, yVals []float64) []CandidateRow {
	names := []struct {
		color  int
		orient string
	}{
		{1, "symmetric"}, {1, "skew"},
		{2, "symmetric"}, {2, "skew"},
		{3, "symmetric"}, {3, "skew"},
	}
	out := make([]CandidateRow, 0, len(names))
	for _, n := range names {
		name := fmt.Sprintf("LQ_color%d_%s", n.color, n.orient)
		out = append(out, CandidateRow{
			Name:                     name,
			ColorIndex:               n.color,
			RealOrientation:          n.orient,
			BMinusLDifference:        diff,
			HalfBMinusLDifference:    half,
			MatterT3RCandidateValues: append([]float64(nil), t3Vals...),
			HypotheticalYValues:      append([]float64(nil), yVals...),
			MatterT3RDiagnostic:      true,
			ContactT3RDerived:        false,
			ContactChiralityDerived:  false,
			SignedBLDerived:          false,
			WeakSU2Derived:           false,
			HyperchargeDerived:       false,
			ElectricChargeDerived:    false,
			LocalFieldDerived:        false,
			RepresentationComplete:   false,
			BetaPermitted:            false,
			RequiresS6Choice:         true,
			Obstruction:              fmt.Sprintf("%s has matter-side T3R candidates ±1/2 and |Δ(B-L)|/2=2/3, giving hypothetical Y values only after unsigned/sign and contact-pullback choices", name),
		})
	}
	return out
}

func uniqueSorted(vals []float64) []float64 {
	out := make([]float64, 0, len(vals))
	for _, v := range vals {
		found := false
		for _, w := range out {
			if math.Abs(v-w) < 1e-12 {
				found = true
				break
			}
		}
		if !found {
			out = append(out, v)
		}
	}
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j] < out[i] {
				out[i], out[j] = out[j], out[i]
			}
		}
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

func FormatSources(xs []SourceCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s matter=%t current=%t contact=%t canonical=%t diag=%t matterSelected=%t pullback=%t T3R=%t chi=%t Y=%t observed=%t hypY=%s s6=%t obstruction=%s)", x.Name, x.Kind, x.MatterSide, x.CurrentSide, x.ContactSide, x.Canonical, x.Diagnostic, x.SelectedInMatterAudit, x.ContactPullbackDerived, x.T3RDerivedOnContact, x.ChiralityDerivedOnContact, x.HyperchargeDerived, x.UsesObservedInput, FormatFloats(x.HypotheticalYValues), x.RequiresS6Assignment, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(xs []CandidateRow) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(color=%d orient=%s dBL=%.10f half=%.10f T3=%s hypY=%s matterDiag=%t contactT3R=%t chi=%t signedBL=%t SU2=%t Y=%t Q=%t local=%t rep=%t beta=%t s6=%t obstruction=%s)", x.Name, x.ColorIndex, x.RealOrientation, x.BMinusLDifference, x.HalfBMinusLDifference, FormatFloats(x.MatterT3RCandidateValues), FormatFloats(x.HypotheticalYValues), x.MatterT3RDiagnostic, x.ContactT3RDerived, x.ContactChiralityDerived, x.SignedBLDerived, x.WeakSU2Derived, x.HyperchargeDerived, x.ElectricChargeDerived, x.LocalFieldDerived, x.RepresentationComplete, x.BetaPermitted, x.RequiresS6Choice, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d lq=%d matterT3R=%t chiralRestrictions=%d mirror=%t matterOrientation=%t hyperaudit=%q rightSinglet=%t fullSM=%t dBL=%.10f half=%.10f hypY=%d pullbacks=%d contactT3R=%d chi=%d signedBL=%d SU2=%d Y=%d Q=%d rep=%d beta=%d zero=%d s6=%d nullity=%d->%d", s.ContactRows, s.LeptoquarkRows, s.MatterT3ROperatorFound, s.MatterChiralRestrictions, s.MatterMirrorAmbiguity, s.MatterOrientationSelected, s.HyperauditPreferredBranchName, s.RightSingletAuditSelected, s.FullSMTableDerived, s.BMinusLDifference, s.HalfBMinusLDifference, s.HypotheticalYValueCount, s.ContactPullbacksDerived, s.ContactT3RRowsDerived, s.ContactChiralityRowsDerived, s.SignedBLRowsDerived, s.WeakSU2RowsDerived, s.HyperchargeRowsDerived, s.ElectricChargeRowsDerived, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func FormatFloats(xs []float64) string {
	if len(xs) == 0 {
		return "[]"
	}
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.10f", x)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
