// Package contactsignsplit implements Gate 140: contact T3R sign-split
// naturality / spectral-cut obstruction theorem.
//
// Gate 139 showed that a quotient-side T3R operator on the seven contact rows
// would require assigning +/-1/2 signs to the contact rows, but no target
// operator was selected. Gate 140 tests the strongest contact-only rescue:
// using spectral cuts of the seven distinct contact partial-overlap values to
// select a sign split.
//
// The contact spectrum does contain a unique largest spectral gap. This gives a
// canonical diagnostic 3|4 partition of the seven contact rows. However, this
// still does not derive T3R: the +/- orientation of the split is not selected,
// the split is odd-dimensional rather than a pulled-back matter 8|8 chiral
// spectrum, and no local field, chirality, hypercharge, mass activation, or
// decoupling row is attached. The result is therefore a sharper firewall rather
// than a physical threshold row.
package contactsignsplit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contacttargetoperator"
)

type CutKind string

const (
	ProperSpectralCut CutKind = "proper-spectral-cut"
	LargestGapCut     CutKind = "largest-gap-cut"
	MedianCut         CutKind = "median-cut"
	OrientationFlip   CutKind = "orientation-flip"
)

type SpectralCut struct {
	Name                   string
	Kind                   CutKind
	LowerIndex             int
	UpperIndex             int
	Gap                    float64
	HighRows               int
	LowRows                int
	MultiplicityPattern    string
	CanonicalDiagnostic    bool
	UniqueLargestGap       bool
	OrientationSelected    bool
	T3RSemanticsDerived    bool
	HyperchargeRowsDerived int
	ContactBetaRowsAllowed int
	Obstruction            string
}

