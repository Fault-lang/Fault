(set-logic QF_NRA)
(declare-fun ador_a_choice_0 () Bool)
(declare-fun ador_a_option1_0 () Bool)
(declare-fun ador_a_option2_0 () Bool)
(declare-fun ador_a_choice_1 () Bool)
(declare-fun ador_a_option1_1 () Bool)
(declare-fun ador_a_option2_1 () Bool)
(declare-fun ador_a_option1_2 () Bool)
(declare-fun ador_a_option2_2 () Bool)
(declare-fun ador_a_choice__state-%8_0 () Bool)
(declare-fun ador_a_choice__state-%8_1 () Bool)
(declare-fun ador_a_option1_3 () Bool)
(declare-fun ador_a_option2_3 () Bool)
(declare-fun block82true_0 () Bool)
(declare-fun block82false_0 () Bool)
(declare-fun ador_a_option1_4 () Bool)
(declare-fun ador_a_option1_5 () Bool)
(declare-fun block86true_0 () Bool)
(declare-fun block86false_0 () Bool)
(declare-fun ador_a_option2_4 () Bool)
(declare-fun ador_a_option2_5 () Bool)
(declare-fun block90true_0 () Bool)
(declare-fun block90false_0 () Bool)
(assert (= ador_a_choice_0 false))
(assert (= ador_a_option1_0 false))
(assert (= ador_a_option2_0 false))
(assert (= ador_a_choice_1 true))

(assert (or (and ador_a_choice__state-%8_0
(not ador_a_choice__state-%8_1))
(and (not ador_a_choice__state-%8_0)
ador_a_choice__state-%8_1)))

(assert (ite (= ador_a_choice_1 true) (= block82true_0 (and (= ador_a_option1_3 ador_a_option1_2) (= ador_a_option2_3 ador_a_option2_2))) (= block82false_0 (and (= ador_a_option1_3 ador_a_option1_0)
(= ador_a_option2_3 ador_a_option2_0)))))
(assert (or (and block82true_0
(not block82false_0))
(and (not block82true_0)
block82false_0)))


(assert (= ador_a_option1_4 true))

(assert (ite (= ador_a_option1_3 true) (= block86true_0 (= ador_a_option1_5 ador_a_option1_4)) (= block86false_0 (= ador_a_option1_5 ador_a_option1_3))))
(assert (or (and block86true_0
(not block86false_0))
(and (not block86true_0)
block86false_0)))


(assert (= ador_a_option2_4 true))

(assert (ite (= ador_a_option2_3 true) (= block90true_0 (= ador_a_option2_5 ador_a_option2_4)) (= block90false_0 (= ador_a_option2_5 ador_a_option2_3))))
(assert (or (and block90true_0
(not block90false_0))
(and (not block90true_0)
block90false_0)))