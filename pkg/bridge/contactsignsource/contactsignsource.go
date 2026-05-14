// Package contactsignsource implements Gate 142: contact sign-orientation
// source / charge-conjugation symmetry obstruction theorem.
//
// Gate 140 found a canonical 3|4 largest-gap split in the seven contact
// partial-overlap rows. Gate 141 showed that the split has two compatible sign
// orientations and that neither is selected by the finite data. Gate 142 asks
// whether a stronger source exists: charge-conjugation symmetry, B-L, chirality,
// hypercharge consistency, source/current pairing, or spectral monotonicity.
//
// The answer remains no. Charge conjugation exchanges the two sign orientations,
// so it proves a Z2 degeneracy rather than choosing a branch. B-L, chirality,
// and T3R are still matter-side diagnostics without a Fock-contact pullback.
// Spectral monotonicity and uniform source/current pairings are contact-side
// diagnostics, not physical sign sources. The contact beta firewall therefore
// stays closed.
package contactsignsource

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactorientation"
)

type OrientationSourceKind string

const (
	SourceSpectralMonotonicity OrientationSourceKind = "spectral-monotonicity"
	SourceChargeConjugation    OrientationSourceKind = "charge-conjugation"
	SourceBMinusL              OrientationSourceKind = "B-L-pullback"
	SourceChirality            OrientationSourceKind = "chirality-pullback"
	SourceHypercharge          OrientationSourceKind = "hypercharge-consistency"
	SourceDualPairing          OrientationSourceKind = "source-current-dual-pairing"
	SourceObservedConstants    OrientationSourceKind = "observed-constants"
)

type OrientationSourceAudit struct {
	Name                  string
	Kind                  OrientationSourceKind
	Available             bool
	Finite                bool
	ContactSide           bool
	MatterSide            bool
	SelectsOrientation    bool
	ExchangesOrientations bool
	RequiresPullback      bool
	RequiresObservedInput bool
	BranchesRemaining     int
	Verdict               string
}

type ChargeConjugationAudit struct {
	OrientationCandidates      int
	ChargeConjugationAvailable bool
	ActsAsInvolution           bool
	ExchangesOrientations      bool
	FixedOrientations          int
	SelectedOrientations       int
	Z2Degeneracy               bool
	Verdict                    string
}

type Summary struct {
	ContactRows                       int
	LargestGapHighRows                int
	LargestGapLowRows                 int
	OrientationCandidates             int
	OrientationSourcesAudited         int
	SourcesAvailable                  int
	SourcesSelectingOrientation       int
	ChargeConjugationInvolutions      int
	ChargeConjugationSelectedBranches int
	Z2OrientationDegeneracy           bool
	T3RSemanticOrientations           int
	TracelessOrientations             int
	T3RPullbackRowsDerived            int
	ChiralityPullbackRowsDerived      int
	BMinusLPullbackRowsDerived        int
	SU2LPullbackRowsDerived           int
	HyperchargeRowsDerived            int
	RepresentationCompleteRows        int
	RepresentationOpenRows            int
	ContactBetaRowsAllowed            int
	ContactZeroRowsProved             int
	ResidualS6Choices                 int
	ResidualNullityBefore             int
	ResidualNullityAfter              int
}

