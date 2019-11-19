CREATE DATABASE fair;


USE fair;

CREATE TABLE fairfood (
    FoodID int auto_increment primary key,
    FairID int,
    FoodName varchar(300),
    FoodDesc varchar(600),
    FoodLocation varchar(300),
    FoodType varchar(300),
    VendorName varchar(300),
    FoodPrice double

);


INSERT INTO fair.fairfood (FairID, FoodName, FoodDesc, FoodLocation, FoodType, VendorName, FoodPrice) VALUES
(1, 'Fried Dough', 'Dough that has been fried with powdered sugar or cinamon', 'Whites Farm Road', 'Desert', 'Food For Fun', 5.50),
(1, 'Fried Oreos', 'Oreos that have been fried', 'Whites Farm Road', 'Desert', 'Food For Fun', 6.00),
(1, 'Cheeseburger Sliders', 'Oreos that have been fried', 'Canfield  Road', 'Dinner', 'Harrys Sliders', 8.70),
(1, 'Corn Dogs', 'Fried hot dogs on a stick', 'Canfield  Road', 'Dinner', 'Durham Elks Club', 3.50),
(1, 'Pulled Pork', 'Smoked pulled pork sandwitch topped with pickles ', 'Town Hall  Road', 'Dinner', 'Porkys Barbecue', 11.00)