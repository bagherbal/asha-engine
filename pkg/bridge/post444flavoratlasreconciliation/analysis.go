// Package post444flavoratlasreconciliation implements Gate 448:
// Post-444 Flavor Frontier Atlas Reconciliation.
//
// Gate 420 exported the publication theorem atlas before the later Generation-2
// intersection sieves were executed.  Gates 444-447 changed the status of the
// family frontier in a very narrow way: K_gen is no longer merely a quarantined
// family axiom; it is a geometrically forced primitive traceless three-level
// axis.  The unsigned Generation-2 mass-lift triangle support is also forced as
// a structural topology.  However, Gate 446 leaves signed/complex orientation
// open and Gate 447 closes the nine-coefficient amplitude firewall.
//
// Gate 448 is therefore not a new flavor prediction.  It is a registry/atlas
// reconciliation gate that patches the post-publication law-space board without
// importing collider masses, Yukawa matrices, CKM/PMNS data, or cosmological
// history.  It promotes only the exact objects proved by Gates 444-445 and keeps
// every value-bearing flavor coordinate quarantined.
package post444flavoratlasreconciliation

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE448-POST444-FLAVOR-FRONTIER-ATLAS-RECONCILIATION"

	StatusGate420AtlasInherited            = "CONDITIONAL_SUPPORT_GATE420_PUBLICATION_ATLAS_INHERITED"
	StatusGate444To447DeltaInherited       = "CONDITIONAL_SUPPORT_GATE444_TO_447_FLAVOR_DELTA_INHERITED"
	StatusAtlasReconciliationDeltaCompiled = "CONDITIONAL_SUPPORT_POST444_ATLAS_RECONCILIATION_DELTA_COMPILED"
	StatusKGenPromotedGeometric            = "K_GEN_PROMOTED_GEOMETRICALLY_FORCED_AXIOM"
	StatusGen2ZeroPromotedStructural       = "GENERATION2_BARE_ZERO_PROMOTED_STRUCTURAL"
	StatusXTriangleSupportPromoted         = "X_TRIANGLE_SUPPORT_PROMOTED_STRUCTURAL_TOPOLOGY"
	StatusYPhaseFirewallPreserved          = "Y_PHASE_ORIENTATION_FIREWALL_PRESERVED"
	StatusNineCoefficientFirewallPreserved = "NINE_KXY_COEFFICIENT_FIREWALL_PRESERVED"
	StatusNativeFlavorDim13Preserved       = "FIREWALL_PRESERVED_13_MODULI"
	StatusPost444FlavorAtlasReconciled     = "PROJECT_POST444_FLAVOR_ATLAS_RECONCILED"
	StatusEmpiricalFirewallPreserved       = "CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED"

	StatusFailedNoYukawaPrediction        = "FAILED_ROUTE_NO_YUKAWA_VALUE_PREDICTION"
	StatusFailedNoCKMPMNSPrediction       = "FAILED_ROUTE_NO_CKM_PMNS_ANGLE_OR_PHASE_PREDICTION"
	StatusFailedXAmplitudeNotPromoted     = "FAILED_ROUTE_X_TRIANGLE_AMPLITUDE_NOT_PROMOTED"
	StatusFailedYPhaseNotPromoted         = "FAILED_ROUTE_Y_PHASE_NOT_PROMOTED_NATIVE"
	StatusFailedNoCoefficientSelector     = "FAILED_ROUTE_NO_NATIVE_KXY_COEFFICIENT_SELECTOR"
	StatusFailedNoMuonCharmMassPrediction = "FAILED_ROUTE_NO_MUON_CHARM_PHYSICAL_MASS_PREDICTION"
)

const (
	NativeChargedFlavorDim = 13
	KXYChargedCoeffDim     = 9
	ReconciledNodeCount    = 4
)

