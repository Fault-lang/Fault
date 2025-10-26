(set-logic QF_NRA)
(declare-fun statechart_fl_active_0 () Bool)
(declare-fun statechart_fl_vault_value_0 () Real)
(declare-fun statechart_drain_initial_0 () Bool)
(declare-fun statechart_drain_open_0 () Bool)
(declare-fun statechart_drain_close_0 () Bool)
(declare-fun statechart_drain_initial_1 () Bool)
(declare-fun statechart_fl_vault_value_1 () Real)
(declare-fun statechart_fl_vault_value_2 () Real)
(declare-fun statechart_fl_vault_value_3 () Real)
(declare-fun statechart_drain_open_1 () Bool)
(declare-fun statechart_drain_open_2 () Bool)
(declare-fun statechart_drain_close_1 () Bool)
(declare-fun statechart_drain_close_2 () Bool)
(declare-fun statechart_drain_close_3 () Bool)
(declare-fun statechart_drain_close_4 () Bool)
(declare-fun statechart_fl_vault_value_4 () Real)
(declare-fun statechart_fl_vault_value_5 () Real)
(declare-fun statechart_fl_vault_value_6 () Real)
(declare-fun statechart_drain_open_3 () Bool)
(declare-fun statechart_drain_open_4 () Bool)
(declare-fun statechart_drain_close_5 () Bool)
(declare-fun statechart_drain_close_6 () Bool)
(declare-fun statechart_drain_close_7 () Bool)
(declare-fun statechart_drain_close_8 () Bool)
(assert (= statechart_fl_active_0 false))
(assert (= statechart_fl_vault_value_0 30.0))
(assert (= statechart_drain_initial_0 false))
(assert (= statechart_drain_open_0 false))
(assert (= statechart_drain_close_0 false))
(assert (= statechart_drain_initial_1 true))
(assert (= statechart_fl_vault_value_1 (+ statechart_fl_vault_value_0 (- statechart_fl_vault_value_0 2.0))))

(assert (ite (> statechart_fl_vault_value_0 4.0) (= statechart_fl_vault_value_2 statechart_fl_vault_value_1) (= statechart_fl_vault_value_2 statechart_fl_vault_value_0)))


(assert (ite (not statechart_drain_close_0) (= statechart_fl_vault_value_3 statechart_fl_vault_value_2) (= statechart_fl_vault_value_3 statechart_fl_vault_value_0)))
(assert (= statechart_drain_open_1 true))

(assert (ite (and (= statechart_drain_initial_1 true) (not statechart_fl_active_0)) (= statechart_drain_open_2 statechart_drain_open_1) (= statechart_drain_open_2 statechart_drain_open_0)))

(assert (= statechart_drain_close_1 true))

(assert (ite (and (= statechart_drain_open_2 true) (< statechart_fl_vault_value_3 0.0)) (= statechart_drain_close_2 statechart_drain_close_1) (= statechart_drain_close_2 statechart_drain_close_0)))

(assert (= statechart_drain_close_3 true))

(assert (ite (= statechart_drain_close_2 true) (= statechart_drain_close_4 statechart_drain_close_3) (= statechart_drain_close_4 statechart_drain_close_2)))

(assert (= statechart_fl_vault_value_4 (+ statechart_fl_vault_value_3 (- statechart_fl_vault_value_3 2.0))))

(assert (ite (> statechart_fl_vault_value_3 4.0) (= statechart_fl_vault_value_5 statechart_fl_vault_value_4) (= statechart_fl_vault_value_5 statechart_fl_vault_value_3)))


(assert (ite (not statechart_drain_close_4) (= statechart_fl_vault_value_6 statechart_fl_vault_value_5) (= statechart_fl_vault_value_6 statechart_fl_vault_value_3)))
(assert (= statechart_drain_open_3 true))

(assert (ite (and (= statechart_drain_initial_1 true) (not statechart_fl_active_0)) (= statechart_drain_open_4 statechart_drain_open_3) (= statechart_drain_open_4 statechart_drain_open_2)))

(assert (= statechart_drain_close_5 true))

(assert (ite (and (= statechart_drain_open_4 true) (< statechart_fl_vault_value_6 0.0)) (= statechart_drain_close_6 statechart_drain_close_5) (= statechart_drain_close_6 statechart_drain_close_4)))

(assert (= statechart_drain_close_7 true))

(assert (ite (= statechart_drain_close_6 true) (= statechart_drain_close_8 statechart_drain_close_7) (= statechart_drain_close_8 statechart_drain_close_6)))