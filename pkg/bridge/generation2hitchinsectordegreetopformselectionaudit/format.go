package generation2hitchinsectordegreetopformselectionaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate649Inheritance) string {
	return fmt.Sprintf("channelInherited=%t support=%t AAA=%t AAB=%t vanishing=%t off=%t slotFormula=%t slotPrimary=%t d_eq_q=%t symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.ChannelAlgebraInherited, x.TwoComponentSupport, x.AAAChannelAudited, x.AABChannelAudited, x.VanishingAudited, x.OffBlockAudited, x.SlotFormulaDerived, x.SlotTheoremPrimary, x.DEqualsQCoincidence, x.FullSymbolicChannelTheorem, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate649FirewallPreserved, x.Verdict)
}

func FormatLedger(x SectorDegreeLedger) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s deg=%s i+=[%s] i-=[%s]", r.Object, r.SectorDegree.String(), r.InteriorPlus, r.InteriorMinus))
	}
	return fmt.Sprintf("top=%s p=%d q=%d A=%s B=%s A21=%t B03=%t verdict=%q rows=%s", x.TopDegree.String(), x.PositiveDim, x.NegativeDim, x.AName, x.BName, x.AHasDegree21, x.BHasDegree03, x.Verdict, strings.Join(parts, "; "))
}

func formatRows(rows []ChannelDegreeRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s/%s: %s total=%s top=%t survive=%t zero=[%t,%t] mechanism=%s", r.Block, r.Channel, r.XYSectors, r.TotalDegree.String(), r.ReachesTopDegree, r.SurvivesByDegree, r.FirstSlotZero, r.SecondSlotZero, r.DegreeMechanism))
	}
	return strings.Join(parts, "; ")
}

func FormatPositive(x PositiveBlockDegreeAudit) string {
	return fmt.Sprintf("survivors=%v AAAOnly=%t mechanism=%q verdict=%q rows=%s", x.SurvivingChannels, x.AAAOnlySurvives, x.SelectionMechanism, x.Verdict, formatRows(x.Rows))
}

func FormatNegative(x NegativeBlockDegreeAudit) string {
	return fmt.Sprintf("survivors=%v AABOnly=%t count=%d mechanism=%q verdict=%q rows=%s", x.SurvivingChannels, x.AABPlacementsOnly, x.AllowedPlacementCount, x.SelectionMechanism, x.Verdict, formatRows(x.Rows))
}

func FormatMixed(x MixedBlockDegreeAudit) string {
	return fmt.Sprintf("any=%t zeroByDegree=%t mechanism=%q verdict=%q rows=%s", x.AnySurvivesByDegree, x.MixedBlockZeroByDegree, x.SelectionMechanism, x.Verdict, formatRows(x.Rows))
}

func FormatSign(x SignNormalizationAudit) string {
	return fmt.Sprintf("supportByDegree=%t finiteSigns=%t equalUnitByDegree=%t needsCalibration=%t missing=%q verdict=%q", x.DegreeRuleCertifiesSupport, x.Gate649CertifiesFiniteSigns, x.EqualUnitWeightCertifiedByDegree, x.RequiresCalibrationIdentity, x.MissingProofObject, x.Verdict)
}

func FormatTheorem(x SymbolicSelectionTheorem) string {
	return fmt.Sprintf("degreeSupported=%t fullSymbolic=%t remaining=%q theorem=%q verdict=%q", x.DegreeSelectionSupported, x.FullSymbolicTheoremCertified, x.RemainingGap, x.CandidateTheorem, x.Verdict)
}

func FormatSlot(x ResultingSlotFormula) string {
	return fmt.Sprintf("p=%d q=%d d=%d norm=%.12g cos=%.12g rho2=%.12g recovers=%t formula=%q verdict=%q", x.PositiveDim, x.NegativeDim, x.SlotMultiplicity, x.NormSquared, x.Cosine, x.ResidualSquared, x.RecoversGate642Angle, x.Formula, x.Verdict)
}

func FormatResonance(x ResonanceAudit) string {
	return fmt.Sprintf("degree=%d q=%d equal=%t interpretation=%q verdict=%q", x.CubicDegree, x.NegativeDim, x.EqualInASHACarrier, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t higgs=%t ckmPmns=%t gauge=%t verdict=%q", x.ClaimsFullSymbolicDegreeTheorem, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
