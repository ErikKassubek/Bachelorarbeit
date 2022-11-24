% Dynamic Analysis of Message-Passing Go Programs

# Overview

We consider the dynamic analysis of message-passing Go Programs.
This includes

* unbuffered channels

* buffered channels

* selective communication

* closing of channels

The approach shall be based on [Two-Phase Dynamic Analysis of Message-Passing Go Programs based on Vector Clocks](https://arxiv.org/abs/1807.03585). Henceforth, referred to as PPDP18.

Brief summary of PPDP18.

The program's run-time behavior is recorded as a trace.
Interesting aspects are:


* Events contain information about their communication partners.
  For example, a send/receive pair can identify its partner evern via a unique hash index.
  This is important as the assumption is that communication partners remain unchanged.

* Each operation implies a *pre* event. For example, the pre event of a receive operation
  indicates that we attempt to receive a value via some channel.

* If the operation could commit (i.e. took place, we also issue a *post* event.

* Infer a happens-before relation among events.

Thus, we can deal with several analysis scenarios.


* MP Message contention

* SC Send on a closed channel

* DR Deadlock recovery

* AC Alternative communication partners for send/receive pairs

## Happens-before and buffered/unbuffered channels

The following [happens-before](https://en.wikipedia.org/wiki/Happened-before) condition apply in Go. Details see [here](https://go.dev/ref/mem)

A send on a channel is synchronized before the completion of the corresponding receive from that channel.

For *unbuffered channels* the following condition applies:

* A receive from an unbuffered channel is synchronized before the completion of the corresponding send on that channel.

For *buffered channels* the following conditions applies:


* The kth receive on a channel with capacity C is synchronized before the completion of the k+Cth send from that channel completes.


MS: The above conditions ought to be implemented in PPDP18 where the happens-before relation is represented via vector clocks.


# Plan of attack

1. Review of PPDP18 and its implementation (c)

2. Compare PPDP18 and GFuzz

3. Prototype implementation


Specific things to look into are.

* Instrumentation and tracing.

    * Can this be done more systematically using some Go tool?

* Trace replay

    * PPDP18 applies a non-deterministic strategy, can we be more deterministic?


# Related works

## [Who Goes First? Detecting Go Concurrency Bugs via Message Reordering](https://songlh.github.io/paper/gfuzz.pdf)

Introduces GFuzz (dynamic) analysis tool.

"GFuzz only considers messages waited for by the same select as concurrent, and thus it misses many other possible concurrent messages"

This sounds rather limiting.



## [Automatically Detecting and Fixing Concurrency Bugs in Go Software Systems](https://songlh.github.io/paper/gcatch.pdf)

Same authors as "GFuzz".

Applies static analysis methods to identify bugs related to blocked channel operations.
Blocking misuse-of-channel (BMOC).

GCatch is the name of the (static) analysis tool.

May require several hours (relies on the SSA package to analyze Go program code).


Some heuristics such as increasing the channel's buffer size are considered
to automatically fix bugs.


NOTE. Does not cite PPDP18. GoBench doesn't seem to be cited as well.

## [GOAT: Automated Concurrency Analysis and Debugging Tool for Go Paper Type: Tool / Benchmark](https://staheri.github.io/files/iiswc21-paper7.pdf)

Same authors as [Automated Dynamic Concurrency Analysis for Go](https://arxiv.org/ftp/arxiv/papers/2105/2105.11064.pdf)

Extends

https://pkg.go.dev/runtime/trace

to keep track of concurrent events (send, receive, ...) but no reads, writes as it seems.

The extended tracer is used by GOAT.

## Further works

### Go Tracing

https://github.com/divan/gotrace

https://github.com/staheri/ectgo#intsall-ect-patch

### GoBench: A Benchmark Suite of Real-World Go Concurrency Bugs

The paper title says it all.
