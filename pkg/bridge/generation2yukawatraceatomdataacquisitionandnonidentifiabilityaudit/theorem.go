package generation2yukawatraceatomdataacquisitionandnonidentifiabilityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-795-YUKAWA-TRACE-ATOM-DATA-ACQUISITION-NON-IDENTIFIABILITY"
	theoremName = "Gate 795 — Yukawa Trace Atom Data Acquisition and Non-Identifiability Audit"
)

func Generation2YukawaTraceAtomDataAcquisitionAndNonIdentifiabilityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 795 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate795}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 794 decomposed trace ledger interface", Passed: a.Gate794.Inherited && a.Gate794.Verdict == StatusGate794Inherited, Detail: a.Gate794.Verdict},
			{Name: "define Yukawa data source hierarchy", Passed: a.Hierarchy.Defined && strings.Contains(a.Hierarchy.HighestAvailable, "aggregate") && containsAll(a.Hierarchy.Priority, []string{"native Yukawa", "singular-value", "sector", "atom", "aggregate"}), Detail: strings.Join(a.Hierarchy.Priority, "; ")},
			{Name: "require acquisition status table", Passed: a.Acquisition.Required && !a.Acquisition.CanPopulate && hasRow(a.Acquisition.Rows, "a_u", StatusMissing) && hasRow(a.Acquisition.Rows, "trace atoms", StatusMissing) && hasRow(a.Acquisition.Rows, "top channel", StatusMissing) && hasRow(a.Acquisition.Rows, "neutrino", StatusAmbiguous), Detail: FormatAcquisition(a.Acquisition)},
			{Name: "execute validation protocol only if decomposed data exists", Passed: a.Validation.ProtocolExecuted && !a.Validation.DataExists && !a.Validation.Validated && containsAll(a.Validation.Rules, []string{"sum sectors", "a^2/b", "sum atoms"}), Detail: strings.Join(a.Validation.Rules, "; ")},
			{Name: "prove aggregate a,b non-identifiability", Passed: a.NonIdentifiability.Proved && a.NonIdentifiability.ConstraintCount == 2 && a.NonIdentifiability.InfiniteFamilies && containsAll(a.NonIdentifiability.CannotIdentify, []string{"sector", "top", "neutrino", "D4"}), Detail: FormatNonIdentifiability(a.NonIdentifiability)},
			{Name: "audit positivity minimum atom count", Passed: a.MinimumAtom.Completed && a.MinimumAtom.MinimumNonzeroAtoms == 4 && a.MinimumAtom.RequiresRestBeyondThree && strings.Contains(a.MinimumAtom.CompatibleReading, "top-color"), Detail: FormatMinimumAtom(a.MinimumAtom)},
			{Name: "compute aggregate top-channel bounds", Passed: a.TopBounds.Computed && closeAbs(a.TopBounds.AOverThree, 0.9474698380779695, 1e-16) && closeAbs(a.TopBounds.SqrtBOverThree, 0.9471025365183062, 1e-16) && closeAbs(a.TopBounds.UpperBoundT, 0.9471025365183062, 1e-16) && !a.TopBounds.DeterminesT, Detail: FormatTopBounds(a.TopBounds)},
			{Name: "record linearized rest pressure estimate", Passed: a.LinearizedRest.Recorded && closeAbs(a.LinearizedRest.DeltaRatio, -0.0002583937062663466, 1e-15) && closeAbs(a.LinearizedRest.AlphaEstimate, 0.0003875905593995199, 5e-16) && !a.LinearizedRest.IsTheorem, Detail: FormatLinearized(a.LinearizedRest)},
			{Name: "gate top-rest decomposition on typed top channel", Passed: a.TopRest.ExecutedIfTExists && !a.TopRest.TypedTFound && !a.TopRest.AlphaBetaComputed && strings.Contains(a.TopRest.FormulaRatio, "beta") && a.TopRest.Verdict == StatusNoTopRestWithoutSelector, Detail: a.TopRest.FormulaDelta},
			{Name: "require explicit neutrino convention", Passed: a.Neutrino.Required && a.Neutrino.Implicit && a.Neutrino.Status == "Y_nu unknown" && a.Neutrino.Verdict == StatusNeutrinoConventionImplicitBlocked, Detail: a.Neutrino.Status},
			{Name: "require scale-locality and multi-scale stability check", Passed: a.Scale.Required && a.Scale.Scale == "M_Z" && !a.Scale.MultiScaleLedger && !a.Scale.ScaleStabilityCertified && a.Scale.Verdict == StatusScaleStabilityRequiresMultiScale, Detail: a.Scale.Verdict},
			{Name: "record C_Higgs impact status", Passed: a.Impact.Recorded && a.Impact.ValidatedAtomLedgerWouldImprove && a.Impact.NEffAggregateSealed && !a.Impact.CHiggsLevelC, Detail: a.Impact.Verdict},
			{Name: "record branch decision", Passed: a.Branch.Recorded && !a.Branch.ValidatedDataFound && !a.Branch.NativeYukawaOperatorsFound && strings.Contains(a.Branch.Recommended, "External Yukawa Ledger") && containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "Native Yukawa"}), Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.AggregateIsAtomLedger && !a.Firewalls.TopDominanceTopYukawaTheorem && !a.Firewalls.MinimumAtomGenerationTheorem && !a.Firewalls.NEffD4Triality && !a.Firewalls.SectorDataNativeYukawa && !a.Firewalls.ValidatedAtomsPMNSCKM && !a.Firewalls.SingleScaleStable && !a.Firewalls.CHiggsLevelC && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate795, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, FormatAcquisition(a.Acquisition), FormatNonIdentifiability(a.NonIdentifiability), FormatMinimumAtom(a.MinimumAtom), FormatTopBounds(a.TopBounds), FormatLinearized(a.LinearizedRest), FormatBranch(a.Branch), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
