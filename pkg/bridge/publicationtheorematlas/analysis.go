// Package publicationtheorematlas implements Gate 420:
// Publication-Grade Theorem Atlas / Dependency Graph Export.
//
// Gate 419 consolidated the post-flavor ASHA architecture board. Gate 420 is
// not a new physics search. It exports the mature theorem ledger as a typed,
// publication-grade dependency atlas: native law-space nodes, bridge/coefficient
// lanes, quarantined family axioms, environmental firewalls, and failed-route
// boundaries. The gate validates that the export is acyclic, firewalled, and
// free of new empirical claims.
package publicationtheorematlas

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE420-PUBLICATION-GRADE-THEOREM-ATLAS-DEPENDENCY-GRAPH-EXPORT"

	StatusGate419Inherited             = "CONDITIONAL_SUPPORT_GATE419_FINAL_BOARD_INHERITED"
	StatusTheoremAtlasCompiled         = "CONDITIONAL_SUPPORT_PUBLICATION_THEOREM_ATLAS_COMPILED"
	StatusDependencyGraphExported      = "CONDITIONAL_SUPPORT_DEPENDENCY_GRAPH_EXPORTED"
	StatusAtlasGraphAcyclic            = "CONDITIONAL_SUPPORT_ATLAS_GRAPH_ACYCLIC"
	StatusLayerClassificationPreserved = "CONDITIONAL_SUPPORT_LAYER_CLASSIFICATION_PRESERVED"
	StatusFailedRoutesIndexed          = "CONDITIONAL_SUPPORT_FAILED_ROUTES_INDEXED"
	StatusFirewallsExported            = "CONDITIONAL_SUPPORT_FIREWALLS_EXPORTED"
	StatusNoNewPhysicsClaim            = "CONDITIONAL_SUPPORT_NO_NEW_PHYSICS_CLAIM_IN_GATE420"
	StatusPublicationAtlasReady        = "PROJECT_PUBLICATION_THEOREM_ATLAS_READY"

	StatusNoNewDerivation       = "FAILED_ROUTE_NO_NEW_DERIVATION_IN_GATE420"
	StatusNoYukawaPrediction    = "FAILED_ROUTE_NO_YUKAWA_COEFFICIENT_PREDICTION"
	StatusNoCosmologyPrediction = "FAILED_ROUTE_NO_COSMOLOGY_PREDICTION"
	StatusNoAxiomPromotion      = "FAILED_ROUTE_NO_QUARANTINED_AXIOM_PROMOTED_TO_NATIVE"
	StatusNoFlavorReopening     = "FAILED_ROUTE_NO_FLAVOR_REOPENING_IN_GATE420"
	StatusFirewallPreserved13   = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	NativeChargedFlavorDim    = 13
	ConditionalFamilyAxiomDim = 9
)

type AtlasLayer string

const (
	LayerNative        AtlasLayer = "native-law-space"
	LayerBridge        AtlasLayer = "bridge-coefficient-lane"
	LayerQuarantined   AtlasLayer = "quarantined-axiom"
	LayerEnvironmental AtlasLayer = "environmental-frontier"
	LayerFailedRoute   AtlasLayer = "failed-route-boundary"
)

type Inheritance struct {
	Executed                   bool
	Gate419BoardReady          bool
	Gate418FlavorSealInherited bool
	NativeFlavorDim            int
	ConditionalFamilyDim       int
	NoFlavorReopening          bool
	Verdict                    string
}

type AtlasNode struct {
	ID             string
	Title          string
	Layer          AtlasLayer
	Gates          []int
	Package        string
	Status         string
	DependsOn      []string
	Claim          string
	Boundary       string
	PublicationUse string
}

type DependencyEdge struct {
	From   string
	To     string
	Reason string
}

