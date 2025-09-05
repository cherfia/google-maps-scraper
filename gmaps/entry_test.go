package gmaps_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"

	"github.com/gosom/google-maps-scraper/gmaps"
)

func createGoQueryFromFile(t *testing.T, path string) *goquery.Document {
	t.Helper()

	fd, err := os.Open(path)
	require.NoError(t, err)

	defer fd.Close()

	doc, err := goquery.NewDocumentFromReader(fd)
	require.NoError(t, err)

	return doc
}

func Test_EntryFromJSON(t *testing.T) {
	expected := gmaps.Entry{
		Link:       "https://www.google.com/maps/place/Kipriakon/data=!4m2!3m1!1s0x14e732fd76f0d90d:0xe5415928d6702b47!10m1!1e1",
		Title:      "Kipriakon",
		Category:   "Restaurant",
		Categories: []string{"Restaurant"},
		Address:    "Old port, Limassol 3042",
		OpenHours: map[string][]string{
			"Monday":    {"12:30–10 pm"},
			"Tuesday":   {"12:30–10 pm"},
			"Wednesday": {"12:30–10 pm"},
			"Thursday":  {"12:30–10 pm"},
			"Friday":    {"12:30–10 pm"},
			"Saturday":  {"12:30–10 pm"},
			"Sunday":    {"12:30–10 pm"},
		},
		WebSite:      "",
		Phone:        "25 101555",
		PlusCode:     "M2CR+6X Limassol",
		ReviewCount:  396,
		ReviewRating: 4.2,
		Latitude:     34.670595399999996,
		Longtitude:   33.042456699999995,
		Cid:          "16519582940102929223",
		Status:       "Closed ⋅ Opens 12:30\u202fpm Tue",
		ReviewsLink:  "https://search.google.com/local/reviews?placeid=ChIJDdnwdv0y5xQRRytw1ihZQeU&q=Kipriakon&authuser=0&hl=en&gl=CY",
		PlaceID:      "ChIJDdnwdv0y5xQRRytw1ihZQeU",
		Thumbnail:    "https://lh5.googleusercontent.com/p/AF1QipP4Y7A8nYL3KKXznSl69pXSq9p2IXCYUjVvOh0F=w408-h408-k-no",
		Timezone:     "Asia/Nicosia",
		PriceRange:   "€€",
		PriceCategory: gmaps.PriceCategoryModerate,
		DataID:       "0x14e732fd76f0d90d:0xe5415928d6702b47",
		Images: []gmaps.Image{
			{
				Title: "All",
				Image: "https://lh5.googleusercontent.com/p/AF1QipP4Y7A8nYL3KKXznSl69pXSq9p2IXCYUjVvOh0F=w298-h298-k-no",
			},
			{
				Title: "Latest",
				Image: "https://lh5.googleusercontent.com/p/AF1QipNgMqyaQs2MqH1oiGC44eDcvudurxQfNb2RuDsd=w224-h298-k-no",
			},
			{
				Title: "Videos",
				Image: "https://lh5.googleusercontent.com/p/AF1QipPZbq8v8K8RZfvL6gZ_4Dw6qwNJ_MUxxOOfBo7h=w224-h398-k-no",
			},
			{
				Title: "Menu",
				Image: "https://lh5.googleusercontent.com/p/AF1QipNhoFtPcaLCIhdN3GhlJ6sQIvdhaESnRG8nyeC8=w397-h298-k-no",
			},
			{
				Title: "Food & drink",
				Image: "https://lh5.googleusercontent.com/p/AF1QipMbu-iiWkE4DsXx3aI7nGaqyXJKbBYCrBXvzOnu=w298-h298-k-no",
			},
			{
				Title: "Vibe",
				Image: "https://lh5.googleusercontent.com/p/AF1QipOGg_vrD4bzkOre5Ly6CFXuO3YCOGfFxQ-EiEkW=w224-h398-k-no",
			},
			{
				Title: "Fried green tomatoes",
				Image: "https://lh5.googleusercontent.com/p/AF1QipOziHd2hqM1jnK9KfCGf1zVhcOrx8Bj7VdJXj0=w397-h298-k-no",
			},
			{
				Title: "French fries",
				Image: "https://lh5.googleusercontent.com/p/AF1QipNJyq7nAlKtsxxbNy4PHUZOhJ0k7HPP8tTAlwcV=w397-h298-k-no",
			},
			{
				Title: "By owner",
				Image: "https://lh5.googleusercontent.com/p/AF1QipNRE2R5k13zT-0WG4b6XOD_BES9-nMK04hlCMVV=w298-h298-k-no",
			},
			{
				Title: "Street View & 360°",
				Image: "https://lh5.googleusercontent.com/p/AF1QipMwkHP8GmDCSuwnWS7pYVQvtDWdsdk-CUwxtsXL=w224-h298-k-no-pi-23.425545-ya289.20517-ro-8.658787-fo100",
			},
		},
		OrderOnline: []gmaps.LinkSource{
			{
				Link:   "https://foody.com.cy/delivery/lemesos/to-kypriakon?utm_source=google&utm_medium=organic&utm_campaign=google_reserve_place_order_action",
				Source: "foody.com.cy",
			},
			{
				Link:   "https://wolt.com/en/cyp/limassol/restaurant/kypriakon?utm_source=googlemapreserved&utm_campaign=kypriakon",
				Source: "wolt.com",
			},
		},
		Owner: gmaps.Owner{
			ID:   "102769814432182832009",
			Name: "Kipriakon (Owner)",
			Link: "https://www.google.com/maps/contrib/102769814432182832009",
		},
		CompleteAddress: gmaps.Address{
			Borough:    "",
			Street:     "Old port",
			City:       "Limassol",
			PostalCode: "3042",
			State:      "",
			Country:    "CY",
		},
		ReviewsPerRating: map[int]int{
			1: 37,
			2: 16,
			3: 27,
			4: 60,
			5: 256,
		},
	}

	raw, err := os.ReadFile("../testdata/raw.json")
	require.NoError(t, err)
	require.NotEmpty(t, raw)

	entry, err := gmaps.EntryFromJSON(raw)
	require.NoError(t, err)

	require.Len(t, entry.About, 10)

	for _, about := range entry.About {
		require.NotEmpty(t, about.ID)
		require.NotEmpty(t, about.Name)
		require.NotEmpty(t, about.Options)
	}

	entry.About = nil

	require.Len(t, entry.PopularTimes, 7)

	for k, v := range entry.PopularTimes {
		require.Contains(t,
			[]string{
				"Monday",
				"Tuesday",
				"Wednesday",
				"Thursday",
				"Friday",
				"Saturday",
				"Sunday",
			}, k)

		for _, traffic := range v {
			require.GreaterOrEqual(t, traffic, 0)
			require.LessOrEqual(t, traffic, 100)
		}
	}

	monday := entry.PopularTimes["Monday"]
	require.Equal(t, 100, monday[20])

	entry.PopularTimes = nil
	entry.UserReviews = nil

	require.Equal(t, expected, entry)
}

