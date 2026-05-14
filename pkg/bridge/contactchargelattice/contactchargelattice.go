// Package contactchargelattice implements Gate 146: contact charge lattice
// embedding / rational-spectrum obstruction theorem.
//
// Gate 145 showed that the centered contact spectral current is canonical as a
// trace-zero diagnostic but not as a charge operator. Gate 146 asks whether its
// eigenvalues can nevertheless be embedded into a finite rational charge lattice
// (half-integer, sixth-integer, seventh-balanced, or fitted rational lattice)
// strongly enough to become T3R, B-L, hypercharge, or threshold beta data.
//
// The answer is no. Common finite charge lattices do not contain the raw
// centered spectrum. A balanced 3|4 split lies in a seventh lattice, but it is a
// two-level diagnostic summary rather than the raw contact spectrum and carries
// no B-L/T3R/chirality/local-field semantics. Arbitrary rational fitting can
// approximate finite real data, but only by choosing denominators/scales not
// selected by the finite action.
package contactchargelattice

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactchargenorm"
)

type LatticeKind string

const (
	LatticeHalfInteger       LatticeKind = "half-integer"
	LatticeSixthInteger      LatticeKind = "sixth-integer"
	LatticeSeventhBalanced   LatticeKind = "seventh-balanced-split"
	LatticeBoundedRational   LatticeKind = "bounded-rational-approximation"
	LatticeFreeScaled        LatticeKind = "free-scaled-rational"
	LatticeObservedChargeFit LatticeKind = "observed-charge-fit"
)

type RationalApproximation struct {
	Value       float64
	Numerator   int
	Denominator int
	ApproxValue float64
	Error       float64
}

type LatticeCandidate struct {
	Name                     string
	Kind                     LatticeKind
	Available                bool
	CanonicalAsDiagnostic    bool
	AppliesToRawSpectrum     bool
	AppliesToBalancedSummary bool
	Step                     float64
	MaxDenominator           int
	RequiresScaleChoice      bool
	RequiresDenominatorFit   bool
	RequiresObservedInput    bool
	ExactEmbedding           bool
	ApproximateEmbedding     bool
	ExactRows                int
	ApproxRows               int
	RowsTested               int
	MaxError                 float64
	Approximations           []RationalApproximation
	T3RSemantic              bool
	BMinusLSemantic          bool
	HyperchargeSemantic      bool
	ChargeOperatorSemantic   bool
	OpensBetaPermission      bool
	Verdict                  string
}

type LatticeRequirements struct {
	FiniteSelectedLattice   bool
	RawSpectrumEmbedded     bool
	PhysicalChargeSemantics bool
	OperatorPullback        bool
	LocalFieldMap           bool
	MassActivation          bool
	DecouplingRule          bool
	ObservedInputFree       bool
	AllSatisfied            bool
	Verdict                 string
}