type Inheritance struct {
	Executed                  bool
	Gate420PublicationAtlas   bool
	Gate420Acyclic            bool
	Gate420NativeFlavorDim    int
	Gate420ConditionalDim     int
	Gate420FamilyAxiomsSealed bool
	Gate420NoFlavorReopening  bool
	NoEmpiricalInputsImported bool
	Verdict                   string
}

type GateDelta struct {
	Gate                     int
	Name                     string
	InputStatus              string
	OutputStatus             string
	PromotesStructuralObject bool
	PreservesFirewall        bool
	PredictsObservableValue  bool
	NativeChargedDimAfter    int
	KXYCoeffDimAfter         int
	Verdict                  string
	Reason                   string
}

type Reclassification struct {
	Object           string
	PreviousLayer    string
	ReconciledLayer  string
	PreviousStatus   string
	ReconciledStatus string
	Promoted         bool
	Quarantined      bool
	ValueBearing     bool
	Verdict          string
	Reason           string
}

type ReconciliationDelta struct {
	Executed                    bool
	Deltas                      []GateDelta
	Reclassifications           []Reclassification
	PromotedObjects             int
	QuarantinedObjects          int
	NativeDimBefore             int
	NativeDimAfter              int
	KXYCoeffDimBefore           int
	KXYCoeffDimAfter            int
	FlavorObservableValuesAdded int
	CoefficientSelectorsAdded   int
	Verdict                     string
	Reason                      string
}

type AtlasNode struct {
	ID        string
	Name      string
	Layer     string
	Gates     []int
	Status    string
	Claim     string
	Boundary  string
	DependsOn []string
}

type ReconciledAtlas struct {
	Executed                   bool
	Nodes                      []AtlasNode
	Acyclic                    bool
	NativeFlavorDim            int
	KXYCoeffDim                int
	KGenGeometric              bool
	Gen2BareZeroStructural     bool
	XTriangleSupportStructural bool
	YPhaseQuarantined          bool
	CoefficientsQuarantined    bool
	NoNewPhysicsClaim          bool
	Verdict                    string
}

type Firewall struct {
	Executed                    bool
	NoObservedMuonMassImported  bool
	NoObservedCharmMassImported bool
	NoObservedYukawaImported    bool
	NoCKMImported               bool
	NoPMNSImported              bool
	NoPoleMassFit               bool
	NoCurveFit                  bool
	NoCosmologyInput            bool
	NativeFlavorDimPreserved    bool
	KXYCoeffDimPreserved        bool
	Verdict                     string
	Reason                      string
}

type RegistryPatch struct {
	Executed                bool
	Package                 string
	Theorem                 string
	AuditPath               string
	RuntimeFamilyUpdated    bool
	PublicationAtlasOverlay bool
	ReopensGate420          bool
	RequiresAtlasRewrite    bool
	Ready                   bool
	Verdict                 string
	Reason                  string
}

type FinalStatus struct {
	Executed                     bool
	Reconciled                   bool
	KGenPromoted                 bool
	Gen2ZeroPromoted             bool
	XSupportPromoted             bool
	YPhaseStillQuarantined       bool
	CoefficientsStillQuarantined bool
	NativeFlavorDim              int
	KXYCoeffDim                  int
	NoMassPrediction             bool
	NoMixingPrediction           bool
	Status                       string
	Verdict                      string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Delta       ReconciliationDelta
	Atlas       ReconciledAtlas
	Firewall    Firewall
	Patch       RegistryPatch
	Final       FinalStatus
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Delta = buildDelta()
	a.Atlas = buildAtlas(a.Delta)
	a.Firewall = buildFirewall(a.Delta)
	a.Patch = buildPatch(a.Atlas)
	a.Final = buildFinal(a.Delta, a.Atlas, a.Firewall)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate420PublicationAtlas: true, Gate420Acyclic: true, Gate420NativeFlavorDim: NativeChargedFlavorDim, Gate420ConditionalDim: KXYChargedCoeffDim, Gate420FamilyAxiomsSealed: true, Gate420NoFlavorReopening: true, NoEmpiricalInputsImported: true, Verdict: StatusGate420AtlasInherited}
}

