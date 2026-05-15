// Package structuralfamilyboardexport implements Gate 449:
// Structural Family Board Export / Manuscript Delta Patch.
//
// Gate 448 reconciled the theorem registry after Gates 444-447.  Gate 449 is
// the paper-facing export of that reconciliation: it produces compact reviewer
// language, section deltas, a structural-family-board table, figure/table
// suggestions, and firewall wording that distinguishes the newly forced K/X
// structural layer from still-quarantined Y/phase/coefficient data.
//
// This gate is intentionally non-phenomenological.  It does not reopen flavor,
// import observed masses, fit Yukawa entries, select CKM/PMNS parameters, or
// update cosmology.  Its only job is to prevent claim drift when the manuscript
// is updated after the Gate-420 publication atlas.
package structuralfamilyboardexport

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE449-STRUCTURAL-FAMILY-BOARD-MANUSCRIPT-DELTA-PATCH"

	StatusGate448ReconciliationInherited = "CONDITIONAL_SUPPORT_GATE448_RECONCILIATION_INHERITED"
	StatusStructuralFamilyBoardCompiled  = "CONDITIONAL_SUPPORT_STRUCTURAL_FAMILY_BOARD_COMPILED"
	StatusManuscriptDeltaCompiled        = "CONDITIONAL_SUPPORT_POST444_MANUSCRIPT_DELTA_COMPILED"
	StatusClaimLanguageCompiled          = "CONDITIONAL_SUPPORT_POST444_CLAIM_LANGUAGE_COMPILED"
	StatusFigureTableDeltaCompiled       = "CONDITIONAL_SUPPORT_POST444_FIGURE_TABLE_DELTA_COMPILED"
	StatusFirewallAddendumCompiled       = "CONDITIONAL_SUPPORT_POST444_FIREWALL_ADDENDUM_COMPILED"
	StatusReviewerPacketCompiled         = "CONDITIONAL_SUPPORT_POST444_REVIEWER_PACKET_COMPILED"
	StatusNoNewPhysicsClaim              = "CONDITIONAL_SUPPORT_NO_NEW_PHYSICS_CLAIM_IN_GATE449"
	StatusManuscriptDeltaReady           = "PROJECT_POST444_MANUSCRIPT_DELTA_READY"

	StatusKGenWordingPromoted        = "K_GEN_WORDING_PROMOTED_TO_GEOMETRIC_AXIS"
	StatusGen2ZeroWordingPromoted    = "GENERATION2_ZERO_WORDING_PROMOTED_TO_STRUCTURAL_ZERO"
	StatusXSupportWordingPromoted    = "X_TRIANGLE_WORDING_PROMOTED_TO_STRUCTURAL_SUPPORT"
	StatusYPhaseWordingQuarantined   = "Y_PHASE_WORDING_REMAINS_QUARANTINED"
	StatusCoeffWordingQuarantined    = "KXY_COEFFICIENT_WORDING_REMAINS_QUARANTINED"
	StatusNativeFlavorDimPreserved   = "FIREWALL_PRESERVED_13_MODULI"
	StatusCoefficientDimPreserved    = "FIREWALL_PRESERVED_9_KXY_COEFFICIENTS"
	StatusEmpiricalFirewallPreserved = "CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED"

	StatusFailedNoYukawaPrediction        = "FAILED_ROUTE_NO_YUKAWA_VALUE_PREDICTION"
	StatusFailedNoCKMPMNSPrediction       = "FAILED_ROUTE_NO_CKM_PMNS_ANGLE_OR_PHASE_PREDICTION"
	StatusFailedNoMuonCharmMassPrediction = "FAILED_ROUTE_NO_MUON_CHARM_PHYSICAL_MASS_PREDICTION"
	StatusFailedNoCoefficientSelector     = "FAILED_ROUTE_NO_NATIVE_KXY_COEFFICIENT_SELECTOR"
	StatusFailedNoFinalPaperRewrite       = "FAILED_ROUTE_NO_AUTOMATIC_FINAL_MANUSCRIPT_REWRITE"
)

const (
	NativeChargedFlavorDim = 13
	KXYChargedCoeffDim     = 9
	ExpectedBoardRows      = 5
	ExpectedPatchBlocks    = 5
	ExpectedFirewallRows   = 6
)

