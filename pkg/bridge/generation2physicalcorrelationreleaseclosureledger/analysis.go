// Package generation2physicalcorrelationreleaseclosureledger implements Gate 548:
// Physical Correlation Import/Release Sector Closure Ledger.
//
// Gate 547 proved that even a checksum-verified synthetic release-review
// manifest cannot release bridge evidence. Gate 548 closes the entire
// Schwinger/source/comparator/release pipeline as a frontier ledger: it records
// which objects remain native, which are bridge sockets, which are environmental
// inputs, and which shortcuts are permanently failed routes. No physical
// correlation data is imported, no bridge evidence is released, and no native
// theorem is written.
package generation2physicalcorrelationreleaseclosureledger

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticreleasereviewmanifestadapter"
)

const (
	AuditID = "GATE548-PHYSICAL-CORRELATION-IMPORT-RELEASE-SECTOR-CLOSURE-LEDGER"

	StatusGate547SyntheticReleaseInherited = "CONDITIONAL_SUPPORT_GATE547_SYNTHETIC_RELEASE_REVIEW_INHERITED"
	StatusClosureLedgerEmitted             = "CONDITIONAL_SUPPORT_PHYSICAL_CORRELATION_IMPORT_RELEASE_SECTOR_CLOSURE_LEDGER_EMITTED"
	StatusNativeFrontierFrozen             = "CONDITIONAL_SUPPORT_PHYSICAL_CORRELATION_NATIVE_FRONTIER_FROZEN"
	StatusBridgeFrontierMapped             = "CONDITIONAL_SUPPORT_PHYSICAL_CORRELATION_BRIDGE_FRONTIER_MAPPED"
	StatusEnvironmentalFrontierMapped      = "CONDITIONAL_SUPPORT_PHYSICAL_CORRELATION_ENVIRONMENTAL_FRONTIER_MAPPED"
	StatusSchwingerSourceBlockClosed       = "CONDITIONAL_SUPPORT_SCHWINGER_SOURCE_SCHEMA_BLOCK_CLOSED"
	StatusAuthenticityBlockClosed          = "CONDITIONAL_SUPPORT_SOURCE_AUTHENTICITY_BLOCK_CLOSED"
	StatusImportSwitchBlockClosed          = "CONDITIONAL_SUPPORT_REAL_IMPORT_SWITCH_BLOCK_CLOSED"
	StatusComparatorHarnessBlockClosed     = "CONDITIONAL_SUPPORT_COMPARATOR_HARNESS_BLOCK_CLOSED"
	StatusReleaseReviewBlockClosed         = "CONDITIONAL_SUPPORT_RELEASE_REVIEW_BLOCK_CLOSED"
	StatusNoBridgeEvidenceReleased         = "CONDITIONAL_SUPPORT_NO_PHYSICAL_CORRELATION_BRIDGE_EVIDENCE_RELEASED_IN_GATE548"
	StatusClosureFirewallMatrixComplete    = "CONDITIONAL_SUPPORT_GATE548_PHYSICAL_CORRELATION_FIREWALL_MATRIX_COMPLETE"

	StatusFailedClosureDoesNotDeriveSchwinger = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedClosureDoesNotProveOS         = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedClosureDoesNotGrantWick       = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedClosureDoesNotSelectHilbert   = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedClosureDoesNotDeriveHamilton  = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedClosureDoesNotGrantUnitary    = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedClosureDoesNotGrantGlobal     = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedClosureDoesNotSelectArrow     = "FAILED_ROUTE_SECTOR_CLOSURE_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedClosureNoRealSource           = "FAILED_ROUTE_NO_AUTHENTICATED_NON_SYNTHETIC_CORRELATION_SOURCE_IN_GATE548"
	StatusFailedClosureNoRelease              = "FAILED_ROUTE_NO_RELEASED_PHYSICAL_CORRELATION_BRIDGE_EVIDENCE_IN_GATE548"
	StatusFirewallPreserved                   = "FIREWALL_PRESERVED_GATE548_PHYSICAL_CORRELATION_SECTOR_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked          = "FIREWALL_BLOCKED_GATE548_PHYSICAL_CORRELATION_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate547ManifestParsed        bool
	Gate547RowsAccepted          int
	Gate547ChecksumVerified      bool
	Gate547HumanReviewParsed     bool
	Gate547ReproducibilityParsed bool
	Gate547SourceChainParsed     bool
	Gate547SyntheticBlocked      bool
	Gate547NoBridgeEvidence      bool
	Gate547NoRealSource          bool
	Gate547NativeWriteLocked     bool
	Gate547NoPhysicalClaims      bool
	Gate547RedirectsToGate548    bool

	Verdict string
	Reason  string
}

