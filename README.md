# hdls - Hardware Description Language Server
An attempt to create versatile language server for hardware description.
The project is currently in the design phase, any suggestions are very welcome.

## Design phase

### Requirements
1. **Orthogonality**.
2. Support for both VHDL and SystemVerilog.
3. Support for libraries, which structure is defined in different ways ([FuseSoc](https://github.com/olofk/fusesoc) `.core` files, [TOML](https://github.com/toml-lang/toml)).
4. Speed/performance.
5. Libraries cyclic dependency detection.
Cyclic dependencies are bad, *really bad*.
They introduce chaos.
With cyclic dependencies it is also much harder to analyze libraries in parallel.
Cyclic libraries dependencies must generate fatal error to enforce user to fix libraries design.
6. Libraries dependency graph generation - graph generation is needed for parallel libraries analysis anyway.
7. Caching of libraries analysis results and automatic changes detection.
8. There must be only single global configuration file, and single configuration file per workspace.

### Implementation decisions

#### Language
[Go](https://golang.org/).
Why?
Orthogonal, efficient, simple, stable, oriented on concurrency.

#### VHDL analysis
[GHDL](https://github.com/ghdl/ghdl) coupled via shared library.
Why?
Nothing even close to what GHDL offers, plus it has been actively maintained for few years.

#### SystemVerilog analysis
I don't know, I have no experience.

#### Configuration file format
YAML or TOML - to be decided (slightly *biased* towards TOML, simpler, easier to read).
JSON is too verbose, does not allow comments.

### Prototypes
Before any target work gets started some prototypes are needed.
It is worth to remind what prototype is.
Prototype is temporary, disposable code.
Its main purpose is to gain knowledge and answer questions.
If you want to reuse your prototype code within the final implementation, then it is not a prototype!
This is explained in a wonderful way in the [The Pragmatic Programmer](https://en.wikipedia.org/wiki/The_Pragmatic_Programmer) book.
Independent prototypes answering following questions are needed:
1. How to cache GHDL analysis results?
2. Does GHDL even allow to parse independent files in parallel?
3. What is the *best/fastest* way to detect changes of files within libraries (not workspaces).
Should libraries files be loaded only on start, or periodically during the work?
Maybe `SIGUSR*` should enforce libraries reload?
