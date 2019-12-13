# Go GitLab Template

A simple hello golang app configured to run on GitLab.

## Getting Started

    make install-tools build test-behavior test-security

### Gitlab Configuration

The `.gitlab-ci.yml` in this repo is derived from GitLab's Auto DevOps
Template. It is tweaked to deploy apps on the same cluster that GitLab
is running. It leverages the nginx-ingress-controller and cert-manager
(with both a staging and production cluster-issuer) as GitLab. If
it does not work out of the box for you, the first place to look
would be at the following variables in `.gitlab-ci.yml`:

- `HELM_UPGRADE_EXTRA_ARGS`
- `HTTPS_PORT`
