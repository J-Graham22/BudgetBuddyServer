create table `Households` (
  `id` int(11) not null auto_increment,
  `name` varchar(100) not null,
  PRIMARY KEY(`id`)
);

create table `Users` (
  `id` int(11) not null auto_increment,
  `username` varchar(50) not null,
  `name` varchar(50) not null,
  `email` varchar(100) not null,
  PRIMARY KEY(`id`)
);

create table `Accounts` (
  `id` int(11) not null auto_increment,
  `name` varchar(50) not null,
  `account_owner` int(11) null, -- if set, the account is owned by this user
  `household_id` int(11) null, -- if set, the account is owned by the household
  PRIMARY KEY(`id`)
);

create table `Budgets` (
  `id` int(11) not null auto_increment,
  `name` varchar(100) not null,
  `start_date` date not null,
  `end_date` date not null,
  `household_id` int(11) not null,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`household_id`) REFERENCES Households(`id`)
);

create table `UserHouseholdMapping` (
  `id` int(11) not null auto_increment,
  `household_id` int(11) not null,
  `user_id` int(11) not null,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`household_id`) REFERENCES Households(`id`),
  FOREIGN KEY(`user_id`) REFERENCES Users(`id`)
);

create table `Categories` (
  `id` int(11) not null auto_increment,
  `name` varchar(50) not null,
  `description` varchar(255) null,
  `household_id` int(11) not null,
  `parent_category_id` int(11) null,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`household_id`) REFERENCES Households(`id`)
);

create table `Transactions` (
  `id` int(11) not null auto_increment,
  `created_by_user_id` int(11) not null,
  `updated_by_user_id` int(11) not null,
  `transaction_date` date not null,
  `name` varchar(50) not null,
  `amount` decimal(10, 2) not null,
  `description` varchar(255) default null,
  `category` int(11) null,
  `transaction_type` enum('income', 'expense') not null,
  `account_id` int(11) not null,
  `created_at` timestamp null default current_timestamp(),
  `updated_at` timestamp null default current_timestamp(),
  PRIMARY KEY(`id`),
  FOREIGN KEY(`created_by_user_id`) REFERENCES Users(`id`),
  FOREIGN KEY(`updated_by_user_id`) REFERENCES Users(`id`),
  FOREIGN KEY(`category`) REFERENCES Categories(`id`),
  FOREIGN KEY(`account_id`) REFERENCES Accounts(`id`)
);

