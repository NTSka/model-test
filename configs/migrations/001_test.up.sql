CREATE TABLE IF NOT EXISTS test (
                      date DATE,
                      assets Array(String),
                      event_host String,
                      event_type LowCardinality(String),
                      action LowCardinality(String),
                      importance Int8,
                      object String,
                      meta1 String,
                      meta2 String,
                      meta3 String,
                      meta4 String,
                      meta5 String
) ENGINE = MergeTree()
      ORDER BY (toStartOfDay(date), event_type, action, importance,
                event_host, assets, object, meta1, meta2, meta3, meta4, meta5)
     PARTITION BY toYYYYMM(date)