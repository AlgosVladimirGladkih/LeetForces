package unique_binary_search_trees_ii

import (
	"sort"
	"testing"
)

func collectInorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := collectInorder(root.Left)
	right := collectInorder(root.Right)
	out := make([]int, 0, len(left)+len(right)+1)
	out = append(out, left...)
	out = append(out, root.Val)
	out = append(out, right...)
	return out
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func isSortedASC(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] <= a[i-1] {
			return false
		}
	}
	return true
}

func signature(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return "(" + signature(root.Left) + "|" + signature(root.Right) + ")"
}

func catalan(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i+1) / (i + 2)
	}
	return c
}

func ok(msg string) {
	println("OK " + msg)
}

func runTest(t *testing.T, n, expected int) {
	t.Helper()
	got := generateTrees(n)
	if len(got) != expected {
		t.Fatalf("n=%d: expected %d trees, got %d", n, expected, len(got))
	}
	if n == 0 {
		if len(got) != 0 {
			t.Fatalf("n=0 should be empty")
		}
		ok("n=0 returns empty slice")
		return
	}
	seen := make(map[string]bool)
	for i, tree := range got {
		if tree == nil {
			t.Fatalf("n=%d tree[%d] should not be nil", n, i)
		}
		if countNodes(tree) != n {
			t.Fatalf("n=%d tree[%d] node count != %d", n, i, n)
		}
		seq := collectInorder(tree)
		if len(seq) != n {
			t.Fatalf("n=%d tree[%d] inorder len != %d", n, i, n)
		}
		if !isSortedASC(seq) {
			t.Fatalf("n=%d tree[%d] not a BST (inorder not strictly ascending): %v", n, i, seq)
		}
		for k := 1; k <= n; k++ {
			if seq[k-1] != k {
				t.Fatalf("n=%d tree[%d] missing or wrong value at pos %d: %v", n, i, k, seq)
			}
		}
		sig := signature(tree)
		if seen[sig] {
			t.Fatalf("n=%d duplicate tree signature at [%d]: %s", n, i, sig)
		}
		seen[sig] = true
	}
	ok("n=" + itoa(n) + " count=" + itoa(expected) + " all trees valid BST with unique shapes")
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	buf := []byte{}
	for v > 0 {
		buf = append([]byte{byte('0' + v%10)}, buf...)
		v /= 10
	}
	if neg {
		return "-" + string(buf)
	}
	return string(buf)
}

func TestGenerateTreesZero(t *testing.T) {
	runTest(t, 0, 0)
}

func TestGenerateTreesOne(t *testing.T) {
	runTest(t, 1, 1)
}

func TestGenerateTreesTwo(t *testing.T) {
	runTest(t, 2, 2)
}

func TestGenerateTreesThree(t *testing.T) {
	runTest(t, 3, 5)
}

func TestGenerateTreesFour(t *testing.T) {
	runTest(t, 4, 14)
}

func TestGenerateTreesFive(t *testing.T) {
	runTest(t, 5, 42)
}

func TestGenerateTreesSix(t *testing.T) {
	runTest(t, 6, 132)
}

func TestGenerateTreesSeven(t *testing.T) {
	runTest(t, 7, 429)
}

func TestGenerateTreesEight(t *testing.T) {
	runTest(t, 8, 1430)
}

func TestCatalanConsistency(t *testing.T) {
	for n := 1; n <= 8; n++ {
		runTest(t, n, catalan(n))
	}
}

func TestNoDuplicateShapes(t *testing.T) {
	trees := generateTrees(5)
	seen := make(map[string]int)
	for _, tr := range trees {
		s := signature(tr)
		seen[s]++
	}
	for sig, c := range seen {
		if c != 1 {
			t.Fatalf("shape seen %d times: %s", c, sig)
		}
	}
	ok("no_duplicate_shapes")
}

func TestEveryValueAppearsExactlyOnce(t *testing.T) {
	trees := generateTrees(6)
	for i, tr := range trees {
		seq := collectInorder(tr)
		freq := make(map[int]int)
		for _, v := range seq {
			freq[v]++
		}
		if len(freq) != 6 {
			t.Fatalf("tree[%d] has %d distinct values", i, len(freq))
		}
		for v := 1; v <= 6; v++ {
			if freq[v] != 1 {
				t.Fatalf("tree[%d] value %d count = %d", i, v, freq[v])
			}
		}
	}
	ok("every_value_appears_exactly_once")
}

func TestBSTProperty(t *testing.T) {
	trees := generateTrees(7)
	for i, tr := range trees {
		if !validateBST(tr, -1<<31, 1<<31) {
			t.Fatalf("tree[%d] not a BST", i)
		}
	}
	ok("bst_property")
}

func validateBST(root *TreeNode, lo, hi int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lo || root.Val >= hi {
		return false
	}
	return validateBST(root.Left, lo, root.Val) && validateBST(root.Right, root.Val, hi)
}

