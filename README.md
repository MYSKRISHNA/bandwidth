# bandwidth
**Communication Flow Stabilization Across Multi Node Enterprise Systems**

### Paper Information
- **Author(s):** SaiKrishna Mylavarapu
- **Published In:** International Journal of Innovative Research in Engineering & Multidisciplinary Physical Sciences (IJIRMPS)
- **Publication Date:** August, 2025
- **ISSN:** E-ISSN: 2349-7300
- **DOI:** https://doi.org/10.37082/IJIRMPS.v13.i4.233139
- **Impact Factor:** 9.907

### Abstract
Efficient CPU resource management is essential for maintaining high performance in distributed systems as workloads and cluster sizes increase. Static workload allocation results in uneven CPU utilization, while conventional load balancing techniques cannot effectively adapt to dynamic workload variations. The study analyzes CPU usage across clusters with 3, 5, 7, 9, and 11 nodes to evaluate the limitations of existing allocation strategies. It demonstrates how workload imbalance and coordination overhead reduce resource utilization, system efficiency, and scalability. The findings highlight the need for adaptive workload allocation mechanisms based on real-time CPU usage to achieve balanced resource utilization and improved overall performance.

### Core Technical Contributions

- **Adaptive CPU Resource Allocation Framework:** Designed a dynamic workload allocation framework that distributes tasks based on real time CPU utilization, improving resource efficiency and balanced node utilization.

- **Intelligent Load Balancing Mechanism:** Implemented a load balancing strategy that selects the least loaded active server, reducing workload imbalance and improving CPU utilization across distributed nodes.

- **Adaptive Auto Scaling Strategy:** Developed an auto scaling mechanism that provisions additional worker nodes when CPU load exceeds predefined thresholds, ensuring sustained performance under increasing workloads.

- **Distributed System Simulation Model:** Implemented a Go based distributed system simulator integrating load balancing, server health monitoring, and adaptive scaling to evaluate CPU resource management strategies.

- **Scalability Analysis Across Cluster Sizes:** Evaluated CPU utilization across clusters of 3, 5, 7, 9, and 11 nodes, demonstrating the scalability benefits of adaptive resource allocation over deterministic approaches. 

### Experimental Results (Summary)

  | Cluster Size | Deterministic Allocation (%) | Load Distributed Allocation (%) | Improvement (%) |
  |--------------|------------------------------| --------------------------------| ---------------|
  | 3            |  78                          | 82                              | 5.13           |
  | 5            |  72                          | 79                              | 9.72           |
  | 7            |  66                          | 75                              | 13.64          |
  | 9            |  61                          | 72                              | 18.03          |
  | 11           |  57                          | 70                              | 22.81          |

### Citation
Adaptive CPU Resource Management in Distributed Systems
* SaiKrishna Mylavarapu
* International Journal of Intelligent Systems and Applications in Engineering (IJISAE) 
* ISSN E-ISSN: 2147-6799
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.
* Resources \
https://www.ijisae.org/ 
* Author Contact \
  **LinkedIn**: linkedin.com/in/saikrishna-mylavarapu-35479114 | **Email**: krishnamysap@gmail.com






