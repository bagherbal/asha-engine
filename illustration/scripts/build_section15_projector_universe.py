#!/usr/bin/env python3
from __future__ import annotations

import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
if str(ROOT) not in sys.path:
    sys.path.insert(0, str(ROOT))

from asha_figures.projector_universe import write_outputs


def main() -> None:
    out_dir = ROOT / "outputs" / "section15"
    paths = write_outputs(out_dir)
    for key, value in paths.items():
        print(f"{key}: {value}")


if __name__ == "__main__":
    main()
