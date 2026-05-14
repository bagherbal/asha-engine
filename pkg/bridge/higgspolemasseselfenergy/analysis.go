// Package higgspolemasseselfenergy implements Gate 332:
// One-Loop Higgs Pole Self-Energy Target / Minimal Precision Correction Audit.
//
// Gate 331 quantified the remaining difference between the native doubled-trace
// tree/running Higgs proxy and the nominal collider pole mass. Gate 332 converts
// that residual into the exact mass-squared/self-energy target required by a
// one-loop pole conversion, and audits whether the correction is perturbatively
// natural without executing a full Standard Model self-energy calculation.
package higgspolemasseselfenergy

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE332-HIGGS-POLE-SELF-ENERGY-TARGET-MINIMAL-PRECISION-CORRECTION-AUDIT"

	StatusGate331Inherited           = "CONDITIONAL_SUPPORT_GATE331_PRECISION_GAP_INHERITED"
	StatusPoleEquationFormalized     = "CONDITIONAL_SUPPORT_POLE_EQUATION_SELF_ENERGY_TARGET_FORMALIZED"
	StatusRequiredSelfEnergyComputed = "CONDITIONAL_SUPPORT_REQUIRED_SELF_ENERGY_TARGET_COMPUTED"
	StatusLoopScaleCapacityAudited   = "CONDITIONAL_SUPPORT_ONE_LOOP_SCALE_CAPACITY_AUDITED"
	StatusPrecisionLedgerFormalized  = "CONDITIONAL_SUPPORT_MINIMAL_PRECISION_CORRECTION_LEDGER_FORMALIZED"
	StatusNoColliderClaimPreserved   = "CONDITIONAL_SUPPORT_EXACT_COLLIDER_CLAIM_FIREWALL_PRESERVED"

	StatusTensionSelfEnergyNotDerived = "CONDITIONAL_TENSION_SELF_ENERGY_TARGET_NOT_DERIVED_FROM_LOOP_INTEGRALS"
	StatusTensionSchemeDependent      = "CONDITIONAL_TENSION_POLE_CONVERSION_SCHEME_DEPENDENT"

	StatusFailedFullOneLoopNotComputed  = "FAILED_ROUTE_FULL_ONE_LOOP_SELF_ENERGY_NOT_COMPUTED"
	StatusFailedTwoLoopNotComputed      = "FAILED_ROUTE_TWO_LOOP_PRECISION_NOT_COMPUTED"
	StatusFailedSMInputsNotInstalled    = "FAILED_ROUTE_FULL_SM_INPUT_SET_NOT_INSTALLED"
	StatusFailedExactPoleMassNotClaimed = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate     = 331
	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	gStarSquaredNative       = 0.5
	electroweakVEVGeV        = 246.22
	observedHiggsGeV         = 125.10
)

type Inputs struct {
	HighestInheritedGate  int
	LambdaH               float64
	TreeMassGeV           float64
	ObservedPoleGeV       float64
	UsesObservedForTarget bool
	Status                string
}

type PoleEquation struct {
	Convention             string
	RunningMassSquaredGeV2 float64
	PoleMassSquaredGeV2    float64
	DeltaPoleMinusRunGeV2  float64
	RequiredRePiGeV2       float64
	Status                 string
}

type SelfEnergyTarget struct {
	RequiredMassShiftGeV         float64
	RequiredMassSquaredShiftGeV2 float64
	RequiredRePiGeV2             float64
	RequiredRePiOverRunMass2     float64
	RequiredRePiOverLoopUnit     float64
	Sign                         string
	Status                       string
}

type LoopCapacity struct {
	LoopUnitGeV2            float64
	LoopUnitFormula         string
	RequiredIsOrderOneLoop  bool
	RequiredIsSmallFraction bool
	CapacityExplanation     string
	Status                  string
}

type PrecisionLedger struct {
	NeedsTopLoop             bool
	NeedsWZLoops             bool
	NeedsHiggsGoldstoneLoops bool
	NeedsCounterterms        bool
	NeedsSchemeChoice        bool
	FullCalculationExecuted  bool
	Status                   string
}

type FirewallAudit struct {
	NoLoopIntegralsEvaluated bool
	NoTwoLoopClaim           bool
	NoExactPoleMassClaim     bool
	NoFitParameterIntroduced bool
	Status                   string
}

type Summary struct {
	RequiredSelfEnergyGeV2 float64
	RequiredMassShiftGeV   float64
	LoopUnits              float64
	DirectAnswer           string
	NextObligation         string
	Status                 string
}

