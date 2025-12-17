(set-logic QF_NRA)
(declare-fun trigger_fl_active_0 () Bool)
(declare-fun trigger_fl_vault_value_0 () Real)
(declare-fun trigger_x_foo_0 () Bool)
(declare-fun trigger_x_bar_0 () Bool)
(declare-fun trigger_x_bash_0 () Bool)
(declare-fun trigger_x_foo_1 () Bool)
(declare-fun trigger_fl_vault_value_1 () Real)
(declare-fun trigger_fl_vault_value_2 () Real)
(declare-fun block145true_0 () Bool)
(declare-fun block145false_0 () Bool)
(declare-fun trigger_fl_vault_value_3 () Real)
(declare-fun block143true_0 () Bool)
(declare-fun block143false_0 () Bool)
(declare-fun trigger_fl_vault_value_4 () Real)
(declare-fun trigger_fl_vault_value_5 () Real)
(declare-fun block150true_0 () Bool)
(declare-fun block150false_0 () Bool)
(declare-fun trigger_fl_vault_value_6 () Real)
(declare-fun block148true_0 () Bool)
(declare-fun block148false_0 () Bool)
(declare-fun trigger_fl_vault_value_7 () Real)
(declare-fun trigger_fl_vault_value_8 () Real)
(declare-fun block155true_0 () Bool)
(declare-fun block155false_0 () Bool)
(declare-fun trigger_fl_vault_value_9 () Real)
(declare-fun block153true_0 () Bool)
(declare-fun block153false_0 () Bool)
(declare-fun trigger_x_foo_2 () Bool)
(declare-fun trigger_x_foo_3 () Bool)
(declare-fun block157true_0 () Bool)
(declare-fun block157false_0 () Bool)
(assert (= trigger_fl_active_0 false))
(assert (= trigger_fl_vault_value_0 30.0))
(assert (= trigger_x_foo_0 false))
(assert (= trigger_x_bar_0 false))
(assert (= trigger_x_bash_0 false))
(assert (= trigger_x_foo_1 true))

(assert (= trigger_fl_vault_value_1 (+ trigger_fl_vault_value_0 (- trigger_fl_vault_value_0 2.0))))

(assert (ite (> trigger_fl_vault_value_0 4.0) (= block145true_0 (= trigger_fl_vault_value_2 trigger_fl_vault_value_1)) (= block145false_0 (= trigger_fl_vault_value_2 trigger_fl_vault_value_0))))
(assert (or (and block145true_0
(not block145false_0))
(and (not block145true_0)
block145false_0)))

(assert (ite (= trigger_x_foo_1 true) (= block143true_0 (= trigger_fl_vault_value_3 trigger_fl_vault_value_2)) (= block143false_0 (= trigger_fl_vault_value_3 trigger_fl_vault_value_0))))
(assert (or (and block143true_0
(not block143false_0))
(and (not block143true_0)
block143false_0)))


(assert (= trigger_fl_vault_value_4 (+ trigger_fl_vault_value_3 (- trigger_fl_vault_value_3 2.0))))

(assert (ite (> trigger_fl_vault_value_3 4.0) (= block150true_0 (= trigger_fl_vault_value_5 trigger_fl_vault_value_4)) (= block150false_0 (= trigger_fl_vault_value_5 trigger_fl_vault_value_3))))
(assert (or (and block150true_0
(not block150false_0))
(and (not block150true_0)
block150false_0)))

(assert (ite (and (= trigger_x_bar_0 true) (= trigger_fl_active_0 true)) (= block148true_0 (= trigger_fl_vault_value_6 trigger_fl_vault_value_5)) (= block148false_0 (= trigger_fl_vault_value_6 trigger_fl_vault_value_3))))
(assert (or (and block148true_0
(not block148false_0))
(and (not block148true_0)
block148false_0)))


(assert (= trigger_fl_vault_value_7 (+ trigger_fl_vault_value_6 (- trigger_fl_vault_value_6 2.0))))

(assert (ite (> trigger_fl_vault_value_6 4.0) (= block155true_0 (= trigger_fl_vault_value_8 trigger_fl_vault_value_7)) (= block155false_0 (= trigger_fl_vault_value_8 trigger_fl_vault_value_6))))
(assert (or (and block155true_0
(not block155false_0))
(and (not block155true_0)
block155false_0)))

(assert (ite (= trigger_x_bash_0 true) (= block153true_0 (= trigger_fl_vault_value_9 trigger_fl_vault_value_8)) (= block153false_0 (= trigger_fl_vault_value_9 trigger_fl_vault_value_6))))
(assert (= trigger_x_foo_2 true))

(assert (ite (and (= trigger_x_bash_0 true) (= trigger_fl_active_0 true)) (= block157true_0 (= trigger_x_foo_3 trigger_x_foo_2)) (= block157false_0 (= trigger_x_foo_3 trigger_x_foo_1))))
(assert (or (and block157true_0
(not block157false_0))
(and (not block157true_0)
block157false_0)))
(assert (or (and block153true_0
(not block153false_0))
(and (not block153true_0)
block153false_0)))