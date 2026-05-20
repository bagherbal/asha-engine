package generation2flavorspectralbalancefunctionaltypeadmissibilityaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedBalance) string {
	return fmt.Sprintf("functional=%q kappa=%.15g trace=%.15g quarter=%.15g J=%.15g B=%.15g delta590=%.15g insideSigma=%t verdict=%q", a.Functional, a.Kappa, a.PMNSProjectorTrace, a.PMNSQuarter, a.JCKM, a.BFlav, a.Delta590, a.ResidualInsideSigma, a.Verdict)
}

func FormatTerm(a TermType) string {
	return fmt.Sprintf("name=%q expr=%q inputs=%q labels=%q spectralCalc=%t fractional=%t observed=%t native=%t obstruction=%q verdict=%q", a.Name, a.Expression, strings.Join(a.Inputs, ","), strings.Join(a.RequiresLabels, ","), a.RequiresSpectralCalc, a.RequiresFractional, a.RequiresObservedData, a.NativePresent, a.PrimaryObstruction, a.Verdict)
}

func FormatTermTyping(a TermTypingAudit) string {
	return fmt.Sprintf("complete=%t epsilon={%s} pmns={%s} ckm={%s} verdict=%q", a.Complete, FormatTerm(a.Epsilon), FormatTerm(a.PMNS), FormatTerm(a.CKM), a.Verdict)
}

func FormatAdmissibilityItem(a AdmissibilityItem) string {
	return fmt.Sprintf("object=%q admits=%t native=%t ledger=%t reason=%q verdict=%q", a.Object, a.CurrentASHAAdmits, a.Native, a.EnvironmentalLedger, a.Reason, a.Verdict)
}

func FormatAdmissibility(a AdmissibilityAudit) string {
	items := make([]string, 0, len(a.Items))
	for _, item := range a.Items {
		items = append(items, FormatAdmissibilityItem(item))
	}
	return fmt.Sprintf("poly=%t detPf=%t projectors=%t fourthRoot=%t rootTrace=%t chamber=%t ckm=%t crossSector=%t items=[%s] verdict=%q", a.PolynomialInvariantsAdmitted, a.DeterminantPfaffianAdmitted, a.SpectralProjectorsAdmitted, a.FractionalFourthRootAdmitted, a.RootTraceAdmitted, a.ChamberWallFunctionalAdmitted, a.NormalizedCKMAdmitted, a.CrossSectorEquationAdmitted, strings.Join(items, "; "), a.Verdict)
}

func FormatObstruction(a NativeObstructionAudit) string {
	return fmt.Sprintf("primary=%q epsBlocked=%t pmnsMoreAdmissible=%t ckmMoreAdmissible=%t bZeroBlocked=%t explanation=%q verdict=%q", a.PrimaryObstruction, a.EpsilonHEBlocked, a.PMNSMoreAdmissible, a.CKMMoreAdmissible, a.BFlavZeroBlocked, a.Explanation, a.Verdict)
}

func FormatPromotionRequirement(a PromotionRequirement) string {
	return fmt.Sprintf("requirement=%q why=%q present=%t verdict=%q", a.Requirement, a.WhyNeeded, a.Present, a.Verdict)
}

func FormatRequirements(a PromotionRequirements) string {
	items := make([]string, 0, len(a.Items))
	for _, item := range a.Items {
		items = append(items, FormatPromotionRequirement(item))
	}
	return fmt.Sprintf("allPresent=%t theorem=%q items=[%s] verdict=%q", a.AllPresent, a.ExactMissingTheorem, strings.Join(items, "; "), a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("fitsResiduals=%t koide=%t pmns=%t ckm=%t yukawas=%t neutrino=%t texture=%t ledgersNative=%t carrier=%t selector=%t gate352=%t verdict=%q", a.FitsNewResiduals, a.DerivesKoide, a.DerivesPMNS, a.DerivesCKM, a.DerivesYukawas, a.DerivesNeutrinoPhysics, a.DerivesFlavorTexture, a.PromotesLedgers, a.AddsCarrier, a.AddsSelector, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("envWellDefined=%t primary=%q projectorCommutatorMoreAdmissible=%t nativeBZero=%t theorem=%q decision=%q verdict=%q", a.BFlavEnvironmentalWellDefined, a.PrimaryNativeObstruction, a.ProjectorAndCommutatorMoreAdmissible, a.NativeBFlavZeroTheoremPresent, a.RequiredTheorem, a.Decision, a.Verdict)
}
