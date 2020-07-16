# DevOps Project

# Ansible

-Install dependencies

```bash
ansible-galaxy collection install ericsysmin.kubernetes
ansible-galaxy install geerlingguy.jenkins
ansible-galaxy install andrewrothstein.clair_scanner
```

- Run playbooks

```bash
ansible-playbook kubernetes/playbook.yml
ansible-playbook jenkins/playbook.yml
ansible-playbook clair/playbook.yml
```

# Jenkins

1. make build-base # (build an image localy) with make build-base
2. make build # (build an image tag) of the build-base image
3. make build-test # (build image with junit for testing)
4. make test-unit # Run unit tests and create report.xml file
