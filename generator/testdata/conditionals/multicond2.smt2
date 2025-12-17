(set-logic QF_NRA)
(declare-fun multicond_t_base_cond_0 () Real)
(declare-fun multicond_t_base_value_0 () Real)
(declare-fun multicond_t_base_value_1 () Real)
(declare-fun multicond_t_base_value_2 () Real)
(declare-fun block192true_0 () Bool)
(declare-fun block192false_0 () Bool)
(assert (= multicond_t_base_cond_0 1.0))
(assert (= multicond_t_base_value_0 10.0))

(assert (= multicond_t_base_value_1 (+ multicond_t_base_value_0 20.0)))

(assert (ite (and (> multicond_t_base_cond_0 0.0) (< multicond_t_base_cond_0 4.0)) (= block192true_0 (= multicond_t_base_value_2 multicond_t_base_value_1)) (= block192false_0 (= multicond_t_base_value_2 multicond_t_base_value_0))))
(assert (or (and block192true_0
(not block192false_0))
(and (not block192true_0)
block192false_0)))
