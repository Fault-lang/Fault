package util

import (
	"os"
	"testing"
)

func TestFilepath(t *testing.T) {
	var host string
	var ok bool

	if host, ok = os.LookupEnv("FAULT_HOST"); !ok {
		host = ""
	}
	os.Setenv("FAULT_HOST", "/host")
	filepath1 := "foo/test/file/system../test.spec"
	filepath1a := Filepath(filepath1)
	if filepath1a != "/host/foo/test/file/test.spec" {
		t.Fatalf("filepath not correct. want=/host/foo/test/file/test.spec got=%s", filepath1a)
	}

	filepath2 := "foo/test/file/system../../test.spec"
	filepath2a := Filepath(filepath2)
	if filepath2a != "/host/foo/test/test.spec" {
		t.Fatalf("filepath not correct. want=/host/foo/test/test.spec got=%s", filepath2a)
	}

	filepath3 := "foo/test/file/system../../../test.spec"
	filepath3a := Filepath(filepath3)
	if filepath3a != "/host/foo/test.spec" {
		t.Fatalf("filepath not correct. want=/host/foo/test.spec got=%s", filepath3a)
	}

	filepath4 := "foo/test/file/system/~/test.spec"
	filepath4a := Filepath(filepath4)
	if filepath4a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath4a)
	}

	filepath5 := "foo/test/file/system/~test.spec"
	filepath5a := Filepath(filepath5)
	if filepath5a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath5a)
	}

	filepath6 := "test.spec"
	filepath6a := Filepath(filepath6)
	if filepath6a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath6a)
	}

	filepath7 := "/.."
	filepath7a := Filepath(filepath7)
	if filepath7a != "/host" {
		t.Fatalf("filepath not correct. want=/host got=%s", filepath7a)
	}

	filepath8 := "/host/test.spec"
	filepath8a := Filepath(filepath8)
	if filepath8a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath8a)
	}

	filepath9 := "foo/../test.spec"
	filepath9a := Filepath(filepath9)
	if filepath9a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath9a)
	}

	filepath10 := "/foo/../test.spec"
	filepath10a := Filepath(filepath10)
	if filepath10a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath10a)
	}

	filepath11 := "../test.spec"
	filepath11a := Filepath(filepath11)
	if filepath11a != "test.spec" {
		t.Fatalf("filepath not correct. want=test.spec got=%s", filepath11a)
	}

	os.Setenv("FAULT_HOST", host)
}

func TestCartesian(t *testing.T) {
	list1 := []string{"a", "b", "c"}
	list2 := []string{"1", "2"}
	r := Cartesian(list1, list2)

	if r[0][0] != "a" || r[0][1] != "1" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[1][0] != "a" || r[1][1] != "2" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[2][0] != "b" || r[2][1] != "1" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[3][0] != "b" || r[3][1] != "2" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

}

func TestCartesianMulti(t *testing.T) {
	list1 := []string{"a", "b", "c"}
	list2 := []string{"1", "2"}
	list3 := []string{"3", "4"}
	r := CartesianMulti([][]string{list1, list2, list3})

	if r[0][0] != "a" || r[0][1] != "1" || r[0][2] != "3" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[1][0] != "a" || r[1][1] != "1" || r[1][2] != "4" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[2][0] != "a" || r[2][1] != "2" || r[2][2] != "3" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[3][0] != "a" || r[3][1] != "2" || r[3][2] != "4" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

}

func TestMergeStringMaps(t *testing.T) {
	m1 := make(map[string]string)
	m1["foo"] = "this"
	m1["bar"] = "is"

	m2 := make(map[string]string)
	m2["test"] = "test"

	m3 := MergeStringMaps(m1, m2)

	if len(m3) != 3 {
		t.Fatalf("merged string map has the wrong length got=%d", len(m3))
	}

	if m3["test"] != "test" || m3["foo"] != "this" {
		t.Fatalf("string map not merged correctly")
	}

}

func TestMergeStrSlices(t *testing.T) {
	sl1 := []string{"here", "there", "everywhere"}
	sl2 := []string{"here", "roy", "kent"}

	merged := MergeStrSlices(sl1, sl2)
	if merged[0] != "here" {
		t.Fatalf("first value of MergeStrSlices not correct. got=%s", merged[0])
	}
	if merged[1] != "there" {
		t.Fatalf("second value of MergeStrSlices not correct. got=%s", merged[1])
	}

	if merged[3] == "here" {
		t.Fatalf("duplicate value detected from MergeStrSlices. got=%s", merged[3])
	}

}

