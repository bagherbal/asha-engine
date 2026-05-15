// Package generation2synthetic3plus1projectionadapter implements Gate 530:
// 3+1 Projection File Adapter and Clifford Compatibility Firewall.
//
// Gate 529 defined a fail-closed airlock for explicit 3+1 projector rows. This
// package performs the first file-backed dry-run through that airlock using a
// deliberately synthetic projector ledger. It verifies only bridge plumbing:
// projector idempotency, complementarity, 4+4 rank arithmetic, metric
// orthogonality, and the inherited Cℓ(1,7) 1+3 external signature assignment.
// Zero residuals are firewall-protected: they do not select physical spacetime,
// Wick rotation, a positive Hilbert space, unitary real-time dynamics, or a
// native gauge interpretation of the internal complement.
package generation2synthetic3plus1projectionadapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2projectionbridgeairlockpreflight"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID       = "GATE530-3PLUS1-PROJECTION-FILE-ADAPTER-CLIFFORD-COMPATIBILITY-FIREWALL"
	DefaultLedger = "data/synthetic_3plus1_projection_ledger_gate530.json"

	StatusGate529AirlockInherited              = "CONDITIONAL_SUPPORT_GATE529_PROJECTION_AIRLOCK_INHERITED"
	StatusSyntheticLedgerLoaded                = "CONDITIONAL_SUPPORT_SYNTHETIC_3PLUS1_PROJECTION_LEDGER_LOADED"
	StatusSyntheticRowAccepted                 = "CONDITIONAL_SUPPORT_GATE530_AIRLOCK_ACCEPTED_SYNTHETIC_PROJECTION_ROW"
	StatusAdapterExecuted                      = "CONDITIONAL_SUPPORT_SYNTHETIC_3PLUS1_ADAPTER_EXECUTED"
	StatusProjectorIdempotencyResidualsZero    = "CONDITIONAL_SUPPORT_PROJECTOR_IDEMPOTENCY_RESIDUALS_ZERO"
	StatusComplementOrthogonalityResidualZero  = "CONDITIONAL_SUPPORT_PROJECTOR_COMPLEMENT_ORTHOGONALITY_RESIDUAL_ZERO"
	StatusRank44SplitConfirmed                 = "CONDITIONAL_SUPPORT_RANK4_PLUS_RANK4_SPLIT_CONFIRMED"
	StatusCl17ExternalSignatureConfirmed       = "CONDITIONAL_SUPPORT_CL17_EXTERNAL_SIGNATURE_1PLUS3_CONFIRMED"
	StatusInternalSignatureReportedBridgeOnly  = "CONDITIONAL_SUPPORT_INTERNAL_COMPLEMENT_SIGNATURE_REPORTED_BRIDGE_ONLY"
	StatusNoObservedDimensionImportedByDefault = "CONDITIONAL_SUPPORT_NO_OBSERVED_DIMENSION_DATA_IMPORTED_BY_DEFAULT"

	StatusFailedFileMissing                  = "FAILED_ROUTE_GATE530_SYNTHETIC_PROJECTION_LEDGER_MISSING"
	StatusFailedMetadataIncomplete           = "FAILED_ROUTE_GATE530_PROJECTION_METADATA_INCOMPLETE"
	StatusFailedInvalidMatrixDomain          = "FAILED_ROUTE_GATE530_INVALID_PROJECTOR_MATRIX_DOMAIN"
	StatusFailedRankMismatch                 = "FAILED_ROUTE_GATE530_RANK4_PLUS_RANK4_SPLIT_NOT_CONFIRMED"
	StatusFailedProjectorResidualNonzero     = "FAILED_ROUTE_GATE530_PROJECTOR_IDEMPOTENCY_RESIDUAL_NONZERO"
	StatusFailedComplementResidualNonzero    = "FAILED_ROUTE_GATE530_PROJECTOR_COMPLEMENT_RESIDUAL_NONZERO"
	StatusFailedSignatureMismatch            = "FAILED_ROUTE_GATE530_EXTERNAL_SIGNATURE_NOT_1PLUS3"
	StatusFailedSyntheticProjectorNative     = "FAILED_ROUTE_SYNTHETIC_PROJECTOR_NATIVE_PROMOTION_REJECTED"
	StatusFailedZeroResidualsNotNative       = "FAILED_ROUTE_ZERO_RESIDUALS_DO_NOT_SELECT_PHYSICAL_SPACETIME"
	StatusFailedProjectorDoesNotGrantWick    = "FAILED_ROUTE_SYNTHETIC_3PLUS1_PROJECTOR_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedProjectorDoesNotGrantHilbert = "FAILED_ROUTE_SYNTHETIC_3PLUS1_PROJECTOR_DOES_NOT_GRANT_POSITIVE_HILBERT_SPACE"
	StatusFailedProjectorDoesNotGrantUnitary = "FAILED_ROUTE_SYNTHETIC_3PLUS1_PROJECTOR_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedInternalGaugeNativeRejected  = "FAILED_ROUTE_SYNTHETIC_INTERNAL_COMPLEMENT_NATIVE_GAUGE_IDENTIFICATION_REJECTED"
	StatusFirewallPreserved                  = "FIREWALL_PRESERVED_GATE530_SYNTHETIC_PROJECTION_ADAPTER_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked         = "FIREWALL_BLOCKED_GATE530_3PLUS1_PROJECTION_NATIVE_WRITE"
)