type Analysis struct {
	Inputs   Inputs
	Equation PoleEquation
	Target   SelfEnergyTarget
	Capacity LoopCapacity
	Ledger   PrecisionLedger
	Audit    FirewallAudit
	Summary  Summary
	Truth    string
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
	equation := formalizePoleEquation(inputs)
	target := computeSelfEnergyTarget(inputs, equation)
	capacity := auditLoopCapacity(inputs, target)
	ledger := formalizePrecisionLedger()
	audit := auditFirewalls()
	summary := compileSummary(target)
	truth := "Gate 332 converts the remaining Gate 331 +0.174 GeV native tree-proxy excess into a required one-loop pole conversion target. Under the convention p²-m_run²+ReΠ=0, the required finite self-energy is ReΠ≈+43.604 GeV², equivalently a pole-minus-running shift of -43.604 GeV². This is about 0.88 times the natural λv²/(16π²) one-loop scalar scale, so it is a precision-sized correction, but the gate does not compute the Standard Model self-energy integrals or claim the exact collider pole mass."
	return Analysis{Inputs: inputs, Equation: equation, Target: target, Capacity: capacity, Ledger: ledger, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	lambda := (contactScalarNumerator / contactScalarDenominator) * gStarSquaredNative
	treeMass := electroweakVEVGeV * math.Sqrt(contactScalarNumerator/contactScalarDenominator)
	return Inputs{HighestInheritedGate: inheritedHighestGate, LambdaH: lambda, TreeMassGeV: treeMass, ObservedPoleGeV: observedHiggsGeV, UsesObservedForTarget: true, Status: StatusGate331Inherited}
}

func formalizePoleEquation(in Inputs) PoleEquation {
	run2 := in.TreeMassGeV * in.TreeMassGeV
	pole2 := in.ObservedPoleGeV * in.ObservedPoleGeV
	delta := pole2 - run2
	// With p² - m_run² + ReΠ(p²)=0 at p²=M², ReΠ_required = m_run² - M².
	requiredPi := run2 - pole2
	return PoleEquation{Convention: "M_H² - m_run² + ReΠ_HH(M_H²)=0, hence ReΠ_required=m_run²-M_H²", RunningMassSquaredGeV2: run2, PoleMassSquaredGeV2: pole2, DeltaPoleMinusRunGeV2: delta, RequiredRePiGeV2: requiredPi, Status: StatusPoleEquationFormalized}
}

func computeSelfEnergyTarget(in Inputs, eq PoleEquation) SelfEnergyTarget {
	shift := in.ObservedPoleGeV - in.TreeMassGeV
	loopUnit := in.LambdaH * electroweakVEVGeV * electroweakVEVGeV / (16.0 * math.Pi * math.Pi)
	sign := "negative pole-mass shift; positive ReΠ in the chosen pole equation convention"
	return SelfEnergyTarget{RequiredMassShiftGeV: shift, RequiredMassSquaredShiftGeV2: eq.DeltaPoleMinusRunGeV2, RequiredRePiGeV2: eq.RequiredRePiGeV2, RequiredRePiOverRunMass2: eq.RequiredRePiGeV2 / eq.RunningMassSquaredGeV2, RequiredRePiOverLoopUnit: eq.RequiredRePiGeV2 / loopUnit, Sign: sign, Status: StatusRequiredSelfEnergyComputed}
}

func auditLoopCapacity(in Inputs, target SelfEnergyTarget) LoopCapacity {
	loopUnit := in.LambdaH * electroweakVEVGeV * electroweakVEVGeV / (16.0 * math.Pi * math.Pi)
	orderOne := math.Abs(target.RequiredRePiOverLoopUnit) > 0.1 && math.Abs(target.RequiredRePiOverLoopUnit) < 10.0
	smallFraction := math.Abs(target.RequiredRePiOverRunMass2) < 0.01
	explanation := "The required correction is sub-percent relative to m² and order-one in the natural λv²/(16π²) loop unit. This is compatible with a pole-conversion precision ledger, but it is not a derivation of the finite self-energy coefficient."
	return LoopCapacity{LoopUnitGeV2: loopUnit, LoopUnitFormula: "λ_H v²/(16π²)", RequiredIsOrderOneLoop: orderOne, RequiredIsSmallFraction: smallFraction, CapacityExplanation: explanation, Status: StatusLoopScaleCapacityAudited}
}

func formalizePrecisionLedger() PrecisionLedger {
	return PrecisionLedger{NeedsTopLoop: true, NeedsWZLoops: true, NeedsHiggsGoldstoneLoops: true, NeedsCounterterms: true, NeedsSchemeChoice: true, FullCalculationExecuted: false, Status: StatusPrecisionLedgerFormalized}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{NoLoopIntegralsEvaluated: true, NoTwoLoopClaim: true, NoExactPoleMassClaim: true, NoFitParameterIntroduced: true, Status: StatusNoColliderClaimPreserved}
}

func compileSummary(target SelfEnergyTarget) Summary {
	return Summary{RequiredSelfEnergyGeV2: target.RequiredRePiGeV2, RequiredMassShiftGeV: target.RequiredMassShiftGeV, LoopUnits: target.RequiredRePiOverLoopUnit, DirectAnswer: "The residual 0.174 GeV gap corresponds to ReΠ≈+43.604 GeV² in the chosen pole equation, a natural one-loop-sized precision correction.", NextObligation: "Install the explicit one-loop Standard Model Higgs self-energy functions and a fixed renormalization scheme; then add two-loop/RG precision before making a collider pole-mass claim.", Status: StatusRequiredSelfEnergyComputed}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Equation.Status,
		a.Target.Status,
		a.Capacity.Status,
		a.Ledger.Status,
		a.Audit.Status,
		StatusTensionSelfEnergyNotDerived,
		StatusTensionSchemeDependent,
		StatusFailedFullOneLoopNotComputed,
		StatusFailedTwoLoopNotComputed,
		StatusFailedSMInputsNotInstalled,
		StatusFailedExactPoleMassNotClaimed,
	}
}

