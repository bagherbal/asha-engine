// Package contactlqcharge implements Gate 135: leptoquark contact
// hypercharge source / B-L and charge-lattice obstruction theorem.
//
// Gate 134 closed beta permission for the six contact-leptoquark candidates
// because no weak SU(2)L action, hypercharge row, local field map, Lorentz
// kinetic data, mass activation, or decoupling rule was derived. Gate 135 tests
// the most tempting remaining shortcut: use the already-valid B-L charge split,
// or a finite charge lattice built from it, as the missing hypercharge source.
//
// The result is a disciplined obstruction. B-L really supplies a lepton/color
// polarization, and lepton-color off-diagonal currents have a natural B-L
// difference of 4/3. But B-L is not hypercharge, does not supply T3R or weak
// chirality, does not choose the symmetric/skew orientation as an SU(2)L
// component, and does not assign the six current slots to contact rows. Charge
// lattices made from B-L and electric-charge templates are diagnostics only
// until a contact-side local field and representation row is derived.
package contactlqcharge

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactlqbetapermission"
	"github.com/bagherbal/asha-engine/pkg/matter/charge"
)

type SourceKind string

const (
	BLDifferenceSource     SourceKind = "B-L-difference-source"
	ChargeLatticeSource    SourceKind = "finite-charge-lattice-source"
	MatterHyperchargeTable SourceKind = "matter-hypercharge-table-borrowing"
	ObservedChargeFit      SourceKind = "observed-charge-fit"
)

type ChargeSourceCandidate struct {
	Name                      string
	Kind                      SourceKind
	Canonical                 bool
	CurrentSideDiagnostic     bool
	ContactSideSemantic       bool
	UsesObservedInput         bool
	RequiresObservedInput     bool
	RequiresT3R               bool
	RequiresWeakChirality     bool
	RequiresSU2Action         bool
	RequiresLocalField        bool
	RequiresContactAssignment bool
	BMinusLValue              float64
	BMinusLDifference         float64
	HyperchargeDerived        bool
	ElectricChargeDerived     bool
	RepresentationRow         bool
	BetaPermitted             bool
	ZeroContributionProved    bool
	HiddenChoices             int
	Obstruction               string
}

type CandidateRow struct {
	Name                   string
	ColorIndex             int
	RealOrientation        string
	BMinusLDifference      float64
	BLMagnitudeDiagnostic  bool
	SignedBLDerived        bool
	T3RDerived             bool
	WeakChiralityDerived   bool
	WeakSU2Derived         bool
	HyperchargeDerived     bool
	ElectricChargeDerived  bool
	LocalFieldDerived      bool
	RepresentationComplete bool
	BetaPermitted          bool
	RequiresS6Choice       bool
	Obstruction            string
}

type Summary struct {
	ContactRows                  int
	LeptoquarkRows               int
	CurrentLQSlots               int
	BMinusLOneParticleTrace      float64
	BMinusLOneParticleTrace2     float64
	BMinusLPolarizesOnePlusThree bool
	LeptonCharge                 float64
	ColorCharge                  float64
	LeptonColorBLDifference      float64
	BLDiagnosticRows             int
	SignedBLRowsDerived          int
	T3RRowsDerived               int
	WeakChiralityRowsDerived     int
	WeakSU2RowsDerived           int
	HyperchargeRowsDerived       int
	ElectricChargeRowsDerived    int
	RepresentationCompleteRows   int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	ResidualS6Choices            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
}

type Analysis struct {
	Previous contactlqbetapermission.Analysis
	Charge   charge.Analysis

	Sources []ChargeSourceCandidate
	Rows    []CandidateRow
	Summary Summary

	ContactRows                   int
	LeptoquarkRows                int
	CurrentLQSlots                int
	BMinusLPolarizesOnePlusThree  bool
	BMinusLChargeBridgeValid      bool
	LeptonColorBLDifference       float64
	BLDifferenceDiagnostic        bool
	SignedBLRowsDerived           int
	T3RRowsDerived                int
	WeakChiralityRowsDerived      int
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
		prev, err := contactlqbetapermission.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ch, err := charge.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, ch)
	})
	return defaultValue, defaultErr
}