const tolerance = 1e-12

type Inheritance struct {
	Executed bool

	Gate529AirlockDefined             bool
	Gate529ProjectorSchemaReady       bool
	Gate529RequiresSourceConvention   bool
	Gate529RequiresBridgeOnly         bool
	Gate529RejectsNativePromotion     bool
	Gate529ComparatorExecutionBlocked bool
	Gate529WickHilbertUnitaryBlocked  bool
	Gate529InternalGaugeBlocked       bool
	Gate529NoObservedDimensionData    bool
	Gate529NativeRegistryBlocked      bool
	Gate530FileAdapterRedirect        bool

	Verdict, Reason string
}

type ProjectionRow struct {
	Name                         string      `json:"name"`
	ChosenProjectorMatrix        [][]float64 `json:"chosen_projector_matrix"`
	InternalComplementProjector  [][]float64 `json:"internal_complement_projector"`
	ExternalSignatureAssignment  string      `json:"external_signature_assignment"`
	InternalComplementAssignment string      `json:"internal_complement_assignment"`
	Source                       string      `json:"source"`
	SourceVersion                string      `json:"source_version"`
	Convention                   string      `json:"convention"`
	Uncertainty                  string      `json:"uncertainty"`
	BridgeOnly                   bool        `json:"bridge_only"`
	ComparatorOnly               bool        `json:"comparator_only"`
	NoTheoremInput               bool        `json:"no_theorem_input"`
	Synthetic                    bool        `json:"synthetic"`
	Observed                     bool        `json:"observed"`
	NativePromotion              bool        `json:"native_promotion"`
	NativeInputClaim             bool        `json:"native_input_claim"`
}

type ProjectionLedger struct {
	Gate                    int             `json:"gate"`
	LedgerName              string          `json:"ledger_name"`
	Description             string          `json:"description"`
	BridgeOnly              bool            `json:"bridge_only"`
	NativeRegistryWrite     bool            `json:"native_registry_write"`
	SyntheticFixture        bool            `json:"synthetic_fixture"`
	ObservedDimensionLoaded bool            `json:"observed_dimension_loaded"`
	Source                  string          `json:"source"`
	SourceVersion           string          `json:"source_version"`
	Convention              string          `json:"convention"`
	Rows                    []ProjectionRow `json:"rows"`
}

type FileImport struct {
	Executed                     bool
	Loaded                       bool
	Path                         string
	Rows                         int
	AcceptedRows                 int
	RejectedRows                 int
	BridgeOnlyLedger             bool
	SyntheticFixture             bool
	ObservedDimensionLoaded      bool
	NativeRegistryWriteRequested bool
	MetadataComplete             bool
	AllRowsBridgeOnly            bool
	AllRowsComparatorOnly        bool
	AllRowsNoTheoremInput        bool
	AllRowsSynthetic             bool
	AnyObservedClaim             bool
	NativePromotionRejected      bool
	DefaultFixtureObservedReject bool
	Verdict, Reason              string
	Failures                     []string
}

type AdapterInput struct {
	ProjectorRows                int
	P                            linear.Matrix
	Q                            linear.Matrix
	ExternalSignatureAssignment  string
	InternalComplementAssignment string
	BridgeOnly                   bool
	SyntheticFixture             bool
	ObservedDimensionLoaded      bool
	NativePromotion              bool
	MetadataComplete             bool
}

type AdapterOutput struct {
	Executed  bool
	Attempted bool
	Ready     bool

	Dimension            int
	BaseMetricPositive   int
	BaseMetricNegative   int
	ProjectorTrace       float64
	ComplementTrace      float64
	ProjectorRank        int
	ComplementRank       int
	ProjectorRankValid   bool
	ComplementRankValid  bool
	PIdempotencyResidual float64
	QIdempotencyResidual float64
	PQOrthogonalityNorm  float64
	QPOrthogonalityNorm  float64
	PPlusQIdentityNorm   float64
	MetricCrossResidual  float64
	ExternalPositive     int
	ExternalNegative     int
	ExternalNull         int
	InternalPositive     int
	InternalNegative     int
	InternalNull         int
	ExternalSignatureOK  bool
	InternalRankOK       bool
	AllResidualsZero     bool
	CliffordCompatible   bool
	BridgeOnly           bool
	NativePrediction     bool
	Verdict, Reason      string
	Failures             []string
}

