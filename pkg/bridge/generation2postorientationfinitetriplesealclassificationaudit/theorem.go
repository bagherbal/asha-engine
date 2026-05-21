package generation2postorientationfinitetriplesealclassificationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_POST_ORIENTATION_FINITE_TRIPLE_SEAL_CLASSIFICATION_AUDIT"
	theoremName = "Gate 863 — Post-Orientation FiniteTriple Seal Classification Audit"
)

func Generation2PostOrientationFiniteTripleSealClassificationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "classify five-layer post-orientation stack", Passed: layersOK(a.Layers), Detail: FormatLayers(a.Layers)},
			{Name: "separate ambient 16/32 carrier from minimal active 15/30 carrier", Passed: a.Carrier.AmbientPartRank == HPartAmbientRank && a.Carrier.AmbientFullRank == HFAmbientRank && a.Carrier.MinimalPartRank == HPartMinRank && a.Carrier.MinimalFullRank == HFMinRank && a.Carrier.MinimalBranchSealed && !a.Carrier.AmbientBranchNative && containsAll(a.Carrier.Supports, []string{SupportMinimalCarrier, SupportPunctureKernelPair}), Detail: FormatCarrier(a.Carrier)},
			{Name: "classify scalar edge-socket matrix as symbolic post-orientation operator seal", Passed: a.Edge.ColorCentral && a.Edge.ScalarSocket && a.Edge.CharacterMatchedBySeal && a.Edge.FirstOrderConditional && !a.Edge.NumericalYukawa && !a.Edge.NativeFiniteTriple && a.Edge.YRank == YRank && a.Edge.DSymRank == DSymRank && a.Edge.KernelRank == KernelRank && containsAll(a.Edge.Supports, []string{SupportScalarEdgeMatrix, SupportStabilizerFirstOrder}) && containsAll(a.Edge.Failures, []string{FailureSymbolicYNotMagnitude, FailureNoTraceMagnitudeReadout}), Detail: FormatEdge(a.Edge)},
			{Name: "block R3 while identifying YdaggerY trace-magnitude readout as next wound", Passed: a.R3.FiniteBodySealPresent && !a.R3.SectorTraceLedgerPresent && !a.R3.TraceMagnitudeReadoutPresent && a.R3.YDaggerYShapeCandidate && !a.R3.YSocketMagnitudesDerived && !a.R3.AlphaNative && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && containsAll(a.R3.Supports, []string{SupportYdaggerYNext}) && containsAll(a.R3.Failures, []string{FailureNoTraceMagnitudeReadout, FailureNoYToAggTheorem, FailureNoAlphaSource, FailureNoR3}), Detail: FormatR3(a.R3)},
			{Name: "classify branch as post-orientation finite-triple seal, not native theorem", Passed: a.Impact.Classification == Classification && a.Impact.Subtype == Subtype && a.Impact.PostOrientationFiniteTripleSeal && a.Impact.StabilizerFirstOrderConditional && !a.Impact.NativeFiniteTriple && !a.Impact.YukawaMagnitudes && !a.Impact.SectorTraceReadout && !a.Impact.AlphaNative && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatImpact(a.Impact)},
			{Name: "preserve official ledger freeze", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 863 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatLayers(a.Layers), FormatCarrier(a.Carrier), FormatEdge(a.Edge), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func layersOK(layers []Layer) bool {
	seen := map[string]bool{}
	for _, l := range layers {
		seen[l.Name] = true
		if l.Native {
			return false
		}
		if l.Name == "native/full unbroken finite algebra" && (l.SealSuccess || !l.FullUnbroken || !containsAll(l.Failures, []string{FailureCurrentDFNotFullAF})) {
			return false
		}
		if l.Name != "native/full unbroken finite algebra" && !l.SealSuccess {
			return false
		}
	}
	return seen["native/full unbroken finite algebra"] && seen["post-orientation stabilizer algebra"] && seen["minimal finite carrier"] && seen["edge operator"] && seen["first-order compatibility"]
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.CurrentDFNotFullAF && f.SocketIDNotNative && f.HiggsOrientationNotNative && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NoFullUnbrokenFirstOrder && f.SymbolicYNotMagnitude && f.NoNumericalYukawa && f.NoTraceReadout && f.NoAlphaSource && f.NoYToAggTheorem && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoParticleAssignment && f.NoNeutrinoTheorem && f.NoThreeGenerationTheorem && f.NoR3 && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
