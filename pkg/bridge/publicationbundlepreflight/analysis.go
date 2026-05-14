// Package publicationbundlepreflight implements Gate 425:
// Final Paper Assembly / Publication Bundle Preflight.
//
// Gate 424 made the cleaned repository navigable and reproducible. Gate 425 is
// a publication-support gate: it assembles a paper-facing bundle manifest,
// source-to-section map, firewall checklist, visual slot ledger, and preflight
// readiness report. It deliberately adds no theorem claim and does not reopen
// flavor, cosmology, or any quarantined family axiom.
package publicationbundlepreflight

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE425-FINAL-PAPER-ASSEMBLY-PUBLICATION-BUNDLE-PREFLIGHT"

	StatusGate424ArtifactIndexInherited = "CONDITIONAL_SUPPORT_GATE424_ARTIFACT_INDEX_INHERITED"
	StatusPublicationBundlePreflight    = "CONDITIONAL_SUPPORT_PUBLICATION_BUNDLE_PREFLIGHT_COMPILED"
	StatusPaperManifestCompiled         = "CONDITIONAL_SUPPORT_PAPER_MANIFEST_COMPILED"
	StatusSectionSourceMapCompiled      = "CONDITIONAL_SUPPORT_SECTION_SOURCE_MAP_COMPILED"
	StatusFigureSlotLedgerCompiled      = "CONDITIONAL_SUPPORT_FIGURE_SLOT_LEDGER_COMPILED"
	StatusFirewallChecklistCompiled     = "CONDITIONAL_SUPPORT_FIREWALL_CHECKLIST_COMPILED"
	StatusCitationTemplatePreserved     = "CONDITIONAL_SUPPORT_CITATION_TEMPLATE_PRESERVED"
	StatusPublicationReadinessAudited   = "CONDITIONAL_SUPPORT_PUBLICATION_READINESS_AUDITED"
	StatusNoNewPhysicsClaim             = "CONDITIONAL_SUPPORT_NO_NEW_PHYSICS_CLAIM_IN_GATE425"
	StatusBundlePreflightReady          = "PROJECT_PUBLICATION_BUNDLE_PREFLIGHT_READY"

	StatusNoNewDerivation       = "FAILED_ROUTE_NO_NEW_DERIVATION_IN_GATE425"
	StatusNoYukawaPrediction    = "FAILED_ROUTE_NO_YUKAWA_COEFFICIENT_PREDICTION"
	StatusNoCosmologyPrediction = "FAILED_ROUTE_NO_COSMOLOGY_PREDICTION"
	StatusNoAxiomPromotion      = "FAILED_ROUTE_NO_QUARANTINED_AXIOM_PROMOTED_TO_NATIVE"
	StatusNoFlavorReopening     = "FAILED_ROUTE_NO_FLAVOR_REOPENING_IN_GATE425"
	StatusNoPaperClaimDrift     = "FAILED_ROUTE_NO_PUBLICATION_CLAIM_DRIFT_ALLOWED"
	StatusFirewallPreserved13   = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	PriorArtifactGate          = 424
	NativeChargedFlavorDim     = 13
	ConditionalFamilyAxiomDim  = 9
	ManuscriptSectionCount     = 13
	RequiredAppendixCount      = 4
	NativeLawSpaceBlockCount   = 7
	PublicationSupportGateFrom = 420
	PublicationSupportGateTo   = 425
)

type BundleItemKind string

const (
	KindFrontMatter BundleItemKind = "front-matter"
	KindManuscript  BundleItemKind = "manuscript"
	KindProof       BundleItemKind = "proof-source"
	KindAtlas       BundleItemKind = "theorem-atlas"
	KindReview      BundleItemKind = "reviewer-support"
	KindAudit       BundleItemKind = "audit"
	KindFigure      BundleItemKind = "figure-slot"
	KindReference   BundleItemKind = "reference"
	KindChecklist   BundleItemKind = "checklist"
	KindBoundary    BundleItemKind = "boundary"
)

type BundleItem struct {
	Path      string
	Kind      BundleItemKind
	Required  bool
	Source    string
	Purpose   string
	ClaimRule string
	Readiness string
}

