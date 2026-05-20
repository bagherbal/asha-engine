package generation2crosssectororientationintertwinerminimalityaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedRelation) string {
	return fmt.Sprintf("eps=%.15g rad %.15g deg kappa=%.15g A=%.15g J=%.15g B=%.15g delta590=%.15g rel=%.15g epsPred=%.15g rad epsResidual=%.15g rad oneSigma=[%.15g,%.15g] covers=%t belowR=%t belowQ=%t Rdef=%.15g absQ=%.15g thetaDominates=%t verdict=%q", a.EpsilonObsRad, a.EpsilonObsDeg, a.KappaObs, a.ReactorQuarter, a.JCKM, a.OrientationCandidate, a.Delta590, a.RelativeDelta590, a.EpsilonPredictionRad, a.EpsilonResidualRad, a.CombinedOneSigmaLow, a.CombinedOneSigmaHigh, a.ResidualInsideOneSigma, a.ResidualBelowRDefect, a.ResidualBelowQDefect, a.RDefect, a.AbsQResidual, a.Theta13DominatesUncertainty, a.Verdict)
}

func FormatTypedObject(a TypedObject) string {
	return fmt.Sprintf("symbol=%q carrier=%q role=%q equation=%q value=%.15g nativeStatus=%q verdict=%q", a.Symbol, a.Carrier, a.Role, a.Equation, a.RuntimeValue, a.NativeStatus, a.Verdict)
}

func FormatTyped(a TypedObjectAudit) string {
	parts := make([]string, 0, len(a.Objects))
	for _, obj := range a.Objects {
		parts = append(parts, FormatTypedObject(obj))
	}
	return fmt.Sprintf("objects=[%s] verdict=%q", strings.Join(parts, "; "), a.Verdict)
}

func FormatRequired(a RequiredBridge) string {
	return fmt.Sprintf("names=%q domain=%q codomain=%q map=%q minimality=%q basisInvariant=%t rootSpace=%t sectorTypes=%t verdict=%q", strings.Join(a.NameCandidates, ","), strings.Join(a.Domain, ","), a.Codomain, a.RequiredMap, a.Minimality, a.MustBeBasisInvariant, a.MustHandleRootSpace, a.MustRespectSectorTypes, a.Verdict)
}

func FormatRepositoryObject(a RepositoryObjectAudit) string {
	return fmt.Sprintf("name=%q present=%t canIntertwine=%t reason=%q verdict=%q", a.Name, a.PresentInASHA, a.CanSupplyIntertwiner, a.Reason, a.Verdict)
}

func FormatRepository(a RepositoryAudit) string {
	parts := make([]string, 0, len(a.Objects))
	for _, obj := range a.Objects {
		parts = append(parts, FormatRepositoryObject(obj))
	}
	return fmt.Sprintf("anyCrossSector=%t rootTraceOrAbsDirac=%t objects=[%s] verdict=%q", a.AnyNativeCrossSectorIntertwiner, a.NativeRootTraceOrAbsoluteDirac, strings.Join(parts, "; "), a.Verdict)
}

func FormatSeal(a OrientationBalanceSeal) string {
	return fmt.Sprintf("name=%q kappaDef=%q epsilonDef=%q kappa=%.15g candidate=%.15g residual=%.15g eps=%.15g epsPred=%.15g epsResidual=%.15g residualStatus=%q native=%t interpretation=%q verdict=%q", a.Name, a.KappaDefinition, a.EpsilonDefinition, a.KappaValue, a.KappaCandidate, a.KappaResidual, a.EpsilonValueRad, a.EpsilonPredRad, a.EpsilonResidualRad, a.ResidualStatus, a.Native, a.Interpretation, a.Verdict)
}

func FormatPrecision(a PrecisionAudit) string {
	return fmt.Sprintf("delta590=%.15g abs=%.15g oneSigmaDistances=[%.15g,%.15g] sigmaFractions=[%.15g,%.15g] Rdef=%.15g absQ=%.15g belowR=%t belowQ=%t addCorrection=%t verdict=%q", a.Delta590, a.AbsDelta590, a.OneSigmaLowDistance, a.OneSigmaHighDistance, a.SigmaFractionMinus, a.SigmaFractionPlus, a.RDefect, a.AbsQResidual, a.DeltaSmallerThanRDefect, a.DeltaSmallerThanQResidual, a.AdditionalCorrectionJustified, a.Verdict)
}

func FormatLawfulness(a LawfulnessAudit) string {
	return fmt.Sprintf("crossSector=%t balance=%t rootMap=%t rootTrace=%t absDirac=%t deriveWall=%t derivePMNS=%t deriveCKM=%t deriveKappa=%t verdict=%q", a.CrossSectorOrientationIntertwinerPresent, a.FlavorOrientationBalanceOperatorPresent, a.RootSpaceOrientationMapPresent, a.NativeRootTraceOperatorPresent, a.AbsoluteDiracObservablePresent, a.DerivesKoideWallCoordinate, a.DerivesPMNSReactorAngle, a.DerivesCKMJarlskog, a.DerivesKappaRelation, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t pmns=%t ckm=%t theta13=%t neutrino=%t masses=%t texture=%t observedNative=%t newCarrier=%t newSelector=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesPMNS, a.DerivesCKM, a.DerivesTheta13, a.DerivesNeutrinoPhysics, a.DerivesChargedLeptonMasses, a.DerivesFlavorTexture, a.PromotesObservedAsNative, a.AddsNewCarrier, a.AddsNewSelector, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("typed=%q nativeIntertwiner=%t residualMeaningful=%t minimalSeal=%q environmental=%t decision=%q verdict=%q", a.TypedObjectsConnected, a.NativeIntertwinerPresent, a.ResidualMeaningfulBeyondV1, a.MinimalSeal, a.KappaRemainsEnvironmental, a.Decision, a.Verdict)
}