type Summary struct {
	ContactRows                  int
	DistinctSpectralRows         int
	ProperSpectralCuts           int
	UniqueLargestGapCuts         int
	LargestGapHighRows           int
	LargestGapLowRows            int
	LargestGapOrientations       int
	MedianRowsBelow              int
	MedianRowsAt                 int
	MedianRowsAbove              int
	AbstractSignAssignments      int
	SpectralCutSignAssignments   int
	CanonicalDiagnosticSplits    int
	OrientationSelectedSplits    int
	T3RSemanticSplits            int
	CanonicalT3RTargetOperators  int
	QuotientInducedT3RTargetOps  int
	FullOperatorIntertwiners     int
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
	Previous contacttargetoperator.Analysis

	SpectrumDescending []float64
	Cuts               []SpectralCut
	LargestGap         SpectralCut
	Summary            Summary

	ContactRows                  int
	DistinctSpectralRows         int
	ProperSpectralCuts           int
	UniqueLargestGapCuts         int
	LargestGapHighRows           int
	LargestGapLowRows            int
	LargestGapOrientations       int
	MedianRowsBelow              int
	MedianRowsAt                 int
	MedianRowsAbove              int
	AbstractSignAssignments      int
	SpectralCutSignAssignments   int
	CanonicalDiagnosticSplits    int
	OrientationSelectedSplits    int
	T3RSemanticSplits            int
	T3RTargetOperatorsDerived    int
	QuotientInducedT3RTargetOps  int
	FullOperatorIntertwiners     int
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
		prev, err := contacttargetoperator.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contacttargetoperator.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.ContactSpectralDistinctRows != 7 {
		return Analysis{}, fmt.Errorf("Gate 140 requires Gate 139 seven-row contact target-operator obstruction with closed firewall")
	}
	if prev.CanonicalT3RTargetOperators != 0 || prev.QuotientInducedT3RTargetOperators != 0 || prev.T3RPullbackRowsDerived != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 140 requires no contact T3R/hypercharge/beta rows from Gate 139")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 140 refuses hidden observed physical input")
	}

	spectrum := append([]float64(nil), prev.ContactSpectrum...)
	sort.Sort(sort.Reverse(sort.Float64Slice(spectrum)))
	cuts, largest, uniqueLargest := buildSpectralCuts(spectrum)
	if !uniqueLargest {
		return Analysis{}, fmt.Errorf("Gate 140 expected a unique largest spectral gap diagnostic")
	}

	contactRows := prev.ContactRows
	properCuts := contactRows - 1
	cutSigns := 2 * properCuts
	abstractSigns := prev.T3RRowSignAssignments
	canonicalDiagnosticSplits := count(cuts, func(c SpectralCut) bool { return c.CanonicalDiagnostic })
	orientationSelected := count(cuts, func(c SpectralCut) bool { return c.OrientationSelected })
	t3rSemantic := count(cuts, func(c SpectralCut) bool { return c.T3RSemanticsDerived })
	medianBelow, medianAt, medianAbove := medianAudit(spectrum)

	summary := Summary{
		ContactRows:                  contactRows,
		DistinctSpectralRows:         prev.ContactSpectralDistinctRows,
		ProperSpectralCuts:           properCuts,
		UniqueLargestGapCuts:         1,
		LargestGapHighRows:           largest.HighRows,
		LargestGapLowRows:            largest.LowRows,
		LargestGapOrientations:       2,
		MedianRowsBelow:              medianBelow,
		MedianRowsAt:                 medianAt,
		MedianRowsAbove:              medianAbove,
		AbstractSignAssignments:      abstractSigns,
		SpectralCutSignAssignments:   cutSigns,
		CanonicalDiagnosticSplits:    canonicalDiagnosticSplits,
		OrientationSelectedSplits:    orientationSelected,
		T3RSemanticSplits:            t3rSemantic,
		CanonicalT3RTargetOperators:  0,
		QuotientInducedT3RTargetOps:  0,
		FullOperatorIntertwiners:     prev.FullOperatorIntertwinersDerived,
		T3RPullbackRowsDerived:       0,
		ChiralityPullbackRowsDerived: 0,
		BMinusLPullbackRowsDerived:   0,
		SU2LPullbackRowsDerived:      0,
		HyperchargeRowsDerived:       0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       contactRows,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
	}

	truth := "Gate 140 tests spectral-cut naturality for contact-side T3R signs. The seven contact rows have a unique largest spectral gap, giving a canonical diagnostic 3|4 split, but this is still not a T3R operator: the sign orientation is not selected, the cut is spectral rather than representation-theoretic, and no Fock-contact intertwiner, chirality, B-L pullback, SU(2)L action, hypercharge row, local field map, mass activation, or decoupling rule is derived. The contact beta firewall remains closed."

	return Analysis{
		Previous:                     prev,
		SpectrumDescending:           spectrum,
		Cuts:                         cuts,
		LargestGap:                   largest,
		Summary:                      summary,
		ContactRows:                  contactRows,
		DistinctSpectralRows:         prev.ContactSpectralDistinctRows,
		ProperSpectralCuts:           properCuts,
		UniqueLargestGapCuts:         1,
		LargestGapHighRows:           largest.HighRows,
		LargestGapLowRows:            largest.LowRows,
		LargestGapOrientations:       2,
		MedianRowsBelow:              medianBelow,
		MedianRowsAt:                 medianAt,
		MedianRowsAbove:              medianAbove,
		AbstractSignAssignments:      abstractSigns,
		SpectralCutSignAssignments:   cutSigns,
		CanonicalDiagnosticSplits:    canonicalDiagnosticSplits,
		OrientationSelectedSplits:    orientationSelected,
		T3RSemanticSplits:            t3rSemantic,
		T3RTargetOperatorsDerived:    0,
		QuotientInducedT3RTargetOps:  0,
		FullOperatorIntertwiners:     prev.FullOperatorIntertwinersDerived,
		T3RPullbackRowsDerived:       0,
		ChiralityPullbackRowsDerived: 0,
		BMinusLPullbackRowsDerived:   0,
		SU2LPullbackRowsDerived:      0,
		HyperchargeRowsDerived:       0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       contactRows,
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
			"the unique largest spectral gap is a T3R sign theorem",
			"the 3|4 spectral cut selects the plus/minus orientation of T3R",
			"a median or largest-gap cut supplies chirality or hypercharge rows",
			"spectral ordering is equivalent to local field representation semantics",
			"contact spectral cuts allow threshold beta rows before a decoupling theorem",
		},
		RemainingUnknowns: []string{
			"orientation rule for any contact sign split",
			"operator equation P A_T3R = B_T3R P selecting the target sign operator",
			"contact chirality, B-L, SU(2)L, and hypercharge rows",
			"local field map, Lorentz kinetic row, mass activation, and decoupling rule",
			"threshold beta tensor and physical-flow scale/coupling data",
		},
		RecommendedNextGate: "Gate 141 — contact spectral-gap orientation / sign-choice obstruction theorem",
	}, nil
}

