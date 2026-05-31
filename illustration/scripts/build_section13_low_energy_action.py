#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

TOOLKIT_ROOT = Path(__file__).resolve().parents[1]
if str(TOOLKIT_ROOT) not in sys.path:
    sys.path.insert(0, str(TOOLKIT_ROOT))

from asha_figures.readme_low_energy_action import low_energy_action_contract, export_contract, validate_contract
from asha_figures.section13_svg import render_low_energy_action
from asha_figures.style import STYLE


def main() -> None:
    parser = argparse.ArgumentParser(description="Build ASHA README current low-energy action skeleton visual.")
    parser.add_argument("--out", default="outputs/section13", help="Output directory, relative to toolkit root unless absolute.")
    parser.add_argument("--png", action="store_true", help="Also render PNG via cairosvg when available.")
    args = parser.parse_args()
    out_dir = Path(args.out)
    if not out_dir.is_absolute():
        out_dir = TOOLKIT_ROOT / out_dir
    out_dir.mkdir(parents=True, exist_ok=True)

    c = low_energy_action_contract()
    checks = validate_contract(c)
    if checks["status"] == "FAIL":
        raise SystemExit(f"validation failed for {c.figure_id}: {checks}")

    svg_path = out_dir / f"{c.figure_id}.svg"
    json_path = out_dir / f"{c.figure_id}.geometry.json"
    manifest_path = out_dir / f"{c.figure_id}.manifest.json"
    render_low_energy_action(c, svg_path)
    json_path.write_text(json.dumps(export_contract(c), indent=2, ensure_ascii=False), encoding="utf-8")
    manifest = {
        "artifact": svg_path.name,
        "companion_geometry": json_path.name,
        "status": checks["status"],
        "checks": checks,
        "aesthetic": "Transcendental Geometric Minimalism",
        "readme_anchor": c.readme_anchor,
        "notes": c.notes,
    }
    manifest_path.write_text(json.dumps(manifest, indent=2, ensure_ascii=False), encoding="utf-8")
    if args.png:
        try:
            import cairosvg
            png_path = out_dir / f"{c.figure_id}.png"
            cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=STYLE.width, output_height=STYLE.height)
            print(png_path)
        except Exception as exc:
            print(f"PNG export skipped for {c.figure_id}: {exc}")
    print(svg_path)
    print(json_path)
    print(manifest_path)

if __name__ == "__main__":
    main()