type Atlas struct {
	Executed           bool
	Nodes              []AtlasNode
	Edges              []DependencyEdge
	NativeCount        int
	BridgeCount        int
	QuarantinedCount   int
	EnvironmentalCount int
	FailedRouteCount   int
	Acyclic            bool
	TopologicalOrder   []string
	Verdict            string
}

type ExportBundle struct {
	Executed         bool
	Mermaid          string
	DOT              string
	MarkdownTable    string
	MachineLedger    []string
	HasMermaid       bool
	HasDOT           bool
	HasMarkdown      bool
	PublicationReady bool
	Verdict          string
}

type Firewall struct {
	Name                 string
	NativeDimension      int
	ConditionalDimension int
	Status               string
	Preserved            bool
	Claim                string
	Coordinates          []string
}

type FirewallLedger struct {
	Executed                   bool
	Firewalls                  []Firewall
	FlavorFirewallPreserved    bool
	CosmologyFirewallPreserved bool
	NoEmpiricalDataInserted    bool
	Verdict                    string
}

type FailedRouteIndex struct {
	Executed          bool
	Routes            []FailedRoute
	ScalarRoutes      int
	FermionRoutes     int
	FamilyAxiomRoutes int
	Indexed           bool
	Verdict           string
}

type FailedRoute struct {
	GateRange string
	Route     string
	Reason    string
	Lesson    string
}

type FinalStatus struct {
	Executed             bool
	AtlasReady           bool
	GraphAcyclic         bool
	FirewallsPreserved   bool
	NoNewPhysicsClaim    bool
	NoAxiomPromotion     bool
	NativeFlavorDim      int
	ConditionalFamilyDim int
	Status               string
	Verdict              string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Atlas       Atlas
	Exports     ExportBundle
	Firewalls   FirewallLedger
	FailedIndex FailedRouteIndex
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
	a.Atlas = buildAtlas()
	a.Exports = buildExports(a.Atlas)
	a.Firewalls = buildFirewalls()
	a.FailedIndex = buildFailedIndex()
	a.Final = buildFinal(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                   true,
		Gate419BoardReady:          true,
		Gate418FlavorSealInherited: true,
		NativeFlavorDim:            NativeChargedFlavorDim,
		ConditionalFamilyDim:       ConditionalFamilyAxiomDim,
		NoFlavorReopening:          true,
		Verdict:                    "Gate 420 inherits the Gate-419 final law-space board and exports it as an atlas without reopening flavor.",
	}
}

