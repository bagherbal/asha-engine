package app

import (
	"fmt"

	"github.com/bagherbal/asha-engine/internal/report"
	"github.com/bagherbal/asha-engine/pkg/clifford"
	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/exterior"
	"github.com/bagherbal/asha-engine/pkg/gauge"
	"github.com/bagherbal/asha-engine/pkg/geometry/boolean"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/geometry/g2"
	"github.com/bagherbal/asha-engine/pkg/phase"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Run() error {
	registry := theorem.NewRegistry(
		exterior.GradeStructureTheorem(8),
		clifford.StructureTheorem(clifford.Signature{Positive: 1, Negative: 7}),
		phase.CovariantPhaseSpaceTheorem(4),
		boolean.IncidenceSupportTheorem(8, 3, 4),
		g2.CalibrationSupportTheorem(),
		contact.ContactSpaceTheorem(),
		bsector.ContactVacuumTheorem(),
		gauge.ContactCentralizerTheorem(),
	)

	results := registry.RunAll()
	fmt.Print(report.RenderTerminal(results))
	return nil
}
