package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"istio.io/istio/istioctl/pkg/util/handlers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func getPodsNameInDefaultNamespace(toComplete string) ([]string, error) {
	kubeClient, err := kubeClient(kubeconfig, configContext)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	ns := handlers.HandleNamespace(namespace, defaultNamespace)
	podList, err := kubeClient.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var podsName []string
	for _, pod := range podList.Items {
		if toComplete == "" || strings.HasPrefix(pod.Name, toComplete) {
			podsName = append(podsName, pod.Name)
		}
	}

	return podsName, nil
}

func validPodsNameArgsFunction(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	podsName, err := getPodsNameInDefaultNamespace(toComplete)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return podsName, cobra.ShellCompDirectiveNoFileComp
}

func getNamespacesName(toComplete string) ([]string, error) {
	kubeClient, err := kubeClient(kubeconfig, configContext)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	nsList, err := getNamespaces(ctx, kubeClient)
	if err != nil {
		return nil, err
	}

	var nsNameList []string
	for _, ns := range nsList {
		if toComplete == "" || strings.HasPrefix(ns.Name, toComplete) {
			nsNameList = append(nsNameList, ns.Name)
		}
	}

	return nsNameList, nil
}

func validNamespaceArgsFunction(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	nsName, err := getNamespacesName(toComplete)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return nsName, cobra.ShellCompDirectiveNoFileComp
}
