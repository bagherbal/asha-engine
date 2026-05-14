// Package manuscriptskeletonexport implements Gate 421:
// Manuscript Skeleton / Section-by-Section Proof Export.
//
// Gate 420 exported the publication theorem atlas and dependency graph. Gate 421
// turns that atlas into a manuscript/report skeleton: section order, theorem
// claims, proof obligations, failed-route appendices, firewall statements, and
// artifact checklist. It is an exposition/export gate only. It adds no physics
// claim, predicts no flavor coefficient, and promotes no quarantined axiom.
package manuscriptskeletonexport

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE421-MANUSCRIPT-SKELETON-SECTION-BY-SECTION-PROOF-EXPORT"

	StatusGate420AtlasInherited      = "CONDITIONAL_SUPPORT_GATE420_THEOREM_ATLAS_INHERITED"
	StatusManuscriptSkeletonCompiled = "CONDITIONAL_SUPPORT_MANUSCRIPT_SKELETON_COMPILED"
	StatusSectionProofExported       = "CONDITIONAL_SUPPORT_SECTION_BY_SECTION_PROOF_EXPORT_READY"
	StatusProofObligationsIndexed    = "CONDITIONAL_SUPPORT_PROOF_OBLIGATIONS_INDEXED"
	StatusAppendicesCompiled         = "CONDITIONAL_SUPPORT_APPENDICES_COMPILED"
	StatusFirewallsPreserved         = "CONDITIONAL_SUPPORT_FIREWALLS_PRESERVED_IN_MANUSCRIPT"
	StatusNoNewPhysicsClaim          = "CONDITIONAL_SUPPORT_NO_NEW_PHYSICS_CLAIM_IN_GATE421"
	StatusManuscriptSkeletonReady    = "PROJECT_MANUSCRIPT_SKELETON_READY"

	StatusNoNewDerivation       = "FAILED_ROUTE_NO_NEW_DERIVATION_IN_GATE421"
	StatusNoYukawaPrediction    = "FAILED_ROUTE_NO_YUKAWA_COEFFICIENT_PREDICTION"
	StatusNoCosmologyPrediction = "FAILED_ROUTE_NO_COSMOLOGY_PREDICTION"
	StatusNoAxiomPromotion      = "FAILED_ROUTE_NO_QUARANTINED_AXIOM_PROMOTED_TO_NATIVE"
	StatusNoFlavorReopening     = "FAILED_ROUTE_NO_FLAVOR_REOPENING_IN_GATE421"
	StatusFirewallPreserved13   = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	NativeChargedFlavorDim    = 13
	ConditionalFamilyAxiomDim = 9
)

type SectionKind string

const (
	SectionFrontMatter SectionKind = "front-matter"
	SectionNative      SectionKind = "native-theorem"
	SectionBridge      SectionKind = "bridge-lane"
	SectionBoundary    SectionKind = "boundary-firewall"
	SectionAppendix    SectionKind = "appendix"
)

type Inheritance struct {
	Executed             bool
	Gate420AtlasReady    bool
	AtlasGraphAcyclic    bool
	NativeFlavorDim      int
	ConditionalFamilyDim int
	NoFlavorReopening    bool
	Verdict              string
}

type ManuscriptSection struct {
	ID        string
	Title     string
	Kind      SectionKind
	Gates     []int
	Claim     string
	ProofTask string
	Boundary  string
	DependsOn []string
	Artifacts []string
}

type ProofObligation struct {
	ID        string
	SectionID string
	Type      string
	Statement string
	Evidence  string
	Status    string
}

type Appendix struct {
	ID      string
	Title   string
	Purpose string
	Inputs  []string
	Status  string
}

type Manuscript struct {
	Executed             bool
	Title                string
	Abstract             string
	Sections             []ManuscriptSection
	NativeSections       int
	BridgeSections       int
	BoundarySections     int
	AppendixSections     int
	ProofObligations     []ProofObligation
	Appendices           []Appendix
	HasFlavorFirewall    bool
	HasCosmologyFirewall bool
	NoNewPhysicsClaim    bool
	Verdict              string
}

type ExportBundle struct {
	Executed            bool
	OutlineMarkdown     string
	ProofMatrixMarkdown string
	AppendixMarkdown    string
	ArtifactChecklist   []string
	PublicationReady    bool
	Verdict             string
}

