-- Sample dataset schemas for benchmarking

-- TPC-H inspired customer table
CREATE TABLE IF NOT EXISTS customer_hive (
    c_custkey BIGINT,
    c_name VARCHAR(25),
    c_address VARCHAR(40),
    c_nationkey BIGINT,
    c_phone VARCHAR(15),
    c_acctbal DECIMAL(15,2),
    c_mktsegment VARCHAR(10),
    c_comment VARCHAR(117)
) 
USING HIVE
PARTITIONED BY (c_nationkey)
LOCATION 's3a://benchmark-data/customer_hive/';

-- Same table in Iceberg format
CREATE TABLE IF NOT EXISTS customer_iceberg (
    c_custkey BIGINT,
    c_name VARCHAR(25),
    c_address VARCHAR(40),
    c_nationkey BIGINT,
    c_phone VARCHAR(15),
    c_acctbal DECIMAL(15,2),
    c_mktsegment VARCHAR(10),
    c_comment VARCHAR(117)
) 
USING ICEBERG
PARTITIONED BY (c_nationkey)
LOCATION 's3a://benchmark-data/customer_iceberg/';

-- Orders table for join benchmarks
CREATE TABLE IF NOT EXISTS orders_hive (
    o_orderkey BIGINT,
    o_custkey BIGINT,
    o_orderstatus VARCHAR(1),
    o_totalprice DECIMAL(15,2),
    o_orderdate DATE,
    o_orderpriority VARCHAR(15),
    o_clerk VARCHAR(15),
    o_shippriority INTEGER,
    o_comment VARCHAR(79)
)
USING HIVE
PARTITIONED BY (year(o_orderdate))
LOCATION 's3a://benchmark-data/orders_hive/';

CREATE TABLE IF NOT EXISTS orders_iceberg (
    o_orderkey BIGINT,
    o_custkey BIGINT,
    o_orderstatus VARCHAR(1),
    o_totalprice DECIMAL(15,2),
    o_orderdate DATE,
    o_orderpriority VARCHAR(15),
    o_clerk VARCHAR(15),
    o_shippriority INTEGER,
    o_comment VARCHAR(79)
)
USING ICEBERG
PARTITIONED BY (year(o_orderdate))
LOCATION 's3a://benchmark-data/orders_iceberg/';

-- Lineitem table for large dataset benchmarks
CREATE TABLE IF NOT EXISTS lineitem_hive (
    l_orderkey BIGINT,
    l_partkey BIGINT,
    l_suppkey BIGINT,
    l_linenumber INTEGER,
    l_quantity DECIMAL(15,2),
    l_extendedprice DECIMAL(15,2),
    l_discount DECIMAL(15,2),
    l_tax DECIMAL(15,2),
    l_returnflag VARCHAR(1),
    l_linestatus VARCHAR(1),
    l_shipdate DATE,
    l_commitdate DATE,
    l_receiptdate DATE,
    l_shipinstruct VARCHAR(25),
    l_shipmode VARCHAR(10),
    l_comment VARCHAR(44)
)
USING HIVE
PARTITIONED BY (year(l_shipdate), month(l_shipdate))
LOCATION 's3a://benchmark-data/lineitem_hive/';

CREATE TABLE IF NOT EXISTS lineitem_iceberg (
    l_orderkey BIGINT,
    l_partkey BIGINT,
    l_suppkey BIGINT,
    l_linenumber INTEGER,
    l_quantity DECIMAL(15,2),
    l_extendedprice DECIMAL(15,2),
    l_discount DECIMAL(15,2),
    l_tax DECIMAL(15,2),
    l_returnflag VARCHAR(1),
    l_linestatus VARCHAR(1),
    l_shipdate DATE,
    l_commitdate DATE,
    l_receiptdate DATE,
    l_shipinstruct VARCHAR(25),
    l_shipmode VARCHAR(10),
    l_comment VARCHAR(44)
)
USING ICEBERG
PARTITIONED BY (year(l_shipdate), month(l_shipdate))
LOCATION 's3a://benchmark-data/lineitem_iceberg/';
