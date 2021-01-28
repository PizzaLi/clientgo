package main

import (
	"fmt"
	"k8s/clientgo/k8s"
	"k8s/clientgo/util"
	"strings"
)

func main() {
	/*
		args := len(os.Args)
		var inOrOut string
		if args <= 1 {
			fmt.Println("Usage: need argument in or out")
			return
		} else if len(os.Args) > 2 {
			fmt.Println("Two arguments is enough")
			return
		} else {
			inOrOut = os.Args[1]
		}

		if inOrOut == "in" {
			incluster.GetK8sInfoFromIncluster()
		} else if inOrOut == "out" {
			outcluster.GetK8sInfoFromOutcluster()
		} else {
			fmt.Println("Usage: need argument in or out")
			return
		}
	*/
	/*
		var err error

		err = incluster.GetK8sInfoFromIncluster()
		if err != nil {
			fmt.Println(err.Error())
		}

		err = outcluster.GetK8sInfoFromOutcluster()
		if err != nil {
			fmt.Println(err.Error())
		}
	*/

	clientset, err := k8s.NewKubeClient("")
	if clientset == nil || err != nil {
		fmt.Printf("error clientset[%v], err[%s]\n", clientset, err.Error())
		return
	}
	podName, err := clientset.GetPodWithPattern("", "kubefate")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/*
		cmd := "kubectl get pods -n kube-fate | grep kubefate|grep Running| awk '{print $1}'"
		result, _ := util.ExecCommand(cmd)
		if len(result) == 0 {
			return
		}
	*/
	// fmt.Println(podName)
	// fmt.Println(strings.TrimSpace(result))
	/*
		if strings.TrimSpace(podName) == strings.TrimSpace(result) {
			fmt.Printf("podname[%s] equals to result[%s]\n", podName, strings.TrimSpace(result))
		}
		fmt.Println(podName)
		fmt.Println(strings.TrimSpace(result))
	*/
	/*
		cmd = fmt.Sprintf("kubectl logs -n kube-fate --tail 500 %s > ./testLog/kubefate.log", podName)
		result, _ = util.ExecCommand(cmd)
		fmt.Println("result:", result)
	*/

	err = clientset.WriteLogsIntoFile("kube-fate", podName, "./testLog/kubefate1.log", 500)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nodes, err := clientset.GetNodes()
	// nodes, err := clientset.GetNodesWithoutMaster()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/*
		cmd := "cnt=0;for i in `kubectl get node -o wide | grep -v NAME | awk -va=$cnt '{print $1\"tempfm-node-\"a\"=\"$6}'`;do ret=`echo $i | sed 's/temp/ /g'`;cnt=`expr $cnt + 1`;kubectl label node $ret --overwrite; done"
		util.ExecCommand(cmd)
	*/
	labels := clientset.GenerateFMNodeLabel(nodes, "fm-node-")
	err = clientset.SetLabelsForNode(nodes, labels)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	/*
		cmd := "iplist=\"\";for i in `kubectl get nodes --show-labels | grep -v NAME |awk '{split($6,a,\",\");{for(j=0;j<length(a);j++){if(length(a[j])>9 && substr(a[j],0,8)==\"fm-node-\"){split(a[j],b,\"=\");{print b[1]\":\"b[2]}}}}}'`;do iplist=$i,$iplist;done;echo ${iplist%,*}"
		result, _ := util.ExecCommand(cmd)
		fmt.Println("result:", result)
		fmt.Println("labelFromClient:", clientset.GetNodeLabelOfFM(nodes, "fm-node-"))
	*/
	cmd := fmt.Sprintf("kubectl get pods -n fate-%d |grep %s* | grep Running |wc -l", 10002, "python")
	result, _ := util.ExecCommand(cmd)
	fmt.Println("result:", strings.TrimSpace(result))

	podNameList, _ := clientset.GetPodListWithPattern("fate-10002", "python")
	if len(podNameList) > 0 {
		fmt.Println("result from client: ", len(podNameList))
	}

	/*
		cmd = fmt.Sprintf("kubectl get namespace |awk '{if($1==\"%s\"){print $0}}' |grep Active|wc -l", "fate-10002")
		value, _ := util.ExecCommand(cmd)
		fmt.Println("value:", value)

		nsList, _ := clientset.ListNamespaceWithPattern("fate-10002")
		if len(nsList) > 0 {
			fmt.Println("nsList: ", len(nsList))
		}
	*/
	cmd = fmt.Sprintf("kubectl create namespace %s", "test")
	value, _ := util.ExecCommand(cmd)
	fmt.Println("value:", value)

	namespace, _ := clientset.CreateNamespace("test1")
	fmt.Println("namespace: ", namespace)
}
