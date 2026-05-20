package generation2u2invariantrenormalizablehiggspotentialformandcoefficientsealaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2U2InvariantRenormalizableHiggsPotentialFormAndCoefficientSealAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 769 — U(2)-Invariant Renormalizable Higgs Potential Form and Coefficient-Seal Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate769U2InvariantPotentialFormBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 768 Hessian spectral projector firewall", Passed: a.Gate768.Inherited && strings.Contains(a.Gate768.RadialReplacement, "P_rad := supp") && strings.Contains(a.Gate768.HessianSupport, "2 lambda v^2") && strings.Contains(a.Gate768.LHopfSource, "supp") && a.Gate768.PotentialStillSupplied && !a.Gate768.NativePotentialTheorem && !a.Gate768.NativeHistoryLoopTheorem, Detail: FormatGate768(a.Gate768)},
			{Name: "inherit C2 Higgs carrier after sealed twistor selector", Passed: a.Carrier.Selector == "sealed twistor selector n" && a.Carrier.ComplexStructure == "J_H(n)" && strings.Contains(a.Carrier.Carrier, "C^2") && a.Carrier.ComplexDimension == complexHiggsDim && a.Carrier.RealDimension == realHiggsDim && strings.Contains(a.Carrier.RepresentationSocket, "U(2)") && !a.Carrier.NativeSelectorTheorem, Detail: FormatCarrier(a.Carrier)},
			{Name: "reduce U(2)-invariant function to phi dagger phi", Passed: strings.Contains(a.Reduction.Action, "transitive") && a.Reduction.InvariantCoordinate == "r^2=phi^dagger phi" && a.Reduction.TransitiveOnFixedRadiusSpheres && a.Reduction.PotentialFunctionForm == "V(phi)=f(phi^dagger phi)" && a.Reduction.DependsOnlyOnPhiDaggerPhi && !a.Reduction.SelectsCP1Point && a.Reduction.RequiresNoAnisotropicHermitianAxis, Detail: FormatReduction(a.Reduction)},
			{Name: "audit renormalizable polynomial normal form", Passed: a.Polynomial.RealPolynomialPremise && a.Polynomial.QuarticTruncation && a.Polynomial.QuarticDegreeInRealFields == quarticDegree && a.Polynomial.Coordinate == "r^2=phi^dagger phi" && a.Polynomial.MaxPowerInCoordinate == maxFPower && strings.Contains(a.Polynomial.FunctionForm, "c_0+mu^2") && strings.Contains(a.Polynomial.PotentialForm, "lambda(phi^dagger phi)^2") && a.Polynomial.UniqueUnderPremises && !a.Polynomial.NativeSpectralActionTheorem && !a.Polynomial.NativeScalarPotentialTheorem, Detail: FormatPolynomial(a.Polynomial)},
			{Name: "separate constant offset from local Hessian/radial dynamics", Passed: a.Constant.ConstantSymbol == "c_0" && !a.Constant.AffectsGradient && !a.Constant.AffectsHessian && !a.Constant.AffectsRadialEvent && a.Constant.IgnoredForLocalScalarDynamics && !a.Constant.CosmologicalConstantTheorem, Detail: FormatConstant(a.Constant)},
			{Name: "audit coefficient seals", Passed: strings.Contains(a.Coefficients.MuSquaredRole, "quadratic") && !a.Coefficients.MuSquaredDerived && !a.Coefficients.MuSquaredSignDerived && strings.Contains(a.Coefficients.LambdaRole, "quartic") && !a.Coefficients.LambdaDerived && a.Coefficients.LambdaRuntimeBridgeMaySupply && !a.Coefficients.RuntimeLambdaIndependent && strings.Contains(a.Coefficients.C0Role, "vacuum-energy offset") && !a.Coefficients.C0CosmologicalTheorem, Detail: FormatCoefficients(a.Coefficients)},
			{Name: "preserve CP1 flatness", Passed: strings.Contains(a.CP1.Reason, "fixed-radius") && a.CP1.FlatAtFixedRadius && !a.CP1.CP1SelectedByPotential && a.CP1.RadialDirectionNonzero && a.CP1.AngularDirectionsFlat == 3 && a.CP1.PreservesGate764765, Detail: FormatCP1(a.CP1)},
			{Name: "record Hessian compatibility with Gate 766/768", Passed: a.Hessian.RealCoordinateConvention == "phi^dagger phi=(1/2)||x||^2" && strings.Contains(a.Hessian.RealPotentialForm, "lambda/4") && a.Hessian.Gate766HessianFormula == "H_V(x_0)=2 lambda v^2 P_rad" && strings.Contains(a.Hessian.Gate768SupportReplacement, "supp") && a.Hessian.HessianNormalizationBelongsHere && !a.Hessian.NativeVEVTheorem && !a.Hessian.PoleMassTheorem, Detail: FormatHessian(a.Hessian)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.NativeScalarPotentialTheorem && !a.Firewalls.NativeMuSquaredTheorem && !a.Firewalls.NativeQuarticCoefficientTheorem && !a.Firewalls.NativeVEVTheorem && !a.Firewalls.NativeSpectralActionTruncationTheorem && !a.Firewalls.C0CosmologicalConstantTheorem && !a.Firewalls.NativeHistoryLoopUnitTheorem && !a.Firewalls.TreeProxyPoleMassTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate769U2InvariantPotentialFormBoundary), Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