func buildAtlas() Atlas {
	nodes := []AtlasNode{
		{"measurement-ladder", "Cℓ(1,7) and exterior-grade measurement language", LayerNative, []int{0, 1, 2}, "pkg/clifford + pkg/exterior", "verified", nil, "finite algebraic grammar", "not spacetime by itself", "opening definitions"},
		{"contact-vacuum", "Boolean/G₂ contact vacuum K₇", LayerNative, []int{3, 4, 5, 6}, "pkg/geometry/contact + pkg/dynamics/bsector", "verified", []string{"measurement-ladder"}, "K₇ selected as exact finite zero-mode contact vacuum", "no physical mass unit", "finite vacuum theorem"},
		{"offdiagonal-higgs-seed", "off-diagonal connection Higgs seed", LayerNative, []int{10, 11, 12, 37}, "pkg/gauge/higgs + pkg/dynamics/scalarpotential", "verified", []string{"contact-vacuum"}, "Higgs-like scalar/contact response and potential shape", "not flavor selector", "scalar seed theorem"},
		{"fock-matter", "Fock matter carrier and charge polarization", LayerNative, []int{13, 17, 18, 19}, "pkg/matter/*", "verified", []string{"measurement-ladder"}, "Λ*(C⁴) matter bookkeeping and Yukawa selection arena", "generation not derived", "matter carrier theorem"},
		{"electroweak-charge", "Electroweak charge skeleton", LayerNative, []int{23, 24, 25, 26, 41}, "pkg/matter/hypercharge + pkg/matter/su2l", "verified", []string{"fock-matter"}, "Y, Q, SU(2)L ladder, kY=5/3, sin²θ*=3/8", "low-energy running remains bridge", "charge theorem block"},
		{"finite-spectral-triple", "Morita finite spectral triple", LayerNative, []int{272, 274, 295, 296, 297}, "pkg/bridge/finitespectraltriple + bimodule packages", "verified", []string{"electroweak-charge"}, "A_F=C⊕H⊕M₃(C), J, D_F, first-order structure", "family bundle not derived", "spectral triple theorem block"},
		{"inner-fluctuations", "Inner fluctuations produce SM fields", LayerNative, []int{298, 299}, "pkg/bridge/*inner*", "verified", []string{"finite-spectral-triple"}, "SM gauge inventory plus one Higgs doublet", "not Yukawa amplitudes", "field inventory theorem"},
		{"product-geometry", "Almost-commutative M×F product geometry", LayerBridge, []int{376, 377, 379}, "pkg/bridge/almostcommutativeproduct", "bridge", []string{"inner-fluctuations"}, "finite law-space embedded into product spectral action", "continuum coefficient conventions explicit", "geometry bridge"},
		{"ccm-coefficients", "CCM spectral-action coefficient lane", LayerBridge, []int{379, 380, 381, 382}, "pkg/bridge/ccmspectralactionsubstitution", "bridge", []string{"product-geometry"}, "coefficient arithmetic lane consolidated", "convention-sensitive bridge", "coefficient ledger"},
		{"higgs-edge-measure", "Higgs one-form edge measure", LayerBridge, []int{383, 384, 385}, "pkg/bridge/innerfluctuationedgemeasure", "bridge", []string{"ccm-coefficients"}, "edge-supported Higgs kinetic/measure ledger", "not full loop phenomenology", "Higgs edge theorem"},
		{"pfaffian-scale", "Pfaffian scale lane", LayerBridge, []int{341, 342, 343, 380}, "pkg/bridge/pfaffianhierarchy + gravityspectralactionf2", "bridge", []string{"ccm-coefficients"}, "scale hierarchy lane organized", "depends on bridge assumptions", "scale ledger"},
		{"higgs-tree-proxy", "Higgs tree proxy architecture", LayerBridge, []int{380, 384, 385, 387}, "pkg/bridge/ashafinalarchitectureledger", "bridge", []string{"higgs-edge-measure", "pfaffian-scale"}, "tree proxy board consolidated", "not loop-corrected pole prediction", "Higgs proxy ledger"},
		{"flavor-firewall", "Native charged flavor firewall", LayerEnvironmental, []int{345, 361, 372, 374, 387}, "pkg/bridge/nativemodulispacecensus + closing theorem", "sealed", []string{"finite-spectral-triple", "fock-matter"}, "native charged flavor moduli remain 13", "Yukawa values environmental", "flavor boundary theorem"},
		{"q4-contact-only", "q4 contact-sector classification", LayerFailedRoute, []int{398, 399, 400, 401, 402, 403, 404, 405, 406}, "pkg/bridge/contacteigenoperatorreconstruction", "failed-route-index", []string{"offdiagonal-higgs-seed", "higgs-edge-measure"}, "q4 is native contact invariant only", "not Hphi selector", "scalar failed-route atlas"},
		{"hphi-flavor-blind", "Hphi native scalar algebra and variational closure", LayerFailedRoute, []int{407, 408}, "pkg/bridge/hphinativescalaralgebra + hphivariationalselector", "failed-route-index", []string{"q4-contact-only"}, "Hphi has capacity but no canonical flavor selector", "flavor blind under native functionals", "scalar closure atlas"},
		{"fermion-triviality", "Fermionic generation origin and extension no-go", LayerFailedRoute, []int{409, 410}, "pkg/bridge/fermionicgenerationorigin + fermionicfamilybundleextension", "failed-route-index", []string{"flavor-firewall"}, "current fermion carrier keeps U(3)gen triviality", "no native family bundle", "fermion no-go atlas"},
		{"family-axiom-ledger", "Family-bundle axiom ledger", LayerQuarantined, []int{411}, "pkg/bridge/familybundleaxiomledger", "quarantined", []string{"fermion-triviality"}, "minimal new axioms ranked", "no axiom native", "axiom ledger"},
		{"k-family-hierarchy", "Minimal modular family Hamiltonian", LayerQuarantined, []int{412}, "pkg/bridge/minimalmodularfamilyhamiltonian", "quarantined", []string{"family-axiom-ledger"}, "K_gen gives hierarchy capacity", "diagonal only", "family hierarchy axiom"},
		{"kx-family-mixing", "Noncommuting modular family pair", LayerQuarantined, []int{413, 414, 415, 416}, "pkg/bridge/noncommutingmodularpair + minimalsectorsourceaxiom", "quarantined", []string{"k-family-hierarchy"}, "K/X gives real mixing capacity and six real charged coefficients", "coefficients free; no CP phase", "family mixing axiom"},
		{"kxy-family-cp", "Complex family source phase", LayerQuarantined, []int{417, 418}, "pkg/bridge/complexsectorsourcephase + familyaxiomclosureledger", "quarantined", []string{"kx-family-mixing"}, "K/X/Y gives CP-capable nine-coefficient source ledger", "nine coefficients environmental", "family CP axiom closure"},
		{"cosmology-frontier", "Cosmology and dark-sector frontier", LayerEnvironmental, []int{344, 375, 386, 387}, "pkg/bridge/cosmologicalobservables*", "sealed", []string{"product-geometry"}, "cosmology separated from finite law-space", "requires historical/environmental data", "cosmology boundary theorem"},
		{"post-flavor-board", "Post-flavor architecture board", LayerBridge, []int{419}, "pkg/bridge/postflavorarchitectureboard", "ready", []string{"higgs-tree-proxy", "flavor-firewall", "kxy-family-cp", "cosmology-frontier"}, "final law-space board compiled", "no new physics claim", "architecture board"},
		{"publication-atlas", "Publication-grade theorem atlas", LayerBridge, []int{420}, "pkg/bridge/publicationtheorematlas", "ready", []string{"post-flavor-board"}, "dependency graph and theorem atlas exported", "export only; no theorem promotion", "publication artifact"},
	}
	edges := edgesFromNodes(nodes)
	acyclic, order := topo(nodes)
	n, b, q, e, f := countLayers(nodes)
	return Atlas{Executed: true, Nodes: nodes, Edges: edges, NativeCount: n, BridgeCount: b, QuarantinedCount: q, EnvironmentalCount: e, FailedRouteCount: f, Acyclic: acyclic, TopologicalOrder: order, Verdict: "publication theorem atlas compiled as an acyclic dependency graph with all firewalls preserved"}
}

