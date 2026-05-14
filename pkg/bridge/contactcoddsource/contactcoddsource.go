// Package contactcoddsource implements Gate 144: contact C-odd source
// functional / finite signed-current construction attempt.
//
// Gate 143 found real finite asymmetry diagnostics in the contact spectrum, but
// no charge-conjugation-breaking source. Gate 144 asks a constructive question:
// can the seven contact rows themselves generate a finite signed current or
// C-odd functional that orients the 3|4 split and opens the beta firewall?
//
// The strongest current construction available is the centered contact spectral
// functional J = D_contact - tr(D_contact)/7. It is canonical, trace-zero, and
// signed: the three high-overlap rows are positive and the four low-overlap rows
// are negative. This is genuine finite diagnostic structure. But it is not yet a
// physical C-odd source: it has no charge-conjugation action, no source-current
// coupling, no Fock/contact pullback, no local field map, no T3R/B-L/hypercharge
// semantics, and no decoupling rule. Therefore the signed-current construction
// remains diagnostic only and the contact beta firewall remains closed.
package contactcoddsource

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactasymmetry"
)

type SignedSourceKind string

const (
	SourceCenteredSpectrum    SignedSourceKind = "centered-contact-spectrum"
	SourceBinaryGapSplit      SignedSourceKind = "binary-largest-gap-split"
	SourceTracelessGapSplit   SignedSourceKind = "traceless-balanced-gap-split"
	SourceChargeOddCurrent    SignedSourceKind = "charge-conjugation-odd-current"
	SourcePullbackCharge      SignedSourceKind = "B-L-T3R-hypercharge-pullback"
	SourceLocalPoleResidue    SignedSourceKind = "local-field-pole-residue-current"
	SourceObservedOrientation SignedSourceKind = "observed-orientation-fit"
)

type SignedSourceAudit struct {
	Name                    string
	Kind                    SignedSourceKind
	Available               bool
	Finite                  bool
	ContactSide             bool
	MatterSide              bool
	Signed                  bool
	TraceZero               bool
	CoddProved              bool
	CBreaking               bool
	SelectsMathematicalSign bool
	SelectsPhysicalBranch   bool
	T3RSemantic             bool
	HyperchargeSemantic     bool
	RequiresOrientation     bool
	RequiresPullback        bool
	RequiresLocalFieldMap   bool
	RequiresObservedInput   bool
	BranchesRemaining       int
	Verdict                 string
}

type CenteredSpectralFunctional struct {
	Mean                  float64
	Values                []float64
	Trace                 float64
	PositiveRows          int
	NegativeRows          int
	ZeroRows              int
	MatchesLargestGap     bool
	CanonicalAsDiagnostic bool
	PhysicalCoddSource    bool
	Verdict               string
}

type BinarySplitAudit struct {
	HighRows                  int
	LowRows                   int
	HalfSignTraceHighPositive float64
	HalfSignTraceLowPositive  float64
	TracelessHighWeight       float64
	TracelessLowWeight        float64
	BinaryT3RAvailable        bool
	TracelessSignedAvailable  bool
	T3RSemantic               bool
	SelectedPhysicalBranch    bool
	Verdict                   string
}