func TestCombinations(t *testing.T) {
	input := [][]string{{"a", "", "b"}, {"c", "", "d"}, {"h", "", "i"}, {"r", "", "s"}}
	expected := [][][]string{
		{
			{"a", "", "b"}, {"c", "", "d"},
		},
		{
			{"a", "", "b"}, {"h", "", "i"},
		},
		{
			{"a", "", "b"}, {"r", "", "s"},
		},
		{
			{"c", "", "d"}, {"h", "", "i"},
		},
		{
			{"c", "", "d"}, {"r", "", "s"},
		},
		{
			{"h", "", "i"}, {"r", "", "s"},
		},
	}

	results := Combinations(input, 2)
	if results[0][0][0] != expected[0][0][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[0][0][0], expected[0][0][0])
	}

	if results[0][1][0] != expected[0][1][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[0][1][0], expected[0][1][0])
	}

	if results[2][1][0] != expected[2][1][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[2][1][0], expected[2][1][0])
	}

	if results[2][1][1] != expected[2][1][1] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[2][1][0], expected[2][1][0])
	}

	if results[4][1][2] != expected[4][1][2] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[4][1][2], expected[4][1][2])
	}

	if results[5][0][2] != expected[5][0][2] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[5][0][2], expected[5][0][2])
	}
}

func TestMoreCombinations(t *testing.T) {
	input := [][]string{{"a", "b"}, {"c", "d"}, {"h", "i"}, {"r", "s"}}
	expected := [][][]string{
		{
			{"a", "b"}, {"c", "d"}, {"h", "i"},
		},
		{
			{"a", "b"}, {"h", "i"}, {"r", "s"},
		},
		{
			{"c", "d"}, {"h", "i"}, {"r", "s"},
		},
	}

	results := Combinations(input, 3)
	if results[0][0][0] != expected[0][0][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[0][0][0], expected[0][0][0])
	}

	if results[0][1][0] != expected[0][1][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[0][1][0], expected[0][1][0])
	}
	if results[0][2][0] != expected[0][2][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[0][2][0], expected[0][2][0])
	}

	if results[2][1][0] != expected[2][1][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[2][1][0], expected[2][1][0])
	}

	if results[2][1][1] != expected[2][1][1] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[2][1][1], expected[2][1][1])
	}

	if results[2][2][1] != expected[2][2][1] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[2][2][1], expected[2][2][1])
	}

}

func TestCombinationsN1(t *testing.T) {
	input := [][]string{{"a", "b"}, {"c", "d"}, {"h", "i"}, {"r", "s"}}
	expected := [][][]string{
		{
			{"a", "b"},
		},
		{
			{"c", "d"},
		},
		{
			{"h", "i"},
		},
		{
			{"r", "s"},
		},
	}

	results := Combinations(input, 1)
	if results[0][0][0] != expected[0][0][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[0][0][0], expected[0][0][0])
	}

	if results[1][0][0] != expected[1][0][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[1][0][0], expected[1][0][0])
	}

	if results[2][0][0] != expected[2][0][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[2][0][0], expected[2][0][0])
	}

	if results[3][0][0] != expected[3][0][0] {
		t.Fatalf("combinations not calculated correctly got=%s want=%s", results[3][0][0], expected[3][0][0])
	}

}

func TestStableSortKeys(t *testing.T) {
	test := []string{"this", "is", "a", "test", "ok?"}
	result := StableSortKeys(test)
	if len(result) != 5 || result[0] != "a" || result[1] != "is" || result[2] != "ok?" || result[3] != "test" || result[4] != "this" {
		t.Fatalf("StableSortKey returned the wrong result")
	}
}

func TestCaptureState(t *testing.T) {
	test1 := "this_is_a_test"

	r1, a1, c1 := CaptureState(test1)
	if r1 != "" || !a1 || c1 {
		t.Fatal("first test of CaptureState is incorrect")
	}

	test2 := "this_is_a_3"

	r2, a2, c2 := CaptureState(test2)
	if r2 != "3" || a2 || c2 {
		t.Fatal("second test of CaptureState is incorrect")
	}

	test3 := "this_is"

	r3, a3, c3 := CaptureState(test3)
	if r3 != "" || a3 || !c3 {
		t.Fatal("third test of CaptureState is incorrect")
	}
}

