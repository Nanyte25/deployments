{
  local defaultConfig = self,

  name: 'observatorium-xyz',
  namespace: 'observatorium',
  thanosVersion: 'master-2020-08-12-70f89d83',  // Fixes a blocker issue in v0.13.0-rc.0
  thanosImage: 'quay.io/thanos/thanos:' + defaultConfig.thanosVersion,
  objectStorageConfig: {
    thanos: {
      name: 'thanos-objectstorage',
      key: 'thanos.yaml',
    },
    loki: {
      secretName: 'loki-objectstorage',
      endpointKey: 'endpoint',
      bucketsKey: 'buckets',
      regionKey: 'region',
      accessKeyIdKey: 'aws_access_key_id',
      secretAccessKeyKey: 'aws_secret_access_key',
    },
  },

  hashrings: [
    {
      hashring: 'default',
      tenants: [
        // Match all for now
        // 'foo',
        // 'bar',
      ],
    },
  ],

  compact: {
    image: defaultConfig.thanosImage,
    version: defaultConfig.thanosVersion,
    objectStorageConfig: defaultConfig.objectStorageConfig.thanos,
    retentionResolutionRaw: '14d',
    retentionResolution5m: '1s',
    retentionResolution1h: '1s',
    replicas: 1,
    enableDownsampling: false,
    volumeClaimTemplate: {
      spec: {
        accessModes: ['ReadWriteOnce'],
        resources: {
          requests: {
            storage: '50Gi',
          },
        },
      },
    },
    logLevel: 'info',
  },

  thanosReceiveController: {
    local trcConfig = self,
    version: 'master-2020-02-06-b66e0c8',
    image: 'quay.io/observatorium/thanos-receive-controller:' + trcConfig.version,
    hashrings: defaultConfig.hashrings,
  },

  receivers: {
    image: defaultConfig.thanosImage,
    version: defaultConfig.thanosVersion,
    hashrings: defaultConfig.hashrings,
    objectStorageConfig: defaultConfig.objectStorageConfig.thanos,
    replicas: 1,
    volumeClaimTemplate: {
      spec: {
        accessModes: ['ReadWriteOnce'],
        resources: {
          requests: {
            storage: '50Gi',
          },
        },
      },
    },
    logLevel: 'info',
    debug: '',
  },

  rule: {
    image: defaultConfig.thanosImage,
    version: defaultConfig.thanosVersion,
    objectStorageConfig: defaultConfig.objectStorageConfig.thanos,
    replicas: 1,
    volumeClaimTemplate: {
      spec: {
        accessModes: ['ReadWriteOnce'],
        resources: {
          requests: {
            storage: '50Gi',
          },
        },
      },
    },
  },

  store: {
    image: defaultConfig.thanosImage,
    version: defaultConfig.thanosVersion,
    objectStorageConfig: defaultConfig.objectStorageConfig.thanos,
    shards: 1,
    volumeClaimTemplate: {
      spec: {
        accessModes: ['ReadWriteOnce'],
        resources: {
          requests: {
            storage: '50Gi',
          },
        },
      },
    },
    logLevel: 'info',
  },

  storeCache: {
    local scConfig = self,
    replicas: 1,
    version: '1.6.3-alpine',
    image: 'docker.io/memcached:' + scConfig.version,
    exporterVersion: 'v0.6.0',
    exporterImage: 'prom/memcached-exporter:' + scConfig.exporterVersion,
    memoryLimitMb: 1024,
  },

  query: {
    image: defaultConfig.thanosImage,
    version: defaultConfig.thanosVersion,
    replicas: 1,
  },

  queryFrontend: {
    image: defaultConfig.thanosImage,
    version: defaultConfig.thanosVersion,
    replicas: 1,
  },

  api: {
    local apiConfig = self,
    version: 'master-2020-10-27-v0.1.1-187-g76a4c28',
    image: 'quay.io/observatorium/observatorium:' + apiConfig.version,
    replicas: 1,
  },

  gubernator: {
    local gubernatorConfig = self,
    version: '1.0.0-rc.1',
    image: 'thrawn01/gubernator:' + gubernatorConfig.version,
    replicas: 1,
  },

  apiQuery: {
    image: defaultConfig.thanosImage,
    version: defaultConfig.thanosVersion,
  },

  up: {
    local upConfig = self,
    version: 'master-2020-06-03-8a20b4e',
    image: 'quay.io/observatorium/up:' + upConfig.version,
  },

  lokiRingStore: {
    local lokiRingStoreConfig = self,
    version: 'v3.4.9',
    image: 'quay.io/coreos/etcd:' + lokiRingStoreConfig.version,
    replicas: 1,
  },

  lokiCaches: {
    local scConfig = self,
    version: '1.6.3-alpine',
    image: 'docker.io/memcached:' + scConfig.version,
    exporterVersion: 'v0.6.0',
    exporterImage: 'prom/memcached-exporter:' + scConfig.exporterVersion,
    replicas: {
      chunk_cache: 1,
      index_query_cache: 1,
      results_cache: 1,
    },
  },

  loki+: {
    local lokiConfig = self,
    version: '2.0.0',
    image: 'docker.io/grafana/loki:' + lokiConfig.version,
    objectStorageConfig: defaultConfig.objectStorageConfig.loki,
    replicas: {
      compactor: 1,
      distributor: 1,
      ingester: 1,
      querier: 1,
      query_frontend: 1,
    },
    volumeClaimTemplate: {
      spec: {
        accessModes: ['ReadWriteOnce'],
        resources: {
          requests: {
            storage: '250Mi',
          },
        },
      },
    },
  },
}
