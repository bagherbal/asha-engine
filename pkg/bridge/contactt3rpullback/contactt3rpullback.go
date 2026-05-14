// Package contactt3rpullback implements Gate 137: contact T3R pullback
// obstruction / Fock-to-contact intertwiner search.
//
// Gate 136 established that the matter/Fock sector has a genuine temporal
// T3R candidate family, chiral restrictions, and a hyperaudit diagnostic, but
// no map transporting those operators onto the seven contact partial-overlap
// rows. Gate 137 searches for that missing map explicitly.
//
// The result is a disciplined obstruction. Generic linear maps from the
// 16-dimensional Fock carrier, or the 64-dimensional matter-scalar tensor
// carrier, to the seven contact rows exist. But none is canonical, none is
// shown to intertwine T3R/chirality/B-L/SU2L, and the leptoquark six-block still
// carries an S6 assignment ambiguity. Therefore matter-side T3R cannot yet be
// pulled back to contact hypercharge rows.
package contactt3rpullback

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactlqt3r"
)

type IntertwinerKind string

const (
	FockToContactLinearMap     IntertwinerKind = "fock-to-contact-linear-map"
	TensorToContactLinearMap   IntertwinerKind = "tensor-to-contact-linear-map"
	LeptoquarkSixBlockMap      IntertwinerKind = "leptoquark-six-block-map"
	CurrentOnePlusSixBranchMap IntertwinerKind = "current-1-plus-6-branch-map"
	ScalarToContactMap         IntertwinerKind = "scalar-to-contact-map"
	ContactSpectralIdentityMap IntertwinerKind = "contact-spectral-identity-map"
)

type Candidate struct {
	Name                         string
	Kind                         IntertwinerKind
	Domain                       string
	Codomain                     string
	DomainDim                    int
	CodomainDim                  int
	RankTarget                   int
	KernelDimIfSurjective        int
	GenericMapsExist             bool
	Canonical                    bool
	ContactSideOnly              bool
	MatterSideOnly               bool
	CurrentSideOnly              bool
	FockToContactIntertwiner     bool
	IntertwinesT3R               bool
	IntertwinesChirality         bool
	IntertwinesBMinusL           bool
	IntertwinesSU2L              bool
	PreservesContactRows         bool
	RepresentationRowDerived     bool
	RequiresContinuousChoice     bool
	RequiresS6Choice             bool
	RequiresOnePlusSixRefinement bool
	HiddenDiscreteChoices        int
	Obstruction                  string
}

type PullbackRow struct {
	Name                     string
	ContactRow               string
	MatterT3RCandidateValues []float64
	T3RPullbackDerived       bool
	ChiralityPullbackDerived bool
	BMinusLPullbackDerived   bool
	SU2LPullbackDerived      bool
	HyperchargeRowDerived    bool
	LocalFieldDerived        bool
	RepresentationComplete   bool
	BetaPermitted            bool
	RequiresFockContactMap   bool
	RequiresS6Choice         bool
	Obstruction              string
}

type Summary struct {
	MatterDimension                   int
	ScalarDimension                   int
	TensorDimension                   int
	ContactRows                       int
	LeptoquarkRows                    int
	MatterT3ROperatorFound            bool
	MatterChiralRestrictionsAvailable bool
	MatterMirrorAmbiguity             bool
	FockToContactGenericKernelDim     int
	TensorToContactGenericKernelDim   int
	GenericFockToContactMapsExist     bool
	CanonicalFockToContactMaps        int
	FockToContactIntertwinersDerived  int
	T3RPullbackRowsDerived            int
	ChiralityPullbackRowsDerived      int
	BMinusLPullbackRowsDerived        int
	SU2LPullbackRowsDerived           int
	HyperchargeRowsDerived            int
	RepresentationCompleteRows        int
	ContactBetaRowsAllowed            int
	ContactZeroRowsProved             int
	ResidualS6Choices                 int
	ResidualNullityBefore             int
	ResidualNullityAfter              int
}