func TestCopy(t *testing.T) {
	test := []string{"here", "it", "is", "folks"}
	r := Copy(test)
	if len(test) != len(r) {
		t.Fatal("copy dropped a value")
	}
}

func TestMaxInt16(t *testing.T) {
	test := []int16{2, 6, 8, 4, 10, 8, 3}
	r := MaxInt16(test)
	if r != 10 {
		t.Fatal("MaxInt16 returned an incorrect value")
	}
}

func TestNotInSet(t *testing.T) {
	inputC := [][]string{{"a", "b"}, {"c", "d"}, {"h", "i"}, {"r", "s"}}
	iOn1 := [][]string{
		{"a", "b"}, {"c", "d"}, {"h", "i"},
	}

	e1 := [][]string{{"r", "s"}}

	r1 := NotInSet(iOn1, inputC)
	if r1[0][0] != e1[0][0] {
		t.Fatalf("incorrect value returned for NotInSet1 got=%s want=%s", r1[0][0], e1[0][0])
	}
	if r1[0][1] != e1[0][1] {
		t.Fatalf("incorrect value returned for NotInSet1 got=%s want=%s", r1[0][1], e1[0][1])
	}

	iOn2 := [][]string{
		{"a", "b"}, {"h", "i"}, {"r", "s"},
	}

	e2 := [][]string{{"c", "d"}}

	r2 := NotInSet(iOn2, inputC)
	if r2[0][0] != e2[0][0] {
		t.Fatalf("incorrect value returned for NotInSet2 got=%s want=%s", r2[0][0], e2[0][0])
	}
	if r2[0][1] != e2[0][1] {
		t.Fatalf("incorrect value returned for NotInSet2 got=%s want=%s", r2[0][1], e2[0][1])
	}

	iOn3 := [][]string{
		{"c", "d"}, {"h", "i"}, {"r", "s"},
	}

	e3 := [][]string{{"a", "b"}}

	r3 := NotInSet(iOn3, inputC)
	if r3[0][0] != e3[0][0] {
		t.Fatalf("incorrect value returned for NotInSet3 got=%s want=%s", r3[0][0], e3[0][0])
	}
	if r3[0][1] != e3[0][1] {
		t.Fatalf("incorrect value returned for NotInSet3 got=%s want=%s", r3[0][1], e3[0][1])
	}

}

func TestDetectMode(t *testing.T) {
	t1 := DetectMode("test.fspec")
	if t1 != "fspec" {
		t.Fatalf("incorrect value returned from DetectMode got=%s want=%s", t1, "fspec")
	}

	t2 := DetectMode("test.fsystem")
	if t2 != "fsystem" {
		t.Fatalf("incorrect value returned from DetectMode got=%s want=%s", t2, "fsystem")
	}

	t3 := DetectMode("test.mp4")
	if t3 != "" {
		t.Fatalf("incorrect value returned from DetectMode got=%s want=%s", t3, "")
	}
}

func TestImportTrail(t *testing.T) {
	it := ImportTrail{}
	it = it.PushSpec("test")
	it = it.PushSpec("this")
	it = it.PushSpec("trail")

	if len(it) != 3 {
		t.Fatal("specs not added to trail correctly")
	}

	i, it2 := it.PopSpec()
	if i != "trail" {
		t.Fatalf("trail entry incorrect. got=%s, want=trail", i)
	}

	if len(it2) != 2 {
		t.Fatal("specs not popped off trail correctly")
	}
}

func TestIntersection(t *testing.T) {
	t1 := Intersection([]string{"a", "b", "c"}, []string{"b"}, true)
	if t1[0] != "a" || t1[1] != "c" {
		t.Fatalf("intersection delivered an incorrect result got=%s", t1)
	}

	t2 := Intersection([]string{"a", "b"}, []string{"b", "c"}, true)
	if t2[0] != "a" || t2[1] != "c" {
		t.Fatalf("intersection delivered an incorrect result got=%s", t2)
	}

	t3 := Intersection([]string{"b"}, []string{"a", "b", "c"}, true)
	if t3[0] != "a" || t3[1] != "c" {
		t.Fatalf("intersection delivered an incorrect result got=%s", t3)
	}
}

func TestFromEnd(t *testing.T) {
	test := "helloworld"

	if FromEnd(test, 2) != "ld" {
		t.Fatal("FromEnd produces wrong substring")
	}
}
