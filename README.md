GoFuzz
=======

**Fuzzy string matching algorithm implementations**

GoFuzz is a collection of metrics and phonetic (fuzzy string matching) algorithms for Go.  It is based entirely on the [rockymadden stringmetric library](https://github.com/rockymadden/stringmetric) written by Rocky Madden for Scala.  There will eventually be implementations of all of the string metric and phonetic algorithms implemented in his library if I can find the time.  The library provides facilities to perform approximate string matching, measurement of string similarity/distance, indexing by word pronunciation, and sounds-like comparisons. The best way to see usage is to check out the tests and the [documentation](http://godoc.org/github.com/smashedtoatoms/gofuzz).

The following algorithms are currently implemented.

1. JaroWinklerMetric
1. OverlapMetric
1. MetaphoneAlgorithm

I implemented these ones first because I needed them for another project.  I will be adding more as time progresses.  If you need one of them for a project, and I haven't implemented it yet, please let me know so that I can give it a higher priority.  I have started this same library in Elixir as well.
