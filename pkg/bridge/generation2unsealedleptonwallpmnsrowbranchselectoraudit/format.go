package generation2unsealedleptonwallpmnsrowbranchselectoraudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate601) string {
	return fmt.Sprintf("observedMinimal=%t selectsP3=%t selectsPlusJ=%t selectsLeptonOrdering=%t unique=%t native=%t minClass=%d verdict=%q", a.ObservedMinimal, a.SelectsP3, a.SelectsPositiveJ, a.SelectsLeptonOrdering, a.UniqueBranchSelector, a.NativeBranchSelector, a.MinimalClassSize, a.Verdict)
}

func FormatDefinition(a BranchRowBalanceDefinition) string {
	return fmt.Sprintf("formula=%q sigma=%q alpha=%q neutrino=%q sign=%q environmental=%t native=%t verdict=%q", a.Formula, a.SigmaDomain, a.AlphaDomain, a.NeutrinoDomain, a.CKMSignDomain, a.Environmental, a.Native, a.Verdict)
}

func FormatWallCandidate(a ChargedLeptonWallCandidate) string {
	return fmt.Sprintf("sigma=%q alpha=%q idx=%d delta=%.15g R=%.15g wall=%.15g epsDeg=%.15g epsRad=%.15g kappa=%.15g positive=%t observed=%t verdict=%q", a.Sigma, a.Alpha, a.ComponentIndex, a.DeltaDeg, a.R, a.WallDeg, a.EpsilonDeg, a.EpsilonRad, a.Kappa, a.PositiveChamber, a.ObservedWall, a.Verdict)
}

func FormatWallCandidates(rows []ChargedLeptonWallCandidate, n int) string {
	if n > len(rows) {
		n = len(rows)
	}
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		parts = append(parts, FormatWallCandidate(rows[i]))
	}
	return strings.Join(parts, " | ")
}

func FormatPMNS(a PMNSRowProjectorOverlap) string {
	return fmt.Sprintf("alpha=%q i=%d projector=%q |U|^2=%.15g L=%.15g verdict=%q", a.Alpha, a.Index, a.Projector, a.UAbs2, a.Li, a.Verdict)
}

func FormatPMNSTable(rows []PMNSRowProjectorOverlap) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatPMNS(r))
	}
	return strings.Join(parts, " | ")
}

func FormatCKMSign(a CKMSign) string {
	return fmt.Sprintf("sign=%+d convention=%q value=%.15g verdict=%q", a.Sign, a.Convention, a.Value, a.Verdict)
}

func FormatCKMSigns(rows []CKMSign) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCKMSign(r))
	}
	return strings.Join(parts, " | ")
}

func FormatBranchRow(a BranchRowBalanceRow) string {
	return fmt.Sprintf("sigma=%q alpha=%q i=%d sJ=%+d kappa=%.15g L=%.15g J=%.15g B=%.15g |B|=%.15g observed=%t", a.Sigma, a.Alpha, a.NeutrinoI, a.CKMSign, a.Kappa, a.Li, a.JTerm, a.BFlav, a.AbsBFlav, a.Observed)
}

func FormatTopBranchRows(rows []BranchRowBalanceRow, n int) string {
	if n > len(rows) {
		n = len(rows)
	}
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		parts = append(parts, FormatBranchRow(rows[i]))
	}
	return strings.Join(parts, " | ")
}

func FormatObservedRank(a ObservedTupleRank) string {
	return fmt.Sprintf("observed=(%s,alpha=%s,i=%d,sJ=%+d) B=%.15g |B|=%.15g rank=%d minimalClass=%d unique=%t summary=%q verdict=%q", a.ObservedSigma, a.ObservedAlpha, a.ObservedNeutrinoI, a.ObservedCKMSign, a.ObservedBFlav, a.ObservedAbsBFlav, a.Rank, a.MinimalClassSize, a.Unique, a.MinimalClassSummary, a.Verdict)
}

func FormatGap(a GapAudit) string {
	return fmt.Sprintf("best=%.15g nextDistinct=%.15g gap=%.15g large=%t meaning=%q verdict=%q", a.BestAbsResidual, a.NextDistinctAbsResidual, a.GapToNextDistinct, a.GapLarge, a.GapMeaning, a.Verdict)
}

func FormatDegeneracy(a DegeneracyLedger) string {
	return fmt.Sprintf("minimalRows=%d alphas=%v neutrinos=%v signs=%v sigmas=%v electron=%t p3=%t plusJ=%t sigmaDegenerate=%t verdict=%q", a.MinimalRows, a.DistinctAlphas, a.DistinctNeutrinoProjectors, a.DistinctCKMSigns, a.DistinctSigmas, a.ElectronRowSelected, a.P3Selected, a.PositiveJSelected, a.SigmaStillDegenerate, a.Verdict)
}

func FormatSelectorVerdict(a SelectorVerdict) string {
	return fmt.Sprintf("observedMinimal=%t electron=%t p3=%t plusJ=%t fullSigma=%t unique=%t native=%t decision=%q verdict=%q", a.ObservedInMinimalClass, a.SelectsElectronRow, a.SelectsThirdNeutrinoProjector, a.SelectsPositiveCKMSign, a.SelectsFullChargedLeptonSigma, a.UniqueSelector, a.NativeSelector, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t masses=%t pmns=%t ckm=%t neutrino=%t flavor=%t bflavNative=%t observedPromotion=%t carrier=%t selector=%t gate352=%t gate596=%t gate600=%t gate601=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNS, a.DerivesCKM, a.DerivesNeutrinoData, a.DerivesFlavor, a.DerivesBFlavZeroNative, a.PromotesObservedData, a.AddsCarrier, a.AddsSelector, a.PreservesGate352, a.PreservesGate596, a.PreservesGate600, a.PreservesGate601, a.Verdict)
}