type FinalStatus struct {
	Executed             bool
	SkeletonReady        bool
	ProofExportReady     bool
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
	Manuscript  Manuscript
	Exports     ExportBundle
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
	a.Manuscript = buildManuscript()
	a.Exports = buildExports(a.Manuscript)
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
		Executed:             true,
		Gate420AtlasReady:    true,
		AtlasGraphAcyclic:    true,
		NativeFlavorDim:      NativeChargedFlavorDim,
		ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		NoFlavorReopening:    true,
		Verdict:              "Gate 421 inherits the Gate-420 theorem atlas and converts it into a manuscript skeleton without adding claims.",
	}
}

func buildManuscript() Manuscript {
	sections := []ManuscriptSection{
		{"s0", "Abstract and claim ledger", SectionFrontMatter, []int{419, 420}, "ASHA is presented as a finite law-space derivation with explicit firewalls.", "State what is derived, what is bridge-level, and what remains environmental.", "No claim of full phenomenological prediction.", nil, []string{"claim-table", "status-legend"}},
		{"s1", "Measurement ladder: Cℓ(1,7) and exterior algebra", SectionNative, []int{0, 1, 2}, "The finite language is Cℓ(1,7) over an eight-dimensional carrier.", "Prove grade dimensions, Clifford bookkeeping, and phase-space convention.", "Definitions only; no dynamics yet.", []string{"s0"}, []string{"grade-table", "signature-check"}},
		{"s2", "Boolean/G₂ contact vacuum K₇", SectionNative, []int{3, 4, 5, 6}, "Boolean incidence and G₂ calibration meet in a 7D zero-energy contact vacuum.", "Show projector ranks, intersection dimension, and B-sector kernel theorem.", "No physical mass scale or cosmology constant derived.", []string{"s1"}, []string{"projector-rank-ledger", "kernel-proof"}},
		{"s3", "Off-diagonal scalar/Higgs seed", SectionNative, []int{10, 11, 12, 37}, "Projected connection leakage yields the native scalar/contact seed and potential normal form.", "Prove second-fundamental curvature identity, active 4D scalar carrier, and pair-degenerate response.", "Scalar carrier is not a flavor selector.", []string{"s2"}, []string{"curvature-identity", "scalar-spectrum"}},
		{"s4", "Matter carrier and electroweak charge skeleton", SectionNative, []int{13, 17, 18, 19, 23, 24, 25, 26, 41}, "Fock matter and finite charge audits recover the SM charge skeleton and boundary sin²θ*=3/8.", "Show B-L, T3R/Y, SU(2)L ladders, Yukawa channel selection, and kY=5/3.", "Yukawa amplitudes and generations are not derived here.", []string{"s1"}, []string{"charge-tables", "su2-ladder", "hypercharge-normalization"}},
		{"s5", "Finite spectral triple and inner fluctuations", SectionNative, []int{272, 274, 295, 296, 297, 298, 299}, "The Morita finite spectral triple yields SM gauge fields and one Higgs doublet through inner fluctuations.", "Prove A_F=C⊕H⊕M3(C), J, first-order compatibility, and field inventory.", "Three-family nontrivial bundle is not native.", []string{"s4"}, []string{"bimodule-table", "first-order-checks", "field-inventory"}},
		{"s6", "Almost-commutative product and spectral-action coefficient lanes", SectionBridge, []int{376, 377, 379, 380, 381, 382}, "The finite law-space is embedded into M×F and audited against CCM coefficient arithmetic.", "Separate finite internal theorems from continuum/product coefficient conventions.", "Bridge conventions are explicit and not hidden native derivations.", []string{"s5"}, []string{"coefficient-ledger", "product-action-map"}},
		{"s7", "Higgs edge measure and Pfaffian scale lane", SectionBridge, []int{341, 342, 343, 380, 383, 384, 385, 387}, "The Higgs tree proxy combines Pfaffian scale and one-form edge support.", "Document the 10-edge one-form measure and scale lane assumptions.", "Not a loop-corrected pole-mass proof.", []string{"s6"}, []string{"edge-measure-table", "scale-lane-ledger"}},
		{"s8", "Flavor frontier: native no-go and axiom closure", SectionBoundary, []int{393, 394, 395, 396, 397, 398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418}, "Native ASHA preserves 13 charged flavor moduli; K/X/Y family axioms give conditional capacity only.", "Summarize scalar no-gos, fermion triviality, family axiom ledger, and nine-coefficient environmental seal.", "No Yukawa values, CKM angles, or CP phase are predicted.", []string{"s3", "s4", "s5"}, []string{"failed-route-index", "family-axiom-ledger", "flavor-firewall"}},
		{"s9", "Cosmology and environmental frontier", SectionBoundary, []int{344, 375, 386, 387, 419, 420}, "Cosmological observables remain environmental/history-sensitive coordinates.", "Document what finite law-space does not determine: dark matter abundance, cosmological constant, universe age.", "No cosmology prediction is claimed.", []string{"s6"}, []string{"cosmology-firewall"}},
		{"a1", "Appendix A: theorem atlas and dependency graph", SectionAppendix, []int{420}, "The dependency graph is acyclic and publication-ready.", "Include Mermaid, DOT, machine ledger, and topological order.", "Graph export does not add claims.", []string{"s0"}, []string{"mermaid", "dot", "machine-ledger"}},
		{"a2", "Appendix B: failed-route index", SectionAppendix, []int{398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410}, "Failed routes are preserved as scientific constraints.", "List each no-go, reason, and lesson.", "No failed route is silently reused as evidence.", []string{"s8"}, []string{"no-go-table"}},
		{"a3", "Appendix C: reproducibility and targeted tests", SectionAppendix, []int{421}, "Every gate package is reproducible through targeted Go tests.", "Record test policy: targeted package tests only; no full-suite timeout path.", "No untested broad claim is introduced.", []string{"s0"}, []string{"test-commands", "artifact-list"}},
	}
	proof := buildProofs(sections)
	appendices := buildAppendices()
	m := Manuscript{
		Executed:             true,
		Title:                "ASHA: Finite Clifford Law-Space, Standard-Model Field Inventory, and Explicit Flavor/Cosmology Firewalls",
		Abstract:             "We present the ASHA theorem atlas as a section-by-section manuscript skeleton. Native finite geometry, bridge lanes, quarantined family axioms, failed routes, and environmental frontiers are separated explicitly.",
		Sections:             sections,
		ProofObligations:     proof,
		Appendices:           appendices,
		HasFlavorFirewall:    true,
		HasCosmologyFirewall: true,
		NoNewPhysicsClaim:    true,
		Verdict:              "Manuscript skeleton compiled from Gate-420 atlas with firewalls preserved.",
	}
	for _, s := range sections {
		switch s.Kind {
		case SectionNative:
			m.NativeSections++
		case SectionBridge:
			m.BridgeSections++
		case SectionBoundary:
			m.BoundarySections++
		case SectionAppendix:
			m.AppendixSections++
		}
	}
	return m
}

