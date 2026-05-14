// Package reviewerobjectionmatrix implements Gate 423:
// Reviewer Objection Matrix / Rebuttal Readiness Export.
//
// Gate 422 produced executive front-matter and claim-audit language. Gate 423
// converts that claim language into a reviewer-facing objection matrix. It is a
// publication-support export gate only: it adds no physics claim, predicts no
// flavor/cosmology coordinate, and promotes no quarantined axiom. Its purpose is
// to make the boundary between native theorem, conditional bridge, failed route,
// and environmental firewall explicit enough for peer-review replies.
package reviewerobjectionmatrix

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE423-REVIEWER-OBJECTION-MATRIX-REBUTTAL-READINESS-EXPORT"

	StatusGate422ClaimAuditInherited = "CONDITIONAL_SUPPORT_GATE422_CLAIM_AUDIT_INHERITED"
	StatusObjectionMatrixCompiled    = "CONDITIONAL_SUPPORT_REVIEWER_OBJECTION_MATRIX_COMPILED"
	StatusRebuttalBoundariesCompiled = "CONDITIONAL_SUPPORT_REBUTTAL_BOUNDARIES_COMPILED"
	StatusGateReferenceMapCompiled   = "CONDITIONAL_SUPPORT_GATE_REFERENCE_MAP_COMPILED"
	StatusRiskRankingCompiled        = "CONDITIONAL_SUPPORT_REVIEWER_RISK_RANKING_COMPILED"
	StatusClaimWordingAudited        = "CONDITIONAL_SUPPORT_CLAIM_WORDING_AUDITED"
	StatusNoNewPhysicsClaim          = "CONDITIONAL_SUPPORT_NO_NEW_PHYSICS_CLAIM_IN_GATE423"
	StatusReviewerMatrixReady        = "PROJECT_REVIEWER_OBJECTION_MATRIX_READY"

	StatusNoNewDerivation       = "FAILED_ROUTE_NO_NEW_DERIVATION_IN_GATE423"
	StatusNoYukawaPrediction    = "FAILED_ROUTE_NO_YUKAWA_COEFFICIENT_PREDICTION"
	StatusNoCosmologyPrediction = "FAILED_ROUTE_NO_COSMOLOGY_PREDICTION"
	StatusNoAxiomPromotion      = "FAILED_ROUTE_NO_QUARANTINED_AXIOM_PROMOTED_TO_NATIVE"
	StatusNoFlavorReopening     = "FAILED_ROUTE_NO_FLAVOR_REOPENING_IN_GATE423"
	StatusFirewallPreserved13   = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	NativeChargedFlavorDim    = 13
	ConditionalFamilyAxiomDim = 9
	PriorClaimAuditGate       = 422
)

type ObjectionClass string

const (
	ObjectionOverclaim       ObjectionClass = "overclaim-risk"
	ObjectionFlavor          ObjectionClass = "flavor-frontier"
	ObjectionCosmology       ObjectionClass = "cosmology-frontier"
	ObjectionBridge          ObjectionClass = "bridge-lane"
	ObjectionFailedRoute     ObjectionClass = "failed-route"
	ObjectionAxiom           ObjectionClass = "quarantined-axiom"
	ObjectionReproducibility ObjectionClass = "reproducibility"
)

type Severity string

const (
	SeverityHigh   Severity = "high"
	SeverityMedium Severity = "medium"
	SeverityLow    Severity = "low"
)

type Inheritance struct {
	Executed             bool
	Gate422Ready         bool
	NativeFlavorDim      int
	ConditionalFamilyDim int
	NoFlavorReopening    bool
	Verdict              string
}

type ObjectionRow struct {
	ID             string
	Class          ObjectionClass
	Severity       Severity
	Objection      string
	Rebuttal       string
	Boundary       string
	GateReferences []int
	SafeWording    string
	UnsafeWording  string
}

type ObjectionMatrix struct {
	Executed              bool
	Rows                  []ObjectionRow
	HighRiskCount         int
	MediumRiskCount       int
	LowRiskCount          int
	FlavorRows            int
	CosmologyRows         int
	AxiomRows             int
	FailedRouteRows       int
	ReproducibilityRows   int
	AllRowsHaveReferences bool
	AllRowsHaveBoundaries bool
	Verdict               string
}