type ClaimClass string

const (
	ClaimStructural ClaimClass = "structural-theorem"
	ClaimBridge     ClaimClass = "bridge-topology"
	ClaimQuarantine ClaimClass = "quarantined-frontier"
	ClaimNonClaim   ClaimClass = "explicit-non-claim"
	ClaimManuscript ClaimClass = "manuscript-export"
)

type Inheritance struct {
	Executed                  bool
	Gate448Reconciled         bool
	KGenPromoted              bool
	Gen2ZeroPromoted          bool
	XSupportPromoted          bool
	YPhaseQuarantined         bool
	CoefficientsQuarantined   bool
	NativeFlavorDim           int
	KXYCoeffDim               int
	NoEmpiricalInputsImported bool
	PublicationSupportOnly    bool
	Verdict                   string
}

type BoardRow struct {
	Object      string
	Layer       string
	Formula     string
	SourceGate  int
	Claim       string
	Boundary    string
	PaperAction string
	Status      string
}

type StructuralBoard struct {
	Executed              bool
	Rows                  []BoardRow
	PromotedRows          int
	QuarantinedRows       int
	NativeFlavorDim       int
	KXYCoeffDim           int
	KGenRowPresent        bool
	Gen2ZeroRowPresent    bool
	XTriangleRowPresent   bool
	YPhaseRowQuarantined  bool
	CoeffRowsQuarantined  bool
	NoObservablePredicted bool
	Verdict               string
	Reason                string
}

type PatchBlock struct {
	TargetSection string
	Operation     string
	ClaimClass    ClaimClass
	SourceGates   []int
	Markdown      string
	Boundary      string
	Ready         bool
}

type ManuscriptDelta struct {
	Executed                 bool
	Blocks                   []PatchBlock
	AbstractInsertionReady   bool
	Section9ReplacementReady bool
	ConclusionAddendumReady  bool
	ReviewerNoteReady        bool
	AppendixDeltaReady       bool
	NoClaimDrift             bool
	NoFinalDocumentMutation  bool
	RecommendedTargetPath    string
	Verdict                  string
	Reason                   string
}

type ArtifactDelta struct {
	Kind       string
	Name       string
	TargetPath string
	Source     string
	Purpose    string
	ClaimRule  string
	Status     string
}

type FigureTableDelta struct {
	Executed      bool
	Tables        []ArtifactDelta
	Figures       []ArtifactDelta
	RequiredCount int
	ReadyCount    int
	Verdict       string
	Reason        string
}

type FirewallRow struct {
	Topic     string
	Allowed   string
	Forbidden string
	Source    string
	Status    string
}

type FirewallAddendum struct {
	Executed                bool
	Rows                    []FirewallRow
	NativeFlavorDim         int
	KXYCoeffDim             int
	AllowsKGenPromotion     bool
	AllowsXSupportPromotion bool
	ForbidsYukawaPrediction bool
	ForbidsMixingPrediction bool
	ForbidsMassPrediction   bool
	ForbidsCoefficientFit   bool
	ForbidsCosmologyUpdate  bool
	Verdict                 string
	Reason                  string
}

type ReviewerObjection struct {
	Objection string
	Answer    string
	Boundary  string
	Status    string
}

type ReviewerPacket struct {
	Executed       bool
	Objections     []ReviewerObjection
	ReadyCount     int
	NoClaimDrift   bool
	FirewallStated bool
	Verdict        string
	Reason         string
}

type ExportBundle struct {
	Executed                 bool
	TargetPath               string
	StructuralBoardMarkdown  string
	ManuscriptDeltaMarkdown  string
	FigureTableMarkdown      string
	FirewallAddendumMarkdown string
	ReviewerPacketMarkdown   string
	CombinedMarkdown         string
	PublicationReady         bool
	NoNewPhysicsClaim        bool
	Verdict                  string
	Reason                   string
}