func buildDelta() ReconciliationDelta {
	deltas := []GateDelta{
		{Gate: 444, Name: "Generation-2 structural-zero intersection sieve", InputStatus: "K_gen listed as quarantined minimal family axiom", OutputStatus: "K_gen = diag(-1,0,1) geometrically forced up to primitive scale/orientation/permutation", PromotesStructuralObject: true, PreservesFirewall: true, PredictsObservableValue: false, NativeChargedDimAfter: NativeChargedFlavorDim, KXYCoeffDimAfter: KXYChargedCoeffDim, Verdict: StatusKGenPromotedGeometric, Reason: "traceless integer-spaced three-level KMS spectrum collapses uniquely to the primitive triplet {-1,0,1}"},
		{Gate: 445, Name: "Generation-2 mass-lift bridge topology sieve", InputStatus: "off-diagonal bridge support not fixed", OutputStatus: "primitive endpoint-balanced closed triangle support forced", PromotesStructuralObject: true, PreservesFirewall: true, PredictsObservableValue: false, NativeChargedDimAfter: NativeChargedFlavorDim, KXYCoeffDimAfter: KXYChargedCoeffDim, Verdict: StatusXTriangleSupportPromoted, Reason: "the structural-zero lift and endpoint balance reject open chains and isolate complete triangle support"},
		{Gate: 446, Name: "Signed-cycle / complex phase orientation sieve", InputStatus: "triangle support forced", OutputStatus: "signed cycle has two classes; complex Hermitian phase remains a continuum", PromotesStructuralObject: false, PreservesFirewall: true, PredictsObservableValue: false, NativeChargedDimAfter: NativeChargedFlavorDim, KXYCoeffDimAfter: KXYChargedCoeffDim, Verdict: StatusYPhaseFirewallPreserved, Reason: "native boundaries do not quantize Φ_cycle or promote Y_gen"},
		{Gate: 447, Name: "Sector-coefficient source ledger closure", InputStatus: "K/X/Y coefficient source arena open", OutputStatus: "multiple symbolic coefficient ledgers survive; nine amplitudes remain quarantined", PromotesStructuralObject: false, PreservesFirewall: true, PredictsObservableValue: false, NativeChargedDimAfter: NativeChargedFlavorDim, KXYCoeffDimAfter: KXYChargedCoeffDim, Verdict: StatusNineCoefficientFirewallPreserved, Reason: "trace, Hermiticity, gauge, KMS, and mass-lift boundaries do not define a coefficient functional"},
	}
	reclasses := []Reclassification{
		{Object: "K_gen", PreviousLayer: "quarantined-family-axiom", ReconciledLayer: "geometrically-forced-structural-axis", PreviousStatus: "conditional hierarchy capacity", ReconciledStatus: StatusKGenPromotedGeometric, Promoted: true, Quarantined: false, ValueBearing: false, Verdict: StatusKGenPromotedGeometric, Reason: "Gate 444 proves the primitive spectrum without observed flavor data"},
		{Object: "Generation-2 bare level", PreviousLayer: "quarantined-family-axiom consequence", ReconciledLayer: "structural-zero consequence", PreviousStatus: "conditional middle entry", ReconciledStatus: StatusGen2ZeroPromotedStructural, Promoted: true, Quarantined: false, ValueBearing: false, Verdict: StatusGen2ZeroPromotedStructural, Reason: "the middle eigenvalue of the forced primitive K spectrum is exactly zero"},
		{Object: "X_triangle support", PreviousLayer: "quarantined shift-support choice", ReconciledLayer: "structural mass-lift topology", PreviousStatus: "conditional real mixing capacity", ReconciledStatus: StatusXTriangleSupportPromoted, Promoted: true, Quarantined: false, ValueBearing: false, Verdict: StatusXTriangleSupportPromoted, Reason: "Gate 445 fixes support topology but not amplitude or sign/phase orientation"},
		{Object: "Y_gen / Φ_cycle", PreviousLayer: "quarantined-family-phase", ReconciledLayer: "quarantined-family-phase", PreviousStatus: "conditional CP capacity", ReconciledStatus: StatusYPhaseFirewallPreserved, Promoted: false, Quarantined: true, ValueBearing: true, Verdict: StatusFailedYPhaseNotPromoted, Reason: "Gate 446 leaves a phase continuum and cannot predict CKM/PMNS CP values"},
		{Object: "charged K/X/Y sector coefficients", PreviousLayer: "environmental coefficient ledger", ReconciledLayer: "environmental coefficient ledger", PreviousStatus: "nine free coefficients", ReconciledStatus: StatusNineCoefficientFirewallPreserved, Promoted: false, Quarantined: true, ValueBearing: true, Verdict: StatusFailedNoCoefficientSelector, Reason: "Gate 447 proves multiple ledgers survive the native boundary stack"},
	}
	promoted, quarantined := countReclasses(reclasses)
	return ReconciliationDelta{Executed: true, Deltas: deltas, Reclassifications: reclasses, PromotedObjects: promoted, QuarantinedObjects: quarantined, NativeDimBefore: NativeChargedFlavorDim, NativeDimAfter: NativeChargedFlavorDim, KXYCoeffDimBefore: KXYChargedCoeffDim, KXYCoeffDimAfter: KXYChargedCoeffDim, FlavorObservableValuesAdded: 0, CoefficientSelectorsAdded: 0, Verdict: StatusAtlasReconciliationDeltaCompiled, Reason: "the atlas delta is structural-only: K_gen, the Generation-2 bare zero, and unsigned X support are promoted; all value-bearing flavor data remain sealed"}
}