func Test_EntryFromJSON2(t *testing.T) {
	fnames := []string{
		"../testdata/panic.json",
		"../testdata/panic2.json",
	}
	for _, fname := range fnames {
		raw, err := os.ReadFile(fname)
		require.NoError(t, err)
		require.NotEmpty(t, raw)

		_, err = gmaps.EntryFromJSON(raw)
		require.NoError(t, err)
	}
}

func Test_EntryFromJSONRaw2(t *testing.T) {
	raw, err := os.ReadFile("../testdata/raw2.json")

	require.NoError(t, err)
	require.NotEmpty(t, raw)

	entry, err := gmaps.EntryFromJSON(raw)

	require.NoError(t, err)
	require.Greater(t, len(entry.About), 0)
}

func Test_EntryFromJsonC(t *testing.T) {
	raw, err := os.ReadFile("../testdata/output.json")

	require.NoError(t, err)
	require.NotEmpty(t, raw)

	entries, err := gmaps.ParseSearchResults(raw)

	require.NoError(t, err)

	for _, entry := range entries {
		fmt.Printf("%+v\n", entry)
	}
}

func Test_parseRelativeTime(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"23 hours ago", ""}, // Will be empty because we can't predict the exact date
		{"a day ago", ""},    // Will be empty because we can't predict the exact date
		{"2 days ago", ""},   // Will be empty because we can't predict the exact date
		{"3 days ago", ""},   // Will be empty because we can't predict the exact date
		{"1 week ago", ""},   // Will be empty because we can't predict the exact date
		{"a month ago", ""},  // Will be empty because we can't predict the exact date
		{"1 year ago", ""},   // Will be empty because we can't predict the exact date
		{"Edited 2 months ago", ""}, // Will be empty because we can't predict the exact date
		{"Edited a day ago", ""},    // Will be empty because we can't predict the exact date
		{"Edited 1 week ago", ""},   // Will be empty because we can't predict the exact date
		{"invalid", "invalid"}, // Should return original string if can't parse
		{"", ""},             // Should return empty string for empty input
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// Since we can't predict the exact date, we just check that it's not empty
			// and follows the expected format (YYYY-M-D)
			result := gmaps.ParseRelativeTime(tt.input)
			
			if tt.input == "" {
				require.Equal(t, "", result)
			} else if tt.input == "invalid" {
				require.Equal(t, "invalid", result)
			} else {
				// For valid relative time strings, check that we get a full timestamp format
				require.NotEmpty(t, result)
				require.Regexp(t, `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{6}\+00$`, result)
			}
		})
	}
}

