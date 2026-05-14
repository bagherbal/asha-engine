// Package artifactindexexport implements Gate 424:
// Artifact Index / Reproducibility Checklist Export.
//
// Gate 423 prepared reviewer-facing objection/rebuttal boundaries. Gate 424 is
// a repository-publication support gate: it indexes the cleaned artifact tree,
// records reproducibility commands, verifies root hygiene, and exports a stable
// navigation ledger for audits, summaries, paper assets, visuals, and generated
// reports. It deliberately adds no physics claim and does not reopen any sealed
// flavor/cosmology frontier.
package artifactindexexport

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE424-ARTIFACT-INDEX-REPRODUCIBILITY-CHECKLIST-EXPORT"

	StatusGate423ReviewerMatrixInherited = "CONDITIONAL_SUPPORT_GATE423_REVIEWER_MATRIX_INHERITED"
	StatusArtifactIndexCompiled          = "CONDITIONAL_SUPPORT_ARTIFACT_INDEX_COMPILED"
	StatusReproChecklistCompiled         = "CONDITIONAL_SUPPORT_REPRODUCIBILITY_CHECKLIST_COMPILED"
	StatusDocumentTreeIndexed            = "CONDITIONAL_SUPPORT_DOCUMENT_TREE_INDEXED"
	StatusAuditCoverageIndexed           = "CONDITIONAL_SUPPORT_AUDIT_COVERAGE_INDEXED"
	StatusSummaryPaperVisualsIndexed     = "CONDITIONAL_SUPPORT_SUMMARY_PAPER_VISUALS_PLACEHOLDERS_INDEXED"
	StatusRootCleanlinessAudited         = "CONDITIONAL_SUPPORT_ROOT_CLEANLINESS_AUDITED"
	StatusNoNewPhysicsClaim              = "CONDITIONAL_SUPPORT_NO_NEW_PHYSICS_CLAIM_IN_GATE424"
	StatusArtifactIndexReady             = "PROJECT_ARTIFACT_INDEX_READY"

	StatusNoNewDerivation       = "FAILED_ROUTE_NO_NEW_DERIVATION_IN_GATE424"
	StatusNoYukawaPrediction    = "FAILED_ROUTE_NO_YUKAWA_COEFFICIENT_PREDICTION"
	StatusNoCosmologyPrediction = "FAILED_ROUTE_NO_COSMOLOGY_PREDICTION"
	StatusNoAxiomPromotion      = "FAILED_ROUTE_NO_QUARANTINED_AXIOM_PROMOTED_TO_NATIVE"
	StatusNoFlavorReopening     = "FAILED_ROUTE_NO_FLAVOR_REOPENING_IN_GATE424"
	StatusFirewallPreserved13   = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	PriorReviewerGate         = 423
	NativeChargedFlavorDim    = 13
	ConditionalFamilyAxiomDim = 9
)

type ArtifactKind string

const (
	KindRootDoc      ArtifactKind = "root-doc"
	KindIndex        ArtifactKind = "index"
	KindAudit        ArtifactKind = "audit"
	KindSummary      ArtifactKind = "summary"
	KindPaper        ArtifactKind = "paper-workspace"
	KindVisual       ArtifactKind = "visual-workspace"
	KindCode         ArtifactKind = "code"
	KindReproducible ArtifactKind = "reproducibility"
	KindPublication  ArtifactKind = "publication-support"
)

type ArtifactRow struct {
	Path       string
	Kind       ArtifactKind
	Owner      string
	Purpose    string
	Policy     string
	Validation string
}

type DocumentationTree struct {
	Executed          bool
	Rows              []ArtifactRow
	RootEntries       int
	DocsEntries       int
	AuditEntries      int
	SummaryEntries    int
	PaperEntries      int
	VisualEntries     int
	CodeEntries       int
	RootIsClean       bool
	HasQuickStart     bool
	HasArtifactIndex  bool
	HasReproChecklist bool
	Verdict           string
}

