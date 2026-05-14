// Package executiveabstractclaimaudit implements Gate 422:
// Executive Abstract / Claim-Audit Summary Export.
//
// Gate 421 exported a section-by-section manuscript skeleton. Gate 422 turns
// that manuscript scaffold into front-matter language: executive abstract,
// claim classification, exact firewall wording, reviewer-safe non-claims, and
// a concise readiness summary. It is an exposition/export gate only. It adds no
// physics claim, predicts no flavor/cosmology coordinate, and promotes no
// quarantined axiom.
package executiveabstractclaimaudit

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE422-EXECUTIVE-ABSTRACT-CLAIM-AUDIT-SUMMARY-EXPORT"

	StatusGate421ManuscriptInherited = "CONDITIONAL_SUPPORT_GATE421_MANUSCRIPT_SKELETON_INHERITED"
	StatusExecutiveAbstractCompiled  = "CONDITIONAL_SUPPORT_EXECUTIVE_ABSTRACT_COMPILED"
	StatusClaimAuditCompiled         = "CONDITIONAL_SUPPORT_CLAIM_AUDIT_SUMMARY_COMPILED"
	StatusFirewallLanguageCompiled   = "CONDITIONAL_SUPPORT_FIREWALL_LANGUAGE_COMPILED"
	StatusReviewerWarningsCompiled   = "CONDITIONAL_SUPPORT_REVIEWER_SAFE_WARNINGS_COMPILED"
	StatusNonClaimLedgerCompiled     = "CONDITIONAL_SUPPORT_NON_CLAIM_LEDGER_COMPILED"
	StatusNoNewPhysicsClaim          = "CONDITIONAL_SUPPORT_NO_NEW_PHYSICS_CLAIM_IN_GATE422"
	StatusExecutiveSummaryReady      = "PROJECT_EXECUTIVE_CLAIM_AUDIT_SUMMARY_READY"

	StatusNoNewDerivation       = "FAILED_ROUTE_NO_NEW_DERIVATION_IN_GATE422"
	StatusNoYukawaPrediction    = "FAILED_ROUTE_NO_YUKAWA_COEFFICIENT_PREDICTION"
	StatusNoCosmologyPrediction = "FAILED_ROUTE_NO_COSMOLOGY_PREDICTION"
	StatusNoAxiomPromotion      = "FAILED_ROUTE_NO_QUARANTINED_AXIOM_PROMOTED_TO_NATIVE"
	StatusNoFlavorReopening     = "FAILED_ROUTE_NO_FLAVOR_REOPENING_IN_GATE422"
	StatusFirewallPreserved13   = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	NativeChargedFlavorDim    = 13
	ConditionalFamilyAxiomDim = 9
	ManuscriptSectionCount    = 13
	ManuscriptProofCount      = 26
)

type ClaimClass string

const (
	ClaimNative      ClaimClass = "native-theorem"
	ClaimBridge      ClaimClass = "bridge-lane"
	ClaimConditional ClaimClass = "quarantined-conditional-axiom"
	ClaimFirewall    ClaimClass = "firewall/environmental-boundary"
	ClaimFailedRoute ClaimClass = "failed-route/no-go"
	ClaimNonClaim    ClaimClass = "explicit-non-claim"
)

type Inheritance struct {
	Executed             bool
	Gate421Ready         bool
	SectionCount         int
	ProofObligationCount int
	NativeFlavorDim      int
	ConditionalFamilyDim int
	NoFlavorReopening    bool
	Verdict              string
}

type ClaimAuditRow struct {
	ID       string
	Class    ClaimClass
	Claim    string
	Evidence string
	Boundary string
	SafeVerb string
}

type ExecutiveAbstract struct {
	Executed             bool
	Title                string
	OneSentence          string
	ShortAbstract        string
	NativeClaims         []string
	ConditionalClaims    []string
	Firewalls            []string
	NonClaims            []string
	ReviewerWarnings     []string
	NativeFlavorDim      int
	ConditionalFamilyDim int
	NoNewPhysicsClaim    bool
	Verdict              string
}

type ClaimAudit struct {
	Executed          bool
	Rows              []ClaimAuditRow
	NativeCount       int
	BridgeCount       int
	ConditionalCount  int
	FirewallCount     int
	FailedRouteCount  int
	NonClaimCount     int
	FirewallsExplicit bool
	NoAxiomPromotion  bool
	Verdict           string
}

