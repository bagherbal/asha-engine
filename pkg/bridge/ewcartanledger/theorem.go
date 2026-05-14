package ewcartanledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ElectroweakCartanLedgerRetrievalAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-ELECTROWEAK-CARTAN-LEDGER-RETRIEVAL-AUDIT"
	const name = "Electroweak Cartan Ledger Retrieval / Native T3L-Y_phi Coefficient Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 254 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 253 Witt dictionary inherited without weakening", Passed: a.PreviousGate253.WittPairingRetrieved && a.PreviousGate253.NumberSO8Coordinates && a.PreviousGate253.KnownFockLedgersCoordinateReady && !a.PreviousGate253.T3LYPhiSO8Coordinates && !a.PreviousGate253.Q8vCConstructed && !a.PreviousGate253.Neutral3PlaneDerived, Detail: FormatInherited(a.PreviousGate253)},
			{Name: "registry search finds nearby electroweak ledgers", Passed: a.RegistrySearch.BMinusLRetrieved && a.RegistrySearch.NativeU1Retrieved && a.RegistrySearch.TemporalT0Retrieved && a.RegistrySearch.ScalarTPhiRetrieved && a.RegistrySearch.MatterT3RDiagnosticRetrieved && a.RegistrySearch.LeftDoubletT3LRetrieved && a.RegistrySearch.CandidateWeakCartansRetrieved, Detail: FormatRegistrySearch(a.RegistrySearch)},
			{Name: "Fock-number ledgers translate through Witt dictionary", Passed: a.Translation.BMinusLSO8Coordinate && a.Translation.TemporalT0SO8Coordinate && a.Translation.CandidateWeakSO8Coordinate && a.Translation.TranslatedLedgerCount >= 3, Detail: FormatTranslation(a.Translation)},
			{Name: "physical T3L/Y_phi native number ledgers are absent", Passed: !a.RegistrySearch.T3LAsNativeNumberLedger && !a.RegistrySearch.YPhiAsNativeNumberLedger && !a.RegistrySearch.CompleteEWLedgerFound && !a.Translation.T3LSO8Coordinate && !a.Translation.YPhiSO8Coordinate && !a.Translation.QSO8Coordinate, Detail: FormatRegistrySearch(a.RegistrySearch)},
			{Name: "carrier typing rejects T3R/T3L/Y_phi conflation", Passed: a.CarrierTyping.T3LBridgeKnown && a.CarrierTyping.YPhiBridgeKnown && a.CarrierTyping.MatterT3RNumberLedger && a.CarrierTyping.ConflationRejected && !a.CarrierTyping.T3LNumberLedgerFound && !a.CarrierTyping.YPhiNumberLedgerFound && !a.CarrierTyping.T3LDirectSO8Found && !a.CarrierTyping.YPhiDirectSO8Found, Detail: FormatCarrierTyping(a.CarrierTyping)},
			{Name: "candidate weak Cartans are audited but not selected", Passed: len(a.WeakCartans) == 6 && !anySelectedWeakCartan(a.WeakCartans) && countSpatialU1(a.WeakCartans) == 3, Detail: FormatWeakCartans(a.WeakCartans)},
			{Name: "triality branch remains blocked by missing physical weights", Passed: a.Triality.CandidateBranchCount == 2 && !a.Triality.RepresentationWeightsAvailable && !a.Triality.CanSelect8sTo8v && !a.Triality.SelectedByOutcome, Detail: FormatTriality(a.Triality)},
			{Name: "Q8vC and neutral three-plane remain blocked", Passed: !a.Kernel.Q8vCConstructed && !a.Kernel.EigensystemComputed && !a.Kernel.KernelDimensionKnown && !a.Kernel.ExactlyThree && !a.Kernel.ThreePlaneDerived, Detail: FormatKernel(a.Kernel)},
			{Name: "downstream flavor route remains sealed", Passed: !a.Downstream.Neutral3PlaneAvailable && !a.Downstream.VTauConstructed && !a.Downstream.YukawaTextureDerived && !a.Downstream.CKMPMNSDerived && !a.Downstream.FermionMassesDerived, Detail: FormatDownstream(a.Downstream)},
			{Name: "firewall preserved", Passed: !a.Firewall.ImportedSMHyperchargeAsLedger && !a.Firewall.ConflatedT3RWithT3L && !a.Firewall.ConflatedScalarYPhiWithFockY && !a.Firewall.ForcedWeakPlane && !a.Firewall.SelectedTrialityByKernel && !a.Firewall.ForcedKernelDim3 && !a.Firewall.ConstructedVTauByHand && !a.Firewall.InsertedYukawaTexture && !a.Firewall.ImportedObservedMasses && !a.Firewall.PollutedFiniteCore, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records complete ledger search and exact remaining obstruction", Passed: a.Summary.Gate253DictionaryInherited && a.Summary.RegistrySearchCompleted && a.Summary.FockLedgersRetrieved && !a.Summary.T3LNumberLedgerRetrieved && !a.Summary.YPhiNumberLedgerRetrieved && !a.Summary.T3LYPhiSO8Coordinates && !a.Summary.TrialityBranchSelected && !a.Summary.Q8vCConstructed && !a.Summary.Neutral3PlaneDerived && !a.Summary.YukawaTextureDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 254 is a ledger-retrieval theorem, not a physics-fitting step.",
			"The audit retrieves real Fock number ledgers B-L and T0 and translates them to Cartan so(8) coordinates via Gate 253.",
			"The physical T3L/Y_phi pair remains blocked because current project data place T3L on a derived left-doublet carrier and Y_phi on the scalar/contact carrier, not as native N_k ledgers on S_C.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func anySelectedWeakCartan(xs []CandidateWeakCartan) bool {
	for _, x := range xs {
		if x.SelectedPhysicalT3L {
			return true
		}
	}
	return false
}

func countSpatialU1(xs []CandidateWeakCartan) int {
	count := 0
	for _, x := range xs {
		if x.SpatialPreservingU1 {
			count++
		}
	}
	return count
}