type FrontierRow struct {
	Gate        int
	Block       string
	Native      string
	Bridge      string
	Environment string
	Firewall    string
}

type ClosureLedger struct {
	Executed bool
	Rows     []FrontierRow

	RowCount             int
	NativeRows           int
	BridgeRows           int
	EnvironmentalRows    int
	FailedRouteRows      int
	SchwingerBlockClosed bool
	AuthenticityClosed   bool
	ImportSwitchClosed   bool
	ComparatorClosed     bool
	ReleaseClosed        bool
	NativeFrontierFrozen bool
	BridgeFrontierMapped bool
	EnvironmentalMapped  bool
	Verdict              string
	Reason               string
}

type SectorGuard struct {
	Executed bool

	PhysicalSchwingerFunctionsImported bool
	AuthenticatedNonSyntheticSource    bool
	SourceAuthenticityAccepted         bool
	RealImportSwitchEnabled            bool
	OperatorIntentForRealImport        bool
	ComparatorAuthorized               bool
	ComparatorExecutedOnRealSource     bool
	ComparatorOutputQuarantined        bool
	ComparatorOutputReleased           bool
	BridgeEvidenceReleased             bool
	ReleaseReviewAccepted              bool
	NativeWriteLocked                  bool
	NativeWriteAuthorization           bool
	NativeRegistryWrite                bool
	ClosureOnly                        bool

	Verdict  string
	Reason   string
	Failures []string
}

type Firewall struct {
	Executed bool

	PhysicalSchwingerFunctionsLoaded bool
	PhysicalOSCertificateLoaded      bool
	PhysicalWickMapLoaded            bool
	PhysicalHilbertSpaceLoaded       bool
	PhysicalHamiltonianLoaded        bool
	UnitaryDynamicsLoaded            bool
	GlobalCausalityLoaded            bool
	TimeArrowLoaded                  bool
	ReleasedBridgeEvidence           bool
	NativeSchwingerFunctionWrite     bool
	NativeOSPositivityWrite          bool
	NativeWickWrite                  bool
	NativeHilbertWrite               bool
	NativeHamiltonianWrite           bool
	NativeUnitaryDynamicsWrite       bool
	NativeGlobalCausalWrite          bool
	NativeTimeArrowWrite             bool
	NativeRegistryWritten            bool

	Verdict string
	Reason  string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Closure     ClosureLedger
	Guard       SectorGuard
	Firewall    Firewall
	Registry    RegistryUpdate
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
	g547, err := generation2syntheticreleasereviewmanifestadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate547 synthetic release-review manifest adapter: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g547)}
	a.Closure = buildClosureLedger()
	a.Guard = buildGuard(a.Inheritance, a.Closure)
	a.Firewall = buildFirewall(a.Guard)
	a.Registry = buildRegistry()
	a.Next = buildNext()
	a.Truth = truth()
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2syntheticreleasereviewmanifestadapter.Analysis) Inheritance {
	noPhysical := !g.Firewall.PhysicalSchwingerFunctionsLoaded && !g.Firewall.OSPositivityCertificateLoaded && !g.Firewall.WickMapLoaded && !g.Firewall.HilbertSpaceReconstructed && !g.Firewall.HamiltonianSpectrumLoaded && !g.Firewall.UnitaryDynamicsLoaded && !g.Firewall.GlobalCausalityLoaded && !g.Firewall.TimeArrowLoaded
	return Inheritance{
		Executed:                     true,
		Gate547ManifestParsed:        g.Import.Loaded && g.Import.ManifestImported,
		Gate547RowsAccepted:          g.Import.AcceptedRows,
		Gate547ChecksumVerified:      g.Import.ChecksumVerified,
		Gate547HumanReviewParsed:     g.Review.HumanReviewMetadataParsed,
		Gate547ReproducibilityParsed: g.Review.ReproducibilityMetadataParsed,
		Gate547SourceChainParsed:     g.Review.SourceChainMetadataParsed,
		Gate547SyntheticBlocked:      g.Review.BlockedBecauseSynthetic && !g.Review.ReleaseAllowed && !g.Review.BridgeEvidenceReleased,
		Gate547NoBridgeEvidence:      !g.Firewall.ComparatorOutputReleased && !g.Firewall.BridgeEvidenceClaimReleased,
		Gate547NoRealSource:          !g.Firewall.RealSchwingerSourceImported && !g.Firewall.AuthenticatedRealSource,
		Gate547NativeWriteLocked:     g.Review.NativeWriteLocked && !g.Review.NativeWriteAuthorization && !g.Firewall.NativeRegistryWritten,
		Gate547NoPhysicalClaims:      noPhysical,
		Gate547RedirectsToGate548:    g.Next.Gate == 548,
		Verdict:                      StatusGate547SyntheticReleaseInherited,
		Reason:                       "Gate548 inherits Gate547's checksum-verified synthetic release review, parsed review/reproducibility/source-chain metadata, synthetic release block, no bridge evidence release, and native-write lock.",
	}
}

