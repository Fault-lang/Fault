(set-logic QF_NRA)
(declare-fun swaps_s2_v_0 () Real)
(declare-fun swaps_s2_v_1 () Real)
(declare-fun swaps_s2_v_2 () Real)
(declare-fun swaps_s2_v_3 () Real)
(declare-fun swaps_s2_v_4 () Real)
(declare-fun swaps_s2_v_5 () Real)
(declare-fun swaps_s2_v_6 () Real)
(declare-fun swaps_s2_v_7 () Real)
(declare-fun swaps_s2_v_8 () Real)
(declare-fun swaps_s2_v_9 () Real)
(declare-fun swaps_s2_v_10 () Real)
(assert (= swaps_s2_v_0 10.0))
(assert (= swaps_s2_v_1 (+ swaps_s2_v_0 2.0)))
(assert (= swaps_s2_v_2 (- swaps_s2_v_1 5.0)))
(assert (= swaps_s2_v_3 (- swaps_s2_v_0 5.0)))
(assert (= swaps_s2_v_4 (+ swaps_s2_v_3 2.0)))
(assert (or (= swaps_s2_v_5 swaps_s2_v_2) (= swaps_s2_v_5 swaps_s2_v_4)))
(assert (= swaps_s2_v_6 (+ swaps_s2_v_5 2.0)))
(assert (= swaps_s2_v_7 (- swaps_s2_v_6 5.0)))
(assert (= swaps_s2_v_8 (- swaps_s2_v_5 5.0)))
(assert (= swaps_s2_v_9 (+ swaps_s2_v_8 2.0)))
(assert (or (= swaps_s2_v_10 swaps_s2_v_7) (= swaps_s2_v_10 swaps_s2_v_9)))