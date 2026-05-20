package generation2flavorspectralorientationbalancefunctionalaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedRelation) string {
	return fmt.Sprintf("eps=%.15g kappa=%.15g trace=%.15g quarter=%.15g J=%.15g rhs=%.15g delta590=%.15g Bflav=%.15g epsPred=%.15g epsResidual=%.15g insideSigma=%t verdict=%q", a.EpsilonObsRad, a.KappaObs, a.PMNSTrace, a.PMNSQuarter, a.JCKM, a.RightHandCandidate, a.Delta590, a.BFlavValue, a.EpsilonPredictionRad, a.EpsilonResidualRad, a.ResidualInsideSigma, a.Verdict)
}

func FormatAlgebra(a SpectralAlgebraAudit) string {
	return fmt.Sprintf("name=%q generators=%q definition=%q carrier=%q observedOnly=%t native=%t verdict=%q", a.Name, strings.Join(a.Generators, ","), a.Definition, a.CommonCarrier, a.ObservedOnly, a.NativeAlgebraPresent, a.Verdict)
}

func FormatChargedLepton(a ChargedLeptonRootFunctional) string {
	return fmt.Sprintf("H=%q spectral=%q singular=%q x=%q chamber=%q epsilon=%q spectralCalc=%t fourthRoot=%t chamberSeal=%t native=%t gate352=%t verdict=%q", a.HermitianObject, a.SpectralData, a.SingularValueExtraction, a.SquareRootVectorExtraction, a.KoideChamberCoordinate, a.EpsilonFunctional, a.RequiresSpectralCalculus, a.RequiresFourthRootOfHE, a.RequiresChamberOrderSeal, a.NativeFunctionalPresent, a.Gate352ObstructionPreserved, a.Verdict)
}

func FormatPMNS(a PMNSSpectralProjectorAudit) string {
	return fmt.Sprintf("Hnu=%q Pe=%q P3=%q overlap=%q trace=%.15g quarter=%.15g labels=%q basis=%q native=%t observed=%t verdict=%q", a.HermitianObject, a.ElectronProjector, a.NeutrinoProjector, a.ProjectorOverlap, a.TraceValue, a.QuarterValue, strings.Join(a.RequiredLabels, ","), a.BasisStatement, a.NativeDerivation, a.ObservedLedger, a.Verdict)
}

func FormatCKM(a CKMSpectralCommutatorAudit) string {
	return fmt.Sprintf("Hu=%q Hd=%q Jexpr=%q comm=%q J=%.15g nondegenerate=%t labels=%q native=%t observed=%t verdict=%q", a.UpHermitian, a.DownHermitian, a.JarlskogArea, a.NormalizedCommutator, a.JValue, a.RequiresNondegenerateSpectra, strings.Join(a.RequiredLabels, ","), a.NativeDerivation, a.ObservedLedger, a.Verdict)
}

func FormatBalance(a BalanceFunctionalAudit) string {
	return fmt.Sprintf("name=%q def=%q expanded=%q kappa=%.15g quarter=%.15g J=%.15g rhs=%.15g B=%.15g delta=%.15g sign=%q insideSigma=%t verdict=%q", a.Name, a.Definition, a.ExpandedDefinition, a.LeftKappa, a.ProjectorQuarter, a.JCKM, a.RightHandCandidate, a.BFlav, a.Delta590, a.SignConvention, a.ResidualInsideSigma, a.Verdict)
}

func FormatInvarianceItem(a InvarianceItem) string {
	return fmt.Sprintf("transformation=%q invariant=%q seal=%q conclusion=%q verdict=%q", a.Transformation, a.InvariantPart, a.RequiredSeal, a.Conclusion, a.Verdict)
}

func FormatInvariance(a InvarianceAudit) string {
	items := make([]string, 0, len(a.Items))
	for _, item := range a.Items {
		items = append(items, FormatInvarianceItem(item))
	}
	return fmt.Sprintf("sealsNamed=%t invariantWithSeals=%t items=[%s] verdict=%q", a.AllRequiredSealsNamed, a.BasisInvariantWithSeals, strings.Join(items, "; "), a.Verdict)
}

func FormatAvailabilityItem(a AvailabilityItem) string {
	return fmt.Sprintf("object=%q observed=%t native=%t forceB0=%t reason=%q verdict=%q", a.Object, a.ObservedLedger, a.NativeOperator, a.CanForceBZero, a.Reason, a.Verdict)
}

func FormatAvailability(a AvailabilityAudit) string {
	items := make([]string, 0, len(a.Items))
	for _, item := range a.Items {
		items = append(items, FormatAvailabilityItem(item))
	}
	return fmt.Sprintf("nativeAlg=%t nativeEps=%t nativePMNS=%t nativeCKM=%t nativeB0=%t items=[%s] verdict=%q", a.NativeFlavorSpectralAlgebra, a.NativeEpsilonFunctional, a.NativePMNSProjector, a.NativeCKMCommutator, a.NativeBFlavZeroTheorem, strings.Join(items, "; "), a.Verdict)
}

func FormatTarget(a OperatorTarget) string {
	return fmt.Sprintf("name=%q domain=%q codomain=%q functional=%q theorem=%q spectral=%t invariant=%t sectorTyping=%t native=%t verdict=%q", a.Name, strings.Join(a.Domain, ","), a.Codomain, a.Functional, a.NativeTheoremRequired, a.MustUseSpectralCalculus, a.MustBeBasisInvariantWithSeals, a.MustPreserveSectorTyping, a.NativePresent, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t pmns=%t ckm=%t yukawas=%t neutrino=%t texture=%t observedNative=%t newCarrier=%t newSelector=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesPMNS, a.DerivesCKM, a.DerivesYukawas, a.DerivesNeutrinoData, a.DerivesFlavorTexture, a.PromotesObservedData, a.AddsNewCarrier, a.AddsNewSelector, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("oneAlgebra=%t labels=%q nativeB=%t environmental=%t theorem=%q decision=%q verdict=%q", a.AllTermsInOneSpectralAlgebra, a.RequiredLabels, a.NativeBFlavOperatorPresent, a.BFlavEnvironmental, a.PromotionTheoremNeeded, a.Decision, a.Verdict)
}
