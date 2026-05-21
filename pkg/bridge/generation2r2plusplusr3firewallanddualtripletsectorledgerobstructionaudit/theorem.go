package generation2r2plusplusr3firewallanddualtripletsectorledgerobstructionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-831-R2PLUSPLUS-R3-FIREWALL-DUAL-TRIPLET-SECTOR-LEDGER-OBSTRUCTION"
	theoremName = "Gate 831 — R2++ / R3 Firewall and Dual-Triplet Sector Ledger Obstruction Audit"
)

func Generation2R2PlusPlusR3FirewallAndDualTripletSectorLedgerObstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 831 R2++/R3 firewall audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit sealed aggregate operator and classify R2++ status", Passed: a.Ledger.R2PlusPlusConsolidated && a.Ledger.AlphaSealed && !a.Ledger.R3SectorLedgerCertified && math.Abs(a.Ledger.OperatorNEff-OperatorNEff) < 5e-16 && containsAll(a.Impact.Verdicts, []string{StatusGate829830Inherited, StatusR2PlusPlusClassified}), Detail: FormatLedger(a.Ledger)},
			{Name: "audit I_3 top block without promoting it to generations or triality", Passed: a.TopBlock.SourceTypeAudited && a.TopBlock.Rank == TopBlockDim && !a.TopBlock.IsGenerationTheorem && !a.TopBlock.IsD4TrialityTheorem && !a.TopBlock.IsThreeYukawaFamilies && containsAll(a.TopBlock.Failures, []string{FailureTopBlockNotGenerationTheorem, FailureTopBlockNotD4Triality, FailureTopBlockNotThreeFamilies}), Detail: FormatTopBlock(a.TopBlock)},
			{Name: "audit Fock B-L rest block without promoting it to sector ledger", Passed: a.RestBlock.SourceTypeAudited && a.RestBlock.P1Rank == FockP1Rank && a.RestBlock.P3Rank == FockP3Rank && a.RestBlock.CarrierDim == RestBlockDim && !a.RestBlock.IsSMSectorAssignment && !a.RestBlock.IsObservedFlavorHierarchy && !a.RestBlock.IsYukawaSectorLedger && containsAll(a.RestBlock.Failures, []string{FailureFockSelectorNotSectorLedger, FailureRestBlockNotFlavorHierarchy}), Detail: FormatRestBlock(a.RestBlock)},
			{Name: "separate the two triplets despite equal dimension", Passed: a.DualTriplet.SameDimension && !a.DualTriplet.TypedMapCertified && !a.DualTriplet.Identified && containsAll(a.DualTriplet.Failures, []string{FailureColorTripletNotFockTriplet, FailureNoSectorTraceLedgerMap}), Detail: FormatDualTriplet(a.DualTriplet)},
			{Name: "audit seven-count resonance without identifying K7", Passed: a.Seven.CountMatchesK7 && a.Seven.ClassifiedAsResonanceOnly && !a.Seven.ProjectorTheoremCertified && !a.Seven.AggregateToK7MapCertified && containsAll(a.Seven.Failures, []string{FailureSevenAtomsNotK7, FailureNoAggregateToK7Map}), Detail: FormatSeven(a.Seven)},
			{Name: "state R3 sector-ledger requirements and reject promotion", Passed: a.Requirements.RequiresTypedSectorProjectors && a.Requirements.RequiresPositiveTraceAtoms && a.Requirements.RequiresCarrierCompatibility && a.Requirements.RequiresFiniteAlgebraCommutation && a.Requirements.RequiresNonCircularAssignment && a.Requirements.RequiresNoObservedYukawaFit && a.Requirements.RequiresReadoutMap && !a.Requirements.RequirementsSatisfied && !a.Requirements.SectorLedgerCertified && containsAll(a.Requirements.Failures, []string{FailureNoSectorTraceLedgerMap, FailureR2NotR3, FailureNoNativeYukawaOperator, FailureNoObservedYukawaFit}), Detail: FormatRequirements(a.Requirements)},
			{Name: "block all ledger updates and native Yukawa promotion", Passed: !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.NextMissingObject, "SectorTraceLedgerMap") && containsAll(a.Impact.Failures, []string{FailureAlphaStillSealed, FailureNoBoundaryAlphaMap, FailureNoSectorTraceLedgerMap, FailureNoNEffUpdate, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve R2++/R3 and physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.TopNotGeneration && a.Firewalls.TopNotD4 && a.Firewalls.TopNotFamilies && a.Firewalls.FockNotSectorLedger && a.Firewalls.DualTripletSeparated && a.Firewalls.SevenNotK7 && a.Firewalls.NoSectorLedgerMap && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoPMNSCKM && a.Firewalls.NoSectorAssignment && a.Firewalls.Verdict == StatusFirewallGate831, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatTopBlock(a.TopBlock), FormatRestBlock(a.RestBlock), FormatDualTriplet(a.DualTriplet), FormatSeven(a.Seven), FormatRequirements(a.Requirements), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