func TestCategorizePriceRange(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected gmaps.PriceCategory
	}{
		// Empty or missing input
		{"empty string", "", gmaps.PriceCategoryUnspecified},
		{"whitespace only", "   ", gmaps.PriceCategoryUnspecified},
		{"tab and spaces", "\t  \n", gmaps.PriceCategoryUnspecified},

		// Symbol-based pricing
		{"single dollar", "$", gmaps.PriceCategoryBudget},
		{"double dollar", "$$", gmaps.PriceCategoryModerate},
		{"triple dollar", "$$$", gmaps.PriceCategoryExpensive},
		{"quadruple dollar", "$$$$", gmaps.PriceCategoryLuxury},
		{"five dollars", "$$$$$", gmaps.PriceCategoryUnspecified},
		{"dollar with spaces", "  $$  ", gmaps.PriceCategoryModerate},

		// Numeric ranges
		{"range 10-20", "10-20", gmaps.PriceCategoryBudget},
		{"range 15-25", "15-25", gmaps.PriceCategoryBudget},
		{"range 20-40", "20-40", gmaps.PriceCategoryModerate},
		{"range 30-50", "30-50", gmaps.PriceCategoryModerate},
		{"range 40-100", "40-100", gmaps.PriceCategoryExpensive},
		{"range 60-80", "60-80", gmaps.PriceCategoryExpensive},
		{"range 100-200", "100-200", gmaps.PriceCategoryLuxury},
		{"range 150-300", "150-300", gmaps.PriceCategoryLuxury},

		// Range with different separators
		{"range with en dash", "10–20", gmaps.PriceCategoryBudget},
		{"range with em dash", "10—20", gmaps.PriceCategoryBudget},
		{"range with spaces", "10 - 20", gmaps.PriceCategoryBudget},

		// Plus format
		{"plus 100", "100+", gmaps.PriceCategoryLuxury},
		{"plus 50", "50+", gmaps.PriceCategoryExpensive},
		{"plus 25", "25+", gmaps.PriceCategoryModerate},
		{"plus 15", "15+", gmaps.PriceCategoryBudget},
		{"plus with spaces", " 100+ ", gmaps.PriceCategoryLuxury},

		// Single numbers
		{"single 10", "10", gmaps.PriceCategoryBudget},
		{"single 20", "20", gmaps.PriceCategoryBudget},
		{"single 30", "30", gmaps.PriceCategoryModerate},
		{"single 50", "50", gmaps.PriceCategoryExpensive},
		{"single 100", "100", gmaps.PriceCategoryLuxury},
		{"single 150", "150", gmaps.PriceCategoryLuxury},

		// Decimal numbers
		{"decimal range", "15.5-25.5", gmaps.PriceCategoryModerate},
		{"decimal single", "25.99", gmaps.PriceCategoryModerate},
		{"decimal plus", "99.99+", gmaps.PriceCategoryExpensive},

		// Edge cases
		{"exactly 20", "20", gmaps.PriceCategoryBudget},
		{"exactly 40", "40", gmaps.PriceCategoryModerate},
		{"exactly 100", "100", gmaps.PriceCategoryLuxury},

		// Currency symbols with numeric values
		{"dollar with number", "$25", gmaps.PriceCategoryModerate},
		{"dollar with decimal", "$15.50", gmaps.PriceCategoryBudget},
		{"dollar with range", "$10-20", gmaps.PriceCategoryBudget},
		{"euro with number", "€30", gmaps.PriceCategoryModerate},
		{"pound with number", "£50", gmaps.PriceCategoryExpensive},

		// Invalid inputs
		{"invalid text", "expensive", gmaps.PriceCategoryUnspecified},
		{"invalid range", "a-b", gmaps.PriceCategoryUnspecified},
		{"negative number", "-10", gmaps.PriceCategoryUnspecified},
		{"zero", "0", gmaps.PriceCategoryBudget},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gmaps.CategorizePriceRange(tt.input)
			require.Equal(t, tt.expected, result, "Input: %q", tt.input)
		})
	}
}

