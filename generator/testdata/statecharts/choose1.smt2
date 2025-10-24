(set-logic QF_NRA)
(declare-fun choose1_a_choice_0 () Bool)
(declare-fun choose1_a_option1_0 () Bool)
(declare-fun choose1_a_option2_0 () Bool)
(declare-fun choose1_a_choice_1 () Bool)
(declare-fun choose1_a_option1_1 () Bool)
(declare-fun choose1_a_choice_2 () Bool)
(declare-fun choose1_a_option2_1 () Bool)
(declare-fun choose1_a_choice_3 () Bool)
(declare-fun choose1_a_option1_2 () Bool)
(declare-fun choose1_a_choice_4 () Bool)
(declare-fun choose1_a_option2_2 () Bool)
(declare-fun choose1_a_choice_5 () Bool)
(declare-fun choose1_a_choice_6 () Bool)
(declare-fun choose1_a_option1_3 () Bool)
(declare-fun choose1_a_option2_3 () Bool)
(declare-fun choose1_a_option2_4 () Bool)
(declare-fun choose1_a_choice_7 () Bool)
(declare-fun choose1_a_option1_4 () Bool)
(assert (= choose1_a_choice_0 false))
(assert (= choose1_a_option1_0 false))
(assert (= choose1_a_option2_0 false))
(assert (= choose1_a_choice_1 true))

(assert 
    (or 
        (and 
            (and 
                (= choose1_a_option1_1 true)
                (= choose1_a_choice_2 false))
            (and 
                (not (= choose1_a_option2_1 true))
                (not (= choose1_a_choice_3 false)))
            (= choose1_a_choice_6 choose1_a_choice_3)
            (= choose1_a_option1_3 choose1_a_option1_1)
            (= choose1_a_option2_3 choose1_a_option2_1))

        (and
            (and 
                (not (= choose1_a_option1_2 true))
                (not (= choose1_a_choice_4 false)))
            (and
                (= choose1_a_option2_2 true)
                (= choose1_a_choice_5 false))
        (= choose1_a_choice_6 choose1_a_choice_5)
        (= choose1_a_option1_3 choose1_a_option1_2)
        (= choose1_a_option2_3 choose1_a_option2_2))))

(assert 
    (ite
        (= choose1_a_choice_1 true)
        (and (= choose1_a_option2_4 choose1_a_option2_3) (= choose1_a_choice_7 choose1_a_choice_6) (= choose1_a_option1_4 choose1_a_option1_3))
        (and (= choose1_a_choice_7 choose1_a_choice_1) (= choose1_a_option1_4 choose1_a_option1_0) (= choose1_a_option2_4 choose1_a_option2_0))))