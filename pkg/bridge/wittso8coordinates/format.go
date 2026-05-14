package wittso8coordinates

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate252Audit) string {
	return fmt.Sprintf("trialityCapacity=%t spinorEW=%t spinorSO8=%t trialityMap=%t vectorEW=%t Q=%t plane=%t J=%t vtau=%t unblocked=%t yukawa=%t", a.InfinitesimalTrialityCapacity, a.SpinorEWBridgeKnown, a.SpinorSO8Coordinates, a.ExplicitLieTrialityMap, a.VectorEWMatriciesDerived, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.JCompatibleTransport, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived)
}

func FormatWittBasis(a WittBasisAudit) string {
	return fmt.Sprintf("source=%q realDim=%d modes=%d split=%q retrieved=%t native=%t pairs=%s convention=%q verdict=%q", a.SourceGate, a.RealDimension, a.ComplexModeCount, a.TemporalSpatialSplit, a.Retrieved, a.AllPairsNative, FormatWittPairs(a.Pairs), a.Convention, a.Verdict)
}

func FormatWittPairs(xs []WittPairAudit) string {
	parts := make([]string, len(xs))
	for i, p := range xs {
		parts[i] = fmt.Sprintf("k=%d %s/%s kind=%s plane=%s biv=%s native=%t", p.ModeIndex, p.CreationName, p.AnnihilationName, p.Kind, p.RealPlane, p.Bivector, p.NativePairing)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatNumberOperators(a NumberOperatorExpansionAudit) string {
	return fmt.Sprintf("formula=%q count=%d torus=%d pureAfterShift=%t centralRejected=%t derived=%t coords=%s verdict=%q", a.Formula, a.CoordinateCount, a.MaximalTorusDimension, a.AllPureBivectorAfterShift, a.CentralPartRejectedBySO8, a.Derived, FormatNumberCoordinates(a.Coordinates), a.Verdict)
}

func FormatNumberCoordinates(xs []NumberOperatorCoordinate) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s central=%.6g bivCoeff=%.6g biv=%s lie=%s", x.NumberOperator, x.CentralIdentityShift, x.ImaginaryBivectorCoeff, x.Bivector, x.LieCoordinateFormula)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatFockLedgers(a KnownFockLedgerAudit) string {
	return fmt.Sprintf("B-L=%t T0=%t weakCandidate=%t all=%t ledgers=%s verdict=%q", a.BMinusLCoordinatesDerived, a.TemporalT0CoordinatesDerived, a.WeakPlaneCandidateDerived, a.AllDerivedFromNumberOps, FormatLedgerCoordinates(a.Ledgers), a.Verdict)
}

func FormatLedgerCoordinates(xs []FockLedgerCoordinate) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s expr=%q central=%.6g biv=%s derived=%t physicalEW=%t", x.Name, x.Expression, x.CentralIdentityShift, x.BivectorFormula, x.CoordinateDerived, x.PhysicalEWGenerator)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatElectroweak(a ElectroweakGeneratorCoordinateAudit) string {
	return fmt.Sprintf("requested=%v T3name=%t Yname=%t T3coeff=%t Ycoeff=%t T3so8=%t Yso8=%t Q=%t Z=%t candidates=%v obstruction=%q verdict=%q", a.RequestedGenerators, a.T3LBridgeNameKnown, a.YPhiBridgeNameKnown, a.T3LNumberOperatorCoefficients, a.YPhiNumberOperatorCoefficients, a.T3LSO8CoordinatesDerived, a.YPhiSO8CoordinatesDerived, a.QSO8CoordinatesDerived, a.ZSO8CoordinatesDerived, a.CandidateFockLedgersAvailable, a.Obstruction, a.Verdict)
}

func FormatTriality(a TrialityAudit) string {
	return fmt.Sprintf("group=%q candidates=%d selected=%t wrongChoiceRisk=%t canApplyEW=%t obstruction=%q maps=%s verdict=%q", a.OuterAutomorphismGroup, a.CandidateCount, a.SpecificSpinorToVectorChoiceDerived, a.UsesWrongChoiceRiskAudited, a.CanApplyToPhysicalEW, a.Obstruction, FormatTrialityCandidates(a.Candidates), a.Verdict)
}

func FormatTrialityCandidates(xs []TrialityCartanCandidate) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s orth=%t involutive=%t detAbs1=%t D4=%t selected=%t", x.Name, x.Orthogonal, x.Involutive, x.DetAbsOne, x.MapsD4RootLattice, x.Selected)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatKernel(a Q8VCKernelAudit) string {
	return fmt.Sprintf("def=%q T3=%t Y=%t triality=%t Q=%t eig=%t dimKnown=%t dim=%d exact3=%t plane=%t reason=%q verdict=%q", a.Definition, a.T3LCoordinatesAvailable, a.YPhiCoordinatesAvailable, a.TrialityChoiceAvailable, a.Q8vCConstructed, a.EigensystemComputed, a.KernelDimensionKnown, a.KernelComplexDimension, a.ExactlyThree, a.ThreePlaneDerived, a.DiagnosticOnlyReason, a.Verdict)
}

func FormatDownstream(a DownstreamAudit) string {
	return fmt.Sprintf("plane=%t tau=%v vtau=%t yukawa=%t CKM/PMNS=%t masses=%t verdict=%q", a.Neutral3PlaneAvailable, a.TauEta, a.VTauConstructed, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("witt=%t T3=%t Y=%t trialityByOutcome=%t force3=%t vtau=%t yukawa=%t masses=%t polluted=%t verdict=%q", a.InventedWittPairing, a.InventedT3LCoefficients, a.InventedYPhiCoefficients, a.SelectedTrialityByOutcome, a.ForcedKernelDim3, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.ImportedObservedMasses, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("witt=%t numbers=%t knownLedgers=%t T3Y=%t triality=%t Q=%t plane=%t vtau=%t yukawa=%t status=%q next=%q comment=%q", a.WittPairingRetrieved, a.NumberSO8Coordinates, a.KnownFockLedgersCoordinateReady, a.T3LYPhiSO8Coordinates, a.ExplicitTrialitySelected, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.VTauConstructed, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