func TestCategorizePriceRange_EdgeCases(t *testing.T) {
	// Test boundary values
	boundaryTests := []struct {
		value    float64
		expected gmaps.PriceCategory
	}{
		{0, gmaps.PriceCategoryBudget},
		{20, gmaps.PriceCategoryBudget},
		{20.01, gmaps.PriceCategoryModerate},
		{40, gmaps.PriceCategoryModerate},
		{40.01, gmaps.PriceCategoryExpensive},
		{100, gmaps.PriceCategoryLuxury},
		{100.01, gmaps.PriceCategoryLuxury},
	}

	for _, tt := range boundaryTests {
		t.Run(fmt.Sprintf("boundary_%.2f", tt.value), func(t *testing.T) {
			result := gmaps.CategorizePriceRange(fmt.Sprintf("%.2f", tt.value))
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractPlaceIDFromReviewsLink(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid reviews link",
			input:    "https://search.google.com/local/reviews?placeid=ChIJgxwpY1_MHkcRFKqXgftRhPY&q=La+Cantina&authuser=0&hl=en&gl=PL",
			expected: "ChIJgxwpY1_MHkcRFKqXgftRhPY",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no placeid parameter",
			input:    "https://search.google.com/local/reviews?q=La+Cantina&authuser=0&hl=en&gl=PL",
			expected: "",
		},
		{
			name:     "placeid at end of URL",
			input:    "https://search.google.com/local/reviews?q=La+Cantina&authuser=0&hl=en&gl=PL&placeid=ChIJgxwpY1_MHkcRFKqXgftRhPY",
			expected: "ChIJgxwpY1_MHkcRFKqXgftRhPY",
		},
		{
			name:     "placeid in middle of URL",
			input:    "https://search.google.com/local/reviews?q=La+Cantina&placeid=ChIJgxwpY1_MHkcRFKqXgftRhPY&authuser=0&hl=en&gl=PL",
			expected: "ChIJgxwpY1_MHkcRFKqXgftRhPY",
		},
		{
			name:     "invalid URL format",
			input:    "not a URL",
			expected: "",
		},
		{
			name:     "placeid with special characters",
			input:    "https://search.google.com/local/reviews?placeid=ChIJgxwpY1_MHkcRFKqXgftRhPY%20encoded&q=La+Cantina",
			expected: "ChIJgxwpY1_MHkcRFKqXgftRhPY%20encoded",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gmaps.ExtractPlaceIDFromReviewsLink(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