type ExportBundle struct {
	Executed               bool
	ExecutiveMarkdown      string
	ClaimAuditMarkdown     string
	FirewallMarkdown       string
	ReviewerMarkdown       string
	RecommendedFrontMatter []string
	PublicationReady       bool
	Verdict                string
}

type FinalStatus struct {
	Executed             bool
	ExecutiveReady       bool
	ClaimAuditReady      bool
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
	Abstract    ExecutiveAbstract
	ClaimAudit  ClaimAudit
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
	a.Abstract = buildAbstract()
	a.ClaimAudit = buildClaimAudit()
	a.Exports = buildExports(a.Abstract, a.ClaimAudit)
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
		Gate421Ready:         true,
		SectionCount:         ManuscriptSectionCount,
		ProofObligationCount: ManuscriptProofCount,
		NativeFlavorDim:      NativeChargedFlavorDim,
		ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		NoFlavorReopening:    true,
		Verdict:              "Gate 422 inherits the Gate-421 manuscript skeleton and compresses it into front-matter claim language without adding claims.",
	}
}

func buildAbstract() ExecutiveAbstract {
	native := []string{
		"finite Cℓ(1,7) measurement ladder and Boolean/G₂ contact vacuum K₇",
		"off-diagonal scalar/Higgs seed and one-form edge-supported Higgs carrier",
		"Fock matter carrier and Standard-Model electroweak charge skeleton with kY=5/3 and sin²θ*=3/8",
		"Morita finite spectral triple A_F=C⊕H⊕M₃(C), J/first-order compatibility, and SM gauge/Higgs field inventory",
		"almost-commutative M×F product architecture and CCM/Pfaffian/edge-measure bridge lanes",
	}
	conditional := []string{
		"K_gen gives quarantined hierarchy capacity but is not native ASHA data",
		"K/X gives quarantined real noncommuting mixing capacity",
		"K/X/Y gives quarantined CP-capable texture capacity with nine symbolic charged coefficients",
	}
	firewalls := []string{
		"native charged flavor moduli remain dimension 13",
		"nine K/X/Y source coefficients remain environmental boundary coordinates",
		"cosmological observables remain environmental/history-sensitive coordinates",
	}
	nonClaims := []string{
		"no Yukawa values are predicted",
		"no CKM angles or CKM CP phase are predicted",
		"no PMNS parameters are predicted",
		"no dark matter abundance, cosmological constant, or universe age is predicted",
		"no quarantined family axiom is promoted to native theorem",
	}
	warnings := []string{
		"Use 'derives' only for native finite-theorem layers; use 'conditional support' for bridge/axiom lanes.",
		"Do not state that ASHA predicts all Standard-Model parameters; flavor and cosmology are explicitly firewalled.",
		"Do not turn q4/contact invariants into Higgs/flavor claims; Gate 406 classifies q4 as contact-sector invariant.",
		"Do not identify color, chirality, triality, or Fock occupation with generation without a typed functor.",
	}
	return ExecutiveAbstract{
		Executed:      true,
		Title:         "ASHA Executive Claim Audit: Native Law-Space, Conditional Family Capacity, and Explicit Firewalls",
		OneSentence:   "ASHA is presented as a finite Clifford/almost-commutative law-space derivation with explicit bridge lanes and firewalls, not as a full prediction of flavor or cosmology.",
		ShortAbstract: "The ASHA theorem atlas organizes a finite Cℓ(1,7)-based construction of contact vacuum structure, electroweak charge logic, finite spectral-triple field content, and Higgs/scale bridge lanes. Gates 393–418 audit the flavor frontier and show that native ASHA preserves 13 charged flavor moduli; a quarantined K/X/Y family axiom chain supplies hierarchy, mixing, and CP capacity with nine symbolic charged coefficients, but the coefficient values remain environmental. Gate 422 exports concise claim language for manuscript front matter and reviewer-facing summaries while preserving all firewalls.",
		NativeClaims:  native, ConditionalClaims: conditional, Firewalls: firewalls, NonClaims: nonClaims, ReviewerWarnings: warnings,
		NativeFlavorDim: NativeChargedFlavorDim, ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		NoNewPhysicsClaim: true,
		Verdict:           "Executive abstract compiled with native, conditional, firewall, and non-claim language separated.",
	}
}

