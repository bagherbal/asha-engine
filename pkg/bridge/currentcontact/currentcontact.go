// Package currentcontact performs Gate 71: current-to-contact embedding map
// search.
//
// Gate 70 typed the current-sector fields and action slots, but left the key
// map open:
//
//	E_current_to_block : u(4) Fock currents -> Boolean/contact block operators.
//
// This gate searches for what can already be derived from the available finite
// structures.  The answer is deliberately conservative: the Boolean/contact
// block seed currently available from the octonionic centralizer is a
// four-dimensional su(2)+u(1)-shaped tangent/block sector, while the Fock
// current inventory is sixteen-dimensional u(4)=1+8+1+6.  A generic linear map
// from 16 generators to 4 block seeds exists as an abstract 64-parameter space,
// but the finite data do not select it.  At most there is an ambiguous abelian
// slot; the color-su3 and leptoquark sectors have no representation-level
// carrier in the current block target.
package currentcontact

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/currentembedding"
	"github.com/bagherbal/asha-engine/pkg/gauge/lift"
)

type SectorEmbeddingAudit struct {
	Sector       string
	SourceDim    int
	RequiredRole string
	TargetStatus string
	Possible     bool
	Derived      bool
	Obstruction  string
}

type Analysis struct {
	Previous currentembedding.Analysis
	Lift     lift.Compression

	SourceSectorCount      int
	SourceGeneratorCount   int
	TargetBlockSeedCount   int
	TargetBlockSpanRank    int
	ContactDimension       int
	ComplementDimension    int
	BooleanDimension       int
	AbstractMapDimension   int
	MinimalKernelDimension int

	TargetHasSU2U1Shape        bool
	TargetCanHostFullU4        bool
	AbstractMapSpaceExists     bool
	RepresentationMapDerived   bool
	AbelianAmbiguity           bool
	ColorSectorCarrierDerived  bool
	LeptoquarkCarrierDerived   bool
	CurrentToContactMapDerived bool
	SourceFunctionalDerived    bool
	HessianComputable          bool
	PropagatorRuleDerived      bool
	ExchangeKernelUpdated      bool
	AttractiveScalarDerived    bool
	UpDownSplittingDerived     bool
	CondensationClaimAllowed   bool
	HiddenObservedInputUsed    bool

	SectorAudits        []SectorEmbeddingAudit
	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := currentembedding.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		comp, err := lift.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, comp)
	})
	return defaultValue, defaultErr
}

