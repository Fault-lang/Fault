(set-logic QF_NRA)
(declare-fun multicond_t_base_cond_0 () Real)
(declare-fun multicond_t_base_value_0 () Real)
(declare-fun multicond_t_base_value_1 () Real)
(declare-fun multicond_t_base_value_2 () Real)
(declare-fun multicond_t_base_value_3 () Real)
(declare-fun block216true_0 () Bool)
(declare-fun block216false_0 () Bool)
(declare-fun multicond_t_base_value_4 () Real)
(declare-fun block214true_0 () Bool)
(declare-fun block214false_0 () Bool)
(declare-fun multicond_t_base_cond_1 () Real)
(declare-fun multicond_t_base_value_5 () Real)
(declare-fun multicond_t_base_cond_2 () Real)
(declare-fun block212true_0 () Bool)
(declare-fun block212false_0 () Bool)
(declare-fun multicond_t_base_value_6 () Real)
(declare-fun multicond_t_base_value_7 () Real)
(declare-fun multicond_t_base_value_8 () Real)
(declare-fun multicond_t_base_value_9 () Real)
(declare-fun multicond_t_base_cond_3 () Real)
(declare-fun multicond_t_base_cond_4 () Real)
(declare-fun multicond_t_base_value_10 () Real)
(assert (= multicond_t_base_cond_0 1.0))
(assert (= multicond_t_base_value_0 10.0))

(assert (= multicond_t_base_value_1 (+ multicond_t_base_value_0 20.0)))
(assert (= multicond_t_base_value_2 0.0))

(assert (ite (> multicond_t_base_value_1 100.0) (= block216true_0 (= multicond_t_base_value_3 multicond_t_base_value_2)) (= block216false_0 (= multicond_t_base_value_3 multicond_t_base_value_1))))
(assert (or (and block216true_0
(not block216false_0))
(and (not block216true_0)
block216false_0)))

(assert (ite (< multicond_t_base_cond_0 4.0) (= block214true_0 (= multicond_t_base_value_4 multicond_t_base_value_3)) (= block214false_0 (= multicond_t_base_value_4 multicond_t_base_value_0))))
(assert (= multicond_t_base_cond_1 (+ multicond_t_base_cond_0 1.0)))
(assert (or (and block214true_0
(not block214false_0))
(and (not block214true_0)
block214false_0)))

(assert (ite (> multicond_t_base_cond_0 0.0) (= block212true_0 (and (= multicond_t_base_value_5 multicond_t_base_value_4) (= multicond_t_base_cond_2 multicond_t_base_cond_1))) (= block212false_0 (and (= multicond_t_base_value_5 multicond_t_base_value_0)
(= multicond_t_base_cond_2 multicond_t_base_cond_0)))))
(assert (or (and block212true_0
(not block212false_0))
(and (not block212true_0)
block212false_0)))


(assert (= multicond_t_base_value_6 (+ multicond_t_base_value_5 20.0)))
(assert (= multicond_t_base_value_7 0.0))

(assert (ite (> multicond_t_base_value_6 100.0) (= block216true_0 (= multicond_t_base_value_8 multicond_t_base_value_7)) (= block216false_0 (= multicond_t_base_value_8 multicond_t_base_value_6))))
(assert (or (and block216true_0
(not block216false_0))
(and (not block216true_0)
block216false_0)))

(assert (ite (< multicond_t_base_cond_2 4.0) (= block214true_0 (= multicond_t_base_value_9 multicond_t_base_value_8)) (= block214false_0 (= multicond_t_base_value_9 multicond_t_base_value_5))))
(assert (= multicond_t_base_cond_3 (+ multicond_t_base_cond_2 1.0)))
(assert (or (and block214true_0
(not block214false_0))
(and (not block214true_0)
block214false_0)))

(assert (ite (> multicond_t_base_cond_2 0.0) (= block212true_0 (and (= multicond_t_base_cond_4 multicond_t_base_cond_3) (= multicond_t_base_value_10 multicond_t_base_value_9))) (= block212false_0 (and (= multicond_t_base_cond_4 multicond_t_base_cond_2)
(= multicond_t_base_value_10 multicond_t_base_value_5)))))
(assert (or (and block212true_0
(not block212false_0))
(and (not block212true_0)
block212false_0)))