func buildClosureLedger() ClosureLedger {
	rows := []FrontierRow{
		{Gate: 536, Block: "physical Schwinger source schema", Native: "no", Bridge: "source-ledger airlock only", Environment: "actual S_n family and constructive measure", Firewall: "schema does not derive correlators"},
		{Gate: 537, Block: "synthetic Schwinger parser", Native: "no", Bridge: "synthetic plumbing accepted", Environment: "none imported", Firewall: "fake S_n cannot become physics"},
		{Gate: 538, Block: "source authenticity schema", Native: "no", Bridge: "provenance/authenticity sieve", Environment: "non-synthetic source identity", Firewall: "authenticity schema is not a source"},
		{Gate: 539, Block: "synthetic authenticity fixture", Native: "no", Bridge: "checksum/provenance parser", Environment: "none imported", Firewall: "synthetic fixture rejected as physical source"},
		{Gate: 540, Block: "real import switch", Native: "no", Bridge: "default-off switch", Environment: "operator intent and access grant", Firewall: "no import without switch and intent"},
		{Gate: 541, Block: "real-looking negative control", Native: "no", Bridge: "default-deny proof", Environment: "unauthenticated source rejected", Firewall: "real-looking is not real"},
		{Gate: 542, Block: "authorization manifest schema", Native: "no", Bridge: "14-row authorization airlock", Environment: "human/operator authorization", Firewall: "schema does not run comparator"},
		{Gate: 543, Block: "synthetic authorization manifest", Native: "no", Bridge: "quarantined dry-run authorization", Environment: "none imported", Firewall: "synthetic authorization cannot authorize real import"},
		{Gate: 544, Block: "comparator execution harness", Native: "no", Bridge: "16-row execution contract", Environment: "authenticated non-synthetic source", Firewall: "harness does not execute physics"},
		{Gate: 545, Block: "synthetic comparator output", Native: "no", Bridge: "quarantine result plumbing", Environment: "none imported", Firewall: "synthetic output remains quarantined"},
		{Gate: 546, Block: "release-review airlock", Native: "no", Bridge: "15-row release schema", Environment: "human review and reproducibility", Firewall: "release schema is not bridge evidence"},
		{Gate: 547, Block: "synthetic release-review manifest", Native: "no", Bridge: "release-review parser", Environment: "none imported", Firewall: "synthetic output cannot be released"},
	}
	c := ClosureLedger{Executed: true, Rows: rows, RowCount: len(rows), NativeFrontierFrozen: true, BridgeFrontierMapped: true, EnvironmentalMapped: true, SchwingerBlockClosed: true, AuthenticityClosed: true, ImportSwitchClosed: true, ComparatorClosed: true, ReleaseClosed: true, Verdict: StatusClosureLedgerEmitted, Reason: "Gates 536-547 are closed as a physical-correlation import/release frontier ledger; every stage remains bridge-only until authenticated non-synthetic sources, comparator execution, release review, and zero-native-write manifests exist."}
	for _, r := range rows {
		if r.Native != "" {
			c.NativeRows++
		}
		if r.Bridge != "" {
			c.BridgeRows++
		}
		if r.Environment != "" {
			c.EnvironmentalRows++
		}
		if r.Firewall != "" {
			c.FailedRouteRows++
		}
	}
	return c
}

