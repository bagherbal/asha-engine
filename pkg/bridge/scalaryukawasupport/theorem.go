package scalaryukawasupport

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TensorLiftedScalarFundamentalClassYukawaBilinearSupportTheorem() theorem.Theorem {
	const id = "BRIDGE-TENSOR-LIFTED-SCALAR-FUNDAMENTAL-CLASS-YUKAWA-BILINEAR-SUPPORT-AUDIT"
	const name = "tensor-lifted scalar fundamental class / Yukawa bilinear support audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build tensor-lifted Yukawa support audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 193 finite scalar fundamental class is inherited", Passed: a.Summary.InheritedScalarFundamentalClass && a.ScalarFundamental.Summary.FullMatrixEtaTraceRejected && !a.ScalarFundamental.Firewall.AbsoluteCouplingPromoted, Detail: "tau_eta is available only on the audited scalar support/observable domain; no continuum integral or coupling leaked"},
			{Name: "tensor-lifted matter-scalar support functional is constructed", Passed: a.TensorLift.DoublyGradedFunctionalConstructed && a.TensorLift.UsesScalarFundamentalClass && a.TensorLift.UsesGate25YukawaChannels && !a.TensorLift.ContinuumIntegralImported && !a.TensorLift.YukawaAmplitudeInserted, Detail: FormatTensorLift(a.TensorLift)},
			{Name: "sealed scalar branches have nonzero opposite native tau_eta support", Passed: len(a.ScalarBranches) == 2 && a.ScalarBranches[0].NonzeroSupport && a.ScalarBranches[1].NonzeroSupport && a.ScalarBranches[0].TauEtaProjector*a.ScalarBranches[1].TauEtaProjector < 0, Detail: FormatBranches(a.ScalarBranches)},
			{Name: "all eight Gate-25 Yukawa bilinear channels survive as support", Passed: a.BilinearSupport.ChannelsAudited == 8 && a.BilinearSupport.SupportedChannels == 8 && a.BilinearSupport.UpSupport == 3 && a.BilinearSupport.DownSupport == 3 && a.BilinearSupport.NeutrinoSupport == 1 && a.BilinearSupport.ElectronSupport == 1 && a.BilinearSupport.AllHyperchargeNeutral && a.BilinearSupport.AllColorLeptonValid && a.BilinearSupport.AllSupportNonzero && a.BilinearSupport.SupportOnlyTheorem && !a.BilinearSupport.YukawaAmplitudesDerived && !a.BilinearSupport.MassTermsDerived, Detail: FormatBilinear(a.BilinearSupport)},
			{Name: "eta-signed support balances without becoming an anomaly or mass theorem", Passed: a.Neutrality.TotalEtaSupportBalances && a.Neutrality.UpDownQuarkBalance && a.Neutrality.NeutrinoElectronBalance && a.Neutrality.BLWeightedEtaSupportBalances && a.Neutrality.HyperchargeResidualSumAbs == 0 && !a.Neutrality.AnomalyCancellationTheoremDerived && a.Neutrality.NeutralityPreflightOnly, Detail: FormatNeutrality(a.Neutrality)},
			{Name: "firewall keeps amplitudes, masses, generation textures, mixings, thresholds, couplings, and constants sealed", Passed: a.Firewall.TensorSupportDerived && !a.Firewall.PhysicalYukawaAmplitudesDerived && !a.Firewall.FermionMassesDerived && !a.Firewall.GenerationTextureValuesDerived && !a.Firewall.CKMMatrixDerived && !a.Firewall.PMNSMatrixDerived && !a.Firewall.ObservedMassInputUsed && !a.Firewall.ObservedMixingInputUsed && !a.Firewall.HiggsVEVValueInserted && !a.Firewall.PhysicalScalarVEVAmplitudeDerived && !a.Firewall.SpectralActionEvaluated && !a.Firewall.HeatKernelMatchingDerived && !a.Firewall.ThresholdBetaRowsDerived && !a.Firewall.AbsoluteCouplingPromoted && !a.Firewall.PhysicalConstantsDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records constructive support only", Passed: a.Summary.TestsAudited == 5 && a.Summary.TensorLiftConstructed && a.Summary.EightGate25ChannelsSupported && a.Summary.AllChannelsHaveNonzeroScalarSupport && a.Summary.EtaSignedSupportBalances && a.Summary.OnlySupportNotAmplitude && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 194 closes the support geometry question: the sealed scalar fundamental class can integrate the finite Yukawa incidence support.",
			"It does not compute a Yukawa coupling, fermion mass, Higgs VEV amplitude, generation texture, CKM/PMNS matrix, threshold row, or physical constant.",
		}}
	}}
}
