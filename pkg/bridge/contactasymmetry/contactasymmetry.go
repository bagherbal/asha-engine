// Package contactasymmetry implements Gate 143: contact charge-conjugation
// breaking source / asymmetry selector search.
//
// Gate 140 found a canonical 3|4 largest-gap split in the seven contact
// partial-overlap rows. Gate 141 showed that the split has two possible sign
// orientations. Gate 142 proved that charge conjugation exchanges those two
// orientations and therefore gives a Z2 degeneracy rather than a branch
// selector. Gate 143 asks the sharper question: does the finite system contain
// any charge-conjugation-breaking source or asymmetry functional that selects
// one branch without importing observed constants?
//
// The answer remains no. The contact spectrum contains real finite asymmetry
// diagnostics: a 3|4 cardinal imbalance and different high/low spectral
// moments. But those diagnostics are unordered or C-even until a signed source,
// pullback, current, local field map, or representation row exists. The gate
// therefore records the asymmetry as useful contact data while keeping the beta
// firewall closed.
package contactasymmetry

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactsignsource"
)

type AsymmetrySourceKind string

const (
	SourceCardinalityImbalance AsymmetrySourceKind = "contact-cardinality-imbalance"
	SourceSpectralMoment       AsymmetrySourceKind = "contact-spectral-moment"
	SourceCBreakingFunctional  AsymmetrySourceKind = "C-breaking-functional"
	SourceBLOrT3RPullback      AsymmetrySourceKind = "B-L-or-T3R-pullback"
	SourceHyperchargeRow       AsymmetrySourceKind = "hypercharge-row"
	SourceLocalFieldCurrent    AsymmetrySourceKind = "local-field-current"
	SourceObservedFit          AsymmetrySourceKind = "observed-constant-fit"
)

type AsymmetrySourceAudit struct {
	Name                  string
	Kind                  AsymmetrySourceKind
	Available             bool
	Finite                bool
	ContactSide           bool
	MatterSide            bool
	CInvariant            bool
	CBreaking             bool
	SelectsOrientation    bool
	RequiresPullback      bool
	RequiresLocalFieldMap bool
	RequiresObservedInput bool
	BranchesRemaining     int
	Verdict               string
}

type SplitAsymmetryAudit struct {
	HighRows              int
	LowRows               int
	CardinalityImbalance  int
	HighMean              float64
	LowMean               float64
	MomentSeparation      float64
	AsymmetryDiagnostics  int
	CInvariantDiagnostics int
	CBreakingDiagnostics  int
	SelectedOrientations  int
	Verdict               string
}

type Summary struct {
	ContactRows                  int
	LargestGapHighRows           int
	LargestGapLowRows            int
	OrientationCandidates        int
	AsymmetrySourcesAudited      int
	AsymmetrySourcesAvailable    int
	CBreakingSources             int
	SourcesSelectingOrientation  int
	CardinalityImbalance         int
	AsymmetryDiagnostics         int
	CInvariantDiagnostics        int
	CBreakingDiagnostics         int
	Z2OrientationDegeneracy      bool
	ChargeConjugationBroken      bool
	CoddContactFunctionals       int
	T3RPullbackRowsDerived       int
	ChiralityPullbackRowsDerived int
	BMinusLPullbackRowsDerived   int
	SU2LPullbackRowsDerived      int
	HyperchargeRowsDerived       int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	ResidualS6Choices            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
}

