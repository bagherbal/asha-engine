package ewcartanledger

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate253Audit) string {
	return fmt.Sprintf("witt=%t numberSO8=%t knownLedgers=%t T3Y=%t triality=%t Q=%t plane=%t vtau=%t yukawa=%t", a.WittPairingRetrieved, a.NumberSO8Coordinates, a.KnownFockLedgersCoordinateReady, a.T3LYPhiSO8Coordinates, a.ExplicitTrialitySelected, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.VTauConstructed, a.YukawaTextureDerived)
}

func FormatRegistrySearch(a RegistrySearchAudit) string {
	return fmt.Sprintf("sources=%s BL=%t nativeU1=%t T0=%t Tphi=%t T3Rdiag=%t T3L=%t weakCartans=%t T3L_N=%t Yphi_N=%t complete=%t mismatch=%t verdict=%q", strings.Join(a.SourcesSearched, " | "), a.BMinusLRetrieved, a.NativeU1Retrieved, a.TemporalT0Retrieved, a.ScalarTPhiRetrieved, a.MatterT3RDiagnosticRetrieved, a.LeftDoubletT3LRetrieved, a.CandidateWeakCartansRetrieved, a.T3LAsNativeNumberLedger, a.YPhiAsNativeNumberLedger, a.CompleteEWLedgerFound, a.CarrierMismatchDetected, a.Verdict)
}

func FormatLedgers(xs []RetrievedLedger) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s source=%q carrier=%q expr=%q coeff=%v central=%.6g biv=%q Nledger=%t ready=%t T3L=%t Yphi=%t hyper=%t obstruction=%q", x.Name, x.Source, x.Carrier, x.Expression, x.NumberCoefficients, x.CentralIdentityShift, x.BivectorFormula, x.NativeNumberOperatorForm, x.CoordinateReady, x.PhysicalT3L, x.PhysicalYPhi, x.PhysicalHypercharge, x.Obstruction)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatWeakCartans(xs []CandidateWeakCartan) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s plane=%s coeff=%v biv=%q spatialU1=%t selected=%t", x.Name, x.Plane, x.NumberCoefficients, x.BivectorFormula, x.SpatialPreservingU1, x.SelectedPhysicalT3L)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatTranslation(a TranslationAudit) string {
	return fmt.Sprintf("translated=%d BL=%t T0=%t weak=%t T3L=%t Yphi=%t Q=%t Z=%t obstruction=%q verdict=%q ledgers=%s", a.TranslatedLedgerCount, a.BMinusLSO8Coordinate, a.TemporalT0SO8Coordinate, a.CandidateWeakSO8Coordinate, a.T3LSO8Coordinate, a.YPhiSO8Coordinate, a.QSO8Coordinate, a.ZSO8Coordinate, a.Obstruction, a.Verdict, FormatLedgers(a.TranslatedLedgers))
}

func FormatCarrierTyping(a CarrierTypingAudit) string {
	return fmt.Sprintf("T3Lknown=%t T3Lcarrier=%q dim=%d T3L_N=%t T3L_so8=%t Yknown=%t Ycarrier=%q dim=%d Y_N=%t Y_so8=%t T3R=%q T3R_N=%t reject=%t verdict=%q", a.T3LBridgeKnown, a.T3LCarrier, a.T3LDimension, a.T3LNumberLedgerFound, a.T3LDirectSO8Found, a.YPhiBridgeKnown, a.YPhiCarrier, a.YPhiDimension, a.YPhiNumberLedgerFound, a.YPhiDirectSO8Found, a.MatterT3RCarrier, a.MatterT3RNumberLedger, a.ConflationRejected, a.Verdict)
}

func FormatTriality(a TrialityBranchAudit) string {
	return fmt.Sprintf("candidates=%d weights=%t T3=%t Y=%t canSelect=%t branch=%q outcome=%t obstruction=%q verdict=%q", a.CandidateBranchCount, a.RepresentationWeightsAvailable, a.T3LWeightsAvailable, a.YPhiWeightsAvailable, a.CanSelect8sTo8v, a.SelectedBranch, a.SelectedByOutcome, a.Obstruction, a.Verdict)
}

func FormatKernel(a Q8VCKernelAudit) string {
	return fmt.Sprintf("def=%q T3=%t Y=%t triality=%t Q=%t eig=%t known=%t dim=%d exact3=%t plane=%t reason=%q verdict=%q", a.Definition, a.T3LCoordinatesAvailable, a.YPhiCoordinatesAvailable, a.TrialityBranchAvailable, a.Q8vCConstructed, a.EigensystemComputed, a.KernelDimensionKnown, a.KernelComplexDimension, a.ExactlyThree, a.ThreePlaneDerived, a.DiagnosticOnlyReason, a.Verdict)
}

func FormatDownstream(a DownstreamAudit) string {
	return fmt.Sprintf("plane=%t tau=%v vtau=%t yukawa=%t CKM=%t masses=%t verdict=%q", a.Neutral3PlaneAvailable, a.TauEta, a.VTauConstructed, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("SMYasLedger=%t T3RasT3L=%t scalarYasFockY=%t weakPlane=%t trialityByKernel=%t kernel3=%t vtau=%t yukawa=%t masses=%t polluted=%t verdict=%q", a.ImportedSMHyperchargeAsLedger, a.ConflatedT3RWithT3L, a.ConflatedScalarYPhiWithFockY, a.ForcedWeakPlane, a.SelectedTrialityByKernel, a.ForcedKernelDim3, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.ImportedObservedMasses, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate253=%t search=%t fock=%t T3ledger=%t Yledger=%t T3Y=%t triality=%t Q=%t plane=%t vtau=%t yukawa=%t status=%q next=%q comment=%q", a.Gate253DictionaryInherited, a.RegistrySearchCompleted, a.FockLedgersRetrieved, a.T3LNumberLedgerRetrieved, a.YPhiNumberLedgerRetrieved, a.T3LYPhiSO8Coordinates, a.TrialityBranchSelected, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.VTauConstructed, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
