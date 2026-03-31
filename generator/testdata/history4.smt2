(set-logic QF_NRA)
(declare-fun history4_s_value_0 () Real)
(declare-fun history4_s_A_0 () Bool)
(declare-fun history4_s_B_0 () Bool)
(declare-fun history4_s_value_1 () Real)
(declare-fun history4_s_value_2 () Real)
(declare-fun block3true_1 () Bool)
(declare-fun block3false_1 () Bool)
(declare-fun history4_s_value_3 () Real)
(declare-fun history4_s_value_4 () Real)
(declare-fun block3true_2 () Bool)
(declare-fun block3false_2 () Bool)
(declare-fun history4_s_value_5 () Real)
(declare-fun history4_s_value_6 () Real)
(declare-fun block3true_3 () Bool)
(declare-fun block3false_3 () Bool)
(declare-fun history4_s_value_7 () Real)
(declare-fun history4_s_value_8 () Real)
(declare-fun block3true_4 () Bool)
(declare-fun block3false_4 () Bool)
(declare-fun history4_s_value_9 () Real)
(declare-fun history4_s_value_10 () Real)
(declare-fun block3true_5 () Bool)
(declare-fun block3false_5 () Bool)
(assert (= history4_s_value_0 1.0))
(assert (= history4_s_A_0 true))
(assert (= history4_s_B_0 true))

(assert (= history4_s_value_1 (+ history4_s_value_0 history4_s_value_0)))

(assert (ite (and (= history4_s_A_0 true) (= history4_s_B_0 true)) (and (= block3true_1 true) (= block3false_1 false) (= history4_s_value_2 history4_s_value_1)) (and (= block3true_1 false) (= block3false_1 true) (= history4_s_value_2 history4_s_value_0))))
(assert (or (and block3true_1
(not block3false_1))
(and (not block3true_1)
block3false_1)))


(assert (= history4_s_value_3 (+ history4_s_value_2 history4_s_value_0)))

(assert (ite (and (= history4_s_A_0 true) (= history4_s_B_0 true)) (and (= block3true_2 true) (= block3false_2 false) (= history4_s_value_4 history4_s_value_3)) (and (= block3true_2 false) (= block3false_2 true) (= history4_s_value_4 history4_s_value_2))))
(assert (or (and block3true_2
(not block3false_2))
(and (not block3true_2)
block3false_2)))


(assert (= history4_s_value_5 (+ history4_s_value_4 history4_s_value_2)))

(assert (ite (and (= history4_s_A_0 true) (= history4_s_B_0 true)) (and (= block3true_3 true) (= block3false_3 false) (= history4_s_value_6 history4_s_value_5)) (and (= block3true_3 false) (= block3false_3 true) (= history4_s_value_6 history4_s_value_4))))
(assert (or (and block3true_3
(not block3false_3))
(and (not block3true_3)
block3false_3)))


(assert (= history4_s_value_7 (+ history4_s_value_6 history4_s_value_4)))

(assert (ite (and (= history4_s_A_0 true) (= history4_s_B_0 true)) (and (= block3true_4 true) (= block3false_4 false) (= history4_s_value_8 history4_s_value_7)) (and (= block3true_4 false) (= block3false_4 true) (= history4_s_value_8 history4_s_value_6))))
(assert (or (and block3true_4
(not block3false_4))
(and (not block3true_4)
block3false_4)))


(assert (= history4_s_value_9 (+ history4_s_value_8 history4_s_value_6)))

(assert (ite (and (= history4_s_A_0 true) (= history4_s_B_0 true)) (and (= block3true_5 true) (= block3false_5 false) (= history4_s_value_10 history4_s_value_9)) (and (= block3true_5 false) (= block3false_5 true) (= history4_s_value_10 history4_s_value_8))))
(assert (or (and block3true_5
(not block3false_5))
(and (not block3true_5)
block3false_5)))