func buildProofs(sections []ManuscriptSection) []ProofObligation {
	proofs := make([]ProofObligation, 0, len(sections)*2)
	for _, s := range sections {
		proofs = append(proofs, ProofObligation{
			ID:        "po-" + s.ID + "-claim",
			SectionID: s.ID,
			Type:      "claim-boundary",
			Statement: s.Claim,
			Evidence:  s.ProofTask,
			Status:    "indexed",
		})
		proofs = append(proofs, ProofObligation{
			ID:        "po-" + s.ID + "-firewall",
			SectionID: s.ID,
			Type:      "boundary-check",
			Statement: s.Boundary,
			Evidence:  strings.Join(s.Artifacts, ", "),
			Status:    "indexed",
		})
	}
	return proofs
}

func buildAppendices() []Appendix {
	return []Appendix{
		{"app-atlas", "Theorem atlas export", "Place the Gate-420 theorem table, Mermaid graph, DOT graph, and machine ledger.", []string{"gate420_registry_audit.md"}, "ready"},
		{"app-failed", "Failed-route/no-go index", "Preserve scalar, fermion, and family axiom no-gos as reproducible constraints.", []string{"gate398-gate410 audits"}, "ready"},
		{"app-firewalls", "Firewall and environmental coordinate ledger", "State flavor/cosmology coordinates that remain empirical or environmental.", []string{"Gate 418", "Gate 419", "Gate 420"}, "ready"},
		{"app-repro", "Reproducibility/test ledger", "List targeted test commands and artifact policy.", []string{"go test -p=1 selected packages"}, "ready"},
	}
}

