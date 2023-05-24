(set-logic QF_NRA)
(declare-fun test_str1_0 () Bool)
(declare-fun test_str2_0 () Bool)
(declare-fun test_str3_0 () Bool)
(declare-fun test_str1_neg_0 () Bool)
(declare-fun test_str4_0 () Bool)
(assert (= test_str1_neg_0 (not test_str1_0)))
(assert (= test_str4_0 (and test_str2_0 test_str1_neg_0)))
(assert (not test_str3_0))
(assert (or (and test_str1_0 test_str3_0) test_str4_0))