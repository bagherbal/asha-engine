// Package contactchargenorm implements Gate 145: centered contact spectral
// current / charge-operator normalization obstruction theorem.
//
// Gate 144 constructed a canonical trace-zero signed diagnostic
// J = D_contact - mean(D_contact) I. Gate 145 asks the next sharper question:
// can this diagnostic be normalized into a physical charge operator such as
// T3R, B-L, hypercharge, or a threshold beta representation row?
//
// The answer is no. Several canonical normalizations exist (max-absolute,
// Frobenius, range, centered spectral units), and the 3|4 split admits binary
// or balanced two-level summaries. But each candidate fails at least one
// required charge-operator condition: selected physical sign, rational/finite
// charge semantics, pullback to B-L/T3R/chirality/SU2L, local field variables,
// mass activation, and decoupling. The centered current remains diagnostic only.
package contactchargenorm

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactcoddsource"
)

type NormalizationKind string

const (
	NormCenteredRaw       NormalizationKind = "centered-raw"
	NormMaxAbs            NormalizationKind = "max-absolute"
	NormFrobenius         NormalizationKind = "frobenius"
	NormRange             NormalizationKind = "range"
	NormBinaryHalfCharge  NormalizationKind = "binary-half-charge"
	NormBalancedTraceZero NormalizationKind = "balanced-trace-zero-split"
	NormObservedChargeFit NormalizationKind = "observed-charge-fit"
)

type NormalizationAudit struct {
	Name                   string
	Kind                   NormalizationKind
	Available              bool
	CanonicalAsDiagnostic  bool
	TraceZero              bool
	Signed                 bool
	Finite                 bool
	Eigenvalues            []float64
	DistinctEigenvalues    int
	PositiveRows           int
	NegativeRows           int
	ZeroRows               int
	TwoLevel               bool
	UniformMagnitude       bool
	T3RSemantic            bool
	BMinusLSemantic        bool
	HyperchargeSemantic    bool
	ChargeOperatorSemantic bool
	RequiresOrientation    bool
	RequiresPullback       bool
	RequiresLocalFieldMap  bool
	RequiresObservedInput  bool
	OpensBetaPermission    bool
	Verdict                string
}

type ChargeOperatorRequirements struct {
	TraceControlled         bool
	SelectedOrientation     bool
	FiniteChargeLattice     bool
	OperatorPullback        bool
	LocalFieldMap           bool
	GaugeRepresentationRows bool
	MassActivation          bool
	DecouplingRule          bool
	ObservedInputFree       bool
	AllSatisfied            bool
	Verdict                 string
}

