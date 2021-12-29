OVERVIEW: llvm .bc -> .bc modular optimizer and analysis printer

USAGE: opt [options] <input bitcode file>

OPTIONS:

Color Options:

  --color                                                               - Use colors in output (default=autodetect)

General options:

  --O0                                                                  - Optimization level 0. Similar to clang -O0
  --O1                                                                  - Optimization level 1. Similar to clang -O1
  --O2                                                                  - Optimization level 2. Similar to clang -O2
  --O3                                                                  - Optimization level 3. Similar to clang -O3
  --Os                                                                  - Like -O2 with extra optimizations for size. Similar to clang -Os
  --Oz                                                                  - Like -Os but reduces code size further. Similar to clang -Oz
  -S                                                                    - Write output as LLVM assembly
  --aarch64-neon-syntax=<value>                                         - Choose style of NEON code to emit from AArch64 backend:
    =generic                                                            -   Emit generic NEON assembly
    =apple                                                              -   Emit Apple-style NEON assembly
  --abort-on-max-devirt-iterations-reached                              - Abort when the max iterations for devirtualization CGSCC repeat pass is reached
  --addrsig                                                             - Emit an address-significance table
  --amdgpu-bypass-slow-div                                              - Skip 64-bit divide for dynamic 32-bit values
  --amdgpu-disable-loop-alignment                                       - Do not align and prefetch loops
  --amdgpu-disable-power-sched                                          - Disable scheduling to minimize mAI power bursts
  --amdgpu-dpp-combine                                                  - Enable DPP combiner
  --amdgpu-dump-hsa-metadata                                            - Dump AMDGPU HSA Metadata
  --amdgpu-enable-flat-scratch                                          - Use flat scratch instructions
  --amdgpu-enable-merge-m0                                              - Merge and hoist M0 initializations
  --amdgpu-promote-alloca-to-vector-limit=<uint>                        - Maximum byte size to consider promote alloca to vector
  --amdgpu-reserve-vgpr-for-sgpr-spill                                  - Allocates one VGPR for future SGPR Spill
  --amdgpu-sdwa-peephole                                                - Enable SDWA peepholer
  --amdgpu-use-aa-in-codegen                                            - Enable the use of AA during codegen.
  --amdgpu-verify-hsa-metadata                                          - Verify AMDGPU HSA Metadata
  --amdgpu-vgpr-index-mode                                              - Use GPR indexing mode instead of movrel for vector indexing
  --analyze                                                             - Only perform analysis, no optimization
  --arm-add-build-attributes                                            - 
  --arm-implicit-it=<value>                                             - Allow conditional instructions outdside of an IT block
    =always                                                             -   Accept in both ISAs, emit implicit ITs in Thumb
    =never                                                              -   Warn in ARM, reject in Thumb
    =arm                                                                -   Accept in ARM, reject in Thumb
    =thumb                                                              -   Warn in ARM, emit implicit ITs in Thumb
  --asm-show-inst                                                       - Emit internal instruction representation to assembly file
  --atomic-counter-update-promoted                                      - Do counter update using atomic fetch add  for promoted counters only
  --atomic-first-counter                                                - Use atomic fetch add for first counter in a function (usually the entry counter)
  --basic-block-sections=<all | <function list (file)> | labels | none> - Emit basic blocks into separate sections
  --bounds-checking-single-trap                                         - Use one trap block per function
  --cfg-hide-deoptimize-paths                                           - 
  --cfg-hide-unreachable-paths                                          - 
  --code-model=<value>                                                  - Choose code model
    =tiny                                                               -   Tiny code model
    =small                                                              -   Small code model
    =kernel                                                             -   Kernel code model
    =medium                                                             -   Medium code model
    =large                                                              -   Large code model
  --codegen-opt-level=<uint>                                            - Override optimization level for codegen hooks
  --cost-kind=<value>                                                   - Target cost kind
    =throughput                                                         -   Reciprocal throughput
    =latency                                                            -   Instruction latency
    =code-size                                                          -   Code size
    =size-latency                                                       -   Code size and latency
  --cvp-dont-add-nowrap-flags                                           - 
  --data-layout=<layout-string>                                         - data layout string to use
  --data-sections                                                       - Emit data into separate sections
  --debug-entry-values                                                  - Enable debug info for the debug entry values.
  --debugger-tune=<value>                                               - Tune debug info for a particular debugger
    =gdb                                                                -   gdb
    =lldb                                                               -   lldb
    =sce                                                                -   SCE targets (e.g. PS4)
  --debugify-each                                                       - Start each pass with debugify and end it with check-debugify
  --debugify-export=<filename>                                          - Export per-pass debugify statistics to this file
  --debugify-level=<value>                                              - Kind of debug info to add
    =locations                                                          -   Locations only
    =location+variables                                                 -   Locations and Variables
  --debugify-quiet                                                      - Suppress verbose debugify output
  --denormal-fp-math=<value>                                            - Select which denormal numbers the code is permitted to require
    =ieee                                                               -   IEEE 754 denormal numbers
    =preserve-sign                                                      -   the sign of a  flushed-to-zero number is preserved in the sign of 0
    =positive-zero                                                      -   denormals are flushed to positive zero
  --denormal-fp-math-f32=<value>                                        - Select which denormal numbers the code is permitted to require for float
    =ieee                                                               -   IEEE 754 denormal numbers
    =preserve-sign                                                      -   the sign of a  flushed-to-zero number is preserved in the sign of 0
    =positive-zero                                                      -   denormals are flushed to positive zero
  --disable-builtin=<string>                                            - Disable specific target library builtin function
  --disable-debug-info-type-map                                         - Don't use a uniquing type map for debug info
  --disable-inlining                                                    - Do not run the inliner pass
  --disable-loop-unrolling                                              - Disable loop unrolling in all relevant passes
  --disable-opt                                                         - Do not run any optimization passes
  --disable-promote-alloca-to-lds                                       - Disable promote alloca to LDS
  --disable-promote-alloca-to-vector                                    - Disable promote alloca to vector
  --disable-simplify-libcalls                                           - Disable simplify-libcalls
  --disable-tail-calls                                                  - Never emit tail calls
  --do-counter-promotion                                                - Do counter register promotion
  --dot-cfg-mssa=<file name for generated dot file>                     - file name for generated dot file
  --dwarf-version=<int>                                                 - Dwarf version
  --dwarf64                                                             - Generate debugging info in the 64-bit DWARF format
  --emit-call-site-info                                                 - Emit call site debug information, if debug information is enabled.
  --emscripten-cxx-exceptions-allowed=<string>                          - The list of function names in which Emscripten-style exception handling is enabled (see emscripten EMSCRIPTEN_CATCHING_ALLOWED options)
  --emulated-tls                                                        - Use emulated TLS model
  --enable-cse-in-irtranslator                                          - Should enable CSE in irtranslator
  --enable-cse-in-legalizer                                             - Should enable CSE in Legalizer
  --enable-debugify                                                     - Start the pipeline with debugify and end it with check-debugify
  --enable-emscripten-cxx-exceptions                                    - WebAssembly Emscripten-style exception handling
  --enable-emscripten-sjlj                                              - WebAssembly Emscripten-style setjmp/longjmp handling
  --enable-gvn-hoist                                                    - Enable the GVN hoisting pass (default = off)
  --enable-gvn-memdep                                                   - 
  --enable-gvn-sink                                                     - Enable the GVN sinking pass (default = off)
  --enable-load-in-loop-pre                                             - 
  --enable-load-pre                                                     - 
  --enable-loop-simplifycfg-term-folding                                - 
  --enable-name-compression                                             - Enable name/filename string compression
  --enable-new-pm                                                       - Enable the new pass manager
  --enable-no-infs-fp-math                                              - Enable FP math optimizations that assume no +-Infs
  --enable-no-nans-fp-math                                              - Enable FP math optimizations that assume no NaNs
  --enable-no-signed-zeros-fp-math                                      - Enable FP math optimizations that assume the sign of 0 is insignificant
  --enable-no-trapping-fp-math                                          - Enable setting the FP exceptions build attribute not to use exceptions
  --enable-split-backedge-in-load-pre                                   - 
  --enable-unsafe-fp-math                                               - Enable optimizations that may decrease FP precision
  --exception-model=<value>                                             - exception model
    =default                                                            -   default exception handling model
    =dwarf                                                              -   DWARF-like CFI based exception handling
    =sjlj                                                               -   SjLj exception handling
    =arm                                                                -   ARM EHABI exceptions
    =wineh                                                              -   Windows exception model
    =wasm                                                               -   WebAssembly exception handling
  --experimental-debug-variable-locations                               - Use experimental new value-tracking variable locations
  -f                                                                    - Enable binary output on terminals
  --fatal-warnings                                                      - Treat warnings as errors
  --filetype=<value>                                                    - Choose a file type (not all types are supported by all targets):
    =asm                                                                -   Emit an assembly ('.s') file
    =obj                                                                -   Emit a native object ('.o') file
    =null                                                               -   Emit nothing, for performance testing
  --float-abi=<value>                                                   - Choose float ABI type
    =default                                                            -   Target default float ABI type
    =soft                                                               -   Soft float ABI (implied by -soft-float)
    =hard                                                               -   Hard float ABI (uses FP registers)
  --force-dwarf-frame-section                                           - Always emit a debug frame section.
  --fp-contract=<value>                                                 - Enable aggressive formation of fused FP ops
    =fast                                                               -   Fuse FP ops whenever profitable
    =on                                                                 -   Only fuse 'blessed' FP ops.
    =off                                                                -   Only fuse FP ops when the result won't be affected.
  --frame-pointer=<value>                                               - Specify frame pointer elimination optimization
    =all                                                                -   Disable frame pointer elimination
    =non-leaf                                                           -   Disable frame pointer elimination for non-leaf frame
    =none                                                               -   Enable frame pointer elimination
  --function-sections                                                   - Emit functions into separate sections
  --gpsize=<uint>                                                       - Global Pointer Addressing Size.  The default size is 8.
  --hash-based-counter-split                                            - Rename counter variable of a comdat function based on cfg hash
  --hot-cold-split                                                      - Enable hot-cold splitting pass
  --ignore-xcoff-visibility                                             - Not emit the visibility attribute for asm in AIX OS or give all symbols 'unspecified' visibility in XCOFF object file
  --import-all-index                                                    - Import all external functions in index.
  --incremental-linker-compatible                                       - When used with filetype=obj, emit an object file which can be used with an incremental linker
  --instcombine-code-sinking                                            - Enable code sinking
  --instcombine-guard-widening-window=<uint>                            - How wide an instruction window to bypass looking for another guard
  --instcombine-max-iterations=<uint>                                   - Limit the maximum number of instruction combining iterations
  --instcombine-max-num-phis=<uint>                                     - Maximum number phis to handle in intptr/ptrint folding
  --instcombine-maxarray-size=<uint>                                    - Maximum array size considered when doing a combine
  --instcombine-negator-enabled                                         - Should we attempt to sink negations?
  --instcombine-negator-max-depth=<uint>                                - What is the maximal lookup depth when trying to check for viability of negation sinking.
  --instcombine-unsafe-select-transform                                 - Enable poison-unsafe select to and/or transform
  --instrprof-atomic-counter-update-all                                 - Make all profile counter updates atomic (for testing only)
  --internalize-public-api-file=<filename>                              - A file containing list of symbol names to preserve
  --internalize-public-api-list=<list>                                  - A list of symbol names to preserve
  --iterative-counter-promotion                                         - Allow counter promotion across the whole loop nest.
  --load=<pluginfilename>                                               - Load the specified plugin
  --load-pass-plugin=<string>                                           - Load passes from plugin library
  Optimizations available:
      --aa                                                                 - Function Alias Analysis Results
      --aa-eval                                                            - Exhaustive Alias Analysis Precision Evaluator
      --aarch64-a57-fp-load-balancing                                      - AArch64 A57 FP Load-Balancing
      --aarch64-branch-targets                                             - AArch64 Branch Targets
      --aarch64-ccmp                                                       - AArch64 CCMP Pass
      --aarch64-collect-loh                                                - AArch64 Collect Linker Optimization Hint (LOH)
      --aarch64-condopt                                                    - AArch64 CondOpt Pass
      --aarch64-copyelim                                                   - AArch64 redundant copy elimination pass
      --aarch64-dead-defs                                                  - AArch64 Dead register definitions
      --aarch64-expand-pseudo                                              - AArch64 pseudo instruction expansion pass
      --aarch64-falkor-hwpf-fix                                            - Falkor HW Prefetch Fix
      --aarch64-falkor-hwpf-fix-late                                       - Falkor HW Prefetch Fix Late Phase
      --aarch64-fix-cortex-a53-835769-pass                                 - AArch64 fix for A53 erratum 835769
      --aarch64-jump-tables                                                - AArch64 compress jump tables pass
      --aarch64-ldst-opt                                                   - AArch64 load / store optimization pass
      --aarch64-local-dynamic-tls-cleanup                                  - AArch64 Local Dynamic TLS Access Clean-up
      --aarch64-post-select-optimize                                       - Optimize AArch64 selected instructions
      --aarch64-postlegalizer-combiner                                     - Combine AArch64 MachineInstrs after legalization
      --aarch64-postlegalizer-lowering                                     - Lower AArch64 MachineInstrs after legalization
      --aarch64-prelegalizer-combiner                                      - Combine AArch64 machine instrs before legalization
      --aarch64-promote-const                                              - AArch64 Promote Constant Pass
      --aarch64-simd-scalar                                                - AdvSIMD Scalar Operation Optimization
      --aarch64-simdinstr-opt                                              - AArch64 SIMD instructions optimization pass
      --aarch64-sls-hardening                                              - AArch64 sls hardening pass
      --aarch64-speculation-hardening                                      - AArch64 speculation hardening pass
      --aarch64-stack-tagging                                              - AArch64 Stack Tagging
      --aarch64-stack-tagging-pre-ra                                       - AArch64 Stack Tagging PreRA Pass
      --aarch64-stp-suppress                                               - AArch64 Store Pair Suppression
      --aarch64-sve-intrinsic-opts                                         - SVE intrinsics optimizations
      --adce                                                               - Aggressive Dead Code Elimination
      --add-discriminators                                                 - Add DWARF path discriminators
      --aggressive-instcombine                                             - Combine pattern based expressions
      --alignment-from-assumptions                                         - Alignment from assumptions
      --alloca-hoisting                                                    - Hoisting alloca instructions in non-entry blocks to the entry block
      --always-inline                                                      - Inliner for always_inline functions
      --amdgpu-aa                                                          - AMDGPU Address space based Alias Analysis
      --amdgpu-aa-wrapper                                                  - AMDGPU Address space based Alias Analysis Wrapper
      --amdgpu-always-inline                                               - AMDGPU Inline All Functions
      --amdgpu-annotate-kernel-features                                    - Add AMDGPU function attributes
      --amdgpu-annotate-uniform                                            - Add AMDGPU uniform metadata
      --amdgpu-argument-reg-usage-info                                     - Argument Register Usage Information Storage
      --amdgpu-atomic-optimizer                                            - AMDGPU atomic optimizations
      --amdgpu-codegenprepare                                              - AMDGPU IR optimizations
      --amdgpu-fix-function-bitcasts                                       - Fix function bitcasts for AMDGPU
      --amdgpu-isel                                                        - AMDGPU DAG->DAG Pattern Instruction Selection
      --amdgpu-late-codegenprepare                                         - AMDGPU IR late optimizations
      --amdgpu-lower-enqueued-block                                        - Lower OpenCL enqueued blocks
      --amdgpu-lower-intrinsics                                            - Lower intrinsics
      --amdgpu-lower-kernel-arguments                                      - AMDGPU Lower Kernel Arguments
      --amdgpu-lower-kernel-attributes                                     - AMDGPU IR optimizations
      --amdgpu-nsa-reassign                                                - GCN NSA Reassign
      --amdgpu-perf-hint                                                   - Analysis if a function is memory bound
      --amdgpu-postlegalizer-combiner                                      - Combine AMDGPU machine instrs after legalization
      --amdgpu-prelegalizer-combiner                                       - Combine AMDGPU machine instrs before legalization
      --amdgpu-printf-runtime-binding                                      - AMDGPU Printf lowering
      --amdgpu-promote-alloca                                              - AMDGPU promote alloca to vector or LDS
      --amdgpu-promote-alloca-to-vector                                    - AMDGPU promote alloca to vector
      --amdgpu-propagate-attributes-early                                  - Early propagate attributes from kernels to functions
      --amdgpu-propagate-attributes-late                                   - Late propagate attributes from kernels to functions
      --amdgpu-regbanks-reassign                                           - GCN RegBank Reassign
      --amdgpu-rewrite-out-arguments                                       - AMDGPU Rewrite Out Arguments
      --amdgpu-simplifylib                                                 - Simplify well-known AMD library calls
      --amdgpu-unify-divergent-exit-nodes                                  - Unify divergent function exit nodes
      --amdgpu-unify-metadata                                              - Unify multiple OpenCL metadata due to linking
      --amdgpu-usenative                                                   - Replace builtin math calls with that native versions.
      --amode-opt                                                          - Optimize addressing mode
      --annotation-remarks                                                 - Annotation Remarks
      --annotation2metadata                                                - Annotation2Metadata
      --argpromotion                                                       - Promote 'by reference' arguments to scalars
      --arm-block-placement                                                - ARM block placement
      --arm-cp-islands                                                     - ARM constant island placement and branch shortening pass
      --arm-execution-domain-fix                                           - ARM Execution Domain Fix
      --arm-ldst-opt                                                       - ARM load / store optimization pass
      --arm-low-overhead-loops                                             - ARM Low Overhead Loops pass
      --arm-mve-gather-scatter-lowering                                    - MVE gather/scattering lowering pass
      --arm-mve-vpt                                                        - ARM MVE VPT block pass
      --arm-mve-vpt-opts                                                   - ARM MVE TailPred and VPT Optimisations pass
      --arm-parallel-dsp                                                   - Transform functions to use DSP intrinsics
      --arm-prera-ldst-opt                                                 - ARM pre- register allocation load / store optimization pass
      --arm-pseudo                                                         - ARM pseudo instruction expansion pass
      --arm-sls-hardening                                                  - ARM sls hardening pass
      --asan                                                               - AddressSanitizer: detects use-after-free and out-of-bounds bugs.
      --asan-globals-md                                                    - Read metadata to mark which globals should be instrumented when running ASan.
      --asan-module                                                        - AddressSanitizer: detects use-after-free and out-of-bounds bugs.ModulePass
      --assume-builder                                                     - Assume Builder
      --assume-simplify                                                    - Assume Simplify
      --assumption-cache-tracker                                           - Assumption Cache Tracker
      --atomic-expand                                                      - Expand Atomic instructions
      --attributor                                                         - Deduce and propagate attributes
      --attributor-cgscc                                                   - Deduce and propagate attributes (CGSCC pass)
      --avr-expand-pseudo                                                  - AVR pseudo instruction expansion pass
      --avr-relax-mem                                                      - AVR memory operation relaxation pass
      --barrier                                                            - A No-Op Barrier Pass
      --basic-aa                                                           - Basic Alias Analysis (stateless AA impl)
      --basiccg                                                            - CallGraph Construction
      --bdce                                                               - Bit-Tracking Dead Code Elimination
      --block-freq                                                         - Block Frequency Analysis
      --bounds-checking                                                    - Run-time bounds checking
      --bpf-abstract-member-access                                         - BPF Abstract Member Access
      --bpf-adjust-opt                                                     - BPF Adjust Optimization
      --bpf-check-and-opt-ir                                               - BPF Check And Adjust IR
      --bpf-mi-trunc-elim                                                  - BPF MachineSSA Peephole Optimization For TRUNC Eliminate
      --bpf-mi-zext-elim                                                   - BPF MachineSSA Peephole Optimization For ZEXT Eliminate
      --bpf-preserve-di-type                                               - BPF Preserve Debuginfo Type
      --branch-prob                                                        - Branch Probability Analysis
      --break-crit-edges                                                   - Break critical edges in CFG
      --called-value-propagation                                           - Called Value Propagation
      --callsite-splitting                                                 - Call-site splitting
      --canon-freeze                                                       - Canonicalize Freeze Instructions in Loops
      --canonicalize-aliases                                               - Canonicalize aliases
      --cfl-anders-aa                                                      - Inclusion-Based CFL Alias Analysis
      --cfl-steens-aa                                                      - Unification-Based CFL Alias Analysis
      --cg-profile                                                         - Call Graph Profile
      --check-debugify                                                     - Check debug info from -debugify
      --check-debugify-function                                            - Check debug info from -debugify-function
      --chr                                                                - Reduce control height in the hot paths
      --codegenprepare                                                     - Optimize for code generation
      --consthoist                                                         - Constant Hoisting
      --constmerge                                                         - Merge Duplicate Global Constants
      --constraint-elimination                                             - Constraint Elimination
      --coro-cleanup                                                       - Lower all coroutine related intrinsics
      --coro-early                                                         - Lower early coroutine intrinsics
      --coro-elide                                                         - Coroutine frame allocation elision and indirect calls replacement
      --coro-split                                                         - Split coroutine into a set of functions driving its state machine
      --correlated-propagation                                             - Value Propagation
      --cost-model                                                         - Cost Model Analysis
      --cross-dso-cfi                                                      - Cross-DSO CFI
      --cseinfo                                                            - Analysis containing CSE Info
      --da                                                                 - Dependence Analysis
      --dce                                                                - Dead Code Elimination
      --deadargelim                                                        - Dead Argument Elimination
      --deadarghaX0r                                                       - Dead Argument Hacking (BUGPOINT USE ONLY; DO NOT USE)
      --debugify                                                           - Attach debug info to everything
      --debugify-function                                                  - Attach debug info to a function
      --delinearize                                                        - Delinearization
      --demanded-bits                                                      - Demanded bits analysis
      --dfsan                                                              - DataFlowSanitizer: dynamic data flow analysis.
      --div-rem-pairs                                                      - Hoist/decompose integer division and remainder
      --divergence                                                         - Legacy Divergence Analysis
      --domfrontier                                                        - Dominance Frontier Construction
      --domtree                                                            - Dominator Tree Construction
      --dot-callgraph                                                      - Print call graph to 'dot' file
      --dot-cfg                                                            - Print CFG of function to 'dot' file
      --dot-cfg-only                                                       - Print CFG of function to 'dot' file (with no function bodies)
      --dot-dom                                                            - Print dominance tree of function to 'dot' file
      --dot-dom-only                                                       - Print dominance tree of function to 'dot' file (with no function bodies)
      --dot-postdom                                                        - Print postdominance tree of function to 'dot' file
      --dot-postdom-only                                                   - Print postdominance tree of function to 'dot' file (with no function bodies)
      --dot-regions                                                        - Print regions of function to 'dot' file
      --dot-regions-only                                                   - Print regions of function to 'dot' file (with no function bodies)
      --dot-scops                                                          - Polly - Print Scops of function
      --dot-scops-only                                                     - Polly - Print Scops of function (with no function bodies)
      --dse                                                                - Dead Store Elimination
      --dwarfehprepare                                                     - Prepare DWARF exceptions
      --early-cse                                                          - Early CSE
      --early-cse-memssa                                                   - Early CSE w/ MemorySSA
      --edge-bundles                                                       - Bundle Machine CFG Edges
      --ee-instrument                                                      - Instrument function entry/exit with calls to e.g. mcount() (pre inlining)
      --elim-avail-extern                                                  - Eliminate Available Externally Globals
      --expand-reductions                                                  - Expand reduction intrinsics
      --expandmemcmp                                                       - Expand memcmp() to load/stores
      --external-aa                                                        - External Alias Analysis
      --extract-blocks                                                     - Extract basic blocks from module
      --fix-irreducible                                                    - Convert irreducible control-flow into natural loops
      --flattencfg                                                         - Flatten the CFG
      --float2int                                                          - Float to int
      --forceattrs                                                         - Force set function attributes
      --function-attrs                                                     - Deduce function attributes
      --function-import                                                    - Summary Based Function Import
      --gcn-dpp-combine                                                    - GCN DPP Combine
      --generic-to-nvvm                                                    - Ensure that the global variables are in the global address space
      --gisel-known-bits                                                   - Analysis for ComputingKnownBits
      --global-merge                                                       - Merge global variables
      --globaldce                                                          - Dead Global Elimination
      --globalopt                                                          - Global Variable Optimizer
      --globals-aa                                                         - Globals Alias Analysis
      --globalsplit                                                        - Global splitter
      --guard-widening                                                     - Widen guards
      --gvn                                                                - Global Value Numbering
      --gvn-hoist                                                          - Early GVN Hoisting of Expressions
      --gvn-sink                                                           - Early GVN sinking of Expressions
      --hardware-loops                                                     - Hardware Loop Insertion
      --hexagon-bit-simplify                                               - Hexagon bit simplification
      --hexagon-cext-opt                                                   - Hexagon constant-extender optimization
      --hexagon-constp                                                     - Hexagon Constant Propagation
      --hexagon-early-if                                                   - Hexagon early if conversion
      --hexagon-gen-mux                                                    - Hexagon generate mux instructions
      --hexagon-loop-idiom                                                 - Recognize Hexagon-specific loop idioms
      --hexagon-nvj                                                        - Hexagon NewValueJump
      --hexagon-packetizer                                                 - Hexagon Packetizer
      --hexagon-rdf-opt                                                    - Hexagon RDF optimizations
      --hexagon-split-double                                               - Hexagon Split Double Registers
      --hexagon-vc                                                         - Hexagon Vector Combine
      --hexagon-vextract                                                   - Hexagon optimize vextract
      --hexagon-vlcr                                                       - Hexagon-specific predictive commoning for HVX vectors
      --hotcoldsplit                                                       - Hot Cold Splitting
      --hwasan                                                             - HWAddressSanitizer: detect memory bugs using tagged addressing.
      --hwloops                                                            - Hexagon Hardware Loops
      --indirectbr-expand                                                  - Expand indirectbr instructions
      --indvars                                                            - Induction Variable Simplification
      --infer-address-spaces                                               - Infer address spaces
      --inferattrs                                                         - Infer set function attributes
      --inject-tli-mappings                                                - Inject TLI Mappings
      --inline                                                             - Function Integration/Inlining
      --insert-gcov-profiling                                              - Insert instrumentation for GCOV profiling
      --instcombine                                                        - Combine redundant instructions
      --instcount                                                          - Counts the various types of Instructions
      --instnamer                                                          - Assign names to anonymous instructions
      --instrorderfile                                                     - Instrumentation for Order File
      --instrprof                                                          - Frontend instrumentation-based coverage lowering.
      --instruction-select                                                 - Select target instructions out of generic instructions
      --instsimplify                                                       - Remove redundant instructions
      --interleaved-access                                                 - Lower interleaved memory accesses to target specific intrinsics
      --interleaved-load-combine                                           - Combine interleaved loads into wide loads and shufflevector instructions
      --internalize                                                        - Internalize Global Symbols
      --intervals                                                          - Interval Partition Construction
      --ipsccp                                                             - Interprocedural Sparse Conditional Constant Propagation
      --ir-similarity-identifier                                           - ir-similarity-identifier
      --irce                                                               - Inductive range check elimination
      --iroutliner                                                         - IR Outliner
      --irtranslator                                                       - IRTranslator LLVM IR -> MI
      --iv-users                                                           - Induction Variable Users
      --jump-threading                                                     - Jump Threading
      --lazy-block-freq                                                    - Lazy Block Frequency Analysis
      --lazy-branch-prob                                                   - Lazy Branch Probability Analysis
      --lazy-value-info                                                    - Lazy Value Information Analysis
      --lcssa                                                              - Loop-Closed SSA Form Pass
      --lcssa-verification                                                 - LCSSA Verifier
      --legalizer                                                          - Legalize the Machine IR a function's Machine IR
      --libcalls-shrinkwrap                                                - Conditionally eliminate dead library calls
      --licm                                                               - Loop Invariant Code Motion
      --lint                                                               - Statically lint-checks LLVM IR
      --liveintervals                                                      - Live Interval Analysis
      --liveregmatrix                                                      - Live Register Matrix
      --load-store-vectorizer                                              - Vectorize load and store instructions
      --localizer                                                          - Move/duplicate certain instructions close to their use
      --loop-accesses                                                      - Loop Access Analysis
      --loop-data-prefetch                                                 - Loop Data Prefetch
      --loop-deletion                                                      - Delete dead loops
      --loop-distribute                                                    - Loop Distribution
      --loop-extract                                                       - Extract loops into new functions
      --loop-extract-single                                                - Extract at most one loop into a new function
      --loop-flatten                                                       - Flattens loops
      --loop-fusion                                                        - Loop Fusion
      --loop-guard-widening                                                - Widen guards (within a single loop, as a loop pass)
      --loop-idiom                                                         - Recognize loop idioms
      --loop-instsimplify                                                  - Simplify instructions in loops
      --loop-interchange                                                   - Interchanges loops for cache reuse
      --loop-load-elim                                                     - Loop Load Elimination
      --loop-predication                                                   - Loop predication
      --loop-reduce                                                        - Loop Strength Reduction
      --loop-reroll                                                        - Reroll loops
      --loop-rotate                                                        - Rotate Loops
      --loop-simplify                                                      - Canonicalize natural loops
      --loop-simplifycfg                                                   - Simplify loop CFG
      --loop-sink                                                          - Loop Sink
      --loop-unroll                                                        - Unroll loops
      --loop-unroll-and-jam                                                - Unroll and Jam loops
      --loop-unswitch                                                      - Unswitch loops
      --loop-vectorize                                                     - Loop Vectorization
      --loop-versioning                                                    - Loop Versioning
      --loop-versioning-licm                                               - Loop Versioning For LICM
      --loops                                                              - Natural Loop Information
      --lower-amx-type                                                     - Lower AMX type for load/store
      --lower-constant-intrinsics                                          - Lower constant intrinsics
      --lower-expect                                                       - Lower 'expect' Intrinsics
      --lower-guard-intrinsic                                              - Lower the guard intrinsic to normal control flow
      --lower-matrix-intrinsics                                            - Lower the matrix intrinsics
      --lower-matrix-intrinsics-minimal                                    - Lower the matrix intrinsics (minimal)
      --lower-widenable-condition                                          - Lower the widenable condition to default true value
      --loweratomic                                                        - Lower atomic intrinsics to non-atomic form
      --lowerinvoke                                                        - Lower invoke and unwind, for unwindless code generators
      --lowerswitch                                                        - Lower SwitchInst's to branches
      --lowertypetests                                                     - Lower type metadata
      --machine-block-freq                                                 - Machine Block Frequency Analysis
      --machine-branch-prob                                                - Machine Branch Probability Analysis
      --machine-domfrontier                                                - Machine Dominance Frontier Construction
      --machine-loops                                                      - Machine Natural Loop Construction
      --machine-trace-metrics                                              - Machine Trace Metrics
      --machinedomtree                                                     - MachineDominator Tree Construction
      --machinepostdomtree                                                 - MachinePostDominator Tree Construction
      --make-guards-explicit                                               - Lower the guard intrinsic to explicit control flow form
      --mem2reg                                                            - Promote Memory to Register
      --memcpyopt                                                          - MemCpy Optimization
      --memdep                                                             - Memory Dependence Analysis
      --memoryssa                                                          - Memory SSA
      --memprof                                                            - MemProfiler: profile memory allocations and accesses.
      --memprof-module                                                     - MemProfiler: profile memory allocations and accesses.ModulePass
      --mergefunc                                                          - Merge Functions
      --mergeicmps                                                         - Merge contiguous icmps into a memcmp
      --mergereturn                                                        - Unify function exit nodes
      --metarenamer                                                        - Assign new names to everything
      --micromips-reduce-size                                              - MicroMips instruction size reduce pass
      --mips-branch-expansion                                              - Expand out of range branch instructions and fix forbidden slot hazards
      --mips-delay-slot-filler                                             - Fill delay slot for MIPS
      --mips-prelegalizer-combiner                                         - Combine Mips machine instrs before legalization
      --mldst-motion                                                       - MergedLoadStoreMotion
      --module-debuginfo                                                   - Decodes module-level debug info
      --module-summary-analysis                                            - Module Summary Analysis
      --module-summary-info                                                - Module summary info
      --msan                                                               - MemorySanitizer: detects uninitialized reads.
      --mve-tail-predication                                               - Transform predicated vector loops to use MVE tail predication
      --name-anon-globals                                                  - Provide a name to nameless globals
      --nary-reassociate                                                   - Nary reassociation
      --newgvn                                                             - Global Value Numbering
      --nvptx-assign-valid-global-names                                    - Assign valid PTX names to globals
      --nvptx-lower-aggr-copies                                            - Lower aggregate copies, and llvm.mem* intrinsics into loops
      --nvptx-lower-alloca                                                 - Lower Alloca
      --nvptx-lower-args                                                   - Lower arguments (NVPTX)
      --nvptx-proxyreg-erasure                                             - NVPTX ProxyReg Erasure
      --nvvm-intr-range                                                    - Add !range metadata to NVVM intrinsics.
      --nvvm-reflect                                                       - Replace occurrences of __nvvm_reflect() calls with 0/1
      --objc-arc                                                           - ObjC ARC optimization
      --objc-arc-aa                                                        - ObjC-ARC-Based Alias Analysis
      --objc-arc-apelim                                                    - ObjC ARC autorelease pool elimination
      --objc-arc-contract                                                  - ObjC ARC contraction
      --objc-arc-expand                                                    - ObjC ARC expansion
      --openmpopt                                                          - OpenMP specific optimizations
      --opt-remark-emitter                                                 - Optimization Remark Emitter
      --pa-eval                                                            - Evaluate ProvenanceAnalysis on all pairs
      --packets                                                            - R600 Packetizer
      --partial-inliner                                                    - Partial Inliner
      --partially-inline-libcalls                                          - Partially inline calls to library functions
      --pgo-icall-prom                                                     - Use PGO instrumentation profile to promote indirect calls to direct calls.
      --pgo-instr-gen                                                      - PGO instrumentation.
      --pgo-instr-use                                                      - Read PGO instrumentation profile.
      --pgo-memop-opt                                                      - Optimize memory intrinsic using its size value profile
      --phi-values                                                         - Phi Values Analysis
      --place-backedge-safepoints-impl                                     - Place Backedge Safepoints
      --place-safepoints                                                   - Place Safepoints
      --polly-ast                                                          - Polly - Generate an AST from the SCoP (isl)
      --polly-canonicalize                                                 - Polly - Run canonicalization passes
      --polly-cleanup                                                      - Polly - Cleanup after code generation
      --polly-codegen                                                      - Polly - Create LLVM-IR from SCoPs
      --polly-dce                                                          - Polly - Remove dead iterations
      --polly-delicm                                                       - Polly - DeLICM/DePRE
      --polly-dependences                                                  - Polly - Calculate dependences
      --polly-detect                                                       - Polly - Detect static control parts (SCoPs)
      --polly-dump-module                                                  - Polly - Dump Module
      --polly-export-jscop                                                 - Polly - Export Scops as JSON (Writes a .jscop file for each Scop)
      --polly-flatten-schedule                                             - Polly - Flatten schedule
      --polly-function-dependences                                         - Polly - Calculate dependences for all the SCoPs of a function
      --polly-function-scops                                               - Polly - Create polyhedral description of all Scops of a function
      --polly-import-jscop                                                 - Polly - Import Scops from JSON (Reads a .jscop file for each Scop)
      --polly-mse                                                          - Polly - Maximal static expansion of SCoP
      --polly-opt-isl                                                      - Polly - Optimize schedule of SCoP
      --polly-optree                                                       - Polly - Forward operand tree
      --polly-prepare                                                      - Polly - Prepare code for polly
      --polly-prune-unprofitable                                           - Polly - Prune unprofitable SCoPs
      --polly-rewrite-byref-params                                         - Polly - Rewrite by reference parameters
      --polly-scop-inliner                                                 - inline functions based on how much of the function is a scop.
      --polly-scops                                                        - Polly - Create polyhedral description of Scops
      --polly-simplify                                                     - Polly - Simplify
      --polyhedral-info                                                    - Polly - Interface to polyhedral analysis engine
      --post-inline-ee-instrument                                          - Instrument function entry/exit with calls to e.g. mcount() (post inlining)
      --postdomtree                                                        - Post-Dominator Tree Construction
      --ppc-bool-ret-to-int                                                - Convert i1 constants to i32/i64 if they are returned
      --ppc-branch-coalescing                                              - Branch Coalescing
      --ppc-branch-select                                                  - PowerPC Branch Selector
      --ppc-early-ret                                                      - PowerPC Early-Return Creation
      --ppc-expand-isel                                                    - PowerPC Expand ISEL Generation
      --ppc-loop-instr-form-prep                                           - Prepare loop for ppc preferred instruction forms
      --ppc-lower-massv-entries                                            - Lower MASSV entries
      --ppc-mi-peepholes                                                   - PowerPC MI Peephole Optimization
      --ppc-pre-emit-peephole                                              - PowerPC Pre-Emit Peephole
      --ppc-reduce-cr-ops                                                  - PowerPC Reduce CR logical Operation
      --ppc-tls-dynamic-call                                               - PowerPC TLS Dynamic Call Fixup
      --ppc-toc-reg-deps                                                   - PowerPC TOC Register Dependencies
      --ppc-vsx-copy                                                       - PowerPC VSX Copy Legalization
      --ppc-vsx-fma-mutate                                                 - PowerPC VSX FMA Mutation
      --ppc-vsx-swaps                                                      - PowerPC VSX Swap Removal
      --pre-isel-intrinsic-lowering                                        - Pre-ISel Intrinsic Lowering
      --print-alias-sets                                                   - Alias Set Printer
      --print-callgraph                                                    - Print a call graph
      --print-callgraph-sccs                                               - Print SCCs of the Call Graph
      --print-cfg-sccs                                                     - Print SCCs of each function CFG
      --print-dom-info                                                     - Dominator Info Printer
      --print-externalfnconstants                                          - Print external fn callsites passed constants
      --print-function                                                     - Print function to stderr
      --print-lazy-value-info                                              - Lazy Value Info Printer Pass
      --print-memdeps                                                      - Print MemDeps of function
      --print-memderefs                                                    - Memory Dereferenciblity of pointers in function
      --print-memoryssa                                                    - Memory SSA Printer
      --print-module                                                       - Print module to stderr
      --print-must-be-executed-contexts                                    - print the must-be-executed-context for all instructions
      --print-mustexecute                                                  - Instructions which execute on loop entry
      --print-predicateinfo                                                - PredicateInfo Printer
      --profile-summary-info                                               - Profile summary info
      --prune-eh                                                           - Remove unused exception handling info
      --pseudo-probe-inserter                                              - Insert pseudo probe annotations for value profiling
      --r600-expand-special-instrs                                         - R600ExpandSpecialInstrs
      --r600cf                                                             - R600 Control Flow Finalizer
      --r600mergeclause                                                    - R600 Clause Merge
      --reaching-deps-analysis                                             - ReachingDefAnalysis
      --reassociate                                                        - Reassociate expressions
      --redundant-dbg-inst-elim                                            - Redundant Dbg Instruction Elimination
      --reg2mem                                                            - Demote all values to stack slots
      --regbankselect                                                      - Assign register bank of generic virtual registers
      --regions                                                            - Detect single entry single exit regions
      --rewrite-statepoints-for-gc                                         - Make relocations explicit at statepoints
      --rewrite-symbols                                                    - Rewrite Symbols
      --riscv-cleanup-vsetvli                                              - RISCV Cleanup VSETVLI pass
      --riscv-expand-pseudo                                                - RISCV pseudo instruction expansion pass
      --riscv-merge-base-offset                                            - RISCV Merge Base Offset
      --rpo-function-attrs                                                 - Deduce function attributes in RPO
      --safe-stack                                                         - Safe Stack instrumentation pass
      --sample-profile                                                     - Sample Profile loader
      --sancov                                                             - Pass for instrumenting coverage on functions
      --scalar-evolution                                                   - Scalar Evolution Analysis
      --scalarize-masked-mem-intrin                                        - Scalarize unsupported masked memory intrinsics
      --scalarizer                                                         - Scalarize vector operations
      --sccp                                                               - Sparse Conditional Constant Propagation
      --scev-aa                                                            - ScalarEvolution-based Alias Analysis
      --scoped-noalias-aa                                                  - Scoped NoAlias Alias Analysis
      --separate-const-offset-from-gep                                     - Split GEPs to a variadic base and a constant offset for better CSE
      --si-annotate-control-flow                                           - Annotate SI Control Flow
      --si-fix-sgpr-copies                                                 - SI Fix SGPR copies
      --si-fix-vgpr-copies                                                 - SI Fix VGPR copies
      --si-fold-operands                                                   - SI Fold Operands
      --si-form-memory-clauses                                             - SI Form memory clauses
      --si-i1-copies                                                       - SI Lower i1 Copies
      --si-img-init                                                        - SI Add IMG Init
      --si-insert-hard-clauses                                             - SI Insert Hard Clauses
      --si-insert-skips                                                    - SI insert s_cbranch_execz instructions
      --si-insert-waitcnts                                                 - SI Insert Waitcnts
      --si-load-store-opt                                                  - SI Load Store Optimizer
      --si-lower-control-flow                                              - SI lower control flow
      --si-lower-sgpr-spills                                               - SI lower SGPR spill instructions
      --si-memory-legalizer                                                - SI Memory Legalizer
      --si-mode-register                                                   - Insert required mode register values
      --si-optimize-exec-masking                                           - SI optimize exec mask operations
      --si-optimize-exec-masking-pre-ra                                    - SI optimize exec mask operations pre-RA
      --si-peephole-sdwa                                                   - SI Peephole SDWA
      --si-post-ra-bundler                                                 - SI post-RA bundler
      --si-pre-allocate-wwm-regs                                           - SI Pre-allocate WWM Registers
      --si-pre-emit-peephole                                               - SI peephole optimizations
      --si-remove-short-exec-branches                                      - SI remove short exec branches
      --si-shrink-instructions                                             - SI Shrink Instructions
      --si-wqm                                                             - SI Whole Quad Mode
      --simple-loop-unswitch                                               - Simple unswitch loops
      --simplifycfg                                                        - Simplify the CFG
      --sink                                                               - Code sinking
      --sjljehprepare                                                      - Prepare SjLj exceptions
      --slotindexes                                                        - Slot index numbering
      --slp-vectorizer                                                     - SLP Vectorizer
      --slsr                                                               - Straight line strength reduction
      --speculative-execution                                              - Speculatively execute instructions
      --sroa                                                               - Scalar Replacement Of Aggregates
      --stack-protector                                                    - Insert stack protectors
      --stack-safety                                                       - Stack Safety Analysis
      --stack-safety-local                                                 - Stack Safety Local Analysis
      --strip                                                              - Strip all symbols from a module
      --strip-dead-debug-info                                              - Strip debug info for unused symbols
      --strip-dead-prototypes                                              - Strip Unused Function Prototypes
      --strip-debug-declare                                                - Strip all llvm.dbg.declare intrinsics
      --strip-gc-relocates                                                 - Strip gc.relocates inserted through RewriteStatepointsForGC
      --strip-nondebug                                                     - Strip all symbols, except dbg symbols, from a module
      --strip-nonlinetable-debuginfo                                       - Strip all debug info except linetables
      --structurizecfg                                                     - Structurize the CFG
      --tailcallelim                                                       - Tail Call Elimination
      --targetlibinfo                                                      - Target Library Information
      --targetpassconfig                                                   - Target Pass Configuration
      --tbaa                                                               - Type-Based Alias Analysis
      --thumb2-reduce-size                                                 - Thumb2 instruction size reduce pass
      --tileconfig                                                         - Tile Register Configure
      --transform-warning                                                  - Warn about non-applied transformations
      --tsan                                                               - ThreadSanitizer: detects data races.
      --tti                                                                - Target Transform Information
      --type-promotion                                                     - Type Promotion
      --unify-loop-exits                                                   - Fixup each natural loop to have a single exit block
      --unique-internal-linkage-names                                      - Uniqueify Internal linkage names
      --unreachableblockelim                                               - Remove unreachable blocks from the CFG
      --vec-merger                                                         - R600 Vector Reg Merger
      --vector-combine                                                     - Optimize scalar/vector ops
      --verify                                                             - Module Verifier
      --verify-safepoint-ir                                                - Safepoint IR Verifier
      --view-callgraph                                                     - View call graph
      --view-cfg                                                           - View CFG of function
      --view-cfg-only                                                      - View CFG of function (with no function bodies)
      --view-dom                                                           - View dominance tree of function
      --view-dom-only                                                      - View dominance tree of function (with no function bodies)
      --view-postdom                                                       - View postdominance tree of function
      --view-postdom-only                                                  - View postdominance tree of function (with no function bodies)
      --view-regions                                                       - View regions of function
      --view-regions-only                                                  - View regions of function (with no function bodies)
      --view-scops                                                         - Polly - View Scops of function
      --view-scops-only                                                    - Polly - View Scops of function (with no function bodies)
      --virtregmap                                                         - Virtual Register Map
      --wasm-add-missing-prototypes                                        - Add prototypes to prototypes-less functions
      --wasm-argument-move                                                 - Move ARGUMENT instructions for WebAssembly
      --wasm-cfg-sort                                                      - Reorders blocks in topological order
      --wasm-cfg-stackify                                                  - Insert BLOCK/LOOP/TRY markers for WebAssembly scopes
      --wasm-debug-fixup                                                   - Ensures debug_value's that have been stackified become stack relative
      --wasm-exception-info                                                - WebAssembly Exception Information
      --wasm-explicit-locals                                               - Convert registers to WebAssembly locals
      --wasm-fix-function-bitcasts                                         - Fix mismatching bitcasts for WebAssembly
      --wasm-fix-irreducible-control-flow                                  - Removes irreducible control flow
      --wasm-late-eh-prepare                                               - WebAssembly Late Exception Preparation
      --wasm-lower-br_unless                                               - Lowers br_unless into inverted br_if
      --wasm-lower-em-ehsjlj                                               - WebAssembly Lower Emscripten Exceptions / Setjmp / Longjmp
      --wasm-lower-global-dtors                                            - Lower @llvm.global_dtors for WebAssembly
      --wasm-mem-intrinsic-results                                         - Optimize memory intrinsic result values for WebAssembly
      --wasm-optimize-live-intervals                                       - Optimize LiveIntervals for WebAssembly
      --wasm-optimize-returned                                             - Optimize calls with "returned" attributes for WebAssembly
      --wasm-peephole                                                      - WebAssembly peephole optimizations
      --wasm-prepare-for-live-intervals                                    - Fix up code for LiveIntervals
      --wasm-reg-coloring                                                  - Minimize number of registers used
      --wasm-reg-numbering                                                 - Assigns WebAssembly register numbers for virtual registers
      --wasm-reg-stackify                                                  - Reorder instructions to use the WebAssembly value stack
      --wasm-replace-phys-regs                                             - Replace physical registers with virtual registers
      --wasm-set-p2align-operands                                          - Set the p2align operands for WebAssembly loads and stores
      --wasmehprepare                                                      - Prepare WebAssembly exceptions
      --wholeprogramdevirt                                                 - Whole program devirtualization
      --winehprepare                                                       - Prepare Windows exceptions
      --write-bitcode                                                      - Write Bitcode
      --x86-avoid-SFB                                                      - Machine code sinking
      --x86-avoid-trailing-call                                            - X86 avoid trailing call pass
      --x86-cf-opt                                                         - X86 Call Frame Optimization
      --x86-cmov-conversion                                                - X86 cmov Conversion
      --x86-codegen                                                        - X86 FP Stackifier
      --x86-domain-reassignment                                            - X86 Domain Reassignment Pass
      --x86-evex-to-vex-compress                                           - Compressing EVEX instrs to VEX encoding when possible
      --x86-execution-domain-fix                                           - X86 Execution Domain Fix
      --x86-fixup-LEAs                                                     - X86 LEA Fixup
      --x86-fixup-bw-insts                                                 - X86 Byte/Word Instruction Fixup
      --x86-fixup-setcc                                                    - x86-fixup-setcc
      --x86-flags-copy-lowering                                            - X86 EFLAGS copy lowering
      --x86-lvi-load                                                       - X86 LVI load hardening
      --x86-lvi-ret                                                        - X86 LVI ret hardener
      --x86-optimize-LEAs                                                  - X86 optimize LEA pass
      --x86-partial-reduction                                              - X86 Partial Reduction
      --x86-pseudo                                                         - X86 pseudo instruction expansion pass
      --x86-seses                                                          - X86 Speculative Execution Side Effect Suppression
      --x86-slh                                                            - X86 speculative load hardener
      --x86-winehstate                                                     - Insert stores for EH state numbers
  --lto-embed-bitcode=<value>                                           - Embed LLVM bitcode in object files produced by LTO
    =none                                                               -   Do not embed
    =optimized                                                          -   Embed after all optimization passes
    =post-merge-pre-opt                                                 -   Embed post merge, but before optimizations
  --lto-pass-remarks-filter=<regex>                                     - Only record optimization remarks from passes whose names match the given regular expression
  --lto-pass-remarks-format=<format>                                    - The format used for serializing remarks (default: YAML)
  --lto-pass-remarks-output=<filename>                                  - Output filename for pass remarks
  --march=<string>                                                      - Architecture to generate code for (see --version)
  --matrix-default-layout=<value>                                       - Sets the default matrix layout
    =column-major                                                       -   Use column-major layout
    =row-major                                                          -   Use row-major layout
  --mattr=<a1,+a2,-a3,...>                                              - Target specific attributes (-mattr=help for details)
  --max-counter-promotions=<int>                                        - Max number of allowed counter promotions
  --max-counter-promotions-per-loop=<uint>                              - Max number counter promotions per loop to avoid increasing register pressure too much
  --mc-relax-all                                                        - When used with filetype=obj, relax all fixups in the emitted object file
  --mcpu=<cpu-name>                                                     - Target a specific cpu type (-mcpu=help for details)
  --meabi=<value>                                                       - Set EABI type (default depends on triple):
    =default                                                            -   Triple default EABI version
    =4                                                                  -   EABI version 4
    =5                                                                  -   EABI version 5
    =gnu                                                                -   EABI GNU
  --merror-missing-parenthesis                                          - Error for missing parenthesis around predicate registers
  --merror-noncontigious-register                                       - Error for register names that aren't contigious
  --mhvx                                                                - Enable Hexagon Vector eXtensions
  --mhvx=<value>                                                        - Enable Hexagon Vector eXtensions
    =v60                                                                -   Build for HVX v60
    =v62                                                                -   Build for HVX v62
    =v65                                                                -   Build for HVX v65
    =v66                                                                -   Build for HVX v66
    =v67                                                                -   Build for HVX v67
  --mips-compact-branches=<value>                                       - MIPS Specific: Compact branch policy.
    =never                                                              -   Do not use compact branches if possible.
    =optimal                                                            -   Use compact branches where appropriate (default).
    =always                                                             -   Always use compact branches if possible.
  --mips16-constant-islands                                             - Enable mips16 constant islands.
  --mips16-hard-float                                                   - Enable mips16 hard float.
  --mir-strip-debugify-only                                             - Should mir-strip-debug only strip debug info from debugified modules by default
  --mno-compound                                                        - Disable looking for compound instructions for Hexagon
  --mno-fixup                                                           - Disable fixing up resolved relocations for Hexagon
  --mno-ldc1-sdc1                                                       - Expand double precision loads and stores to their single precision counterparts
  --mno-pairing                                                         - Disable looking for duplex instructions for Hexagon
  --module-hash                                                         - Emit module hash
  --module-summary                                                      - Emit module summary index
  --mtriple=<string>                                                    - Override target triple for module
  --mwarn-missing-parenthesis                                           - Warn for missing parenthesis around predicate registers
  --mwarn-noncontigious-register                                        - Warn for register names that arent contigious
  --mwarn-sign-mismatch                                                 - Warn for mismatching a signed and unsigned value
  --no-deprecated-warn                                                  - Suppress all deprecated warnings
  --no-discriminators                                                   - Disable generation of discriminator information.
  --no-warn                                                             - Suppress all warnings
  --no-xray-index                                                       - Don't emit xray_fn_idx section
  --nozero-initialized-in-bss                                           - Don't place zero-initialized symbols into bss section
  --nvptx-sched4reg                                                     - NVPTX Specific: schedule for register pressue
  -o=<filename>                                                         - Override output filename
  -p                                                                    - Print module after each transformation
  --pass-remarks-filter=<regex>                                         - Only record optimization remarks from passes whose names match the given regular expression
  --pass-remarks-format=<format>                                        - The format used for serializing remarks (default: YAML)
  --pass-remarks-output=<filename>                                      - Output filename for pass remarks
  --poison-checking-function-local                                      - Check that returns are non-poison (for testing)
  --print-breakpoints-for-testing                                       - Print select breakpoints location for testing
  --pseudo-probe-for-profiling                                          - Emit pseudo probes for AutoFDO
  --r600-ir-structurize                                                 - Use StructurizeCFG IR pass
  --rdf-dump                                                            - 
  --rdf-limit=<uint>                                                    - 
  --relax-elf-relocations                                               - Emit GOTPCRELX/REX_GOTPCRELX instead of GOTPCREL on x86-64 ELF
  --relocation-model=<value>                                            - Choose relocation model
    =static                                                             -   Non-relocatable code
    =pic                                                                -   Fully relocatable, position independent code
    =dynamic-no-pic                                                     -   Relocatable external references, non-relocatable code
    =ropi                                                               -   Code and read-only data relocatable, accessed PC-relative
    =rwpi                                                               -   Read-write data relocatable, accessed relative to static base
    =ropi-rwpi                                                          -   Combination of ropi and rwpi
  --runtime-counter-relocation                                          - Enable relocating counters at runtime.
  --safepoint-ir-verifier-print-only                                    - 
  --sample-profile-check-record-coverage=<N>                            - Emit a warning if less than N% of records in the input profile are matched to the IR.
  --sample-profile-check-sample-coverage=<N>                            - Emit a warning if less than N% of samples in the input profile are matched to the IR.
  --sample-profile-max-propagate-iterations=<uint>                      - Maximum number of iterations to go through when propagating sample block/edge weights through the CFG.
  --skip-ret-exit-block                                                 - Suppress counter promotion if exit blocks contain ret.
  --speculative-counter-promotion-max-exiting=<uint>                    - The max number of exiting blocks of a loop to allow  speculative counter promotion
  --speculative-counter-promotion-to-loop                               - When the option is false, if the target block is in a loop, the promotion will be disallowed unless the promoted counter  update can be further/iteratively promoted into an acyclic  region.
  --split-machine-functions                                             - Split out cold basic blocks from machine functions based on profile information
  --stack-alignment=<uint>                                              - Override default stack alignment
  --stack-protector-guard=<string>                                      - Stack protector guard mode
  --stack-protector-guard-offset=<uint>                                 - Stack protector guard offset
  --stack-protector-guard-reg=<string>                                  - Stack protector guard register
  --stack-size-section                                                  - Emit a section containing stack size metadata
  --stack-symbol-ordering                                               - Order local stack symbols.
  --stackrealign                                                        - Force align the stack to the minimum alignment
  --std-link-opts                                                       - Include the standard link time optimizations
  --strip-debug                                                         - Strip debugger symbol info from translation unit
  --strip-named-metadata                                                - Strip module-level named metadata
  --summary-file=<string>                                               - The summary file to use for function importing.
  --tail-predication=<value>                                            - MVE tail-predication pass options
    =disabled                                                           -   Don't tail-predicate loops
    =enabled-no-reductions                                              -   Enable tail-predication, but not for reduction loops
    =enabled                                                            -   Enable tail-predication, including reduction loops
    =force-enabled-no-reductions                                        -   Enable tail-predication, but not for reduction loops, and force this which might be unsafe
    =force-enabled                                                      -   Enable tail-predication, including reduction loops, and force this which might be unsafe
  --tailcallopt                                                         - Turn fastcc calls into tail calls by (potentially) changing ABI.
  --thin-link-bitcode-file=<filename>                                   - A file in which to write minimized bitcode for the thin link only
  --thinlto-assume-merged                                               - Assume the input has already undergone ThinLTO function importing and the other pre-optimization pipeline changes.
  --thinlto-bc                                                          - Write output as ThinLTO-ready bitcode
  --thinlto-split-lto-unit                                              - Enable splitting of a ThinLTO LTOUnit
  --thread-model=<value>                                                - Choose threading model
    =posix                                                              -   POSIX thread model
    =single                                                             -   Single thread model
  --threads=<int>                                                       - 
  --time-trace                                                          - Record time trace
  --time-trace-file=<filename>                                          - Specify time trace file destination
  --tls-size=<uint>                                                     - Bit size of immediate TLS offsets
  --unique-basic-block-section-names                                    - Give unique names to every basic block section
  --unique-section-names                                                - Give unique names to every section
  --use-ctors                                                           - Use .ctors instead of .init_array.
  --vec-extabi                                                          - Enable the AIX Extended Altivec ABI.
  --verify-each                                                         - Verify after each transform
  --verify-region-info                                                  - Verify region info (time consuming)
  --vp-counters-per-site=<number>                                       - The average number of profile counters allocated per value profiling site.
  --vp-static-alloc                                                     - Do static counter allocation for value profiler
  --x86-align-branch=<string>                                           - Specify types of branches to align (plus separated list of types):
                                                                          jcc      indicates conditional jumps
                                                                          fused    indicates fused conditional jumps
                                                                          jmp      indicates direct unconditional jumps
                                                                          call     indicates direct and indirect calls
                                                                          ret      indicates rets
                                                                          indirect indicates indirect unconditional jumps
  --x86-align-branch-boundary=<uint>                                    - Control how the assembler should align branches with NOP. If the boundary's size is not 0, it should be a power of 2 and no less than 32. Branches will be aligned to prevent from being across or against the boundary of specified size. The default value 0 does not align branches.
  --x86-branches-within-32B-boundaries                                  - Align selected instructions to mitigate negative performance impact of Intel's micro code update for errata skx102.  May break assumptions about labels corresponding to particular instructions, and should be used with caution.
  --x86-pad-max-prefix-size=<uint>                                      - Maximum number of prefixes to use for padding
  --xcoff-traceback-table                                               - Emit the XCOFF traceback table

