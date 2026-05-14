// Package conditionalrgbranch implements Gate 176: conditional RG boundary-scale
// solvability under the quarantined u=1 instanton branch.
//
// Gate 175 preserved the branch u=1/g_*^2=1 only as a conditional matching rule,
// not as a strict finite-to-continuum theorem. Gate 176 asks the deliberately
// quarantined phenomenological question: if that branch is assumed, do the
// one-loop boundary equations produce a sensible low-energy electroweak/QCD
// point near M_Z? It also records a ratio-only running check in GUT-normalized
// variables that is independent of the absolute intercept u.
//
// Observed numbers enter this package only as an explicitly marked comparison
// ledger. They are not used to select a theorem, not used to tune thresholds,
// and not allowed to reduce strict nullity.
package conditionalrgbranch

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/instantontracebridge"
)

type InputAudit struct {
	Gate175ConditionalBranchAvailable bool
	Gate175StrictAbsoluteUDerived     bool
	RelativeGaugeRatioClosed          bool
	WeakAngleSeedClosed               bool
	ConditionalUInverseGStar          float64
	KY                                float64
	B1                                float64
	B2                                float64
	B3                                float64
	UsesObservedInputForDerivation    bool
	ObservedComparisonQuarantined     bool
	Verdict                           string
}

type ObservedReference struct {
	Name                  string
	ScaleGeV              float64
	AlphaEMInverse        float64
	Sin2MSbar             float64
	AlphaS                float64
	Alpha2Inverse         float64
	AlphaYInverse         float64
	Alpha1GUTInverse      float64
	Alpha3Inverse         float64
	Source                string
	UsedForDerivation     bool
	QuarantinedComparison bool
}

type RunningPoint struct {
	Name               string
	LogIntervalL       float64
	MStarOverMu        float64
	MStarGeV           float64
	A1GUTInverseG2     float64
	AYInverseG2        float64
	A2InverseG2        float64
	A3InverseG2        float64
	Alpha1GUTInverse   float64
	AlphaYInverse      float64
	Alpha2Inverse      float64
	Alpha3Inverse      float64
	AlphaEMInverse     float64
	Sin2               float64
	PositiveKinetic    bool
	PhysicalScaleOrder bool
	Detail             string
}

type FitResult struct {
	Name              string
	TargetObservable  string
	LogIntervalL      float64
	MStarGeV          float64
	FeasiblePositive  bool
	FeasibleLPositive bool
	Point             RunningPoint
	AlphaEMResidual   float64
	Sin2Residual      float64
	Alpha2Residual    float64
	Alpha3Residual    float64
	PassesLooseRange  bool
	Verdict           string
}

type RatioAudit struct {
	Formula                 string
	TheoryRatio             float64
	ObservedRatio           float64
	ObservedL23             float64
	ObservedL12             float64
	RelativeMismatch        float64
	LIntervalsAgree         bool
	IndependentOfU          bool
	UsesGUTNormalizedAlpha1 bool
	ObservedComparisonOnly  bool
	Verdict                 string
}

type ConditionalBranchFirewall struct {
	ConditionalUAssumed                  bool
	StrictUStillOpen                     bool
	OneLoopContinuumAssumed              bool
	ThresholdCorrectionsIncluded         bool
	ObservedComparisonQuarantined        bool
	AnyObservedFitViable                 bool
	RatioCheckPasses                     bool
	ConditionalBranchRejectedByMZ        bool
	StrictNullityBefore                  int
	StrictNullityAfter                   int
	ConditionalNullityBefore             int
	ConditionalNullityAfter              int
	PhysicalConstantsDerived             bool
	BoundaryScaleDerivedStrict           bool
	BoundaryScaleDerivedConditional      bool
	HiddenObservedInputUsedForDerivation bool
	RemainingStrictUnknowns              []string
	RecommendedNextGate                  string
	Verdict                              string
}