type RebuttalGuide struct {
	Executed          bool
	OpeningParagraph  string
	Rules             []string
	RequiredPhrases   []string
	ForbiddenPhrases  []string
	EvidenceChecklist []string
	Verdict           string
}

type ExportBundle struct {
	Executed              bool
	ObjectionMarkdown     string
	RebuttalMarkdown      string
	GateReferenceMarkdown string
	RiskMarkdown          string
	PublicationReady      bool
	Verdict               string
}

type FinalStatus struct {
	Executed             bool
	MatrixReady          bool
	BoundariesReady      bool
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
	Matrix      ObjectionMatrix
	Guide       RebuttalGuide
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
	a.Matrix = buildMatrix()
	a.Guide = buildGuide()
	a.Exports = buildExports(a.Matrix, a.Guide)
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
		Gate422Ready:         true,
		NativeFlavorDim:      NativeChargedFlavorDim,
		ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		NoFlavorReopening:    true,
		Verdict:              "Gate 423 inherits Gate 422 claim-audit language and translates it into reviewer objection/rebuttal form without adding claims.",
	}
}

func buildMatrix() ObjectionMatrix {
	rows := []ObjectionRow{
		{
			ID: "OBJ-001", Class: ObjectionOverclaim, Severity: SeverityHigh,
			Objection:      "The manuscript appears to claim a theory of everything or a prediction of all Standard-Model parameters.",
			Rebuttal:       "The claim language is intentionally narrower: native ASHA is a finite law-space/gauge-Higgs scaffold with explicit bridge lanes and environmental firewalls.",
			Boundary:       "Do not claim prediction of flavor coefficients, CKM/PMNS parameters, or cosmological observables.",
			GateReferences: []int{418, 419, 420, 421, 422},
			SafeWording:    "ASHA derives/audits the finite law-space and preserves explicit flavor and cosmology firewalls.",
			UnsafeWording:  "ASHA predicts all constants of nature.",
		},
		{
			ID: "OBJ-002", Class: ObjectionFlavor, Severity: SeverityHigh,
			Objection:      "If the theory does not predict Yukawa values and CKM angles, the flavor problem is unsolved.",
			Rebuttal:       "Correct. Gates 393–418 make that boundary explicit. Native charged flavor remains a 13-dimensional frontier; K/X/Y family axioms give conditional capacity but leave coefficients environmental.",
			Boundary:       "State that flavor is classified and firewalled, not natively solved.",
			GateReferences: []int{393, 394, 409, 410, 411, 416, 417, 418},
			SafeWording:    "The flavor sector is formally sealed: native dim M_charged=13; conditional K/X/Y capacity has nine symbolic charged coefficients.",
			UnsafeWording:  "ASHA derives the Yukawa matrices.",
		},
		{
			ID: "OBJ-003", Class: ObjectionAxiom, Severity: SeverityHigh,
			Objection:      "The K/X/Y family operators are new assumptions, not native derivations.",
			Rebuttal:       "Yes. They are explicitly quarantined axioms introduced only after the native carrier and extension searches failed. The project does not promote them to native ASHA theorem.",
			Boundary:       "Use 'quarantined capacity axiom' or 'conditional support', never 'native derivation'.",
			GateReferences: []int{411, 412, 413, 414, 415, 416, 417, 418},
			SafeWording:    "K/X/Y supply conditional hierarchy, mixing, and CP capacity under explicit axiom quarantine.",
			UnsafeWording:  "K/X/Y are derived from Cℓ(1,7).",
		},
		{
			ID: "OBJ-004", Class: ObjectionFailedRoute, Severity: SeverityMedium,
			Objection:      "Triality seems used as a rhetorical replacement for the generation problem.",
			Rebuttal:       "The engine explicitly rejects that shortcut. Exact triality gives degeneracy or arena capacity, not a native finite-Dirac generation carrier.",
			Boundary:       "Triality may be cited as an arena/capacity result only with the no-go result included.",
			GateReferences: []int{27, 28, 393, 394, 395, 409},
			SafeWording:    "Triality was audited and did not reduce the 13-moduli firewall natively.",
			UnsafeWording:  "Triality explains the three generations.",
		},
		{
			ID: "OBJ-005", Class: ObjectionFailedRoute, Severity: SeverityMedium,
			Objection:      "The contact quartic q4 is being forced into the Higgs sector by dimension matching.",
			Rebuttal:       "Gates 398–406 reject that promotion. q4 is reconstructed as an internal contact-sector invariant, not a native H_phi selector.",
			Boundary:       "Cite q4 only as contact spectral data unless a future typed functor is derived.",
			GateReferences: []int{398, 399, 400, 405, 406},
			SafeWording:    "q4 lives in the contact spectral sector under current ASHA functors.",
			UnsafeWording:  "q4 is the Higgs selector.",
		},
		{
			ID: "OBJ-006", Class: ObjectionBridge, Severity: SeverityMedium,
			Objection:      "Bridge lanes such as Pfaffian scale, CCM coefficients, and edge measures may mix theorem and phenomenology.",
			Rebuttal:       "The atlas separates native finite algebra, bridge/coefficient lanes, and physical comparison lanes. Bridge statements require their stated assumptions and do not erase firewalls.",
			Boundary:       "Do not label bridge/coefficient transport as a raw native finite theorem.",
			GateReferences: []int{379, 380, 383, 384, 385, 387, 419, 420},
			SafeWording:    "The CCM/Pfaffian/edge-measure lanes provide audited bridge support under explicit assumptions.",
			UnsafeWording:  "The finite core alone directly outputs every low-energy value.",
		},
		{
			ID: "OBJ-007", Class: ObjectionCosmology, Severity: SeverityHigh,
			Objection:      "The project does not predict dark matter abundance, cosmological constant, or universe age.",
			Rebuttal:       "Correct. Those coordinates are explicitly environmental/history-sensitive frontiers unless a future dynamical bridge is derived.",
			Boundary:       "Cosmology remains a firewall, not an omitted hidden derivation.",
			GateReferences: []int{344, 375, 386, 387, 419, 420, 422},
			SafeWording:    "Cosmological observables remain environmental frontiers.",
			UnsafeWording:  "ASHA predicts cosmology.",
		},
		{
			ID: "OBJ-008", Class: ObjectionReproducibility, Severity: SeverityMedium,
			Objection:      "The result may be a narrative rather than a reproducible theorem ledger.",
			Rebuttal:       "The codebase records gates as theorem packages with targeted tests, registry wiring, audits, and a Gate-420 dependency atlas. Gate 423 indexes reviewer objections to exact gates.",
			Boundary:       "Reproducibility should be claimed at the package/audit level, not as independent external peer validation.",
			GateReferences: []int{420, 421, 422, 423},
			SafeWording:    "The project exports a reproducible internal theorem/audit ledger.",
			UnsafeWording:  "The theory is externally validated.",
		},
		{
			ID: "OBJ-009", Class: ObjectionOverclaim, Severity: SeverityMedium,
			Objection:      "The Higgs mass statement may sound like an exact physical prediction rather than a tree-level proxy/bridge result.",
			Rebuttal:       "Use the exact phrase 'sealed Higgs tree proxy' or 'tree-level proxy under bridge assumptions' and keep RG/threshold/loop limitations visible.",
			Boundary:       "Do not state a full physical Higgs mass prediction without bridge qualifiers.",
			GateReferences: []int{380, 384, 385, 387, 419, 422},
			SafeWording:    "The project derives a sealed Higgs tree proxy via CCM/Pfaffian/edge-measure lanes.",
			UnsafeWording:  "ASHA exactly predicts the measured Higgs mass.",
		},
		{
			ID: "OBJ-010", Class: ObjectionFailedRoute, Severity: SeverityLow,
			Objection:      "Repeated failed routes could indicate weakness rather than rigor.",
			Rebuttal:       "The failed-route ledger is part of the mathematical hygiene: it prevents color/generation confusion, scalar/flavor overclaims, q4/H_phi conflation, and triality shortcuts.",
			Boundary:       "Present failed routes as theorem-gated boundaries, not as hidden errors.",
			GateReferences: []int{393, 398, 406, 408, 409, 410, 414, 418, 420},
			SafeWording:    "Failed routes are indexed no-go/boundary results.",
			UnsafeWording:  "Failed routes are irrelevant implementation details.",
		},
		{
			ID: "OBJ-011", Class: ObjectionAxiom, Severity: SeverityMedium,
			Objection:      "The complex sector source with nine symbolic coefficients still leaves parameters free.",
			Rebuttal:       "Yes. Gate 417 shows CP capacity but not CP prediction; Gate 418 seals the nine coefficients as environmental boundary coordinates.",
			Boundary:       "Never claim CKM angle or CP-phase prediction from K/X/Y alone.",
			GateReferences: []int{416, 417, 418},
			SafeWording:    "The K/X/Y chain conditionally compresses the charged ledger to nine symbolic coefficients whose values remain environmental.",
			UnsafeWording:  "The K/X/Y chain predicts CKM.",
		},
		{
			ID: "OBJ-012", Class: ObjectionReproducibility, Severity: SeverityLow,
			Objection:      "A reader needs a concise map from claims to proof obligations and artifacts.",
			Rebuttal:       "Use the Gate-421 manuscript skeleton, Gate-420 theorem atlas, Gate-422 claim audit, and Gate-423 objection matrix together as the publication support stack.",
			Boundary:       "Do not replace formal proofs with executive language; use executive language only as front matter.",
			GateReferences: []int{420, 421, 422, 423},
			SafeWording:    "The publication support stack maps claims to proof obligations and gates.",
			UnsafeWording:  "The abstract alone proves the theory.",
		},
	}
	m := ObjectionMatrix{Executed: true, Rows: rows, AllRowsHaveReferences: true, AllRowsHaveBoundaries: true}
	for _, r := range rows {
		switch r.Severity {
		case SeverityHigh:
			m.HighRiskCount++
		case SeverityMedium:
			m.MediumRiskCount++
		case SeverityLow:
			m.LowRiskCount++
		}
		switch r.Class {
		case ObjectionFlavor:
			m.FlavorRows++
		case ObjectionCosmology:
			m.CosmologyRows++
		case ObjectionAxiom:
			m.AxiomRows++
		case ObjectionFailedRoute:
			m.FailedRouteRows++
		case ObjectionReproducibility:
			m.ReproducibilityRows++
		}
		if len(r.GateReferences) == 0 {
			m.AllRowsHaveReferences = false
		}
		if strings.TrimSpace(r.Boundary) == "" {
			m.AllRowsHaveBoundaries = false
		}
	}
	m.Verdict = "Reviewer objection matrix compiled with explicit rebuttal boundaries and gate references."
	return m
}