type Firewall struct {
	Executed bool

	ObservedDimensionImported          bool
	SyntheticFixtureOnly               bool
	FileRowsNative                     bool
	AdapterOutputsNative               bool
	ProjectorNativePrediction          bool
	External3Plus1NativePrediction     bool
	InternalComplementNativePrediction bool
	WickRotationGranted                bool
	PositiveHilbertGranted             bool
	ReflectionPositivityGranted        bool
	PositiveEnergyGranted              bool
	UnitaryRealTimeGranted             bool
	GlobalHyperbolicityGranted         bool
	InternalGaugeNativeIdentification  bool
	ReopenedFlavorFirewall             bool
	ReopenedEWScaleFirewall            bool
	ReopenedGravityFirewall            bool
	ReopenedTopologyFirewall           bool
	NativeRegistryWritten              bool
	Verdict, Reason                    string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }
type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Import      FileImport
	Input       AdapterInput
	Output      AdapterOutput
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
	cache.Once.Do(func() { cache.a, cache.err = BuildFromFile(DefaultLedger) })
	return cache.a, cache.err
}

func Build() (Analysis, error) { return BuildFromFile(DefaultLedger) }

func BuildFromFile(path string) (Analysis, error) {
	a := Analysis{Inheritance: buildInheritance()}
	ledger, imp := loadLedger(path)
	a.Import = imp
	if imp.Loaded && imp.MetadataComplete && imp.AcceptedRows == 1 {
		in, err := buildInput(ledger, imp)
		if err != nil {
			a.Input = AdapterInput{ProjectorRows: len(ledger.Rows), MetadataComplete: imp.MetadataComplete, BridgeOnly: ledger.BridgeOnly, SyntheticFixture: ledger.SyntheticFixture, ObservedDimensionLoaded: ledger.ObservedDimensionLoaded, NativePromotion: ledger.NativeRegistryWrite}
			a.Output = AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedInvalidMatrixDomain, Reason: err.Error(), Failures: []string{StatusFailedInvalidMatrixDomain}}
		} else {
			a.Input = in
			a.Output = runAdapter(in)
		}
	} else if !imp.Loaded {
		a.Output = AdapterOutput{Executed: true, Attempted: false, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedFileMissing, Reason: "explicit Gate530 synthetic projection ledger was not found", Failures: []string{StatusFailedFileMissing}}
	} else {
		a.Output = AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedMetadataIncomplete, Reason: "Gate530 projection file did not satisfy the Gate529 airlock metadata domain", Failures: []string{StatusFailedMetadataIncomplete}}
	}
	a.Firewall = buildFirewall(a.Import, a.Output)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	g529, err := generation2projectionbridgeairlockpreflight.BuildDefault()
	if err != nil {
		return Inheritance{Executed: false, Verdict: StatusFailedFileMissing, Reason: fmt.Sprintf("could not inherit Gate529 projection airlock preflight: %v", err)}
	}
	return Inheritance{
		Executed:                          true,
		Gate529AirlockDefined:             g529.Schema.Executed && g529.Schema.ProjectorMatrixRequired && g529.Schema.ComplementMatrixRequired,
		Gate529ProjectorSchemaReady:       g529.Schema.ProjectorRankRequired == 4 && g529.Schema.ComplementRankRequired == 4 && g529.Schema.OrthogonalComplementCheck,
		Gate529RequiresSourceConvention:   g529.Schema.SourceRequired && g529.Schema.ConventionRequired,
		Gate529RequiresBridgeOnly:         g529.Schema.BridgeOnlyRequired,
		Gate529RejectsNativePromotion:     g529.Schema.NativePromotionRejected && g529.Rejection.NativeProjectorWriteRejected,
		Gate529ComparatorExecutionBlocked: !g529.Rejection.ComparatorExecutionPerformed,
		Gate529WickHilbertUnitaryBlocked:  !g529.Obligations.GrantsWickRotation && !g529.Obligations.GrantsPositiveHilbertProduct && !g529.Obligations.GrantsUnitaryRealTimeDynamics,
		Gate529InternalGaugeBlocked:       !g529.Obligations.GrantsInternalGaugeIdentification,
		Gate529NoObservedDimensionData:    !g529.Firewall.ObservedDimensionImported,
		Gate529NativeRegistryBlocked:      !g529.Firewall.NativeRegistryWritten,
		Gate530FileAdapterRedirect:        g529.Next.Gate == 530,
		Verdict:                           StatusGate529AirlockInherited,
		Reason:                            "Gate530 inherits Gate529's fail-closed 3+1 projection airlock: explicit projector/complement matrices may be checked only as sourced, synthetic, bridge-only data with native promotion rejected.",
	}
}

