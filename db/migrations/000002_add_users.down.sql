alter table accounts drop constraint "accounts_owner_fkey";

alter table accounts drop constraint "unique_owner_currency";

drop table users;