func Build(prev contactlqbetapermission.Analysis, ch charge.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.LeptoquarkRows != 6 || prev.CurrentLQSlots != 6 || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 135 requires Gate 134 leptoquark beta permission firewall to be closed")
	}
	if prev.HyperchargeRowDerived || prev.LocalFieldMapDerived || prev.MassActivationDerived || prev.DecouplingRuleDerived {
		return Analysis{}, fmt.Errorf("Gate 135 requires Gate 134 hypercharge/local/mass/decoupling rows to remain underived")
	}
	if !ch.ChargePolarizesOnePlusThree || len(ch.OneParticleChargeSpectrum) != 4 {
		return Analysis{}, fmt.Errorf("Gate 135 requires valid B-L 1+3 matter charge polarization")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 135 refuses hidden physical input from Gate 134")
	}

	lepton := -1.0
	color := 1.0 / 3.0
	diff := color - lepton
	if math.Abs(diff-4.0/3.0) > 1e-10 {
		return Analysis{}, fmt.Errorf("unexpected B-L lepton-color difference: %.12f", diff)
	}

	sources := buildSources(diff, prev.ResidualS6Choices)
	rows := buildRows(diff)

	blRows := count(rows, func(r CandidateRow) bool { return r.BLMagnitudeDiagnostic })
	signedRows := count(rows, func(r CandidateRow) bool { return r.SignedBLDerived })
	t3rRows := count(rows, func(r CandidateRow) bool { return r.T3RDerived })
	chirRows := count(rows, func(r CandidateRow) bool { return r.WeakChiralityDerived })
	su2Rows := count(rows, func(r CandidateRow) bool { return r.WeakSU2Derived })
	yRows := count(rows, func(r CandidateRow) bool { return r.HyperchargeDerived })
	qRows := count(rows, func(r CandidateRow) bool { return r.ElectricChargeDerived })
	repRows := count(rows, func(r CandidateRow) bool { return r.RepresentationComplete })
	betaRows := count(rows, func(r CandidateRow) bool { return r.BetaPermitted })
	zeroRows := count(rows, func(r CandidateRow) bool { return false })

	summary := Summary{
		ContactRows:                  7,
		LeptoquarkRows:               6,
		CurrentLQSlots:               6,
		BMinusLOneParticleTrace:      ch.TraceOneParticleCharge,
		BMinusLOneParticleTrace2:     ch.TraceOneParticleChargeSquared,
		BMinusLPolarizesOnePlusThree: ch.ChargePolarizesOnePlusThree,
		LeptonCharge:                 lepton,
		ColorCharge:                  color,
		LeptonColorBLDifference:      diff,
		BLDiagnosticRows:             blRows,
		SignedBLRowsDerived:          signedRows,
		T3RRowsDerived:               t3rRows,
		WeakChiralityRowsDerived:     chirRows,
		WeakSU2RowsDerived:           su2Rows,
		HyperchargeRowsDerived:       yRows,
		ElectricChargeRowsDerived:    qRows,
		RepresentationCompleteRows:   repRows,
		ContactBetaRowsAllowed:       betaRows,
		ContactZeroRowsProved:        zeroRows,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
	}

	truth := "Gate 135 audits the B-L and charge-lattice shortcut for contact leptoquark hypercharge. The matter-side B-L operator is valid and polarizes the Fock one-particle sector into a 1+3 lepton/color split. It also gives a lepton-color B-L difference of 4/3, so the leptoquark current slots have a real charge diagnostic. But B-L difference is not hypercharge: Y = T3R + (B-L)/2 still needs a contact-side T3R or weak-chirality operator, a non-abelian SU(2)L action, a local field map, and the S6 contact-row assignment. Therefore B-L and finite charge-lattice diagnostics do not open the contact beta firewall."

	return Analysis{
		Previous: prev,
		Charge:   ch,
		Sources:  sources,
		Rows:     rows,
		Summary:  summary,

		ContactRows:                   7,
		LeptoquarkRows:                6,
		CurrentLQSlots:                6,
		BMinusLPolarizesOnePlusThree:  ch.ChargePolarizesOnePlusThree,
		BMinusLChargeBridgeValid:      ch.ChargePolarizesOnePlusThree && math.Abs(ch.TraceOneParticleCharge) < 1e-10,
		LeptonColorBLDifference:       diff,
		BLDifferenceDiagnostic:        blRows == 6,
		SignedBLRowsDerived:           signedRows,
		T3RRowsDerived:                t3rRows,
		WeakChiralityRowsDerived:      chirRows,
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
			"B-L difference alone is a contact leptoquark hypercharge row",
			"Y can be borrowed from matter/Fock tables without a contact local field map",
			"electric charge can be computed before SU(2)L and T3 data exist on the contact carrier",
			"the finite charge lattice selects the S6 contact-row assignment",
			"the B-L value 4/3 is a threshold beta contribution",
			"observed electric charges may choose the missing contact rows",
		},
		RemainingUnknowns: []string{
			"contact-side T3R or weak chirality operator",
			"contact-side SU(2)L action on the six leptoquark candidates",
			"signed orientation of B-L raising/lowering on real symmetric/skew slots",
			"local field map and Lorentz kinetic row",
			"mass activation and decoupling rule",
			"canonical S6 assignment of current leptoquark slots to contact rows",
		},
		RecommendedNextGate: "Gate 136 — contact T3R/chirality source search for leptoquark hypercharge",
	}, nil
}

