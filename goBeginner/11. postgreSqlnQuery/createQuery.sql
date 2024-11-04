--create database
create database "E-Learning";

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
        "user_id" integer references "User" ("id"),
        "class_id" integer references "User" ("id")
    );

create table
    if not exists "Student" (
        "id" serial primary key,
        "user_id" integer references "User" ("id"),
        "name" varchar(50) not null,
        "phone_number" varchar(13) not null,
        "address" varchar(255) not null
    );

create table
    if not exists "Mentor" (
        "id" serial primary key,
        "name" varchar(50) not null,
        "degree" varchar(50) not null,
        "experience" integer,
        "user_id" integer references "User" ("id")
    );

create table
    if not exists "Class" (
        "id" serial primary key,
        "name" varchar(255) not null
    );

create table
    if not exists "Class_Student" (
        "id" serial primary key,
        "class_id" integer references "Class" ("id"),
        "student_id" integer references "Student" ("id"),
        "mentor_id" integer references "Mentor" ("id")
    );

create table
    if not exists "Class_Mentor" (
        "id" serial primary key,
        "class_id" integer references "Class" ("id"),
        "mentor_id" integer references "Mentor" ("id")
    );

create table
    if not exists "Schedule" (
        "id" serial primary key,
        "class_id" integer references "Class" ("id"),
        "date" date not null,
        "start_time" time not null,
        "end_time" time not null
    );

create table
    if not exists "Material" (
        "id" serial primary key,
        "class_id" integer references "Class" ("id"),
        "title" varchar(255) not null,
        "content" varchar(255) not null
    );

create table
    if not exists "Assignment" (
        "id" serial primary key,
        "class_id" integer references "Class" ("id"),
        "title" varchar(255) not null,
        "description" varchar(255) not null,
        "deadline" time not null
    );

create table
    if not exists "Attendance" (
        "id" serial primary key,
        "schedule_id" integer references "Schedule" ("id"),
        "user_id" integer references "User" ("id"),
        "status" boolean
    );

create table
    if not exists "Announcement" (
        "id" serial primary key,
        "admin_id" integer references "Admin" ("id"),
        "title" varchar(255) not null,
        "content" varchar(255) not null
    );

create table
    if not exists "Announcement_Visibility" (
        "id" serial primary key,
        "announcement_id" integer references "Announcement" ("id"),
        "user_id" integer references "User" ("id")
    );

create table
    if not exists "Grade" (
        "id" serial primary key,
        "class_id" integer references "Class" ("id"),
        "student_id" integer references "Student" ("id"),
        "grade" integer
    );

create table
    if not exists "LeaderBoard" (
        "id" serial primary key,
        "student_id" integer references "Student" ("id"),
        "point" integer
    );