// Package contactorientation implements Gate 141: contact spectral-gap
// orientation / sign-choice obstruction theorem.
//
// Gate 140 found a real contact-side diagnostic: the seven distinct contact
// partial-overlap values have a unique largest spectral gap, producing a
// canonical 3|4 partition. Gate 141 asks the next necessary question: does the
// finite project select which side of that split should be +T3R and which side
// should be -T3R?
//
// The answer remains no. The two orientations are both mathematically
// compatible with the finite contact spectrum. Spectral monotonicity can label
// them, but it supplies no operator equation, no Fock-contact intertwiner, no
// B-L/chirality pullback, no hypercharge row, no local field map, and no
// decoupling rule. Moreover a pure +/- half-sign operator on seven rows has
// nonzero trace for either 3|4 orientation, so it cannot be promoted to a
// traceless su(2)-style generator on this seven-row carrier without additional
// structure. The gate therefore preserves the firewall and turns the sign
// choice into an explicit obstruction.
package contactorientation

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactsignsplit"
)

type OrientationKind string

const (
	HighPositive OrientationKind = "high-overlap-positive"
	LowPositive  OrientationKind = "low-overlap-positive"
)

type OrientationCandidate struct {
	Name                   string
	Kind                   OrientationKind
	PositiveRows           int
	NegativeRows           int
	PositiveSide           string
	NegativeSide           string
	Trace                  float64
	AbsoluteTrace          float64
	Traceless              bool
	SpectrallyMonotone     bool
	OrientationSelected    bool
	T3RSemanticsDerived    bool
	HyperchargeRowsDerived int
	ContactBetaRowsAllowed int
	Obstruction            string
}

type SourceAudit struct {
	Name               string
	Available          bool
	SelectsOrientation bool
	Verdict            string
}

type Summary struct {
	ContactRows                         int
	LargestGapHighRows                  int
	LargestGapLowRows                   int
	OrientationCandidates               int
	SpectrallyMonotoneOrientations      int
	SelectedOrientations                int
	T3RSemanticOrientations             int
	TracelessOrientations               int
	PureHalfSignTraceMagnitudeNumerator int
	PureHalfSignTraceMagnitudeDenom     int
	FockContactIntertwiners             int
	T3RPullbackRowsDerived              int
	ChiralityPullbackRowsDerived        int
	BMinusLPullbackRowsDerived          int
	SU2LPullbackRowsDerived             int
	HyperchargeRowsDerived              int
	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ResidualS6Choices                   int
	ResidualNullityBefore               int
	ResidualNullityAfter                int
}