type Summary struct {
	ContactRows                int
	CenteredPositiveRows       int
	CenteredNegativeRows       int
	CenteredZeroRows           int
	LatticeCandidatesAudited   int
	AvailableCandidates        int
	RawExactEmbeddings         int
	RawApproxEmbeddings        int
	BalancedExactEmbeddings    int
	ChargeSemanticEmbeddings   int
	ScaleDependentCandidates   int
	ObservedFitCandidates      int
	T3RRowsDerived             int
	BMinusLRowsDerived         int
	HyperchargeRowsDerived     int
	RepresentationCompleteRows int
	RepresentationOpenRows     int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualS6Choices          int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous contactchargenorm.Analysis

	CenteredValues       []float64
	BalancedSplitValues  []float64
	HalfIntegerAudit     LatticeCandidate
	SixthIntegerAudit    LatticeCandidate
	SeventhBalancedAudit LatticeCandidate
	BoundedRationalAudit LatticeCandidate
	FreeScaledAudit      LatticeCandidate
	ObservedFitAudit     LatticeCandidate
	Candidates           []LatticeCandidate
	Requirements         LatticeRequirements
	Summary              Summary

	ContactRows                  int
	CenteredPositiveRows         int
	CenteredNegativeRows         int
	CenteredZeroRows             int
	LatticeCandidatesAudited     int
	AvailableCandidates          int
	RawExactEmbeddings           int
	RawApproxEmbeddings          int
	BalancedExactEmbeddings      int
	ChargeSemanticEmbeddings     int
	ScaleDependentCandidates     int
	ObservedFitCandidates        int
	T3RRowsDerived               int
	ChiralityRowsDerived         int
	BMinusLRowsDerived           int
	SU2LRowsDerived              int
	HyperchargeRowsDerived       int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	BetaPermissionFirewallClosed bool
	ThresholdCorrectedBeta       bool
	FullBetaMatchingTensor       bool
	ResidualS6Choices            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
	HiddenObservedInputUsed      bool
	PhysicalWeakAngleDerived     bool
	FineStructureDerived         bool
	PhysicalMassesDerived        bool
	PhysicalScaleDerived         bool

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
		prev, err := contactchargenorm.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactchargenorm.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.CenteredPositiveRows != 3 || prev.CenteredNegativeRows != 4 || prev.CenteredZeroRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 146 requires Gate 145 closed-firewall centered contact diagnostic")
	}
	if prev.ChargeSemanticNormalizations != 0 || prev.T3RRowsDerived != 0 || prev.BMinusLRowsDerived != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 146 requires no contact charge semantics or beta permission from Gate 145")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 146 refuses hidden observed physical input")
	}

	centered := append([]float64(nil), prev.CenteredValues...)
	balanced := append([]float64(nil), prev.BalancedSplitAudit.Eigenvalues...)

	half := fixedStepCandidate("half-integer lattice on raw centered current", LatticeHalfInteger, centered, 0.5, true, false, "raw centered spectrum is not contained in the half-integer lattice; binary ±1/2 was only a non-trace-zero summary")
	sixth := fixedStepCandidate("sixth-integer charge lattice on raw centered current", LatticeSixthInteger, centered, 1.0/6.0, true, false, "raw centered spectrum is not contained in the Standard-Model-like sixth-integer lattice")
	seventh := fixedStepCandidate("seventh-balanced trace-zero split lattice", LatticeSeventhBalanced, balanced, 1.0/7.0, false, true, "balanced 3|4 summary lies in a 1/7 lattice, but it is not the raw spectrum and has no charge semantics")
	bounded := boundedRationalCandidate("bounded rational approximation of raw centered spectrum", LatticeBoundedRational, centered, 840, "bounded rational approximations exist for some rows, but exact embedding and canonical denominator selection fail")
	free := freeScaledCandidate("free scaled rational lattice", LatticeFreeScaled, centered, "abstract scaling/denominator fitting can be imposed on finite data, but the scale is not selected by the finite action")
	observed := observedCandidate("observed charge lattice fit", LatticeObservedChargeFit, "forbidden: fitting to observed charges or low-energy constants would bypass the theorem ladder")

	candidates := []LatticeCandidate{half, sixth, seventh, bounded, free, observed}

	reqs := LatticeRequirements{
		FiniteSelectedLattice:   false,
		RawSpectrumEmbedded:     false,
		PhysicalChargeSemantics: false,
		OperatorPullback:        false,
		LocalFieldMap:           false,
		MassActivation:          false,
		DecouplingRule:          false,
		ObservedInputFree:       true,
		AllSatisfied:            false,
		Verdict:                 "no finite charge lattice is selected for the raw centered contact spectrum; rational summaries or fits remain diagnostics/conventions, not contact T3R, B-L, hypercharge, or beta rows",
	}

	summary := Summary{
		ContactRows:                prev.ContactRows,
		CenteredPositiveRows:       prev.CenteredPositiveRows,
		CenteredNegativeRows:       prev.CenteredNegativeRows,
		CenteredZeroRows:           prev.CenteredZeroRows,
		LatticeCandidatesAudited:   len(candidates),
		AvailableCandidates:        count(candidates, func(c LatticeCandidate) bool { return c.Available }),
		RawExactEmbeddings:         count(candidates, func(c LatticeCandidate) bool { return c.AppliesToRawSpectrum && c.ExactEmbedding }),
		RawApproxEmbeddings:        count(candidates, func(c LatticeCandidate) bool { return c.AppliesToRawSpectrum && c.ApproximateEmbedding }),
		BalancedExactEmbeddings:    count(candidates, func(c LatticeCandidate) bool { return c.AppliesToBalancedSummary && c.ExactEmbedding }),
		ChargeSemanticEmbeddings:   count(candidates, func(c LatticeCandidate) bool { return c.ChargeOperatorSemantic }),
		ScaleDependentCandidates:   count(candidates, func(c LatticeCandidate) bool { return c.RequiresScaleChoice || c.RequiresDenominatorFit }),
		ObservedFitCandidates:      count(candidates, func(c LatticeCandidate) bool { return c.RequiresObservedInput }),
		T3RRowsDerived:             0,
		BMinusLRowsDerived:         0,
		HyperchargeRowsDerived:     0,
		RepresentationCompleteRows: 0,
		RepresentationOpenRows:     prev.ContactRows,
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualS6Choices:          prev.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 146 proves that the centered contact spectral current does not select a physical rational charge lattice. Half-integer and sixth-integer lattices do not contain the raw spectrum; the balanced 1/7 split is only a two-level diagnostic summary; bounded rational approximations require denominator choices; and free or observed charge fitting is forbidden. Contact T3R, B-L, hypercharge, local field rows, mass activation, decoupling, and threshold beta permission remain sealed."

	return Analysis{
		Previous:                     prev,
		CenteredValues:               centered,
		BalancedSplitValues:          balanced,
		HalfIntegerAudit:             half,
		SixthIntegerAudit:            sixth,
		SeventhBalancedAudit:         seventh,
		BoundedRationalAudit:         bounded,
		FreeScaledAudit:              free,
		ObservedFitAudit:             observed,
		Candidates:                   candidates,
		Requirements:                 reqs,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		CenteredPositiveRows:         prev.CenteredPositiveRows,
		CenteredNegativeRows:         prev.CenteredNegativeRows,
		CenteredZeroRows:             prev.CenteredZeroRows,
		LatticeCandidatesAudited:     summary.LatticeCandidatesAudited,
		AvailableCandidates:          summary.AvailableCandidates,
		RawExactEmbeddings:           summary.RawExactEmbeddings,
		RawApproxEmbeddings:          summary.RawApproxEmbeddings,
		BalancedExactEmbeddings:      summary.BalancedExactEmbeddings,
		ChargeSemanticEmbeddings:     summary.ChargeSemanticEmbeddings,
		ScaleDependentCandidates:     summary.ScaleDependentCandidates,
		ObservedFitCandidates:        summary.ObservedFitCandidates,
		T3RRowsDerived:               0,
		ChiralityRowsDerived:         0,
		BMinusLRowsDerived:           0,
		SU2LRowsDerived:              0,
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
			"the centered contact current lies on the half-integer T3R lattice",
			"the centered contact current lies on the sixth-integer hypercharge/electric charge lattice",
			"the balanced 1/7 split is a physical charge operator",
			"rational approximation selects a finite charge lattice",
			"observed charges may choose the contact lattice scale",
		},
		RemainingUnknowns: []string{
			"finite selected contact charge lattice, if any",
			"operator pullback from B-L, chirality, T3R, SU2L, or hypercharge",
			"local field map, pole residues, mass activation, and decoupling",
			"threshold beta row permission for contact modes",
		},
		RecommendedNextGate: "Gate 147 — contact irrational-spectrum algebraic-origin / minimal-polynomial obstruction theorem",
	}, nil
}

