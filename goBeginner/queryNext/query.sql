--create database
create database "Ojol";

-- droptable
DROP TABLE IF EXISTS "Order";

DROP TABLE IF EXISTS "LoginLog";

DROP TABLE IF EXISTS "RegionArea";

DROP TABLE IF EXISTS "Region";

DROP TABLE IF EXISTS "Driver";

DROP TABLE IF EXISTS "Customer";

DROP TABLE IF EXISTS "Admin";

DROP TABLE IF EXISTS "User";

--create table
create table
    if not exists "User" (
        "id" serial primary key,
        "username" varchar(50) not null,
        "password" varchar(50) not null,
        "role" varchar(20) not null
    );

create table
    if not exists "Admin" (
        "id" serial primary key,
        "name" varchar(50) not null,
        "user_id" integer references "User" ("id")
    );

create table
    if not exists "Customer" (
        "id" serial primary key,
        "name" varchar(50) not null,
        "phoneNumber" varchar(13) not null,
        "address" varchar(255) not null,
        "user_id" integer references "User" ("id")
    );

create table
    if not exists "Driver" (
        "id" serial primary key,
        "name" varchar(50) not null,
        "phoneNumber" varchar(13) not null,
        "address" varchar(255) not null,
        "vehicle" varchar(100) not null,
        "user_id" integer references "User" ("id")
    );

create table
    if not exists "Region" (
        "id" serial primary key,
        "name" varchar(100) not null
    );

create table
    if not exists "RegionArea" (
        "id" serial primary key,
        "area" varchar(255) not null,
        "region_id" integer references "Region" ("id")
    );

create table
    if not exists "LoginLog" (
        "id" serial primary key,
        "status" bool,
        "admin_id" integer references "Admin" ("id"),
        "customer_id" integer references "Customer" ("id")
    );

create table
    if not exists "Order" (
        "id" serial primary key,
        "admin_id" integer references "Admin" ("id"),
        "customer_id" integer references "Customer" ("id"),
        "driver_id" integer references "Driver" ("id"),
        "regionArea_id" integer references "RegionArea" ("id"),
        "dateOrder" timestamp,
        "status" bool
    );

--
--
-- Dummy data untuk tabel "User"
INSERT INTO
    "User" ("username", "password", "role")
VALUES
    ('admin1', 'password123', 'admin'),
    ('cust1', 'password123', 'customer'),
    ('cust2', 'password123', 'customer'),
    ('driver1', 'password123', 'driver'),
    ('driver2', 'password123', 'driver');

-- Dummy data untuk tabel "Admin"
INSERT INTO
    "Admin" ("name", "user_id")
VALUES
    ('Admin One', 1);

-- Dummy data untuk tabel "Customer"
INSERT INTO
    "Customer" ("name", "phoneNumber", "address", "user_id")
VALUES
    (
        'Customer One',
        '081234567890',
        'Jl. Merdeka No.1',
        2
    ),
    (
        'Customer Two',
        '081234567891',
        'Jl. Sudirman No.2',
        3
    );

-- Dummy data untuk tabel "Driver"
INSERT INTO
    "Driver" (
        "name",
        "phoneNumber",
        "address",
        "vehicle",
        "user_id"
    )
VALUES
    (
        'Driver One',
        '081234567892',
        'Jl. Gatot Subroto No.3',
        'car',
        4
    ),
    (
        'Driver Two',
        '081234567893',
        'Jl. Thamrin No.4',
        'motor',
        5
    );

-- Dummy data untuk tabel "Region"
INSERT INTO
    "Region" ("name")
VALUES
    ('Jakarta'),
    ('Bogor'),
    ('Depok'),
    ('Karawang');

INSERT INTO
    "RegionArea" ("area", "region_id")
VALUES
    ('Central Jakarta', 1),
    ('West Jakarta', 1),
    ('North Jakarta', 1),
    ('West Bogor', 2),
    ('Nort Bogor', 2),
    ('West Karawang', 4),
    ('East Karawang', 4);

-- Dummy data untuk tabel "LoginLog"
INSERT INTO
    "LoginLog" ("status", "admin_id", "customer_id")
VALUES
    (true, 1, 2),
    (false, 1, 1);

--
---- Dummy data untuk tabel "Order"
INSERT INTO
    "Order" (
        "admin_id",
        "customer_id",
        "driver_id",
        "regionArea_id",
        "dateOrder",
        "status"
    )
VALUES
    (1, 2, 1, 1, '2024-10-29 01:30:00', TRUE),
    (1, 2, 1, 2, '2024-11-29 09:30:00', TRUE),
    (1, 1, 2, 3, '2024-10-29 08:20:00', FALSE),
    (1, 1, 2, 1, '2024-12-29 10:00:00', TRUE);

-- Login log
SELECT
    (
        SELECT
            COUNT(*)
        FROM
            "LoginLog"
        WHERE
            "status" = true
    ) AS "active_customers",
    (
        SELECT
            COUNT(*)
        FROM
            "LoginLog"
        WHERE
            "status" = false
    ) AS "inactive_customers";

-- Total data order every month
SELECT
    TO_CHAR ("dateOrder", 'YYYY-MM') AS "month",
    COUNT("id") AS "total_order"
FROM
    "Order"
WHERE
    "status" = TRUE
GROUP BY
    "month"
order by
    "month" desc;

-- Total order by customer every month
SELECT
    c."name" AS "customer_name",
    TO_CHAR ("dateOrder", 'YYYY-MM') AS "month",
    COUNT(o."id") AS "total_order"
FROM
    "Order" o
    JOIN "Customer" c ON o."customer_id" = c."id"
WHERE
    "status" = TRUE
GROUP BY
    "month",
    "customer_name"
ORDER BY
    "month",
    "total_order" DESC;

-- Often order every month based on region
SELECT
    r."name" AS "region_name",
    ra."area" AS "region_area",
    TO_CHAR (o."dateOrder", 'YYYY-MM') AS "month",
    COUNT(DISTINCT o."id") AS "total_order"
FROM
    "Order" o
    JOIN "RegionArea" ra ON o."regionArea_id" = ra."id"
    JOIN "Region" r ON ra."region_id" = r."id"
where
    "status" = true
GROUP BY
    r."name",
    ra."area",
    "month"
ORDER by
    "month" desc,
    "total_order" DESC;

-- Often order by time
SELECT
    TO_CHAR ("dateOrder", 'HH24:MI:SS') AS "hour",
    COUNT("id") AS "total_order"
FROM
    "Order"
GROUP BY
    "hour"
ORDER BY
    "total_order" DESC;

-- Driver often take orders every month
SELECT
    d."name" AS "driver_name",
    TO_CHAR ("dateOrder", 'YYYY-MM-DD') AS "month",
    COUNT(o."id") AS "total_order"
FROM
    "Order" o
    JOIN "Driver" d ON o."driver_id" = d."id"
GROUP BY
    "month",
    "driver_name"
ORDER BY
    "month",
    "total_order" DESC;