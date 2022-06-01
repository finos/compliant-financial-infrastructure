#! bin/bash
oc annotate compliancescans/ocp4-finos-finos compliance.openshift.io/rescan=
oc annotate compliancescans/ocp4-finos-finos-node-master compliance.openshift.io/rescan=
oc annotate compliancescans/ocp4-finos-finos-node-worker compliance.openshift.io/rescan=
oc annotate compliancescans/rhcos4-finos-finos-master compliance.openshift.io/rescan=
oc annotate compliancescans/rhcos4-finos-finos-worker  compliance.openshift.io/rescan=