func buildGuard(i Inheritance, c ClosureLedger) SectorGuard {
	g := SectorGuard{Executed: true,
		PhysicalSchwingerFunctionsImported: false,
		AuthenticatedNonSyntheticSource:    false,
		SourceAuthenticityAccepted:         false,
		RealImportSwitchEnabled:            false,
		OperatorIntentForRealImport:        false,
		ComparatorAuthorized:               false,
		ComparatorExecutedOnRealSource:     false,
		ComparatorOutputQuarantined:        false,
		ComparatorOutputReleased:           false,
		BridgeEvidenceReleased:             false,
		ReleaseReviewAccepted:              false,
		NativeWriteLocked:                  true,
		NativeWriteAuthorization:           false,
		NativeRegistryWrite:                false,
		ClosureOnly:                        true,
		Verdict:                            StatusNoBridgeEvidenceReleased,
		Reason:                             "The sector closes as a ledger only: no authenticated source, real import switch, comparator authorization, comparator run, release acceptance, bridge evidence release, or native registry write exists.",
	}
	if !i.Gate547SyntheticBlocked || !i.Gate547NoBridgeEvidence || !i.Gate547NoRealSource || !i.Gate547NativeWriteLocked || !i.Gate547NoPhysicalClaims {
		g.Failures = append(g.Failures, "Gate547 inheritance leaked")
	}
	if !c.NativeFrontierFrozen || !c.BridgeFrontierMapped || !c.EnvironmentalMapped || c.RowCount != 12 {
		g.Failures = append(g.Failures, "closure ledger incomplete")
	}
	if g.BridgeEvidenceReleased || g.NativeRegistryWrite || g.ComparatorExecutedOnRealSource || g.AuthenticatedNonSyntheticSource {
		g.Failures = append(g.Failures, "closure guard leaked physical evidence")
	}
	if len(g.Failures) > 0 {
		g.Verdict = strings.Join(g.Failures, ";")
		g.Reason = "Physical-correlation sector closure failed."
	}
	return g
}

func buildFirewall(g SectorGuard) Firewall {
	return Firewall{Executed: true,
		PhysicalSchwingerFunctionsLoaded: false,
		PhysicalOSCertificateLoaded:      false,
		PhysicalWickMapLoaded:            false,
		PhysicalHilbertSpaceLoaded:       false,
		PhysicalHamiltonianLoaded:        false,
		UnitaryDynamicsLoaded:            false,
		GlobalCausalityLoaded:            false,
		TimeArrowLoaded:                  false,
		ReleasedBridgeEvidence:           g.BridgeEvidenceReleased,
		NativeSchwingerFunctionWrite:     false,
		NativeOSPositivityWrite:          false,
		NativeWickWrite:                  false,
		NativeHilbertWrite:               false,
		NativeHamiltonianWrite:           false,
		NativeUnitaryDynamicsWrite:       false,
		NativeGlobalCausalWrite:          false,
		NativeTimeArrowWrite:             false,
		NativeRegistryWritten:            g.NativeRegistryWrite,
		Verdict:                          StatusFirewallPreserved,
		Reason:                           "The closure ledger freezes the physical-correlation pipeline without importing correlators, releasing bridge evidence, or writing native OS/Wick/Hilbert/Hamiltonian/dynamics facts.",
	}
}