type Analysis struct {
	Previous contactorientation.Analysis

	SpectrumDescending []float64
	SplitPattern       string
	SourceAudits       []OrientationSourceAudit
	ChargeConjugation  ChargeConjugationAudit
	Summary            Summary

	ContactRows                       int
	LargestGapHighRows                int
	LargestGapLowRows                 int
	OrientationCandidates             int
	OrientationSourcesAudited         int
	SourcesAvailable                  int
	SourcesSelectingOrientation       int
	ChargeConjugationInvolutions      int
	ChargeConjugationSelectedBranches int
	Z2OrientationDegeneracy           bool
	T3RSemanticOrientations           int
	TracelessOrientations             int
	T3RPullbackRowsDerived            int
	ChiralityPullbackRowsDerived      int
	BMinusLPullbackRowsDerived        int
	SU2LPullbackRowsDerived           int
	HyperchargeRowsDerived            int
	RepresentationCompleteRows        int
	RepresentationOpenRows            int
	ContactBetaRowsAllowed            int
	ContactZeroRowsProved             int
	BetaPermissionFirewallClosed      bool
	ThresholdCorrectedBeta            bool
	FullBetaMatchingTensor            bool

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
		prev, err := contactorientation.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactorientation.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.LargestGapHighRows != 3 || prev.LargestGapLowRows != 4 || prev.OrientationCandidates != 2 {
		return Analysis{}, fmt.Errorf("Gate 142 requires Gate 141 closed-firewall two-orientation 3|4 diagnostic")
	}
	if prev.SelectedOrientations != 0 || prev.T3RSemanticOrientations != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 142 requires Gate 141 to leave orientation, T3R semantics, hypercharge rows, and contact beta rows unresolved")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 142 refuses hidden observed physical input")
	}

	charge := ChargeConjugationAudit{
		OrientationCandidates:      prev.OrientationCandidates,
		ChargeConjugationAvailable: true,
		ActsAsInvolution:           true,
		ExchangesOrientations:      true,
		FixedOrientations:          0,
		SelectedOrientations:       0,
		Z2Degeneracy:               true,
		Verdict:                    "charge conjugation exchanges the two sign orientations; it proves a Z2 branch degeneracy, not a preferred branch",
	}

	audits := []OrientationSourceAudit{
		{
			Name:               "spectral monotonicity / largest-gap order",
			Kind:               SourceSpectralMonotonicity,
			Available:          true,
			Finite:             true,
			ContactSide:        true,
			SelectsOrientation: false,
			BranchesRemaining:  2,
			Verdict:            "orders high versus low overlap but does not say which side is +T3R",
		},
		{
			Name:                  "charge-conjugation symmetry",
			Kind:                  SourceChargeConjugation,
			Available:             true,
			Finite:                true,
			ContactSide:           true,
			SelectsOrientation:    false,
			ExchangesOrientations: true,
			BranchesRemaining:     2,
			Verdict:               charge.Verdict,
		},
		{
			Name:               "B-L pullback",
			Kind:               SourceBMinusL,
			Available:          false,
			Finite:             true,
			MatterSide:         true,
			SelectsOrientation: false,
			RequiresPullback:   true,
			BranchesRemaining:  2,
			Verdict:            "B-L exists as a matter/current diagnostic but no signed contact-row pullback is derived",
		},
		{
			Name:               "chirality / T3R pullback",
			Kind:               SourceChirality,
			Available:          false,
			Finite:             true,
			MatterSide:         true,
			SelectsOrientation: false,
			RequiresPullback:   true,
			BranchesRemaining:  2,
			Verdict:            "matter-side chirality and T3R diagnostics do not yet define contact-row signs",
		},
		{
			Name:               "hypercharge consistency",
			Kind:               SourceHypercharge,
			Available:          false,
			Finite:             true,
			SelectsOrientation: false,
			RequiresPullback:   true,
			BranchesRemaining:  2,
			Verdict:            "Y=T3R+(B-L)/2 cannot orient contact rows until both signed inputs exist on the contact carrier",
		},
		{
			Name:               "source-current dual pairing",
			Kind:               SourceDualPairing,
			Available:          true,
			Finite:             true,
			ContactSide:        true,
			SelectsOrientation: false,
			BranchesRemaining:  2,
			Verdict:            "uniform pairings are row-blind and spectral self-pairings are diagnostic rather than semantic sign sources",
		},
		{
			Name:                  "observed physical constants",
			Kind:                  SourceObservedConstants,
			Available:             false,
			Finite:                false,
			SelectsOrientation:    false,
			RequiresObservedInput: true,
			BranchesRemaining:     2,
			Verdict:               "forbidden: observed constants cannot be used to orient the finite sign split",
		},
	}

	available := count(audits, func(a OrientationSourceAudit) bool { return a.Available })
	selecting := count(audits, func(a OrientationSourceAudit) bool { return a.SelectsOrientation })

	summary := Summary{
		ContactRows:                       prev.ContactRows,
		LargestGapHighRows:                prev.LargestGapHighRows,
		LargestGapLowRows:                 prev.LargestGapLowRows,
		OrientationCandidates:             prev.OrientationCandidates,
		OrientationSourcesAudited:         len(audits),
		SourcesAvailable:                  available,
		SourcesSelectingOrientation:       selecting,
		ChargeConjugationInvolutions:      1,
		ChargeConjugationSelectedBranches: charge.SelectedOrientations,
		Z2OrientationDegeneracy:           charge.Z2Degeneracy,
		T3RSemanticOrientations:           0,
		TracelessOrientations:             prev.TracelessOrientations,
		T3RPullbackRowsDerived:            0,
		ChiralityPullbackRowsDerived:      0,
		BMinusLPullbackRowsDerived:        0,
		SU2LPullbackRowsDerived:           0,
		HyperchargeRowsDerived:            0,
		RepresentationCompleteRows:        0,
		RepresentationOpenRows:            prev.ContactRows,
		ContactBetaRowsAllowed:            0,
		ContactZeroRowsProved:             0,
		ResidualS6Choices:                 prev.ResidualS6Choices,
		ResidualNullityBefore:             prev.ResidualNullityAfter,
		ResidualNullityAfter:              prev.ResidualNullityAfter,
	}

	truth := "Gate 142 searches for a source that orients the two compatible signs of the canonical 3|4 contact spectral split. Spectral order and uniform source-current pairings are finite but do not select +T3R. Charge conjugation is available as an involution but exchanges the two orientations, giving a Z2 degeneracy rather than a branch selector. B-L, chirality, T3R, SU(2)L, and hypercharge remain matter-side or current-side diagnostics without a contact-row pullback. The contact beta firewall remains closed."

	return Analysis{
		Previous:                          prev,
		SpectrumDescending:                append([]float64(nil), prev.SpectrumDescending...),
		SplitPattern:                      prev.SplitPattern,
		SourceAudits:                      audits,
		ChargeConjugation:                 charge,
		Summary:                           summary,
		ContactRows:                       prev.ContactRows,
		LargestGapHighRows:                prev.LargestGapHighRows,
		LargestGapLowRows:                 prev.LargestGapLowRows,
		OrientationCandidates:             prev.OrientationCandidates,
		OrientationSourcesAudited:         len(audits),
		SourcesAvailable:                  available,
		SourcesSelectingOrientation:       selecting,
		ChargeConjugationInvolutions:      1,
		ChargeConjugationSelectedBranches: charge.SelectedOrientations,
		Z2OrientationDegeneracy:           charge.Z2Degeneracy,
		T3RSemanticOrientations:           0,
		TracelessOrientations:             prev.TracelessOrientations,
		T3RPullbackRowsDerived:            0,
		ChiralityPullbackRowsDerived:      0,
		BMinusLPullbackRowsDerived:        0,
		SU2LPullbackRowsDerived:           0,
		HyperchargeRowsDerived:            0,
		RepresentationCompleteRows:        0,
		RepresentationOpenRows:            prev.ContactRows,
		ContactBetaRowsAllowed:            0,
		ContactZeroRowsProved:             0,
		BetaPermissionFirewallClosed:      true,
		ThresholdCorrectedBeta:            false,
		FullBetaMatchingTensor:            false,
		ResidualS6Choices:                 prev.ResidualS6Choices,
		ResidualNullityBefore:             prev.ResidualNullityAfter,
		ResidualNullityAfter:              prev.ResidualNullityAfter,
		HiddenObservedInputUsed:           false,
		PhysicalWeakAngleDerived:          false,
		FineStructureDerived:              false,
		PhysicalMassesDerived:             false,
		PhysicalScaleDerived:              false,
		TruthStatement:                    truth,
		RejectedClaims: []string{
			"charge conjugation selects the high-overlap-positive orientation",
			"charge conjugation selects the low-overlap-positive orientation",
			"spectral monotonicity is a physical T3R sign source",
			"B-L or chirality can be copied to contact rows without a pullback map",
			"hypercharge consistency can orient contact rows before contact T3R and B-L exist",
			"contact sign orientation permits threshold beta rows",
		},
		RemainingUnknowns: []string{
			"contact-row charge-conjugation breaking source, if any",
			"Fock-contact pullback for B-L, chirality, T3R, and SU(2)L",
			"contact hypercharge rows and local field map",
			"mass activation, decoupling, and threshold beta tensor",
		},
		RecommendedNextGate: "Gate 143 — contact charge-conjugation breaking source / asymmetry selector search",
	}, nil
}

func FormatSourceAudits(items []OrientationSourceAudit) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s available=%t selects=%t exchanges=%t (%s)", item.Name, item.Available, item.SelectsOrientation, item.ExchangesOrientations, item.Verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatChargeConjugation(a ChargeConjugationAudit) string {
	return fmt.Sprintf("available=%t involution=%t exchanges=%t fixed=%d selected=%d Z2=%t (%s)", a.ChargeConjugationAvailable, a.ActsAsInvolution, a.ExchangesOrientations, a.FixedOrientations, a.SelectedOrientations, a.Z2Degeneracy, a.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d split=%d|%d orientations=%d sources=%d available=%d selected=%d C-selected=%d Z2=%t T3R=%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.LargestGapHighRows, s.LargestGapLowRows, s.OrientationCandidates, s.OrientationSourcesAudited, s.SourcesAvailable, s.SourcesSelectingOrientation, s.ChargeConjugationSelectedBranches, s.Z2OrientationDegeneracy, s.T3RSemanticOrientations, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
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
