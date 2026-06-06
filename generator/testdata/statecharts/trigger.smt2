(set-logic QF_NRA)
(declare-fun trigger_fl_vault_value_0 () Real)
(declare-fun trigger_x_foo_0 () Bool)
(declare-fun trigger_x_bar_0 () Bool)
(declare-fun trigger_x_bash_0 () Bool)
(declare-fun trigger_x_foo_1 () Bool)
(declare-fun trigger_fl_vault_value_1 () Real)
(declare-fun trigger_fl_vault_value_2 () Real)
(declare-fun block85true_1 () Bool)
(declare-fun block85false_1 () Bool)
(declare-fun trigger_fl_vault_value_3 () Real)
(declare-fun block83true_1 () Bool)
(declare-fun block83false_1 () Bool)
(declare-fun trigger_fl_vault_value_4 () Real)
(declare-fun trigger_fl_vault_value_5 () Real)
(declare-fun block90true_1 () Bool)
(declare-fun block90false_1 () Bool)
(declare-fun trigger_fl_vault_value_6 () Real)
(declare-fun block88true_1 () Bool)
(declare-fun block88false_1 () Bool)
(declare-fun trigger_fl_vault_value_7 () Real)
(declare-fun trigger_fl_vault_value_8 () Real)
(declare-fun block95true_1 () Bool)
(declare-fun block95false_1 () Bool)
(declare-fun trigger_fl_vault_value_9 () Real)
(declare-fun block93true_1 () Bool)
(declare-fun block93false_1 () Bool)
(declare-fun trigger_x_foo_2 () Bool)
(declare-fun trigger_x_foo_3 () Bool)
(declare-fun block97true_1 () Bool)
(declare-fun block97false_1 () Bool)
(declare-fun simple_active_0 () Bool)
(assert (= simple_active_0 true))

(assert (= trigger_fl_vault_value_0 30.0))
(assert (= trigger_x_foo_0 false))
(assert (= trigger_x_bar_0 false))
(assert (= trigger_x_bash_0 false))
(assert (= trigger_x_foo_1 true))

(assert (= trigger_fl_vault_value_1 (+ trigger_fl_vault_value_0 (- trigger_fl_vault_value_0 2.0))))

(assert (ite (> trigger_fl_vault_value_0 4.0) (and (= block85true_1 true) (= block85false_1 false) (= trigger_fl_vault_value_2 trigger_fl_vault_value_1)) (and (= block85true_1 false) (= block85false_1 true) (= trigger_fl_vault_value_2 trigger_fl_vault_value_0))))
(assert (or (and block85true_1
(not block85false_1))
(and (not block85true_1)
block85false_1)))

(assert (ite (= trigger_x_foo_1 true) (and (= block83true_1 true) (= block83false_1 false) (= trigger_fl_vault_value_3 trigger_fl_vault_value_2)) (and (= block83true_1 false) (= block83false_1 true) (= trigger_fl_vault_value_3 trigger_fl_vault_value_0))))
(assert (or (and block83true_1
(not block83false_1))
(and (not block83true_1)
block83false_1)))


(assert (= trigger_fl_vault_value_4 (+ trigger_fl_vault_value_3 (- trigger_fl_vault_value_3 2.0))))

(assert (ite (> trigger_fl_vault_value_3 4.0) (and (= block90true_1 true) (= block90false_1 false) (= trigger_fl_vault_value_5 trigger_fl_vault_value_4)) (and (= block90true_1 false) (= block90false_1 true) (= trigger_fl_vault_value_5 trigger_fl_vault_value_3))))
(assert (or (and block90true_1
(not block90false_1))
(and (not block90true_1)
block90false_1)))

(assert (ite (and (= trigger_x_bar_0 true) (= simple_active_0 true)) (and (= block88true_1 true) (= block88false_1 false) (= trigger_fl_vault_value_6 trigger_fl_vault_value_5)) (and (= block88true_1 false) (= block88false_1 true) (= trigger_fl_vault_value_6 trigger_fl_vault_value_3))))
(assert (or (and block88true_1
(not block88false_1))
(and (not block88true_1)
block88false_1)))


(assert (= trigger_fl_vault_value_7 (+ trigger_fl_vault_value_6 (- trigger_fl_vault_value_6 2.0))))

(assert (ite (> trigger_fl_vault_value_6 4.0) (and (= block95true_1 true) (= block95false_1 false) (= trigger_fl_vault_value_8 trigger_fl_vault_value_7)) (and (= block95true_1 false) (= block95false_1 true) (= trigger_fl_vault_value_8 trigger_fl_vault_value_6))))
(assert (or (and block95true_1
(not block95false_1))
(and (not block95true_1)
block95false_1)))

(assert (ite (= trigger_x_bash_0 true) (and (= block93true_1 true) (= block93false_1 false) (= trigger_fl_vault_value_9 trigger_fl_vault_value_8)) (and (= block93true_1 false) (= block93false_1 true) (= trigger_fl_vault_value_9 trigger_fl_vault_value_6))))
(assert (= trigger_x_foo_2 true))

(assert (ite (and (= trigger_x_bash_0 true) (= simple_active_0 true)) (and (= block97true_1 true) (= block97false_1 false) (= trigger_x_foo_3 trigger_x_foo_2)) (and (= block97true_1 false) (= block97false_1 true) (= trigger_x_foo_3 trigger_x_foo_1))))
(assert (or (and block97true_1
(not block97false_1))
(and (not block97true_1)
block97false_1)))
(assert (or (and block93true_1
(not block93false_1))
(and (not block93true_1)
block93false_1)))
