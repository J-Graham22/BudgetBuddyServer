create table Households (
  id serial primary key,
  name varchar(100) not null
);

create table Users (
  id serial primary key,
  username varchar(50) not null,
  name varchar(50) not null,
  email varchar(100) not null,
  password bytea not null
);

create table Accounts (
  id serial primary key,
  name varchar(50) not null,
  account_owner integer null, -- if set, the account is owned by this user
  household_id integer null -- if set, the account is owned by the household
);

create table Budgets (
  id serial primary key,
  name varchar(100) not null,
  start_date date not null,
  end_date date not null,
  household_id integer not null,
  FOREIGN KEY(household_id) REFERENCES Households(id)
);

create table UserHouseholdMapping (
  id serial primary key,
  household_id integer not null,
  user_id integer not null,
  FOREIGN KEY(household_id) REFERENCES Households(id),
  FOREIGN KEY(user_id) REFERENCES Users(id)
);

create table Categories (
  id serial primary key,
  name varchar(50) not null,
  description varchar(255) null,
  household_id integer not null,
  parent_category_id integer null,
  FOREIGN KEY(household_id) REFERENCES Households(id)
);

create table Transactions (
  id serial primary key,
  created_by_user_id integer not null,
  updated_by_user_id integer not null,
  transaction_date date not null,
  name varchar(50) not null,
  amount decimal(10, 2) not null,
  description varchar(255) default null,
  category integer null,
  household_id integer not null,
  is_income boolean not null,
  account_id integer not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  FOREIGN KEY(created_by_user_id) REFERENCES Users(id),
  FOREIGN KEY(updated_by_user_id) REFERENCES Users(id),
  FOREIGN KEY(category) REFERENCES Categories(id),
  FOREIGN KEY(account_id) REFERENCES Accounts(id)
);