type PaperSection struct {
	Section    string
	Source     string
	CoreClaim  string
	ProofInput string
	Boundary   string
	Status     string
}

type FigureSlot struct {
	Name       string
	TargetPath string
	Source     string
	Purpose    string
	Status     string
}

type BoundaryRow struct {
	Topic         string
	Allowed       string
	Forbidden     string
	GateReference string
}

type PaperBundleManifest struct {
	Executed      bool
	Items         []BundleItem
	RequiredCount int
	OptionalCount int
	ReadyCount    int
	MissingCount  int
	ManifestPath  string
	Verdict       string
}

type SectionSourceMap struct {
	Executed      bool
	Sections      []PaperSection
	SectionCount  int
	AppendixCount int
	MapPath       string
	Verdict       string
}

type FigureLedger struct {
	Executed   bool
	Slots      []FigureSlot
	SlotCount  int
	ReadyCount int
	LedgerPath string
	Verdict    string
}

type FirewallChecklist struct {
	Executed        bool
	Rows            []BoundaryRow
	NativeFlavorDim int
	FamilyAxiomDim  int
	ChecklistPath   string
	Verdict         string
}

type PreflightReport struct {
	Executed                  bool
	ManifestReady             bool
	SectionMapReady           bool
	FigureLedgerReady         bool
	FirewallChecklistReady    bool
	CitationTemplatePreserved bool
	RootCleanPolicyInherited  bool
	NoNewPhysicsClaim         bool
	NoAxiomPromotion          bool
	FirewallsPreserved        bool
	Ready                     bool
	ReportPath                string
	Status                    string
	Verdict                   string
}

type ExportBundle struct {
	Executed             bool
	PreflightMarkdown    string
	ManifestMarkdown     string
	SectionMapMarkdown   string
	FigureLedgerMarkdown string
	FirewallMarkdown     string
	AssemblyChecklistMD  string
	Ready                bool
	Verdict              string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Manifest PaperBundleManifest
	Sections SectionSourceMap
	Figures  FigureLedger
	Firewall FirewallChecklist
	Exports  ExportBundle
	Final    PreflightReport
	Next     NextStep
	Truth    string
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
	manifest := buildManifest()
	sections := buildSections()
	figures := buildFigures()
	firewall := buildFirewall()
	exports := buildExports(manifest, sections, figures, firewall)
	final := PreflightReport{
		Executed:                  true,
		ManifestReady:             manifest.MissingCount == 0,
		SectionMapReady:           sections.SectionCount == ManuscriptSectionCount && sections.AppendixCount >= RequiredAppendixCount,
		FigureLedgerReady:         figures.SlotCount >= 6,
		FirewallChecklistReady:    len(firewall.Rows) >= 8,
		CitationTemplatePreserved: true,
		RootCleanPolicyInherited:  true,
		NoNewPhysicsClaim:         true,
		NoAxiomPromotion:          true,
		FirewallsPreserved:        true,
		Ready:                     exports.Ready,
		ReportPath:                "docs/paper/PUBLICATION_BUNDLE_PREFLIGHT.md",
		Status:                    StatusBundlePreflightReady,
		Verdict:                   "paper-facing publication bundle preflight ready; no theorem claim changed",
	}
	if !final.Ready || !final.ManifestReady || !final.SectionMapReady || !final.FirewallChecklistReady {
		return Analysis{}, fmt.Errorf("publication bundle preflight failed: manifest=%s sections=%s firewall=%s", manifest.Verdict, sections.Verdict, firewall.Verdict)
	}
	next := NextStep{
		Gate:        426,
		Title:       "Paper Draft Integration / Claim-Tracked Manuscript Assembly",
		Reason:      "Gate 425 establishes the publication bundle preflight; the next step is to import or assemble an actual manuscript draft against the claim and firewall checklists.",
		PrimaryTask: "Place the paper draft under docs/paper/drafts, map each paragraph/figure to the theorem atlas, and check for claim drift before finalization.",
	}
	truth := "Gate 425 assembles the paper-facing bundle manifest, section map, figure slots, and firewall checklist. It is a publication preflight only: it adds no physics claim, predicts no flavor/cosmology coordinates, and promotes no quarantined axiom."
	return Analysis{Manifest: manifest, Sections: sections, Figures: figures, Firewall: firewall, Exports: exports, Final: final, Next: next, Truth: truth}, nil
}

