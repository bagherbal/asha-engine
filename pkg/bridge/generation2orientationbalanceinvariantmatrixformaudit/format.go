package generation2orientationbalanceinvariantmatrixformaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedRelation) string {
	return fmt.Sprintf("eps=%.15g rad %.15g deg kappa=%.15g trace=%.15g A=%.15g J=%.15g candidate=%.15g delta=%.15g epsPred=%.15g epsResidual=%.15g insideSigma=%t belowR=%t belowQ=%t verdict=%q", a.EpsilonObsRad, a.EpsilonObsDeg, a.KappaObs, a.ReactorTrace, a.ReactorQuarter, a.JCKM, a.OrientationCandidate, a.Delta590, a.EpsilonPredRad, a.EpsilonResidualRad, a.ResidualInsideSigma, a.ResidualBelowRDefect, a.ResidualBelowQDefect, a.Verdict)
}

func FormatRootSpace(a RootSpaceMapAudit) string {
	return fmt.Sprintf("source=%q singularMap=%q x=%q chamber=%q epsilon=%q requiresRoot=%t rootTrace=%t absDirac=%t gate352=%t verdict=%q", a.Source, a.SingularValueMap, a.SquareRootVector, a.FourierKoideChamberCoordinate, a.EpsilonFunctional, a.RequiresRootSpectrumOperation, a.NativeRootTraceOperatorPresent, a.NativeAbsoluteDiracPresent, a.Gate352ObstructionPreserved, a.Verdict)
}

func FormatPMNS(a PMNSProjectorAudit) string {
	return fmt.Sprintf("expr=%q projectors=%q labels=%q trace=%.15g quarter=%.15g convention=%q native=%t observed=%t verdict=%q", a.Expression, strings.Join(a.Projectors, ","), strings.Join(a.RequiredLabels, ","), a.TraceValue, a.ReactorQuarter, a.Convention, a.NativeOperator, a.ObservedLedger, a.Verdict)
}

func FormatCKM(a CKMOrientationAudit) string {
	return fmt.Sprintf("Jexpr=%q J=%.15g commutator=%q sign=%q labels=%q basisInvariant=%t native=%t observed=%t verdict=%q", a.RephasingInvariantExpression, a.JCKM, a.CommutatorExpression, a.CommutatorSignConvention, strings.Join(a.RequiredLabels, ","), a.BasisInvariantGivenSpectra, a.NativeOperator, a.ObservedLedger, a.Verdict)
}

func FormatBalance(a InvariantBalanceAudit) string {
	return fmt.Sprintf("projectorEq=%q functionalEq=%q left=%.15g right=%.15g residual=%.15g abs=%.15g epsEq=%q epsPred=%.15g epsResidual=%.15g insideSigma=%t verdict=%q", a.EquationProjector, a.EquationFunctional, a.LeftKappa, a.RightProjectorMinusCKM, a.Residual, a.AbsResidual, a.EpsilonEquation, a.EpsilonPredictionRad, a.EpsilonResidualRad, a.ResidualInsideSigma, a.Verdict)
}

func FormatLabelDependency(a LabelDependency) string {
	return fmt.Sprintf("object=%q invariant=%q seal=%q reason=%q verdict=%q", a.Object, a.InvariantPart, a.RequiredSeal, a.Reason, a.Verdict)
}

func FormatLabels(a LabelAudit) string {
	parts := make([]string, 0, len(a.Labels))
	for _, label := range a.Labels {
		parts = append(parts, FormatLabelDependency(label))
	}
	return fmt.Sprintf("explicit=%t labels=[%s] verdict=%q", a.AllLabelSealsExplicit, strings.Join(parts, "; "), a.Verdict)
}

func FormatAvailabilityItem(a AvailabilityItem) string {
	return fmt.Sprintf("object=%q observed=%t native=%t canSupply=%t reason=%q verdict=%q", a.Object, a.ObservedLedger, a.NativeOperator, a.CanSupplyBalance, a.Reason, a.Verdict)
}

func FormatAvailability(a AvailabilityAudit) string {
	parts := make([]string, 0, len(a.Items))
	for _, item := range a.Items {
		parts = append(parts, FormatAvailabilityItem(item))
	}
	return fmt.Sprintf("nativeBalance=%t nativeRoot=%t nativeCommutator=%t items=[%s] verdict=%q", a.AnyNativeBalanceOperator, a.AnyNativeRootSpectrumMap, a.AnyNativeFlavorCommutatorMap, strings.Join(parts, "; "), a.Verdict)
}

func FormatTarget(a OperatorTarget) string {
	return fmt.Sprintf("name=%q domain=%q codomain=%q zero=%q equation=%q root=%t projectors=%t jarlskog=%t rephasing=%t native=%t verdict=%q", a.Name, strings.Join(a.Domain, ","), a.Codomain, a.ZeroCondition, a.RequiredEquation, a.MustHandleRootSpectrum, a.MustHandleProjectors, a.MustHandleJarlskogArea, a.MustBeRephasingInvariant, a.NativePresent, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t pmns=%t ckm=%t yukawas=%t neutrino=%t texture=%t observedNative=%t newCarrier=%t newSelector=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesPMNS, a.DerivesCKM, a.DerivesYukawas, a.DerivesNeutrinoPhysics, a.DerivesFlavorTexture, a.PromotesObservedData, a.AddsNewCarrier, a.AddsNewSelector, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("invariant=%t labels=%q native=%t missing=%q environmental=%t decision=%q verdict=%q", a.InvariantFormAvailable, a.RequiredLabels, a.NativeOperatorPresent, a.MissingOperatorTarget, a.OrientationBalanceSealEnvironmental, a.Decision, a.Verdict)
}
