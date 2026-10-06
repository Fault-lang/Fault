# Fault Agent Changelog

A structured record of language and compiler changes, written for agents. Each entry answers:
- What changed (precisely, with old→new)
- What files/patterns need updating when this change is relevant
- What the linter should enforce

Entries are ordered newest-first. Date format: YYYY-MM-DD.

---

## 2026-10-05

### `::inconsistent` annotation added (paraconsistent logic)

**Area:** grammar, AST, type checker, LLVM, SMT generator, execute, TUI
**Kind:** additive (new syntax)

Adds a `::inconsistent` annotation that attaches Belnap four-valued logic to an existing boolean variable. The variable type becomes `INCONSISTENT`, which behaves as `bool` everywhere except in `assert`/`assume` context where `both` and `neither` literals are additionally valid on the RHS.

**Syntax:**
```fault
target::inconsistent{
    true-when  <evidence> [<weight>],
    false-when <evidence> [<weight>],
};
```

- `target` — a bare `IDENT` or dotted path (`x`, `x.y`, `orders.eligible`); covers imported names
- `true-when` — evidence that supports the claim being true
- `false-when` — evidence that supports the claim being false
- `<weight>` — optional bare positive numeric literal after the evidence (default: 1.0); no keyword
- Trailing comma on every rule including the last (consistent with `stock{}` / `flow{}`)
- Semicolon after the closing brace

**New keywords:** `inconsistent`, `true-when`, `false-when`, `both`, `neither`

**Full example:**
```fault
x           = "this is a duck";
hasFeathers = "has feathers";
quacks      = "quacks";
isPlastic   = "is made of plastic";

x::inconsistent{
    true-when  hasFeathers 3,
    true-when  quacks,
    false-when isPlastic 2,
};

assert x = both;    // conflicting evidence?
assert x = true;    // resolves true by weight?
assume x = neither; // no evidence fires?
```

**Belnap states:**

| `supported` | `defeated` | State     |
|---|---|---|
| true  | false | `true`    |
| false | true  | `false`   |
| true  | true  | `both`    |
| false | false | `neither` |

**Boolean projection (used for `true`/`false` assert/assume):**
- `both` → resolves `true` if `support_score >= defeat_score`, else `false`
- `neither` → resolves `false`
- `true`/`false` → direct (no conflict)

**assert/assume RHS expansions per round N:**

| Literal   | SMT expansion |
|---|---|
| `true`    | `(not x_resolved_N)` (violation check) |
| `false`   | `x_resolved_N` (violation check) |
| `both`    | `(and x_supported_N x_defeated_N)` |
| `neither` | `(not x_supported_N) (not x_defeated_N)` (neither state) |

All temporal modifiers (`always`, `eventually`, `eventually-always`, `nmt`, `nft`) apply and combine across rounds the same way as regular `assert`/`assume`.

**Nesting:** An `INCONSISTENT`-typed variable may be evidence in another annotation; its `_resolved_N` value is used as the bool condition. Annotation dependency graph must be acyclic (cycles are a compile-time error). Annotations are emitted in topological order.

**Constraints:**
- Both `true-when` and `false-when` sides must each have at least one rule
- Weights must be bare strictly-positive numeric literals (not `const` references)
- Annotation target may not appear as evidence in its own block (self-reference is a compile-time error)
- Later redeclaration overrides with a compiler warning

**Valid in both `.fspec` and `.fsystem` files**, at the top level alongside `assume`/`assert`.

**SMT encoding:** five `declare-fun` + `assert` entries per annotation per round: `x_supported_N`, `x_defeated_N`, `x_support_score_N`, `x_defeat_score_N`, `x_resolved_N`.

**Output:** TUI displays per-round Belnap state and score summary using the variable's string declaration as the label.

---

## 2026-08-30

### Fix: excluded+redeclared stock properties sever parent assume/assert propagation

**Area:** compiler (`llvm/compiler.go`)
**Kind:** bug fix

When a child stock excludes a parent field and immediately redeclares it as a fresh solvable (`exclude age, age,`), the parent spec's assumes and asserts on that field no longer propagate to the child's instances. Previously they did, making it impossible to replace a property from a library stock with clean semantics.

**Old behavior:** `assume entity.age <= 6` in a base spec would appear in the child's SMT output even after `exclude age, age,`.

**New behavior:** The parent rule is dropped. Only rules written in the child spec (e.g. `assume person.age <= 8`) apply to the redeclared property.

