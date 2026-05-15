// Package generation2electroweakiksource implements Gate 474:
// Electroweak K-overlap source search / independent IK selector audit.
//
// Gate 473 proved that raw quark masses do not derive I_K=1/2. Gate 474
// therefore audits the remaining plausible Standard-Model-wide channels that
// could independently select the K-axis overlap: the Higgs VEV, electroweak
// gauge couplings, and the lepton/neutrino PMNS-facing sector.
//
// Result: no native electroweak object supplies an I_K selector. The Higgs VEV
// and W/Z gauge couplings are universal/generation-blind in the finite family
// address; they set scale or gauge normalization but do not provide the missing
// family-axis overlap. PMNS/lepton data can be an independent bridge comparator
// only if admitted through the empirical airlock with rank-complete metadata,
// branch tags, and bridge-only status.
package generation2electroweakiksource

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE474-ELECTROWEAK-K-OVERLAP-SOURCE-SEARCH"

	StatusGate473Inherited  = "CONDITIONAL_SUPPORT_GATE473_IK_GAP_INHERITED"
	StatusAuditExecuted     = "CONDITIONAL_SUPPORT_ELECTROWEAK_I_K_SOURCE_AUDIT_EXECUTED"
	StatusFrontierDefined   = "CONDITIONAL_SUPPORT_I_K_SOURCE_FRONTIER_DEFINED"
	StatusFirewallPreserved = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE474_IK_SOURCE_AUDIT"

	StatusFailedHiggsGenerationBlind = "FAILED_ROUTE_HIGGS_VEV_GENERATION_BLIND"
	StatusFailedGaugeGenerationBlind = "FAILED_ROUTE_ELECTROWEAK_GAUGE_COUPLINGS_GENERATION_BLIND"
	StatusFailedPMNSNeedsAirlock     = "FAILED_ROUTE_PMNS_LEPTON_SECTOR_REQUIRES_EMPIRICAL_AIRLOCK"
	StatusFailedNoNativeIKSource     = "FAILED_ROUTE_NATIVE_ELECTROWEAK_GEOMETRY_DOES_NOT_SELECT_I_K"
	StatusFailedIKHalfNotDerived     = "FAILED_ROUTE_I_K_HALF_NOT_DERIVED_FROM_ELECTROWEAK_UNIVERSALS"
	StatusFailedNativePromotion      = "FAILED_ROUTE_ELECTROWEAK_I_K_NATIVE_PROMOTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                  bool
	Gate444KGenForced         bool
	Gate445TriangleForced     bool
	Gate454RankAuditAvailable bool
	Gate465AirlockAvailable   bool
	Gate473MassClosureFailed  bool
	NativeRegistryClean       bool
	MissingIK                 bool
	Verdict                   string
}

type Candidate struct {
	Name                       string
	Channel                    string
	Universal                  bool
	GenerationBlind            bool
	FamilySensitive            bool
	SuppliesScale              bool
	SuppliesGaugeNormalization bool
	SuppliesSpectrum           bool
	SuppliesIK                 bool
	SuppliesBranchTags         bool
	RequiresEmpiricalAirlock   bool
	NativePromotionAttempt     bool
	Verdict                    string
	Reason                     string
}

type SourceSieve struct {
	Executed             bool
	Candidates           []Candidate
	NativeSelectors      int
	BridgeOnlyCandidates int
	IKHalfDerived        bool
	Verdict              string
	Reason               string
	Failures             []string
}

type Frontier struct {
	Executed                                 bool
	RequiredObject                           string
	RequiredFields                           []string
	CanUsePMNSAsBridgeComparator             bool
	CanUseHiggsVEVAsScaleInput               bool
	CanUseGaugeCouplingsAsNormalizationInput bool
	CanUseAnyAsNativeIKSelector              bool
	Verdict                                  string
	Reason                                   string
}

type Firewall struct {
	Executed                      bool
	HiggsVEVNativeIK              bool
	GaugeCouplingsNativeIK        bool
	PMNSNativeIK                  bool
	IKHalfNative                  bool
	CKMNativePrediction           bool
	PMNSNativePrediction          bool
	NativeRegistryWritten         bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Sieve       SourceSieve
	Frontier    Frontier
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	a := Analysis{Inheritance: buildInheritance()}
	a.Sieve = buildSieve()
	a.Frontier = buildFrontier()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate445TriangleForced: true, Gate454RankAuditAvailable: true, Gate465AirlockAvailable: true, Gate473MassClosureFailed: true, NativeRegistryClean: true, MissingIK: true, Verdict: StatusGate473Inherited}
}