func FormatInputs(x Inputs) string {
	return fmt.Sprintf("gate=%d λ=%.15f m_tree=%.12f observed=%.12f usesObservedForTarget=%v status=%s", x.HighestInheritedGate, x.LambdaH, x.TreeMassGeV, x.ObservedPoleGeV, x.UsesObservedForTarget, x.Status)
}
func FormatEquation(x PoleEquation) string {
	return fmt.Sprintf("run²=%.12f pole²=%.12f Δ(pole-run)=%+.12f ReΠ_req=%+.12f convention=%s status=%s", x.RunningMassSquaredGeV2, x.PoleMassSquaredGeV2, x.DeltaPoleMinusRunGeV2, x.RequiredRePiGeV2, x.Convention, x.Status)
}
func FormatTarget(x SelfEnergyTarget) string {
	return fmt.Sprintf("Δm=%+.12f GeV Δm²=%+.12f ReΠ=%+.12f ReΠ/m²=%.9f ReΠ/loop=%.9f sign=%s status=%s", x.RequiredMassShiftGeV, x.RequiredMassSquaredShiftGeV2, x.RequiredRePiGeV2, x.RequiredRePiOverRunMass2, x.RequiredRePiOverLoopUnit, x.Sign, x.Status)
}
func FormatCapacity(x LoopCapacity) string {
	return fmt.Sprintf("loopUnit=%.12f formula=%s orderOne=%v smallFraction=%v status=%s reason=%s", x.LoopUnitGeV2, x.LoopUnitFormula, x.RequiredIsOrderOneLoop, x.RequiredIsSmallFraction, x.Status, x.CapacityExplanation)
}
func FormatLedger(x PrecisionLedger) string {
	return fmt.Sprintf("top=%v WZ=%v scalar=%v counterterms=%v scheme=%v executed=%v status=%s", x.NeedsTopLoop, x.NeedsWZLoops, x.NeedsHiggsGoldstoneLoops, x.NeedsCounterterms, x.NeedsSchemeChoice, x.FullCalculationExecuted, x.Status)
}
func FormatAudit(x FirewallAudit) string {
	return fmt.Sprintf("noLoops=%v noTwoLoop=%v noExactPole=%v noFit=%v status=%s", x.NoLoopIntegralsEvaluated, x.NoTwoLoopClaim, x.NoExactPoleMassClaim, x.NoFitParameterIntroduced, x.Status)
}
func FormatSummary(x Summary) string {
	return fmt.Sprintf("ReΠ=%.12f Δm=%+.12f loopUnits=%.9f answer=%s next=%s status=%s", x.RequiredSelfEnergyGeV2, x.RequiredMassShiftGeV, x.LoopUnits, x.DirectAnswer, x.NextObligation, x.Status)
}
func FormatStatuses(ss []string) string  { return "statuses=" + strings.Join(ss, "; ") }
func nearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