type Analysis struct {
	Previous       instantontracebridge.Analysis
	Input          InputAudit
	Observed       ObservedReference
	BoundaryPoint  RunningPoint
	SamplePoints   []RunningPoint
	Fits           []FitResult
	Ratio          RatioAudit
	Firewall       ConditionalBranchFirewall
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := instantontracebridge.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev instantontracebridge.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.Firewall.ConditionalAbsoluteUPreserved || prev.Firewall.ConditionalNullityAfter != 2 {
		return Analysis{}, fmt.Errorf("Gate 176 requires the Gate 175 quarantined u=1 conditional branch")
	}
	if prev.Firewall.StrictAbsoluteUDerived || prev.Firewall.StrictNullityAfter != 3 {
		return Analysis{}, fmt.Errorf("Gate 176 requires strict absolute coupling to remain open")
	}
	if !prev.Input.RelativeGaugeRatioClosed || !prev.Input.WeakAngleSeedClosed {
		return Analysis{}, fmt.Errorf("Gate 176 requires the closed Gate 167 relative gauge ratio")
	}

	const (
		ky = 5.0 / 3.0
		b1 = 41.0 / 10.0
		b2 = -19.0 / 6.0
		b3 = -7.0
		u  = 1.0
	)
	obs := buildObservedReference()
	input := InputAudit{
		Gate175ConditionalBranchAvailable: prev.Firewall.ConditionalAbsoluteUPreserved,
		Gate175StrictAbsoluteUDerived:     prev.Firewall.StrictAbsoluteUDerived,
		RelativeGaugeRatioClosed:          prev.Input.RelativeGaugeRatioClosed,
		WeakAngleSeedClosed:               prev.Input.WeakAngleSeedClosed,
		ConditionalUInverseGStar:          prev.Input.ConditionalUInverseGStar,
		KY:                                ky,
		B1:                                b1,
		B2:                                b2,
		B3:                                b3,
		UsesObservedInputForDerivation:    false,
		ObservedComparisonQuarantined:     true,
		Verdict:                           "conditional RG branch may be evaluated, but strict normalization and observed comparison remain quarantined",
	}
	if math.Abs(prev.Input.ConditionalUInverseGStar-u) > eps {
		return Analysis{}, fmt.Errorf("Gate 176 expected inherited conditional u=1, got %.12g", prev.Input.ConditionalUInverseGStar)
	}

	boundary := runningPoint("conditional boundary witness", ky, b1, b2, b3, u, 0, obs.ScaleGeV)
	samples := []RunningPoint{
		boundary,
		runningPoint("alpha3-fit witness", ky, b1, b2, b3, u, fitLForAlpha3(b3, u, obs.Alpha3Inverse), obs.ScaleGeV),
		runningPoint("weak-angle-fit witness", ky, b1, b2, b3, u, fitLForSin2(ky, b1, b2, u, obs.Sin2MSbar), obs.ScaleGeV),
	}
	fits := []FitResult{
		buildFit("fit alpha3 only", "α3(M_Z)", ky, b1, b2, b3, u, fitLForAlpha3(b3, u, obs.Alpha3Inverse), obs, eps),
		buildFit("fit alpha2 only", "α2(M_Z)", ky, b1, b2, b3, u, fitLForAlpha2(b2, u, obs.Alpha2Inverse), obs, eps),
		buildFit("fit electromagnetic alpha only", "α_em(M_Z)", ky, b1, b2, b3, u, fitLForAlphaEM(ky, b1, b2, u, obs.AlphaEMInverse), obs, eps),
		buildFit("fit weak angle only", "sin²θ(M_Z)", ky, b1, b2, b3, u, fitLForSin2(ky, b1, b2, u, obs.Sin2MSbar), obs, eps),
	}
	ratio := buildRatioAudit(b1, b2, b3, obs, eps)
	anyViable := false
	for _, f := range fits {
		if f.PassesLooseRange {
			anyViable = true
			break
		}
	}
	firewall := ConditionalBranchFirewall{
		ConditionalUAssumed:                  true,
		StrictUStillOpen:                     !prev.Firewall.StrictAbsoluteUDerived,
		OneLoopContinuumAssumed:              true,
		ThresholdCorrectionsIncluded:         false,
		ObservedComparisonQuarantined:        true,
		AnyObservedFitViable:                 anyViable,
		RatioCheckPasses:                     ratio.LIntervalsAgree,
		ConditionalBranchRejectedByMZ:        !anyViable,
		StrictNullityBefore:                  3,
		StrictNullityAfter:                   3,
		ConditionalNullityBefore:             2,
		ConditionalNullityAfter:              2,
		PhysicalConstantsDerived:             false,
		BoundaryScaleDerivedStrict:           false,
		BoundaryScaleDerivedConditional:      false,
		HiddenObservedInputUsedForDerivation: false,
		RemainingStrictUnknowns: []string{
			"u=1/g_*²: still strict-open because Gate 175 did not derive the finite-to-continuum normalization bridge",
			"L=ln(M*/M_Z): no finite scale selector; observed comparisons are diagnostics only",
			"Δb_i: threshold activation/decoupling remains open and is not fitted here",
		},
		RecommendedNextGate: "Gate 177 — normalization-prefactor or threshold-deformation branch audit after conditional u=1 M_Z rejection",
		Verdict:             "the quarantined u=1 branch is internally computable but does not land near the observed M_Z coupling pattern under unthresholded one-loop running",
	}

	return Analysis{
		Previous:       prev,
		Input:          input,
		Observed:       obs,
		BoundaryPoint:  boundary,
		SamplePoints:   samples,
		Fits:           fits,
		Ratio:          ratio,
		Firewall:       firewall,
		TruthStatement: "Gate 176 shows that the conditional instanton branch u=1 can be propagated through the one-loop RG equations, but it does not produce a viable M_Z coupling pattern without additional normalization or threshold data. Fitting α3 gives a positive log interval but predicts α2 and α_em far too strong; fitting α2 requires a negative log interval; fitting α_em destroys positivity; fitting sin² leaves α3 far off. The GUT-normalized ratio audit also fails against the comparison ledger, so the branch remains quarantined and strict nullity does not reduce.",
	}, nil
}