type Analysis struct {
	Previous contactsignsplit.Analysis

	SpectrumDescending []float64
	SplitPattern       string
	Candidates         []OrientationCandidate
	SourceAudits       []SourceAudit
	Summary            Summary

	ContactRows                         int
	LargestGapHighRows                  int
	LargestGapLowRows                   int
	OrientationCandidates               int
	SpectrallyMonotoneOrientations      int
	SelectedOrientations                int
	T3RSemanticOrientations             int
	TracelessOrientations               int
	PureHalfSignTraceMagnitudeNumerator int
	PureHalfSignTraceMagnitudeDenom     int
	FockContactIntertwiners             int
	T3RPullbackRowsDerived              int
	ChiralityPullbackRowsDerived        int
	BMinusLPullbackRowsDerived          int
	SU2LPullbackRowsDerived             int
	HyperchargeRowsDerived              int
	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	BetaPermissionFirewallClosed        bool
	ThresholdCorrectedBeta              bool
	FullBetaMatchingTensor              bool

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
		prev, err := contactsignsplit.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactsignsplit.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.UniqueLargestGapCuts != 1 || prev.LargestGapHighRows != 3 || prev.LargestGapLowRows != 4 {
		return Analysis{}, fmt.Errorf("Gate 141 requires Gate 140 unique 3|4 largest-gap diagnostic with closed firewall")
	}
	if prev.OrientationSelectedSplits != 0 || prev.T3RSemanticSplits != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 141 requires no selected orientation, T3R semantics, hypercharge rows, or contact beta rows from Gate 140")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 141 refuses hidden observed physical input")
	}

	high := prev.LargestGapHighRows
	low := prev.LargestGapLowRows
	candidates := []OrientationCandidate{
		makeCandidate("high-overlap side positive", HighPositive, high, low, "high-overlap 3-row side", "low-overlap 4-row side"),
		makeCandidate("low-overlap side positive", LowPositive, low, high, "low-overlap 4-row side", "high-overlap 3-row side"),
	}
	sourceAudits := []SourceAudit{
		{Name: "spectral monotonicity", Available: true, SelectsOrientation: false, Verdict: "orders rows but does not decide whether high overlap is +T3R or -T3R"},
		{Name: "largest-gap partition", Available: true, SelectsOrientation: false, Verdict: "selects 3|4 partition only; orientation remains two-valued"},
		{Name: "Fock-contact intertwiner", Available: false, SelectsOrientation: false, Verdict: "no equation P A_T3R = B_T3R P has been derived"},
		{Name: "contact B-L/chirality pullback", Available: false, SelectsOrientation: false, Verdict: "no signed contact B-L or chirality rows exist"},
		{Name: "hypercharge consistency", Available: false, SelectsOrientation: false, Verdict: "hypercharge rows are still absent"},
		{Name: "local field / decoupling rule", Available: false, SelectsOrientation: false, Verdict: "no local variables, pole residues, mass activation, or threshold matching rule"},
		{Name: "observed constants", Available: false, SelectsOrientation: false, Verdict: "forbidden as a selector"},
	}

	selected := count(candidates, func(c OrientationCandidate) bool { return c.OrientationSelected })
	semantic := count(candidates, func(c OrientationCandidate) bool { return c.T3RSemanticsDerived })
	monotone := count(candidates, func(c OrientationCandidate) bool { return c.SpectrallyMonotone })
	traceless := count(candidates, func(c OrientationCandidate) bool { return c.Traceless })

	summary := Summary{
		ContactRows:                         prev.ContactRows,
		LargestGapHighRows:                  high,
		LargestGapLowRows:                   low,
		OrientationCandidates:               len(candidates),
		SpectrallyMonotoneOrientations:      monotone,
		SelectedOrientations:                selected,
		T3RSemanticOrientations:             semantic,
		TracelessOrientations:               traceless,
		PureHalfSignTraceMagnitudeNumerator: 1,
		PureHalfSignTraceMagnitudeDenom:     2,
		FockContactIntertwiners:             prev.FullOperatorIntertwiners,
		T3RPullbackRowsDerived:              0,
		ChiralityPullbackRowsDerived:        0,
		BMinusLPullbackRowsDerived:          0,
		SU2LPullbackRowsDerived:             0,
		HyperchargeRowsDerived:              0,
		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              prev.ContactRows,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ResidualS6Choices:                   prev.ResidualS6Choices,
		ResidualNullityBefore:               prev.ResidualNullityAfter,
		ResidualNullityAfter:                prev.ResidualNullityAfter,
	}

	truth := "Gate 141 tests whether the unique 3|4 contact largest-gap split has a canonical sign orientation. It does not. The two assignments high-overlap=+T3R and low-overlap=+T3R are both compatible spectral diagnostics; neither is selected by an operator equation, Fock-contact intertwiner, B-L/chirality pullback, hypercharge row, local field map, mass activation, or decoupling rule. A pure +/-1/2 sign operator on seven rows is also non-traceless for either 3|4 orientation, so it cannot be promoted to an su(2)-style T3R generator on this carrier without additional structure. The contact beta firewall remains closed."

	return Analysis{
		Previous:                            prev,
		SpectrumDescending:                  append([]float64(nil), prev.SpectrumDescending...),
		SplitPattern:                        prev.LargestGap.MultiplicityPattern,
		Candidates:                          candidates,
		SourceAudits:                        sourceAudits,
		Summary:                             summary,
		ContactRows:                         prev.ContactRows,
		LargestGapHighRows:                  high,
		LargestGapLowRows:                   low,
		OrientationCandidates:               len(candidates),
		SpectrallyMonotoneOrientations:      monotone,
		SelectedOrientations:                selected,
		T3RSemanticOrientations:             semantic,
		TracelessOrientations:               traceless,
		PureHalfSignTraceMagnitudeNumerator: 1,
		PureHalfSignTraceMagnitudeDenom:     2,
		FockContactIntertwiners:             prev.FullOperatorIntertwiners,
		T3RPullbackRowsDerived:              0,
		ChiralityPullbackRowsDerived:        0,
		BMinusLPullbackRowsDerived:          0,
		SU2LPullbackRowsDerived:             0,
		HyperchargeRowsDerived:              0,
		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              prev.ContactRows,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		BetaPermissionFirewallClosed:        true,
		ThresholdCorrectedBeta:              false,
		FullBetaMatchingTensor:              false,
		ResidualS6Choices:                   prev.ResidualS6Choices,
		ResidualNullityBefore:               prev.ResidualNullityAfter,
		ResidualNullityAfter:                prev.ResidualNullityAfter,
		HiddenObservedInputUsed:             false,
		PhysicalWeakAngleDerived:            false,
		FineStructureDerived:                false,
		PhysicalMassesDerived:               false,
		PhysicalScaleDerived:                false,
		TruthStatement:                      truth,
		RejectedClaims: []string{
			"high-overlap side is automatically +T3R",
			"low-overlap side is automatically +T3R",
			"largest-gap orientation is a physical chirality theorem",
			"a nontraceless seven-row +/- half-sign diagnostic is an su(2)-style T3R generator",
			"contact spectral orientation permits threshold beta rows",
		},
		RemainingUnknowns: []string{
			"orientation source for the 3|4 contact split",
			"Fock-contact operator equation selecting the quotient-side sign operator",
			"contact chirality, B-L, SU(2)L, and hypercharge rows",
			"local field map, Lorentz kinetic row, mass activation, and decoupling rule",
			"threshold beta tensor and physical-flow scale/coupling data",
		},
		RecommendedNextGate: "Gate 142 — contact sign-orientation source / charge-conjugation symmetry obstruction theorem",
	}, nil
}

