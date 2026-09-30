# `inconsistent` Annotation — Implementation Spec

## Summary

Add a new annotation type `::inconsistent` to the Fault language that implements Belnap's
four-valued paraconsistent logic. The annotation isolates contradictory rules so they don't
cause the model to explode, tracks four truth states (true, false, both, neither), resolves
to a single boolean for use in the rest of the model, and reports the Belnap state in results.

This is the first of a family of annotation types (`::obligation`, `::permission`, ...) that
will share the same `variable::annotation{ }` syntax.

---

## Syntax

```
x::inconsistent{
    true-when  <evidence> [<weight>],
    false-when <evidence> [<weight>],
};
```

- any `ParameterCall` — a bare `IDENT` or arbitrarily deep dotted path (`x`, `x.y`, `x.y.z`, etc.); covers imported names like `orders.eligible::inconsistent`
- `::inconsistent` — the annotation type
- `true-when` — evidence that supports the claim being true
- `false-when` — evidence that supports the claim being false
- `<weight>` — optional bare numeric literal after the evidence (default: 1.0); no keyword needed

### Docstring

The docstring displayed in results comes from `x`'s string declaration if one exists, otherwise
the variable name is used. The annotation itself carries no docstring — this allows annotations
to be applied to variables defined in imported specs without requiring a re-declaration.

```
x = "this is a duck";  // string declaration is separate, optional
x::inconsistent{ ... };
```

### Trailing comma convention

Consistent with `stock{}` and `flow{}`: each rule ends with a comma, including the last.

### Full example

```
x           = "this is a duck";
hasFeathers = "has feathers";
quacks      = "quacks";
swims       = "swims";
isPlastic   = "is made of plastic";

x::inconsistent{
    true-when  hasFeathers 3,
    true-when  quacks,
    true-when  swims,
    false-when isPlastic 2,
};
```

### Cross-import example

```
// orders.fspec defines: eligible = "order is eligible"
// fulfillment.fspec defines: eligible = "fulfillment can proceed"

// system.fsystem
import orders;
import fulfillment;

orders.eligible::inconsistent{
    true-when  orders.eligible,
    false-when fulfillment.eligible,
};
```

The annotation is valid wherever the variable is in scope — in `.fspec` or `.fsystem` files,
including on names introduced via import.

---

## Evidence forms

Each `true-when` / `false-when` rule accepts one of:

1. **Bare identifier** — a boolean variable, string-labeled boolean, stock property, or constant
   ```
   true-when swims,
   false-when hasTentacles,
   ```

2. **Comparison expression** — any expression that resolves to `bool`
   ```
   true-when mortal.status == true,
   true-when body_temp > 36.5,
   ```

3. **Another `::inconsistent` annotation** — reference by name; its `_resolved` value is used
   ```
   true-when duck,   // duck has a ::inconsistent annotation
   ```

Non-boolean references are a **type error**.

---

## Where it lives

Valid in both `.fspec` and `.fsystem` files, at the top level alongside `assume` / `assert`.
Valid anywhere the annotated variable is in scope — including on names introduced via `import`.

---

## Belnap Semantics

Each `::inconsistent` annotation tracks two scores:

```
support_score = sum(weight_i * condition_i)   for all true-when rules
defeat_score  = sum(weight_j * condition_j)   for all false-when rules
```

These produce four truth states:

| `support_score > 0` | `defeat_score > 0` | Belnap   | Meaning |
|---|---|---|---|
| true  | false | `true`    | Established, uncontested |
| false | true  | `false`   | Rebutted, uncontested |
| true  | true  | `both`    | Contested — potential loophole |
| false | false | `neither` | No relevant evidence |

### Resolution for the rest of the model

The annotation exposes a single boolean `_resolved` used as evidence in other rules or annotations:

| Belnap    | `_resolved` |
|---|---|
| `true`    | `true` |
| `false`   | `false` |
| `neither` | `false` |
| `both`    | `support_score >= defeat_score` |

`both` defaults to `true` when support outweighs or equals defeat, `false` otherwise.

---

## assert / assume on inconsistent annotations

You can assert or assume a specific Belnap state using the four lowercase literals:

```
assert x = both;
assume x = true;
assert x = neither always;
```

These expand in SMT to:

| Literal   | SMT expansion |
|---|---|
| `true`    | `(and x_supported (not x_defeated))` |
| `false`   | `(and (not x_supported) x_defeated)` |
| `both`    | `(and x_supported x_defeated)` |
| `neither` | `(and (not x_supported) (not x_defeated))` |

