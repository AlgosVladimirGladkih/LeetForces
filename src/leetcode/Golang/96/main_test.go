package mainunique_binary_search_trees

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want int
	}{
		{"basic_n0_empty", 0, 1},
		{"basic_n1_single", 1, 1},
		{"basic_n2_two", 2, 2},
		{"basic_n3_three", 3, 5},
		{"basic_n4_four", 4, 14},
		{"basic_n5_five", 5, 42},
		{"basic_n6_six", 6, 132},
		{"basic_n7_seven", 7, 429},
		{"basic_n8_eight", 8, 1430},
		{"basic_n9_nine", 9, 4862},
		{"basic_n10_ten", 10, 16796},
		{"basic_n11_eleven", 11, 58786},
		{"basic_n12_twelve", 12, 208012},
		{"basic_n13_thirteen", 13, 742900},
		{"basic_n14_fourteen", 14, 2674440},
		{"basic_n15_fifteen", 15, 9694845},
		{"basic_n16_sixteen", 16, 35357670},
		{"basic_n17_seventeen", 17, 129644790},
		{"basic_n18_eighteen", 18, 477638700},
		{"basic_n19_nineteen_max", 19, 1767263190},
		{"verify_n0_lower", 0, 1},
		{"verify_n1_min", 1, 1},
		{"verify_n2_smallest_nonzero", 2, 2},
		{"verify_n3_catalan5", 3, 5},
		{"verify_n4_catalan14", 4, 14},
		{"verify_n5_catalan42", 5, 42},
		{"verify_n6_catalan132", 6, 132},
		{"verify_n7_catalan429", 7, 429},
		{"verify_n8_catalan1430", 8, 1430},
		{"verify_n9_catalan4862", 9, 4862},
		{"verify_n10_catalan16796", 10, 16796},
		{"verify_n11_catalan58786", 11, 58786},
		{"verify_n12_catalan208012", 12, 208012},
		{"verify_n13_catalan742900", 13, 742900},
		{"verify_n14_catalan2674440", 14, 2674440},
		{"verify_n15_catalan9694845", 15, 9694845},
		{"verify_n16_catalan35357670", 16, 35357670},
		{"verify_n17_catalan129644790", 17, 129644790},
		{"verify_n18_catalan477638700", 18, 477638700},
		{"verify_n19_catalan1767263190", 19, 1767263190},
		{"stress_n0", 0, 1},
		{"stress_n1", 1, 1},
		{"stress_n2", 2, 2},
		{"stress_n3", 3, 5},
		{"stress_n4", 4, 14},
		{"stress_n5", 5, 42},
		{"stress_n6", 6, 132},
		{"stress_n7", 7, 429},
		{"stress_n8", 8, 1430},
		{"stress_n9", 9, 4862},
		{"stress_n10", 10, 16796},
		{"stress_n11", 11, 58786},
		{"stress_n12", 12, 208012},
		{"stress_n13", 13, 742900},
		{"stress_n14", 14, 2674440},
		{"stress_n15", 15, 9694845},
		{"stress_n16", 16, 35357670},
		{"stress_n17", 17, 129644790},
		{"stress_n18", 18, 477638700},
		{"stress_n19", 19, 1767263190},
		{"check_n0", 0, 1},
		{"check_n1", 1, 1},
		{"check_n2", 2, 2},
		{"check_n3", 3, 5},
		{"check_n4", 4, 14},
		{"check_n5", 5, 42},
		{"check_n6", 6, 132},
		{"check_n7", 7, 429},
		{"check_n8", 8, 1430},
		{"check_n9", 9, 4862},
		{"check_n10", 10, 16796},
		{"check_n11", 11, 58786},
		{"check_n12", 12, 208012},
		{"check_n13", 13, 742900},
		{"check_n14", 14, 2674440},
		{"check_n15", 15, 9694845},
		{"check_n16", 16, 35357670},
		{"check_n17", 17, 129644790},
		{"check_n18", 18, 477638700},
		{"check_n19", 19, 1767263190},
		{"final_n0", 0, 1},
		{"final_n1", 1, 1},
		{"final_n2", 2, 2},
		{"final_n3", 3, 5},
		{"final_n4", 4, 14},
		{"final_n5", 5, 42},
		{"final_n6", 6, 132},
		{"final_n7", 7, 429},
		{"final_n8", 8, 1430},
		{"final_n9", 9, 4862},
		{"final_n10", 10, 16796},
		{"final_n11", 11, 58786},
		{"final_n12", 12, 208012},
		{"final_n13", 13, 742900},
		{"final_n14", 14, 2674440},
		{"final_n15", 15, 9694845},
		{"final_n16", 16, 35357670},
		{"final_n17", 17, 129644790},
		{"final_n18", 18, 477638700},
		{"final_n19", 19, 1767263190},
	}
	passed := 0
	for _, c := range cases {
		got := numTrees(c.n)
		if got != c.want {
			t.Errorf("FAIL %s: numTrees(%d)=%d want %d", c.name, c.n, got, c.want)
			continue
		}
		passed++
		fmt.Printf("OK %s n=%d -> %d\n", c.name, c.n, c.want)
	}
	fmt.Printf("Total: %d/%d passed\n", passed, len(cases))
}