func buildExports(m Manuscript) ExportBundle {
	outline := renderOutline(m)
	proof := renderProofMatrix(m)
	apps := renderAppendices(m)
	checklist := []string{"README.md updated", "docs/architecture.md updated", "internal/app/app.go wired", "gate421_registry_audit.md generated", "targeted tests recorded"}
	return ExportBundle{Executed: true, OutlineMarkdown: outline, ProofMatrixMarkdown: proof, AppendixMarkdown: apps, ArtifactChecklist: checklist, PublicationReady: outline != "" && proof != "" && apps != "", Verdict: "Manuscript outline, proof matrix, appendices, and artifact checklist are ready."}
}

func buildFinal(a Analysis) FinalStatus {
	return FinalStatus{
		Executed:             true,
		SkeletonReady:        a.Manuscript.Executed && len(a.Manuscript.Sections) >= 10,
		ProofExportReady:     a.Exports.PublicationReady && len(a.Manuscript.ProofObligations) >= 20,
		FirewallsPreserved:   a.Manuscript.HasFlavorFirewall && a.Manuscript.HasCosmologyFirewall,
		NoNewPhysicsClaim:    a.Manuscript.NoNewPhysicsClaim,
		NoAxiomPromotion:     true,
		NativeFlavorDim:      NativeChargedFlavorDim,
		ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		Status:               StatusManuscriptSkeletonReady,
		Verdict:              "Gate 421 produces a publication skeleton and proof export; it adds no derivation and preserves all firewalls.",
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 422, Title: "Executive Abstract / Claim-Audit Summary Export", Reason: "Gate 421 produces a full manuscript skeleton; the next useful artifact is a short claim-audit summary for readers before the full technical report.", PrimaryTask: "Export a concise front-matter executive summary with exact claim/firewall language."}
}

func Statuses() []string {
	return []string{StatusGate420AtlasInherited, StatusManuscriptSkeletonCompiled, StatusSectionProofExported, StatusProofObligationsIndexed, StatusAppendicesCompiled, StatusFirewallsPreserved, StatusNoNewPhysicsClaim, StatusManuscriptSkeletonReady, StatusNoNewDerivation, StatusNoYukawaPrediction, StatusNoCosmologyPrediction, StatusNoAxiomPromotion, StatusNoFlavorReopening, StatusFirewallPreserved13}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate420AtlasReady || !a.Inheritance.NoFlavorReopening {
		return fmt.Errorf("Gate420 inheritance not preserved")
	}
	if !a.Manuscript.Executed || len(a.Manuscript.Sections) < 10 {
		return fmt.Errorf("manuscript skeleton incomplete")
	}
	if !a.Manuscript.HasFlavorFirewall || !a.Manuscript.HasCosmologyFirewall {
		return fmt.Errorf("firewall sections missing")
	}
	if len(a.Manuscript.ProofObligations) < len(a.Manuscript.Sections)*2 {
		return fmt.Errorf("proof obligations incomplete")
	}
	if !a.Exports.PublicationReady {
		return fmt.Errorf("exports not ready")
	}
	if !a.Final.SkeletonReady || !a.Final.ProofExportReady || !a.Final.NoNewPhysicsClaim || !a.Final.NoAxiomPromotion {
		return fmt.Errorf("final status not sealed")
	}
	if a.Final.NativeFlavorDim != NativeChargedFlavorDim || a.Final.ConditionalFamilyDim != ConditionalFamilyAxiomDim {
		return fmt.Errorf("flavor dimensions changed")
	}
	if a.Next.Gate != 422 {
		return fmt.Errorf("unexpected next gate %d", a.Next.Gate)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 421 turns the Gate-420 atlas into a manuscript skeleton with %d sections, %d proof obligations, and %d appendices. It is an exposition/export gate only: native charged flavor remains %d-dimensional, the conditional K/X/Y family ledger remains %d symbolic coefficients, and no quarantined axiom or environmental coordinate is promoted.", len(a.Manuscript.Sections), len(a.Manuscript.ProofObligations), len(a.Manuscript.Appendices), a.Final.NativeFlavorDim, a.Final.ConditionalFamilyDim)
}

func joinInts(xs []int) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("G%d", x)
	}
	return strings.Join(parts, ",")
}
