package generation2flavorbranchcompatibilityselectoraudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate600) string {
	return fmt.Sprintf("branchSeal=%t epsilonAlgebraic=%t nativeBranch=%t nativeFourthRoot=%t nativeChamber=%t bflavNative=%t verdict=%q", a.BranchSealDefined, a.EpsilonBranchAlgebraic, a.NativeBranchTheorem, a.NativeFourthRoot, a.NativeChamberSelector, a.BFlavNative, a.Verdict)
}

func FormatDefinition(a BalanceDefinition) string {
	return fmt.Sprintf("formula=%q sigma=%q neutrino=%q sign=%q environmental=%t native=%t verdict=%q", a.Formula, a.SigmaDomain, a.NeutrinoDomain, a.CKMSignDomain, a.Environmental, a.Native, a.Verdict)
}

func FormatLeptonBranch(a ChargedLeptonBranch) string {
	return fmt.Sprintf("sigma=%q delta=%.15g R=%.15g eIndex=%d eWall=%.15g epsDeg=%.15g epsRad=%.15g kappa=%.15g positive=%t verdict=%q", a.Sigma, a.DeltaDeg, a.R, a.ElectronIndex, a.ElectronWallDeg, a.EpsilonDeg, a.EpsilonRad, a.Kappa, a.PositiveChamber, a.Verdict)
}

func FormatLeptonBranches(rows []ChargedLeptonBranch) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatLeptonBranch(r))
	}
	return strings.Join(parts, " | ")
}

func FormatPMNS(a PMNSProjectorOverlap) string {
	return fmt.Sprintf("i=%d projector=%q |Uei|^2=%.15g L_i=%.15g verdict=%q", a.Index, a.Projector, a.UeiAbs2, a.Li, a.Verdict)
}

func FormatPMNSTable(rows []PMNSProjectorOverlap) string {
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

func FormatBranchBalanceRow(a BranchBalanceRow) string {
	return fmt.Sprintf("sigma=%q i=%d sJ=%+d kappa=%.15g L=%.15g J=%.15g B=%.15g |B|=%.15g observed=%t", a.Sigma, a.NeutrinoI, a.CKMSign, a.Kappa, a.Li, a.JTerm, a.BFlav, a.AbsBFlav, a.Observed)
}

func FormatTopBranchBalances(rows []BranchBalanceRow, n int) string {
	if n > len(rows) {
		n = len(rows)
	}
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		parts = append(parts, FormatBranchBalanceRow(rows[i]))
	}
	return strings.Join(parts, " | ")
}

func FormatObservedRank(a ObservedBranchRank) string {
	return fmt.Sprintf("observed=(%s,i=%d,sJ=%+d) B=%.15g |B|=%.15g rank=%d minimalClass=%d unique=%t desc=%q verdict=%q", a.ObservedSigma, a.ObservedNeutrinoI, a.ObservedCKMSign, a.ObservedBFlav, a.ObservedAbsBFlav, a.Rank, a.MinimalClassSize, a.Unique, a.MinimalClassDescription, a.Verdict)
}

func FormatGap(a GapAudit) string {
	return fmt.Sprintf("best=%.15g nextDistinct=%.15g gap=%.15g large=%t meaning=%q verdict=%q", a.BestAbsResidual, a.NextDistinctAbsResidual, a.GapToNextDistinct, a.GapLarge, a.GapMeaning, a.Verdict)
}

func FormatSelectorVerdict(a BranchSelectorVerdict) string {
	return fmt.Sprintf("observedMinimal=%t selectsP3=%t selectsPlusJ=%t selectsLeptonOrdering=%t unique=%t native=%t decision=%q verdict=%q", a.ObservedInMinimalClass, a.SelectsNeutrinoThirdProjector, a.SelectsPositiveCKMSign, a.SelectsChargedLeptonOrdering, a.UniqueBranchSelector, a.NativeBranchSelector, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t masses=%t pmns=%t ckm=%t neutrino=%t bflavNative=%t observedPromotion=%t carrier=%t selector=%t gate352=%t gate596=%t gate600=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNS, a.DerivesCKM, a.DerivesNeutrinoData, a.DerivesBFlavZeroNative, a.PromotesObservedData, a.AddsCarrier, a.AddsSelector, a.PreservesGate352, a.PreservesGate596, a.PreservesGate600, a.Verdict)
}
