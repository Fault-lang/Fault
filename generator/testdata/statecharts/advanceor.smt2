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
(declare-fun ador_a_option2_3 () Bool)
(declare-fun ador_a_option1_3 () Bool)
(declare-fun block22true_1 () Bool)
(declare-fun block22false_1 () Bool)
(declare-fun ador_a_option1_4 () Bool)
(declare-fun ador_a_option1_5 () Bool)
(declare-fun block26true_1 () Bool)
(declare-fun block26false_1 () Bool)
(declare-fun ador_a_option2_4 () Bool)
(declare-fun ador_a_option2_5 () Bool)
(declare-fun block30true_1 () Bool)
(declare-fun block30false_1 () Bool)
(assert (= ador_a_choice_0 false))
(assert (= ador_a_option1_0 false))
(assert (= ador_a_option2_0 false))
(assert (= ador_a_choice_1 true))

(assert (=> ador_a_choice__state-%8_0 (= ador_a_option1_1 true)))
(assert (=> ador_a_choice__state-%8_1 (= ador_a_option2_1 true)))
(assert (= ador_a_choice__state-%8_0 (and (= ador_a_option1_2 ador_a_option1_1)
(= ador_a_option2_2 ador_a_option2_0))))
(assert (= ador_a_choice__state-%8_1 (and (= ador_a_option2_2 ador_a_option2_1)
(= ador_a_option1_2 ador_a_option1_0))))
(assert (or (and ador_a_choice__state-%8_0
(not ador_a_choice__state-%8_1))
(and (not ador_a_choice__state-%8_0)
ador_a_choice__state-%8_1)))

(assert (ite (= ador_a_choice_1 true) (and (= block22true_1 true) (= block22false_1 false) (and (= ador_a_option2_3 ador_a_option2_2) (= ador_a_option1_3 ador_a_option1_2))) (and (= block22true_1 false) (= block22false_1 true) (and (= ador_a_option2_3 ador_a_option2_0)
(= ador_a_option1_3 ador_a_option1_0)))))
(assert (or (and block22true_1
(not block22false_1))
(and (not block22true_1)
block22false_1)))


(assert (= ador_a_option1_4 true))

(assert (ite (= ador_a_option1_3 true) (and (= block26true_1 true) (= block26false_1 false) (= ador_a_option1_5 ador_a_option1_4)) (and (= block26true_1 false) (= block26false_1 true) (= ador_a_option1_5 ador_a_option1_3))))
(assert (or (and block26true_1
(not block26false_1))
(and (not block26true_1)
block26false_1)))


(assert (= ador_a_option2_4 true))

(assert (ite (= ador_a_option2_3 true) (and (= block30true_1 true) (= block30false_1 false) (= ador_a_option2_5 ador_a_option2_4)) (and (= block30true_1 false) (= block30false_1 true) (= ador_a_option2_5 ador_a_option2_3))))
(assert (or (and block30true_1
(not block30false_1))
(and (not block30true_1)
block30false_1)))