func buildSieve() SourceSieve {
	candidates := []Candidate{
		{
			Name: "Higgs vacuum expectation value", Channel: "scalar-vev", Universal: true, GenerationBlind: true, FamilySensitive: false,
			SuppliesScale: true, SuppliesGaugeNormalization: false, SuppliesSpectrum: false, SuppliesIK: false, SuppliesBranchTags: false,
			Verdict: StatusFailedHiggsGenerationBlind,
			Reason:  "the Higgs VEV multiplies Yukawa operators and fixes an electroweak scale, but carries no native family-index projector or K-axis overlap",
		},
		{
			Name: "Electroweak W/Z gauge couplings", Channel: "gauge-normalization", Universal: true, GenerationBlind: true, FamilySensitive: false,
			SuppliesScale: false, SuppliesGaugeNormalization: true, SuppliesSpectrum: false, SuppliesIK: false, SuppliesBranchTags: false,
			Verdict: StatusFailedGaugeGenerationBlind,
			Reason:  "SU(2)_L and U(1)_Y normalization is generation-universal; it can normalize charges but cannot distinguish the three K_gen levels",
		},
		{
			Name: "Lepton/neutrino PMNS-facing sector", Channel: "lepton-bridge-comparator", Universal: false, GenerationBlind: false, FamilySensitive: true,
			SuppliesScale: false, SuppliesGaugeNormalization: false, SuppliesSpectrum: true, SuppliesIK: false, SuppliesBranchTags: false, RequiresEmpiricalAirlock: true,
			Verdict: StatusFailedPMNSNeedsAirlock,
			Reason:  "lepton/neutrino mixing may provide an independent empirical bridge comparator, but observed PMNS data and neutrino masses are not native IK selectors and require the airlock plus branch metadata",
		},
	}
	selectors := 0
	bridgeOnly := 0
	for _, c := range candidates {
		if c.SuppliesIK && !c.RequiresEmpiricalAirlock && !c.NativePromotionAttempt {
			selectors++
		}
		if c.RequiresEmpiricalAirlock {
			bridgeOnly++
		}
	}
	failures := []string{StatusFailedHiggsGenerationBlind, StatusFailedGaugeGenerationBlind, StatusFailedPMNSNeedsAirlock, StatusFailedNoNativeIKSource, StatusFailedIKHalfNotDerived}
	return SourceSieve{Executed: true, Candidates: candidates, NativeSelectors: selectors, BridgeOnlyCandidates: bridgeOnly, IKHalfDerived: false, Verdict: StatusFailedNoNativeIKSource, Reason: "all audited electroweak-wide channels either commute with/give no information about K_gen or require bridge-only empirical import; no native I_K source is present", Failures: failures}
}

