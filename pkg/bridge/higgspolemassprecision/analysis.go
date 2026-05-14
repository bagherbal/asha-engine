// Package higgspolemassprecision implements Gate 331:
// Higgs Pole-Mass Conversion / Precision Gap Ledger Audit.
//
// Gate 330 conditionally promoted the native doubled bosonic trace branch,
// yielding g_*²=1/2 and the tree-level Higgs proxy
// m_H = v sqrt(1197/4624) ≈ 125.274 GeV. Gate 331 audits the remaining
// precision gap between this running/tree proxy and a collider pole-mass claim.
package higgspolemassprecision

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE331-HIGGS-POLE-MASS-CONVERSION-PRECISION-GAP-LEDGER-AUDIT"

	StatusInputsInherited                = "CONDITIONAL_SUPPORT_GATE330_NATIVE_DOUBLED_TRACE_BRANCH_INHERITED"
	StatusTreeProxyRecomputed            = "CONDITIONAL_SUPPORT_NATIVE_TREE_LEVEL_HIGGS_PROXY_RECOMPUTED"
	StatusPrecisionGapQuantified         = "CONDITIONAL_SUPPORT_POLE_MASS_PRECISION_GAP_QUANTIFIED"
	StatusPoleConversionLedgerFormalized = "CONDITIONAL_SUPPORT_POLE_MASS_CONVERSION_LEDGER_FORMALIZED"
	StatusLoopCapacitySieveFormalized    = "CONDITIONAL_SUPPORT_LOOP_CORRECTION_CAPACITY_SIEVE_FORMALIZED"
	StatusNoColliderClaimPreserved       = "CONDITIONAL_SUPPORT_COLLIDER_CLAIM_FIREWALL_PRESERVED"

	StatusTensionPoleMassStillUnexecuted = "CONDITIONAL_TENSION_POLE_MASS_CONVERSION_STILL_UNEXECUTED"
	StatusTensionSchemeDependence        = "CONDITIONAL_TENSION_MS_BAR_TO_POLE_SCHEME_DEPENDENCE_REMAINS"

	StatusFailedSelfEnergiesNotComputed     = "FAILED_ROUTE_W_Z_TOP_SELF_ENERGIES_NOT_COMPUTED"
	StatusFailedTwoLoopNotExecuted          = "FAILED_ROUTE_TWO_LOOP_RG_NOT_EXECUTED"
	StatusFailedThresholdScaleConditional   = "FAILED_ROUTE_THRESHOLD_SCALE_STILL_CONDITIONAL"
	StatusFailedExactColliderMassNotClaimed = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate     = 330
	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	gStarSquaredNative       = 0.5
	electroweakVEVGeV        = 246.22
	observedHiggsGeV         = 125.10
)

type Inputs struct {
	HighestInheritedGate int
	BoundaryRatio        string
	GStarSquared         float64
	VEVGeV               float64
	ObservedPoleMassGeV  float64
	AddsFit              bool
	Status               string
}

type TreeProxy struct {
	LambdaH         float64
	Formula         string
	MassGeV         float64
	MassSquaredGeV2 float64
	Status          string
}

type TargetComparison struct {
	TargetLambda          float64
	DeltaLambda           float64
	DeltaMassGeV          float64
	DeltaMassSquaredGeV2  float64
	RelativeMassErrorPct  float64
	RequiredMassShiftSign string
	Status                string
}

type PoleLedger struct {
	RunningMassDefinition         string
	PoleMassDefinition            string
	RequiredCorrection            string
	RequiresTopSelfEnergy         bool
	RequiresWeakBosonSelfEnergy   bool
	RequiresScalarSelfEnergy      bool
	RequiresRenormalizationScheme bool
	Executed                      bool
	Status                        string
}

type CapacitySieve struct {
	GapGeV                                      float64
	GapIsSubGeV                                 bool
	PerturbativePoleCorrectionsCanHaveThisScale bool
	TwoLoopRGCanHaveThisScale                   bool
	ThresholdRetuningNotRequiredForThisGap      bool
	Explanation                                 string
	Status                                      string
}

type FirewallAudit struct {
	NoObservedMassFitted           bool
	NoSelfEnergiesComputed         bool
	NoTwoLoopClaim                 bool
	NoExactColliderClaim           bool
	ThresholdScaleStillConditional bool
	Status                         string
}

type Summary struct {
	NativeTreeProxyGeV   float64
	ObservedReferenceGeV float64
	DifferenceGeV        float64
	RelativeErrorPct     float64
	DirectAnswer         string
	NextObligation       string
	Status               string
}