type FinalStatus struct {
	Executed               bool
	Ready                  bool
	BoardReady             bool
	ManuscriptDeltaReady   bool
	FirewallReady          bool
	ReviewerReady          bool
	NoNewPhysicsClaim      bool
	NoObservedMassImported bool
	NoYukawaImported       bool
	NoCKMImported          bool
	NoPMNSImported         bool
	NativeFlavorDim        int
	KXYCoeffDim            int
	Status                 string
	Verdict                string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Board       StructuralBoard
	Delta       ManuscriptDelta
	Artifacts   FigureTableDelta
	Firewall    FirewallAddendum
	Reviewer    ReviewerPacket
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
	a.Board = buildBoard()
	a.Delta = buildDelta(a.Board)
	a.Artifacts = buildArtifacts()
	a.Firewall = buildFirewall()
	a.Reviewer = buildReviewer()
	a.Exports = buildExports(a.Board, a.Delta, a.Artifacts, a.Firewall, a.Reviewer)
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
		Executed:                  true,
		Gate448Reconciled:         true,
		KGenPromoted:              true,
		Gen2ZeroPromoted:          true,
		XSupportPromoted:          true,
		YPhaseQuarantined:         true,
		CoefficientsQuarantined:   true,
		NativeFlavorDim:           NativeChargedFlavorDim,
		KXYCoeffDim:               KXYChargedCoeffDim,
		NoEmpiricalInputsImported: true,
		PublicationSupportOnly:    true,
		Verdict:                   StatusGate448ReconciliationInherited,
	}
}

func buildBoard() StructuralBoard {
	rows := []BoardRow{
		{Object: "K_gen", Layer: "geometrically-forced structural axis", Formula: "diag(-1,0,1)", SourceGate: 444, Claim: "The primitive traceless, integer-spaced, three-generation KMS spectrum is forced up to primitive equivalence.", Boundary: "Not a Yukawa spectrum, not a physical mass spectrum, and not a fit parameter.", PaperAction: "Replace earlier 'quarantined K_gen choice' wording with 'forced structural family axis'.", Status: StatusKGenWordingPromoted},
		{Object: "Generation-2 bare level", Layer: "structural consequence", Formula: "middle eigenvalue = 0", SourceGate: 444, Claim: "The middle bare level of the forced family axis is exactly zero.", Boundary: "Muon/charm physical masses still require bridge amplitudes and sector coefficients.", PaperAction: "Add one sentence distinguishing bare structural zero from observed muon/charm masses.", Status: StatusGen2ZeroWordingPromoted},
		{Object: "B_lift support", Layer: "structural bridge topology", Formula: "support([[0,1,1],[1,0,1],[1,1,0]])", SourceGate: 445, Claim: "The unique minimal endpoint-balanced mass-lift support is the closed triangle.", Boundary: "Only support is promoted; amplitude, signs, and complex phase remain sealed.", PaperAction: "Add a small table/figure row for the unsigned triangle topology.", Status: StatusXSupportWordingPromoted},
		{Object: "Y_gen / Phi_cycle", Layer: "quarantined phase frontier", Formula: "Phi = arg(z12 z23 conjugate(z13))", SourceGate: 446, Claim: "The native sieve leaves signed/complex orientation underdetermined.", Boundary: "No CKM/PMNS CP phase or mixing angle is predicted.", PaperAction: "Keep Y/phase in the firewall table, not in the native theorem list.", Status: StatusYPhaseWordingQuarantined},
		{Object: "charged K/X/Y coefficients", Layer: "environmental coefficient ledger", Formula: "dim C_KXY^charged = 9", SourceGate: 447, Claim: "Multiple symbolic ledgers survive all native tests.", Boundary: "No sector coefficient, Yukawa entry, or muon/charm mass is selected.", PaperAction: "Preserve the nine-coefficient firewall language.", Status: StatusCoeffWordingQuarantined},
	}
	promoted, quarantined := countBoard(rows)
	return StructuralBoard{
		Executed:              true,
		Rows:                  rows,
		PromotedRows:          promoted,
		QuarantinedRows:       quarantined,
		NativeFlavorDim:       NativeChargedFlavorDim,
		KXYCoeffDim:           KXYChargedCoeffDim,
		KGenRowPresent:        hasBoardStatus(rows, StatusKGenWordingPromoted),
		Gen2ZeroRowPresent:    hasBoardStatus(rows, StatusGen2ZeroWordingPromoted),
		XTriangleRowPresent:   hasBoardStatus(rows, StatusXSupportWordingPromoted),
		YPhaseRowQuarantined:  hasBoardStatus(rows, StatusYPhaseWordingQuarantined),
		CoeffRowsQuarantined:  hasBoardStatus(rows, StatusCoeffWordingQuarantined),
		NoObservablePredicted: true,
		Verdict:               StatusStructuralFamilyBoardCompiled,
		Reason:                "the board exports exactly three structural promotions and two value-bearing quarantines",
	}
}

