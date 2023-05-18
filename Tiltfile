# Tiltfile

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_resource', 'helm_resource', 'helm_repo')


compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux'

services = [
  struct(
    name='mano-server',
    service_name='mano-server',
    port=['9090:9090'],
    main_path='cmd/main.go',
    helm_path='mano-server/k8s/chart',
    dockerfile='mano-server/docker/developer.Dockerfile',
    value_file='mano-server/k8s/chart/env/local/values.yaml',
    docker_allow_list=[
      'mano-server/',
    ],
    live_updates=[
      sync('mano-server/bin/mano-server', '/app/bin/mano-server'),
    ],
  ),
]

for service in services:
  # Compile command
  compile_service_cmd = '{compile_opt} go build -o bin/{service_name} {main_path}'.format(
      compile_opt=compile_opt,
      service_name=service.service_name,
      main_path=service.main_path
    )

  # Compile locally
  local_resource(
    '{service_name}-compile'.format(service_name=service.service_name),
    compile_service_cmd,
    dir=service.name,
    deps=[service.name],
    ignore=[service.name + '/bin'],
    labels=[service.service_name],
  )

  # Build service docker image
  docker_build_with_restart(
    '{service_name}-image'.format(service_name=service.service_name),
    '.',
    entrypoint='bin/{service_name}'.format(service_name=service.service_name),
    dockerfile=service.dockerfile,
    only=service.docker_allow_list,
    live_update=service.live_updates,
  )

  # Deploy service using helm
  helm_resource(
    service.service_name,
    service.helm_path,
    image_deps=['{name}-image'.format(name=service.service_name)],
    image_keys=[('image.repository', 'image.tag')],
    port_forwards=service.port,
    flags=[
      '-f',
      service.value_file,
    ],
    labels=service.service_name,
  )