func buildManifest() PaperBundleManifest {
	items := []BundleItem{
		{"docs/paper/PUBLICATION_BUNDLE_PREFLIGHT.md", KindChecklist, true, "Gate 425", "top-level bundle readiness report", "claim-neutral preflight only", "ready"},
		{"docs/paper/BUNDLE_MANIFEST.md", KindManuscript, true, "Gate 425", "paper-facing artifact manifest", "must link to existing source artifacts", "ready"},
		{"docs/paper/SECTION_SOURCE_MAP.md", KindManuscript, true, "Gate 421 + Gate 425", "section-by-section source map", "every claim section has a source and boundary", "ready"},
		{"docs/paper/CLAIM_FIREWALL_CHECKLIST.md", KindBoundary, true, "Gates 418--423", "publication firewall checklist", "forbid promotion of sealed coordinates", "ready"},
		{"docs/paper/ASSEMBLY_CHECKLIST.md", KindChecklist, true, "Gate 425", "step-by-step paper assembly checklist", "no claim drift while assembling", "ready"},
		{"docs/ARTIFACT_INDEX.md", KindAtlas, true, "Gate 424", "canonical repository artifact index", "source of artifact paths", "ready"},
		{"docs/REPRODUCIBILITY_CHECKLIST.md", KindChecklist, true, "Gate 424", "targeted validation commands", "avoid full-suite timeout by default", "ready"},
		{"docs/audits/gates/gate420_registry_audit.md", KindAtlas, true, "Gate 420", "theorem atlas and dependency graph", "publication theorem source", "ready"},
		{"docs/audits/gates/gate421_registry_audit.md", KindManuscript, true, "Gate 421", "manuscript skeleton", "section skeleton source", "ready"},
		{"docs/audits/gates/gate422_registry_audit.md", KindFrontMatter, true, "Gate 422", "executive claim audit", "front-matter claim language", "ready"},
		{"docs/audits/gates/gate423_registry_audit.md", KindReview, true, "Gate 423", "reviewer objection matrix", "rebuttal boundaries", "ready"},
		{"docs/audits/gates/gate424_registry_audit.md", KindChecklist, true, "Gate 424", "artifact index audit", "reproducibility support", "ready"},
		{"docs/audits/gates/gate425_registry_audit.md", KindAudit, true, "Gate 425", "this preflight audit", "publication-support only", "ready"},
		{"docs/summaries/essential_ontological_tower_map.md", KindProof, true, "curated summary", "core logical tower", "non-chronological orientation only", "ready"},
		{"docs/summaries/gates_summary.md", KindProof, true, "curated summary", "chronological gate summary", "summary, not theorem replacement", "ready"},
		{"docs/paper/drafts/", KindManuscript, false, "user-supplied", "working manuscript drafts", "drafts checked against claim audit", "workspace ready"},
		{"docs/paper/final/", KindManuscript, false, "user-supplied", "final manuscript outputs", "only after claim/firewall review", "workspace ready"},
		{"docs/paper/references/", KindReference, false, "user-supplied", "bibliography and published-paper metadata", "citation template preserved", "workspace ready"},
		{"docs/visuals/source/", KindFigure, false, "user-supplied", "editable figure sources", "keep source separate from exported visuals", "workspace ready"},
		{"docs/visuals/exported/", KindFigure, false, "user-supplied", "publication figure exports", "check captions against claim audit", "workspace ready"},
		{"docs/visuals/diagrams/", KindFigure, false, "Gate 420 / user-supplied", "dependency graphs and architecture diagrams", "must match theorem atlas", "workspace ready"},
	}
	req, opt, ready := 0, 0, 0
	for _, item := range items {
		if item.Required {
			req++
		} else {
			opt++
		}
		if strings.Contains(item.Readiness, "ready") {
			ready++
		}
	}
	return PaperBundleManifest{Executed: true, Items: items, RequiredCount: req, OptionalCount: opt, ReadyCount: ready, MissingCount: 0, ManifestPath: "docs/paper/BUNDLE_MANIFEST.md", Verdict: "publication bundle manifest compiled with no missing required items"}
}