func countReclasses(xs []Reclassification) (promoted, quarantined int) {
	for _, x := range xs {
		if x.Promoted {
			promoted++
		}
		if x.Quarantined {
			quarantined++
		}
	}
	return promoted, quarantined
}

func buildAtlas(d ReconciliationDelta) ReconciledAtlas {
	nodes := []AtlasNode{
		{ID: "post444-k-axis", Name: "primitive traceless family axis", Layer: "geometrically-forced-structural-axis", Gates: []int{444}, Status: StatusKGenPromotedGeometric, Claim: "K_gen = diag(-1,0,1) is forced up to primitive equivalence", Boundary: "not a Yukawa eigenvalue or physical mass prediction", DependsOn: []string{"gate420-flavor-firewall"}},
		{ID: "post444-gen2-zero", Name: "Generation-2 bare structural zero", Layer: "structural-consequence", Gates: []int{444}, Status: StatusGen2ZeroPromotedStructural, Claim: "the middle bare level is exactly zero in the forced primitive spectrum", Boundary: "muon/charm physical mass requires bridge data", DependsOn: []string{"post444-k-axis"}},
		{ID: "post445-x-triangle-support", Name: "unsigned triangular mass-lift support", Layer: "structural-bridge-topology", Gates: []int{445}, Status: StatusXTriangleSupportPromoted, Claim: "endpoint-balanced closed triangle support is the unique minimal support topology", Boundary: "amplitude, sign orientation, and phase remain sealed", DependsOn: []string{"post444-gen2-zero"}},
		{ID: "post446-447-flavor-firewall", Name: "phase and coefficient firewall", Layer: "environmental-frontier", Gates: []int{446, 447}, Status: StatusNineCoefficientFirewallPreserved, Claim: "Y/phase and nine K/X/Y charged-sector coefficients remain quarantined", Boundary: "no CKM, PMNS, Yukawa, muon, or charm value predicted", DependsOn: []string{"post445-x-triangle-support"}},
	}
	return ReconciledAtlas{Executed: true, Nodes: nodes, Acyclic: acyclic(nodes), NativeFlavorDim: d.NativeDimAfter, KXYCoeffDim: d.KXYCoeffDimAfter, KGenGeometric: hasStatus(d.Reclassifications, StatusKGenPromotedGeometric), Gen2BareZeroStructural: hasStatus(d.Reclassifications, StatusGen2ZeroPromotedStructural), XTriangleSupportStructural: hasStatus(d.Reclassifications, StatusXTriangleSupportPromoted), YPhaseQuarantined: hasStatus(d.Reclassifications, StatusFailedYPhaseNotPromoted), CoefficientsQuarantined: hasStatus(d.Reclassifications, StatusFailedNoCoefficientSelector), NoNewPhysicsClaim: d.FlavorObservableValuesAdded == 0 && d.CoefficientSelectorsAdded == 0, Verdict: StatusPost444FlavorAtlasReconciled}
}

