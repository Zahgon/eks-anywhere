package cilium

import (
	v1 "k8s.io/api/apps/v1"
)

func CheckDaemonSetReady(daemonSet *v1.DaemonSet) error { _ = "STUB: not implemented"; return nil }

func CheckPreflightDaemonSetReady(ciliumDaemonSet, preflightDaemonSet *v1.DaemonSet) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckDeploymentReady(deployment *v1.Deployment) error { _ = "STUB: not implemented"; return nil }

func checkDaemonSetObservedGeneration(daemonSet *v1.DaemonSet) error {
	_ = "STUB: not implemented"
	return nil
}

func checkDeploymentObservedGeneration(deployment *v1.Deployment) error {
	_ = "STUB: not implemented"
	return nil
}