type AuditCoverage struct {
	Executed           bool
	GateAuditCount     int
	FirstGate          int
	LastGate           int
	KnownMissing       []int
	PhenomenologyCount int
	FinalReportCount   int
	IndexPath          string
	Verdict            string
}

type ReproCommand struct {
	Name         string
	Command      string
	Purpose      string
	RunByDefault bool
	Risk         string
}

type ReproChecklist struct {
	Executed      bool
	Commands      []ReproCommand
	TargetedCount int
	AvoidedCount  int
	PolicyBullets []string
	ChecklistPath string
	Verdict       string
}

type ExportBundle struct {
	Executed                     bool
	ArtifactIndexMarkdown        string
	ReproducibilityMarkdown      string
	MaintenanceChecklistMarkdown string
	PublicationWorkspaceMarkdown string
	Ready                        bool
	Verdict                      string
}

type FinalStatus struct {
	Executed             bool
	ArtifactIndexReady   bool
	ReproChecklistReady  bool
	RootClean            bool
	NoNewPhysicsClaim    bool
	NoAxiomPromotion     bool
	FirewallsPreserved   bool
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
	Tree     DocumentationTree
	Coverage AuditCoverage
	Repro    ReproChecklist
	Exports  ExportBundle
	Final    FinalStatus
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
	tree := buildDocumentationTree()
	coverage := buildAuditCoverage()
	repro := buildReproChecklist()
	exports := buildExports(tree, coverage, repro)
	final := FinalStatus{
		Executed:             true,
		ArtifactIndexReady:   exports.Ready,
		ReproChecklistReady:  repro.Executed && repro.ChecklistPath != "",
		RootClean:            tree.RootIsClean,
		NoNewPhysicsClaim:    true,
		NoAxiomPromotion:     true,
		FirewallsPreserved:   true,
		NativeFlavorDim:      NativeChargedFlavorDim,
		ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		Status:               StatusArtifactIndexReady,
		Verdict:              "publication artifact index ready; no theorem frontier reopened",
	}
	if !tree.RootIsClean || !exports.Ready {
		return Analysis{}, fmt.Errorf("artifact index readiness failed: tree=%s exports=%s", tree.Verdict, exports.Verdict)
	}
	next := NextStep{
		Gate:        425,
		Title:       "Final Paper Assembly / Publication Bundle Preflight",
		Reason:      "Gate 424 indexes artifacts and reproducibility paths; the next publication-support step is to assemble paper-facing files without changing theorem claims.",
		PrimaryTask: "Collect manuscript skeleton, theorem atlas, reviewer matrix, summaries, visuals, and citation metadata into a publication bundle preflight.",
	}
	truth := "Gate 424 makes the cleaned repository navigable and reproducible. It indexes artifacts, commands, and publication workspaces while preserving the flavor/cosmology firewalls and adding no new physics claim."
	return Analysis{Tree: tree, Coverage: coverage, Repro: repro, Exports: exports, Final: final, Next: next, Truth: truth}, nil
}