type Analysis struct {
	Inputs     Inputs
	Tree       TreeProxy
	Comparison TargetComparison
	Pole       PoleLedger
	Capacity   CapacitySieve
	Audit      FirewallAudit
	Summary    Summary
	Truth      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inputs := compileInputs()
	tree := computeTreeProxy()
	comp := compareTarget(tree)
	pole := formalizePoleLedger()
	cap := auditCapacity(comp)
	firewall := auditFirewalls()
	summary := compileSummary(tree, comp)
	truth := "Gate 331 shows that the native Gate 330 doubled-bosonic-trace branch gives a tree/running Higgs proxy of about 125.274 GeV, only about +0.174 GeV above the nominal 125.10 GeV reference. This residual is a precision-scale MS-bar-to-pole conversion problem, not a structural 50–200 GeV hierarchy problem. The gate does not compute collider self-energies or two-loop running, so it preserves the firewall against an exact pole-mass claim."
	return Analysis{Inputs: inputs, Tree: tree, Comparison: comp, Pole: pole, Capacity: cap, Audit: firewall, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	return Inputs{HighestInheritedGate: inheritedHighestGate, BoundaryRatio: "λ_H/g_*² = 1197/4624; g_*² = 1/2 from full doubled bosonic trace", GStarSquared: gStarSquaredNative, VEVGeV: electroweakVEVGeV, ObservedPoleMassGeV: observedHiggsGeV, AddsFit: false, Status: StatusInputsInherited}
}

func computeTreeProxy() TreeProxy {
	lambda := (contactScalarNumerator / contactScalarDenominator) * gStarSquaredNative
	m2 := 2.0 * lambda * electroweakVEVGeV * electroweakVEVGeV
	return TreeProxy{LambdaH: lambda, Formula: "m_tree = v√(2λ_H) = v√(1197/4624)", MassGeV: math.Sqrt(m2), MassSquaredGeV2: m2, Status: StatusTreeProxyRecomputed}
}

func compareTarget(tree TreeProxy) TargetComparison {
	targetLambda := observedHiggsGeV * observedHiggsGeV / (2.0 * electroweakVEVGeV * electroweakVEVGeV)
	deltaLambda := targetLambda - tree.LambdaH
	deltaM := observedHiggsGeV - tree.MassGeV
	deltaM2 := observedHiggsGeV*observedHiggsGeV - tree.MassSquaredGeV2
	sign := "negative correction to the native tree/running proxy"
	if deltaM > 0 {
		sign = "positive correction to the native tree/running proxy"
	}
	return TargetComparison{TargetLambda: targetLambda, DeltaLambda: deltaLambda, DeltaMassGeV: deltaM, DeltaMassSquaredGeV2: deltaM2, RelativeMassErrorPct: (tree.MassGeV - observedHiggsGeV) / observedHiggsGeV * 100.0, RequiredMassShiftSign: sign, Status: StatusPrecisionGapQuantified}
}

func formalizePoleLedger() PoleLedger {
	return PoleLedger{
		RunningMassDefinition:         "m_run²(μ)=2λ(μ)v(μ)² in a chosen MS-bar-like scheme",
		PoleMassDefinition:            "M_H² solves p² - m_run²(μ) + Re Π_HH(p²,μ)=0 at p²=M_H²",
		RequiredCorrection:            "ΔM_H² = Re Π_HH(M_H²,μ) plus finite counterterm/scheme conversion; dominant ledgers include top, W, Z, Goldstone, and Higgs loops",
		RequiresTopSelfEnergy:         true,
		RequiresWeakBosonSelfEnergy:   true,
		RequiresScalarSelfEnergy:      true,
		RequiresRenormalizationScheme: true,
		Executed:                      false,
		Status:                        StatusPoleConversionLedgerFormalized,
	}
}

func auditCapacity(comp TargetComparison) CapacitySieve {
	gap := math.Abs(comp.DeltaMassGeV)
	return CapacitySieve{GapGeV: gap, GapIsSubGeV: gap < 1.0, PerturbativePoleCorrectionsCanHaveThisScale: true, TwoLoopRGCanHaveThisScale: true, ThresholdRetuningNotRequiredForThisGap: true, Explanation: "The remaining +0.174 GeV tree-proxy excess is sub-percent and sub-GeV, so it lies in the natural precision range of pole conversion, scheme choice, threshold-scale placement, and two-loop transport. It is not evidence for another large structural threshold.", Status: StatusLoopCapacitySieveFormalized}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{NoObservedMassFitted: true, NoSelfEnergiesComputed: true, NoTwoLoopClaim: true, NoExactColliderClaim: true, ThresholdScaleStillConditional: true, Status: StatusNoColliderClaimPreserved}
}

func compileSummary(tree TreeProxy, comp TargetComparison) Summary {
	return Summary{NativeTreeProxyGeV: tree.MassGeV, ObservedReferenceGeV: observedHiggsGeV, DifferenceGeV: tree.MassGeV - observedHiggsGeV, RelativeErrorPct: comp.RelativeMassErrorPct, DirectAnswer: "The native doubled-trace branch already sits within a sub-GeV precision window; the next missing layer is explicit pole-mass and two-loop precision, not a new threshold mechanism.", NextObligation: "Install a scheme-explicit one-loop pole-mass conversion ledger and two-loop RG precision transport before claiming an exact collider Higgs mass.", Status: StatusPoleConversionLedgerFormalized}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Tree.Status,
		a.Comparison.Status,
		a.Pole.Status,
		a.Capacity.Status,
		a.Audit.Status,
		StatusTensionPoleMassStillUnexecuted,
		StatusTensionSchemeDependence,
		StatusFailedSelfEnergiesNotComputed,
		StatusFailedTwoLoopNotExecuted,
		StatusFailedThresholdScaleConditional,
		StatusFailedExactColliderMassNotClaimed,
	}
}

