
GO
DROP PROC IF EXISTS columns_find_by_names
GO
ALTER PROCEDURE [dbo].[columns_find_by_names](@columns TEXT , @table VARCHAR(100))
	AS
    BEGIN
        SELECT 
            COLUMN_NAME ,
            DATA_TYPE ,
            IS_NULLABLE ,
            ISNULL(CHARACTER_MAXIMUM_LENGTH,'0') CHARACTER_MAXIMUM_LENGTH 
        FROM INFORMATION_SCHEMA.COLUMNS 
        JOIN dbo.ConvertDelimitedListIntoStrTable(@columns,',') arr
                 ON COLUMN_NAME = arr.item
        WHERE  TABLE_NAME = @table
    END


	