func fixedStepCandidate(name string, kind LatticeKind, values []float64, step float64, raw, balanced bool, verdict string) LatticeCandidate {
	exactRows, maxErr := stepMembership(values, step, 1e-9)
	return LatticeCandidate{
		Name:                     name,
		Kind:                     kind,
		Available:                true,
		CanonicalAsDiagnostic:    balanced,
		AppliesToRawSpectrum:     raw,
		AppliesToBalancedSummary: balanced,
		Step:                     step,
		MaxDenominator:           denominatorFromStep(step),
		ExactEmbedding:           len(values) > 0 && exactRows == len(values),
		ApproximateEmbedding:     false,
		ExactRows:                exactRows,
		ApproxRows:               exactRows,
		RowsTested:               len(values),
		MaxError:                 maxErr,
		ChargeOperatorSemantic:   false,
		Verdict:                  verdict,
	}
}

func boundedRationalCandidate(name string, kind LatticeKind, values []float64, maxDenom int, verdict string) LatticeCandidate {
	approxs := make([]RationalApproximation, 0, len(values))
	exactRows, approxRows := 0, 0
	maxErr := 0.0
	for _, v := range values {
		ap := bestRational(v, maxDenom)
		approxs = append(approxs, ap)
		if ap.Error <= 1e-9 {
			exactRows++
		}
		if ap.Error <= 1e-6 {
			approxRows++
		}
		if ap.Error > maxErr {
			maxErr = ap.Error
		}
	}
	return LatticeCandidate{
		Name:                   name,
		Kind:                   kind,
		Available:              true,
		CanonicalAsDiagnostic:  false,
		AppliesToRawSpectrum:   true,
		MaxDenominator:         maxDenom,
		RequiresDenominatorFit: true,
		ExactEmbedding:         len(values) > 0 && exactRows == len(values),
		ApproximateEmbedding:   len(values) > 0 && approxRows == len(values),
		ExactRows:              exactRows,
		ApproxRows:             approxRows,
		RowsTested:             len(values),
		MaxError:               maxErr,
		Approximations:         approxs,
		ChargeOperatorSemantic: false,
		Verdict:                verdict,
	}
}