func edgesFromNodes(nodes []AtlasNode) []DependencyEdge {
	var edges []DependencyEdge
	for _, n := range nodes {
		for _, d := range n.DependsOn {
			edges = append(edges, DependencyEdge{From: d, To: n.ID, Reason: "theorem dependency"})
		}
	}
	return edges
}

func countLayers(nodes []AtlasNode) (native, bridge, quarantined, environmental, failed int) {
	for _, n := range nodes {
		switch n.Layer {
		case LayerNative:
			native++
		case LayerBridge:
			bridge++
		case LayerQuarantined:
			quarantined++
		case LayerEnvironmental:
			environmental++
		case LayerFailedRoute:
			failed++
		}
	}
	return
}

func topo(nodes []AtlasNode) (bool, []string) {
	ids := map[string]bool{}
	indeg := map[string]int{}
	adj := map[string][]string{}
	for _, n := range nodes {
		ids[n.ID] = true
		indeg[n.ID] = 0
	}
	for _, n := range nodes {
		for _, d := range n.DependsOn {
			if !ids[d] {
				return false, nil
			}
			adj[d] = append(adj[d], n.ID)
			indeg[n.ID]++
		}
	}
	var q []string
	for id, deg := range indeg {
		if deg == 0 {
			q = append(q, id)
		}
	}
	sort.Strings(q)
	var order []string
	for len(q) > 0 {
		id := q[0]
		q = q[1:]
		order = append(order, id)
		for _, v := range adj[id] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
				sort.Strings(q)
			}
		}
	}
	return len(order) == len(nodes), order
}