type Summary struct {
	ContactRows                  int
	LargestGapHighRows           int
	LargestGapLowRows            int
	OrientationCandidates        int
	SignedSourcesAudited         int
	AvailableSignedDiagnostics   int
	TraceZeroDiagnostics         int
	CanonicalSignedDiagnostics   int
	CoddContactFunctionals       int
	CBreakingSources             int
	SourcesSelectingPhysicalSign int
	CenteredPositiveRows         int
	CenteredNegativeRows         int
	CenteredZeroRows             int
	BinaryT3RRowsDerived         int
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
	Previous contactasymmetry.Analysis

	SpectrumDescending []float64
	SplitPattern       string
	CenteredFunctional CenteredSpectralFunctional
	BinarySplit        BinarySplitAudit
	SourceAudits       []SignedSourceAudit
	Summary            Summary

	ContactRows                  int
	LargestGapHighRows           int
	LargestGapLowRows            int
	OrientationCandidates        int
	SignedSourcesAudited         int
	AvailableSignedDiagnostics   int
	TraceZeroDiagnostics         int
	CanonicalSignedDiagnostics   int
	CoddContactFunctionals       int
	CBreakingSources             int
	SourcesSelectingPhysicalSign int
	CenteredPositiveRows         int
	CenteredNegativeRows         int
	CenteredZeroRows             int
	BinaryT3RRowsDerived         int
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
		prev, err := contactasymmetry.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactasymmetry.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.LargestGapHighRows != 3 || prev.LargestGapLowRows != 4 || prev.OrientationCandidates != 2 {
		return Analysis{}, fmt.Errorf("Gate 144 requires Gate 143 closed-firewall 3|4 split with two orientations")
	}
	if !prev.Z2OrientationDegeneracy || prev.ChargeConjugationBroken || prev.CBreakingSources != 0 || prev.CoddContactFunctionals != 0 || prev.SourcesSelectingOrientation != 0 {
		return Analysis{}, fmt.Errorf("Gate 144 requires Gate 143 unbroken C-degeneracy with no C-odd contact source")
	}
	if prev.T3RPullbackRowsDerived != 0 || prev.ChiralityPullbackRowsDerived != 0 || prev.BMinusLPullbackRowsDerived != 0 || prev.SU2LPullbackRowsDerived != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 144 requires no contact charge pullbacks, hypercharge rows, or beta rows")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 144 refuses hidden observed physical input")
	}

	spectrum := append([]float64(nil), prev.SpectrumDescending...)
	mu := mean(spectrum)
	centered := make([]float64, len(spectrum))
	trace := 0.0
	positive, negative, zero := 0, 0, 0
	const eps = 1e-10
	for i, v := range spectrum {
		c := v - mu
		centered[i] = c
		trace += c
		switch {
		case c > eps:
			positive++
		case c < -eps:
			negative++
		default:
			zero++
		}
	}

	centeredFunctional := CenteredSpectralFunctional{
		Mean:                  mu,
		Values:                centered,
		Trace:                 trace,
		PositiveRows:          positive,
		NegativeRows:          negative,
		ZeroRows:              zero,
		MatchesLargestGap:     positive == prev.LargestGapHighRows && negative == prev.LargestGapLowRows,
		CanonicalAsDiagnostic: true,
		PhysicalCoddSource:    false,
		Verdict:               "J = D_contact - mean(D_contact) I is a canonical trace-zero signed contact diagnostic, but not a proven C-odd source current or T3R/hypercharge operator",
	}

	high := prev.LargestGapHighRows
	low := prev.LargestGapLowRows
	binary := BinarySplitAudit{
		HighRows:                  high,
		LowRows:                   low,
		HalfSignTraceHighPositive: float64(high-low) / 2.0,
		HalfSignTraceLowPositive:  float64(low-high) / 2.0,
		TracelessHighWeight:       float64(low) / float64(high+low),
		TracelessLowWeight:        -float64(high) / float64(high+low),
		BinaryT3RAvailable:        false,
		TracelessSignedAvailable:  true,
		T3RSemantic:               false,
		SelectedPhysicalBranch:    false,
		Verdict:                   "the 3|4 split admits signed/traceless diagnostics, but a pure ±1/2 T3R operator is non-traceless and no physical sign branch is selected",
	}

	audits := []SignedSourceAudit{
		{
			Name:                    "centered contact spectral functional J = D - mean(D)I",
			Kind:                    SourceCenteredSpectrum,
			Available:               true,
			Finite:                  true,
			ContactSide:             true,
			Signed:                  true,
			TraceZero:               true,
			SelectsMathematicalSign: true,
			SelectsPhysicalBranch:   false,
			BranchesRemaining:       2,
			Verdict:                 centeredFunctional.Verdict,
		},
		{
			Name:                    "binary largest-gap sign split",
			Kind:                    SourceBinaryGapSplit,
			Available:               true,
			Finite:                  true,
			ContactSide:             true,
			Signed:                  true,
			TraceZero:               false,
			SelectsMathematicalSign: false,
			SelectsPhysicalBranch:   false,
			RequiresOrientation:     true,
			BranchesRemaining:       2,
			Verdict:                 "uses the canonical 3|4 split but still has two signs and its ±1/2 trace is nonzero",
		},
		{
			Name:                    "traceless balanced 3|4 split functional",
			Kind:                    SourceTracelessGapSplit,
			Available:               true,
			Finite:                  true,
			ContactSide:             true,
			Signed:                  true,
			TraceZero:               true,
			SelectsMathematicalSign: false,
			SelectsPhysicalBranch:   false,
			RequiresOrientation:     true,
			BranchesRemaining:       2,
			Verdict:                 "a trace-zero split functional exists, but both signs are equivalent without a C-breaking source and the weights are not T3R charges",
		},
		{
			Name:                  "C-odd contact source current",
			Kind:                  SourceChargeOddCurrent,
			Available:             false,
			Finite:                true,
			ContactSide:           true,
			Signed:                true,
			CoddProved:            false,
			CBreaking:             false,
			SelectsPhysicalBranch: false,
			BranchesRemaining:     2,
			Verdict:               "no finite C-odd source current with a coupling functional has been derived",
		},
		{
			Name:                  "B-L/T3R/hypercharge pullback signed current",
			Kind:                  SourcePullbackCharge,
			Available:             false,
			Finite:                true,
			MatterSide:            true,
			Signed:                true,
			RequiresPullback:      true,
			SelectsPhysicalBranch: false,
			BranchesRemaining:     2,
			Verdict:               "matter-side charges do not yet pull back to the contact carrier",
		},
		{
			Name:                  "local field pole/residue signed current",
			Kind:                  SourceLocalPoleResidue,
			Available:             false,
			Finite:                true,
			ContactSide:           true,
			Signed:                true,
			RequiresLocalFieldMap: true,
			SelectsPhysicalBranch: false,
			BranchesRemaining:     2,
			Verdict:               "no Lorentz local field, pole residue, mass activation, or decoupling sign source is available",
		},
		{
			Name:                  "observed orientation / low-energy fit",
			Kind:                  SourceObservedOrientation,
			Available:             false,
			Finite:                false,
			RequiresObservedInput: true,
			SelectsPhysicalBranch: false,
			BranchesRemaining:     2,
			Verdict:               "forbidden: observed constants cannot orient the finite contact split",
		},
	}

	availableSigned := count(audits, func(a SignedSourceAudit) bool { return a.Available && a.Signed })
	traceZero := count(audits, func(a SignedSourceAudit) bool { return a.Available && a.TraceZero })
	canonicalDiagnostics := count(audits, func(a SignedSourceAudit) bool {
		return a.Available && a.ContactSide && a.Signed && !a.RequiresOrientation
	})
	codd := count(audits, func(a SignedSourceAudit) bool { return a.Available && a.CoddProved })
	breaking := count(audits, func(a SignedSourceAudit) bool { return a.Available && a.CBreaking })
	selecting := count(audits, func(a SignedSourceAudit) bool { return a.SelectsPhysicalBranch })

	summary := Summary{
		ContactRows:                  prev.ContactRows,
		LargestGapHighRows:           prev.LargestGapHighRows,
		LargestGapLowRows:            prev.LargestGapLowRows,
		OrientationCandidates:        prev.OrientationCandidates,
		SignedSourcesAudited:         len(audits),
		AvailableSignedDiagnostics:   availableSigned,
		TraceZeroDiagnostics:         traceZero,
		CanonicalSignedDiagnostics:   canonicalDiagnostics,
		CoddContactFunctionals:       codd,
		CBreakingSources:             breaking,
		SourcesSelectingPhysicalSign: selecting,
		CenteredPositiveRows:         positive,
		CenteredNegativeRows:         negative,
		CenteredZeroRows:             zero,
		BinaryT3RRowsDerived:         0,
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

	truth := "Gate 144 constructs the strongest signed object currently available on the contact carrier: the centered spectral functional J = D_contact - mean(D_contact) I. It is finite, canonical as a diagnostic, trace-zero, and signed with a 3|4 pattern matching the largest-gap split. But it is not a proven C-odd source current, not a T3R operator, not a hypercharge row, and not a local-field residue or decoupling rule. Therefore it cannot orient the physical branch or open contact beta matching."

	return Analysis{
		Previous:                     prev,
		SpectrumDescending:           spectrum,
		SplitPattern:                 prev.SplitPattern,
		CenteredFunctional:           centeredFunctional,
		BinarySplit:                  binary,
		SourceAudits:                 audits,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		LargestGapHighRows:           prev.LargestGapHighRows,
		LargestGapLowRows:            prev.LargestGapLowRows,
		OrientationCandidates:        prev.OrientationCandidates,
		SignedSourcesAudited:         len(audits),
		AvailableSignedDiagnostics:   availableSigned,
		TraceZeroDiagnostics:         traceZero,
		CanonicalSignedDiagnostics:   canonicalDiagnostics,
		CoddContactFunctionals:       codd,
		CBreakingSources:             breaking,
		SourcesSelectingPhysicalSign: selecting,
		CenteredPositiveRows:         positive,
		CenteredNegativeRows:         negative,
		CenteredZeroRows:             zero,
		BinaryT3RRowsDerived:         0,
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
			"the centered spectral functional is a physical C-odd source",
			"the centered spectral functional is contact T3R",
			"the 3|4 binary split gives a traceless ±1/2 generator",
			"a trace-zero split diagnostic supplies hypercharge rows",
			"signed contact diagnostics permit threshold beta rows",
		},
		RemainingUnknowns: []string{
			"physical C-odd source current or signed coupling, if any",
			"Fock-contact pullback for B-L, chirality, T3R, and SU(2)L",
			"contact local field map, pole residues, mass activation, and decoupling",
			"contact hypercharge rows and threshold beta tensor",
		},
		RecommendedNextGate: "Gate 145 — centered contact spectral current / charge-operator normalization obstruction theorem",
	}, nil
}

