# Bachelor Thesis: Dynamic Analysis of Massage-Passing Go Programs
Code at:
https://github.com/ErikKassubek/GoChan

### 2022-10-31
Welche Tracing Möglichkeiten gibt es?
Vorherige Arbeiten und deren eventuelle Einschränkungen,
z.B. Behandlung von select, ...

Vorschlag für Ansatz.
Textuelle Transformation des Go Programmcodes, siehe PPDP18.

### 2022-11-09
Beschreibung des Tracers (-> Bachelor Thesis).
Vergleiche zu bestehenden Ansaetzten? Z.B. wieso wurde nicht das "go-trace" package verwendet ...

Technische Details.

1. Kann es zu Data Races beim Tracing kommen? Jeder Thread hat seine eigene Position im Slice, aber im Fall von Updates
gibt es vlt Verwaltungstrukturen?

2. Skaliert der Tracer fuer groessere Programme? Overhead?

3. Behandlung weiterer Sprachfeatures, z.B. RWMutex?
    In diesem Zusammenhang, GFuzz scheint nur "channels" zu behandeln.
    Aber in vielen "real world" Programmen wird aber sicher ein Mix aus channels und Mutex verwendet.

Ausblick.

1. Welche Analyseszenarien wollen wir betrachten?

2. Der "two-phase" Ansatz in PPDP'18 berechnet fuer jedes Event seine Vectorclock.

  Auszug aus PPD"18   "We summarize. The vector clock annotated trace E contains a wealth of information."
  Je nach Szenario, sollten wir nur fuer gewisse Events deren Vectorclock speichern.

### 2022-11-28

1. Tracing "non-blocking vs blocking"
Source-code Instrumentierung
Vergleich zu GFuzz Instrumentierung.

2. Instrumentierung Performanz.
  - Laufzeit Overhead
  - Source-to-source Transformation

3. Anpassung von Bachelorprojekt "Deadlockanalyse fuer RW-Mutex"
Annahme "globaler" Zaehler der Traceposition im Trace.

4. Deadlock-recovery.
 Motivierende Beispiele.
 Was genau wollen wir machen?
 Z.B. Deadlock wurde erkannt, wie kann dieser Deadlock aufgeloest werden.
 Spezielle Szenarien, RW-Mutex, Ungepufferte channels, "select", Kombinationen, ...
 Heuristische Methoden.