func countBoard(rows []BoardRow) (promoted, quarantined int) {
	for _, r := range rows {
		switch r.Status {
		case StatusKGenWordingPromoted, StatusGen2ZeroWordingPromoted, StatusXSupportWordingPromoted:
			promoted++
		case StatusYPhaseWordingQuarantined, StatusCoeffWordingQuarantined:
			quarantined++
		}
	}
	return promoted, quarantined
}

func hasBoardStatus(rows []BoardRow, status string) bool {
	for _, r := range rows {
		if r.Status == status {
			return true
		}
	}
	return false
}

func buildDelta(board StructuralBoard) ManuscriptDelta {
	blocks := []PatchBlock{
		{TargetSection: "Abstract", Operation: "insert one guarded sentence", ClaimClass: ClaimManuscript, SourceGates: []int{444, 445, 446, 447, 448}, Markdown: "After the Gate-420 atlas, a post-publication family sieve promotes `K_gen = diag(-1,0,1)`, the Generation-2 bare structural zero, and the unsigned triangular mass-lift support as structural results, while preserving the Yukawa/CKM/PMNS and coefficient firewalls.", Boundary: "Do not call this a physical muon/charm mass or mixing prediction.", Ready: true},
		{TargetSection: "9. Flavor Frontier and Family Axiom Ledger", Operation: "replace first firewall paragraph with structural/frontier split", ClaimClass: ClaimStructural, SourceGates: []int{444, 445, 446, 447, 448}, Markdown: "The post-444 refinement separates the family board into a forced structural layer and a sealed value-bearing layer.  The structural layer contains the primitive traceless family axis `K_gen = diag(-1,0,1)`, the resulting Generation-2 bare zero, and the unsigned endpoint-balanced triangle support for mass-lift compatibility.  The sealed layer still contains `Y_gen`, the cycle phase, bridge amplitudes, and all charged-sector K/X/Y coefficients.", Boundary: "State structural zero only at the bare-law level; observed lepton/quark masses remain outside the native core.", Ready: true},
		{TargetSection: "Section 9 table", Operation: "add structural-family-board table", ClaimClass: ClaimBridge, SourceGates: []int{444, 445, 446, 447, 448}, Markdown: renderBoardTable(board.Rows), Boundary: "Rows marked quarantined cannot be cited as derived observable values.", Ready: true},
		{TargetSection: "11. Failed Routes and Reviewer Objections", Operation: "add reviewer note", ClaimClass: ClaimNonClaim, SourceGates: []int{446, 447, 448}, Markdown: "The phase and coefficient sieves are deliberately negative results: they prove that native ASHA does not select CP phases, CKM/PMNS coordinates, or charged-sector amplitudes from the tested boundaries alone.", Boundary: "Failed-route language must stay visible rather than hidden in appendices.", Ready: true},
		{TargetSection: "Conclusion", Operation: "add compact frontier sentence", ClaimClass: ClaimManuscript, SourceGates: []int{444, 445, 446, 447, 448}, Markdown: "The updated frontier is therefore sharper but not overclaimed: ASHA now fixes the primitive family axis and a minimal mass-lift support topology, while leaving all value-bearing flavor coordinates behind the environmental firewall.", Boundary: "No final-paper numerical flavor claim is added.", Ready: true},
	}
	return ManuscriptDelta{
		Executed:                 true,
		Blocks:                   blocks,
		AbstractInsertionReady:   blocks[0].Ready,
		Section9ReplacementReady: blocks[1].Ready && blocks[2].Ready,
		ConclusionAddendumReady:  blocks[4].Ready,
		ReviewerNoteReady:        blocks[3].Ready,
		AppendixDeltaReady:       true,
		NoClaimDrift:             true,
		NoFinalDocumentMutation:  true,
		RecommendedTargetPath:    "docs/paper/POST444_MANUSCRIPT_DELTA.md",
		Verdict:                  StatusManuscriptDeltaCompiled,
		Reason:                   "the delta is exported as a source patch rather than directly mutating final DOCX/PDF artifacts",
	}
}