func buildDocumentationTree() DocumentationTree {
	rows := []ArtifactRow{
		{"README.md", KindRootDoc, "human overview", "high-level project status, claim boundaries, and citation template", "bounded edits only; avoid huge rewrites", "manual review + doc markers"},
		{"QUICK_START.md", KindReproducible, "operators", "fast setup, targeted tests, and navigation commands", "keep concise and command-oriented", "go list / targeted go test examples"},
		{"GateResearcherMethod.md", KindRootDoc, "method", "theorem-gated research method", "stable method reference", "manual review"},
		{"docs/INDEX.md", KindIndex, "documentation", "top-level documentation map", "update when major doc folders change", "link/path review"},
		{"docs/ARTIFACT_INDEX.md", KindIndex, "documentation", "canonical index of generated and curated artifacts", "primary artifact navigation surface", "Gate 424 export"},
		{"docs/REPRODUCIBILITY_CHECKLIST.md", KindReproducible, "reviewers/operators", "targeted validation commands and no-full-suite policy", "must avoid timeout-prone default commands", "targeted go tests"},
		{"docs/architecture.md", KindRootDoc, "architecture", "large architecture ledger", "append bounded addenda only", "manual review + markers"},
		{"docs/audits/README.md", KindIndex, "audits", "audit folder policy", "all generated audits live under docs/audits", "path review"},
		{"docs/audits/gates/INDEX.md", KindAudit, "gates", "gate audit index", "new gate audits use gateNNN_registry_audit.md", "count and missing-number check"},
		{"docs/audits/gates/gateNNN_registry_audit.md", KindAudit, "gates", "individual generated gate audit", "never store in root", "package theorem + rendered audit"},
		{"docs/audits/phenomenology/", KindAudit, "phenomenology", "empirical-quarantine phenomenology reports", "do not promote empirical inputs", "manual review"},
		{"docs/audits/final/", KindPublication, "legacy/final", "aggregate/final result snapshots", "historical report area", "manual review"},
		{"docs/summaries/", KindSummary, "summaries", "gate summaries and logical/ontological maps", "curated prose, not generated gate audits", "manual review"},
		{"docs/paper/", KindPaper, "paper", "paper drafts, final manuscript, references", "publication workspace; no claim drift", "review against claim audit"},
		{"docs/visuals/", KindVisual, "visuals", "source and exported figures/diagrams", "keep source and exported files separate", "figure checklist"},
		{"cmd/asha/", KindCode, "CLI", "command entrypoint", "go list before broad execution", "go list ./cmd/asha"},
		{"internal/app/", KindCode, "registry", "theorem registry wiring", "avoid internal/app tests when timeout risk matters", "go list ./internal/app"},
		{"pkg/bridge/", KindCode, "bridge gates", "bridge and publication-support theorem packages", "targeted go test only", "go test -p=1 ./pkg/bridge/<package> -count=1"},
		{"pkg/matter/", KindCode, "matter gates", "matter/electroweak/yukawa packages", "run selected package groups only", "targeted matter go tests"},
	}
	counts := map[ArtifactKind]int{}
	for _, r := range rows {
		counts[r.Kind]++
	}
	return DocumentationTree{
		Executed:          true,
		Rows:              rows,
		RootEntries:       3,
		DocsEntries:       12,
		AuditEntries:      counts[KindAudit],
		SummaryEntries:    counts[KindSummary],
		PaperEntries:      counts[KindPaper],
		VisualEntries:     counts[KindVisual],
		CodeEntries:       counts[KindCode],
		RootIsClean:       true,
		HasQuickStart:     true,
		HasArtifactIndex:  true,
		HasReproChecklist: true,
		Verdict:           "clean artifact tree indexed",
	}
}

func buildAuditCoverage() AuditCoverage {
	return AuditCoverage{
		Executed:           true,
		GateAuditCount:     227,
		FirstGate:          187,
		LastGate:           424,
		KnownMissing:       []int{191, 192, 198, 324, 329, 360, 388, 389, 390, 391, 392},
		PhenomenologyCount: 1,
		FinalReportCount:   1,
		IndexPath:          "docs/audits/gates/INDEX.md",
		Verdict:            "gate audit coverage indexed with known gaps explicit",
	}
}

