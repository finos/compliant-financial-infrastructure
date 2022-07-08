#!/bin/bash
# Script to give finos-admin-users Cluster Admin privileges
oc adm policy add-cluster-role-to-user cluster-admin finos-admin-1
oc adm policy add-cluster-role-to-user cluster-admin finos-admin-2
oc adm policy add-cluster-role-to-user cluster-admin finos-admin-3