func buildObservedReference() ObservedReference {
	// Comparison values are ordinary precision-electroweak reference values near
	// M_Z. They are intentionally hard-coded here as a quarantined external audit
	// vector, not as finite engine input.
	const (
		mz       = 91.1876
		alphaInv = 127.955
		sin2     = 0.23122
		alphaS   = 0.1179
	)
	alpha2Inv := sin2 * alphaInv
	alphaYInv := (1 - sin2) * alphaInv
	alpha1Inv := (3.0 / 5.0) * alphaYInv
	alpha3Inv := 1 / alphaS
	return ObservedReference{
		Name:                  "M_Z comparison ledger",
		ScaleGeV:              mz,
		AlphaEMInverse:        alphaInv,
		Sin2MSbar:             sin2,
		AlphaS:                alphaS,
		Alpha2Inverse:         alpha2Inv,
		AlphaYInverse:         alphaYInv,
		Alpha1GUTInverse:      alpha1Inv,
		Alpha3Inverse:         alpha3Inv,
		Source:                "PDG-style M_Z electroweak/QCD comparison values; used only for external viability audit",
		UsedForDerivation:     false,
		QuarantinedComparison: true,
	}
}

func runningPoint(name string, ky, b1, b2, b3, u, L, muGeV float64) RunningPoint {
	s1 := b1 / (8 * math.Pi * math.Pi)
	s2 := b2 / (8 * math.Pi * math.Pi)
	s3 := b3 / (8 * math.Pi * math.Pi)
	a1 := u + s1*L
	aY := ky*u + s1*L
	a2 := u + s2*L
	a3 := u + s3*L
	emInv := 4 * math.Pi * (aY + a2)
	sin2 := math.NaN()
	if math.Abs(aY+a2) > 0 {
		sin2 = a2 / (aY + a2)
	}
	ratio := math.Exp(L)
	mstar := muGeV * ratio
	return RunningPoint{
		Name:               name,
		LogIntervalL:       L,
		MStarOverMu:        ratio,
		MStarGeV:           mstar,
		A1GUTInverseG2:     a1,
		AYInverseG2:        aY,
		A2InverseG2:        a2,
		A3InverseG2:        a3,
		Alpha1GUTInverse:   4 * math.Pi * a1,
		AlphaYInverse:      4 * math.Pi * aY,
		Alpha2Inverse:      4 * math.Pi * a2,
		Alpha3Inverse:      4 * math.Pi * a3,
		AlphaEMInverse:     emInv,
		Sin2:               sin2,
		PositiveKinetic:    aY > 0 && a2 > 0 && a3 > 0,
		PhysicalScaleOrder: L >= 0,
		Detail:             "conditional u=1, unthresholded one-loop flow; diagnostic only",
	}
}