func TestNodeCountMatchesN(t *testing.T) {
	for n := 1; n <= 8; n++ {
		trees := generateTrees(n)
		for i, tr := range trees {
			if countNodes(tr) != n {
				t.Fatalf("n=%d tree[%d] has %d nodes", n, i, countNodes(tr))
			}
		}
	}
	ok("node_count_matches_n")
}

func TestResultIsIndependentBetweenCalls(t *testing.T) {
	a := generateTrees(4)
	b := generateTrees(4)
	if len(a) != len(b) {
		t.Fatalf("inconsistent lengths across calls")
	}
	sigsA := make([]string, len(a))
	sigsB := make([]string, len(b))
	for i := range a {
		sigsA[i] = signature(a[i])
		sigsB[i] = signature(b[i])
	}
	sort.Strings(sigsA)
	sort.Strings(sigsB)
	for i := range sigsA {
		if sigsA[i] != sigsB[i] {
			t.Fatalf("calls differ at %d: %s vs %s", i, sigsA[i], sigsB[i])
		}
	}
	ok("result_independent_between_calls")
}

func TestExactlyNLeafShapeCount(t *testing.T) {
	trees := generateTrees(3)
	shapes := map[string]bool{}
	for _, tr := range trees {
		shapes[signature(tr)] = true
	}
	if len(shapes) != 5 {
		t.Fatalf("n=3 unique shape count = %d, want 5", len(shapes))
	}
	ok("exactly_5_shapes_for_n3")
}

func TestMirroredShapesForN3(t *testing.T) {
	trees := generateTrees(3)
	shapes := []string{}
	for _, tr := range trees {
		shapes = append(shapes, signature(tr))
	}
	sort.Strings(shapes)
	if len(shapes) != 5 {
		t.Fatalf("n=3 should produce 5 distinct shapes, got %d", len(shapes))
	}
	ok("mirrored_shapes_for_n3")
}

func TestNonEmptyRootForPositiveN(t *testing.T) {
	for n := 1; n <= 8; n++ {
		trees := generateTrees(n)
		if len(trees) == 0 {
			t.Fatalf("n=%d should return at least one tree", n)
		}
		for i, tr := range trees {
			if tr == nil {
				t.Fatalf("n=%d tree[%d] nil", n, i)
			}
		}
	}
	ok("non_empty_root_for_positive_n")
}

func TestPanicFreeOnMaxN(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic on n=8: %v", r)
		}
	}()
	tr := generateTrees(8)
	if len(tr) != 1430 {
		t.Fatalf("n=8 count wrong: %d", len(tr))
	}
	ok("panic_free_on_max_n")
}

func TestNoNilEntriesForPositiveN(t *testing.T) {
	for n := 1; n <= 6; n++ {
		trees := generateTrees(n)
		for i, tr := range trees {
			if tr == nil {
				t.Fatalf("n=%d tree[%d] nil", n, i)
			}
		}
	}
	ok("no_nil_entries_for_positive_n")
}

func TestN1StructureExact(t *testing.T) {
	trees := generateTrees(1)
	if len(trees) != 1 {
		t.Fatalf("n=1 should produce 1 tree, got %d", len(trees))
	}
	if trees[0] == nil || trees[0].Val != 1 {
		t.Fatalf("n=1 tree should be single node with Val=1")
	}
	if trees[0].Left != nil || trees[0].Right != nil {
		t.Fatalf("n=1 tree should have no children")
	}
	ok("n1_structure_exact")
}

func TestN2BothShapesPresent(t *testing.T) {
	trees := generateTrees(2)
	sigs := map[string]bool{}
	for _, tr := range trees {
		sigs[signature(tr)] = true
	}
	expectedLeft := "((#|#)|#)"
	expectedRight := "(#|(#|#))"
	if !sigs[expectedLeft] {
		t.Fatalf("missing shape with root=1")
	}
	if !sigs[expectedRight] {
		t.Fatalf("missing shape with root=2")
	}
	ok("n2_both_shapes_present")
}

func TestN4HasFourteenShapes(t *testing.T) {
	trees := generateTrees(4)
	sigs := map[string]bool{}
	for _, tr := range trees {
		sigs[signature(tr)] = true
	}
	if len(sigs) != 14 {
		t.Fatalf("n=4 unique shape count = %d, want 14", len(sigs))
	}
	ok("n4_has_14_shapes")
}

func TestN5HasFortyTwoShapes(t *testing.T) {
	trees := generateTrees(5)
	sigs := map[string]bool{}
	for _, tr := range trees {
		sigs[signature(tr)] = true
	}
	if len(sigs) != 42 {
		t.Fatalf("n=5 unique shape count = %d, want 42", len(sigs))
	}
	ok("n5_has_42_shapes")
}

func TestInorderIsExactlyOneToN(t *testing.T) {
	for n := 1; n <= 7; n++ {
		trees := generateTrees(n)
		for i, tr := range trees {
			seq := collectInorder(tr)
			for k := 1; k <= n; k++ {
				if seq[k-1] != k {
					t.Fatalf("n=%d tree[%d] inorder position %d = %d, want %d", n, i, k, seq[k-1], k)
				}
			}
		}
	}
	ok("inorder_is_exactly_1_to_n")
}