type Analysis struct {
	Previous   contactlqt3r.Analysis
	Candidates []Candidate
	Rows       []PullbackRow
	Summary    Summary

	MatterDimension                  int
	ScalarDimension                  int
	TensorDimension                  int
	ContactRows                      int
	LeptoquarkRows                   int
	MatterT3ROperatorFound           bool
	MatterChiralRestricted           bool
	MatterMirrorAmbiguity            bool
	GenericFockToContactMapsExist    bool
	GenericTensorToContactMapsExist  bool
	CanonicalFockToContactMaps       int
	FockToContactIntertwinersDerived int
	T3RPullbackRowsDerived           int
	ChiralityPullbackRowsDerived     int
	BMinusLPullbackRowsDerived       int
	SU2LPullbackRowsDerived          int
	HyperchargeRowsDerived           int
	ElectricChargeRowsDerived        int
	RepresentationCompleteRows       int
	RepresentationOpenRows           int
	ContactBetaRowsAllowed           int
	ContactZeroRowsProved            int
	FullBetaMatchingTensorDerived    bool
	ThresholdCorrectedBetaDerived    bool
	BetaPermissionFirewallClosed     bool

	ResidualS6Choices        int
	ResidualNullityBefore    int
	ResidualNullityAfter     int
	HiddenObservedInputUsed  bool
	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	PhysicalScaleDerived     bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := contactlqt3r.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactlqt3r.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.LeptoquarkRows != 6 || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 137 requires Gate 136 contact T3R firewall to be closed")
	}
	if !prev.MatterT3ROperatorFound || !prev.MatterChiralRestricted || !prev.MatterMirrorAmbiguity || prev.MatterFullSMTableDerived {
		return Analysis{}, fmt.Errorf("Gate 137 requires matter-side T3R diagnostics to remain bridge-open")
	}
	if prev.ContactPullbackRowsDerived != 0 || prev.ContactT3RRowsDerived != 0 || prev.ContactChiralityRowsDerived != 0 || prev.SignedBLRowsDerived != 0 || prev.WeakSU2RowsDerived != 0 || prev.HyperchargeRowsDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 137 requires Gate 136 to derive no contact pullback rows")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 137 refuses hidden physical input from Gate 136")
	}

	matterDim := prev.T3R.MatterDimension
	scalarDim := prev.T3R.ScalarDimension
	tensorDim := prev.T3R.TensorDimension
	if matterDim != 16 || scalarDim != 4 || tensorDim != 64 {
		return Analysis{}, fmt.Errorf("unexpected matter/scalar/tensor dimensions: %d/%d/%d", matterDim, scalarDim, tensorDim)
	}
	if math.Abs(prev.T3R.TemporalTrace) > 1e-10 || math.Abs(prev.T3R.TemporalTraceSquared-4.0) > 1e-10 {
		return Analysis{}, fmt.Errorf("unexpected matter temporal T3R trace data")
	}

	contactRows := 7
	lqRows := 6
	fockKernel := matterDim - contactRows
	tensorKernel := tensorDim - contactRows

	candidates := buildCandidates(matterDim, scalarDim, tensorDim, contactRows, lqRows, prev.ResidualS6Choices)
	rows := buildRows(prev.Rows)

	canonicalFockMaps := count(candidates, func(c Candidate) bool { return c.Kind == FockToContactLinearMap && c.Canonical })
	fockIntertwiners := count(candidates, func(c Candidate) bool { return c.FockToContactIntertwiner })
	t3rPullbacks := count(rows, func(r PullbackRow) bool { return r.T3RPullbackDerived })
	chirPullbacks := count(rows, func(r PullbackRow) bool { return r.ChiralityPullbackDerived })
	blPullbacks := count(rows, func(r PullbackRow) bool { return r.BMinusLPullbackDerived })
	su2Pullbacks := count(rows, func(r PullbackRow) bool { return r.SU2LPullbackDerived })
	yRows := count(rows, func(r PullbackRow) bool { return r.HyperchargeRowDerived })
	repRows := count(rows, func(r PullbackRow) bool { return r.RepresentationComplete })
	betaRows := count(rows, func(r PullbackRow) bool { return r.BetaPermitted })

	summary := Summary{
		MatterDimension:                   matterDim,
		ScalarDimension:                   scalarDim,
		TensorDimension:                   tensorDim,
		ContactRows:                       contactRows,
		LeptoquarkRows:                    lqRows,
		MatterT3ROperatorFound:            prev.MatterT3ROperatorFound,
		MatterChiralRestrictionsAvailable: prev.MatterChiralRestricted,
		MatterMirrorAmbiguity:             prev.MatterMirrorAmbiguity,
		FockToContactGenericKernelDim:     fockKernel,
		TensorToContactGenericKernelDim:   tensorKernel,
		GenericFockToContactMapsExist:     true,
		CanonicalFockToContactMaps:        canonicalFockMaps,
		FockToContactIntertwinersDerived:  fockIntertwiners,
		T3RPullbackRowsDerived:            t3rPullbacks,
		ChiralityPullbackRowsDerived:      chirPullbacks,
		BMinusLPullbackRowsDerived:        blPullbacks,
		SU2LPullbackRowsDerived:           su2Pullbacks,
		HyperchargeRowsDerived:            yRows,
		RepresentationCompleteRows:        repRows,
		ContactBetaRowsAllowed:            betaRows,
		ContactZeroRowsProved:             0,
		ResidualS6Choices:                 prev.ResidualS6Choices,
		ResidualNullityBefore:             prev.ResidualNullityAfter,
		ResidualNullityAfter:              prev.ResidualNullityAfter,
	}

	truth := "Gate 137 searches for the missing Fock-to-contact intertwiner that would pull matter-side T3R/chirality onto the seven contact rows. Generic linear maps H_Fock→R^7 and (H_Fock⊗H_phi)→R^7 exist, but they require arbitrary kernels and do not intertwine T3R, chirality, B-L, SU(2)L, or contact row semantics. The leptoquark six-block route remains an S6 assignment problem. Therefore matter-side T3R cannot yet become contact hypercharge, and the contact beta firewall remains closed."

	return Analysis{
		Previous:   prev,
		Candidates: candidates,
		Rows:       rows,
		Summary:    summary,

		MatterDimension:                  matterDim,
		ScalarDimension:                  scalarDim,
		TensorDimension:                  tensorDim,
		ContactRows:                      contactRows,
		LeptoquarkRows:                   lqRows,
		MatterT3ROperatorFound:           prev.MatterT3ROperatorFound,
		MatterChiralRestricted:           prev.MatterChiralRestricted,
		MatterMirrorAmbiguity:            prev.MatterMirrorAmbiguity,
		GenericFockToContactMapsExist:    true,
		GenericTensorToContactMapsExist:  true,
		CanonicalFockToContactMaps:       canonicalFockMaps,
		FockToContactIntertwinersDerived: fockIntertwiners,
		T3RPullbackRowsDerived:           t3rPullbacks,
		ChiralityPullbackRowsDerived:     chirPullbacks,
		BMinusLPullbackRowsDerived:       blPullbacks,
		SU2LPullbackRowsDerived:          su2Pullbacks,
		HyperchargeRowsDerived:           yRows,
		ElectricChargeRowsDerived:        0,
		RepresentationCompleteRows:       repRows,
		RepresentationOpenRows:           contactRows,
		ContactBetaRowsAllowed:           betaRows,
		ContactZeroRowsProved:            0,
		FullBetaMatchingTensorDerived:    false,
		ThresholdCorrectedBetaDerived:    false,
		BetaPermissionFirewallClosed:     repRows == 0 && betaRows == 0,
		ResidualS6Choices:                prev.ResidualS6Choices,
		ResidualNullityBefore:            prev.ResidualNullityAfter,
		ResidualNullityAfter:             prev.ResidualNullityAfter,
		HiddenObservedInputUsed:          false,
		PhysicalWeakAngleDerived:         false,
		FineStructureDerived:             false,
		PhysicalMassesDerived:            false,
		PhysicalScaleDerived:             false,
		TruthStatement:                   truth,
		RejectedClaims: []string{
			"any rank-seven linear map from H_Fock to contact rows is a physical pullback",
			"a 9-dimensional generic kernel inside H_Fock is canonical",
			"the tensor carrier H_Fock⊗H_phi selects contact rows by dimension alone",
			"the leptoquark six-block supplies contact chirality without an S6 assignment",
			"contact spectral identity can pull back matter T3R",
			"matter-side hyperaudit orientation can be borrowed as contact chirality",
		},
		RemainingUnknowns: []string{
			"canonical Fock-to-contact map or relation",
			"operator intertwining equations for T3R, chirality, B-L, and SU(2)L",
			"selected kernel/quotient from H_Fock or H_Fock⊗H_phi to seven contact rows",
			"canonical S6 assignment of leptoquark current slots to contact rows",
			"local field map, hypercharge row, mass activation, and decoupling rule",
		},
		RecommendedNextGate: "Gate 138 — Fock-contact kernel selection / operator-intertwining obstruction theorem",
	}, nil
}