func buildFit(name, obsName string, ky, b1, b2, b3, u, L float64, obs ObservedReference, eps float64) FitResult {
	p := runningPoint(name, ky, b1, b2, b3, u, L, obs.ScaleGeV)
	alphaEMResidual := relResidual(p.AlphaEMInverse, obs.AlphaEMInverse)
	sin2Residual := relResidual(p.Sin2, obs.Sin2MSbar)
	alpha2Residual := relResidual(p.Alpha2Inverse, obs.Alpha2Inverse)
	alpha3Residual := relResidual(p.Alpha3Inverse, obs.Alpha3Inverse)
	// A deliberately loose viability window. The branch should not need exact
	// precision to survive a first audit, but it must at least land in the same
	// broad coupling regime without negative kinetic terms or a reversed scale.
	pass := p.PositiveKinetic && p.PhysicalScaleOrder &&
		math.Abs(alphaEMResidual) < 0.15 &&
		math.Abs(sin2Residual) < 0.15 &&
		math.Abs(alpha2Residual) < 0.15 &&
		math.Abs(alpha3Residual) < 0.15
	verdict := "not viable as a simultaneous M_Z fit"
	if pass {
		verdict = "viable broad-range simultaneous M_Z fit"
	}
	if !p.PositiveKinetic {
		verdict = "rejected: running point has negative kinetic coefficient"
	} else if !p.PhysicalScaleOrder {
		verdict = "rejected: requires M* below the comparison scale"
	}
	return FitResult{
		Name:              name,
		TargetObservable:  obsName,
		LogIntervalL:      L,
		MStarGeV:          p.MStarGeV,
		FeasiblePositive:  p.PositiveKinetic,
		FeasibleLPositive: p.PhysicalScaleOrder,
		Point:             p,
		AlphaEMResidual:   alphaEMResidual,
		Sin2Residual:      sin2Residual,
		Alpha2Residual:    alpha2Residual,
		Alpha3Residual:    alpha3Residual,
		PassesLooseRange:  pass,
		Verdict:           verdict,
	}
}

func buildRatioAudit(b1, b2, b3 float64, obs ObservedReference, eps float64) RatioAudit {
	theory := (b2 - b3) / (b1 - b2)
	obsNum := obs.Alpha2Inverse - obs.Alpha3Inverse
	obsDen := obs.Alpha1GUTInverse - obs.Alpha2Inverse
	observed := obsNum / obsDen
	L23 := (2 * math.Pi * obsNum) / (b2 - b3)
	L12 := (2 * math.Pi * obsDen) / (b1 - b2)
	mismatch := relResidual(theory, observed)
	agree := math.Abs(L23-L12)/math.Max(1, math.Abs(L12)) < 0.15
	return RatioAudit{
		Formula:                 "(α₂⁻¹-α₃⁻¹)/(α₁⁻¹-α₂⁻¹) = (b₂-b₃)/(b₁-b₂) in GUT-normalized α₁ variables",
		TheoryRatio:             theory,
		ObservedRatio:           observed,
		ObservedL23:             L23,
		ObservedL12:             L12,
		RelativeMismatch:        mismatch,
		LIntervalsAgree:         agree && math.Abs(mismatch) < 0.15,
		IndependentOfU:          true,
		UsesGUTNormalizedAlpha1: true,
		ObservedComparisonOnly:  true,
		Verdict:                 "ratio-only check fails for the unthresholded one-loop SM beta vector; thresholds or new normalization data would be required",
	}
}

func fitLForAlpha3(b3, u, alpha3Inv float64) float64 {
	s3 := b3 / (8 * math.Pi * math.Pi)
	return (alpha3Inv/(4*math.Pi) - u) / s3
}

func fitLForAlpha2(b2, u, alpha2Inv float64) float64 {
	s2 := b2 / (8 * math.Pi * math.Pi)
	return (alpha2Inv/(4*math.Pi) - u) / s2
}

func fitLForAlphaEM(ky, b1, b2, u, alphaEMInv float64) float64 {
	s := (b1 + b2) / (8 * math.Pi * math.Pi)
	return (alphaEMInv/(4*math.Pi) - (ky+1)*u) / s
}

func fitLForSin2(ky, b1, b2, u, sin2 float64) float64 {
	s1 := b1 / (8 * math.Pi * math.Pi)
	s2 := b2 / (8 * math.Pi * math.Pi)
	return (sin2*(ky+1)*u - u) / (s2 - sin2*(s1+s2))
}