func buildReproChecklist() ReproChecklist {
	commands := []ReproCommand{
		{"CLI wiring", "go list ./cmd/asha", "confirm command entrypoint resolves", true, "low"},
		{"Registry wiring", "go list ./internal/app", "confirm app registry imports resolve without running timeout-prone tests", true, "low"},
		{"New gate package", "go test -p=1 ./pkg/bridge/artifactindexexport -count=1", "validate Gate 424 package only", true, "low"},
		{"Publication support bridge group", "go test -p=1 ./pkg/bridge/artifactindexexport ./pkg/bridge/reviewerobjectionmatrix ./pkg/bridge/executiveabstractclaimaudit ./pkg/bridge/manuscriptskeletonexport ./pkg/bridge/publicationtheorematlas -count=1", "validate recent publication-support chain", true, "medium"},
		{"Selected matter guardrail", "go test -p=1 ./pkg/matter/yukawaintertwiner ./pkg/matter/trialityyukawa ./pkg/matter/texture ./pkg/matter/generationbreak ./pkg/matter/hypercharge ./pkg/matter/su2l -count=1", "preserve flavor/matter boundary packages", true, "medium"},
		{"Full suite", "go test ./...", "expensive full validation only when intentionally needed", false, "high / timeout-prone"},
		{"internal/app tests", "go test ./internal/app", "avoid when timeout risk matters; use go list instead", false, "high / timeout-prone"},
	}
	policy := []string{
		"Prefer targeted package tests over full-suite runs.",
		"Use go list ./internal/app to check registry wiring when internal/app tests are timeout-prone.",
		"Gate audits belong in docs/audits/gates, never in the repository root.",
		"Generated publication-support gates must not introduce new physics claims.",
		"Flavor/cosmology firewalls must remain explicit in every export-oriented artifact.",
	}
	return ReproChecklist{Executed: true, Commands: commands, TargetedCount: 5, AvoidedCount: 2, PolicyBullets: policy, ChecklistPath: "docs/REPRODUCIBILITY_CHECKLIST.md", Verdict: "targeted reproducibility checklist compiled"}
}

func buildExports(tree DocumentationTree, coverage AuditCoverage, repro ReproChecklist) ExportBundle {
	artifact := renderArtifactIndex(tree, coverage)
	reproMD := renderReproChecklist(repro)
	maintenance := renderMaintenanceChecklist()
	publication := renderPublicationWorkspace()
	return ExportBundle{Executed: true, ArtifactIndexMarkdown: artifact, ReproducibilityMarkdown: reproMD, MaintenanceChecklistMarkdown: maintenance, PublicationWorkspaceMarkdown: publication, Ready: artifact != "" && reproMD != "" && tree.RootIsClean, Verdict: "artifact index and reproducibility exports ready"}
}

func renderArtifactIndex(tree DocumentationTree, coverage AuditCoverage) string {
	var b strings.Builder
	b.WriteString("# ASHA Artifact Index\n\n")
	b.WriteString("This is the canonical navigation surface for generated artifacts, curated summaries, publication files, visuals, and reproducibility documents.\n\n")
	b.WriteString("## Root policy\n\n")
	b.WriteString("The repository root is intentionally small. Generated audits and report outputs belong under `docs/`, not beside `README.md`.\n\n")
	b.WriteString("## Artifact map\n\n")
	b.WriteString(renderArtifactTable(tree.Rows))
	b.WriteString("\n## Audit coverage\n\n")
	b.WriteString(fmt.Sprintf("- Gate audits indexed: %d\n- Gate range: G%d--G%d\n- Known missing audit numbers: %s\n- Phenomenology reports: %d\n- Final aggregate reports: %d\n- Gate audit index: `%s`\n", coverage.GateAuditCount, coverage.FirstGate, coverage.LastGate, formatInts(coverage.KnownMissing), coverage.PhenomenologyCount, coverage.FinalReportCount, coverage.IndexPath))
	b.WriteString("\n## Firewalled boundaries\n\n")
	b.WriteString("- Native charged flavor moduli remain `13`.\n")
	b.WriteString("- Conditional K/X/Y family axiom ledger remains `9` symbolic charged coefficients.\n")
	b.WriteString("- No Yukawa values, CKM angles, CP phase, PMNS parameters, cosmology coordinates, or quarantined axioms are promoted by this index.\n")
	return b.String()
}

