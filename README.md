<!--
parent:
  order: false
-->

<div align="center">
  <h1> Cosmos SDK </h1>
</div>
 

The Reapchain Cosmos SDK is a framework for building blockchain applications. [Reapchain Core](https://github.com/reapchain/reapchain-core) and the Cosmos SDK are written in the Golang programming language. Cosmos SDK is used to build [Reapchain MainNet](https://github.com/reapchain/reapchain).

**WARNING**: The project has mostly stabilized, but we are still making some breaking changes.

**Note**: Requires [Go 1.18+](https://golang.org/dl/)


## Introduction

The Reapchain Cosmos-SDK is a fork of the Cosmos-SDK [v0.45.7](https://github.com/cosmos/cosmos-sdk/tree/v0.45.7). We, the Reapchain Team, have integrated our Consensus Engine which we call the Reapchain Core on top of the Open Source code available in the Cosmos Ecosystem. By utilizing the Reapchain Cosmos-SDK as well as the Reapchain Core Consensus Engine, you will be able to create your own Proof-of-Stake blockchains.

## Quick Start

**Build**
```
make build
make install

# you can see the version!
simd version
```

&nbsp;