type Summary struct {
	ContactRows                  int
	CenteredPositiveRows         int
	CenteredNegativeRows         int
	CenteredZeroRows             int
	NormalizationsAudited        int
	AvailableNormalizations      int
	CanonicalDiagnosticNorms     int
	TraceZeroNormalizations      int
	TwoLevelNormalizations       int
	ChargeSemanticNormalizations int
	T3RRowsDerived               int
	BMinusLRowsDerived           int
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
	Previous contactcoddsource.Analysis

	CenteredValues     []float64
	MaxAbs             float64
	FrobeniusNorm      float64
	SpectralRange      float64
	RawAudit           NormalizationAudit
	MaxAbsAudit        NormalizationAudit
	FrobeniusAudit     NormalizationAudit
	RangeAudit         NormalizationAudit
	BinaryHalfAudit    NormalizationAudit
	BalancedSplitAudit NormalizationAudit
	ObservedFitAudit   NormalizationAudit
	Audits             []NormalizationAudit
	Requirements       ChargeOperatorRequirements
	Summary            Summary

	ContactRows                  int
	CenteredPositiveRows         int
	CenteredNegativeRows         int
	CenteredZeroRows             int
	NormalizationsAudited        int
	AvailableNormalizations      int
	CanonicalDiagnosticNorms     int
	TraceZeroNormalizations      int
	TwoLevelNormalizations       int
	ChargeSemanticNormalizations int
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
		prev, err := contactcoddsource.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactcoddsource.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.CenteredPositiveRows != 3 || prev.CenteredNegativeRows != 4 || prev.CenteredZeroRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 145 requires Gate 144 closed-firewall centered 3|4 signed diagnostic")
	}
	if !prev.CenteredFunctional.CanonicalAsDiagnostic || prev.CenteredFunctional.PhysicalCoddSource || prev.CoddContactFunctionals != 0 || prev.CBreakingSources != 0 || prev.SourcesSelectingPhysicalSign != 0 {
		return Analysis{}, fmt.Errorf("Gate 145 requires Gate 144 canonical diagnostic but no physical C-odd source")
	}
	if prev.T3RPullbackRowsDerived != 0 || prev.BMinusLPullbackRowsDerived != 0 || prev.HyperchargeRowsDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 145 requires no contact charge rows or beta rows")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 145 refuses hidden observed physical input")
	}

	centered := append([]float64(nil), prev.CenteredFunctional.Values...)
	maxAbs := maxAbs(centered)
	frob := frobenius(centered)
	rng := max(centered) - min(centered)

	rawVals := append([]float64(nil), centered...)
	maxVals := scale(centered, reciprocal(maxAbs))
	frobVals := scale(centered, reciprocal(frob))
	rangeVals := scale(centered, reciprocal(rng))
	binaryVals := binaryHalf(prev.CenteredPositiveRows, prev.CenteredNegativeRows)
	balancedVals := balancedSplit(prev.CenteredPositiveRows, prev.CenteredNegativeRows)

	raw := makeAudit("raw centered spectral current J", NormCenteredRaw, true, true, true, rawVals, false, false, false, false, false, false, false, "canonical trace-zero signed diagnostic, but seven unequal eigenvalues and no charge semantics")
	maxAudit := makeAudit("max-absolute normalized centered current", NormMaxAbs, true, true, true, maxVals, false, false, false, false, false, false, false, "canonical diagnostic scale with max |eigenvalue| = 1, but no finite charge lattice or physical current coupling")
	frobAudit := makeAudit("Frobenius normalized centered current", NormFrobenius, true, true, true, frobVals, false, false, false, false, false, false, false, "canonical diagnostic scale with Tr(J^2)=1, but no T3R/B-L/hypercharge semantics")
	rangeAudit := makeAudit("range normalized centered current", NormRange, true, true, true, rangeVals, false, false, false, false, false, false, false, "canonical diagnostic scale by spectral range, but still a real spectral diagnostic rather than a charge operator")
	binaryAudit := makeAudit("binary ±1/2 largest-gap split", NormBinaryHalfCharge, true, true, false, binaryVals, true, true, false, false, false, false, false, "has the desired ±1/2 magnitudes but is non-trace-zero on a 3|4 split and still lacks orientation and charge semantics")
	balancedAudit := makeAudit("balanced trace-zero 3|4 split", NormBalancedTraceZero, true, true, true, balancedVals, true, false, false, false, false, false, false, "trace-zero two-level split exists with weights +4/7 and -3/7, but those are not T3R, B-L, or hypercharge eigenvalues")
	observedAudit := makeAudit("observed charge normalization fit", NormObservedChargeFit, false, false, false, nil, false, false, false, false, false, false, true, "forbidden: fitting normalization to observed charges or low-energy constants would bypass the theorem ladder")

	audits := []NormalizationAudit{raw, maxAudit, frobAudit, rangeAudit, binaryAudit, balancedAudit, observedAudit}
	reqs := ChargeOperatorRequirements{
		TraceControlled:         true,
		SelectedOrientation:     false,
		FiniteChargeLattice:     false,
		OperatorPullback:        false,
		LocalFieldMap:           false,
		GaugeRepresentationRows: false,
		MassActivation:          false,
		DecouplingRule:          false,
		ObservedInputFree:       true,
		AllSatisfied:            false,
		Verdict:                 "normalization alone cannot turn a canonical contact diagnostic into a physical charge operator; orientation, pullback, representation, locality, mass activation, and decoupling remain missing",
	}

	available := count(audits, func(a NormalizationAudit) bool { return a.Available })
	canonical := count(audits, func(a NormalizationAudit) bool { return a.Available && a.CanonicalAsDiagnostic })
	traceZero := count(audits, func(a NormalizationAudit) bool { return a.Available && a.TraceZero })
	twoLevel := count(audits, func(a NormalizationAudit) bool { return a.Available && a.TwoLevel })
	semantic := count(audits, func(a NormalizationAudit) bool { return a.Available && a.ChargeOperatorSemantic })

	summary := Summary{
		ContactRows:                  prev.ContactRows,
		CenteredPositiveRows:         prev.CenteredPositiveRows,
		CenteredNegativeRows:         prev.CenteredNegativeRows,
		CenteredZeroRows:             prev.CenteredZeroRows,
		NormalizationsAudited:        len(audits),
		AvailableNormalizations:      available,
		CanonicalDiagnosticNorms:     canonical,
		TraceZeroNormalizations:      traceZero,
		TwoLevelNormalizations:       twoLevel,
		ChargeSemanticNormalizations: semantic,
		T3RRowsDerived:               0,
		BMinusLRowsDerived:           0,
		HyperchargeRowsDerived:       0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.ContactRows,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
	}

	truth := "Gate 145 proves that the centered contact spectral current is canonical only as a finite diagnostic. Max-absolute, Frobenius, and range normalizations preserve the seven unequal spectral rows; binary ±1/2 normalization loses trace control on the 3|4 split; and balanced trace-zero normalization gives +4/7 and -3/7 rather than physical charge eigenvalues. No normalization supplies orientation, B-L/T3R/chirality pullback, local field variables, mass activation, decoupling, or threshold beta permission."

	return Analysis{
		Previous:                     prev,
		CenteredValues:               centered,
		MaxAbs:                       maxAbs,
		FrobeniusNorm:                frob,
		SpectralRange:                rng,
		RawAudit:                     raw,
		MaxAbsAudit:                  maxAudit,
		FrobeniusAudit:               frobAudit,
		RangeAudit:                   rangeAudit,
		BinaryHalfAudit:              binaryAudit,
		BalancedSplitAudit:           balancedAudit,
		ObservedFitAudit:             observedAudit,
		Audits:                       audits,
		Requirements:                 reqs,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		CenteredPositiveRows:         prev.CenteredPositiveRows,
		CenteredNegativeRows:         prev.CenteredNegativeRows,
		CenteredZeroRows:             prev.CenteredZeroRows,
		NormalizationsAudited:        len(audits),
		AvailableNormalizations:      available,
		CanonicalDiagnosticNorms:     canonical,
		TraceZeroNormalizations:      traceZero,
		TwoLevelNormalizations:       twoLevel,
		ChargeSemanticNormalizations: semantic,
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
			"normalizing the centered contact current derives T3R",
			"the ±1/2 binary split is a valid contact charge operator",
			"the balanced +4/7,-3/7 split is hypercharge",
			"spectral normalization supplies threshold beta rows",
			"observed constants may choose the normalization",
		},
		RemainingUnknowns: []string{
			"contact charge-operator semantics for centered spectral current, if any",
			"orientation and charge-conjugation-breaking source",
			"B-L, chirality, T3R, SU2L, and hypercharge pullbacks",
			"local field map, pole residues, mass activation, and decoupling",
		},
		RecommendedNextGate: "Gate 146 — contact charge lattice embedding / rational-spectrum obstruction theorem",
	}, nil
}