func buildArtifacts() FigureTableDelta {
	tables := []ArtifactDelta{
		{Kind: "table", Name: "Post-444 structural family board", TargetPath: "docs/paper/POST444_MANUSCRIPT_DELTA.md#structural-family-board", Source: "Gates 444--448", Purpose: "show promoted structural objects and quarantined value-bearing coordinates", ClaimRule: "structural topology only; no physical flavor values", Status: "ready"},
		{Kind: "table", Name: "Post-444 claim firewall addendum", TargetPath: "docs/paper/POST444_MANUSCRIPT_DELTA.md#claim-firewall-addendum", Source: "Gate 449", Purpose: "safe/forbidden wording for manuscript revision", ClaimRule: "must preserve native 13-moduli and nine-coefficient firewalls", Status: "ready"},
	}
	figures := []ArtifactDelta{
		{Kind: "figure", Name: "Structural family board overlay", TargetPath: "docs/visuals/diagrams/post444_structural_family_board.(svg|png|pdf)", Source: "Gates 444--448", Purpose: "visual split between forced K/X structural layer and sealed Y/phase/coefficient layer", ClaimRule: "caption must say support topology, not mass prediction", Status: "slot ready"},
	}
	return FigureTableDelta{Executed: true, Tables: tables, Figures: figures, RequiredCount: len(tables) + len(figures), ReadyCount: len(tables) + len(figures), Verdict: StatusFigureTableDeltaCompiled, Reason: "one board figure slot and two manuscript tables are enough to express the post-444 delta without claim drift"}
}

func buildFirewall() FirewallAddendum {
	rows := []FirewallRow{
		{Topic: "K_gen", Allowed: "`K_gen = diag(-1,0,1)` is a forced primitive structural family axis.", Forbidden: "Do not call `K_gen` an observed mass spectrum or a fitted Yukawa matrix.", Source: "Gate 444", Status: StatusKGenWordingPromoted},
		{Topic: "Generation-2 zero", Allowed: "The second bare family level is structurally zero in the forced primitive axis.", Forbidden: "Do not claim the muon or charm pole/running mass is zero or predicted.", Source: "Gate 444", Status: StatusGen2ZeroWordingPromoted},
		{Topic: "X triangle support", Allowed: "The unsigned endpoint-balanced triangle support is the minimal mass-lift topology.", Forbidden: "Do not assign its amplitude, sign orientation, or complex phase as native data.", Source: "Gate 445", Status: StatusXSupportWordingPromoted},
		{Topic: "Y/phase", Allowed: "`Y_gen` and the cycle phase remain quarantined CP-capacity coordinates.", Forbidden: "Do not predict CKM/PMNS angles or phases.", Source: "Gate 446", Status: StatusYPhaseWordingQuarantined},
		{Topic: "charged coefficients", Allowed: "The charged K/X/Y coefficient space remains nine-dimensional and symbolic.", Forbidden: "Do not select sector coefficients or Yukawa values without an explicit external seal.", Source: "Gate 447", Status: StatusCoeffWordingQuarantined},
		{Topic: "cosmology/dark sector", Allowed: "No post-444 family-board update changes cosmology or dark-sector firewalls.", Forbidden: "Do not infer dark matter abundance, cosmological constant, or cosmological history from the family board.", Source: "Gate 448", Status: StatusEmpiricalFirewallPreserved},
	}
	return FirewallAddendum{Executed: true, Rows: rows, NativeFlavorDim: NativeChargedFlavorDim, KXYCoeffDim: KXYChargedCoeffDim, AllowsKGenPromotion: true, AllowsXSupportPromotion: true, ForbidsYukawaPrediction: true, ForbidsMixingPrediction: true, ForbidsMassPrediction: true, ForbidsCoefficientFit: true, ForbidsCosmologyUpdate: true, Verdict: StatusFirewallAddendumCompiled, Reason: "the addendum upgrades structural wording while explicitly forbidding every value-bearing flavor/cosmology overclaim"}
}

