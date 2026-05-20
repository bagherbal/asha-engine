package generation2koidefouriercirculantphaseaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("mu0=%.12g lambda12=%.12g gate581_dtheta=%.15g gate581_dphi=%.15g source=%q verdict=%q", a.Mu0GeV, a.Lambda12GeV, a.Gate581MZDThetaDeg, a.Gate581MZDPhiDeg, a.RuntimeSource, a.Verdict)
}

func FormatFormula(a FourierFormulaAudit) string {
	return fmt.Sprintf("formula=%q A=%q R=%q delta=%q koide=%q q=%q order=%v convention=%q verdict=%q", a.Formula, a.ADefinition, a.RDefinition, a.DeltaDefinition, a.KoideEquivalence, a.QFromR, a.CanonicalOrder, a.Convention, a.Verdict)
}

func FormatPoint(a FourierPoint) string {
	return fmt.Sprintf("name=%q labels=%v A=%.15g R=%.15g Rminus1=%.15g Q=%.15g Qminus2over3=%.15g deltaDeg=%.15g deltaTurn=%.15g maxReconErr=%.15g verdict=%q", a.Name, a.Labels, a.A, a.PlaneAmplitudeR, a.PlaneAmplitudeResidual, a.Q, a.DeltaQ, a.DeltaDeg, a.DeltaTurn, a.MaxReconstructionError, a.Verdict)
}

func FormatTransport(a PhaseTransportAudit) string {
	return fmt.Sprintf("mzDelta=%.15g lambdaDelta=%.15g drift=%.15g absDrift=%.15g mzRminus1=%.15g lambdaRminus1=%.15g amplitudeTowardOne=%t stable=%t verdict=%q", a.MZDeltaDeg, a.LambdaDeltaDeg, a.SignedDriftDeg, a.AbsDriftDeg, a.MZAmplitudeResidual, a.LambdaAmplitudeResidual, a.AmplitudeMovesTowardOne, a.PhaseStable, a.Verdict)
}

func FormatPermutationPhase(a PermutationPhase) string {
	return fmt.Sprintf("order=%v deltaDeg=%.15g deltaTurn=%.15g nearest=%q nearestDeg=%.15g residual=%.15g", a.Order, a.DeltaDeg, a.DeltaTurn, a.NearestRational, a.NearestRationalDeg, a.RationalResidualDeg)
}

func FormatPermutation(a PermutationAudit) string {
	parts := make([]string, 0, len(a.Phases))
	for _, p := range a.Phases {
		parts = append(parts, FormatPermutationPhase(p))
	}
	return fmt.Sprintf("canonical=%v canonicalDelta=%.15g bestOrder=%v best=%q bestDeg=%.15g bestResidual=%.15g cert=%.15g unique=%t certified=%t phases=[%s] explanation=%q verdict=%q", a.CanonicalOrder, a.CanonicalDeltaDeg, a.BestRationalOrder, a.BestRational, a.BestRationalDeg, a.BestResidualDeg, a.CertificationDeg, a.UniqueWithoutOrdering, a.SimplePhaseCertified, strings.Join(parts, "; "), a.Explanation, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("masses=%t yukawas=%t koide=%t phase=%t ckm=%t pmns=%t generation=%t carrier=%t observedNative=%t gate352=%t verdict=%q", a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesKoide, a.DerivesFourierPhase, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.AddsNewCarrier, a.PromotesObservedAsNative, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q deltaMZ=%.15g deltaLambda=%.15g Rmz=%.15g Rlambda=%.15g stable=%t rational=%t native=%t next=%q verdict=%q", a.SealName, a.CanonicalDeltaMZDeg, a.CanonicalDeltaLambdaDeg, a.FourierAmplitudeMZ, a.FourierAmplitudeLambda, a.PhaseStableInV1, a.SimpleRationalCertified, a.NativeSelectorCertified, a.MinimalNextRequirement, a.Verdict)
}
