(set-logic QF_NRA)
(declare-fun renamed_l_active_0 () Bool)
(declare-fun renamed_l_vault_value_0 () Real)
(declare-fun renamed_l_vault_value_1 () Real)
(declare-fun renamed_l_vault_value_2 () Real)
(declare-fun block59true_0 () Bool)
(declare-fun block59false_0 () Bool)
(assert (= renamed_l_active_0 false))
(assert (= renamed_l_vault_value_0 30.0))

(assert (= renamed_l_vault_value_1 (+ renamed_l_vault_value_0 (- renamed_l_vault_value_0 2.0))))

(assert (ite (> renamed_l_vault_value_0 4.0) (= block59true_0 (= renamed_l_vault_value_2 renamed_l_vault_value_1)) (= block59false_0 (= renamed_l_vault_value_2 renamed_l_vault_value_0))))
(assert (or (and block59true_0
(not block59false_0))
(and (not block59true_0)
block59false_0)))