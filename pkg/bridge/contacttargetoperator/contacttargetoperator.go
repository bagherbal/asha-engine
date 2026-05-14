// Package contacttargetoperator implements Gate 139: contact target-operator
// reconstruction / quotient-side T3R spectrum search.
//
// Gate 138 showed that generic Fock-to-contact quotients exist, but no
// canonical 9-dimensional kernel or operator-intertwiner is selected. Gate 139
// therefore searches from the opposite side: can the seven contact rows
// reconstruct a target operator with T3R-like spectrum by themselves?
//
// The result is another firewall. The contact spectrum has a canonical
// diagnostic diagonal operator with seven distinct values, but it is not a T3R
// operator. A quotient-side T3R target would require assigning ±1/2 signs to
// seven contact rows, giving 2^7 possible row-sign assignments and 8 abstract
// multiplicity splits. No finite rule selects one, and no Fock kernel is paired
// with any target operator. Therefore matter T3R/chirality still cannot be
// pulled into contact hypercharge rows.
package contacttargetoperator

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fockcontactkernel"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
)

type OperatorKind string

const (
	ContactSpectralDiagonal OperatorKind = "contact-spectral-diagonal"
	ZeroContactOperator     OperatorKind = "zero-contact-operator"
	ScalarContactOperator   OperatorKind = "scalar-contact-operator"
	AbstractT3RSplit        OperatorKind = "abstract-T3R-split"
	SpectralThresholdSign   OperatorKind = "spectral-threshold-sign"
	QuotientInducedT3R      OperatorKind = "quotient-induced-T3R"
)

type Candidate struct {
	Name                      string
	Kind                      OperatorKind
	Dimension                 int
	Eigenvalues               []float64
	CanonicalContactSide      bool
	DiagnosticOnly            bool
	T3RSpectrum               bool
	T3RMultiplicityPattern    string
	NeedsRowSignAssignment    bool
	NeedsFockKernel           bool
	NeedsObservedInput        bool
	TargetT3ROperatorDerived  bool
	HyperchargeRowsDerived    int
	RepresentationRowsDerived int
	ContactBetaRowsAllowed    int
	HiddenDiscreteChoices     int
	Obstruction               string
}

type SplitCandidate struct {
	PositiveRows          int
	NegativeRows          int
	MultiplicityPattern   string
	RowAssignments        int
	NonScalarT3RLike      bool
	CanonicalSelected     bool
	NeedsContactRowLabels bool
	NeedsFockKernel       bool
	Obstruction           string
}

type Summary struct {
	ContactRows                       int
	MatterPositiveMultiplicity        int
	MatterNegativeMultiplicity        int
	ContactSpectralRows               int
	ContactSpectralDistinctRows       int
	DiagnosticContactOperators        int
	AbstractT3RMultiplicitySplits     int
	T3RRowSignAssignments             int
	NonScalarT3RRowSignAssignments    int
	CanonicalT3RTargetOperators       int
	QuotientInducedT3RTargetOperators int
	CanonicalFockKernels              int
	FullOperatorIntertwinersDerived   int
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
	Previous  fockcontactkernel.Analysis
	Threshold threshold.Analysis

	ContactSpectrum []float64
	Candidates      []Candidate
	Splits          []SplitCandidate
	Summary         Summary

	ContactRows                       int
	MatterPositiveMultiplicity        int
	MatterNegativeMultiplicity        int
	ContactSpectralRows               int
	ContactSpectralDistinctRows       int
	DiagnosticContactOperators        int
	AbstractT3RMultiplicitySplits     int
	T3RRowSignAssignments             int
	NonScalarT3RRowSignAssignments    int
	CanonicalT3RTargetOperators       int
	QuotientInducedT3RTargetOperators int
	CanonicalFockKernels              int
	FullOperatorIntertwinersDerived   int
	T3RPullbackRowsDerived            int
	ChiralityPullbackRowsDerived      int
	BMinusLPullbackRowsDerived        int
	SU2LPullbackRowsDerived           int
	HyperchargeRowsDerived            int
	RepresentationCompleteRows        int
	RepresentationOpenRows            int
	ContactBetaRowsAllowed            int
	ContactZeroRowsProved             int
	FullBetaMatchingTensorDerived     bool
	ThresholdCorrectedBetaDerived     bool
	BetaPermissionFirewallClosed      bool

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
		prev, err := fockcontactkernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		thr, err := threshold.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, thr)
	})
	return defaultValue, defaultErr
}

