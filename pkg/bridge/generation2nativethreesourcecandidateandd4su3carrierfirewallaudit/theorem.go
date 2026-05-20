package generation2nativethreesourcecandidateandd4su3carrierfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-799-NATIVE-THREE-SOURCE-CANDIDATE-D4-SU3-CARRIER-FIREWALL"
	theoremName = "Gate 799 — Native Three-Source Candidate Ranking and D4/SU3 Carrier Firewall Audit"
)

func Generation2NativeThreeSourceCandidateAndD4SU3CarrierFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 799 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate799}}
		}
		color, okColor := CandidateByName(a.Candidates, "Color SU(3)")
		gj, okGJ := CandidateByName(a.Candidates, "Georgi")
		d4, okD4 := CandidateByName(a.Candidates, "D4")
		a2, okA2 := CandidateByName(a.Candidates, "SU(3) / A2")
		gen, okGen := CandidateByName(a.Candidates, "Generation carrier")
		k7, okK7 := CandidateByName(a.Candidates, "K7")
		checks := []theorem.Check{
			{Name: "inherit Gate 798 three-source status", Passed: a.Gate798.Inherited && a.Gate798.CurrentCertifiedSource == "color-tripled top dominance" && a.Gate798.NotGenerationTheorem && a.Gate798.NotD4Theorem && a.Gate798.NotNativeYukawaTheorem && closeAbs(a.Gate798.NEff, NEffSnapshot, 1e-15) && closeAbs(a.Gate798.CYukawa, CYukawaSnapshot, 1e-15), Detail: StatusGate798Inherited + "; " + StatusCurrentThreeInherited},
			{Name: "define native three-source package requirements", Passed: a.Requirement.Defined && a.Requirement.RejectNoMap && a.Requirement.RejectNoOps && containsAll(a.Requirement.Fields, []string{"typed carrier", "trace/readout", "breaking", "noncircularity"}) && containsAll(a.Requirement.NEffReadout, []string{"a=sum", "b=sum", "N_eff"}) && containsAll(a.Requirement.GJReadout, []string{"R_GJ_3", "R_GJ_2", "R_GJ_1"}), Detail: FormatRequirement(a.Requirement)},
			{Name: "audit color SU3 multiplicity source", Passed: okColor && color.Rank == 1 && strings.Contains(color.TypedSource, "trace") && containsAll(color.Strengths, []string{"a_u=3", "N_eff_top=3"}) && containsAll(color.Limitations, []string{"does not derive y_t", "generation"}) && containsAll(color.Failures, []string{StatusColorNoEigenvalues, StatusColorNoGeneration}), Detail: FormatCandidate(color)},
			{Name: "audit Georgi-Jarlskog Clebsch-three source", Passed: okGJ && gj.Rank == 5 && containsAll(gj.RequiredMap, []string{"multi-scale", "RG", "trace-readout"}) && containsAll(gj.Failures, []string{StatusGJNeedsLedger, StatusGJNoNEffWithoutMap, StatusGJNotGUT}), Detail: FormatCandidate(gj)},
			{Name: "audit D4/Spin8 triality source candidate", Passed: okD4 && d4.Rank == 3 && containsAll(d4.RequiredMap, []string{"D4TrialityCarrierPackage", "real-form", "trace-readout", "breaking"}) && containsAll(d4.Failures, []string{StatusNoD4Package, StatusNoD4TraceMap, StatusCompactSpin8Firewall, StatusTrialityNotGeneration}), Detail: FormatCandidate(d4)},
			{Name: "audit SU3/A2 hexagonal carrier candidate", Passed: okA2 && a2.Rank == 6 && containsAll(a2.RequiredMap, []string{"A2SU3CarrierPackage", "root/weight", "trace-readout"}) && containsAll(a2.Failures, []string{StatusHexMotifNotEvidence, StatusColorNotFlavorSU3, StatusNoA2TraceMap}), Detail: FormatCandidate(a2)},
			{Name: "audit generation count candidate", Passed: okGen && gen.Rank == 4 && containsAll(gen.RequiredMap, []string{"GenerationCarrierPackage", "sector operators", "PMNS/CKM", "trace atom"}) && containsAll(gen.Failures, []string{StatusNoNativeGeneration, StatusGenerationsNoNEff, StatusNoPMNSCKM}), Detail: FormatCandidate(gen)},
			{Name: "audit K7 and projective structural resonances", Passed: okK7 && k7.Rank == 7 && containsAll(k7.Strengths, []string{"K7 Hodge", "projective"}) && containsAll(k7.Failures, []string{StatusK7MinusNotGeneration, StatusNoK7YukawaMap, StatusNoK7PMNSCKMMap, StatusProjectiveNotYukawa, StatusProjectiveNotMixing}), Detail: FormatCandidate(k7)},
			{Name: "rank all three-source candidates", Passed: a.Ranking.Recorded && len(a.Ranking.Ranks) == 7 && containsAll(a.Ranking.Ranks, []string{"1 Color", "2 External", "3 D4", "4 Generation", "5 Georgi", "6 SU(3)", "7 K7"}), Detail: FormatRanking(a.Ranking)},
			{Name: "record methodological branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Recommended, "D4 Triality") && strings.Contains(a.Branch.RecommendationWhy, "Cl(1,7)") && containsAll(a.Branch.EmpiricalPath, []string{"external Yukawa", "sector contributions"}) && containsAll(a.Branch.NativePath, []string{"D4TrialityCarrierPackage", "trace-readout"}) && strings.Contains(a.Branch.ForbiddenPath, "symbolic"), Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.ThreeIsProof && !a.Firewalls.NEffD4Theorem && !a.Firewalls.NEffGenerationTheorem && !a.Firewalls.ColorFullFlavorTheorem && !a.Firewalls.GJGUTTheorem && !a.Firewalls.HexagramEvidence && !a.Firewalls.K7GenerationTheorem && !a.Firewalls.ProjectiveYukawaTheorem && !a.Firewalls.D4WithoutRealFormTheorem && !a.Firewalls.CHiggsLevelC && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate799, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatRequirement(a.Requirement)}
		for _, c := range a.Candidates {
			notes = append(notes, FormatCandidate(c))
		}
		notes = append(notes, FormatRanking(a.Ranking), FormatBranch(a.Branch), a.Final)
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