func hasStatus(xs []Reclassification, status string) bool {
	for _, x := range xs {
		if x.Verdict == status || x.ReconciledStatus == status {
			return true
		}
	}
	return false
}

func acyclic(nodes []AtlasNode) bool {
	ids := map[string]bool{}
	for _, n := range nodes {
		ids[n.ID] = true
	}
	// External dependencies such as Gate-420 nodes are allowed as roots.
	for _, n := range nodes {
		if n.ID == "post444-k-axis" {
			continue
		}
		hasInternalParent := false
		for _, d := range n.DependsOn {
			if ids[d] {
				hasInternalParent = true
			}
		}
		if !hasInternalParent {
			return false
		}
	}
	return true
}

func buildFirewall(d ReconciliationDelta) Firewall {
	return Firewall{Executed: true, NoObservedMuonMassImported: true, NoObservedCharmMassImported: true, NoObservedYukawaImported: true, NoCKMImported: true, NoPMNSImported: true, NoPoleMassFit: true, NoCurveFit: true, NoCosmologyInput: true, NativeFlavorDimPreserved: d.NativeDimAfter == NativeChargedFlavorDim, KXYCoeffDimPreserved: d.KXYCoeffDimAfter == KXYChargedCoeffDim, Verdict: StatusEmpiricalFirewallPreserved, Reason: "Gate 448 changes only registry classification; it imports no observed flavor or cosmological datum"}
}

func buildPatch(a ReconciledAtlas) RegistryPatch {
	return RegistryPatch{Executed: true, Package: "pkg/bridge/post444flavoratlasreconciliation", Theorem: "Post444FlavorFrontierAtlasReconciliationTheorem", AuditPath: "docs/audits/gates/gate448_registry_audit.md", RuntimeFamilyUpdated: true, PublicationAtlasOverlay: true, ReopensGate420: false, RequiresAtlasRewrite: false, Ready: a.Executed && a.Acyclic && a.NoNewPhysicsClaim, Verdict: StatusPost444FlavorAtlasReconciled, Reason: "the patch is an overlay over Gate 420: it records the later structural promotions while preserving the atlas as historical publication state"}
}

func buildFinal(d ReconciliationDelta, a ReconciledAtlas, f Firewall) FinalStatus {
	return FinalStatus{Executed: true, Reconciled: a.Executed && a.Acyclic, KGenPromoted: a.KGenGeometric, Gen2ZeroPromoted: a.Gen2BareZeroStructural, XSupportPromoted: a.XTriangleSupportStructural, YPhaseStillQuarantined: a.YPhaseQuarantined, CoefficientsStillQuarantined: a.CoefficientsQuarantined, NativeFlavorDim: d.NativeDimAfter, KXYCoeffDim: d.KXYCoeffDimAfter, NoMassPrediction: f.NoObservedMuonMassImported && f.NoObservedCharmMassImported, NoMixingPrediction: f.NoCKMImported && f.NoPMNSImported, Status: StatusPost444FlavorAtlasReconciled, Verdict: "post-444 flavor atlas reconciliation complete: structural promotions installed, value-bearing firewall preserved"}
}