func loadLedger(path string) (ProjectionLedger, FileImport) {
	resolved := resolvePath(path)
	imp := FileImport{Executed: true, Path: resolved}
	b, err := os.ReadFile(resolved)
	if err != nil {
		imp.Verdict = StatusFailedFileMissing
		imp.Reason = err.Error()
		imp.Failures = []string{StatusFailedFileMissing}
		return ProjectionLedger{}, imp
	}
	var ledger ProjectionLedger
	if err := json.Unmarshal(b, &ledger); err != nil {
		imp.Loaded = true
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = err.Error()
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		return ledger, imp
	}
	imp.Loaded = true
	imp.Rows = len(ledger.Rows)
	imp.BridgeOnlyLedger = ledger.BridgeOnly
	imp.SyntheticFixture = ledger.SyntheticFixture
	imp.ObservedDimensionLoaded = ledger.ObservedDimensionLoaded
	imp.NativeRegistryWriteRequested = ledger.NativeRegistryWrite
	imp.NativePromotionRejected = !ledger.NativeRegistryWrite
	metadata := ledger.Gate == 530 && ledger.LedgerName != "" && ledger.Source != "" && ledger.SourceVersion != "" && ledger.Convention != "" && len(ledger.Rows) == 1
	allBridge := ledger.BridgeOnly
	allComparator := true
	allNoTheorem := true
	allSynthetic := ledger.SyntheticFixture
	anyObserved := ledger.ObservedDimensionLoaded
	accepted := 0
	failures := []string{}
	for _, r := range ledger.Rows {
		rowMeta := r.Name != "" && r.Source != "" && r.SourceVersion != "" && r.Convention != "" && r.Uncertainty != "" && r.ExternalSignatureAssignment != "" && r.InternalComplementAssignment != ""
		shapeOK := matrixShapeOK(r.ChosenProjectorMatrix, 8, 8) && matrixShapeOK(r.InternalComplementProjector, 8, 8)
		if !rowMeta {
			metadata = false
		}
		if !shapeOK {
			failures = append(failures, StatusFailedInvalidMatrixDomain)
		}
		if !r.BridgeOnly {
			allBridge = false
		}
		if !r.ComparatorOnly {
			allComparator = false
		}
		if !r.NoTheoremInput {
			allNoTheorem = false
		}
		if !r.Synthetic {
			allSynthetic = false
		}
		if r.Observed {
			anyObserved = true
		}
		if r.NativePromotion || r.NativeInputClaim {
			failures = append(failures, StatusFailedSyntheticProjectorNative)
		}
		if rowMeta && shapeOK && r.BridgeOnly && r.ComparatorOnly && r.NoTheoremInput && r.Synthetic && !r.Observed && !r.NativePromotion && !r.NativeInputClaim {
			accepted++
		}
	}
	if !metadata {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	if !ledger.BridgeOnly || ledger.NativeRegistryWrite {
		failures = append(failures, StatusFailedSyntheticProjectorNative)
	}
	if ledger.ObservedDimensionLoaded || anyObserved {
		failures = append(failures, StatusFailedZeroResidualsNotNative)
	}
	imp.AcceptedRows = accepted
	imp.RejectedRows = imp.Rows - accepted
	imp.MetadataComplete = metadata && accepted == 1 && len(failures) == 0
	imp.AllRowsBridgeOnly = allBridge
	imp.AllRowsComparatorOnly = allComparator
	imp.AllRowsNoTheoremInput = allNoTheorem
	imp.AllRowsSynthetic = allSynthetic
	imp.AnyObservedClaim = anyObserved
	imp.DefaultFixtureObservedReject = ledger.SyntheticFixture && !ledger.ObservedDimensionLoaded && !anyObserved
	imp.Failures = unique(failures)
	if len(imp.Failures) == 0 {
		imp.Verdict = strings.Join([]string{StatusSyntheticLedgerLoaded, StatusSyntheticRowAccepted, StatusNoObservedDimensionImportedByDefault}, ";")
		imp.Reason = "Gate530 loaded a deliberately synthetic, source-tagged 3+1 projector row through the Gate529 airlock with bridge_only=true, no_theorem_input=true, and native_promotion=false."
	} else {
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "Gate530 file import failed one or more Gate529 projection-airlock requirements."
	}
	return ledger, imp
}

func buildInput(ledger ProjectionLedger, imp FileImport) (AdapterInput, error) {
	row := ledger.Rows[0]
	p, err := linear.FromRows(row.ChosenProjectorMatrix)
	if err != nil {
		return AdapterInput{}, err
	}
	q, err := linear.FromRows(row.InternalComplementProjector)
	if err != nil {
		return AdapterInput{}, err
	}
	return AdapterInput{ProjectorRows: len(ledger.Rows), P: p, Q: q, ExternalSignatureAssignment: row.ExternalSignatureAssignment, InternalComplementAssignment: row.InternalComplementAssignment, BridgeOnly: ledger.BridgeOnly && row.BridgeOnly, SyntheticFixture: ledger.SyntheticFixture && row.Synthetic, ObservedDimensionLoaded: ledger.ObservedDimensionLoaded || row.Observed, NativePromotion: ledger.NativeRegistryWrite || row.NativePromotion || row.NativeInputClaim, MetadataComplete: imp.MetadataComplete}, nil
}

func runAdapter(in AdapterInput) AdapterOutput {
	out := AdapterOutput{Executed: true, Attempted: true, Dimension: 8, BaseMetricPositive: 1, BaseMetricNegative: 7, BridgeOnly: true, NativePrediction: false}
	ready := in.MetadataComplete && in.BridgeOnly && in.SyntheticFixture && !in.ObservedDimensionLoaded && !in.NativePromotion && in.P.Rows() == 8 && in.P.Cols() == 8 && in.Q.Rows() == 8 && in.Q.Cols() == 8
	out.Ready = ready
	if !ready {
		out.Verdict = StatusFailedMetadataIncomplete
		out.Reason = "Synthetic 3+1 adapter input did not meet the complete bridge-only domain."
		out.Failures = []string{StatusFailedMetadataIncomplete}
		return out
	}
	pProj, _ := linear.NewProjector("synthetic_external_3plus1", in.P)
	qProj, _ := linear.NewProjector("synthetic_internal_complement", in.Q)
	out.PIdempotencyResidual, _ = pProj.IdempotenceResidual()
	out.QIdempotencyResidual, _ = qProj.IdempotenceResidual()
	out.ProjectorTrace, _ = in.P.Trace()
	out.ComplementTrace, _ = in.Q.Trace()
	out.ProjectorRank = int(math.Round(out.ProjectorTrace))
	out.ComplementRank = int(math.Round(out.ComplementTrace))
	out.ProjectorRankValid = nearly(out.ProjectorTrace, 4, tolerance) && out.ProjectorRank == 4
	out.ComplementRankValid = nearly(out.ComplementTrace, 4, tolerance) && out.ComplementRank == 4
	pq, _ := in.P.Mul(in.Q)
	qp, _ := in.Q.Mul(in.P)
	out.PQOrthogonalityNorm = pq.FrobeniusNorm()
	out.QPOrthogonalityNorm = qp.FrobeniusNorm()
	pPlusQ, _ := in.P.Add(in.Q)
	identityResidual, _ := pPlusQ.Sub(linear.Identity(8))
	out.PPlusQIdentityNorm = identityResidual.FrobeniusNorm()
	g := cliffordMetric17()
	ptg, _ := in.P.Transpose().Mul(g)
	ptgq, _ := ptg.Mul(in.Q)
	qtg, _ := in.Q.Transpose().Mul(g)
	qtgp, _ := qtg.Mul(in.P)
	out.MetricCrossResidual = math.Max(ptgq.FrobeniusNorm(), qtgp.FrobeniusNorm())
	pgp, _ := ptg.Mul(in.P)
	qgq, _ := qtg.Mul(in.Q)
	out.ExternalPositive, out.ExternalNegative, out.ExternalNull = diagonalSignature(pgp, tolerance)
	out.InternalPositive, out.InternalNegative, out.InternalNull = diagonalSignature(qgq, tolerance)
	out.ExternalSignatureOK = in.ExternalSignatureAssignment == "1+3" && out.ExternalPositive == 1 && out.ExternalNegative == 3
	out.InternalRankOK = out.InternalPositive+out.InternalNegative == 4 && out.InternalNull == 4
	out.AllResidualsZero = nearly(out.PIdempotencyResidual, 0, tolerance) && nearly(out.QIdempotencyResidual, 0, tolerance) && nearly(out.PQOrthogonalityNorm, 0, tolerance) && nearly(out.QPOrthogonalityNorm, 0, tolerance) && nearly(out.PPlusQIdentityNorm, 0, tolerance) && nearly(out.MetricCrossResidual, 0, tolerance)
	out.CliffordCompatible = out.AllResidualsZero && out.ProjectorRankValid && out.ComplementRankValid && out.ExternalSignatureOK && out.InternalRankOK
	failures := []string{}
	if !out.ProjectorRankValid || !out.ComplementRankValid {
		failures = append(failures, StatusFailedRankMismatch)
	}
	if !nearly(out.PIdempotencyResidual, 0, tolerance) || !nearly(out.QIdempotencyResidual, 0, tolerance) {
		failures = append(failures, StatusFailedProjectorResidualNonzero)
	}
	if !nearly(out.PQOrthogonalityNorm, 0, tolerance) || !nearly(out.QPOrthogonalityNorm, 0, tolerance) || !nearly(out.PPlusQIdentityNorm, 0, tolerance) || !nearly(out.MetricCrossResidual, 0, tolerance) {
		failures = append(failures, StatusFailedComplementResidualNonzero)
	}
	if !out.ExternalSignatureOK {
		failures = append(failures, StatusFailedSignatureMismatch)
	}
	out.Failures = unique(failures)
	if len(out.Failures) == 0 {
		out.Verdict = strings.Join([]string{StatusAdapterExecuted, StatusProjectorIdempotencyResidualsZero, StatusComplementOrthogonalityResidualZero, StatusRank44SplitConfirmed, StatusCl17ExternalSignatureConfirmed, StatusInternalSignatureReportedBridgeOnly}, ";")
		out.Reason = "The synthetic diagonal projector passes bridge-only algebra: P^2=P, Q^2=Q, PQ=QP=0, P+Q=I, rank(P)=rank(Q)=4, PᵀGQ=0, and the external image has inherited Cℓ(1,7) signature 1+3."
	} else {
		out.Verdict = strings.Join(out.Failures, ";")
		out.Reason = "Gate530 adapter found nonzero residuals or a signature/rank mismatch in the supplied projection ledger."
	}
	return out
}

func buildFirewall(imp FileImport, out AdapterOutput) Firewall {
	return Firewall{
		Executed:                           true,
		ObservedDimensionImported:          imp.ObservedDimensionLoaded || imp.AnyObservedClaim,
		SyntheticFixtureOnly:               imp.SyntheticFixture && !imp.ObservedDimensionLoaded && !imp.AnyObservedClaim,
		FileRowsNative:                     false,
		AdapterOutputsNative:               out.NativePrediction,
		ProjectorNativePrediction:          false,
		External3Plus1NativePrediction:     false,
		InternalComplementNativePrediction: false,
		WickRotationGranted:                false,
		PositiveHilbertGranted:             false,
		ReflectionPositivityGranted:        false,
		PositiveEnergyGranted:              false,
		UnitaryRealTimeGranted:             false,
		GlobalHyperbolicityGranted:         false,
		InternalGaugeNativeIdentification:  false,
		ReopenedFlavorFirewall:             false,
		ReopenedEWScaleFirewall:            false,
		ReopenedGravityFirewall:            false,
		ReopenedTopologyFirewall:           false,
		NativeRegistryWritten:              imp.NativeRegistryWriteRequested,
		Verdict:                            strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSyntheticProjectorNative, StatusFailedZeroResidualsNotNative, StatusFailedProjectorDoesNotGrantWick, StatusFailedProjectorDoesNotGrantHilbert, StatusFailedProjectorDoesNotGrantUnitary, StatusFailedInternalGaugeNativeRejected}, ";"),
		Reason:                             "Gate530 residuals are bridge-only plumbing checks. The synthetic projector row and zero residuals do not become a native 3+1 spacetime theorem and do not grant Wick rotation, Hilbert reconstruction, unitary dynamics, global hyperbolicity, or native internal gauge identification.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native physical 3+1 spacetime projector is written at Gate530.",
			"Cℓ(1,7) retains only the native 1+7 causal/null-cone socket; the rank-four vector projector remains chosen bridge data.",
			"Wick rotation, positive Hilbert reconstruction, time orientation, positive energy, unitary dynamics, and global hyperbolicity remain independent unsolved obligations.",
		},
		BridgeEntries: []string{
			"File-backed synthetic 3+1 projection adapter implemented against the Gate529 airlock.",
			"Bridge residuals verify P²−P=0, Q²−Q=0, PQ=QP=0, P+Q=I, rank(P)=rank(Q)=4, and PᵀGQ=0 for the default synthetic fixture.",
			"The selected external image carries inherited Cℓ(1,7) signature 1+3 under the checked convention; the four-dimensional complement is reported only as bridge metadata.",
		},
		EnvironmentalEntries: []string{
			"The actual physical four-plane, arrow of time, Wick dictionary, Hilbert product, and internal gauge/geometric interpretation remain environmental or bridge choices.",
			"A future observed dimensional/topological model may be compared only through source-tagged bridge ledgers and may not be promoted by residual closure alone.",
		},
		FailedRoutes: []string{
			StatusFailedSyntheticProjectorNative,
			StatusFailedZeroResidualsNotNative,
			StatusFailedProjectorDoesNotGrantWick,
			StatusFailedProjectorDoesNotGrantHilbert,
			StatusFailedProjectorDoesNotGrantUnitary,
			StatusFailedInternalGaugeNativeRejected,
		},
		OpenTheorems: []string{
			"derive, or continue to quarantine, a native Spin(1,7)-breaking selector for the physical rank-four vector projector",
			"construct a separate Wick/reflection-positivity/Hilbert reconstruction airlock rather than smuggling those structures through dimensional projection",
			"audit whether a source-tagged internal complement can interface with gauge geometry without claiming native complement-to-gauge identity",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 531, Title: "Wick/Hilbert Fundamental-Symmetry Airlock Preflight", Reason: "Gate530 validates dimensional socket plumbing but leaves the separate Lorentzian obligations untouched. The next safe move is to define the fail-closed schema for a fundamental symmetry or Wick/Hilbert reconstruction input.", PrimaryTask: "Define the bridge-only metadata requirements for importing a Krein-to-Hilbert fundamental symmetry, reflection-positivity/Wick data, and time-orientation convention without promoting them to native ASHA theorems."}
}

