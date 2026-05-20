package generation2boundaryweighteddeficitclosurescalesweepaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryWeightedDeficitClosureScaleSweepAndSensitivityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 662 — BoundaryWeightedDeficitClosure Scale-Sweep and Sensitivity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate662 scale-sweep audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate661 noncircular closure", Passed: a.Inherited.ClosureInherited && math.Abs(a.Inherited.E72-8.525834413464217e-10) < 5e-18 && a.Inherited.Lambda12OnlyComputed && a.Inherited.FormulaLiftCircular && a.Inherited.NoNativeSevenOver72 && a.Inherited.NoNativeTransport && a.Inherited.NoIndependentEndpoint && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "construct v1 transport seed", Passed: a.Seed.Mu0GeV > 0 && a.Seed.Lambda12GeV > 9e13 && a.Seed.T13 > a.Seed.T12 && a.Seed.T23 > a.Seed.T13 && len(a.Seed.InitialVector) == 13, Detail: FormatSeed(a.Seed)},
			{Name: "compute scale sweep and Lambda12 selection", Passed: len(a.ScaleSweep.Rows) == 4 && a.ScaleSweep.BestEWMeanScale == "Lambda_12" && a.ScaleSweep.BestPairScale == "Lambda_12" && a.ScaleSweep.BestEWMeanResidual < 1e-9 && a.ScaleSweep.BestPairResidual < 1e-9 && a.ScaleSweep.Lambda12UniquelyMinimalEW && a.ScaleSweep.Lambda12UniquelyMinimalPair, Detail: FormatScaleSweep(a.ScaleSweep)},
			{Name: "compute local Lambda12 perturbation sweep", Passed: len(a.LocalSweep.Rows) == 9 && a.LocalSweep.LocalGridSelectsLambda12 && math.Abs(a.LocalSweep.MinimumDeltaLog) < 1e-15 && a.LocalSweep.MinimumAbsResidual < 1e-9 && a.LocalSweep.Threshold1eMinus4Width >= 0.1 && a.LocalSweep.FiniteDifferenceSlope > 0.0009, Detail: FormatLocalPerturbation(a.LocalSweep)},
			{Name: "audit best weight and orientation sensitivity", Passed: math.Abs(a.Weight.WBestExact-0.09722288188941036) < 5e-15 && math.Abs(a.Weight.WBestExactMinus7Over72-6.596671881381466e-07) < 5e-15 && math.Abs(a.Weight.WBestOrientation-0.09937065106104444) < 5e-15 && a.Weight.ExactWeightNear7Over72 && a.Weight.OrientationWeightNear7Over72, Detail: FormatWeight(a.Weight)},
			{Name: "compute input sensitivity Jacobian", Passed: a.Jacobian.DE_DKappaE == 1 && math.Abs(a.Jacobian.DE_DAbsLambda+65.0/72.0) < 1e-15 && math.Abs(a.Jacobian.DE_DR3Minus1+7.0/72.0) < 1e-15 && a.Jacobian.DKappa_DLambdaRuntime < -200 && a.Jacobian.DKappa_DLambdaProxy > 200 && a.Jacobian.DKappa_DL > 20 && len(a.Jacobian.Notes) == 3, Detail: FormatJacobian(a.Jacobian)},
			{Name: "audit orientation substitution scale effect", Passed: math.Abs(a.Orientation.OrientationE72AtLambda12-2.7767257213331953e-06) < 5e-18 && a.Orientation.ClosureResidualAmplification > 3000 && a.Orientation.BestWeightShift > 0.002, Detail: FormatOrientation(a.Orientation)},
			{Name: "preserve scale-sweep firewalls", Passed: !a.Discipline.ClaimsNativeScaleSelection && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsFullUncertaintyPropagation && !a.Discipline.ClaimsNativeTransportTheorem && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate662Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func notesContain(notes []string, want string) bool {
	return strings.Contains(strings.Join(notes, "\n"), want)
}