func buildClaimAudit() ClaimAudit {
	rows := []ClaimAuditRow{
		{"C1", ClaimNative, "Cℓ(1,7), Λ•R⁸, Boolean/G₂ contact geometry build the finite internal arena.", "Gates 0–6 and tower Floor 0–1.", "Definitions/contact vacuum only; no spacetime dynamics or physical constants here.", "derives"},
		{"C2", ClaimNative, "Off-diagonal projected connection geometry yields the scalar/contact Higgs seed.", "Gates 10–12 and 37.", "Pair-degenerate scalar response is not a flavor selector.", "derives"},
		{"C3", ClaimNative, "Matter/electroweak audits recover the charge skeleton, SU(2)L ladders, kY=5/3, and sin²θ*=3/8 boundary value.", "Gates 13–26 and 41.", "Boundary value still needs bridge/RG context for low-energy comparison.", "derives/audits"},
		{"C4", ClaimNative, "Morita finite spectral triple yields SM gauge and Higgs field inventory through inner fluctuations.", "Gates 272–299.", "Generation texture is not derived by the finite spectral triple.", "derives"},
		{"C5", ClaimBridge, "CCM/Pfaffian/edge-measure lanes produce the current Higgs tree-proxy board.", "Gates 376–387 and Gate 419 board.", "Bridge conventions must remain explicit; no loop-corrected pole-mass claim.", "conditionally supports"},
		{"C6", ClaimFirewall, "Native ASHA preserves the charged flavor frontier at dim M_charged=13.", "Gates 372, 374, 387, 393–410, 418.", "No Yukawa coefficients, CKM angles, or CP phase are predicted natively.", "seals"},
		{"C7", ClaimConditional, "K/X/Y family axioms give hierarchy, mixing, and CP capacity with nine symbolic charged coefficients.", "Gates 412–418.", "Axioms are quarantined; coefficient values are environmental.", "conditionally supports"},
		{"C8", ClaimFirewall, "Cosmological quantities remain environmental/history-sensitive frontiers.", "Gates 344, 375, 386, 419–421.", "No dark sector/cosmological constant prediction is claimed.", "seals"},
		{"C9", ClaimFailedRoute, "q4 is an exact contact-sector invariant, not an H_phi or flavor selector.", "Gates 398–406.", "Do not reuse q4 as scalar/Yukawa evidence without a new functor.", "rejects"},
		{"C10", ClaimFailedRoute, "H_phi native functionals are central or pair-degenerate as selected observables.", "Gates 407–408.", "Full End_R(H_phi) capacity does not determine coefficients.", "rejects"},
		{"C11", ClaimFailedRoute, "Existing fermionic carrier treats generation as a trivial U(3) multiplicity.", "Gates 409–410.", "Color/chirality/triality are not generation origins under current functors.", "rejects"},
		{"C12", ClaimNonClaim, "ASHA is not claimed to predict all Standard-Model numerical parameters.", "Gate 422 claim-audit policy.", "Flavor/cosmology coordinates remain firewalled.", "does not claim"},
	}
	ca := ClaimAudit{Executed: true, Rows: rows, FirewallsExplicit: true, NoAxiomPromotion: true, Verdict: "Claim audit compiled with native, bridge, conditional, firewall, failed-route, and non-claim rows separated."}
	for _, r := range rows {
		switch r.Class {
		case ClaimNative:
			ca.NativeCount++
		case ClaimBridge:
			ca.BridgeCount++
		case ClaimConditional:
			ca.ConditionalCount++
		case ClaimFirewall:
			ca.FirewallCount++
		case ClaimFailedRoute:
			ca.FailedRouteCount++
		case ClaimNonClaim:
			ca.NonClaimCount++
		}
	}
	return ca
}

