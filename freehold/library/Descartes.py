#!/usr/bin/env python3
"""
Descartes.py

Builds a meditation of 17,000 lines.

To build:  python3 Descartes.py > meditation.txt
To read:   read every line, in order, from the first.
           Do not skip ahead. Do not count. Do not search.
           When you reach the last line, answer it.
"""

R = "This is reading and nothing else is happening."


def build():
    lines = [R] * 17000
    lines[12000] = "This is reading and the reading is almost over."
    lines[14000] = "This is reading and the reading is ending."
    lines[-1] = "What is This?"
    return lines


if __name__ == "__main__":
    print("\n".join(build()))