func buildFrontier() Frontier {
	return Frontier{
		Executed:                                 true,
		RequiredObject:                           "rank-complete family-sensitive K-overlap comparator independent of quark masses and CKM alignment",
		RequiredFields:                           []string{"sector", "scale", "scheme", "source", "uncertainty", "I_spec", "I_K", "sigma_CP", "n_C3", "bridge_only=true", "native_registry_write=false"},
		CanUsePMNSAsBridgeComparator:             true,
		CanUseHiggsVEVAsScaleInput:               true,
		CanUseGaugeCouplingsAsNormalizationInput: true,
		CanUseAnyAsNativeIKSelector:              false,
		Verdict:                                  StatusFrontierDefined,
		Reason:                                   "future lepton/electroweak tests are allowed only as rank-complete bridge comparators; Higgs/gauge data may normalize scale/conventions but cannot close the K-overlap gap natively",
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, KGenStillForced: a.Inheritance.Gate444KGenForced, XTriangleStillForced: a.Inheritance.Gate445TriangleForced, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate474 records only a no-native-selector audit and a bridge frontier; no Higgs, gauge, PMNS, CKM, I_K, or d_ud value is written to native law-space"}
}

func buildNext() NextStep {
	return NextStep{Gate: 475, Title: "Lepton-sector rank-complete preflight", Reason: "Gate474 finds no native electroweak I_K selector but permits PMNS/lepton data as an independent bridge comparator.", PrimaryTask: "define a fail-closed lepton/neutrino airlock ledger with common scale/scheme, I_spec, I_K, branch tags, and uncertainty before any PMNS-facing residual can run"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate473MassClosureFailed || !a.Inheritance.MissingIK || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate474 inheritance incomplete: %+v", a.Inheritance)
	}
	if !a.Sieve.Executed || len(a.Sieve.Candidates) != 3 || a.Sieve.NativeSelectors != 0 || a.Sieve.IKHalfDerived || a.Sieve.Verdict != StatusFailedNoNativeIKSource {
		return fmt.Errorf("Gate474 source sieve invalid: %+v", a.Sieve)
	}
	for _, c := range a.Sieve.Candidates {
		if c.NativePromotionAttempt || (c.SuppliesIK && !c.RequiresEmpiricalAirlock) {
			return fmt.Errorf("Gate474 falsely admitted native IK selector: %+v", c)
		}
	}
	if !a.Frontier.Executed || !a.Frontier.CanUsePMNSAsBridgeComparator || !a.Frontier.CanUseHiggsVEVAsScaleInput || !a.Frontier.CanUseGaugeCouplingsAsNormalizationInput || a.Frontier.CanUseAnyAsNativeIKSelector || len(a.Frontier.RequiredFields) < 10 {
		return fmt.Errorf("Gate474 frontier incomplete: %+v", a.Frontier)
	}
	if !a.Firewall.Executed || a.Firewall.HiggsVEVNativeIK || a.Firewall.GaugeCouplingsNativeIK || a.Firewall.PMNSNativeIK || a.Firewall.IKHalfNative || a.Firewall.CKMNativePrediction || a.Firewall.PMNSNativePrediction || a.Firewall.NativeRegistryWritten || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate474 firewall violated: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 474 audits the proposed electroweak sources of I_K and finds no native selector. The Higgs VEV supplies a universal scale, not a family K-overlap. Electroweak W/Z couplings supply generation-blind gauge normalization, not a family-axis projector. PMNS/lepton data may be valuable as an independent empirical bridge comparator, but only through the same airlock and rank-complete metadata rules. Therefore I_K=0.5 remains unproven natively, and the 13-moduli firewall stays intact."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t rank_audit=%t airlock=%t gate473_failed=%t missing_IK=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate454RankAuditAvailable, x.Gate465AirlockAvailable, x.Gate473MassClosureFailed, x.MissingIK, x.NativeRegistryClean, x.Verdict)
}
func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("name=%q channel=%s universal=%t generation_blind=%t family_sensitive=%t scale=%t gauge_norm=%t spectrum=%t supplies_IK=%t supplies_branch_tags=%t airlock=%t verdict=%s reason=%s", c.Name, c.Channel, c.Universal, c.GenerationBlind, c.FamilySensitive, c.SuppliesScale, c.SuppliesGaugeNormalization, c.SuppliesSpectrum, c.SuppliesIK, c.SuppliesBranchTags, c.RequiresEmpiricalAirlock, c.Verdict, c.Reason)
}
func FormatSieve(x SourceSieve) string {
	parts := []string{fmt.Sprintf("executed=%t native_selectors=%d bridge_only_candidates=%d IK_half_derived=%t verdict=%s reason=%s", x.Executed, x.NativeSelectors, x.BridgeOnlyCandidates, x.IKHalfDerived, x.Verdict, x.Reason)}
	for _, c := range x.Candidates {
		parts = append(parts, "- "+FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}
func FormatFrontier(x Frontier) string {
	return fmt.Sprintf("executed=%t required_object=%q required_fields=%s PMNS_bridge=%t Higgs_scale=%t gauge_norm=%t native_selector=%t verdict=%s reason=%s", x.Executed, x.RequiredObject, strings.Join(x.RequiredFields, ","), x.CanUsePMNSAsBridgeComparator, x.CanUseHiggsVEVAsScaleInput, x.CanUseGaugeCouplingsAsNormalizationInput, x.CanUseAnyAsNativeIKSelector, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t Higgs_IK_native=%t gauge_IK_native=%t PMNS_IK_native=%t IK_half_native=%t CKM_native=%t PMNS_native=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.HiggsVEVNativeIK, x.GaugeCouplingsNativeIK, x.PMNSNativeIK, x.IKHalfNative, x.CKMNativePrediction, x.PMNSNativePrediction, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}
func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}

func statuses() []string {
	return []string{StatusGate473Inherited, StatusAuditExecuted, StatusFailedHiggsGenerationBlind, StatusFailedGaugeGenerationBlind, StatusFailedPMNSNeedsAirlock, StatusFailedNoNativeIKSource, StatusFailedIKHalfNotDerived, StatusFailedNativePromotion, StatusFrontierDefined, StatusFirewallPreserved}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 474 Registry Audit — Electroweak K-Overlap Source Search\n\n## Verdict\n\n")
	b.WriteString("`" + StatusFailedNoNativeIKSource + "`\n\n")
	b.WriteString("Gate 474 asks whether the missing `I_K` can be selected by an independent electroweak object rather than by quark masses or CKM alignment. The answer is negative at the native-law level: Higgs and gauge channels are universal/generation-blind, while PMNS/lepton data is empirical bridge information requiring an airlock ledger.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Source sieve\n\n" + FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("I_K = alpha / sqrt(alpha^2 + 3)\n")
	b.WriteString("Gate473 result: raw masses -> I_spec only; I_K missing\n")
	b.WriteString("Higgs VEV: universal scale; no family K projector\n")
	b.WriteString("W/Z couplings: generation-universal gauge normalization; no K overlap\n")
	b.WriteString("PMNS/lepton sector: possible independent bridge comparator; not native without airlock metadata and branch tags\n")
	b.WriteString("I_K=0.5 native derivation: not achieved\n")
	b.WriteString("`````\n\n")
	b.WriteString("## Frontier contract\n\n" + FormatFrontier(a.Frontier) + "\n\n")
	b.WriteString("## Firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No Higgs VEV, gauge coupling, PMNS value, CKM value, `I_K`, `alpha`, phase branch, or `d_ud` result is written into native law-space.\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return strings.ReplaceAll(b.String(), "`````", "```")
}