func buildCandidates(matterDim, scalarDim, tensorDim, contactRows, lqRows, s6 int) []Candidate {
	return []Candidate{
		{Name: "generic H_Fock → R7_contact rank-seven map", Kind: FockToContactLinearMap, Domain: "H_Fock", Codomain: "R7_contact", DomainDim: matterDim, CodomainDim: contactRows, RankTarget: contactRows, KernelDimIfSurjective: matterDim - contactRows, GenericMapsExist: true, Canonical: false, MatterSideOnly: false, RequiresContinuousChoice: true, HiddenDiscreteChoices: 0, Obstruction: "surjective maps exist, but choosing a 9-dimensional kernel in H_Fock is not finite-canonical and no operator-intertwining law is derived"},
		{Name: "generic H_Fock⊗H_phi → R7_contact rank-seven map", Kind: TensorToContactLinearMap, Domain: "H_Fock⊗H_phi", Codomain: "R7_contact", DomainDim: tensorDim, CodomainDim: contactRows, RankTarget: contactRows, KernelDimIfSurjective: tensorDim - contactRows, GenericMapsExist: true, Canonical: false, RequiresContinuousChoice: true, Obstruction: "the 64-dimensional tensor carrier admits many maps to seven rows; no selected tensor quotient or row semantics is derived"},
		{Name: "leptoquark six current slots → six contact rows", Kind: LeptoquarkSixBlockMap, Domain: "LQ_current_six", Codomain: "contact_six_subblock", DomainDim: lqRows, CodomainDim: lqRows, RankTarget: lqRows, KernelDimIfSurjective: 0, GenericMapsExist: true, Canonical: false, CurrentSideOnly: true, RequiresS6Choice: true, HiddenDiscreteChoices: s6, Obstruction: fmt.Sprintf("requires an S6 assignment of %d possible permutations", s6)},
		{Name: "current 1+6 quotient branch → contact 1+6 refinement", Kind: CurrentOnePlusSixBranchMap, Domain: "u4_current_quotient", Codomain: "contact_rows", DomainDim: contactRows, CodomainDim: contactRows, RankTarget: contactRows, KernelDimIfSurjective: 0, GenericMapsExist: true, Canonical: false, CurrentSideOnly: true, RequiresOnePlusSixRefinement: true, HiddenDiscreteChoices: 7 * s6, Obstruction: "choosing which contact row is singlet and ordering the remaining six rows is not selected"},
		{Name: "scalar H_phi → seven contact rows", Kind: ScalarToContactMap, Domain: "H_phi", Codomain: "R7_contact", DomainDim: scalarDim, CodomainDim: contactRows, RankTarget: contactRows, KernelDimIfSurjective: -1, GenericMapsExist: false, Canonical: false, MatterSideOnly: true, Obstruction: "a four-dimensional scalar carrier cannot surject onto seven independent contact rows"},
		{Name: "contact spectral identity on R7_contact", Kind: ContactSpectralIdentityMap, Domain: "R7_contact", Codomain: "R7_contact", DomainDim: contactRows, CodomainDim: contactRows, RankTarget: contactRows, KernelDimIfSurjective: 0, GenericMapsExist: true, Canonical: true, ContactSideOnly: true, PreservesContactRows: true, Obstruction: "identity preserves contact diagnostics, but does not pull back matter-side T3R/chirality"},
	}
}