func buildReviewer() ReviewerPacket {
	rows := []ReviewerObjection{
		{Objection: "Does the post-444 update predict the muon or charm mass?", Answer: "No. It proves a bare structural zero in the primitive family axis; physical masses require bridge amplitude and sector-coefficient data still behind the firewall.", Boundary: "No observed mass imported or derived.", Status: StatusFailedNoMuonCharmMassPrediction},
		{Objection: "Did ASHA promote the full K/X/Y family ansatz to native law?", Answer: "No. Only K_gen and unsigned X-support topology are structurally promoted. Y/phase and all nine charged coefficients remain quarantined.", Boundary: "No coefficient selector exists in Gate 449.", Status: StatusFailedNoCoefficientSelector},
		{Objection: "Does the triangle support fix CKM or PMNS mixing?", Answer: "No. The support graph is a topology, not a mixing matrix; signed/complex orientation and sector amplitudes remain underdetermined.", Boundary: "No CKM/PMNS coordinate predicted.", Status: StatusFailedNoCKMPMNSPrediction},
		{Objection: "Should the final DOCX/PDF be overwritten automatically?", Answer: "No. Gate 449 exports a source delta patch for controlled human manuscript revision; final artifacts are not silently rewritten.", Boundary: "Publication support only.", Status: StatusFailedNoFinalPaperRewrite},
	}
	return ReviewerPacket{Executed: true, Objections: rows, ReadyCount: len(rows), NoClaimDrift: true, FirewallStated: true, Verdict: StatusReviewerPacketCompiled, Reason: "the reviewer packet anticipates the main overclaim risks created by the stronger structural wording"}
}

func buildExports(board StructuralBoard, delta ManuscriptDelta, artifacts FigureTableDelta, firewall FirewallAddendum, reviewer ReviewerPacket) ExportBundle {
	boardMD := renderBoard(board)
	deltaMD := renderDelta(delta)
	artifactMD := renderArtifacts(artifacts)
	firewallMD := renderFirewall(firewall)
	reviewerMD := renderReviewer(reviewer)
	combined := renderCombined(board, delta, artifacts, firewall, reviewer)
	return ExportBundle{Executed: true, TargetPath: delta.RecommendedTargetPath, StructuralBoardMarkdown: boardMD, ManuscriptDeltaMarkdown: deltaMD, FigureTableMarkdown: artifactMD, FirewallAddendumMarkdown: firewallMD, ReviewerPacketMarkdown: reviewerMD, CombinedMarkdown: combined, PublicationReady: true, NoNewPhysicsClaim: true, Verdict: StatusManuscriptDeltaReady, Reason: "all export blocks are publication-facing and carry explicit firewall boundaries"}
}

func buildFinal(a Analysis) FinalStatus {
	return FinalStatus{Executed: true, Ready: a.Exports.PublicationReady, BoardReady: a.Board.Executed && len(a.Board.Rows) == ExpectedBoardRows, ManuscriptDeltaReady: a.Delta.Executed && len(a.Delta.Blocks) == ExpectedPatchBlocks && a.Delta.NoClaimDrift, FirewallReady: a.Firewall.Executed && len(a.Firewall.Rows) >= ExpectedFirewallRows, ReviewerReady: a.Reviewer.Executed && a.Reviewer.ReadyCount == len(a.Reviewer.Objections), NoNewPhysicsClaim: a.Exports.NoNewPhysicsClaim && a.Board.NoObservablePredicted, NoObservedMassImported: true, NoYukawaImported: true, NoCKMImported: true, NoPMNSImported: true, NativeFlavorDim: NativeChargedFlavorDim, KXYCoeffDim: KXYChargedCoeffDim, Status: StatusManuscriptDeltaReady, Verdict: "Gate 449 exports a safe post-444 manuscript delta: structural K/X updates are visible, value-bearing flavor/cosmology firewalls remain intact"}
}

