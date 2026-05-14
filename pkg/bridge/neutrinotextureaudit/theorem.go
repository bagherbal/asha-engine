package neutrinotextureaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NeutrinoFlavorTextureAuditSealTheorem() theorem.Theorem {
	const id = "BRIDGE-NEUTRINO-FLAVOR-TEXTURE-SEAL-AUDIT"
	const name = "Neutrino flavor texture audit / NeutrinoTextureSeal activation"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 232 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 231 sealed seesaw preflight is inherited", Passed: a.Gate231.Gate231Inherited && a.Gate231.IntermediateSealActive && a.Gate231.SmallYukawaSeesawConditionallyOK && !a.Gate231.FiniteNeutrinoMatrixDerived && a.Gate231.MIntGeV > 0 && a.Gate231.MDirac3GeV > 0, Detail: FormatGate231(a.Gate231)},
			{Name: "NeutrinoTextureSeal is activated only as a phenomenological boundary", Passed: a.Seal.Active && a.Seal.PhenomenologicalBoundary && !a.Seal.FiniteDerived && !a.Seal.DerivesPMNS && !a.Seal.DerivesMassOrdering && !a.Seal.DerivesMajoranaMatrix && !a.Seal.DerivesDiracMatrix, Detail: FormatSeal(a.Seal)},
			{Name: "solar/atmospheric target ratio is used only as a comparison target", Passed: a.Ratio.Ratio > 0.16 && a.Ratio.Ratio < 0.18 && a.Ratio.ToleranceRelative > 0 && a.Firewall.UsesObservedRatioAsTargetOnly, Detail: FormatRatio(a.Ratio)},
			{Name: "direct SM mass hierarchy proxies are too hierarchical", Passed: !a.Audit.AnySMMassProxySupported && a.Audit.BestStandardSMMassProxy.RatioM2ToM3 < 0.07 && a.Audit.BestStandardSMMassProxy.RatioError > 0.5, Detail: FormatCandidate(a.Audit.BestStandardSMMassProxy)},
			{Name: "simple quadratic generation-index texture gives conditional ratio support", Passed: a.Audit.AnyGenerationProxySupported && a.Audit.BestGenerationIndexProxy.Name == "generation-index quadratic" && a.Audit.BestGenerationIndexProxy.RatioWithinTolerance && a.Audit.BestGenerationIndexProxy.RatioM2ToM3 > 0.18 && a.Audit.BestGenerationIndexProxy.RatioM2ToM3 < 0.22 && a.Audit.Verdict == StatusGenerationQuadraticSupport, Detail: FormatCandidate(a.Audit.BestGenerationIndexProxy)},
			{Name: "required second Dirac mass is recorded without claiming a texture derivation", Passed: a.Audit.RequiredM2DiracGeV > 2 && a.Audit.RequiredM2DiracGeV < 3 && a.Audit.RequiredY2 > 0.009 && a.Audit.RequiredY2 < 0.011 && a.Audit.RequiredPowerIndexLaw > 2 && a.Audit.RequiredPowerIndexLaw < 2.4, Detail: FormatAudit(a.Audit)},
			{Name: "finite neutrino mass matrices and PMNS data remain obstructed", Passed: !a.Matrix.RightHandedNeutrinoFieldsDerived && !a.Matrix.DegenerateMajoranaMatrixDerived && !a.Matrix.DiracTextureDerived && !a.Matrix.PMNSMatrixDerived && !a.Matrix.CPPhasesDerived && !a.Matrix.MassOrderingDerived && !a.Matrix.ThreeActiveEigenvaluesDerived && a.Matrix.OnlyRatioPreflightAvailable, Detail: FormatMatrix(a.Matrix)},
			{Name: "firewalls remain closed", Passed: a.Firewall.UsesGate231SealedScale && a.Firewall.ActivatesNeutrinoTextureSeal && !a.Firewall.ClaimsFiniteTexture && !a.Firewall.ClaimsFinitePMNS && !a.Firewall.ClaimsFiniteMajoranaMatrix && !a.Firewall.TunesToObservedMixingAngles && a.Firewall.UsesObservedRatioAsTargetOnly && !a.Firewall.ReopensIntermediateDynamics && !a.Firewall.ReopensPatiSalam && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "Gate 232 supports a mild quadratic generation-index Dirac texture as a ratio-level phenomenological resonance, while rejecting direct SM mass proxies and preserving the finite flavor firewall."}}
	}}
}
