package rabbitmq

import (
	"context"
	"fmt"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type rabbitContainer struct {
	testcontainers.Container
}

func startContainer(ctx context.Context) (*rabbitContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/rabbitmq:3-management",
		ExposedPorts: []string{"15672/tcp"},
		Env: map[string]string{
			"RABBITMQ_DEFAULT_USER": "user",
			"RABBITMQ_DEFAULT_PASS": "pass",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog(" completed with 4 plugins."),
			wait.ForListeningPort("15672/tcp"),
		),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ProviderType:     testcontainers.ProviderPodman,
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return &rabbitContainer{Container: container}, nil
}
func TestRabbitMq(t *testing.T) {
	ctx := context.Background()
	container, err := startContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})
	p, _ := container.MappedPort(ctx, "15672/tcp")
	fmt.Println(p)
	err = Discover(p.Port(), "user", "pass")
	if err != nil {
		t.Fatal(err)
	}
}
