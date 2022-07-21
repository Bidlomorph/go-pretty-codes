Pretty codes project

Funcs `prettyNumber` and `prettyNumberSlice` may take 3 params for return pretty code-numbers.
Example: 
`result, err = prettyNumberSlice(9, 3, 2)` may return 152333666, 444215552 ...

Pretty code dependent on `codeLength`(length of returned code),
`orderLength`(length of maximum same numbers) and `orderCount`(number of sets of identical numbers).

Difference between `prettyNumber` and `prettyNumberSlice`:
`prettyNumber` - works with maps.
`prettyNumberSlice` - works with slices.

Performance:
`prettyNumber` - 8950-9100 ns/op.
`prettyNumberSlice` - 8200-8550 ns/op.