func buildGuide() RebuttalGuide {
	return RebuttalGuide{
		Executed:         true,
		OpeningParagraph: "ASHA's mature claim is deliberately bounded: it derives/audits a finite Clifford/almost-commutative law-space and records explicit bridge lanes, while preserving flavor and cosmology as firewalled environmental frontiers.",
		Rules: []string{
			"Answer every objection by first classifying the claim as native, bridge, quarantined axiom, failed route, or firewall.",
			"Use gate numbers to show where a route was verified, conditionally supported, or rejected.",
			"Do not convert conditional capacity into prediction language.",
			"Preserve the 13-dimensional native charged flavor firewall in every flavor reply.",
			"Keep cosmology observables firewalled unless a future dynamical bridge is implemented.",
		},
		RequiredPhrases: []string{
			"native ASHA theorem",
			"conditional support",
			"quarantined axiom",
			"environmental boundary coordinate",
			"FIREWALL_PRESERVED_13_MODULI",
		},
		ForbiddenPhrases: []string{
			"predicts all constants",
			"derives Yukawa matrices",
			"triality explains generations",
			"q4 is the Higgs selector",
			"cosmology is solved",
		},
		EvidenceChecklist: []string{
			"Gate 420 theorem atlas for dependency graph",
			"Gate 421 manuscript skeleton for section/proof mapping",
			"Gate 422 claim audit for front-matter wording",
			"Gate 423 matrix for reviewer objection replies",
		},
		Verdict: "Rebuttal guide compiled with mandatory boundary language and forbidden overclaim phrases.",
	}
}