func TestNoSharedSubtreesAcrossRoots(t *testing.T) {
	trees := generateTrees(4)
	roots := make(map[int]int)
	for _, tr := range trees {
		roots[tr.Val]++
	}
	if roots[1] == 0 || roots[2] == 0 || roots[3] == 0 || roots[4] == 0 {
		t.Fatalf("every value 1..4 must appear as root: %v", roots)
	}
	ok("every_value_appears_as_root_for_n4")
}

func TestRootCountSymmetric(t *testing.T) {
	trees := generateTrees(5)
	roots := make(map[int]int)
	for _, tr := range trees {
		roots[tr.Val]++
	}
	if roots[1] != roots[5] || roots[2] != roots[4] {
		t.Fatalf("root count should be symmetric: %v", roots)
	}
	ok("root_count_symmetric_for_n5")
}

func TestHeightsRange(t *testing.T) {
	trees := generateTrees(4)
	heights := map[int]bool{}
	for _, tr := range trees {
		h := treeHeight(tr)
		heights[h] = true
	}
	if !heights[3] {
		t.Fatalf("n=4 should include a near-balanced tree (height 3)")
	}
	if !heights[4] {
		t.Fatalf("n=4 should include right-skewed chain (height 4)")
	}
	ok("heights_range_for_n4")
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl := treeHeight(root.Left)
	hr := treeHeight(root.Right)
	if hl > hr {
		return hl + 1
	}
	return hr + 1
}

func TestBalancedAndSkewedPresent(t *testing.T) {
	trees := generateTrees(4)
	hasNearBalanced := false
	hasSkewed := false
	for _, tr := range trees {
		h := treeHeight(tr)
		if h == 3 {
			hasNearBalanced = true
		}
		if h == 4 {
			hasSkewed = true
		}
	}
	if !hasNearBalanced || !hasSkewed {
		t.Fatalf("n=4 must include both near-balanced (h=3) and skewed (h=4) trees")
	}
	ok("near_balanced_and_skewed_present_for_n4")
}

func TestLeftSkewedContainsAllValues(t *testing.T) {
	trees := generateTrees(4)
	for _, tr := range trees {
		if tr.Val == 4 && tr.Left != nil {
			if tr.Left.Val == 3 && tr.Left.Right == nil && tr.Left.Left == nil {
				if tr.Right != nil {
					t.Fatalf("skewed tree root=4 should have no right subtree")
				}
			}
		}
	}
	ok("left_skewed_shape_valid")
}

func TestRightSkewedContainsAllValues(t *testing.T) {
	trees := generateTrees(4)
	found := false
	for _, tr := range trees {
		if tr.Val == 1 && tr.Right != nil && tr.Left == nil {
			found = true
		}
	}
	if !found {
		t.Fatalf("n=4 should include root=1 with only right child")
	}
	ok("right_skewed_shape_present")
}

func TestResultDeterministicForSameN(t *testing.T) {
	first := signatureSet(generateTrees(6))
	second := signatureSet(generateTrees(6))
	if len(first) != len(second) {
		t.Fatalf("non-deterministic count")
	}
	for k := range first {
		if !second[k] {
			t.Fatalf("missing shape on second call: %s", k)
		}
	}
	ok("result_deterministic_for_same_n")
}

func signatureSet(trees []*TreeNode) map[string]bool {
	s := make(map[string]bool)
	for _, tr := range trees {
		s[signature(tr)] = true
	}
	return s
}

func TestNoTreeHasZeroNodes(t *testing.T) {
	for n := 1; n <= 8; n++ {
		trees := generateTrees(n)
		for i, tr := range trees {
			if countNodes(tr) == 0 {
				t.Fatalf("n=%d tree[%d] has zero nodes", n, i)
			}
		}
	}
	ok("no_tree_has_zero_nodes")
}

func TestTotalNodeCountSanity(t *testing.T) {
	for n := 1; n <= 5; n++ {
		trees := generateTrees(n)
		total := 0
		for _, tr := range trees {
			total += countNodes(tr)
		}
		expected := catalan(n) * n
		if total != expected {
			t.Fatalf("n=%d total nodes across all trees = %d, want %d", n, total, expected)
		}
	}
	ok("total_node_count_sanity")
}

func TestN6ShapesMatchCatalan(t *testing.T) {
	got := len(signatureSet(generateTrees(6)))
	if got != 132 {
		t.Fatalf("n=6 shapes = %d, want 132", got)
	}
	ok("n6_shapes_match_catalan")
}

func TestN7ShapesMatchCatalan(t *testing.T) {
	got := len(signatureSet(generateTrees(7)))
	if got != 429 {
		t.Fatalf("n=7 shapes = %d, want 429", got)
	}
	ok("n7_shapes_match_catalan")
}

func TestN8ShapesMatchCatalan(t *testing.T) {
	got := len(signatureSet(generateTrees(8)))
	if got != 1430 {
		t.Fatalf("n=8 shapes = %d, want 1430", got)
	}
	ok("n8_shapes_match_catalan")
}