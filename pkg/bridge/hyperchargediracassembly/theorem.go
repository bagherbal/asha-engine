package hyperchargediracassembly

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HyperchargeLedgerSieveCanonicalFiniteDiracAssemblyAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HYPERCHARGE-LEDGER-SIEVE-CANONICAL-FINITE-DIRAC-ASSEMBLY-AUDIT"
	const name = "Hypercharge Ledger Sieve / Canonical Finite Dirac Assembly Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 296 hypercharge/Dirac assembly audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 295 true bimodule is inherited", Passed: a.Input.TrueBimoduleDerived && a.Input.ZeroOrderVerified, Detail: FormatInput(a.Input)},
			{Name: "anomaly/Yukawa/unimodularity equations recover the SM hypercharge ray", Passed: a.Hypercharge.RayRecovered && a.Hypercharge.FractionalLedgerGenerated, Detail: FormatHypercharge(a.Hypercharge)},
			{Name: "absolute U(1) normalization remains un-derived", Passed: !a.Hypercharge.AbsoluteNormalizationFixed, Detail: FormatHypercharge(a.Hypercharge)},
			{Name: "canonical D_F edge graph is assembled as structural adjacency", Passed: a.Summary.DFEdgeGraphAssembled && len(a.Dirac.AllowedEdges) == 4, Detail: FormatDirac(a.Dirac)},
			{Name: "first-order preflight forbids color-changing and lepton-quark edges", Passed: a.FirstOrder.ColorIntertwinerVerified && a.FirstOrder.LeptonQuarkForbiddenByModule && a.FirstOrder.ChargeViolatingEdgesForbidden, Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "full first-order spectral triple and canonical D_F remain unselected", Passed: !a.FirstOrder.FullFirstOrderVerified && !a.Summary.CanonicalDFSelected && !a.Dirac.NumericalYukawas, Detail: FormatSummary(a.Summary)},
			{Name: "B-gap Majorana activation remains sealed", Passed: !a.Dirac.BGapActivated && !a.Dirac.IncludesMajorana, Detail: FormatDirac(a.Dirac)},
			{Name: "firewalls preserve Higgs/B-gap dynamics", Passed: !a.Firewalls.FiniteCorePolluted && a.Firewalls.DoesNotUnlockHiggs && a.Firewalls.DoesNotUnlockBGap, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
