package leader

import (
	"context"
	"fmt"
	"github.com/go-liquor/liquor/v2/config"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"os"
	"path"
	"strings"
	"time"
)

func getNamespace() string {
	namespaceBytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		return "default"
	}
	namespace := strings.TrimSpace(string(namespaceBytes))
	return namespace
}

func NewLeaderElection(cfg *config.Config, logger *zap.Logger, run func(ctx context.Context)) {
	logger = logger.Named("leader-election")
	var k8sConfig *rest.Config
	var err error
	k8sConfig, err = rest.InClusterConfig()
	if err != nil {
		homeDir, _ := os.UserHomeDir()
		kubeconfig := path.Join(homeDir, ".kube", "config")
		k8sConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	clientset, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		logger.Error("Error building clientset", zap.Error(err))
		return
	}

	appName := cfg.GetString(config.AppName)
	if appName == "" {
		logger.Error("Error getting AppName (app.name)")
		return
	}
	id, _ := os.Hostname()
	id = fmt.Sprintf("%s-%s", id, time.Now().Format("20060102150405"))
	logger.Info("Running leader election", zap.String("id", id),
		zap.String("appName", appName),
		zap.String("namespace", getNamespace()))
	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("leader-election-%s", cfg.GetString(config.AppName)),
			Namespace: getNamespace(),
		},
		Client: clientset.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: id,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:          lock,
		LeaseDuration: 15 * time.Second,
		RenewDeadline: 10 * time.Second,
		RetryPeriod:   2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				logger.Debug("leader election lost", zap.String("id", id))
			},
		},
	})

}