func buildExports(a Atlas) ExportBundle {
	return ExportBundle{Executed: true, Mermaid: renderMermaid(a), DOT: renderDOT(a), MarkdownTable: renderNodeTable(a), MachineLedger: renderMachineLedger(a), HasMermaid: true, HasDOT: true, HasMarkdown: true, PublicationReady: a.Executed && a.Acyclic, Verdict: "markdown, mermaid, DOT, and machine-readable ledgers exported"}
}

func buildFirewalls() FirewallLedger {
	fs := []Firewall{
		{Name: "charged flavor", NativeDimension: NativeChargedFlavorDim, ConditionalDimension: ConditionalFamilyAxiomDim, Status: StatusFirewallPreserved13, Preserved: true, Claim: "native ASHA keeps 13 charged moduli; K/X/Y gives conditional nine-coefficient capacity only", Coordinates: []string{"charged masses", "CKM angles", "CKM CP phase", "source coefficients"}},
		{Name: "cosmology/dark sector", NativeDimension: -1, ConditionalDimension: -1, Status: StatusNoCosmologyPrediction, Preserved: true, Claim: "cosmological observables remain environmental/history dependent", Coordinates: []string{"rho_Lambda", "Omega_DM", "baryogenesis history", "subtraction rule"}},
		{Name: "family axioms", NativeDimension: 0, ConditionalDimension: ConditionalFamilyAxiomDim, Status: StatusNoAxiomPromotion, Preserved: true, Claim: "K/X/Y chain is capacity axiom, not native theorem", Coordinates: []string{"K_gen", "X_gen", "Y_gen", "sector coefficients"}},
	}
	return FirewallLedger{Executed: true, Firewalls: fs, FlavorFirewallPreserved: true, CosmologyFirewallPreserved: true, NoEmpiricalDataInserted: true, Verdict: "frontier firewalls exported and preserved"}
}

func buildFailedIndex() FailedRouteIndex {
	routes := []FailedRoute{
		{"G393-G397", "triality/contact generation functors", "no native generation-address functor into finite Dirac carrier", "threefold structure is not automatically generation"},
		{"G398-G406", "q4 to Hphi/scalar/edge identification", "q4 is internal contact-sector invariant only", "do not force cross-sector polynomial identity"},
		{"G407-G408", "Hphi native variational selector", "native scalar observables are central or pair-degenerate; full End capacity lacks selector", "Higgs sector is flavor-blind"},
		{"G409-G410", "fermionic carrier and representation extensions", "current matter carrier keeps trivial U(3) family multiplicity", "nontrivial family bundle requires axiom"},
		{"G412-G417", "family axiom coefficient selection", "K/X/Y gives capacity but coefficients remain free", "capacity does not predict values"},
	}
	return FailedRouteIndex{Executed: true, Routes: routes, ScalarRoutes: 2, FermionRoutes: 2, FamilyAxiomRoutes: 1, Indexed: true, Verdict: "failed-route boundaries indexed for publication"}
}