func buildRows(prevRows []contactlqt3r.CandidateRow) []PullbackRow {
	out := make([]PullbackRow, 0, len(prevRows))
	for _, r := range prevRows {
		out = append(out, PullbackRow{
			Name:                     r.Name,
			ContactRow:               r.Name,
			MatterT3RCandidateValues: append([]float64(nil), r.MatterT3RCandidateValues...),
			T3RPullbackDerived:       false,
			ChiralityPullbackDerived: false,
			BMinusLPullbackDerived:   false,
			SU2LPullbackDerived:      false,
			HyperchargeRowDerived:    false,
			LocalFieldDerived:        false,
			RepresentationComplete:   false,
			BetaPermitted:            false,
			RequiresFockContactMap:   true,
			RequiresS6Choice:         r.RequiresS6Choice,
			Obstruction:              "matter-side ±T3R candidates remain hypothetical because no Fock-to-contact intertwiner is derived",
		})
	}
	return out
}

func FormatCandidates(c []Candidate) string {
	parts := make([]string, 0, len(c))
	for _, x := range c {
		parts = append(parts, fmt.Sprintf("%s[%s→%s dim=%d→%d rank=%d kernel=%d canonical=%t intertwinesT3R=%t hidden=%d; %s]", x.Name, x.Domain, x.Codomain, x.DomainDim, x.CodomainDim, x.RankTarget, x.KernelDimIfSurjective, x.Canonical, x.IntertwinesT3R, x.HiddenDiscreteChoices, x.Obstruction))
	}
	return strings.Join(parts, "; ")
}

