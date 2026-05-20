package generation2chargedleptonsigmadegeneracygaugeorientationaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate602) string {
	return fmt.Sprintf("electron=%t p3=%t plusJ=%t fullSigma=%t minClass=%d best=%.15g next=%.15g verdict=%q", a.SelectsElectronRow, a.SelectsP3, a.SelectsPositiveJ, a.SelectsFullSigma, a.MinimalClassSize, a.BestResidual, a.NextDistinctResidual, a.Verdict)
}
func FormatS3Action(a S3ActionRow) string {
	return fmt.Sprintf("sigma=%q delta=%.15g R=%.15g Q=%.15g epsE=%.15gdeg Vx=%.15g signVx=%+d parity=%s B=%.15g verdict=%q", a.Sigma, a.DeltaDeg, a.R, a.Q, a.ElectronWallEpsilonDeg, a.VandermondeX, a.SignVandermondeX, a.OrientationParity, a.BFlavInvariantValue, a.Verdict)
}
func FormatS3Table(rows []S3ActionRow, n int) string {
	if n > len(rows) {
		n = len(rows)
	}
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		parts = append(parts, FormatS3Action(rows[i]))
	}
	return strings.Join(parts, " | ")
}
func FormatInvariant(a InvariantVsOrientationSensitive) string {
	return fmt.Sprintf("quantity=%q s3Invariant=%t orientationSensitive=%t native=%t explanation=%q verdict=%q", a.Quantity, a.InvariantUnderS3, a.OrientationSensitive, a.Native, a.Explanation, a.Verdict)
}
func FormatInvariants(rows []InvariantVsOrientationSensitive) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatInvariant(r))
	}
	return strings.Join(parts, " | ")
}
func FormatPhysicalLabels(a PhysicalLabelAudit) string {
	return fmt.Sprintf("observed=%q yukawa=%q native=%t environmental=%t verdict=%q", a.ObservedOrdering, a.YukawaOrdering, a.NativeOrdering, a.EnvironmentalOrdering, a.Verdict)
}
func FormatSignedDiscriminant(a SignedDiscriminantAudit) string {
	return fmt.Sprintf("Delta=%q V=%q Vx=%q deltaOnly=%t signedRequiresOrder=%t signDistinguishes=%t nativeSignTheorem=%t seal=%q verdict=%q", a.DeltaFormula, a.VandermondeFormula, a.RootVandermondeFormula, a.TraceRingSuppliesDeltaOnly, a.SignedVRequiresOrdering, a.SignDistinguishesOrientation, a.NativeSignedVTheorem, a.MinimalSeal, a.Verdict)
}
func FormatFourierCyclic(a FourierCyclicOrientationAudit) string {
	return fmt.Sprintf("requiresCycle=%t canonical=%q reversed=%q unsignedWall=%t dependsOnCyclic=%t verdict=%q", a.RequiresCyclicConvention, a.CanonicalCycle, a.ReversedCycle, a.BFlavUsesUnsignedWallDistance, a.BFlavDependsOnCyclicOrientation, a.Verdict)
}
func FormatOrientationCoupling(a PMNSCKMOrientationCouplingAudit) string {
	return fmt.Sprintf("sgnJckm=%+d sgnJpmns=%+d sgnVx=%+d cand1=%q value=%+d cand2=%q value=%+d nativeOperator=%t verdict=%q", a.JCKMSign, a.JPMNSSign, a.CanonicalVxSign, a.Candidate1, a.Candidate1Value, a.Candidate2, a.Candidate2Value, a.TypedASHAOperatorPresent, a.Verdict)
}
func FormatMinimalSeal(a MinimalRemainingSeal) string {
	return fmt.Sprintf("sigmaGauge=%t physicalSeal=%t seal=%q data=%v statement=%q verdict=%q", a.SigmaGaugeForBFlav, a.PhysicalFullOrderingRequiresSeal, a.SealName, a.SealData, a.Statement, a.Verdict)
}
func FormatFirewalls(a Firewalls) string {
	return fmt.Sprintf("koide=%t masses=%t pmns=%t ckm=%t bflavNative=%t carrier=%t selector=%t gate352=%t gate596=%t gate600=%t gate602=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNS, a.DerivesCKM, a.DerivesBFlavZero, a.AddsCarrier, a.AddsSelector, a.PreservesGate352, a.PreservesGate596, a.PreservesGate600, a.PreservesGate602, a.Verdict)
}
