# Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG CUDA_SAMPLE_IMAGE=undefined
ARG GOLANG_VERSION=x.x.x

FROM nvcr.io/nvidia/cuda:12.8.1-base-ubi9 as build

RUN dnf install -y wget make git gcc

ARG GOLANG_VERSION=0.0.0
RUN set -eux; \
    \
    arch="$(uname -m)"; \
    case "${arch##*-}" in \
        x86_64 | amd64) ARCH='amd64' ;; \
        ppc64el | ppc64le) ARCH='ppc64le' ;; \
        aarch64 | arm64) ARCH='arm64' ;; \
        *) echo "unsupported architecture" ; exit 1 ;; \
    esac; \
    wget -nv -O - https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-${ARCH}.tar.gz \
    | tar -C /usr/local -xz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR /build
COPY . .

ARG VERSION="unknown"
ARG GIT_COMMIT="unknown"

RUN mkdir /artifacts
RUN make -C validator/ validator
RUN cp /build/validator/validator /artifacts/nvidia-validator
RUN cp /build/validator/plugin-workload-validation.yaml /artifacts/
RUN cp /build/validator/cuda-workload-validation.yaml /artifacts/

FROM ${CUDA_SAMPLE_IMAGE} AS sample-builder
RUN mkdir /artifacts
RUN cp /cuda-samples/vectorAdd /artifacts/vectorAdd

FROM nvcr.io/nvidia/cuda:12.8.1-base-ubi9

# Remove CUDA libs(compat etc) in favor of libs installed by the NVIDIA driver
RUN dnf remove -y cuda-*

RUN dnf install -y \
        kmod \
        pciutils && \
    rm -rf /var/cache/yum/*

RUN mkdir /licenses && mv /NGC-DL-CONTAINER-LICENSE /licenses/NGC-DL-CONTAINER-LICENSE
COPY --from=build /artifacts/nvidia-validator  /usr/bin/nvidia-validator
COPY --from=sample-builder /artifacts/vectorAdd /usr/bin/vectorAdd
RUN mkdir -p /var/nvidia/manifests
COPY --from=build /artifacts/plugin-workload-validation.yaml /var/nvidia/manifests
COPY --from=build /artifacts/cuda-workload-validation.yaml /var/nvidia/manifests

ENV NVIDIA_DISABLE_REQUIRE="true"
ENV NVIDIA_VISIBLE_DEVICES=void
ENV NVIDIA_DRIVER_CAPABILITIES=utility,compute

ARG VERSION="unknown"
ARG GIT_COMMIT="unknown"

LABEL version="${VERSION}"
LABEL release="N/A"
LABEL vendor="NVIDIA"
LABEL io.k8s.display-name="NVIDIA Validators for the GPU Operator Components"
LABEL name="NVIDIA Validator for the GPU Operator"
LABEL summary="NVIDIA Validator for the GPU Operator"
LABEL description="See summary"
LABEL vsc-ref=${GIT_COMMIT}

# Install / upgrade packages here that are required to resolve CVEs
ARG CVE_UPDATES
RUN if [ -n "${CVE_UPDATES}" ]; then \
        dnf update -y ${CVE_UPDATES} && \
        rm -rf /var/cache/yum/*; \
    fi

ENTRYPOINT ["/bin/bash"]