func FormatRows(rows []PullbackRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s:T3R=%s pullback=%t chirality=%t Y=%t rep=%t", r.Name, formatFloats(r.MatterT3RCandidateValues), r.T3RPullbackDerived, r.ChiralityPullbackDerived, r.HyperchargeRowDerived, r.RepresentationComplete))
	}
	return strings.Join(parts, "; ")
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("matter=%d scalar=%d tensor=%d contact=%d fockKernel=%d tensorKernel=%d canonicalMaps=%d intertwiners=%d T3R=%d chirality=%d beta=%d nullity=%d→%d", s.MatterDimension, s.ScalarDimension, s.TensorDimension, s.ContactRows, s.FockToContactGenericKernelDim, s.TensorToContactGenericKernelDim, s.CanonicalFockToContactMaps, s.FockToContactIntertwinersDerived, s.T3RPullbackRowsDerived, s.ChiralityPullbackRowsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(items []string) string { return strings.Join(items, "; ") }

func count[T any](items []T, pred func(T) bool) int {
	n := 0
	for _, item := range items {
		if pred(item) {
			n++
		}
	}
	return n
}

func formatFloats(vals []float64) string {
	cp := append([]float64(nil), vals...)
	sort.Float64s(cp)
	parts := make([]string, 0, len(cp))
	for _, v := range cp {
		parts = append(parts, fmt.Sprintf("%+.6f", v))
	}
	return "[" + strings.Join(parts, ",") + "]"
}
