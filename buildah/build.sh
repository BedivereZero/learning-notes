#!/bin/bash

container=$(buildah from alpine:3.14)
buildah copy "${container}" runecho /usr/bin
buildah config --cmd /usr/bin/runecho "${container}"
buildah config --created-by "ipbabble" "${container}"
buildah config --author "BedivereZero" "${container}"
buildah config --label name=alpine-bashecho "${container}"
buildah commit "${container}" image
