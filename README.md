Metago                                                              
====

A meta-language for building Go types with some interesting built in functionality:

- Forwards and backwards compatibility through statically assigned type and attribute identifiers
- A form of subtyping polymorphism
- Self-differencing: err, d = a.Diff(b) produces a record of the differences between two objects of the same type
- Compact binary serialization/deserialization of both objects and difference records to io.Writer/io.Reader

Docs: [![GoDoc](https://godoc.org/github.com/idawes/metago?status.svg)](https://godoc.org/github.com/idawes/metago)

[![views](https://sourcegraph.com/api/repos/github.com/idawes/metago/.counters/views.svg)](https://sourcegraph.com/github.com/idawes/metago) [![views 24h](https://sourcegraph.com/api/repos/github.com/idawes/metago/.counters/views-24h.svg)](https://sourcegraph.com/github.com/idawes/metago)
