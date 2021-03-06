package image // import "moby/integration/image"

import (
	"context"
	"testing"

	"moby/api/types"
	"moby/integration/internal/container"
	"moby/internal/test/request"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestRemoveImageOrphaning(t *testing.T) {
	defer setupTest(t)()
	ctx := context.Background()
	client := request.NewAPIClient(t)

	img := "test-container-orphaning"

	// Create a container from busybox, and commit a small change so we have a new image
	cID1 := container.Create(t, ctx, client, container.WithCmd(""))
	commitResp1, err := client.ContainerCommit(ctx, cID1, types.ContainerCommitOptions{
		Changes:   []string{`ENTRYPOINT ["true"]`},
		Reference: img,
	})
	assert.NilError(t, err)

	// verifies that reference now points to first image
	resp, _, err := client.ImageInspectWithRaw(ctx, img)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(resp.ID, commitResp1.ID))

	// Create a container from created image, and commit a small change with same reference name
	cID2 := container.Create(t, ctx, client, container.WithImage(img), container.WithCmd(""))
	commitResp2, err := client.ContainerCommit(ctx, cID2, types.ContainerCommitOptions{
		Changes:   []string{`LABEL Maintainer="Integration Tests"`},
		Reference: img,
	})
	assert.NilError(t, err)

	// verifies that reference now points to second image
	resp, _, err = client.ImageInspectWithRaw(ctx, img)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(resp.ID, commitResp2.ID))

	// try to remove the image, should not error out.
	_, err = client.ImageRemove(ctx, img, types.ImageRemoveOptions{})
	assert.NilError(t, err)

	// check if the first image is still there
	resp, _, err = client.ImageInspectWithRaw(ctx, commitResp1.ID)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(resp.ID, commitResp1.ID))

	// check if the second image has been deleted
	_, _, err = client.ImageInspectWithRaw(ctx, commitResp2.ID)
	assert.Check(t, is.ErrorContains(err, "No such image:"))
}
