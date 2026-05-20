package generation2paulihopfscalarmomentmapaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedBoundaryAudit) string {
	return fmt.Sprintf("gate558Eta=%t gate5582plus2=%t gate559NoLinear=%t gate559NoTraceRank=%t verdict=%q", a.Gate558EtaIsSigma3Candidate, a.Gate558HPhiSplitTwoPlusTwo, a.Gate559NoLinearTransfer, a.Gate559NoTraceRankTransfer, a.Verdict)
}

func FormatScalar(a ScalarComplexStructureAudit) string {
	return fmt.Sprintf("basis=%q realDim=%d complexDim=%d coords=%v identity=%t J=%q Jmat=%s sealed=%t nativeUnsealed=%t verdict=%q", a.BasisName, a.RealDimension, a.ComplexDimension, a.Coordinates, a.IdentityCertified, a.ComplexStructureName, a.ComplexStructure, a.SealedCarrierOnly, a.NativeUnsealed, a.Verdict)
}

func FormatPauli(a PauliMatrixAudit) string {
	return fmt.Sprintf("%s M=%s symmetric=%t squareResidual=%.3g trace=%.3g rank=%d spectrum=%q nativeUnsealed=%t sealedConstructible=%t verdict=%q", a.Name, a.MatrixString, a.Symmetric, a.SquareResidual, a.Trace, a.Rank, a.SpectrumSummary, a.NativeUnsealed, a.ConstructibleSealed, a.Verdict)
}

func FormatPauliList(xs []PauliMatrixAudit) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = FormatPauli(x)
	}
	return strings.Join(parts, "; ")
}

func FormatRelations(a PauliRelationsAudit) string {
	return fmt.Sprintf("matrices=%t squares=%t maxSquareResidual=%.3g maxAntiResidual=%.3g signature=%q nativeUnsealed=%t sealed=%t verdict=%q", a.MatricesAvailable, a.SquaresIdentity, a.MaxSquareResidual, a.MaxAnticommutatorResidual, a.CliffordSignature, a.NativeUnsealed, a.ConstructedUnderScalarSeal, a.Verdict)
}

func FormatSample(a MomentSample) string {
	return fmt.Sprintf("%s x=%v r2=%.6g mu=%s", a.Name, a.X, a.R2, formatFloatVec(a.Mu))
}

func FormatSamples(xs []MomentSample) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = FormatSample(x)
	}
	return strings.Join(parts, "; ")
}

func FormatMoment(a MomentCoordinateAudit) string {
	return fmt.Sprintf("r2=%q mu1=%q mu2=%q mu3=%q samples=[%s] match=%t verdict=%q", a.R2Formula, a.Mu1Formula, a.Mu2Formula, a.Mu3Formula, FormatSamples(a.SamplePoints), a.CoordinatesMatch, a.Verdict)
}

func FormatHopf(a HopfIdentityAudit) string {
	return fmt.Sprintf("identity=%q sampleResidualMax=%.3g projector=%q reconstruction=%q verified=%t sealedC2=%t verdict=%q", a.IdentitySymbolic, a.SampleResidualMax, a.ProjectorIdentity, a.PhiPhiDaggerReconstruction, a.IdentityVerified, a.ReliesOnSealedScalarC2Carrier, a.Verdict)
}

func FormatScalarSplit(a ScalarDecompositionAudit) string {
	return fmt.Sprintf("map=%q domain=%q codomain=%q radius=%q moment=%q 4to1plus3=%t gauge=%t W=%t weakIso=%t flavor=%t verdict=%q", a.MapName, a.Domain, a.Codomain, a.RadiusComponent, a.MomentTripletComponent, a.ScalarSectorFourToOnePlus3, a.IdentifiesGaugeBosons, a.IdentifiesWSpatial, a.IdentifiesWeakIsospin, a.IdentifiesFlavor, a.Verdict)
}

func FormatOrbit(a MomentOrbitSplitAudit) string {
	return fmt.Sprintf("nonzero=%t split=%q radialLine=%t planeGivenMu=%t scalarOnly=%t selectsWPlane=%t selectsGenPlane=%t verdict=%q", a.NonzeroMomentCondition, a.Split, a.RadialLineCanonical, a.OrthogonalPlaneCanonicalGivenMu, a.ScalarSectorOnly, a.SelectsWSpatialWeakPlane, a.SelectsGenerationPlane, a.Verdict)
}

func FormatEtaRelation(a EtaRelationAudit) string {
	return fmt.Sprintf("etaSigma3=%t O1=%q res=%.3g O2=%q res=%.3g O3=%q res=%.3g tau=%s axisShadow=%t pauliTriplet=%t promotedSpectrum=%t verdict=%q", a.EtaEqualsSigma3, a.O1Expression, a.O1Residual, a.O2Expression, a.O2Residual, a.O3Expression, a.O3Residual, formatFloatVec(a.TauEtaTraceList), a.Sigma3AxisShadowOnly, a.LargerPauliTripletAvailable, a.TauEtaPromotedToSpectrum, a.Verdict)
}

func FormatTransfer(a TransferFirewallAudit) string {
	return fmt.Sprintf("triplet=%t toW=%t toWeakPlanes=%t toGen=%t weakPlane=%t hierarchy=%t yukawa=%t ckm=%t observed=%t gaugeBoson=%t allowed=%t verdict=%q", a.PauliMomentTripletAvailable, a.FunctorToWSpatial, a.FunctorToWeakPlaneCandidates, a.FunctorToGeneration, a.WeakPlaneSelected, a.GenerationHierarchyDerived, a.YukawaTextureDerived, a.CKMPMNSDerived, a.ObservedFlavorImported, a.GaugeBosonIdentification, a.TransferAllowed, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("pauli=%t hopf=%t scalar4to1plus3=%t moment3to1plus2=%t etaAxis=%t transfer=%t next=%q verdict=%q", a.SealedPauliTripletExists, a.HopfMomentIdentityHolds, a.ScalarFourToOnePlusThree, a.NonzeroMomentThreeToOnePlusTwo, a.EtaIsSigma3Axis, a.LawfulTransferToWOrGeneration, a.MissingNextTheorem, a.Verdict)
}