func buildSpectralCuts(spectrum []float64) ([]SpectralCut, SpectralCut, bool) {
	cuts := make([]SpectralCut, 0, len(spectrum)-1)
	maxGap := math.Inf(-1)
	maxCount := 0
	maxIndex := -1
	for i := 0; i < len(spectrum)-1; i++ {
		gap := spectrum[i] - spectrum[i+1]
		if gap > maxGap+1e-12 {
			maxGap = gap
			maxCount = 1
			maxIndex = i
		} else if math.Abs(gap-maxGap) <= 1e-12 {
			maxCount++
		}
	}
	for i := 0; i < len(spectrum)-1; i++ {
		gap := spectrum[i] - spectrum[i+1]
		high := i + 1
		low := len(spectrum) - high
		largest := i == maxIndex && maxCount == 1
		kind := ProperSpectralCut
		if largest {
			kind = LargestGapCut
		}
		cuts = append(cuts, SpectralCut{
			Name:                fmt.Sprintf("cut after row %d", high),
			Kind:                kind,
			LowerIndex:          i,
			UpperIndex:          i + 1,
			Gap:                 gap,
			HighRows:            high,
			LowRows:             low,
			MultiplicityPattern: fmt.Sprintf("%d|%d", high, low),
			CanonicalDiagnostic: largest,
			UniqueLargestGap:    largest,
			OrientationSelected: false,
			T3RSemanticsDerived: false,
			Obstruction:         "spectral partition is diagnostic; T3R sign orientation and representation semantics are not selected",
		})
	}
	if maxIndex < 0 {
		return cuts, SpectralCut{}, false
	}
	return cuts, cuts[maxIndex], maxCount == 1
}

func medianAudit(spectrum []float64) (below, at, above int) {
	if len(spectrum) == 0 {
		return 0, 0, 0
	}
	mid := len(spectrum) / 2
	return len(spectrum) - mid - 1, 1, mid
}

func FormatSpectrum(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%.10f", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatCuts(cuts []SpectralCut) string {
	parts := make([]string, 0, len(cuts))
	for _, c := range cuts {
		mark := ""
		if c.UniqueLargestGap {
			mark = " largest"
		}
		parts = append(parts, fmt.Sprintf("%s gap=%.10f pattern=%s%s orientation=%t T3R=%t", c.Name, c.Gap, c.MultiplicityPattern, mark, c.OrientationSelected, c.T3RSemanticsDerived))
	}
	return strings.Join(parts, "; ")
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d distinct=%d cuts=%d largest=%d|%d orientations=%d signs=%d cutSigns=%d diagnosticSplits=%d oriented=%d T3R=%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.DistinctSpectralRows, s.ProperSpectralCuts, s.LargestGapHighRows, s.LargestGapLowRows, s.LargestGapOrientations, s.AbstractSignAssignments, s.SpectralCutSignAssignments, s.CanonicalDiagnosticSplits, s.OrientationSelectedSplits, s.T3RSemanticSplits, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
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
