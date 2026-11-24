# Check if Two Trees are Identical

> Striver SDE Sheet 121 · Day 18 Binary Tree II — structural recursion: equal roots and recursively equal children; the nil handling is the whole exam. Go, stdlib only, tested.

## Problem

Are two binary trees exactly equal — same shape *and* same values at every position?

## Approach

The recursion mirrors the definition: both nil → identical; exactly one nil → not (this is what catches shape differences); otherwise equal values and identical subtrees pairwise. The one-line nil case `if a == nil || b == nil { return a == b }` folds three base cases into one comparison — the kind of compression that's clear once seen and buggy when hand-rolled as separate ifs. Go's short-circuit `&&` stops at the first mismatch, so unequal trees are cheap to reject.

**Complexity:** O(min(n, m)) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
