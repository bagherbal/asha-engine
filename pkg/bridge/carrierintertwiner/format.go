package carrierintertwiner

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate254Audit) string {
	return fmt.Sprintf("gate253=%t search=%t fock=%t T3ledger=%t Yledger=%t T3Yso8=%t triality=%t Q=%t plane=%t vtau=%t yukawa=%t status=%q", a.Gate253DictionaryInherited, a.RegistrySearchCompleted, a.FockLedgersRetrieved, a.T3LNumberLedgerRetrieved, a.YPhiNumberLedgerRetrieved, a.T3LYPhiSO8Coordinates, a.TrialityBranchSelected, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.VTauConstructed, a.YukawaTextureDerived, a.Status)
}

func FormatCarrierObject(x CarrierObject) string {
	return fmt.Sprintf("%s dim=%d type=%q source=%q available=%t ownCanonical=%t matter=%t scalar=%t contact=%t subSC=%t embSC=%t projSC=%t intertwines=%t coords=%t obstruction=%q verdict=%q", x.Name, x.Dimension, x.CarrierType, x.Source, x.Available, x.CanonicalOnOwnCarrier, x.ActsOnMatter, x.ActsOnScalar, x.ActsOnContact, x.SubspaceOfSC, x.EmbeddingIntoSC, x.ProjectionFromSC, x.IntertwiningVerified, x.CoordinateComplete, x.Obstruction, x.Verdict)
}

func FormatCarriers(a CarrierInventoryAudit) string {
	parts := make([]string, len(a.ObjectsAudited))
	for i, x := range a.ObjectsAudited {
		parts[i] = FormatCarrierObject(x)
	}
	return fmt.Sprintf("SC=%t dim=%d T3L=%t/%d Yphi=%t/%d HphiSubSC=%t leftSubSC=%t commonSC=%t commonAction=%t verdict=%q objects=[%s]", a.SCAvailable, a.SCDimension, a.T3LAvailable, a.T3LDimension, a.YPhiAvailable, a.YPhiDimension, a.HphiSubspaceOfSC, a.LeftDoubletSubspaceOfSC, a.CommonSCCarrierAvailable, a.CommonActionCarrierAvailable, a.Verdict, strings.Join(parts, "; "))
}

func FormatIntertwinerCandidate(x IntertwinerCandidate) string {
	return fmt.Sprintf("%s from=%q to=%q source=%q avail=%t canonical=%t branchFree=%t injective=%t surjImage=%t T3=%t Yphi=%t intoSC=%t fromSC=%t isom=%t orientation=%t observed=%t gaugeFrame=%t rejected=%q verdict=%q", x.Name, x.From, x.To, x.Source, x.Available, x.Canonical, x.BranchFree, x.Injective, x.SurjectiveOntoImage, x.IntertwinesT3L, x.IntertwinesYPhi, x.MapsIntoSC, x.MapsFromSC, x.Isometric, x.UsesOrientationSeal, x.UsesObservedInput, x.RequiresGaugeFrame, x.RejectedReason, x.Verdict)
}

func FormatIntertwiners(a IntertwinerSearchAudit) string {
	parts := make([]string, len(a.Candidates))
	for i, x := range a.Candidates {
		parts[i] = FormatIntertwinerCandidate(x)
	}
	return fmt.Sprintf("count=%d available=%d canonical=%d SC=%d T3=%d Yphi=%d joint=%d lawful=%t formalRejected=%d verdict=%q candidates=[%s]", a.CandidateCount, a.AvailableCandidates, a.CanonicalCandidates, a.SCEmbeddingCandidates, a.T3LIntertwiningCandidates, a.YPhiIntertwiningCandidates, a.JointIntertwiningCandidates, a.LawfulCommonIntertwiner, a.RejectedFormalAssemblies, a.Verdict, strings.Join(parts, "; "))
}

func FormatUnifiedLedger(a UnifiedLedgerAudit) string {
	return fmt.Sprintf("carrier=%q common=%t T3proj=%t Yproj=%t T3coeff=%t Ycoeff=%t T3=%v Y=%v unified=%t obstruction=%q verdict=%q", a.CommonCarrier, a.CommonCarrierDerived, a.T3LProjectedToSC, a.YPhiProjectedToSC, a.T3LNumberCoefficientsAvailable, a.YPhiNumberCoefficientsAvailable, a.T3LNumberCoefficients, a.YPhiNumberCoefficients, a.UnifiedLedgerConstructed, a.Obstruction, a.Verdict)
}

func FormatSO8(a SO8TranslationAudit) string {
	return fmt.Sprintf("witt=%t T3ledger=%t Yledger=%t T3so8=%t Yso8=%t Q=%t Z=%t obstruction=%q verdict=%q", a.WittDictionaryAvailable, a.T3LUnifiedLedger, a.YPhiUnifiedLedger, a.T3LSO8Coordinates, a.YPhiSO8Coordinates, a.QSO8Coordinates, a.ZSO8Coordinates, a.Obstruction, a.Verdict)
}

func FormatTrialityKernel(a TrialityKernelAudit) string {
	return fmt.Sprintf("trialityCandidates=%t weights=%t branch=%t Q=%t eig=%t known=%t dim=%d exact3=%t plane=%t reason=%q verdict=%q", a.TrialityCandidatesKnown, a.RepresentationWeightsKnown, a.PhysicalBranchSelected, a.Q8vCConstructed, a.EigensystemComputed, a.KernelDimensionKnown, a.KernelComplexDimension, a.ExactlyThree, a.NeutralThreePlaneDerived, a.DiagnosticOnlyReason, a.Verdict)
}

func FormatDownstream(a DownstreamAudit) string {
	return fmt.Sprintf("plane=%t tau=%v vtau=%t trialityTexture=%t yukawa=%t CKM=%t masses=%t verdict=%q", a.Neutral3PlaneAvailable, a.TauEta, a.VTauConstructed, a.TrialityTextureOpened, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("HphiByDim=%t leftByLabel=%t tensorAsSC=%t directSum=%t Connes=%t SMY=%t weakPlane=%t trialityByKernel=%t kernel3=%t vtau=%t yukawa=%t masses=%t polluted=%t verdict=%q", a.EmbeddedHphiIntoSCByDimension, a.EmbeddedLeftDoubletByLabel, a.TreatedTensorProductAsSC, a.TreatedDirectSumAsIntertwiner, a.ImportedConnesRepresentation, a.InsertedSMHyperchargeConvention, a.ForcedWeakPlane, a.SelectedTrialityByKernel, a.ForcedKernelDim3, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.ImportedObservedMasses, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate254=%t SC=%t local=%t common=%t intertwiner=%t ledger=%t so8=%t triality=%t Q=%t plane=%t vtau=%t yukawa=%t status=%q next=%q comment=%q", a.Gate254Inherited, a.SCCarrierKnown, a.LocalActionsAudited, a.CommonCarrierDerived, a.CarrierIntertwinerDerived, a.UnifiedLedgerConstructed, a.T3LYPhiSO8Coordinates, a.TrialityBranchSelected, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.VTauConstructed, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
