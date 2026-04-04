(set-logic QF_NRA)
(declare-fun increment_f_place_value_0 () Real)
(declare-fun increment_f_place_value_1 () Real)
(declare-fun increment_f_place_value_2 () Real)
(declare-fun increment_f_place_value_3 () Real)
(declare-fun block3true_1 () Bool)
(declare-fun block4false_1 () Bool)
(declare-fun increment_f_place_value_4 () Real)
(declare-fun increment_f_place_value_5 () Real)
(declare-fun increment_f_place_value_6 () Real)
(declare-fun block3true_2 () Bool)
(declare-fun block4false_2 () Bool)
(declare-fun increment_f_place_value_7 () Real)
(declare-fun increment_f_place_value_8 () Real)
(declare-fun increment_f_place_value_9 () Real)
(declare-fun block3true_3 () Bool)
(declare-fun block4false_3 () Bool)
(declare-fun increment_f_place_value_10 () Real)
(declare-fun increment_f_place_value_11 () Real)
(declare-fun increment_f_place_value_12 () Real)
(declare-fun block3true_4 () Bool)
(declare-fun block4false_4 () Bool)
(declare-fun increment_f_place_value_13 () Real)
(declare-fun increment_f_place_value_14 () Real)
(declare-fun increment_f_place_value_15 () Real)
(declare-fun block3true_5 () Bool)
(declare-fun block4false_5 () Bool)
(assert (= increment_f_place_value_0 0.0))

(assert (= increment_f_place_value_1 1.0))
(assert (= increment_f_place_value_2 (+ increment_f_place_value_0 increment_f_place_value_0)))
(assert (ite (= increment_f_place_value_0 0.0) (and (= block3true_1 true) (= block4false_1 false) (= increment_f_place_value_3 increment_f_place_value_1)) (and (= block3true_1 false) (= block4false_1 true) (= increment_f_place_value_3 increment_f_place_value_2))))
(assert (or (and block3true_1
(not block4false_1))
(and (not block3true_1)
block4false_1)))


(assert (= increment_f_place_value_4 1.0))
(assert (= increment_f_place_value_5 (+ increment_f_place_value_3 increment_f_place_value_0)))
(assert (ite (= increment_f_place_value_3 0.0) (and (= block3true_2 true) (= block4false_2 false) (= increment_f_place_value_6 increment_f_place_value_4)) (and (= block3true_2 false) (= block4false_2 true) (= increment_f_place_value_6 increment_f_place_value_5))))
(assert (or (and block3true_2
(not block4false_2))
(and (not block3true_2)
block4false_2)))


(assert (= increment_f_place_value_7 1.0))
(assert (= increment_f_place_value_8 (+ increment_f_place_value_6 increment_f_place_value_3)))
(assert (ite (= increment_f_place_value_6 0.0) (and (= block3true_3 true) (= block4false_3 false) (= increment_f_place_value_9 increment_f_place_value_7)) (and (= block3true_3 false) (= block4false_3 true) (= increment_f_place_value_9 increment_f_place_value_8))))
(assert (or (and block3true_3
(not block4false_3))
(and (not block3true_3)
block4false_3)))


(assert (= increment_f_place_value_10 1.0))
(assert (= increment_f_place_value_11 (+ increment_f_place_value_9 increment_f_place_value_6)))
(assert (ite (= increment_f_place_value_9 0.0) (and (= block3true_4 true) (= block4false_4 false) (= increment_f_place_value_12 increment_f_place_value_10)) (and (= block3true_4 false) (= block4false_4 true) (= increment_f_place_value_12 increment_f_place_value_11))))
(assert (or (and block3true_4
(not block4false_4))
(and (not block3true_4)
block4false_4)))


(assert (= increment_f_place_value_13 1.0))
(assert (= increment_f_place_value_14 (+ increment_f_place_value_12 increment_f_place_value_9)))
(assert (ite (= increment_f_place_value_12 0.0) (and (= block3true_5 true) (= block4false_5 false) (= increment_f_place_value_15 increment_f_place_value_13)) (and (= block3true_5 false) (= block4false_5 true) (= increment_f_place_value_15 increment_f_place_value_14))))
(assert (or (and block3true_5
(not block4false_5))
(and (not block3true_5)
block4false_5)))