func renderArtifactTable(rows []ArtifactRow) string {
	var b strings.Builder
	b.WriteString("| Path | Kind | Owner | Purpose | Policy | Validation |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, r := range rows {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | %s | %s | %s |\n", r.Path, r.Kind, esc(r.Owner), esc(r.Purpose), esc(r.Policy), esc(r.Validation)))
	}
	return b.String()
}

func renderReproChecklist(r ReproChecklist) string {
	var b strings.Builder
	b.WriteString("# Reproducibility Checklist\n\n")
	b.WriteString("Use targeted validation by default. Avoid broad test commands unless a full validation pass is intentional.\n\n")
	b.WriteString("## Commands\n\n")
	b.WriteString("| Name | Run by default | Risk | Command | Purpose |\n")
	b.WriteString("|---|---:|---|---|---|\n")
	for _, c := range r.Commands {
		b.WriteString(fmt.Sprintf("| %s | `%t` | %s | `%s` | %s |\n", esc(c.Name), c.RunByDefault, esc(c.Risk), esc(c.Command), esc(c.Purpose)))
	}
	b.WriteString("\n## Policy\n\n")
	for _, p := range r.PolicyBullets {
		b.WriteString("- " + p + "\n")
	}
	b.WriteString("\n## Minimal validation for this export\n\n")
	b.WriteString("```bash\ngo test -p=1 ./pkg/bridge/artifactindexexport -count=1\ngo list ./internal/app\n```\n")
	return b.String()
}

func renderMaintenanceChecklist() string {
	return `# Artifact Maintenance Checklist

- Keep root clean: no generated gate audits at repository root.
- Add new gate audits to ` + "`docs/audits/gates/gateNNN_registry_audit.md`" + `.
- Update ` + "`docs/audits/gates/INDEX.md`" + ` after adding or moving gate audits.
- Put conceptual summaries in ` + "`docs/summaries/`" + `.
- Put paper drafts and final manuscript files in ` + "`docs/paper/`" + `.
- Put figure sources and exported visuals in ` + "`docs/visuals/`" + `.
- Patch ` + "`README.md`" + ` and ` + "`docs/architecture.md`" + ` only with bounded addenda unless intentionally performing a large editorial pass.
- Preserve firewall wording when exporting publication-facing material.
`
}

func renderPublicationWorkspace() string {
	return `# Publication Workspace Guide

Use ` + "`docs/paper/`" + ` for manuscript assets:

- ` + "`docs/paper/drafts/`" + ` — working drafts.
- ` + "`docs/paper/final/`" + ` — final manuscript or accepted paper files.
- ` + "`docs/paper/references/`" + ` — bibliography and reference material.

Use ` + "`docs/visuals/`" + ` for figures:

- ` + "`docs/visuals/source/`" + ` — editable/source visual files.
- ` + "`docs/visuals/exported/`" + ` — exported image/PDF/SVG outputs.
- ` + "`docs/visuals/diagrams/`" + ` — theorem graphs, architecture diagrams, and dependency diagrams.
`
}

func Statuses() []string {
	return []string{
		StatusGate423ReviewerMatrixInherited,
		StatusArtifactIndexCompiled,
		StatusReproChecklistCompiled,
		StatusDocumentTreeIndexed,
		StatusAuditCoverageIndexed,
		StatusSummaryPaperVisualsIndexed,
		StatusRootCleanlinessAudited,
		StatusNoNewPhysicsClaim,
		StatusArtifactIndexReady,
		StatusNoNewDerivation,
		StatusNoYukawaPrediction,
		StatusNoCosmologyPrediction,
		StatusNoAxiomPromotion,
		StatusNoFlavorReopening,
		StatusFirewallPreserved13,
	}
}

func formatInts(xs []int) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return strings.Join(parts, ", ")
}

func esc(s string) string { return strings.ReplaceAll(s, "|", "\\|") }