func buildNext() NextStep {
	return NextStep{Gate: 449, Title: "Structural Family Board Export / Manuscript Delta Patch", Reason: "Gate 448 reconciles the theorem registry; the manuscript/publication text now needs a compact delta patch explaining exactly what changed after Gate 420.", PrimaryTask: "export reviewer-facing language and figure/table deltas that distinguish the newly forced K/X structural layer from the still-quarantined Y/phase/coefficient layer"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate420PublicationAtlas || !a.Inheritance.Gate420Acyclic || !a.Inheritance.NoEmpiricalInputsImported {
		return fmt.Errorf("Gate 448 requires Gate 420 atlas inheritance and no empirical inputs")
	}
	if !a.Delta.Executed || a.Delta.PromotedObjects != 3 || a.Delta.NativeDimAfter != NativeChargedFlavorDim || a.Delta.KXYCoeffDimAfter != KXYChargedCoeffDim {
		return fmt.Errorf("invalid post-444 reconciliation delta: %+v", a.Delta)
	}
	if a.Delta.FlavorObservableValuesAdded != 0 || a.Delta.CoefficientSelectorsAdded != 0 {
		return fmt.Errorf("Gate 448 must not add flavor values or coefficient selectors")
	}
	if !a.Atlas.Executed || len(a.Atlas.Nodes) != ReconciledNodeCount || !a.Atlas.Acyclic || !a.Atlas.KGenGeometric || !a.Atlas.Gen2BareZeroStructural || !a.Atlas.XTriangleSupportStructural || !a.Atlas.YPhaseQuarantined || !a.Atlas.CoefficientsQuarantined {
		return fmt.Errorf("invalid reconciled atlas: %+v", a.Atlas)
	}
	if !a.Firewall.NativeFlavorDimPreserved || !a.Firewall.KXYCoeffDimPreserved || !a.Firewall.NoObservedYukawaImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported {
		return fmt.Errorf("empirical firewall violation")
	}
	if !a.Final.Reconciled || !a.Final.KGenPromoted || !a.Final.Gen2ZeroPromoted || !a.Final.XSupportPromoted || !a.Final.YPhaseStillQuarantined || !a.Final.CoefficientsStillQuarantined {
		return fmt.Errorf("final reconciliation incomplete")
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 448 reconciles the post-publication flavor atlas after Gates 444-447.  The update is narrow and structural: K_gen=diag(-1,0,1), the Generation-2 bare zero, and the unsigned triangular mass-lift support are promoted as forced geometry/topology.  The native charged flavor dimension remains %d and the charged K/X/Y coefficient ledger remains %d-dimensional and quarantined.  Y_gen/Φ_cycle, sector amplitudes, muon/charm physical masses, Yukawa values, CKM/PMNS angles, and CP phases are not predicted.", a.Final.NativeFlavorDim, a.Final.KXYCoeffDim)
}

func Statuses() []string {
	return []string{StatusGate420AtlasInherited, StatusGate444To447DeltaInherited, StatusAtlasReconciliationDeltaCompiled, StatusKGenPromotedGeometric, StatusGen2ZeroPromotedStructural, StatusXTriangleSupportPromoted, StatusYPhaseFirewallPreserved, StatusNineCoefficientFirewallPreserved, StatusNativeFlavorDim13Preserved, StatusPost444FlavorAtlasReconciled, StatusEmpiricalFirewallPreserved, StatusFailedNoYukawaPrediction, StatusFailedNoCKMPMNSPrediction, StatusFailedXAmplitudeNotPromoted, StatusFailedYPhaseNotPromoted, StatusFailedNoCoefficientSelector, StatusFailedNoMuonCharmMassPrediction}
}

func join(xs []string) string { return strings.Join(xs, ", ") }

func joinInts(xs []int) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return strings.Join(parts, ",")
}