func makeCandidate(name string, kind OrientationKind, positiveRows, negativeRows int, positiveSide, negativeSide string) OrientationCandidate {
	trace := 0.5 * float64(positiveRows-negativeRows)
	absTrace := trace
	if absTrace < 0 {
		absTrace = -absTrace
	}
	return OrientationCandidate{
		Name:                   name,
		Kind:                   kind,
		PositiveRows:           positiveRows,
		NegativeRows:           negativeRows,
		PositiveSide:           positiveSide,
		NegativeSide:           negativeSide,
		Trace:                  trace,
		AbsoluteTrace:          absTrace,
		Traceless:              absTrace < 1e-12,
		SpectrallyMonotone:     true,
		OrientationSelected:    false,
		T3RSemanticsDerived:    false,
		HyperchargeRowsDerived: 0,
		ContactBetaRowsAllowed: 0,
		Obstruction:            "orientation is compatible with the spectral split but not selected by a finite operator, source, or representation theorem",
	}
}

func FormatCandidates(candidates []OrientationCandidate) string {
	parts := make([]string, 0, len(candidates))
	for _, c := range candidates {
		parts = append(parts, fmt.Sprintf("%s: +rows=%d -rows=%d trace=%+.1f selected=%t T3R=%t", c.Name, c.PositiveRows, c.NegativeRows, c.Trace, c.OrientationSelected, c.T3RSemanticsDerived))
	}
	return strings.Join(parts, "; ")
}

func FormatSourceAudits(items []SourceAudit) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s available=%t selects=%t (%s)", item.Name, item.Available, item.SelectsOrientation, item.Verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d split=%d|%d orientations=%d monotone=%d selected=%d T3R=%d traceless=%d traceAbs=%d/%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.LargestGapHighRows, s.LargestGapLowRows, s.OrientationCandidates, s.SpectrallyMonotoneOrientations, s.SelectedOrientations, s.T3RSemanticOrientations, s.TracelessOrientations, s.PureHalfSignTraceMagnitudeNumerator, s.PureHalfSignTraceMagnitudeDenom, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
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
