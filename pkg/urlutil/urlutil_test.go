package urlutil // import "moby/pkg/urlutil"

import "testing"

var (
	gitUrls = []string{
		"git://moby",
		"git@github.com:docker/docker.git",
		"git@bitbucket.org:atlassianlabs/atlassian-docker.git",
		"https://moby.git",
		"http://moby.git",
		"http://moby.git#branch",
		"http://moby.git#:dir",
	}
	incompleteGitUrls = []string{
		"moby",
	}
	invalidGitUrls = []string{
		"http://moby.git:#branch",
	}
	transportUrls = []string{
		"tcp://example.com",
		"tcp+tls://example.com",
		"udp://example.com",
		"unix:///example",
		"unixgram:///example",
	}
)

func TestIsGIT(t *testing.T) {
	for _, url := range gitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range incompleteGitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range invalidGitUrls {
		if IsGitURL(url) {
			t.Fatalf("%q should not be detected as valid Git prefix", url)
		}
	}
}

func TestIsTransport(t *testing.T) {
	for _, url := range transportUrls {
		if !IsTransportURL(url) {
			t.Fatalf("%q should be detected as valid Transport url", url)
		}
	}
}