func Build(prev fockcontactkernel.Analysis, thr threshold.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.MatterDimension != 16 || prev.CanonicalKernelCandidates != 1 || prev.FullOperatorIntertwinersDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 139 requires Gate 138 kernel obstruction with closed firewall")
	}
	if prev.TargetContactOperatorsDerived != 0 || prev.T3RPullbackRowsDerived != 0 || prev.ChiralityPullbackRowsDerived != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 139 requires no target contact operator or contact pullback rows from Gate 138")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 139 refuses hidden physical input from Gate 138")
	}
	if len(thr.ContactPartialOverlap) != 7 {
		return Analysis{}, fmt.Errorf("Gate 139 expects seven contact partial-overlap rows, got %d", len(thr.ContactPartialOverlap))
	}

	contactRows := prev.ContactRows
	spectrum := append([]float64(nil), thr.ContactPartialOverlap...)
	distinct := countDistinct(spectrum, 1e-10)
	if distinct != contactRows {
		return Analysis{}, fmt.Errorf("Gate 139 expects seven distinct contact spectral rows, got %d", distinct)
	}

	matterPos := 8
	matterNeg := 8
	splits := buildSplits(contactRows)
	rowSignAssignments := pow2(contactRows)
	nonScalarAssignments := rowSignAssignments - 2 // all-plus/all-minus are scalar sign operators, not split T3R.
	candidates := buildCandidates(spectrum, splits, prev)

	diagnosticOps := count(candidates, func(c Candidate) bool { return c.DiagnosticOnly && c.CanonicalContactSide })
	canonicalT3RTargets := count(candidates, func(c Candidate) bool { return c.TargetT3ROperatorDerived })
	quotientTargets := count(candidates, func(c Candidate) bool { return c.Kind == QuotientInducedT3R && c.TargetT3ROperatorDerived })

	summary := Summary{
		ContactRows:                       contactRows,
		MatterPositiveMultiplicity:        matterPos,
		MatterNegativeMultiplicity:        matterNeg,
		ContactSpectralRows:               len(spectrum),
		ContactSpectralDistinctRows:       distinct,
		DiagnosticContactOperators:        diagnosticOps,
		AbstractT3RMultiplicitySplits:     len(splits),
		T3RRowSignAssignments:             rowSignAssignments,
		NonScalarT3RRowSignAssignments:    nonScalarAssignments,
		CanonicalT3RTargetOperators:       canonicalT3RTargets,
		QuotientInducedT3RTargetOperators: quotientTargets,
		CanonicalFockKernels:              prev.CanonicalKernelCandidates,
		FullOperatorIntertwinersDerived:   prev.FullOperatorIntertwinersDerived,
		T3RPullbackRowsDerived:            0,
		ChiralityPullbackRowsDerived:      0,
		BMinusLPullbackRowsDerived:        0,
		SU2LPullbackRowsDerived:           0,
		HyperchargeRowsDerived:            0,
		RepresentationCompleteRows:        0,
		RepresentationOpenRows:            contactRows,
		ContactBetaRowsAllowed:            0,
		ContactZeroRowsProved:             0,
		ResidualS6Choices:                 prev.ResidualS6Choices,
		ResidualNullityBefore:             prev.ResidualNullityAfter,
		ResidualNullityAfter:              prev.ResidualNullityAfter,
	}

	truth := "Gate 139 searches from the contact side for a quotient target operator. The seven contact partial-overlap rows define a canonical spectral diagonal diagnostic, but its seven distinct eigenvalues are not the ±1/2 T3R spectrum. A quotient-side T3R operator would require choosing a split of seven contact rows into positive and negative T3R sectors, giving 8 abstract multiplicity splits and 2^7 row-sign assignments. No Fock kernel, row-sign rule, or operator equation selects one, so contact T3R/hypercharge rows remain underived."

	return Analysis{
		Previous:                          prev,
		Threshold:                         thr,
		ContactSpectrum:                   spectrum,
		Candidates:                        candidates,
		Splits:                            splits,
		Summary:                           summary,
		ContactRows:                       contactRows,
		MatterPositiveMultiplicity:        matterPos,
		MatterNegativeMultiplicity:        matterNeg,
		ContactSpectralRows:               len(spectrum),
		ContactSpectralDistinctRows:       distinct,
		DiagnosticContactOperators:        diagnosticOps,
		AbstractT3RMultiplicitySplits:     len(splits),
		T3RRowSignAssignments:             rowSignAssignments,
		NonScalarT3RRowSignAssignments:    nonScalarAssignments,
		CanonicalT3RTargetOperators:       canonicalT3RTargets,
		QuotientInducedT3RTargetOperators: quotientTargets,
		CanonicalFockKernels:              prev.CanonicalKernelCandidates,
		FullOperatorIntertwinersDerived:   prev.FullOperatorIntertwinersDerived,
		T3RPullbackRowsDerived:            0,
		ChiralityPullbackRowsDerived:      0,
		BMinusLPullbackRowsDerived:        0,
		SU2LPullbackRowsDerived:           0,
		HyperchargeRowsDerived:            0,
		RepresentationCompleteRows:        0,
		RepresentationOpenRows:            contactRows,
		ContactBetaRowsAllowed:            0,
		ContactZeroRowsProved:             0,
		FullBetaMatchingTensorDerived:     false,
		ThresholdCorrectedBetaDerived:     false,
		BetaPermissionFirewallClosed:      true,
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
			"the contact spectral diagonal is a T3R operator",
			"a ±1/2 row-sign assignment is selected by the contact spectrum alone",
			"abstract quotient multiplicities select a Fock-to-contact kernel",
			"median/min/max spectral cuts are representation semantics",
			"contact T3R rows can be used to form hypercharge before a target operator is derived",
		},
		RemainingUnknowns: []string{
			"canonical quotient-side T3R target operator on R7_contact",
			"row-sign assignment or split rule for the seven contact rows",
			"Fock kernel and operator equation P A_T3R = B_T3R P",
			"contact chirality, B-L, SU(2)L, and hypercharge rows",
			"local field map, mass activation, decoupling rule, and threshold beta tensor",
		},
		RecommendedNextGate: "Gate 140 — contact T3R sign-split naturality / spectral-cut obstruction theorem",
	}, nil
}

