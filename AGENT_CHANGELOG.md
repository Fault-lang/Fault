# Fault Agent Changelog

A structured record of language and compiler changes, written for agents. Each entry answers:
- What changed (precisely, with old→new)
- What files/patterns need updating when this change is relevant
- What the linter should enforce

Entries are ordered newest-first. Date format: YYYY-MM-DD.

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
