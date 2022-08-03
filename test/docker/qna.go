//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const QnATransformers = "qna-transformers"

func startQnATransformers(ctx context.Context, networkName, qnaImage string) (*DockerContainer, error) {
	image := "semitechnologies/qna-transformers:distilbert-base-uncased-distilled-squad"
	if len(qnaImage) > 0 {
		image = qnaImage
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:    image,
			Hostname: QnATransformers,
			Networks: []string{networkName},
			NetworkAliases: map[string][]string{
				networkName: {QnATransformers},
			},
			ExposedPorts: []string{"8080/tcp"},
			AutoRemove:   true,
			WaitingFor: wait.
				ForHTTP("/.well-known/ready").
				WithPort(nat.Port("8080")).
				WithStatusCodeMatcher(func(status int) bool {
					return status == 204
				}).
				WithStartupTimeout(240 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		return nil, err
	}
	uri, err := container.Endpoint(ctx, "")
	if err != nil {
		return nil, err
	}
	envSettings := make(map[string]string)
	envSettings["QNA_INFERENCE_API"] = fmt.Sprintf("http://%s:%s", QnATransformers, "8080")
	return &DockerContainer{QnATransformers, uri, container, envSettings}, nil
}