func freeScaledCandidate(name string, kind LatticeKind, values []float64, verdict string) LatticeCandidate {
	return LatticeCandidate{
		Name:                   name,
		Kind:                   kind,
		Available:              true,
		CanonicalAsDiagnostic:  false,
		AppliesToRawSpectrum:   true,
		RequiresScaleChoice:    true,
		RequiresDenominatorFit: true,
		ExactEmbedding:         false,
		ApproximateEmbedding:   true,
		ExactRows:              0,
		ApproxRows:             len(values),
		RowsTested:             len(values),
		ChargeOperatorSemantic: false,
		Verdict:                verdict,
	}
}

func observedCandidate(name string, kind LatticeKind, verdict string) LatticeCandidate {
	return LatticeCandidate{
		Name:                   name,
		Kind:                   kind,
		Available:              false,
		RequiresObservedInput:  true,
		ChargeOperatorSemantic: false,
		Verdict:                verdict,
	}
}

func FormatCandidate(c LatticeCandidate) string {
	return fmt.Sprintf("%s available=%t raw=%t balanced=%t exact=%t approx=%t exactRows=%d/%d approxRows=%d/%d scale=%t denomFit=%t observed=%t chargeSemantic=%t maxErr=%.3g (%s)", c.Name, c.Available, c.AppliesToRawSpectrum, c.AppliesToBalancedSummary, c.ExactEmbedding, c.ApproximateEmbedding, c.ExactRows, c.RowsTested, c.ApproxRows, c.RowsTested, c.RequiresScaleChoice, c.RequiresDenominatorFit, c.RequiresObservedInput, c.ChargeOperatorSemantic, c.MaxError, c.Verdict)
}

func FormatCandidates(items []LatticeCandidate) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, FormatCandidate(item))
	}
	return strings.Join(parts, "; ")
}

func FormatApproximations(items []RationalApproximation) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%.10f≈%d/%d(err=%.3g)", item.Value, item.Numerator, item.Denominator, item.Error))
	}
	return strings.Join(parts, ", ")
}

func FormatRequirements(r LatticeRequirements) string {
	return fmt.Sprintf("selectedLattice=%t rawEmbedded=%t semantics=%t pullback=%t local=%t mass=%t decoupling=%t observedFree=%t all=%t (%s)", r.FiniteSelectedLattice, r.RawSpectrumEmbedded, r.PhysicalChargeSemantics, r.OperatorPullback, r.LocalFieldMap, r.MassActivation, r.DecouplingRule, r.ObservedInputFree, r.AllSatisfied, r.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d centered=%d/%d/%d candidates=%d available=%d rawExact=%d rawApprox=%d balancedExact=%d semantic=%d scaleDependent=%d observedFit=%d T3R=%d B-L=%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.CenteredPositiveRows, s.CenteredNegativeRows, s.CenteredZeroRows, s.LatticeCandidatesAudited, s.AvailableCandidates, s.RawExactEmbeddings, s.RawApproxEmbeddings, s.BalancedExactEmbeddings, s.ChargeSemanticEmbeddings, s.ScaleDependentCandidates, s.ObservedFitCandidates, s.T3RRowsDerived, s.BMinusLRowsDerived, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(items []string) string { return strings.Join(items, "; ") }

func stepMembership(values []float64, step float64, eps float64) (int, float64) {
	if step <= 0 {
		return 0, 0
	}
	rows := 0
	maxErr := 0.0
	for _, v := range values {
		n := math.Round(v / step)
		err := math.Abs(v - n*step)
		if err <= eps {
			rows++
		}
		if err > maxErr {
			maxErr = err
		}
	}
	return rows, maxErr
}

func denominatorFromStep(step float64) int {
	if step <= 0 {
		return 0
	}
	return int(math.Round(1 / step))
}

func bestRational(v float64, maxDenom int) RationalApproximation {
	bestN, bestD := 0, 1
	bestErr := math.Inf(1)
	for d := 1; d <= maxDenom; d++ {
		n := int(math.Round(v * float64(d)))
		approx := float64(n) / float64(d)
		err := math.Abs(v - approx)
		if err < bestErr {
			bestErr = err
			bestN = n
			bestD = d
		}
	}
	return RationalApproximation{Value: v, Numerator: bestN, Denominator: bestD, ApproxValue: float64(bestN) / float64(bestD), Error: bestErr}
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