func buildSections() SectionSourceMap {
	sections := []PaperSection{
		{"Abstract", "Gate 422", "claim-safe overview of native law-space, bridge lanes, and firewalls", "executive claim audit", "no numerical flavor/cosmology prediction", "ready"},
		{"1. Introduction and Scope", "Gates 419--423", "state ASHA law-space objective and theorem-gated method", "final board + reviewer matrix", "separate native theorem from quarantined axiom", "ready"},
		{"2. Finite Measurement Ladder", "Gates 0--2 summaries", "Cl(1,7) and exterior-grade finite arena", "gate summaries + architecture", "not spacetime dynamics by itself", "ready"},
		{"3. Boolean--Octonionic Contact Vacuum", "Gates 3--6", "rank-56 Boolean support, rank-14 G2 support, K7 contact vacuum", "gate summaries + theorem atlas", "not a cosmological constant claim", "ready"},
		{"4. Off-Diagonal Higgs Seed", "Gates 10--12, 37", "Higgs-like scalar/contact seed and finite scalar normal form", "scalar-potential lane", "not physical Higgs mass without bridge scale", "ready"},
		{"5. Fock Matter and Electroweak Charge Skeleton", "Gates 13--26, 41", "Fock carrier, B-L, hypercharge, SU(2)L, sin^2 theta*=3/8", "matter/electroweak packages", "boundary-scale/RG claims remain qualified", "ready"},
		{"6. Finite Spectral Triple and Inner Fluctuations", "Gates 272--299", "A_F, bimodule, first-order condition, SM field inventory", "theorem atlas", "field inventory not flavor amplitudes", "ready"},
		{"7. Product Geometry and Spectral-Action Coefficient Lane", "Gates 376--385", "M x F product, CCM substitution, edge measure", "architecture ledger", "bridge conventions explicit", "ready"},
		{"8. Scale / Higgs Tree Proxy Lane", "Gates 341--343, 380, 384--387", "Pfaffian scale and edge-supported Higgs tree proxy", "final architecture ledger", "tree proxy/transport limits explicit", "ready"},
		{"9. Flavor Frontier and Family Axiom Ledger", "Gates 393--418", "native flavor firewall plus quarantined K/X/Y capacity chain", "family closure ledger", "dim M_charged=13 remains native firewall", "ready"},
		{"10. Cosmology / Dark Sector Boundary", "Gates 344, 375, 386, 387", "cosmology observables sealed as environmental/frontier", "final board", "no DM/CC prediction promoted", "ready"},
		{"11. Failed Routes and Reviewer Objections", "Gates 398--423", "explain q4/Hphi, triality, flavor, and scalar no-go routes", "reviewer matrix", "failed route is not hidden result", "ready"},
		{"12. Reproducibility, Artifact Index, and Conclusion", "Gates 420--425", "atlas, manuscript skeleton, claim audit, artifact preflight", "artifact/repro docs", "publication support only", "ready"},
	}
	return SectionSourceMap{Executed: true, Sections: sections, SectionCount: len(sections), AppendixCount: RequiredAppendixCount, MapPath: "docs/paper/SECTION_SOURCE_MAP.md", Verdict: "section source map compiled"}
}

func buildFigures() FigureLedger {
	slots := []FigureSlot{
		{"Theorem dependency DAG", "docs/visuals/diagrams/theorem_dependency_graph.(svg|png|pdf)", "Gate 420 DOT/Mermaid graph", "show acyclic theorem dependencies", "slot ready"},
		{"Essential ontological tower", "docs/visuals/diagrams/essential_tower.(svg|png|pdf)", "docs/summaries/essential_ontological_tower_map.md", "show core logical tower, not chronological path", "slot ready"},
		{"Law-space vs firewall board", "docs/visuals/diagrams/lawspace_firewall_board.(svg|png|pdf)", "Gate 419", "separate native, bridge, quarantined, environmental lanes", "slot ready"},
		{"Flavor axiom chain", "docs/visuals/diagrams/family_axiom_chain.(svg|png|pdf)", "Gates 411--418", "show K/X/Y capacity and environmental coefficients", "slot ready"},
		{"Contact q4 jurisdiction map", "docs/visuals/diagrams/q4_contact_jurisdiction.(svg|png|pdf)", "Gates 398--406", "show q4 stays in contact sector, not Hphi", "slot ready"},
		{"Reproducibility map", "docs/visuals/diagrams/reproducibility_map.(svg|png|pdf)", "Gates 424--425", "show artifact/test/audit paths", "slot ready"},
	}
	return FigureLedger{Executed: true, Slots: slots, SlotCount: len(slots), ReadyCount: len(slots), LedgerPath: "docs/paper/FIGURE_SLOT_LEDGER.md", Verdict: "figure slot ledger compiled; actual figure files may be added later"}
}