type Analysis struct {
	Previous contactsignsource.Analysis

	SpectrumDescending []float64
	SplitPattern       string
	SplitAudit         SplitAsymmetryAudit
	SourceAudits       []AsymmetrySourceAudit
	Summary            Summary

	ContactRows                  int
	LargestGapHighRows           int
	LargestGapLowRows            int
	OrientationCandidates        int
	AsymmetrySourcesAudited      int
	AsymmetrySourcesAvailable    int
	CBreakingSources             int
	SourcesSelectingOrientation  int
	CardinalityImbalance         int
	AsymmetryDiagnostics         int
	CInvariantDiagnostics        int
	CBreakingDiagnostics         int
	Z2OrientationDegeneracy      bool
	ChargeConjugationBroken      bool
	CoddContactFunctionals       int
	T3RPullbackRowsDerived       int
	ChiralityPullbackRowsDerived int
	BMinusLPullbackRowsDerived   int
	SU2LPullbackRowsDerived      int
	HyperchargeRowsDerived       int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	BetaPermissionFirewallClosed bool
	ThresholdCorrectedBeta       bool
	FullBetaMatchingTensor       bool

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
		prev, err := contactsignsource.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactsignsource.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.LargestGapHighRows != 3 || prev.LargestGapLowRows != 4 || prev.OrientationCandidates != 2 {
		return Analysis{}, fmt.Errorf("Gate 143 requires Gate 142 closed-firewall 3|4 split with two charge-conjugation-related orientations")
	}
	if !prev.Z2OrientationDegeneracy || prev.ChargeConjugationSelectedBranches != 0 || prev.SourcesSelectingOrientation != 0 {
		return Analysis{}, fmt.Errorf("Gate 143 requires Gate 142 Z2 orientation degeneracy with no selected branch")
	}
	if prev.T3RPullbackRowsDerived != 0 || prev.ChiralityPullbackRowsDerived != 0 || prev.BMinusLPullbackRowsDerived != 0 || prev.SU2LPullbackRowsDerived != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 143 requires no contact charge pullbacks, hypercharge rows, or beta rows before asymmetry selection")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 143 refuses hidden observed physical input")
	}

	highRows := prev.LargestGapHighRows
	lowRows := prev.LargestGapLowRows
	highMean := mean(prev.SpectrumDescending[:highRows])
	lowMean := mean(prev.SpectrumDescending[highRows:])
	momentSeparation := highMean - lowMean
	cardinalityImbalance := abs(highRows - lowRows)

	split := SplitAsymmetryAudit{
		HighRows:              highRows,
		LowRows:               lowRows,
		CardinalityImbalance:  cardinalityImbalance,
		HighMean:              highMean,
		LowMean:               lowMean,
		MomentSeparation:      momentSeparation,
		AsymmetryDiagnostics:  2,
		CInvariantDiagnostics: 2,
		CBreakingDiagnostics:  0,
		SelectedOrientations:  0,
		Verdict:               "the contact split has real finite asymmetry diagnostics, but they are unordered/C-even until a C-breaking source or signed pullback exists",
	}

	audits := []AsymmetrySourceAudit{
		{
			Name:               "3|4 contact cardinality imbalance",
			Kind:               SourceCardinalityImbalance,
			Available:          true,
			Finite:             true,
			ContactSide:        true,
			CInvariant:         true,
			CBreaking:          false,
			SelectsOrientation: false,
			BranchesRemaining:  2,
			Verdict:            "detects that the largest-gap split is unbalanced, but not which side is +T3R",
		},
		{
			Name:               "high/low spectral moment separation",
			Kind:               SourceSpectralMoment,
			Available:          true,
			Finite:             true,
			ContactSide:        true,
			CInvariant:         true,
			CBreaking:          false,
			SelectsOrientation: false,
			BranchesRemaining:  2,
			Verdict:            "orders high-overlap versus low-overlap rows, but sign orientation remains two-valued",
		},
		{
			Name:               "contact C-breaking finite source",
			Kind:               SourceCBreakingFunctional,
			Available:          false,
			Finite:             true,
			ContactSide:        true,
			CInvariant:         false,
			CBreaking:          false,
			SelectsOrientation: false,
			BranchesRemaining:  2,
			Verdict:            "no C-odd contact functional or signed source has been derived",
		},
		{
			Name:               "B-L / T3R / chirality pullback",
			Kind:               SourceBLOrT3RPullback,
			Available:          false,
			Finite:             true,
			MatterSide:         true,
			SelectsOrientation: false,
			RequiresPullback:   true,
			BranchesRemaining:  2,
			Verdict:            "matter-side charge diagnostics still lack a Fock-contact pullback",
		},
		{
			Name:               "hypercharge or electric-charge consistency",
			Kind:               SourceHyperchargeRow,
			Available:          false,
			Finite:             true,
			SelectsOrientation: false,
			RequiresPullback:   true,
			BranchesRemaining:  2,
			Verdict:            "cannot break the sign symmetry before contact T3R and B-L rows exist",
		},
		{
			Name:                  "local field current / pole-residue asymmetry",
			Kind:                  SourceLocalFieldCurrent,
			Available:             false,
			Finite:                true,
			ContactSide:           true,
			SelectsOrientation:    false,
			RequiresLocalFieldMap: true,
			BranchesRemaining:     2,
			Verdict:               "no local contact current, Lorentz kinetic residue, mass activation, or decoupling asymmetry is available",
		},
		{
			Name:                  "observed constants / measured chirality selector",
			Kind:                  SourceObservedFit,
			Available:             false,
			Finite:                false,
			SelectsOrientation:    false,
			RequiresObservedInput: true,
			BranchesRemaining:     2,
			Verdict:               "forbidden: observed constants cannot be used to break the finite charge-conjugation degeneracy",
		},
	}

	available := count(audits, func(a AsymmetrySourceAudit) bool { return a.Available })
	breaking := count(audits, func(a AsymmetrySourceAudit) bool { return a.CBreaking && a.Available })
	selecting := count(audits, func(a AsymmetrySourceAudit) bool { return a.SelectsOrientation })

	summary := Summary{
		ContactRows:                  prev.ContactRows,
		LargestGapHighRows:           highRows,
		LargestGapLowRows:            lowRows,
		OrientationCandidates:        prev.OrientationCandidates,
		AsymmetrySourcesAudited:      len(audits),
		AsymmetrySourcesAvailable:    available,
		CBreakingSources:             breaking,
		SourcesSelectingOrientation:  selecting,
		CardinalityImbalance:         cardinalityImbalance,
		AsymmetryDiagnostics:         split.AsymmetryDiagnostics,
		CInvariantDiagnostics:        split.CInvariantDiagnostics,
		CBreakingDiagnostics:         split.CBreakingDiagnostics,
		Z2OrientationDegeneracy:      prev.Z2OrientationDegeneracy,
		ChargeConjugationBroken:      false,
		CoddContactFunctionals:       0,
		T3RPullbackRowsDerived:       0,
		ChiralityPullbackRowsDerived: 0,
		BMinusLPullbackRowsDerived:   0,
		SU2LPullbackRowsDerived:      0,
		HyperchargeRowsDerived:       0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.ContactRows,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
	}

	truth := "Gate 143 searches for a finite source that breaks the charge-conjugation exchange between the two signs of the canonical 3|4 contact split. The contact spectrum contains real asymmetry diagnostics: a 3|4 cardinal imbalance and separated high/low spectral moments. But both diagnostics are C-even unless a signed source, contact charge pullback, local current, or representation row is derived. No C-breaking source is currently present, so the Z2 orientation degeneracy and contact beta firewall remain closed."

	return Analysis{
		Previous:                     prev,
		SpectrumDescending:           append([]float64(nil), prev.SpectrumDescending...),
		SplitPattern:                 prev.SplitPattern,
		SplitAudit:                   split,
		SourceAudits:                 audits,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		LargestGapHighRows:           highRows,
		LargestGapLowRows:            lowRows,
		OrientationCandidates:        prev.OrientationCandidates,
		AsymmetrySourcesAudited:      len(audits),
		AsymmetrySourcesAvailable:    available,
		CBreakingSources:             breaking,
		SourcesSelectingOrientation:  selecting,
		CardinalityImbalance:         cardinalityImbalance,
		AsymmetryDiagnostics:         split.AsymmetryDiagnostics,
		CInvariantDiagnostics:        split.CInvariantDiagnostics,
		CBreakingDiagnostics:         split.CBreakingDiagnostics,
		Z2OrientationDegeneracy:      prev.Z2OrientationDegeneracy,
		ChargeConjugationBroken:      false,
		CoddContactFunctionals:       0,
		T3RPullbackRowsDerived:       0,
		ChiralityPullbackRowsDerived: 0,
		BMinusLPullbackRowsDerived:   0,
		SU2LPullbackRowsDerived:      0,
		HyperchargeRowsDerived:       0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.ContactRows,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		BetaPermissionFirewallClosed: true,
		ThresholdCorrectedBeta:       false,
		FullBetaMatchingTensor:       false,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
		HiddenObservedInputUsed:      false,
		PhysicalWeakAngleDerived:     false,
		FineStructureDerived:         false,
		PhysicalMassesDerived:        false,
		PhysicalScaleDerived:         false,
		TruthStatement:               truth,
		RejectedClaims: []string{
			"3|4 cardinal imbalance selects the +T3R side",
			"spectral moment separation is a C-breaking source",
			"charge conjugation is broken by the current finite contact data",
			"B-L, chirality, or T3R can orient contact rows without pullback",
			"hypercharge consistency can be used before contact T3R and B-L exist",
			"contact asymmetry diagnostics permit threshold beta rows",
		},
		RemainingUnknowns: []string{
			"contact C-odd source functional, if any",
			"Fock-contact pullback for B-L, chirality, T3R, and SU(2)L",
			"contact local field map, pole residues, mass activation, and decoupling",
			"contact hypercharge rows and threshold beta tensor",
		},
		RecommendedNextGate: "Gate 144 — contact C-odd source functional / finite signed-current construction attempt",
	}, nil
}