func buildRegistry() RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Cℓ(1,7) finite algebra, projectors, anomaly/topology sockets, spectral-action shape, and synthetic-plumbing validators remain native/structural only where previously proved.",
			"No new native physical Schwinger, OS, Wick, Hilbert, Hamiltonian, unitary, global-causal, or arrow-of-time theorem is added at Gate548.",
		},
		BridgeEntries: []string{
			"Gates 536-547 are recorded as a complete bridge-only physical-correlation import/release pipeline: source schema, authenticity, import switch, authorization, comparator harness, quarantine output, and release review.",
			"Synthetic fixtures remain parser/plumbing evidence only; released bridge evidence requires authenticated non-synthetic sources and release review.",
		},
		EnvironmentalEntries: []string{
			"Actual S_n families, constructive measures, renormalization schemes, OS certificates, Wick/iε maps, Hamiltonian spectrum domains, reproducibility reports, and human release attestations remain environmental/source data.",
		},
		FailedRoutes: []string{
			StatusFailedClosureDoesNotDeriveSchwinger,
			StatusFailedClosureDoesNotProveOS,
			StatusFailedClosureDoesNotGrantWick,
			StatusFailedClosureDoesNotSelectHilbert,
			StatusFailedClosureDoesNotDeriveHamilton,
			StatusFailedClosureDoesNotGrantUnitary,
			StatusFailedClosureDoesNotGrantGlobal,
			StatusFailedClosureDoesNotSelectArrow,
			StatusFailedClosureNoRealSource,
			StatusFailedClosureNoRelease,
			StatusFirewallPreserved,
			StatusFirewallNativeWriteBlocked,
		},
		OpenTheorems: []string{
			"A future non-synthetic source can only proceed by importing a fully authenticated source chain and producing quarantined bridge evidence; native promotion remains forbidden unless a separate theorem proves it.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 549, Title: "Physical Correlation Evidence Board Airlock", Reason: "Gate548 closes the source/comparator/release pipeline. The next safe gate is an evidence-board schema for organizing future released bridge evidence without modifying native ASHA law.", PrimaryTask: "Define a zero-native-write evidence-board schema with citations, uncertainty, reproducibility, environmental classification, revocation hooks, and native-delta checks."}
}

