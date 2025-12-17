(set-logic QF_NRA)
(declare-fun swaps_s2_v_0 () Real)
(declare-fun swaps_s2_v_1 () Real)
(declare-fun swaps_s2_v_2 () Real)
(declare-fun @__run_0 () Bool)
(declare-fun swaps_s2_v_4 () Real)
(declare-fun swaps_s2_v_5 () Real)
(declare-fun @__run_1 () Bool)
(declare-fun swaps_s2_v_7 () Real)
(declare-fun swaps_s2_v_8 () Real)
(declare-fun swaps_s2_v_10 () Real)
(declare-fun swaps_s2_v_11 () Real)
(assert (= swaps_s2_v_0 10.0))

(assert (= swaps_s2_v_1 (+ swaps_s2_v_0 2.0)))


(assert (= swaps_s2_v_2 (- swaps_s2_v_1 5.0)))

(assert (= @__run_0 (= swaps_s2_v_3 swaps_s2_v_2)))

(assert (= swaps_s2_v_4 (- swaps_s2_v_0 5.0)))


(assert (= swaps_s2_v_5 (+ swaps_s2_v_4 2.0)))

(assert (= @__run_1 (and (= swaps_s2_v_6 swaps_s2_v_2)
(= swaps_s2_v_6 swaps_s2_v_5))))
(assert (or (and @__run_0
(not @__run_1))
(and (not @__run_0)
@__run_1)))

(assert (= swaps_s2_v_7 (+ swaps_s2_v_6 2.0)))


(assert (= swaps_s2_v_8 (- swaps_s2_v_7 5.0)))

(assert (= @__run_0 (= swaps_s2_v_9 swaps_s2_v_8)))

(assert (= swaps_s2_v_10 (- swaps_s2_v_6 5.0)))


(assert (= swaps_s2_v_11 (+ swaps_s2_v_10 2.0)))

(assert (= @__run_1 (and (= swaps_s2_v_12 swaps_s2_v_8)
(= swaps_s2_v_12 swaps_s2_v_11))))
(assert (or (and @__run_0
(not @__run_1))
(and (not @__run_0)
@__run_1)))