func makeAudit(name string, kind NormalizationKind, available, canonical, traceZero bool, values []float64, twoLevel, uniformMagnitude, t3r, bl, hyper, beta, observed bool, verdict string) NormalizationAudit {
	positive, negative, zero := signCounts(values)
	chargeSemantic := t3r || bl || hyper
	return NormalizationAudit{
		Name:                   name,
		Kind:                   kind,
		Available:              available,
		CanonicalAsDiagnostic:  canonical,
		TraceZero:              traceZero,
		Signed:                 positive > 0 && negative > 0,
		Finite:                 !observed,
		Eigenvalues:            append([]float64(nil), values...),
		DistinctEigenvalues:    distinct(values, 1e-9),
		PositiveRows:           positive,
		NegativeRows:           negative,
		ZeroRows:               zero,
		TwoLevel:               twoLevel,
		UniformMagnitude:       uniformMagnitude,
		T3RSemantic:            t3r,
		BMinusLSemantic:        bl,
		HyperchargeSemantic:    hyper,
		ChargeOperatorSemantic: chargeSemantic,
		RequiresOrientation:    kind == NormBinaryHalfCharge || kind == NormBalancedTraceZero,
		RequiresPullback:       chargeSemantic,
		RequiresLocalFieldMap:  beta,
		RequiresObservedInput:  observed,
		OpensBetaPermission:    beta,
		Verdict:                verdict,
	}
}

