(set-logic QF_NRA)
(declare-fun increment_f_value_0 () Real)
(declare-fun increment_f_value_1 () Real)
(declare-fun increment_f_value_2 () Real)
(declare-fun increment_f_value_3 () Real)
(declare-fun block3true_0 () Bool)
(declare-fun block4false_0 () Bool)
(declare-fun increment_f_value_4 () Real)
(declare-fun increment_f_value_5 () Real)
(declare-fun increment_f_value_6 () Real)
(declare-fun increment_f_value_7 () Real)
(declare-fun increment_f_value_8 () Real)
(declare-fun increment_f_value_9 () Real)
(declare-fun increment_f_value_10 () Real)
(declare-fun increment_f_value_11 () Real)
(declare-fun increment_f_value_12 () Real)
(declare-fun increment_f_value_13 () Real)
(declare-fun increment_f_value_14 () Real)
(declare-fun increment_f_value_15 () Real)
(assert (= increment_f_value_0 0.0))

(assert (= increment_f_value_1 1.0))
(assert (= increment_f_value_2 (+ increment_f_value_0 increment_f_value_0)))
(assert (ite (= increment_f_value_0 0.0) (and (= block3true_0 true) (= block4false_0 false) (= increment_f_value_3 increment_f_value_1)) (and (= block3true_0 false) (= block4false_0 true) (= increment_f_value_3 increment_f_value_2))))
(assert (or (and block3true_0
(not block4false_0))
(and (not block3true_0)
block4false_0)))


(assert (= increment_f_value_4 1.0))
(assert (= increment_f_value_5 (+ increment_f_value_3 increment_f_value_3)))
(assert (ite (= increment_f_value_3 0.0) (and (= block3true_0 true) (= block4false_0 false) (= increment_f_value_6 increment_f_value_4)) (and (= block3true_0 false) (= block4false_0 true) (= increment_f_value_6 increment_f_value_5))))
(assert (or (and block3true_0
(not block4false_0))
(and (not block3true_0)
block4false_0)))


(assert (= increment_f_value_7 1.0))
(assert (= increment_f_value_8 (+ increment_f_value_6 increment_f_value_6)))
(assert (ite (= increment_f_value_6 0.0) (and (= block3true_0 true) (= block4false_0 false) (= increment_f_value_9 increment_f_value_7)) (and (= block3true_0 false) (= block4false_0 true) (= increment_f_value_9 increment_f_value_8))))
(assert (or (and block3true_0
(not block4false_0))
(and (not block3true_0)
block4false_0)))


(assert (= increment_f_value_10 1.0))
(assert (= increment_f_value_11 (+ increment_f_value_9 increment_f_value_9)))
(assert (ite (= increment_f_value_9 0.0) (and (= block3true_0 true) (= block4false_0 false) (= increment_f_value_12 increment_f_value_10)) (and (= block3true_0 false) (= block4false_0 true) (= increment_f_value_12 increment_f_value_11))))
(assert (or (and block3true_0
(not block4false_0))
(and (not block3true_0)
block4false_0)))


(assert (= increment_f_value_13 1.0))
(assert (= increment_f_value_14 (+ increment_f_value_12 increment_f_value_12)))
(assert (ite (= increment_f_value_12 0.0) (and (= block3true_0 true) (= block4false_0 false) (= increment_f_value_15 increment_f_value_13)) (and (= block3true_0 false) (= block4false_0 true) (= increment_f_value_15 increment_f_value_14))))
(assert (or (and block3true_0
(not block4false_0))
(and (not block3true_0)
block4false_0)))