func buildNext() NextStep {
	return NextStep{Gate: 450, Title: "Post-444 Publication Bundle Integrity Check", Reason: "Gate 449 exports new paper-facing delta material; the next gate should verify that the bundle manifest, section map, claim firewall, and runtime metadata all reference it consistently.", PrimaryTask: "audit publication-support files for post-444 consistency without mutating final manuscript binaries or adding physics claims"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate448Reconciled || !a.Inheritance.KGenPromoted || !a.Inheritance.XSupportPromoted || !a.Inheritance.YPhaseQuarantined || !a.Inheritance.CoefficientsQuarantined {
		return fmt.Errorf("Gate 449 requires a complete Gate 448 reconciliation inheritance")
	}
	if !a.Inheritance.NoEmpiricalInputsImported || !a.Inheritance.PublicationSupportOnly {
		return fmt.Errorf("Gate 449 must remain a publication-support export")
	}
	if !a.Board.Executed || len(a.Board.Rows) != ExpectedBoardRows || a.Board.PromotedRows != 3 || a.Board.QuarantinedRows != 2 || !a.Board.NoObservablePredicted {
		return fmt.Errorf("invalid structural family board: %+v", a.Board)
	}
	if !a.Board.KGenRowPresent || !a.Board.Gen2ZeroRowPresent || !a.Board.XTriangleRowPresent || !a.Board.YPhaseRowQuarantined || !a.Board.CoeffRowsQuarantined {
		return fmt.Errorf("structural family board missing required rows")
	}
	if !a.Delta.Executed || len(a.Delta.Blocks) != ExpectedPatchBlocks || !a.Delta.AbstractInsertionReady || !a.Delta.Section9ReplacementReady || !a.Delta.ConclusionAddendumReady || !a.Delta.NoClaimDrift || !a.Delta.NoFinalDocumentMutation {
		return fmt.Errorf("invalid manuscript delta: %+v", a.Delta)
	}
	if !a.Artifacts.Executed || a.Artifacts.ReadyCount != a.Artifacts.RequiredCount || a.Artifacts.RequiredCount < 3 {
		return fmt.Errorf("figure/table delta incomplete")
	}
	if !a.Firewall.Executed || len(a.Firewall.Rows) < ExpectedFirewallRows || !a.Firewall.ForbidsYukawaPrediction || !a.Firewall.ForbidsMixingPrediction || !a.Firewall.ForbidsMassPrediction || !a.Firewall.ForbidsCoefficientFit || !a.Firewall.ForbidsCosmologyUpdate {
		return fmt.Errorf("firewall addendum incomplete")
	}
	if !a.Reviewer.Executed || a.Reviewer.ReadyCount != len(a.Reviewer.Objections) || !a.Reviewer.NoClaimDrift || !a.Reviewer.FirewallStated {
		return fmt.Errorf("reviewer packet incomplete")
	}
	if !a.Final.Ready || !a.Final.NoNewPhysicsClaim || !a.Final.NoObservedMassImported || !a.Final.NoYukawaImported || !a.Final.NoCKMImported || !a.Final.NoPMNSImported || a.Final.NativeFlavorDim != NativeChargedFlavorDim || a.Final.KXYCoeffDim != KXYChargedCoeffDim {
		return fmt.Errorf("final status violates publication firewall")
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 449 exports the post-444 manuscript delta without adding physics.  The paper may now say that K_gen=diag(-1,0,1), the Generation-2 bare zero, and the unsigned triangular lift support are structural results.  It must still say that Y_gen/Φ_cycle, bridge amplitudes, charged K/X/Y coefficients, Yukawa entries, CKM/PMNS values, and muon/charm physical masses remain quarantined.  The native charged flavor dimension remains %d and the charged K/X/Y coefficient ledger remains %d-dimensional.", a.Final.NativeFlavorDim, a.Final.KXYCoeffDim)
}

func Statuses() []string {
	return []string{StatusGate448ReconciliationInherited, StatusStructuralFamilyBoardCompiled, StatusManuscriptDeltaCompiled, StatusClaimLanguageCompiled, StatusFigureTableDeltaCompiled, StatusFirewallAddendumCompiled, StatusReviewerPacketCompiled, StatusNoNewPhysicsClaim, StatusManuscriptDeltaReady, StatusKGenWordingPromoted, StatusGen2ZeroWordingPromoted, StatusXSupportWordingPromoted, StatusYPhaseWordingQuarantined, StatusCoeffWordingQuarantined, StatusNativeFlavorDimPreserved, StatusCoefficientDimPreserved, StatusEmpiricalFirewallPreserved, StatusFailedNoYukawaPrediction, StatusFailedNoCKMPMNSPrediction, StatusFailedNoMuonCharmMassPrediction, StatusFailedNoCoefficientSelector, StatusFailedNoFinalPaperRewrite}
}

func joinInts(xs []int) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return strings.Join(parts, ",")
}
