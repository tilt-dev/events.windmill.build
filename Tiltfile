local_resource(
  'build', './hack/build.sh')

docker_build(
  'events-windmill-build', '.')

k8s_yaml(
  helm('./chart',
       name='events-windmill-build',
       set=[
         'imageName=events-windmill-build',
         'numReplicas=1',
       ]))

k8s_resource(
  'wmstats-prod',
  port_forwards='8080:8080')