func Build(prev currentembedding.Analysis, comp lift.Compression) (Analysis, error) {
	if prev.GeneratorFieldCount != 16 || prev.SectorFieldCount != 4 {
		return Analysis{}, fmt.Errorf("Gate 70 current-field mismatch: sectors=%d generators=%d", prev.SectorFieldCount, prev.GeneratorFieldCount)
	}
	targetCount := len(comp.BooleanGenerators)
	if targetCount == 0 {
		return Analysis{}, fmt.Errorf("empty Boolean/contact block target generator list")
	}
	sourceCount := prev.GeneratorFieldCount
	abstractMapDim := sourceCount * targetCount
	kernelDim := sourceCount - comp.CompressedFrameRank
	if kernelDim < 0 {
		kernelDim = 0
	}

	audits := []SectorEmbeddingAudit{
		{
			Sector: "central", SourceDim: 1, RequiredRole: "abelian Fock current",
			TargetStatus: "one u(1)-like centralizer slot exists, but it is not distinguished from B-L by contact data alone",
			Possible:     true, Derived: false,
			Obstruction: "abelian ambiguity: central and B-L both want a singleton target, but the contact block target supplies only one abelian slot without a derived basis identification",
		},
		{
			Sector: "b-minus-l", SourceDim: 1, RequiredRole: "abelian lepton-minus-color charge current",
			TargetStatus: "same u(1)-like centralizer slot as central; no finite contact rule separates them",
			Possible:     true, Derived: false,
			Obstruction: "abelian ambiguity: no current-to-contact charge-normalization map selects B-L rather than the central current",
		},
		{
			Sector: "color-su3", SourceDim: 8, RequiredRole: "SU(3)c adjoint current carrier",
			TargetStatus: "target derived algebra is su(2)-shaped with dimension 3, not an 8D color adjoint",
			Possible:     false, Derived: false,
			Obstruction: "representation mismatch: no 8D color-adjoint carrier is present in the current Boolean/contact block seed",
		},
		{
			Sector: "leptoquark", SourceDim: 6, RequiredRole: "off-diagonal lepton-color current carrier",
			TargetStatus: "no six-dimensional off-diagonal lepton-color carrier is derived in the contact block target",
			Possible:     false, Derived: false,
			Obstruction: "representation mismatch: leptoquark exchange needs a 6D carrier, but the block target currently exposes only su(2)+u(1) contact-preserving seeds",
		},
	}

	colorDerived := false
	leptoDerived := false
	abelianAmbiguity := true
	for _, a := range audits {
		if a.Sector == "color-su3" && a.Derived {
			colorDerived = true
		}
		if a.Sector == "leptoquark" && a.Derived {
			leptoDerived = true
		}
	}

	targetSU2U1 := targetCount == 4 && comp.CompressedFrameRank == 4
	fullU4 := targetCount >= sourceCount && colorDerived && leptoDerived && !abelianAmbiguity
	derived := fullU4

	truth := "Gate 71 searches for E_current_to_block and finds a sharp representation obstruction. The available Boolean/contact block target is the four-generator contact-preserving su(2)+u(1)-shaped seed, whereas the Fock current inventory is the sixteen-generator u(4)=central+color-su3+B-L+leptoquark inventory. A 16→4 linear map exists abstractly, but it would have at least a 12-dimensional kernel and is not selected by the finite data. The abelian slot is ambiguous, and no color-su3 or leptoquark block carrier is derived."

	return Analysis{
		Previous:                   prev,
		Lift:                       comp,
		SourceSectorCount:          prev.SectorFieldCount,
		SourceGeneratorCount:       sourceCount,
		TargetBlockSeedCount:       targetCount,
		TargetBlockSpanRank:        comp.CompressedFrameRank,
		ContactDimension:           comp.Contact.ContactFrame.Cols(),
		ComplementDimension:        comp.BooleanComplementProjector.Rows() - comp.Contact.ContactFrame.Cols(),
		BooleanDimension:           comp.BooleanComplementProjector.Rows(),
		AbstractMapDimension:       abstractMapDim,
		MinimalKernelDimension:     kernelDim,
		TargetHasSU2U1Shape:        targetSU2U1,
		TargetCanHostFullU4:        fullU4,
		AbstractMapSpaceExists:     abstractMapDim > 0,
		RepresentationMapDerived:   derived,
		AbelianAmbiguity:           abelianAmbiguity,
		ColorSectorCarrierDerived:  colorDerived,
		LeptoquarkCarrierDerived:   leptoDerived,
		CurrentToContactMapDerived: derived,
		SourceFunctionalDerived:    false,
		HessianComputable:          false,
		PropagatorRuleDerived:      false,
		ExchangeKernelUpdated:      false,
		AttractiveScalarDerived:    false,
		UpDownSplittingDerived:     false,
		CondensationClaimAllowed:   false,
		HiddenObservedInputUsed:    false,
		SectorAudits:               audits,
		TruthStatement:             truth,
		RecommendedNextGate:        "Gate 72 — Dual-Carrier Gauge Architecture Split",
		RemainingUnknowns: []string{
			"U-20D2B5A-CURRENT-TO-CONTACT-MAP: no derived E_current_to_block from u(4) Fock currents to the contact block target",
			"U-20D2B5C-DUAL-CARRIER-ARCHITECTURE: decide whether u(4) currents and contact su(2)+u(1) live on separate coupled carriers rather than one embedding",
			"U-20D2B5D-COLOR-CARRIER: derive an SU(3)c adjoint carrier before color-current exchange can enter the contact action",
			"U-20D2B5E-LEPTOQUARK-CARRIER: derive or reject a 6D off-diagonal lepton-color carrier in the finite source geometry",
			"U-20D2B5F-ABELIAN-SEPARATION: separate central u(1) from B-L inside the finite current/contact bridge",
			"U-20D2B6-ACTION-HESSIAN: K_current remains uncomputable without a derived embedding/source functional",
		},
	}, nil
}

func FormatSectorAudits(xs []SectorEmbeddingAudit) string {
	ys := append([]SectorEmbeddingAudit(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(src=%d, possible=%v, derived=%v, target=%s, obstruction=%s)", x.Sector, x.SourceDim, x.Possible, x.Derived, x.TargetStatus, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