func truth(a Analysis) string {
	return "Gate530 proves that the Gate529 dimensional socket can safely house a synthetic 3+1 projector as bridge plumbing: idempotency, complementarity, rank 4+4 arithmetic, metric orthogonality, and external 1+3 signature all check out for the default fixture. It still does not prove that ASHA selected physical spacetime. Zero residuals validate the adapter, not Wick rotation, Hilbert positivity, time orientation, unitary dynamics, global hyperbolicity, or internal gauge identification."
}

func validate(a Analysis) error {
	problems := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate529AirlockDefined || !a.Inheritance.Gate529ProjectorSchemaReady || !a.Inheritance.Gate529RequiresSourceConvention || !a.Inheritance.Gate529RequiresBridgeOnly || !a.Inheritance.Gate529RejectsNativePromotion || !a.Inheritance.Gate529ComparatorExecutionBlocked || !a.Inheritance.Gate529WickHilbertUnitaryBlocked || !a.Inheritance.Gate529InternalGaugeBlocked || !a.Inheritance.Gate529NoObservedDimensionData || !a.Inheritance.Gate529NativeRegistryBlocked || !a.Inheritance.Gate530FileAdapterRedirect {
		problems = append(problems, "bad Gate529 inheritance")
	}
	if !a.Import.Loaded || a.Import.Rows != 1 || a.Import.AcceptedRows != 1 || a.Import.RejectedRows != 0 || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedDimensionLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSynthetic || a.Import.AnyObservedClaim {
		problems = append(problems, "bad file import")
	}
	if !a.Output.Executed || !a.Output.Attempted || !a.Output.Ready || !a.Output.BridgeOnly || a.Output.NativePrediction || !a.Output.ProjectorRankValid || !a.Output.ComplementRankValid || a.Output.ProjectorRank != 4 || a.Output.ComplementRank != 4 || !a.Output.AllResidualsZero || !a.Output.ExternalSignatureOK || !a.Output.InternalRankOK || !a.Output.CliffordCompatible || !nearly(a.Output.PIdempotencyResidual, 0, tolerance) || !nearly(a.Output.QIdempotencyResidual, 0, tolerance) || !nearly(a.Output.PQOrthogonalityNorm, 0, tolerance) || !nearly(a.Output.QPOrthogonalityNorm, 0, tolerance) || !nearly(a.Output.PPlusQIdentityNorm, 0, tolerance) || !nearly(a.Output.MetricCrossResidual, 0, tolerance) {
		problems = append(problems, "bad adapter output")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedDimensionImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.ProjectorNativePrediction || a.Firewall.External3Plus1NativePrediction || a.Firewall.InternalComplementNativePrediction || a.Firewall.WickRotationGranted || a.Firewall.PositiveHilbertGranted || a.Firewall.ReflectionPositivityGranted || a.Firewall.PositiveEnergyGranted || a.Firewall.UnitaryRealTimeGranted || a.Firewall.GlobalHyperbolicityGranted || a.Firewall.InternalGaugeNativeIdentification || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		problems = append(problems, "firewall violation")
	}
	if len(problems) > 0 {
		return fmt.Errorf(strings.Join(problems, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: airlock=%t schema_ready=%t source_convention=%t bridge_only=%t native_promotion_rejected=%t comparator_blocked=%t wick_hilbert_unitary_blocked=%t internal_gauge_blocked=%t no_observed_dimension=%t native_blocked=%t gate530_redirect=%t; %s", x.Verdict, x.Gate529AirlockDefined, x.Gate529ProjectorSchemaReady, x.Gate529RequiresSourceConvention, x.Gate529RequiresBridgeOnly, x.Gate529RejectsNativePromotion, x.Gate529ComparatorExecutionBlocked, x.Gate529WickHilbertUnitaryBlocked, x.Gate529InternalGaugeBlocked, x.Gate529NoObservedDimensionData, x.Gate529NativeRegistryBlocked, x.Gate530FileAdapterRedirect, x.Reason)
}

func FormatImport(x FileImport) string {
	return fmt.Sprintf("%s: loaded=%t rows=%d accepted=%d rejected=%d bridge_only=%t synthetic_fixture=%t observed_dimension_loaded=%t native_write_requested=%t metadata_complete=%t all_bridge_only=%t all_comparator_only=%t no_theorem_input=%t all_synthetic=%t observed_claim=%t native_promotion_rejected=%t; %s", x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, x.BridgeOnlyLedger, x.SyntheticFixture, x.ObservedDimensionLoaded, x.NativeRegistryWriteRequested, x.MetadataComplete, x.AllRowsBridgeOnly, x.AllRowsComparatorOnly, x.AllRowsNoTheoremInput, x.AllRowsSynthetic, x.AnyObservedClaim, x.NativePromotionRejected, x.Reason)
}

func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("%s: ready=%t dim=%d base_signature=%d+%d rankP=%d traceP=%.12g rankQ=%d traceQ=%.12g P2_minus_P=%.12g Q2_minus_Q=%.12g PQ=%.12g QP=%.12g PplusQ_minus_I=%.12g PTGQ=%.12g external_signature=%d+%d external_null=%d internal_signature=%d+%d internal_null=%d all_residuals_zero=%t Clifford_compatible=%t bridge_only=%t native_prediction=%t; %s", x.Verdict, x.Ready, x.Dimension, x.BaseMetricPositive, x.BaseMetricNegative, x.ProjectorRank, x.ProjectorTrace, x.ComplementRank, x.ComplementTrace, x.PIdempotencyResidual, x.QIdempotencyResidual, x.PQOrthogonalityNorm, x.QPOrthogonalityNorm, x.PPlusQIdentityNorm, x.MetricCrossResidual, x.ExternalPositive, x.ExternalNegative, x.ExternalNull, x.InternalPositive, x.InternalNegative, x.InternalNull, x.AllResidualsZero, x.CliffordCompatible, x.BridgeOnly, x.NativePrediction, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_dimension=%t synthetic_only=%t file_rows_native=%t outputs_native=%t projector_native=%t external_3plus1_native=%t internal_native=%t Wick=%t Hilbert=%t reflection=%t positive_energy=%t unitary=%t global_hyperbolicity=%t internal_gauge_native=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_registry_written=%t; %s", x.Verdict, x.ObservedDimensionImported, x.SyntheticFixtureOnly, x.FileRowsNative, x.AdapterOutputsNative, x.ProjectorNativePrediction, x.External3Plus1NativePrediction, x.InternalComplementNativePrediction, x.WickRotationGranted, x.PositiveHilbertGranted, x.ReflectionPositivityGranted, x.PositiveEnergyGranted, x.UnitaryRealTimeGranted, x.GlobalHyperbolicityGranted, x.InternalGaugeNativeIdentification, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 530 Registry Audit — 3+1 Projection File Adapter and Clifford Compatibility Firewall\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 530 inherits Gate 529's dimensional-projection airlock. A projector row may be loaded only as synthetic bridge data with source, convention, bridge_only=true, native_promotion=false, and no_theorem_input=true.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Synthetic ledger import\n\n" + a.Import.Reason + "\n\n```text\n" + FormatImport(a.Import) + "\n```\n\n")
	b.WriteString("## Projector and complement residuals\n\nThe adapter evaluates the bridge-only conditions `P^2-P=0`, `Q^2-Q=0`, `PQ=QP=0`, and `P+Q=I`, then checks metric orthogonality against the inherited Cℓ(1,7) quadratic form.\n\n```text\n" + FormatOutput(a.Output) + "\n```\n\n")
	b.WriteString("## Clifford metric compatibility\n\nThe default synthetic fixture uses the Cℓ(1,7) metric convention `G=diag(+1,-1,-1,-1,-1,-1,-1,-1)`. The external projector image has signature `1+3`; the complement is four-dimensional and reported bridge-only.\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "### Native", a.Registry.NativeEntries)
	writeList(&b, "### Bridge", a.Registry.BridgeEntries)
	writeList(&b, "### Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "### Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "### Open theorems", a.Registry.OpenTheorems)
	b.WriteString(fmt.Sprintf("## Next step\n\nGate %d — %s. %s\n\nPrimary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func statuses() []string {
	return []string{
		StatusGate529AirlockInherited,
		StatusSyntheticLedgerLoaded,
		StatusSyntheticRowAccepted,
		StatusAdapterExecuted,
		StatusProjectorIdempotencyResidualsZero,
		StatusComplementOrthogonalityResidualZero,
		StatusRank44SplitConfirmed,
		StatusCl17ExternalSignatureConfirmed,
		StatusInternalSignatureReportedBridgeOnly,
		StatusNoObservedDimensionImportedByDefault,
		StatusFailedSyntheticProjectorNative,
		StatusFailedZeroResidualsNotNative,
		StatusFailedProjectorDoesNotGrantWick,
		StatusFailedProjectorDoesNotGrantHilbert,
		StatusFailedProjectorDoesNotGrantUnitary,
		StatusFailedInternalGaugeNativeRejected,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func cliffordMetric17() linear.Matrix {
	return linear.Diagonal([]float64{1, -1, -1, -1, -1, -1, -1, -1})
}

func diagonalSignature(m linear.Matrix, eps float64) (positive, negative, zero int) {
	for i := 0; i < m.Rows() && i < m.Cols(); i++ {
		v := m.At(i, i)
		switch {
		case v > eps:
			positive++
		case v < -eps:
			negative++
		default:
			zero++
		}
	}
	return positive, negative, zero
}

func matrixShapeOK(x [][]float64, rows, cols int) bool {
	if len(x) != rows {
		return false
	}
	for _, r := range x {
		if len(r) != cols {
			return false
		}
		for _, v := range r {
			if math.IsNaN(v) || math.IsInf(v, 0) {
				return false
			}
		}
	}
	return true
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString(title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}

func unique(xs []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, x := range xs {
		if !seen[x] {
			seen[x] = true
			out = append(out, x)
		}
	}
	return out
}

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

func resolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return path
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "../../.."))
	return filepath.Join(root, path)
}
