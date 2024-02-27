# tippecanoe-concat-distinct

A simple program intended to be used with [tippecanoe](https://github.com/felt/tippecanoe).

Reads from `std in`, modifies the geo json by removing any duplicates from the value of the target field passed to the program, and writes the modified geo json back to `std out`.

When using tippecanoe's `concat` accumulate attribute when clustering points. Points which share the same value for the accumulated attribute will each be concatenated, leading to duplicates.

Using this program via tippecanoe's post-filter functionality allows for use cases where the duplicates are not useful and tile size needs to be reduced.
