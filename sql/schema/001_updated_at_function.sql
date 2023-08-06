-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION before_update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE apply_before_update_updated_at_to_table(
    t text
) AS
$$
BEGIN
    EXECUTE format('
     CREATE TRIGGER before_update_updated_at_%s BEFORE UPDATE ON %I
                FOR EACH ROW EXECUTE PROCEDURE before_update_updated_at();
     ', t, t);
END;
$$ language 'plpgsql';
-- +goose StatementEnd

-- +goose Down
DROP PROCEDURE IF EXISTS apply_before_update_updated_at_to_table;
DROP FUNCTION IF EXISTS before_update_updated_at;