func relResidual(predicted, observed float64) float64 {
	if observed == 0 || math.IsNaN(predicted) || math.IsInf(predicted, 0) {
		return math.Inf(1)
	}
	return (predicted - observed) / observed
}

func FormatInput(x InputAudit) string {
	return fmt.Sprintf("conditional=%t strict-u=%t ratio=%t sin2=%t u=%.12g ky=%.12g b=(%.12g,%.12g,%.12g) observed-derivation=%t comparison-quarantined=%t verdict=%s", x.Gate175ConditionalBranchAvailable, x.Gate175StrictAbsoluteUDerived, x.RelativeGaugeRatioClosed, x.WeakAngleSeedClosed, x.ConditionalUInverseGStar, x.KY, x.B1, x.B2, x.B3, x.UsesObservedInputForDerivation, x.ObservedComparisonQuarantined, x.Verdict)
}

func FormatObserved(x ObservedReference) string {
	return fmt.Sprintf("%s: MZ=%.6g GeV, αem⁻¹=%.6g, sin²=%.6g, αs=%.6g, α2⁻¹=%.6g, α1_GUT⁻¹=%.6g, α3⁻¹=%.6g, derivation=%t, quarantined=%t", x.Name, x.ScaleGeV, x.AlphaEMInverse, x.Sin2MSbar, x.AlphaS, x.Alpha2Inverse, x.Alpha1GUTInverse, x.Alpha3Inverse, x.UsedForDerivation, x.QuarantinedComparison)
}

func FormatPoint(p RunningPoint) string {
	return fmt.Sprintf("%s: L=%.9g, M*/MZ=%.9g, M*=%.9g GeV, invα=(α1GUT %.6g, αY %.6g, α2 %.6g, α3 %.6g, αem %.6g), sin²=%.6g, positive=%t, L>=0=%t", p.Name, p.LogIntervalL, p.MStarOverMu, p.MStarGeV, p.Alpha1GUTInverse, p.AlphaYInverse, p.Alpha2Inverse, p.Alpha3Inverse, p.AlphaEMInverse, p.Sin2, p.PositiveKinetic, p.PhysicalScaleOrder)
}

func FormatPoints(xs []RunningPoint) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatPoint(x))
	}
	return strings.Join(parts, " | ")
}

func FormatFit(f FitResult) string {
	return fmt.Sprintf("%s targeting %s: L=%.9g, M*=%.9g GeV, positive=%t L>=0=%t residuals=(αem %.3g, sin² %.3g, α2 %.3g, α3 %.3g), loose-pass=%t verdict=%s", f.Name, f.TargetObservable, f.LogIntervalL, f.MStarGeV, f.FeasiblePositive, f.FeasibleLPositive, f.AlphaEMResidual, f.Sin2Residual, f.Alpha2Residual, f.Alpha3Residual, f.PassesLooseRange, f.Verdict)
}

func FormatFits(xs []FitResult) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatFit(x))
	}
	return strings.Join(parts, " | ")
}

func FormatRatio(x RatioAudit) string {
	return fmt.Sprintf("%s: theory=%.9g observed=%.9g mismatch=%.3g L23=%.9g L12=%.9g agree=%t independent-u=%t comparison-only=%t verdict=%s", x.Formula, x.TheoryRatio, x.ObservedRatio, x.RelativeMismatch, x.ObservedL23, x.ObservedL12, x.LIntervalsAgree, x.IndependentOfU, x.ObservedComparisonOnly, x.Verdict)
}

func FormatFirewall(x ConditionalBranchFirewall) string {
	return fmt.Sprintf("u-assumed=%t strict-u-open=%t one-loop=%t thresholds=%t observed-quarantined=%t any-fit=%t ratio-pass=%t rejected=%t strict-nullity=%d->%d conditional-nullity=%d->%d physical-derived=%t hidden-observed-derivation=%t next=%s verdict=%s", x.ConditionalUAssumed, x.StrictUStillOpen, x.OneLoopContinuumAssumed, x.ThresholdCorrectionsIncluded, x.ObservedComparisonQuarantined, x.AnyObservedFitViable, x.RatioCheckPasses, x.ConditionalBranchRejectedByMZ, x.StrictNullityBefore, x.StrictNullityAfter, x.ConditionalNullityBefore, x.ConditionalNullityAfter, x.PhysicalConstantsDerived, x.HiddenObservedInputUsedForDerivation, x.RecommendedNextGate, x.Verdict)
}
