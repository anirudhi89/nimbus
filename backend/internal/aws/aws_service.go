package aws

// Handles data fetch and logic related to AWS (EC2, EKS)

// // Function to initialize AWS SDK with default config
// func initAWSConfig() (aws.Config, error) {
// 	// Load default configuration from AWS credentials and region
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
// 	if err != nil {
// 		return aws.Config{}, fmt.Errorf("unable to load SDK config, %v", err)
// 	}
// 	return cfg, nil
// }

// // Fetch EC2 Instances and their status
// func getEC2Instances(cfg aws.Config) ([]*ec2.Instance, error) {
// 	ec2Svc := ec2.NewFromConfig(cfg)

// 	// Describe EC2 instances
// 	resp, err := ec2Svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to describe EC2 instances, %v", err)
// 	}

// 	var instances []*ec2.Instance
// 	for _, reservation := range resp.Reservations {
// 		for _, instance := range reservation.Instances {
// 			instances = append(instances, instance)
// 		}
// 	}
// 	return instances, nil
// }

// // Fetch EKS Clusters and Node Groups
// func getEKSClusterInfo(cfg aws.Config) ([]*eks.Cluster, error) {
// 	eksSvc := eks.NewFromConfig(cfg)

// 	// List EKS clusters
// 	resp, err := eksSvc.ListClusters(context.TODO(), &eks.ListClustersInput{})
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to list EKS clusters, %v", err)
// 	}

// 	var clusters []*eks.Cluster
// 	for _, clusterName := range resp.Clusters {
// 		clusterResp, err := eksSvc.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{
// 			Name: &clusterName,
// 		})
// 		if err != nil {
// 			log.Printf("Failed to describe EKS cluster %s: %v", clusterName, err)
// 			continue
// 		}
// 		clusters = append(clusters, clusterResp.Cluster)
// 	}
// 	return clusters, nil
// }

// // Print EC2 Instances Info
// func printEC2Info(instances []*ec2.Instance) {
// 	fmt.Println("EC2 Instance Statuses:")
// 	for _, instance := range instances {
// 		fmt.Printf("- Instance ID: %s, State: %s, Type: %s\n", *instance.InstanceId, instance.State.Name, *instance.InstanceType)
// 	}
// }

// // Print EKS Cluster Info
// func printEKSInfo(clusters []*eks.Cluster) {
// 	fmt.Println("EKS Cluster Info:")
// 	for _, cluster := range clusters {
// 		fmt.Printf("- Cluster Name: %s, Status: %s, Version: %s\n", *cluster.Name, cluster.Status, *cluster.Version)
// 	}
// }

// func main() {
// 	// Initialize AWS SDK
// 	cfg, err := initAWSConfig()
// 	if err != nil {
// 		log.Fatalf("Error initializing AWS SDK: %v", err)
// 	}

// 	// Fetch EC2 instances
// 	instances, err := getEC2Instances(cfg)
// 	if err != nil {
// 		log.Fatalf("Error getting EC2 instances: %v", err)
// 	}

// 	// Fetch EKS clusters
// 	clusters, err := getEKSClusterInfo(cfg)
// 	if err != nil {
// 		log.Fatalf("Error getting EKS clusters: %v", err)
// 	}

// 	// Output the EC2 instance information
// 	printEC2Info(instances)

// 	// Output the EKS cluster information
// 	printEKSInfo(clusters)
// }
