(set-logic QF_NRA)
(declare-fun simple_l_active_0 () Bool)
(declare-fun simple_l_vault_value_0 () Real)
(declare-fun simple_l_vault_value_2 () Real)
(declare-fun simple_l_vault_value_1 () Real)
(assert 
    (= simple_l_active_0 false))
(assert
    (= simple_l_vault_value_0 30.0))
(assert
    (= simple_l_vault_value_1 (+ simple_l_vault_value_0 (- simple_l_vault_value_0 2.0))))
(assert
    (ite (> simple_l_vault_value_0 4.0)
    (= simple_l_vault_value_2 simple_l_vault_value_1)
    (= simple_l_vault_value_2 simple_l_vault_value_0)))
