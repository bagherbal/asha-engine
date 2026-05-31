#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

TOOLKIT_ROOT = Path(__file__).resolve().parents[1]
if str(TOOLKIT_ROOT) not in sys.path:
    sys.path.insert(0, str(TOOLKIT_ROOT))

from asha_figures.contact_depth import build_contact_depth_geometry, export_contact_depth_json_dict, validate_contact_depth_geometry
from asha_figures.style import STYLE
from asha_figures.section3_svg import render_contact_depth_svg


def main() -> None:
    parser = argparse.ArgumentParser(description="Build ASHA README Contact Seven / Depth Triple figure artifacts.")
    parser.add_argument("--out", default="outputs/section3", help="Output directory, relative to toolkit root unless absolute.")
    parser.add_argument("--png", action="store_true", help="Also render PNG via cairosvg when available.")
    args = parser.parse_args()

    out_dir = Path(args.out)
    if not out_dir.is_absolute():
        out_dir = TOOLKIT_ROOT / out_dir
    out_dir.mkdir(parents=True, exist_ok=True)

    geometry = build_contact_depth_geometry(width=STYLE.width, height=STYLE.height)
    checks = validate_contact_depth_geometry(geometry)
    if checks["status"] != "PASS_README_CONTACT_SEVEN_DEPTH_TRIPLE":
        raise SystemExit(f"validation failed: {checks}")

    svg_path = out_dir / "asha_readme_contact_seven_depth_triple.svg"
    json_path = out_dir / "asha_readme_contact_seven_depth_triple.geometry.json"
    manifest_path = out_dir / "asha_readme_contact_seven_depth_triple.manifest.json"

    render_contact_depth_svg(geometry, svg_path)
    json_path.write_text(json.dumps(export_contact_depth_json_dict(geometry), indent=2, ensure_ascii=False), encoding="utf-8")
    manifest_path.write_text(json.dumps({
        "artifact": svg_path.name,
        "companion_geometry": json_path.name,
        "status": checks["status"],
        "checks": checks,
        "aesthetic": "Transcendental Geometric Minimalism",
        "readme_anchor": "Contact seven and the depth triple",
        "notes": [
            "This is the next README-critical visual after the measurement-octave / contact-vacuum figures.",
            "The carrier contract is exact: V8 basis, x0 observer reference, seven contact directions, and three phase planes.",
            "The N_Q diagonal entries and W_Q numerical evaluations are generated from the exact fractions 1/3, 1/2, 2/3.",
            "The W_Q bars are intentionally logarithmic because linear rendering would visually erase the later layers.",
            "No particle identities, flavor claims, or mass predictions are encoded in this figure."
        ],
    }, indent=2, ensure_ascii=False), encoding="utf-8")

    if args.png:
        try:
            import cairosvg
            png_path = out_dir / "asha_readme_contact_seven_depth_triple.png"
            cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=STYLE.width, output_height=STYLE.height)
            print(png_path)
        except Exception as exc:
            print(f"PNG export skipped: {exc}")

    print(svg_path)
    print(json_path)
    print(manifest_path)


if __name__ == "__main__":
    main()
