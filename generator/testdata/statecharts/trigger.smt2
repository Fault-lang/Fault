(set-logic QF_NRA)
(declare-fun trigger_fl_vault_value_0 () Real)
(declare-fun trigger_x_foo_0 () Bool)
(declare-fun trigger_x_bar_0 () Bool)
(declare-fun trigger_x_bash_0 () Bool)
(declare-fun trigger_x_foo_1 () Bool)
(declare-fun trigger_fl_vault_value_1 () Real)
(declare-fun trigger_fl_vault_value_2 () Real)
(declare-fun block5true_0 () Bool)
(declare-fun block5false_0 () Bool)
(declare-fun trigger_fl_vault_value_3 () Real)
(declare-fun block3true_0 () Bool)
(declare-fun block3false_0 () Bool)
(declare-fun trigger_fl_vault_value_4 () Real)
(declare-fun trigger_fl_vault_value_5 () Real)
(declare-fun block10true_0 () Bool)
(declare-fun block10false_0 () Bool)
(declare-fun trigger_fl_vault_value_6 () Real)
(declare-fun block8true_0 () Bool)
(declare-fun block8false_0 () Bool)
(declare-fun trigger_fl_vault_value_7 () Real)
(declare-fun trigger_fl_vault_value_8 () Real)
(declare-fun block15true_0 () Bool)
(declare-fun block15false_0 () Bool)
(declare-fun trigger_fl_vault_value_9 () Real)
(declare-fun block13true_0 () Bool)
(declare-fun block13false_0 () Bool)
(declare-fun trigger_x_foo_2 () Bool)
(declare-fun trigger_x_foo_3 () Bool)
(declare-fun block17true_0 () Bool)
(declare-fun block17false_0 () Bool)
(declare-fun simple_active_0 () Bool)
(assert (= simple_active_0 true))

(assert (= trigger_fl_vault_value_0 30.0))
(assert (= trigger_x_foo_0 false))
(assert (= trigger_x_bar_0 false))
(assert (= trigger_x_bash_0 false))
(assert (= trigger_x_foo_1 true))

(assert (= trigger_fl_vault_value_1 (+ trigger_fl_vault_value_0 (- trigger_fl_vault_value_0 2.0))))

(assert (ite (> trigger_fl_vault_value_0 4.0) (and (= block5true_0 true) (= block5false_0 false) (= trigger_fl_vault_value_2 trigger_fl_vault_value_1)) (and (= block5true_0 false) (= block5false_0 true) (= trigger_fl_vault_value_2 trigger_fl_vault_value_0))))
(assert (or (and block5true_0
(not block5false_0))
(and (not block5true_0)
block5false_0)))

(assert (ite (= trigger_x_foo_1 true) (and (= block3true_0 true) (= block3false_0 false) (= trigger_fl_vault_value_3 trigger_fl_vault_value_2)) (and (= block3true_0 false) (= block3false_0 true) (= trigger_fl_vault_value_3 trigger_fl_vault_value_0))))
(assert (or (and block3true_0
(not block3false_0))
(and (not block3true_0)
block3false_0)))


(assert (= trigger_fl_vault_value_4 (+ trigger_fl_vault_value_3 (- trigger_fl_vault_value_3 2.0))))

(assert (ite (> trigger_fl_vault_value_3 4.0) (and (= block10true_0 true) (= block10false_0 false) (= trigger_fl_vault_value_5 trigger_fl_vault_value_4)) (and (= block10true_0 false) (= block10false_0 true) (= trigger_fl_vault_value_5 trigger_fl_vault_value_3))))
(assert (or (and block10true_0
(not block10false_0))
(and (not block10true_0)
block10false_0)))

(assert (ite (and (= trigger_x_bar_0 true) (= simple_active_0 true)) (and (= block8true_0 true) (= block8false_0 false) (= trigger_fl_vault_value_6 trigger_fl_vault_value_5)) (and (= block8true_0 false) (= block8false_0 true) (= trigger_fl_vault_value_6 trigger_fl_vault_value_3))))
(assert (or (and block8true_0
(not block8false_0))
(and (not block8true_0)
block8false_0)))


(assert (= trigger_fl_vault_value_7 (+ trigger_fl_vault_value_6 (- trigger_fl_vault_value_6 2.0))))

(assert (ite (> trigger_fl_vault_value_6 4.0) (and (= block15true_0 true) (= block15false_0 false) (= trigger_fl_vault_value_8 trigger_fl_vault_value_7)) (and (= block15true_0 false) (= block15false_0 true) (= trigger_fl_vault_value_8 trigger_fl_vault_value_6))))
(assert (or (and block15true_0
(not block15false_0))
(and (not block15true_0)
block15false_0)))

(assert (ite (= trigger_x_bash_0 true) (and (= block13true_0 true) (= block13false_0 false) (= trigger_fl_vault_value_9 trigger_fl_vault_value_8)) (and (= block13true_0 false) (= block13false_0 true) (= trigger_fl_vault_value_9 trigger_fl_vault_value_6))))
(assert (= trigger_x_foo_2 true))

(assert (ite (and (= trigger_x_bash_0 true) (= simple_active_0 true)) (and (= block17true_0 true) (= block17false_0 false) (= trigger_x_foo_3 trigger_x_foo_2)) (and (= block17true_0 false) (= block17false_0 true) (= trigger_x_foo_3 trigger_x_foo_1))))
(assert (or (and block17true_0
(not block17false_0))
(and (not block17true_0)
block17false_0)))
(assert (or (and block13true_0
(not block13false_0))
(and (not block13true_0)
block13false_0)))

