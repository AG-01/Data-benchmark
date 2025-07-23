-- Sample benchmark queries for comparing Hive vs Iceberg performance

-- Query 1: Simple SELECT with WHERE clause
-- Tests basic filtering performance
SELECT c_custkey, c_name, c_acctbal 
FROM customer_hive 
WHERE c_nationkey = 15 AND c_acctbal > 1000.00;

-- Query 2: Aggregation query
-- Tests aggregation performance
SELECT 
    c_mktsegment,
    COUNT(*) as customer_count,
    AVG(c_acctbal) as avg_balance,
    SUM(c_acctbal) as total_balance
FROM customer_hive 
GROUP BY c_mktsegment
ORDER BY total_balance DESC;

-- Query 3: JOIN query
-- Tests join performance between tables
SELECT 
    c.c_name,
    o.o_orderdate,
    o.o_totalprice
FROM customer_hive c
JOIN orders_hive o ON c.c_custkey = o.o_custkey
WHERE c.c_nationkey = 1 
  AND o.o_orderdate >= DATE '2023-01-01'
ORDER BY o.o_totalprice DESC
LIMIT 100;

-- Query 4: Complex aggregation with multiple JOINs
-- Tests complex query performance
SELECT 
    c.c_mktsegment,
    YEAR(o.o_orderdate) as order_year,
    COUNT(DISTINCT o.o_orderkey) as order_count,
    SUM(o.o_totalprice) as total_revenue,
    AVG(o.o_totalprice) as avg_order_value
FROM customer_hive c
JOIN orders_hive o ON c.c_custkey = o.o_custkey
WHERE o.o_orderdate >= DATE '2022-01-01'
GROUP BY c.c_mktsegment, YEAR(o.o_orderdate)
HAVING COUNT(DISTINCT o.o_orderkey) > 10
ORDER BY total_revenue DESC;

-- Query 5: Window function query
-- Tests analytical function performance
SELECT 
    c_custkey,
    c_name,
    c_acctbal,
    c_mktsegment,
    RANK() OVER (PARTITION BY c_mktsegment ORDER BY c_acctbal DESC) as balance_rank,
    AVG(c_acctbal) OVER (PARTITION BY c_mktsegment) as segment_avg_balance
FROM customer_hive
WHERE c_acctbal > 0;

-- Query 6: Time-based analysis
-- Tests partition pruning and date filtering
SELECT 
    DATE_TRUNC('month', o_orderdate) as order_month,
    o_orderstatus,
    COUNT(*) as order_count,
    SUM(o_totalprice) as monthly_revenue,
    MIN(o_totalprice) as min_order,
    MAX(o_totalprice) as max_order
FROM orders_hive
WHERE o_orderdate BETWEEN DATE '2023-01-01' AND DATE '2023-12-31'
GROUP BY DATE_TRUNC('month', o_orderdate), o_orderstatus
ORDER BY order_month, o_orderstatus;

-- Query 7: Large table scan with filters
-- Tests performance on larger datasets
SELECT 
    l_returnflag,
    l_linestatus,
    SUM(l_quantity) as sum_qty,
    SUM(l_extendedprice) as sum_base_price,
    SUM(l_extendedprice * (1 - l_discount)) as sum_disc_price,
    SUM(l_extendedprice * (1 - l_discount) * (1 + l_tax)) as sum_charge,
    AVG(l_quantity) as avg_qty,
    AVG(l_extendedprice) as avg_price,
    AVG(l_discount) as avg_disc,
    COUNT(*) as count_order
FROM lineitem_hive
WHERE l_shipdate <= DATE '2023-09-01'
GROUP BY l_returnflag, l_linestatus
ORDER BY l_returnflag, l_linestatus;

-- Query 8: Complex 3-way JOIN
-- Tests multi-table join performance
SELECT 
    c.c_name,
    c.c_custkey,
    o.o_orderkey,
    o.o_orderdate,
    o.o_totalprice,
    SUM(l.l_quantity) as total_quantity
FROM customer_hive c
JOIN orders_hive o ON c.c_custkey = o.o_custkey
JOIN lineitem_hive l ON o.o_orderkey = l.l_orderkey
WHERE c.c_mktsegment = 'BUILDING'
  AND o.o_orderdate >= DATE '2023-01-01'
  AND l.l_returnflag = 'R'
GROUP BY c.c_name, c.c_custkey, o.o_orderkey, o.o_orderdate, o.o_totalprice
HAVING SUM(l.l_quantity) > 100
ORDER BY o.o_totalprice DESC
LIMIT 50;

-- Query 9: Subquery with aggregation
-- Tests subquery performance
SELECT 
    c.c_custkey,
    c.c_name,
    c.c_acctbal,
    (SELECT COUNT(*) 
     FROM orders_hive o 
     WHERE o.o_custkey = c.c_custkey) as order_count,
    (SELECT SUM(o.o_totalprice) 
     FROM orders_hive o 
     WHERE o.o_custkey = c.c_custkey) as total_spent
FROM customer_hive c
WHERE c.c_acctbal > (
    SELECT AVG(c2.c_acctbal) 
    FROM customer_hive c2 
    WHERE c2.c_mktsegment = c.c_mktsegment
)
ORDER BY total_spent DESC NULLS LAST
LIMIT 100;

-- Query 10: Advanced analytics with percentiles
-- Tests advanced analytical functions
SELECT 
    c_mktsegment,
    COUNT(*) as customer_count,
    MIN(c_acctbal) as min_balance,
    PERCENTILE_CONT(0.25) WITHIN GROUP (ORDER BY c_acctbal) as q1_balance,
    PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY c_acctbal) as median_balance,
    PERCENTILE_CONT(0.75) WITHIN GROUP (ORDER BY c_acctbal) as q3_balance,
    MAX(c_acctbal) as max_balance,
    STDDEV(c_acctbal) as balance_stddev
FROM customer_hive
WHERE c_acctbal > -999
GROUP BY c_mktsegment
ORDER BY median_balance DESC;
