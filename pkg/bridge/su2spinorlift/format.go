package su2spinorlift

import (
	"fmt"
	"strings"
)

func FormatPlanes(rows []PlaneAudit) string {
	parts := make([]string, 0, len(rows))
	for _, p := range rows {
		parts = append(parts, fmt.Sprintf("%s modes=%v comp=%v ΛU=%d ΛV=%d doublets=%d doubletDimC=%d singlets=%d singletDimC=%d localH=%t", p.Plane, p.ModeIndices, p.ComplementIndices, p.LambdaUDimC, p.LambdaVDimC, p.DoubletCopies, p.DoubletStateDimC, p.SingletCopies, p.SingletStateDimC, p.LocalHModule))
	}
	return strings.Join(parts, " | ")
}

func FormatLift(a SpinorLiftAudit) string {
	return fmt.Sprintf("carrier=%q dimC=%d dimR=%d input=%q contactMatricesSC=%t candidateLifts=%t planeCount=%d closureResidual=%.3e nativeContact=%t selectedPlane=%t verdict=%s", a.Carrier, a.ComplexDimension, a.RealDimension, a.DerivedContactSU2Input, a.ExplicitContactMatricesOnSC, a.CandidateWedgeLiftsComputed, a.CandidatePlaneCount, a.ClosureResidual, a.NativeIdentifiesContactSU2, a.NativeWeakPlaneSelected, a.Verdict)
}

func FormatDoublets(a DoubletProjectionAudit) string {
	return fmt.Sprintf("planes=%d doubletCopies=%d singletCopies=%d doubletDimC=%d singletDimC=%d SMLeftDimC=%d dimMatch=%t hypercharge=%t colorMultiplicity=%t physicalProjection=%t verdict=%s", a.CandidatePlanes, a.DoubletCopiesPerPlane, a.SingletCopiesPerPlane, a.DoubletStateDimCPerPlane, a.SingletStateDimCPerPlane, a.StandardOneGenerationLeftDimC, a.DimensionalMatchToQLPlusLL, a.HyperchargeAssignmentDerived, a.ColorMultiplicityAssignment, a.PhysicalLeftDoubletProjection, a.Verdict)
}

func FormatQuaternionic(a QuaternionicClosureAudit) string {
	return fmt.Sprintf("pseudoReal=%t localH=%t image=%q closureDimC=%d HdimR=%d singletRemainder=%t globalH=%t planeSelection=%t opposite=%t orderOne=%t verdict=%s", a.FundamentalDoubletPseudoReal, a.LocalQuaternionicStructureOnDoublet, a.AssociativeImagePerSelectedPlane, a.ComplexClosureDimension, a.RealQuaternionicDimension, a.SingletScalarRemainder, a.GlobalHSummandDerived, a.PlaneSelectionRequired, a.OppositeActionDerived, a.OrderOneReady, a.Verdict)
}

func FormatAlgebra(a AlgebraCompletionAudit) string {
	return fmt.Sprintf("prevCplusM3=%t u1C=%t localH=%t globalH=%t exactCplusHplusM3=%t faithfulSC=%t orderOne=%t majoranaSieve=%t verdict=%s", a.PreviousCPlusM3Preflight, a.U1ComplexPreflight, a.LocalHPreflight, a.NativeHGlobalSummand, a.ExactCPlusHPlusM3Derived, a.FaithfulRepresentationOnSC, a.FullOrderOneCalculusReady, a.MajoranaSieveReady, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("pauliImported=%t connesImported=%t weakPlaneForced=%t hyperchargeForced=%t smGaugeInserted=%t bGapMass=%t claimedExactH=%t claimedOrderOne=%t polluted=%t verdict=%s", a.PauliMatricesImportedAsAnswer, a.ConnesAlgebraImported, a.WeakPlaneForced, a.HyperchargeForced, a.SMGaugeGroupInserted, a.BGapPromotedToMass, a.ClaimedExactH, a.ClaimedOrderOne, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("candidateLifts=%t doubletSupport=%t localH=%t nativeContactLift=%t canonicalPlane=%t globalH=%t exactSMAlg=%t status=%s next=%q comment=%q", a.CandidateSU2Lifts, a.DoubletDimensionalSupport, a.PseudoRealLocalHSupport, a.NativeContactLiftDerived, a.CanonicalWeakPlane, a.GlobalHDerived, a.ExactSMAlgebraDerived, a.Status, a.NextGate, a.Comment)
}
