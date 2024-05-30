[![main](https://github.com/flowerinthenight/memx/actions/workflows/main.yml/badge.svg)](https://github.com/flowerinthenight/memx/actions/workflows/main.yml)

`memx` returns the calling process' memory usage (private + shared) in kB. Linux-specific only.

It does this by reading [`PSS`](https://en.wikipedia.org/wiki/Proportional_set_size) from `/proc/{pid}/smaps[_rollup]`.