func buildSources(diff float64, hidden int) []ChargeSourceCandidate {
	return []ChargeSourceCandidate{
		{Name: "matter B-L lepton-color difference", Kind: BLDifferenceSource, Canonical: true, CurrentSideDiagnostic: true, ContactSideSemantic: false, BMinusLValue: diff, BMinusLDifference: diff, RequiresT3R: true, RequiresWeakChirality: true, RequiresSU2Action: true, RequiresLocalField: true, RequiresContactAssignment: true, HiddenChoices: hidden, Obstruction: "B-L supplies lepton/color polarization and a 4/3 current diagnostic, but not contact hypercharge"},
		{Name: "finite charge lattice generated by B-L and orientation sign", Kind: ChargeLatticeSource, Canonical: false, CurrentSideDiagnostic: true, ContactSideSemantic: false, BMinusLDifference: diff, RequiresT3R: true, RequiresWeakChirality: true, RequiresSU2Action: true, RequiresLocalField: true, RequiresContactAssignment: true, HiddenChoices: hidden, Obstruction: "orientation sign on real symmetric/skew slots is not a selected weak-isospin or chirality label"},
		{Name: "borrow matter hypercharge table", Kind: MatterHyperchargeTable, Canonical: false, CurrentSideDiagnostic: false, ContactSideSemantic: false, RequiresT3R: true, RequiresWeakChirality: true, RequiresSU2Action: true, RequiresLocalField: true, RequiresContactAssignment: true, HiddenChoices: hidden, Obstruction: "matter hypercharge rows live on Fock states, not on contact partial-overlap rows"},
		{Name: "observed electric-charge fit", Kind: ObservedChargeFit, Canonical: false, CurrentSideDiagnostic: false, ContactSideSemantic: false, UsesObservedInput: true, RequiresObservedInput: true, HiddenChoices: hidden, Obstruction: "observed charge fitting is forbidden as a selector"},
	}
}

func buildRows(diff float64) []CandidateRow {
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
			Name:                   name,
			ColorIndex:             n.color,
			RealOrientation:        n.orient,
			BMinusLDifference:      diff,
			BLMagnitudeDiagnostic:  true,
			SignedBLDerived:        false,
			T3RDerived:             false,
			WeakChiralityDerived:   false,
			WeakSU2Derived:         false,
			HyperchargeDerived:     false,
			ElectricChargeDerived:  false,
			LocalFieldDerived:      false,
			RepresentationComplete: false,
			BetaPermitted:          false,
			RequiresS6Choice:       true,
			Obstruction:            fmt.Sprintf("%s has the B-L magnitude |Δ(B-L)|=4/3 as a current diagnostic, but no signed contact orientation, T3R, SU(2)L, local field, or contact-row assignment", name),
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

func FormatSources(xs []ChargeSourceCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s canonical=%t currentDiag=%t contactSemantic=%t observed=%t BL=%.10f diff=%.10f Y=%t Q=%t rep=%t beta=%t zero=%t hidden=%d obstruction=%s)", x.Name, x.Kind, x.Canonical, x.CurrentSideDiagnostic, x.ContactSideSemantic, x.UsesObservedInput, x.BMinusLValue, x.BMinusLDifference, x.HyperchargeDerived, x.ElectricChargeDerived, x.RepresentationRow, x.BetaPermitted, x.ZeroContributionProved, x.HiddenChoices, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(xs []CandidateRow) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(color=%d orient=%s dBL=%.10f BDiag=%t signedBL=%t T3R=%t chi=%t SU2=%t Y=%t Q=%t local=%t rep=%t beta=%t s6=%t obstruction=%s)", x.Name, x.ColorIndex, x.RealOrientation, x.BMinusLDifference, x.BLMagnitudeDiagnostic, x.SignedBLDerived, x.T3RDerived, x.WeakChiralityDerived, x.WeakSU2Derived, x.HyperchargeDerived, x.ElectricChargeDerived, x.LocalFieldDerived, x.RepresentationComplete, x.BetaPermitted, x.RequiresS6Choice, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d lq=%d slots=%d TrBL=%.3e TrBL2=%.10f BL1+3=%t lepton=%.10f color=%.10f dBL=%.10f BDiag=%d signedBL=%d T3R=%d chi=%d SU2=%d Y=%d Q=%d rep=%d beta=%d zero=%d s6=%d nullity=%d->%d", s.ContactRows, s.LeptoquarkRows, s.CurrentLQSlots, s.BMinusLOneParticleTrace, s.BMinusLOneParticleTrace2, s.BMinusLPolarizesOnePlusThree, s.LeptonCharge, s.ColorCharge, s.LeptonColorBLDifference, s.BLDiagnosticRows, s.SignedBLRowsDerived, s.T3RRowsDerived, s.WeakChiralityRowsDerived, s.WeakSU2RowsDerived, s.HyperchargeRowsDerived, s.ElectricChargeRowsDerived, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
