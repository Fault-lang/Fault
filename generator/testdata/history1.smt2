(set-logic QF_NRA)
(declare-fun history1_c_place_value_0 () Real)
(declare-fun history1_c_place_value_1 () Real)
(declare-fun history1_c_place_value_2 () Real)
(declare-fun history1_c_place_value_3 () Real)
(declare-fun block3true_1 () Bool)
(declare-fun block4false_1 () Bool)
(declare-fun history1_c_place_value_4 () Real)
(declare-fun history1_c_place_value_5 () Real)
(declare-fun history1_c_place_value_6 () Real)
(declare-fun block3true_2 () Bool)
(declare-fun block4false_2 () Bool)
(declare-fun history1_c_place_value_7 () Real)
(declare-fun history1_c_place_value_8 () Real)
(declare-fun history1_c_place_value_9 () Real)
(declare-fun block3true_3 () Bool)
(declare-fun block4false_3 () Bool)
(declare-fun history1_c_place_value_10 () Real)
(declare-fun history1_c_place_value_11 () Real)
(declare-fun history1_c_place_value_12 () Real)
(declare-fun block3true_4 () Bool)
(declare-fun block4false_4 () Bool)
(assert (= history1_c_place_value_0 1.0))

(assert (= history1_c_place_value_1 (+ history1_c_place_value_0 history1_c_place_value_0)))
(assert (= history1_c_place_value_2 1.0))
(assert (ite (> history1_c_place_value_0 0.0) (and (= block3true_1 true) (= block4false_1 false) (= history1_c_place_value_3 history1_c_place_value_1)) (and (= block3true_1 false) (= block4false_1 true) (= history1_c_place_value_3 history1_c_place_value_2))))
(assert (or (and block3true_1
(not block4false_1))
(and (not block3true_1)
block4false_1)))


(assert (= history1_c_place_value_4 (+ history1_c_place_value_3 history1_c_place_value_0)))
(assert (= history1_c_place_value_5 1.0))
(assert (ite (> history1_c_place_value_3 0.0) (and (= block3true_2 true) (= block4false_2 false) (= history1_c_place_value_6 history1_c_place_value_4)) (and (= block3true_2 false) (= block4false_2 true) (= history1_c_place_value_6 history1_c_place_value_5))))
(assert (or (and block3true_2
(not block4false_2))
(and (not block3true_2)
block4false_2)))


(assert (= history1_c_place_value_7 (+ history1_c_place_value_6 history1_c_place_value_3)))
(assert (= history1_c_place_value_8 1.0))
(assert (ite (> history1_c_place_value_6 0.0) (and (= block3true_3 true) (= block4false_3 false) (= history1_c_place_value_9 history1_c_place_value_7)) (and (= block3true_3 false) (= block4false_3 true) (= history1_c_place_value_9 history1_c_place_value_8))))
(assert (or (and block3true_3
(not block4false_3))
(and (not block3true_3)
block4false_3)))


(assert (= history1_c_place_value_10 (+ history1_c_place_value_9 history1_c_place_value_6)))
(assert (= history1_c_place_value_11 1.0))
(assert (ite (> history1_c_place_value_9 0.0) (and (= block3true_4 true) (= block4false_4 false) (= history1_c_place_value_12 history1_c_place_value_10)) (and (= block3true_4 false) (= block4false_4 true) (= history1_c_place_value_12 history1_c_place_value_11))))
(assert (or (and block3true_4
(not block4false_4))
(and (not block3true_4)
block4false_4)))