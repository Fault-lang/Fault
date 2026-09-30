# `::inconsistent` Annotation — Implementation Spec

## Summary

Add a new annotation type `::inconsistent` to the Fault language that implements Belnap's
four-valued paraconsistent logic. The annotation isolates contradictory rules so they don't
cause the model to explode, tracks four truth states (`true`, `false`, `both`, `neither`),
resolves to a single boolean per round for use in the rest of the model, and reports the
Belnap state in results.

This is the first of a family of annotation types (`::obligation`, `::permission`, ...) that
will share the same `variable::annotation{ }` syntax.

---

## Syntax

```
target::inconsistent{
    true-when  <evidence> [<weight>],
    false-when <evidence> [<weight>],
};
```

- `target` — any `ParameterCall`: a bare `IDENT` or arbitrarily deep dotted path (`x`, `x.y`,
  `x.y.z.w`, etc.); covers imported names like `orders.eligible::inconsistent`
- `::inconsistent` — the annotation type keyword
- `true-when` — evidence that supports the claim being true
- `false-when` — evidence that supports the claim being false
- `<weight>` — optional bare numeric literal after the evidence (default: 1.0); no keyword
- Trailing comma on every rule including the last, consistent with `stock{}` / `flow{}`
- Semicolon after the closing brace, consistent with other top-level statements

### Docstring

The docstring displayed in results comes from `target`'s string declaration if one exists;
otherwise the variable name is used. The annotation itself carries no docstring — this allows
annotations to be applied to variables defined in imported specs without requiring a
re-declaration.

```
x = "this is a duck";  // optional — string declaration is separate from the annotation
x::inconsistent{ ... };
```

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
// fulfillment.fspec defines its own conditions

// system.fsystem
import orders;
import fulfillment;

orders.eligible::inconsistent{
    true-when  orders.hasValidPayment,
    false-when fulfillment.outOfStock,
};
```

---

## Evidence forms

Each `true-when` / `false-when` rule accepts one of:

1. **Bare `ParameterCall`** — a boolean variable, string-labeled boolean, stock property, or constant
   ```
   true-when swims,
   false-when hasTentacles,
   true-when orders.hasValidPayment,
   ```

2. **Comparison expression** — any expression that resolves to `bool`
   ```
   true-when mortal.status == true,
   true-when body_temp > 36.5,
   ```

3. **Another `::inconsistent` annotation target** — reference by name; its `_resolved` value
   is used as the condition
   ```
   true-when duck,   // duck has a ::inconsistent annotation; duck_resolved is used
   ```

Non-boolean references are a **compile-time type error**.
The annotation target itself may not appear as evidence within its own block — **compile-time error**.

---

## Where it lives

Valid in both `.fspec` and `.fsystem` files, at the top level alongside `assume` / `assert`.
Valid wherever the annotated variable is in scope, including names introduced via `import`.

---

## Keywords

`both` and `neither` are new **reserved keywords** added to the lexer.
`true` and `false` are reused as-is. All four are lowercase.
`true-when` and `false-when` are new reserved keywords (hyphenated, like `eventually-always`).
`inconsistent` is a new reserved keyword.

---

## Belnap Semantics

Each `::inconsistent` annotation is evaluated **per round** (SSA-versioned). For round `N`:

```
support_score_N = sum(weight_i * condition_i_N)   for all true-when rules
defeat_score_N  = sum(weight_j * condition_j_N)   for all false-when rules
```

These produce four truth states per round:

| `support_score_N > 0` | `defeat_score_N > 0` | Belnap    | Meaning |
|---|---|---|---|
| true  | false | `true`    | Established, uncontested |
| false | true  | `false`   | Rebutted, uncontested |
| true  | true  | `both`    | Contested — potential loophole |
| false | false | `neither` | No relevant evidence |

### Resolution for the rest of the model

The annotation exposes a boolean `resolved_N` per round, used as evidence in other rules
and annotations:

| Belnap    | `resolved_N` |
|---|---|
| `true`    | `true` |
| `false`   | `false` |
| `neither` | `false` |
| `both`    | `support_score_N >= defeat_score_N` |

`both` defaults to `true` when support outweighs or equals defeat, `false` otherwise.

---

## Empty block rules

- A block with **no rules at all** is a compile-time error.
- A block with **only `true-when` rules** is a compile-time error.
- A block with **only `false-when` rules** is a compile-time error.
- Both sides must have at least one rule.

---

## `assert` / `assume` on inconsistent annotations

You can assert or assume a specific Belnap state using the four literals:

```
assert x = both;
assume x = true;
assert x = neither always;
```

These expand in SMT (for a given SSA round N) to:

| Literal   | SMT expansion |
|---|---|
| `true`    | `(and x_supported_N (not x_defeated_N))` |
| `false`   | `(and (not x_supported_N) x_defeated_N)` |
| `both`    | `(and x_supported_N x_defeated_N)` |
| `neither` | `(and (not x_supported_N) (not x_defeated_N))` |

All standard temporal modifiers (`always`, `eventually`, `eventually-always`, `nmt`, `nft`) apply,
combining across rounds the same way as regular `assert`/`assume`.

The type checker identifies the LHS as `INCONSISTENT`-typed and accepts `both`/`neither` on the
RHS only in that context. On any other LHS type, `both` and `neither` are not valid in this
position.

---

## SMT Encoding

Per annotation per round `N`, emit five `define-fun` entries. These are emitted in the **Init
block** so they always precede any `assert` / `assume` that references them.

```smt2
(define-fun x_support_score_N () Real
  (+ (* 3.0 (ite hasFeathers_N 1.0 0.0))
     (* 1.0 (ite quacks_N 1.0 0.0))
     (* 1.0 (ite swims_N 1.0 0.0))))

