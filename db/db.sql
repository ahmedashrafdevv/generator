
DROP PROC IF EXISTS StagingCustomerslist 
GO
CREATE PROC StagingCustomerslist
    (
    @customer_id INT ,
    @created_at DATETIME ,
    @name VARCHAR(250) ,
    @phone VARCHAR(250)  = NULL
)
AS
BEGIN

    SELECT
        customer_id ,
        created_at ,
        name ,
        ISNULL(phone , '')
    FROM staging.customers
    WHERE
 tomer_id =  @customer_id
        AND created_at LIKE CONCAT('%' , @created_at , '%')
        AND name LIKE CONCAT('%' , @name , '%')
        AND phone LIKE CONCAT('%' , ISNULL(@phone , phone)  , '%')
END
