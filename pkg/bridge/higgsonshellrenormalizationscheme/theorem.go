package higgsonshellrenormalizationscheme

import "github.com/bagherbal/asha-engine/pkg/theorem"

func OnShellRenormalizationSchemePassarinoVeltmanPoleMatchingAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-ON-SHELL-RENORMALIZATION-SCHEME-PASSARINO-VELTMAN-POLE-MATCHING-AUDIT"
	const name = "On-Shell Renormalization Scheme / Passarino-Veltman Pole Matching Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 338 on-shell renormalization scheme audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 337 precision route inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.RequiredRePiGeV2 > 40 && a.Inputs.FiniteRemainderGeV2 > 1000, Detail: FormatInputs(a.Inputs)},
			{Name: "PV self-energy structure formalized", Passed: len(a.PV.Definitions) == 3 && a.PV.PoleEquation != "" && a.PV.SelfEnergy != "", Detail: FormatPVStructure(a.PV)},
			{Name: "finite PV blocks computed on real below-threshold branch", Passed: len(a.Blocks.Blocks) == 4 && allRealFiniteBlocks(a.Blocks), Detail: FormatPVBlocks(a.Blocks)},
			{Name: "renormalization scheme dependency audited", Passed: len(a.Schemes.Lanes) == 3 && hasOnShellAndMSBar(a.Schemes), Detail: FormatSchemes(a.Schemes)},
			{Name: "counterterm target mapped to on-shell ledger", Passed: a.Counterterm.RequiredFiniteRemainder > 1000 && a.Counterterm.RemainderOverTarget > 20, Detail: FormatCounterterm(a.Counterterm)},
			{Name: "geometric alignment does not select IR scheme", Passed: a.Alignment.UVBoundaryFixed && a.Alignment.ContactShapeImmutable && !a.Alignment.IRSchemeSelected && !a.Alignment.NativeCountertermFound, Detail: FormatAlignment(a.Alignment)},
			{Name: "firewalls preserve no exact collider pole-mass claim", Passed: a.Firewalls.NoFullSMCoefficientTable && a.Firewalls.NoNativeCounterterms && a.Firewalls.NoNativeIRScheme && a.Firewalls.NoExactPoleMassClaim, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}

func allRealFiniteBlocks(blocks PVBlockLedger) bool {
	for _, b := range blocks.Blocks {
		if !b.RealBranch || b.A0FiniteGeV2 != b.A0FiniteGeV2 || b.B0Finite != b.B0Finite {
			return false
		}
	}
	return true
}

func hasOnShellAndMSBar(s SchemeAudit) bool {
	hasOS := false
	hasMS := false
	for _, l := range s.Lanes {
		if l.Name == "On-Shell" && l.CanHitTarget && !l.NativeSelected {
			hasOS = true
		}
		if l.Name == "MS-bar" && l.CanHitTarget && !l.NativeSelected {
			hasMS = true
		}
	}
	return hasOS && hasMS
}