func truth() string {
	return "Gate548 closes Gates 536-547 as a physical-correlation import/release frontier ledger: the full Schwinger source, authenticity, switch, authorization, comparator, quarantine, and release path is mapped, but no real source, bridge evidence release, or native physical dynamics theorem exists."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Gate547ManifestParsed || a.Inheritance.Gate547RowsAccepted != 15 || !a.Inheritance.Gate547ChecksumVerified || !a.Inheritance.Gate547HumanReviewParsed || !a.Inheritance.Gate547ReproducibilityParsed || !a.Inheritance.Gate547SourceChainParsed || !a.Inheritance.Gate547SyntheticBlocked || !a.Inheritance.Gate547NoBridgeEvidence || !a.Inheritance.Gate547NoRealSource || !a.Inheritance.Gate547NativeWriteLocked || !a.Inheritance.Gate547NoPhysicalClaims || !a.Inheritance.Gate547RedirectsToGate548 {
		bad = append(bad, "Gate547 inheritance incomplete")
	}
	if !a.Closure.Executed || a.Closure.RowCount != 12 || !a.Closure.NativeFrontierFrozen || !a.Closure.BridgeFrontierMapped || !a.Closure.EnvironmentalMapped || !a.Closure.SchwingerBlockClosed || !a.Closure.AuthenticityClosed || !a.Closure.ImportSwitchClosed || !a.Closure.ComparatorClosed || !a.Closure.ReleaseClosed {
		bad = append(bad, "closure ledger incomplete")
	}
	if a.Guard.PhysicalSchwingerFunctionsImported || a.Guard.AuthenticatedNonSyntheticSource || a.Guard.SourceAuthenticityAccepted || a.Guard.RealImportSwitchEnabled || a.Guard.ComparatorAuthorized || a.Guard.ComparatorExecutedOnRealSource || a.Guard.ComparatorOutputReleased || a.Guard.BridgeEvidenceReleased || a.Guard.ReleaseReviewAccepted || !a.Guard.NativeWriteLocked || a.Guard.NativeWriteAuthorization || a.Guard.NativeRegistryWrite || !a.Guard.ClosureOnly {
		bad = append(bad, "sector guard leaked")
	}
	if a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.PhysicalOSCertificateLoaded || a.Firewall.PhysicalWickMapLoaded || a.Firewall.PhysicalHilbertSpaceLoaded || a.Firewall.PhysicalHamiltonianLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.ReleasedBridgeEvidence || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate548 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate547SyntheticReleaseInherited,
		StatusClosureLedgerEmitted,
		StatusNativeFrontierFrozen,
		StatusBridgeFrontierMapped,
		StatusEnvironmentalFrontierMapped,
		StatusSchwingerSourceBlockClosed,
		StatusAuthenticityBlockClosed,
		StatusImportSwitchBlockClosed,
		StatusComparatorHarnessBlockClosed,
		StatusReleaseReviewBlockClosed,
		StatusNoBridgeEvidenceReleased,
		StatusClosureFirewallMatrixComplete,
		StatusFailedClosureDoesNotDeriveSchwinger,
		StatusFailedClosureDoesNotProveOS,
		StatusFailedClosureDoesNotGrantWick,
		StatusFailedClosureDoesNotSelectHilbert,
		StatusFailedClosureDoesNotDeriveHamilton,
		StatusFailedClosureDoesNotGrantUnitary,
		StatusFailedClosureDoesNotGrantGlobal,
		StatusFailedClosureDoesNotSelectArrow,
		StatusFailedClosureNoRealSource,
		StatusFailedClosureNoRelease,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("gate547_rows=%d checksum=%t human_review=%t reproducibility=%t source_chain=%t synthetic_blocked=%t bridge_evidence=%t real_source=%t native_locked=%t",
		i.Gate547RowsAccepted, i.Gate547ChecksumVerified, i.Gate547HumanReviewParsed, i.Gate547ReproducibilityParsed, i.Gate547SourceChainParsed, i.Gate547SyntheticBlocked, !i.Gate547NoBridgeEvidence, !i.Gate547NoRealSource, i.Gate547NativeWriteLocked)
}

func FormatClosure(c ClosureLedger) string {
	return fmt.Sprintf("rows=%d native_frozen=%t bridge_mapped=%t environmental_mapped=%t blocks=[schwinger:%t authenticity:%t switch:%t comparator:%t release:%t]",
		c.RowCount, c.NativeFrontierFrozen, c.BridgeFrontierMapped, c.EnvironmentalMapped, c.SchwingerBlockClosed, c.AuthenticityClosed, c.ImportSwitchClosed, c.ComparatorClosed, c.ReleaseClosed)
}

func FormatGuard(g SectorGuard) string {
	return fmt.Sprintf("real_source=%t source_auth=%t switch=%t comparator_auth=%t comparator_real=%t released=%t bridge_evidence=%t native_locked=%t native_write=%t closure_only=%t failures=%d",
		g.AuthenticatedNonSyntheticSource, g.SourceAuthenticityAccepted, g.RealImportSwitchEnabled, g.ComparatorAuthorized, g.ComparatorExecutedOnRealSource, g.ComparatorOutputReleased, g.BridgeEvidenceReleased, g.NativeWriteLocked, g.NativeRegistryWrite, g.ClosureOnly, len(g.Failures))
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("schwinger=%t os=%t wick=%t hilbert=%t hamiltonian=%t unitary=%t global=%t arrow=%t bridge_evidence=%t native_registry=%t",
		f.PhysicalSchwingerFunctionsLoaded, f.PhysicalOSCertificateLoaded, f.PhysicalWickMapLoaded, f.PhysicalHilbertSpaceLoaded, f.PhysicalHamiltonianLoaded, f.UnitaryDynamicsLoaded, f.GlobalCausalityLoaded, f.TimeArrowLoaded, f.ReleasedBridgeEvidence, f.NativeRegistryWritten)
}