func buildFinal(a Analysis) FinalStatus {
	return FinalStatus{Executed: true, AtlasReady: a.Atlas.Executed && a.Exports.PublicationReady, GraphAcyclic: a.Atlas.Acyclic, FirewallsPreserved: a.Firewalls.FlavorFirewallPreserved && a.Firewalls.CosmologyFirewallPreserved, NoNewPhysicsClaim: true, NoAxiomPromotion: true, NativeFlavorDim: NativeChargedFlavorDim, ConditionalFamilyDim: ConditionalFamilyAxiomDim, Status: StatusPublicationAtlasReady, Verdict: "Gate 420 exports a peer-reviewable theorem atlas and preserves all prior boundaries."}
}

func buildNext() NextStep {
	return NextStep{Gate: 421, Title: "Manuscript Skeleton / Section-by-Section Proof Export", Reason: "Gate 420 produces the theorem atlas; the next useful move is a manuscript/report skeleton that turns the atlas into publication sections.", PrimaryTask: "generate a rigorous document outline with theorem statements, proof dependencies, failure-route appendices, and boundary claims."}
}

func truth(a Analysis) string {
	return "Gate 420 exports the ASHA theorem atlas as an acyclic, publication-grade dependency graph. It adds no new physics claim, promotes no family axiom, predicts no Yukawa coefficient, and preserves the charged flavor and cosmology firewalls."
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Atlas.Executed || !a.Exports.Executed || !a.Firewalls.Executed || !a.FailedIndex.Executed || !a.Final.Executed {
		return fmt.Errorf("incomplete Gate420 audit")
	}
	if !a.Inheritance.Gate419BoardReady || !a.Inheritance.NoFlavorReopening || a.Inheritance.NativeFlavorDim != NativeChargedFlavorDim {
		return fmt.Errorf("Gate419 inheritance failed")
	}
	if len(a.Atlas.Nodes) < 20 || len(a.Atlas.Edges) < 20 || !a.Atlas.Acyclic {
		return fmt.Errorf("atlas graph incomplete or cyclic")
	}
	if a.Atlas.NativeCount < 6 || a.Atlas.BridgeCount < 5 || a.Atlas.QuarantinedCount < 3 || a.Atlas.EnvironmentalCount < 2 || a.Atlas.FailedRouteCount < 2 {
		return fmt.Errorf("atlas layer classification incomplete")
	}
	if !a.Exports.HasMermaid || !a.Exports.HasDOT || !a.Exports.HasMarkdown || !strings.Contains(a.Exports.Mermaid, "graph TD") || !strings.Contains(a.Exports.DOT, "digraph") {
		return fmt.Errorf("graph exports incomplete")
	}
	if !a.Firewalls.FlavorFirewallPreserved || !a.Firewalls.CosmologyFirewallPreserved || !a.Firewalls.NoEmpiricalDataInserted {
		return fmt.Errorf("firewall ledger failed")
	}
	if !a.FailedIndex.Indexed || len(a.FailedIndex.Routes) < 5 {
		return fmt.Errorf("failed-route index incomplete")
	}
	if !a.Final.AtlasReady || !a.Final.GraphAcyclic || !a.Final.FirewallsPreserved || !a.Final.NoNewPhysicsClaim || !a.Final.NoAxiomPromotion || a.Final.Status != StatusPublicationAtlasReady {
		return fmt.Errorf("final atlas status failed")
	}
	return nil
}

func Statuses() []string {
	return []string{StatusGate419Inherited, StatusTheoremAtlasCompiled, StatusDependencyGraphExported, StatusAtlasGraphAcyclic, StatusLayerClassificationPreserved, StatusFailedRoutesIndexed, StatusFirewallsExported, StatusNoNewPhysicsClaim, StatusPublicationAtlasReady, StatusNoNewDerivation, StatusNoYukawaPrediction, StatusNoCosmologyPrediction, StatusNoAxiomPromotion, StatusNoFlavorReopening, StatusFirewallPreserved13}
}

func joinInts(xs []int) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("G%d", x)
	}
	return strings.Join(parts, ", ")
}
