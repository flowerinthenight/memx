[![main](https://github.com/flowerinthenight/memx/actions/workflows/main.yml/badge.svg)](https://github.com/flowerinthenight/memx/actions/workflows/main.yml)

Returns the calling process' memory usage (internal + shared) in kB. Linux-specific only. It does this by reading [`PSS`](https://en.wikipedia.org/wiki/Proportional_set_size) from `/proc/{pid}/smaps`.