(define-fun x_defeat_score_N () Real
  (* 2.0 (ite isPlastic_N 1.0 0.0)))

(define-fun x_supported_N () Bool
  (> x_support_score_N 0.0))

(define-fun x_defeated_N () Bool
  (> x_defeat_score_N 0.0))

(define-fun x_resolved_N () Bool
  (ite (and x_supported_N x_defeated_N)
       (>= x_support_score_N x_defeat_score_N)
       (and x_supported_N (not x_defeated_N))))
```

Boolean conditions are cast to Real via `(ite cond 1.0 0.0)`.
Nested annotation references use `duck_resolved_N` directly — already `Bool`, no cast needed.
For a `ParameterCall` target like `orders.eligible`, join path segments with `_`:
`orders_eligible_support_score_N`, etc.

---

## Results / Output

The TUI displays `::inconsistent` annotations separately from assert/assume violations.
The display label is the variable's string declaration if one exists, otherwise the variable name.
Per-round Belnap states are shown, plus a summary:

```
x  "this is a duck"
  round 0 — both    (support: 4.0, defeat: 2.0, resolved: true)
  round 1 — true    (support: 3.0, defeat: 0.0, resolved: true)
  round 2 — neither (support: 0.0, defeat: 0.0, resolved: false)
```

- `both` is highlighted as a potential loophole
- `neither` is a warning (no evidence found in that round — possible dead rule)

---

## Implementation Layers

| Layer | Change | Key files |
|---|---|---|
| 1. Grammar | Add `inconsistent`, `true-when`, `false-when`, `both`, `neither` tokens; annotation rule with `ParameterCall::inconsistent{ }` | `grammar/FaultLexer.g4`, `FaultParser.g4` |
| 2. AST | Add `AnnotationStatement` node (Kind field for future `::obligation` etc.) with Target, TrueWhen, FalseWhen | `ast/ast.go` |
| 3. Listener | `ExitAnnotationStatement`, `ExitTrueWhen`, `ExitFalseWhen`; validate target not in own evidence | `listener/listener_exit.go` |
| 4. Type checker | Evidence must resolve to `bool`; `INCONSISTENT`-typed LHS in assert/assume accepts `both`/`neither`/`true`/`false` | `types/types.go` |
| 5. LLVM | Collect `AnnotationStatement` nodes into `RawInputs.Annotations []AnnotationStatement` | `llvm/compiler.go` |
| 6. SMT generation | For each annotation and each round N, emit five `define-fun` entries in the Init block; expand Belnap literals in assert/assume | `generator/asserts/asserts.go`, `generator/generator.go` |
| 7. Violation eval | Read `x_supported_N` / `x_defeated_N` from solver model per round; compute Belnap state; attach to annotation | `execute/violations.go` |
| 8. Output | Display per-round Belnap state + score summary using variable's docstring | `tui/` |

---

## Type System Additions

- New AST node: `*ast.BelnapLiteral` with value one of `{"true", "false", "both", "neither"}`
  - `both` and `neither` are new reserved keywords parsed directly as `BelnapLiteral`
  - `true` and `false` are promoted to `BelnapLiteral` when LHS type is `INCONSISTENT`
- New type string: `"INCONSISTENT"` — the inferred type of a variable with a `::inconsistent`
  annotation attached
- Type rule: an `INCONSISTENT`-typed variable resolves to `bool` (via `_resolved_N`) when used
  as evidence inside another annotation's `true-when` / `false-when` rule

---

## Constraints and Non-Goals

- **No non-boolean evidence** — compile-time type error
- **Weights are bare numeric literals only** — not `const` references or expressions
- **Weights must be non-negative** — compile-time error otherwise
- **Both sides required** — a block with only `true-when` or only `false-when` rules is a compile-time error
- **No self-reference** — annotation target may not appear as evidence in its own block
- **No anonymous annotations** — must reference a named variable/selector
- **Nesting is supported** — an `INCONSISTENT` variable used as evidence contributes its `_resolved_N` value
- **`::obligation` / `::permission`** — reserved for future issues; share the `::` syntax but are not implemented here