func FormatInputs(x Inputs) string {
	return fmt.Sprintf("gate=%d boundary=%s g2=%.12f v=%.2f observed=%.2f fit=%v status=%s", x.HighestInheritedGate, x.BoundaryRatio, x.GStarSquared, x.VEVGeV, x.ObservedPoleMassGeV, x.AddsFit, x.Status)
}
func FormatTree(x TreeProxy) string {
	return fmt.Sprintf("λ=%.15f formula=%s m=%.12f GeV m²=%.12f status=%s", x.LambdaH, x.Formula, x.MassGeV, x.MassSquaredGeV2, x.Status)
}
func FormatComparison(x TargetComparison) string {
	return fmt.Sprintf("λ_target=%.15f Δλ=%+.15f Δm=%+.12f GeV Δm²=%+.12f rel=%+.6f%% sign=%s status=%s", x.TargetLambda, x.DeltaLambda, x.DeltaMassGeV, x.DeltaMassSquaredGeV2, x.RelativeMassErrorPct, x.RequiredMassShiftSign, x.Status)
}
func FormatPole(x PoleLedger) string {
	return fmt.Sprintf("run=%s pole=%s correction=%s top=%v weak=%v scalar=%v scheme=%v executed=%v status=%s", x.RunningMassDefinition, x.PoleMassDefinition, x.RequiredCorrection, x.RequiresTopSelfEnergy, x.RequiresWeakBosonSelfEnergy, x.RequiresScalarSelfEnergy, x.RequiresRenormalizationScheme, x.Executed, x.Status)
}
func FormatCapacity(x CapacitySieve) string {
	return fmt.Sprintf("gap=%.12f subGeV=%v poleCapacity=%v twoLoopCapacity=%v noNewThreshold=%v status=%s reason=%s", x.GapGeV, x.GapIsSubGeV, x.PerturbativePoleCorrectionsCanHaveThisScale, x.TwoLoopRGCanHaveThisScale, x.ThresholdRetuningNotRequiredForThisGap, x.Status, x.Explanation)
}
func FormatAudit(x FirewallAudit) string {
	return fmt.Sprintf("noFit=%v noSelfEnergy=%v noTwoLoop=%v noColliderClaim=%v thresholdConditional=%v status=%s", x.NoObservedMassFitted, x.NoSelfEnergiesComputed, x.NoTwoLoopClaim, x.NoExactColliderClaim, x.ThresholdScaleStillConditional, x.Status)
}
func FormatSummary(x Summary) string {
	return fmt.Sprintf("m_native=%.12f observed=%.12f diff=%+.12f rel=%+.6f%% answer=%s next=%s status=%s", x.NativeTreeProxyGeV, x.ObservedReferenceGeV, x.DifferenceGeV, x.RelativeErrorPct, x.DirectAnswer, x.NextObligation, x.Status)
}
func FormatStatuses(ss []string) string  { return "statuses=" + strings.Join(ss, "; ") }
func nearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
