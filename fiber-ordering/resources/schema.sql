DROP DATABASE shop;
CREATE DATABASE shop;
USE shop;

CREATE TABLE IF NOT EXISTS customer (
    `customer_id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_full_name` varchar(30) NOT NULL,
    `customer_email` varchar(50) NOT NULL,
    `customer_phone_number` varchar(15) NOT NULL,
    `customer_username` varchar(30) NOT NULL,
    `customer_password` varchar(255) NOT NULL,
    `account_status` int(1) NOT NULL DEFAULT 1,
    PRIMARY KEY (`customer_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS product (
    `product_id` int(11) NOT NULL AUTO_INCREMENT,
    `product_category_id` int(11) NOT NULL,
    `product_code` varchar(100) NOT NULL,
    `product_name` varchar(100) NOT NULL,
    `product_image` varchar(100) NOT NULL,
    `product_price` float NOT NULL,
    `product_status` int(1) NOT NULL,
    PRIMARY KEY (`product_id`),
    KEY `product_category_id` (`product_category_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS product_category (
    `product_category_id` int(11) NOT NULL AUTO_INCREMENT,
    `product_category_name` varchar(50) NOT NULL,
    PRIMARY KEY (`product_category_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS orders (
    `order_id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL,
    `total_price` float,
    PRIMARY KEY (`order_id`),
    KEY `customer_id` (`customer_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS order_details (
    `order_details_id` int(11) NOT NULL AUTO_INCREMENT,
    `order_id` int(11) NOT NULL,
    `product_id` int(11) NOT NULL,
    `quantity` int(11) NOT NULL,
    `total_price` float,
    PRIMARY KEY (`order_details_id`),
    KEY `order_id` (`order_id`,`product_id`),
    KEY `product_id` (`product_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS cart (
    `cart_id` int(11) NOT NULL AUTO_INCREMENT,
    `quantity` int(11) NOT NULL,
    `total_price` float,
    `product_id` int(11) NOT NULL,
    PRIMARY KEY (`cart_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS payment_type (
    `payment_type_id` int(11) NOT NULL AUTO_INCREMENT,
    `payment_method` VARCHAR(50) NOT NULL,
    `payment_company` VARCHAR(50) NOT NULL,
    PRIMARY KEY(`payment_type_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS payment (
    `payment_id` int(11) NOT NULL AUTO_INCREMENT,
    `order_id` int(11) NOT NULL,
    `amount` float NOT NULL,
    `payment_category_id` int(11) NOT NULL,
    `payment_date` datetime NOT NULL,
    `payment_status` int(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`payment_id`),
    KEY `order_id` (`order_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1;

ALTER TABLE product
    ADD CONSTRAINT `tblproduct_ibfk_1` FOREIGN KEY (`product_category_id`) REFERENCES product_category (`product_category_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE orders
    ADD CONSTRAINT `tblorder_ibfk_2` FOREIGN KEY (`customer_id`) REFERENCES customer (`customer_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE order_details
    ADD CONSTRAINT `tblorderdetails_ibfk_2` FOREIGN KEY (`order_id`) REFERENCES orders (`order_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `tblorderdetails_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES product (`product_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE cart
    ADD CONSTRAINT `tblcart_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES product (`product_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE payment
    ADD CONSTRAINT `tblpayment_ibfk_2` FOREIGN KEY (`order_id`) REFERENCES orders (`order_id`) ON DELETE CASCADE ON UPDATE CASCADE;

