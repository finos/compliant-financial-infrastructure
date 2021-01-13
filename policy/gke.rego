package main

has_field(obj, field) {
    obj[field]
}

warn[msg] {
    np := input.resource.google_container_cluster[gke].network_policy
    np.enabled == false
    msg = sprintf("GKE cluster `%v` network policy needs to be enabled", [gke])
}

deny[msg] {
    pcc := input.resource.google_container_cluster[gke]
    has_field(pcc, "private_cluster_config")
    pcc.private_cluster_config.enable_private_nodes == false
    msg = sprintf("GKE cluster `%v` private nodes are not enabled", [gke])
}