func buildCandidates(spectrum []float64, splits []SplitCandidate, prev fockcontactkernel.Analysis) []Candidate {
	rows := len(spectrum)
	return []Candidate{
		{Name: "canonical contact spectral diagonal", Kind: ContactSpectralDiagonal, Dimension: rows, Eigenvalues: append([]float64(nil), spectrum...), CanonicalContactSide: true, DiagnosticOnly: true, T3RSpectrum: false, Obstruction: "canonical on contact rows, but eigenvalues are seven overlap diagnostics rather than ±1/2 T3R charges"},
		{Name: "zero contact operator", Kind: ZeroContactOperator, Dimension: rows, Eigenvalues: make([]float64, rows), CanonicalContactSide: true, DiagnosticOnly: true, T3RSpectrum: false, Obstruction: "canonical but annihilates row distinctions and cannot encode T3R"},
		{Name: "+1/2 scalar sign operator", Kind: ScalarContactOperator, Dimension: rows, Eigenvalues: filled(rows, 0.5), CanonicalContactSide: true, DiagnosticOnly: false, T3RSpectrum: true, T3RMultiplicityPattern: fmt.Sprintf("%d+0", rows), Obstruction: "has ±1/2 form but scalar action only; no chirality or hypercharge split"},
		{Name: "abstract quotient-side T3R split family", Kind: AbstractT3RSplit, Dimension: rows, CanonicalContactSide: false, DiagnosticOnly: false, T3RSpectrum: true, NeedsRowSignAssignment: true, HiddenDiscreteChoices: pow2(rows), Obstruction: fmt.Sprintf("%d abstract multiplicity splits and %d row-sign assignments exist, but none is selected", len(splits), pow2(rows))},
		{Name: "spectral cut sign operator", Kind: SpectralThresholdSign, Dimension: rows, CanonicalContactSide: false, DiagnosticOnly: true, T3RSpectrum: true, NeedsRowSignAssignment: true, HiddenDiscreteChoices: rows - 1, Obstruction: "cutting the ordered contact spectrum gives diagnostic signs only; cut orientation and threshold are not representation semantics"},
		{Name: "quotient-induced contact T3R target", Kind: QuotientInducedT3R, Dimension: rows, CanonicalContactSide: false, DiagnosticOnly: false, T3RSpectrum: true, NeedsFockKernel: true, HiddenDiscreteChoices: prev.ResidualS6Choices, Obstruction: "requires a canonical Fock kernel and target operator B_T3R satisfying P A_T3R = B_T3R P; Gate 138 derived none"},
	}
}

