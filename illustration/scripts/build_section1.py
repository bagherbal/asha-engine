#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

TOOLKIT_ROOT = Path(__file__).resolve().parents[1]
if str(TOOLKIT_ROOT) not in sys.path:
    sys.path.insert(0, str(TOOLKIT_ROOT))

from asha_figures.boolean_lattice import build_nodes, build_edges, export_json_dict, validate
from asha_figures.style import STYLE
from asha_figures.svg import render_section1_svg


def main() -> None:
    parser = argparse.ArgumentParser(description="Build ASHA Section 1 Measurement Ladder figure artifacts.")
    parser.add_argument("--out", default="outputs/section1", help="Output directory, relative to toolkit root unless absolute.")
    parser.add_argument("--png", action="store_true", help="Also render PNG via cairosvg when available.")
    args = parser.parse_args()

    toolkit_root = TOOLKIT_ROOT
    out_dir = Path(args.out)
    if not out_dir.is_absolute():
        out_dir = toolkit_root / out_dir
    out_dir.mkdir(parents=True, exist_ok=True)

    nodes = build_nodes(
        width=STYLE.width,
        margin_x=STYLE.margin_x,
        tier_top=STYLE.tier_top,
        tier_bottom=STYLE.tier_bottom,
        node_radius=STYLE.node_radius,
        middle_node_radius=STYLE.middle_node_radius,
        terminal_node_radius=STYLE.terminal_node_radius,
    )
    edges = build_edges(nodes)
    checks = validate(nodes, edges)
    if not (checks["node_count_exact"] and checks["tier_counts_exact"] and checks["edge_count_exact"]):
        raise SystemExit(f"validation failed: {checks}")

    svg_path = out_dir / "asha_section1_measurement_ladder_cl_1_7.svg"
    json_path = out_dir / "asha_section1_measurement_ladder_cl_1_7.coordinates.json"
    manifest_path = out_dir / "asha_section1_measurement_ladder_cl_1_7.manifest.json"

    render_section1_svg(nodes, edges, svg_path)
    json_path.write_text(json.dumps(export_json_dict(nodes, edges), indent=2, ensure_ascii=False), encoding="utf-8")
    manifest_path.write_text(json.dumps({
        "artifact": svg_path.name,
        "companion_coordinates": json_path.name,
        "status": "PASS_EXACT_PASCAL_TIER_LADDER",
        "checks": checks,
        "aesthetic": "Transcendental Geometric Minimalism",
        "notes": [
            "The figure encodes Cℓ(1,7) as an exterior-grade basis ladder.",
            "Warm signature cue means the blade contains e₀; cyan means space-like-only blade.",
            "The middle Λ⁴ℝ⁸ chamber is visually emphasized because it is the 70-dimensional arena used by later Boolean/G₂ contact-vacuum gates.",
            "This is intentionally not a particle catalogue."
        ],
    }, indent=2, ensure_ascii=False), encoding="utf-8")

    if args.png:
        try:
            import cairosvg
            png_path = out_dir / "asha_section1_measurement_ladder_cl_1_7.png"
            cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=STYLE.width, output_height=STYLE.height)
            print(png_path)
        except Exception as exc:
            print(f"PNG export skipped: {exc}")

    print(svg_path)
    print(json_path)
    print(manifest_path)

if __name__ == "__main__":
    main()