func FormatAudit(a NormalizationAudit) string {
	return fmt.Sprintf("%s available=%t canonicalDiag=%t traceZero=%t signed=%t distinct=%d twoLevel=%t uniformMag=%t chargeSemantic=%t beta=%t values=[%s] (%s)", a.Name, a.Available, a.CanonicalAsDiagnostic, a.TraceZero, a.Signed, a.DistinctEigenvalues, a.TwoLevel, a.UniformMagnitude, a.ChargeOperatorSemantic, a.OpensBetaPermission, formatFloats(a.Eigenvalues), a.Verdict)
}

func FormatAudits(items []NormalizationAudit) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, FormatAudit(item))
	}
	return strings.Join(parts, "; ")
}

func FormatRequirements(r ChargeOperatorRequirements) string {
	return fmt.Sprintf("trace=%t orientation=%t lattice=%t pullback=%t local=%t reps=%t mass=%t decoupling=%t observedFree=%t all=%t (%s)", r.TraceControlled, r.SelectedOrientation, r.FiniteChargeLattice, r.OperatorPullback, r.LocalFieldMap, r.GaugeRepresentationRows, r.MassActivation, r.DecouplingRule, r.ObservedInputFree, r.AllSatisfied, r.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d centered=%d/%d/%d norms=%d available=%d canonical=%d traceZero=%d twoLevel=%d semantic=%d T3R=%d B-L=%d Y=%d beta=%d nullity=%d→%d", s.ContactRows, s.CenteredPositiveRows, s.CenteredNegativeRows, s.CenteredZeroRows, s.NormalizationsAudited, s.AvailableNormalizations, s.CanonicalDiagnosticNorms, s.TraceZeroNormalizations, s.TwoLevelNormalizations, s.ChargeSemanticNormalizations, s.T3RRowsDerived, s.BMinusLRowsDerived, s.HyperchargeRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(items []string) string { return strings.Join(items, "; ") }

func signCounts(values []float64) (int, int, int) {
	positive, negative, zero := 0, 0, 0
	const eps = 1e-10
	for _, v := range values {
		switch {
		case v > eps:
			positive++
		case v < -eps:
			negative++
		default:
			zero++
		}
	}
	return positive, negative, zero
}

func distinct(values []float64, eps float64) int {
	if len(values) == 0 {
		return 0
	}
	unique := make([]float64, 0, len(values))
	for _, v := range values {
		seen := false
		for _, u := range unique {
			if math.Abs(v-u) <= eps {
				seen = true
				break
			}
		}
		if !seen {
			unique = append(unique, v)
		}
	}
	return len(unique)
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

func scale(values []float64, factor float64) []float64 {
	out := make([]float64, len(values))
	for i, v := range values {
		out[i] = factor * v
	}
	return out
}

func reciprocal(v float64) float64 {
	if math.Abs(v) < 1e-12 {
		return 0
	}
	return 1 / v
}

func maxAbs(values []float64) float64 {
	m := 0.0
	for _, v := range values {
		if math.Abs(v) > m {
			m = math.Abs(v)
		}
	}
	return m
}

func frobenius(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v * v
	}
	return math.Sqrt(sum)
}

func max(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func min(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func binaryHalf(pos, neg int) []float64 {
	out := make([]float64, 0, pos+neg)
	for i := 0; i < pos; i++ {
		out = append(out, 0.5)
	}
	for i := 0; i < neg; i++ {
		out = append(out, -0.5)
	}
	return out
}

func balancedSplit(pos, neg int) []float64 {
	total := pos + neg
	out := make([]float64, 0, total)
	if total == 0 || pos == 0 || neg == 0 {
		return out
	}
	high := float64(neg) / float64(total)
	low := -float64(pos) / float64(total)
	for i := 0; i < pos; i++ {
		out = append(out, high)
	}
	for i := 0; i < neg; i++ {
		out = append(out, low)
	}
	return out
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