func buildSplits(rows int) []SplitCandidate {
	out := make([]SplitCandidate, 0, rows+1)
	for p := 0; p <= rows; p++ {
		n := rows - p
		assignments := binomial(rows, p)
		nonScalar := p != 0 && n != 0
		out = append(out, SplitCandidate{
			PositiveRows:          p,
			NegativeRows:          n,
			MultiplicityPattern:   fmt.Sprintf("%d+%d", p, n),
			RowAssignments:        assignments,
			NonScalarT3RLike:      nonScalar,
			CanonicalSelected:     false,
			NeedsContactRowLabels: true,
			NeedsFockKernel:       true,
			Obstruction:           "multiplicity split is abstract until contact rows and a Fock kernel are selected",
		})
	}
	return out
}

func FormatCandidates(c []Candidate) string {
	parts := make([]string, 0, len(c))
	for _, x := range c {
		parts = append(parts, fmt.Sprintf("%s[%s dim=%d canonical=%t diagnostic=%t T3R=%t pattern=%s hidden=%d beta=%d; %s]", x.Name, x.Kind, x.Dimension, x.CanonicalContactSide, x.DiagnosticOnly, x.T3RSpectrum, x.T3RMultiplicityPattern, x.HiddenDiscreteChoices, x.ContactBetaRowsAllowed, x.Obstruction))
	}
	return strings.Join(parts, "; ")
}

func FormatSplits(s []SplitCandidate) string {
	parts := make([]string, 0, len(s))
	for _, x := range s {
		parts = append(parts, fmt.Sprintf("%s assignments=%d nonScalar=%t selected=%t", x.MultiplicityPattern, x.RowAssignments, x.NonScalarT3RLike, x.CanonicalSelected))
	}
	return strings.Join(parts, "; ")
}

func FormatSpectrum(xs []float64) string {
	vals := append([]float64(nil), xs...)
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	parts := make([]string, 0, len(vals))
	for _, x := range vals {
		parts = append(parts, fmt.Sprintf("%.10f", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d spectrum=%d distinct=%d diagnosticOps=%d splits=%d signs=%d nonScalarSigns=%d targetT3R=%d quotientT3R=%d kernels=%d intertwiners=%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.ContactSpectralRows, s.ContactSpectralDistinctRows, s.DiagnosticContactOperators, s.AbstractT3RMultiplicitySplits, s.T3RRowSignAssignments, s.NonScalarT3RRowSignAssignments, s.CanonicalT3RTargetOperators, s.QuotientInducedT3RTargetOperators, s.CanonicalFockKernels, s.FullOperatorIntertwinersDerived, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(items []string) string { return strings.Join(items, "; ") }

func filled(n int, v float64) []float64 {
	out := make([]float64, n)
	for i := range out {
		out[i] = v
	}
	return out
}

func countDistinct(xs []float64, eps float64) int {
	if len(xs) == 0 {
		return 0
	}
	vals := append([]float64(nil), xs...)
	sort.Float64s(vals)
	n := 1
	prev := vals[0]
	for _, v := range vals[1:] {
		if math.Abs(v-prev) > eps {
			n++
			prev = v
		}
	}
	return n
}

func binomial(n int, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	out := 1
	for i := 1; i <= k; i++ {
		out = out * (n - k + i) / i
	}
	return out
}

func pow2(n int) int {
	out := 1
	for i := 0; i < n; i++ {
		out *= 2
	}
	return out
}

func count[T any](items []T, pred func(T) bool) int {
	n := 0
	for _, item := range items {
		if pred(item) {
			n++
		}
	}
	return n
}
