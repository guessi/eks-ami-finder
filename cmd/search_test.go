package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func TestEc2EndpointHostname(t *testing.T) {
	tests := map[string]string{
		"us-east-1":      "ec2.us-east-1.amazonaws.com",
		"cn-north-1":     "ec2.cn-north-1.amazonaws.com.cn",
		"us-iso-east-1":  "ec2.us-iso-east-1.c2s.ic.gov",
		"us-isob-east-1": "ec2.us-isob-east-1.sc2s.sgov.gov",
	}
	for region, want := range tests {
		if got := ec2EndpointHostname(region); got != want {
			t.Errorf("region %q: got %q, want %q", region, got, want)
		}
	}
}

func TestToWindowsReleaseDate(t *testing.T) {
	tests := map[string]string{
		"2026":     "2026",
		"202605":   "2026.05",
		"20260527": "2026.05.27",
	}
	for in, want := range tests {
		if got := toWindowsReleaseDate(in); got != want {
			t.Errorf("input %q: got %q, want %q", in, got, want)
		}
	}
}

// fakeDescribeImagesClient serves pre-defined pages one by one
type fakeDescribeImagesClient struct {
	pages [][]types.Image
	calls int
}

func (f *fakeDescribeImagesClient) DescribeImages(ctx context.Context, in *ec2.DescribeImagesInput, _ ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
	out := &ec2.DescribeImagesOutput{Images: f.pages[f.calls]}
	f.calls++
	if f.calls < len(f.pages) {
		out.NextToken = aws.String(fmt.Sprintf("token-%d", f.calls))
	}
	return out, nil
}

func mkImage(id, creationDate string) types.Image {
	return types.Image{ImageId: aws.String(id), CreationDate: aws.String(creationDate)}
}

func TestFindAmiMatches(t *testing.T) {
	// two pages, unsorted, with a creation date tie across pages
	svc := &fakeDescribeImagesClient{
		pages: [][]types.Image{
			{
				mkImage("ami-b", "2026-05-20T00:00:00.000Z"),
				mkImage("ami-d", "2026-07-24T00:00:00.000Z"),
				mkImage("ami-e", "2026-03-01T00:00:00.000Z"),
			},
			{
				mkImage("ami-c", "2026-05-20T00:00:00.000Z"), // tie with ami-b
				mkImage("ami-a", "2026-06-30T00:00:00.000Z"),
			},
		},
	}

	images, err := findAmiMatches(context.Background(), svc, &ec2.DescribeImagesInput{}, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// expect newest first, tie broken by ImageId ascending
	want := []string{"ami-d", "ami-a", "ami-b"}
	if len(images) != len(want) {
		t.Fatalf("got %d images, want %d", len(images), len(want))
	}
	for i, id := range want {
		if got := aws.ToString(images[i].ImageId); got != id {
			t.Errorf("position %d: got %q, want %q", i, got, id)
		}
	}

	if svc.calls != 2 {
		t.Errorf("expected all pages fetched, got %d calls", svc.calls)
	}
}