All standard temporal modifiers (`always`, `eventually`, `eventually-always`, `nmt`, `nft`) apply.

The type checker distinguishes `both` / `neither` from `true` / `false` by the LHS type:
if `x` is typed `INCONSISTENT`, the RHS is parsed as a Belnap literal.

---

## SMT Encoding

Per annotation, introduce five synthetic definitions:

```
x_support_score  : Real
x_defeat_score   : Real
x_supported      : Bool   (= x_support_score > 0)
x_defeated       : Bool   (= x_defeat_score  > 0)
x_resolved       : Bool   (the value used as evidence in nested annotations)
```

Emitted as `define-fun` (not free variables — fully determined by the model):

```smt2
(define-fun x_support_score () Real
  (+ (* 3.0 (ite hasFeathers 1.0 0.0))
     (* 1.0 (ite quacks 1.0 0.0))
     (* 1.0 (ite swims 1.0 0.0))))

(define-fun x_defeat_score () Real
  (* 2.0 (ite isPlastic 1.0 0.0)))

(define-fun x_supported () Bool
  (> x_support_score 0.0))

(define-fun x_defeated () Bool
  (> x_defeat_score 0.0))

(define-fun x_resolved () Bool
  (ite (and x_supported x_defeated)
       (>= x_support_score x_defeat_score)
       (and x_supported (not x_defeated))))
```

Boolean conditions are cast to Real via `(ite cond 1.0 0.0)`.
Nested annotation references use `x_resolved` directly (already `Bool`, no cast needed).

---

## Results / Output

The TUI displays `::inconsistent` annotations separately from assert/assume violations.
The display label is the variable's string declaration if one exists, otherwise the variable name:

```
x  "this is a duck"
  Belnap state: both
  support: 4.0  (hasFeathers×3=3.0, quacks×1=1.0, swims×1=0.0 — 2 of 3 fired)
  defeat:  2.0  (isPlastic×2=2.0 — fired)
  resolved: true
```

- `both` is highlighted as a potential loophole
- `neither` is a warning (no evidence found — possible dead rule)

---

## Implementation Layers

| Layer | Change | Key files |
|---|---|---|
| 1. Grammar | Add `::inconsistent`, `true-when`, `false-when` tokens + rules; annotation target is any `ParameterCall` (bare `IDENT` or arbitrarily deep dotted path); weight is a bare `NUMBER` after evidence | `grammar/FaultLexer.g4`, `FaultParser.g4` |
| 2. AST | Add `AnnotationStatement` node (named for the family) with `Kind = "inconsistent"` | `ast/ast.go` |
| 3. Listener | `ExitAnnotationStatement`, `ExitTrueWhen`, `ExitFalseWhen` | `listener/listener_exit.go` |
| 4. Type checker | Evidence must resolve to `bool`; LHS of `INCONSISTENT` type in assert/assume accepts `both`/`neither`/`true`/`false` | `types/types.go` |
| 5. LLVM | Collect `AnnotationStatement` nodes; emit synthetic variable names into `RawInputs` | `llvm/compiler.go` |
| 6. SMT generation | Emit `define-fun` blocks for scores, supported, defeated, resolved; expand Belnap literals in assert/assume | `generator/asserts/asserts.go`, `generator/generator.go` |
| 7. Violation eval | Read `x_supported` / `x_defeated` from solver model; compute Belnap state; attach to annotation | `execute/violations.go` |
| 8. Output | Display Belnap state + score breakdown using variable's docstring | `tui/` |

---

## Type System Additions

- New AST node: `*ast.BelnapLiteral` with value one of `{"true", "false", "both", "neither"}`
  - `both` and `neither` are new; `true` and `false` reuse existing boolean token values
  - Disambiguated from plain booleans by context: only valid when LHS type is `INCONSISTENT`
- New type string: `"INCONSISTENT"` — the inferred type of a variable with a `::inconsistent` annotation
- Type rule: `INCONSISTENT`-typed variable resolves to `bool` (via `_resolved`) in boolean evidence contexts

---

## Constraints and Non-Goals

- **No non-boolean evidence** — type error at compile time
- **Weights are bare numeric literals only** — not `const` references or expressions
- **Weights must be non-negative** — compile-time error otherwise
- **No anonymous annotations** — must reference a named variable
- **Nesting is supported** — an `INCONSISTENT` variable used as evidence contributes its `_resolved` value
- **`::obligation` / `::permission`** — reserved for future issues; share the `::` syntax but are not implemented here