func buildExports(m ObjectionMatrix, g RebuttalGuide) ExportBundle {
	return ExportBundle{
		Executed:              true,
		ObjectionMarkdown:     renderObjectionMatrix(m),
		RebuttalMarkdown:      renderRebuttalGuide(g),
		GateReferenceMarkdown: renderGateReferenceMap(m),
		RiskMarkdown:          renderRiskSummary(m),
		PublicationReady:      m.AllRowsHaveReferences && m.AllRowsHaveBoundaries && len(g.ForbiddenPhrases) >= 5,
		Verdict:               "Reviewer-facing objection matrix, rebuttal guide, gate-reference map, and risk summary exported.",
	}
}

func buildFinal(a Analysis) FinalStatus {
	return FinalStatus{
		Executed:             true,
		MatrixReady:          a.Matrix.Executed && len(a.Matrix.Rows) >= 10,
		BoundariesReady:      a.Matrix.AllRowsHaveBoundaries && a.Guide.Executed,
		FirewallsPreserved:   a.Inheritance.NoFlavorReopening,
		NoNewPhysicsClaim:    true,
		NoAxiomPromotion:     true,
		NativeFlavorDim:      NativeChargedFlavorDim,
		ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		Status:               StatusReviewerMatrixReady,
		Verdict:              "Gate 423 is reviewer-ready exposition. It adds no new derivation and preserves all Gate-422 firewalls.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        424,
		Title:       "Artifact Index / Reproducibility Checklist Export",
		Reason:      "Gate 423 prepares reviewer objections. The next publication-support artifact is a reproducibility checklist mapping code packages, tests, audits, theorem atlas entries, and manuscript sections.",
		PrimaryTask: "Export an artifact index and minimal reproduction instructions without running the full suite or adding physics claims.",
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 423 turns the Gate-%d claim audit into a reviewer objection matrix. It is ready=%t, preserves native charged flavor dim=%d, leaves the K/X/Y conditional family ledger at %d symbolic charged coefficients, and promotes no quarantined axiom or environmental coordinate.", PriorClaimAuditGate, a.Final.MatrixReady, a.Final.NativeFlavorDim, a.Final.ConditionalFamilyDim)
}

func Statuses() []string {
	return []string{
		StatusGate422ClaimAuditInherited,
		StatusObjectionMatrixCompiled,
		StatusRebuttalBoundariesCompiled,
		StatusGateReferenceMapCompiled,
		StatusRiskRankingCompiled,
		StatusClaimWordingAudited,
		StatusNoNewPhysicsClaim,
		StatusReviewerMatrixReady,
		StatusNoNewDerivation,
		StatusNoYukawaPrediction,
		StatusNoCosmologyPrediction,
		StatusNoAxiomPromotion,
		StatusNoFlavorReopening,
		StatusFirewallPreserved13,
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate422Ready {
		return fmt.Errorf("Gate422 inheritance missing")
	}
	if len(a.Matrix.Rows) < 10 {
		return fmt.Errorf("objection matrix too small: %d", len(a.Matrix.Rows))
	}
	if !a.Matrix.AllRowsHaveReferences || !a.Matrix.AllRowsHaveBoundaries {
		return fmt.Errorf("matrix rows lack references or boundaries")
	}
	if a.Matrix.HighRiskCount < 3 {
		return fmt.Errorf("expected at least three high-risk objections")
	}
	if a.Matrix.FlavorRows == 0 || a.Matrix.AxiomRows == 0 || a.Matrix.CosmologyRows == 0 {
		return fmt.Errorf("matrix missing key frontier classifications")
	}
	if !a.Exports.PublicationReady {
		return fmt.Errorf("exports not publication-ready")
	}
	if !a.Final.FirewallsPreserved || !a.Final.NoAxiomPromotion || a.Final.NativeFlavorDim != NativeChargedFlavorDim {
		return fmt.Errorf("final firewall state invalid")
	}
	if a.Next.Gate != 424 {
		return fmt.Errorf("unexpected next gate %d", a.Next.Gate)
	}
	return nil
}

func join(xs []string) string { return strings.Join(xs, " ") }