func FormatAsymmetrySources(items []AsymmetrySourceAudit) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s available=%t C-breaking=%t selects=%t (%s)", item.Name, item.Available, item.CBreaking, item.SelectsOrientation, item.Verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatSplitAudit(a SplitAsymmetryAudit) string {
	return fmt.Sprintf("split=%d|%d imbalance=%d highMean=%.10f lowMean=%.10f momentSep=%.10f diagnostics=%d C-even=%d C-breaking=%d selected=%d (%s)", a.HighRows, a.LowRows, a.CardinalityImbalance, a.HighMean, a.LowMean, a.MomentSeparation, a.AsymmetryDiagnostics, a.CInvariantDiagnostics, a.CBreakingDiagnostics, a.SelectedOrientations, a.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d split=%d|%d sources=%d available=%d C-breaking=%d selected=%d C-broken=%t C-odd=%d T3R=%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.LargestGapHighRows, s.LargestGapLowRows, s.AsymmetrySourcesAudited, s.AsymmetrySourcesAvailable, s.CBreakingSources, s.SourcesSelectingOrientation, s.ChargeConjugationBroken, s.CoddContactFunctionals, s.T3RPullbackRowsDerived, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(items []string) string { return strings.Join(items, "; ") }

func count[T any](items []T, pred func(T) bool) int {
	n := 0
	for _, item := range items {
		if pred(item) {
			n++
		}
	}
	return n
}

func mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
