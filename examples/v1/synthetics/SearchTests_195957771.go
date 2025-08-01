// Search Synthetic tests with boolean query parameters

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.SearchTests(ctx, *datadogV1.NewSearchTestsOptionalParameters().WithText("tag:value").WithIncludeFullConfig(true).WithSearchSuites(true).WithFacetsOnly(true).WithStart(10).WithCount(5).WithSort("name,desc"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.SearchTests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.SearchTests`:\n%s\n", responseContent)
}
