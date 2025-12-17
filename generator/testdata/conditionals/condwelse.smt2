(set-logic QF_NRA)
(declare-fun condwe_t_base_cond_0 () Real)
(declare-fun condwe_t_base_value_0 () Real)
(declare-fun condwe_t_base_value_1 () Real)
(declare-fun condwe_t_base_value_2 () Real)
(declare-fun block220true_0 () Bool)
(declare-fun block220false_0 () Bool)
(declare-fun condwe_t_base_value_3 () Real)
(declare-fun condwe_t_base_value_4 () Real)
(declare-fun block222true_0 () Bool)
(declare-fun block222false_0 () Bool)
(assert (= condwe_t_base_cond_0 1.0))
(assert (= condwe_t_base_value_0 10.0))

(assert (= condwe_t_base_value_1 (- condwe_t_base_value_0 30.0)))

(assert (ite (> condwe_t_base_cond_0 0.0) (= block220true_0 (= condwe_t_base_value_2 condwe_t_base_value_1)) (= block220false_0 (= condwe_t_base_value_2 condwe_t_base_value_0))))
(assert (= condwe_t_base_value_3 (+ condwe_t_base_value_2 20.0)))

(assert (ite (and (> condwe_t_base_cond_0 0.0) (< condwe_t_base_cond_0 4.0)) (= block222true_0 (= condwe_t_base_value_4 condwe_t_base_value_3)) (= block222false_0 (= condwe_t_base_value_4 condwe_t_base_value_2))))
(assert (or (and block222true_0
(not block222false_0))
(and (not block222true_0)
block222false_0)))
(assert (or (and block220true_0
(not block220false_0))
(and (not block220true_0)
block220false_0)))