Generic Options:

  --help                                                                - Display available options (--help-hidden for more)
  --help-list                                                           - Display list of available options (--help-list-hidden for more)
  --version                                                             - Display the version of this program

Polly Options:
Configure the polly loop optimizer

  --polly                                                               - Enable the polly optimizer (only at -O3)
  --polly-2nd-level-tiling                                              - Enable a 2nd level loop of loop tiling
  --polly-ast-print-accesses                                            - Print memory access functions
  --polly-context=<isl parameter set>                                   - Provide additional constraints on the context parameters
  --polly-dce-precise-steps=<int>                                       - The number of precise steps between two approximating iterations. (A value of -1 schedules another approximation stage before the actual dead code elimination.
  --polly-delicm-max-ops=<int>                                          - Maximum number of isl operations to invest for lifetime analysis; 0=no limit
  --polly-detect-full-functions                                         - Allow the detection of full functions
  --polly-dump-after                                                    - Dump module after Polly transformations into a file suffixed with "-after"
  --polly-dump-after-file=<string>                                      - Dump module after Polly transformations to the given file
  --polly-dump-before                                                   - Dump module before Polly transformations into a file suffixed with "-before"
  --polly-dump-before-file=<string>                                     - Dump module before Polly transformations to the given file
  --polly-enable-simplify                                               - Simplify SCoP after optimizations
  --polly-ignore-func=<string>                                          - Ignore functions that match a regex. Multiple regexes can be comma separated. Scop detection will ignore all functions that match ANY of the regexes provided.
  --polly-isl-arg=<argument>                                            - Option passed to ISL
  --polly-on-isl-error-abort                                            - Abort if an isl error is encountered
  --polly-only-func=<string>                                            - Only run on functions that match a regex. Multiple regexes can be comma separated. Scop detection will run on all functions that match ANY of the regexes provided.
  --polly-only-region=<identifier>                                      - Only run on certain regions (The provided identifier must appear in the name of the region's entry block
  --polly-only-scop-detection                                           - Only run scop detection, but no other optimizations
  --polly-optimized-scops                                               - Polly - Dump polyhedral description of Scops optimized with the isl scheduling optimizer and the set of post-scheduling transformations is applied on the schedule tree
  --polly-parallel                                                      - Generate thread parallel code (isl codegen only)
  --polly-parallel-force                                                - Force generation of thread parallel code ignoring any cost model
  --polly-pattern-matching-based-opts                                   - Perform optimizations based on pattern matching
  --polly-process-unprofitable                                          - Process scops that are unlikely to benefit from Polly optimizations.
  --polly-register-tiling                                               - Enable register tiling
  --polly-report                                                        - Print information about the activities of Polly
  --polly-show                                                          - Highlight the code regions that will be optimized in a (CFG BBs and LLVM-IR instructions)
  --polly-show-only                                                     - Highlight the code regions that will be optimized in a (CFG only BBs)
  --polly-stmt-granularity=<value>                                      - Algorithm to use for splitting basic blocks into multiple statements
    =bb                                                                 -   One statement per basic block
    =scalar-indep                                                       -   Scalar independence heuristic
    =store                                                              -   Store-level granularity
  --polly-target=<value>                                                - The hardware to target
    =cpu                                                                -   generate CPU code
  --polly-tiling                                                        - Enable loop tiling
  --polly-vectorizer=<value>                                            - Select the vectorization strategy
    =none                                                               -   No Vectorization
    =polly                                                              -   Polly internal vectorizer
    =stripmine                                                          -   Strip-mine outer loops for the loop-vectorizer to trigger