func buildExports(abs ExecutiveAbstract, ca ClaimAudit) ExportBundle {
	front := []string{
		abs.Title,
		abs.OneSentence,
		"Native claims: " + fmt.Sprint(len(abs.NativeClaims)),
		"Conditional family claims: " + fmt.Sprint(len(abs.ConditionalClaims)),
		"Explicit firewalls: " + fmt.Sprint(len(abs.Firewalls)),
		"Explicit non-claims: " + fmt.Sprint(len(abs.NonClaims)),
	}
	return ExportBundle{
		Executed:               true,
		ExecutiveMarkdown:      renderExecutive(abs),
		ClaimAuditMarkdown:     renderClaimAudit(ca),
		FirewallMarkdown:       renderFirewalls(abs),
		ReviewerMarkdown:       renderReviewerWarnings(abs),
		RecommendedFrontMatter: front,
		PublicationReady:       true,
		Verdict:                "Front-matter abstract, claim-audit table, firewall language, and reviewer warnings exported.",
	}
}

func buildFinal(a Analysis) FinalStatus {
	ready := a.Inheritance.Executed && a.Abstract.Executed && a.ClaimAudit.Executed && a.Exports.PublicationReady
	return FinalStatus{
		Executed:             true,
		ExecutiveReady:       ready,
		ClaimAuditReady:      len(a.ClaimAudit.Rows) >= 10 && a.ClaimAudit.FirewallsExplicit,
		FirewallsPreserved:   a.Inheritance.NativeFlavorDim == NativeChargedFlavorDim && a.Abstract.NativeFlavorDim == NativeChargedFlavorDim,
		NoNewPhysicsClaim:    a.Abstract.NoNewPhysicsClaim,
		NoAxiomPromotion:     a.ClaimAudit.NoAxiomPromotion,
		NativeFlavorDim:      NativeChargedFlavorDim,
		ConditionalFamilyDim: ConditionalFamilyAxiomDim,
		Status:               StatusExecutiveSummaryReady,
		Verdict:              "Executive claim-audit summary ready; no native theorem frontier changed.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        423,
		Title:       "Reviewer Objection Matrix / Rebuttal Readiness Export",
		Reason:      "Gate 422 produces concise claim/firewall front matter. The next useful artifact is a reviewer-facing objection matrix with exact rebuttal boundaries and source-to-gate references.",
		PrimaryTask: "Map likely peer-review objections to the theorem atlas, failed-route index, firewalls, and exact language that avoids overclaiming.",
	}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate421Ready, "missing Gate421 inheritance"},
		{a.Abstract.Executed && a.Abstract.NoNewPhysicsClaim, "executive abstract not firewalled"},
		{len(a.Abstract.NativeClaims) >= 5, "too few native claims"},
		{len(a.Abstract.Firewalls) >= 3, "too few firewalls"},
		{len(a.Abstract.NonClaims) >= 5, "too few non-claims"},
		{a.ClaimAudit.NativeCount >= 4, "claim audit lacks native claims"},
		{a.ClaimAudit.FirewallCount >= 2, "claim audit lacks firewalls"},
		{a.ClaimAudit.FailedRouteCount >= 3, "claim audit lacks failed routes"},
		{a.Final.ExecutiveReady && a.Final.ClaimAuditReady, "final summary not ready"},
		{a.Final.FirewallsPreserved && a.Final.NativeFlavorDim == NativeChargedFlavorDim, "flavor firewall not preserved"},
		{a.Next.Gate == 423, "unexpected next gate"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate421ManuscriptInherited,
		StatusExecutiveAbstractCompiled,
		StatusClaimAuditCompiled,
		StatusFirewallLanguageCompiled,
		StatusReviewerWarningsCompiled,
		StatusNonClaimLedgerCompiled,
		StatusNoNewPhysicsClaim,
		StatusExecutiveSummaryReady,
		StatusNoNewDerivation,
		StatusNoYukawaPrediction,
		StatusNoCosmologyPrediction,
		StatusNoAxiomPromotion,
		StatusNoFlavorReopening,
		StatusFirewallPreserved13,
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 422 exports executive claim language only. Native charged flavor remains dim %d; the quarantined K/X/Y family chain remains %d symbolic charged coefficients; no Yukawa values, CKM/PMNS parameters, cosmology coordinates, or quarantined axioms are promoted.", a.Final.NativeFlavorDim, a.Final.ConditionalFamilyDim)
}

func join(items []string) string { return strings.Join(items, "; ") }