**How it works:**
- `compileStruct` populates `stockSeveredProps map[string]map[string]bool` — keyed by stock name (`"specName_stockName"`), value is the set of severed field names.
- A field is severed when it appears in both `stock.Excludes` and `stock.Pairs` (i.e., excluded then redeclared).
- `bfsFetchInstances` replaces the inline BFS in `convertAssertVariables`. When traversing the inheritance graph to resolve assert/assume targets, it skips any non-origin node that has severed the requested field, stopping propagation through that stock and its descendants.
- The origin node is never skipped — `assume person.age <= 8` (direct reference to the child's own field) still resolves correctly.

**Applies to both `assume` and `assert`:** both go through `convertAssertVariables` and `assertionRefsActive`, both of which now use `bfsFetchInstances`.

**Pattern:** to cleanly replace a library stock's property with your own rules:
```fault
def child = stock{
    extends base_lib.entity,
    exclude age,
    age,            // fresh solvable — no parent rules apply
};

assume child.age <= 8;  // only this rule applies
```

---

## 2026-08-30

### Named import alias required for cross-spec stock extension

**Area:** syntax, imports
**Kind:** clarification (not a code change)

`import "file.fspec"` without an alias uses the **declared spec name** (the `spec foo;` at the top of the file) as the namespace — not the filename. When extending a stock from an imported spec, use a named alias that matches, or be explicit:

**Correct:**
```fault
import baselib "02a_base_lib.fspec";   // spec inside declares: spec baselib;

def child = stock{
    extends baselib.entity,
    ...
};
```

**Wrong (produces resolution error):**
```fault
import "02a_base_lib.fspec";   // namespace is "baselib", not "02a_base_lib"

def child = stock{
    extends 02a_base_lib.entity,  // fails — no spec named "02a_base_lib"
};
```

**Rule:** always use `import <alias> "file.fspec"` and match `<alias>` to the `spec <name>;` declaration inside the file.

---

## 2026-08-17

### BREAKING: `func{}` renamed to `sfunc{}` in state chart bodies

**Area:** grammar, syntax
**Kind:** breaking rename

State functions (inside `component ... = states{}` blocks) now use the keyword `sfunc` instead of `func`. Flow functions (inside `def ... = flow{}` blocks) still use `func`.

**Old:**
```
component foo = states{
    idle: func{
        advance(this.running);
    },
};
```

**New:**
```
component foo = states{
    idle: sfunc{
        advance(this.running);
    },
};
```

**How to distinguish:** `sfunc{}` bodies may contain `stay()`, `advance()`, `leave()`. `func{}` bodies contain arithmetic expressions and stock/flow mutations.

**Files to check when updating `.fsystem` files:**
- Search for `func{` inside any `states{` block and replace with `sfunc{`
- Flow `func{` (inside `flow{` blocks) is unchanged

**Grammar rule:** `stateLit : 'sfunc' stateBlock` (was `'func' stateBlock`)
**Lexer token:** `SFUNC: 'sfunc'` added to `FaultLexer.g4`

**Linter should:** reject `func{` inside a `states{}` block and suggest `sfunc{}`

---

## 2026-08-07

### `run{}` replaces `for N run{}`; `init{}` clause added

**Area:** grammar, syntax
**Kind:** breaking rename

The `for N run{}` construction was removed. Use `run{}` directly. An optional `init{}` clause inside `run{}` handles initializations that previously had to live in the run loop.

**Old:**
```
for 5 run {
    foo.initial;
}
```

**New:**
```
run {
    foo.initial;
}
```

With initialization:
```
run {
    init{
        foo.x = 10;
    }
    foo.initial;
}
```

**Linter should:** reject `for N run{` syntax and suggest the new form

---

## 2026-08-07

### `assume` replaces `produce`

**Area:** grammar, syntax
**Kind:** breaking rename

The keyword `produce` was replaced by `assume` in `unfunc{}` blocks.

**Old:** `produce: expr`
**New:** `assume: expr`

**Linter should:** reject `produce:` inside `unfunc{}` and suggest `assume:`

---

## 2026-08-07

### `available` added as a temporal clause

**Area:** grammar
**Kind:** additive

`available` is now a valid temporal keyword alongside `always`, `eventually`, `eventually-always`. Used in assert blocks to check whether a variable is reachable.

---

## 2026-06-23

### `start{}` blocks replaced by `run{}`

**Area:** grammar, syntax
**Kind:** breaking rename

`start{}` was removed and replaced by `run{}`.

**Linter should:** reject `start{` and suggest `run{`

---

## 2026-06-05

### `unfunc{}` added

**Area:** grammar
**Kind:** additive

New block type for program synthesis. Syntax:

```
unfunc foo {
    requires: expr,
    emits: expr,
    assume: expr,
}
```

Valid inside `flow{}` definitions alongside regular `func{}` properties.

---

## Notes for agents

- `.fspec` files define stocks and flows (use `func{}`)
- `.fsystem` files define components and run blocks (use `sfunc{}` for state bodies)
- The linter runs through the preprocess stage only — it does not invoke the SMT solver
- Grammar source of truth: `grammar/FaultParser.g4` and `grammar/FaultLexer.g4`
- After any grammar change: run `make java && make golang` to regenerate the ANTLR parser