func FormatCenteredFunctional(c CenteredSpectralFunctional) string {
	return fmt.Sprintf("mean=%.10f trace=%.3e pos=%d neg=%d zero=%d matchesGap=%t canonicalDiagnostic=%t physicalCodd=%t values=[%s] (%s)", c.Mean, c.Trace, c.PositiveRows, c.NegativeRows, c.ZeroRows, c.MatchesLargestGap, c.CanonicalAsDiagnostic, c.PhysicalCoddSource, formatFloats(c.Values), c.Verdict)
}

func FormatBinarySplit(b BinarySplitAudit) string {
	return fmt.Sprintf("split=%d|%d ±1/2 traces=(high+=%.1f low+=%.1f) tracelessWeights=(high=%.10f low=%.10f) binaryT3R=%t tracelessSigned=%t T3Rsemantic=%t selected=%t (%s)", b.HighRows, b.LowRows, b.HalfSignTraceHighPositive, b.HalfSignTraceLowPositive, b.TracelessHighWeight, b.TracelessLowWeight, b.BinaryT3RAvailable, b.TracelessSignedAvailable, b.T3RSemantic, b.SelectedPhysicalBranch, b.Verdict)
}

func FormatSignedSources(items []SignedSourceAudit) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s available=%t signed=%t traceZero=%t Codd=%t Cbreak=%t selectsPhysical=%t (%s)", item.Name, item.Available, item.Signed, item.TraceZero, item.CoddProved, item.CBreaking, item.SelectsPhysicalBranch, item.Verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d split=%d|%d signedSources=%d availableSigned=%d traceZero=%d canonicalDiag=%d Codd=%d Cbreak=%d selected=%d centered=%d/%d/%d T3R=%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.LargestGapHighRows, s.LargestGapLowRows, s.SignedSourcesAudited, s.AvailableSignedDiagnostics, s.TraceZeroDiagnostics, s.CanonicalSignedDiagnostics, s.CoddContactFunctionals, s.CBreakingSources, s.SourcesSelectingPhysicalSign, s.CenteredPositiveRows, s.CenteredNegativeRows, s.CenteredZeroRows, s.T3RPullbackRowsDerived, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
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

func formatFloats(values []float64) string {
	parts := make([]string, 0, len(values))
	for _, v := range values {
		if math.Abs(v) < 1e-12 {
			v = 0
		}
		parts = append(parts, fmt.Sprintf("%.10f", v))
	}
	return strings.Join(parts, ", ")
}