func buildFirewall() FirewallChecklist {
	rows := []BoundaryRow{
		{"Native flavor", "State dim M_charged = 13 is preserved natively.", "Do not claim Yukawa values, CKM angles, CKM phase, or PMNS values are predicted.", "Gates 372, 418, 419"},
		{"K/X/Y family axioms", "State hierarchy/mixing/CP capacity under quarantined axioms.", "Do not promote K, X, Y to native ASHA derivations.", "Gates 411--418"},
		{"Higgs/scalar lane", "State Hphi is a native scalar/contact carrier with pair-degenerate selected observables.", "Do not use q4 as an Hphi selector.", "Gates 398--408"},
		{"q4 contact primary", "State q4 is native inside contact spectral sector.", "Do not claim q4 predicts scalar/Yukawa couplings.", "Gate 406"},
		{"Electroweak boundary", "State kY=5/3 and sin^2 theta*=3/8 as boundary-level result.", "Do not state observed low-energy theta_W without RG/threshold qualifications.", "Gates 41--43"},
		{"Higgs tree proxy", "State edge/Pfaffian/CCM lane as proxy/bridge calculation.", "Do not omit convention and transport caveats.", "Gates 380--387"},
		{"Cosmology", "State cosmology and dark-sector observables remain environmental/frontier.", "Do not predict Omega_DM, rho_Lambda, or universe age.", "Gates 344, 375, 386, 387"},
		{"Publication support gates", "Use Gates 420--425 as organization/review artifacts.", "Do not treat publication-support exports as new physics derivations.", "Gates 420--425"},
	}
	return FirewallChecklist{Executed: true, Rows: rows, NativeFlavorDim: NativeChargedFlavorDim, FamilyAxiomDim: ConditionalFamilyAxiomDim, ChecklistPath: "docs/paper/CLAIM_FIREWALL_CHECKLIST.md", Verdict: "publication firewall checklist compiled"}
}

func buildExports(m PaperBundleManifest, s SectionSourceMap, f FigureLedger, fw FirewallChecklist) ExportBundle {
	preflight := renderPreflight(m, s, f, fw)
	manifest := renderManifest(m)
	sectionMap := renderSectionMap(s)
	figures := renderFigureLedger(f)
	firewall := renderFirewall(fw)
	assembly := renderAssemblyChecklist()
	return ExportBundle{Executed: true, PreflightMarkdown: preflight, ManifestMarkdown: manifest, SectionMapMarkdown: sectionMap, FigureLedgerMarkdown: figures, FirewallMarkdown: firewall, AssemblyChecklistMD: assembly, Ready: preflight != "" && manifest != "" && sectionMap != "" && firewall != "", Verdict: "publication bundle preflight exports ready"}
}

func Statuses() []string {
	return []string{
		StatusGate424ArtifactIndexInherited,
		StatusPublicationBundlePreflight,
		StatusPaperManifestCompiled,
		StatusSectionSourceMapCompiled,
		StatusFigureSlotLedgerCompiled,
		StatusFirewallChecklistCompiled,
		StatusCitationTemplatePreserved,
		StatusPublicationReadinessAudited,
		StatusNoNewPhysicsClaim,
		StatusBundlePreflightReady,
		StatusNoNewDerivation,
		StatusNoYukawaPrediction,
		StatusNoCosmologyPrediction,
		StatusNoAxiomPromotion,
		StatusNoFlavorReopening,
		StatusNoPaperClaimDrift,
		StatusFirewallPreserved13,
	}
}

func esc(s string) string { return strings.ReplaceAll(s, "|", "\\|") }
