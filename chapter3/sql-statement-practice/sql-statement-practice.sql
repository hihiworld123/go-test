-- 题目1：基本CRUD操作
INSERT INTO students (NAME, age, grade) VALUES('张三', 20, '三年级');

SELECT * FROM students t WHERE t.age >18;

UPDATE students SET grade = '四年级' WHERE `NAME`='张三';

DELETE FROM students WHERE age < 15;

-- 题目2：事务语句

drop procedure if exists transfer1;
delimiter $$
create PROCEDURE transfer1(in a_acount_id int,in b_acount_id int,in money decimal)
begin
    declare v_balance decimal;

START TRANSACTION;
IF v_balance >= money THEN
UPDATE accounts SET balance = balance - money WHERE id = a_acount_id;
UPDATE accounts SET balance = balance + money WHERE id = b_acount_id;
INSERT INTO transactions(from_account_id,to_account_id,amount) VALUES(a_acount_id,b_acount_id,money);
COMMIT;
ELSE
      ROLLBACK;
END IF;

end $$
DELIMITER ;

